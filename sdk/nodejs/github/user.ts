// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages policy mappings for Github Users authenticated via Github. See the [Vault
 * documentation](https://www.vaultproject.io/docs/auth/github/) for more
 * information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const example = new vault.github.AuthBackend("example", {organization: "myorg"});
 * const tfUser = new vault.github.User("tfUser", {
 *     backend: example.id,
 *     user: "john.doe",
 *     policies: [
 *         "developer",
 *         "read-only",
 *     ],
 * });
 * ```
 *
 * ## Import
 *
 * Github user mappings can be imported using the `path`, e.g.
 *
 * ```sh
 *  $ pulumi import vault:github/user:User tf_user auth/github/map/users/john.doe
 * ```
 */
export class User extends pulumi.CustomResource {
    /**
     * Get an existing User resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: UserState, opts?: pulumi.CustomResourceOptions): User {
        return new User(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:github/user:User';

    /**
     * Returns true if the given object is an instance of User.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is User {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === User.__pulumiType;
    }

    /**
     * Path where the github auth backend is mounted. Defaults to `github`
     * if not specified.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * An array of strings specifying the policies to be set on tokens issued
     * using this role.
     */
    public readonly policies!: pulumi.Output<string[] | undefined>;
    /**
     * Specifies the blocks of IP addresses which are allowed to use the generated token
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    public readonly tokenBoundCidrs!: pulumi.Output<string[] | undefined>;
    /**
     * Generated Token's Explicit Maximum TTL in seconds
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    public readonly tokenExplicitMaxTtl!: pulumi.Output<number | undefined>;
    /**
     * The maximum lifetime of the generated token
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    public readonly tokenMaxTtl!: pulumi.Output<number | undefined>;
    /**
     * If true, the 'default' policy will not automatically be added to generated tokens
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    public readonly tokenNoDefaultPolicy!: pulumi.Output<boolean | undefined>;
    /**
     * The maximum number of times a token may be used, a value of zero means unlimited
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    public readonly tokenNumUses!: pulumi.Output<number | undefined>;
    /**
     * Generated Token's Period
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    public readonly tokenPeriod!: pulumi.Output<number | undefined>;
    /**
     * Generated Token's Policies
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    public readonly tokenPolicies!: pulumi.Output<string[] | undefined>;
    /**
     * The initial ttl of the token to generate in seconds
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    public readonly tokenTtl!: pulumi.Output<number | undefined>;
    /**
     * The type of token to generate, service or batch
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    public readonly tokenType!: pulumi.Output<string | undefined>;
    /**
     * GitHub user name.
     */
    public readonly user!: pulumi.Output<string>;

    /**
     * Create a User resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: UserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: UserArgs | UserState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as UserState | undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["policies"] = state ? state.policies : undefined;
            inputs["tokenBoundCidrs"] = state ? state.tokenBoundCidrs : undefined;
            inputs["tokenExplicitMaxTtl"] = state ? state.tokenExplicitMaxTtl : undefined;
            inputs["tokenMaxTtl"] = state ? state.tokenMaxTtl : undefined;
            inputs["tokenNoDefaultPolicy"] = state ? state.tokenNoDefaultPolicy : undefined;
            inputs["tokenNumUses"] = state ? state.tokenNumUses : undefined;
            inputs["tokenPeriod"] = state ? state.tokenPeriod : undefined;
            inputs["tokenPolicies"] = state ? state.tokenPolicies : undefined;
            inputs["tokenTtl"] = state ? state.tokenTtl : undefined;
            inputs["tokenType"] = state ? state.tokenType : undefined;
            inputs["user"] = state ? state.user : undefined;
        } else {
            const args = argsOrState as UserArgs | undefined;
            if ((!args || args.user === undefined) && !opts.urn) {
                throw new Error("Missing required property 'user'");
            }
            inputs["backend"] = args ? args.backend : undefined;
            inputs["policies"] = args ? args.policies : undefined;
            inputs["tokenBoundCidrs"] = args ? args.tokenBoundCidrs : undefined;
            inputs["tokenExplicitMaxTtl"] = args ? args.tokenExplicitMaxTtl : undefined;
            inputs["tokenMaxTtl"] = args ? args.tokenMaxTtl : undefined;
            inputs["tokenNoDefaultPolicy"] = args ? args.tokenNoDefaultPolicy : undefined;
            inputs["tokenNumUses"] = args ? args.tokenNumUses : undefined;
            inputs["tokenPeriod"] = args ? args.tokenPeriod : undefined;
            inputs["tokenPolicies"] = args ? args.tokenPolicies : undefined;
            inputs["tokenTtl"] = args ? args.tokenTtl : undefined;
            inputs["tokenType"] = args ? args.tokenType : undefined;
            inputs["user"] = args ? args.user : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(User.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering User resources.
 */
export interface UserState {
    /**
     * Path where the github auth backend is mounted. Defaults to `github`
     * if not specified.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * An array of strings specifying the policies to be set on tokens issued
     * using this role.
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the blocks of IP addresses which are allowed to use the generated token
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    readonly tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Generated Token's Explicit Maximum TTL in seconds
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    readonly tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime of the generated token
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    readonly tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If true, the 'default' policy will not automatically be added to generated tokens
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    readonly tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The maximum number of times a token may be used, a value of zero means unlimited
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    readonly tokenNumUses?: pulumi.Input<number>;
    /**
     * Generated Token's Period
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    readonly tokenPeriod?: pulumi.Input<number>;
    /**
     * Generated Token's Policies
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    readonly tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The initial ttl of the token to generate in seconds
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    readonly tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token to generate, service or batch
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    readonly tokenType?: pulumi.Input<string>;
    /**
     * GitHub user name.
     */
    readonly user?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a User resource.
 */
export interface UserArgs {
    /**
     * Path where the github auth backend is mounted. Defaults to `github`
     * if not specified.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * An array of strings specifying the policies to be set on tokens issued
     * using this role.
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Specifies the blocks of IP addresses which are allowed to use the generated token
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    readonly tokenBoundCidrs?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Generated Token's Explicit Maximum TTL in seconds
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    readonly tokenExplicitMaxTtl?: pulumi.Input<number>;
    /**
     * The maximum lifetime of the generated token
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    readonly tokenMaxTtl?: pulumi.Input<number>;
    /**
     * If true, the 'default' policy will not automatically be added to generated tokens
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    readonly tokenNoDefaultPolicy?: pulumi.Input<boolean>;
    /**
     * The maximum number of times a token may be used, a value of zero means unlimited
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    readonly tokenNumUses?: pulumi.Input<number>;
    /**
     * Generated Token's Period
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    readonly tokenPeriod?: pulumi.Input<number>;
    /**
     * Generated Token's Policies
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    readonly tokenPolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The initial ttl of the token to generate in seconds
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    readonly tokenTtl?: pulumi.Input<number>;
    /**
     * The type of token to generate, service or batch
     *
     * @deprecated This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.
     */
    readonly tokenType?: pulumi.Input<string>;
    /**
     * GitHub user name.
     */
    readonly user: pulumi.Input<string>;
}
