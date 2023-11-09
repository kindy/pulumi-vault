// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages policies for an Identity Entity for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.
 *
 * ## Example Usage
 * ### Exclusive Policies
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const entity = new vault.identity.Entity("entity", {externalPolicies: true});
 * const policies = new vault.identity.EntityPolicies("policies", {
 *     policies: [
 *         "default",
 *         "test",
 *     ],
 *     exclusive: true,
 *     entityId: entity.id,
 * });
 * ```
 * ### Non-exclusive Policies
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const entity = new vault.identity.Entity("entity", {externalPolicies: true});
 * const _default = new vault.identity.EntityPolicies("default", {
 *     policies: [
 *         "default",
 *         "test",
 *     ],
 *     exclusive: false,
 *     entityId: entity.id,
 * });
 * const others = new vault.identity.EntityPolicies("others", {
 *     policies: ["others"],
 *     exclusive: false,
 *     entityId: entity.id,
 * });
 * ```
 */
export class EntityPolicies extends pulumi.CustomResource {
    /**
     * Get an existing EntityPolicies resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EntityPoliciesState, opts?: pulumi.CustomResourceOptions): EntityPolicies {
        return new EntityPolicies(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:identity/entityPolicies:EntityPolicies';

    /**
     * Returns true if the given object is an instance of EntityPolicies.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EntityPolicies {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EntityPolicies.__pulumiType;
    }

    /**
     * Entity ID to assign policies to.
     */
    public readonly entityId!: pulumi.Output<string>;
    /**
     * The name of the entity that are assigned the policies.
     */
    public /*out*/ readonly entityName!: pulumi.Output<string>;
    /**
     * Defaults to `true`.
     *
     * If `true`, this resource will take exclusive control of the policies assigned to the entity and will set it equal to what is specified in the resource.
     *
     * If set to `false`, this resource will simply ensure that the policies specified in the resource are present in the entity. When destroying the resource, the resource will ensure that the policies specified in the resource are removed.
     */
    public readonly exclusive!: pulumi.Output<boolean | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * List of policies to assign to the entity
     */
    public readonly policies!: pulumi.Output<string[]>;

    /**
     * Create a EntityPolicies resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EntityPoliciesArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EntityPoliciesArgs | EntityPoliciesState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EntityPoliciesState | undefined;
            resourceInputs["entityId"] = state ? state.entityId : undefined;
            resourceInputs["entityName"] = state ? state.entityName : undefined;
            resourceInputs["exclusive"] = state ? state.exclusive : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["policies"] = state ? state.policies : undefined;
        } else {
            const args = argsOrState as EntityPoliciesArgs | undefined;
            if ((!args || args.entityId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'entityId'");
            }
            if ((!args || args.policies === undefined) && !opts.urn) {
                throw new Error("Missing required property 'policies'");
            }
            resourceInputs["entityId"] = args ? args.entityId : undefined;
            resourceInputs["exclusive"] = args ? args.exclusive : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["policies"] = args ? args.policies : undefined;
            resourceInputs["entityName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EntityPolicies.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EntityPolicies resources.
 */
export interface EntityPoliciesState {
    /**
     * Entity ID to assign policies to.
     */
    entityId?: pulumi.Input<string>;
    /**
     * The name of the entity that are assigned the policies.
     */
    entityName?: pulumi.Input<string>;
    /**
     * Defaults to `true`.
     *
     * If `true`, this resource will take exclusive control of the policies assigned to the entity and will set it equal to what is specified in the resource.
     *
     * If set to `false`, this resource will simply ensure that the policies specified in the resource are present in the entity. When destroying the resource, the resource will ensure that the policies specified in the resource are removed.
     */
    exclusive?: pulumi.Input<boolean>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * List of policies to assign to the entity
     */
    policies?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a EntityPolicies resource.
 */
export interface EntityPoliciesArgs {
    /**
     * Entity ID to assign policies to.
     */
    entityId: pulumi.Input<string>;
    /**
     * Defaults to `true`.
     *
     * If `true`, this resource will take exclusive control of the policies assigned to the entity and will set it equal to what is specified in the resource.
     *
     * If set to `false`, this resource will simply ensure that the policies specified in the resource are present in the entity. When destroying the resource, the resource will ensure that the policies specified in the resource are removed.
     */
    exclusive?: pulumi.Input<boolean>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * List of policies to assign to the entity
     */
    policies: pulumi.Input<pulumi.Input<string>[]>;
}
