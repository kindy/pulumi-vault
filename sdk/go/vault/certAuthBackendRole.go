// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a resource to create a role in an [Cert auth backend within Vault](https://www.vaultproject.io/docs/auth/cert.html).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"io/ioutil"
//
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func readFileOrPanic(path string) pulumi.StringPtrInput {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	return pulumi.String(string(data))
// }
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		certAuthBackend, err := vault.NewAuthBackend(ctx, "certAuthBackend", &vault.AuthBackendArgs{
// 			Path: pulumi.String("cert"),
// 			Type: pulumi.String("cert"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vault.NewCertAuthBackendRole(ctx, "certCertAuthBackendRole", &vault.CertAuthBackendRoleArgs{
// 			Certificate: readFileOrPanic("/path/to/certs/ca-cert.pem"),
// 			Backend:     certAuthBackend.Path,
// 			AllowedNames: pulumi.StringArray{
// 				pulumi.String("foo.example.org"),
// 				pulumi.String("baz.example.org"),
// 			},
// 			TokenTtl:    pulumi.Int(300),
// 			TokenMaxTtl: pulumi.Int(600),
// 			TokenPolicies: pulumi.StringArray{
// 				pulumi.String("foo"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type CertAuthBackendRole struct {
	pulumi.CustomResourceState

	// Allowed the common names for authenticated client certificates
	AllowedCommonNames pulumi.StringArrayOutput `pulumi:"allowedCommonNames"`
	// Allowed alternative dns names for authenticated client certificates
	AllowedDnsSans pulumi.StringArrayOutput `pulumi:"allowedDnsSans"`
	// Allowed emails for authenticated client certificates
	AllowedEmailSans pulumi.StringArrayOutput `pulumi:"allowedEmailSans"`
	// Allowed subject names for authenticated client certificates
	AllowedNames pulumi.StringArrayOutput `pulumi:"allowedNames"`
	// Allowed organization units for authenticated client certificates
	AllowedOrganizationUnits pulumi.StringArrayOutput `pulumi:"allowedOrganizationUnits"`
	// Allowed URIs for authenticated client certificates
	AllowedUriSans pulumi.StringArrayOutput `pulumi:"allowedUriSans"`
	// Path to the mounted Cert auth backend
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// CA certificate used to validate client certificates
	Certificate pulumi.StringOutput `pulumi:"certificate"`
	// The name to display on tokens issued under this role.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Name of the role
	Name pulumi.StringOutput `pulumi:"name"`
	// TLS extensions required on client certificates
	RequiredExtensions pulumi.StringArrayOutput `pulumi:"requiredExtensions"`
	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs pulumi.StringArrayOutput `pulumi:"tokenBoundCidrs"`
	// If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl pulumi.IntPtrOutput `pulumi:"tokenExplicitMaxTtl"`
	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl pulumi.IntPtrOutput `pulumi:"tokenMaxTtl"`
	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy pulumi.BoolPtrOutput `pulumi:"tokenNoDefaultPolicy"`
	// The
	// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
	// if any, in number of seconds to set on the token.
	TokenNumUses pulumi.IntPtrOutput `pulumi:"tokenNumUses"`
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod pulumi.IntPtrOutput `pulumi:"tokenPeriod"`
	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies pulumi.StringArrayOutput `pulumi:"tokenPolicies"`
	// The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl pulumi.IntPtrOutput `pulumi:"tokenTtl"`
	// The type of token that should be generated. Can be `service`,
	// `batch`, or `default` to use the mount's tuned default (which unless changed will be
	// `service` tokens). For token store roles, there are two additional possibilities:
	// `default-service` and `default-batch` which specify the type to return unless the client
	// requests a different type at generation time.
	TokenType pulumi.StringPtrOutput `pulumi:"tokenType"`
}

// NewCertAuthBackendRole registers a new resource with the given unique name, arguments, and options.
func NewCertAuthBackendRole(ctx *pulumi.Context,
	name string, args *CertAuthBackendRoleArgs, opts ...pulumi.ResourceOption) (*CertAuthBackendRole, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Certificate == nil {
		return nil, errors.New("invalid value for required argument 'Certificate'")
	}
	var resource CertAuthBackendRole
	err := ctx.RegisterResource("vault:index/certAuthBackendRole:CertAuthBackendRole", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCertAuthBackendRole gets an existing CertAuthBackendRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCertAuthBackendRole(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CertAuthBackendRoleState, opts ...pulumi.ResourceOption) (*CertAuthBackendRole, error) {
	var resource CertAuthBackendRole
	err := ctx.ReadResource("vault:index/certAuthBackendRole:CertAuthBackendRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering CertAuthBackendRole resources.
type certAuthBackendRoleState struct {
	// Allowed the common names for authenticated client certificates
	AllowedCommonNames []string `pulumi:"allowedCommonNames"`
	// Allowed alternative dns names for authenticated client certificates
	AllowedDnsSans []string `pulumi:"allowedDnsSans"`
	// Allowed emails for authenticated client certificates
	AllowedEmailSans []string `pulumi:"allowedEmailSans"`
	// Allowed subject names for authenticated client certificates
	AllowedNames []string `pulumi:"allowedNames"`
	// Allowed organization units for authenticated client certificates
	AllowedOrganizationUnits []string `pulumi:"allowedOrganizationUnits"`
	// Allowed URIs for authenticated client certificates
	AllowedUriSans []string `pulumi:"allowedUriSans"`
	// Path to the mounted Cert auth backend
	Backend *string `pulumi:"backend"`
	// CA certificate used to validate client certificates
	Certificate *string `pulumi:"certificate"`
	// The name to display on tokens issued under this role.
	DisplayName *string `pulumi:"displayName"`
	// Name of the role
	Name *string `pulumi:"name"`
	// TLS extensions required on client certificates
	RequiredExtensions []string `pulumi:"requiredExtensions"`
	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs []string `pulumi:"tokenBoundCidrs"`
	// If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl *int `pulumi:"tokenExplicitMaxTtl"`
	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl *int `pulumi:"tokenMaxTtl"`
	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy *bool `pulumi:"tokenNoDefaultPolicy"`
	// The
	// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
	// if any, in number of seconds to set on the token.
	TokenNumUses *int `pulumi:"tokenNumUses"`
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod *int `pulumi:"tokenPeriod"`
	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies []string `pulumi:"tokenPolicies"`
	// The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl *int `pulumi:"tokenTtl"`
	// The type of token that should be generated. Can be `service`,
	// `batch`, or `default` to use the mount's tuned default (which unless changed will be
	// `service` tokens). For token store roles, there are two additional possibilities:
	// `default-service` and `default-batch` which specify the type to return unless the client
	// requests a different type at generation time.
	TokenType *string `pulumi:"tokenType"`
}

type CertAuthBackendRoleState struct {
	// Allowed the common names for authenticated client certificates
	AllowedCommonNames pulumi.StringArrayInput
	// Allowed alternative dns names for authenticated client certificates
	AllowedDnsSans pulumi.StringArrayInput
	// Allowed emails for authenticated client certificates
	AllowedEmailSans pulumi.StringArrayInput
	// Allowed subject names for authenticated client certificates
	AllowedNames pulumi.StringArrayInput
	// Allowed organization units for authenticated client certificates
	AllowedOrganizationUnits pulumi.StringArrayInput
	// Allowed URIs for authenticated client certificates
	AllowedUriSans pulumi.StringArrayInput
	// Path to the mounted Cert auth backend
	Backend pulumi.StringPtrInput
	// CA certificate used to validate client certificates
	Certificate pulumi.StringPtrInput
	// The name to display on tokens issued under this role.
	DisplayName pulumi.StringPtrInput
	// Name of the role
	Name pulumi.StringPtrInput
	// TLS extensions required on client certificates
	RequiredExtensions pulumi.StringArrayInput
	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs pulumi.StringArrayInput
	// If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl pulumi.IntPtrInput
	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl pulumi.IntPtrInput
	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy pulumi.BoolPtrInput
	// The
	// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
	// if any, in number of seconds to set on the token.
	TokenNumUses pulumi.IntPtrInput
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod pulumi.IntPtrInput
	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies pulumi.StringArrayInput
	// The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl pulumi.IntPtrInput
	// The type of token that should be generated. Can be `service`,
	// `batch`, or `default` to use the mount's tuned default (which unless changed will be
	// `service` tokens). For token store roles, there are two additional possibilities:
	// `default-service` and `default-batch` which specify the type to return unless the client
	// requests a different type at generation time.
	TokenType pulumi.StringPtrInput
}

func (CertAuthBackendRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*certAuthBackendRoleState)(nil)).Elem()
}

type certAuthBackendRoleArgs struct {
	// Allowed the common names for authenticated client certificates
	AllowedCommonNames []string `pulumi:"allowedCommonNames"`
	// Allowed alternative dns names for authenticated client certificates
	AllowedDnsSans []string `pulumi:"allowedDnsSans"`
	// Allowed emails for authenticated client certificates
	AllowedEmailSans []string `pulumi:"allowedEmailSans"`
	// Allowed subject names for authenticated client certificates
	AllowedNames []string `pulumi:"allowedNames"`
	// Allowed organization units for authenticated client certificates
	AllowedOrganizationUnits []string `pulumi:"allowedOrganizationUnits"`
	// Allowed URIs for authenticated client certificates
	AllowedUriSans []string `pulumi:"allowedUriSans"`
	// Path to the mounted Cert auth backend
	Backend *string `pulumi:"backend"`
	// CA certificate used to validate client certificates
	Certificate string `pulumi:"certificate"`
	// The name to display on tokens issued under this role.
	DisplayName *string `pulumi:"displayName"`
	// Name of the role
	Name *string `pulumi:"name"`
	// TLS extensions required on client certificates
	RequiredExtensions []string `pulumi:"requiredExtensions"`
	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs []string `pulumi:"tokenBoundCidrs"`
	// If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl *int `pulumi:"tokenExplicitMaxTtl"`
	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl *int `pulumi:"tokenMaxTtl"`
	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy *bool `pulumi:"tokenNoDefaultPolicy"`
	// The
	// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
	// if any, in number of seconds to set on the token.
	TokenNumUses *int `pulumi:"tokenNumUses"`
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod *int `pulumi:"tokenPeriod"`
	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies []string `pulumi:"tokenPolicies"`
	// The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl *int `pulumi:"tokenTtl"`
	// The type of token that should be generated. Can be `service`,
	// `batch`, or `default` to use the mount's tuned default (which unless changed will be
	// `service` tokens). For token store roles, there are two additional possibilities:
	// `default-service` and `default-batch` which specify the type to return unless the client
	// requests a different type at generation time.
	TokenType *string `pulumi:"tokenType"`
}

// The set of arguments for constructing a CertAuthBackendRole resource.
type CertAuthBackendRoleArgs struct {
	// Allowed the common names for authenticated client certificates
	AllowedCommonNames pulumi.StringArrayInput
	// Allowed alternative dns names for authenticated client certificates
	AllowedDnsSans pulumi.StringArrayInput
	// Allowed emails for authenticated client certificates
	AllowedEmailSans pulumi.StringArrayInput
	// Allowed subject names for authenticated client certificates
	AllowedNames pulumi.StringArrayInput
	// Allowed organization units for authenticated client certificates
	AllowedOrganizationUnits pulumi.StringArrayInput
	// Allowed URIs for authenticated client certificates
	AllowedUriSans pulumi.StringArrayInput
	// Path to the mounted Cert auth backend
	Backend pulumi.StringPtrInput
	// CA certificate used to validate client certificates
	Certificate pulumi.StringInput
	// The name to display on tokens issued under this role.
	DisplayName pulumi.StringPtrInput
	// Name of the role
	Name pulumi.StringPtrInput
	// TLS extensions required on client certificates
	RequiredExtensions pulumi.StringArrayInput
	// List of CIDR blocks; if set, specifies blocks of IP
	// addresses which can authenticate successfully, and ties the resulting token to these blocks
	// as well.
	TokenBoundCidrs pulumi.StringArrayInput
	// If set, will encode an
	// [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
	// onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
	// `tokenMaxTtl` would otherwise allow a renewal.
	TokenExplicitMaxTtl pulumi.IntPtrInput
	// The maximum lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenMaxTtl pulumi.IntPtrInput
	// If set, the default policy will not be set on
	// generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy pulumi.BoolPtrInput
	// The
	// [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
	// if any, in number of seconds to set on the token.
	TokenNumUses pulumi.IntPtrInput
	// If set, indicates that the
	// token generated using this role should never expire. The token should be renewed within the
	// duration specified by this value. At each renewal, the token's TTL will be set to the
	// value of this field. Specified in seconds.
	TokenPeriod pulumi.IntPtrInput
	// List of policies to encode onto generated tokens. Depending
	// on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies pulumi.StringArrayInput
	// The incremental lifetime for generated tokens in number of seconds.
	// Its current value will be referenced at renewal time.
	TokenTtl pulumi.IntPtrInput
	// The type of token that should be generated. Can be `service`,
	// `batch`, or `default` to use the mount's tuned default (which unless changed will be
	// `service` tokens). For token store roles, there are two additional possibilities:
	// `default-service` and `default-batch` which specify the type to return unless the client
	// requests a different type at generation time.
	TokenType pulumi.StringPtrInput
}

func (CertAuthBackendRoleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*certAuthBackendRoleArgs)(nil)).Elem()
}

type CertAuthBackendRoleInput interface {
	pulumi.Input

	ToCertAuthBackendRoleOutput() CertAuthBackendRoleOutput
	ToCertAuthBackendRoleOutputWithContext(ctx context.Context) CertAuthBackendRoleOutput
}

func (*CertAuthBackendRole) ElementType() reflect.Type {
	return reflect.TypeOf((**CertAuthBackendRole)(nil)).Elem()
}

func (i *CertAuthBackendRole) ToCertAuthBackendRoleOutput() CertAuthBackendRoleOutput {
	return i.ToCertAuthBackendRoleOutputWithContext(context.Background())
}

func (i *CertAuthBackendRole) ToCertAuthBackendRoleOutputWithContext(ctx context.Context) CertAuthBackendRoleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertAuthBackendRoleOutput)
}

// CertAuthBackendRoleArrayInput is an input type that accepts CertAuthBackendRoleArray and CertAuthBackendRoleArrayOutput values.
// You can construct a concrete instance of `CertAuthBackendRoleArrayInput` via:
//
//          CertAuthBackendRoleArray{ CertAuthBackendRoleArgs{...} }
type CertAuthBackendRoleArrayInput interface {
	pulumi.Input

	ToCertAuthBackendRoleArrayOutput() CertAuthBackendRoleArrayOutput
	ToCertAuthBackendRoleArrayOutputWithContext(context.Context) CertAuthBackendRoleArrayOutput
}

type CertAuthBackendRoleArray []CertAuthBackendRoleInput

func (CertAuthBackendRoleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertAuthBackendRole)(nil)).Elem()
}

func (i CertAuthBackendRoleArray) ToCertAuthBackendRoleArrayOutput() CertAuthBackendRoleArrayOutput {
	return i.ToCertAuthBackendRoleArrayOutputWithContext(context.Background())
}

func (i CertAuthBackendRoleArray) ToCertAuthBackendRoleArrayOutputWithContext(ctx context.Context) CertAuthBackendRoleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertAuthBackendRoleArrayOutput)
}

// CertAuthBackendRoleMapInput is an input type that accepts CertAuthBackendRoleMap and CertAuthBackendRoleMapOutput values.
// You can construct a concrete instance of `CertAuthBackendRoleMapInput` via:
//
//          CertAuthBackendRoleMap{ "key": CertAuthBackendRoleArgs{...} }
type CertAuthBackendRoleMapInput interface {
	pulumi.Input

	ToCertAuthBackendRoleMapOutput() CertAuthBackendRoleMapOutput
	ToCertAuthBackendRoleMapOutputWithContext(context.Context) CertAuthBackendRoleMapOutput
}

type CertAuthBackendRoleMap map[string]CertAuthBackendRoleInput

func (CertAuthBackendRoleMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertAuthBackendRole)(nil)).Elem()
}

func (i CertAuthBackendRoleMap) ToCertAuthBackendRoleMapOutput() CertAuthBackendRoleMapOutput {
	return i.ToCertAuthBackendRoleMapOutputWithContext(context.Background())
}

func (i CertAuthBackendRoleMap) ToCertAuthBackendRoleMapOutputWithContext(ctx context.Context) CertAuthBackendRoleMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertAuthBackendRoleMapOutput)
}

type CertAuthBackendRoleOutput struct{ *pulumi.OutputState }

func (CertAuthBackendRoleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CertAuthBackendRole)(nil)).Elem()
}

func (o CertAuthBackendRoleOutput) ToCertAuthBackendRoleOutput() CertAuthBackendRoleOutput {
	return o
}

func (o CertAuthBackendRoleOutput) ToCertAuthBackendRoleOutputWithContext(ctx context.Context) CertAuthBackendRoleOutput {
	return o
}

type CertAuthBackendRoleArrayOutput struct{ *pulumi.OutputState }

func (CertAuthBackendRoleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*CertAuthBackendRole)(nil)).Elem()
}

func (o CertAuthBackendRoleArrayOutput) ToCertAuthBackendRoleArrayOutput() CertAuthBackendRoleArrayOutput {
	return o
}

func (o CertAuthBackendRoleArrayOutput) ToCertAuthBackendRoleArrayOutputWithContext(ctx context.Context) CertAuthBackendRoleArrayOutput {
	return o
}

func (o CertAuthBackendRoleArrayOutput) Index(i pulumi.IntInput) CertAuthBackendRoleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *CertAuthBackendRole {
		return vs[0].([]*CertAuthBackendRole)[vs[1].(int)]
	}).(CertAuthBackendRoleOutput)
}

type CertAuthBackendRoleMapOutput struct{ *pulumi.OutputState }

func (CertAuthBackendRoleMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*CertAuthBackendRole)(nil)).Elem()
}

func (o CertAuthBackendRoleMapOutput) ToCertAuthBackendRoleMapOutput() CertAuthBackendRoleMapOutput {
	return o
}

func (o CertAuthBackendRoleMapOutput) ToCertAuthBackendRoleMapOutputWithContext(ctx context.Context) CertAuthBackendRoleMapOutput {
	return o
}

func (o CertAuthBackendRoleMapOutput) MapIndex(k pulumi.StringInput) CertAuthBackendRoleOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *CertAuthBackendRole {
		return vs[0].(map[string]*CertAuthBackendRole)[vs[1].(string)]
	}).(CertAuthBackendRoleOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CertAuthBackendRoleInput)(nil)).Elem(), &CertAuthBackendRole{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertAuthBackendRoleArrayInput)(nil)).Elem(), CertAuthBackendRoleArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CertAuthBackendRoleMapInput)(nil)).Elem(), CertAuthBackendRoleMap{})
	pulumi.RegisterOutputType(CertAuthBackendRoleOutput{})
	pulumi.RegisterOutputType(CertAuthBackendRoleArrayOutput{})
	pulumi.RegisterOutputType(CertAuthBackendRoleMapOutput{})
}
