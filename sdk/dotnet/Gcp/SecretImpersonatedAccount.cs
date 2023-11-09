// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Gcp
{
    /// <summary>
    /// Creates a Impersonated Account in the [GCP Secrets Engine](https://www.vaultproject.io/docs/secrets/gcp/index.html) for Vault.
    /// 
    /// Each [impersonated account](https://www.vaultproject.io/docs/secrets/gcp/index.html#impersonated-accounts) is tied to a separately managed
    /// Service Account.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.IO;
    /// using System.Linq;
    /// using Pulumi;
    /// using Gcp = Pulumi.Gcp;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @this = new Gcp.ServiceAccount.Account("this", new()
    ///     {
    ///         AccountId = "my-awesome-account",
    ///     });
    /// 
    ///     var gcp = new Vault.Gcp.SecretBackend("gcp", new()
    ///     {
    ///         Path = "gcp",
    ///         Credentials = File.ReadAllText("credentials.json"),
    ///     });
    /// 
    ///     var impersonatedAccount = new Vault.Gcp.SecretImpersonatedAccount("impersonatedAccount", new()
    ///     {
    ///         Backend = gcp.Path,
    ///         ImpersonatedAccount = "this",
    ///         ServiceAccountEmail = @this.Email,
    ///         TokenScopes = new[]
    ///         {
    ///             "https://www.googleapis.com/auth/cloud-platform",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// A impersonated account can be imported using its Vault Path. For example, referencing the example above,
    /// 
    /// ```sh
    ///  $ pulumi import vault:gcp/secretImpersonatedAccount:SecretImpersonatedAccount impersonated_account gcp/impersonated-account/project_viewer
    /// ```
    /// </summary>
    [VaultResourceType("vault:gcp/secretImpersonatedAccount:SecretImpersonatedAccount")]
    public partial class SecretImpersonatedAccount : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Path where the GCP Secrets Engine is mounted
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// Name of the Impersonated Account to create
        /// </summary>
        [Output("impersonatedAccount")]
        public Output<string> ImpersonatedAccount { get; private set; } = null!;

        /// <summary>
        /// Target namespace. (requires Enterprise)
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Email of the GCP service account to impersonate.
        /// </summary>
        [Output("serviceAccountEmail")]
        public Output<string> ServiceAccountEmail { get; private set; } = null!;

        /// <summary>
        /// Project the service account belongs to.
        /// </summary>
        [Output("serviceAccountProject")]
        public Output<string> ServiceAccountProject { get; private set; } = null!;

        /// <summary>
        /// List of OAuth scopes to assign to access tokens generated under this impersonated account.
        /// </summary>
        [Output("tokenScopes")]
        public Output<ImmutableArray<string>> TokenScopes { get; private set; } = null!;


        /// <summary>
        /// Create a SecretImpersonatedAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretImpersonatedAccount(string name, SecretImpersonatedAccountArgs args, CustomResourceOptions? options = null)
            : base("vault:gcp/secretImpersonatedAccount:SecretImpersonatedAccount", name, args ?? new SecretImpersonatedAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretImpersonatedAccount(string name, Input<string> id, SecretImpersonatedAccountState? state = null, CustomResourceOptions? options = null)
            : base("vault:gcp/secretImpersonatedAccount:SecretImpersonatedAccount", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretImpersonatedAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretImpersonatedAccount Get(string name, Input<string> id, SecretImpersonatedAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretImpersonatedAccount(name, id, state, options);
        }
    }

    public sealed class SecretImpersonatedAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Path where the GCP Secrets Engine is mounted
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// Name of the Impersonated Account to create
        /// </summary>
        [Input("impersonatedAccount", required: true)]
        public Input<string> ImpersonatedAccount { get; set; } = null!;

        /// <summary>
        /// Target namespace. (requires Enterprise)
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Email of the GCP service account to impersonate.
        /// </summary>
        [Input("serviceAccountEmail", required: true)]
        public Input<string> ServiceAccountEmail { get; set; } = null!;

        [Input("tokenScopes")]
        private InputList<string>? _tokenScopes;

        /// <summary>
        /// List of OAuth scopes to assign to access tokens generated under this impersonated account.
        /// </summary>
        public InputList<string> TokenScopes
        {
            get => _tokenScopes ?? (_tokenScopes = new InputList<string>());
            set => _tokenScopes = value;
        }

        public SecretImpersonatedAccountArgs()
        {
        }
        public static new SecretImpersonatedAccountArgs Empty => new SecretImpersonatedAccountArgs();
    }

    public sealed class SecretImpersonatedAccountState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Path where the GCP Secrets Engine is mounted
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// Name of the Impersonated Account to create
        /// </summary>
        [Input("impersonatedAccount")]
        public Input<string>? ImpersonatedAccount { get; set; }

        /// <summary>
        /// Target namespace. (requires Enterprise)
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Email of the GCP service account to impersonate.
        /// </summary>
        [Input("serviceAccountEmail")]
        public Input<string>? ServiceAccountEmail { get; set; }

        /// <summary>
        /// Project the service account belongs to.
        /// </summary>
        [Input("serviceAccountProject")]
        public Input<string>? ServiceAccountProject { get; set; }

        [Input("tokenScopes")]
        private InputList<string>? _tokenScopes;

        /// <summary>
        /// List of OAuth scopes to assign to access tokens generated under this impersonated account.
        /// </summary>
        public InputList<string> TokenScopes
        {
            get => _tokenScopes ?? (_tokenScopes = new InputList<string>());
            set => _tokenScopes = value;
        }

        public SecretImpersonatedAccountState()
        {
        }
        public static new SecretImpersonatedAccountState Empty => new SecretImpersonatedAccountState();
    }
}
