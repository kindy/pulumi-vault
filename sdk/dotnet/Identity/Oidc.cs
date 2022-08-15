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
    /// Configure the [Identity Tokens Backend](https://www.vaultproject.io/docs/secrets/identity/index.html#identity-tokens).
    /// 
    /// The Identity secrets engine is the identity management solution for Vault. It internally maintains
    /// the clients who are recognized by Vault.
    /// 
    /// &gt; **NOTE:** Each Vault server may only have one Identity Tokens Backend configuration. Multiple configurations of the resource against the same Vault server will cause a perpetual difference.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var server = new Vault.Identity.Oidc("server", new()
    ///     {
    ///         Issuer = "https://www.acme.com",
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [VaultResourceType("vault:identity/oidc:Oidc")]
    public partial class Oidc : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Issuer URL to be used in the iss claim of the token. If not set, Vault's
        /// `api_addr` will be used. The issuer is a case sensitive URL using the https scheme that contains
        /// scheme, host, and optionally, port number and path components, but no query or fragment
        /// components.
        /// </summary>
        [Output("issuer")]
        public Output<string> Issuer { get; private set; } = null!;


        /// <summary>
        /// Create a Oidc resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Oidc(string name, OidcArgs? args = null, CustomResourceOptions? options = null)
            : base("vault:identity/oidc:Oidc", name, args ?? new OidcArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Oidc(string name, Input<string> id, OidcState? state = null, CustomResourceOptions? options = null)
            : base("vault:identity/oidc:Oidc", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Oidc resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Oidc Get(string name, Input<string> id, OidcState? state = null, CustomResourceOptions? options = null)
        {
            return new Oidc(name, id, state, options);
        }
    }

    public sealed class OidcArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Issuer URL to be used in the iss claim of the token. If not set, Vault's
        /// `api_addr` will be used. The issuer is a case sensitive URL using the https scheme that contains
        /// scheme, host, and optionally, port number and path components, but no query or fragment
        /// components.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        public OidcArgs()
        {
        }
        public static new OidcArgs Empty => new OidcArgs();
    }

    public sealed class OidcState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Issuer URL to be used in the iss claim of the token. If not set, Vault's
        /// `api_addr` will be used. The issuer is a case sensitive URL using the https scheme that contains
        /// scheme, host, and optionally, port number and path components, but no query or fragment
        /// components.
        /// </summary>
        [Input("issuer")]
        public Input<string>? Issuer { get; set; }

        public OidcState()
        {
        }
        public static new OidcState Empty => new OidcState();
    }
}
