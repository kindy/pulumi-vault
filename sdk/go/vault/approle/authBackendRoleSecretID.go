// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package approle

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an AppRole auth backend SecretID in a Vault server. See the [Vault
// documentation](https://www.vaultproject.io/docs/auth/approle) for more
// information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-vault/sdk/v4/go/vault"
// 	"github.com/pulumi/pulumi-vault/sdk/v4/go/vault/appRole"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		approle, err := vault.NewAuthBackend(ctx, "approle", &vault.AuthBackendArgs{
// 			Type: pulumi.String("approle"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		example, err := appRole.NewAuthBackendRole(ctx, "example", &appRole.AuthBackendRoleArgs{
// 			Backend:  approle.Path,
// 			RoleName: pulumi.String("test-role"),
// 			Policies: pulumi.StringArray{
// 				pulumi.String("default"),
// 				pulumi.String("dev"),
// 				pulumi.String("prod"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = appRole.NewAuthBackendRoleSecretID(ctx, "id", &appRole.AuthBackendRoleSecretIDArgs{
// 			Backend:  approle.Path,
// 			RoleName: example.RoleName,
// 			Metadata: pulumi.String(fmt.Sprintf("%v%v%v", "  {\n", "    \"hello\": \"world\"\n", "  }\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AuthBackendRoleSecretID struct {
	pulumi.CustomResourceState

	// The unique ID for this SecretID that can be safely logged.
	Accessor pulumi.StringOutput `pulumi:"accessor"`
	// Unique name of the auth backend to configure.
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// If set, specifies blocks of IP addresses which can
	// perform the login operation using this SecretID.
	CidrLists pulumi.StringArrayOutput `pulumi:"cidrLists"`
	// A JSON-encoded string containing metadata in
	// key-value pairs to be set on tokens issued with this SecretID.
	Metadata pulumi.StringPtrOutput `pulumi:"metadata"`
	// The name of the role to create the SecretID for.
	RoleName pulumi.StringOutput `pulumi:"roleName"`
	// The SecretID to be created. If set, uses "Push"
	// mode.  Defaults to Vault auto-generating SecretIDs.
	SecretId pulumi.StringOutput `pulumi:"secretId"`
	// The unique ID for the response-wrapped SecretID that can
	// be safely logged.
	WrappingAccessor pulumi.StringOutput `pulumi:"wrappingAccessor"`
	// The token used to retrieve a response-wrapped SecretID.
	WrappingToken pulumi.StringOutput `pulumi:"wrappingToken"`
	// If set, the SecretID response will be
	// [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
	// and available for the duration specified. Only a single unwrapping of the
	// token is allowed.
	WrappingTtl pulumi.StringPtrOutput `pulumi:"wrappingTtl"`
}

// NewAuthBackendRoleSecretID registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendRoleSecretID(ctx *pulumi.Context,
	name string, args *AuthBackendRoleSecretIDArgs, opts ...pulumi.ResourceOption) (*AuthBackendRoleSecretID, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RoleName == nil {
		return nil, errors.New("invalid value for required argument 'RoleName'")
	}
	var resource AuthBackendRoleSecretID
	err := ctx.RegisterResource("vault:appRole/authBackendRoleSecretID:AuthBackendRoleSecretID", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackendRoleSecretID gets an existing AuthBackendRoleSecretID resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendRoleSecretID(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendRoleSecretIDState, opts ...pulumi.ResourceOption) (*AuthBackendRoleSecretID, error) {
	var resource AuthBackendRoleSecretID
	err := ctx.ReadResource("vault:appRole/authBackendRoleSecretID:AuthBackendRoleSecretID", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendRoleSecretID resources.
type authBackendRoleSecretIDState struct {
	// The unique ID for this SecretID that can be safely logged.
	Accessor *string `pulumi:"accessor"`
	// Unique name of the auth backend to configure.
	Backend *string `pulumi:"backend"`
	// If set, specifies blocks of IP addresses which can
	// perform the login operation using this SecretID.
	CidrLists []string `pulumi:"cidrLists"`
	// A JSON-encoded string containing metadata in
	// key-value pairs to be set on tokens issued with this SecretID.
	Metadata *string `pulumi:"metadata"`
	// The name of the role to create the SecretID for.
	RoleName *string `pulumi:"roleName"`
	// The SecretID to be created. If set, uses "Push"
	// mode.  Defaults to Vault auto-generating SecretIDs.
	SecretId *string `pulumi:"secretId"`
	// The unique ID for the response-wrapped SecretID that can
	// be safely logged.
	WrappingAccessor *string `pulumi:"wrappingAccessor"`
	// The token used to retrieve a response-wrapped SecretID.
	WrappingToken *string `pulumi:"wrappingToken"`
	// If set, the SecretID response will be
	// [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
	// and available for the duration specified. Only a single unwrapping of the
	// token is allowed.
	WrappingTtl *string `pulumi:"wrappingTtl"`
}

type AuthBackendRoleSecretIDState struct {
	// The unique ID for this SecretID that can be safely logged.
	Accessor pulumi.StringPtrInput
	// Unique name of the auth backend to configure.
	Backend pulumi.StringPtrInput
	// If set, specifies blocks of IP addresses which can
	// perform the login operation using this SecretID.
	CidrLists pulumi.StringArrayInput
	// A JSON-encoded string containing metadata in
	// key-value pairs to be set on tokens issued with this SecretID.
	Metadata pulumi.StringPtrInput
	// The name of the role to create the SecretID for.
	RoleName pulumi.StringPtrInput
	// The SecretID to be created. If set, uses "Push"
	// mode.  Defaults to Vault auto-generating SecretIDs.
	SecretId pulumi.StringPtrInput
	// The unique ID for the response-wrapped SecretID that can
	// be safely logged.
	WrappingAccessor pulumi.StringPtrInput
	// The token used to retrieve a response-wrapped SecretID.
	WrappingToken pulumi.StringPtrInput
	// If set, the SecretID response will be
	// [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
	// and available for the duration specified. Only a single unwrapping of the
	// token is allowed.
	WrappingTtl pulumi.StringPtrInput
}

func (AuthBackendRoleSecretIDState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendRoleSecretIDState)(nil)).Elem()
}

type authBackendRoleSecretIDArgs struct {
	// Unique name of the auth backend to configure.
	Backend *string `pulumi:"backend"`
	// If set, specifies blocks of IP addresses which can
	// perform the login operation using this SecretID.
	CidrLists []string `pulumi:"cidrLists"`
	// A JSON-encoded string containing metadata in
	// key-value pairs to be set on tokens issued with this SecretID.
	Metadata *string `pulumi:"metadata"`
	// The name of the role to create the SecretID for.
	RoleName string `pulumi:"roleName"`
	// The SecretID to be created. If set, uses "Push"
	// mode.  Defaults to Vault auto-generating SecretIDs.
	SecretId *string `pulumi:"secretId"`
	// If set, the SecretID response will be
	// [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
	// and available for the duration specified. Only a single unwrapping of the
	// token is allowed.
	WrappingTtl *string `pulumi:"wrappingTtl"`
}

// The set of arguments for constructing a AuthBackendRoleSecretID resource.
type AuthBackendRoleSecretIDArgs struct {
	// Unique name of the auth backend to configure.
	Backend pulumi.StringPtrInput
	// If set, specifies blocks of IP addresses which can
	// perform the login operation using this SecretID.
	CidrLists pulumi.StringArrayInput
	// A JSON-encoded string containing metadata in
	// key-value pairs to be set on tokens issued with this SecretID.
	Metadata pulumi.StringPtrInput
	// The name of the role to create the SecretID for.
	RoleName pulumi.StringInput
	// The SecretID to be created. If set, uses "Push"
	// mode.  Defaults to Vault auto-generating SecretIDs.
	SecretId pulumi.StringPtrInput
	// If set, the SecretID response will be
	// [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
	// and available for the duration specified. Only a single unwrapping of the
	// token is allowed.
	WrappingTtl pulumi.StringPtrInput
}

func (AuthBackendRoleSecretIDArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendRoleSecretIDArgs)(nil)).Elem()
}

type AuthBackendRoleSecretIDInput interface {
	pulumi.Input

	ToAuthBackendRoleSecretIDOutput() AuthBackendRoleSecretIDOutput
	ToAuthBackendRoleSecretIDOutputWithContext(ctx context.Context) AuthBackendRoleSecretIDOutput
}

func (*AuthBackendRoleSecretID) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthBackendRoleSecretID)(nil))
}

func (i *AuthBackendRoleSecretID) ToAuthBackendRoleSecretIDOutput() AuthBackendRoleSecretIDOutput {
	return i.ToAuthBackendRoleSecretIDOutputWithContext(context.Background())
}

func (i *AuthBackendRoleSecretID) ToAuthBackendRoleSecretIDOutputWithContext(ctx context.Context) AuthBackendRoleSecretIDOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendRoleSecretIDOutput)
}

func (i *AuthBackendRoleSecretID) ToAuthBackendRoleSecretIDPtrOutput() AuthBackendRoleSecretIDPtrOutput {
	return i.ToAuthBackendRoleSecretIDPtrOutputWithContext(context.Background())
}

func (i *AuthBackendRoleSecretID) ToAuthBackendRoleSecretIDPtrOutputWithContext(ctx context.Context) AuthBackendRoleSecretIDPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendRoleSecretIDPtrOutput)
}

type AuthBackendRoleSecretIDPtrInput interface {
	pulumi.Input

	ToAuthBackendRoleSecretIDPtrOutput() AuthBackendRoleSecretIDPtrOutput
	ToAuthBackendRoleSecretIDPtrOutputWithContext(ctx context.Context) AuthBackendRoleSecretIDPtrOutput
}

type authBackendRoleSecretIDPtrType AuthBackendRoleSecretIDArgs

func (*authBackendRoleSecretIDPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendRoleSecretID)(nil))
}

func (i *authBackendRoleSecretIDPtrType) ToAuthBackendRoleSecretIDPtrOutput() AuthBackendRoleSecretIDPtrOutput {
	return i.ToAuthBackendRoleSecretIDPtrOutputWithContext(context.Background())
}

func (i *authBackendRoleSecretIDPtrType) ToAuthBackendRoleSecretIDPtrOutputWithContext(ctx context.Context) AuthBackendRoleSecretIDPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendRoleSecretIDPtrOutput)
}

// AuthBackendRoleSecretIDArrayInput is an input type that accepts AuthBackendRoleSecretIDArray and AuthBackendRoleSecretIDArrayOutput values.
// You can construct a concrete instance of `AuthBackendRoleSecretIDArrayInput` via:
//
//          AuthBackendRoleSecretIDArray{ AuthBackendRoleSecretIDArgs{...} }
type AuthBackendRoleSecretIDArrayInput interface {
	pulumi.Input

	ToAuthBackendRoleSecretIDArrayOutput() AuthBackendRoleSecretIDArrayOutput
	ToAuthBackendRoleSecretIDArrayOutputWithContext(context.Context) AuthBackendRoleSecretIDArrayOutput
}

type AuthBackendRoleSecretIDArray []AuthBackendRoleSecretIDInput

func (AuthBackendRoleSecretIDArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AuthBackendRoleSecretID)(nil))
}

func (i AuthBackendRoleSecretIDArray) ToAuthBackendRoleSecretIDArrayOutput() AuthBackendRoleSecretIDArrayOutput {
	return i.ToAuthBackendRoleSecretIDArrayOutputWithContext(context.Background())
}

func (i AuthBackendRoleSecretIDArray) ToAuthBackendRoleSecretIDArrayOutputWithContext(ctx context.Context) AuthBackendRoleSecretIDArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendRoleSecretIDArrayOutput)
}

// AuthBackendRoleSecretIDMapInput is an input type that accepts AuthBackendRoleSecretIDMap and AuthBackendRoleSecretIDMapOutput values.
// You can construct a concrete instance of `AuthBackendRoleSecretIDMapInput` via:
//
//          AuthBackendRoleSecretIDMap{ "key": AuthBackendRoleSecretIDArgs{...} }
type AuthBackendRoleSecretIDMapInput interface {
	pulumi.Input

	ToAuthBackendRoleSecretIDMapOutput() AuthBackendRoleSecretIDMapOutput
	ToAuthBackendRoleSecretIDMapOutputWithContext(context.Context) AuthBackendRoleSecretIDMapOutput
}

type AuthBackendRoleSecretIDMap map[string]AuthBackendRoleSecretIDInput

func (AuthBackendRoleSecretIDMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AuthBackendRoleSecretID)(nil))
}

func (i AuthBackendRoleSecretIDMap) ToAuthBackendRoleSecretIDMapOutput() AuthBackendRoleSecretIDMapOutput {
	return i.ToAuthBackendRoleSecretIDMapOutputWithContext(context.Background())
}

func (i AuthBackendRoleSecretIDMap) ToAuthBackendRoleSecretIDMapOutputWithContext(ctx context.Context) AuthBackendRoleSecretIDMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendRoleSecretIDMapOutput)
}

type AuthBackendRoleSecretIDOutput struct {
	*pulumi.OutputState
}

func (AuthBackendRoleSecretIDOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthBackendRoleSecretID)(nil))
}

func (o AuthBackendRoleSecretIDOutput) ToAuthBackendRoleSecretIDOutput() AuthBackendRoleSecretIDOutput {
	return o
}

func (o AuthBackendRoleSecretIDOutput) ToAuthBackendRoleSecretIDOutputWithContext(ctx context.Context) AuthBackendRoleSecretIDOutput {
	return o
}

func (o AuthBackendRoleSecretIDOutput) ToAuthBackendRoleSecretIDPtrOutput() AuthBackendRoleSecretIDPtrOutput {
	return o.ToAuthBackendRoleSecretIDPtrOutputWithContext(context.Background())
}

func (o AuthBackendRoleSecretIDOutput) ToAuthBackendRoleSecretIDPtrOutputWithContext(ctx context.Context) AuthBackendRoleSecretIDPtrOutput {
	return o.ApplyT(func(v AuthBackendRoleSecretID) *AuthBackendRoleSecretID {
		return &v
	}).(AuthBackendRoleSecretIDPtrOutput)
}

type AuthBackendRoleSecretIDPtrOutput struct {
	*pulumi.OutputState
}

func (AuthBackendRoleSecretIDPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendRoleSecretID)(nil))
}

func (o AuthBackendRoleSecretIDPtrOutput) ToAuthBackendRoleSecretIDPtrOutput() AuthBackendRoleSecretIDPtrOutput {
	return o
}

func (o AuthBackendRoleSecretIDPtrOutput) ToAuthBackendRoleSecretIDPtrOutputWithContext(ctx context.Context) AuthBackendRoleSecretIDPtrOutput {
	return o
}

type AuthBackendRoleSecretIDArrayOutput struct{ *pulumi.OutputState }

func (AuthBackendRoleSecretIDArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AuthBackendRoleSecretID)(nil))
}

func (o AuthBackendRoleSecretIDArrayOutput) ToAuthBackendRoleSecretIDArrayOutput() AuthBackendRoleSecretIDArrayOutput {
	return o
}

func (o AuthBackendRoleSecretIDArrayOutput) ToAuthBackendRoleSecretIDArrayOutputWithContext(ctx context.Context) AuthBackendRoleSecretIDArrayOutput {
	return o
}

func (o AuthBackendRoleSecretIDArrayOutput) Index(i pulumi.IntInput) AuthBackendRoleSecretIDOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AuthBackendRoleSecretID {
		return vs[0].([]AuthBackendRoleSecretID)[vs[1].(int)]
	}).(AuthBackendRoleSecretIDOutput)
}

type AuthBackendRoleSecretIDMapOutput struct{ *pulumi.OutputState }

func (AuthBackendRoleSecretIDMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AuthBackendRoleSecretID)(nil))
}

func (o AuthBackendRoleSecretIDMapOutput) ToAuthBackendRoleSecretIDMapOutput() AuthBackendRoleSecretIDMapOutput {
	return o
}

func (o AuthBackendRoleSecretIDMapOutput) ToAuthBackendRoleSecretIDMapOutputWithContext(ctx context.Context) AuthBackendRoleSecretIDMapOutput {
	return o
}

func (o AuthBackendRoleSecretIDMapOutput) MapIndex(k pulumi.StringInput) AuthBackendRoleSecretIDOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AuthBackendRoleSecretID {
		return vs[0].(map[string]AuthBackendRoleSecretID)[vs[1].(string)]
	}).(AuthBackendRoleSecretIDOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthBackendRoleSecretIDOutput{})
	pulumi.RegisterOutputType(AuthBackendRoleSecretIDPtrOutput{})
	pulumi.RegisterOutputType(AuthBackendRoleSecretIDArrayOutput{})
	pulumi.RegisterOutputType(AuthBackendRoleSecretIDMapOutput{})
}
