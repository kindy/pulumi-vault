// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * Provides a resource to manage [Namespaces](https://www.vaultproject.io/docs/enterprise/namespaces/index.html).
 *
 * **Note** this feature is available only with Vault Enterprise.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const ns1 = new vault.Namespace("ns1", {
 *     path: "ns1",
 * });
 * ```
 */
export class Namespace extends pulumi.CustomResource {
    /**
     * Get an existing Namespace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NamespaceState, opts?: pulumi.CustomResourceOptions): Namespace {
        return new Namespace(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:index/namespace:Namespace';

    /**
     * Returns true if the given object is an instance of Namespace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Namespace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Namespace.__pulumiType;
    }

    /**
     * ID of the namepsace.
     */
    public /*out*/ readonly namespaceId!: pulumi.Output<string>;
    /**
     * The path of the namespace. Must not have a trailing `/`
     */
    public readonly path!: pulumi.Output<string>;

    /**
     * Create a Namespace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NamespaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NamespaceArgs | NamespaceState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as NamespaceState | undefined;
            inputs["namespaceId"] = state ? state.namespaceId : undefined;
            inputs["path"] = state ? state.path : undefined;
        } else {
            const args = argsOrState as NamespaceArgs | undefined;
            if (!args || args.path === undefined) {
                throw new Error("Missing required property 'path'");
            }
            inputs["path"] = args ? args.path : undefined;
            inputs["namespaceId"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Namespace.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Namespace resources.
 */
export interface NamespaceState {
    /**
     * ID of the namepsace.
     */
    readonly namespaceId?: pulumi.Input<string>;
    /**
     * The path of the namespace. Must not have a trailing `/`
     */
    readonly path?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Namespace resource.
 */
export interface NamespaceArgs {
    /**
     * The path of the namespace. Must not have a trailing `/`
     */
    readonly path: pulumi.Input<string>;
}
