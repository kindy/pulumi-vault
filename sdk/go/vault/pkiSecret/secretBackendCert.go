// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pkiSecret

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type SecretBackendCert struct {
	pulumi.CustomResourceState

	// List of alternative names
	AltNames pulumi.StringArrayOutput `pulumi:"altNames"`
	// If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
	AutoRenew pulumi.BoolPtrOutput `pulumi:"autoRenew"`
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// The CA chain
	CaChain pulumi.StringOutput `pulumi:"caChain"`
	// The certificate
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// CN of certificate to create
	CommonName pulumi.StringOutput `pulumi:"commonName"`
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
	// The private key
	PrivateKey pulumi.StringOutput `pulumi:"privateKey"`
	// The private key format
	PrivateKeyFormat pulumi.StringPtrOutput `pulumi:"privateKeyFormat"`
	// The private key type
	PrivateKeyType pulumi.StringOutput `pulumi:"privateKeyType"`
	// The serial number
	SerialNumber pulumi.StringOutput `pulumi:"serialNumber"`
	// Time to live
	Ttl pulumi.StringPtrOutput `pulumi:"ttl"`
	// List of alternative URIs
	UriSans pulumi.StringArrayOutput `pulumi:"uriSans"`
}

// NewSecretBackendCert registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendCert(ctx *pulumi.Context,
	name string, args *SecretBackendCertArgs, opts ...pulumi.ResourceOption) (*SecretBackendCert, error) {
	if args == nil || args.Backend == nil {
		return nil, errors.New("missing required argument 'Backend'")
	}
	if args == nil || args.CommonName == nil {
		return nil, errors.New("missing required argument 'CommonName'")
	}
	if args == nil {
		args = &SecretBackendCertArgs{}
	}
	var resource SecretBackendCert
	err := ctx.RegisterResource("vault:pkiSecret/secretBackendCert:SecretBackendCert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendCert gets an existing SecretBackendCert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendCert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendCertState, opts ...pulumi.ResourceOption) (*SecretBackendCert, error) {
	var resource SecretBackendCert
	err := ctx.ReadResource("vault:pkiSecret/secretBackendCert:SecretBackendCert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendCert resources.
type secretBackendCertState struct {
	// List of alternative names
	AltNames []string `pulumi:"altNames"`
	// If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
	AutoRenew *bool `pulumi:"autoRenew"`
	// The PKI secret backend the resource belongs to.
	Backend *string `pulumi:"backend"`
	// The CA chain
	CaChain *string `pulumi:"caChain"`
	// The certificate
	Certificate *string `pulumi:"certificate"`
	// CN of certificate to create
	CommonName *string `pulumi:"commonName"`
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
	// The private key
	PrivateKey *string `pulumi:"privateKey"`
	// The private key format
	PrivateKeyFormat *string `pulumi:"privateKeyFormat"`
	// The private key type
	PrivateKeyType *string `pulumi:"privateKeyType"`
	// The serial number
	SerialNumber *string `pulumi:"serialNumber"`
	// Time to live
	Ttl *string `pulumi:"ttl"`
	// List of alternative URIs
	UriSans []string `pulumi:"uriSans"`
}

type SecretBackendCertState struct {
	// List of alternative names
	AltNames pulumi.StringArrayInput
	// If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
	AutoRenew pulumi.BoolPtrInput
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringPtrInput
	// The CA chain
	CaChain pulumi.StringPtrInput
	// The certificate
	Certificate pulumi.StringPtrInput
	// CN of certificate to create
	CommonName pulumi.StringPtrInput
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
	// The private key
	PrivateKey pulumi.StringPtrInput
	// The private key format
	PrivateKeyFormat pulumi.StringPtrInput
	// The private key type
	PrivateKeyType pulumi.StringPtrInput
	// The serial number
	SerialNumber pulumi.StringPtrInput
	// Time to live
	Ttl pulumi.StringPtrInput
	// List of alternative URIs
	UriSans pulumi.StringArrayInput
}

func (SecretBackendCertState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendCertState)(nil)).Elem()
}

type secretBackendCertArgs struct {
	// List of alternative names
	AltNames []string `pulumi:"altNames"`
	// If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
	AutoRenew *bool `pulumi:"autoRenew"`
	// The PKI secret backend the resource belongs to.
	Backend string `pulumi:"backend"`
	// CN of certificate to create
	CommonName string `pulumi:"commonName"`
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
	// The private key format
	PrivateKeyFormat *string `pulumi:"privateKeyFormat"`
	// Time to live
	Ttl *string `pulumi:"ttl"`
	// List of alternative URIs
	UriSans []string `pulumi:"uriSans"`
}

// The set of arguments for constructing a SecretBackendCert resource.
type SecretBackendCertArgs struct {
	// List of alternative names
	AltNames pulumi.StringArrayInput
	// If set to `true`, certs will be renewed if the expiration is within `minSecondsRemaining`. Default `false`
	AutoRenew pulumi.BoolPtrInput
	// The PKI secret backend the resource belongs to.
	Backend pulumi.StringInput
	// CN of certificate to create
	CommonName pulumi.StringInput
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
	// The private key format
	PrivateKeyFormat pulumi.StringPtrInput
	// Time to live
	Ttl pulumi.StringPtrInput
	// List of alternative URIs
	UriSans pulumi.StringArrayInput
}

func (SecretBackendCertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendCertArgs)(nil)).Elem()
}
