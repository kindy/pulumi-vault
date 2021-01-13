// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package consul

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// ## Import
//
// Consul secret backends can be imported using the `path`, e.g.
//
// ```sh
//  $ pulumi import vault:consul/secretBackend:SecretBackend example consul
// ```
type SecretBackend struct {
	pulumi.CustomResourceState

	// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
	Address pulumi.StringOutput `pulumi:"address"`
	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds pulumi.IntPtrOutput `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds pulumi.IntPtrOutput `pulumi:"maxLeaseTtlSeconds"`
	// The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults to `consul`.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// Specifies the URL scheme to use. Defaults to `http`.
	Scheme pulumi.StringPtrOutput `pulumi:"scheme"`
	// The Consul management token this backend should use to issue new tokens.
	Token pulumi.StringOutput `pulumi:"token"`
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
	if args.Token == nil {
		return nil, errors.New("invalid value for required argument 'Token'")
	}
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
	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description *string `pulumi:"description"`
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults to `consul`.
	Path *string `pulumi:"path"`
	// Specifies the URL scheme to use. Defaults to `http`.
	Scheme *string `pulumi:"scheme"`
	// The Consul management token this backend should use to issue new tokens.
	Token *string `pulumi:"token"`
}

type SecretBackendState struct {
	// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
	Address pulumi.StringPtrInput
	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// A human-friendly description for this backend.
	Description pulumi.StringPtrInput
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults to `consul`.
	Path pulumi.StringPtrInput
	// Specifies the URL scheme to use. Defaults to `http`.
	Scheme pulumi.StringPtrInput
	// The Consul management token this backend should use to issue new tokens.
	Token pulumi.StringPtrInput
}

func (SecretBackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendState)(nil)).Elem()
}

type secretBackendArgs struct {
	// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
	Address string `pulumi:"address"`
	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds *int `pulumi:"defaultLeaseTtlSeconds"`
	// A human-friendly description for this backend.
	Description *string `pulumi:"description"`
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults to `consul`.
	Path *string `pulumi:"path"`
	// Specifies the URL scheme to use. Defaults to `http`.
	Scheme *string `pulumi:"scheme"`
	// The Consul management token this backend should use to issue new tokens.
	Token string `pulumi:"token"`
}

// The set of arguments for constructing a SecretBackend resource.
type SecretBackendArgs struct {
	// Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
	Address pulumi.StringInput
	// The default TTL for credentials issued by this backend.
	DefaultLeaseTtlSeconds pulumi.IntPtrInput
	// A human-friendly description for this backend.
	Description pulumi.StringPtrInput
	// The maximum TTL that can be requested
	// for credentials issued by this backend.
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults to `consul`.
	Path pulumi.StringPtrInput
	// Specifies the URL scheme to use. Defaults to `http`.
	Scheme pulumi.StringPtrInput
	// The Consul management token this backend should use to issue new tokens.
	Token pulumi.StringInput
}

func (SecretBackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendArgs)(nil)).Elem()
}

type SecretBackendInput interface {
	pulumi.Input

	ToSecretBackendOutput() SecretBackendOutput
	ToSecretBackendOutputWithContext(ctx context.Context) SecretBackendOutput
}

func (SecretBackend) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretBackend)(nil)).Elem()
}

func (i SecretBackend) ToSecretBackendOutput() SecretBackendOutput {
	return i.ToSecretBackendOutputWithContext(context.Background())
}

func (i SecretBackend) ToSecretBackendOutputWithContext(ctx context.Context) SecretBackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendOutput)
}

type SecretBackendOutput struct {
	*pulumi.OutputState
}

func (SecretBackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretBackendOutput)(nil)).Elem()
}

func (o SecretBackendOutput) ToSecretBackendOutput() SecretBackendOutput {
	return o
}

func (o SecretBackendOutput) ToSecretBackendOutputWithContext(ctx context.Context) SecretBackendOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SecretBackendOutput{})
}
