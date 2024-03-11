// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.PkiSecret
{
    public static class GetBackendKeys
    {
        /// <summary>
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var pki = new Vault.Mount("pki", new()
        ///     {
        ///         Path = "pki",
        ///         Type = "pki",
        ///         Description = "PKI secret engine mount",
        ///     });
        /// 
        ///     var root = new Vault.PkiSecret.SecretBackendRootCert("root", new()
        ///     {
        ///         Backend = pki.Path,
        ///         Type = "internal",
        ///         CommonName = "example",
        ///         Ttl = "86400",
        ///         KeyName = "example",
        ///     });
        /// 
        ///     var example = Vault.PkiSecret.GetBackendKeys.Invoke(new()
        ///     {
        ///         Backend = root.Backend,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Task<GetBackendKeysResult> InvokeAsync(GetBackendKeysArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetBackendKeysResult>("vault:pkiSecret/getBackendKeys:getBackendKeys", args ?? new GetBackendKeysArgs(), options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// 
        /// &lt;!--Start PulumiCodeChooser --&gt;
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var pki = new Vault.Mount("pki", new()
        ///     {
        ///         Path = "pki",
        ///         Type = "pki",
        ///         Description = "PKI secret engine mount",
        ///     });
        /// 
        ///     var root = new Vault.PkiSecret.SecretBackendRootCert("root", new()
        ///     {
        ///         Backend = pki.Path,
        ///         Type = "internal",
        ///         CommonName = "example",
        ///         Ttl = "86400",
        ///         KeyName = "example",
        ///     });
        /// 
        ///     var example = Vault.PkiSecret.GetBackendKeys.Invoke(new()
        ///     {
        ///         Backend = root.Backend,
        ///     });
        /// 
        /// });
        /// ```
        /// &lt;!--End PulumiCodeChooser --&gt;
        /// </summary>
        public static Output<GetBackendKeysResult> Invoke(GetBackendKeysInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetBackendKeysResult>("vault:pkiSecret/getBackendKeys:getBackendKeys", args ?? new GetBackendKeysInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetBackendKeysArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The path to the PKI secret backend to
        /// read the keys from, with no leading or trailing `/`s.
        /// </summary>
        [Input("backend", required: true)]
        public string Backend { get; set; } = null!;

        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public string? Namespace { get; set; }

        public GetBackendKeysArgs()
        {
        }
        public static new GetBackendKeysArgs Empty => new GetBackendKeysArgs();
    }

    public sealed class GetBackendKeysInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The path to the PKI secret backend to
        /// read the keys from, with no leading or trailing `/`s.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// The namespace of the target resource.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public GetBackendKeysInvokeArgs()
        {
        }
        public static new GetBackendKeysInvokeArgs Empty => new GetBackendKeysInvokeArgs();
    }


    [OutputType]
    public sealed class GetBackendKeysResult
    {
        public readonly string Backend;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Map of key strings read from Vault.
        /// </summary>
        public readonly ImmutableDictionary<string, object> KeyInfo;
        /// <summary>
        /// JSON-encoded key data read from Vault.
        /// </summary>
        public readonly string KeyInfoJson;
        /// <summary>
        /// Keys used under the backend path.
        /// </summary>
        public readonly ImmutableArray<string> Keys;
        public readonly string? Namespace;

        [OutputConstructor]
        private GetBackendKeysResult(
            string backend,

            string id,

            ImmutableDictionary<string, object> keyInfo,

            string keyInfoJson,

            ImmutableArray<string> keys,

            string? @namespace)
        {
            Backend = backend;
            Id = id;
            KeyInfo = keyInfo;
            KeyInfoJson = keyInfoJson;
            Keys = keys;
            Namespace = @namespace;
        }
    }
}
