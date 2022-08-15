// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Ldap
{
    /// <summary>
    /// Provides a resource to create a group in an [LDAP auth backend within Vault](https://www.vaultproject.io/docs/auth/ldap.html).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var ldap = new Vault.Ldap.AuthBackend("ldap", new()
    ///     {
    ///         Path = "ldap",
    ///         Url = "ldaps://dc-01.example.org",
    ///         Userdn = "OU=Users,OU=Accounts,DC=example,DC=org",
    ///         Userattr = "sAMAccountName",
    ///         Upndomain = "EXAMPLE.ORG",
    ///         Discoverdn = false,
    ///         Groupdn = "OU=Groups,DC=example,DC=org",
    ///         Groupfilter = "(&amp;(objectClass=group)(member:1.2.840.113556.1.4.1941:={{.UserDN}}))",
    ///     });
    /// 
    ///     var @group = new Vault.Ldap.AuthBackendGroup("group", new()
    ///     {
    ///         Groupname = "dba",
    ///         Policies = new[]
    ///         {
    ///             "dba",
    ///         },
    ///         Backend = ldap.Path,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// LDAP authentication backend groups can be imported using the `path`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import vault:ldap/authBackendGroup:AuthBackendGroup foo auth/ldap/groups/foo
    /// ```
    /// </summary>
    [VaultResourceType("vault:ldap/authBackendGroup:AuthBackendGroup")]
    public partial class AuthBackendGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Path to the authentication backend
        /// </summary>
        [Output("backend")]
        public Output<string?> Backend { get; private set; } = null!;

        /// <summary>
        /// The LDAP groupname
        /// </summary>
        [Output("groupname")]
        public Output<string> Groupname { get; private set; } = null!;

        /// <summary>
        /// Policies which should be granted to members of the group
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<string>> Policies { get; private set; } = null!;


        /// <summary>
        /// Create a AuthBackendGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthBackendGroup(string name, AuthBackendGroupArgs args, CustomResourceOptions? options = null)
            : base("vault:ldap/authBackendGroup:AuthBackendGroup", name, args ?? new AuthBackendGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthBackendGroup(string name, Input<string> id, AuthBackendGroupState? state = null, CustomResourceOptions? options = null)
            : base("vault:ldap/authBackendGroup:AuthBackendGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AuthBackendGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthBackendGroup Get(string name, Input<string> id, AuthBackendGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthBackendGroup(name, id, state, options);
        }
    }

    public sealed class AuthBackendGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Path to the authentication backend
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// The LDAP groupname
        /// </summary>
        [Input("groupname", required: true)]
        public Input<string> Groupname { get; set; } = null!;

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// Policies which should be granted to members of the group
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        public AuthBackendGroupArgs()
        {
        }
        public static new AuthBackendGroupArgs Empty => new AuthBackendGroupArgs();
    }

    public sealed class AuthBackendGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Path to the authentication backend
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// The LDAP groupname
        /// </summary>
        [Input("groupname")]
        public Input<string>? Groupname { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// Policies which should be granted to members of the group
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        public AuthBackendGroupState()
        {
        }
        public static new AuthBackendGroupState Empty => new AuthBackendGroupState();
    }
}
