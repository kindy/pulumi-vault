// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Import
 *
 * Nomad secret backend can be imported using the `backend`, e.g.
 *
 * ```sh
 *  $ pulumi import vault:index/nomadSecretBackend:NomadSecretBackend nomad nomad
 * ```
 */
export class NomadSecretBackend extends pulumi.CustomResource {
    /**
     * Get an existing NomadSecretBackend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NomadSecretBackendState, opts?: pulumi.CustomResourceOptions): NomadSecretBackend {
        return new NomadSecretBackend(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:index/nomadSecretBackend:NomadSecretBackend';

    /**
     * Returns true if the given object is an instance of NomadSecretBackend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NomadSecretBackend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NomadSecretBackend.__pulumiType;
    }

    /**
     * Specifies the address of the Nomad instance, provided
     * as "protocol://host:port" like "http://127.0.0.1:4646".
     */
    public readonly address!: pulumi.Output<string | undefined>;
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `nomad`.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * CA certificate to use when verifying the Nomad server certificate, must be
     * x509 PEM encoded.
     */
    public readonly caCert!: pulumi.Output<string | undefined>;
    /**
     * Client certificate to provide to the Nomad server, must be x509 PEM encoded.
     */
    public readonly clientCert!: pulumi.Output<string | undefined>;
    /**
     * Client certificate key to provide to the Nomad server, must be x509 PEM encoded.
     */
    public readonly clientKey!: pulumi.Output<string | undefined>;
    /**
     * Default lease duration for secrets in seconds.
     */
    public readonly defaultLeaseTtlSeconds!: pulumi.Output<number>;
    /**
     * Human-friendly description of the mount for the Active Directory backend.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Mark the secrets engine as local-only. Local engines are not replicated or removed by
     * replication.Tolerance duration to use when checking the last rotation time.
     */
    public readonly local!: pulumi.Output<boolean | undefined>;
    /**
     * Maximum possible lease duration for secrets in seconds.
     */
    public readonly maxLeaseTtlSeconds!: pulumi.Output<number>;
    /**
     * Specifies the maximum length to use for the name of the Nomad token
     * generated with Generate Credential. If omitted, 0 is used and ignored, defaulting to the max value allowed
     * by the Nomad version.
     */
    public readonly maxTokenNameLength!: pulumi.Output<number>;
    /**
     * Maximum possible lease duration for secrets in seconds.
     */
    public readonly maxTtl!: pulumi.Output<number>;
    /**
     * Specifies the Nomad Management token to use.
     */
    public readonly token!: pulumi.Output<string | undefined>;
    /**
     * Specifies the ttl of the lease for the generated token.
     */
    public readonly ttl!: pulumi.Output<number>;

    /**
     * Create a NomadSecretBackend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NomadSecretBackendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NomadSecretBackendArgs | NomadSecretBackendState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NomadSecretBackendState | undefined;
            resourceInputs["address"] = state ? state.address : undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["caCert"] = state ? state.caCert : undefined;
            resourceInputs["clientCert"] = state ? state.clientCert : undefined;
            resourceInputs["clientKey"] = state ? state.clientKey : undefined;
            resourceInputs["defaultLeaseTtlSeconds"] = state ? state.defaultLeaseTtlSeconds : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["local"] = state ? state.local : undefined;
            resourceInputs["maxLeaseTtlSeconds"] = state ? state.maxLeaseTtlSeconds : undefined;
            resourceInputs["maxTokenNameLength"] = state ? state.maxTokenNameLength : undefined;
            resourceInputs["maxTtl"] = state ? state.maxTtl : undefined;
            resourceInputs["token"] = state ? state.token : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
        } else {
            const args = argsOrState as NomadSecretBackendArgs | undefined;
            resourceInputs["address"] = args ? args.address : undefined;
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["caCert"] = args ? args.caCert : undefined;
            resourceInputs["clientCert"] = args ? args.clientCert : undefined;
            resourceInputs["clientKey"] = args ? args.clientKey : undefined;
            resourceInputs["defaultLeaseTtlSeconds"] = args ? args.defaultLeaseTtlSeconds : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["local"] = args ? args.local : undefined;
            resourceInputs["maxLeaseTtlSeconds"] = args ? args.maxLeaseTtlSeconds : undefined;
            resourceInputs["maxTokenNameLength"] = args ? args.maxTokenNameLength : undefined;
            resourceInputs["maxTtl"] = args ? args.maxTtl : undefined;
            resourceInputs["token"] = args ? args.token : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NomadSecretBackend.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NomadSecretBackend resources.
 */
export interface NomadSecretBackendState {
    /**
     * Specifies the address of the Nomad instance, provided
     * as "protocol://host:port" like "http://127.0.0.1:4646".
     */
    address?: pulumi.Input<string>;
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `nomad`.
     */
    backend?: pulumi.Input<string>;
    /**
     * CA certificate to use when verifying the Nomad server certificate, must be
     * x509 PEM encoded.
     */
    caCert?: pulumi.Input<string>;
    /**
     * Client certificate to provide to the Nomad server, must be x509 PEM encoded.
     */
    clientCert?: pulumi.Input<string>;
    /**
     * Client certificate key to provide to the Nomad server, must be x509 PEM encoded.
     */
    clientKey?: pulumi.Input<string>;
    /**
     * Default lease duration for secrets in seconds.
     */
    defaultLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * Human-friendly description of the mount for the Active Directory backend.
     */
    description?: pulumi.Input<string>;
    /**
     * Mark the secrets engine as local-only. Local engines are not replicated or removed by
     * replication.Tolerance duration to use when checking the last rotation time.
     */
    local?: pulumi.Input<boolean>;
    /**
     * Maximum possible lease duration for secrets in seconds.
     */
    maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * Specifies the maximum length to use for the name of the Nomad token
     * generated with Generate Credential. If omitted, 0 is used and ignored, defaulting to the max value allowed
     * by the Nomad version.
     */
    maxTokenNameLength?: pulumi.Input<number>;
    /**
     * Maximum possible lease duration for secrets in seconds.
     */
    maxTtl?: pulumi.Input<number>;
    /**
     * Specifies the Nomad Management token to use.
     */
    token?: pulumi.Input<string>;
    /**
     * Specifies the ttl of the lease for the generated token.
     */
    ttl?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a NomadSecretBackend resource.
 */
export interface NomadSecretBackendArgs {
    /**
     * Specifies the address of the Nomad instance, provided
     * as "protocol://host:port" like "http://127.0.0.1:4646".
     */
    address?: pulumi.Input<string>;
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `nomad`.
     */
    backend?: pulumi.Input<string>;
    /**
     * CA certificate to use when verifying the Nomad server certificate, must be
     * x509 PEM encoded.
     */
    caCert?: pulumi.Input<string>;
    /**
     * Client certificate to provide to the Nomad server, must be x509 PEM encoded.
     */
    clientCert?: pulumi.Input<string>;
    /**
     * Client certificate key to provide to the Nomad server, must be x509 PEM encoded.
     */
    clientKey?: pulumi.Input<string>;
    /**
     * Default lease duration for secrets in seconds.
     */
    defaultLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * Human-friendly description of the mount for the Active Directory backend.
     */
    description?: pulumi.Input<string>;
    /**
     * Mark the secrets engine as local-only. Local engines are not replicated or removed by
     * replication.Tolerance duration to use when checking the last rotation time.
     */
    local?: pulumi.Input<boolean>;
    /**
     * Maximum possible lease duration for secrets in seconds.
     */
    maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * Specifies the maximum length to use for the name of the Nomad token
     * generated with Generate Credential. If omitted, 0 is used and ignored, defaulting to the max value allowed
     * by the Nomad version.
     */
    maxTokenNameLength?: pulumi.Input<number>;
    /**
     * Maximum possible lease duration for secrets in seconds.
     */
    maxTtl?: pulumi.Input<number>;
    /**
     * Specifies the Nomad Management token to use.
     */
    token?: pulumi.Input<string>;
    /**
     * Specifies the ttl of the lease for the generated token.
     */
    ttl?: pulumi.Input<number>;
}
