// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package secrets

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configures the secret sync global config.
// The config is global and can only be managed in the root namespace.
//
// > **Important** The config is global so the secrets.SyncConfig resource must not be defined
// multiple times for the same Vault server. If multiple definition exists, the last one applied will be
// effective.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/secrets"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := secrets.NewSyncConfig(ctx, "globalConfig", &secrets.SyncConfigArgs{
//				Disabled:      pulumi.Bool(true),
//				QueueCapacity: pulumi.Int(500000),
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
// ```sh
//
//	$ pulumi import vault:secrets/syncConfig:SyncConfig config global_config
//
// ```
type SyncConfig struct {
	pulumi.CustomResourceState

	// Disables the syncing process between Vault and external destinations. Defaults to `false`.
	Disabled pulumi.BoolPtrOutput `pulumi:"disabled"`
	// The namespace to provision the resource in.
	// This resource can only be configured in the root namespace.
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Maximum number of pending sync operations allowed on the queue. Defaults to `1000000`.
	QueueCapacity pulumi.IntPtrOutput `pulumi:"queueCapacity"`
}

// NewSyncConfig registers a new resource with the given unique name, arguments, and options.
func NewSyncConfig(ctx *pulumi.Context,
	name string, args *SyncConfigArgs, opts ...pulumi.ResourceOption) (*SyncConfig, error) {
	if args == nil {
		args = &SyncConfigArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SyncConfig
	err := ctx.RegisterResource("vault:secrets/syncConfig:SyncConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSyncConfig gets an existing SyncConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSyncConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SyncConfigState, opts ...pulumi.ResourceOption) (*SyncConfig, error) {
	var resource SyncConfig
	err := ctx.ReadResource("vault:secrets/syncConfig:SyncConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SyncConfig resources.
type syncConfigState struct {
	// Disables the syncing process between Vault and external destinations. Defaults to `false`.
	Disabled *bool `pulumi:"disabled"`
	// The namespace to provision the resource in.
	// This resource can only be configured in the root namespace.
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Maximum number of pending sync operations allowed on the queue. Defaults to `1000000`.
	QueueCapacity *int `pulumi:"queueCapacity"`
}

type SyncConfigState struct {
	// Disables the syncing process between Vault and external destinations. Defaults to `false`.
	Disabled pulumi.BoolPtrInput
	// The namespace to provision the resource in.
	// This resource can only be configured in the root namespace.
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Maximum number of pending sync operations allowed on the queue. Defaults to `1000000`.
	QueueCapacity pulumi.IntPtrInput
}

func (SyncConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*syncConfigState)(nil)).Elem()
}

type syncConfigArgs struct {
	// Disables the syncing process between Vault and external destinations. Defaults to `false`.
	Disabled *bool `pulumi:"disabled"`
	// The namespace to provision the resource in.
	// This resource can only be configured in the root namespace.
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Maximum number of pending sync operations allowed on the queue. Defaults to `1000000`.
	QueueCapacity *int `pulumi:"queueCapacity"`
}

// The set of arguments for constructing a SyncConfig resource.
type SyncConfigArgs struct {
	// Disables the syncing process between Vault and external destinations. Defaults to `false`.
	Disabled pulumi.BoolPtrInput
	// The namespace to provision the resource in.
	// This resource can only be configured in the root namespace.
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Maximum number of pending sync operations allowed on the queue. Defaults to `1000000`.
	QueueCapacity pulumi.IntPtrInput
}

func (SyncConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*syncConfigArgs)(nil)).Elem()
}

type SyncConfigInput interface {
	pulumi.Input

	ToSyncConfigOutput() SyncConfigOutput
	ToSyncConfigOutputWithContext(ctx context.Context) SyncConfigOutput
}

func (*SyncConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncConfig)(nil)).Elem()
}

func (i *SyncConfig) ToSyncConfigOutput() SyncConfigOutput {
	return i.ToSyncConfigOutputWithContext(context.Background())
}

func (i *SyncConfig) ToSyncConfigOutputWithContext(ctx context.Context) SyncConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncConfigOutput)
}

// SyncConfigArrayInput is an input type that accepts SyncConfigArray and SyncConfigArrayOutput values.
// You can construct a concrete instance of `SyncConfigArrayInput` via:
//
//	SyncConfigArray{ SyncConfigArgs{...} }
type SyncConfigArrayInput interface {
	pulumi.Input

	ToSyncConfigArrayOutput() SyncConfigArrayOutput
	ToSyncConfigArrayOutputWithContext(context.Context) SyncConfigArrayOutput
}

type SyncConfigArray []SyncConfigInput

func (SyncConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyncConfig)(nil)).Elem()
}

func (i SyncConfigArray) ToSyncConfigArrayOutput() SyncConfigArrayOutput {
	return i.ToSyncConfigArrayOutputWithContext(context.Background())
}

func (i SyncConfigArray) ToSyncConfigArrayOutputWithContext(ctx context.Context) SyncConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncConfigArrayOutput)
}

// SyncConfigMapInput is an input type that accepts SyncConfigMap and SyncConfigMapOutput values.
// You can construct a concrete instance of `SyncConfigMapInput` via:
//
//	SyncConfigMap{ "key": SyncConfigArgs{...} }
type SyncConfigMapInput interface {
	pulumi.Input

	ToSyncConfigMapOutput() SyncConfigMapOutput
	ToSyncConfigMapOutputWithContext(context.Context) SyncConfigMapOutput
}

type SyncConfigMap map[string]SyncConfigInput

func (SyncConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyncConfig)(nil)).Elem()
}

func (i SyncConfigMap) ToSyncConfigMapOutput() SyncConfigMapOutput {
	return i.ToSyncConfigMapOutputWithContext(context.Background())
}

func (i SyncConfigMap) ToSyncConfigMapOutputWithContext(ctx context.Context) SyncConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncConfigMapOutput)
}

type SyncConfigOutput struct{ *pulumi.OutputState }

func (SyncConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncConfig)(nil)).Elem()
}

func (o SyncConfigOutput) ToSyncConfigOutput() SyncConfigOutput {
	return o
}

func (o SyncConfigOutput) ToSyncConfigOutputWithContext(ctx context.Context) SyncConfigOutput {
	return o
}

// Disables the syncing process between Vault and external destinations. Defaults to `false`.
func (o SyncConfigOutput) Disabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SyncConfig) pulumi.BoolPtrOutput { return v.Disabled }).(pulumi.BoolPtrOutput)
}

// The namespace to provision the resource in.
// This resource can only be configured in the root namespace.
// *Available only for Vault Enterprise*.
func (o SyncConfigOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncConfig) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Maximum number of pending sync operations allowed on the queue. Defaults to `1000000`.
func (o SyncConfigOutput) QueueCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SyncConfig) pulumi.IntPtrOutput { return v.QueueCapacity }).(pulumi.IntPtrOutput)
}

type SyncConfigArrayOutput struct{ *pulumi.OutputState }

func (SyncConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SyncConfig)(nil)).Elem()
}

func (o SyncConfigArrayOutput) ToSyncConfigArrayOutput() SyncConfigArrayOutput {
	return o
}

func (o SyncConfigArrayOutput) ToSyncConfigArrayOutputWithContext(ctx context.Context) SyncConfigArrayOutput {
	return o
}

func (o SyncConfigArrayOutput) Index(i pulumi.IntInput) SyncConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SyncConfig {
		return vs[0].([]*SyncConfig)[vs[1].(int)]
	}).(SyncConfigOutput)
}

type SyncConfigMapOutput struct{ *pulumi.OutputState }

func (SyncConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SyncConfig)(nil)).Elem()
}

func (o SyncConfigMapOutput) ToSyncConfigMapOutput() SyncConfigMapOutput {
	return o
}

func (o SyncConfigMapOutput) ToSyncConfigMapOutputWithContext(ctx context.Context) SyncConfigMapOutput {
	return o
}

func (o SyncConfigMapOutput) MapIndex(k pulumi.StringInput) SyncConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SyncConfig {
		return vs[0].(map[string]*SyncConfig)[vs[1].(string)]
	}).(SyncConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SyncConfigInput)(nil)).Elem(), &SyncConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncConfigArrayInput)(nil)).Elem(), SyncConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SyncConfigMapInput)(nil)).Elem(), SyncConfigMap{})
	pulumi.RegisterOutputType(SyncConfigOutput{})
	pulumi.RegisterOutputType(SyncConfigArrayOutput{})
	pulumi.RegisterOutputType(SyncConfigMapOutput{})
}
