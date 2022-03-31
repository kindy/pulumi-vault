// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package gcp

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Reads a GCP auth role from a Vault server.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/gcp"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		role, err := gcp.LookupAuthBackendRole(ctx, &gcp.LookupAuthBackendRoleArgs{
// 			Backend:  pulumi.StringRef("my-gcp-backend"),
// 			RoleName: "my-role",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		ctx.Export("role-id", role.RoleId)
// 		return nil
// 	})
// }
// ```
func LookupAuthBackendRole(ctx *pulumi.Context, args *LookupAuthBackendRoleArgs, opts ...pulumi.InvokeOption) (*LookupAuthBackendRoleResult, error) {
	var rv LookupAuthBackendRoleResult
	err := ctx.Invoke("vault:gcp/getAuthBackendRole:getAuthBackendRole", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAuthBackendRole.
type LookupAuthBackendRoleArgs struct {
	// The unique name for the GCP backend from which to fetch the role. Defaults to "gcp".
	Backend *string `pulumi:"backend"`
	// The name of the role to retrieve the Role ID for.
	RoleName string `pulumi:"roleName"`
	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs []string `pulumi:"tokenBoundCidrs"`
	// If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl *int `pulumi:"tokenExplicitMaxTtl"`
	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl *int `pulumi:"tokenMaxTtl"`
	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy *bool `pulumi:"tokenNoDefaultPolicy"`
	// The
	// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
	// if any, in number of seconds to set on the token.
	TokenNumUses *int `pulumi:"tokenNumUses"`
	// (Optional) If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod *int `pulumi:"tokenPeriod"`
	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies []string `pulumi:"tokenPolicies"`
	// The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl *int `pulumi:"tokenTtl"`
	// The type of token that should be generated. Can be `service`,
	// `batch`, or `default` to use the mount's tuned default (which unless changed will be
	// `service` tokens). For token store roles, there are two additional possibilities:
	// `default-service` and `default-batch` which specify the type to return unless the client
	// requests a different type at generation time.
	TokenType *string `pulumi:"tokenType"`
}

// A collection of values returned by getAuthBackendRole.
type LookupAuthBackendRoleResult struct {
	Backend *string `pulumi:"backend"`
	// GCP regions bound to the role. Returned when `type` is `gce`.
	BoundInstanceGroups []string `pulumi:"boundInstanceGroups"`
	// GCP labels bound to the role. Returned when `type` is `gce`.
	BoundLabels []string `pulumi:"boundLabels"`
	// GCP projects bound to the role.
	BoundProjects []string `pulumi:"boundProjects"`
	// GCP regions bound to the role. Returned when `type` is `gce`.
	BoundRegions []string `pulumi:"boundRegions"`
	// GCP service accounts bound to the role. Returned when `type` is `iam`.
	BoundServiceAccounts []string `pulumi:"boundServiceAccounts"`
	// GCP zones bound to the role. Returned when `type` is `gce`.
	BoundZones []string `pulumi:"boundZones"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The RoleID of the GCP role.
	RoleId   string `pulumi:"roleId"`
	RoleName string `pulumi:"roleName"`
	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs []string `pulumi:"tokenBoundCidrs"`
	// If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl *int `pulumi:"tokenExplicitMaxTtl"`
	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl *int `pulumi:"tokenMaxTtl"`
	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy *bool `pulumi:"tokenNoDefaultPolicy"`
	// The
	// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
	// if any, in number of seconds to set on the token.
	TokenNumUses *int `pulumi:"tokenNumUses"`
	// (Optional) If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod *int `pulumi:"tokenPeriod"`
	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies []string `pulumi:"tokenPolicies"`
	// The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl *int `pulumi:"tokenTtl"`
	// The type of token that should be generated. Can be `service`,
	// `batch`, or `default` to use the mount's tuned default (which unless changed will be
	// `service` tokens). For token store roles, there are two additional possibilities:
	// `default-service` and `default-batch` which specify the type to return unless the client
	// requests a different type at generation time.
	TokenType *string `pulumi:"tokenType"`
	// Type of GCP role. Expected values are `iam` or `gce`.
	Type string `pulumi:"type"`
}

func LookupAuthBackendRoleOutput(ctx *pulumi.Context, args LookupAuthBackendRoleOutputArgs, opts ...pulumi.InvokeOption) LookupAuthBackendRoleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAuthBackendRoleResult, error) {
			args := v.(LookupAuthBackendRoleArgs)
			r, err := LookupAuthBackendRole(ctx, &args, opts...)
			return *r, err
		}).(LookupAuthBackendRoleResultOutput)
}

// A collection of arguments for invoking getAuthBackendRole.
type LookupAuthBackendRoleOutputArgs struct {
	// The unique name for the GCP backend from which to fetch the role. Defaults to "gcp".
	Backend pulumi.StringPtrInput `pulumi:"backend"`
	// The name of the role to retrieve the Role ID for.
	RoleName pulumi.StringInput `pulumi:"roleName"`
	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs pulumi.StringArrayInput `pulumi:"tokenBoundCidrs"`
	// If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl pulumi.IntPtrInput `pulumi:"tokenExplicitMaxTtl"`
	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl pulumi.IntPtrInput `pulumi:"tokenMaxTtl"`
	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy pulumi.BoolPtrInput `pulumi:"tokenNoDefaultPolicy"`
	// The
	// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
	// if any, in number of seconds to set on the token.
	TokenNumUses pulumi.IntPtrInput `pulumi:"tokenNumUses"`
	// (Optional) If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod pulumi.IntPtrInput `pulumi:"tokenPeriod"`
	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies pulumi.StringArrayInput `pulumi:"tokenPolicies"`
	// The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl pulumi.IntPtrInput `pulumi:"tokenTtl"`
	// The type of token that should be generated. Can be `service`,
	// `batch`, or `default` to use the mount's tuned default (which unless changed will be
	// `service` tokens). For token store roles, there are two additional possibilities:
	// `default-service` and `default-batch` which specify the type to return unless the client
	// requests a different type at generation time.
	TokenType pulumi.StringPtrInput `pulumi:"tokenType"`
}

func (LookupAuthBackendRoleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthBackendRoleArgs)(nil)).Elem()
}

// A collection of values returned by getAuthBackendRole.
type LookupAuthBackendRoleResultOutput struct{ *pulumi.OutputState }

func (LookupAuthBackendRoleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthBackendRoleResult)(nil)).Elem()
}

func (o LookupAuthBackendRoleResultOutput) ToLookupAuthBackendRoleResultOutput() LookupAuthBackendRoleResultOutput {
	return o
}

func (o LookupAuthBackendRoleResultOutput) ToLookupAuthBackendRoleResultOutputWithContext(ctx context.Context) LookupAuthBackendRoleResultOutput {
	return o
}

func (o LookupAuthBackendRoleResultOutput) Backend() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthBackendRoleResult) *string { return v.Backend }).(pulumi.StringPtrOutput)
}

// GCP regions bound to the role. Returned when `type` is `gce`.
func (o LookupAuthBackendRoleResultOutput) BoundInstanceGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthBackendRoleResult) []string { return v.BoundInstanceGroups }).(pulumi.StringArrayOutput)
}

// GCP labels bound to the role. Returned when `type` is `gce`.
func (o LookupAuthBackendRoleResultOutput) BoundLabels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthBackendRoleResult) []string { return v.BoundLabels }).(pulumi.StringArrayOutput)
}

// GCP projects bound to the role.
func (o LookupAuthBackendRoleResultOutput) BoundProjects() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthBackendRoleResult) []string { return v.BoundProjects }).(pulumi.StringArrayOutput)
}

// GCP regions bound to the role. Returned when `type` is `gce`.
func (o LookupAuthBackendRoleResultOutput) BoundRegions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthBackendRoleResult) []string { return v.BoundRegions }).(pulumi.StringArrayOutput)
}

// GCP service accounts bound to the role. Returned when `type` is `iam`.
func (o LookupAuthBackendRoleResultOutput) BoundServiceAccounts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthBackendRoleResult) []string { return v.BoundServiceAccounts }).(pulumi.StringArrayOutput)
}

// GCP zones bound to the role. Returned when `type` is `gce`.
func (o LookupAuthBackendRoleResultOutput) BoundZones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthBackendRoleResult) []string { return v.BoundZones }).(pulumi.StringArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAuthBackendRoleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthBackendRoleResult) string { return v.Id }).(pulumi.StringOutput)
}

// The RoleID of the GCP role.
func (o LookupAuthBackendRoleResultOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthBackendRoleResult) string { return v.RoleId }).(pulumi.StringOutput)
}

func (o LookupAuthBackendRoleResultOutput) RoleName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthBackendRoleResult) string { return v.RoleName }).(pulumi.StringOutput)
}

// List of CIDR blocks; if set, specifies blocks of IP
// addresses which can authenticate successfully, and ties the resulting token to these blocks
// as well.
func (o LookupAuthBackendRoleResultOutput) TokenBoundCidrs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthBackendRoleResult) []string { return v.TokenBoundCidrs }).(pulumi.StringArrayOutput)
}

// If set, will encode an
// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
// `tokenMaxTtl` would otherwise allow a renewal.
func (o LookupAuthBackendRoleResultOutput) TokenExplicitMaxTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAuthBackendRoleResult) *int { return v.TokenExplicitMaxTtl }).(pulumi.IntPtrOutput)
}

// The maximum lifetime for generated tokens in number of seconds.
// Its current value will be referenced at renewal time.
func (o LookupAuthBackendRoleResultOutput) TokenMaxTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAuthBackendRoleResult) *int { return v.TokenMaxTtl }).(pulumi.IntPtrOutput)
}

// If set, the default policy will not be set on
// generated tokens; otherwise it will be added to the policies set in token_policies.
func (o LookupAuthBackendRoleResultOutput) TokenNoDefaultPolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupAuthBackendRoleResult) *bool { return v.TokenNoDefaultPolicy }).(pulumi.BoolPtrOutput)
}

// The
// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
// if any, in number of seconds to set on the token.
func (o LookupAuthBackendRoleResultOutput) TokenNumUses() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAuthBackendRoleResult) *int { return v.TokenNumUses }).(pulumi.IntPtrOutput)
}

// (Optional) If set, indicates that the
// token generated using this role should never expire. The token should be renewed within the
// duration specified by this value. At each renewal, the token's TTL will be set to the
// value of this field. Specified in seconds.
func (o LookupAuthBackendRoleResultOutput) TokenPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAuthBackendRoleResult) *int { return v.TokenPeriod }).(pulumi.IntPtrOutput)
}

// List of policies to encode onto generated tokens. Depending
// on the auth method, this list may be supplemented by user/group/other values.
func (o LookupAuthBackendRoleResultOutput) TokenPolicies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthBackendRoleResult) []string { return v.TokenPolicies }).(pulumi.StringArrayOutput)
}

// The incremental lifetime for generated tokens in number of seconds.
// Its current value will be referenced at renewal time.
func (o LookupAuthBackendRoleResultOutput) TokenTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupAuthBackendRoleResult) *int { return v.TokenTtl }).(pulumi.IntPtrOutput)
}

// The type of token that should be generated. Can be `service`,
// `batch`, or `default` to use the mount's tuned default (which unless changed will be
// `service` tokens). For token store roles, there are two additional possibilities:
// `default-service` and `default-batch` which specify the type to return unless the client
// requests a different type at generation time.
func (o LookupAuthBackendRoleResultOutput) TokenType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthBackendRoleResult) *string { return v.TokenType }).(pulumi.StringPtrOutput)
}

// Type of GCP role. Expected values are `iam` or `gce`.
func (o LookupAuthBackendRoleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthBackendRoleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAuthBackendRoleResultOutput{})
}
