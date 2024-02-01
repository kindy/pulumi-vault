// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package consul

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
// ### Creating a standard backend resource:
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/consul"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := consul.NewSecretBackend(ctx, "test", &consul.SecretBackendArgs{
//				Address:     pulumi.String("127.0.0.1:8500"),
//				Description: pulumi.String("Manages the Consul backend"),
//				Path:        pulumi.String("consul"),
//				Token:       pulumi.String("4240861b-ce3d-8530-115a-521ff070dd29"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// ### Creating a backend resource to bootstrap a new Consul instance:
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/consul"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := consul.NewSecretBackend(ctx, "test", &consul.SecretBackendArgs{
//				Address:     pulumi.String("127.0.0.1:8500"),
//				Bootstrap:   pulumi.Bool(true),
//				Description: pulumi.String("Bootstrap the Consul backend"),
//				Path:        pulumi.String("consul"),
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
// Consul secret backends can be imported using the `path`, e.g.
//
// ```sh
//
//	$ pulumi import vault:consul/secretBackend:SecretBackend example consul
//
// ```
type SecretBackend struct {
	pulumi.CustomResourceState

	// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
	Address pulumi.StringOutput `pulumi:"address"`
	// Denotes a backend resource that is used to bootstrap the Consul ACL system. Only one resource may be used to bootstrap.
	Bootstrap pulumi.BoolPtrOutput `pulumi:"bootstrap"`
	// CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
	CaCert pulumi.StringPtrOutput `pulumi:"caCert"`
	// Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if
	// this is set you need to also set client_key.
	ClientCert pulumi.StringPtrOutput `pulumi:"clientCert"`
	// Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set
	// you need to also set client_cert.
	ClientKey pulumi.StringPtrOutput `pulumi:"clientKey"`
	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds pulumi.IntPtrOutput `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount pulumi.BoolPtrOutput `pulumi:"disableRemount"`
	// Specifies if the secret backend is local only.
	Local pulumi.BoolPtrOutput `pulumi:"local"`
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds pulumi.IntPtrOutput `pulumi:"maxLeaseTtlSeconds"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults
	// to `consul`.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// Specifies the URL scheme to use. Defaults to `http`.
	Scheme pulumi.StringPtrOutput `pulumi:"scheme"`
	// Specifies the Consul token to use when managing or issuing new tokens.
	Token pulumi.StringPtrOutput `pulumi:"token"`
}

// NewSecretBackend registers a new resource with the given unique name, arguments, and options.
func NewSecretBackend(ctx *pulumi.Context,
	name string, args *SecretBackendArgs, opts ...pulumi.ResourceOption) (*SecretBackend, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Address == nil {
		return nil, errors.New("invalid value for required argument 'Address'")
	}
	if args.ClientCert != nil {
		args.ClientCert = pulumi.ToSecret(args.ClientCert).(pulumi.StringPtrInput)
	}
	if args.ClientKey != nil {
		args.ClientKey = pulumi.ToSecret(args.ClientKey).(pulumi.StringPtrInput)
	}
	if args.Token != nil {
		args.Token = pulumi.ToSecret(args.Token).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"clientCert",
		"clientKey",
		"token",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretBackend
	err := ctx.RegisterResource("vault:consul/secretBackend:SecretBackend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackend gets an existing SecretBackend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendState, opts ...pulumi.ResourceOption) (*SecretBackend, error) {
	var resource SecretBackend
	err := ctx.ReadResource("vault:consul/secretBackend:SecretBackend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackend resources.
type secretBackendState struct {
	// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
	Address *string `pulumi:"address"`
	// Denotes a backend resource that is used to bootstrap the Consul ACL system. Only one resource may be used to bootstrap.
	Bootstrap *bool `pulumi:"bootstrap"`
	// CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
	CaCert *string `pulumi:"caCert"`
	// Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if
	// this is set you need to also set client_key.
	ClientCert *string `pulumi:"clientCert"`
	// Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set
	// you need to also set client_cert.
	ClientKey *string `pulumi:"clientKey"`
	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description *string `pulumi:"description"`
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount *bool `pulumi:"disableRemount"`
	// Specifies if the secret backend is local only.
	Local *bool `pulumi:"local"`
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults
	// to `consul`.
	Path *string `pulumi:"path"`
	// Specifies the URL scheme to use. Defaults to `http`.
	Scheme *string `pulumi:"scheme"`
	// Specifies the Consul token to use when managing or issuing new tokens.
	Token *string `pulumi:"token"`
}

type SecretBackendState struct {
	// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
	Address pulumi.StringPtrInput
	// Denotes a backend resource that is used to bootstrap the Consul ACL system. Only one resource may be used to bootstrap.
	Bootstrap pulumi.BoolPtrInput
	// CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
	CaCert pulumi.StringPtrInput
	// Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if
	// this is set you need to also set client_key.
	ClientCert pulumi.StringPtrInput
	// Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set
	// you need to also set client_cert.
	ClientKey pulumi.StringPtrInput
	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// A human-friendly description for this backend.
	Description pulumi.StringPtrInput
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount pulumi.BoolPtrInput
	// Specifies if the secret backend is local only.
	Local pulumi.BoolPtrInput
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults
	// to `consul`.
	Path pulumi.StringPtrInput
	// Specifies the URL scheme to use. Defaults to `http`.
	Scheme pulumi.StringPtrInput
	// Specifies the Consul token to use when managing or issuing new tokens.
	Token pulumi.StringPtrInput
}

func (SecretBackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendState)(nil)).Elem()
}

type secretBackendArgs struct {
	// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
	Address string `pulumi:"address"`
	// Denotes a backend resource that is used to bootstrap the Consul ACL system. Only one resource may be used to bootstrap.
	Bootstrap *bool `pulumi:"bootstrap"`
	// CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
	CaCert *string `pulumi:"caCert"`
	// Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if
	// this is set you need to also set client_key.
	ClientCert *string `pulumi:"clientCert"`
	// Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set
	// you need to also set client_cert.
	ClientKey *string `pulumi:"clientKey"`
	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description *string `pulumi:"description"`
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount *bool `pulumi:"disableRemount"`
	// Specifies if the secret backend is local only.
	Local *bool `pulumi:"local"`
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults
	// to `consul`.
	Path *string `pulumi:"path"`
	// Specifies the URL scheme to use. Defaults to `http`.
	Scheme *string `pulumi:"scheme"`
	// Specifies the Consul token to use when managing or issuing new tokens.
	Token *string `pulumi:"token"`
}

// The set of arguments for constructing a SecretBackend resource.
type SecretBackendArgs struct {
	// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
	Address pulumi.StringInput
	// Denotes a backend resource that is used to bootstrap the Consul ACL system. Only one resource may be used to bootstrap.
	Bootstrap pulumi.BoolPtrInput
	// CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
	CaCert pulumi.StringPtrInput
	// Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if
	// this is set you need to also set client_key.
	ClientCert pulumi.StringPtrInput
	// Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set
	// you need to also set client_cert.
	ClientKey pulumi.StringPtrInput
	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// A human-friendly description for this backend.
	Description pulumi.StringPtrInput
	// If set, opts out of mount migration on path updates.
	// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
	DisableRemount pulumi.BoolPtrInput
	// Specifies if the secret backend is local only.
	Local pulumi.BoolPtrInput
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults
	// to `consul`.
	Path pulumi.StringPtrInput
	// Specifies the URL scheme to use. Defaults to `http`.
	Scheme pulumi.StringPtrInput
	// Specifies the Consul token to use when managing or issuing new tokens.
	Token pulumi.StringPtrInput
}

func (SecretBackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendArgs)(nil)).Elem()
}

type SecretBackendInput interface {
	pulumi.Input

	ToSecretBackendOutput() SecretBackendOutput
	ToSecretBackendOutputWithContext(ctx context.Context) SecretBackendOutput
}

func (*SecretBackend) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackend)(nil)).Elem()
}

func (i *SecretBackend) ToSecretBackendOutput() SecretBackendOutput {
	return i.ToSecretBackendOutputWithContext(context.Background())
}

func (i *SecretBackend) ToSecretBackendOutputWithContext(ctx context.Context) SecretBackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendOutput)
}

// SecretBackendArrayInput is an input type that accepts SecretBackendArray and SecretBackendArrayOutput values.
// You can construct a concrete instance of `SecretBackendArrayInput` via:
//
//	SecretBackendArray{ SecretBackendArgs{...} }
type SecretBackendArrayInput interface {
	pulumi.Input

	ToSecretBackendArrayOutput() SecretBackendArrayOutput
	ToSecretBackendArrayOutputWithContext(context.Context) SecretBackendArrayOutput
}

type SecretBackendArray []SecretBackendInput

func (SecretBackendArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackend)(nil)).Elem()
}

func (i SecretBackendArray) ToSecretBackendArrayOutput() SecretBackendArrayOutput {
	return i.ToSecretBackendArrayOutputWithContext(context.Background())
}

func (i SecretBackendArray) ToSecretBackendArrayOutputWithContext(ctx context.Context) SecretBackendArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendArrayOutput)
}

// SecretBackendMapInput is an input type that accepts SecretBackendMap and SecretBackendMapOutput values.
// You can construct a concrete instance of `SecretBackendMapInput` via:
//
//	SecretBackendMap{ "key": SecretBackendArgs{...} }
type SecretBackendMapInput interface {
	pulumi.Input

	ToSecretBackendMapOutput() SecretBackendMapOutput
	ToSecretBackendMapOutputWithContext(context.Context) SecretBackendMapOutput
}

type SecretBackendMap map[string]SecretBackendInput

func (SecretBackendMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackend)(nil)).Elem()
}

func (i SecretBackendMap) ToSecretBackendMapOutput() SecretBackendMapOutput {
	return i.ToSecretBackendMapOutputWithContext(context.Background())
}

func (i SecretBackendMap) ToSecretBackendMapOutputWithContext(ctx context.Context) SecretBackendMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendMapOutput)
}

type SecretBackendOutput struct{ *pulumi.OutputState }

func (SecretBackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackend)(nil)).Elem()
}

func (o SecretBackendOutput) ToSecretBackendOutput() SecretBackendOutput {
	return o
}

func (o SecretBackendOutput) ToSecretBackendOutputWithContext(ctx context.Context) SecretBackendOutput {
	return o
}

// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
func (o SecretBackendOutput) Address() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringOutput { return v.Address }).(pulumi.StringOutput)
}

// Denotes a backend resource that is used to bootstrap the Consul ACL system. Only one resource may be used to bootstrap.
func (o SecretBackendOutput) Bootstrap() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.BoolPtrOutput { return v.Bootstrap }).(pulumi.BoolPtrOutput)
}

// CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
func (o SecretBackendOutput) CaCert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.CaCert }).(pulumi.StringPtrOutput)
}

// Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if
// this is set you need to also set client_key.
func (o SecretBackendOutput) ClientCert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.ClientCert }).(pulumi.StringPtrOutput)
}

// Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set
// you need to also set client_cert.
func (o SecretBackendOutput) ClientKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.ClientKey }).(pulumi.StringPtrOutput)
}

// The default TTL for credentials issued by this backend.
func (o SecretBackendOutput) DefaultLeaseTtlSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.IntPtrOutput { return v.DefaultLeaseTtlSeconds }).(pulumi.IntPtrOutput)
}

// A human-friendly description for this backend.
func (o SecretBackendOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// If set, opts out of mount migration on path updates.
// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
func (o SecretBackendOutput) DisableRemount() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.BoolPtrOutput { return v.DisableRemount }).(pulumi.BoolPtrOutput)
}

// Specifies if the secret backend is local only.
func (o SecretBackendOutput) Local() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.BoolPtrOutput { return v.Local }).(pulumi.BoolPtrOutput)
}

// The maximum TTL that can be requested
// for credentials issued by this backend.
func (o SecretBackendOutput) MaxLeaseTtlSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.IntPtrOutput { return v.MaxLeaseTtlSeconds }).(pulumi.IntPtrOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o SecretBackendOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults
// to `consul`.
func (o SecretBackendOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.Path }).(pulumi.StringPtrOutput)
}

// Specifies the URL scheme to use. Defaults to `http`.
func (o SecretBackendOutput) Scheme() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.Scheme }).(pulumi.StringPtrOutput)
}

// Specifies the Consul token to use when managing or issuing new tokens.
func (o SecretBackendOutput) Token() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackend) pulumi.StringPtrOutput { return v.Token }).(pulumi.StringPtrOutput)
}

type SecretBackendArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackend)(nil)).Elem()
}

func (o SecretBackendArrayOutput) ToSecretBackendArrayOutput() SecretBackendArrayOutput {
	return o
}

func (o SecretBackendArrayOutput) ToSecretBackendArrayOutputWithContext(ctx context.Context) SecretBackendArrayOutput {
	return o
}

func (o SecretBackendArrayOutput) Index(i pulumi.IntInput) SecretBackendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackend {
		return vs[0].([]*SecretBackend)[vs[1].(int)]
	}).(SecretBackendOutput)
}

type SecretBackendMapOutput struct{ *pulumi.OutputState }

func (SecretBackendMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackend)(nil)).Elem()
}

func (o SecretBackendMapOutput) ToSecretBackendMapOutput() SecretBackendMapOutput {
	return o
}

func (o SecretBackendMapOutput) ToSecretBackendMapOutputWithContext(ctx context.Context) SecretBackendMapOutput {
	return o
}

func (o SecretBackendMapOutput) MapIndex(k pulumi.StringInput) SecretBackendOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackend {
		return vs[0].(map[string]*SecretBackend)[vs[1].(string)]
	}).(SecretBackendOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendInput)(nil)).Elem(), &SecretBackend{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendArrayInput)(nil)).Elem(), SecretBackendArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendMapInput)(nil)).Elem(), SecretBackendMap{})
	pulumi.RegisterOutputType(SecretBackendOutput{})
	pulumi.RegisterOutputType(SecretBackendArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendMapOutput{})
}
