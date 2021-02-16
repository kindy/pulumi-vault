// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pkisecret

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type SecretBackendSign struct {
	pulumi.CustomResourceState

	// List of alternative names
	AltNames pulumi.StringArrayOutput `pulumi:"altNames"`
	// If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
	AutoRenew pulumi.BoolPtrOutput `pulumi:"autoRenew"`
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// The CA chain
	CaChains pulumi.StringArrayOutput `pulumi:"caChains"`
	// The certificate
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// CN of certificate to create
	CommonName pulumi.StringOutput `pulumi:"commonName"`
	// The CSR
	Csr pulumi.StringOutput `pulumi:"csr"`
	// Flag to exclude CN from SANs
	ExcludeCnFromSans pulumi.BoolPtrOutput `pulumi:"excludeCnFromSans"`
	// The expiration date of the certificate in unix epoch format
	Expiration pulumi.IntOutput `pulumi:"expiration"`
	// The format of data
	Format pulumi.StringPtrOutput `pulumi:"format"`
	// List of alternative IPs
	IpSans pulumi.StringArrayOutput `pulumi:"ipSans"`
	// The issuing CA
	IssuingCa pulumi.StringOutput `pulumi:"issuingCa"`
	// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
	MinSecondsRemaining pulumi.IntPtrOutput `pulumi:"minSecondsRemaining"`
	// Name of the role to create the certificate against
	Name pulumi.StringOutput `pulumi:"name"`
	// List of other SANs
	OtherSans pulumi.StringArrayOutput `pulumi:"otherSans"`
	// The serial
	Serial pulumi.StringOutput `pulumi:"serial"`
	// Time to live
	Ttl pulumi.StringPtrOutput `pulumi:"ttl"`
	// List of alterative URIs
	UriSans pulumi.StringArrayOutput `pulumi:"uriSans"`
}

// NewSecretBackendSign registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendSign(ctx *pulumi.Context,
	name string, args *SecretBackendSignArgs, opts ...pulumi.ResourceOption) (*SecretBackendSign, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	if args.CommonName == nil {
		return nil, errors.New("invalid value for required argument 'CommonName'")
	}
	if args.Csr == nil {
		return nil, errors.New("invalid value for required argument 'Csr'")
	}
	var resource SecretBackendSign
	err := ctx.RegisterResource("vault:pkiSecret/secretBackendSign:SecretBackendSign", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendSign gets an existing SecretBackendSign resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendSign(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendSignState, opts ...pulumi.ResourceOption) (*SecretBackendSign, error) {
	var resource SecretBackendSign
	err := ctx.ReadResource("vault:pkiSecret/secretBackendSign:SecretBackendSign", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendSign resources.
type secretBackendSignState struct {
	// List of alternative names
	AltNames []string `pulumi:"altNames"`
	// If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
	AutoRenew *bool `pulumi:"autoRenew"`
	// The PKI secret backend the resource belongs to.
	Backend *string `pulumi:"backend"`
	// The CA chain
	CaChains []string `pulumi:"caChains"`
	// The certificate
	Certificate *string `pulumi:"certificate"`
	// CN of certificate to create
	CommonName *string `pulumi:"commonName"`
	// The CSR
	Csr *string `pulumi:"csr"`
	// Flag to exclude CN from SANs
	ExcludeCnFromSans *bool `pulumi:"excludeCnFromSans"`
	// The expiration date of the certificate in unix epoch format
	Expiration *int `pulumi:"expiration"`
	// The format of data
	Format *string `pulumi:"format"`
	// List of alternative IPs
	IpSans []string `pulumi:"ipSans"`
	// The issuing CA
	IssuingCa *string `pulumi:"issuingCa"`
	// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
	MinSecondsRemaining *int `pulumi:"minSecondsRemaining"`
	// Name of the role to create the certificate against
	Name *string `pulumi:"name"`
	// List of other SANs
	OtherSans []string `pulumi:"otherSans"`
	// The serial
	Serial *string `pulumi:"serial"`
	// Time to live
	Ttl *string `pulumi:"ttl"`
	// List of alterative URIs
	UriSans []string `pulumi:"uriSans"`
}

type SecretBackendSignState struct {
	// List of alternative names
	AltNames pulumi.StringArrayInput
	// If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
	AutoRenew pulumi.BoolPtrInput
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringPtrInput
	// The CA chain
	CaChains pulumi.StringArrayInput
	// The certificate
	Certificate pulumi.StringPtrInput
	// CN of certificate to create
	CommonName pulumi.StringPtrInput
	// The CSR
	Csr pulumi.StringPtrInput
	// Flag to exclude CN from SANs
	ExcludeCnFromSans pulumi.BoolPtrInput
	// The expiration date of the certificate in unix epoch format
	Expiration pulumi.IntPtrInput
	// The format of data
	Format pulumi.StringPtrInput
	// List of alternative IPs
	IpSans pulumi.StringArrayInput
	// The issuing CA
	IssuingCa pulumi.StringPtrInput
	// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
	MinSecondsRemaining pulumi.IntPtrInput
	// Name of the role to create the certificate against
	Name pulumi.StringPtrInput
	// List of other SANs
	OtherSans pulumi.StringArrayInput
	// The serial
	Serial pulumi.StringPtrInput
	// Time to live
	Ttl pulumi.StringPtrInput
	// List of alterative URIs
	UriSans pulumi.StringArrayInput
}

func (SecretBackendSignState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendSignState)(nil)).Elem()
}

type secretBackendSignArgs struct {
	// List of alternative names
	AltNames []string `pulumi:"altNames"`
	// If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
	AutoRenew *bool `pulumi:"autoRenew"`
	// The PKI secret backend the resource belongs to.
	Backend string `pulumi:"backend"`
	// CN of certificate to create
	CommonName string `pulumi:"commonName"`
	// The CSR
	Csr string `pulumi:"csr"`
	// Flag to exclude CN from SANs
	ExcludeCnFromSans *bool `pulumi:"excludeCnFromSans"`
	// The format of data
	Format *string `pulumi:"format"`
	// List of alternative IPs
	IpSans []string `pulumi:"ipSans"`
	// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
	MinSecondsRemaining *int `pulumi:"minSecondsRemaining"`
	// Name of the role to create the certificate against
	Name *string `pulumi:"name"`
	// List of other SANs
	OtherSans []string `pulumi:"otherSans"`
	// Time to live
	Ttl *string `pulumi:"ttl"`
	// List of alterative URIs
	UriSans []string `pulumi:"uriSans"`
}

// The set of arguments for constructing a SecretBackendSign resource.
type SecretBackendSignArgs struct {
	// List of alternative names
	AltNames pulumi.StringArrayInput
	// If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
	AutoRenew pulumi.BoolPtrInput
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringInput
	// CN of certificate to create
	CommonName pulumi.StringInput
	// The CSR
	Csr pulumi.StringInput
	// Flag to exclude CN from SANs
	ExcludeCnFromSans pulumi.BoolPtrInput
	// The format of data
	Format pulumi.StringPtrInput
	// List of alternative IPs
	IpSans pulumi.StringArrayInput
	// Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
	MinSecondsRemaining pulumi.IntPtrInput
	// Name of the role to create the certificate against
	Name pulumi.StringPtrInput
	// List of other SANs
	OtherSans pulumi.StringArrayInput
	// Time to live
	Ttl pulumi.StringPtrInput
	// List of alterative URIs
	UriSans pulumi.StringArrayInput
}

func (SecretBackendSignArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendSignArgs)(nil)).Elem()
}

type SecretBackendSignInput interface {
	pulumi.Input

	ToSecretBackendSignOutput() SecretBackendSignOutput
	ToSecretBackendSignOutputWithContext(ctx context.Context) SecretBackendSignOutput
}

func (*SecretBackendSign) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretBackendSign)(nil))
}

func (i *SecretBackendSign) ToSecretBackendSignOutput() SecretBackendSignOutput {
	return i.ToSecretBackendSignOutputWithContext(context.Background())
}

func (i *SecretBackendSign) ToSecretBackendSignOutputWithContext(ctx context.Context) SecretBackendSignOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendSignOutput)
}

func (i *SecretBackendSign) ToSecretBackendSignPtrOutput() SecretBackendSignPtrOutput {
	return i.ToSecretBackendSignPtrOutputWithContext(context.Background())
}

func (i *SecretBackendSign) ToSecretBackendSignPtrOutputWithContext(ctx context.Context) SecretBackendSignPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendSignPtrOutput)
}

type SecretBackendSignPtrInput interface {
	pulumi.Input

	ToSecretBackendSignPtrOutput() SecretBackendSignPtrOutput
	ToSecretBackendSignPtrOutputWithContext(ctx context.Context) SecretBackendSignPtrOutput
}

type secretBackendSignPtrType SecretBackendSignArgs

func (*secretBackendSignPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendSign)(nil))
}

func (i *secretBackendSignPtrType) ToSecretBackendSignPtrOutput() SecretBackendSignPtrOutput {
	return i.ToSecretBackendSignPtrOutputWithContext(context.Background())
}

func (i *secretBackendSignPtrType) ToSecretBackendSignPtrOutputWithContext(ctx context.Context) SecretBackendSignPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendSignPtrOutput)
}

// SecretBackendSignArrayInput is an input type that accepts SecretBackendSignArray and SecretBackendSignArrayOutput values.
// You can construct a concrete instance of `SecretBackendSignArrayInput` via:
//
//          SecretBackendSignArray{ SecretBackendSignArgs{...} }
type SecretBackendSignArrayInput interface {
	pulumi.Input

	ToSecretBackendSignArrayOutput() SecretBackendSignArrayOutput
	ToSecretBackendSignArrayOutputWithContext(context.Context) SecretBackendSignArrayOutput
}

type SecretBackendSignArray []SecretBackendSignInput

func (SecretBackendSignArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*SecretBackendSign)(nil))
}

func (i SecretBackendSignArray) ToSecretBackendSignArrayOutput() SecretBackendSignArrayOutput {
	return i.ToSecretBackendSignArrayOutputWithContext(context.Background())
}

func (i SecretBackendSignArray) ToSecretBackendSignArrayOutputWithContext(ctx context.Context) SecretBackendSignArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendSignArrayOutput)
}

// SecretBackendSignMapInput is an input type that accepts SecretBackendSignMap and SecretBackendSignMapOutput values.
// You can construct a concrete instance of `SecretBackendSignMapInput` via:
//
//          SecretBackendSignMap{ "key": SecretBackendSignArgs{...} }
type SecretBackendSignMapInput interface {
	pulumi.Input

	ToSecretBackendSignMapOutput() SecretBackendSignMapOutput
	ToSecretBackendSignMapOutputWithContext(context.Context) SecretBackendSignMapOutput
}

type SecretBackendSignMap map[string]SecretBackendSignInput

func (SecretBackendSignMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*SecretBackendSign)(nil))
}

func (i SecretBackendSignMap) ToSecretBackendSignMapOutput() SecretBackendSignMapOutput {
	return i.ToSecretBackendSignMapOutputWithContext(context.Background())
}

func (i SecretBackendSignMap) ToSecretBackendSignMapOutputWithContext(ctx context.Context) SecretBackendSignMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendSignMapOutput)
}

type SecretBackendSignOutput struct {
	*pulumi.OutputState
}

func (SecretBackendSignOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretBackendSign)(nil))
}

func (o SecretBackendSignOutput) ToSecretBackendSignOutput() SecretBackendSignOutput {
	return o
}

func (o SecretBackendSignOutput) ToSecretBackendSignOutputWithContext(ctx context.Context) SecretBackendSignOutput {
	return o
}

func (o SecretBackendSignOutput) ToSecretBackendSignPtrOutput() SecretBackendSignPtrOutput {
	return o.ToSecretBackendSignPtrOutputWithContext(context.Background())
}

func (o SecretBackendSignOutput) ToSecretBackendSignPtrOutputWithContext(ctx context.Context) SecretBackendSignPtrOutput {
	return o.ApplyT(func(v SecretBackendSign) *SecretBackendSign {
		return &v
	}).(SecretBackendSignPtrOutput)
}

type SecretBackendSignPtrOutput struct {
	*pulumi.OutputState
}

func (SecretBackendSignPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendSign)(nil))
}

func (o SecretBackendSignPtrOutput) ToSecretBackendSignPtrOutput() SecretBackendSignPtrOutput {
	return o
}

func (o SecretBackendSignPtrOutput) ToSecretBackendSignPtrOutputWithContext(ctx context.Context) SecretBackendSignPtrOutput {
	return o
}

type SecretBackendSignArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendSignArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretBackendSign)(nil))
}

func (o SecretBackendSignArrayOutput) ToSecretBackendSignArrayOutput() SecretBackendSignArrayOutput {
	return o
}

func (o SecretBackendSignArrayOutput) ToSecretBackendSignArrayOutputWithContext(ctx context.Context) SecretBackendSignArrayOutput {
	return o
}

func (o SecretBackendSignArrayOutput) Index(i pulumi.IntInput) SecretBackendSignOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecretBackendSign {
		return vs[0].([]SecretBackendSign)[vs[1].(int)]
	}).(SecretBackendSignOutput)
}

type SecretBackendSignMapOutput struct{ *pulumi.OutputState }

func (SecretBackendSignMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]SecretBackendSign)(nil))
}

func (o SecretBackendSignMapOutput) ToSecretBackendSignMapOutput() SecretBackendSignMapOutput {
	return o
}

func (o SecretBackendSignMapOutput) ToSecretBackendSignMapOutputWithContext(ctx context.Context) SecretBackendSignMapOutput {
	return o
}

func (o SecretBackendSignMapOutput) MapIndex(k pulumi.StringInput) SecretBackendSignOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) SecretBackendSign {
		return vs[0].(map[string]SecretBackendSign)[vs[1].(string)]
	}).(SecretBackendSignOutput)
}

func init() {
	pulumi.RegisterOutputType(SecretBackendSignOutput{})
	pulumi.RegisterOutputType(SecretBackendSignPtrOutput{})
	pulumi.RegisterOutputType(SecretBackendSignArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendSignMapOutput{})
}
