// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package jwt

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource for managing an
// [JWT auth backend within Vault](https://www.vaultproject.io/docs/auth/jwt.html).
//
// ## Example Usage
//
// Manage JWT auth backend:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/jwt"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := jwt.NewAuthBackend(ctx, "example", &jwt.AuthBackendArgs{
// 			BoundIssuer:      pulumi.String("https://myco.auth0.com/"),
// 			Description:      pulumi.String("Demonstration of the Terraform JWT auth backend"),
// 			OidcDiscoveryUrl: pulumi.String("https://myco.auth0.com/"),
// 			Path:             pulumi.String("jwt"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Manage OIDC auth backend:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/jwt"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := jwt.NewAuthBackend(ctx, "example", &jwt.AuthBackendArgs{
// 			BoundIssuer:      pulumi.String("https://myco.auth0.com/"),
// 			Description:      pulumi.String("Demonstration of the Terraform JWT auth backend"),
// 			OidcClientId:     pulumi.String("1234567890"),
// 			OidcClientSecret: pulumi.String("secret123456"),
// 			OidcDiscoveryUrl: pulumi.String("https://myco.auth0.com/"),
// 			Path:             pulumi.String("oidc"),
// 			Tune: &jwt.AuthBackendTuneArgs{
// 				ListingVisibility: pulumi.String("unauth"),
// 			},
// 			Type: pulumi.String("oidc"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// Configuring the auth backend with a `provider_config:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/jwt"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := jwt.NewAuthBackend(ctx, "gsuite", &jwt.AuthBackendArgs{
// 			Description:      pulumi.String("OIDC backend"),
// 			OidcDiscoveryUrl: pulumi.String("https://accounts.google.com"),
// 			Path:             pulumi.String("oidc"),
// 			ProviderConfig: pulumi.StringMap{
// 				"fetch_groups":             pulumi.String("true"),
// 				"fetch_user_info":          pulumi.String("true"),
// 				"groups_recurse_max_depth": pulumi.String("1"),
// 				"provider":                 pulumi.String("gsuite"),
// 			},
// 			Type: pulumi.String("oidc"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// JWT auth backend can be imported using the `type`, e.g.
//
// ```sh
//  $ pulumi import vault:jwt/authBackend:AuthBackend oidc oidc
// ```
//
//  or
//
// ```sh
//  $ pulumi import vault:jwt/authBackend:AuthBackend jwt jwt
// ```
type AuthBackend struct {
	pulumi.CustomResourceState

	// The accessor for this auth method
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
	// Specifies if the auth method is local only.
	Local pulumi.BoolPtrOutput `pulumi:"local"`
	// Pass namespace in the OIDC state parameter instead of as a separate query parameter. With this setting, the allowed redirect URL(s) in Vault and on the provider side should not contain a namespace query parameter. This means only one redirect URL entry needs to be maintained on the OIDC provider side for all vault namespaces that will be authenticating against it. Defaults to true for new configs
	NamespaceInState pulumi.BoolPtrOutput `pulumi:"namespaceInState"`
	// Client ID used for OIDC backends
	OidcClientId pulumi.StringPtrOutput `pulumi:"oidcClientId"`
	// Client Secret used for OIDC backends
	OidcClientSecret pulumi.StringPtrOutput `pulumi:"oidcClientSecret"`
	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
	OidcDiscoveryCaPem pulumi.StringPtrOutput `pulumi:"oidcDiscoveryCaPem"`
	// The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwtValidationPubkeys`
	OidcDiscoveryUrl pulumi.StringPtrOutput `pulumi:"oidcDiscoveryUrl"`
	// The response mode to be used in the OAuth2 request. Allowed values are `query` and `formPost`. Defaults to `query`. If using Vault namespaces, and `oidcResponseMode` is `formPost`, then `namespaceInState` should be set to `false`.
	OidcResponseMode pulumi.StringPtrOutput `pulumi:"oidcResponseMode"`
	// List of response types to request. Allowed values are 'code' and 'id_token'. Defaults to `["code"]`. Note: `idToken` may only be used if `oidcResponseMode` is set to `formPost`.
	OidcResponseTypes pulumi.StringArrayOutput `pulumi:"oidcResponseTypes"`
	// Path to mount the JWT/OIDC auth backend
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// Provider specific handling configuration. All values may be strings, and the provider will convert to the appropriate type when configuring Vault.
	ProviderConfig pulumi.StringMapOutput `pulumi:"providerConfig"`
	Tune           AuthBackendTuneOutput  `pulumi:"tune"`
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
	// The accessor for this auth method
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
	// Specifies if the auth method is local only.
	Local *bool `pulumi:"local"`
	// Pass namespace in the OIDC state parameter instead of as a separate query parameter. With this setting, the allowed redirect URL(s) in Vault and on the provider side should not contain a namespace query parameter. This means only one redirect URL entry needs to be maintained on the OIDC provider side for all vault namespaces that will be authenticating against it. Defaults to true for new configs
	NamespaceInState *bool `pulumi:"namespaceInState"`
	// Client ID used for OIDC backends
	OidcClientId *string `pulumi:"oidcClientId"`
	// Client Secret used for OIDC backends
	OidcClientSecret *string `pulumi:"oidcClientSecret"`
	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
	OidcDiscoveryCaPem *string `pulumi:"oidcDiscoveryCaPem"`
	// The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwtValidationPubkeys`
	OidcDiscoveryUrl *string `pulumi:"oidcDiscoveryUrl"`
	// The response mode to be used in the OAuth2 request. Allowed values are `query` and `formPost`. Defaults to `query`. If using Vault namespaces, and `oidcResponseMode` is `formPost`, then `namespaceInState` should be set to `false`.
	OidcResponseMode *string `pulumi:"oidcResponseMode"`
	// List of response types to request. Allowed values are 'code' and 'id_token'. Defaults to `["code"]`. Note: `idToken` may only be used if `oidcResponseMode` is set to `formPost`.
	OidcResponseTypes []string `pulumi:"oidcResponseTypes"`
	// Path to mount the JWT/OIDC auth backend
	Path *string `pulumi:"path"`
	// Provider specific handling configuration. All values may be strings, and the provider will convert to the appropriate type when configuring Vault.
	ProviderConfig map[string]string `pulumi:"providerConfig"`
	Tune           *AuthBackendTune  `pulumi:"tune"`
	// Type of auth backend. Should be one of `jwt` or `oidc`. Default - `jwt`
	Type *string `pulumi:"type"`
}

type AuthBackendState struct {
	// The accessor for this auth method
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
	// Specifies if the auth method is local only.
	Local pulumi.BoolPtrInput
	// Pass namespace in the OIDC state parameter instead of as a separate query parameter. With this setting, the allowed redirect URL(s) in Vault and on the provider side should not contain a namespace query parameter. This means only one redirect URL entry needs to be maintained on the OIDC provider side for all vault namespaces that will be authenticating against it. Defaults to true for new configs
	NamespaceInState pulumi.BoolPtrInput
	// Client ID used for OIDC backends
	OidcClientId pulumi.StringPtrInput
	// Client Secret used for OIDC backends
	OidcClientSecret pulumi.StringPtrInput
	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
	OidcDiscoveryCaPem pulumi.StringPtrInput
	// The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwtValidationPubkeys`
	OidcDiscoveryUrl pulumi.StringPtrInput
	// The response mode to be used in the OAuth2 request. Allowed values are `query` and `formPost`. Defaults to `query`. If using Vault namespaces, and `oidcResponseMode` is `formPost`, then `namespaceInState` should be set to `false`.
	OidcResponseMode pulumi.StringPtrInput
	// List of response types to request. Allowed values are 'code' and 'id_token'. Defaults to `["code"]`. Note: `idToken` may only be used if `oidcResponseMode` is set to `formPost`.
	OidcResponseTypes pulumi.StringArrayInput
	// Path to mount the JWT/OIDC auth backend
	Path pulumi.StringPtrInput
	// Provider specific handling configuration. All values may be strings, and the provider will convert to the appropriate type when configuring Vault.
	ProviderConfig pulumi.StringMapInput
	Tune           AuthBackendTunePtrInput
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
	// Specifies if the auth method is local only.
	Local *bool `pulumi:"local"`
	// Pass namespace in the OIDC state parameter instead of as a separate query parameter. With this setting, the allowed redirect URL(s) in Vault and on the provider side should not contain a namespace query parameter. This means only one redirect URL entry needs to be maintained on the OIDC provider side for all vault namespaces that will be authenticating against it. Defaults to true for new configs
	NamespaceInState *bool `pulumi:"namespaceInState"`
	// Client ID used for OIDC backends
	OidcClientId *string `pulumi:"oidcClientId"`
	// Client Secret used for OIDC backends
	OidcClientSecret *string `pulumi:"oidcClientSecret"`
	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
	OidcDiscoveryCaPem *string `pulumi:"oidcDiscoveryCaPem"`
	// The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwtValidationPubkeys`
	OidcDiscoveryUrl *string `pulumi:"oidcDiscoveryUrl"`
	// The response mode to be used in the OAuth2 request. Allowed values are `query` and `formPost`. Defaults to `query`. If using Vault namespaces, and `oidcResponseMode` is `formPost`, then `namespaceInState` should be set to `false`.
	OidcResponseMode *string `pulumi:"oidcResponseMode"`
	// List of response types to request. Allowed values are 'code' and 'id_token'. Defaults to `["code"]`. Note: `idToken` may only be used if `oidcResponseMode` is set to `formPost`.
	OidcResponseTypes []string `pulumi:"oidcResponseTypes"`
	// Path to mount the JWT/OIDC auth backend
	Path *string `pulumi:"path"`
	// Provider specific handling configuration. All values may be strings, and the provider will convert to the appropriate type when configuring Vault.
	ProviderConfig map[string]string `pulumi:"providerConfig"`
	Tune           *AuthBackendTune  `pulumi:"tune"`
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
	// Specifies if the auth method is local only.
	Local pulumi.BoolPtrInput
	// Pass namespace in the OIDC state parameter instead of as a separate query parameter. With this setting, the allowed redirect URL(s) in Vault and on the provider side should not contain a namespace query parameter. This means only one redirect URL entry needs to be maintained on the OIDC provider side for all vault namespaces that will be authenticating against it. Defaults to true for new configs
	NamespaceInState pulumi.BoolPtrInput
	// Client ID used for OIDC backends
	OidcClientId pulumi.StringPtrInput
	// Client Secret used for OIDC backends
	OidcClientSecret pulumi.StringPtrInput
	// The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
	OidcDiscoveryCaPem pulumi.StringPtrInput
	// The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwtValidationPubkeys`
	OidcDiscoveryUrl pulumi.StringPtrInput
	// The response mode to be used in the OAuth2 request. Allowed values are `query` and `formPost`. Defaults to `query`. If using Vault namespaces, and `oidcResponseMode` is `formPost`, then `namespaceInState` should be set to `false`.
	OidcResponseMode pulumi.StringPtrInput
	// List of response types to request. Allowed values are 'code' and 'id_token'. Defaults to `["code"]`. Note: `idToken` may only be used if `oidcResponseMode` is set to `formPost`.
	OidcResponseTypes pulumi.StringArrayInput
	// Path to mount the JWT/OIDC auth backend
	Path pulumi.StringPtrInput
	// Provider specific handling configuration. All values may be strings, and the provider will convert to the appropriate type when configuring Vault.
	ProviderConfig pulumi.StringMapInput
	Tune           AuthBackendTunePtrInput
	// Type of auth backend. Should be one of `jwt` or `oidc`. Default - `jwt`
	Type pulumi.StringPtrInput
}

func (AuthBackendArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendArgs)(nil)).Elem()
}

type AuthBackendInput interface {
	pulumi.Input

	ToAuthBackendOutput() AuthBackendOutput
	ToAuthBackendOutputWithContext(ctx context.Context) AuthBackendOutput
}

func (*AuthBackend) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthBackend)(nil))
}

func (i *AuthBackend) ToAuthBackendOutput() AuthBackendOutput {
	return i.ToAuthBackendOutputWithContext(context.Background())
}

func (i *AuthBackend) ToAuthBackendOutputWithContext(ctx context.Context) AuthBackendOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendOutput)
}

func (i *AuthBackend) ToAuthBackendPtrOutput() AuthBackendPtrOutput {
	return i.ToAuthBackendPtrOutputWithContext(context.Background())
}

func (i *AuthBackend) ToAuthBackendPtrOutputWithContext(ctx context.Context) AuthBackendPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendPtrOutput)
}

type AuthBackendPtrInput interface {
	pulumi.Input

	ToAuthBackendPtrOutput() AuthBackendPtrOutput
	ToAuthBackendPtrOutputWithContext(ctx context.Context) AuthBackendPtrOutput
}

type authBackendPtrType AuthBackendArgs

func (*authBackendPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackend)(nil))
}

func (i *authBackendPtrType) ToAuthBackendPtrOutput() AuthBackendPtrOutput {
	return i.ToAuthBackendPtrOutputWithContext(context.Background())
}

func (i *authBackendPtrType) ToAuthBackendPtrOutputWithContext(ctx context.Context) AuthBackendPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendPtrOutput)
}

// AuthBackendArrayInput is an input type that accepts AuthBackendArray and AuthBackendArrayOutput values.
// You can construct a concrete instance of `AuthBackendArrayInput` via:
//
//          AuthBackendArray{ AuthBackendArgs{...} }
type AuthBackendArrayInput interface {
	pulumi.Input

	ToAuthBackendArrayOutput() AuthBackendArrayOutput
	ToAuthBackendArrayOutputWithContext(context.Context) AuthBackendArrayOutput
}

type AuthBackendArray []AuthBackendInput

func (AuthBackendArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackend)(nil)).Elem()
}

func (i AuthBackendArray) ToAuthBackendArrayOutput() AuthBackendArrayOutput {
	return i.ToAuthBackendArrayOutputWithContext(context.Background())
}

func (i AuthBackendArray) ToAuthBackendArrayOutputWithContext(ctx context.Context) AuthBackendArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendArrayOutput)
}

// AuthBackendMapInput is an input type that accepts AuthBackendMap and AuthBackendMapOutput values.
// You can construct a concrete instance of `AuthBackendMapInput` via:
//
//          AuthBackendMap{ "key": AuthBackendArgs{...} }
type AuthBackendMapInput interface {
	pulumi.Input

	ToAuthBackendMapOutput() AuthBackendMapOutput
	ToAuthBackendMapOutputWithContext(context.Context) AuthBackendMapOutput
}

type AuthBackendMap map[string]AuthBackendInput

func (AuthBackendMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackend)(nil)).Elem()
}

func (i AuthBackendMap) ToAuthBackendMapOutput() AuthBackendMapOutput {
	return i.ToAuthBackendMapOutputWithContext(context.Background())
}

func (i AuthBackendMap) ToAuthBackendMapOutputWithContext(ctx context.Context) AuthBackendMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendMapOutput)
}

type AuthBackendOutput struct{ *pulumi.OutputState }

func (AuthBackendOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AuthBackend)(nil))
}

func (o AuthBackendOutput) ToAuthBackendOutput() AuthBackendOutput {
	return o
}

func (o AuthBackendOutput) ToAuthBackendOutputWithContext(ctx context.Context) AuthBackendOutput {
	return o
}

func (o AuthBackendOutput) ToAuthBackendPtrOutput() AuthBackendPtrOutput {
	return o.ToAuthBackendPtrOutputWithContext(context.Background())
}

func (o AuthBackendOutput) ToAuthBackendPtrOutputWithContext(ctx context.Context) AuthBackendPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AuthBackend) *AuthBackend {
		return &v
	}).(AuthBackendPtrOutput)
}

type AuthBackendPtrOutput struct{ *pulumi.OutputState }

func (AuthBackendPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackend)(nil))
}

func (o AuthBackendPtrOutput) ToAuthBackendPtrOutput() AuthBackendPtrOutput {
	return o
}

func (o AuthBackendPtrOutput) ToAuthBackendPtrOutputWithContext(ctx context.Context) AuthBackendPtrOutput {
	return o
}

func (o AuthBackendPtrOutput) Elem() AuthBackendOutput {
	return o.ApplyT(func(v *AuthBackend) AuthBackend {
		if v != nil {
			return *v
		}
		var ret AuthBackend
		return ret
	}).(AuthBackendOutput)
}

type AuthBackendArrayOutput struct{ *pulumi.OutputState }

func (AuthBackendArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AuthBackend)(nil))
}

func (o AuthBackendArrayOutput) ToAuthBackendArrayOutput() AuthBackendArrayOutput {
	return o
}

func (o AuthBackendArrayOutput) ToAuthBackendArrayOutputWithContext(ctx context.Context) AuthBackendArrayOutput {
	return o
}

func (o AuthBackendArrayOutput) Index(i pulumi.IntInput) AuthBackendOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AuthBackend {
		return vs[0].([]AuthBackend)[vs[1].(int)]
	}).(AuthBackendOutput)
}

type AuthBackendMapOutput struct{ *pulumi.OutputState }

func (AuthBackendMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AuthBackend)(nil))
}

func (o AuthBackendMapOutput) ToAuthBackendMapOutput() AuthBackendMapOutput {
	return o
}

func (o AuthBackendMapOutput) ToAuthBackendMapOutputWithContext(ctx context.Context) AuthBackendMapOutput {
	return o
}

func (o AuthBackendMapOutput) MapIndex(k pulumi.StringInput) AuthBackendOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AuthBackend {
		return vs[0].(map[string]AuthBackend)[vs[1].(string)]
	}).(AuthBackendOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendInput)(nil)).Elem(), &AuthBackend{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendPtrInput)(nil)).Elem(), &AuthBackend{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendArrayInput)(nil)).Elem(), AuthBackendArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendMapInput)(nil)).Elem(), AuthBackendMap{})
	pulumi.RegisterOutputType(AuthBackendOutput{})
	pulumi.RegisterOutputType(AuthBackendPtrOutput{})
	pulumi.RegisterOutputType(AuthBackendArrayOutput{})
	pulumi.RegisterOutputType(AuthBackendMapOutput{})
}
