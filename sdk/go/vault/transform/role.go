// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package transform

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This resource supports the "/transform/role/{name}" Vault endpoint.
//
// It creates or updates the role with the given name. If a role with the name does not exist, it will be created.
// If the role exists, it will be updated with the new attributes.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v2/go/vault"
// 	"github.com/pulumi/pulumi-vault/sdk/v2/go/vault/transform"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		mountTransform, err := vault.NewMount(ctx, "mountTransform", &vault.MountArgs{
// 			Path: pulumi.String("transform"),
// 			Type: pulumi.String("transform"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = transform.NewRole(ctx, "test", &transform.RoleArgs{
// 			Path: mountTransform.Path,
// 			Transformations: pulumi.StringArray{
// 				pulumi.String("ccn-fpe"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Role struct {
	pulumi.CustomResourceState

	// The name of the role.
	Name pulumi.StringOutput `pulumi:"name"`
	// Path to where the back-end is mounted within Vault.
	Path pulumi.StringOutput `pulumi:"path"`
	// A comma separated string or slice of transformations to use.
	Transformations pulumi.StringArrayOutput `pulumi:"transformations"`
}

// NewRole registers a new resource with the given unique name, arguments, and options.
func NewRole(ctx *pulumi.Context,
	name string, args *RoleArgs, opts ...pulumi.ResourceOption) (*Role, error) {
	if args == nil || args.Path == nil {
		return nil, errors.New("missing required argument 'Path'")
	}
	if args == nil {
		args = &RoleArgs{}
	}
	var resource Role
	err := ctx.RegisterResource("vault:transform/role:Role", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRole gets an existing Role resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoleState, opts ...pulumi.ResourceOption) (*Role, error) {
	var resource Role
	err := ctx.ReadResource("vault:transform/role:Role", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Role resources.
type roleState struct {
	// The name of the role.
	Name *string `pulumi:"name"`
	// Path to where the back-end is mounted within Vault.
	Path *string `pulumi:"path"`
	// A comma separated string or slice of transformations to use.
	Transformations []string `pulumi:"transformations"`
}

type RoleState struct {
	// The name of the role.
	Name pulumi.StringPtrInput
	// Path to where the back-end is mounted within Vault.
	Path pulumi.StringPtrInput
	// A comma separated string or slice of transformations to use.
	Transformations pulumi.StringArrayInput
}

func (RoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*roleState)(nil)).Elem()
}

type roleArgs struct {
	// The name of the role.
	Name *string `pulumi:"name"`
	// Path to where the back-end is mounted within Vault.
	Path string `pulumi:"path"`
	// A comma separated string or slice of transformations to use.
	Transformations []string `pulumi:"transformations"`
}

// The set of arguments for constructing a Role resource.
type RoleArgs struct {
	// The name of the role.
	Name pulumi.StringPtrInput
	// Path to where the back-end is mounted within Vault.
	Path pulumi.StringInput
	// A comma separated string or slice of transformations to use.
	Transformations pulumi.StringArrayInput
}

func (RoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*roleArgs)(nil)).Elem()
}
