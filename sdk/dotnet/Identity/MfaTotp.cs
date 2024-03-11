// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Identity
{
    /// <summary>
    /// Resource for configuring the totp MFA method.
    /// 
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
    ///     var example = new Vault.Identity.MfaTotp("example", new()
    ///     {
    ///         Issuer = "issuer1",
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Resource can be imported using its `uuid` field, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:identity/mfaTotp:MfaTotp example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec
    /// ```
    /// </summary>
    [VaultResourceType("vault:identity/mfaTotp:MfaTotp")]
    public partial class MfaTotp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the hashing algorithm used to generate the TOTP code. Options include SHA1, SHA256, SHA512.
        /// </summary>
        [Output("algorithm")]
        public Output<string?> Algorithm { get; private set; } = null!;

        /// <summary>
        /// The number of digits in the generated TOTP token. This value can either be 6 or 8
        /// </summary>
        [Output("digits")]
        public Output<int?> Digits { get; private set; } = null!;

        /// <summary>
        /// The name of the key's issuing organization.
        /// </summary>
        [Output("issuer")]
        public Output<string> Issuer { get; private set; } = null!;

        /// <summary>
        /// Specifies the size in bytes of the generated key.
        /// </summary>
        [Output("keySize")]
        public Output<int?> KeySize { get; private set; } = null!;

        /// <summary>
        /// The maximum number of consecutive failed validation attempts allowed.
        /// </summary>
        [Output("maxValidationAttempts")]
        public Output<int?> MaxValidationAttempts { get; private set; } = null!;

        /// <summary>
        /// Method ID.
        /// </summary>
        [Output("methodId")]
        public Output<string> MethodId { get; private set; } = null!;

        /// <summary>
        /// Mount accessor.
        /// </summary>
        [Output("mountAccessor")]
        public Output<string> MountAccessor { get; private set; } = null!;

        /// <summary>
        /// Method name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Target namespace. (requires Enterprise)
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Method's namespace ID.
        /// </summary>
        [Output("namespaceId")]
        public Output<string> NamespaceId { get; private set; } = null!;

        /// <summary>
        /// Method's namespace path.
        /// </summary>
        [Output("namespacePath")]
        public Output<string> NamespacePath { get; private set; } = null!;

        /// <summary>
        /// The length of time in seconds used to generate a counter for the TOTP token calculation.
        /// </summary>
        [Output("period")]
        public Output<int?> Period { get; private set; } = null!;

        /// <summary>
        /// The pixel size of the generated square QR code.
        /// </summary>
        [Output("qrSize")]
        public Output<int> QrSize { get; private set; } = null!;

        /// <summary>
        /// The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1.
        /// </summary>
        [Output("skew")]
        public Output<int?> Skew { get; private set; } = null!;

        /// <summary>
        /// MFA type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Resource UUID.
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;


        /// <summary>
        /// Create a MfaTotp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MfaTotp(string name, MfaTotpArgs args, CustomResourceOptions? options = null)
            : base("vault:identity/mfaTotp:MfaTotp", name, args ?? new MfaTotpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MfaTotp(string name, Input<string> id, MfaTotpState? state = null, CustomResourceOptions? options = null)
            : base("vault:identity/mfaTotp:MfaTotp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MfaTotp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MfaTotp Get(string name, Input<string> id, MfaTotpState? state = null, CustomResourceOptions? options = null)
        {
            return new MfaTotp(name, id, state, options);
        }
    }

    public sealed class MfaTotpArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the hashing algorithm used to generate the TOTP code. Options include SHA1, SHA256, SHA512.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// The number of digits in the generated TOTP token. This value can either be 6 or 8
        /// </summary>
        [Input("digits")]
        public Input<int>? Digits { get; set; }

        /// <summary>
        /// The name of the key's issuing organization.
        /// </summary>
        [Input("issuer", required: true)]
        public Input<string> Issuer { get; set; } = null!;

        /// <summary>
        /// Specifies the size in bytes of the generated key.
        /// </summary>
        [Input("keySize")]
        public Input<int>? KeySize { get; set; }

        /// <summary>
        /// The maximum number of consecutive failed validation attempts allowed.
        /// </summary>
        [Input("maxValidationAttempts")]
        public Input<int>? MaxValidationAttempts { get; set; }

        /// <summary>
        /// Target namespace. (requires Enterprise)
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// The length of time in seconds used to generate a counter for the TOTP token calculation.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The pixel size of the generated square QR code.
        /// </summary>
        [Input("qrSize")]
        public Input<int>? QrSize { get; set; }

        /// <summary>
        /// The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1.
        /// </summary>
        [Input("skew")]
        public Input<int>? Skew { get; set; }

        public MfaTotpArgs()
        {
        }
        public static new MfaTotpArgs Empty => new MfaTotpArgs();
    }

    public sealed class MfaTotpState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the hashing algorithm used to generate the TOTP code. Options include SHA1, SHA256, SHA512.
        /// </summary>
        [Input("algorithm")]
        public Input<string>? Algorithm { get; set; }

        /// <summary>
        /// The number of digits in the generated TOTP token. This value can either be 6 or 8
        /// </summary>
        [Input("digits")]
        public Input<int>? Digits { get; set; }

        /// <summary>
        /// The name of the key's issuing organization.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        /// <summary>
        /// Specifies the size in bytes of the generated key.
        /// </summary>
        [Input("keySize")]
        public Input<int>? KeySize { get; set; }

        /// <summary>
        /// The maximum number of consecutive failed validation attempts allowed.
        /// </summary>
        [Input("maxValidationAttempts")]
        public Input<int>? MaxValidationAttempts { get; set; }

        /// <summary>
        /// Method ID.
        /// </summary>
        [Input("methodId")]
        public Input<string>? MethodId { get; set; }

        /// <summary>
        /// Mount accessor.
        /// </summary>
        [Input("mountAccessor")]
        public Input<string>? MountAccessor { get; set; }

        /// <summary>
        /// Method name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Target namespace. (requires Enterprise)
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Method's namespace ID.
        /// </summary>
        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        /// <summary>
        /// Method's namespace path.
        /// </summary>
        [Input("namespacePath")]
        public Input<string>? NamespacePath { get; set; }

        /// <summary>
        /// The length of time in seconds used to generate a counter for the TOTP token calculation.
        /// </summary>
        [Input("period")]
        public Input<int>? Period { get; set; }

        /// <summary>
        /// The pixel size of the generated square QR code.
        /// </summary>
        [Input("qrSize")]
        public Input<int>? QrSize { get; set; }

        /// <summary>
        /// The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1.
        /// </summary>
        [Input("skew")]
        public Input<int>? Skew { get; set; }

        /// <summary>
        /// MFA type.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Resource UUID.
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        public MfaTotpState()
        {
        }
        public static new MfaTotpState Empty => new MfaTotpState();
    }
}
