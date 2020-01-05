// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package vault

import (
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

func LookupAuthBackend(ctx *pulumi.Context, args *GetAuthBackendArgs) (*GetAuthBackendResult, error) {
	inputs := make(map[string]interface{})
	if args != nil {
		inputs["path"] = args.Path
	}
	outputs, err := ctx.Invoke("vault:index/getAuthBackend:getAuthBackend", inputs)
	if err != nil {
		return nil, err
	}
	return &GetAuthBackendResult{
		Accessor: outputs["accessor"],
		DefaultLeaseTtlSeconds: outputs["defaultLeaseTtlSeconds"],
		Description: outputs["description"],
		ListingVisibility: outputs["listingVisibility"],
		Local: outputs["local"],
		MaxLeaseTtlSeconds: outputs["maxLeaseTtlSeconds"],
		Path: outputs["path"],
		Type: outputs["type"],
		Id: outputs["id"],
	}, nil
}

// A collection of arguments for invoking getAuthBackend.
type GetAuthBackendArgs struct {
	// The auth backend mount point.
	Path interface{}
}

// A collection of values returned by getAuthBackend.
type GetAuthBackendResult struct {
	// The accessor for this auth method
	Accessor interface{}
	// The default lease duration in seconds.
	DefaultLeaseTtlSeconds interface{}
	// A description of the auth method.
	Description interface{}
	// Speficies whether to show this mount in the UI-specific listing endpoint.
	ListingVisibility interface{}
	// Specifies if the auth method is local only.
	Local interface{}
	// The maximum lease duration in seconds.
	MaxLeaseTtlSeconds interface{}
	Path interface{}
	// The name of the auth method type.
	Type interface{}
	// id is the provider-assigned unique ID for this managed resource.
	Id interface{}
}