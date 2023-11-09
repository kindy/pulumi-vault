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
    ///     var test = new Vault.Identity.EntityAlias("test", new()
    ///     {
    ///         CanonicalId = "49877D63-07AD-4B85-BDA8-B61626C477E8",
    ///         MountAccessor = "token_1f2bd5",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Identity entity alias can be imported using the `id`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import vault:identity/entityAlias:EntityAlias test "3856fb4d-3c91-dcaf-2401-68f446796bfb"
    /// ```
    /// </summary>
    [VaultResourceType("vault:identity/entityAlias:EntityAlias")]
    public partial class EntityAlias : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Entity ID to which this alias belongs to.
        /// </summary>
        [Output("canonicalId")]
        public Output<string> CanonicalId { get; private set; } = null!;

        /// <summary>
        /// Custom metadata to be associated with this alias.
        /// </summary>
        [Output("customMetadata")]
        public Output<ImmutableDictionary<string, string>?> CustomMetadata { get; private set; } = null!;

        /// <summary>
        /// Accessor of the mount to which the alias should belong to.
        /// </summary>
        [Output("mountAccessor")]
        public Output<string> MountAccessor { get; private set; } = null!;

        /// <summary>
        /// Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;


        /// <summary>
        /// Create a EntityAlias resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public EntityAlias(string name, EntityAliasArgs args, CustomResourceOptions? options = null)
            : base("vault:identity/entityAlias:EntityAlias", name, args ?? new EntityAliasArgs(), MakeResourceOptions(options, ""))
        {
        }

        private EntityAlias(string name, Input<string> id, EntityAliasState? state = null, CustomResourceOptions? options = null)
            : base("vault:identity/entityAlias:EntityAlias", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing EntityAlias resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static EntityAlias Get(string name, Input<string> id, EntityAliasState? state = null, CustomResourceOptions? options = null)
        {
            return new EntityAlias(name, id, state, options);
        }
    }

    public sealed class EntityAliasArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Entity ID to which this alias belongs to.
        /// </summary>
        [Input("canonicalId", required: true)]
        public Input<string> CanonicalId { get; set; } = null!;

        [Input("customMetadata")]
        private InputMap<string>? _customMetadata;

        /// <summary>
        /// Custom metadata to be associated with this alias.
        /// </summary>
        public InputMap<string> CustomMetadata
        {
            get => _customMetadata ?? (_customMetadata = new InputMap<string>());
            set => _customMetadata = value;
        }

        /// <summary>
        /// Accessor of the mount to which the alias should belong to.
        /// </summary>
        [Input("mountAccessor", required: true)]
        public Input<string> MountAccessor { get; set; } = null!;

        /// <summary>
        /// Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public EntityAliasArgs()
        {
        }
        public static new EntityAliasArgs Empty => new EntityAliasArgs();
    }

    public sealed class EntityAliasState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Entity ID to which this alias belongs to.
        /// </summary>
        [Input("canonicalId")]
        public Input<string>? CanonicalId { get; set; }

        [Input("customMetadata")]
        private InputMap<string>? _customMetadata;

        /// <summary>
        /// Custom metadata to be associated with this alias.
        /// </summary>
        public InputMap<string> CustomMetadata
        {
            get => _customMetadata ?? (_customMetadata = new InputMap<string>());
            set => _customMetadata = value;
        }

        /// <summary>
        /// Accessor of the mount to which the alias should belong to.
        /// </summary>
        [Input("mountAccessor")]
        public Input<string>? MountAccessor { get; set; }

        /// <summary>
        /// Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public EntityAliasState()
        {
        }
        public static new EntityAliasState Empty => new EntityAliasState();
    }
}
