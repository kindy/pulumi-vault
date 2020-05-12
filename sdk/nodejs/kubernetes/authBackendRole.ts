// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Kubernetes auth backend role in a Vault server. See the [Vault
 * documentation](https://www.vaultproject.io/docs/auth/kubernetes.html) for more
 * information.
 * 
 * ## Example Usage
 * 
 * 
 * 
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 * 
 * const kubernetes = new vault.AuthBackend("kubernetes", {type: "kubernetes"});
 * const example = new vault.kubernetes.AuthBackendRole("example", {
 *     backend: kubernetes.path,
 *     roleName: "example-role",
 *     boundServiceAccountNames: ["example"],
 *     boundServiceAccountNamespaces: ["example"],
 *     tokenTtl: 3600,
 *     tokenPolicies: [
 *         "default",
 *         "dev",
 *         "prod",
 *     ],
 *     audience: "vault",
 * });
 * ```
 *
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/kubernetes_auth_backend_role.html.md.
 */
export class AuthBackendRole extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackendRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendRoleState, opts?: pulumi.CustomResourceOptions): AuthBackendRole {
        return new AuthBackendRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:kubernetes/authBackendRole:AuthBackendRole';

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

    /**
     * Audience claim to verify in the JWT.
     */
    public readonly audience!: pulumi.Output<string | undefined>;
    /**
     * Unique name of the kubernetes backend to configure.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * If set, a list of
     * CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
     */
    public readonly boundCidrs!: pulumi.Output<string[] | undefined>;
    /**
     * List of service account names able to access this role. If set to `["*"]` all names are allowed, both this and boundServiceAccountNamespaces can not be "*".
     */
    public readonly boundServiceAccountNames!: pulumi.Output<string[]>;
    /**
     * List of namespaces allowed to access this role. If set to `["*"]` all namespaces are allowed, both this and boundServiceAccountNames can not be set to "*".
     */
    public readonly boundServiceAccountNamespaces!: pulumi.Output<string[]>;
    /**
     * The maximum allowed lifetime of tokens
     * issued using this role, provided as a number of seconds.
     */
    public readonly maxTtl!: pulumi.Output<number | undefined>;
    /**
     * If set, puts a use-count
     * limitation on the issued token.
     */
    public readonly numUses!: pulumi.Output<number | undefined>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * An array of strings
     * specifying the policies to be set on tokens issued using this role.
     */
    public readonly policies!: pulumi.Output<string[] | undefined>;
    /**
     * Name of the role.
     */
    public readonly roleName!: pulumi.Output<string>;
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
     */
    public readonly ttl!: pulumi.Output<number | undefined>;

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
            inputs["audience"] = state ? state.audience : undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["boundCidrs"] = state ? state.boundCidrs : undefined;
            inputs["boundServiceAccountNames"] = state ? state.boundServiceAccountNames : undefined;
            inputs["boundServiceAccountNamespaces"] = state ? state.boundServiceAccountNamespaces : undefined;
            inputs["maxTtl"] = state ? state.maxTtl : undefined;
            inputs["numUses"] = state ? state.numUses : undefined;
            inputs["period"] = state ? state.period : undefined;
            inputs["policies"] = state ? state.policies : undefined;
            inputs["roleName"] = state ? state.roleName : undefined;
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
        } else {
            const args = argsOrState as AuthBackendRoleArgs | undefined;
            if (!args || args.boundServiceAccountNames === undefined) {
                throw new Error("Missing required property 'boundServiceAccountNames'");
            }
            if (!args || args.boundServiceAccountNamespaces === undefined) {
                throw new Error("Missing required property 'boundServiceAccountNamespaces'");
            }
            if (!args || args.roleName === undefined) {
                throw new Error("Missing required property 'roleName'");
            }
            inputs["audience"] = args ? args.audience : undefined;
            inputs["backend"] = args ? args.backend : undefined;
            inputs["boundCidrs"] = args ? args.boundCidrs : undefined;
            inputs["boundServiceAccountNames"] = args ? args.boundServiceAccountNames : undefined;
            inputs["boundServiceAccountNamespaces"] = args ? args.boundServiceAccountNamespaces : undefined;
            inputs["maxTtl"] = args ? args.maxTtl : undefined;
            inputs["numUses"] = args ? args.numUses : undefined;
            inputs["period"] = args ? args.period : undefined;
            inputs["policies"] = args ? args.policies : undefined;
            inputs["roleName"] = args ? args.roleName : undefined;
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
    /**
     * Audience claim to verify in the JWT.
     */
    readonly audience?: pulumi.Input<string>;
    /**
     * Unique name of the kubernetes backend to configure.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * If set, a list of
     * CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
     * 
     * @deprecated use `token_bound_cidrs` instead if you are running Vault >= 1.2
     */
    readonly boundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of service account names able to access this role. If set to `["*"]` all names are allowed, both this and boundServiceAccountNamespaces can not be "*".
     */
    readonly boundServiceAccountNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of namespaces allowed to access this role. If set to `["*"]` all namespaces are allowed, both this and boundServiceAccountNames can not be set to "*".
     */
    readonly boundServiceAccountNamespaces?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The maximum allowed lifetime of tokens
     * issued using this role, provided as a number of seconds.
     * 
     * @deprecated use `token_max_ttl` instead if you are running Vault >= 1.2
     */
    readonly maxTtl?: pulumi.Input<number>;
    /**
     * If set, puts a use-count
     * limitation on the issued token.
     * 
     * @deprecated use `token_num_uses` instead if you are running Vault >= 1.2
     */
    readonly numUses?: pulumi.Input<number>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     * 
     * @deprecated use `token_period` instead if you are running Vault >= 1.2
     */
    readonly period?: pulumi.Input<number>;
    /**
     * An array of strings
     * specifying the policies to be set on tokens issued using this role.
     * 
     * @deprecated use `token_policies` instead if you are running Vault >= 1.2
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the role.
     */
    readonly roleName?: pulumi.Input<string>;
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
    readonly ttl?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a AuthBackendRole resource.
 */
export interface AuthBackendRoleArgs {
    /**
     * Audience claim to verify in the JWT.
     */
    readonly audience?: pulumi.Input<string>;
    /**
     * Unique name of the kubernetes backend to configure.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * If set, a list of
     * CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
     * 
     * @deprecated use `token_bound_cidrs` instead if you are running Vault >= 1.2
     */
    readonly boundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of service account names able to access this role. If set to `["*"]` all names are allowed, both this and boundServiceAccountNamespaces can not be "*".
     */
    readonly boundServiceAccountNames: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * List of namespaces allowed to access this role. If set to `["*"]` all namespaces are allowed, both this and boundServiceAccountNames can not be set to "*".
     */
    readonly boundServiceAccountNamespaces: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The maximum allowed lifetime of tokens
     * issued using this role, provided as a number of seconds.
     * 
     * @deprecated use `token_max_ttl` instead if you are running Vault >= 1.2
     */
    readonly maxTtl?: pulumi.Input<number>;
    /**
     * If set, puts a use-count
     * limitation on the issued token.
     * 
     * @deprecated use `token_num_uses` instead if you are running Vault >= 1.2
     */
    readonly numUses?: pulumi.Input<number>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     * 
     * @deprecated use `token_period` instead if you are running Vault >= 1.2
     */
    readonly period?: pulumi.Input<number>;
    /**
     * An array of strings
     * specifying the policies to be set on tokens issued using this role.
     * 
     * @deprecated use `token_policies` instead if you are running Vault >= 1.2
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the role.
     */
    readonly roleName: pulumi.Input<string>;
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
    readonly ttl?: pulumi.Input<number>;
}
