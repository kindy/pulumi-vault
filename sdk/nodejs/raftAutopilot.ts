// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Autopilot enables automated workflows for managing Raft clusters. The
 * current feature set includes 3 main features: Server Stabilization, Dead
 * Server Cleanup and State API. **These three features are introduced in
 * Vault 1.7.**
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const autopilot = new vault.RaftAutopilot("autopilot", {
 *     cleanupDeadServers: true,
 *     deadServerLastContactThreshold: "24h0m0s",
 *     lastContactThreshold: "10s",
 *     maxTrailingLogs: 1000,
 *     minQuorum: 3,
 *     serverStabilizationTime: "10s",
 * });
 * ```
 *
 * ## Import
 *
 * Raft Autopilot config can be imported using the ID, e.g.
 *
 * ```sh
 *  $ pulumi import vault:index/raftAutopilot:RaftAutopilot autopilot sys/storage/raft/autopilot/configuration
 * ```
 */
export class RaftAutopilot extends pulumi.CustomResource {
    /**
     * Get an existing RaftAutopilot resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RaftAutopilotState, opts?: pulumi.CustomResourceOptions): RaftAutopilot {
        return new RaftAutopilot(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:index/raftAutopilot:RaftAutopilot';

    /**
     * Returns true if the given object is an instance of RaftAutopilot.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RaftAutopilot {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RaftAutopilot.__pulumiType;
    }

    /**
     * Specifies whether to remove dead server nodes
     * periodically or when a new server joins. This requires that `min-quorum` is also set.
     */
    public readonly cleanupDeadServers!: pulumi.Output<boolean | undefined>;
    /**
     * Limit the amount of time a 
     * server can go without leader contact before being considered failed. This only takes
     * effect when `cleanupDeadServers` is set.
     */
    public readonly deadServerLastContactThreshold!: pulumi.Output<string | undefined>;
    /**
     * Disables automatically upgrading Vault using autopilot. (Enterprise-only)
     */
    public readonly disableUpgradeMigration!: pulumi.Output<boolean | undefined>;
    /**
     * Limit the amount of time a server can go 
     * without leader contact before being considered unhealthy.
     */
    public readonly lastContactThreshold!: pulumi.Output<string | undefined>;
    /**
     * Maximum number of log entries in the Raft log 
     * that a server can be behind its leader before being considered unhealthy.
     */
    public readonly maxTrailingLogs!: pulumi.Output<number | undefined>;
    /**
     * Minimum number of servers allowed in a cluster before 
     * autopilot can prune dead servers. This should at least be 3. Applicable only for
     * voting nodes.
     */
    public readonly minQuorum!: pulumi.Output<number | undefined>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Minimum amount of time a server must be 
     * stable in the 'healthy' state before being added to the cluster.
     */
    public readonly serverStabilizationTime!: pulumi.Output<string | undefined>;

    /**
     * Create a RaftAutopilot resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: RaftAutopilotArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RaftAutopilotArgs | RaftAutopilotState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RaftAutopilotState | undefined;
            resourceInputs["cleanupDeadServers"] = state ? state.cleanupDeadServers : undefined;
            resourceInputs["deadServerLastContactThreshold"] = state ? state.deadServerLastContactThreshold : undefined;
            resourceInputs["disableUpgradeMigration"] = state ? state.disableUpgradeMigration : undefined;
            resourceInputs["lastContactThreshold"] = state ? state.lastContactThreshold : undefined;
            resourceInputs["maxTrailingLogs"] = state ? state.maxTrailingLogs : undefined;
            resourceInputs["minQuorum"] = state ? state.minQuorum : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["serverStabilizationTime"] = state ? state.serverStabilizationTime : undefined;
        } else {
            const args = argsOrState as RaftAutopilotArgs | undefined;
            resourceInputs["cleanupDeadServers"] = args ? args.cleanupDeadServers : undefined;
            resourceInputs["deadServerLastContactThreshold"] = args ? args.deadServerLastContactThreshold : undefined;
            resourceInputs["disableUpgradeMigration"] = args ? args.disableUpgradeMigration : undefined;
            resourceInputs["lastContactThreshold"] = args ? args.lastContactThreshold : undefined;
            resourceInputs["maxTrailingLogs"] = args ? args.maxTrailingLogs : undefined;
            resourceInputs["minQuorum"] = args ? args.minQuorum : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["serverStabilizationTime"] = args ? args.serverStabilizationTime : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RaftAutopilot.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RaftAutopilot resources.
 */
export interface RaftAutopilotState {
    /**
     * Specifies whether to remove dead server nodes
     * periodically or when a new server joins. This requires that `min-quorum` is also set.
     */
    cleanupDeadServers?: pulumi.Input<boolean>;
    /**
     * Limit the amount of time a 
     * server can go without leader contact before being considered failed. This only takes
     * effect when `cleanupDeadServers` is set.
     */
    deadServerLastContactThreshold?: pulumi.Input<string>;
    /**
     * Disables automatically upgrading Vault using autopilot. (Enterprise-only)
     */
    disableUpgradeMigration?: pulumi.Input<boolean>;
    /**
     * Limit the amount of time a server can go 
     * without leader contact before being considered unhealthy.
     */
    lastContactThreshold?: pulumi.Input<string>;
    /**
     * Maximum number of log entries in the Raft log 
     * that a server can be behind its leader before being considered unhealthy.
     */
    maxTrailingLogs?: pulumi.Input<number>;
    /**
     * Minimum number of servers allowed in a cluster before 
     * autopilot can prune dead servers. This should at least be 3. Applicable only for
     * voting nodes.
     */
    minQuorum?: pulumi.Input<number>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Minimum amount of time a server must be 
     * stable in the 'healthy' state before being added to the cluster.
     */
    serverStabilizationTime?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RaftAutopilot resource.
 */
export interface RaftAutopilotArgs {
    /**
     * Specifies whether to remove dead server nodes
     * periodically or when a new server joins. This requires that `min-quorum` is also set.
     */
    cleanupDeadServers?: pulumi.Input<boolean>;
    /**
     * Limit the amount of time a 
     * server can go without leader contact before being considered failed. This only takes
     * effect when `cleanupDeadServers` is set.
     */
    deadServerLastContactThreshold?: pulumi.Input<string>;
    /**
     * Disables automatically upgrading Vault using autopilot. (Enterprise-only)
     */
    disableUpgradeMigration?: pulumi.Input<boolean>;
    /**
     * Limit the amount of time a server can go 
     * without leader contact before being considered unhealthy.
     */
    lastContactThreshold?: pulumi.Input<string>;
    /**
     * Maximum number of log entries in the Raft log 
     * that a server can be behind its leader before being considered unhealthy.
     */
    maxTrailingLogs?: pulumi.Input<number>;
    /**
     * Minimum number of servers allowed in a cluster before 
     * autopilot can prune dead servers. This should at least be 3. Applicable only for
     * voting nodes.
     */
    minQuorum?: pulumi.Input<number>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Minimum amount of time a server must be 
     * stable in the 'healthy' state before being added to the cluster.
     */
    serverStabilizationTime?: pulumi.Input<string>;
}
