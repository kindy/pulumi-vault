// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package aws

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/d/aws_access_credentials.html.markdown.
func LookupAccessCredentials(ctx *pulumi.Context, args *GetAccessCredentialsArgs) (*GetAccessCredentialsResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["backend"] = args.Backend
		inputs["role"] = args.Role
		inputs["type"] = args.Type
	}
	outputs, err := ctx.Invoke("vault:aws/getAccessCredentials:getAccessCredentials", inputs)
	if err != nil {
		return nil, err
	}
	return &GetAccessCredentialsResult{
		AccessKey: outputs["accessKey"],
		Backend: outputs["backend"],
		LeaseDuration: outputs["leaseDuration"],
		LeaseId: outputs["leaseId"],
		LeaseRenewable: outputs["leaseRenewable"],
		LeaseStartTime: outputs["leaseStartTime"],
		Role: outputs["role"],
		SecretKey: outputs["secretKey"],
		SecurityToken: outputs["securityToken"],
		Type: outputs["type"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getAccessCredentials.
type GetAccessCredentialsArgs struct {
	// The path to the AWS secret backend to
	// read credentials from, with no leading or trailing `/`s.
	Backend interface{}
	// The name of the AWS secret backend role to read
	// credentials from, with no leading or trailing `/`s.
	Role interface{}
	// The type of credentials to read. Defaults
	// to `"creds"`, which just returns an AWS Access Key ID and Secret
	// Key. Can also be set to `"sts"`, which will return a security token
	// in addition to the keys.
	Type interface{}
}

// A collection of values returned by getAccessCredentials.
type GetAccessCredentialsResult struct {
	// The AWS Access Key ID returned by Vault.
	AccessKey interface{}
	Backend interface{}
	// The duration of the secret lease, in seconds relative
	// to the time the data was requested. Once this time has passed any plan
	// generated with this data may fail to apply.
	LeaseDuration interface{}
	// The lease identifier assigned by Vault.
	LeaseId interface{}
	LeaseRenewable interface{}
	LeaseStartTime interface{}
	Role interface{}
	// The AWS Secret Key returned by Vault.
	SecretKey interface{}
	// The STS token returned by Vault, if any.
	SecurityToken interface{}
	Type interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}
