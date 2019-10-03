// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package pki

import (
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/go/pulumi"
)

// Creates a role on an PKI Secret Backend for Vault.
//
// > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/pki_secret_backend_role.html.markdown.
type SecretBackendRole struct {
	s *pulumi.ResourceState
}

// NewSecretBackendRole registers a new resource with the given unique name, arguments, and options.
func NewSecretBackendRole(ctx *pulumi.Context,
	name string, args *SecretBackendRoleArgs, opts ...pulumi.ResourceOpt) (*SecretBackendRole, error) {
	if args == nil || args.Backend == nil {
		return nil, errors.New("missing required argument 'Backend'")
	}
	inputs := make(map[string]interface{})
	if args == nil {
		inputs["allowAnyName"] = nil
		inputs["allowBareDomains"] = nil
		inputs["allowGlobDomains"] = nil
		inputs["allowIpSans"] = nil
		inputs["allowLocalhost"] = nil
		inputs["allowSubdomains"] = nil
		inputs["allowedDomains"] = nil
		inputs["allowedOtherSans"] = nil
		inputs["allowedUriSans"] = nil
		inputs["backend"] = nil
		inputs["basicConstraintsValidForNonCa"] = nil
		inputs["clientFlag"] = nil
		inputs["codeSigningFlag"] = nil
		inputs["countries"] = nil
		inputs["emailProtectionFlag"] = nil
		inputs["enforceHostnames"] = nil
		inputs["extKeyUsages"] = nil
		inputs["generateLease"] = nil
		inputs["keyBits"] = nil
		inputs["keyType"] = nil
		inputs["keyUsages"] = nil
		inputs["localities"] = nil
		inputs["maxTtl"] = nil
		inputs["name"] = nil
		inputs["noStore"] = nil
		inputs["organizations"] = nil
		inputs["organizationUnit"] = nil
		inputs["policyIdentifiers"] = nil
		inputs["postalCodes"] = nil
		inputs["provinces"] = nil
		inputs["requireCn"] = nil
		inputs["serverFlag"] = nil
		inputs["streetAddresses"] = nil
		inputs["ttl"] = nil
		inputs["useCsrCommonName"] = nil
		inputs["useCsrSans"] = nil
	} else {
		inputs["allowAnyName"] = args.AllowAnyName
		inputs["allowBareDomains"] = args.AllowBareDomains
		inputs["allowGlobDomains"] = args.AllowGlobDomains
		inputs["allowIpSans"] = args.AllowIpSans
		inputs["allowLocalhost"] = args.AllowLocalhost
		inputs["allowSubdomains"] = args.AllowSubdomains
		inputs["allowedDomains"] = args.AllowedDomains
		inputs["allowedOtherSans"] = args.AllowedOtherSans
		inputs["allowedUriSans"] = args.AllowedUriSans
		inputs["backend"] = args.Backend
		inputs["basicConstraintsValidForNonCa"] = args.BasicConstraintsValidForNonCa
		inputs["clientFlag"] = args.ClientFlag
		inputs["codeSigningFlag"] = args.CodeSigningFlag
		inputs["countries"] = args.Countries
		inputs["emailProtectionFlag"] = args.EmailProtectionFlag
		inputs["enforceHostnames"] = args.EnforceHostnames
		inputs["extKeyUsages"] = args.ExtKeyUsages
		inputs["generateLease"] = args.GenerateLease
		inputs["keyBits"] = args.KeyBits
		inputs["keyType"] = args.KeyType
		inputs["keyUsages"] = args.KeyUsages
		inputs["localities"] = args.Localities
		inputs["maxTtl"] = args.MaxTtl
		inputs["name"] = args.Name
		inputs["noStore"] = args.NoStore
		inputs["organizations"] = args.Organizations
		inputs["organizationUnit"] = args.OrganizationUnit
		inputs["policyIdentifiers"] = args.PolicyIdentifiers
		inputs["postalCodes"] = args.PostalCodes
		inputs["provinces"] = args.Provinces
		inputs["requireCn"] = args.RequireCn
		inputs["serverFlag"] = args.ServerFlag
		inputs["streetAddresses"] = args.StreetAddresses
		inputs["ttl"] = args.Ttl
		inputs["useCsrCommonName"] = args.UseCsrCommonName
		inputs["useCsrSans"] = args.UseCsrSans
	}
	s, err := ctx.RegisterResource("vault:pki/secretBackendRole:SecretBackendRole", name, true, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SecretBackendRole{s: s}, nil
}

// GetSecretBackendRole gets an existing SecretBackendRole resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecretBackendRole(ctx *pulumi.Context,
	name string, id pulumi.ID, state *SecretBackendRoleState, opts ...pulumi.ResourceOpt) (*SecretBackendRole, error) {
	inputs := make(map[string]interface{})
	if state != nil {
		inputs["allowAnyName"] = state.AllowAnyName
		inputs["allowBareDomains"] = state.AllowBareDomains
		inputs["allowGlobDomains"] = state.AllowGlobDomains
		inputs["allowIpSans"] = state.AllowIpSans
		inputs["allowLocalhost"] = state.AllowLocalhost
		inputs["allowSubdomains"] = state.AllowSubdomains
		inputs["allowedDomains"] = state.AllowedDomains
		inputs["allowedOtherSans"] = state.AllowedOtherSans
		inputs["allowedUriSans"] = state.AllowedUriSans
		inputs["backend"] = state.Backend
		inputs["basicConstraintsValidForNonCa"] = state.BasicConstraintsValidForNonCa
		inputs["clientFlag"] = state.ClientFlag
		inputs["codeSigningFlag"] = state.CodeSigningFlag
		inputs["countries"] = state.Countries
		inputs["emailProtectionFlag"] = state.EmailProtectionFlag
		inputs["enforceHostnames"] = state.EnforceHostnames
		inputs["extKeyUsages"] = state.ExtKeyUsages
		inputs["generateLease"] = state.GenerateLease
		inputs["keyBits"] = state.KeyBits
		inputs["keyType"] = state.KeyType
		inputs["keyUsages"] = state.KeyUsages
		inputs["localities"] = state.Localities
		inputs["maxTtl"] = state.MaxTtl
		inputs["name"] = state.Name
		inputs["noStore"] = state.NoStore
		inputs["organizations"] = state.Organizations
		inputs["organizationUnit"] = state.OrganizationUnit
		inputs["policyIdentifiers"] = state.PolicyIdentifiers
		inputs["postalCodes"] = state.PostalCodes
		inputs["provinces"] = state.Provinces
		inputs["requireCn"] = state.RequireCn
		inputs["serverFlag"] = state.ServerFlag
		inputs["streetAddresses"] = state.StreetAddresses
		inputs["ttl"] = state.Ttl
		inputs["useCsrCommonName"] = state.UseCsrCommonName
		inputs["useCsrSans"] = state.UseCsrSans
	}
	s, err := ctx.ReadResource("vault:pki/secretBackendRole:SecretBackendRole", name, id, inputs, opts...)
	if err != nil {
		return nil, err
	}
	return &SecretBackendRole{s: s}, nil
}

// URN is this resource's unique name assigned by Pulumi.
func (r *SecretBackendRole) URN() *pulumi.URNOutput {
	return r.s.URN()
}

// ID is this resource's unique identifier assigned by its provider.
func (r *SecretBackendRole) ID() *pulumi.IDOutput {
	return r.s.ID()
}

// Flag to allow any name
func (r *SecretBackendRole) AllowAnyName() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["allowAnyName"])
}

// Flag to allow certificates matching the actual domain
func (r *SecretBackendRole) AllowBareDomains() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["allowBareDomains"])
}

// Flag to allow names containing glob patterns.
func (r *SecretBackendRole) AllowGlobDomains() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["allowGlobDomains"])
}

// Flag to allow IP SANs
func (r *SecretBackendRole) AllowIpSans() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["allowIpSans"])
}

// Flag to allow certificates for localhost
func (r *SecretBackendRole) AllowLocalhost() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["allowLocalhost"])
}

// Flag to allow certificates matching subdomains
func (r *SecretBackendRole) AllowSubdomains() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["allowSubdomains"])
}

// List of allowed domains for certificates
func (r *SecretBackendRole) AllowedDomains() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["allowedDomains"])
}

// Defines allowed custom SANs
func (r *SecretBackendRole) AllowedOtherSans() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["allowedOtherSans"])
}

// Defines allowed URI SANs
func (r *SecretBackendRole) AllowedUriSans() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["allowedUriSans"])
}

// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
func (r *SecretBackendRole) Backend() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["backend"])
}

// Flag to mark basic constraints valid when issuing non-CA certificates
func (r *SecretBackendRole) BasicConstraintsValidForNonCa() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["basicConstraintsValidForNonCa"])
}

// Flag to specify certificates for client use
func (r *SecretBackendRole) ClientFlag() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["clientFlag"])
}

// Flag to specify certificates for code signing use
func (r *SecretBackendRole) CodeSigningFlag() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["codeSigningFlag"])
}

// The country of generated certificates
func (r *SecretBackendRole) Countries() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["countries"])
}

// Flag to specify certificates for email protection use
func (r *SecretBackendRole) EmailProtectionFlag() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["emailProtectionFlag"])
}

// Flag to allow only valid host names
func (r *SecretBackendRole) EnforceHostnames() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["enforceHostnames"])
}

// Specify the allowed extended key usage constraint on issued certificates
func (r *SecretBackendRole) ExtKeyUsages() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["extKeyUsages"])
}

// Flag to generate leases with certificates
func (r *SecretBackendRole) GenerateLease() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["generateLease"])
}

// The number of bits of generated keys
func (r *SecretBackendRole) KeyBits() *pulumi.IntOutput {
	return (*pulumi.IntOutput)(r.s.State["keyBits"])
}

// The type of generated keys
func (r *SecretBackendRole) KeyType() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["keyType"])
}

// Specify the allowed key usage constraint on issued certificates
func (r *SecretBackendRole) KeyUsages() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["keyUsages"])
}

// The locality of generated certificates
func (r *SecretBackendRole) Localities() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["localities"])
}

// The maximum TTL
func (r *SecretBackendRole) MaxTtl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["maxTtl"])
}

// The name to identify this role within the backend. Must be unique within the backend.
func (r *SecretBackendRole) Name() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["name"])
}

// Flag to not store certificates in the storage backend
func (r *SecretBackendRole) NoStore() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["noStore"])
}

// The organization of generated certificates
func (r *SecretBackendRole) Organizations() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["organizations"])
}

// The organization unit of generated certificates
func (r *SecretBackendRole) OrganizationUnit() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["organizationUnit"])
}

// Specify the list of allowed policies IODs
func (r *SecretBackendRole) PolicyIdentifiers() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["policyIdentifiers"])
}

// The postal code of generated certificates
func (r *SecretBackendRole) PostalCodes() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["postalCodes"])
}

// The province of generated certificates
func (r *SecretBackendRole) Provinces() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["provinces"])
}

// Flag to force CN usage
func (r *SecretBackendRole) RequireCn() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["requireCn"])
}

// Flag to specify certificates for server use
func (r *SecretBackendRole) ServerFlag() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["serverFlag"])
}

// The street address of generated certificates
func (r *SecretBackendRole) StreetAddresses() *pulumi.ArrayOutput {
	return (*pulumi.ArrayOutput)(r.s.State["streetAddresses"])
}

// The TTL
func (r *SecretBackendRole) Ttl() *pulumi.StringOutput {
	return (*pulumi.StringOutput)(r.s.State["ttl"])
}

// Flag to use the CN in the CSR
func (r *SecretBackendRole) UseCsrCommonName() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["useCsrCommonName"])
}

// Flag to use the SANs in the CSR
func (r *SecretBackendRole) UseCsrSans() *pulumi.BoolOutput {
	return (*pulumi.BoolOutput)(r.s.State["useCsrSans"])
}

// Input properties used for looking up and filtering SecretBackendRole resources.
type SecretBackendRoleState struct {
	// Flag to allow any name
	AllowAnyName interface{}
	// Flag to allow certificates matching the actual domain
	AllowBareDomains interface{}
	// Flag to allow names containing glob patterns.
	AllowGlobDomains interface{}
	// Flag to allow IP SANs
	AllowIpSans interface{}
	// Flag to allow certificates for localhost
	AllowLocalhost interface{}
	// Flag to allow certificates matching subdomains
	AllowSubdomains interface{}
	// List of allowed domains for certificates
	AllowedDomains interface{}
	// Defines allowed custom SANs
	AllowedOtherSans interface{}
	// Defines allowed URI SANs
	AllowedUriSans interface{}
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend interface{}
	// Flag to mark basic constraints valid when issuing non-CA certificates
	BasicConstraintsValidForNonCa interface{}
	// Flag to specify certificates for client use
	ClientFlag interface{}
	// Flag to specify certificates for code signing use
	CodeSigningFlag interface{}
	// The country of generated certificates
	Countries interface{}
	// Flag to specify certificates for email protection use
	EmailProtectionFlag interface{}
	// Flag to allow only valid host names
	EnforceHostnames interface{}
	// Specify the allowed extended key usage constraint on issued certificates
	ExtKeyUsages interface{}
	// Flag to generate leases with certificates
	GenerateLease interface{}
	// The number of bits of generated keys
	KeyBits interface{}
	// The type of generated keys
	KeyType interface{}
	// Specify the allowed key usage constraint on issued certificates
	KeyUsages interface{}
	// The locality of generated certificates
	Localities interface{}
	// The maximum TTL
	MaxTtl interface{}
	// The name to identify this role within the backend. Must be unique within the backend.
	Name interface{}
	// Flag to not store certificates in the storage backend
	NoStore interface{}
	// The organization of generated certificates
	Organizations interface{}
	// The organization unit of generated certificates
	OrganizationUnit interface{}
	// Specify the list of allowed policies IODs
	PolicyIdentifiers interface{}
	// The postal code of generated certificates
	PostalCodes interface{}
	// The province of generated certificates
	Provinces interface{}
	// Flag to force CN usage
	RequireCn interface{}
	// Flag to specify certificates for server use
	ServerFlag interface{}
	// The street address of generated certificates
	StreetAddresses interface{}
	// The TTL
	Ttl interface{}
	// Flag to use the CN in the CSR
	UseCsrCommonName interface{}
	// Flag to use the SANs in the CSR
	UseCsrSans interface{}
}

// The set of arguments for constructing a SecretBackendRole resource.
type SecretBackendRoleArgs struct {
	// Flag to allow any name
	AllowAnyName interface{}
	// Flag to allow certificates matching the actual domain
	AllowBareDomains interface{}
	// Flag to allow names containing glob patterns.
	AllowGlobDomains interface{}
	// Flag to allow IP SANs
	AllowIpSans interface{}
	// Flag to allow certificates for localhost
	AllowLocalhost interface{}
	// Flag to allow certificates matching subdomains
	AllowSubdomains interface{}
	// List of allowed domains for certificates
	AllowedDomains interface{}
	// Defines allowed custom SANs
	AllowedOtherSans interface{}
	// Defines allowed URI SANs
	AllowedUriSans interface{}
	// The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
	Backend interface{}
	// Flag to mark basic constraints valid when issuing non-CA certificates
	BasicConstraintsValidForNonCa interface{}
	// Flag to specify certificates for client use
	ClientFlag interface{}
	// Flag to specify certificates for code signing use
	CodeSigningFlag interface{}
	// The country of generated certificates
	Countries interface{}
	// Flag to specify certificates for email protection use
	EmailProtectionFlag interface{}
	// Flag to allow only valid host names
	EnforceHostnames interface{}
	// Specify the allowed extended key usage constraint on issued certificates
	ExtKeyUsages interface{}
	// Flag to generate leases with certificates
	GenerateLease interface{}
	// The number of bits of generated keys
	KeyBits interface{}
	// The type of generated keys
	KeyType interface{}
	// Specify the allowed key usage constraint on issued certificates
	KeyUsages interface{}
	// The locality of generated certificates
	Localities interface{}
	// The maximum TTL
	MaxTtl interface{}
	// The name to identify this role within the backend. Must be unique within the backend.
	Name interface{}
	// Flag to not store certificates in the storage backend
	NoStore interface{}
	// The organization of generated certificates
	Organizations interface{}
	// The organization unit of generated certificates
	OrganizationUnit interface{}
	// Specify the list of allowed policies IODs
	PolicyIdentifiers interface{}
	// The postal code of generated certificates
	PostalCodes interface{}
	// The province of generated certificates
	Provinces interface{}
	// Flag to force CN usage
	RequireCn interface{}
	// Flag to specify certificates for server use
	ServerFlag interface{}
	// The street address of generated certificates
	StreetAddresses interface{}
	// The TTL
	Ttl interface{}
	// Flag to use the CN in the CSR
	UseCsrCommonName interface{}
	// Flag to use the SANs in the CSR
	UseCsrSans interface{}
}
