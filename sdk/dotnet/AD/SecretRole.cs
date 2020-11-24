// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.AD
{
    public partial class SecretRole : Pulumi.CustomResource
    {
        /// <summary>
        /// The mount path for the AD backend.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// Last time Vault rotated this service account's password.
        /// </summary>
        [Output("lastVaultRotation")]
        public Output<string> LastVaultRotation { get; private set; } = null!;

        /// <summary>
        /// Last time Vault set this service account's password.
        /// </summary>
        [Output("passwordLastSet")]
        public Output<string> PasswordLastSet { get; private set; } = null!;

        /// <summary>
        /// Name of the role.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// The username/logon name for the service account with which this role will be associated.
        /// </summary>
        [Output("serviceAccountName")]
        public Output<string> ServiceAccountName { get; private set; } = null!;

        /// <summary>
        /// In seconds, the default password time-to-live.
        /// </summary>
        [Output("ttl")]
        public Output<int?> Ttl { get; private set; } = null!;


        /// <summary>
        /// Create a SecretRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretRole(string name, SecretRoleArgs args, CustomResourceOptions? options = null)
            : base("vault:ad/secretRole:SecretRole", name, args ?? new SecretRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretRole(string name, Input<string> id, SecretRoleState? state = null, CustomResourceOptions? options = null)
            : base("vault:ad/secretRole:SecretRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretRole Get(string name, Input<string> id, SecretRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretRole(name, id, state, options);
        }
    }

    public sealed class SecretRoleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The mount path for the AD backend.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// Name of the role.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// The username/logon name for the service account with which this role will be associated.
        /// </summary>
        [Input("serviceAccountName", required: true)]
        public Input<string> ServiceAccountName { get; set; } = null!;

        /// <summary>
        /// In seconds, the default password time-to-live.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        public SecretRoleArgs()
        {
        }
    }

    public sealed class SecretRoleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The mount path for the AD backend.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// Last time Vault rotated this service account's password.
        /// </summary>
        [Input("lastVaultRotation")]
        public Input<string>? LastVaultRotation { get; set; }

        /// <summary>
        /// Last time Vault set this service account's password.
        /// </summary>
        [Input("passwordLastSet")]
        public Input<string>? PasswordLastSet { get; set; }

        /// <summary>
        /// Name of the role.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// The username/logon name for the service account with which this role will be associated.
        /// </summary>
        [Input("serviceAccountName")]
        public Input<string>? ServiceAccountName { get; set; }

        /// <summary>
        /// In seconds, the default password time-to-live.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        public SecretRoleState()
        {
        }
    }
}
