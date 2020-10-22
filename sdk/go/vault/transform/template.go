// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package transform

import (
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
	if args == nil || args.Path == nil {
		return nil, errors.New("missing required argument 'Path'")
	}
	if args == nil {
		args = &TemplateArgs{}
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
