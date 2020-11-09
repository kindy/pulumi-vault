// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rabbitmq

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type SecretBackendRoleVhost struct {
	Configure string `pulumi:"configure"`
	Host      string `pulumi:"host"`
	Read      string `pulumi:"read"`
	Write     string `pulumi:"write"`
}

// SecretBackendRoleVhostInput is an input type that accepts SecretBackendRoleVhostArgs and SecretBackendRoleVhostOutput values.
// You can construct a concrete instance of `SecretBackendRoleVhostInput` via:
//
//          SecretBackendRoleVhostArgs{...}
type SecretBackendRoleVhostInput interface {
	pulumi.Input

	ToSecretBackendRoleVhostOutput() SecretBackendRoleVhostOutput
	ToSecretBackendRoleVhostOutputWithContext(context.Context) SecretBackendRoleVhostOutput
}

type SecretBackendRoleVhostArgs struct {
	Configure pulumi.StringInput `pulumi:"configure"`
	Host      pulumi.StringInput `pulumi:"host"`
	Read      pulumi.StringInput `pulumi:"read"`
	Write     pulumi.StringInput `pulumi:"write"`
}

func (SecretBackendRoleVhostArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretBackendRoleVhost)(nil)).Elem()
}

func (i SecretBackendRoleVhostArgs) ToSecretBackendRoleVhostOutput() SecretBackendRoleVhostOutput {
	return i.ToSecretBackendRoleVhostOutputWithContext(context.Background())
}

func (i SecretBackendRoleVhostArgs) ToSecretBackendRoleVhostOutputWithContext(ctx context.Context) SecretBackendRoleVhostOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRoleVhostOutput)
}

// SecretBackendRoleVhostArrayInput is an input type that accepts SecretBackendRoleVhostArray and SecretBackendRoleVhostArrayOutput values.
// You can construct a concrete instance of `SecretBackendRoleVhostArrayInput` via:
//
//          SecretBackendRoleVhostArray{ SecretBackendRoleVhostArgs{...} }
type SecretBackendRoleVhostArrayInput interface {
	pulumi.Input

	ToSecretBackendRoleVhostArrayOutput() SecretBackendRoleVhostArrayOutput
	ToSecretBackendRoleVhostArrayOutputWithContext(context.Context) SecretBackendRoleVhostArrayOutput
}

type SecretBackendRoleVhostArray []SecretBackendRoleVhostInput

func (SecretBackendRoleVhostArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretBackendRoleVhost)(nil)).Elem()
}

func (i SecretBackendRoleVhostArray) ToSecretBackendRoleVhostArrayOutput() SecretBackendRoleVhostArrayOutput {
	return i.ToSecretBackendRoleVhostArrayOutputWithContext(context.Background())
}

func (i SecretBackendRoleVhostArray) ToSecretBackendRoleVhostArrayOutputWithContext(ctx context.Context) SecretBackendRoleVhostArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRoleVhostArrayOutput)
}

type SecretBackendRoleVhostOutput struct{ *pulumi.OutputState }

func (SecretBackendRoleVhostOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretBackendRoleVhost)(nil)).Elem()
}

func (o SecretBackendRoleVhostOutput) ToSecretBackendRoleVhostOutput() SecretBackendRoleVhostOutput {
	return o
}

func (o SecretBackendRoleVhostOutput) ToSecretBackendRoleVhostOutputWithContext(ctx context.Context) SecretBackendRoleVhostOutput {
	return o
}

func (o SecretBackendRoleVhostOutput) Configure() pulumi.StringOutput {
	return o.ApplyT(func(v SecretBackendRoleVhost) string { return v.Configure }).(pulumi.StringOutput)
}

func (o SecretBackendRoleVhostOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v SecretBackendRoleVhost) string { return v.Host }).(pulumi.StringOutput)
}

func (o SecretBackendRoleVhostOutput) Read() pulumi.StringOutput {
	return o.ApplyT(func(v SecretBackendRoleVhost) string { return v.Read }).(pulumi.StringOutput)
}

func (o SecretBackendRoleVhostOutput) Write() pulumi.StringOutput {
	return o.ApplyT(func(v SecretBackendRoleVhost) string { return v.Write }).(pulumi.StringOutput)
}

type SecretBackendRoleVhostArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendRoleVhostArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretBackendRoleVhost)(nil)).Elem()
}

func (o SecretBackendRoleVhostArrayOutput) ToSecretBackendRoleVhostArrayOutput() SecretBackendRoleVhostArrayOutput {
	return o
}

func (o SecretBackendRoleVhostArrayOutput) ToSecretBackendRoleVhostArrayOutputWithContext(ctx context.Context) SecretBackendRoleVhostArrayOutput {
	return o
}

func (o SecretBackendRoleVhostArrayOutput) Index(i pulumi.IntInput) SecretBackendRoleVhostOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecretBackendRoleVhost {
		return vs[0].([]SecretBackendRoleVhost)[vs[1].(int)]
	}).(SecretBackendRoleVhostOutput)
}

func init() {
	pulumi.RegisterOutputType(SecretBackendRoleVhostOutput{})
	pulumi.RegisterOutputType(SecretBackendRoleVhostArrayOutput{})
}