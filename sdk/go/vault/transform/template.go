// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package transform

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Template struct {
	pulumi.CustomResourceState

	// The alphabet to use for this template. This is only used during FPE transformations.
	Alphabet pulumi.StringPtrOutput `pulumi:"alphabet"`
	// The name of the template.
	Name pulumi.StringOutput `pulumi:"name"`
	// Path to where the back-end is mounted within Vault.
	Path pulumi.StringOutput `pulumi:"path"`
	// The pattern used for matching. Currently, only regular expression pattern is supported.
	Pattern pulumi.StringPtrOutput `pulumi:"pattern"`
	// The pattern type to use for match detection. Currently, only regex is supported.
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewTemplate registers a new resource with the given unique name, arguments, and options.
func NewTemplate(ctx *pulumi.Context,
	name string, args *TemplateArgs, opts ...pulumi.ResourceOption) (*Template, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Path == nil {
		return nil, errors.New("invalid value for required argument 'Path'")
	}
	var resource Template
	err := ctx.RegisterResource("vault:transform/template:Template", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTemplate gets an existing Template resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TemplateState, opts ...pulumi.ResourceOption) (*Template, error) {
	var resource Template
	err := ctx.ReadResource("vault:transform/template:Template", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Template resources.
type templateState struct {
	// The alphabet to use for this template. This is only used during FPE transformations.
	Alphabet *string `pulumi:"alphabet"`
	// The name of the template.
	Name *string `pulumi:"name"`
	// Path to where the back-end is mounted within Vault.
	Path *string `pulumi:"path"`
	// The pattern used for matching. Currently, only regular expression pattern is supported.
	Pattern *string `pulumi:"pattern"`
	// The pattern type to use for match detection. Currently, only regex is supported.
	Type *string `pulumi:"type"`
}

type TemplateState struct {
	// The alphabet to use for this template. This is only used during FPE transformations.
	Alphabet pulumi.StringPtrInput
	// The name of the template.
	Name pulumi.StringPtrInput
	// Path to where the back-end is mounted within Vault.
	Path pulumi.StringPtrInput
	// The pattern used for matching. Currently, only regular expression pattern is supported.
	Pattern pulumi.StringPtrInput
	// The pattern type to use for match detection. Currently, only regex is supported.
	Type pulumi.StringPtrInput
}

func (TemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*templateState)(nil)).Elem()
}

type templateArgs struct {
	// The alphabet to use for this template. This is only used during FPE transformations.
	Alphabet *string `pulumi:"alphabet"`
	// The name of the template.
	Name *string `pulumi:"name"`
	// Path to where the back-end is mounted within Vault.
	Path string `pulumi:"path"`
	// The pattern used for matching. Currently, only regular expression pattern is supported.
	Pattern *string `pulumi:"pattern"`
	// The pattern type to use for match detection. Currently, only regex is supported.
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a Template resource.
type TemplateArgs struct {
	// The alphabet to use for this template. This is only used during FPE transformations.
	Alphabet pulumi.StringPtrInput
	// The name of the template.
	Name pulumi.StringPtrInput
	// Path to where the back-end is mounted within Vault.
	Path pulumi.StringInput
	// The pattern used for matching. Currently, only regular expression pattern is supported.
	Pattern pulumi.StringPtrInput
	// The pattern type to use for match detection. Currently, only regex is supported.
	Type pulumi.StringPtrInput
}

func (TemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*templateArgs)(nil)).Elem()
}

type TemplateInput interface {
	pulumi.Input

	ToTemplateOutput() TemplateOutput
	ToTemplateOutputWithContext(ctx context.Context) TemplateOutput
}

func (*Template) ElementType() reflect.Type {
	return reflect.TypeOf((*Template)(nil))
}

func (i *Template) ToTemplateOutput() TemplateOutput {
	return i.ToTemplateOutputWithContext(context.Background())
}

func (i *Template) ToTemplateOutputWithContext(ctx context.Context) TemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateOutput)
}

func (i *Template) ToTemplatePtrOutput() TemplatePtrOutput {
	return i.ToTemplatePtrOutputWithContext(context.Background())
}

func (i *Template) ToTemplatePtrOutputWithContext(ctx context.Context) TemplatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplatePtrOutput)
}

type TemplatePtrInput interface {
	pulumi.Input

	ToTemplatePtrOutput() TemplatePtrOutput
	ToTemplatePtrOutputWithContext(ctx context.Context) TemplatePtrOutput
}

type templatePtrType TemplateArgs

func (*templatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Template)(nil))
}

func (i *templatePtrType) ToTemplatePtrOutput() TemplatePtrOutput {
	return i.ToTemplatePtrOutputWithContext(context.Background())
}

func (i *templatePtrType) ToTemplatePtrOutputWithContext(ctx context.Context) TemplatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplatePtrOutput)
}

// TemplateArrayInput is an input type that accepts TemplateArray and TemplateArrayOutput values.
// You can construct a concrete instance of `TemplateArrayInput` via:
//
//          TemplateArray{ TemplateArgs{...} }
type TemplateArrayInput interface {
	pulumi.Input

	ToTemplateArrayOutput() TemplateArrayOutput
	ToTemplateArrayOutputWithContext(context.Context) TemplateArrayOutput
}

type TemplateArray []TemplateInput

func (TemplateArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Template)(nil))
}

func (i TemplateArray) ToTemplateArrayOutput() TemplateArrayOutput {
	return i.ToTemplateArrayOutputWithContext(context.Background())
}

func (i TemplateArray) ToTemplateArrayOutputWithContext(ctx context.Context) TemplateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateArrayOutput)
}

// TemplateMapInput is an input type that accepts TemplateMap and TemplateMapOutput values.
// You can construct a concrete instance of `TemplateMapInput` via:
//
//          TemplateMap{ "key": TemplateArgs{...} }
type TemplateMapInput interface {
	pulumi.Input

	ToTemplateMapOutput() TemplateMapOutput
	ToTemplateMapOutputWithContext(context.Context) TemplateMapOutput
}

type TemplateMap map[string]TemplateInput

func (TemplateMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Template)(nil))
}

func (i TemplateMap) ToTemplateMapOutput() TemplateMapOutput {
	return i.ToTemplateMapOutputWithContext(context.Background())
}

func (i TemplateMap) ToTemplateMapOutputWithContext(ctx context.Context) TemplateMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateMapOutput)
}

type TemplateOutput struct {
	*pulumi.OutputState
}

func (TemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Template)(nil))
}

func (o TemplateOutput) ToTemplateOutput() TemplateOutput {
	return o
}

func (o TemplateOutput) ToTemplateOutputWithContext(ctx context.Context) TemplateOutput {
	return o
}

func (o TemplateOutput) ToTemplatePtrOutput() TemplatePtrOutput {
	return o.ToTemplatePtrOutputWithContext(context.Background())
}

func (o TemplateOutput) ToTemplatePtrOutputWithContext(ctx context.Context) TemplatePtrOutput {
	return o.ApplyT(func(v Template) *Template {
		return &v
	}).(TemplatePtrOutput)
}

type TemplatePtrOutput struct {
	*pulumi.OutputState
}

func (TemplatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Template)(nil))
}

func (o TemplatePtrOutput) ToTemplatePtrOutput() TemplatePtrOutput {
	return o
}

func (o TemplatePtrOutput) ToTemplatePtrOutputWithContext(ctx context.Context) TemplatePtrOutput {
	return o
}

type TemplateArrayOutput struct{ *pulumi.OutputState }

func (TemplateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Template)(nil))
}

func (o TemplateArrayOutput) ToTemplateArrayOutput() TemplateArrayOutput {
	return o
}

func (o TemplateArrayOutput) ToTemplateArrayOutputWithContext(ctx context.Context) TemplateArrayOutput {
	return o
}

func (o TemplateArrayOutput) Index(i pulumi.IntInput) TemplateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Template {
		return vs[0].([]Template)[vs[1].(int)]
	}).(TemplateOutput)
}

type TemplateMapOutput struct{ *pulumi.OutputState }

func (TemplateMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Template)(nil))
}

func (o TemplateMapOutput) ToTemplateMapOutput() TemplateMapOutput {
	return o
}

func (o TemplateMapOutput) ToTemplateMapOutputWithContext(ctx context.Context) TemplateMapOutput {
	return o
}

func (o TemplateMapOutput) MapIndex(k pulumi.StringInput) TemplateOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Template {
		return vs[0].(map[string]Template)[vs[1].(string)]
	}).(TemplateOutput)
}

func init() {
	pulumi.RegisterOutputType(TemplateOutput{})
	pulumi.RegisterOutputType(TemplatePtrOutput{})
	pulumi.RegisterOutputType(TemplateArrayOutput{})
	pulumi.RegisterOutputType(TemplateMapOutput{})
}
