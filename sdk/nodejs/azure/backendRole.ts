// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export class BackendRole extends pulumi.CustomResource {
    /**
     * Get an existing BackendRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: BackendRoleState, opts?: pulumi.CustomResourceOptions): BackendRole {
        return new BackendRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:azure/backendRole:BackendRole';

    /**
     * Returns true if the given object is an instance of BackendRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is BackendRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === BackendRole.__pulumiType;
    }

    /**
     * Application Object ID for an existing service principal that will
     * be used instead of creating dynamic service principals. If present, `azureRoles` will be ignored.
     */
    public readonly applicationObjectId!: pulumi.Output<string | undefined>;
    /**
     * List of Azure groups to be assigned to the generated service principal.
     */
    public readonly azureGroups!: pulumi.Output<outputs.azure.BackendRoleAzureGroup[] | undefined>;
    /**
     * List of Azure roles to be assigned to the generated service principal.
     */
    public readonly azureRoles!: pulumi.Output<outputs.azure.BackendRoleAzureRole[] | undefined>;
    /**
     * Path to the mounted Azure auth backend
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * Human-friendly description of the mount for the backend.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies the maximum TTL for service principals generated using this role. Accepts time
     * suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine max TTL time.
     */
    public readonly maxTtl!: pulumi.Output<string | undefined>;
    /**
     * Name of the Azure role
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * Specifies the default TTL for service principals generated using this role.
     * Accepts time suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine default TTL time.
     */
    public readonly ttl!: pulumi.Output<string | undefined>;

    /**
     * Create a BackendRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: BackendRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: BackendRoleArgs | BackendRoleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as BackendRoleState | undefined;
            inputs["applicationObjectId"] = state ? state.applicationObjectId : undefined;
            inputs["azureGroups"] = state ? state.azureGroups : undefined;
            inputs["azureRoles"] = state ? state.azureRoles : undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["maxTtl"] = state ? state.maxTtl : undefined;
            inputs["role"] = state ? state.role : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
        } else {
            const args = argsOrState as BackendRoleArgs | undefined;
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            inputs["applicationObjectId"] = args ? args.applicationObjectId : undefined;
            inputs["azureGroups"] = args ? args.azureGroups : undefined;
            inputs["azureRoles"] = args ? args.azureRoles : undefined;
            inputs["backend"] = args ? args.backend : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["maxTtl"] = args ? args.maxTtl : undefined;
            inputs["role"] = args ? args.role : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(BackendRole.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering BackendRole resources.
 */
export interface BackendRoleState {
    /**
     * Application Object ID for an existing service principal that will
     * be used instead of creating dynamic service principals. If present, `azureRoles` will be ignored.
     */
    readonly applicationObjectId?: pulumi.Input<string>;
    /**
     * List of Azure groups to be assigned to the generated service principal.
     */
    readonly azureGroups?: pulumi.Input<pulumi.Input<inputs.azure.BackendRoleAzureGroup>[]>;
    /**
     * List of Azure roles to be assigned to the generated service principal.
     */
    readonly azureRoles?: pulumi.Input<pulumi.Input<inputs.azure.BackendRoleAzureRole>[]>;
    /**
     * Path to the mounted Azure auth backend
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * Human-friendly description of the mount for the backend.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies the maximum TTL for service principals generated using this role. Accepts time
     * suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine max TTL time.
     */
    readonly maxTtl?: pulumi.Input<string>;
    /**
     * Name of the Azure role
     */
    readonly role?: pulumi.Input<string>;
    /**
     * Specifies the default TTL for service principals generated using this role.
     * Accepts time suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine default TTL time.
     */
    readonly ttl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a BackendRole resource.
 */
export interface BackendRoleArgs {
    /**
     * Application Object ID for an existing service principal that will
     * be used instead of creating dynamic service principals. If present, `azureRoles` will be ignored.
     */
    readonly applicationObjectId?: pulumi.Input<string>;
    /**
     * List of Azure groups to be assigned to the generated service principal.
     */
    readonly azureGroups?: pulumi.Input<pulumi.Input<inputs.azure.BackendRoleAzureGroup>[]>;
    /**
     * List of Azure roles to be assigned to the generated service principal.
     */
    readonly azureRoles?: pulumi.Input<pulumi.Input<inputs.azure.BackendRoleAzureRole>[]>;
    /**
     * Path to the mounted Azure auth backend
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * Human-friendly description of the mount for the backend.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Specifies the maximum TTL for service principals generated using this role. Accepts time
     * suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine max TTL time.
     */
    readonly maxTtl?: pulumi.Input<string>;
    /**
     * Name of the Azure role
     */
    readonly role: pulumi.Input<string>;
    /**
     * Specifies the default TTL for service principals generated using this role.
     * Accepts time suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine default TTL time.
     */
    readonly ttl?: pulumi.Input<string>;
}
