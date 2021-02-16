// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

// If true, adds the value of the `address` argument to the Terraform process environment.
func GetAddAddressToEnv(ctx *pulumi.Context) string {
	return config.Get(ctx, "vault:addAddressToEnv")
}

// URL of the root of the target Vault server.
func GetAddress(ctx *pulumi.Context) string {
	return config.Get(ctx, "vault:address")
}

// Login to vault with an existing auth method using auth/<mount>/login
func GetAuthLogins(ctx *pulumi.Context) string {
	return config.Get(ctx, "vault:authLogins")
}

// Path to directory containing CA certificate files to validate the server's certificate.
func GetCaCertDir(ctx *pulumi.Context) string {
	return config.Get(ctx, "vault:caCertDir")
}

// Path to a CA certificate file to validate the server's certificate.
func GetCaCertFile(ctx *pulumi.Context) string {
	return config.Get(ctx, "vault:caCertFile")
}

// Client authentication credentials.
func GetClientAuths(ctx *pulumi.Context) string {
	return config.Get(ctx, "vault:clientAuths")
}

// The headers to send with each Vault request.
func GetHeaders(ctx *pulumi.Context) string {
	return config.Get(ctx, "vault:headers")
}

// Maximum TTL for secret leases requested by this provider
func GetMaxLeaseTtlSeconds(ctx *pulumi.Context) int {
	v, err := config.TryInt(ctx, "vault:maxLeaseTtlSeconds")
	if err == nil {
		return v
	}
	return getEnvOrDefault(1200, parseEnvInt, "TERRAFORM_VAULT_MAX_TTL").(int)
}

// Maximum number of retries when a 5xx error code is encountered.
func GetMaxRetries(ctx *pulumi.Context) int {
	v, err := config.TryInt(ctx, "vault:maxRetries")
	if err == nil {
		return v
	}
	return getEnvOrDefault(2, parseEnvInt, "VAULT_MAX_RETRIES").(int)
}

// The namespace to use. Available only for Vault Enterprise
func GetNamespace(ctx *pulumi.Context) string {
	return config.Get(ctx, "vault:namespace")
}

// Set this to true only if the target Vault server is an insecure development instance.
func GetSkipTlsVerify(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "vault:skipTlsVerify")
	if err == nil {
		return v
	}
	return getEnvOrDefault(false, parseEnvBool, "VAULT_SKIP_VERIFY").(bool)
}

// Token to use to authenticate to Vault.
func GetToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "vault:token")
}

// Token name to use for creating the Vault child token.
func GetTokenName(ctx *pulumi.Context) string {
	return config.Get(ctx, "vault:tokenName")
}
