// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Import
//
// AWS auth backend STS roles can be imported using `auth/`, the `backend` path, `/config/sts/`, and the `account_id` e.g.
//
// ```sh
//  $ pulumi import vault:aws/authBackendStsRole:AuthBackendStsRole example auth/aws/config/sts/1234567890
// ```
type AuthBackendStsRole struct {
	pulumi.CustomResourceState

	// The AWS account ID to configure the STS role for.
	AccountId pulumi.StringOutput `pulumi:"accountId"`
	// The path the AWS auth backend being configured was
	// mounted at.  Defaults to `aws`.
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// The STS role to assume when verifying requests made
	// by EC2 instances in the account specified by `accountId`.
	StsRole pulumi.StringOutput `pulumi:"stsRole"`
}

// NewAuthBackendStsRole registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendStsRole(ctx *pulumi.Context,
	name string, args *AuthBackendStsRoleArgs, opts ...pulumi.ResourceOption) (*AuthBackendStsRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountId == nil {
		return nil, errors.New("invalid value for required argument 'AccountId'")
	}
	if args.StsRole == nil {
		return nil, errors.New("invalid value for required argument 'StsRole'")
	}
	var resource AuthBackendStsRole
	err := ctx.RegisterResource("vault:aws/authBackendStsRole:AuthBackendStsRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackendStsRole gets an existing AuthBackendStsRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendStsRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendStsRoleState, opts ...pulumi.ResourceOption) (*AuthBackendStsRole, error) {
	var resource AuthBackendStsRole
	err := ctx.ReadResource("vault:aws/authBackendStsRole:AuthBackendStsRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendStsRole resources.
type authBackendStsRoleState struct {
	// The AWS account ID to configure the STS role for.
	AccountId *string `pulumi:"accountId"`
	// The path the AWS auth backend being configured was
	// mounted at.  Defaults to `aws`.
	Backend *string `pulumi:"backend"`
	// The STS role to assume when verifying requests made
	// by EC2 instances in the account specified by `accountId`.
	StsRole *string `pulumi:"stsRole"`
}

type AuthBackendStsRoleState struct {
	// The AWS account ID to configure the STS role for.
	AccountId pulumi.StringPtrInput
	// The path the AWS auth backend being configured was
	// mounted at.  Defaults to `aws`.
	Backend pulumi.StringPtrInput
	// The STS role to assume when verifying requests made
	// by EC2 instances in the account specified by `accountId`.
	StsRole pulumi.StringPtrInput
}

func (AuthBackendStsRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendStsRoleState)(nil)).Elem()
}

type authBackendStsRoleArgs struct {
	// The AWS account ID to configure the STS role for.
	AccountId string `pulumi:"accountId"`
	// The path the AWS auth backend being configured was
	// mounted at.  Defaults to `aws`.
	Backend *string `pulumi:"backend"`
	// The STS role to assume when verifying requests made
	// by EC2 instances in the account specified by `accountId`.
	StsRole string `pulumi:"stsRole"`
}

// The set of arguments for constructing a AuthBackendStsRole resource.
type AuthBackendStsRoleArgs struct {
	// The AWS account ID to configure the STS role for.
	AccountId pulumi.StringInput
	// The path the AWS auth backend being configured was
	// mounted at.  Defaults to `aws`.
	Backend pulumi.StringPtrInput
	// The STS role to assume when verifying requests made
	// by EC2 instances in the account specified by `accountId`.
	StsRole pulumi.StringInput
}

func (AuthBackendStsRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendStsRoleArgs)(nil)).Elem()
}

type AuthBackendStsRoleInput interface {
	pulumi.Input

	ToAuthBackendStsRoleOutput() AuthBackendStsRoleOutput
	ToAuthBackendStsRoleOutputWithContext(ctx context.Context) AuthBackendStsRoleOutput
}

func (*AuthBackendStsRole) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendStsRole)(nil)).Elem()
}

func (i *AuthBackendStsRole) ToAuthBackendStsRoleOutput() AuthBackendStsRoleOutput {
	return i.ToAuthBackendStsRoleOutputWithContext(context.Background())
}

func (i *AuthBackendStsRole) ToAuthBackendStsRoleOutputWithContext(ctx context.Context) AuthBackendStsRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendStsRoleOutput)
}

// AuthBackendStsRoleArrayInput is an input type that accepts AuthBackendStsRoleArray and AuthBackendStsRoleArrayOutput values.
// You can construct a concrete instance of `AuthBackendStsRoleArrayInput` via:
//
//          AuthBackendStsRoleArray{ AuthBackendStsRoleArgs{...} }
type AuthBackendStsRoleArrayInput interface {
	pulumi.Input

	ToAuthBackendStsRoleArrayOutput() AuthBackendStsRoleArrayOutput
	ToAuthBackendStsRoleArrayOutputWithContext(context.Context) AuthBackendStsRoleArrayOutput
}

type AuthBackendStsRoleArray []AuthBackendStsRoleInput

func (AuthBackendStsRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackendStsRole)(nil)).Elem()
}

func (i AuthBackendStsRoleArray) ToAuthBackendStsRoleArrayOutput() AuthBackendStsRoleArrayOutput {
	return i.ToAuthBackendStsRoleArrayOutputWithContext(context.Background())
}

func (i AuthBackendStsRoleArray) ToAuthBackendStsRoleArrayOutputWithContext(ctx context.Context) AuthBackendStsRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendStsRoleArrayOutput)
}

// AuthBackendStsRoleMapInput is an input type that accepts AuthBackendStsRoleMap and AuthBackendStsRoleMapOutput values.
// You can construct a concrete instance of `AuthBackendStsRoleMapInput` via:
//
//          AuthBackendStsRoleMap{ "key": AuthBackendStsRoleArgs{...} }
type AuthBackendStsRoleMapInput interface {
	pulumi.Input

	ToAuthBackendStsRoleMapOutput() AuthBackendStsRoleMapOutput
	ToAuthBackendStsRoleMapOutputWithContext(context.Context) AuthBackendStsRoleMapOutput
}

type AuthBackendStsRoleMap map[string]AuthBackendStsRoleInput

func (AuthBackendStsRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackendStsRole)(nil)).Elem()
}

func (i AuthBackendStsRoleMap) ToAuthBackendStsRoleMapOutput() AuthBackendStsRoleMapOutput {
	return i.ToAuthBackendStsRoleMapOutputWithContext(context.Background())
}

func (i AuthBackendStsRoleMap) ToAuthBackendStsRoleMapOutputWithContext(ctx context.Context) AuthBackendStsRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendStsRoleMapOutput)
}

type AuthBackendStsRoleOutput struct{ *pulumi.OutputState }

func (AuthBackendStsRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendStsRole)(nil)).Elem()
}

func (o AuthBackendStsRoleOutput) ToAuthBackendStsRoleOutput() AuthBackendStsRoleOutput {
	return o
}

func (o AuthBackendStsRoleOutput) ToAuthBackendStsRoleOutputWithContext(ctx context.Context) AuthBackendStsRoleOutput {
	return o
}

type AuthBackendStsRoleArrayOutput struct{ *pulumi.OutputState }

func (AuthBackendStsRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackendStsRole)(nil)).Elem()
}

func (o AuthBackendStsRoleArrayOutput) ToAuthBackendStsRoleArrayOutput() AuthBackendStsRoleArrayOutput {
	return o
}

func (o AuthBackendStsRoleArrayOutput) ToAuthBackendStsRoleArrayOutputWithContext(ctx context.Context) AuthBackendStsRoleArrayOutput {
	return o
}

func (o AuthBackendStsRoleArrayOutput) Index(i pulumi.IntInput) AuthBackendStsRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuthBackendStsRole {
		return vs[0].([]*AuthBackendStsRole)[vs[1].(int)]
	}).(AuthBackendStsRoleOutput)
}

type AuthBackendStsRoleMapOutput struct{ *pulumi.OutputState }

func (AuthBackendStsRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackendStsRole)(nil)).Elem()
}

func (o AuthBackendStsRoleMapOutput) ToAuthBackendStsRoleMapOutput() AuthBackendStsRoleMapOutput {
	return o
}

func (o AuthBackendStsRoleMapOutput) ToAuthBackendStsRoleMapOutputWithContext(ctx context.Context) AuthBackendStsRoleMapOutput {
	return o
}

func (o AuthBackendStsRoleMapOutput) MapIndex(k pulumi.StringInput) AuthBackendStsRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuthBackendStsRole {
		return vs[0].(map[string]*AuthBackendStsRole)[vs[1].(string)]
	}).(AuthBackendStsRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendStsRoleInput)(nil)).Elem(), &AuthBackendStsRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendStsRoleArrayInput)(nil)).Elem(), AuthBackendStsRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendStsRoleMapInput)(nil)).Elem(), AuthBackendStsRoleMap{})
	pulumi.RegisterOutputType(AuthBackendStsRoleOutput{})
	pulumi.RegisterOutputType(AuthBackendStsRoleArrayOutput{})
	pulumi.RegisterOutputType(AuthBackendStsRoleMapOutput{})
}
