// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package ssh

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to manage roles in an SSH secret backend
// [SSH secret backend within Vault](https://www.vaultproject.io/docs/secrets/ssh/index.html).
//
// ## Example Usage
//
// <!--Start PulumiCodeChooser -->
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/ssh"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := vault.NewMount(ctx, "example", &vault.MountArgs{
//				Type: pulumi.String("ssh"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ssh.NewSecretBackendRole(ctx, "foo", &ssh.SecretBackendRoleArgs{
//				Backend:               example.Path,
//				KeyType:               pulumi.String("ca"),
//				AllowUserCertificates: pulumi.Bool(true),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = ssh.NewSecretBackendRole(ctx, "bar", &ssh.SecretBackendRoleArgs{
//				Backend:      example.Path,
//				KeyType:      pulumi.String("otp"),
//				DefaultUser:  pulumi.String("default"),
//				AllowedUsers: pulumi.String("default,baz"),
//				CidrList:     pulumi.String("0.0.0.0/0"),
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
// <!--End PulumiCodeChooser -->
//
// ## Import
//
// SSH secret backend roles can be imported using the `path`, e.g.
//
// ```sh
// $ pulumi import vault:ssh/secretBackendRole:SecretBackendRole foo ssh/roles/my-role
// ```
type SecretBackendRole struct {
	pulumi.CustomResourceState

	// When supplied, this value specifies a signing algorithm for the key. Possible values: ssh-rsa, rsa-sha2-256, rsa-sha2-512.
	AlgorithmSigner pulumi.StringOutput `pulumi:"algorithmSigner"`
	// Specifies if host certificates that are requested are allowed to use the base domains listed in `allowedDomains`.
	AllowBareDomains pulumi.BoolPtrOutput `pulumi:"allowBareDomains"`
	// Specifies if certificates are allowed to be signed for use as a 'host'.
	AllowHostCertificates pulumi.BoolPtrOutput `pulumi:"allowHostCertificates"`
	// Specifies if host certificates that are requested are allowed to be subdomains of those listed in `allowedDomains`.
	AllowSubdomains pulumi.BoolPtrOutput `pulumi:"allowSubdomains"`
	// Specifies if certificates are allowed to be signed for use as a 'user'.
	AllowUserCertificates pulumi.BoolPtrOutput `pulumi:"allowUserCertificates"`
	// Specifies if users can override the key ID for a signed certificate with the `keyId` field.
	AllowUserKeyIds pulumi.BoolPtrOutput `pulumi:"allowUserKeyIds"`
	// Specifies a comma-separated list of critical options that certificates can have when signed.
	AllowedCriticalOptions pulumi.StringPtrOutput `pulumi:"allowedCriticalOptions"`
	// The list of domains for which a client can request a host certificate.
	AllowedDomains pulumi.StringPtrOutput `pulumi:"allowedDomains"`
	// Specifies if `allowedDomains` can be declared using
	// identity template policies. Non-templated domains are also permitted.
	AllowedDomainsTemplate pulumi.BoolOutput `pulumi:"allowedDomainsTemplate"`
	// Specifies a comma-separated list of extensions that certificates can have when signed.
	AllowedExtensions pulumi.StringPtrOutput `pulumi:"allowedExtensions"`
	// Set of configuration blocks to define allowed\
	// user key configuration, like key type and their lengths. Can be specified multiple times.
	// *See Configuration-Options for more info*
	AllowedUserKeyConfigs SecretBackendRoleAllowedUserKeyConfigArrayOutput `pulumi:"allowedUserKeyConfigs"`
	// Specifies a map of ssh key types and their expected sizes which
	// are allowed to be signed by the CA type.
	// *Deprecated: use* allowedUserKeyConfig *instead*
	//
	// Deprecated: Set in allowed_user_key_config
	AllowedUserKeyLengths pulumi.IntMapOutput `pulumi:"allowedUserKeyLengths"`
	// Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.
	AllowedUsers pulumi.StringPtrOutput `pulumi:"allowedUsers"`
	// Specifies if `allowedUsers` can be declared using identity template policies. Non-templated users are also permitted.
	AllowedUsersTemplate pulumi.BoolPtrOutput `pulumi:"allowedUsersTemplate"`
	// The path where the SSH secret backend is mounted.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// The comma-separated string of CIDR blocks for which this role is applicable.
	CidrList pulumi.StringPtrOutput `pulumi:"cidrList"`
	// Specifies a map of critical options that certificates have when signed.
	DefaultCriticalOptions pulumi.MapOutput `pulumi:"defaultCriticalOptions"`
	// Specifies a map of extensions that certificates have when signed.
	DefaultExtensions pulumi.MapOutput `pulumi:"defaultExtensions"`
	// Specifies the default username for which a credential will be generated.
	DefaultUser pulumi.StringPtrOutput `pulumi:"defaultUser"`
	// If set, `defaultUsers` can be specified using identity template values. A non-templated user is also permitted.
	DefaultUserTemplate pulumi.BoolPtrOutput `pulumi:"defaultUserTemplate"`
	// Specifies a custom format for the key id of a signed certificate.
	KeyIdFormat pulumi.StringPtrOutput `pulumi:"keyIdFormat"`
	// Specifies the type of credentials generated by this role. This can be either `otp`, `dynamic` or `ca`.
	KeyType pulumi.StringOutput `pulumi:"keyType"`
	// Specifies the maximum Time To Live value.
	MaxTtl pulumi.StringOutput `pulumi:"maxTtl"`
	// Specifies the name of the role to create.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// Specifies the duration by which to backdate the ValidAfter property. Uses duration format strings.
	NotBeforeDuration pulumi.StringOutput `pulumi:"notBeforeDuration"`
	// Specifies the Time To Live value.
	Ttl pulumi.StringOutput `pulumi:"ttl"`
}

// NewSecretBackendRole registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendRole(ctx *pulumi.Context,
	name string, args *SecretBackendRoleArgs, opts ...pulumi.ResourceOption) (*SecretBackendRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Backend == nil {
		return nil, errors.New("invalid value for required argument 'Backend'")
	}
	if args.KeyType == nil {
		return nil, errors.New("invalid value for required argument 'KeyType'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecretBackendRole
	err := ctx.RegisterResource("vault:ssh/secretBackendRole:SecretBackendRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecretBackendRole gets an existing SecretBackendRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecretBackendRoleState, opts ...pulumi.ResourceOption) (*SecretBackendRole, error) {
	var resource SecretBackendRole
	err := ctx.ReadResource("vault:ssh/secretBackendRole:SecretBackendRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendRole resources.
type secretBackendRoleState struct {
	// When supplied, this value specifies a signing algorithm for the key. Possible values: ssh-rsa, rsa-sha2-256, rsa-sha2-512.
	AlgorithmSigner *string `pulumi:"algorithmSigner"`
	// Specifies if host certificates that are requested are allowed to use the base domains listed in `allowedDomains`.
	AllowBareDomains *bool `pulumi:"allowBareDomains"`
	// Specifies if certificates are allowed to be signed for use as a 'host'.
	AllowHostCertificates *bool `pulumi:"allowHostCertificates"`
	// Specifies if host certificates that are requested are allowed to be subdomains of those listed in `allowedDomains`.
	AllowSubdomains *bool `pulumi:"allowSubdomains"`
	// Specifies if certificates are allowed to be signed for use as a 'user'.
	AllowUserCertificates *bool `pulumi:"allowUserCertificates"`
	// Specifies if users can override the key ID for a signed certificate with the `keyId` field.
	AllowUserKeyIds *bool `pulumi:"allowUserKeyIds"`
	// Specifies a comma-separated list of critical options that certificates can have when signed.
	AllowedCriticalOptions *string `pulumi:"allowedCriticalOptions"`
	// The list of domains for which a client can request a host certificate.
	AllowedDomains *string `pulumi:"allowedDomains"`
	// Specifies if `allowedDomains` can be declared using
	// identity template policies. Non-templated domains are also permitted.
	AllowedDomainsTemplate *bool `pulumi:"allowedDomainsTemplate"`
	// Specifies a comma-separated list of extensions that certificates can have when signed.
	AllowedExtensions *string `pulumi:"allowedExtensions"`
	// Set of configuration blocks to define allowed\
	// user key configuration, like key type and their lengths. Can be specified multiple times.
	// *See Configuration-Options for more info*
	AllowedUserKeyConfigs []SecretBackendRoleAllowedUserKeyConfig `pulumi:"allowedUserKeyConfigs"`
	// Specifies a map of ssh key types and their expected sizes which
	// are allowed to be signed by the CA type.
	// *Deprecated: use* allowedUserKeyConfig *instead*
	//
	// Deprecated: Set in allowed_user_key_config
	AllowedUserKeyLengths map[string]int `pulumi:"allowedUserKeyLengths"`
	// Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.
	AllowedUsers *string `pulumi:"allowedUsers"`
	// Specifies if `allowedUsers` can be declared using identity template policies. Non-templated users are also permitted.
	AllowedUsersTemplate *bool `pulumi:"allowedUsersTemplate"`
	// The path where the SSH secret backend is mounted.
	Backend *string `pulumi:"backend"`
	// The comma-separated string of CIDR blocks for which this role is applicable.
	CidrList *string `pulumi:"cidrList"`
	// Specifies a map of critical options that certificates have when signed.
	DefaultCriticalOptions map[string]interface{} `pulumi:"defaultCriticalOptions"`
	// Specifies a map of extensions that certificates have when signed.
	DefaultExtensions map[string]interface{} `pulumi:"defaultExtensions"`
	// Specifies the default username for which a credential will be generated.
	DefaultUser *string `pulumi:"defaultUser"`
	// If set, `defaultUsers` can be specified using identity template values. A non-templated user is also permitted.
	DefaultUserTemplate *bool `pulumi:"defaultUserTemplate"`
	// Specifies a custom format for the key id of a signed certificate.
	KeyIdFormat *string `pulumi:"keyIdFormat"`
	// Specifies the type of credentials generated by this role. This can be either `otp`, `dynamic` or `ca`.
	KeyType *string `pulumi:"keyType"`
	// Specifies the maximum Time To Live value.
	MaxTtl *string `pulumi:"maxTtl"`
	// Specifies the name of the role to create.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Specifies the duration by which to backdate the ValidAfter property. Uses duration format strings.
	NotBeforeDuration *string `pulumi:"notBeforeDuration"`
	// Specifies the Time To Live value.
	Ttl *string `pulumi:"ttl"`
}

type SecretBackendRoleState struct {
	// When supplied, this value specifies a signing algorithm for the key. Possible values: ssh-rsa, rsa-sha2-256, rsa-sha2-512.
	AlgorithmSigner pulumi.StringPtrInput
	// Specifies if host certificates that are requested are allowed to use the base domains listed in `allowedDomains`.
	AllowBareDomains pulumi.BoolPtrInput
	// Specifies if certificates are allowed to be signed for use as a 'host'.
	AllowHostCertificates pulumi.BoolPtrInput
	// Specifies if host certificates that are requested are allowed to be subdomains of those listed in `allowedDomains`.
	AllowSubdomains pulumi.BoolPtrInput
	// Specifies if certificates are allowed to be signed for use as a 'user'.
	AllowUserCertificates pulumi.BoolPtrInput
	// Specifies if users can override the key ID for a signed certificate with the `keyId` field.
	AllowUserKeyIds pulumi.BoolPtrInput
	// Specifies a comma-separated list of critical options that certificates can have when signed.
	AllowedCriticalOptions pulumi.StringPtrInput
	// The list of domains for which a client can request a host certificate.
	AllowedDomains pulumi.StringPtrInput
	// Specifies if `allowedDomains` can be declared using
	// identity template policies. Non-templated domains are also permitted.
	AllowedDomainsTemplate pulumi.BoolPtrInput
	// Specifies a comma-separated list of extensions that certificates can have when signed.
	AllowedExtensions pulumi.StringPtrInput
	// Set of configuration blocks to define allowed\
	// user key configuration, like key type and their lengths. Can be specified multiple times.
	// *See Configuration-Options for more info*
	AllowedUserKeyConfigs SecretBackendRoleAllowedUserKeyConfigArrayInput
	// Specifies a map of ssh key types and their expected sizes which
	// are allowed to be signed by the CA type.
	// *Deprecated: use* allowedUserKeyConfig *instead*
	//
	// Deprecated: Set in allowed_user_key_config
	AllowedUserKeyLengths pulumi.IntMapInput
	// Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.
	AllowedUsers pulumi.StringPtrInput
	// Specifies if `allowedUsers` can be declared using identity template policies. Non-templated users are also permitted.
	AllowedUsersTemplate pulumi.BoolPtrInput
	// The path where the SSH secret backend is mounted.
	Backend pulumi.StringPtrInput
	// The comma-separated string of CIDR blocks for which this role is applicable.
	CidrList pulumi.StringPtrInput
	// Specifies a map of critical options that certificates have when signed.
	DefaultCriticalOptions pulumi.MapInput
	// Specifies a map of extensions that certificates have when signed.
	DefaultExtensions pulumi.MapInput
	// Specifies the default username for which a credential will be generated.
	DefaultUser pulumi.StringPtrInput
	// If set, `defaultUsers` can be specified using identity template values. A non-templated user is also permitted.
	DefaultUserTemplate pulumi.BoolPtrInput
	// Specifies a custom format for the key id of a signed certificate.
	KeyIdFormat pulumi.StringPtrInput
	// Specifies the type of credentials generated by this role. This can be either `otp`, `dynamic` or `ca`.
	KeyType pulumi.StringPtrInput
	// Specifies the maximum Time To Live value.
	MaxTtl pulumi.StringPtrInput
	// Specifies the name of the role to create.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Specifies the duration by which to backdate the ValidAfter property. Uses duration format strings.
	NotBeforeDuration pulumi.StringPtrInput
	// Specifies the Time To Live value.
	Ttl pulumi.StringPtrInput
}

func (SecretBackendRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendRoleState)(nil)).Elem()
}

type secretBackendRoleArgs struct {
	// When supplied, this value specifies a signing algorithm for the key. Possible values: ssh-rsa, rsa-sha2-256, rsa-sha2-512.
	AlgorithmSigner *string `pulumi:"algorithmSigner"`
	// Specifies if host certificates that are requested are allowed to use the base domains listed in `allowedDomains`.
	AllowBareDomains *bool `pulumi:"allowBareDomains"`
	// Specifies if certificates are allowed to be signed for use as a 'host'.
	AllowHostCertificates *bool `pulumi:"allowHostCertificates"`
	// Specifies if host certificates that are requested are allowed to be subdomains of those listed in `allowedDomains`.
	AllowSubdomains *bool `pulumi:"allowSubdomains"`
	// Specifies if certificates are allowed to be signed for use as a 'user'.
	AllowUserCertificates *bool `pulumi:"allowUserCertificates"`
	// Specifies if users can override the key ID for a signed certificate with the `keyId` field.
	AllowUserKeyIds *bool `pulumi:"allowUserKeyIds"`
	// Specifies a comma-separated list of critical options that certificates can have when signed.
	AllowedCriticalOptions *string `pulumi:"allowedCriticalOptions"`
	// The list of domains for which a client can request a host certificate.
	AllowedDomains *string `pulumi:"allowedDomains"`
	// Specifies if `allowedDomains` can be declared using
	// identity template policies. Non-templated domains are also permitted.
	AllowedDomainsTemplate *bool `pulumi:"allowedDomainsTemplate"`
	// Specifies a comma-separated list of extensions that certificates can have when signed.
	AllowedExtensions *string `pulumi:"allowedExtensions"`
	// Set of configuration blocks to define allowed\
	// user key configuration, like key type and their lengths. Can be specified multiple times.
	// *See Configuration-Options for more info*
	AllowedUserKeyConfigs []SecretBackendRoleAllowedUserKeyConfig `pulumi:"allowedUserKeyConfigs"`
	// Specifies a map of ssh key types and their expected sizes which
	// are allowed to be signed by the CA type.
	// *Deprecated: use* allowedUserKeyConfig *instead*
	//
	// Deprecated: Set in allowed_user_key_config
	AllowedUserKeyLengths map[string]int `pulumi:"allowedUserKeyLengths"`
	// Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.
	AllowedUsers *string `pulumi:"allowedUsers"`
	// Specifies if `allowedUsers` can be declared using identity template policies. Non-templated users are also permitted.
	AllowedUsersTemplate *bool `pulumi:"allowedUsersTemplate"`
	// The path where the SSH secret backend is mounted.
	Backend string `pulumi:"backend"`
	// The comma-separated string of CIDR blocks for which this role is applicable.
	CidrList *string `pulumi:"cidrList"`
	// Specifies a map of critical options that certificates have when signed.
	DefaultCriticalOptions map[string]interface{} `pulumi:"defaultCriticalOptions"`
	// Specifies a map of extensions that certificates have when signed.
	DefaultExtensions map[string]interface{} `pulumi:"defaultExtensions"`
	// Specifies the default username for which a credential will be generated.
	DefaultUser *string `pulumi:"defaultUser"`
	// If set, `defaultUsers` can be specified using identity template values. A non-templated user is also permitted.
	DefaultUserTemplate *bool `pulumi:"defaultUserTemplate"`
	// Specifies a custom format for the key id of a signed certificate.
	KeyIdFormat *string `pulumi:"keyIdFormat"`
	// Specifies the type of credentials generated by this role. This can be either `otp`, `dynamic` or `ca`.
	KeyType string `pulumi:"keyType"`
	// Specifies the maximum Time To Live value.
	MaxTtl *string `pulumi:"maxTtl"`
	// Specifies the name of the role to create.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// Specifies the duration by which to backdate the ValidAfter property. Uses duration format strings.
	NotBeforeDuration *string `pulumi:"notBeforeDuration"`
	// Specifies the Time To Live value.
	Ttl *string `pulumi:"ttl"`
}

// The set of arguments for constructing a SecretBackendRole resource.
type SecretBackendRoleArgs struct {
	// When supplied, this value specifies a signing algorithm for the key. Possible values: ssh-rsa, rsa-sha2-256, rsa-sha2-512.
	AlgorithmSigner pulumi.StringPtrInput
	// Specifies if host certificates that are requested are allowed to use the base domains listed in `allowedDomains`.
	AllowBareDomains pulumi.BoolPtrInput
	// Specifies if certificates are allowed to be signed for use as a 'host'.
	AllowHostCertificates pulumi.BoolPtrInput
	// Specifies if host certificates that are requested are allowed to be subdomains of those listed in `allowedDomains`.
	AllowSubdomains pulumi.BoolPtrInput
	// Specifies if certificates are allowed to be signed for use as a 'user'.
	AllowUserCertificates pulumi.BoolPtrInput
	// Specifies if users can override the key ID for a signed certificate with the `keyId` field.
	AllowUserKeyIds pulumi.BoolPtrInput
	// Specifies a comma-separated list of critical options that certificates can have when signed.
	AllowedCriticalOptions pulumi.StringPtrInput
	// The list of domains for which a client can request a host certificate.
	AllowedDomains pulumi.StringPtrInput
	// Specifies if `allowedDomains` can be declared using
	// identity template policies. Non-templated domains are also permitted.
	AllowedDomainsTemplate pulumi.BoolPtrInput
	// Specifies a comma-separated list of extensions that certificates can have when signed.
	AllowedExtensions pulumi.StringPtrInput
	// Set of configuration blocks to define allowed\
	// user key configuration, like key type and their lengths. Can be specified multiple times.
	// *See Configuration-Options for more info*
	AllowedUserKeyConfigs SecretBackendRoleAllowedUserKeyConfigArrayInput
	// Specifies a map of ssh key types and their expected sizes which
	// are allowed to be signed by the CA type.
	// *Deprecated: use* allowedUserKeyConfig *instead*
	//
	// Deprecated: Set in allowed_user_key_config
	AllowedUserKeyLengths pulumi.IntMapInput
	// Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.
	AllowedUsers pulumi.StringPtrInput
	// Specifies if `allowedUsers` can be declared using identity template policies. Non-templated users are also permitted.
	AllowedUsersTemplate pulumi.BoolPtrInput
	// The path where the SSH secret backend is mounted.
	Backend pulumi.StringInput
	// The comma-separated string of CIDR blocks for which this role is applicable.
	CidrList pulumi.StringPtrInput
	// Specifies a map of critical options that certificates have when signed.
	DefaultCriticalOptions pulumi.MapInput
	// Specifies a map of extensions that certificates have when signed.
	DefaultExtensions pulumi.MapInput
	// Specifies the default username for which a credential will be generated.
	DefaultUser pulumi.StringPtrInput
	// If set, `defaultUsers` can be specified using identity template values. A non-templated user is also permitted.
	DefaultUserTemplate pulumi.BoolPtrInput
	// Specifies a custom format for the key id of a signed certificate.
	KeyIdFormat pulumi.StringPtrInput
	// Specifies the type of credentials generated by this role. This can be either `otp`, `dynamic` or `ca`.
	KeyType pulumi.StringInput
	// Specifies the maximum Time To Live value.
	MaxTtl pulumi.StringPtrInput
	// Specifies the name of the role to create.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// Specifies the duration by which to backdate the ValidAfter property. Uses duration format strings.
	NotBeforeDuration pulumi.StringPtrInput
	// Specifies the Time To Live value.
	Ttl pulumi.StringPtrInput
}

func (SecretBackendRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendRoleArgs)(nil)).Elem()
}

type SecretBackendRoleInput interface {
	pulumi.Input

	ToSecretBackendRoleOutput() SecretBackendRoleOutput
	ToSecretBackendRoleOutputWithContext(ctx context.Context) SecretBackendRoleOutput
}

func (*SecretBackendRole) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendRole)(nil)).Elem()
}

func (i *SecretBackendRole) ToSecretBackendRoleOutput() SecretBackendRoleOutput {
	return i.ToSecretBackendRoleOutputWithContext(context.Background())
}

func (i *SecretBackendRole) ToSecretBackendRoleOutputWithContext(ctx context.Context) SecretBackendRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRoleOutput)
}

// SecretBackendRoleArrayInput is an input type that accepts SecretBackendRoleArray and SecretBackendRoleArrayOutput values.
// You can construct a concrete instance of `SecretBackendRoleArrayInput` via:
//
//	SecretBackendRoleArray{ SecretBackendRoleArgs{...} }
type SecretBackendRoleArrayInput interface {
	pulumi.Input

	ToSecretBackendRoleArrayOutput() SecretBackendRoleArrayOutput
	ToSecretBackendRoleArrayOutputWithContext(context.Context) SecretBackendRoleArrayOutput
}

type SecretBackendRoleArray []SecretBackendRoleInput

func (SecretBackendRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendRole)(nil)).Elem()
}

func (i SecretBackendRoleArray) ToSecretBackendRoleArrayOutput() SecretBackendRoleArrayOutput {
	return i.ToSecretBackendRoleArrayOutputWithContext(context.Background())
}

func (i SecretBackendRoleArray) ToSecretBackendRoleArrayOutputWithContext(ctx context.Context) SecretBackendRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRoleArrayOutput)
}

// SecretBackendRoleMapInput is an input type that accepts SecretBackendRoleMap and SecretBackendRoleMapOutput values.
// You can construct a concrete instance of `SecretBackendRoleMapInput` via:
//
//	SecretBackendRoleMap{ "key": SecretBackendRoleArgs{...} }
type SecretBackendRoleMapInput interface {
	pulumi.Input

	ToSecretBackendRoleMapOutput() SecretBackendRoleMapOutput
	ToSecretBackendRoleMapOutputWithContext(context.Context) SecretBackendRoleMapOutput
}

type SecretBackendRoleMap map[string]SecretBackendRoleInput

func (SecretBackendRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendRole)(nil)).Elem()
}

func (i SecretBackendRoleMap) ToSecretBackendRoleMapOutput() SecretBackendRoleMapOutput {
	return i.ToSecretBackendRoleMapOutputWithContext(context.Background())
}

func (i SecretBackendRoleMap) ToSecretBackendRoleMapOutputWithContext(ctx context.Context) SecretBackendRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretBackendRoleMapOutput)
}

type SecretBackendRoleOutput struct{ *pulumi.OutputState }

func (SecretBackendRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecretBackendRole)(nil)).Elem()
}

func (o SecretBackendRoleOutput) ToSecretBackendRoleOutput() SecretBackendRoleOutput {
	return o
}

func (o SecretBackendRoleOutput) ToSecretBackendRoleOutputWithContext(ctx context.Context) SecretBackendRoleOutput {
	return o
}

// When supplied, this value specifies a signing algorithm for the key. Possible values: ssh-rsa, rsa-sha2-256, rsa-sha2-512.
func (o SecretBackendRoleOutput) AlgorithmSigner() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringOutput { return v.AlgorithmSigner }).(pulumi.StringOutput)
}

// Specifies if host certificates that are requested are allowed to use the base domains listed in `allowedDomains`.
func (o SecretBackendRoleOutput) AllowBareDomains() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.BoolPtrOutput { return v.AllowBareDomains }).(pulumi.BoolPtrOutput)
}

// Specifies if certificates are allowed to be signed for use as a 'host'.
func (o SecretBackendRoleOutput) AllowHostCertificates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.BoolPtrOutput { return v.AllowHostCertificates }).(pulumi.BoolPtrOutput)
}

// Specifies if host certificates that are requested are allowed to be subdomains of those listed in `allowedDomains`.
func (o SecretBackendRoleOutput) AllowSubdomains() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.BoolPtrOutput { return v.AllowSubdomains }).(pulumi.BoolPtrOutput)
}

// Specifies if certificates are allowed to be signed for use as a 'user'.
func (o SecretBackendRoleOutput) AllowUserCertificates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.BoolPtrOutput { return v.AllowUserCertificates }).(pulumi.BoolPtrOutput)
}

// Specifies if users can override the key ID for a signed certificate with the `keyId` field.
func (o SecretBackendRoleOutput) AllowUserKeyIds() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.BoolPtrOutput { return v.AllowUserKeyIds }).(pulumi.BoolPtrOutput)
}

// Specifies a comma-separated list of critical options that certificates can have when signed.
func (o SecretBackendRoleOutput) AllowedCriticalOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringPtrOutput { return v.AllowedCriticalOptions }).(pulumi.StringPtrOutput)
}

// The list of domains for which a client can request a host certificate.
func (o SecretBackendRoleOutput) AllowedDomains() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringPtrOutput { return v.AllowedDomains }).(pulumi.StringPtrOutput)
}

// Specifies if `allowedDomains` can be declared using
// identity template policies. Non-templated domains are also permitted.
func (o SecretBackendRoleOutput) AllowedDomainsTemplate() pulumi.BoolOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.BoolOutput { return v.AllowedDomainsTemplate }).(pulumi.BoolOutput)
}

// Specifies a comma-separated list of extensions that certificates can have when signed.
func (o SecretBackendRoleOutput) AllowedExtensions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringPtrOutput { return v.AllowedExtensions }).(pulumi.StringPtrOutput)
}

// Set of configuration blocks to define allowed\
// user key configuration, like key type and their lengths. Can be specified multiple times.
// *See Configuration-Options for more info*
func (o SecretBackendRoleOutput) AllowedUserKeyConfigs() SecretBackendRoleAllowedUserKeyConfigArrayOutput {
	return o.ApplyT(func(v *SecretBackendRole) SecretBackendRoleAllowedUserKeyConfigArrayOutput {
		return v.AllowedUserKeyConfigs
	}).(SecretBackendRoleAllowedUserKeyConfigArrayOutput)
}

// Specifies a map of ssh key types and their expected sizes which
// are allowed to be signed by the CA type.
// *Deprecated: use* allowedUserKeyConfig *instead*
//
// Deprecated: Set in allowed_user_key_config
func (o SecretBackendRoleOutput) AllowedUserKeyLengths() pulumi.IntMapOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.IntMapOutput { return v.AllowedUserKeyLengths }).(pulumi.IntMapOutput)
}

// Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.
func (o SecretBackendRoleOutput) AllowedUsers() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringPtrOutput { return v.AllowedUsers }).(pulumi.StringPtrOutput)
}

// Specifies if `allowedUsers` can be declared using identity template policies. Non-templated users are also permitted.
func (o SecretBackendRoleOutput) AllowedUsersTemplate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.BoolPtrOutput { return v.AllowedUsersTemplate }).(pulumi.BoolPtrOutput)
}

// The path where the SSH secret backend is mounted.
func (o SecretBackendRoleOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringOutput { return v.Backend }).(pulumi.StringOutput)
}

// The comma-separated string of CIDR blocks for which this role is applicable.
func (o SecretBackendRoleOutput) CidrList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringPtrOutput { return v.CidrList }).(pulumi.StringPtrOutput)
}

// Specifies a map of critical options that certificates have when signed.
func (o SecretBackendRoleOutput) DefaultCriticalOptions() pulumi.MapOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.MapOutput { return v.DefaultCriticalOptions }).(pulumi.MapOutput)
}

// Specifies a map of extensions that certificates have when signed.
func (o SecretBackendRoleOutput) DefaultExtensions() pulumi.MapOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.MapOutput { return v.DefaultExtensions }).(pulumi.MapOutput)
}

// Specifies the default username for which a credential will be generated.
func (o SecretBackendRoleOutput) DefaultUser() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringPtrOutput { return v.DefaultUser }).(pulumi.StringPtrOutput)
}

// If set, `defaultUsers` can be specified using identity template values. A non-templated user is also permitted.
func (o SecretBackendRoleOutput) DefaultUserTemplate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.BoolPtrOutput { return v.DefaultUserTemplate }).(pulumi.BoolPtrOutput)
}

// Specifies a custom format for the key id of a signed certificate.
func (o SecretBackendRoleOutput) KeyIdFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringPtrOutput { return v.KeyIdFormat }).(pulumi.StringPtrOutput)
}

// Specifies the type of credentials generated by this role. This can be either `otp`, `dynamic` or `ca`.
func (o SecretBackendRoleOutput) KeyType() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringOutput { return v.KeyType }).(pulumi.StringOutput)
}

// Specifies the maximum Time To Live value.
func (o SecretBackendRoleOutput) MaxTtl() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringOutput { return v.MaxTtl }).(pulumi.StringOutput)
}

// Specifies the name of the role to create.
func (o SecretBackendRoleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o SecretBackendRoleOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Specifies the duration by which to backdate the ValidAfter property. Uses duration format strings.
func (o SecretBackendRoleOutput) NotBeforeDuration() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringOutput { return v.NotBeforeDuration }).(pulumi.StringOutput)
}

// Specifies the Time To Live value.
func (o SecretBackendRoleOutput) Ttl() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringOutput { return v.Ttl }).(pulumi.StringOutput)
}

type SecretBackendRoleArrayOutput struct{ *pulumi.OutputState }

func (SecretBackendRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecretBackendRole)(nil)).Elem()
}

func (o SecretBackendRoleArrayOutput) ToSecretBackendRoleArrayOutput() SecretBackendRoleArrayOutput {
	return o
}

func (o SecretBackendRoleArrayOutput) ToSecretBackendRoleArrayOutputWithContext(ctx context.Context) SecretBackendRoleArrayOutput {
	return o
}

func (o SecretBackendRoleArrayOutput) Index(i pulumi.IntInput) SecretBackendRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecretBackendRole {
		return vs[0].([]*SecretBackendRole)[vs[1].(int)]
	}).(SecretBackendRoleOutput)
}

type SecretBackendRoleMapOutput struct{ *pulumi.OutputState }

func (SecretBackendRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecretBackendRole)(nil)).Elem()
}

func (o SecretBackendRoleMapOutput) ToSecretBackendRoleMapOutput() SecretBackendRoleMapOutput {
	return o
}

func (o SecretBackendRoleMapOutput) ToSecretBackendRoleMapOutputWithContext(ctx context.Context) SecretBackendRoleMapOutput {
	return o
}

func (o SecretBackendRoleMapOutput) MapIndex(k pulumi.StringInput) SecretBackendRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecretBackendRole {
		return vs[0].(map[string]*SecretBackendRole)[vs[1].(string)]
	}).(SecretBackendRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRoleInput)(nil)).Elem(), &SecretBackendRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRoleArrayInput)(nil)).Elem(), SecretBackendRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecretBackendRoleMapInput)(nil)).Elem(), SecretBackendRoleMap{})
	pulumi.RegisterOutputType(SecretBackendRoleOutput{})
	pulumi.RegisterOutputType(SecretBackendRoleArrayOutput{})
	pulumi.RegisterOutputType(SecretBackendRoleMapOutput{})
}
