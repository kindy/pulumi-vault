// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The provider type for the vault package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}
	if args.Address == nil {
		args.Address = pulumi.StringPtr(getEnvOrDefault("", nil, "VAULT_ADDR").(string))
	}
	if args.CaCertDir == nil {
		args.CaCertDir = pulumi.StringPtr(getEnvOrDefault("", nil, "VAULT_CAPATH").(string))
	}
	if args.CaCertFile == nil {
		args.CaCertFile = pulumi.StringPtr(getEnvOrDefault("", nil, "VAULT_CACERT").(string))
	}
	if args.MaxLeaseTtlSeconds == nil {
		args.MaxLeaseTtlSeconds = pulumi.IntPtr(getEnvOrDefault(1200, parseEnvInt, "TERRAFORM_VAULT_MAX_TTL").(int))
	}
	if args.MaxRetries == nil {
		args.MaxRetries = pulumi.IntPtr(getEnvOrDefault(2, parseEnvInt, "VAULT_MAX_RETRIES").(int))
	}
	if args.Namespace == nil {
		args.Namespace = pulumi.StringPtr(getEnvOrDefault("", nil, "VAULT_NAMESPACE").(string))
	}
	if args.SkipTlsVerify == nil {
		args.SkipTlsVerify = pulumi.BoolPtr(getEnvOrDefault(false, parseEnvBool, "VAULT_SKIP_VERIFY").(bool))
	}
	if args.Token == nil {
		args.Token = pulumi.StringPtr(getEnvOrDefault("", nil, "VAULT_TOKEN").(string))
	}
	if args.TokenName == nil {
		args.TokenName = pulumi.StringPtr(getEnvOrDefault("", nil, "VAULT_TOKEN_NAME").(string))
	}
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:vault", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// If true, adds the value of the `address` argument to the Terraform process environment.
	AddAddressToEnv *string `pulumi:"addAddressToEnv"`
	// URL of the root of the target Vault server.
	Address *string `pulumi:"address"`
	// Login to vault with an existing auth method using auth/<mount>/login
	AuthLogins []ProviderAuthLogin `pulumi:"authLogins"`
	// Path to directory containing CA certificate files to validate the server's certificate.
	CaCertDir *string `pulumi:"caCertDir"`
	// Path to a CA certificate file to validate the server's certificate.
	CaCertFile *string `pulumi:"caCertFile"`
	// Client authentication credentials.
	ClientAuths []ProviderClientAuth `pulumi:"clientAuths"`
	// The headers to send with each Vault request.
	Headers []ProviderHeader `pulumi:"headers"`
	// Maximum TTL for secret leases requested by this provider
	MaxLeaseTtlSeconds *int `pulumi:"maxLeaseTtlSeconds"`
	// Maximum number of retries when a 5xx error code is encountered.
	MaxRetries *int `pulumi:"maxRetries"`
	// The namespace to use. Available only for Vault Enterprise
	Namespace *string `pulumi:"namespace"`
	// Set this to true only if the target Vault server is an insecure development instance.
	SkipTlsVerify *bool `pulumi:"skipTlsVerify"`
	// Token to use to authenticate to Vault.
	Token *string `pulumi:"token"`
	// Token name to use for creating the Vault child token.
	TokenName *string `pulumi:"tokenName"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// If true, adds the value of the `address` argument to the Terraform process environment.
	AddAddressToEnv pulumi.StringPtrInput
	// URL of the root of the target Vault server.
	Address pulumi.StringPtrInput
	// Login to vault with an existing auth method using auth/<mount>/login
	AuthLogins ProviderAuthLoginArrayInput
	// Path to directory containing CA certificate files to validate the server's certificate.
	CaCertDir pulumi.StringPtrInput
	// Path to a CA certificate file to validate the server's certificate.
	CaCertFile pulumi.StringPtrInput
	// Client authentication credentials.
	ClientAuths ProviderClientAuthArrayInput
	// The headers to send with each Vault request.
	Headers ProviderHeaderArrayInput
	// Maximum TTL for secret leases requested by this provider
	MaxLeaseTtlSeconds pulumi.IntPtrInput
	// Maximum number of retries when a 5xx error code is encountered.
	MaxRetries pulumi.IntPtrInput
	// The namespace to use. Available only for Vault Enterprise
	Namespace pulumi.StringPtrInput
	// Set this to true only if the target Vault server is an insecure development instance.
	SkipTlsVerify pulumi.BoolPtrInput
	// Token to use to authenticate to Vault.
	Token pulumi.StringPtrInput
	// Token name to use for creating the Vault child token.
	TokenName pulumi.StringPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}
