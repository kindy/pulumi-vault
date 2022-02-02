// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := vault.NewMount(ctx, "example", &vault.MountArgs{
// 			Description: pulumi.String("This is an example mount"),
// 			Path:        pulumi.String("dummy"),
// 			Type:        pulumi.String("generic"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := vault.NewMount(ctx, "kvv2-example", &vault.MountArgs{
// 			Description: pulumi.String("This is an example KV Version 2 secret engine mount"),
// 			Path:        pulumi.String("version2-example"),
// 			Type:        pulumi.String("kv-v2"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := vault.NewMount(ctx, "transit-example", &vault.MountArgs{
// 			Description: pulumi.String("This is an example transit secret engine mount"),
// 			Options: pulumi.AnyMap{
// 				"convergent_encryption": pulumi.Any(false),
// 			},
// 			Path: pulumi.String("transit-example"),
// 			Type: pulumi.String("transit"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := vault.NewMount(ctx, "pki-example", &vault.MountArgs{
// 			DefaultLeaseTtlSeconds: pulumi.Int(3600),
// 			Description:            pulumi.String("This is an example PKI mount"),
// 			MaxLeaseTtlSeconds:     pulumi.Int(86400),
// 			Path:                   pulumi.String("pki-example"),
// 			Type:                   pulumi.String("pki"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Mounts can be imported using the `path`, e.g.
//
// ```sh
//  $ pulumi import vault:index/mount:Mount example dummy
// ```
type Mount struct {
	pulumi.CustomResourceState

	// The accessor for this mount.
	Accessor pulumi.StringOutput `pulumi:"accessor"`
	// Default lease duration for tokens and secrets in seconds
	DefaultLeaseTtlSeconds pulumi.IntOutput `pulumi:"defaultLeaseTtlSeconds"`
	// Human-friendly description of the mount
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Boolean flag that can be explicitly set to true to enable the secrets engine to access Vault's external entropy source
	ExternalEntropyAccess pulumi.BoolPtrOutput `pulumi:"externalEntropyAccess"`
	// Boolean flag that can be explicitly set to true to enforce local mount in HA environment
	Local pulumi.BoolPtrOutput `pulumi:"local"`
	// Maximum possible lease duration for tokens and secrets in seconds
	MaxLeaseTtlSeconds pulumi.IntOutput `pulumi:"maxLeaseTtlSeconds"`
	// Specifies mount type specific options that are passed to the backend
	Options pulumi.MapOutput `pulumi:"options"`
	// Where the secret backend will be mounted
	Path pulumi.StringOutput `pulumi:"path"`
	// Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
	SealWrap pulumi.BoolOutput `pulumi:"sealWrap"`
	// Type of the backend, such as "aws"
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewMount registers a new resource with the given unique name, arguments, and options.
func NewMount(ctx *pulumi.Context,
	name string, args *MountArgs, opts ...pulumi.ResourceOption) (*Mount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	var resource Mount
	err := ctx.RegisterResource("vault:index/mount:Mount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMount gets an existing Mount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MountState, opts ...pulumi.ResourceOption) (*Mount, error) {
	var resource Mount
	err := ctx.ReadResource("vault:index/mount:Mount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Mount resources.
type mountState struct {
	// The accessor for this mount.
	Accessor *string `pulumi:"accessor"`
	// Default lease duration for tokens and secrets in seconds
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// Human-friendly description of the mount
	Description *string `pulumi:"description"`
	// Boolean flag that can be explicitly set to true to enable the secrets engine to access Vault's external entropy source
	ExternalEntropyAccess *bool `pulumi:"externalEntropyAccess"`
	// Boolean flag that can be explicitly set to true to enforce local mount in HA environment
	Local *bool `pulumi:"local"`
	// Maximum possible lease duration for tokens and secrets in seconds
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// Specifies mount type specific options that are passed to the backend
	Options map[string]interface{} `pulumi:"options"`
	// Where the secret backend will be mounted
	Path *string `pulumi:"path"`
	// Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
	SealWrap *bool `pulumi:"sealWrap"`
	// Type of the backend, such as "aws"
	Type *string `pulumi:"type"`
}

type MountState struct {
	// The accessor for this mount.
	Accessor pulumi.StringPtrInput
	// Default lease duration for tokens and secrets in seconds
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// Human-friendly description of the mount
	Description pulumi.StringPtrInput
	// Boolean flag that can be explicitly set to true to enable the secrets engine to access Vault's external entropy source
	ExternalEntropyAccess pulumi.BoolPtrInput
	// Boolean flag that can be explicitly set to true to enforce local mount in HA environment
	Local pulumi.BoolPtrInput
	// Maximum possible lease duration for tokens and secrets in seconds
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// Specifies mount type specific options that are passed to the backend
	Options pulumi.MapInput
	// Where the secret backend will be mounted
	Path pulumi.StringPtrInput
	// Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
	SealWrap pulumi.BoolPtrInput
	// Type of the backend, such as "aws"
	Type pulumi.StringPtrInput
}

func (MountState) ElementType() reflect.Type {
	return reflect.TypeOf((*mountState)(nil)).Elem()
}

type mountArgs struct {
	// Default lease duration for tokens and secrets in seconds
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// Human-friendly description of the mount
	Description *string `pulumi:"description"`
	// Boolean flag that can be explicitly set to true to enable the secrets engine to access Vault's external entropy source
	ExternalEntropyAccess *bool `pulumi:"externalEntropyAccess"`
	// Boolean flag that can be explicitly set to true to enforce local mount in HA environment
	Local *bool `pulumi:"local"`
	// Maximum possible lease duration for tokens and secrets in seconds
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// Specifies mount type specific options that are passed to the backend
	Options map[string]interface{} `pulumi:"options"`
	// Where the secret backend will be mounted
	Path string `pulumi:"path"`
	// Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
	SealWrap *bool `pulumi:"sealWrap"`
	// Type of the backend, such as "aws"
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Mount resource.
type MountArgs struct {
	// Default lease duration for tokens and secrets in seconds
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// Human-friendly description of the mount
	Description pulumi.StringPtrInput
	// Boolean flag that can be explicitly set to true to enable the secrets engine to access Vault's external entropy source
	ExternalEntropyAccess pulumi.BoolPtrInput
	// Boolean flag that can be explicitly set to true to enforce local mount in HA environment
	Local pulumi.BoolPtrInput
	// Maximum possible lease duration for tokens and secrets in seconds
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// Specifies mount type specific options that are passed to the backend
	Options pulumi.MapInput
	// Where the secret backend will be mounted
	Path pulumi.StringInput
	// Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
	SealWrap pulumi.BoolPtrInput
	// Type of the backend, such as "aws"
	Type pulumi.StringInput
}

func (MountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mountArgs)(nil)).Elem()
}

type MountInput interface {
	pulumi.Input

	ToMountOutput() MountOutput
	ToMountOutputWithContext(ctx context.Context) MountOutput
}

func (*Mount) ElementType() reflect.Type {
	return reflect.TypeOf((**Mount)(nil)).Elem()
}

func (i *Mount) ToMountOutput() MountOutput {
	return i.ToMountOutputWithContext(context.Background())
}

func (i *Mount) ToMountOutputWithContext(ctx context.Context) MountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountOutput)
}

// MountArrayInput is an input type that accepts MountArray and MountArrayOutput values.
// You can construct a concrete instance of `MountArrayInput` via:
//
//          MountArray{ MountArgs{...} }
type MountArrayInput interface {
	pulumi.Input

	ToMountArrayOutput() MountArrayOutput
	ToMountArrayOutputWithContext(context.Context) MountArrayOutput
}

type MountArray []MountInput

func (MountArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Mount)(nil)).Elem()
}

func (i MountArray) ToMountArrayOutput() MountArrayOutput {
	return i.ToMountArrayOutputWithContext(context.Background())
}

func (i MountArray) ToMountArrayOutputWithContext(ctx context.Context) MountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountArrayOutput)
}

// MountMapInput is an input type that accepts MountMap and MountMapOutput values.
// You can construct a concrete instance of `MountMapInput` via:
//
//          MountMap{ "key": MountArgs{...} }
type MountMapInput interface {
	pulumi.Input

	ToMountMapOutput() MountMapOutput
	ToMountMapOutputWithContext(context.Context) MountMapOutput
}

type MountMap map[string]MountInput

func (MountMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Mount)(nil)).Elem()
}

func (i MountMap) ToMountMapOutput() MountMapOutput {
	return i.ToMountMapOutputWithContext(context.Background())
}

func (i MountMap) ToMountMapOutputWithContext(ctx context.Context) MountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MountMapOutput)
}

type MountOutput struct{ *pulumi.OutputState }

func (MountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Mount)(nil)).Elem()
}

func (o MountOutput) ToMountOutput() MountOutput {
	return o
}

func (o MountOutput) ToMountOutputWithContext(ctx context.Context) MountOutput {
	return o
}

type MountArrayOutput struct{ *pulumi.OutputState }

func (MountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Mount)(nil)).Elem()
}

func (o MountArrayOutput) ToMountArrayOutput() MountArrayOutput {
	return o
}

func (o MountArrayOutput) ToMountArrayOutputWithContext(ctx context.Context) MountArrayOutput {
	return o
}

func (o MountArrayOutput) Index(i pulumi.IntInput) MountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Mount {
		return vs[0].([]*Mount)[vs[1].(int)]
	}).(MountOutput)
}

type MountMapOutput struct{ *pulumi.OutputState }

func (MountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Mount)(nil)).Elem()
}

func (o MountMapOutput) ToMountMapOutput() MountMapOutput {
	return o
}

func (o MountMapOutput) ToMountMapOutputWithContext(ctx context.Context) MountMapOutput {
	return o
}

func (o MountMapOutput) MapIndex(k pulumi.StringInput) MountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Mount {
		return vs[0].(map[string]*Mount)[vs[1].(string)]
	}).(MountOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MountInput)(nil)).Elem(), &Mount{})
	pulumi.RegisterInputType(reflect.TypeOf((*MountArrayInput)(nil)).Elem(), MountArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MountMapInput)(nil)).Elem(), MountMap{})
	pulumi.RegisterOutputType(MountOutput{})
	pulumi.RegisterOutputType(MountArrayOutput{})
	pulumi.RegisterOutputType(MountMapOutput{})
}
