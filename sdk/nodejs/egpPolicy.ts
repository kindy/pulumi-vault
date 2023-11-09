// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a resource to manage Endpoint Governing Policy (EGP) via [Sentinel](https://www.vaultproject.io/docs/enterprise/sentinel/index.html).
 *
 * **Note** this feature is available only with Vault Enterprise.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const allow_all = new vault.EgpPolicy("allow-all", {
 *     enforcementLevel: "soft-mandatory",
 *     paths: ["*"],
 *     policy: `main = rule {
 *   true
 * }
 *
 * `,
 * });
 * ```
 */
export class EgpPolicy extends pulumi.CustomResource {
    /**
     * Get an existing EgpPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EgpPolicyState, opts?: pulumi.CustomResourceOptions): EgpPolicy {
        return new EgpPolicy(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:index/egpPolicy:EgpPolicy';

    /**
     * Returns true if the given object is an instance of EgpPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EgpPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EgpPolicy.__pulumiType;
    }

    /**
     * Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
     */
    public readonly enforcementLevel!: pulumi.Output<string>;
    /**
     * The name of the policy
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * List of paths to which the policy will be applied to
     */
    public readonly paths!: pulumi.Output<string[]>;
    /**
     * String containing a Sentinel policy
     */
    public readonly policy!: pulumi.Output<string>;

    /**
     * Create a EgpPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EgpPolicyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EgpPolicyArgs | EgpPolicyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EgpPolicyState | undefined;
            resourceInputs["enforcementLevel"] = state ? state.enforcementLevel : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["paths"] = state ? state.paths : undefined;
            resourceInputs["policy"] = state ? state.policy : undefined;
        } else {
            const args = argsOrState as EgpPolicyArgs | undefined;
            if ((!args || args.enforcementLevel === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enforcementLevel'");
            }
            if ((!args || args.paths === undefined) && !opts.urn) {
                throw new Error("Missing required property 'paths'");
            }
            if ((!args || args.policy === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policy'");
            }
            resourceInputs["enforcementLevel"] = args ? args.enforcementLevel : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["paths"] = args ? args.paths : undefined;
            resourceInputs["policy"] = args ? args.policy : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EgpPolicy.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EgpPolicy resources.
 */
export interface EgpPolicyState {
    /**
     * Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
     */
    enforcementLevel?: pulumi.Input<string>;
    /**
     * The name of the policy
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * List of paths to which the policy will be applied to
     */
    paths?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * String containing a Sentinel policy
     */
    policy?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EgpPolicy resource.
 */
export interface EgpPolicyArgs {
    /**
     * Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
     */
    enforcementLevel: pulumi.Input<string>;
    /**
     * The name of the policy
     */
    name?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * List of paths to which the policy will be applied to
     */
    paths: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * String containing a Sentinel policy
     */
    policy: pulumi.Input<string>;
}
