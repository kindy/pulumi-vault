// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Provides a resource to manage [Duo MFA](https://www.vaultproject.io/docs/enterprise/mfa/mfa-duo.html).
//
// **Note** this feature is available only with Vault Enterprise.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v2/go/vault"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		userpass, err := vault.NewAuthBackend(ctx, "userpass", &vault.AuthBackendArgs{
// 			Path: pulumi.String("userpass"),
// 			Type: pulumi.String("userpass"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = vault.NewMfaDuo(ctx, "myDuo", &vault.MfaDuoArgs{
// 			ApiHostname:    pulumi.String("api-2b5c39f5.duosecurity.com"),
// 			IntegrationKey: pulumi.String("BIACEUEAXI20BNWTEYXT"),
// 			MountAccessor:  userpass.Accessor,
// 			SecretKey:      pulumi.String("8C7THtrIigh2rPZQMbguugt8IUftWhMRCOBzbuyz"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type MfaDuo struct {
	pulumi.CustomResourceState

	// `(string: <required>)` - API hostname for Duo.
	ApiHostname pulumi.StringOutput `pulumi:"apiHostname"`
	// `(string: <required>)` - Integration key for Duo.
	IntegrationKey pulumi.StringOutput `pulumi:"integrationKey"`
	// `(string: <required>)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	MountAccessor pulumi.StringOutput `pulumi:"mountAccessor"`
	// `(string: <required>)` – Name of the MFA method.
	Name pulumi.StringOutput `pulumi:"name"`
	// `(string)` - Push information for Duo.
	PushInfo pulumi.StringPtrOutput `pulumi:"pushInfo"`
	// `(string: <required>)` - Secret key for Duo.
	SecretKey pulumi.StringOutput `pulumi:"secretKey"`
	// `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
	// - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
	// - entity.name: The name configured for the Entity
	// - alias.metadata.`<key>`: The value of the Alias's metadata parameter
	// - entity.metadata.`<key>`: The value of the Entity's metadata parameter
	UsernameFormat pulumi.StringPtrOutput `pulumi:"usernameFormat"`
}

// NewMfaDuo registers a new resource with the given unique name, arguments, and options.
func NewMfaDuo(ctx *pulumi.Context,
	name string, args *MfaDuoArgs, opts ...pulumi.ResourceOption) (*MfaDuo, error) {
	if args == nil || args.ApiHostname == nil {
		return nil, errors.New("missing required argument 'ApiHostname'")
	}
	if args == nil || args.IntegrationKey == nil {
		return nil, errors.New("missing required argument 'IntegrationKey'")
	}
	if args == nil || args.MountAccessor == nil {
		return nil, errors.New("missing required argument 'MountAccessor'")
	}
	if args == nil || args.SecretKey == nil {
		return nil, errors.New("missing required argument 'SecretKey'")
	}
	if args == nil {
		args = &MfaDuoArgs{}
	}
	var resource MfaDuo
	err := ctx.RegisterResource("vault:index/mfaDuo:MfaDuo", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMfaDuo gets an existing MfaDuo resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMfaDuo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MfaDuoState, opts ...pulumi.ResourceOption) (*MfaDuo, error) {
	var resource MfaDuo
	err := ctx.ReadResource("vault:index/mfaDuo:MfaDuo", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MfaDuo resources.
type mfaDuoState struct {
	// `(string: <required>)` - API hostname for Duo.
	ApiHostname *string `pulumi:"apiHostname"`
	// `(string: <required>)` - Integration key for Duo.
	IntegrationKey *string `pulumi:"integrationKey"`
	// `(string: <required>)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	MountAccessor *string `pulumi:"mountAccessor"`
	// `(string: <required>)` – Name of the MFA method.
	Name *string `pulumi:"name"`
	// `(string)` - Push information for Duo.
	PushInfo *string `pulumi:"pushInfo"`
	// `(string: <required>)` - Secret key for Duo.
	SecretKey *string `pulumi:"secretKey"`
	// `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
	// - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
	// - entity.name: The name configured for the Entity
	// - alias.metadata.`<key>`: The value of the Alias's metadata parameter
	// - entity.metadata.`<key>`: The value of the Entity's metadata parameter
	UsernameFormat *string `pulumi:"usernameFormat"`
}

type MfaDuoState struct {
	// `(string: <required>)` - API hostname for Duo.
	ApiHostname pulumi.StringPtrInput
	// `(string: <required>)` - Integration key for Duo.
	IntegrationKey pulumi.StringPtrInput
	// `(string: <required>)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	MountAccessor pulumi.StringPtrInput
	// `(string: <required>)` – Name of the MFA method.
	Name pulumi.StringPtrInput
	// `(string)` - Push information for Duo.
	PushInfo pulumi.StringPtrInput
	// `(string: <required>)` - Secret key for Duo.
	SecretKey pulumi.StringPtrInput
	// `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
	// - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
	// - entity.name: The name configured for the Entity
	// - alias.metadata.`<key>`: The value of the Alias's metadata parameter
	// - entity.metadata.`<key>`: The value of the Entity's metadata parameter
	UsernameFormat pulumi.StringPtrInput
}

func (MfaDuoState) ElementType() reflect.Type {
	return reflect.TypeOf((*mfaDuoState)(nil)).Elem()
}

type mfaDuoArgs struct {
	// `(string: <required>)` - API hostname for Duo.
	ApiHostname string `pulumi:"apiHostname"`
	// `(string: <required>)` - Integration key for Duo.
	IntegrationKey string `pulumi:"integrationKey"`
	// `(string: <required>)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	MountAccessor string `pulumi:"mountAccessor"`
	// `(string: <required>)` – Name of the MFA method.
	Name *string `pulumi:"name"`
	// `(string)` - Push information for Duo.
	PushInfo *string `pulumi:"pushInfo"`
	// `(string: <required>)` - Secret key for Duo.
	SecretKey string `pulumi:"secretKey"`
	// `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
	// - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
	// - entity.name: The name configured for the Entity
	// - alias.metadata.`<key>`: The value of the Alias's metadata parameter
	// - entity.metadata.`<key>`: The value of the Entity's metadata parameter
	UsernameFormat *string `pulumi:"usernameFormat"`
}

// The set of arguments for constructing a MfaDuo resource.
type MfaDuoArgs struct {
	// `(string: <required>)` - API hostname for Duo.
	ApiHostname pulumi.StringInput
	// `(string: <required>)` - Integration key for Duo.
	IntegrationKey pulumi.StringInput
	// `(string: <required>)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
	MountAccessor pulumi.StringInput
	// `(string: <required>)` – Name of the MFA method.
	Name pulumi.StringPtrInput
	// `(string)` - Push information for Duo.
	PushInfo pulumi.StringPtrInput
	// `(string: <required>)` - Secret key for Duo.
	SecretKey pulumi.StringInput
	// `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
	// - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
	// - entity.name: The name configured for the Entity
	// - alias.metadata.`<key>`: The value of the Alias's metadata parameter
	// - entity.metadata.`<key>`: The value of the Entity's metadata parameter
	UsernameFormat pulumi.StringPtrInput
}

func (MfaDuoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mfaDuoArgs)(nil)).Elem()
}
