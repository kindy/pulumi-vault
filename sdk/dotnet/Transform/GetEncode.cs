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
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var transform = new Vault.Mount("transform", new()
        ///     {
        ///         Path = "transform",
        ///         Type = "transform",
        ///     });
        /// 
        ///     var ccn_fpe = new Vault.Transform.Transformation("ccn-fpe", new()
        ///     {
        ///         Path = transform.Path,
        ///         Type = "fpe",
        ///         Template = "builtin/creditcardnumber",
        ///         TweakSource = "internal",
        ///         AllowedRoles = new[]
        ///         {
        ///             "payments",
        ///         },
        ///     });
        /// 
        ///     var payments = new Vault.Transform.Role("payments", new()
        ///     {
        ///         Path = ccn_fpe.Path,
        ///         Transformations = new[]
        ///         {
        ///             "ccn-fpe",
        ///         },
        ///     });
        /// 
        ///     var test = Vault.Transform.GetEncode.Invoke(new()
        ///     {
        ///         Path = payments.Path,
        ///         RoleName = "payments",
        ///         BatchInputs = new[]
        ///         {
        ///             
        ///             {
        ///                 { "value", "1111-2222-3333-4444" },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEncodeResult> InvokeAsync(GetEncodeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetEncodeResult>("vault:transform/getEncode:getEncode", args ?? new GetEncodeArgs(), options.WithDefaults());

        /// <summary>
        /// This data source supports the "/transform/encode/{role_name}" Vault endpoint.
        /// 
        /// It encodes the provided value using a named role.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var transform = new Vault.Mount("transform", new()
        ///     {
        ///         Path = "transform",
        ///         Type = "transform",
        ///     });
        /// 
        ///     var ccn_fpe = new Vault.Transform.Transformation("ccn-fpe", new()
        ///     {
        ///         Path = transform.Path,
        ///         Type = "fpe",
        ///         Template = "builtin/creditcardnumber",
        ///         TweakSource = "internal",
        ///         AllowedRoles = new[]
        ///         {
        ///             "payments",
        ///         },
        ///     });
        /// 
        ///     var payments = new Vault.Transform.Role("payments", new()
        ///     {
        ///         Path = ccn_fpe.Path,
        ///         Transformations = new[]
        ///         {
        ///             "ccn-fpe",
        ///         },
        ///     });
        /// 
        ///     var test = Vault.Transform.GetEncode.Invoke(new()
        ///     {
        ///         Path = payments.Path,
        ///         RoleName = "payments",
        ///         BatchInputs = new[]
        ///         {
        ///             
        ///             {
        ///                 { "value", "1111-2222-3333-4444" },
        ///             },
        ///         },
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetEncodeResult> Invoke(GetEncodeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetEncodeResult>("vault:transform/getEncode:getEncode", args ?? new GetEncodeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetEncodeArgs : global::Pulumi.InvokeArgs
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
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public string? Namespace { get; set; }

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
        public static new GetEncodeArgs Empty => new GetEncodeArgs();
    }

    public sealed class GetEncodeInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("batchInputs")]
        private InputList<ImmutableDictionary<string, object>>? _batchInputs;

        /// <summary>
        /// Specifies a list of items to be encoded in a single batch. If this parameter is set, the parameters 'value', 'transformation' and 'tweak' will be ignored. Each batch item within the list can specify these parameters instead.
        /// </summary>
        public InputList<ImmutableDictionary<string, object>> BatchInputs
        {
            get => _batchInputs ?? (_batchInputs = new InputList<ImmutableDictionary<string, object>>());
            set => _batchInputs = value;
        }

        [Input("batchResults")]
        private InputList<ImmutableDictionary<string, object>>? _batchResults;

        /// <summary>
        /// The result of encoding a batch.
        /// </summary>
        public InputList<ImmutableDictionary<string, object>> BatchResults
        {
            get => _batchResults ?? (_batchResults = new InputList<ImmutableDictionary<string, object>>());
            set => _batchResults = value;
        }

        /// <summary>
        /// The result of encoding a value.
        /// </summary>
        [Input("encodedValue")]
        public Input<string>? EncodedValue { get; set; }

        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Path to where the back-end is mounted within Vault.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// The name of the role.
        /// </summary>
        [Input("roleName", required: true)]
        public Input<string> RoleName { get; set; } = null!;

        /// <summary>
        /// The transformation to perform. If no value is provided and the role contains a single transformation, this value will be inferred from the role.
        /// </summary>
        [Input("transformation")]
        public Input<string>? Transformation { get; set; }

        /// <summary>
        /// The tweak value to use. Only applicable for FPE transformations
        /// </summary>
        [Input("tweak")]
        public Input<string>? Tweak { get; set; }

        /// <summary>
        /// The value in which to encode.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public GetEncodeInvokeArgs()
        {
        }
        public static new GetEncodeInvokeArgs Empty => new GetEncodeInvokeArgs();
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
        public readonly string? Namespace;
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

            string? @namespace,

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
            Namespace = @namespace;
            Path = path;
            RoleName = roleName;
            Transformation = transformation;
            Tweak = tweak;
            Value = value;
        }
    }
}
