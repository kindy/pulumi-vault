// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Manage rate limit quotas which enforce API rate limiting using a token bucket algorithm.
 * A rate limit quota can be created at the root level or defined on a namespace or mount by
 * specifying a path when creating the quota.
 *
 * See [Vault's Documentation](https://www.vaultproject.io/docs/concepts/resource-quotas) for more
 * information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const global = new vault.QuotaRateLimit("global", {
 *     path: "",
 *     rate: 100,
 * });
 * ```
 *
 * ## Import
 *
 * Rate limit quotas can be imported using their names
 *
 * ```sh
 *  $ pulumi import vault:index/quotaRateLimit:QuotaRateLimit global global
 * ```
 */
export class QuotaRateLimit extends pulumi.CustomResource {
    /**
     * Get an existing QuotaRateLimit resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: QuotaRateLimitState, opts?: pulumi.CustomResourceOptions): QuotaRateLimit {
        return new QuotaRateLimit(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:index/quotaRateLimit:QuotaRateLimit';

    /**
     * Returns true if the given object is an instance of QuotaRateLimit.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is QuotaRateLimit {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === QuotaRateLimit.__pulumiType;
    }

    /**
     * Name of the rate limit quota
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Path of the mount or namespace to apply the quota. A blank path configures a
     * global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
     * `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
     * Updating this field on an existing quota can have "moving" effects. For example, updating
     * `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
     * a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
     */
    public readonly path!: pulumi.Output<string | undefined>;
    /**
     * The maximum number of requests at any given second to be allowed by the quota
     * rule. The `rate` must be positive.
     */
    public readonly rate!: pulumi.Output<number>;

    /**
     * Create a QuotaRateLimit resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: QuotaRateLimitArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: QuotaRateLimitArgs | QuotaRateLimitState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as QuotaRateLimitState | undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["rate"] = state ? state.rate : undefined;
        } else {
            const args = argsOrState as QuotaRateLimitArgs | undefined;
            if ((!args || args.rate === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rate'");
            }
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["rate"] = args ? args.rate : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(QuotaRateLimit.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering QuotaRateLimit resources.
 */
export interface QuotaRateLimitState {
    /**
     * Name of the rate limit quota
     */
    name?: pulumi.Input<string>;
    /**
     * Path of the mount or namespace to apply the quota. A blank path configures a
     * global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
     * `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
     * Updating this field on an existing quota can have "moving" effects. For example, updating
     * `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
     * a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
     */
    path?: pulumi.Input<string>;
    /**
     * The maximum number of requests at any given second to be allowed by the quota
     * rule. The `rate` must be positive.
     */
    rate?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a QuotaRateLimit resource.
 */
export interface QuotaRateLimitArgs {
    /**
     * Name of the rate limit quota
     */
    name?: pulumi.Input<string>;
    /**
     * Path of the mount or namespace to apply the quota. A blank path configures a
     * global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
     * `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
     * Updating this field on an existing quota can have "moving" effects. For example, updating
     * `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
     * a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
     */
    path?: pulumi.Input<string>;
    /**
     * The maximum number of requests at any given second to be allowed by the quota
     * rule. The `rate` must be positive.
     */
    rate: pulumi.Input<number>;
}
