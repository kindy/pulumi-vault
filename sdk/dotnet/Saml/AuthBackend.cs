// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Saml
{
    /// <summary>
    /// Manages a SAML Auth mount in a Vault server. See the [Vault
    /// documentation](https://www.vaultproject.io/docs/auth/saml/) for more
    /// information.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var test = new Vault.Saml.AuthBackend("test", new()
    ///     {
    ///         AcsUrls = new[]
    ///         {
    ///             "https://my.vault.primary/v1/auth/saml/callback",
    ///         },
    ///         DefaultRole = "admin",
    ///         EntityId = "https://my.vault/v1/auth/saml",
    ///         IdpMetadataUrl = "https://company.okta.com/app/abc123eb9xnIfzlaf697/sso/saml/metadata",
    ///         Path = "saml",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// SAML authentication mounts can be imported using the `path`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import vault:saml/authBackend:AuthBackend example saml
    /// ```
    /// </summary>
    [VaultResourceType("vault:saml/authBackend:AuthBackend")]
    public partial class AuthBackend : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The well-formatted URLs of your Assertion Consumer Service (ACS)
        /// that should receive a response from the identity provider.
        /// </summary>
        [Output("acsUrls")]
        public Output<ImmutableArray<string>> AcsUrls { get; private set; } = null!;

        /// <summary>
        /// The role to use if no role is provided during login.
        /// </summary>
        [Output("defaultRole")]
        public Output<string?> DefaultRole { get; private set; } = null!;

        /// <summary>
        /// If set to `true`, opts out of mount migration on path updates.
        /// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        /// </summary>
        [Output("disableRemount")]
        public Output<bool?> DisableRemount { get; private set; } = null!;

        /// <summary>
        /// The entity ID of the SAML authentication service provider.
        /// </summary>
        [Output("entityId")]
        public Output<string> EntityId { get; private set; } = null!;

        /// <summary>
        /// The PEM encoded certificate of the identity provider. Mutually exclusive
        /// with `idp_metadata_url`.
        /// </summary>
        [Output("idpCert")]
        public Output<string?> IdpCert { get; private set; } = null!;

        /// <summary>
        /// The entity ID of the identity provider. Mutually exclusive with
        /// `idp_metadata_url`.
        /// </summary>
        [Output("idpEntityId")]
        public Output<string?> IdpEntityId { get; private set; } = null!;

        /// <summary>
        /// The metadata URL of the identity provider.
        /// </summary>
        [Output("idpMetadataUrl")]
        public Output<string?> IdpMetadataUrl { get; private set; } = null!;

        /// <summary>
        /// The SSO URL of the identity provider. Mutually exclusive with 
        /// `idp_metadata_url`.
        /// </summary>
        [Output("idpSsoUrl")]
        public Output<string?> IdpSsoUrl { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Path where the auth backend will be mounted. Defaults to `auth/saml`
        /// if not specified.
        /// </summary>
        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;

        /// <summary>
        /// If set to `true`, logs additional, potentially sensitive
        /// information during the SAML exchange according to the current logging level. Not
        /// recommended for production.
        /// </summary>
        [Output("verboseLogging")]
        public Output<bool> VerboseLogging { get; private set; } = null!;


        /// <summary>
        /// Create a AuthBackend resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AuthBackend(string name, AuthBackendArgs args, CustomResourceOptions? options = null)
            : base("vault:saml/authBackend:AuthBackend", name, args ?? new AuthBackendArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AuthBackend(string name, Input<string> id, AuthBackendState? state = null, CustomResourceOptions? options = null)
            : base("vault:saml/authBackend:AuthBackend", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AuthBackend resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AuthBackend Get(string name, Input<string> id, AuthBackendState? state = null, CustomResourceOptions? options = null)
        {
            return new AuthBackend(name, id, state, options);
        }
    }

    public sealed class AuthBackendArgs : global::Pulumi.ResourceArgs
    {
        [Input("acsUrls", required: true)]
        private InputList<string>? _acsUrls;

        /// <summary>
        /// The well-formatted URLs of your Assertion Consumer Service (ACS)
        /// that should receive a response from the identity provider.
        /// </summary>
        public InputList<string> AcsUrls
        {
            get => _acsUrls ?? (_acsUrls = new InputList<string>());
            set => _acsUrls = value;
        }

        /// <summary>
        /// The role to use if no role is provided during login.
        /// </summary>
        [Input("defaultRole")]
        public Input<string>? DefaultRole { get; set; }

        /// <summary>
        /// If set to `true`, opts out of mount migration on path updates.
        /// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        /// </summary>
        [Input("disableRemount")]
        public Input<bool>? DisableRemount { get; set; }

        /// <summary>
        /// The entity ID of the SAML authentication service provider.
        /// </summary>
        [Input("entityId", required: true)]
        public Input<string> EntityId { get; set; } = null!;

        /// <summary>
        /// The PEM encoded certificate of the identity provider. Mutually exclusive
        /// with `idp_metadata_url`.
        /// </summary>
        [Input("idpCert")]
        public Input<string>? IdpCert { get; set; }

        /// <summary>
        /// The entity ID of the identity provider. Mutually exclusive with
        /// `idp_metadata_url`.
        /// </summary>
        [Input("idpEntityId")]
        public Input<string>? IdpEntityId { get; set; }

        /// <summary>
        /// The metadata URL of the identity provider.
        /// </summary>
        [Input("idpMetadataUrl")]
        public Input<string>? IdpMetadataUrl { get; set; }

        /// <summary>
        /// The SSO URL of the identity provider. Mutually exclusive with 
        /// `idp_metadata_url`.
        /// </summary>
        [Input("idpSsoUrl")]
        public Input<string>? IdpSsoUrl { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Path where the auth backend will be mounted. Defaults to `auth/saml`
        /// if not specified.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// If set to `true`, logs additional, potentially sensitive
        /// information during the SAML exchange according to the current logging level. Not
        /// recommended for production.
        /// </summary>
        [Input("verboseLogging")]
        public Input<bool>? VerboseLogging { get; set; }

        public AuthBackendArgs()
        {
        }
        public static new AuthBackendArgs Empty => new AuthBackendArgs();
    }

    public sealed class AuthBackendState : global::Pulumi.ResourceArgs
    {
        [Input("acsUrls")]
        private InputList<string>? _acsUrls;

        /// <summary>
        /// The well-formatted URLs of your Assertion Consumer Service (ACS)
        /// that should receive a response from the identity provider.
        /// </summary>
        public InputList<string> AcsUrls
        {
            get => _acsUrls ?? (_acsUrls = new InputList<string>());
            set => _acsUrls = value;
        }

        /// <summary>
        /// The role to use if no role is provided during login.
        /// </summary>
        [Input("defaultRole")]
        public Input<string>? DefaultRole { get; set; }

        /// <summary>
        /// If set to `true`, opts out of mount migration on path updates.
        /// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        /// </summary>
        [Input("disableRemount")]
        public Input<bool>? DisableRemount { get; set; }

        /// <summary>
        /// The entity ID of the SAML authentication service provider.
        /// </summary>
        [Input("entityId")]
        public Input<string>? EntityId { get; set; }

        /// <summary>
        /// The PEM encoded certificate of the identity provider. Mutually exclusive
        /// with `idp_metadata_url`.
        /// </summary>
        [Input("idpCert")]
        public Input<string>? IdpCert { get; set; }

        /// <summary>
        /// The entity ID of the identity provider. Mutually exclusive with
        /// `idp_metadata_url`.
        /// </summary>
        [Input("idpEntityId")]
        public Input<string>? IdpEntityId { get; set; }

        /// <summary>
        /// The metadata URL of the identity provider.
        /// </summary>
        [Input("idpMetadataUrl")]
        public Input<string>? IdpMetadataUrl { get; set; }

        /// <summary>
        /// The SSO URL of the identity provider. Mutually exclusive with 
        /// `idp_metadata_url`.
        /// </summary>
        [Input("idpSsoUrl")]
        public Input<string>? IdpSsoUrl { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Path where the auth backend will be mounted. Defaults to `auth/saml`
        /// if not specified.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// If set to `true`, logs additional, potentially sensitive
        /// information during the SAML exchange according to the current logging level. Not
        /// recommended for production.
        /// </summary>
        [Input("verboseLogging")]
        public Input<bool>? VerboseLogging { get; set; }

        public AuthBackendState()
        {
        }
        public static new AuthBackendState Empty => new AuthBackendState();
    }
}
