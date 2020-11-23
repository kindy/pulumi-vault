// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Consul secret backends can be imported using the `path`, e.g.
 *
 * ```sh
 *  $ pulumi import vault:consul/secretBackend:SecretBackend example consul
 * ```
 */
export class SecretBackend extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendState, opts?: pulumi.CustomResourceOptions): SecretBackend {
        return new SecretBackend(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:consul/secretBackend:SecretBackend';

    /**
     * Returns true if the given object is an instance of SecretBackend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackend.__pulumiType;
    }

    /**
     * Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
     */
    public readonly address!: pulumi.Output<string>;
    /**
     * The default TTL for credentials issued by this backend.
     */
    public readonly defaultLeaseTtlSeconds!: pulumi.Output<number | undefined>;
    /**
     * A human-friendly description for this backend.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The maximum TTL that can be requested
     * for credentials issued by this backend.
     */
    public readonly maxLeaseTtlSeconds!: pulumi.Output<number | undefined>;
    /**
     * The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults to `consul`.
     */
    public readonly path!: pulumi.Output<string | undefined>;
    /**
     * Specifies the URL scheme to use. Defaults to `http`.
     */
    public readonly scheme!: pulumi.Output<string | undefined>;
    /**
     * The Consul management token this backend should use to issue new tokens.
     */
    public readonly token!: pulumi.Output<string>;

    /**
     * Create a SecretBackend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretBackendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendArgs | SecretBackendState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecretBackendState | undefined;
            inputs["address"] = state ? state.address : undefined;
            inputs["defaultLeaseTtlSeconds"] = state ? state.defaultLeaseTtlSeconds : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["maxLeaseTtlSeconds"] = state ? state.maxLeaseTtlSeconds : undefined;
            inputs["path"] = state ? state.path : undefined;
            inputs["scheme"] = state ? state.scheme : undefined;
            inputs["token"] = state ? state.token : undefined;
        } else {
            const args = argsOrState as SecretBackendArgs | undefined;
            if (!args || args.address === undefined) {
                throw new Error("Missing required property 'address'");
            }
            if (!args || args.token === undefined) {
                throw new Error("Missing required property 'token'");
            }
            inputs["address"] = args ? args.address : undefined;
            inputs["defaultLeaseTtlSeconds"] = args ? args.defaultLeaseTtlSeconds : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["maxLeaseTtlSeconds"] = args ? args.maxLeaseTtlSeconds : undefined;
            inputs["path"] = args ? args.path : undefined;
            inputs["scheme"] = args ? args.scheme : undefined;
            inputs["token"] = args ? args.token : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SecretBackend.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackend resources.
 */
export interface SecretBackendState {
    /**
     * Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
     */
    readonly address?: pulumi.Input<string>;
    /**
     * The default TTL for credentials issued by this backend.
     */
    readonly defaultLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * A human-friendly description for this backend.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The maximum TTL that can be requested
     * for credentials issued by this backend.
     */
    readonly maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults to `consul`.
     */
    readonly path?: pulumi.Input<string>;
    /**
     * Specifies the URL scheme to use. Defaults to `http`.
     */
    readonly scheme?: pulumi.Input<string>;
    /**
     * The Consul management token this backend should use to issue new tokens.
     */
    readonly token?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretBackend resource.
 */
export interface SecretBackendArgs {
    /**
     * Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
     */
    readonly address: pulumi.Input<string>;
    /**
     * The default TTL for credentials issued by this backend.
     */
    readonly defaultLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * A human-friendly description for this backend.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The maximum TTL that can be requested
     * for credentials issued by this backend.
     */
    readonly maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults to `consul`.
     */
    readonly path?: pulumi.Input<string>;
    /**
     * Specifies the URL scheme to use. Defaults to `http`.
     */
    readonly scheme?: pulumi.Input<string>;
    /**
     * The Consul management token this backend should use to issue new tokens.
     */
    readonly token: pulumi.Input<string>;
}
