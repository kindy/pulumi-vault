// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Writes a KV-V1 secret to a given path in Vault.
 *
 * For more information on Vault's KV-V1 secret backend
 * [see here](https://www.vaultproject.io/docs/secrets/kv/kv-v1).
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const kvv1 = new vault.Mount("kvv1", {
 *     path: "kvv1",
 *     type: "kv",
 *     options: {
 *         version: "1",
 *     },
 *     description: "KV Version 1 secret engine mount",
 * });
 * const secret = new vault.kv.Secret("secret", {
 *     path: pulumi.interpolate`${kvv1.path}/secret`,
 *     dataJson: JSON.stringify({
 *         zip: "zap",
 *         foo: "bar",
 *     }),
 * });
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Required Vault Capabilities
 *
 * Use of this resource requires the `create` or `update` capability
 * (depending on whether the resource already exists) on the given path,
 * the `delete` capability if the resource is removed from configuration,
 * and the `read` capability for drift detection (by default).
 *
 * ## Import
 *
 * KV-V1 secrets can be imported using the `path`, e.g.
 *
 * ```sh
 * $ pulumi import vault:kv/secret:Secret secret kvv1/secret
 * ```
 */
export class Secret extends pulumi.CustomResource {
    /**
     * Get an existing Secret resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretState, opts?: pulumi.CustomResourceOptions): Secret {
        return new Secret(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:kv/secret:Secret';

    /**
     * Returns true if the given object is an instance of Secret.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Secret {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Secret.__pulumiType;
    }

    /**
     * A mapping whose keys are the top-level data keys returned from
     * Vault and whose values are the corresponding values. This map can only
     * represent string data, so any non-string values returned from Vault are
     * serialized as JSON.
     */
    public /*out*/ readonly data!: pulumi.Output<{[key: string]: any}>;
    /**
     * JSON-encoded string that will be
     * written as the secret data at the given path.
     */
    public readonly dataJson!: pulumi.Output<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Full path of the KV-V1 secret.
     */
    public readonly path!: pulumi.Output<string>;

    /**
     * Create a Secret resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretArgs | SecretState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretState | undefined;
            resourceInputs["data"] = state ? state.data : undefined;
            resourceInputs["dataJson"] = state ? state.dataJson : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
        } else {
            const args = argsOrState as SecretArgs | undefined;
            if ((!args || args.dataJson === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dataJson'");
            }
            if ((!args || args.path === undefined) && !opts.urn) {
                throw new Error("Missing required property 'path'");
            }
            resourceInputs["dataJson"] = args?.dataJson ? pulumi.secret(args.dataJson) : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["data"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["data", "dataJson"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(Secret.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Secret resources.
 */
export interface SecretState {
    /**
     * A mapping whose keys are the top-level data keys returned from
     * Vault and whose values are the corresponding values. This map can only
     * represent string data, so any non-string values returned from Vault are
     * serialized as JSON.
     */
    data?: pulumi.Input<{[key: string]: any}>;
    /**
     * JSON-encoded string that will be
     * written as the secret data at the given path.
     */
    dataJson?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Full path of the KV-V1 secret.
     */
    path?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a Secret resource.
 */
export interface SecretArgs {
    /**
     * JSON-encoded string that will be
     * written as the secret data at the given path.
     */
    dataJson: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Full path of the KV-V1 secret.
     */
    path: pulumi.Input<string>;
}
