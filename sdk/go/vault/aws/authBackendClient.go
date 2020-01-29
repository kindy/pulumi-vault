// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

// nolint: lll
package aws

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/aws_auth_backend_client.html.markdown.
type AuthBackendClient struct {
	pulumi.CustomResourceState

	// The AWS access key that Vault should use for the
	// auth backend.
	AccessKey pulumi.StringPtrOutput `pulumi:"accessKey"`
	// The path the AWS auth backend being configured was
	// mounted at.  Defaults to `aws`.
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// Override the URL Vault uses when making EC2 API
	// calls.
	Ec2Endpoint pulumi.StringPtrOutput `pulumi:"ec2Endpoint"`
	// Override the URL Vault uses when making IAM API
	// calls.
	IamEndpoint pulumi.StringPtrOutput `pulumi:"iamEndpoint"`
	// The value to require in the
	// `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
	// that are used in the IAM auth method.
	IamServerIdHeaderValue pulumi.StringPtrOutput `pulumi:"iamServerIdHeaderValue"`
	// The AWS secret key that Vault should use for the
	// auth backend.
	SecretKey pulumi.StringPtrOutput `pulumi:"secretKey"`
	// Override the URL Vault uses when making STS API
	// calls.
	StsEndpoint pulumi.StringPtrOutput `pulumi:"stsEndpoint"`
}

// NewAuthBackendClient registers a new resource with the given unique name, arguments, and options.
func NewAuthBackendClient(ctx *pulumi.Context,
	name string, args *AuthBackendClientArgs, opts ...pulumi.ResourceOption) (*AuthBackendClient, error) {
	if args == nil {
		args = &AuthBackendClientArgs{}
	}
	var resource AuthBackendClient
	err := ctx.RegisterResource("vault:aws/authBackendClient:AuthBackendClient", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAuthBackendClient gets an existing AuthBackendClient resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAuthBackendClient(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AuthBackendClientState, opts ...pulumi.ResourceOption) (*AuthBackendClient, error) {
	var resource AuthBackendClient
	err := ctx.ReadResource("vault:aws/authBackendClient:AuthBackendClient", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AuthBackendClient resources.
type authBackendClientState struct {
	// The AWS access key that Vault should use for the
	// auth backend.
	AccessKey *string `pulumi:"accessKey"`
	// The path the AWS auth backend being configured was
	// mounted at.  Defaults to `aws`.
	Backend *string `pulumi:"backend"`
	// Override the URL Vault uses when making EC2 API
	// calls.
	Ec2Endpoint *string `pulumi:"ec2Endpoint"`
	// Override the URL Vault uses when making IAM API
	// calls.
	IamEndpoint *string `pulumi:"iamEndpoint"`
	// The value to require in the
	// `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
	// that are used in the IAM auth method.
	IamServerIdHeaderValue *string `pulumi:"iamServerIdHeaderValue"`
	// The AWS secret key that Vault should use for the
	// auth backend.
	SecretKey *string `pulumi:"secretKey"`
	// Override the URL Vault uses when making STS API
	// calls.
	StsEndpoint *string `pulumi:"stsEndpoint"`
}

type AuthBackendClientState struct {
	// The AWS access key that Vault should use for the
	// auth backend.
	AccessKey pulumi.StringPtrInput
	// The path the AWS auth backend being configured was
	// mounted at.  Defaults to `aws`.
	Backend pulumi.StringPtrInput
	// Override the URL Vault uses when making EC2 API
	// calls.
	Ec2Endpoint pulumi.StringPtrInput
	// Override the URL Vault uses when making IAM API
	// calls.
	IamEndpoint pulumi.StringPtrInput
	// The value to require in the
	// `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
	// that are used in the IAM auth method.
	IamServerIdHeaderValue pulumi.StringPtrInput
	// The AWS secret key that Vault should use for the
	// auth backend.
	SecretKey pulumi.StringPtrInput
	// Override the URL Vault uses when making STS API
	// calls.
	StsEndpoint pulumi.StringPtrInput
}

func (AuthBackendClientState) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendClientState)(nil)).Elem()
}

type authBackendClientArgs struct {
	// The AWS access key that Vault should use for the
	// auth backend.
	AccessKey *string `pulumi:"accessKey"`
	// The path the AWS auth backend being configured was
	// mounted at.  Defaults to `aws`.
	Backend *string `pulumi:"backend"`
	// Override the URL Vault uses when making EC2 API
	// calls.
	Ec2Endpoint *string `pulumi:"ec2Endpoint"`
	// Override the URL Vault uses when making IAM API
	// calls.
	IamEndpoint *string `pulumi:"iamEndpoint"`
	// The value to require in the
	// `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
	// that are used in the IAM auth method.
	IamServerIdHeaderValue *string `pulumi:"iamServerIdHeaderValue"`
	// The AWS secret key that Vault should use for the
	// auth backend.
	SecretKey *string `pulumi:"secretKey"`
	// Override the URL Vault uses when making STS API
	// calls.
	StsEndpoint *string `pulumi:"stsEndpoint"`
}

// The set of arguments for constructing a AuthBackendClient resource.
type AuthBackendClientArgs struct {
	// The AWS access key that Vault should use for the
	// auth backend.
	AccessKey pulumi.StringPtrInput
	// The path the AWS auth backend being configured was
	// mounted at.  Defaults to `aws`.
	Backend pulumi.StringPtrInput
	// Override the URL Vault uses when making EC2 API
	// calls.
	Ec2Endpoint pulumi.StringPtrInput
	// Override the URL Vault uses when making IAM API
	// calls.
	IamEndpoint pulumi.StringPtrInput
	// The value to require in the
	// `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
	// that are used in the IAM auth method.
	IamServerIdHeaderValue pulumi.StringPtrInput
	// The AWS secret key that Vault should use for the
	// auth backend.
	SecretKey pulumi.StringPtrInput
	// Override the URL Vault uses when making STS API
	// calls.
	StsEndpoint pulumi.StringPtrInput
}

func (AuthBackendClientArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*authBackendClientArgs)(nil)).Elem()
}

