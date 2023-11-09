// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kubernetes

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Reads the Role of an Kubernetes from a Vault server. See the [Vault
// documentation](https://www.vaultproject.io/api-docs/auth/kubernetes#read-config) for more
// information.
func LookupAuthBackendConfig(ctx *pulumi.Context, args *LookupAuthBackendConfigArgs, opts ...pulumi.InvokeOption) (*LookupAuthBackendConfigResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAuthBackendConfigResult
	err := ctx.Invoke("vault:kubernetes/getAuthBackendConfig:getAuthBackendConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAuthBackendConfig.
type LookupAuthBackendConfigArgs struct {
	// The unique name for the Kubernetes backend the config to
	// retrieve Role attributes for resides in. Defaults to "kubernetes".
	Backend              *string `pulumi:"backend"`
	DisableIssValidation *bool   `pulumi:"disableIssValidation"`
	DisableLocalCaJwt    *bool   `pulumi:"disableLocalCaJwt"`
	// Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
	Issuer *string `pulumi:"issuer"`
	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	KubernetesCaCert *string `pulumi:"kubernetesCaCert"`
	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	KubernetesHost *string `pulumi:"kubernetesHost"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	PemKeys []string `pulumi:"pemKeys"`
}

// A collection of values returned by getAuthBackendConfig.
type LookupAuthBackendConfigResult struct {
	Backend              *string `pulumi:"backend"`
	DisableIssValidation bool    `pulumi:"disableIssValidation"`
	DisableLocalCaJwt    bool    `pulumi:"disableLocalCaJwt"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
	Issuer string `pulumi:"issuer"`
	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	KubernetesCaCert string `pulumi:"kubernetesCaCert"`
	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	KubernetesHost string  `pulumi:"kubernetesHost"`
	Namespace      *string `pulumi:"namespace"`
	// Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	PemKeys []string `pulumi:"pemKeys"`
}

func LookupAuthBackendConfigOutput(ctx *pulumi.Context, args LookupAuthBackendConfigOutputArgs, opts ...pulumi.InvokeOption) LookupAuthBackendConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAuthBackendConfigResult, error) {
			args := v.(LookupAuthBackendConfigArgs)
			r, err := LookupAuthBackendConfig(ctx, &args, opts...)
			var s LookupAuthBackendConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAuthBackendConfigResultOutput)
}

// A collection of arguments for invoking getAuthBackendConfig.
type LookupAuthBackendConfigOutputArgs struct {
	// The unique name for the Kubernetes backend the config to
	// retrieve Role attributes for resides in. Defaults to "kubernetes".
	Backend              pulumi.StringPtrInput `pulumi:"backend"`
	DisableIssValidation pulumi.BoolPtrInput   `pulumi:"disableIssValidation"`
	DisableLocalCaJwt    pulumi.BoolPtrInput   `pulumi:"disableLocalCaJwt"`
	// Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
	Issuer pulumi.StringPtrInput `pulumi:"issuer"`
	// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
	KubernetesCaCert pulumi.StringPtrInput `pulumi:"kubernetesCaCert"`
	// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
	KubernetesHost pulumi.StringPtrInput `pulumi:"kubernetesHost"`
	// The namespace of the target resource.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured namespace.
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
	// Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
	PemKeys pulumi.StringArrayInput `pulumi:"pemKeys"`
}

func (LookupAuthBackendConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthBackendConfigArgs)(nil)).Elem()
}

// A collection of values returned by getAuthBackendConfig.
type LookupAuthBackendConfigResultOutput struct{ *pulumi.OutputState }

func (LookupAuthBackendConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAuthBackendConfigResult)(nil)).Elem()
}

func (o LookupAuthBackendConfigResultOutput) ToLookupAuthBackendConfigResultOutput() LookupAuthBackendConfigResultOutput {
	return o
}

func (o LookupAuthBackendConfigResultOutput) ToLookupAuthBackendConfigResultOutputWithContext(ctx context.Context) LookupAuthBackendConfigResultOutput {
	return o
}

func (o LookupAuthBackendConfigResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAuthBackendConfigResult] {
	return pulumix.Output[LookupAuthBackendConfigResult]{
		OutputState: o.OutputState,
	}
}

func (o LookupAuthBackendConfigResultOutput) Backend() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthBackendConfigResult) *string { return v.Backend }).(pulumi.StringPtrOutput)
}

func (o LookupAuthBackendConfigResultOutput) DisableIssValidation() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAuthBackendConfigResult) bool { return v.DisableIssValidation }).(pulumi.BoolOutput)
}

func (o LookupAuthBackendConfigResultOutput) DisableLocalCaJwt() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupAuthBackendConfigResult) bool { return v.DisableLocalCaJwt }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupAuthBackendConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthBackendConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

// Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
func (o LookupAuthBackendConfigResultOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthBackendConfigResult) string { return v.Issuer }).(pulumi.StringOutput)
}

// PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
func (o LookupAuthBackendConfigResultOutput) KubernetesCaCert() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthBackendConfigResult) string { return v.KubernetesCaCert }).(pulumi.StringOutput)
}

// Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
func (o LookupAuthBackendConfigResultOutput) KubernetesHost() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAuthBackendConfigResult) string { return v.KubernetesHost }).(pulumi.StringOutput)
}

func (o LookupAuthBackendConfigResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAuthBackendConfigResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
func (o LookupAuthBackendConfigResultOutput) PemKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAuthBackendConfigResult) []string { return v.PemKeys }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAuthBackendConfigResultOutput{})
}
