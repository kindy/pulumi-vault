// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * The provider type for the vault package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'vault';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Provider.__pulumiType;
    }


    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let inputs: pulumi.Inputs = {};
        {
            inputs["addAddressToEnv"] = args ? args.addAddressToEnv : undefined;
            inputs["address"] = (args ? args.address : undefined) || utilities.getEnv("VAULT_ADDR");
            inputs["authLogins"] = pulumi.output(args ? args.authLogins : undefined).apply(JSON.stringify);
            inputs["caCertDir"] = (args ? args.caCertDir : undefined) || utilities.getEnv("VAULT_CAPATH");
            inputs["caCertFile"] = (args ? args.caCertFile : undefined) || utilities.getEnv("VAULT_CACERT");
            inputs["clientAuths"] = pulumi.output(args ? args.clientAuths : undefined).apply(JSON.stringify);
            inputs["maxLeaseTtlSeconds"] = pulumi.output((args ? args.maxLeaseTtlSeconds : undefined) || (utilities.getEnvNumber("TERRAFORM_VAULT_MAX_TTL") || 1200)).apply(JSON.stringify);
            inputs["maxRetries"] = pulumi.output((args ? args.maxRetries : undefined) || (utilities.getEnvNumber("VAULT_MAX_RETRIES") || 2)).apply(JSON.stringify);
            inputs["namespace"] = (args ? args.namespace : undefined) || utilities.getEnv("VAULT_NAMESPACE");
            inputs["skipTlsVerify"] = pulumi.output((args ? args.skipTlsVerify : undefined) || utilities.getEnvBoolean("VAULT_SKIP_VERIFY")).apply(JSON.stringify);
            inputs["token"] = (args ? args.token : undefined) || utilities.getEnv("VAULT_TOKEN");
            inputs["tokenName"] = (args ? args.tokenName : undefined) || utilities.getEnv("VAULT_TOKEN_NAME");
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(Provider.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * If true, adds the value of the `address` argument to the Terraform process environment.
     */
    readonly addAddressToEnv?: pulumi.Input<string>;
    /**
     * URL of the root of the target Vault server.
     */
    readonly address?: pulumi.Input<string>;
    /**
     * Login to vault with an existing auth method using auth/<mount>/login
     */
    readonly authLogins?: pulumi.Input<pulumi.Input<inputs.ProviderAuthLogin>[]>;
    /**
     * Path to directory containing CA certificate files to validate the server's certificate.
     */
    readonly caCertDir?: pulumi.Input<string>;
    /**
     * Path to a CA certificate file to validate the server's certificate.
     */
    readonly caCertFile?: pulumi.Input<string>;
    /**
     * Client authentication credentials.
     */
    readonly clientAuths?: pulumi.Input<pulumi.Input<inputs.ProviderClientAuth>[]>;
    /**
     * Maximum TTL for secret leases requested by this provider
     */
    readonly maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * Maximum number of retries when a 5xx error code is encountered.
     */
    readonly maxRetries?: pulumi.Input<number>;
    /**
     * The namespace to use. Available only for Vault Enterprise
     */
    readonly namespace?: pulumi.Input<string>;
    /**
     * Set this to true only if the target Vault server is an insecure development instance.
     */
    readonly skipTlsVerify?: pulumi.Input<boolean>;
    /**
     * Token to use to authenticate to Vault.
     */
    readonly token?: pulumi.Input<string>;
    /**
     * Token name to use for creating the Vault child token.
     */
    readonly tokenName?: pulumi.Input<string>;
}
