// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Managed
{
    /// <summary>
    /// A resource that manages the lifecycle of all [Managed Keys](https://www.vaultproject.io/docs/enterprise/managed-keys) in Vault.
    /// 
    /// **Note** this feature is available only with Vault Enterprise.
    /// 
    /// ## Import
    /// 
    /// Mounts can be imported using the `id` of `default`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:managed/keys:Keys keys default
    /// ```
    /// </summary>
    [VaultResourceType("vault:managed/keys:Keys")]
    public partial class Keys : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Configuration block for AWS Managed Keys
        /// </summary>
        [Output("aws")]
        public Output<ImmutableArray<Outputs.KeysAw>> Aws { get; private set; } = null!;

        /// <summary>
        /// Configuration block for Azure Managed Keys
        /// </summary>
        [Output("azures")]
        public Output<ImmutableArray<Outputs.KeysAzure>> Azures { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured namespace.
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Configuration block for PKCS Managed Keys
        /// </summary>
        [Output("pkcs")]
        public Output<ImmutableArray<Outputs.KeysPkc>> Pkcs { get; private set; } = null!;


        /// <summary>
        /// Create a Keys resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Keys(string name, KeysArgs? args = null, CustomResourceOptions? options = null)
            : base("vault:managed/keys:Keys", name, args ?? new KeysArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Keys(string name, Input<string> id, KeysState? state = null, CustomResourceOptions? options = null)
            : base("vault:managed/keys:Keys", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Keys resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Keys Get(string name, Input<string> id, KeysState? state = null, CustomResourceOptions? options = null)
        {
            return new Keys(name, id, state, options);
        }
    }

    public sealed class KeysArgs : global::Pulumi.ResourceArgs
    {
        [Input("aws")]
        private InputList<Inputs.KeysAwArgs>? _aws;

        /// <summary>
        /// Configuration block for AWS Managed Keys
        /// </summary>
        public InputList<Inputs.KeysAwArgs> Aws
        {
            get => _aws ?? (_aws = new InputList<Inputs.KeysAwArgs>());
            set => _aws = value;
        }

        [Input("azures")]
        private InputList<Inputs.KeysAzureArgs>? _azures;

        /// <summary>
        /// Configuration block for Azure Managed Keys
        /// </summary>
        public InputList<Inputs.KeysAzureArgs> Azures
        {
            get => _azures ?? (_azures = new InputList<Inputs.KeysAzureArgs>());
            set => _azures = value;
        }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured namespace.
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        [Input("pkcs")]
        private InputList<Inputs.KeysPkcArgs>? _pkcs;

        /// <summary>
        /// Configuration block for PKCS Managed Keys
        /// </summary>
        public InputList<Inputs.KeysPkcArgs> Pkcs
        {
            get => _pkcs ?? (_pkcs = new InputList<Inputs.KeysPkcArgs>());
            set => _pkcs = value;
        }

        public KeysArgs()
        {
        }
        public static new KeysArgs Empty => new KeysArgs();
    }

    public sealed class KeysState : global::Pulumi.ResourceArgs
    {
        [Input("aws")]
        private InputList<Inputs.KeysAwGetArgs>? _aws;

        /// <summary>
        /// Configuration block for AWS Managed Keys
        /// </summary>
        public InputList<Inputs.KeysAwGetArgs> Aws
        {
            get => _aws ?? (_aws = new InputList<Inputs.KeysAwGetArgs>());
            set => _aws = value;
        }

        [Input("azures")]
        private InputList<Inputs.KeysAzureGetArgs>? _azures;

        /// <summary>
        /// Configuration block for Azure Managed Keys
        /// </summary>
        public InputList<Inputs.KeysAzureGetArgs> Azures
        {
            get => _azures ?? (_azures = new InputList<Inputs.KeysAzureGetArgs>());
            set => _azures = value;
        }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured namespace.
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        [Input("pkcs")]
        private InputList<Inputs.KeysPkcGetArgs>? _pkcs;

        /// <summary>
        /// Configuration block for PKCS Managed Keys
        /// </summary>
        public InputList<Inputs.KeysPkcGetArgs> Pkcs
        {
            get => _pkcs ?? (_pkcs = new InputList<Inputs.KeysPkcGetArgs>());
            set => _pkcs = value;
        }

        public KeysState()
        {
        }
        public static new KeysState Empty => new KeysState();
    }
}
