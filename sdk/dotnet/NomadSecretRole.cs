// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault
{
    /// <summary>
    /// ## Import
    /// 
    /// Nomad secret role can be imported using the `backend`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import vault:index/nomadSecretRole:NomadSecretRole bob nomad/role/bob
    /// ```
    /// </summary>
    public partial class NomadSecretRole : Pulumi.CustomResource
    {
        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `nomad`.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// Specifies if the generated token should be global. Defaults to 
        /// false.
        /// </summary>
        [Output("global")]
        public Output<bool> Global { get; private set; } = null!;

        /// <summary>
        /// List of policies attached to the generated token. This setting is only used 
        /// when `type` is 'client'.
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<string>> Policies { get; private set; } = null!;

        /// <summary>
        /// The name to identify this role within the backend.
        /// Must be unique within the backend.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// Specifies the type of token to create when using this role. Valid 
        /// settings are 'client' and 'management'. Defaults to 'client'.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a NomadSecretRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NomadSecretRole(string name, NomadSecretRoleArgs args, CustomResourceOptions? options = null)
            : base("vault:index/nomadSecretRole:NomadSecretRole", name, args ?? new NomadSecretRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NomadSecretRole(string name, Input<string> id, NomadSecretRoleState? state = null, CustomResourceOptions? options = null)
            : base("vault:index/nomadSecretRole:NomadSecretRole", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NomadSecretRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NomadSecretRole Get(string name, Input<string> id, NomadSecretRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new NomadSecretRole(name, id, state, options);
        }
    }

    public sealed class NomadSecretRoleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `nomad`.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// Specifies if the generated token should be global. Defaults to 
        /// false.
        /// </summary>
        [Input("global")]
        public Input<bool>? Global { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// List of policies attached to the generated token. This setting is only used 
        /// when `type` is 'client'.
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        /// <summary>
        /// The name to identify this role within the backend.
        /// Must be unique within the backend.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// Specifies the type of token to create when using this role. Valid 
        /// settings are 'client' and 'management'. Defaults to 'client'.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public NomadSecretRoleArgs()
        {
        }
    }

    public sealed class NomadSecretRoleState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `nomad`.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// Specifies if the generated token should be global. Defaults to 
        /// false.
        /// </summary>
        [Input("global")]
        public Input<bool>? Global { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// List of policies attached to the generated token. This setting is only used 
        /// when `type` is 'client'.
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        /// <summary>
        /// The name to identify this role within the backend.
        /// Must be unique within the backend.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// Specifies the type of token to create when using this role. Valid 
        /// settings are 'client' and 'management'. Defaults to 'client'.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public NomadSecretRoleState()
        {
        }
    }
}