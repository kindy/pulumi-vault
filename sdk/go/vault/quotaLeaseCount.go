// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manage lease count quotas which enforce the number of leases that can be created.
// A lease count quota can be created at the root level or defined on a namespace or mount by
// specifying a path when creating the quota.
//
// See [Vault's Documentation](https://www.vaultproject.io/docs/enterprise/lease-count-quotas) for more
// information.
//
// **Note** this feature is available only with Vault Enterprise.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v4/go/vault"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := vault.NewQuotaLeaseCount(ctx, "global", &vault.QuotaLeaseCountArgs{
// 			MaxLeases: pulumi.Int(100),
// 			Path:      pulumi.String(""),
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
// Lease count quotas can be imported using their names
//
// ```sh
//  $ pulumi import vault:index/quotaLeaseCount:QuotaLeaseCount global global
// ```
type QuotaLeaseCount struct {
	pulumi.CustomResourceState

	// The maximum number of leases to be allowed by the quota
	// rule. The `maxLeases` must be positive.
	MaxLeases pulumi.IntOutput `pulumi:"maxLeases"`
	// Name of the rate limit quota
	Name pulumi.StringOutput `pulumi:"name"`
	// Path of the mount or namespace to apply the quota. A blank path configures a
	// global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
	// `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
	// Updating this field on an existing quota can have "moving" effects. For example, updating
	// `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
	// a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
	Path pulumi.StringPtrOutput `pulumi:"path"`
}

// NewQuotaLeaseCount registers a new resource with the given unique name, arguments, and options.
func NewQuotaLeaseCount(ctx *pulumi.Context,
	name string, args *QuotaLeaseCountArgs, opts ...pulumi.ResourceOption) (*QuotaLeaseCount, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MaxLeases == nil {
		return nil, errors.New("invalid value for required argument 'MaxLeases'")
	}
	var resource QuotaLeaseCount
	err := ctx.RegisterResource("vault:index/quotaLeaseCount:QuotaLeaseCount", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetQuotaLeaseCount gets an existing QuotaLeaseCount resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetQuotaLeaseCount(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *QuotaLeaseCountState, opts ...pulumi.ResourceOption) (*QuotaLeaseCount, error) {
	var resource QuotaLeaseCount
	err := ctx.ReadResource("vault:index/quotaLeaseCount:QuotaLeaseCount", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering QuotaLeaseCount resources.
type quotaLeaseCountState struct {
	// The maximum number of leases to be allowed by the quota
	// rule. The `maxLeases` must be positive.
	MaxLeases *int `pulumi:"maxLeases"`
	// Name of the rate limit quota
	Name *string `pulumi:"name"`
	// Path of the mount or namespace to apply the quota. A blank path configures a
	// global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
	// `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
	// Updating this field on an existing quota can have "moving" effects. For example, updating
	// `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
	// a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
	Path *string `pulumi:"path"`
}

type QuotaLeaseCountState struct {
	// The maximum number of leases to be allowed by the quota
	// rule. The `maxLeases` must be positive.
	MaxLeases pulumi.IntPtrInput
	// Name of the rate limit quota
	Name pulumi.StringPtrInput
	// Path of the mount or namespace to apply the quota. A blank path configures a
	// global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
	// `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
	// Updating this field on an existing quota can have "moving" effects. For example, updating
	// `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
	// a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
	Path pulumi.StringPtrInput
}

func (QuotaLeaseCountState) ElementType() reflect.Type {
	return reflect.TypeOf((*quotaLeaseCountState)(nil)).Elem()
}

type quotaLeaseCountArgs struct {
	// The maximum number of leases to be allowed by the quota
	// rule. The `maxLeases` must be positive.
	MaxLeases int `pulumi:"maxLeases"`
	// Name of the rate limit quota
	Name *string `pulumi:"name"`
	// Path of the mount or namespace to apply the quota. A blank path configures a
	// global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
	// `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
	// Updating this field on an existing quota can have "moving" effects. For example, updating
	// `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
	// a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
	Path *string `pulumi:"path"`
}

// The set of arguments for constructing a QuotaLeaseCount resource.
type QuotaLeaseCountArgs struct {
	// The maximum number of leases to be allowed by the quota
	// rule. The `maxLeases` must be positive.
	MaxLeases pulumi.IntInput
	// Name of the rate limit quota
	Name pulumi.StringPtrInput
	// Path of the mount or namespace to apply the quota. A blank path configures a
	// global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
	// `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
	// Updating this field on an existing quota can have "moving" effects. For example, updating
	// `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
	// a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
	Path pulumi.StringPtrInput
}

func (QuotaLeaseCountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*quotaLeaseCountArgs)(nil)).Elem()
}

type QuotaLeaseCountInput interface {
	pulumi.Input

	ToQuotaLeaseCountOutput() QuotaLeaseCountOutput
	ToQuotaLeaseCountOutputWithContext(ctx context.Context) QuotaLeaseCountOutput
}

func (*QuotaLeaseCount) ElementType() reflect.Type {
	return reflect.TypeOf((*QuotaLeaseCount)(nil))
}

func (i *QuotaLeaseCount) ToQuotaLeaseCountOutput() QuotaLeaseCountOutput {
	return i.ToQuotaLeaseCountOutputWithContext(context.Background())
}

func (i *QuotaLeaseCount) ToQuotaLeaseCountOutputWithContext(ctx context.Context) QuotaLeaseCountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaLeaseCountOutput)
}

func (i *QuotaLeaseCount) ToQuotaLeaseCountPtrOutput() QuotaLeaseCountPtrOutput {
	return i.ToQuotaLeaseCountPtrOutputWithContext(context.Background())
}

func (i *QuotaLeaseCount) ToQuotaLeaseCountPtrOutputWithContext(ctx context.Context) QuotaLeaseCountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaLeaseCountPtrOutput)
}

type QuotaLeaseCountPtrInput interface {
	pulumi.Input

	ToQuotaLeaseCountPtrOutput() QuotaLeaseCountPtrOutput
	ToQuotaLeaseCountPtrOutputWithContext(ctx context.Context) QuotaLeaseCountPtrOutput
}

type quotaLeaseCountPtrType QuotaLeaseCountArgs

func (*quotaLeaseCountPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QuotaLeaseCount)(nil))
}

func (i *quotaLeaseCountPtrType) ToQuotaLeaseCountPtrOutput() QuotaLeaseCountPtrOutput {
	return i.ToQuotaLeaseCountPtrOutputWithContext(context.Background())
}

func (i *quotaLeaseCountPtrType) ToQuotaLeaseCountPtrOutputWithContext(ctx context.Context) QuotaLeaseCountPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaLeaseCountPtrOutput)
}

// QuotaLeaseCountArrayInput is an input type that accepts QuotaLeaseCountArray and QuotaLeaseCountArrayOutput values.
// You can construct a concrete instance of `QuotaLeaseCountArrayInput` via:
//
//          QuotaLeaseCountArray{ QuotaLeaseCountArgs{...} }
type QuotaLeaseCountArrayInput interface {
	pulumi.Input

	ToQuotaLeaseCountArrayOutput() QuotaLeaseCountArrayOutput
	ToQuotaLeaseCountArrayOutputWithContext(context.Context) QuotaLeaseCountArrayOutput
}

type QuotaLeaseCountArray []QuotaLeaseCountInput

func (QuotaLeaseCountArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*QuotaLeaseCount)(nil))
}

func (i QuotaLeaseCountArray) ToQuotaLeaseCountArrayOutput() QuotaLeaseCountArrayOutput {
	return i.ToQuotaLeaseCountArrayOutputWithContext(context.Background())
}

func (i QuotaLeaseCountArray) ToQuotaLeaseCountArrayOutputWithContext(ctx context.Context) QuotaLeaseCountArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaLeaseCountArrayOutput)
}

// QuotaLeaseCountMapInput is an input type that accepts QuotaLeaseCountMap and QuotaLeaseCountMapOutput values.
// You can construct a concrete instance of `QuotaLeaseCountMapInput` via:
//
//          QuotaLeaseCountMap{ "key": QuotaLeaseCountArgs{...} }
type QuotaLeaseCountMapInput interface {
	pulumi.Input

	ToQuotaLeaseCountMapOutput() QuotaLeaseCountMapOutput
	ToQuotaLeaseCountMapOutputWithContext(context.Context) QuotaLeaseCountMapOutput
}

type QuotaLeaseCountMap map[string]QuotaLeaseCountInput

func (QuotaLeaseCountMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*QuotaLeaseCount)(nil))
}

func (i QuotaLeaseCountMap) ToQuotaLeaseCountMapOutput() QuotaLeaseCountMapOutput {
	return i.ToQuotaLeaseCountMapOutputWithContext(context.Background())
}

func (i QuotaLeaseCountMap) ToQuotaLeaseCountMapOutputWithContext(ctx context.Context) QuotaLeaseCountMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuotaLeaseCountMapOutput)
}

type QuotaLeaseCountOutput struct {
	*pulumi.OutputState
}

func (QuotaLeaseCountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QuotaLeaseCount)(nil))
}

func (o QuotaLeaseCountOutput) ToQuotaLeaseCountOutput() QuotaLeaseCountOutput {
	return o
}

func (o QuotaLeaseCountOutput) ToQuotaLeaseCountOutputWithContext(ctx context.Context) QuotaLeaseCountOutput {
	return o
}

func (o QuotaLeaseCountOutput) ToQuotaLeaseCountPtrOutput() QuotaLeaseCountPtrOutput {
	return o.ToQuotaLeaseCountPtrOutputWithContext(context.Background())
}

func (o QuotaLeaseCountOutput) ToQuotaLeaseCountPtrOutputWithContext(ctx context.Context) QuotaLeaseCountPtrOutput {
	return o.ApplyT(func(v QuotaLeaseCount) *QuotaLeaseCount {
		return &v
	}).(QuotaLeaseCountPtrOutput)
}

type QuotaLeaseCountPtrOutput struct {
	*pulumi.OutputState
}

func (QuotaLeaseCountPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QuotaLeaseCount)(nil))
}

func (o QuotaLeaseCountPtrOutput) ToQuotaLeaseCountPtrOutput() QuotaLeaseCountPtrOutput {
	return o
}

func (o QuotaLeaseCountPtrOutput) ToQuotaLeaseCountPtrOutputWithContext(ctx context.Context) QuotaLeaseCountPtrOutput {
	return o
}

type QuotaLeaseCountArrayOutput struct{ *pulumi.OutputState }

func (QuotaLeaseCountArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]QuotaLeaseCount)(nil))
}

func (o QuotaLeaseCountArrayOutput) ToQuotaLeaseCountArrayOutput() QuotaLeaseCountArrayOutput {
	return o
}

func (o QuotaLeaseCountArrayOutput) ToQuotaLeaseCountArrayOutputWithContext(ctx context.Context) QuotaLeaseCountArrayOutput {
	return o
}

func (o QuotaLeaseCountArrayOutput) Index(i pulumi.IntInput) QuotaLeaseCountOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) QuotaLeaseCount {
		return vs[0].([]QuotaLeaseCount)[vs[1].(int)]
	}).(QuotaLeaseCountOutput)
}

type QuotaLeaseCountMapOutput struct{ *pulumi.OutputState }

func (QuotaLeaseCountMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]QuotaLeaseCount)(nil))
}

func (o QuotaLeaseCountMapOutput) ToQuotaLeaseCountMapOutput() QuotaLeaseCountMapOutput {
	return o
}

func (o QuotaLeaseCountMapOutput) ToQuotaLeaseCountMapOutputWithContext(ctx context.Context) QuotaLeaseCountMapOutput {
	return o
}

func (o QuotaLeaseCountMapOutput) MapIndex(k pulumi.StringInput) QuotaLeaseCountOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) QuotaLeaseCount {
		return vs[0].(map[string]QuotaLeaseCount)[vs[1].(string)]
	}).(QuotaLeaseCountOutput)
}

func init() {
	pulumi.RegisterOutputType(QuotaLeaseCountOutput{})
	pulumi.RegisterOutputType(QuotaLeaseCountPtrOutput{})
	pulumi.RegisterOutputType(QuotaLeaseCountArrayOutput{})
	pulumi.RegisterOutputType(QuotaLeaseCountMapOutput{})
}
