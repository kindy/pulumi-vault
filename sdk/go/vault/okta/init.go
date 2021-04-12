// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package okta

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-vault/sdk/v3/go/vault"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "vault:okta/authBackend:AuthBackend":
		r = &AuthBackend{}
	case "vault:okta/authBackendGroup:AuthBackendGroup":
		r = &AuthBackendGroup{}
	case "vault:okta/authBackendUser:AuthBackendUser":
		r = &AuthBackendUser{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := vault.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"vault",
		"okta/authBackend",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"okta/authBackendGroup",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"vault",
		"okta/authBackendUser",
		&module{version},
	)
}
