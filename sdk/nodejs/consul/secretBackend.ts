// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const test = new vault.consul.SecretBackend("test", {
 *     address: "127.0.0.1:8500",
 *     description: "Manages the Consul backend",
 *     path: "consul",
 *     token: "4240861b-ce3d-8530-115a-521ff070dd29",
 * });
 * ```
 *
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
     * CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
     */
    public readonly caCert!: pulumi.Output<string | undefined>;
    /**
     * Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_key.
     */
    public readonly clientCert!: pulumi.Output<string | undefined>;
    /**
     * Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_cert.
     */
    public readonly clientKey!: pulumi.Output<string | undefined>;
    /**
     * The default TTL for credentials issued by this backend.
     */
    public readonly defaultLeaseTtlSeconds!: pulumi.Output<number | undefined>;
    /**
     * A human-friendly description for this backend.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Specifies if the secret backend is local only.
     */
    public readonly local!: pulumi.Output<boolean | undefined>;
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretBackendState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["caCert"] = state ? state.caCert : undefined;
            resourceInputs["clientCert"] = state ? state.clientCert : undefined;
            resourceInputs["clientKey"] = state ? state.clientKey : undefined;
            resourceInputs["defaultLeaseTtlSeconds"] = state ? state.defaultLeaseTtlSeconds : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["local"] = state ? state.local : undefined;
            resourceInputs["maxLeaseTtlSeconds"] = state ? state.maxLeaseTtlSeconds : undefined;
            resourceInputs["path"] = state ? state.path : undefined;
            resourceInputs["scheme"] = state ? state.scheme : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
        } else {
            const args = argsOrState as SecretBackendArgs | undefined;
            if ((!args || args.address === undefined) && !opts.urn) {
                throw new Error("Missing required property 'address'");
            }
            if ((!args || args.token === undefined) && !opts.urn) {
                throw new Error("Missing required property 'token'");
            }
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["caCert"] = args ? args.caCert : undefined;
            resourceInputs["clientCert"] = args ? args.clientCert : undefined;
            resourceInputs["clientKey"] = args ? args.clientKey : undefined;
            resourceInputs["defaultLeaseTtlSeconds"] = args ? args.defaultLeaseTtlSeconds : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["local"] = args ? args.local : undefined;
            resourceInputs["maxLeaseTtlSeconds"] = args ? args.maxLeaseTtlSeconds : undefined;
            resourceInputs["path"] = args ? args.path : undefined;
            resourceInputs["scheme"] = args ? args.scheme : undefined;
            resourceInputs["token"] = args ? args.token : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecretBackend.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackend resources.
 */
export interface SecretBackendState {
    /**
     * Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
     */
    address?: pulumi.Input<string>;
    /**
     * CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
     */
    caCert?: pulumi.Input<string>;
    /**
     * Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_key.
     */
    clientCert?: pulumi.Input<string>;
    /**
     * Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_cert.
     */
    clientKey?: pulumi.Input<string>;
    /**
     * The default TTL for credentials issued by this backend.
     */
    defaultLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * A human-friendly description for this backend.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies if the secret backend is local only.
     */
    local?: pulumi.Input<boolean>;
    /**
     * The maximum TTL that can be requested
     * for credentials issued by this backend.
     */
    maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults to `consul`.
     */
    path?: pulumi.Input<string>;
    /**
     * Specifies the URL scheme to use. Defaults to `http`.
     */
    scheme?: pulumi.Input<string>;
    /**
     * The Consul management token this backend should use to issue new tokens.
     */
    token?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretBackend resource.
 */
export interface SecretBackendArgs {
    /**
     * Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
     */
    address: pulumi.Input<string>;
    /**
     * CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
     */
    caCert?: pulumi.Input<string>;
    /**
     * Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_key.
     */
    clientCert?: pulumi.Input<string>;
    /**
     * Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set you need to also set client_cert.
     */
    clientKey?: pulumi.Input<string>;
    /**
     * The default TTL for credentials issued by this backend.
     */
    defaultLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * A human-friendly description for this backend.
     */
    description?: pulumi.Input<string>;
    /**
     * Specifies if the secret backend is local only.
     */
    local?: pulumi.Input<boolean>;
    /**
     * The maximum TTL that can be requested
     * for credentials issued by this backend.
     */
    maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults to `consul`.
     */
    path?: pulumi.Input<string>;
    /**
     * Specifies the URL scheme to use. Defaults to `http`.
     */
    scheme?: pulumi.Input<string>;
    /**
     * The Consul management token this backend should use to issue new tokens.
     */
    token: pulumi.Input<string>;
}
