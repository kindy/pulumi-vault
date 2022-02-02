// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an Azure auth backend role in a Vault server. Roles constrain the
 * instances or principals that can perform the login operation against the
 * backend. See the [Vault
 * documentation](https://www.vaultproject.io/docs/auth/azure.html) for more
 * information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const azure = new vault.AuthBackend("azure", {type: "azure"});
 * const example = new vault.azure.AuthBackendRole("example", {
 *     backend: azure.path,
 *     role: "test-role",
 *     boundSubscriptionIds: ["11111111-2222-3333-4444-555555555555"],
 *     boundResourceGroups: ["123456789012"],
 *     tokenTtl: 60,
 *     tokenMaxTtl: 120,
 *     tokenPolicies: [
 *         "default",
 *         "dev",
 *         "prod",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Azure auth backend roles can be imported using `auth/`, the `backend` path, `/role/`, and the `role` name e.g.
 *
 * ```sh
 *  $ pulumi import vault:azure/authBackendRole:AuthBackendRole example auth/azure/role/test-role
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
    public static readonly __pulumiType = 'vault:azure/authBackendRole:AuthBackendRole';

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
     * Unique name of the auth backend to configure.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * If set, defines a constraint on the groups
     * that can perform the login operation that they should be using the group
     * ID specified by this field.
     */
    public readonly boundGroupIds!: pulumi.Output<string[] | undefined>;
    /**
     * If set, defines a constraint on the virtual machines
     * that can perform the login operation that the location in their identity
     * document must match the one specified by this field.
     */
    public readonly boundLocations!: pulumi.Output<string[] | undefined>;
    /**
     * If set, defines a constraint on the virtual
     * machines that can perform the login operation that they be associated with
     * the resource group that matches the value specified by this field.
     */
    public readonly boundResourceGroups!: pulumi.Output<string[] | undefined>;
    /**
     * If set, defines a constraint on the virtual
     * machines that can perform the login operation that they must match the scale set
     * specified by this field.
     */
    public readonly boundScaleSets!: pulumi.Output<string[] | undefined>;
    /**
     * If set, defines a constraint on the
     * service principals that can perform the login operation that they should be possess
     * the ids specified by this field.
     */
    public readonly boundServicePrincipalIds!: pulumi.Output<string[] | undefined>;
    /**
     * If set, defines a constraint on the subscriptions
     * that can perform the login operation to ones which  matches the value specified by this
     * field.
     */
    public readonly boundSubscriptionIds!: pulumi.Output<string[] | undefined>;
    /**
     * The name of the role.
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
     * Create a AuthBackendRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthBackendRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendRoleArgs | AuthBackendRoleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthBackendRoleState | undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["boundGroupIds"] = state ? state.boundGroupIds : undefined;
            resourceInputs["boundLocations"] = state ? state.boundLocations : undefined;
            resourceInputs["boundResourceGroups"] = state ? state.boundResourceGroups : undefined;
            resourceInputs["boundScaleSets"] = state ? state.boundScaleSets : undefined;
            resourceInputs["boundServicePrincipalIds"] = state ? state.boundServicePrincipalIds : undefined;
            resourceInputs["boundSubscriptionIds"] = state ? state.boundSubscriptionIds : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["tokenBoundCidrs"] = state ? state.tokenBoundCidrs : undefined;
            resourceInputs["tokenExplicitMaxTtl"] = state ? state.tokenExplicitMaxTtl : undefined;
            resourceInputs["tokenMaxTtl"] = state ? state.tokenMaxTtl : undefined;
            resourceInputs["tokenNoDefaultPolicy"] = state ? state.tokenNoDefaultPolicy : undefined;
            resourceInputs["tokenNumUses"] = state ? state.tokenNumUses : undefined;
            resourceInputs["tokenPeriod"] = state ? state.tokenPeriod : undefined;
            resourceInputs["tokenPolicies"] = state ? state.tokenPolicies : undefined;
            resourceInputs["tokenTtl"] = state ? state.tokenTtl : undefined;
            resourceInputs["tokenType"] = state ? state.tokenType : undefined;
        } else {
            const args = argsOrState as AuthBackendRoleArgs | undefined;
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["boundGroupIds"] = args ? args.boundGroupIds : undefined;
            resourceInputs["boundLocations"] = args ? args.boundLocations : undefined;
            resourceInputs["boundResourceGroups"] = args ? args.boundResourceGroups : undefined;
            resourceInputs["boundScaleSets"] = args ? args.boundScaleSets : undefined;
            resourceInputs["boundServicePrincipalIds"] = args ? args.boundServicePrincipalIds : undefined;
            resourceInputs["boundSubscriptionIds"] = args ? args.boundSubscriptionIds : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["tokenBoundCidrs"] = args ? args.tokenBoundCidrs : undefined;
            resourceInputs["tokenExplicitMaxTtl"] = args ? args.tokenExplicitMaxTtl : undefined;
            resourceInputs["tokenMaxTtl"] = args ? args.tokenMaxTtl : undefined;
            resourceInputs["tokenNoDefaultPolicy"] = args ? args.tokenNoDefaultPolicy : undefined;
            resourceInputs["tokenNumUses"] = args ? args.tokenNumUses : undefined;
            resourceInputs["tokenPeriod"] = args ? args.tokenPeriod : undefined;
            resourceInputs["tokenPolicies"] = args ? args.tokenPolicies : undefined;
            resourceInputs["tokenTtl"] = args ? args.tokenTtl : undefined;
            resourceInputs["tokenType"] = args ? args.tokenType : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AuthBackendRole.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendRole resources.
 */
export interface AuthBackendRoleState {
    /**
     * Unique name of the auth backend to configure.
     */
    backend?: pulumi.Input<string>;
    /**
     * If set, defines a constraint on the groups
     * that can perform the login operation that they should be using the group
     * ID specified by this field.
     */
    boundGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, defines a constraint on the virtual machines
     * that can perform the login operation that the location in their identity
     * document must match the one specified by this field.
     */
    boundLocations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, defines a constraint on the virtual
     * machines that can perform the login operation that they be associated with
     * the resource group that matches the value specified by this field.
     */
    boundResourceGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, defines a constraint on the virtual
     * machines that can perform the login operation that they must match the scale set
     * specified by this field.
     */
    boundScaleSets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, defines a constraint on the
     * service principals that can perform the login operation that they should be possess
     * the ids specified by this field.
     */
    boundServicePrincipalIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, defines a constraint on the subscriptions
     * that can perform the login operation to ones which  matches the value specified by this
     * field.
     */
    boundSubscriptionIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the role.
     */
    role?: pulumi.Input<string>;
    /**
     * List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     */
    tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
     * `tokenMaxTtl` would otherwise allow a renewal.
     */
    tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     */
    tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The
     * [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
     * if any, in number of seconds to set on the token.
     */
    tokenNumUses?: pulumi.Input<number>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    tokenPeriod?: pulumi.Input<number>;
    /**
     * List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     */
    tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount's tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     */
    tokenType?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthBackendRole resource.
 */
export interface AuthBackendRoleArgs {
    /**
     * Unique name of the auth backend to configure.
     */
    backend?: pulumi.Input<string>;
    /**
     * If set, defines a constraint on the groups
     * that can perform the login operation that they should be using the group
     * ID specified by this field.
     */
    boundGroupIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, defines a constraint on the virtual machines
     * that can perform the login operation that the location in their identity
     * document must match the one specified by this field.
     */
    boundLocations?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, defines a constraint on the virtual
     * machines that can perform the login operation that they be associated with
     * the resource group that matches the value specified by this field.
     */
    boundResourceGroups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, defines a constraint on the virtual
     * machines that can perform the login operation that they must match the scale set
     * specified by this field.
     */
    boundScaleSets?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, defines a constraint on the
     * service principals that can perform the login operation that they should be possess
     * the ids specified by this field.
     */
    boundServicePrincipalIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, defines a constraint on the subscriptions
     * that can perform the login operation to ones which  matches the value specified by this
     * field.
     */
    boundSubscriptionIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of the role.
     */
    role: pulumi.Input<string>;
    /**
     * List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     */
    tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `tokenTtl` and
     * `tokenMaxTtl` would otherwise allow a renewal.
     */
    tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     */
    tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The
     * [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
     * if any, in number of seconds to set on the token.
     */
    tokenNumUses?: pulumi.Input<number>;
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token's TTL will be set to the
     * value of this field. Specified in seconds.
     */
    tokenPeriod?: pulumi.Input<number>;
    /**
     * List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     */
    tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     */
    tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount's tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     */
    tokenType?: pulumi.Input<string>;
}
