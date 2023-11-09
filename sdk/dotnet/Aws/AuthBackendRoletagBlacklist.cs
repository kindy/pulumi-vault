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
    /// Configures the periodic tidying operation of the blacklisted role tag entries.
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
    ///     var exampleAuthBackendRoletagBlacklist = new Vault.Aws.AuthBackendRoletagBlacklist("exampleAuthBackendRoletagBlacklist", new()
    ///     {
    ///         Backend = exampleAuthBackend.Path,
    ///         SafetyBuffer = 360,
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [VaultResourceType("vault:aws/authBackendRoletagBlacklist:AuthBackendRoletagBlacklist")]
    public partial class AuthBackendRoletagBlacklist : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The path the AWS auth backend being configured was
        /// mounted at.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// If set to true, disables the periodic
        /// tidying of the roletag blacklist entries. Defaults to false.
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
        /// The amount of extra time that must have passed
        /// beyond the roletag expiration, before it is removed from the backend storage.
        /// Defaults to 259,200 seconds, or 72 hours.
        /// </summary>
        [Output("safetyBuffer")]
        public Output<int?> SafetyBuffer { get; private set; } = null!;


        /// <summary>
        /// Create a AuthBackendRoletagBlacklist resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthBackendRoletagBlacklist(string name, AuthBackendRoletagBlacklistArgs args, CustomResourceOptions? options = null)
            : base("vault:aws/authBackendRoletagBlacklist:AuthBackendRoletagBlacklist", name, args ?? new AuthBackendRoletagBlacklistArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthBackendRoletagBlacklist(string name, Input<string> id, AuthBackendRoletagBlacklistState? state = null, CustomResourceOptions? options = null)
            : base("vault:aws/authBackendRoletagBlacklist:AuthBackendRoletagBlacklist", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AuthBackendRoletagBlacklist resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthBackendRoletagBlacklist Get(string name, Input<string> id, AuthBackendRoletagBlacklistState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthBackendRoletagBlacklist(name, id, state, options);
        }
    }

    public sealed class AuthBackendRoletagBlacklistArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path the AWS auth backend being configured was
        /// mounted at.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// If set to true, disables the periodic
        /// tidying of the roletag blacklist entries. Defaults to false.
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
        /// The amount of extra time that must have passed
        /// beyond the roletag expiration, before it is removed from the backend storage.
        /// Defaults to 259,200 seconds, or 72 hours.
        /// </summary>
        [Input("safetyBuffer")]
        public Input<int>? SafetyBuffer { get; set; }

        public AuthBackendRoletagBlacklistArgs()
        {
        }
        public static new AuthBackendRoletagBlacklistArgs Empty => new AuthBackendRoletagBlacklistArgs();
    }

    public sealed class AuthBackendRoletagBlacklistState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path the AWS auth backend being configured was
        /// mounted at.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// If set to true, disables the periodic
        /// tidying of the roletag blacklist entries. Defaults to false.
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
        /// The amount of extra time that must have passed
        /// beyond the roletag expiration, before it is removed from the backend storage.
        /// Defaults to 259,200 seconds, or 72 hours.
        /// </summary>
        [Input("safetyBuffer")]
        public Input<int>? SafetyBuffer { get; set; }

        public AuthBackendRoletagBlacklistState()
        {
        }
        public static new AuthBackendRoletagBlacklistState Empty => new AuthBackendRoletagBlacklistState();
    }
}
