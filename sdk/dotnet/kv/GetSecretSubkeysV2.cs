// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.kv
{
    public static class GetSecretSubkeysV2
    {
        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var awsSecret = new Vault.Kv.SecretV2("awsSecret", new()
        ///     {
        ///         Mount = kvv2.Path,
        ///         DataJson = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["zip"] = "zap",
        ///             ["foo"] = "bar",
        ///         }),
        ///     });
        /// 
        ///     var test = Vault.kv.GetSecretSubkeysV2.Invoke(new()
        ///     {
        ///         Mount = kvv2.Path,
        ///         Name = awsSecret.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Required Vault Capabilities
        /// 
        /// Use of this resource requires the `read` capability on the given path.
        /// </summary>
        public static Task<GetSecretSubkeysV2Result> InvokeAsync(GetSecretSubkeysV2Args args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecretSubkeysV2Result>("vault:kv/getSecretSubkeysV2:getSecretSubkeysV2", args ?? new GetSecretSubkeysV2Args(), options.WithDefaults());

        /// <summary>
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
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
        ///     var awsSecret = new Vault.Kv.SecretV2("awsSecret", new()
        ///     {
        ///         Mount = kvv2.Path,
        ///         DataJson = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
        ///         {
        ///             ["zip"] = "zap",
        ///             ["foo"] = "bar",
        ///         }),
        ///     });
        /// 
        ///     var test = Vault.kv.GetSecretSubkeysV2.Invoke(new()
        ///     {
        ///         Mount = kvv2.Path,
        ///         Name = awsSecret.Name,
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// ## Required Vault Capabilities
        /// 
        /// Use of this resource requires the `read` capability on the given path.
        /// </summary>
        public static Output<GetSecretSubkeysV2Result> Invoke(GetSecretSubkeysV2InvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecretSubkeysV2Result>("vault:kv/getSecretSubkeysV2:getSecretSubkeysV2", args ?? new GetSecretSubkeysV2InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecretSubkeysV2Args : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the deepest nesting level to provide in the output.
        /// If non-zero, keys that reside at the specified depth value will be
        /// artificially treated as leaves and will thus be `null` even if further
        /// underlying sub-keys exist.
        /// </summary>
        [Input("depth")]
        public int? Depth { get; set; }

        /// <summary>
        /// Path where KV-V2 engine is mounted.
        /// </summary>
        [Input("mount", required: true)]
        public string Mount { get; set; } = null!;

        /// <summary>
        /// Full name of the secret. For a nested secret
        /// the name is the nested path excluding the mount and data
        /// prefix. For example, for a secret at `kvv2/data/foo/bar/baz`
        /// the name is `foo/bar/baz`.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public string? Namespace { get; set; }

        /// <summary>
        /// Specifies the version to return. If not 
        /// set the latest version is returned.
        /// </summary>
        [Input("version")]
        public int? Version { get; set; }

        public GetSecretSubkeysV2Args()
        {
        }
        public static new GetSecretSubkeysV2Args Empty => new GetSecretSubkeysV2Args();
    }

    public sealed class GetSecretSubkeysV2InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Specifies the deepest nesting level to provide in the output.
        /// If non-zero, keys that reside at the specified depth value will be
        /// artificially treated as leaves and will thus be `null` even if further
        /// underlying sub-keys exist.
        /// </summary>
        [Input("depth")]
        public Input<int>? Depth { get; set; }

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
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Specifies the version to return. If not 
        /// set the latest version is returned.
        /// </summary>
        [Input("version")]
        public Input<int>? Version { get; set; }

        public GetSecretSubkeysV2InvokeArgs()
        {
        }
        public static new GetSecretSubkeysV2InvokeArgs Empty => new GetSecretSubkeysV2InvokeArgs();
    }


    [OutputType]
    public sealed class GetSecretSubkeysV2Result
    {
        /// <summary>
        /// Subkeys for the KV-V2 secret stored as a serialized map of strings.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Data;
        /// <summary>
        /// Subkeys for the KV-V2 secret read from Vault.
        /// </summary>
        public readonly string DataJson;
        public readonly int? Depth;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Mount;
        public readonly string Name;
        public readonly string? Namespace;
        /// <summary>
        /// Full path where the KV-V2 secrets are listed.
        /// </summary>
        public readonly string Path;
        public readonly int? Version;

        [OutputConstructor]
        private GetSecretSubkeysV2Result(
            ImmutableDictionary<string, object> data,

            string dataJson,

            int? depth,

            string id,

            string mount,

            string name,

            string? @namespace,

            string path,

            int? version)
        {
            Data = data;
            DataJson = dataJson;
            Depth = depth;
            Id = id;
            Mount = mount;
            Name = name;
            Namespace = @namespace;
            Path = path;
            Version = version;
        }
    }
}
