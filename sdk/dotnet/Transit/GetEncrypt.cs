// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Transit
{
    public static class GetEncrypt
    {
        /// <summary>
        /// This is a data source which can be used to encrypt plaintext using a Vault Transit key.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Vault = Pulumi.Vault;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var testMount = new Vault.Mount("testMount", new Vault.MountArgs
        ///         {
        ///             Description = "This is an example mount",
        ///             Path = "transit",
        ///             Type = "transit",
        ///         });
        ///         var testSecretBackendKey = new Vault.Transit.SecretBackendKey("testSecretBackendKey", new Vault.Transit.SecretBackendKeyArgs
        ///         {
        ///             Backend = testMount.Path,
        ///         });
        ///         var testEncrypt = Output.Tuple(testMount.Path, testSecretBackendKey.Name).Apply(values =&gt;
        ///         {
        ///             var path = values.Item1;
        ///             var name = values.Item2;
        ///             return Vault.Transit.GetEncrypt.InvokeAsync(new Vault.Transit.GetEncryptArgs
        ///             {
        ///                 Backend = path,
        ///                 Key = name,
        ///                 Plaintext = "foobar",
        ///             });
        ///         });
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetEncryptResult> InvokeAsync(GetEncryptArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetEncryptResult>("vault:transit/getEncrypt:getEncrypt", args ?? new GetEncryptArgs(), options.WithVersion());
    }


    public sealed class GetEncryptArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The path the transit secret backend is mounted at, with no leading or trailing `/`.
        /// </summary>
        [Input("backend", required: true)]
        public string Backend { get; set; } = null!;

        /// <summary>
        /// Context for key derivation. This is required if key derivation is enabled for this key.
        /// </summary>
        [Input("context")]
        public string? Context { get; set; }

        /// <summary>
        /// Specifies the name of the transit key to encrypt against.
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        /// <summary>
        /// The version of the key to use for encryption. If not set, uses the latest version. Must be greater than or equal to the key's `min_encryption_version`, if set.
        /// </summary>
        [Input("keyVersion")]
        public int? KeyVersion { get; set; }

        /// <summary>
        /// Plaintext to be encoded.
        /// </summary>
        [Input("plaintext", required: true)]
        public string Plaintext { get; set; } = null!;

        public GetEncryptArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetEncryptResult
    {
        public readonly string Backend;
        /// <summary>
        /// Encrypted ciphertext returned from Vault
        /// </summary>
        public readonly string Ciphertext;
        public readonly string? Context;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Key;
        public readonly int? KeyVersion;
        public readonly string Plaintext;

        [OutputConstructor]
        private GetEncryptResult(
            string backend,

            string ciphertext,

            string? context,

            string id,

            string key,

            int? keyVersion,

            string plaintext)
        {
            Backend = backend;
            Ciphertext = ciphertext;
            Context = context;
            Id = id;
            Key = key;
            KeyVersion = keyVersion;
            Plaintext = plaintext;
        }
    }
}