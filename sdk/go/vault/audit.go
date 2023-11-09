// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// ## Example Usage
// ### File Audit Device)
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vault.NewAudit(ctx, "test", &vault.AuditArgs{
//				Options: pulumi.StringMap{
//					"file_path": pulumi.String("C:/temp/audit.txt"),
//				},
//				Type: pulumi.String("file"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Socket Audit Device)
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vault.NewAudit(ctx, "test", &vault.AuditArgs{
//				Local: pulumi.Bool(false),
//				Options: pulumi.StringMap{
//					"address":     pulumi.String("127.0.0.1:8000"),
//					"description": pulumi.String("application x socket"),
//					"socket_type": pulumi.String("tcp"),
//				},
//				Path: pulumi.String("app_socket"),
//				Type: pulumi.String("socket"),
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
// Audit devices can be imported using the `path`, e.g.
//
// ```sh
//
//	$ pulumi import vault:index/audit:Audit test syslog
//
// ```
type Audit struct {
	pulumi.CustomResourceState

	// Human-friendly description of the audit device.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Specifies if the audit device is a local only. Local audit devices are not replicated nor (if a secondary) removed by replication.
	Local pulumi.BoolPtrOutput `pulumi:"local"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Configuration options to pass to the audit device itself.
	//
	// For a reference of the device types and their options, consult the [Vault documentation.](https://www.vaultproject.io/docs/audit/index.html)
	Options pulumi.StringMapOutput `pulumi:"options"`
	// The path to mount the audit device. This defaults to the type.
	Path pulumi.StringOutput `pulumi:"path"`
	// Type of the audit device, such as 'file'.
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAudit registers a new resource with the given unique name, arguments, and options.
func NewAudit(ctx *pulumi.Context,
	name string, args *AuditArgs, opts ...pulumi.ResourceOption) (*Audit, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Options == nil {
		return nil, errors.New("invalid value for required argument 'Options'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Audit
	err := ctx.RegisterResource("vault:index/audit:Audit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAudit gets an existing Audit resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAudit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuditState, opts ...pulumi.ResourceOption) (*Audit, error) {
	var resource Audit
	err := ctx.ReadResource("vault:index/audit:Audit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Audit resources.
type auditState struct {
	// Human-friendly description of the audit device.
	Description *string `pulumi:"description"`
	// Specifies if the audit device is a local only. Local audit devices are not replicated nor (if a secondary) removed by replication.
	Local *bool `pulumi:"local"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Configuration options to pass to the audit device itself.
	//
	// For a reference of the device types and their options, consult the [Vault documentation.](https://www.vaultproject.io/docs/audit/index.html)
	Options map[string]string `pulumi:"options"`
	// The path to mount the audit device. This defaults to the type.
	Path *string `pulumi:"path"`
	// Type of the audit device, such as 'file'.
	Type *string `pulumi:"type"`
}

type AuditState struct {
	// Human-friendly description of the audit device.
	Description pulumi.StringPtrInput
	// Specifies if the audit device is a local only. Local audit devices are not replicated nor (if a secondary) removed by replication.
	Local pulumi.BoolPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Configuration options to pass to the audit device itself.
	//
	// For a reference of the device types and their options, consult the [Vault documentation.](https://www.vaultproject.io/docs/audit/index.html)
	Options pulumi.StringMapInput
	// The path to mount the audit device. This defaults to the type.
	Path pulumi.StringPtrInput
	// Type of the audit device, such as 'file'.
	Type pulumi.StringPtrInput
}

func (AuditState) ElementType() reflect.Type {
	return reflect.TypeOf((*auditState)(nil)).Elem()
}

type auditArgs struct {
	// Human-friendly description of the audit device.
	Description *string `pulumi:"description"`
	// Specifies if the audit device is a local only. Local audit devices are not replicated nor (if a secondary) removed by replication.
	Local *bool `pulumi:"local"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Configuration options to pass to the audit device itself.
	//
	// For a reference of the device types and their options, consult the [Vault documentation.](https://www.vaultproject.io/docs/audit/index.html)
	Options map[string]string `pulumi:"options"`
	// The path to mount the audit device. This defaults to the type.
	Path *string `pulumi:"path"`
	// Type of the audit device, such as 'file'.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Audit resource.
type AuditArgs struct {
	// Human-friendly description of the audit device.
	Description pulumi.StringPtrInput
	// Specifies if the audit device is a local only. Local audit devices are not replicated nor (if a secondary) removed by replication.
	Local pulumi.BoolPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Configuration options to pass to the audit device itself.
	//
	// For a reference of the device types and their options, consult the [Vault documentation.](https://www.vaultproject.io/docs/audit/index.html)
	Options pulumi.StringMapInput
	// The path to mount the audit device. This defaults to the type.
	Path pulumi.StringPtrInput
	// Type of the audit device, such as 'file'.
	Type pulumi.StringInput
}

func (AuditArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*auditArgs)(nil)).Elem()
}

type AuditInput interface {
	pulumi.Input

	ToAuditOutput() AuditOutput
	ToAuditOutputWithContext(ctx context.Context) AuditOutput
}

func (*Audit) ElementType() reflect.Type {
	return reflect.TypeOf((**Audit)(nil)).Elem()
}

func (i *Audit) ToAuditOutput() AuditOutput {
	return i.ToAuditOutputWithContext(context.Background())
}

func (i *Audit) ToAuditOutputWithContext(ctx context.Context) AuditOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditOutput)
}

func (i *Audit) ToOutput(ctx context.Context) pulumix.Output[*Audit] {
	return pulumix.Output[*Audit]{
		OutputState: i.ToAuditOutputWithContext(ctx).OutputState,
	}
}

// AuditArrayInput is an input type that accepts AuditArray and AuditArrayOutput values.
// You can construct a concrete instance of `AuditArrayInput` via:
//
//	AuditArray{ AuditArgs{...} }
type AuditArrayInput interface {
	pulumi.Input

	ToAuditArrayOutput() AuditArrayOutput
	ToAuditArrayOutputWithContext(context.Context) AuditArrayOutput
}

type AuditArray []AuditInput

func (AuditArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Audit)(nil)).Elem()
}

func (i AuditArray) ToAuditArrayOutput() AuditArrayOutput {
	return i.ToAuditArrayOutputWithContext(context.Background())
}

func (i AuditArray) ToAuditArrayOutputWithContext(ctx context.Context) AuditArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditArrayOutput)
}

func (i AuditArray) ToOutput(ctx context.Context) pulumix.Output[[]*Audit] {
	return pulumix.Output[[]*Audit]{
		OutputState: i.ToAuditArrayOutputWithContext(ctx).OutputState,
	}
}

// AuditMapInput is an input type that accepts AuditMap and AuditMapOutput values.
// You can construct a concrete instance of `AuditMapInput` via:
//
//	AuditMap{ "key": AuditArgs{...} }
type AuditMapInput interface {
	pulumi.Input

	ToAuditMapOutput() AuditMapOutput
	ToAuditMapOutputWithContext(context.Context) AuditMapOutput
}

type AuditMap map[string]AuditInput

func (AuditMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Audit)(nil)).Elem()
}

func (i AuditMap) ToAuditMapOutput() AuditMapOutput {
	return i.ToAuditMapOutputWithContext(context.Background())
}

func (i AuditMap) ToAuditMapOutputWithContext(ctx context.Context) AuditMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuditMapOutput)
}

func (i AuditMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Audit] {
	return pulumix.Output[map[string]*Audit]{
		OutputState: i.ToAuditMapOutputWithContext(ctx).OutputState,
	}
}

type AuditOutput struct{ *pulumi.OutputState }

func (AuditOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Audit)(nil)).Elem()
}

func (o AuditOutput) ToAuditOutput() AuditOutput {
	return o
}

func (o AuditOutput) ToAuditOutputWithContext(ctx context.Context) AuditOutput {
	return o
}

func (o AuditOutput) ToOutput(ctx context.Context) pulumix.Output[*Audit] {
	return pulumix.Output[*Audit]{
		OutputState: o.OutputState,
	}
}

// Human-friendly description of the audit device.
func (o AuditOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Audit) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Specifies if the audit device is a local only. Local audit devices are not replicated nor (if a secondary) removed by replication.
func (o AuditOutput) Local() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Audit) pulumi.BoolPtrOutput { return v.Local }).(pulumi.BoolPtrOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o AuditOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Audit) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Configuration options to pass to the audit device itself.
//
// For a reference of the device types and their options, consult the [Vault documentation.](https://www.vaultproject.io/docs/audit/index.html)
func (o AuditOutput) Options() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Audit) pulumi.StringMapOutput { return v.Options }).(pulumi.StringMapOutput)
}

// The path to mount the audit device. This defaults to the type.
func (o AuditOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v *Audit) pulumi.StringOutput { return v.Path }).(pulumi.StringOutput)
}

// Type of the audit device, such as 'file'.
func (o AuditOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Audit) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type AuditArrayOutput struct{ *pulumi.OutputState }

func (AuditArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Audit)(nil)).Elem()
}

func (o AuditArrayOutput) ToAuditArrayOutput() AuditArrayOutput {
	return o
}

func (o AuditArrayOutput) ToAuditArrayOutputWithContext(ctx context.Context) AuditArrayOutput {
	return o
}

func (o AuditArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Audit] {
	return pulumix.Output[[]*Audit]{
		OutputState: o.OutputState,
	}
}

func (o AuditArrayOutput) Index(i pulumi.IntInput) AuditOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Audit {
		return vs[0].([]*Audit)[vs[1].(int)]
	}).(AuditOutput)
}

type AuditMapOutput struct{ *pulumi.OutputState }

func (AuditMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Audit)(nil)).Elem()
}

func (o AuditMapOutput) ToAuditMapOutput() AuditMapOutput {
	return o
}

func (o AuditMapOutput) ToAuditMapOutputWithContext(ctx context.Context) AuditMapOutput {
	return o
}

func (o AuditMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Audit] {
	return pulumix.Output[map[string]*Audit]{
		OutputState: o.OutputState,
	}
}

func (o AuditMapOutput) MapIndex(k pulumi.StringInput) AuditOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Audit {
		return vs[0].(map[string]*Audit)[vs[1].(string)]
	}).(AuditOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuditInput)(nil)).Elem(), &Audit{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuditArrayInput)(nil)).Elem(), AuditArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuditMapInput)(nil)).Elem(), AuditMap{})
	pulumi.RegisterOutputType(AuditOutput{})
	pulumi.RegisterOutputType(AuditArrayOutput{})
	pulumi.RegisterOutputType(AuditMapOutput{})
}
