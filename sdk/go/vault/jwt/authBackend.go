// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package jwt

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Provides a resource for managing an
// [JWT auth backend within Vault](https://www.vaultproject.io/docs/auth/jwt.html).
// 
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/jwt_auth_backend.html.markdown.
type AuthBackend struct {
	pulumi.CustomResourceState

	// The accessor of the JWT auth backend
	Accessor pulumi.StringOutput `pulumi:"accessor"`
	// The value against which to match the iss claim in a JWT
	BoundIssuer pulumi.StringPtrOutput `pulumi:"boundIssuer"`
	// The default role to use if none is provided during login
	DefaultRole pulumi.StringPtrOutput `pulumi:"defaultRole"`
	// The description of the auth backend
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
	JwksCaPem pulumi.StringPtrOutput `pulumi:"jwksCaPem"`
	// JWKS URL to use to authenticate signatures. Cannot be used with "oidcDiscoveryUrl" or "jwtValidationPubkeys".
	JwksUrl pulumi.StringPtrOutput `pulumi:"jwksUrl"`
	// A list of supported signing algorithms. Vault 1.1.0 defaults to [RS256] but future or past versions of Vault may differ
	JwtSupportedAlgs pulumi.StringArrayOutput `pulumi:"jwtSupportedAlgs"`
	// A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used in combination with `oidcDiscoveryUrl`
	JwtValidationPubkeys pulumi.StringArrayOutput `pulumi:"jwtValidationPubkeys"`
	// Client ID used for OIDC backends
	OidcClientId pulumi.StringPtrOutput `pulumi:"oidcClientId"`
	// Client Secret used for OIDC backends
	OidcClientSecret pulumi.StringPtrOutput `pulumi:"oidcClientSecret"`
	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
	OidcDiscoveryCaPem pulumi.StringPtrOutput `pulumi:"oidcDiscoveryCaPem"`
	// The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwtValidationPubkeys`
	OidcDiscoveryUrl pulumi.StringPtrOutput `pulumi:"oidcDiscoveryUrl"`
	// Path to mount the JWT/OIDC auth backend
	Path pulumi.StringPtrOutput `pulumi:"path"`
	Tune AuthBackendTuneOutput `pulumi:"tune"`
	// Type of auth backend. Should be one of `jwt` or `oidc`. Default - `jwt`
	Type pulumi.StringPtrOutput `pulumi:"type"`
}

// NewAuthBackend registers a new resource with the given unique name, arguments, and options.
func NewAuthBackend(ctx *pulumi.Context,
	name string, args *AuthBackendArgs, opts ...pulumi.ResourceOption) (*AuthBackend, error) {
	if args == nil {
		args = &AuthBackendArgs{}
	}
	var resource AuthBackend
	err := ctx.RegisterResource("vault:jwt/authBackend:AuthBackend", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackend gets an existing AuthBackend resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackend(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendState, opts ...pulumi.ResourceOption) (*AuthBackend, error) {
	var resource AuthBackend
	err := ctx.ReadResource("vault:jwt/authBackend:AuthBackend", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackend resources.
type authBackendState struct {
	// The accessor of the JWT auth backend
	Accessor *string `pulumi:"accessor"`
	// The value against which to match the iss claim in a JWT
	BoundIssuer *string `pulumi:"boundIssuer"`
	// The default role to use if none is provided during login
	DefaultRole *string `pulumi:"defaultRole"`
	// The description of the auth backend
	Description *string `pulumi:"description"`
	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
	JwksCaPem *string `pulumi:"jwksCaPem"`
	// JWKS URL to use to authenticate signatures. Cannot be used with "oidcDiscoveryUrl" or "jwtValidationPubkeys".
	JwksUrl *string `pulumi:"jwksUrl"`
	// A list of supported signing algorithms. Vault 1.1.0 defaults to [RS256] but future or past versions of Vault may differ
	JwtSupportedAlgs []string `pulumi:"jwtSupportedAlgs"`
	// A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used in combination with `oidcDiscoveryUrl`
	JwtValidationPubkeys []string `pulumi:"jwtValidationPubkeys"`
	// Client ID used for OIDC backends
	OidcClientId *string `pulumi:"oidcClientId"`
	// Client Secret used for OIDC backends
	OidcClientSecret *string `pulumi:"oidcClientSecret"`
	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
	OidcDiscoveryCaPem *string `pulumi:"oidcDiscoveryCaPem"`
	// The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwtValidationPubkeys`
	OidcDiscoveryUrl *string `pulumi:"oidcDiscoveryUrl"`
	// Path to mount the JWT/OIDC auth backend
	Path *string `pulumi:"path"`
	Tune *AuthBackendTune `pulumi:"tune"`
	// Type of auth backend. Should be one of `jwt` or `oidc`. Default - `jwt`
	Type *string `pulumi:"type"`
}

type AuthBackendState struct {
	// The accessor of the JWT auth backend
	Accessor pulumi.StringPtrInput
	// The value against which to match the iss claim in a JWT
	BoundIssuer pulumi.StringPtrInput
	// The default role to use if none is provided during login
	DefaultRole pulumi.StringPtrInput
	// The description of the auth backend
	Description pulumi.StringPtrInput
	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
	JwksCaPem pulumi.StringPtrInput
	// JWKS URL to use to authenticate signatures. Cannot be used with "oidcDiscoveryUrl" or "jwtValidationPubkeys".
	JwksUrl pulumi.StringPtrInput
	// A list of supported signing algorithms. Vault 1.1.0 defaults to [RS256] but future or past versions of Vault may differ
	JwtSupportedAlgs pulumi.StringArrayInput
	// A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used in combination with `oidcDiscoveryUrl`
	JwtValidationPubkeys pulumi.StringArrayInput
	// Client ID used for OIDC backends
	OidcClientId pulumi.StringPtrInput
	// Client Secret used for OIDC backends
	OidcClientSecret pulumi.StringPtrInput
	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
	OidcDiscoveryCaPem pulumi.StringPtrInput
	// The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwtValidationPubkeys`
	OidcDiscoveryUrl pulumi.StringPtrInput
	// Path to mount the JWT/OIDC auth backend
	Path pulumi.StringPtrInput
	Tune AuthBackendTunePtrInput
	// Type of auth backend. Should be one of `jwt` or `oidc`. Default - `jwt`
	Type pulumi.StringPtrInput
}

func (AuthBackendState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendState)(nil)).Elem()
}

type authBackendArgs struct {
	// The value against which to match the iss claim in a JWT
	BoundIssuer *string `pulumi:"boundIssuer"`
	// The default role to use if none is provided during login
	DefaultRole *string `pulumi:"defaultRole"`
	// The description of the auth backend
	Description *string `pulumi:"description"`
	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
	JwksCaPem *string `pulumi:"jwksCaPem"`
	// JWKS URL to use to authenticate signatures. Cannot be used with "oidcDiscoveryUrl" or "jwtValidationPubkeys".
	JwksUrl *string `pulumi:"jwksUrl"`
	// A list of supported signing algorithms. Vault 1.1.0 defaults to [RS256] but future or past versions of Vault may differ
	JwtSupportedAlgs []string `pulumi:"jwtSupportedAlgs"`
	// A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used in combination with `oidcDiscoveryUrl`
	JwtValidationPubkeys []string `pulumi:"jwtValidationPubkeys"`
	// Client ID used for OIDC backends
	OidcClientId *string `pulumi:"oidcClientId"`
	// Client Secret used for OIDC backends
	OidcClientSecret *string `pulumi:"oidcClientSecret"`
	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
	OidcDiscoveryCaPem *string `pulumi:"oidcDiscoveryCaPem"`
	// The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwtValidationPubkeys`
	OidcDiscoveryUrl *string `pulumi:"oidcDiscoveryUrl"`
	// Path to mount the JWT/OIDC auth backend
	Path *string `pulumi:"path"`
	Tune *AuthBackendTune `pulumi:"tune"`
	// Type of auth backend. Should be one of `jwt` or `oidc`. Default - `jwt`
	Type *string `pulumi:"type"`
}

// The set of arguments for constructing a AuthBackend resource.
type AuthBackendArgs struct {
	// The value against which to match the iss claim in a JWT
	BoundIssuer pulumi.StringPtrInput
	// The default role to use if none is provided during login
	DefaultRole pulumi.StringPtrInput
	// The description of the auth backend
	Description pulumi.StringPtrInput
	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
	JwksCaPem pulumi.StringPtrInput
	// JWKS URL to use to authenticate signatures. Cannot be used with "oidcDiscoveryUrl" or "jwtValidationPubkeys".
	JwksUrl pulumi.StringPtrInput
	// A list of supported signing algorithms. Vault 1.1.0 defaults to [RS256] but future or past versions of Vault may differ
	JwtSupportedAlgs pulumi.StringArrayInput
	// A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used in combination with `oidcDiscoveryUrl`
	JwtValidationPubkeys pulumi.StringArrayInput
	// Client ID used for OIDC backends
	OidcClientId pulumi.StringPtrInput
	// Client Secret used for OIDC backends
	OidcClientSecret pulumi.StringPtrInput
	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
	OidcDiscoveryCaPem pulumi.StringPtrInput
	// The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwtValidationPubkeys`
	OidcDiscoveryUrl pulumi.StringPtrInput
	// Path to mount the JWT/OIDC auth backend
	Path pulumi.StringPtrInput
	Tune AuthBackendTunePtrInput
	// Type of auth backend. Should be one of `jwt` or `oidc`. Default - `jwt`
	Type pulumi.StringPtrInput
}

func (AuthBackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendArgs)(nil)).Elem()
}

