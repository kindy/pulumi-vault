// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package azure

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault"
// 	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/azure"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleAuthBackend, err := vault.NewAuthBackend(ctx, "exampleAuthBackend", &vault.AuthBackendArgs{
// 			Type: pulumi.String("azure"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = azure.NewAuthBackendConfig(ctx, "exampleAuthBackendConfig", &azure.AuthBackendConfigArgs{
// 			Backend:      exampleAuthBackend.Path,
// 			TenantId:     pulumi.String("11111111-2222-3333-4444-555555555555"),
// 			ClientId:     pulumi.String("11111111-2222-3333-4444-555555555555"),
// 			ClientSecret: pulumi.String("01234567890123456789"),
// 			Resource:     pulumi.String("https://vault.hashicorp.com"),
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
// Azure auth backends can be imported using `auth/`, the `backend` path, and `/config` e.g.
//
// ```sh
//  $ pulumi import vault:azure/authBackendConfig:AuthBackendConfig example auth/azure/config
// ```
type AuthBackendConfig struct {
	pulumi.CustomResourceState

	// The path the Azure auth backend being configured was
	// mounted at.  Defaults to `azure`.
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// The client id for credentials to query the Azure APIs.
	// Currently read permissions to query compute resources are required.
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// The client secret for credentials to query the
	// Azure APIs.
	ClientSecret pulumi.StringPtrOutput `pulumi:"clientSecret"`
	// The Azure cloud environment. Valid values:
	// AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
	// AzureGermanCloud.  Defaults to `AzurePublicCloud`.
	Environment pulumi.StringPtrOutput `pulumi:"environment"`
	// The configured URL for the application registered in
	// Azure Active Directory.
	Resource pulumi.StringOutput `pulumi:"resource"`
	// The tenant id for the Azure Active Directory
	// organization.
	TenantId pulumi.StringOutput `pulumi:"tenantId"`
}

// NewAuthBackendConfig registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendConfig(ctx *pulumi.Context,
	name string, args *AuthBackendConfigArgs, opts ...pulumi.ResourceOption) (*AuthBackendConfig, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.TenantId == nil {
		return nil, errors.New("invalid value for required argument 'TenantId'")
	}
	var resource AuthBackendConfig
	err := ctx.RegisterResource("vault:azure/authBackendConfig:AuthBackendConfig", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackendConfig gets an existing AuthBackendConfig resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendConfigState, opts ...pulumi.ResourceOption) (*AuthBackendConfig, error) {
	var resource AuthBackendConfig
	err := ctx.ReadResource("vault:azure/authBackendConfig:AuthBackendConfig", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendConfig resources.
type authBackendConfigState struct {
	// The path the Azure auth backend being configured was
	// mounted at.  Defaults to `azure`.
	Backend *string `pulumi:"backend"`
	// The client id for credentials to query the Azure APIs.
	// Currently read permissions to query compute resources are required.
	ClientId *string `pulumi:"clientId"`
	// The client secret for credentials to query the
	// Azure APIs.
	ClientSecret *string `pulumi:"clientSecret"`
	// The Azure cloud environment. Valid values:
	// AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
	// AzureGermanCloud.  Defaults to `AzurePublicCloud`.
	Environment *string `pulumi:"environment"`
	// The configured URL for the application registered in
	// Azure Active Directory.
	Resource *string `pulumi:"resource"`
	// The tenant id for the Azure Active Directory
	// organization.
	TenantId *string `pulumi:"tenantId"`
}

type AuthBackendConfigState struct {
	// The path the Azure auth backend being configured was
	// mounted at.  Defaults to `azure`.
	Backend pulumi.StringPtrInput
	// The client id for credentials to query the Azure APIs.
	// Currently read permissions to query compute resources are required.
	ClientId pulumi.StringPtrInput
	// The client secret for credentials to query the
	// Azure APIs.
	ClientSecret pulumi.StringPtrInput
	// The Azure cloud environment. Valid values:
	// AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
	// AzureGermanCloud.  Defaults to `AzurePublicCloud`.
	Environment pulumi.StringPtrInput
	// The configured URL for the application registered in
	// Azure Active Directory.
	Resource pulumi.StringPtrInput
	// The tenant id for the Azure Active Directory
	// organization.
	TenantId pulumi.StringPtrInput
}

func (AuthBackendConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendConfigState)(nil)).Elem()
}

type authBackendConfigArgs struct {
	// The path the Azure auth backend being configured was
	// mounted at.  Defaults to `azure`.
	Backend *string `pulumi:"backend"`
	// The client id for credentials to query the Azure APIs.
	// Currently read permissions to query compute resources are required.
	ClientId *string `pulumi:"clientId"`
	// The client secret for credentials to query the
	// Azure APIs.
	ClientSecret *string `pulumi:"clientSecret"`
	// The Azure cloud environment. Valid values:
	// AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
	// AzureGermanCloud.  Defaults to `AzurePublicCloud`.
	Environment *string `pulumi:"environment"`
	// The configured URL for the application registered in
	// Azure Active Directory.
	Resource string `pulumi:"resource"`
	// The tenant id for the Azure Active Directory
	// organization.
	TenantId string `pulumi:"tenantId"`
}

// The set of arguments for constructing a AuthBackendConfig resource.
type AuthBackendConfigArgs struct {
	// The path the Azure auth backend being configured was
	// mounted at.  Defaults to `azure`.
	Backend pulumi.StringPtrInput
	// The client id for credentials to query the Azure APIs.
	// Currently read permissions to query compute resources are required.
	ClientId pulumi.StringPtrInput
	// The client secret for credentials to query the
	// Azure APIs.
	ClientSecret pulumi.StringPtrInput
	// The Azure cloud environment. Valid values:
	// AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
	// AzureGermanCloud.  Defaults to `AzurePublicCloud`.
	Environment pulumi.StringPtrInput
	// The configured URL for the application registered in
	// Azure Active Directory.
	Resource pulumi.StringInput
	// The tenant id for the Azure Active Directory
	// organization.
	TenantId pulumi.StringInput
}

func (AuthBackendConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendConfigArgs)(nil)).Elem()
}

type AuthBackendConfigInput interface {
	pulumi.Input

	ToAuthBackendConfigOutput() AuthBackendConfigOutput
	ToAuthBackendConfigOutputWithContext(ctx context.Context) AuthBackendConfigOutput
}

func (*AuthBackendConfig) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendConfig)(nil)).Elem()
}

func (i *AuthBackendConfig) ToAuthBackendConfigOutput() AuthBackendConfigOutput {
	return i.ToAuthBackendConfigOutputWithContext(context.Background())
}

func (i *AuthBackendConfig) ToAuthBackendConfigOutputWithContext(ctx context.Context) AuthBackendConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendConfigOutput)
}

// AuthBackendConfigArrayInput is an input type that accepts AuthBackendConfigArray and AuthBackendConfigArrayOutput values.
// You can construct a concrete instance of `AuthBackendConfigArrayInput` via:
//
//          AuthBackendConfigArray{ AuthBackendConfigArgs{...} }
type AuthBackendConfigArrayInput interface {
	pulumi.Input

	ToAuthBackendConfigArrayOutput() AuthBackendConfigArrayOutput
	ToAuthBackendConfigArrayOutputWithContext(context.Context) AuthBackendConfigArrayOutput
}

type AuthBackendConfigArray []AuthBackendConfigInput

func (AuthBackendConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackendConfig)(nil)).Elem()
}

func (i AuthBackendConfigArray) ToAuthBackendConfigArrayOutput() AuthBackendConfigArrayOutput {
	return i.ToAuthBackendConfigArrayOutputWithContext(context.Background())
}

func (i AuthBackendConfigArray) ToAuthBackendConfigArrayOutputWithContext(ctx context.Context) AuthBackendConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendConfigArrayOutput)
}

// AuthBackendConfigMapInput is an input type that accepts AuthBackendConfigMap and AuthBackendConfigMapOutput values.
// You can construct a concrete instance of `AuthBackendConfigMapInput` via:
//
//          AuthBackendConfigMap{ "key": AuthBackendConfigArgs{...} }
type AuthBackendConfigMapInput interface {
	pulumi.Input

	ToAuthBackendConfigMapOutput() AuthBackendConfigMapOutput
	ToAuthBackendConfigMapOutputWithContext(context.Context) AuthBackendConfigMapOutput
}

type AuthBackendConfigMap map[string]AuthBackendConfigInput

func (AuthBackendConfigMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackendConfig)(nil)).Elem()
}

func (i AuthBackendConfigMap) ToAuthBackendConfigMapOutput() AuthBackendConfigMapOutput {
	return i.ToAuthBackendConfigMapOutputWithContext(context.Background())
}

func (i AuthBackendConfigMap) ToAuthBackendConfigMapOutputWithContext(ctx context.Context) AuthBackendConfigMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AuthBackendConfigMapOutput)
}

type AuthBackendConfigOutput struct{ *pulumi.OutputState }

func (AuthBackendConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AuthBackendConfig)(nil)).Elem()
}

func (o AuthBackendConfigOutput) ToAuthBackendConfigOutput() AuthBackendConfigOutput {
	return o
}

func (o AuthBackendConfigOutput) ToAuthBackendConfigOutputWithContext(ctx context.Context) AuthBackendConfigOutput {
	return o
}

type AuthBackendConfigArrayOutput struct{ *pulumi.OutputState }

func (AuthBackendConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AuthBackendConfig)(nil)).Elem()
}

func (o AuthBackendConfigArrayOutput) ToAuthBackendConfigArrayOutput() AuthBackendConfigArrayOutput {
	return o
}

func (o AuthBackendConfigArrayOutput) ToAuthBackendConfigArrayOutputWithContext(ctx context.Context) AuthBackendConfigArrayOutput {
	return o
}

func (o AuthBackendConfigArrayOutput) Index(i pulumi.IntInput) AuthBackendConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AuthBackendConfig {
		return vs[0].([]*AuthBackendConfig)[vs[1].(int)]
	}).(AuthBackendConfigOutput)
}

type AuthBackendConfigMapOutput struct{ *pulumi.OutputState }

func (AuthBackendConfigMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AuthBackendConfig)(nil)).Elem()
}

func (o AuthBackendConfigMapOutput) ToAuthBackendConfigMapOutput() AuthBackendConfigMapOutput {
	return o
}

func (o AuthBackendConfigMapOutput) ToAuthBackendConfigMapOutputWithContext(ctx context.Context) AuthBackendConfigMapOutput {
	return o
}

func (o AuthBackendConfigMapOutput) MapIndex(k pulumi.StringInput) AuthBackendConfigOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AuthBackendConfig {
		return vs[0].(map[string]*AuthBackendConfig)[vs[1].(string)]
	}).(AuthBackendConfigOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendConfigInput)(nil)).Elem(), &AuthBackendConfig{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendConfigArrayInput)(nil)).Elem(), AuthBackendConfigArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AuthBackendConfigMapInput)(nil)).Elem(), AuthBackendConfigMap{})
	pulumi.RegisterOutputType(AuthBackendConfigOutput{})
	pulumi.RegisterOutputType(AuthBackendConfigArrayOutput{})
	pulumi.RegisterOutputType(AuthBackendConfigMapOutput{})
}
