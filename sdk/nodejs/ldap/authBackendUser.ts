// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a user in an [LDAP auth backend within Vault](https://www.vaultproject.io/docs/auth/ldap.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const ldap = new vault.ldap.AuthBackend("ldap", {
 *     discoverdn: false,
 *     groupdn: "OU=Groups,DC=example,DC=org",
 *     groupfilter: "(&(objectClass=group)(member:1.2.840.113556.1.4.1941:={{.UserDN}}))",
 *     path: "ldap",
 *     upndomain: "EXAMPLE.ORG",
 *     url: "ldaps://dc-01.example.org",
 *     userattr: "sAMAccountName",
 *     userdn: "OU=Users,OU=Accounts,DC=example,DC=org",
 * });
 * const user = new vault.ldap.AuthBackendUser("user", {
 *     backend: ldap.path,
 *     policies: [
 *         "dba",
 *         "sysops",
 *     ],
 *     username: "test-user",
 * });
 * ```
 *
 * ## Import
 *
 * LDAP authentication backend users can be imported using the `path`, e.g.
 *
 * ```sh
 *  $ pulumi import vault:ldap/authBackendUser:AuthBackendUser foo auth/ldap/users/foo
 * ```
 */
export class AuthBackendUser extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackendUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendUserState, opts?: pulumi.CustomResourceOptions): AuthBackendUser {
        return new AuthBackendUser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:ldap/authBackendUser:AuthBackendUser';

    /**
     * Returns true if the given object is an instance of AuthBackendUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackendUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackendUser.__pulumiType;
    }

    /**
     * Path to the authentication backend
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * Override LDAP groups which should be granted to user
     */
    public readonly groups!: pulumi.Output<string[]>;
    /**
     * Policies which should be granted to user
     */
    public readonly policies!: pulumi.Output<string[]>;
    /**
     * The LDAP username
     */
    public readonly username!: pulumi.Output<string>;

    /**
     * Create a AuthBackendUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthBackendUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendUserArgs | AuthBackendUserState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthBackendUserState | undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["groups"] = state ? state.groups : undefined;
            inputs["policies"] = state ? state.policies : undefined;
            inputs["username"] = state ? state.username : undefined;
        } else {
            const args = argsOrState as AuthBackendUserArgs | undefined;
            if ((!args || args.username === undefined) && !opts.urn) {
                throw new Error("Missing required property 'username'");
            }
            inputs["backend"] = args ? args.backend : undefined;
            inputs["groups"] = args ? args.groups : undefined;
            inputs["policies"] = args ? args.policies : undefined;
            inputs["username"] = args ? args.username : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AuthBackendUser.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendUser resources.
 */
export interface AuthBackendUserState {
    /**
     * Path to the authentication backend
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * Override LDAP groups which should be granted to user
     */
    readonly groups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Policies which should be granted to user
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The LDAP username
     */
    readonly username?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthBackendUser resource.
 */
export interface AuthBackendUserArgs {
    /**
     * Path to the authentication backend
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * Override LDAP groups which should be granted to user
     */
    readonly groups?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Policies which should be granted to user
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The LDAP username
     */
    readonly username: pulumi.Input<string>;
}
