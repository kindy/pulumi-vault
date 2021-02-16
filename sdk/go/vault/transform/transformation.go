// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package transform

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// This resource supports the "/transform/transformation/{name}" Vault endpoint.
//
// It creates or updates a transformation with the given name. If a transformation with the name does not exist,
// it will be created. If the transformation exists, it will be updated with the new attributes.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v3/go/vault"
// 	"github.com/pulumi/pulumi-vault/sdk/v3/go/vault/transform"
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
// 		_, err = transform.NewTransformation(ctx, "test", &transform.TransformationArgs{
// 			Path:        mountTransform.Path,
// 			Type:        pulumi.String("fpe"),
// 			Template:    pulumi.String("ccn"),
// 			TweakSource: pulumi.String("internal"),
// 			AllowedRoles: pulumi.StringArray{
// 				pulumi.String("payments"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type Transformation struct {
	pulumi.CustomResourceState

	// The set of roles allowed to perform this transformation.
	AllowedRoles pulumi.StringArrayOutput `pulumi:"allowedRoles"`
	// The character used to replace data when in masking mode
	MaskingCharacter pulumi.StringPtrOutput `pulumi:"maskingCharacter"`
	// The name of the transformation.
	Name pulumi.StringOutput `pulumi:"name"`
	// Path to where the back-end is mounted within Vault.
	Path pulumi.StringOutput `pulumi:"path"`
	// The name of the template to use.
	Template pulumi.StringPtrOutput `pulumi:"template"`
	// Templates configured for transformation.
	Templates pulumi.StringArrayOutput `pulumi:"templates"`
	// The source of where the tweak value comes from. Only valid when in FPE mode.
	TweakSource pulumi.StringPtrOutput `pulumi:"tweakSource"`
	// The type of transformation to perform.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewTransformation registers a new resource with the given unique name, arguments, and options.
func NewTransformation(ctx *pulumi.Context,
	name string, args *TransformationArgs, opts ...pulumi.ResourceOption) (*Transformation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	var resource Transformation
	err := ctx.RegisterResource("vault:transform/transformation:Transformation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransformation gets an existing Transformation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransformation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransformationState, opts ...pulumi.ResourceOption) (*Transformation, error) {
	var resource Transformation
	err := ctx.ReadResource("vault:transform/transformation:Transformation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Transformation resources.
type transformationState struct {
	// The set of roles allowed to perform this transformation.
	AllowedRoles []string `pulumi:"allowedRoles"`
	// The character used to replace data when in masking mode
	MaskingCharacter *string `pulumi:"maskingCharacter"`
	// The name of the transformation.
	Name *string `pulumi:"name"`
	// Path to where the back-end is mounted within Vault.
	Path *string `pulumi:"path"`
	// The name of the template to use.
	Template *string `pulumi:"template"`
	// Templates configured for transformation.
	Templates []string `pulumi:"templates"`
	// The source of where the tweak value comes from. Only valid when in FPE mode.
	TweakSource *string `pulumi:"tweakSource"`
	// The type of transformation to perform.
	Type *string `pulumi:"type"`
}

type TransformationState struct {
	// The set of roles allowed to perform this transformation.
	AllowedRoles pulumi.StringArrayInput
	// The character used to replace data when in masking mode
	MaskingCharacter pulumi.StringPtrInput
	// The name of the transformation.
	Name pulumi.StringPtrInput
	// Path to where the back-end is mounted within Vault.
	Path pulumi.StringPtrInput
	// The name of the template to use.
	Template pulumi.StringPtrInput
	// Templates configured for transformation.
	Templates pulumi.StringArrayInput
	// The source of where the tweak value comes from. Only valid when in FPE mode.
	TweakSource pulumi.StringPtrInput
	// The type of transformation to perform.
	Type pulumi.StringPtrInput
}

func (TransformationState) ElementType() reflect.Type {
	return reflect.TypeOf((*transformationState)(nil)).Elem()
}

type transformationArgs struct {
	// The set of roles allowed to perform this transformation.
	AllowedRoles []string `pulumi:"allowedRoles"`
	// The character used to replace data when in masking mode
	MaskingCharacter *string `pulumi:"maskingCharacter"`
	// The name of the transformation.
	Name *string `pulumi:"name"`
	// Path to where the back-end is mounted within Vault.
	Path string `pulumi:"path"`
	// The name of the template to use.
	Template *string `pulumi:"template"`
	// Templates configured for transformation.
	Templates []string `pulumi:"templates"`
	// The source of where the tweak value comes from. Only valid when in FPE mode.
	TweakSource *string `pulumi:"tweakSource"`
	// The type of transformation to perform.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Transformation resource.
type TransformationArgs struct {
	// The set of roles allowed to perform this transformation.
	AllowedRoles pulumi.StringArrayInput
	// The character used to replace data when in masking mode
	MaskingCharacter pulumi.StringPtrInput
	// The name of the transformation.
	Name pulumi.StringPtrInput
	// Path to where the back-end is mounted within Vault.
	Path pulumi.StringInput
	// The name of the template to use.
	Template pulumi.StringPtrInput
	// Templates configured for transformation.
	Templates pulumi.StringArrayInput
	// The source of where the tweak value comes from. Only valid when in FPE mode.
	TweakSource pulumi.StringPtrInput
	// The type of transformation to perform.
	Type pulumi.StringPtrInput
}

func (TransformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transformationArgs)(nil)).Elem()
}

type TransformationInput interface {
	pulumi.Input

	ToTransformationOutput() TransformationOutput
	ToTransformationOutputWithContext(ctx context.Context) TransformationOutput
}

func (*Transformation) ElementType() reflect.Type {
	return reflect.TypeOf((*Transformation)(nil))
}

func (i *Transformation) ToTransformationOutput() TransformationOutput {
	return i.ToTransformationOutputWithContext(context.Background())
}

func (i *Transformation) ToTransformationOutputWithContext(ctx context.Context) TransformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformationOutput)
}

func (i *Transformation) ToTransformationPtrOutput() TransformationPtrOutput {
	return i.ToTransformationPtrOutputWithContext(context.Background())
}

func (i *Transformation) ToTransformationPtrOutputWithContext(ctx context.Context) TransformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformationPtrOutput)
}

type TransformationPtrInput interface {
	pulumi.Input

	ToTransformationPtrOutput() TransformationPtrOutput
	ToTransformationPtrOutputWithContext(ctx context.Context) TransformationPtrOutput
}

type transformationPtrType TransformationArgs

func (*transformationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Transformation)(nil))
}

func (i *transformationPtrType) ToTransformationPtrOutput() TransformationPtrOutput {
	return i.ToTransformationPtrOutputWithContext(context.Background())
}

func (i *transformationPtrType) ToTransformationPtrOutputWithContext(ctx context.Context) TransformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformationPtrOutput)
}

// TransformationArrayInput is an input type that accepts TransformationArray and TransformationArrayOutput values.
// You can construct a concrete instance of `TransformationArrayInput` via:
//
//          TransformationArray{ TransformationArgs{...} }
type TransformationArrayInput interface {
	pulumi.Input

	ToTransformationArrayOutput() TransformationArrayOutput
	ToTransformationArrayOutputWithContext(context.Context) TransformationArrayOutput
}

type TransformationArray []TransformationInput

func (TransformationArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Transformation)(nil))
}

func (i TransformationArray) ToTransformationArrayOutput() TransformationArrayOutput {
	return i.ToTransformationArrayOutputWithContext(context.Background())
}

func (i TransformationArray) ToTransformationArrayOutputWithContext(ctx context.Context) TransformationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformationArrayOutput)
}

// TransformationMapInput is an input type that accepts TransformationMap and TransformationMapOutput values.
// You can construct a concrete instance of `TransformationMapInput` via:
//
//          TransformationMap{ "key": TransformationArgs{...} }
type TransformationMapInput interface {
	pulumi.Input

	ToTransformationMapOutput() TransformationMapOutput
	ToTransformationMapOutputWithContext(context.Context) TransformationMapOutput
}

type TransformationMap map[string]TransformationInput

func (TransformationMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Transformation)(nil))
}

func (i TransformationMap) ToTransformationMapOutput() TransformationMapOutput {
	return i.ToTransformationMapOutputWithContext(context.Background())
}

func (i TransformationMap) ToTransformationMapOutputWithContext(ctx context.Context) TransformationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformationMapOutput)
}

type TransformationOutput struct {
	*pulumi.OutputState
}

func (TransformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Transformation)(nil))
}

func (o TransformationOutput) ToTransformationOutput() TransformationOutput {
	return o
}

func (o TransformationOutput) ToTransformationOutputWithContext(ctx context.Context) TransformationOutput {
	return o
}

func (o TransformationOutput) ToTransformationPtrOutput() TransformationPtrOutput {
	return o.ToTransformationPtrOutputWithContext(context.Background())
}

func (o TransformationOutput) ToTransformationPtrOutputWithContext(ctx context.Context) TransformationPtrOutput {
	return o.ApplyT(func(v Transformation) *Transformation {
		return &v
	}).(TransformationPtrOutput)
}

type TransformationPtrOutput struct {
	*pulumi.OutputState
}

func (TransformationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Transformation)(nil))
}

func (o TransformationPtrOutput) ToTransformationPtrOutput() TransformationPtrOutput {
	return o
}

func (o TransformationPtrOutput) ToTransformationPtrOutputWithContext(ctx context.Context) TransformationPtrOutput {
	return o
}

type TransformationArrayOutput struct{ *pulumi.OutputState }

func (TransformationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Transformation)(nil))
}

func (o TransformationArrayOutput) ToTransformationArrayOutput() TransformationArrayOutput {
	return o
}

func (o TransformationArrayOutput) ToTransformationArrayOutputWithContext(ctx context.Context) TransformationArrayOutput {
	return o
}

func (o TransformationArrayOutput) Index(i pulumi.IntInput) TransformationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Transformation {
		return vs[0].([]Transformation)[vs[1].(int)]
	}).(TransformationOutput)
}

type TransformationMapOutput struct{ *pulumi.OutputState }

func (TransformationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Transformation)(nil))
}

func (o TransformationMapOutput) ToTransformationMapOutput() TransformationMapOutput {
	return o
}

func (o TransformationMapOutput) ToTransformationMapOutputWithContext(ctx context.Context) TransformationMapOutput {
	return o
}

func (o TransformationMapOutput) MapIndex(k pulumi.StringInput) TransformationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Transformation {
		return vs[0].(map[string]Transformation)[vs[1].(string)]
	}).(TransformationOutput)
}

func init() {
	pulumi.RegisterOutputType(TransformationOutput{})
	pulumi.RegisterOutputType(TransformationPtrOutput{})
	pulumi.RegisterOutputType(TransformationArrayOutput{})
	pulumi.RegisterOutputType(TransformationMapOutput{})
}
