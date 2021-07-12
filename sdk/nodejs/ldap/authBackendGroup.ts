// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a group in an [LDAP auth backend within Vault](https://www.vaultproject.io/docs/auth/ldap.html).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const ldap = new vault.ldap.AuthBackend("ldap", {
 *     path: "ldap",
 *     url: "ldaps://dc-01.example.org",
 *     userdn: "OU=Users,OU=Accounts,DC=example,DC=org",
 *     userattr: "sAMAccountName",
 *     upndomain: "EXAMPLE.ORG",
 *     discoverdn: false,
 *     groupdn: "OU=Groups,DC=example,DC=org",
 *     groupfilter: "(&(objectClass=group)(member:1.2.840.113556.1.4.1941:={{.UserDN}}))",
 * });
 * const group = new vault.ldap.AuthBackendGroup("group", {
 *     groupname: "dba",
 *     policies: ["dba"],
 *     backend: ldap.path,
 * });
 * ```
 *
 * ## Import
 *
 * LDAP authentication backend groups can be imported using the `path`, e.g.
 *
 * ```sh
 *  $ pulumi import vault:ldap/authBackendGroup:AuthBackendGroup foo auth/ldap/groups/foo
 * ```
 */
export class AuthBackendGroup extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackendGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendGroupState, opts?: pulumi.CustomResourceOptions): AuthBackendGroup {
        return new AuthBackendGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:ldap/authBackendGroup:AuthBackendGroup';

    /**
     * Returns true if the given object is an instance of AuthBackendGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackendGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackendGroup.__pulumiType;
    }

    /**
     * Path to the authentication backend
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * The LDAP groupname
     */
    public readonly groupname!: pulumi.Output<string>;
    /**
     * Policies which should be granted to members of the group
     */
    public readonly policies!: pulumi.Output<string[]>;

    /**
     * Create a AuthBackendGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthBackendGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendGroupArgs | AuthBackendGroupState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthBackendGroupState | undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["groupname"] = state ? state.groupname : undefined;
            inputs["policies"] = state ? state.policies : undefined;
        } else {
            const args = argsOrState as AuthBackendGroupArgs | undefined;
            if ((!args || args.groupname === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupname'");
            }
            inputs["backend"] = args ? args.backend : undefined;
            inputs["groupname"] = args ? args.groupname : undefined;
            inputs["policies"] = args ? args.policies : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(AuthBackendGroup.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendGroup resources.
 */
export interface AuthBackendGroupState {
    /**
     * Path to the authentication backend
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * The LDAP groupname
     */
    readonly groupname?: pulumi.Input<string>;
    /**
     * Policies which should be granted to members of the group
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a AuthBackendGroup resource.
 */
export interface AuthBackendGroupArgs {
    /**
     * Path to the authentication backend
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * The LDAP groupname
     */
    readonly groupname: pulumi.Input<string>;
    /**
     * Policies which should be granted to members of the group
     */
    readonly policies?: pulumi.Input<pulumi.Input<string>[]>;
}
