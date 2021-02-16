// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * Azure auth backends can be imported using `auth/`, the `backend` path, and `/config` e.g.
 *
 * ```sh
 *  $ pulumi import vault:azure/authBackendConfig:AuthBackendConfig example auth/azure/config
 * ```
 */
export class AuthBackendConfig extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackendConfig resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendConfigState, opts?: pulumi.CustomResourceOptions): AuthBackendConfig {
        return new AuthBackendConfig(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:azure/authBackendConfig:AuthBackendConfig';

    /**
     * Returns true if the given object is an instance of AuthBackendConfig.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackendConfig {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackendConfig.__pulumiType;
    }

    /**
     * The path the Azure auth backend being configured was
     * mounted at.  Defaults to `azure`.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * The client id for credentials to query the Azure APIs.
     * Currently read permissions to query compute resources are required.
     */
    public readonly clientId!: pulumi.Output<string | undefined>;
    /**
     * The client secret for credentials to query the
     * Azure APIs.
     */
    public readonly clientSecret!: pulumi.Output<string | undefined>;
    /**
     * The Azure cloud environment. Valid values:
     * AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
     * AzureGermanCloud.  Defaults to `AzurePublicCloud`.
     */
    public readonly environment!: pulumi.Output<string | undefined>;
    /**
     * The configured URL for the application registered in
     * Azure Active Directory.
     */
    public readonly resource!: pulumi.Output<string>;
    /**
     * The tenant id for the Azure Active Directory
     * organization.
     */
    public readonly tenantId!: pulumi.Output<string>;

    /**
     * Create a AuthBackendConfig resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthBackendConfigArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendConfigArgs | AuthBackendConfigState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthBackendConfigState | undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["clientId"] = state ? state.clientId : undefined;
            inputs["clientSecret"] = state ? state.clientSecret : undefined;
            inputs["environment"] = state ? state.environment : undefined;
            inputs["resource"] = state ? state.resource : undefined;
            inputs["tenantId"] = state ? state.tenantId : undefined;
        } else {
            const args = argsOrState as AuthBackendConfigArgs | undefined;
            if ((!args || args.resource === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resource'");
            }
            if ((!args || args.tenantId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantId'");
            }
            inputs["backend"] = args ? args.backend : undefined;
            inputs["clientId"] = args ? args.clientId : undefined;
            inputs["clientSecret"] = args ? args.clientSecret : undefined;
            inputs["environment"] = args ? args.environment : undefined;
            inputs["resource"] = args ? args.resource : undefined;
            inputs["tenantId"] = args ? args.tenantId : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AuthBackendConfig.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendConfig resources.
 */
export interface AuthBackendConfigState {
    /**
     * The path the Azure auth backend being configured was
     * mounted at.  Defaults to `azure`.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * The client id for credentials to query the Azure APIs.
     * Currently read permissions to query compute resources are required.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * The client secret for credentials to query the
     * Azure APIs.
     */
    readonly clientSecret?: pulumi.Input<string>;
    /**
     * The Azure cloud environment. Valid values:
     * AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
     * AzureGermanCloud.  Defaults to `AzurePublicCloud`.
     */
    readonly environment?: pulumi.Input<string>;
    /**
     * The configured URL for the application registered in
     * Azure Active Directory.
     */
    readonly resource?: pulumi.Input<string>;
    /**
     * The tenant id for the Azure Active Directory
     * organization.
     */
    readonly tenantId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthBackendConfig resource.
 */
export interface AuthBackendConfigArgs {
    /**
     * The path the Azure auth backend being configured was
     * mounted at.  Defaults to `azure`.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * The client id for credentials to query the Azure APIs.
     * Currently read permissions to query compute resources are required.
     */
    readonly clientId?: pulumi.Input<string>;
    /**
     * The client secret for credentials to query the
     * Azure APIs.
     */
    readonly clientSecret?: pulumi.Input<string>;
    /**
     * The Azure cloud environment. Valid values:
     * AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
     * AzureGermanCloud.  Defaults to `AzurePublicCloud`.
     */
    readonly environment?: pulumi.Input<string>;
    /**
     * The configured URL for the application registered in
     * Azure Active Directory.
     */
    readonly resource: pulumi.Input<string>;
    /**
     * The tenant id for the Azure Active Directory
     * organization.
     */
    readonly tenantId: pulumi.Input<string>;
}
