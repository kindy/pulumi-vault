// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kv

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configures KV-V2 backend level settings that are applied to
// every key in the key-value store.
//
// For more information on Vault's KV-V2 secret backend
// [see here](https://www.vaultproject.io/docs/secrets/kv/kv-v2).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/kv"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			kvv2, err := vault.NewMount(ctx, "kvv2", &vault.MountArgs{
//				Path: pulumi.String("kvv2"),
//				Type: pulumi.String("kv"),
//				Options: pulumi.AnyMap{
//					"version": pulumi.Any("2"),
//				},
//				Description: pulumi.String("KV Version 2 secret engine mount"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = kv.NewSecretBackendV2(ctx, "config", &kv.SecretBackendV2Args{
//				Mount:              kvv2.Path,
//				MaxVersions:        pulumi.Int(5),
//				DeleteVersionAfter: pulumi.Int(12600),
//				CasRequired:        pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ## Required Vault Capabilities
//
// Use of this resource requires the `create` or `update` capability
// (depending on whether the resource already exists) on the given path,
// the `delete` capability if the resource is removed from configuration,
// and the `read` capability for drift detection (by default).
//
// ## Import
//
// The KV-V2 secret backend can be imported using the `path`, e.g.
//
// ```sh
//
//	$ pulumi import vault:kv/secretBackendV2:SecretBackendV2 config kvv2/config
//
// ```
type SecretBackendV2 struct {
	pulumi.CustomResourceState

	// If true, all keys will require the cas
	// parameter to be set on all write requests.
	CasRequired pulumi.BoolOutput `pulumi:"casRequired"`
	// If set, specifies the length of time before
	// a version is deleted. Accepts duration in integer seconds.
	DeleteVersionAfter pulumi.IntPtrOutput `pulumi:"deleteVersionAfter"`
	// The number of versions to keep per key.
	MaxVersions pulumi.IntOutput `pulumi:"maxVersions"`
	// Path where KV-V2 engine is mounted.
	Mount pulumi.StringOutput `pulumi:"mount"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
}

// NewSecretBackendV2 registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendV2(ctx *pulumi.Context,
	name string, args *SecretBackendV2Args, opts ...pulumi.ResourceOption) (*SecretBackendV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Mount == nil {
		return nil, errors.New("invalid value for required argument 'Mount'")
	}
	var resource SecretBackendV2
	err := ctx.RegisterResource("vault:kv/secretBackendV2:SecretBackendV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendV2 gets an existing SecretBackendV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendV2State, opts ...pulumi.ResourceOption) (*SecretBackendV2, error) {
	var resource SecretBackendV2
	err := ctx.ReadResource("vault:kv/secretBackendV2:SecretBackendV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendV2 resources.
type secretBackendV2State struct {
	// If true, all keys will require the cas
	// parameter to be set on all write requests.
	CasRequired *bool `pulumi:"casRequired"`
	// If set, specifies the length of time before
	// a version is deleted. Accepts duration in integer seconds.
	DeleteVersionAfter *int `pulumi:"deleteVersionAfter"`
	// The number of versions to keep per key.
	MaxVersions *int `pulumi:"maxVersions"`
	// Path where KV-V2 engine is mounted.
	Mount *string `pulumi:"mount"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

type SecretBackendV2State struct {
	// If true, all keys will require the cas
	// parameter to be set on all write requests.
	CasRequired pulumi.BoolPtrInput
	// If set, specifies the length of time before
	// a version is deleted. Accepts duration in integer seconds.
	DeleteVersionAfter pulumi.IntPtrInput
	// The number of versions to keep per key.
	MaxVersions pulumi.IntPtrInput
	// Path where KV-V2 engine is mounted.
	Mount pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
}

func (SecretBackendV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendV2State)(nil)).Elem()
}

type secretBackendV2Args struct {
	// If true, all keys will require the cas
	// parameter to be set on all write requests.
	CasRequired *bool `pulumi:"casRequired"`
	// If set, specifies the length of time before
	// a version is deleted. Accepts duration in integer seconds.
	DeleteVersionAfter *int `pulumi:"deleteVersionAfter"`
	// The number of versions to keep per key.
	MaxVersions *int `pulumi:"maxVersions"`
	// Path where KV-V2 engine is mounted.
	Mount string `pulumi:"mount"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
}

// The set of arguments for constructing a SecretBackendV2 resource.
type SecretBackendV2Args struct {
	// If true, all keys will require the cas
	// parameter to be set on all write requests.
	CasRequired pulumi.BoolPtrInput
	// If set, specifies the length of time before
	// a version is deleted. Accepts duration in integer seconds.
	DeleteVersionAfter pulumi.IntPtrInput
	// The number of versions to keep per key.
	MaxVersions pulumi.IntPtrInput
	// Path where KV-V2 engine is mounted.
	Mount pulumi.StringInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
}

func (SecretBackendV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendV2Args)(nil)).Elem()
}

type SecretBackendV2Input interface {
	pulumi.Input

	ToSecretBackendV2Output() SecretBackendV2Output
	ToSecretBackendV2OutputWithContext(ctx context.Context) SecretBackendV2Output
}

func (*SecretBackendV2) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendV2)(nil)).Elem()
}

func (i *SecretBackendV2) ToSecretBackendV2Output() SecretBackendV2Output {
	return i.ToSecretBackendV2OutputWithContext(context.Background())
}

func (i *SecretBackendV2) ToSecretBackendV2OutputWithContext(ctx context.Context) SecretBackendV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendV2Output)
}

// SecretBackendV2ArrayInput is an input type that accepts SecretBackendV2Array and SecretBackendV2ArrayOutput values.
// You can construct a concrete instance of `SecretBackendV2ArrayInput` via:
//
//	SecretBackendV2Array{ SecretBackendV2Args{...} }
type SecretBackendV2ArrayInput interface {
	pulumi.Input

	ToSecretBackendV2ArrayOutput() SecretBackendV2ArrayOutput
	ToSecretBackendV2ArrayOutputWithContext(context.Context) SecretBackendV2ArrayOutput
}

type SecretBackendV2Array []SecretBackendV2Input

func (SecretBackendV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendV2)(nil)).Elem()
}

func (i SecretBackendV2Array) ToSecretBackendV2ArrayOutput() SecretBackendV2ArrayOutput {
	return i.ToSecretBackendV2ArrayOutputWithContext(context.Background())
}

func (i SecretBackendV2Array) ToSecretBackendV2ArrayOutputWithContext(ctx context.Context) SecretBackendV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendV2ArrayOutput)
}

// SecretBackendV2MapInput is an input type that accepts SecretBackendV2Map and SecretBackendV2MapOutput values.
// You can construct a concrete instance of `SecretBackendV2MapInput` via:
//
//	SecretBackendV2Map{ "key": SecretBackendV2Args{...} }
type SecretBackendV2MapInput interface {
	pulumi.Input

	ToSecretBackendV2MapOutput() SecretBackendV2MapOutput
	ToSecretBackendV2MapOutputWithContext(context.Context) SecretBackendV2MapOutput
}

type SecretBackendV2Map map[string]SecretBackendV2Input

func (SecretBackendV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendV2)(nil)).Elem()
}

func (i SecretBackendV2Map) ToSecretBackendV2MapOutput() SecretBackendV2MapOutput {
	return i.ToSecretBackendV2MapOutputWithContext(context.Background())
}

func (i SecretBackendV2Map) ToSecretBackendV2MapOutputWithContext(ctx context.Context) SecretBackendV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendV2MapOutput)
}

type SecretBackendV2Output struct{ *pulumi.OutputState }

func (SecretBackendV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendV2)(nil)).Elem()
}

func (o SecretBackendV2Output) ToSecretBackendV2Output() SecretBackendV2Output {
	return o
}

func (o SecretBackendV2Output) ToSecretBackendV2OutputWithContext(ctx context.Context) SecretBackendV2Output {
	return o
}

// If true, all keys will require the cas
// parameter to be set on all write requests.
func (o SecretBackendV2Output) CasRequired() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretBackendV2) pulumi.BoolOutput { return v.CasRequired }).(pulumi.BoolOutput)
}

// If set, specifies the length of time before
// a version is deleted. Accepts duration in integer seconds.
func (o SecretBackendV2Output) DeleteVersionAfter() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretBackendV2) pulumi.IntPtrOutput { return v.DeleteVersionAfter }).(pulumi.IntPtrOutput)
}

// The number of versions to keep per key.
func (o SecretBackendV2Output) MaxVersions() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretBackendV2) pulumi.IntOutput { return v.MaxVersions }).(pulumi.IntOutput)
}

// Path where KV-V2 engine is mounted.
func (o SecretBackendV2Output) Mount() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendV2) pulumi.StringOutput { return v.Mount }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o SecretBackendV2Output) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendV2) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

type SecretBackendV2ArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendV2)(nil)).Elem()
}

func (o SecretBackendV2ArrayOutput) ToSecretBackendV2ArrayOutput() SecretBackendV2ArrayOutput {
	return o
}

func (o SecretBackendV2ArrayOutput) ToSecretBackendV2ArrayOutputWithContext(ctx context.Context) SecretBackendV2ArrayOutput {
	return o
}

func (o SecretBackendV2ArrayOutput) Index(i pulumi.IntInput) SecretBackendV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackendV2 {
		return vs[0].([]*SecretBackendV2)[vs[1].(int)]
	}).(SecretBackendV2Output)
}

type SecretBackendV2MapOutput struct{ *pulumi.OutputState }

func (SecretBackendV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendV2)(nil)).Elem()
}

func (o SecretBackendV2MapOutput) ToSecretBackendV2MapOutput() SecretBackendV2MapOutput {
	return o
}

func (o SecretBackendV2MapOutput) ToSecretBackendV2MapOutputWithContext(ctx context.Context) SecretBackendV2MapOutput {
	return o
}

func (o SecretBackendV2MapOutput) MapIndex(k pulumi.StringInput) SecretBackendV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackendV2 {
		return vs[0].(map[string]*SecretBackendV2)[vs[1].(string)]
	}).(SecretBackendV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendV2Input)(nil)).Elem(), &SecretBackendV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendV2ArrayInput)(nil)).Elem(), SecretBackendV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendV2MapInput)(nil)).Elem(), SecretBackendV2Map{})
	pulumi.RegisterOutputType(SecretBackendV2Output{})
	pulumi.RegisterOutputType(SecretBackendV2ArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendV2MapOutput{})
}