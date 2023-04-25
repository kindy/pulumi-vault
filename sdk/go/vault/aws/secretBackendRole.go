// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package aws

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
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/aws"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			aws, err := aws.NewSecretBackend(ctx, "aws", &aws.SecretBackendArgs{
//				AccessKey: pulumi.String("AKIA....."),
//				SecretKey: pulumi.String("AWS secret key"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = aws.NewSecretBackendRole(ctx, "role", &aws.SecretBackendRoleArgs{
//				Backend:        aws.Path,
//				CredentialType: pulumi.String("iam_user"),
//				PolicyDocument: pulumi.String("{\n  \"Version\": \"2012-10-17\",\n  \"Statement\": [\n    {\n      \"Effect\": \"Allow\",\n      \"Action\": \"iam:*\",\n      \"Resource\": \"*\"\n    }\n  ]\n}\n"),
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
// AWS secret backend roles can be imported using the `path`, e.g.
//
// ```sh
//
//	$ pulumi import vault:aws/secretBackendRole:SecretBackendRole role aws/roles/deploy
//
// ```
type SecretBackendRole struct {
	pulumi.CustomResourceState

	// The path the AWS secret backend is mounted at,
	// with no leading or trailing `/`s.
	Backend pulumi.StringOutput `pulumi:"backend"`
	// Specifies the type of credential to be used when
	// retrieving credentials from the role. Must be one of `iamUser`, `assumedRole`, or
	// `federationToken`.
	CredentialType pulumi.StringOutput `pulumi:"credentialType"`
	// The default TTL in seconds for STS credentials.
	// When a TTL is not specified when STS credentials are requested,
	// and a default TTL is specified on the role,
	// then this default TTL will be used. Valid only when `credentialType` is one of
	// `assumedRole` or `federationToken`.
	DefaultStsTtl pulumi.IntOutput `pulumi:"defaultStsTtl"`
	// A list of IAM group names. IAM users generated
	// against this vault role will be added to these IAM Groups. For a credential
	// type of `assumedRole` or `federationToken`, the policies sent to the
	// corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the
	// policies from each group in `iamGroups` combined with the `policyDocument`
	// and `policyArns` parameters.
	IamGroups pulumi.StringArrayOutput `pulumi:"iamGroups"`
	// The max allowed TTL in seconds for STS credentials
	// (credentials TTL are capped to `maxStsTtl`). Valid only when `credentialType` is
	// one of `assumedRole` or `federationToken`.
	MaxStsTtl pulumi.IntOutput `pulumi:"maxStsTtl"`
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Name pulumi.StringOutput `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// The ARN of the AWS Permissions
	// Boundary to attach to IAM users created in the role. Valid only when
	// `credentialType` is `iamUser`. If not specified, then no permissions boundary
	// policy will be attached.
	PermissionsBoundaryArn pulumi.StringPtrOutput `pulumi:"permissionsBoundaryArn"`
	// Specifies a list of AWS managed policy ARNs. The
	// behavior depends on the credential type. With `iamUser`, the policies will be
	// attached to IAM users when they are requested. With `assumedRole` and
	// `federationToken`, the policy ARNs will act as a filter on what the credentials
	// can do, similar to `policyDocument`. When `credentialType` is `iamUser` or
	// `federationToken`, at least one of `policyDocument` or `policyArns` must
	// be specified.
	PolicyArns pulumi.StringArrayOutput `pulumi:"policyArns"`
	// The IAM policy document for the role. The
	// behavior depends on the credential type. With `iamUser`, the policy document
	// will be attached to the IAM user generated and augment the permissions the IAM
	// user has. With `assumedRole` and `federationToken`, the policy document will
	// act as a filter on what the credentials can do, similar to `policyArns`.
	PolicyDocument pulumi.StringPtrOutput `pulumi:"policyDocument"`
	// Specifies the ARNs of the AWS roles this Vault role
	// is allowed to assume. Required when `credentialType` is `assumedRole` and
	// prohibited otherwise.
	RoleArns pulumi.StringArrayOutput `pulumi:"roleArns"`
	// The path for the user name. Valid only when
	// `credentialType` is `iamUser`. Default is `/`.
	UserPath pulumi.StringPtrOutput `pulumi:"userPath"`
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
	if args.CredentialType == nil {
		return nil, errors.New("invalid value for required argument 'CredentialType'")
	}
	var resource SecretBackendRole
	err := ctx.RegisterResource("vault:aws/secretBackendRole:SecretBackendRole", name, args, &resource, opts...)
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
	err := ctx.ReadResource("vault:aws/secretBackendRole:SecretBackendRole", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecretBackendRole resources.
type secretBackendRoleState struct {
	// The path the AWS secret backend is mounted at,
	// with no leading or trailing `/`s.
	Backend *string `pulumi:"backend"`
	// Specifies the type of credential to be used when
	// retrieving credentials from the role. Must be one of `iamUser`, `assumedRole`, or
	// `federationToken`.
	CredentialType *string `pulumi:"credentialType"`
	// The default TTL in seconds for STS credentials.
	// When a TTL is not specified when STS credentials are requested,
	// and a default TTL is specified on the role,
	// then this default TTL will be used. Valid only when `credentialType` is one of
	// `assumedRole` or `federationToken`.
	DefaultStsTtl *int `pulumi:"defaultStsTtl"`
	// A list of IAM group names. IAM users generated
	// against this vault role will be added to these IAM Groups. For a credential
	// type of `assumedRole` or `federationToken`, the policies sent to the
	// corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the
	// policies from each group in `iamGroups` combined with the `policyDocument`
	// and `policyArns` parameters.
	IamGroups []string `pulumi:"iamGroups"`
	// The max allowed TTL in seconds for STS credentials
	// (credentials TTL are capped to `maxStsTtl`). Valid only when `credentialType` is
	// one of `assumedRole` or `federationToken`.
	MaxStsTtl *int `pulumi:"maxStsTtl"`
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The ARN of the AWS Permissions
	// Boundary to attach to IAM users created in the role. Valid only when
	// `credentialType` is `iamUser`. If not specified, then no permissions boundary
	// policy will be attached.
	PermissionsBoundaryArn *string `pulumi:"permissionsBoundaryArn"`
	// Specifies a list of AWS managed policy ARNs. The
	// behavior depends on the credential type. With `iamUser`, the policies will be
	// attached to IAM users when they are requested. With `assumedRole` and
	// `federationToken`, the policy ARNs will act as a filter on what the credentials
	// can do, similar to `policyDocument`. When `credentialType` is `iamUser` or
	// `federationToken`, at least one of `policyDocument` or `policyArns` must
	// be specified.
	PolicyArns []string `pulumi:"policyArns"`
	// The IAM policy document for the role. The
	// behavior depends on the credential type. With `iamUser`, the policy document
	// will be attached to the IAM user generated and augment the permissions the IAM
	// user has. With `assumedRole` and `federationToken`, the policy document will
	// act as a filter on what the credentials can do, similar to `policyArns`.
	PolicyDocument *string `pulumi:"policyDocument"`
	// Specifies the ARNs of the AWS roles this Vault role
	// is allowed to assume. Required when `credentialType` is `assumedRole` and
	// prohibited otherwise.
	RoleArns []string `pulumi:"roleArns"`
	// The path for the user name. Valid only when
	// `credentialType` is `iamUser`. Default is `/`.
	UserPath *string `pulumi:"userPath"`
}

type SecretBackendRoleState struct {
	// The path the AWS secret backend is mounted at,
	// with no leading or trailing `/`s.
	Backend pulumi.StringPtrInput
	// Specifies the type of credential to be used when
	// retrieving credentials from the role. Must be one of `iamUser`, `assumedRole`, or
	// `federationToken`.
	CredentialType pulumi.StringPtrInput
	// The default TTL in seconds for STS credentials.
	// When a TTL is not specified when STS credentials are requested,
	// and a default TTL is specified on the role,
	// then this default TTL will be used. Valid only when `credentialType` is one of
	// `assumedRole` or `federationToken`.
	DefaultStsTtl pulumi.IntPtrInput
	// A list of IAM group names. IAM users generated
	// against this vault role will be added to these IAM Groups. For a credential
	// type of `assumedRole` or `federationToken`, the policies sent to the
	// corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the
	// policies from each group in `iamGroups` combined with the `policyDocument`
	// and `policyArns` parameters.
	IamGroups pulumi.StringArrayInput
	// The max allowed TTL in seconds for STS credentials
	// (credentials TTL are capped to `maxStsTtl`). Valid only when `credentialType` is
	// one of `assumedRole` or `federationToken`.
	MaxStsTtl pulumi.IntPtrInput
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The ARN of the AWS Permissions
	// Boundary to attach to IAM users created in the role. Valid only when
	// `credentialType` is `iamUser`. If not specified, then no permissions boundary
	// policy will be attached.
	PermissionsBoundaryArn pulumi.StringPtrInput
	// Specifies a list of AWS managed policy ARNs. The
	// behavior depends on the credential type. With `iamUser`, the policies will be
	// attached to IAM users when they are requested. With `assumedRole` and
	// `federationToken`, the policy ARNs will act as a filter on what the credentials
	// can do, similar to `policyDocument`. When `credentialType` is `iamUser` or
	// `federationToken`, at least one of `policyDocument` or `policyArns` must
	// be specified.
	PolicyArns pulumi.StringArrayInput
	// The IAM policy document for the role. The
	// behavior depends on the credential type. With `iamUser`, the policy document
	// will be attached to the IAM user generated and augment the permissions the IAM
	// user has. With `assumedRole` and `federationToken`, the policy document will
	// act as a filter on what the credentials can do, similar to `policyArns`.
	PolicyDocument pulumi.StringPtrInput
	// Specifies the ARNs of the AWS roles this Vault role
	// is allowed to assume. Required when `credentialType` is `assumedRole` and
	// prohibited otherwise.
	RoleArns pulumi.StringArrayInput
	// The path for the user name. Valid only when
	// `credentialType` is `iamUser`. Default is `/`.
	UserPath pulumi.StringPtrInput
}

func (SecretBackendRoleState) ElementType() reflect.Type {
	return reflect.TypeOf((*secretBackendRoleState)(nil)).Elem()
}

type secretBackendRoleArgs struct {
	// The path the AWS secret backend is mounted at,
	// with no leading or trailing `/`s.
	Backend string `pulumi:"backend"`
	// Specifies the type of credential to be used when
	// retrieving credentials from the role. Must be one of `iamUser`, `assumedRole`, or
	// `federationToken`.
	CredentialType string `pulumi:"credentialType"`
	// The default TTL in seconds for STS credentials.
	// When a TTL is not specified when STS credentials are requested,
	// and a default TTL is specified on the role,
	// then this default TTL will be used. Valid only when `credentialType` is one of
	// `assumedRole` or `federationToken`.
	DefaultStsTtl *int `pulumi:"defaultStsTtl"`
	// A list of IAM group names. IAM users generated
	// against this vault role will be added to these IAM Groups. For a credential
	// type of `assumedRole` or `federationToken`, the policies sent to the
	// corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the
	// policies from each group in `iamGroups` combined with the `policyDocument`
	// and `policyArns` parameters.
	IamGroups []string `pulumi:"iamGroups"`
	// The max allowed TTL in seconds for STS credentials
	// (credentials TTL are capped to `maxStsTtl`). Valid only when `credentialType` is
	// one of `assumedRole` or `federationToken`.
	MaxStsTtl *int `pulumi:"maxStsTtl"`
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Name *string `pulumi:"name"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// The ARN of the AWS Permissions
	// Boundary to attach to IAM users created in the role. Valid only when
	// `credentialType` is `iamUser`. If not specified, then no permissions boundary
	// policy will be attached.
	PermissionsBoundaryArn *string `pulumi:"permissionsBoundaryArn"`
	// Specifies a list of AWS managed policy ARNs. The
	// behavior depends on the credential type. With `iamUser`, the policies will be
	// attached to IAM users when they are requested. With `assumedRole` and
	// `federationToken`, the policy ARNs will act as a filter on what the credentials
	// can do, similar to `policyDocument`. When `credentialType` is `iamUser` or
	// `federationToken`, at least one of `policyDocument` or `policyArns` must
	// be specified.
	PolicyArns []string `pulumi:"policyArns"`
	// The IAM policy document for the role. The
	// behavior depends on the credential type. With `iamUser`, the policy document
	// will be attached to the IAM user generated and augment the permissions the IAM
	// user has. With `assumedRole` and `federationToken`, the policy document will
	// act as a filter on what the credentials can do, similar to `policyArns`.
	PolicyDocument *string `pulumi:"policyDocument"`
	// Specifies the ARNs of the AWS roles this Vault role
	// is allowed to assume. Required when `credentialType` is `assumedRole` and
	// prohibited otherwise.
	RoleArns []string `pulumi:"roleArns"`
	// The path for the user name. Valid only when
	// `credentialType` is `iamUser`. Default is `/`.
	UserPath *string `pulumi:"userPath"`
}

// The set of arguments for constructing a SecretBackendRole resource.
type SecretBackendRoleArgs struct {
	// The path the AWS secret backend is mounted at,
	// with no leading or trailing `/`s.
	Backend pulumi.StringInput
	// Specifies the type of credential to be used when
	// retrieving credentials from the role. Must be one of `iamUser`, `assumedRole`, or
	// `federationToken`.
	CredentialType pulumi.StringInput
	// The default TTL in seconds for STS credentials.
	// When a TTL is not specified when STS credentials are requested,
	// and a default TTL is specified on the role,
	// then this default TTL will be used. Valid only when `credentialType` is one of
	// `assumedRole` or `federationToken`.
	DefaultStsTtl pulumi.IntPtrInput
	// A list of IAM group names. IAM users generated
	// against this vault role will be added to these IAM Groups. For a credential
	// type of `assumedRole` or `federationToken`, the policies sent to the
	// corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the
	// policies from each group in `iamGroups` combined with the `policyDocument`
	// and `policyArns` parameters.
	IamGroups pulumi.StringArrayInput
	// The max allowed TTL in seconds for STS credentials
	// (credentials TTL are capped to `maxStsTtl`). Valid only when `credentialType` is
	// one of `assumedRole` or `federationToken`.
	MaxStsTtl pulumi.IntPtrInput
	// The name to identify this role within the backend.
	// Must be unique within the backend.
	Name pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// The ARN of the AWS Permissions
	// Boundary to attach to IAM users created in the role. Valid only when
	// `credentialType` is `iamUser`. If not specified, then no permissions boundary
	// policy will be attached.
	PermissionsBoundaryArn pulumi.StringPtrInput
	// Specifies a list of AWS managed policy ARNs. The
	// behavior depends on the credential type. With `iamUser`, the policies will be
	// attached to IAM users when they are requested. With `assumedRole` and
	// `federationToken`, the policy ARNs will act as a filter on what the credentials
	// can do, similar to `policyDocument`. When `credentialType` is `iamUser` or
	// `federationToken`, at least one of `policyDocument` or `policyArns` must
	// be specified.
	PolicyArns pulumi.StringArrayInput
	// The IAM policy document for the role. The
	// behavior depends on the credential type. With `iamUser`, the policy document
	// will be attached to the IAM user generated and augment the permissions the IAM
	// user has. With `assumedRole` and `federationToken`, the policy document will
	// act as a filter on what the credentials can do, similar to `policyArns`.
	PolicyDocument pulumi.StringPtrInput
	// Specifies the ARNs of the AWS roles this Vault role
	// is allowed to assume. Required when `credentialType` is `assumedRole` and
	// prohibited otherwise.
	RoleArns pulumi.StringArrayInput
	// The path for the user name. Valid only when
	// `credentialType` is `iamUser`. Default is `/`.
	UserPath pulumi.StringPtrInput
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

// The path the AWS secret backend is mounted at,
// with no leading or trailing `/`s.
func (o SecretBackendRoleOutput) Backend() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringOutput { return v.Backend }).(pulumi.StringOutput)
}

// Specifies the type of credential to be used when
// retrieving credentials from the role. Must be one of `iamUser`, `assumedRole`, or
// `federationToken`.
func (o SecretBackendRoleOutput) CredentialType() pulumi.StringOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringOutput { return v.CredentialType }).(pulumi.StringOutput)
}

// The default TTL in seconds for STS credentials.
// When a TTL is not specified when STS credentials are requested,
// and a default TTL is specified on the role,
// then this default TTL will be used. Valid only when `credentialType` is one of
// `assumedRole` or `federationToken`.
func (o SecretBackendRoleOutput) DefaultStsTtl() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.IntOutput { return v.DefaultStsTtl }).(pulumi.IntOutput)
}

// A list of IAM group names. IAM users generated
// against this vault role will be added to these IAM Groups. For a credential
// type of `assumedRole` or `federationToken`, the policies sent to the
// corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the
// policies from each group in `iamGroups` combined with the `policyDocument`
// and `policyArns` parameters.
func (o SecretBackendRoleOutput) IamGroups() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringArrayOutput { return v.IamGroups }).(pulumi.StringArrayOutput)
}

// The max allowed TTL in seconds for STS credentials
// (credentials TTL are capped to `maxStsTtl`). Valid only when `credentialType` is
// one of `assumedRole` or `federationToken`.
func (o SecretBackendRoleOutput) MaxStsTtl() pulumi.IntOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.IntOutput { return v.MaxStsTtl }).(pulumi.IntOutput)
}

// The name to identify this role within the backend.
// Must be unique within the backend.
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

// The ARN of the AWS Permissions
// Boundary to attach to IAM users created in the role. Valid only when
// `credentialType` is `iamUser`. If not specified, then no permissions boundary
// policy will be attached.
func (o SecretBackendRoleOutput) PermissionsBoundaryArn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringPtrOutput { return v.PermissionsBoundaryArn }).(pulumi.StringPtrOutput)
}

// Specifies a list of AWS managed policy ARNs. The
// behavior depends on the credential type. With `iamUser`, the policies will be
// attached to IAM users when they are requested. With `assumedRole` and
// `federationToken`, the policy ARNs will act as a filter on what the credentials
// can do, similar to `policyDocument`. When `credentialType` is `iamUser` or
// `federationToken`, at least one of `policyDocument` or `policyArns` must
// be specified.
func (o SecretBackendRoleOutput) PolicyArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringArrayOutput { return v.PolicyArns }).(pulumi.StringArrayOutput)
}

// The IAM policy document for the role. The
// behavior depends on the credential type. With `iamUser`, the policy document
// will be attached to the IAM user generated and augment the permissions the IAM
// user has. With `assumedRole` and `federationToken`, the policy document will
// act as a filter on what the credentials can do, similar to `policyArns`.
func (o SecretBackendRoleOutput) PolicyDocument() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringPtrOutput { return v.PolicyDocument }).(pulumi.StringPtrOutput)
}

// Specifies the ARNs of the AWS roles this Vault role
// is allowed to assume. Required when `credentialType` is `assumedRole` and
// prohibited otherwise.
func (o SecretBackendRoleOutput) RoleArns() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringArrayOutput { return v.RoleArns }).(pulumi.StringArrayOutput)
}

// The path for the user name. Valid only when
// `credentialType` is `iamUser`. Default is `/`.
func (o SecretBackendRoleOutput) UserPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecretBackendRole) pulumi.StringPtrOutput { return v.UserPath }).(pulumi.StringPtrOutput)
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
