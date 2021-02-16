// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Manages policies for an Identity Group for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.
//
// ## Example Usage
// ### Exclusive Policies
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v3/go/vault/identity"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		internal, err := identity.NewGroup(ctx, "internal", &identity.GroupArgs{
// 			Type:             pulumi.String("internal"),
// 			ExternalPolicies: pulumi.Bool(true),
// 			Metadata: pulumi.StringMap{
// 				"version": pulumi.String("2"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = identity.NewGroupPolicies(ctx, "policies", &identity.GroupPoliciesArgs{
// 			Policies: pulumi.StringArray{
// 				pulumi.String("default"),
// 				pulumi.String("test"),
// 			},
// 			Exclusive: pulumi.Bool(true),
// 			GroupId:   internal.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Non-exclusive Policies
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v3/go/vault/identity"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		internal, err := identity.NewGroup(ctx, "internal", &identity.GroupArgs{
// 			Type:             pulumi.String("internal"),
// 			ExternalPolicies: pulumi.Bool(true),
// 			Metadata: pulumi.StringMap{
// 				"version": pulumi.String("2"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = identity.NewGroupPolicies(ctx, "_default", &identity.GroupPoliciesArgs{
// 			Policies: pulumi.StringArray{
// 				pulumi.String("default"),
// 				pulumi.String("test"),
// 			},
// 			Exclusive: pulumi.Bool(false),
// 			GroupId:   internal.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = identity.NewGroupPolicies(ctx, "others", &identity.GroupPoliciesArgs{
// 			Policies: pulumi.StringArray{
// 				pulumi.String("others"),
// 			},
// 			Exclusive: pulumi.Bool(false),
// 			GroupId:   internal.ID(),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type GroupPolicies struct {
	pulumi.CustomResourceState

	// Defaults to `true`.
	Exclusive pulumi.BoolPtrOutput `pulumi:"exclusive"`
	// Group ID to assign policies to.
	GroupId pulumi.StringOutput `pulumi:"groupId"`
	// The name of the group that are assigned the policies.
	GroupName pulumi.StringOutput `pulumi:"groupName"`
	// List of policies to assign to the group
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
}

// NewGroupPolicies registers a new resource with the given unique name, arguments, and options.
func NewGroupPolicies(ctx *pulumi.Context,
	name string, args *GroupPoliciesArgs, opts ...pulumi.ResourceOption) (*GroupPolicies, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupId == nil {
		return nil, errors.New("invalid value for required argument 'GroupId'")
	}
	if args.Policies == nil {
		return nil, errors.New("invalid value for required argument 'Policies'")
	}
	var resource GroupPolicies
	err := ctx.RegisterResource("vault:identity/groupPolicies:GroupPolicies", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGroupPolicies gets an existing GroupPolicies resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGroupPolicies(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GroupPoliciesState, opts ...pulumi.ResourceOption) (*GroupPolicies, error) {
	var resource GroupPolicies
	err := ctx.ReadResource("vault:identity/groupPolicies:GroupPolicies", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GroupPolicies resources.
type groupPoliciesState struct {
	// Defaults to `true`.
	Exclusive *bool `pulumi:"exclusive"`
	// Group ID to assign policies to.
	GroupId *string `pulumi:"groupId"`
	// The name of the group that are assigned the policies.
	GroupName *string `pulumi:"groupName"`
	// List of policies to assign to the group
	Policies []string `pulumi:"policies"`
}

type GroupPoliciesState struct {
	// Defaults to `true`.
	Exclusive pulumi.BoolPtrInput
	// Group ID to assign policies to.
	GroupId pulumi.StringPtrInput
	// The name of the group that are assigned the policies.
	GroupName pulumi.StringPtrInput
	// List of policies to assign to the group
	Policies pulumi.StringArrayInput
}

func (GroupPoliciesState) ElementType() reflect.Type {
	return reflect.TypeOf((*groupPoliciesState)(nil)).Elem()
}

type groupPoliciesArgs struct {
	// Defaults to `true`.
	Exclusive *bool `pulumi:"exclusive"`
	// Group ID to assign policies to.
	GroupId string `pulumi:"groupId"`
	// List of policies to assign to the group
	Policies []string `pulumi:"policies"`
}

// The set of arguments for constructing a GroupPolicies resource.
type GroupPoliciesArgs struct {
	// Defaults to `true`.
	Exclusive pulumi.BoolPtrInput
	// Group ID to assign policies to.
	GroupId pulumi.StringInput
	// List of policies to assign to the group
	Policies pulumi.StringArrayInput
}

func (GroupPoliciesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*groupPoliciesArgs)(nil)).Elem()
}

type GroupPoliciesInput interface {
	pulumi.Input

	ToGroupPoliciesOutput() GroupPoliciesOutput
	ToGroupPoliciesOutputWithContext(ctx context.Context) GroupPoliciesOutput
}

func (*GroupPolicies) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupPolicies)(nil))
}

func (i *GroupPolicies) ToGroupPoliciesOutput() GroupPoliciesOutput {
	return i.ToGroupPoliciesOutputWithContext(context.Background())
}

func (i *GroupPolicies) ToGroupPoliciesOutputWithContext(ctx context.Context) GroupPoliciesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPoliciesOutput)
}

func (i *GroupPolicies) ToGroupPoliciesPtrOutput() GroupPoliciesPtrOutput {
	return i.ToGroupPoliciesPtrOutputWithContext(context.Background())
}

func (i *GroupPolicies) ToGroupPoliciesPtrOutputWithContext(ctx context.Context) GroupPoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPoliciesPtrOutput)
}

type GroupPoliciesPtrInput interface {
	pulumi.Input

	ToGroupPoliciesPtrOutput() GroupPoliciesPtrOutput
	ToGroupPoliciesPtrOutputWithContext(ctx context.Context) GroupPoliciesPtrOutput
}

type groupPoliciesPtrType GroupPoliciesArgs

func (*groupPoliciesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupPolicies)(nil))
}

func (i *groupPoliciesPtrType) ToGroupPoliciesPtrOutput() GroupPoliciesPtrOutput {
	return i.ToGroupPoliciesPtrOutputWithContext(context.Background())
}

func (i *groupPoliciesPtrType) ToGroupPoliciesPtrOutputWithContext(ctx context.Context) GroupPoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPoliciesPtrOutput)
}

// GroupPoliciesArrayInput is an input type that accepts GroupPoliciesArray and GroupPoliciesArrayOutput values.
// You can construct a concrete instance of `GroupPoliciesArrayInput` via:
//
//          GroupPoliciesArray{ GroupPoliciesArgs{...} }
type GroupPoliciesArrayInput interface {
	pulumi.Input

	ToGroupPoliciesArrayOutput() GroupPoliciesArrayOutput
	ToGroupPoliciesArrayOutputWithContext(context.Context) GroupPoliciesArrayOutput
}

type GroupPoliciesArray []GroupPoliciesInput

func (GroupPoliciesArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*GroupPolicies)(nil))
}

func (i GroupPoliciesArray) ToGroupPoliciesArrayOutput() GroupPoliciesArrayOutput {
	return i.ToGroupPoliciesArrayOutputWithContext(context.Background())
}

func (i GroupPoliciesArray) ToGroupPoliciesArrayOutputWithContext(ctx context.Context) GroupPoliciesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPoliciesArrayOutput)
}

// GroupPoliciesMapInput is an input type that accepts GroupPoliciesMap and GroupPoliciesMapOutput values.
// You can construct a concrete instance of `GroupPoliciesMapInput` via:
//
//          GroupPoliciesMap{ "key": GroupPoliciesArgs{...} }
type GroupPoliciesMapInput interface {
	pulumi.Input

	ToGroupPoliciesMapOutput() GroupPoliciesMapOutput
	ToGroupPoliciesMapOutputWithContext(context.Context) GroupPoliciesMapOutput
}

type GroupPoliciesMap map[string]GroupPoliciesInput

func (GroupPoliciesMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*GroupPolicies)(nil))
}

func (i GroupPoliciesMap) ToGroupPoliciesMapOutput() GroupPoliciesMapOutput {
	return i.ToGroupPoliciesMapOutputWithContext(context.Background())
}

func (i GroupPoliciesMap) ToGroupPoliciesMapOutputWithContext(ctx context.Context) GroupPoliciesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupPoliciesMapOutput)
}

type GroupPoliciesOutput struct {
	*pulumi.OutputState
}

func (GroupPoliciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupPolicies)(nil))
}

func (o GroupPoliciesOutput) ToGroupPoliciesOutput() GroupPoliciesOutput {
	return o
}

func (o GroupPoliciesOutput) ToGroupPoliciesOutputWithContext(ctx context.Context) GroupPoliciesOutput {
	return o
}

func (o GroupPoliciesOutput) ToGroupPoliciesPtrOutput() GroupPoliciesPtrOutput {
	return o.ToGroupPoliciesPtrOutputWithContext(context.Background())
}

func (o GroupPoliciesOutput) ToGroupPoliciesPtrOutputWithContext(ctx context.Context) GroupPoliciesPtrOutput {
	return o.ApplyT(func(v GroupPolicies) *GroupPolicies {
		return &v
	}).(GroupPoliciesPtrOutput)
}

type GroupPoliciesPtrOutput struct {
	*pulumi.OutputState
}

func (GroupPoliciesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupPolicies)(nil))
}

func (o GroupPoliciesPtrOutput) ToGroupPoliciesPtrOutput() GroupPoliciesPtrOutput {
	return o
}

func (o GroupPoliciesPtrOutput) ToGroupPoliciesPtrOutputWithContext(ctx context.Context) GroupPoliciesPtrOutput {
	return o
}

type GroupPoliciesArrayOutput struct{ *pulumi.OutputState }

func (GroupPoliciesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GroupPolicies)(nil))
}

func (o GroupPoliciesArrayOutput) ToGroupPoliciesArrayOutput() GroupPoliciesArrayOutput {
	return o
}

func (o GroupPoliciesArrayOutput) ToGroupPoliciesArrayOutputWithContext(ctx context.Context) GroupPoliciesArrayOutput {
	return o
}

func (o GroupPoliciesArrayOutput) Index(i pulumi.IntInput) GroupPoliciesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GroupPolicies {
		return vs[0].([]GroupPolicies)[vs[1].(int)]
	}).(GroupPoliciesOutput)
}

type GroupPoliciesMapOutput struct{ *pulumi.OutputState }

func (GroupPoliciesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]GroupPolicies)(nil))
}

func (o GroupPoliciesMapOutput) ToGroupPoliciesMapOutput() GroupPoliciesMapOutput {
	return o
}

func (o GroupPoliciesMapOutput) ToGroupPoliciesMapOutputWithContext(ctx context.Context) GroupPoliciesMapOutput {
	return o
}

func (o GroupPoliciesMapOutput) MapIndex(k pulumi.StringInput) GroupPoliciesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) GroupPolicies {
		return vs[0].(map[string]GroupPolicies)[vs[1].(string)]
	}).(GroupPoliciesOutput)
}

func init() {
	pulumi.RegisterOutputType(GroupPoliciesOutput{})
	pulumi.RegisterOutputType(GroupPoliciesPtrOutput{})
	pulumi.RegisterOutputType(GroupPoliciesArrayOutput{})
	pulumi.RegisterOutputType(GroupPoliciesMapOutput{})
}
