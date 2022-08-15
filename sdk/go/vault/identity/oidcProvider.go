// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package identity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages OIDC Providers in a Vault server. See the [Vault documentation](https://www.vaultproject.io/api-docs/secret/identity/oidc-provider#create-or-update-an-assignment)
// for more information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/identity"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			testOidcKey, err := identity.NewOidcKey(ctx, "testOidcKey", &identity.OidcKeyArgs{
//				AllowedClientIds: pulumi.StringArray{
//					pulumi.String("*"),
//				},
//				RotationPeriod:  pulumi.Int(3600),
//				VerificationTtl: pulumi.Int(3600),
//			})
//			if err != nil {
//				return err
//			}
//			testOidcAssignment, err := identity.NewOidcAssignment(ctx, "testOidcAssignment", &identity.OidcAssignmentArgs{
//				EntityIds: pulumi.StringArray{
//					pulumi.String("fake-ascbascas-2231a-sdfaa"),
//				},
//				GroupIds: pulumi.StringArray{
//					pulumi.String("fake-sajkdsad-32414-sfsada"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			testOidcClient, err := identity.NewOidcClient(ctx, "testOidcClient", &identity.OidcClientArgs{
//				Key: testOidcKey.Name,
//				RedirectUris: pulumi.StringArray{
//					pulumi.String("http://127.0.0.1:9200/v1/auth-methods/oidc:authenticate:callback"),
//					pulumi.String("http://127.0.0.1:8251/callback"),
//					pulumi.String("http://127.0.0.1:8080/callback"),
//				},
//				Assignments: pulumi.StringArray{
//					testOidcAssignment.Name,
//				},
//				IdTokenTtl:     pulumi.Int(2400),
//				AccessTokenTtl: pulumi.Int(7200),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"groups": "{{identity.entity.groups.names}}",
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			testOidcScope, err := identity.NewOidcScope(ctx, "testOidcScope", &identity.OidcScopeArgs{
//				Template:    pulumi.String(json0),
//				Description: pulumi.String("Groups scope."),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = identity.NewOidcProvider(ctx, "testOidcProvider", &identity.OidcProviderArgs{
//				HttpsEnabled: pulumi.Bool(false),
//				IssuerHost:   pulumi.String("127.0.0.1:8200"),
//				AllowedClientIds: pulumi.StringArray{
//					testOidcClient.ClientId,
//				},
//				ScopesSupporteds: pulumi.StringArray{
//					testOidcScope.Name,
//				},
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
// OIDC Providers can be imported using the `name`, e.g.
//
// ```sh
//
//	$ pulumi import vault:identity/oidcProvider:OidcProvider test my-provider
//
// ```
type OidcProvider struct {
	pulumi.CustomResourceState

	// The client IDs that are permitted to use the provider.
	// If empty, no clients are allowed. If `*`, all clients are allowed.
	AllowedClientIds pulumi.StringArrayOutput `pulumi:"allowedClientIds"`
	// Set to true if the issuer endpoint uses HTTPS.
	HttpsEnabled pulumi.BoolPtrOutput `pulumi:"httpsEnabled"`
	// Specifies what will be used as the `scheme://host:port`
	// component for the `iss` claim of ID tokens. This value is computed using the
	// `issuerHost` and `httpsEnabled` fields.
	Issuer pulumi.StringOutput `pulumi:"issuer"`
	// The host for the issuer. Can be either host or host:port.
	IssuerHost pulumi.StringPtrOutput `pulumi:"issuerHost"`
	// The name of the provider.
	Name pulumi.StringOutput `pulumi:"name"`
	// The scopes available for requesting on the provider.
	ScopesSupporteds pulumi.StringArrayOutput `pulumi:"scopesSupporteds"`
}

// NewOidcProvider registers a new resource with the given unique name, arguments, and options.
func NewOidcProvider(ctx *pulumi.Context,
	name string, args *OidcProviderArgs, opts ...pulumi.ResourceOption) (*OidcProvider, error) {
	if args == nil {
		args = &OidcProviderArgs{}
	}

	var resource OidcProvider
	err := ctx.RegisterResource("vault:identity/oidcProvider:OidcProvider", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOidcProvider gets an existing OidcProvider resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOidcProvider(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OidcProviderState, opts ...pulumi.ResourceOption) (*OidcProvider, error) {
	var resource OidcProvider
	err := ctx.ReadResource("vault:identity/oidcProvider:OidcProvider", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OidcProvider resources.
type oidcProviderState struct {
	// The client IDs that are permitted to use the provider.
	// If empty, no clients are allowed. If `*`, all clients are allowed.
	AllowedClientIds []string `pulumi:"allowedClientIds"`
	// Set to true if the issuer endpoint uses HTTPS.
	HttpsEnabled *bool `pulumi:"httpsEnabled"`
	// Specifies what will be used as the `scheme://host:port`
	// component for the `iss` claim of ID tokens. This value is computed using the
	// `issuerHost` and `httpsEnabled` fields.
	Issuer *string `pulumi:"issuer"`
	// The host for the issuer. Can be either host or host:port.
	IssuerHost *string `pulumi:"issuerHost"`
	// The name of the provider.
	Name *string `pulumi:"name"`
	// The scopes available for requesting on the provider.
	ScopesSupporteds []string `pulumi:"scopesSupporteds"`
}

type OidcProviderState struct {
	// The client IDs that are permitted to use the provider.
	// If empty, no clients are allowed. If `*`, all clients are allowed.
	AllowedClientIds pulumi.StringArrayInput
	// Set to true if the issuer endpoint uses HTTPS.
	HttpsEnabled pulumi.BoolPtrInput
	// Specifies what will be used as the `scheme://host:port`
	// component for the `iss` claim of ID tokens. This value is computed using the
	// `issuerHost` and `httpsEnabled` fields.
	Issuer pulumi.StringPtrInput
	// The host for the issuer. Can be either host or host:port.
	IssuerHost pulumi.StringPtrInput
	// The name of the provider.
	Name pulumi.StringPtrInput
	// The scopes available for requesting on the provider.
	ScopesSupporteds pulumi.StringArrayInput
}

func (OidcProviderState) ElementType() reflect.Type {
	return reflect.TypeOf((*oidcProviderState)(nil)).Elem()
}

type oidcProviderArgs struct {
	// The client IDs that are permitted to use the provider.
	// If empty, no clients are allowed. If `*`, all clients are allowed.
	AllowedClientIds []string `pulumi:"allowedClientIds"`
	// Set to true if the issuer endpoint uses HTTPS.
	HttpsEnabled *bool `pulumi:"httpsEnabled"`
	// The host for the issuer. Can be either host or host:port.
	IssuerHost *string `pulumi:"issuerHost"`
	// The name of the provider.
	Name *string `pulumi:"name"`
	// The scopes available for requesting on the provider.
	ScopesSupporteds []string `pulumi:"scopesSupporteds"`
}

// The set of arguments for constructing a OidcProvider resource.
type OidcProviderArgs struct {
	// The client IDs that are permitted to use the provider.
	// If empty, no clients are allowed. If `*`, all clients are allowed.
	AllowedClientIds pulumi.StringArrayInput
	// Set to true if the issuer endpoint uses HTTPS.
	HttpsEnabled pulumi.BoolPtrInput
	// The host for the issuer. Can be either host or host:port.
	IssuerHost pulumi.StringPtrInput
	// The name of the provider.
	Name pulumi.StringPtrInput
	// The scopes available for requesting on the provider.
	ScopesSupporteds pulumi.StringArrayInput
}

func (OidcProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*oidcProviderArgs)(nil)).Elem()
}

type OidcProviderInput interface {
	pulumi.Input

	ToOidcProviderOutput() OidcProviderOutput
	ToOidcProviderOutputWithContext(ctx context.Context) OidcProviderOutput
}

func (*OidcProvider) ElementType() reflect.Type {
	return reflect.TypeOf((**OidcProvider)(nil)).Elem()
}

func (i *OidcProvider) ToOidcProviderOutput() OidcProviderOutput {
	return i.ToOidcProviderOutputWithContext(context.Background())
}

func (i *OidcProvider) ToOidcProviderOutputWithContext(ctx context.Context) OidcProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OidcProviderOutput)
}

// OidcProviderArrayInput is an input type that accepts OidcProviderArray and OidcProviderArrayOutput values.
// You can construct a concrete instance of `OidcProviderArrayInput` via:
//
//	OidcProviderArray{ OidcProviderArgs{...} }
type OidcProviderArrayInput interface {
	pulumi.Input

	ToOidcProviderArrayOutput() OidcProviderArrayOutput
	ToOidcProviderArrayOutputWithContext(context.Context) OidcProviderArrayOutput
}

type OidcProviderArray []OidcProviderInput

func (OidcProviderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OidcProvider)(nil)).Elem()
}

func (i OidcProviderArray) ToOidcProviderArrayOutput() OidcProviderArrayOutput {
	return i.ToOidcProviderArrayOutputWithContext(context.Background())
}

func (i OidcProviderArray) ToOidcProviderArrayOutputWithContext(ctx context.Context) OidcProviderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OidcProviderArrayOutput)
}

// OidcProviderMapInput is an input type that accepts OidcProviderMap and OidcProviderMapOutput values.
// You can construct a concrete instance of `OidcProviderMapInput` via:
//
//	OidcProviderMap{ "key": OidcProviderArgs{...} }
type OidcProviderMapInput interface {
	pulumi.Input

	ToOidcProviderMapOutput() OidcProviderMapOutput
	ToOidcProviderMapOutputWithContext(context.Context) OidcProviderMapOutput
}

type OidcProviderMap map[string]OidcProviderInput

func (OidcProviderMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OidcProvider)(nil)).Elem()
}

func (i OidcProviderMap) ToOidcProviderMapOutput() OidcProviderMapOutput {
	return i.ToOidcProviderMapOutputWithContext(context.Background())
}

func (i OidcProviderMap) ToOidcProviderMapOutputWithContext(ctx context.Context) OidcProviderMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OidcProviderMapOutput)
}

type OidcProviderOutput struct{ *pulumi.OutputState }

func (OidcProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OidcProvider)(nil)).Elem()
}

func (o OidcProviderOutput) ToOidcProviderOutput() OidcProviderOutput {
	return o
}

func (o OidcProviderOutput) ToOidcProviderOutputWithContext(ctx context.Context) OidcProviderOutput {
	return o
}

// The client IDs that are permitted to use the provider.
// If empty, no clients are allowed. If `*`, all clients are allowed.
func (o OidcProviderOutput) AllowedClientIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OidcProvider) pulumi.StringArrayOutput { return v.AllowedClientIds }).(pulumi.StringArrayOutput)
}

// Set to true if the issuer endpoint uses HTTPS.
func (o OidcProviderOutput) HttpsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *OidcProvider) pulumi.BoolPtrOutput { return v.HttpsEnabled }).(pulumi.BoolPtrOutput)
}

// Specifies what will be used as the `scheme://host:port`
// component for the `iss` claim of ID tokens. This value is computed using the
// `issuerHost` and `httpsEnabled` fields.
func (o OidcProviderOutput) Issuer() pulumi.StringOutput {
	return o.ApplyT(func(v *OidcProvider) pulumi.StringOutput { return v.Issuer }).(pulumi.StringOutput)
}

// The host for the issuer. Can be either host or host:port.
func (o OidcProviderOutput) IssuerHost() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OidcProvider) pulumi.StringPtrOutput { return v.IssuerHost }).(pulumi.StringPtrOutput)
}

// The name of the provider.
func (o OidcProviderOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *OidcProvider) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The scopes available for requesting on the provider.
func (o OidcProviderOutput) ScopesSupporteds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *OidcProvider) pulumi.StringArrayOutput { return v.ScopesSupporteds }).(pulumi.StringArrayOutput)
}

type OidcProviderArrayOutput struct{ *pulumi.OutputState }

func (OidcProviderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OidcProvider)(nil)).Elem()
}

func (o OidcProviderArrayOutput) ToOidcProviderArrayOutput() OidcProviderArrayOutput {
	return o
}

func (o OidcProviderArrayOutput) ToOidcProviderArrayOutputWithContext(ctx context.Context) OidcProviderArrayOutput {
	return o
}

func (o OidcProviderArrayOutput) Index(i pulumi.IntInput) OidcProviderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OidcProvider {
		return vs[0].([]*OidcProvider)[vs[1].(int)]
	}).(OidcProviderOutput)
}

type OidcProviderMapOutput struct{ *pulumi.OutputState }

func (OidcProviderMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OidcProvider)(nil)).Elem()
}

func (o OidcProviderMapOutput) ToOidcProviderMapOutput() OidcProviderMapOutput {
	return o
}

func (o OidcProviderMapOutput) ToOidcProviderMapOutputWithContext(ctx context.Context) OidcProviderMapOutput {
	return o
}

func (o OidcProviderMapOutput) MapIndex(k pulumi.StringInput) OidcProviderOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OidcProvider {
		return vs[0].(map[string]*OidcProvider)[vs[1].(string)]
	}).(OidcProviderOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OidcProviderInput)(nil)).Elem(), &OidcProvider{})
	pulumi.RegisterInputType(reflect.TypeOf((*OidcProviderArrayInput)(nil)).Elem(), OidcProviderArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OidcProviderMapInput)(nil)).Elem(), OidcProviderMap{})
	pulumi.RegisterOutputType(OidcProviderOutput{})
	pulumi.RegisterOutputType(OidcProviderArrayOutput{})
	pulumi.RegisterOutputType(OidcProviderMapOutput{})
}
