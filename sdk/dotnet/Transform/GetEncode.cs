// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Transform
{
    public static class GetEncode
    {
        /// <summary>
        /// This data source supports the "/transform/encode/{role_name}" Vault endpoint.
        /// 
        /// It encodes the provided value using a named role.
        /// </summary>
        public static Task<GetEncodeResult> InvokeAsync(GetEncodeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEncodeResult>("vault:transform/getEncode:getEncode", args ?? new GetEncodeArgs(), options.WithVersion());
    }


    public sealed class GetEncodeArgs : Pulumi.InvokeArgs
    {
        [Input("batchInputs")]
        private List<ImmutableDictionary<string, object>>? _batchInputs;

        /// <summary>
        /// Specifies a list of items to be encoded in a single batch. If this parameter is set, the parameters 'value', 'transformation' and 'tweak' will be ignored. Each batch item within the list can specify these parameters instead.
        /// </summary>
        public List<ImmutableDictionary<string, object>> BatchInputs
        {
            get => _batchInputs ?? (_batchInputs = new List<ImmutableDictionary<string, object>>());
            set => _batchInputs = value;
        }

        [Input("batchResults")]
        private List<ImmutableDictionary<string, object>>? _batchResults;

        /// <summary>
        /// The result of encoding a batch.
        /// </summary>
        public List<ImmutableDictionary<string, object>> BatchResults
        {
            get => _batchResults ?? (_batchResults = new List<ImmutableDictionary<string, object>>());
            set => _batchResults = value;
        }

        /// <summary>
        /// The result of encoding a value.
        /// </summary>
        [Input("encodedValue")]
        public string? EncodedValue { get; set; }

        /// <summary>
        /// Path to where the back-end is mounted within Vault.
        /// </summary>
        [Input("path", required: true)]
        public string Path { get; set; } = null!;

        /// <summary>
        /// The name of the role.
        /// </summary>
        [Input("roleName", required: true)]
        public string RoleName { get; set; } = null!;

        /// <summary>
        /// The transformation to perform. If no value is provided and the role contains a single transformation, this value will be inferred from the role.
        /// </summary>
        [Input("transformation")]
        public string? Transformation { get; set; }

        /// <summary>
        /// The tweak value to use. Only applicable for FPE transformations
        /// </summary>
        [Input("tweak")]
        public string? Tweak { get; set; }

        /// <summary>
        /// The value in which to encode.
        /// </summary>
        [Input("value")]
        public string? Value { get; set; }

        public GetEncodeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEncodeResult
    {
        public readonly ImmutableArray<ImmutableDictionary<string, object>> BatchInputs;
        public readonly ImmutableArray<ImmutableDictionary<string, object>> BatchResults;
        public readonly string EncodedValue;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Path;
        public readonly string RoleName;
        public readonly string? Transformation;
        public readonly string? Tweak;
        public readonly string? Value;

        [OutputConstructor]
        private GetEncodeResult(
            ImmutableArray<ImmutableDictionary<string, object>> batchInputs,

            ImmutableArray<ImmutableDictionary<string, object>> batchResults,

            string encodedValue,

            string id,

            string path,

            string roleName,

            string? transformation,

            string? tweak,

            string? value)
        {
            BatchInputs = batchInputs;
            BatchResults = batchResults;
            EncodedValue = encodedValue;
            Id = id;
            Path = path;
            RoleName = roleName;
            Transformation = transformation;
            Tweak = tweak;
            Value = value;
        }
    }
}
