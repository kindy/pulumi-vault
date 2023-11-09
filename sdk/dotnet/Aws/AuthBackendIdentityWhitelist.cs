// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Aws
{
    /// <summary>
    /// Configures the periodic tidying operation of the whitelisted identity entries.
    /// 
    /// For more information, see the
    /// [Vault docs](https://www.vaultproject.io/api-docs/auth/aws#configure-identity-whitelist-tidy-operation).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleAuthBackend = new Vault.AuthBackend("exampleAuthBackend", new()
    ///     {
    ///         Type = "aws",
    ///     });
    /// 
    ///     var exampleAuthBackendIdentityWhitelist = new Vault.Aws.AuthBackendIdentityWhitelist("exampleAuthBackendIdentityWhitelist", new()
    ///     {
    ///         Backend = exampleAuthBackend.Path,
    ///         SafetyBuffer = 3600,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// AWS auth backend identity whitelists can be imported using `auth/`, the `backend` path, and `/config/tidy/identity-whitelist` e.g.
    /// 
    /// ```sh
    ///  $ pulumi import vault:aws/authBackendIdentityWhitelist:AuthBackendIdentityWhitelist example auth/aws/config/tidy/identity-whitelist
    /// ```
    /// </summary>
    [VaultResourceType("vault:aws/authBackendIdentityWhitelist:AuthBackendIdentityWhitelist")]
    public partial class AuthBackendIdentityWhitelist : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The path of the AWS backend being configured.
        /// </summary>
        [Output("backend")]
        public Output<string?> Backend { get; private set; } = null!;

        /// <summary>
        /// If set to true, disables the periodic
        /// tidying of the identity-whitelist entries.
        /// </summary>
        [Output("disablePeriodicTidy")]
        public Output<bool?> DisablePeriodicTidy { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// The amount of extra time, in minutes, that must
        /// have passed beyond the roletag expiration, before it is removed from the
        /// backend storage.
        /// </summary>
        [Output("safetyBuffer")]
        public Output<int?> SafetyBuffer { get; private set; } = null!;


        /// <summary>
        /// Create a AuthBackendIdentityWhitelist resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthBackendIdentityWhitelist(string name, AuthBackendIdentityWhitelistArgs? args = null, CustomResourceOptions? options = null)
            : base("vault:aws/authBackendIdentityWhitelist:AuthBackendIdentityWhitelist", name, args ?? new AuthBackendIdentityWhitelistArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthBackendIdentityWhitelist(string name, Input<string> id, AuthBackendIdentityWhitelistState? state = null, CustomResourceOptions? options = null)
            : base("vault:aws/authBackendIdentityWhitelist:AuthBackendIdentityWhitelist", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AuthBackendIdentityWhitelist resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthBackendIdentityWhitelist Get(string name, Input<string> id, AuthBackendIdentityWhitelistState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthBackendIdentityWhitelist(name, id, state, options);
        }
    }

    public sealed class AuthBackendIdentityWhitelistArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path of the AWS backend being configured.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// If set to true, disables the periodic
        /// tidying of the identity-whitelist entries.
        /// </summary>
        [Input("disablePeriodicTidy")]
        public Input<bool>? DisablePeriodicTidy { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The amount of extra time, in minutes, that must
        /// have passed beyond the roletag expiration, before it is removed from the
        /// backend storage.
        /// </summary>
        [Input("safetyBuffer")]
        public Input<int>? SafetyBuffer { get; set; }

        public AuthBackendIdentityWhitelistArgs()
        {
        }
        public static new AuthBackendIdentityWhitelistArgs Empty => new AuthBackendIdentityWhitelistArgs();
    }

    public sealed class AuthBackendIdentityWhitelistState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path of the AWS backend being configured.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// If set to true, disables the periodic
        /// tidying of the identity-whitelist entries.
        /// </summary>
        [Input("disablePeriodicTidy")]
        public Input<bool>? DisablePeriodicTidy { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The amount of extra time, in minutes, that must
        /// have passed beyond the roletag expiration, before it is removed from the
        /// backend storage.
        /// </summary>
        [Input("safetyBuffer")]
        public Input<int>? SafetyBuffer { get; set; }

        public AuthBackendIdentityWhitelistState()
        {
        }
        public static new AuthBackendIdentityWhitelistState Empty => new AuthBackendIdentityWhitelistState();
    }
}
