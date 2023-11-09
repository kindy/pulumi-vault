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
    /// Creates a Static Account in the [GCP Secrets Engine](https://www.vaultproject.io/docs/secrets/gcp/index.html) for Vault.
    /// 
    /// Each [static account](https://www.vaultproject.io/docs/secrets/gcp/index.html#static-accounts) is tied to a separately managed
    /// Service Account, and can have one or more [bindings](https://www.vaultproject.io/docs/secrets/gcp/index.html#bindings) associated with it.
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
    ///     var staticAccount = new Vault.Gcp.SecretStaticAccount("staticAccount", new()
    ///     {
    ///         Backend = gcp.Path,
    ///         StaticAccount = "project_viewer",
    ///         SecretType = "access_token",
    ///         TokenScopes = new[]
    ///         {
    ///             "https://www.googleapis.com/auth/cloud-platform",
    ///         },
    ///         ServiceAccountEmail = @this.Email,
    ///         Bindings = new[]
    ///         {
    ///             new Vault.Gcp.Inputs.SecretStaticAccountBindingArgs
    ///             {
    ///                 Resource = @this.Project.Apply(project =&gt; $"//cloudresourcemanager.googleapis.com/projects/{project}"),
    ///                 Roles = new[]
    ///                 {
    ///                     "roles/viewer",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// A static account can be imported using its Vault Path. For example, referencing the example above,
    /// 
    /// ```sh
    ///  $ pulumi import vault:gcp/secretStaticAccount:SecretStaticAccount static_account gcp/static-account/project_viewer
    /// ```
    /// </summary>
    [VaultResourceType("vault:gcp/secretStaticAccount:SecretStaticAccount")]
    public partial class SecretStaticAccount : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Path where the GCP Secrets Engine is mounted
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// Bindings to create for this static account. This can be specified multiple times for multiple bindings. Structure is documented below.
        /// </summary>
        [Output("bindings")]
        public Output<ImmutableArray<Outputs.SecretStaticAccountBinding>> Bindings { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Type of secret generated for this static account. Accepted values: `access_token`, `service_account_key`. Defaults to `access_token`.
        /// </summary>
        [Output("secretType")]
        public Output<string> SecretType { get; private set; } = null!;

        /// <summary>
        /// Email of the GCP service account to manage.
        /// </summary>
        [Output("serviceAccountEmail")]
        public Output<string> ServiceAccountEmail { get; private set; } = null!;

        /// <summary>
        /// Project the service account belongs to.
        /// </summary>
        [Output("serviceAccountProject")]
        public Output<string> ServiceAccountProject { get; private set; } = null!;

        /// <summary>
        /// Name of the Static Account to create
        /// </summary>
        [Output("staticAccount")]
        public Output<string> StaticAccount { get; private set; } = null!;

        /// <summary>
        /// List of OAuth scopes to assign to `access_token` secrets generated under this static account (`access_token` static accounts only).
        /// </summary>
        [Output("tokenScopes")]
        public Output<ImmutableArray<string>> TokenScopes { get; private set; } = null!;


        /// <summary>
        /// Create a SecretStaticAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretStaticAccount(string name, SecretStaticAccountArgs args, CustomResourceOptions? options = null)
            : base("vault:gcp/secretStaticAccount:SecretStaticAccount", name, args ?? new SecretStaticAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretStaticAccount(string name, Input<string> id, SecretStaticAccountState? state = null, CustomResourceOptions? options = null)
            : base("vault:gcp/secretStaticAccount:SecretStaticAccount", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretStaticAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretStaticAccount Get(string name, Input<string> id, SecretStaticAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretStaticAccount(name, id, state, options);
        }
    }

    public sealed class SecretStaticAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Path where the GCP Secrets Engine is mounted
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        [Input("bindings")]
        private InputList<Inputs.SecretStaticAccountBindingArgs>? _bindings;

        /// <summary>
        /// Bindings to create for this static account. This can be specified multiple times for multiple bindings. Structure is documented below.
        /// </summary>
        public InputList<Inputs.SecretStaticAccountBindingArgs> Bindings
        {
            get => _bindings ?? (_bindings = new InputList<Inputs.SecretStaticAccountBindingArgs>());
            set => _bindings = value;
        }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Type of secret generated for this static account. Accepted values: `access_token`, `service_account_key`. Defaults to `access_token`.
        /// </summary>
        [Input("secretType")]
        public Input<string>? SecretType { get; set; }

        /// <summary>
        /// Email of the GCP service account to manage.
        /// </summary>
        [Input("serviceAccountEmail", required: true)]
        public Input<string> ServiceAccountEmail { get; set; } = null!;

        /// <summary>
        /// Name of the Static Account to create
        /// </summary>
        [Input("staticAccount", required: true)]
        public Input<string> StaticAccount { get; set; } = null!;

        [Input("tokenScopes")]
        private InputList<string>? _tokenScopes;

        /// <summary>
        /// List of OAuth scopes to assign to `access_token` secrets generated under this static account (`access_token` static accounts only).
        /// </summary>
        public InputList<string> TokenScopes
        {
            get => _tokenScopes ?? (_tokenScopes = new InputList<string>());
            set => _tokenScopes = value;
        }

        public SecretStaticAccountArgs()
        {
        }
        public static new SecretStaticAccountArgs Empty => new SecretStaticAccountArgs();
    }

    public sealed class SecretStaticAccountState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Path where the GCP Secrets Engine is mounted
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        [Input("bindings")]
        private InputList<Inputs.SecretStaticAccountBindingGetArgs>? _bindings;

        /// <summary>
        /// Bindings to create for this static account. This can be specified multiple times for multiple bindings. Structure is documented below.
        /// </summary>
        public InputList<Inputs.SecretStaticAccountBindingGetArgs> Bindings
        {
            get => _bindings ?? (_bindings = new InputList<Inputs.SecretStaticAccountBindingGetArgs>());
            set => _bindings = value;
        }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Type of secret generated for this static account. Accepted values: `access_token`, `service_account_key`. Defaults to `access_token`.
        /// </summary>
        [Input("secretType")]
        public Input<string>? SecretType { get; set; }

        /// <summary>
        /// Email of the GCP service account to manage.
        /// </summary>
        [Input("serviceAccountEmail")]
        public Input<string>? ServiceAccountEmail { get; set; }

        /// <summary>
        /// Project the service account belongs to.
        /// </summary>
        [Input("serviceAccountProject")]
        public Input<string>? ServiceAccountProject { get; set; }

        /// <summary>
        /// Name of the Static Account to create
        /// </summary>
        [Input("staticAccount")]
        public Input<string>? StaticAccount { get; set; }

        [Input("tokenScopes")]
        private InputList<string>? _tokenScopes;

        /// <summary>
        /// List of OAuth scopes to assign to `access_token` secrets generated under this static account (`access_token` static accounts only).
        /// </summary>
        public InputList<string> TokenScopes
        {
            get => _tokenScopes ?? (_tokenScopes = new InputList<string>());
            set => _tokenScopes = value;
        }

        public SecretStaticAccountState()
        {
        }
        public static new SecretStaticAccountState Empty => new SecretStaticAccountState();
    }
}
