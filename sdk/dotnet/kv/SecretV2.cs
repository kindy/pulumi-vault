// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.kv
{
    /// <summary>
    /// Writes a KV-V2 secret to a given path in Vault.
    /// 
    /// For more information on Vault's KV-V2 secret backend
    /// [see here](https://www.vaultproject.io/docs/secrets/kv/kv-v2).
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var kvv2 = new Vault.Mount("kvv2", new()
    ///     {
    ///         Path = "kvv2",
    ///         Type = "kv",
    ///         Options = 
    ///         {
    ///             { "version", "2" },
    ///         },
    ///         Description = "KV Version 2 secret engine mount",
    ///     });
    /// 
    ///     var example = new Vault.Kv.SecretV2("example", new()
    ///     {
    ///         Mount = kvv2.Path,
    ///         Cas = 1,
    ///         DeleteAllVersions = true,
    ///         DataJson = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["zip"] = "zap",
    ///             ["foo"] = "bar",
    ///         }),
    ///         CustomMetadata = new Vault.kv.Inputs.SecretV2CustomMetadataArgs
    ///         {
    ///             MaxVersions = 5,
    ///             Data = 
    ///             {
    ///                 { "foo", "vault@example.com" },
    ///                 { "bar", "12345" },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// ## Required Vault Capabilities
    /// 
    /// Use of this resource requires the `create` or `update` capability
    /// (depending on whether the resource already exists) on the given path,
    /// the `delete` capability if the resource is removed from configuration,
    /// and the `read` capability for drift detection (by default).
    /// 
    /// ### Custom Metadata Configuration Options
    /// 
    /// * `max_versions` - (Optional) The number of versions to keep per key.
    /// 
    /// * `cas_required` - (Optional) If true, all keys will require the cas
    /// parameter to be set on all write requests.
    /// 
    /// * `delete_version_after` - (Optional) If set, specifies the length of time before
    /// a version is deleted. Accepts duration in integer seconds.
    /// 
    /// * `data` - (Optional) A string to string map describing the secret.
    /// 
    /// ## Import
    /// 
    /// KV-V2 secrets can be imported using the `path`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import vault:kv/secretV2:SecretV2 example kvv2/data/secret
    /// ```
    /// </summary>
    [VaultResourceType("vault:kv/secretV2:SecretV2")]
    public partial class SecretV2 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// This flag is required if `cas_required` is set to true
        /// on either the secret or the engine's config. In order for a
        /// write operation to be successful, cas must be set to the current version
        /// of the secret.
        /// </summary>
        [Output("cas")]
        public Output<int?> Cas { get; private set; } = null!;

        /// <summary>
        /// A nested block that allows configuring metadata for the
        /// KV secret. Refer to the
        /// Configuration Options for more info.
        /// </summary>
        [Output("customMetadata")]
        public Output<Outputs.SecretV2CustomMetadata> CustomMetadata { get; private set; } = null!;

        /// <summary>
        /// A mapping whose keys are the top-level data keys returned from
        /// Vault and whose values are the corresponding values. This map can only
        /// represent string data, so any non-string values returned from Vault are
        /// serialized as JSON.
        /// </summary>
        [Output("data")]
        public Output<ImmutableDictionary<string, object>> Data { get; private set; } = null!;

        /// <summary>
        /// JSON-encoded string that will be
        /// written as the secret data at the given path.
        /// </summary>
        [Output("dataJson")]
        public Output<string> DataJson { get; private set; } = null!;

        /// <summary>
        /// If set to true, permanently deletes all
        /// versions for the specified key.
        /// </summary>
        [Output("deleteAllVersions")]
        public Output<bool?> DeleteAllVersions { get; private set; } = null!;

        /// <summary>
        /// If set to true, disables reading secret from Vault;
        /// note: drift won't be detected.
        /// </summary>
        [Output("disableRead")]
        public Output<bool?> DisableRead { get; private set; } = null!;

        /// <summary>
        /// Metadata associated with this secret read from Vault.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, object>> Metadata { get; private set; } = null!;

        /// <summary>
        /// Path where KV-V2 engine is mounted.
        /// </summary>
        [Output("mount")]
        public Output<string> Mount { get; private set; } = null!;

        /// <summary>
        /// Full name of the secret. For a nested secret
        /// the name is the nested path excluding the mount and data
        /// prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
        /// the name is `foo/bar/baz`.
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
        /// An object that holds option settings.
        /// </summary>
        [Output("options")]
        public Output<ImmutableDictionary<string, object>?> Options { get; private set; } = null!;

        /// <summary>
        /// Full path where the KV-V2 secret will be written.
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;


        /// <summary>
        /// Create a SecretV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretV2(string name, SecretV2Args args, CustomResourceOptions? options = null)
            : base("vault:kv/secretV2:SecretV2", name, args ?? new SecretV2Args(), MakeResourceOptions(options, ""))
        {
        }

        private SecretV2(string name, Input<string> id, SecretV2State? state = null, CustomResourceOptions? options = null)
            : base("vault:kv/secretV2:SecretV2", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "data",
                    "dataJson",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecretV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretV2 Get(string name, Input<string> id, SecretV2State? state = null, CustomResourceOptions? options = null)
        {
            return new SecretV2(name, id, state, options);
        }
    }

    public sealed class SecretV2Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This flag is required if `cas_required` is set to true
        /// on either the secret or the engine's config. In order for a
        /// write operation to be successful, cas must be set to the current version
        /// of the secret.
        /// </summary>
        [Input("cas")]
        public Input<int>? Cas { get; set; }

        /// <summary>
        /// A nested block that allows configuring metadata for the
        /// KV secret. Refer to the
        /// Configuration Options for more info.
        /// </summary>
        [Input("customMetadata")]
        public Input<Inputs.SecretV2CustomMetadataArgs>? CustomMetadata { get; set; }

        [Input("dataJson", required: true)]
        private Input<string>? _dataJson;

        /// <summary>
        /// JSON-encoded string that will be
        /// written as the secret data at the given path.
        /// </summary>
        public Input<string>? DataJson
        {
            get => _dataJson;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _dataJson = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// If set to true, permanently deletes all
        /// versions for the specified key.
        /// </summary>
        [Input("deleteAllVersions")]
        public Input<bool>? DeleteAllVersions { get; set; }

        /// <summary>
        /// If set to true, disables reading secret from Vault;
        /// note: drift won't be detected.
        /// </summary>
        [Input("disableRead")]
        public Input<bool>? DisableRead { get; set; }

        /// <summary>
        /// Path where KV-V2 engine is mounted.
        /// </summary>
        [Input("mount", required: true)]
        public Input<string> Mount { get; set; } = null!;

        /// <summary>
        /// Full name of the secret. For a nested secret
        /// the name is the nested path excluding the mount and data
        /// prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
        /// the name is `foo/bar/baz`.
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

        [Input("options")]
        private InputMap<object>? _options;

        /// <summary>
        /// An object that holds option settings.
        /// </summary>
        public InputMap<object> Options
        {
            get => _options ?? (_options = new InputMap<object>());
            set => _options = value;
        }

        public SecretV2Args()
        {
        }
        public static new SecretV2Args Empty => new SecretV2Args();
    }

    public sealed class SecretV2State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// This flag is required if `cas_required` is set to true
        /// on either the secret or the engine's config. In order for a
        /// write operation to be successful, cas must be set to the current version
        /// of the secret.
        /// </summary>
        [Input("cas")]
        public Input<int>? Cas { get; set; }

        /// <summary>
        /// A nested block that allows configuring metadata for the
        /// KV secret. Refer to the
        /// Configuration Options for more info.
        /// </summary>
        [Input("customMetadata")]
        public Input<Inputs.SecretV2CustomMetadataGetArgs>? CustomMetadata { get; set; }

        [Input("data")]
        private InputMap<object>? _data;

        /// <summary>
        /// A mapping whose keys are the top-level data keys returned from
        /// Vault and whose values are the corresponding values. This map can only
        /// represent string data, so any non-string values returned from Vault are
        /// serialized as JSON.
        /// </summary>
        public InputMap<object> Data
        {
            get => _data ?? (_data = new InputMap<object>());
            set
            {
                var emptySecret = Output.CreateSecret(ImmutableDictionary.Create<string, object>());
                _data = Output.All(value, emptySecret).Apply(v => v[0]);
            }
        }

        [Input("dataJson")]
        private Input<string>? _dataJson;

        /// <summary>
        /// JSON-encoded string that will be
        /// written as the secret data at the given path.
        /// </summary>
        public Input<string>? DataJson
        {
            get => _dataJson;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _dataJson = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// If set to true, permanently deletes all
        /// versions for the specified key.
        /// </summary>
        [Input("deleteAllVersions")]
        public Input<bool>? DeleteAllVersions { get; set; }

        /// <summary>
        /// If set to true, disables reading secret from Vault;
        /// note: drift won't be detected.
        /// </summary>
        [Input("disableRead")]
        public Input<bool>? DisableRead { get; set; }

        [Input("metadata")]
        private InputMap<object>? _metadata;

        /// <summary>
        /// Metadata associated with this secret read from Vault.
        /// </summary>
        public InputMap<object> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<object>());
            set => _metadata = value;
        }

        /// <summary>
        /// Path where KV-V2 engine is mounted.
        /// </summary>
        [Input("mount")]
        public Input<string>? Mount { get; set; }

        /// <summary>
        /// Full name of the secret. For a nested secret
        /// the name is the nested path excluding the mount and data
        /// prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
        /// the name is `foo/bar/baz`.
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

        [Input("options")]
        private InputMap<object>? _options;

        /// <summary>
        /// An object that holds option settings.
        /// </summary>
        public InputMap<object> Options
        {
            get => _options ?? (_options = new InputMap<object>());
            set => _options = value;
        }

        /// <summary>
        /// Full path where the KV-V2 secret will be written.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        public SecretV2State()
        {
        }
        public static new SecretV2State Empty => new SecretV2State();
    }
}
