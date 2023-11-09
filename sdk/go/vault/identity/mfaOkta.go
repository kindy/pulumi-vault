// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Resource for configuring the okta MFA method.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/identity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := identity.NewMfaOkta(ctx, "example", &identity.MfaOktaArgs{
//				ApiToken: pulumi.String("token1"),
//				BaseUrl:  pulumi.String("qux.baz.com"),
//				OrgName:  pulumi.String("org1"),
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
// Resource can be imported using its `uuid` field, e.g.
//
// ```sh
//
//	$ pulumi import vault:identity/mfaOkta:MfaOkta example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec
//
// ```
type MfaOkta struct {
	pulumi.CustomResourceState

	// Okta API token.
	ApiToken pulumi.StringOutput `pulumi:"apiToken"`
	// The base domain to use for API requests.
	BaseUrl pulumi.StringPtrOutput `pulumi:"baseUrl"`
	// Method ID.
	MethodId pulumi.StringOutput `pulumi:"methodId"`
	// Mount accessor.
	MountAccessor pulumi.StringOutput `pulumi:"mountAccessor"`
	// Method name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Target namespace. (requires Enterprise)
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Method's namespace ID.
	NamespaceId pulumi.StringOutput `pulumi:"namespaceId"`
	// Method's namespace path.
	NamespacePath pulumi.StringOutput `pulumi:"namespacePath"`
	// Name of the organization to be used in the Okta API.
	OrgName pulumi.StringOutput `pulumi:"orgName"`
	// Only match the primary email for the account.
	PrimaryEmail pulumi.BoolPtrOutput `pulumi:"primaryEmail"`
	// MFA type.
	Type pulumi.StringOutput `pulumi:"type"`
	// A template string for mapping Identity names to MFA methods.
	UsernameFormat pulumi.StringPtrOutput `pulumi:"usernameFormat"`
	// Resource UUID.
	Uuid pulumi.StringOutput `pulumi:"uuid"`
}

// NewMfaOkta registers a new resource with the given unique name, arguments, and options.
func NewMfaOkta(ctx *pulumi.Context,
	name string, args *MfaOktaArgs, opts ...pulumi.ResourceOption) (*MfaOkta, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApiToken == nil {
		return nil, errors.New("invalid value for required argument 'ApiToken'")
	}
	if args.OrgName == nil {
		return nil, errors.New("invalid value for required argument 'OrgName'")
	}
	if args.ApiToken != nil {
		args.ApiToken = pulumi.ToSecret(args.ApiToken).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"apiToken",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MfaOkta
	err := ctx.RegisterResource("vault:identity/mfaOkta:MfaOkta", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMfaOkta gets an existing MfaOkta resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMfaOkta(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MfaOktaState, opts ...pulumi.ResourceOption) (*MfaOkta, error) {
	var resource MfaOkta
	err := ctx.ReadResource("vault:identity/mfaOkta:MfaOkta", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MfaOkta resources.
type mfaOktaState struct {
	// Okta API token.
	ApiToken *string `pulumi:"apiToken"`
	// The base domain to use for API requests.
	BaseUrl *string `pulumi:"baseUrl"`
	// Method ID.
	MethodId *string `pulumi:"methodId"`
	// Mount accessor.
	MountAccessor *string `pulumi:"mountAccessor"`
	// Method name.
	Name *string `pulumi:"name"`
	// Target namespace. (requires Enterprise)
	Namespace *string `pulumi:"namespace"`
	// Method's namespace ID.
	NamespaceId *string `pulumi:"namespaceId"`
	// Method's namespace path.
	NamespacePath *string `pulumi:"namespacePath"`
	// Name of the organization to be used in the Okta API.
	OrgName *string `pulumi:"orgName"`
	// Only match the primary email for the account.
	PrimaryEmail *bool `pulumi:"primaryEmail"`
	// MFA type.
	Type *string `pulumi:"type"`
	// A template string for mapping Identity names to MFA methods.
	UsernameFormat *string `pulumi:"usernameFormat"`
	// Resource UUID.
	Uuid *string `pulumi:"uuid"`
}

type MfaOktaState struct {
	// Okta API token.
	ApiToken pulumi.StringPtrInput
	// The base domain to use for API requests.
	BaseUrl pulumi.StringPtrInput
	// Method ID.
	MethodId pulumi.StringPtrInput
	// Mount accessor.
	MountAccessor pulumi.StringPtrInput
	// Method name.
	Name pulumi.StringPtrInput
	// Target namespace. (requires Enterprise)
	Namespace pulumi.StringPtrInput
	// Method's namespace ID.
	NamespaceId pulumi.StringPtrInput
	// Method's namespace path.
	NamespacePath pulumi.StringPtrInput
	// Name of the organization to be used in the Okta API.
	OrgName pulumi.StringPtrInput
	// Only match the primary email for the account.
	PrimaryEmail pulumi.BoolPtrInput
	// MFA type.
	Type pulumi.StringPtrInput
	// A template string for mapping Identity names to MFA methods.
	UsernameFormat pulumi.StringPtrInput
	// Resource UUID.
	Uuid pulumi.StringPtrInput
}

func (MfaOktaState) ElementType() reflect.Type {
	return reflect.TypeOf((*mfaOktaState)(nil)).Elem()
}

type mfaOktaArgs struct {
	// Okta API token.
	ApiToken string `pulumi:"apiToken"`
	// The base domain to use for API requests.
	BaseUrl *string `pulumi:"baseUrl"`
	// Target namespace. (requires Enterprise)
	Namespace *string `pulumi:"namespace"`
	// Name of the organization to be used in the Okta API.
	OrgName string `pulumi:"orgName"`
	// Only match the primary email for the account.
	PrimaryEmail *bool `pulumi:"primaryEmail"`
	// A template string for mapping Identity names to MFA methods.
	UsernameFormat *string `pulumi:"usernameFormat"`
}

// The set of arguments for constructing a MfaOkta resource.
type MfaOktaArgs struct {
	// Okta API token.
	ApiToken pulumi.StringInput
	// The base domain to use for API requests.
	BaseUrl pulumi.StringPtrInput
	// Target namespace. (requires Enterprise)
	Namespace pulumi.StringPtrInput
	// Name of the organization to be used in the Okta API.
	OrgName pulumi.StringInput
	// Only match the primary email for the account.
	PrimaryEmail pulumi.BoolPtrInput
	// A template string for mapping Identity names to MFA methods.
	UsernameFormat pulumi.StringPtrInput
}

func (MfaOktaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mfaOktaArgs)(nil)).Elem()
}

type MfaOktaInput interface {
	pulumi.Input

	ToMfaOktaOutput() MfaOktaOutput
	ToMfaOktaOutputWithContext(ctx context.Context) MfaOktaOutput
}

func (*MfaOkta) ElementType() reflect.Type {
	return reflect.TypeOf((**MfaOkta)(nil)).Elem()
}

func (i *MfaOkta) ToMfaOktaOutput() MfaOktaOutput {
	return i.ToMfaOktaOutputWithContext(context.Background())
}

func (i *MfaOkta) ToMfaOktaOutputWithContext(ctx context.Context) MfaOktaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MfaOktaOutput)
}

func (i *MfaOkta) ToOutput(ctx context.Context) pulumix.Output[*MfaOkta] {
	return pulumix.Output[*MfaOkta]{
		OutputState: i.ToMfaOktaOutputWithContext(ctx).OutputState,
	}
}

// MfaOktaArrayInput is an input type that accepts MfaOktaArray and MfaOktaArrayOutput values.
// You can construct a concrete instance of `MfaOktaArrayInput` via:
//
//	MfaOktaArray{ MfaOktaArgs{...} }
type MfaOktaArrayInput interface {
	pulumi.Input

	ToMfaOktaArrayOutput() MfaOktaArrayOutput
	ToMfaOktaArrayOutputWithContext(context.Context) MfaOktaArrayOutput
}

type MfaOktaArray []MfaOktaInput

func (MfaOktaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MfaOkta)(nil)).Elem()
}

func (i MfaOktaArray) ToMfaOktaArrayOutput() MfaOktaArrayOutput {
	return i.ToMfaOktaArrayOutputWithContext(context.Background())
}

func (i MfaOktaArray) ToMfaOktaArrayOutputWithContext(ctx context.Context) MfaOktaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MfaOktaArrayOutput)
}

func (i MfaOktaArray) ToOutput(ctx context.Context) pulumix.Output[[]*MfaOkta] {
	return pulumix.Output[[]*MfaOkta]{
		OutputState: i.ToMfaOktaArrayOutputWithContext(ctx).OutputState,
	}
}

// MfaOktaMapInput is an input type that accepts MfaOktaMap and MfaOktaMapOutput values.
// You can construct a concrete instance of `MfaOktaMapInput` via:
//
//	MfaOktaMap{ "key": MfaOktaArgs{...} }
type MfaOktaMapInput interface {
	pulumi.Input

	ToMfaOktaMapOutput() MfaOktaMapOutput
	ToMfaOktaMapOutputWithContext(context.Context) MfaOktaMapOutput
}

type MfaOktaMap map[string]MfaOktaInput

func (MfaOktaMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MfaOkta)(nil)).Elem()
}

func (i MfaOktaMap) ToMfaOktaMapOutput() MfaOktaMapOutput {
	return i.ToMfaOktaMapOutputWithContext(context.Background())
}

func (i MfaOktaMap) ToMfaOktaMapOutputWithContext(ctx context.Context) MfaOktaMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MfaOktaMapOutput)
}

func (i MfaOktaMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*MfaOkta] {
	return pulumix.Output[map[string]*MfaOkta]{
		OutputState: i.ToMfaOktaMapOutputWithContext(ctx).OutputState,
	}
}

type MfaOktaOutput struct{ *pulumi.OutputState }

func (MfaOktaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MfaOkta)(nil)).Elem()
}

func (o MfaOktaOutput) ToMfaOktaOutput() MfaOktaOutput {
	return o
}

func (o MfaOktaOutput) ToMfaOktaOutputWithContext(ctx context.Context) MfaOktaOutput {
	return o
}

func (o MfaOktaOutput) ToOutput(ctx context.Context) pulumix.Output[*MfaOkta] {
	return pulumix.Output[*MfaOkta]{
		OutputState: o.OutputState,
	}
}

// Okta API token.
func (o MfaOktaOutput) ApiToken() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaOkta) pulumi.StringOutput { return v.ApiToken }).(pulumi.StringOutput)
}

// The base domain to use for API requests.
func (o MfaOktaOutput) BaseUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MfaOkta) pulumi.StringPtrOutput { return v.BaseUrl }).(pulumi.StringPtrOutput)
}

// Method ID.
func (o MfaOktaOutput) MethodId() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaOkta) pulumi.StringOutput { return v.MethodId }).(pulumi.StringOutput)
}

// Mount accessor.
func (o MfaOktaOutput) MountAccessor() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaOkta) pulumi.StringOutput { return v.MountAccessor }).(pulumi.StringOutput)
}

// Method name.
func (o MfaOktaOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaOkta) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Target namespace. (requires Enterprise)
func (o MfaOktaOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MfaOkta) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Method's namespace ID.
func (o MfaOktaOutput) NamespaceId() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaOkta) pulumi.StringOutput { return v.NamespaceId }).(pulumi.StringOutput)
}

// Method's namespace path.
func (o MfaOktaOutput) NamespacePath() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaOkta) pulumi.StringOutput { return v.NamespacePath }).(pulumi.StringOutput)
}

// Name of the organization to be used in the Okta API.
func (o MfaOktaOutput) OrgName() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaOkta) pulumi.StringOutput { return v.OrgName }).(pulumi.StringOutput)
}

// Only match the primary email for the account.
func (o MfaOktaOutput) PrimaryEmail() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MfaOkta) pulumi.BoolPtrOutput { return v.PrimaryEmail }).(pulumi.BoolPtrOutput)
}

// MFA type.
func (o MfaOktaOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaOkta) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// A template string for mapping Identity names to MFA methods.
func (o MfaOktaOutput) UsernameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MfaOkta) pulumi.StringPtrOutput { return v.UsernameFormat }).(pulumi.StringPtrOutput)
}

// Resource UUID.
func (o MfaOktaOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v *MfaOkta) pulumi.StringOutput { return v.Uuid }).(pulumi.StringOutput)
}

type MfaOktaArrayOutput struct{ *pulumi.OutputState }

func (MfaOktaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MfaOkta)(nil)).Elem()
}

func (o MfaOktaArrayOutput) ToMfaOktaArrayOutput() MfaOktaArrayOutput {
	return o
}

func (o MfaOktaArrayOutput) ToMfaOktaArrayOutputWithContext(ctx context.Context) MfaOktaArrayOutput {
	return o
}

func (o MfaOktaArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*MfaOkta] {
	return pulumix.Output[[]*MfaOkta]{
		OutputState: o.OutputState,
	}
}

func (o MfaOktaArrayOutput) Index(i pulumi.IntInput) MfaOktaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MfaOkta {
		return vs[0].([]*MfaOkta)[vs[1].(int)]
	}).(MfaOktaOutput)
}

type MfaOktaMapOutput struct{ *pulumi.OutputState }

func (MfaOktaMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MfaOkta)(nil)).Elem()
}

func (o MfaOktaMapOutput) ToMfaOktaMapOutput() MfaOktaMapOutput {
	return o
}

func (o MfaOktaMapOutput) ToMfaOktaMapOutputWithContext(ctx context.Context) MfaOktaMapOutput {
	return o
}

func (o MfaOktaMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*MfaOkta] {
	return pulumix.Output[map[string]*MfaOkta]{
		OutputState: o.OutputState,
	}
}

func (o MfaOktaMapOutput) MapIndex(k pulumi.StringInput) MfaOktaOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MfaOkta {
		return vs[0].(map[string]*MfaOkta)[vs[1].(string)]
	}).(MfaOktaOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MfaOktaInput)(nil)).Elem(), &MfaOkta{})
	pulumi.RegisterInputType(reflect.TypeOf((*MfaOktaArrayInput)(nil)).Elem(), MfaOktaArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MfaOktaMapInput)(nil)).Elem(), MfaOktaMap{})
	pulumi.RegisterOutputType(MfaOktaOutput{})
	pulumi.RegisterOutputType(MfaOktaArrayOutput{})
	pulumi.RegisterOutputType(MfaOktaMapOutput{})
}
