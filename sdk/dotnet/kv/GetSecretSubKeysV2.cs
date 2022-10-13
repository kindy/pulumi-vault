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
        public static Task<GetSecretSubkeysV2Result> InvokeAsync(GetSecretSubkeysV2Args args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetSecretSubkeysV2Result>("vault:kv/getSecretSubkeysV2:getSecretSubkeysV2", args ?? new GetSecretSubkeysV2Args(), options.WithDefaults());

        public static Output<GetSecretSubkeysV2Result> Invoke(GetSecretSubkeysV2InvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetSecretSubkeysV2Result>("vault:kv/getSecretSubkeysV2:getSecretSubkeysV2", args ?? new GetSecretSubkeysV2InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetSecretSubkeysV2Args : global::Pulumi.InvokeArgs
    {
        [Input("depth")]
        public int? Depth { get; set; }

        [Input("mount", required: true)]
        public string Mount { get; set; } = null!;

        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("namespace")]
        public string? Namespace { get; set; }

        [Input("version")]
        public int? Version { get; set; }

        public GetSecretSubkeysV2Args()
        {
        }
        public static new GetSecretSubkeysV2Args Empty => new GetSecretSubkeysV2Args();
    }

    public sealed class GetSecretSubkeysV2InvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("depth")]
        public Input<int>? Depth { get; set; }

        [Input("mount", required: true)]
        public Input<string> Mount { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

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
        public readonly ImmutableDictionary<string, object> Data;
        public readonly string DataJson;
        public readonly int? Depth;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Mount;
        public readonly string Name;
        public readonly string? Namespace;
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
