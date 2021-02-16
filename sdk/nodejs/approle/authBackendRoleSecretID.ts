// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Manages an AppRole auth backend SecretID in a Vault server. See the [Vault
 * documentation](https://www.vaultproject.io/docs/auth/approle) for more
 * information.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const approle = new vault.AuthBackend("approle", {
 *     type: "approle",
 * });
 * const example = new vault.appRole.AuthBackendRole("example", {
 *     backend: approle.path,
 *     policies: [
 *         "default",
 *         "dev",
 *         "prod",
 *     ],
 *     roleName: "test-role",
 * });
 * const id = new vault.appRole.AuthBackendRoleSecretID("id", {
 *     backend: approle.path,
 *     metadata: `{
 *   "hello": "world"
 * }
 * `,
 *     roleName: example.roleName,
 * });
 * ```
 */
export class AuthBackendRoleSecretID extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackendRoleSecretID resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendRoleSecretIDState, opts?: pulumi.CustomResourceOptions): AuthBackendRoleSecretID {
        return new AuthBackendRoleSecretID(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:appRole/authBackendRoleSecretID:AuthBackendRoleSecretID';

    /**
     * Returns true if the given object is an instance of AuthBackendRoleSecretID.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackendRoleSecretID {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackendRoleSecretID.__pulumiType;
    }

    /**
     * The unique ID for this SecretID that can be safely logged.
     */
    public /*out*/ readonly accessor!: pulumi.Output<string>;
    /**
     * Unique name of the auth backend to configure.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * If set, specifies blocks of IP addresses which can
     * perform the login operation using this SecretID.
     */
    public readonly cidrLists!: pulumi.Output<string[] | undefined>;
    /**
     * A JSON-encoded string containing metadata in
     * key-value pairs to be set on tokens issued with this SecretID.
     */
    public readonly metadata!: pulumi.Output<string | undefined>;
    /**
     * The name of the role to create the SecretID for.
     */
    public readonly roleName!: pulumi.Output<string>;
    /**
     * The SecretID to be created. If set, uses "Push"
     * mode.  Defaults to Vault auto-generating SecretIDs.
     */
    public readonly secretId!: pulumi.Output<string>;
    /**
     * The unique ID for the response-wrapped SecretID that can
     * be safely logged.
     */
    public /*out*/ readonly wrappingAccessor!: pulumi.Output<string>;
    /**
     * The token used to retrieve a response-wrapped SecretID.
     */
    public /*out*/ readonly wrappingToken!: pulumi.Output<string>;
    /**
     * If set, the SecretID response will be
     * [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
     * and available for the duration specified. Only a single unwrapping of the
     * token is allowed.
     */
    public readonly wrappingTtl!: pulumi.Output<string | undefined>;

    /**
     * Create a AuthBackendRoleSecretID resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthBackendRoleSecretIDArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendRoleSecretIDArgs | AuthBackendRoleSecretIDState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthBackendRoleSecretIDState | undefined;
            inputs["accessor"] = state ? state.accessor : undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["cidrLists"] = state ? state.cidrLists : undefined;
            inputs["metadata"] = state ? state.metadata : undefined;
            inputs["roleName"] = state ? state.roleName : undefined;
            inputs["secretId"] = state ? state.secretId : undefined;
            inputs["wrappingAccessor"] = state ? state.wrappingAccessor : undefined;
            inputs["wrappingToken"] = state ? state.wrappingToken : undefined;
            inputs["wrappingTtl"] = state ? state.wrappingTtl : undefined;
        } else {
            const args = argsOrState as AuthBackendRoleSecretIDArgs | undefined;
            if ((!args || args.roleName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleName'");
            }
            inputs["backend"] = args ? args.backend : undefined;
            inputs["cidrLists"] = args ? args.cidrLists : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["roleName"] = args ? args.roleName : undefined;
            inputs["secretId"] = args ? args.secretId : undefined;
            inputs["wrappingTtl"] = args ? args.wrappingTtl : undefined;
            inputs["accessor"] = undefined /*out*/;
            inputs["wrappingAccessor"] = undefined /*out*/;
            inputs["wrappingToken"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AuthBackendRoleSecretID.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendRoleSecretID resources.
 */
export interface AuthBackendRoleSecretIDState {
    /**
     * The unique ID for this SecretID that can be safely logged.
     */
    readonly accessor?: pulumi.Input<string>;
    /**
     * Unique name of the auth backend to configure.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * If set, specifies blocks of IP addresses which can
     * perform the login operation using this SecretID.
     */
    readonly cidrLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A JSON-encoded string containing metadata in
     * key-value pairs to be set on tokens issued with this SecretID.
     */
    readonly metadata?: pulumi.Input<string>;
    /**
     * The name of the role to create the SecretID for.
     */
    readonly roleName?: pulumi.Input<string>;
    /**
     * The SecretID to be created. If set, uses "Push"
     * mode.  Defaults to Vault auto-generating SecretIDs.
     */
    readonly secretId?: pulumi.Input<string>;
    /**
     * The unique ID for the response-wrapped SecretID that can
     * be safely logged.
     */
    readonly wrappingAccessor?: pulumi.Input<string>;
    /**
     * The token used to retrieve a response-wrapped SecretID.
     */
    readonly wrappingToken?: pulumi.Input<string>;
    /**
     * If set, the SecretID response will be
     * [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
     * and available for the duration specified. Only a single unwrapping of the
     * token is allowed.
     */
    readonly wrappingTtl?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthBackendRoleSecretID resource.
 */
export interface AuthBackendRoleSecretIDArgs {
    /**
     * Unique name of the auth backend to configure.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * If set, specifies blocks of IP addresses which can
     * perform the login operation using this SecretID.
     */
    readonly cidrLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A JSON-encoded string containing metadata in
     * key-value pairs to be set on tokens issued with this SecretID.
     */
    readonly metadata?: pulumi.Input<string>;
    /**
     * The name of the role to create the SecretID for.
     */
    readonly roleName: pulumi.Input<string>;
    /**
     * The SecretID to be created. If set, uses "Push"
     * mode.  Defaults to Vault auto-generating SecretIDs.
     */
    readonly secretId?: pulumi.Input<string>;
    /**
     * If set, the SecretID response will be
     * [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
     * and available for the duration specified. Only a single unwrapping of the
     * token is allowed.
     */
    readonly wrappingTtl?: pulumi.Input<string>;
}
