/*
* Copyright 2018 - Present Okta, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

// Code generated by okta openapi generator. DO NOT EDIT.

package okta

import (
	"context"
	"fmt"
	"time"
)

type AuthorizationServerPolicyRuleResource resource

type AuthorizationServerPolicyRule struct {
	Actions     *AuthorizationServerPolicyRuleActions    `json:"actions,omitempty"`
	Conditions  *AuthorizationServerPolicyRuleConditions `json:"conditions,omitempty"`
	Created     *time.Time                               `json:"created,omitempty"`
	Id          string                                   `json:"id,omitempty"`
	LastUpdated *time.Time                               `json:"lastUpdated,omitempty"`
	Name        string                                   `json:"name,omitempty"`
	Priority    int64                                    `json:"priority,omitempty"`
	Status      string                                   `json:"status,omitempty"`
	System      *bool                                    `json:"system,omitempty"`
	Type        string                                   `json:"type,omitempty"`
}

// Updates the configuration of the Policy Rule defined in the specified Custom Authorization Server and Policy.
func (m *AuthorizationServerPolicyRuleResource) UpdateAuthorizationServerPolicyRule(ctx context.Context, authServerId string, policyId string, ruleId string, body AuthorizationServerPolicyRule) (*AuthorizationServerPolicyRule, *Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/policies/%v/rules/%v", authServerId, policyId, ruleId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("PUT", url, body)
	if err != nil {
		return nil, nil, err
	}

	var authorizationServerPolicyRule *AuthorizationServerPolicyRule

	resp, err := rq.Do(ctx, req, &authorizationServerPolicyRule)
	if err != nil {
		return nil, resp, err
	}

	return authorizationServerPolicyRule, resp, nil
}

// Deletes a Policy Rule defined in the specified Custom Authorization Server and Policy.
func (m *AuthorizationServerPolicyRuleResource) DeleteAuthorizationServerPolicyRule(ctx context.Context, authServerId string, policyId string, ruleId string) (*Response, error) {
	url := fmt.Sprintf("/api/v1/authorizationServers/%v/policies/%v/rules/%v", authServerId, policyId, ruleId)

	rq := m.client.CloneRequestExecutor()

	req, err := rq.WithAccept("application/json").WithContentType("application/json").NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := m.client.requestExecutor.Do(ctx, req, nil)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
