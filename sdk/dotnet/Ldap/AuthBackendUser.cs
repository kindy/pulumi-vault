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
    /// Provides a resource to create a user in an [LDAP auth backend within Vault](https://www.vaultproject.io/docs/auth/ldap.html).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var ldap = new Vault.Ldap.AuthBackend("ldap", new Vault.Ldap.AuthBackendArgs
    ///         {
    ///             Discoverdn = false,
    ///             Groupdn = "OU=Groups,DC=example,DC=org",
    ///             Groupfilter = "(&amp;(objectClass=group)(member:1.2.840.113556.1.4.1941:={{.UserDN}}))",
    ///             Path = "ldap",
    ///             Upndomain = "EXAMPLE.ORG",
    ///             Url = "ldaps://dc-01.example.org",
    ///             Userattr = "sAMAccountName",
    ///             Userdn = "OU=Users,OU=Accounts,DC=example,DC=org",
    ///         });
    ///         var user = new Vault.Ldap.AuthBackendUser("user", new Vault.Ldap.AuthBackendUserArgs
    ///         {
    ///             Backend = ldap.Path,
    ///             Policies = 
    ///             {
    ///                 "dba",
    ///                 "sysops",
    ///             },
    ///             Username = "test-user",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// LDAP authentication backend users can be imported using the `path`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import vault:ldap/authBackendUser:AuthBackendUser foo auth/ldap/users/foo
    /// ```
    /// </summary>
    public partial class AuthBackendUser : Pulumi.CustomResource
    {
        /// <summary>
        /// Path to the authentication backend
        /// </summary>
        [Output("backend")]
        public Output<string?> Backend { get; private set; } = null!;

        /// <summary>
        /// Override LDAP groups which should be granted to user
        /// </summary>
        [Output("groups")]
        public Output<ImmutableArray<string>> Groups { get; private set; } = null!;

        /// <summary>
        /// Policies which should be granted to user
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<string>> Policies { get; private set; } = null!;

        /// <summary>
        /// The LDAP username
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;


        /// <summary>
        /// Create a AuthBackendUser resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthBackendUser(string name, AuthBackendUserArgs args, CustomResourceOptions? options = null)
            : base("vault:ldap/authBackendUser:AuthBackendUser", name, args ?? new AuthBackendUserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthBackendUser(string name, Input<string> id, AuthBackendUserState? state = null, CustomResourceOptions? options = null)
            : base("vault:ldap/authBackendUser:AuthBackendUser", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AuthBackendUser resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthBackendUser Get(string name, Input<string> id, AuthBackendUserState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthBackendUser(name, id, state, options);
        }
    }

    public sealed class AuthBackendUserArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Path to the authentication backend
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// Override LDAP groups which should be granted to user
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// Policies which should be granted to user
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        /// <summary>
        /// The LDAP username
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public AuthBackendUserArgs()
        {
        }
    }

    public sealed class AuthBackendUserState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Path to the authentication backend
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        [Input("groups")]
        private InputList<string>? _groups;

        /// <summary>
        /// Override LDAP groups which should be granted to user
        /// </summary>
        public InputList<string> Groups
        {
            get => _groups ?? (_groups = new InputList<string>());
            set => _groups = value;
        }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// Policies which should be granted to user
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        /// <summary>
        /// The LDAP username
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public AuthBackendUserState()
        {
        }
    }
}
