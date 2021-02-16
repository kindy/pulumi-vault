// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// The key can be imported with the role name, for example
//
// ```sh
//  $ pulumi import vault:identity/oidcRole:OidcRole role role
// ```
type OidcRole struct {
	pulumi.CustomResourceState

	// The value that will be included in the `aud` field of all the OIDC identity
	// tokens issued by this role
	ClientId pulumi.StringOutput `pulumi:"clientId"`
	// A configured named key, the key must already exist
	// before tokens can be issued.
	Key pulumi.StringOutput `pulumi:"key"`
	// Name of the OIDC Role to create.
	Name pulumi.StringOutput `pulumi:"name"`
	// The template string to use for generating tokens. This may be in
	// string-ified JSON or base64 format. See the
	// [documentation](https://www.vaultproject.io/docs/secrets/identity/index.html#token-contents-and-templates)
	// for the template format.
	Template pulumi.StringPtrOutput `pulumi:"template"`
	// TTL of the tokens generated against the role in number of seconds.
	Ttl pulumi.IntPtrOutput `pulumi:"ttl"`
}

// NewOidcRole registers a new resource with the given unique name, arguments, and options.
func NewOidcRole(ctx *pulumi.Context,
	name string, args *OidcRoleArgs, opts ...pulumi.ResourceOption) (*OidcRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Key == nil {
		return nil, errors.New("invalid value for required argument 'Key'")
	}
	var resource OidcRole
	err := ctx.RegisterResource("vault:identity/oidcRole:OidcRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOidcRole gets an existing OidcRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOidcRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OidcRoleState, opts ...pulumi.ResourceOption) (*OidcRole, error) {
	var resource OidcRole
	err := ctx.ReadResource("vault:identity/oidcRole:OidcRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OidcRole resources.
type oidcRoleState struct {
	// The value that will be included in the `aud` field of all the OIDC identity
	// tokens issued by this role
	ClientId *string `pulumi:"clientId"`
	// A configured named key, the key must already exist
	// before tokens can be issued.
	Key *string `pulumi:"key"`
	// Name of the OIDC Role to create.
	Name *string `pulumi:"name"`
	// The template string to use for generating tokens. This may be in
	// string-ified JSON or base64 format. See the
	// [documentation](https://www.vaultproject.io/docs/secrets/identity/index.html#token-contents-and-templates)
	// for the template format.
	Template *string `pulumi:"template"`
	// TTL of the tokens generated against the role in number of seconds.
	Ttl *int `pulumi:"ttl"`
}

type OidcRoleState struct {
	// The value that will be included in the `aud` field of all the OIDC identity
	// tokens issued by this role
	ClientId pulumi.StringPtrInput
	// A configured named key, the key must already exist
	// before tokens can be issued.
	Key pulumi.StringPtrInput
	// Name of the OIDC Role to create.
	Name pulumi.StringPtrInput
	// The template string to use for generating tokens. This may be in
	// string-ified JSON or base64 format. See the
	// [documentation](https://www.vaultproject.io/docs/secrets/identity/index.html#token-contents-and-templates)
	// for the template format.
	Template pulumi.StringPtrInput
	// TTL of the tokens generated against the role in number of seconds.
	Ttl pulumi.IntPtrInput
}

func (OidcRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*oidcRoleState)(nil)).Elem()
}

type oidcRoleArgs struct {
	// The value that will be included in the `aud` field of all the OIDC identity
	// tokens issued by this role
	ClientId *string `pulumi:"clientId"`
	// A configured named key, the key must already exist
	// before tokens can be issued.
	Key string `pulumi:"key"`
	// Name of the OIDC Role to create.
	Name *string `pulumi:"name"`
	// The template string to use for generating tokens. This may be in
	// string-ified JSON or base64 format. See the
	// [documentation](https://www.vaultproject.io/docs/secrets/identity/index.html#token-contents-and-templates)
	// for the template format.
	Template *string `pulumi:"template"`
	// TTL of the tokens generated against the role in number of seconds.
	Ttl *int `pulumi:"ttl"`
}

// The set of arguments for constructing a OidcRole resource.
type OidcRoleArgs struct {
	// The value that will be included in the `aud` field of all the OIDC identity
	// tokens issued by this role
	ClientId pulumi.StringPtrInput
	// A configured named key, the key must already exist
	// before tokens can be issued.
	Key pulumi.StringInput
	// Name of the OIDC Role to create.
	Name pulumi.StringPtrInput
	// The template string to use for generating tokens. This may be in
	// string-ified JSON or base64 format. See the
	// [documentation](https://www.vaultproject.io/docs/secrets/identity/index.html#token-contents-and-templates)
	// for the template format.
	Template pulumi.StringPtrInput
	// TTL of the tokens generated against the role in number of seconds.
	Ttl pulumi.IntPtrInput
}

func (OidcRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*oidcRoleArgs)(nil)).Elem()
}

type OidcRoleInput interface {
	pulumi.Input

	ToOidcRoleOutput() OidcRoleOutput
	ToOidcRoleOutputWithContext(ctx context.Context) OidcRoleOutput
}

func (*OidcRole) ElementType() reflect.Type {
	return reflect.TypeOf((*OidcRole)(nil))
}

func (i *OidcRole) ToOidcRoleOutput() OidcRoleOutput {
	return i.ToOidcRoleOutputWithContext(context.Background())
}

func (i *OidcRole) ToOidcRoleOutputWithContext(ctx context.Context) OidcRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OidcRoleOutput)
}

func (i *OidcRole) ToOidcRolePtrOutput() OidcRolePtrOutput {
	return i.ToOidcRolePtrOutputWithContext(context.Background())
}

func (i *OidcRole) ToOidcRolePtrOutputWithContext(ctx context.Context) OidcRolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OidcRolePtrOutput)
}

type OidcRolePtrInput interface {
	pulumi.Input

	ToOidcRolePtrOutput() OidcRolePtrOutput
	ToOidcRolePtrOutputWithContext(ctx context.Context) OidcRolePtrOutput
}

type oidcRolePtrType OidcRoleArgs

func (*oidcRolePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OidcRole)(nil))
}

func (i *oidcRolePtrType) ToOidcRolePtrOutput() OidcRolePtrOutput {
	return i.ToOidcRolePtrOutputWithContext(context.Background())
}

func (i *oidcRolePtrType) ToOidcRolePtrOutputWithContext(ctx context.Context) OidcRolePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OidcRolePtrOutput)
}

// OidcRoleArrayInput is an input type that accepts OidcRoleArray and OidcRoleArrayOutput values.
// You can construct a concrete instance of `OidcRoleArrayInput` via:
//
//          OidcRoleArray{ OidcRoleArgs{...} }
type OidcRoleArrayInput interface {
	pulumi.Input

	ToOidcRoleArrayOutput() OidcRoleArrayOutput
	ToOidcRoleArrayOutputWithContext(context.Context) OidcRoleArrayOutput
}

type OidcRoleArray []OidcRoleInput

func (OidcRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*OidcRole)(nil))
}

func (i OidcRoleArray) ToOidcRoleArrayOutput() OidcRoleArrayOutput {
	return i.ToOidcRoleArrayOutputWithContext(context.Background())
}

func (i OidcRoleArray) ToOidcRoleArrayOutputWithContext(ctx context.Context) OidcRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OidcRoleArrayOutput)
}

// OidcRoleMapInput is an input type that accepts OidcRoleMap and OidcRoleMapOutput values.
// You can construct a concrete instance of `OidcRoleMapInput` via:
//
//          OidcRoleMap{ "key": OidcRoleArgs{...} }
type OidcRoleMapInput interface {
	pulumi.Input

	ToOidcRoleMapOutput() OidcRoleMapOutput
	ToOidcRoleMapOutputWithContext(context.Context) OidcRoleMapOutput
}

type OidcRoleMap map[string]OidcRoleInput

func (OidcRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*OidcRole)(nil))
}

func (i OidcRoleMap) ToOidcRoleMapOutput() OidcRoleMapOutput {
	return i.ToOidcRoleMapOutputWithContext(context.Background())
}

func (i OidcRoleMap) ToOidcRoleMapOutputWithContext(ctx context.Context) OidcRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OidcRoleMapOutput)
}

type OidcRoleOutput struct {
	*pulumi.OutputState
}

func (OidcRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OidcRole)(nil))
}

func (o OidcRoleOutput) ToOidcRoleOutput() OidcRoleOutput {
	return o
}

func (o OidcRoleOutput) ToOidcRoleOutputWithContext(ctx context.Context) OidcRoleOutput {
	return o
}

func (o OidcRoleOutput) ToOidcRolePtrOutput() OidcRolePtrOutput {
	return o.ToOidcRolePtrOutputWithContext(context.Background())
}

func (o OidcRoleOutput) ToOidcRolePtrOutputWithContext(ctx context.Context) OidcRolePtrOutput {
	return o.ApplyT(func(v OidcRole) *OidcRole {
		return &v
	}).(OidcRolePtrOutput)
}

type OidcRolePtrOutput struct {
	*pulumi.OutputState
}

func (OidcRolePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OidcRole)(nil))
}

func (o OidcRolePtrOutput) ToOidcRolePtrOutput() OidcRolePtrOutput {
	return o
}

func (o OidcRolePtrOutput) ToOidcRolePtrOutputWithContext(ctx context.Context) OidcRolePtrOutput {
	return o
}

type OidcRoleArrayOutput struct{ *pulumi.OutputState }

func (OidcRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]OidcRole)(nil))
}

func (o OidcRoleArrayOutput) ToOidcRoleArrayOutput() OidcRoleArrayOutput {
	return o
}

func (o OidcRoleArrayOutput) ToOidcRoleArrayOutputWithContext(ctx context.Context) OidcRoleArrayOutput {
	return o
}

func (o OidcRoleArrayOutput) Index(i pulumi.IntInput) OidcRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) OidcRole {
		return vs[0].([]OidcRole)[vs[1].(int)]
	}).(OidcRoleOutput)
}

type OidcRoleMapOutput struct{ *pulumi.OutputState }

func (OidcRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]OidcRole)(nil))
}

func (o OidcRoleMapOutput) ToOidcRoleMapOutput() OidcRoleMapOutput {
	return o
}

func (o OidcRoleMapOutput) ToOidcRoleMapOutputWithContext(ctx context.Context) OidcRoleMapOutput {
	return o
}

func (o OidcRoleMapOutput) MapIndex(k pulumi.StringInput) OidcRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) OidcRole {
		return vs[0].(map[string]OidcRole)[vs[1].(string)]
	}).(OidcRoleOutput)
}

func init() {
	pulumi.RegisterOutputType(OidcRoleOutput{})
	pulumi.RegisterOutputType(OidcRolePtrOutput{})
	pulumi.RegisterOutputType(OidcRoleArrayOutput{})
	pulumi.RegisterOutputType(OidcRoleMapOutput{})
}
