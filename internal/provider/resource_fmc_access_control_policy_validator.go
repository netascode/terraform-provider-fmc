// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
)

func (r *AccessControlPolicyResource) ConfigValidators(context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		&AccessControlPolicyResourceValidator{},
	}
}

type AccessControlPolicyResourceValidator struct{}

// Description describes the validation in plain text formatting.
//
// This information may be automatically added to resource plain text
// descriptions by external tooling.
func (_ AccessControlPolicyResourceValidator) Description(_ context.Context) string {
	return "Validates the rules"
}

// MarkdownDescription describes the validation in Markdown formatting.
//
// This information may be automatically added to resource Markdown
// descriptions by external tooling.
func (x *AccessControlPolicyResourceValidator) MarkdownDescription(ctx context.Context) string {
	return x.Description(ctx)
}

// ValidateResource performs the validation.
//
// This method name is separate from the datasource.ConfigValidator
// interface ValidateDataSource method name and provider.ConfigValidator
// interface ValidateProvider method name to allow generic validators.
func (x *AccessControlPolicyResourceValidator) ValidateResource(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	_, diags := NewValidAccessControlPolicy(ctx, tfsdk.Plan(req.Config))
	resp.Diagnostics.Append(diags...)
}
