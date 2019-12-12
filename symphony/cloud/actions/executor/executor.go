// Copyright (c) 2004-present Facebook All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

import (
	"context"

	"github.com/facebookincubator/symphony/cloud/actions/core"
	"github.com/pkg/errors"
)

// Executor will execute all Actions defined in Rules for a Trigger
type Executor struct {
	Context context.Context
	*Registry
	DataLoader
	OnError func(error)
}

// Execute runs all workflows for the specified object/trigger
func (exc Executor) Execute(ctx context.Context, objectID string, triggerToPayload map[core.TriggerID]map[string]interface{}) {

	// Note that we should keep this interface serializable, so if we need to eventually
	// offload this to workers, we can

	for triggerID, inputPayload := range triggerToPayload {
		trigger, err := exc.Registry.TriggerForID(triggerID)
		if err != nil {
			// TODO: Should we bail here, or just log an error and continue
			exc.OnError(errors.Errorf("could not find trigger: %s", triggerID))
			continue
		}

		rules, err := exc.DataLoader.QueryRules(ctx, triggerID)
		if err != nil {
			exc.OnError(errors.Errorf("could not query rules for trigger: %s", triggerID))
		}

		for _, rule := range rules {
			shouldExecute, err := trigger.Evaluate(rule)
			if err != nil {
				exc.OnError(errors.Errorf("evaluating rule %s: %v", rule.ID, err))
				continue
			}
			if !shouldExecute {
				continue
			}
			for _, ruleAction := range rule.RuleActions {
				err := exc.executeAction(rule, ruleAction, inputPayload)
				if err != nil {
					exc.OnError(errors.Errorf("executing action %s: %v", ruleAction.ActionID, err))
				}
			}
		}
	}
}

func (exc Executor) executeAction(rule core.Rule, ruleAction *core.ActionsRuleAction, inputPayload map[string]interface{}) error {
	action, err := exc.Registry.ActionForID(ruleAction.ActionID)
	if err != nil {
		return errors.Errorf("could not find action %v, skipping: %v", ruleAction.ActionID, err)
	}
	actionContext := core.ActionContext{
		TriggerPayload: inputPayload,
		Rule:           rule,
		RuleAction:     ruleAction,
	}
	err = action.Execute(actionContext)
	if err != nil {
		return errors.Errorf("executing %v: %v", ruleAction.ActionID, err)
	}
	return nil
}
