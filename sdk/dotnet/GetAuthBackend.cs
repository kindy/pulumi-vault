// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault
{
    public static class GetAuthBackend
    {
        /// <summary>
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
        ///     var example = Vault.GetAuthBackend.Invoke(new()
        ///     {
        ///         Path = "userpass",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetAuthBackendResult> InvokeAsync(GetAuthBackendArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAuthBackendResult>("vault:index/getAuthBackend:getAuthBackend", args ?? new GetAuthBackendArgs(), options.WithDefaults());

        /// <summary>
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
        ///     var example = Vault.GetAuthBackend.Invoke(new()
        ///     {
        ///         Path = "userpass",
        ///     });
        /// 
        /// });
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetAuthBackendResult> Invoke(GetAuthBackendInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAuthBackendResult>("vault:index/getAuthBackend:getAuthBackend", args ?? new GetAuthBackendInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAuthBackendArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public string? Namespace { get; set; }

        /// <summary>
        /// The auth backend mount point.
        /// </summary>
        [Input("path", required: true)]
        public string Path { get; set; } = null!;

        public GetAuthBackendArgs()
        {
        }
        public static new GetAuthBackendArgs Empty => new GetAuthBackendArgs();
    }

    public sealed class GetAuthBackendInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The auth backend mount point.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        public GetAuthBackendInvokeArgs()
        {
        }
        public static new GetAuthBackendInvokeArgs Empty => new GetAuthBackendInvokeArgs();
    }


    [OutputType]
    public sealed class GetAuthBackendResult
    {
        /// <summary>
        /// The accessor for this auth method.
        /// </summary>
        public readonly string Accessor;
        /// <summary>
        /// The default lease duration in seconds.
        /// </summary>
        public readonly int DefaultLeaseTtlSeconds;
        /// <summary>
        /// A description of the auth method.
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Specifies whether to show this mount in the UI-specific listing endpoint.
        /// </summary>
        public readonly string ListingVisibility;
        /// <summary>
        /// Specifies if the auth method is local only.
        /// </summary>
        public readonly bool Local;
        /// <summary>
        /// The maximum lease duration in seconds.
        /// </summary>
        public readonly int MaxLeaseTtlSeconds;
        public readonly string? Namespace;
        public readonly string Path;
        /// <summary>
        /// The name of the auth method type.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GetAuthBackendResult(
            string accessor,

            int defaultLeaseTtlSeconds,

            string description,

            string id,

            string listingVisibility,

            bool local,

            int maxLeaseTtlSeconds,

            string? @namespace,

            string path,

            string type)
        {
            Accessor = accessor;
            DefaultLeaseTtlSeconds = defaultLeaseTtlSeconds;
            Description = description;
            Id = id;
            ListingVisibility = listingVisibility;
            Local = local;
            MaxLeaseTtlSeconds = maxLeaseTtlSeconds;
            Namespace = @namespace;
            Path = path;
            Type = type;
        }
    }
}
