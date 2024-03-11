// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Configure the cache for the Transit Secret Backend in Vault.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const transit = new vault.Mount("transit", {
 *     path: "transit",
 *     type: "transit",
 *     description: "Example description",
 *     defaultLeaseTtlSeconds: 3600,
 *     maxLeaseTtlSeconds: 86400,
 * });
 * const cfg = new vault.transit.SecretCacheConfig("cfg", {
 *     backend: transit.path,
 *     size: 500,
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 */
export class SecretCacheConfig extends pulumi.CustomResource {
    /**
     * Get an existing SecretCacheConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretCacheConfigState, opts?: pulumi.CustomResourceOptions): SecretCacheConfig {
        return new SecretCacheConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:transit/secretCacheConfig:SecretCacheConfig';

    /**
     * Returns true if the given object is an instance of SecretCacheConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretCacheConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretCacheConfig.__pulumiType;
    }

    /**
     * The path the transit secret backend is mounted at, with no leading or trailing `/`s.
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * The number of cache entries. 0 means unlimited.
     */
    public readonly size!: pulumi.Output<number>;

    /**
     * Create a SecretCacheConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretCacheConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretCacheConfigArgs | SecretCacheConfigState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretCacheConfigState | undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
        } else {
            const args = argsOrState as SecretCacheConfigArgs | undefined;
            if ((!args || args.backend === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backend'");
            }
            if ((!args || args.size === undefined) && !opts.urn) {
                throw new Error("Missing required property 'size'");
            }
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["size"] = args ? args.size : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecretCacheConfig.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretCacheConfig resources.
 */
export interface SecretCacheConfigState {
    /**
     * The path the transit secret backend is mounted at, with no leading or trailing `/`s.
     */
    backend?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The number of cache entries. 0 means unlimited.
     */
    size?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a SecretCacheConfig resource.
 */
export interface SecretCacheConfigArgs {
    /**
     * The path the transit secret backend is mounted at, with no leading or trailing `/`s.
     */
    backend: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * The number of cache entries. 0 means unlimited.
     */
    size: pulumi.Input<number>;
}
