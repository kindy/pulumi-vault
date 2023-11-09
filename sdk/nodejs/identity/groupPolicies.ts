// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages policies for an Identity Group for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.
 *
 * ## Example Usage
 * ### Exclusive Policies
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const internal = new vault.identity.Group("internal", {
 *     type: "internal",
 *     externalPolicies: true,
 *     metadata: {
 *         version: "2",
 *     },
 * });
 * const policies = new vault.identity.GroupPolicies("policies", {
 *     policies: [
 *         "default",
 *         "test",
 *     ],
 *     exclusive: true,
 *     groupId: internal.id,
 * });
 * ```
 * ### Non-exclusive Policies
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const internal = new vault.identity.Group("internal", {
 *     type: "internal",
 *     externalPolicies: true,
 *     metadata: {
 *         version: "2",
 *     },
 * });
 * const _default = new vault.identity.GroupPolicies("default", {
 *     policies: [
 *         "default",
 *         "test",
 *     ],
 *     exclusive: false,
 *     groupId: internal.id,
 * });
 * const others = new vault.identity.GroupPolicies("others", {
 *     policies: ["others"],
 *     exclusive: false,
 *     groupId: internal.id,
 * });
 * ```
 */
export class GroupPolicies extends pulumi.CustomResource {
    /**
     * Get an existing GroupPolicies resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupPoliciesState, opts?: pulumi.CustomResourceOptions): GroupPolicies {
        return new GroupPolicies(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:identity/groupPolicies:GroupPolicies';

    /**
     * Returns true if the given object is an instance of GroupPolicies.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupPolicies {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupPolicies.__pulumiType;
    }

    /**
     * Defaults to `true`.
     *
     * If `true`, this resource will take exclusive control of the policies assigned to the group and will set it equal to what is specified in the resource.
     *
     * If set to `false`, this resource will simply ensure that the policies specified in the resource are present in the group. When destroying the resource, the resource will ensure that the policies specified in the resource are removed.
     */
    public readonly exclusive!: pulumi.Output<boolean | undefined>;
    /**
     * Group ID to assign policies to.
     */
    public readonly groupId!: pulumi.Output<string>;
    /**
     * The name of the group that are assigned the policies.
     */
    public /*out*/ readonly groupName!: pulumi.Output<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * List of policies to assign to the group
     */
    public readonly policies!: pulumi.Output<string[]>;

    /**
     * Create a GroupPolicies resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupPoliciesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupPoliciesArgs | GroupPoliciesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupPoliciesState | undefined;
            resourceInputs["exclusive"] = state ? state.exclusive : undefined;
            resourceInputs["groupId"] = state ? state.groupId : undefined;
            resourceInputs["groupName"] = state ? state.groupName : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["policies"] = state ? state.policies : undefined;
        } else {
            const args = argsOrState as GroupPoliciesArgs | undefined;
            if ((!args || args.groupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupId'");
            }
            if ((!args || args.policies === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policies'");
            }
            resourceInputs["exclusive"] = args ? args.exclusive : undefined;
            resourceInputs["groupId"] = args ? args.groupId : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["policies"] = args ? args.policies : undefined;
            resourceInputs["groupName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GroupPolicies.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupPolicies resources.
 */
export interface GroupPoliciesState {
    /**
     * Defaults to `true`.
     *
     * If `true`, this resource will take exclusive control of the policies assigned to the group and will set it equal to what is specified in the resource.
     *
     * If set to `false`, this resource will simply ensure that the policies specified in the resource are present in the group. When destroying the resource, the resource will ensure that the policies specified in the resource are removed.
     */
    exclusive?: pulumi.Input<boolean>;
    /**
     * Group ID to assign policies to.
     */
    groupId?: pulumi.Input<string>;
    /**
     * The name of the group that are assigned the policies.
     */
    groupName?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * List of policies to assign to the group
     */
    policies?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a GroupPolicies resource.
 */
export interface GroupPoliciesArgs {
    /**
     * Defaults to `true`.
     *
     * If `true`, this resource will take exclusive control of the policies assigned to the group and will set it equal to what is specified in the resource.
     *
     * If set to `false`, this resource will simply ensure that the policies specified in the resource are present in the group. When destroying the resource, the resource will ensure that the policies specified in the resource are removed.
     */
    exclusive?: pulumi.Input<boolean>;
    /**
     * Group ID to assign policies to.
     */
    groupId: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * List of policies to assign to the group
     */
    policies: pulumi.Input<pulumi.Input<string>[]>;
}
