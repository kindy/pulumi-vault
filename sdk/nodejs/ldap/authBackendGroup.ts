// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Provides a resource to create a group in an [LDAP auth backend within Vault](https://www.vaultproject.io/docs/auth/ldap.html).
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
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
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * LDAP authentication backend groups can be imported using the `path`, e.g.
 *
 * ```sh
 * $ pulumi import vault:ldap/authBackendGroup:AuthBackendGroup foo auth/ldap/groups/foo
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
     *
     * For more details on the usage of each argument consult the [Vault LDAP API documentation](https://www.vaultproject.io/api-docs/auth/ldap).
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * The LDAP groupname
     */
    public readonly groupname!: pulumi.Output<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as AuthBackendGroupState | undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["groupname"] = state ? state.groupname : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["policies"] = state ? state.policies : undefined;
        } else {
            const args = argsOrState as AuthBackendGroupArgs | undefined;
            if ((!args || args.groupname === undefined) && !opts.urn) {
                throw new Error("Missing required property 'groupname'");
            }
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["groupname"] = args ? args.groupname : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["policies"] = args ? args.policies : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(AuthBackendGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendGroup resources.
 */
export interface AuthBackendGroupState {
    /**
     * Path to the authentication backend
     *
     * For more details on the usage of each argument consult the [Vault LDAP API documentation](https://www.vaultproject.io/api-docs/auth/ldap).
     */
    backend?: pulumi.Input<string>;
    /**
     * The LDAP groupname
     */
    groupname?: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Policies which should be granted to members of the group
     */
    policies?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a AuthBackendGroup resource.
 */
export interface AuthBackendGroupArgs {
    /**
     * Path to the authentication backend
     *
     * For more details on the usage of each argument consult the [Vault LDAP API documentation](https://www.vaultproject.io/api-docs/auth/ldap).
     */
    backend?: pulumi.Input<string>;
    /**
     * The LDAP groupname
     */
    groupname: pulumi.Input<string>;
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
    /**
     * Policies which should be granted to members of the group
     */
    policies?: pulumi.Input<pulumi.Input<string>[]>;
}
