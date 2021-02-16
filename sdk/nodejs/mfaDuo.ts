// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a resource to manage [Duo MFA](https://www.vaultproject.io/docs/enterprise/mfa/mfa-duo.html).
 *
 * **Note** this feature is available only with Vault Enterprise.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const userpass = new vault.AuthBackend("userpass", {
 *     path: "userpass",
 *     type: "userpass",
 * });
 * const myDuo = new vault.MfaDuo("my_duo", {
 *     apiHostname: "api-2b5c39f5.duosecurity.com",
 *     integrationKey: "BIACEUEAXI20BNWTEYXT",
 *     mountAccessor: userpass.accessor,
 *     secretKey: "8C7THtrIigh2rPZQMbguugt8IUftWhMRCOBzbuyz",
 * });
 * ```
 *
 * ## Import
 *
 * Mounts can be imported using the `path`, e.g.
 *
 * ```sh
 *  $ pulumi import vault:index/mfaDuo:MfaDuo my_duo my_duo
 * ```
 */
export class MfaDuo extends pulumi.CustomResource {
    /**
     * Get an existing MfaDuo resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MfaDuoState, opts?: pulumi.CustomResourceOptions): MfaDuo {
        return new MfaDuo(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:index/mfaDuo:MfaDuo';

    /**
     * Returns true if the given object is an instance of MfaDuo.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MfaDuo {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MfaDuo.__pulumiType;
    }

    /**
     * `(string: <required>)` - API hostname for Duo.
     */
    public readonly apiHostname!: pulumi.Output<string>;
    /**
     * `(string: <required>)` - Integration key for Duo.
     */
    public readonly integrationKey!: pulumi.Output<string>;
    /**
     * `(string: <required>)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
     */
    public readonly mountAccessor!: pulumi.Output<string>;
    /**
     * `(string: <required>)` – Name of the MFA method.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * `(string)` - Push information for Duo.
     */
    public readonly pushInfo!: pulumi.Output<string | undefined>;
    /**
     * `(string: <required>)` - Secret key for Duo.
     */
    public readonly secretKey!: pulumi.Output<string>;
    /**
     * `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
     * - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
     * - entity.name: The name configured for the Entity
     * - alias.metadata.`<key>`: The value of the Alias's metadata parameter
     * - entity.metadata.`<key>`: The value of the Entity's metadata parameter
     */
    public readonly usernameFormat!: pulumi.Output<string | undefined>;

    /**
     * Create a MfaDuo resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MfaDuoArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MfaDuoArgs | MfaDuoState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MfaDuoState | undefined;
            inputs["apiHostname"] = state ? state.apiHostname : undefined;
            inputs["integrationKey"] = state ? state.integrationKey : undefined;
            inputs["mountAccessor"] = state ? state.mountAccessor : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["pushInfo"] = state ? state.pushInfo : undefined;
            inputs["secretKey"] = state ? state.secretKey : undefined;
            inputs["usernameFormat"] = state ? state.usernameFormat : undefined;
        } else {
            const args = argsOrState as MfaDuoArgs | undefined;
            if ((!args || args.apiHostname === undefined) && !opts.urn) {
                throw new Error("Missing required property 'apiHostname'");
            }
            if ((!args || args.integrationKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'integrationKey'");
            }
            if ((!args || args.mountAccessor === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mountAccessor'");
            }
            if ((!args || args.secretKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'secretKey'");
            }
            inputs["apiHostname"] = args ? args.apiHostname : undefined;
            inputs["integrationKey"] = args ? args.integrationKey : undefined;
            inputs["mountAccessor"] = args ? args.mountAccessor : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["pushInfo"] = args ? args.pushInfo : undefined;
            inputs["secretKey"] = args ? args.secretKey : undefined;
            inputs["usernameFormat"] = args ? args.usernameFormat : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(MfaDuo.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MfaDuo resources.
 */
export interface MfaDuoState {
    /**
     * `(string: <required>)` - API hostname for Duo.
     */
    readonly apiHostname?: pulumi.Input<string>;
    /**
     * `(string: <required>)` - Integration key for Duo.
     */
    readonly integrationKey?: pulumi.Input<string>;
    /**
     * `(string: <required>)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
     */
    readonly mountAccessor?: pulumi.Input<string>;
    /**
     * `(string: <required>)` – Name of the MFA method.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * `(string)` - Push information for Duo.
     */
    readonly pushInfo?: pulumi.Input<string>;
    /**
     * `(string: <required>)` - Secret key for Duo.
     */
    readonly secretKey?: pulumi.Input<string>;
    /**
     * `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
     * - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
     * - entity.name: The name configured for the Entity
     * - alias.metadata.`<key>`: The value of the Alias's metadata parameter
     * - entity.metadata.`<key>`: The value of the Entity's metadata parameter
     */
    readonly usernameFormat?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MfaDuo resource.
 */
export interface MfaDuoArgs {
    /**
     * `(string: <required>)` - API hostname for Duo.
     */
    readonly apiHostname: pulumi.Input<string>;
    /**
     * `(string: <required>)` - Integration key for Duo.
     */
    readonly integrationKey: pulumi.Input<string>;
    /**
     * `(string: <required>)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
     */
    readonly mountAccessor: pulumi.Input<string>;
    /**
     * `(string: <required>)` – Name of the MFA method.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * `(string)` - Push information for Duo.
     */
    readonly pushInfo?: pulumi.Input<string>;
    /**
     * `(string: <required>)` - Secret key for Duo.
     */
    readonly secretKey: pulumi.Input<string>;
    /**
     * `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
     * - alias.name: The name returned by the mount configured via the `mountAccessor` parameter
     * - entity.name: The name configured for the Entity
     * - alias.metadata.`<key>`: The value of the Alias's metadata parameter
     * - entity.metadata.`<key>`: The value of the Entity's metadata parameter
     */
    readonly usernameFormat?: pulumi.Input<string>;
}
