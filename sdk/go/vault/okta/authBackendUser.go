// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package okta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to create a user in an
// [Okta auth backend within Vault](https://www.vaultproject.io/docs/auth/okta.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v3/go/vault/okta"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := okta.NewAuthBackend(ctx, "example", &okta.AuthBackendArgs{
// 			Organization: pulumi.String("dummy"),
// 			Path:         pulumi.String("user_okta"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = okta.NewAuthBackendUser(ctx, "foo", &okta.AuthBackendUserArgs{
// 			Groups: pulumi.StringArray{
// 				pulumi.String("one"),
// 				pulumi.String("two"),
// 			},
// 			Path:     example.Path,
// 			Username: pulumi.String("foo"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type AuthBackendUser struct {
	pulumi.CustomResourceState

	// List of Okta groups to associate with this user
	Groups pulumi.StringArrayOutput `pulumi:"groups"`
	// The path where the Okta auth backend is mounted
	Path pulumi.StringOutput `pulumi:"path"`
	// List of Vault policies to associate with this user
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
	// Name of the user within Okta
	Username pulumi.StringOutput `pulumi:"username"`
}

// NewAuthBackendUser registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendUser(ctx *pulumi.Context,
	name string, args *AuthBackendUserArgs, opts ...pulumi.ResourceOption) (*AuthBackendUser, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	var resource AuthBackendUser
	err := ctx.RegisterResource("vault:okta/authBackendUser:AuthBackendUser", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackendUser gets an existing AuthBackendUser resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendUserState, opts ...pulumi.ResourceOption) (*AuthBackendUser, error) {
	var resource AuthBackendUser
	err := ctx.ReadResource("vault:okta/authBackendUser:AuthBackendUser", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendUser resources.
type authBackendUserState struct {
	// List of Okta groups to associate with this user
	Groups []string `pulumi:"groups"`
	// The path where the Okta auth backend is mounted
	Path *string `pulumi:"path"`
	// List of Vault policies to associate with this user
	Policies []string `pulumi:"policies"`
	// Name of the user within Okta
	Username *string `pulumi:"username"`
}

type AuthBackendUserState struct {
	// List of Okta groups to associate with this user
	Groups pulumi.StringArrayInput
	// The path where the Okta auth backend is mounted
	Path pulumi.StringPtrInput
	// List of Vault policies to associate with this user
	Policies pulumi.StringArrayInput
	// Name of the user within Okta
	Username pulumi.StringPtrInput
}

func (AuthBackendUserState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendUserState)(nil)).Elem()
}

type authBackendUserArgs struct {
	// List of Okta groups to associate with this user
	Groups []string `pulumi:"groups"`
	// The path where the Okta auth backend is mounted
	Path string `pulumi:"path"`
	// List of Vault policies to associate with this user
	Policies []string `pulumi:"policies"`
	// Name of the user within Okta
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a AuthBackendUser resource.
type AuthBackendUserArgs struct {
	// List of Okta groups to associate with this user
	Groups pulumi.StringArrayInput
	// The path where the Okta auth backend is mounted
	Path pulumi.StringInput
	// List of Vault policies to associate with this user
	Policies pulumi.StringArrayInput
	// Name of the user within Okta
	Username pulumi.StringInput
}

func (AuthBackendUserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendUserArgs)(nil)).Elem()
}

type AuthBackendUserInput interface {
	pulumi.Input

	ToAuthBackendUserOutput() AuthBackendUserOutput
	ToAuthBackendUserOutputWithContext(ctx context.Context) AuthBackendUserOutput
}

func (*AuthBackendUser) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthBackendUser)(nil))
}

func (i *AuthBackendUser) ToAuthBackendUserOutput() AuthBackendUserOutput {
	return i.ToAuthBackendUserOutputWithContext(context.Background())
}

func (i *AuthBackendUser) ToAuthBackendUserOutputWithContext(ctx context.Context) AuthBackendUserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendUserOutput)
}

func (i *AuthBackendUser) ToAuthBackendUserPtrOutput() AuthBackendUserPtrOutput {
	return i.ToAuthBackendUserPtrOutputWithContext(context.Background())
}

func (i *AuthBackendUser) ToAuthBackendUserPtrOutputWithContext(ctx context.Context) AuthBackendUserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendUserPtrOutput)
}

type AuthBackendUserPtrInput interface {
	pulumi.Input

	ToAuthBackendUserPtrOutput() AuthBackendUserPtrOutput
	ToAuthBackendUserPtrOutputWithContext(ctx context.Context) AuthBackendUserPtrOutput
}

type authBackendUserPtrType AuthBackendUserArgs

func (*authBackendUserPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendUser)(nil))
}

func (i *authBackendUserPtrType) ToAuthBackendUserPtrOutput() AuthBackendUserPtrOutput {
	return i.ToAuthBackendUserPtrOutputWithContext(context.Background())
}

func (i *authBackendUserPtrType) ToAuthBackendUserPtrOutputWithContext(ctx context.Context) AuthBackendUserPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendUserPtrOutput)
}

// AuthBackendUserArrayInput is an input type that accepts AuthBackendUserArray and AuthBackendUserArrayOutput values.
// You can construct a concrete instance of `AuthBackendUserArrayInput` via:
//
//          AuthBackendUserArray{ AuthBackendUserArgs{...} }
type AuthBackendUserArrayInput interface {
	pulumi.Input

	ToAuthBackendUserArrayOutput() AuthBackendUserArrayOutput
	ToAuthBackendUserArrayOutputWithContext(context.Context) AuthBackendUserArrayOutput
}

type AuthBackendUserArray []AuthBackendUserInput

func (AuthBackendUserArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*AuthBackendUser)(nil))
}

func (i AuthBackendUserArray) ToAuthBackendUserArrayOutput() AuthBackendUserArrayOutput {
	return i.ToAuthBackendUserArrayOutputWithContext(context.Background())
}

func (i AuthBackendUserArray) ToAuthBackendUserArrayOutputWithContext(ctx context.Context) AuthBackendUserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendUserArrayOutput)
}

// AuthBackendUserMapInput is an input type that accepts AuthBackendUserMap and AuthBackendUserMapOutput values.
// You can construct a concrete instance of `AuthBackendUserMapInput` via:
//
//          AuthBackendUserMap{ "key": AuthBackendUserArgs{...} }
type AuthBackendUserMapInput interface {
	pulumi.Input

	ToAuthBackendUserMapOutput() AuthBackendUserMapOutput
	ToAuthBackendUserMapOutputWithContext(context.Context) AuthBackendUserMapOutput
}

type AuthBackendUserMap map[string]AuthBackendUserInput

func (AuthBackendUserMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*AuthBackendUser)(nil))
}

func (i AuthBackendUserMap) ToAuthBackendUserMapOutput() AuthBackendUserMapOutput {
	return i.ToAuthBackendUserMapOutputWithContext(context.Background())
}

func (i AuthBackendUserMap) ToAuthBackendUserMapOutputWithContext(ctx context.Context) AuthBackendUserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendUserMapOutput)
}

type AuthBackendUserOutput struct {
	*pulumi.OutputState
}

func (AuthBackendUserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthBackendUser)(nil))
}

func (o AuthBackendUserOutput) ToAuthBackendUserOutput() AuthBackendUserOutput {
	return o
}

func (o AuthBackendUserOutput) ToAuthBackendUserOutputWithContext(ctx context.Context) AuthBackendUserOutput {
	return o
}

func (o AuthBackendUserOutput) ToAuthBackendUserPtrOutput() AuthBackendUserPtrOutput {
	return o.ToAuthBackendUserPtrOutputWithContext(context.Background())
}

func (o AuthBackendUserOutput) ToAuthBackendUserPtrOutputWithContext(ctx context.Context) AuthBackendUserPtrOutput {
	return o.ApplyT(func(v AuthBackendUser) *AuthBackendUser {
		return &v
	}).(AuthBackendUserPtrOutput)
}

type AuthBackendUserPtrOutput struct {
	*pulumi.OutputState
}

func (AuthBackendUserPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendUser)(nil))
}

func (o AuthBackendUserPtrOutput) ToAuthBackendUserPtrOutput() AuthBackendUserPtrOutput {
	return o
}

func (o AuthBackendUserPtrOutput) ToAuthBackendUserPtrOutputWithContext(ctx context.Context) AuthBackendUserPtrOutput {
	return o
}

type AuthBackendUserArrayOutput struct{ *pulumi.OutputState }

func (AuthBackendUserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AuthBackendUser)(nil))
}

func (o AuthBackendUserArrayOutput) ToAuthBackendUserArrayOutput() AuthBackendUserArrayOutput {
	return o
}

func (o AuthBackendUserArrayOutput) ToAuthBackendUserArrayOutputWithContext(ctx context.Context) AuthBackendUserArrayOutput {
	return o
}

func (o AuthBackendUserArrayOutput) Index(i pulumi.IntInput) AuthBackendUserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AuthBackendUser {
		return vs[0].([]AuthBackendUser)[vs[1].(int)]
	}).(AuthBackendUserOutput)
}

type AuthBackendUserMapOutput struct{ *pulumi.OutputState }

func (AuthBackendUserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AuthBackendUser)(nil))
}

func (o AuthBackendUserMapOutput) ToAuthBackendUserMapOutput() AuthBackendUserMapOutput {
	return o
}

func (o AuthBackendUserMapOutput) ToAuthBackendUserMapOutputWithContext(ctx context.Context) AuthBackendUserMapOutput {
	return o
}

func (o AuthBackendUserMapOutput) MapIndex(k pulumi.StringInput) AuthBackendUserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AuthBackendUser {
		return vs[0].(map[string]AuthBackendUser)[vs[1].(string)]
	}).(AuthBackendUserOutput)
}

func init() {
	pulumi.RegisterOutputType(AuthBackendUserOutput{})
	pulumi.RegisterOutputType(AuthBackendUserPtrOutput{})
	pulumi.RegisterOutputType(AuthBackendUserArrayOutput{})
	pulumi.RegisterOutputType(AuthBackendUserMapOutput{})
}
