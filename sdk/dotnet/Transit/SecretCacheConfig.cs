// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Transit
{
    /// <summary>
    /// Configure the cache for the Transit Secret Backend in Vault.
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
    ///         var transit = new Vault.Mount("transit", new Vault.MountArgs
    ///         {
    ///             Path = "transit",
    ///             Type = "transit",
    ///             Description = "Example description",
    ///             DefaultLeaseTtlSeconds = 3600,
    ///             MaxLeaseTtlSeconds = 86400,
    ///         });
    ///         var cfg = new Vault.Transit.SecretCacheConfig("cfg", new Vault.Transit.SecretCacheConfigArgs
    ///         {
    ///             Backend = transit.Path,
    ///             Size = 500,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [VaultResourceType("vault:transit/secretCacheConfig:SecretCacheConfig")]
    public partial class SecretCacheConfig : Pulumi.CustomResource
    {
        /// <summary>
        /// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// The number of cache entries. 0 means unlimited.
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;


        /// <summary>
        /// Create a SecretCacheConfig resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretCacheConfig(string name, SecretCacheConfigArgs args, CustomResourceOptions? options = null)
            : base("vault:transit/secretCacheConfig:SecretCacheConfig", name, args ?? new SecretCacheConfigArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretCacheConfig(string name, Input<string> id, SecretCacheConfigState? state = null, CustomResourceOptions? options = null)
            : base("vault:transit/secretCacheConfig:SecretCacheConfig", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretCacheConfig resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretCacheConfig Get(string name, Input<string> id, SecretCacheConfigState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretCacheConfig(name, id, state, options);
        }
    }

    public sealed class SecretCacheConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// The number of cache entries. 0 means unlimited.
        /// </summary>
        [Input("size", required: true)]
        public Input<int> Size { get; set; } = null!;

        public SecretCacheConfigArgs()
        {
        }
    }

    public sealed class SecretCacheConfigState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The path the transit secret backend is mounted at, with no leading or trailing `/`s.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// The number of cache entries. 0 means unlimited.
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        public SecretCacheConfigState()
        {
        }
    }
}
