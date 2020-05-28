// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Identity
{
    /// <summary>
    /// Creates an Identity Group Alias for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.
    /// 
    /// Group aliases allows entity membership in external groups to be managed semi-automatically. External group serves as a mapping to a group that is outside of the identity store. External groups can have one (and only one) alias. This alias should map to a notion of group that is outside of the identity store. For example, groups in LDAP, and teams in GitHub. A username in LDAP, belonging to a group in LDAP, can get its entity ID added as a member of a group in Vault automatically during logins and token renewals. This works only if the group in Vault is an external group and has an alias that maps to the group in LDAP. If the user is removed from the group in LDAP, that change gets reflected in Vault only upon the subsequent login or renewal operation.
    /// 
    /// ## Example Usage
    /// 
    /// 
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @group = new Vault.Identity.Group("group", new Vault.Identity.GroupArgs
    ///         {
    ///             Policies = 
    ///             {
    ///                 "test",
    ///             },
    ///             Type = "external",
    ///         });
    ///         var github = new Vault.AuthBackend("github", new Vault.AuthBackendArgs
    ///         {
    ///             Path = "github",
    ///             Type = "github",
    ///         });
    ///         var group_alias = new Vault.Identity.GroupAlias("group-alias", new Vault.Identity.GroupAliasArgs
    ///         {
    ///             CanonicalId = @group.Id,
    ///             MountAccessor = github.Accessor,
    ///             Name = "Github_Team_Slug",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class GroupAlias : Pulumi.CustomResource
    {
        /// <summary>
        /// ID of the group to which this is an alias.
        /// </summary>
        [Output("canonicalId")]
        public Output<string> CanonicalId { get; private set; } = null!;

        /// <summary>
        /// Mount accessor of the authentication backend to which this alias belongs to.
        /// </summary>
        [Output("mountAccessor")]
        public Output<string> MountAccessor { get; private set; } = null!;

        /// <summary>
        /// Name of the group alias to create.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a GroupAlias resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupAlias(string name, GroupAliasArgs args, CustomResourceOptions? options = null)
            : base("vault:identity/groupAlias:GroupAlias", name, args ?? new GroupAliasArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupAlias(string name, Input<string> id, GroupAliasState? state = null, CustomResourceOptions? options = null)
            : base("vault:identity/groupAlias:GroupAlias", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing GroupAlias resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupAlias Get(string name, Input<string> id, GroupAliasState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupAlias(name, id, state, options);
        }
    }

    public sealed class GroupAliasArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the group to which this is an alias.
        /// </summary>
        [Input("canonicalId", required: true)]
        public Input<string> CanonicalId { get; set; } = null!;

        /// <summary>
        /// Mount accessor of the authentication backend to which this alias belongs to.
        /// </summary>
        [Input("mountAccessor", required: true)]
        public Input<string> MountAccessor { get; set; } = null!;

        /// <summary>
        /// Name of the group alias to create.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GroupAliasArgs()
        {
        }
    }

    public sealed class GroupAliasState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the group to which this is an alias.
        /// </summary>
        [Input("canonicalId")]
        public Input<string>? CanonicalId { get; set; }

        /// <summary>
        /// Mount accessor of the authentication backend to which this alias belongs to.
        /// </summary>
        [Input("mountAccessor")]
        public Input<string>? MountAccessor { get; set; }

        /// <summary>
        /// Name of the group alias to create.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public GroupAliasState()
        {
        }
    }
}
