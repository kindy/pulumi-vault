// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package transit

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Configure the cache for the Transit Secret Backend in Vault.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v3/go/vault"
// 	"github.com/pulumi/pulumi-vault/sdk/v3/go/vault/transit"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		transit, err := vault.NewMount(ctx, "transit", &vault.MountArgs{
// 			DefaultLeaseTtlSeconds: pulumi.Int(3600),
// 			Description:            pulumi.String("Example description"),
// 			MaxLeaseTtlSeconds:     pulumi.Int(86400),
// 			Path:                   pulumi.String("transit"),
// 			Type:                   pulumi.String("transit"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = transit.NewSecretCacheConfig(ctx, "cfg", &transit.SecretCacheConfigArgs{
// 			Backend: transit.Path,
// 			Size:    pulumi.Int(500),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type SecretCacheConfig struct {
	pulumi.CustomResourceState

	// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// The number of cache entries. 0 means unlimited.
	Size pulumi.IntOutput `pulumi:"size"`
}

// NewSecretCacheConfig registers a new resource with the given unique name, arguments, and options.
func NewSecretCacheConfig(ctx *pulumi.Context,
	name string, args *SecretCacheConfigArgs, opts ...pulumi.ResourceOption) (*SecretCacheConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	if args.Size == nil {
		return nil, errors.New("invalid value for required argument 'Size'")
	}
	var resource SecretCacheConfig
	err := ctx.RegisterResource("vault:transit/secretCacheConfig:SecretCacheConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretCacheConfig gets an existing SecretCacheConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretCacheConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretCacheConfigState, opts ...pulumi.ResourceOption) (*SecretCacheConfig, error) {
	var resource SecretCacheConfig
	err := ctx.ReadResource("vault:transit/secretCacheConfig:SecretCacheConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretCacheConfig resources.
type secretCacheConfigState struct {
	// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
	Backend *string `pulumi:"backend"`
	// The number of cache entries. 0 means unlimited.
	Size *int `pulumi:"size"`
}

type SecretCacheConfigState struct {
	// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
	Backend pulumi.StringPtrInput
	// The number of cache entries. 0 means unlimited.
	Size pulumi.IntPtrInput
}

func (SecretCacheConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretCacheConfigState)(nil)).Elem()
}

type secretCacheConfigArgs struct {
	// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
	Backend string `pulumi:"backend"`
	// The number of cache entries. 0 means unlimited.
	Size int `pulumi:"size"`
}

// The set of arguments for constructing a SecretCacheConfig resource.
type SecretCacheConfigArgs struct {
	// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
	Backend pulumi.StringInput
	// The number of cache entries. 0 means unlimited.
	Size pulumi.IntInput
}

func (SecretCacheConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretCacheConfigArgs)(nil)).Elem()
}

type SecretCacheConfigInput interface {
	pulumi.Input

	ToSecretCacheConfigOutput() SecretCacheConfigOutput
	ToSecretCacheConfigOutputWithContext(ctx context.Context) SecretCacheConfigOutput
}

func (*SecretCacheConfig) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretCacheConfig)(nil))
}

func (i *SecretCacheConfig) ToSecretCacheConfigOutput() SecretCacheConfigOutput {
	return i.ToSecretCacheConfigOutputWithContext(context.Background())
}

func (i *SecretCacheConfig) ToSecretCacheConfigOutputWithContext(ctx context.Context) SecretCacheConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretCacheConfigOutput)
}

func (i *SecretCacheConfig) ToSecretCacheConfigPtrOutput() SecretCacheConfigPtrOutput {
	return i.ToSecretCacheConfigPtrOutputWithContext(context.Background())
}

func (i *SecretCacheConfig) ToSecretCacheConfigPtrOutputWithContext(ctx context.Context) SecretCacheConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretCacheConfigPtrOutput)
}

type SecretCacheConfigPtrInput interface {
	pulumi.Input

	ToSecretCacheConfigPtrOutput() SecretCacheConfigPtrOutput
	ToSecretCacheConfigPtrOutputWithContext(ctx context.Context) SecretCacheConfigPtrOutput
}

type secretCacheConfigPtrType SecretCacheConfigArgs

func (*secretCacheConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretCacheConfig)(nil))
}

func (i *secretCacheConfigPtrType) ToSecretCacheConfigPtrOutput() SecretCacheConfigPtrOutput {
	return i.ToSecretCacheConfigPtrOutputWithContext(context.Background())
}

func (i *secretCacheConfigPtrType) ToSecretCacheConfigPtrOutputWithContext(ctx context.Context) SecretCacheConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretCacheConfigPtrOutput)
}

// SecretCacheConfigArrayInput is an input type that accepts SecretCacheConfigArray and SecretCacheConfigArrayOutput values.
// You can construct a concrete instance of `SecretCacheConfigArrayInput` via:
//
//          SecretCacheConfigArray{ SecretCacheConfigArgs{...} }
type SecretCacheConfigArrayInput interface {
	pulumi.Input

	ToSecretCacheConfigArrayOutput() SecretCacheConfigArrayOutput
	ToSecretCacheConfigArrayOutputWithContext(context.Context) SecretCacheConfigArrayOutput
}

type SecretCacheConfigArray []SecretCacheConfigInput

func (SecretCacheConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SecretCacheConfig)(nil))
}

func (i SecretCacheConfigArray) ToSecretCacheConfigArrayOutput() SecretCacheConfigArrayOutput {
	return i.ToSecretCacheConfigArrayOutputWithContext(context.Background())
}

func (i SecretCacheConfigArray) ToSecretCacheConfigArrayOutputWithContext(ctx context.Context) SecretCacheConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretCacheConfigArrayOutput)
}

// SecretCacheConfigMapInput is an input type that accepts SecretCacheConfigMap and SecretCacheConfigMapOutput values.
// You can construct a concrete instance of `SecretCacheConfigMapInput` via:
//
//          SecretCacheConfigMap{ "key": SecretCacheConfigArgs{...} }
type SecretCacheConfigMapInput interface {
	pulumi.Input

	ToSecretCacheConfigMapOutput() SecretCacheConfigMapOutput
	ToSecretCacheConfigMapOutputWithContext(context.Context) SecretCacheConfigMapOutput
}

type SecretCacheConfigMap map[string]SecretCacheConfigInput

func (SecretCacheConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SecretCacheConfig)(nil))
}

func (i SecretCacheConfigMap) ToSecretCacheConfigMapOutput() SecretCacheConfigMapOutput {
	return i.ToSecretCacheConfigMapOutputWithContext(context.Background())
}

func (i SecretCacheConfigMap) ToSecretCacheConfigMapOutputWithContext(ctx context.Context) SecretCacheConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretCacheConfigMapOutput)
}

type SecretCacheConfigOutput struct {
	*pulumi.OutputState
}

func (SecretCacheConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretCacheConfig)(nil))
}

func (o SecretCacheConfigOutput) ToSecretCacheConfigOutput() SecretCacheConfigOutput {
	return o
}

func (o SecretCacheConfigOutput) ToSecretCacheConfigOutputWithContext(ctx context.Context) SecretCacheConfigOutput {
	return o
}

func (o SecretCacheConfigOutput) ToSecretCacheConfigPtrOutput() SecretCacheConfigPtrOutput {
	return o.ToSecretCacheConfigPtrOutputWithContext(context.Background())
}

func (o SecretCacheConfigOutput) ToSecretCacheConfigPtrOutputWithContext(ctx context.Context) SecretCacheConfigPtrOutput {
	return o.ApplyT(func(v SecretCacheConfig) *SecretCacheConfig {
		return &v
	}).(SecretCacheConfigPtrOutput)
}

type SecretCacheConfigPtrOutput struct {
	*pulumi.OutputState
}

func (SecretCacheConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretCacheConfig)(nil))
}

func (o SecretCacheConfigPtrOutput) ToSecretCacheConfigPtrOutput() SecretCacheConfigPtrOutput {
	return o
}

func (o SecretCacheConfigPtrOutput) ToSecretCacheConfigPtrOutputWithContext(ctx context.Context) SecretCacheConfigPtrOutput {
	return o
}

type SecretCacheConfigArrayOutput struct{ *pulumi.OutputState }

func (SecretCacheConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretCacheConfig)(nil))
}

func (o SecretCacheConfigArrayOutput) ToSecretCacheConfigArrayOutput() SecretCacheConfigArrayOutput {
	return o
}

func (o SecretCacheConfigArrayOutput) ToSecretCacheConfigArrayOutputWithContext(ctx context.Context) SecretCacheConfigArrayOutput {
	return o
}

func (o SecretCacheConfigArrayOutput) Index(i pulumi.IntInput) SecretCacheConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecretCacheConfig {
		return vs[0].([]SecretCacheConfig)[vs[1].(int)]
	}).(SecretCacheConfigOutput)
}

type SecretCacheConfigMapOutput struct{ *pulumi.OutputState }

func (SecretCacheConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SecretCacheConfig)(nil))
}

func (o SecretCacheConfigMapOutput) ToSecretCacheConfigMapOutput() SecretCacheConfigMapOutput {
	return o
}

func (o SecretCacheConfigMapOutput) ToSecretCacheConfigMapOutputWithContext(ctx context.Context) SecretCacheConfigMapOutput {
	return o
}

func (o SecretCacheConfigMapOutput) MapIndex(k pulumi.StringInput) SecretCacheConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SecretCacheConfig {
		return vs[0].(map[string]SecretCacheConfig)[vs[1].(string)]
	}).(SecretCacheConfigOutput)
}

func init() {
	pulumi.RegisterOutputType(SecretCacheConfigOutput{})
	pulumi.RegisterOutputType(SecretCacheConfigPtrOutput{})
	pulumi.RegisterOutputType(SecretCacheConfigArrayOutput{})
	pulumi.RegisterOutputType(SecretCacheConfigMapOutput{})
}
