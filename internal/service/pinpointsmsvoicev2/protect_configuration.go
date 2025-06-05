// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pinpointsmsvoicev2

// **PLEASE DELETE THIS AND ALL TIP COMMENTS BEFORE SUBMITTING A PR FOR REVIEW!**
//
// TIP: ==== INTRODUCTION ====
// Thank you for trying the skaff tool!
//
// You have opted to include these helpful comments. They all include "TIP:"
// to help you find and remove them when you're done with them.
//
// While some aspects of this file are customized to your input, the
// scaffold tool does *not* look at the AWS API and ensure it has correct
// function, structure, and variable names. It makes guesses based on
// commonalities. You will need to make significant adjustments.
//
// In other words, as generated, this is a rough outline of the work you will
// need to do. If something doesn't make sense for your situation, get rid of
// it.

import (
	// TIP: ==== IMPORTS ====
	// This is a common set of imports but not customized to your code since
	// your code hasn't been written yet. Make sure you, your IDE, or
	// goimports -w <file> fixes these imports.
	//
	// The provider linter wants your imports to be in two groups: first,
	// standard library (i.e., "fmt" or "strings"), second, everything else.
	//
	// Also, AWS Go SDK v2 may handle nested structures differently than v1,
	// using the services/pinpointsmsvoicev2/types package. If so, you'll
	// need to import types and reference the nested types, e.g., as
	// awstypes.<Type Name>.
	"context"
	"errors"
	"time"

	"github.com/YakDriver/regexache"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2"
	awstypes "github.com/aws/aws-sdk-go-v2/service/pinpointsmsvoicev2/types"
	"github.com/hashicorp/terraform-plugin-framework-timeouts/resource/timeouts"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/retry"
	"github.com/hashicorp/terraform-provider-aws/internal/conns"
	"github.com/hashicorp/terraform-provider-aws/internal/create"
	"github.com/hashicorp/terraform-provider-aws/internal/errs"
	"github.com/hashicorp/terraform-provider-aws/internal/errs/fwdiag"
	"github.com/hashicorp/terraform-provider-aws/internal/framework"
	"github.com/hashicorp/terraform-provider-aws/internal/framework/flex"
	fwtypes "github.com/hashicorp/terraform-provider-aws/internal/framework/types"
	"github.com/hashicorp/terraform-provider-aws/internal/sweep"
	sweepfw "github.com/hashicorp/terraform-provider-aws/internal/sweep/framework"
	tftags "github.com/hashicorp/terraform-provider-aws/internal/tags"
	"github.com/hashicorp/terraform-provider-aws/internal/tfresource"
	"github.com/hashicorp/terraform-provider-aws/names"
)

// TIP: ==== FILE STRUCTURE ====
// All resources should follow this basic outline. Improve this resource's
// maintainability by sticking to it.
//
// 1. Package declaration
// 2. Imports
// 3. Main resource struct with schema method
// 4. Create, read, update, delete methods (in that order)
// 5. Other functions (flatteners, expanders, waiters, finders, etc.)

// Function annotations are used for resource registration to the Provider. DO NOT EDIT.
// @FrameworkResource("aws_pinpointsmsvoicev2_protect_configuration", name="Protect Configuration")
func newResourceProtectConfiguration(_ context.Context) (resource.ResourceWithConfigure, error) {
	r := &resourceProtectConfiguration{}

	// TIP: ==== CONFIGURABLE TIMEOUTS ====
	// Users can configure timeout lengths but you need to use the times they
	// provide. Access the timeout they configure (or the defaults) using,
	// e.g., r.CreateTimeout(ctx, plan.Timeouts) (see below). The times here are
	// the defaults if they don't configure timeouts.
	r.SetDefaultCreateTimeout(30 * time.Minute)
	r.SetDefaultUpdateTimeout(30 * time.Minute)
	r.SetDefaultDeleteTimeout(30 * time.Minute)

	return r, nil
}

const (
	ResNameProtectConfiguration = "Protect Configuration"
)

type resourceProtectConfiguration struct {
	framework.ResourceWithConfigure
	framework.WithTimeouts
}

// TIP: ==== SCHEMA ====
// In the schema, add each of the attributes in snake case (e.g.,
// delete_automated_backups).
//
// Formatting rules:
// * Alphabetize attributes to make them easier to find.
// * Do not add a blank line between attributes.
//
// Attribute basics:
//   - If a user can provide a value ("configure a value") for an
//     attribute (e.g., instances = 5), we call the attribute an
//     "argument."
//   - You change the way users interact with attributes using:
//   - Required
//   - Optional
//   - Computed
//   - There are only four valid combinations:
//
// 1. Required only - the user must provide a value
// Required: true,
//
//  2. Optional only - the user can configure or omit a value; do not
//     use Default or DefaultFunc
//
// Optional: true,
//
//  3. Computed only - the provider can provide a value but the user
//     cannot, i.e., read-only
//
// Computed: true,
//
//  4. Optional AND Computed - the provider or user can provide a value;
//     use this combination if you are using Default
//
// Optional: true,
// Computed: true,
//
// You will typically find arguments in the input struct
// (e.g., CreateDBInstanceInput) for the create operation. Sometimes
// they are only in the input struct (e.g., ModifyDBInstanceInput) for
// the modify operation.
//
// For more about schema options, visit
// https://developer.hashicorp.com/terraform/plugin/framework/handling-data/schemas?page=schemas
func (r *resourceProtectConfiguration) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"account_default": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
				Description: "Indicates whether this protect configuration is the account default. Only one protect configuration can be the account default.",
			},
			names.AttrARN: framework.ARNAttributeComputedOnly(),
			// TIP: ==== "ID" ATTRIBUTE ====
			// When using the Terraform Plugin Framework, there is no required "id" attribute.
			// This is different from the Terraform Plugin SDK.
			//
			// Only include an "id" attribute if the AWS API has an "Id" field, such as "ProtectConfigurationId"
			"client_token": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.RegexMatches(regexache.MustCompile("^[!-~]{1,63}$"), "must be between 1 and 64 characters and contain printable ASCII characters, excluding whitespace"),
				},
				Description: "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.",
			},
			"deletion_protection_enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Default:     booldefault.StaticBool(false),
				Description: "Indicates whether deletion protection is enabled for this protect configuration. If deletion protection is enabled, you cannot delete the protect configuration unless you first disable deletion protection.",
			},
			names.AttrID:      framework.IDAttribute(),
			names.AttrTags:    tftags.TagsAttribute(),
			names.AttrTagsAll: tftags.TagsAttributeComputedOnly(),
		},
		Blocks: map[string]schema.Block{
			// The request syntax:
			// {
			// 	"CountryRuleSetUpdates": {
			// 			"string" : {
			// 				"ProtectStatus": "string"
			// 			}
			// 	},
			// 	"NumberCapability": "string",
			// 	"ProtectConfigurationId": "string"
			// }
			"sms_country_rules": schema.ListNestedBlock{
				CustomType: fwtypes.NewListNestedObjectTypeOf[countryRuleSetModel](ctx),
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
					listvalidator.SizeAtMost(300),
				},
				Description: "A list of country rules that define the protect status for each country.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"country_code": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexache.MustCompile(`^[A-Z]{2}$`), "must be a valid ISO 3166-1 alpha-2 country code"),
							},
							Description: "The ISO 3166-1 alpha-2 country code for the country to which the rule applies.",
						},
						"protect_status": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"ALLOW",
									"BLOCK",
									"MONITOR",
									"FILTER",
								),
							},
							Description: "The protect status for the country rule. Valid values are ALLOW, BLOCK, MONITOR, and FILTER.",
						},
					},
				},
			},
			"voice_country_rules": schema.ListNestedBlock{
				CustomType: fwtypes.NewListNestedObjectTypeOf[countryRuleSetModel](ctx),
				Validators: []validator.List{
					listvalidator.SizeAtLeast(1),
					listvalidator.SizeAtMost(300),
				},
				Description: "A list of country rules that define the protect status for each country.",
				NestedObject: schema.NestedBlockObject{
					Attributes: map[string]schema.Attribute{
						"country_code": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.RegexMatches(regexache.MustCompile(`^[A-Z]{2}$`), "must be a valid ISO 3166-1 alpha-2 country code"),
							},
							Description: "The ISO 3166-1 alpha-2 country code for the country to which the rule applies.",
						},
						"protect_status": schema.StringAttribute{
							Required: true,
							Validators: []validator.String{
								stringvalidator.OneOf(
									"ALLOW",
									"BLOCK",
								),
							},
							Description: "The protect status for the country rule. Valid values are ALLOW and BLOCK.",
						},
					},
				},
			},
			names.AttrTimeouts: timeouts.Block(ctx, timeouts.Opts{
				Create: true,
				Update: true,
				Delete: true,
			}),
		},
	}
}

func (r *resourceProtectConfiguration) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// TIP: ==== RESOURCE CREATE ====
	// Generally, the Create function should do the following things. Make
	// sure there is a good reason if you don't do one of these.
	//
	// 1. Get a client connection to the relevant service
	// 2. Fetch the plan
	// 3. Populate a create input structure
	// 4. Call the AWS create/put function
	// 5. Using the output from the create function, set the minimum arguments
	//    and attributes for the Read function to work, as well as any computed
	//    only attributes.
	// 6. Use a waiter to wait for create to complete
	// 7. Save the request plan to response state

	// TIP: -- 1. Get a client connection to the relevant service
	conn := r.Meta().PinpointSMSVoiceV2Client(ctx)

	// TIP: -- 2. Fetch the plan
	var plan resourceProtectConfigurationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TIP: -- 3. Populate a Create input structure
	var input pinpointsmsvoicev2.CreateProtectConfigurationInput
	// TIP: Using a field name prefix allows mapping fields such as `ID` to `ProtectConfigurationId`
	resp.Diagnostics.Append(flex.Expand(ctx, plan, &input, flex.WithFieldNamePrefix("ProtectConfiguration"))...)
	if resp.Diagnostics.HasError() {
		return
	}

	input.Tags = getTagsIn(ctx)

	// TIP: -- 4. Call the AWS Create function
	out, err := conn.CreateProtectConfiguration(ctx, &input)
	if err != nil {
		// TIP: Since ID has not been set yet, you cannot use plan.ID.String()
		// in error messages at this point.
		resp.Diagnostics.AddError(
			create.ProblemStandardMessage(names.PinpointSMSVoiceV2, create.ErrActionCreating, ResNameProtectConfiguration, plan.Name.String(), err),
			err.Error(),
		)
		return
	}
	if out == nil || out.ProtectConfigurationId == nil {
		resp.Diagnostics.AddError(
			create.ProblemStandardMessage(names.PinpointSMSVoiceV2, create.ErrActionCreating, ResNameProtectConfiguration, plan.Name.String(), nil),
			errors.New("empty output").Error(),
		)
		return
	}

	// TIP: -- 5. Using the output from the create function, set attributes
	resp.Diagnostics.Append(flex.Flatten(ctx, out, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TIP: -- 6. Use a waiter to wait for create to complete
	createTimeout := r.CreateTimeout(ctx, plan.Timeouts)
	_, err = waitProtectConfigurationCreated(ctx, conn, plan.ID.ValueString(), createTimeout)
	if err != nil {
		resp.Diagnostics.AddError(
			create.ProblemStandardMessage(names.PinpointSMSVoiceV2, create.ErrActionWaitingForCreation, ResNameProtectConfiguration, plan.Name.String(), err),
			err.Error(),
		)
		return
	}

	// TIP: -- 7. Save the request plan to response state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *resourceProtectConfiguration) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// TIP: ==== RESOURCE READ ====
	// Generally, the Read function should do the following things. Make
	// sure there is a good reason if you don't do one of these.
	//
	// 1. Get a client connection to the relevant service
	// 2. Fetch the state
	// 3. Get the resource from AWS
	// 4. Remove resource from state if it is not found
	// 5. Set the arguments and attributes
	// 6. Set the state

	// TIP: -- 1. Get a client connection to the relevant service
	conn := r.Meta().PinpointSMSVoiceV2Client(ctx)

	// TIP: -- 2. Fetch the state
	var state resourceProtectConfigurationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TIP: -- 3. Get the resource from AWS using an API Get, List, or Describe-
	// type function, or, better yet, using a finder.
	out, err := findProtectConfigurationsByID(ctx, conn, state.ID.ValueString())
	// TIP: -- 4. Remove resource from state if it is not found
	if tfresource.NotFound(err) {
		resp.Diagnostics.Append(fwdiag.NewResourceNotFoundWarningDiagnostic(err))
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError(
			create.ProblemStandardMessage(names.PinpointSMSVoiceV2, create.ErrActionReading, ResNameProtectConfiguration, state.ID.String(), err),
			err.Error(),
		)
		return
	}

	// TIP: -- 5. Set the arguments and attributes
	resp.Diagnostics.Append(flex.Flatten(ctx, out, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TIP: -- 6. Set the state
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *resourceProtectConfiguration) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// TIP: ==== RESOURCE UPDATE ====
	// Not all resources have Update functions. There are a few reasons:
	// a. The AWS API does not support changing a resource
	// b. All arguments have RequiresReplace() plan modifiers
	// c. The AWS API uses a create call to modify an existing resource
	//
	// In the cases of a. and b., the resource will not have an update method
	// defined. In the case of c., Update and Create can be refactored to call
	// the same underlying function.
	//
	// The rest of the time, there should be an Update function and it should
	// do the following things. Make sure there is a good reason if you don't
	// do one of these.
	//
	// 1. Get a client connection to the relevant service
	// 2. Fetch the plan and state
	// 3. Populate a modify input structure and check for changes
	// 4. Call the AWS modify/update function
	// 5. Use a waiter to wait for update to complete
	// 6. Save the request plan to response state
	// TIP: -- 1. Get a client connection to the relevant service
	conn := r.Meta().PinpointSMSVoiceV2Client(ctx)

	// TIP: -- 2. Fetch the plan
	var plan, state resourceProtectConfigurationModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TIP: -- 3. Get the difference between the plan and state, if any
	diff, d := flex.Diff(ctx, plan, state)
	resp.Diagnostics.Append(d...)
	if resp.Diagnostics.HasError() {
		return
	}

	if diff.HasChanges() {
		var input pinpointsmsvoicev2.UpdateProtectConfigurationInput
		resp.Diagnostics.Append(flex.Expand(ctx, plan, &input, flex.WithFieldNamePrefix("Test"))...)
		if resp.Diagnostics.HasError() {
			return
		}

		// TIP: -- 4. Call the AWS modify/update function
		out, err := conn.UpdateProtectConfiguration(ctx, &input)
		if err != nil {
			resp.Diagnostics.AddError(
				create.ProblemStandardMessage(names.PinpointSMSVoiceV2, create.ErrActionUpdating, ResNameProtectConfiguration, plan.ID.String(), err),
				err.Error(),
			)
			return
		}
		if out == nil || out.ProtectConfigurationId == nil {
			resp.Diagnostics.AddError(
				create.ProblemStandardMessage(names.PinpointSMSVoiceV2, create.ErrActionUpdating, ResNameProtectConfiguration, plan.ID.String(), nil),
				errors.New("empty output").Error(),
			)
			return
		}

		// TIP: Using the output from the update function, re-set any computed attributes
		resp.Diagnostics.Append(flex.Flatten(ctx, out, &plan)...)
		if resp.Diagnostics.HasError() {
			return
		}
	}

	// TIP: -- 5. Use a waiter to wait for update to complete
	updateTimeout := r.UpdateTimeout(ctx, plan.Timeouts)
	_, err := waitProtectConfigurationUpdated(ctx, conn, plan.ID.ValueString(), updateTimeout)
	if err != nil {
		resp.Diagnostics.AddError(
			create.ProblemStandardMessage(names.PinpointSMSVoiceV2, create.ErrActionWaitingForUpdate, ResNameProtectConfiguration, plan.ID.String(), err),
			err.Error(),
		)
		return
	}

	// TIP: -- 6. Save the request plan to response state
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *resourceProtectConfiguration) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// TIP: ==== RESOURCE DELETE ====
	// Most resources have Delete functions. There are rare situations
	// where you might not need a delete:
	// a. The AWS API does not provide a way to delete the resource
	// b. The point of your resource is to perform an action (e.g., reboot a
	//    server) and deleting serves no purpose.
	//
	// The Delete function should do the following things. Make sure there
	// is a good reason if you don't do one of these.
	//
	// 1. Get a client connection to the relevant service
	// 2. Fetch the state
	// 3. Populate a delete input structure
	// 4. Call the AWS delete function
	// 5. Use a waiter to wait for delete to complete
	// TIP: -- 1. Get a client connection to the relevant service
	conn := r.Meta().PinpointSMSVoiceV2Client(ctx)

	// TIP: -- 2. Fetch the state
	var state resourceProtectConfigurationModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// TIP: -- 3. Populate a delete input structure
	input := pinpointsmsvoicev2.DeleteProtectConfigurationInput{
		ProtectConfigurationId: state.ID.ValueStringPointer(),
	}

	// TIP: -- 4. Call the AWS delete function
	_, err := conn.DeleteProtectConfiguration(ctx, &input)
	// TIP: On rare occassions, the API returns a not found error after deleting a
	// resource. If that happens, we don't want it to show up as an error.
	if err != nil {
		if errs.IsA[*awstypes.ResourceNotFoundException](err) {
			return
		}

		resp.Diagnostics.AddError(
			create.ProblemStandardMessage(names.PinpointSMSVoiceV2, create.ErrActionDeleting, ResNameProtectConfiguration, state.ID.String(), err),
			err.Error(),
		)
		return
	}

	// TIP: -- 5. Use a waiter to wait for delete to complete
	deleteTimeout := r.DeleteTimeout(ctx, state.Timeouts)
	_, err = waitProtectConfigurationDeleted(ctx, conn, state.ID.ValueString(), deleteTimeout)
	if err != nil {
		resp.Diagnostics.AddError(
			create.ProblemStandardMessage(names.PinpointSMSVoiceV2, create.ErrActionWaitingForDeletion, ResNameProtectConfiguration, state.ID.String(), err),
			err.Error(),
		)
		return
	}
}

// TIP: ==== TERRAFORM IMPORTING ====
// If Read can get all the information it needs from the Identifier
// (i.e., path.Root("id")), you can use the PassthroughID importer. Otherwise,
// you'll need a custom import function.
//
// See more:
// https://developer.hashicorp.com/terraform/plugin/framework/resources/import
func (r *resourceProtectConfiguration) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	resource.ImportStatePassthroughID(ctx, path.Root(names.AttrID), req, resp)
}

// TIP: ==== STATUS CONSTANTS ====
// Create constants for states and statuses if the service does not
// already have suitable constants. We prefer that you use the constants
// provided in the service if available (e.g., awstypes.StatusInProgress).
const (
	statusChangePending = "Pending"
	statusDeleting      = "Deleting"
	statusNormal        = "Normal"
	statusUpdated       = "Updated"
)

// TIP: ==== WAITERS ====
// Some resources of some services have waiters provided by the AWS API.
// Unless they do not work properly, use them rather than defining new ones
// here.
//
// Sometimes we define the wait, status, and find functions in separate
// files, wait.go, status.go, and find.go. Follow the pattern set out in the
// service and define these where it makes the most sense.
//
// If these functions are used in the _test.go file, they will need to be
// exported (i.e., capitalized).
//
// You will need to adjust the parameters and names to fit the service.
func waitProtectConfigurationCreated(ctx context.Context, conn *pinpointsmsvoicev2.Client, id string, timeout time.Duration) (*awstypes.ProtectConfigurationInformation, error) {
	stateConf := &retry.StateChangeConf{
		Pending:                   []string{},
		Target:                    []string{statusNormal},
		Refresh:                   statusProtectConfiguration(ctx, conn, id),
		Timeout:                   timeout,
		NotFoundChecks:            20,
		ContinuousTargetOccurence: 2,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)
	if out, ok := outputRaw.(*awstypes.ProtectConfigurationInformation); ok {
		return out, err
	}

	return nil, err
}

// TIP: It is easier to determine whether a resource is updated for some
// resources than others. The best case is a status flag that tells you when
// the update has been fully realized. Other times, you can check to see if a
// key resource argument is updated to a new value or not.
func waitProtectConfigurationUpdated(ctx context.Context, conn *pinpointsmsvoicev2.Client, id string, timeout time.Duration) (*awstypes.ProtectConfigurationInformation, error) {
	stateConf := &retry.StateChangeConf{
		Pending:                   []string{statusChangePending},
		Target:                    []string{statusUpdated},
		Refresh:                   statusProtectConfiguration(ctx, conn, id),
		Timeout:                   timeout,
		NotFoundChecks:            20,
		ContinuousTargetOccurence: 2,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)
	if out, ok := outputRaw.(*awstypes.ProtectConfigurationInformation); ok {
		return out, err
	}

	return nil, err
}

// TIP: A deleted waiter is almost like a backwards created waiter. There may
// be additional pending states, however.
func waitProtectConfigurationDeleted(ctx context.Context, conn *pinpointsmsvoicev2.Client, id string, timeout time.Duration) (*awstypes.ProtectConfigurationInformation, error) {
	stateConf := &retry.StateChangeConf{
		Pending: []string{statusDeleting, statusNormal},
		Target:  []string{},
		Refresh: statusProtectConfiguration(ctx, conn, id),
		Timeout: timeout,
	}

	outputRaw, err := stateConf.WaitForStateContext(ctx)
	if out, ok := outputRaw.(*awstypes.ProtectConfigurationInformation); ok {
		return out, err
	}

	return nil, err
}

// TIP: ==== STATUS ====
// The status function can return an actual status when that field is
// available from the API (e.g., out.Status). Otherwise, you can use custom
// statuses to communicate the states of the resource.
//
// Waiters consume the values returned by status functions. Design status so
// that it can be reused by a create, update, and delete waiter, if possible.
func statusProtectConfiguration(ctx context.Context, conn *pinpointsmsvoicev2.Client, id string) retry.StateRefreshFunc {
	return func() (any, string, error) {
		out, err := findProtectConfigurationsByID(ctx, conn, id)
		if tfresource.NotFound(err) {
			return nil, "", nil
		}

		if err != nil {
			return nil, "", err
		}

		return out, out.CreatedTimestamp.String(), nil
	}
}

// TIP: ==== FINDERS ====
// The find function is not strictly necessary. You could do the API
// request from the status function. However, we have found that find often
// comes in handy in other places besides the status function. As a result, it
// is good practice to define it separately.
func findProtectConfigurationsByID(ctx context.Context, conn *pinpointsmsvoicev2.Client, id string) (*awstypes.ProtectConfigurationInformation, error) {
	input := pinpointsmsvoicev2.DescribeProtectConfigurationsInput{
		ProtectConfigurationIds: []string{id},
	}

	return findProtectConfiguration(ctx, conn, &input)
}

func findProtectConfiguration(ctx context.Context, conn *pinpointsmsvoicev2.Client, input *pinpointsmsvoicev2.DescribeProtectConfigurationsInput) (*awstypes.ProtectConfigurationInformation, error) {
	output, err := findProtectConfigurations(ctx, conn, input)

	if err != nil {
		return nil, err
	}

	return tfresource.AssertSingleValueResult(output)
}

func findProtectConfigurations(ctx context.Context, conn *pinpointsmsvoicev2.Client, input *pinpointsmsvoicev2.DescribeProtectConfigurationsInput) ([]awstypes.ProtectConfigurationInformation, error) {
	var output []awstypes.ProtectConfigurationInformation

	pages := pinpointsmsvoicev2.NewDescribeProtectConfigurationsPaginator(conn, input)
	for pages.HasMorePages() {
		page, err := pages.NextPage(ctx)

		if errs.IsA[*awstypes.ResourceNotFoundException](err) {
			return nil, &retry.NotFoundError{
				LastError:   err,
				LastRequest: input,
			}
		}

		if err != nil {
			return nil, err
		}

		output = append(output, page.ProtectConfigurations...)
	}

	return output, nil
}

// TIP: ==== DATA STRUCTURES ====
// With Terraform Plugin-Framework configurations are deserialized into
// Go types, providing type safety without the need for type assertions.
// These structs should match the schema definition exactly, and the `tfsdk`
// tag value should match the attribute name.
//
// Nested objects are represented in their own data struct. These will
// also have a corresponding attribute type mapping for use inside flex
// functions.
//
// See more:
// https://developer.hashicorp.com/terraform/plugin/framework/handling-data/accessing-values
type resourceProtectConfigurationModel struct {
	AccountDefault      types.Bool                                           `tfsdk:"account_default"`
	ARN                 types.String                                         `tfsdk:"arn"`
	ID                  types.String                                         `tfsdk:"id"`
	SMSCountryRuleSet   fwtypes.ListNestedObjectValueOf[countryRuleSetModel] `tfsdk:"sms_country_rule_set"`
	Tags                tftags.Map                                           `tfsdk:"tags"`
	TagsAll             tftags.Map                                           `tfsdk:"tags_all"`
	Timeouts            timeouts.Value                                       `tfsdk:"timeouts"`
	VoiceCountryRuleSet fwtypes.ListNestedObjectValueOf[countryRuleSetModel] `tfsdk:"voice_country_rule_set"`
}

type countryRuleSetModel struct {
	CountryCode   types.String `tfsdk:"country_code"`
	ProtectStatus types.String `tfsdk:"protect_status"`
}

// TIP: ==== SWEEPERS ====
// When acceptance testing resources, interrupted or failed tests may
// leave behind orphaned resources in an account. To facilitate cleaning
// up lingering resources, each resource implementation should include
// a corresponding "sweeper" function.
//
// The sweeper function lists all resources of a given type and sets the
// appropriate identifers required to delete the resource via the Delete
// method implemented above.
//
// Once the sweeper function is implemented, register it in sweeper.go
// as follows:
//
//	awsv2.Register("aws_pinpointsmsvoicev2_protect_configuration", sweepProtectConfigurations)
//
// See more:
// https://hashicorp.github.io/terraform-provider-aws/running-and-writing-acceptance-tests/#acceptance-test-sweepers
func sweepProtectConfigurations(ctx context.Context, client *conns.AWSClient) ([]sweep.Sweepable, error) {
	input := pinpointsmsvoicev2.DescribeProtectConfigurationsInput{}
	conn := client.PinpointSMSVoiceV2Client(ctx)
	var sweepResources []sweep.Sweepable

	pages := pinpointsmsvoicev2.NewDescribeProtectConfigurationsPaginator(conn, &input)
	for pages.HasMorePages() {
		page, err := pages.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, v := range page.ProtectConfigurations {
			sweepResources = append(sweepResources, sweepfw.NewSweepResource(newResourceProtectConfiguration, client,
				sweepfw.NewAttribute(names.AttrID, aws.ToString(v.ProtectConfigurationId))),
			)
		}
	}

	return sweepResources, nil
}
