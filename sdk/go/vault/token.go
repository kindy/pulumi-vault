// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vault.NewToken(ctx, "example", &vault.TokenArgs{
//				Metadata: pulumi.StringMap{
//					"purpose": pulumi.String("service-account"),
//				},
//				Policies: pulumi.StringArray{
//					pulumi.String("policy1"),
//					pulumi.String("policy2"),
//				},
//				RenewIncrement: pulumi.Int(86400),
//				RenewMinLease:  pulumi.Int(43200),
//				Renewable:      pulumi.Bool(true),
//				RoleName:       pulumi.String("app"),
//				Ttl:            pulumi.String("24h"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Tokens can be imported using its `id` as accessor id, e.g.
//
// ```sh
//
//	$ pulumi import vault:index/token:Token example <accessor_id>
//
// ```
type Token struct {
	pulumi.CustomResourceState

	// String containing the client token if stored in present file
	ClientToken pulumi.StringOutput `pulumi:"clientToken"`
	// String containing the token display name
	DisplayName pulumi.StringPtrOutput `pulumi:"displayName"`
	// The explicit max TTL of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
	ExplicitMaxTtl pulumi.StringPtrOutput `pulumi:"explicitMaxTtl"`
	// String containing the token lease duration if present in state file
	LeaseDuration pulumi.IntOutput `pulumi:"leaseDuration"`
	// String containing the token lease started time if present in state file
	LeaseStarted pulumi.StringOutput `pulumi:"leaseStarted"`
	// Metadata to be set on this token
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Flag to not attach the default policy to this token
	NoDefaultPolicy pulumi.BoolPtrOutput `pulumi:"noDefaultPolicy"`
	// Flag to create a token without parent
	NoParent pulumi.BoolOutput `pulumi:"noParent"`
	// The number of allowed uses of this token
	NumUses pulumi.IntOutput `pulumi:"numUses"`
	// The period of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
	Period pulumi.StringPtrOutput `pulumi:"period"`
	// List of policies to attach to this token
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
	// The renew increment. This is specified in seconds
	RenewIncrement pulumi.IntPtrOutput `pulumi:"renewIncrement"`
	// The minimal lease to renew this token
	RenewMinLease pulumi.IntPtrOutput `pulumi:"renewMinLease"`
	// Flag to allow to renew this token
	Renewable pulumi.BoolOutput `pulumi:"renewable"`
	// The token role name
	RoleName pulumi.StringPtrOutput `pulumi:"roleName"`
	// The TTL period of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
	Ttl pulumi.StringPtrOutput `pulumi:"ttl"`
	// The client wrapped token.
	WrappedToken pulumi.StringOutput `pulumi:"wrappedToken"`
	// The client wrapping accessor.
	WrappingAccessor pulumi.StringOutput `pulumi:"wrappingAccessor"`
	// The TTL period of the wrapped token.
	WrappingTtl pulumi.StringPtrOutput `pulumi:"wrappingTtl"`
}

// NewToken registers a new resource with the given unique name, arguments, and options.
func NewToken(ctx *pulumi.Context,
	name string, args *TokenArgs, opts ...pulumi.ResourceOption) (*Token, error) {
	if args == nil {
		args = &TokenArgs{}
	}

	secrets := pulumi.AdditionalSecretOutputs([]string{
		"clientToken",
		"wrappedToken",
		"wrappingAccessor",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Token
	err := ctx.RegisterResource("vault:index/token:Token", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetToken gets an existing Token resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetToken(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TokenState, opts ...pulumi.ResourceOption) (*Token, error) {
	var resource Token
	err := ctx.ReadResource("vault:index/token:Token", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Token resources.
type tokenState struct {
	// String containing the client token if stored in present file
	ClientToken *string `pulumi:"clientToken"`
	// String containing the token display name
	DisplayName *string `pulumi:"displayName"`
	// The explicit max TTL of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
	ExplicitMaxTtl *string `pulumi:"explicitMaxTtl"`
	// String containing the token lease duration if present in state file
	LeaseDuration *int `pulumi:"leaseDuration"`
	// String containing the token lease started time if present in state file
	LeaseStarted *string `pulumi:"leaseStarted"`
	// Metadata to be set on this token
	Metadata map[string]string `pulumi:"metadata"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Flag to not attach the default policy to this token
	NoDefaultPolicy *bool `pulumi:"noDefaultPolicy"`
	// Flag to create a token without parent
	NoParent *bool `pulumi:"noParent"`
	// The number of allowed uses of this token
	NumUses *int `pulumi:"numUses"`
	// The period of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
	Period *string `pulumi:"period"`
	// List of policies to attach to this token
	Policies []string `pulumi:"policies"`
	// The renew increment. This is specified in seconds
	RenewIncrement *int `pulumi:"renewIncrement"`
	// The minimal lease to renew this token
	RenewMinLease *int `pulumi:"renewMinLease"`
	// Flag to allow to renew this token
	Renewable *bool `pulumi:"renewable"`
	// The token role name
	RoleName *string `pulumi:"roleName"`
	// The TTL period of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
	Ttl *string `pulumi:"ttl"`
	// The client wrapped token.
	WrappedToken *string `pulumi:"wrappedToken"`
	// The client wrapping accessor.
	WrappingAccessor *string `pulumi:"wrappingAccessor"`
	// The TTL period of the wrapped token.
	WrappingTtl *string `pulumi:"wrappingTtl"`
}

type TokenState struct {
	// String containing the client token if stored in present file
	ClientToken pulumi.StringPtrInput
	// String containing the token display name
	DisplayName pulumi.StringPtrInput
	// The explicit max TTL of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
	ExplicitMaxTtl pulumi.StringPtrInput
	// String containing the token lease duration if present in state file
	LeaseDuration pulumi.IntPtrInput
	// String containing the token lease started time if present in state file
	LeaseStarted pulumi.StringPtrInput
	// Metadata to be set on this token
	Metadata pulumi.StringMapInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Flag to not attach the default policy to this token
	NoDefaultPolicy pulumi.BoolPtrInput
	// Flag to create a token without parent
	NoParent pulumi.BoolPtrInput
	// The number of allowed uses of this token
	NumUses pulumi.IntPtrInput
	// The period of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
	Period pulumi.StringPtrInput
	// List of policies to attach to this token
	Policies pulumi.StringArrayInput
	// The renew increment. This is specified in seconds
	RenewIncrement pulumi.IntPtrInput
	// The minimal lease to renew this token
	RenewMinLease pulumi.IntPtrInput
	// Flag to allow to renew this token
	Renewable pulumi.BoolPtrInput
	// The token role name
	RoleName pulumi.StringPtrInput
	// The TTL period of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
	Ttl pulumi.StringPtrInput
	// The client wrapped token.
	WrappedToken pulumi.StringPtrInput
	// The client wrapping accessor.
	WrappingAccessor pulumi.StringPtrInput
	// The TTL period of the wrapped token.
	WrappingTtl pulumi.StringPtrInput
}

func (TokenState) ElementType() reflect.Type {
	return reflect.TypeOf((*tokenState)(nil)).Elem()
}

type tokenArgs struct {
	// String containing the token display name
	DisplayName *string `pulumi:"displayName"`
	// The explicit max TTL of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
	ExplicitMaxTtl *string `pulumi:"explicitMaxTtl"`
	// Metadata to be set on this token
	Metadata map[string]string `pulumi:"metadata"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Flag to not attach the default policy to this token
	NoDefaultPolicy *bool `pulumi:"noDefaultPolicy"`
	// Flag to create a token without parent
	NoParent *bool `pulumi:"noParent"`
	// The number of allowed uses of this token
	NumUses *int `pulumi:"numUses"`
	// The period of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
	Period *string `pulumi:"period"`
	// List of policies to attach to this token
	Policies []string `pulumi:"policies"`
	// The renew increment. This is specified in seconds
	RenewIncrement *int `pulumi:"renewIncrement"`
	// The minimal lease to renew this token
	RenewMinLease *int `pulumi:"renewMinLease"`
	// Flag to allow to renew this token
	Renewable *bool `pulumi:"renewable"`
	// The token role name
	RoleName *string `pulumi:"roleName"`
	// The TTL period of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
	Ttl *string `pulumi:"ttl"`
	// The TTL period of the wrapped token.
	WrappingTtl *string `pulumi:"wrappingTtl"`
}

// The set of arguments for constructing a Token resource.
type TokenArgs struct {
	// String containing the token display name
	DisplayName pulumi.StringPtrInput
	// The explicit max TTL of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
	ExplicitMaxTtl pulumi.StringPtrInput
	// Metadata to be set on this token
	Metadata pulumi.StringMapInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Flag to not attach the default policy to this token
	NoDefaultPolicy pulumi.BoolPtrInput
	// Flag to create a token without parent
	NoParent pulumi.BoolPtrInput
	// The number of allowed uses of this token
	NumUses pulumi.IntPtrInput
	// The period of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
	Period pulumi.StringPtrInput
	// List of policies to attach to this token
	Policies pulumi.StringArrayInput
	// The renew increment. This is specified in seconds
	RenewIncrement pulumi.IntPtrInput
	// The minimal lease to renew this token
	RenewMinLease pulumi.IntPtrInput
	// Flag to allow to renew this token
	Renewable pulumi.BoolPtrInput
	// The token role name
	RoleName pulumi.StringPtrInput
	// The TTL period of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
	Ttl pulumi.StringPtrInput
	// The TTL period of the wrapped token.
	WrappingTtl pulumi.StringPtrInput
}

func (TokenArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*tokenArgs)(nil)).Elem()
}

type TokenInput interface {
	pulumi.Input

	ToTokenOutput() TokenOutput
	ToTokenOutputWithContext(ctx context.Context) TokenOutput
}

func (*Token) ElementType() reflect.Type {
	return reflect.TypeOf((**Token)(nil)).Elem()
}

func (i *Token) ToTokenOutput() TokenOutput {
	return i.ToTokenOutputWithContext(context.Background())
}

func (i *Token) ToTokenOutputWithContext(ctx context.Context) TokenOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenOutput)
}

func (i *Token) ToOutput(ctx context.Context) pulumix.Output[*Token] {
	return pulumix.Output[*Token]{
		OutputState: i.ToTokenOutputWithContext(ctx).OutputState,
	}
}

// TokenArrayInput is an input type that accepts TokenArray and TokenArrayOutput values.
// You can construct a concrete instance of `TokenArrayInput` via:
//
//	TokenArray{ TokenArgs{...} }
type TokenArrayInput interface {
	pulumi.Input

	ToTokenArrayOutput() TokenArrayOutput
	ToTokenArrayOutputWithContext(context.Context) TokenArrayOutput
}

type TokenArray []TokenInput

func (TokenArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Token)(nil)).Elem()
}

func (i TokenArray) ToTokenArrayOutput() TokenArrayOutput {
	return i.ToTokenArrayOutputWithContext(context.Background())
}

func (i TokenArray) ToTokenArrayOutputWithContext(ctx context.Context) TokenArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenArrayOutput)
}

func (i TokenArray) ToOutput(ctx context.Context) pulumix.Output[[]*Token] {
	return pulumix.Output[[]*Token]{
		OutputState: i.ToTokenArrayOutputWithContext(ctx).OutputState,
	}
}

// TokenMapInput is an input type that accepts TokenMap and TokenMapOutput values.
// You can construct a concrete instance of `TokenMapInput` via:
//
//	TokenMap{ "key": TokenArgs{...} }
type TokenMapInput interface {
	pulumi.Input

	ToTokenMapOutput() TokenMapOutput
	ToTokenMapOutputWithContext(context.Context) TokenMapOutput
}

type TokenMap map[string]TokenInput

func (TokenMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Token)(nil)).Elem()
}

func (i TokenMap) ToTokenMapOutput() TokenMapOutput {
	return i.ToTokenMapOutputWithContext(context.Background())
}

func (i TokenMap) ToTokenMapOutputWithContext(ctx context.Context) TokenMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenMapOutput)
}

func (i TokenMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Token] {
	return pulumix.Output[map[string]*Token]{
		OutputState: i.ToTokenMapOutputWithContext(ctx).OutputState,
	}
}

type TokenOutput struct{ *pulumi.OutputState }

func (TokenOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Token)(nil)).Elem()
}

func (o TokenOutput) ToTokenOutput() TokenOutput {
	return o
}

func (o TokenOutput) ToTokenOutputWithContext(ctx context.Context) TokenOutput {
	return o
}

func (o TokenOutput) ToOutput(ctx context.Context) pulumix.Output[*Token] {
	return pulumix.Output[*Token]{
		OutputState: o.OutputState,
	}
}

// String containing the client token if stored in present file
func (o TokenOutput) ClientToken() pulumi.StringOutput {
	return o.ApplyT(func(v *Token) pulumi.StringOutput { return v.ClientToken }).(pulumi.StringOutput)
}

// String containing the token display name
func (o TokenOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Token) pulumi.StringPtrOutput { return v.DisplayName }).(pulumi.StringPtrOutput)
}

// The explicit max TTL of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
func (o TokenOutput) ExplicitMaxTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Token) pulumi.StringPtrOutput { return v.ExplicitMaxTtl }).(pulumi.StringPtrOutput)
}

// String containing the token lease duration if present in state file
func (o TokenOutput) LeaseDuration() pulumi.IntOutput {
	return o.ApplyT(func(v *Token) pulumi.IntOutput { return v.LeaseDuration }).(pulumi.IntOutput)
}

// String containing the token lease started time if present in state file
func (o TokenOutput) LeaseStarted() pulumi.StringOutput {
	return o.ApplyT(func(v *Token) pulumi.StringOutput { return v.LeaseStarted }).(pulumi.StringOutput)
}

// Metadata to be set on this token
func (o TokenOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Token) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o TokenOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Token) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Flag to not attach the default policy to this token
func (o TokenOutput) NoDefaultPolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Token) pulumi.BoolPtrOutput { return v.NoDefaultPolicy }).(pulumi.BoolPtrOutput)
}

// Flag to create a token without parent
func (o TokenOutput) NoParent() pulumi.BoolOutput {
	return o.ApplyT(func(v *Token) pulumi.BoolOutput { return v.NoParent }).(pulumi.BoolOutput)
}

// The number of allowed uses of this token
func (o TokenOutput) NumUses() pulumi.IntOutput {
	return o.ApplyT(func(v *Token) pulumi.IntOutput { return v.NumUses }).(pulumi.IntOutput)
}

// The period of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
func (o TokenOutput) Period() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Token) pulumi.StringPtrOutput { return v.Period }).(pulumi.StringPtrOutput)
}

// List of policies to attach to this token
func (o TokenOutput) Policies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Token) pulumi.StringArrayOutput { return v.Policies }).(pulumi.StringArrayOutput)
}

// The renew increment. This is specified in seconds
func (o TokenOutput) RenewIncrement() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Token) pulumi.IntPtrOutput { return v.RenewIncrement }).(pulumi.IntPtrOutput)
}

// The minimal lease to renew this token
func (o TokenOutput) RenewMinLease() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Token) pulumi.IntPtrOutput { return v.RenewMinLease }).(pulumi.IntPtrOutput)
}

// Flag to allow to renew this token
func (o TokenOutput) Renewable() pulumi.BoolOutput {
	return o.ApplyT(func(v *Token) pulumi.BoolOutput { return v.Renewable }).(pulumi.BoolOutput)
}

// The token role name
func (o TokenOutput) RoleName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Token) pulumi.StringPtrOutput { return v.RoleName }).(pulumi.StringPtrOutput)
}

// The TTL period of this token. This is specified as a numeric string with suffix like "30s" ro "5m"
func (o TokenOutput) Ttl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Token) pulumi.StringPtrOutput { return v.Ttl }).(pulumi.StringPtrOutput)
}

// The client wrapped token.
func (o TokenOutput) WrappedToken() pulumi.StringOutput {
	return o.ApplyT(func(v *Token) pulumi.StringOutput { return v.WrappedToken }).(pulumi.StringOutput)
}

// The client wrapping accessor.
func (o TokenOutput) WrappingAccessor() pulumi.StringOutput {
	return o.ApplyT(func(v *Token) pulumi.StringOutput { return v.WrappingAccessor }).(pulumi.StringOutput)
}

// The TTL period of the wrapped token.
func (o TokenOutput) WrappingTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Token) pulumi.StringPtrOutput { return v.WrappingTtl }).(pulumi.StringPtrOutput)
}

type TokenArrayOutput struct{ *pulumi.OutputState }

func (TokenArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Token)(nil)).Elem()
}

func (o TokenArrayOutput) ToTokenArrayOutput() TokenArrayOutput {
	return o
}

func (o TokenArrayOutput) ToTokenArrayOutputWithContext(ctx context.Context) TokenArrayOutput {
	return o
}

func (o TokenArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Token] {
	return pulumix.Output[[]*Token]{
		OutputState: o.OutputState,
	}
}

func (o TokenArrayOutput) Index(i pulumi.IntInput) TokenOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Token {
		return vs[0].([]*Token)[vs[1].(int)]
	}).(TokenOutput)
}

type TokenMapOutput struct{ *pulumi.OutputState }

func (TokenMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Token)(nil)).Elem()
}

func (o TokenMapOutput) ToTokenMapOutput() TokenMapOutput {
	return o
}

func (o TokenMapOutput) ToTokenMapOutputWithContext(ctx context.Context) TokenMapOutput {
	return o
}

func (o TokenMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Token] {
	return pulumix.Output[map[string]*Token]{
		OutputState: o.OutputState,
	}
}

func (o TokenMapOutput) MapIndex(k pulumi.StringInput) TokenOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Token {
		return vs[0].(map[string]*Token)[vs[1].(string)]
	}).(TokenOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TokenInput)(nil)).Elem(), &Token{})
	pulumi.RegisterInputType(reflect.TypeOf((*TokenArrayInput)(nil)).Elem(), TokenArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TokenMapInput)(nil)).Elem(), TokenMap{})
	pulumi.RegisterOutputType(TokenOutput{})
	pulumi.RegisterOutputType(TokenArrayOutput{})
	pulumi.RegisterOutputType(TokenMapOutput{})
}
