// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a role in an [GCP auth backend within Vault](https://www.vaultproject.io/docs/auth/gcp.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const gcpAuthBackend = new vault.AuthBackend("gcpAuthBackend", {
 *     path: "gcp",
 *     type: "gcp",
 * });
 * const gcpAuthBackendRole = new vault.gcp.AuthBackendRole("gcpAuthBackendRole", {
 *     backend: gcpAuthBackend.path,
 *     projectId: "foo-bar-baz",
 *     boundServiceAccounts: ["database-server@foo-bar-baz.iam.gserviceaccount.com"],
 *     tokenPolicies: ["database-server"],
 * });
 * ```
 *
 * ## Import
 *
 * GCP authentication roles can be imported using the `path`, e.g.
 *
 * ```sh
 *  $ pulumi import vault:gcp/authBackendRole:AuthBackendRole my_role auth/gcp/role/my_role
 * ```
 */
export class AuthBackendRole extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackendRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendRoleState, opts?: pulumi.CustomResourceOptions): AuthBackendRole {
        return new AuthBackendRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:gcp/authBackendRole:AuthBackendRole';

    /**
     * Returns true if the given object is an instance of AuthBackendRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackendRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackendRole.__pulumiType;
    }

    public readonly addGroupAliases!: pulumi.Output<boolean>;
    /**
     * A flag to determine if this role should allow GCE instances to authenticate by inferring service accounts from the GCE identity metadata token.
     */
    public readonly allowGceInference!: pulumi.Output<boolean>;
    /**
     * Path to the mounted GCP auth backend
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * The instance groups that an authorized instance must belong to in order to be authenticated. If specified, either `boundZones` or `boundRegions` must be set too.
     */
    public readonly boundInstanceGroups!: pulumi.Output<string[]>;
    /**
     * A comma-separated list of GCP labels formatted as `"key:value"` strings that must be set on authorized GCE instances. Because GCP labels are not currently ACL'd, we recommend that this be used in conjunction with other restrictions.
     */
    public readonly boundLabels!: pulumi.Output<string[]>;
    /**
     * GCP Projects that the role exists within
     */
    public readonly boundProjects!: pulumi.Output<string[] | undefined>;
    /**
     * The list of regions that a GCE instance must belong to in order to be authenticated. If boundInstanceGroups is provided, it is assumed to be a regional group and the group must belong to this region. If boundZones are provided, this attribute is ignored.
     */
    public readonly boundRegions!: pulumi.Output<string[]>;
    /**
     * GCP Service Accounts allowed to issue tokens under this role. (Note: **Required** if role is `iam`)
     */
    public readonly boundServiceAccounts!: pulumi.Output<string[]>;
    /**
     * The list of zones that a GCE instance must belong to in order to be authenticated. If boundInstanceGroups is provided, it is assumed to be a zonal group and the group must belong to this zone.
     */
    public readonly boundZones!: pulumi.Output<string[]>;
    /**
     * The number of seconds past the time of authentication that the login param JWT must expire within. For example, if a user attempts to login with a token that expires within an hour and this is set to 15 minutes, Vault will return an error prompting the user to create a new signed JWT with a shorter `exp`. The GCE metadata tokens currently do not allow the `exp` claim to be customized.
     */
    public readonly maxJwtExp!: pulumi.Output<string>;
    /**
     * The maximum allowed lifetime of tokens
     * issued using this role, provided as a number of seconds.
     *
     * @deprecated use `token_max_ttl` instead if you are running Vault >= 1.2
     */
    public readonly maxTtl!: pulumi.Output<string>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     *
     * @deprecated use `token_period` instead if you are running Vault >= 1.2
     */
    public readonly period!: pulumi.Output<string>;
    /**
     * An array of strings
     * specifying the policies to be set on tokens issued using this role.
     *
     * @deprecated use `token_policies` instead if you are running Vault >= 1.2
     */
    public readonly policies!: pulumi.Output<string[]>;
    /**
     * Name of the GCP role
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     */
    public readonly tokenBoundCidrs!: pulumi.Output<string[] | undefined>;
    /**
     * If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
     * `tokenMaxTtl` would otherwise allow a renewal.
     */
    public readonly tokenExplicitMaxTtl!: pulumi.Output<number | undefined>;
    /**
     * The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    public readonly tokenMaxTtl!: pulumi.Output<number | undefined>;
    /**
     * If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     */
    public readonly tokenNoDefaultPolicy!: pulumi.Output<boolean | undefined>;
    /**
     * The
     * [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
     * if any, in number of seconds to set on the token.
     */
    public readonly tokenNumUses!: pulumi.Output<number | undefined>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    public readonly tokenPeriod!: pulumi.Output<number | undefined>;
    /**
     * List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     */
    public readonly tokenPolicies!: pulumi.Output<string[] | undefined>;
    /**
     * The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    public readonly tokenTtl!: pulumi.Output<number | undefined>;
    /**
     * The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount's tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     */
    public readonly tokenType!: pulumi.Output<string | undefined>;
    /**
     * The TTL period of tokens issued
     * using this role, provided as a number of seconds.
     *
     * @deprecated use `token_ttl` instead if you are running Vault >= 1.2
     */
    public readonly ttl!: pulumi.Output<string>;
    /**
     * Type of GCP authentication role (either `gce` or `iam`)
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a AuthBackendRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthBackendRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendRoleArgs | AuthBackendRoleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AuthBackendRoleState | undefined;
            inputs["addGroupAliases"] = state ? state.addGroupAliases : undefined;
            inputs["allowGceInference"] = state ? state.allowGceInference : undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["boundInstanceGroups"] = state ? state.boundInstanceGroups : undefined;
            inputs["boundLabels"] = state ? state.boundLabels : undefined;
            inputs["boundProjects"] = state ? state.boundProjects : undefined;
            inputs["boundRegions"] = state ? state.boundRegions : undefined;
            inputs["boundServiceAccounts"] = state ? state.boundServiceAccounts : undefined;
            inputs["boundZones"] = state ? state.boundZones : undefined;
            inputs["maxJwtExp"] = state ? state.maxJwtExp : undefined;
            inputs["maxTtl"] = state ? state.maxTtl : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["policies"] = state ? state.policies : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["tokenBoundCidrs"] = state ? state.tokenBoundCidrs : undefined;
            inputs["tokenExplicitMaxTtl"] = state ? state.tokenExplicitMaxTtl : undefined;
            inputs["tokenMaxTtl"] = state ? state.tokenMaxTtl : undefined;
            inputs["tokenNoDefaultPolicy"] = state ? state.tokenNoDefaultPolicy : undefined;
            inputs["tokenNumUses"] = state ? state.tokenNumUses : undefined;
            inputs["tokenPeriod"] = state ? state.tokenPeriod : undefined;
            inputs["tokenPolicies"] = state ? state.tokenPolicies : undefined;
            inputs["tokenTtl"] = state ? state.tokenTtl : undefined;
            inputs["tokenType"] = state ? state.tokenType : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as AuthBackendRoleArgs | undefined;
            if ((!args || args.role === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'role'");
            }
            if ((!args || args.type === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'type'");
            }
            inputs["addGroupAliases"] = args ? args.addGroupAliases : undefined;
            inputs["allowGceInference"] = args ? args.allowGceInference : undefined;
            inputs["backend"] = args ? args.backend : undefined;
            inputs["boundInstanceGroups"] = args ? args.boundInstanceGroups : undefined;
            inputs["boundLabels"] = args ? args.boundLabels : undefined;
            inputs["boundProjects"] = args ? args.boundProjects : undefined;
            inputs["boundRegions"] = args ? args.boundRegions : undefined;
            inputs["boundServiceAccounts"] = args ? args.boundServiceAccounts : undefined;
            inputs["boundZones"] = args ? args.boundZones : undefined;
            inputs["maxJwtExp"] = args ? args.maxJwtExp : undefined;
            inputs["maxTtl"] = args ? args.maxTtl : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["policies"] = args ? args.policies : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["tokenBoundCidrs"] = args ? args.tokenBoundCidrs : undefined;
            inputs["tokenExplicitMaxTtl"] = args ? args.tokenExplicitMaxTtl : undefined;
            inputs["tokenMaxTtl"] = args ? args.tokenMaxTtl : undefined;
            inputs["tokenNoDefaultPolicy"] = args ? args.tokenNoDefaultPolicy : undefined;
            inputs["tokenNumUses"] = args ? args.tokenNumUses : undefined;
            inputs["tokenPeriod"] = args ? args.tokenPeriod : undefined;
            inputs["tokenPolicies"] = args ? args.tokenPolicies : undefined;
            inputs["tokenTtl"] = args ? args.tokenTtl : undefined;
            inputs["tokenType"] = args ? args.tokenType : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["type"] = args ? args.type : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AuthBackendRole.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendRole resources.
 */
export interface AuthBackendRoleState {
    readonly addGroupAliases?: pulumi.Input<boolean>;
    /**
     * A flag to determine if this role should allow GCE instances to authenticate by inferring service accounts from the GCE identity metadata token.
     */
    readonly allowGceInference?: pulumi.Input<boolean>;
    /**
     * Path to the mounted GCP auth backend
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * The instance groups that an authorized instance must belong to in order to be authenticated. If specified, either `boundZones` or `boundRegions` must be set too.
     */
    readonly boundInstanceGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A comma-separated list of GCP labels formatted as `"key:value"` strings that must be set on authorized GCE instances. Because GCP labels are not currently ACL'd, we recommend that this be used in conjunction with other restrictions.
     */
    readonly boundLabels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * GCP Projects that the role exists within
     */
    readonly boundProjects?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of regions that a GCE instance must belong to in order to be authenticated. If boundInstanceGroups is provided, it is assumed to be a regional group and the group must belong to this region. If boundZones are provided, this attribute is ignored.
     */
    readonly boundRegions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * GCP Service Accounts allowed to issue tokens under this role. (Note: **Required** if role is `iam`)
     */
    readonly boundServiceAccounts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of zones that a GCE instance must belong to in order to be authenticated. If boundInstanceGroups is provided, it is assumed to be a zonal group and the group must belong to this zone.
     */
    readonly boundZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The number of seconds past the time of authentication that the login param JWT must expire within. For example, if a user attempts to login with a token that expires within an hour and this is set to 15 minutes, Vault will return an error prompting the user to create a new signed JWT with a shorter `exp`. The GCE metadata tokens currently do not allow the `exp` claim to be customized.
     */
    readonly maxJwtExp?: pulumi.Input<string>;
    /**
     * The maximum allowed lifetime of tokens
     * issued using this role, provided as a number of seconds.
     *
     * @deprecated use `token_max_ttl` instead if you are running Vault >= 1.2
     */
    readonly maxTtl?: pulumi.Input<string>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     *
     * @deprecated use `token_period` instead if you are running Vault >= 1.2
     */
    readonly period?: pulumi.Input<string>;
    /**
     * An array of strings
     * specifying the policies to be set on tokens issued using this role.
     *
     * @deprecated use `token_policies` instead if you are running Vault >= 1.2
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the GCP role
     */
    readonly role?: pulumi.Input<string>;
    /**
     * List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     */
    readonly tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
     * `tokenMaxTtl` would otherwise allow a renewal.
     */
    readonly tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    readonly tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     */
    readonly tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The
     * [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
     * if any, in number of seconds to set on the token.
     */
    readonly tokenNumUses?: pulumi.Input<number>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    readonly tokenPeriod?: pulumi.Input<number>;
    /**
     * List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     */
    readonly tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    readonly tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount's tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     */
    readonly tokenType?: pulumi.Input<string>;
    /**
     * The TTL period of tokens issued
     * using this role, provided as a number of seconds.
     *
     * @deprecated use `token_ttl` instead if you are running Vault >= 1.2
     */
    readonly ttl?: pulumi.Input<string>;
    /**
     * Type of GCP authentication role (either `gce` or `iam`)
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthBackendRole resource.
 */
export interface AuthBackendRoleArgs {
    readonly addGroupAliases?: pulumi.Input<boolean>;
    /**
     * A flag to determine if this role should allow GCE instances to authenticate by inferring service accounts from the GCE identity metadata token.
     */
    readonly allowGceInference?: pulumi.Input<boolean>;
    /**
     * Path to the mounted GCP auth backend
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * The instance groups that an authorized instance must belong to in order to be authenticated. If specified, either `boundZones` or `boundRegions` must be set too.
     */
    readonly boundInstanceGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A comma-separated list of GCP labels formatted as `"key:value"` strings that must be set on authorized GCE instances. Because GCP labels are not currently ACL'd, we recommend that this be used in conjunction with other restrictions.
     */
    readonly boundLabels?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * GCP Projects that the role exists within
     */
    readonly boundProjects?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of regions that a GCE instance must belong to in order to be authenticated. If boundInstanceGroups is provided, it is assumed to be a regional group and the group must belong to this region. If boundZones are provided, this attribute is ignored.
     */
    readonly boundRegions?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * GCP Service Accounts allowed to issue tokens under this role. (Note: **Required** if role is `iam`)
     */
    readonly boundServiceAccounts?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The list of zones that a GCE instance must belong to in order to be authenticated. If boundInstanceGroups is provided, it is assumed to be a zonal group and the group must belong to this zone.
     */
    readonly boundZones?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The number of seconds past the time of authentication that the login param JWT must expire within. For example, if a user attempts to login with a token that expires within an hour and this is set to 15 minutes, Vault will return an error prompting the user to create a new signed JWT with a shorter `exp`. The GCE metadata tokens currently do not allow the `exp` claim to be customized.
     */
    readonly maxJwtExp?: pulumi.Input<string>;
    /**
     * The maximum allowed lifetime of tokens
     * issued using this role, provided as a number of seconds.
     *
     * @deprecated use `token_max_ttl` instead if you are running Vault >= 1.2
     */
    readonly maxTtl?: pulumi.Input<string>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     *
     * @deprecated use `token_period` instead if you are running Vault >= 1.2
     */
    readonly period?: pulumi.Input<string>;
    /**
     * An array of strings
     * specifying the policies to be set on tokens issued using this role.
     *
     * @deprecated use `token_policies` instead if you are running Vault >= 1.2
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the GCP role
     */
    readonly role: pulumi.Input<string>;
    /**
     * List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     */
    readonly tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
     * `tokenMaxTtl` would otherwise allow a renewal.
     */
    readonly tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    readonly tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     */
    readonly tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The
     * [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
     * if any, in number of seconds to set on the token.
     */
    readonly tokenNumUses?: pulumi.Input<number>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    readonly tokenPeriod?: pulumi.Input<number>;
    /**
     * List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     */
    readonly tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    readonly tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount's tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     */
    readonly tokenType?: pulumi.Input<string>;
    /**
     * The TTL period of tokens issued
     * using this role, provided as a number of seconds.
     *
     * @deprecated use `token_ttl` instead if you are running Vault >= 1.2
     */
    readonly ttl?: pulumi.Input<string>;
    /**
     * Type of GCP authentication role (either `gce` or `iam`)
     */
    readonly type: pulumi.Input<string>;
}
