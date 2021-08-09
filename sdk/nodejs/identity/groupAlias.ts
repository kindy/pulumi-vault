// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates an Identity Group Alias for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.
 *
 * Group aliases allows entity membership in external groups to be managed semi-automatically. External group serves as a mapping to a group that is outside of the identity store. External groups can have one (and only one) alias. This alias should map to a notion of group that is outside of the identity store. For example, groups in LDAP, and teams in GitHub. A username in LDAP, belonging to a group in LDAP, can get its entity ID added as a member of a group in Vault automatically during logins and token renewals. This works only if the group in Vault is an external group and has an alias that maps to the group in LDAP. If the user is removed from the group in LDAP, that change gets reflected in Vault only upon the subsequent login or renewal operation.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const group = new vault.identity.Group("group", {
 *     type: "external",
 *     policies: ["test"],
 * });
 * const github = new vault.AuthBackend("github", {
 *     type: "github",
 *     path: "github",
 * });
 * const group_alias = new vault.identity.GroupAlias("group-alias", {
 *     name: "Github_Team_Slug",
 *     mountAccessor: github.accessor,
 *     canonicalId: group.id,
 * });
 * ```
 *
 * ## Import
 *
 * The group alias can be imported with the group alias `id`, for example
 *
 * ```sh
 *  $ pulumi import vault:identity/groupAlias:GroupAlias group-alias id
 * ```
 *
 *  Group aliases can also be imported using the UUID of the alias record, e.g.
 *
 * ```sh
 *  $ pulumi import vault:identity/groupAlias:GroupAlias alias_name 63104e20-88e4-11eb-8d04-cf7ac9d60157
 * ```
 */
export class GroupAlias extends pulumi.CustomResource {
    /**
     * Get an existing GroupAlias resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GroupAliasState, opts?: pulumi.CustomResourceOptions): GroupAlias {
        return new GroupAlias(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:identity/groupAlias:GroupAlias';

    /**
     * Returns true if the given object is an instance of GroupAlias.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GroupAlias {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GroupAlias.__pulumiType;
    }

    /**
     * ID of the group to which this is an alias.
     */
    public readonly canonicalId!: pulumi.Output<string>;
    /**
     * Mount accessor of the authentication backend to which this alias belongs to.
     */
    public readonly mountAccessor!: pulumi.Output<string>;
    /**
     * Name of the group alias to create.
     */
    public readonly name!: pulumi.Output<string>;

    /**
     * Create a GroupAlias resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GroupAliasArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GroupAliasArgs | GroupAliasState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GroupAliasState | undefined;
            inputs["canonicalId"] = state ? state.canonicalId : undefined;
            inputs["mountAccessor"] = state ? state.mountAccessor : undefined;
            inputs["name"] = state ? state.name : undefined;
        } else {
            const args = argsOrState as GroupAliasArgs | undefined;
            if ((!args || args.canonicalId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'canonicalId'");
            }
            if ((!args || args.mountAccessor === undefined) && !opts.urn) {
                throw new Error("Missing required property 'mountAccessor'");
            }
            if ((!args || args.name === undefined) && !opts.urn) {
                throw new Error("Missing required property 'name'");
            }
            inputs["canonicalId"] = args ? args.canonicalId : undefined;
            inputs["mountAccessor"] = args ? args.mountAccessor : undefined;
            inputs["name"] = args ? args.name : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(GroupAlias.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GroupAlias resources.
 */
export interface GroupAliasState {
    /**
     * ID of the group to which this is an alias.
     */
    readonly canonicalId?: pulumi.Input<string>;
    /**
     * Mount accessor of the authentication backend to which this alias belongs to.
     */
    readonly mountAccessor?: pulumi.Input<string>;
    /**
     * Name of the group alias to create.
     */
    readonly name?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a GroupAlias resource.
 */
export interface GroupAliasArgs {
    /**
     * ID of the group to which this is an alias.
     */
    readonly canonicalId: pulumi.Input<string>;
    /**
     * Mount accessor of the authentication backend to which this alias belongs to.
     */
    readonly mountAccessor: pulumi.Input<string>;
    /**
     * Name of the group alias to create.
     */
    readonly name: pulumi.Input<string>;
}
