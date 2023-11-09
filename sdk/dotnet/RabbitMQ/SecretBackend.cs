// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.RabbitMQ
{
    /// <summary>
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
    ///     var rabbitmq = new Vault.RabbitMQ.SecretBackend("rabbitmq", new()
    ///     {
    ///         ConnectionUri = "https://.....",
    ///         Password = "password",
    ///         Username = "user",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// RabbitMQ secret backends can be imported using the `path`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import vault:rabbitMq/secretBackend:SecretBackend rabbitmq rabbitmq
    /// ```
    /// </summary>
    [VaultResourceType("vault:rabbitMq/secretBackend:SecretBackend")]
    public partial class SecretBackend : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the RabbitMQ connection URI.
        /// </summary>
        [Output("connectionUri")]
        public Output<string> ConnectionUri { get; private set; } = null!;

        /// <summary>
        /// The default TTL for credentials
        /// issued by this backend.
        /// </summary>
        [Output("defaultLeaseTtlSeconds")]
        public Output<int> DefaultLeaseTtlSeconds { get; private set; } = null!;

        /// <summary>
        /// A human-friendly description for this backend.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// If set, opts out of mount migration on path updates.
        /// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        /// </summary>
        [Output("disableRemount")]
        public Output<bool?> DisableRemount { get; private set; } = null!;

        /// <summary>
        /// The maximum TTL that can be requested
        /// for credentials issued by this backend.
        /// </summary>
        [Output("maxLeaseTtlSeconds")]
        public Output<int> MaxLeaseTtlSeconds { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Specifies the RabbitMQ management administrator password.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// Specifies a password policy to use when creating dynamic credentials. Defaults to generating an alphanumeric password if not set.
        /// </summary>
        [Output("passwordPolicy")]
        public Output<string?> PasswordPolicy { get; private set; } = null!;

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `rabbitmq`.
        /// </summary>
        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;

        /// <summary>
        /// Specifies the RabbitMQ management administrator username.
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;

        /// <summary>
        /// Template describing how dynamic usernames are generated.
        /// </summary>
        [Output("usernameTemplate")]
        public Output<string?> UsernameTemplate { get; private set; } = null!;

        /// <summary>
        /// Specifies whether to verify connection URI, username, and password.
        /// Defaults to `true`.
        /// </summary>
        [Output("verifyConnection")]
        public Output<bool?> VerifyConnection { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackend resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackend(string name, SecretBackendArgs args, CustomResourceOptions? options = null)
            : base("vault:rabbitMq/secretBackend:SecretBackend", name, args ?? new SecretBackendArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretBackend(string name, Input<string> id, SecretBackendState? state = null, CustomResourceOptions? options = null)
            : base("vault:rabbitMq/secretBackend:SecretBackend", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "password",
                    "username",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecretBackend resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackend Get(string name, Input<string> id, SecretBackendState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackend(name, id, state, options);
        }
    }

    public sealed class SecretBackendArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the RabbitMQ connection URI.
        /// </summary>
        [Input("connectionUri", required: true)]
        public Input<string> ConnectionUri { get; set; } = null!;

        /// <summary>
        /// The default TTL for credentials
        /// issued by this backend.
        /// </summary>
        [Input("defaultLeaseTtlSeconds")]
        public Input<int>? DefaultLeaseTtlSeconds { get; set; }

        /// <summary>
        /// A human-friendly description for this backend.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If set, opts out of mount migration on path updates.
        /// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        /// </summary>
        [Input("disableRemount")]
        public Input<bool>? DisableRemount { get; set; }

        /// <summary>
        /// The maximum TTL that can be requested
        /// for credentials issued by this backend.
        /// </summary>
        [Input("maxLeaseTtlSeconds")]
        public Input<int>? MaxLeaseTtlSeconds { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// Specifies the RabbitMQ management administrator password.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Specifies a password policy to use when creating dynamic credentials. Defaults to generating an alphanumeric password if not set.
        /// </summary>
        [Input("passwordPolicy")]
        public Input<string>? PasswordPolicy { get; set; }

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `rabbitmq`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("username", required: true)]
        private Input<string>? _username;

        /// <summary>
        /// Specifies the RabbitMQ management administrator username.
        /// </summary>
        public Input<string>? Username
        {
            get => _username;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _username = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Template describing how dynamic usernames are generated.
        /// </summary>
        [Input("usernameTemplate")]
        public Input<string>? UsernameTemplate { get; set; }

        /// <summary>
        /// Specifies whether to verify connection URI, username, and password.
        /// Defaults to `true`.
        /// </summary>
        [Input("verifyConnection")]
        public Input<bool>? VerifyConnection { get; set; }

        public SecretBackendArgs()
        {
        }
        public static new SecretBackendArgs Empty => new SecretBackendArgs();
    }

    public sealed class SecretBackendState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the RabbitMQ connection URI.
        /// </summary>
        [Input("connectionUri")]
        public Input<string>? ConnectionUri { get; set; }

        /// <summary>
        /// The default TTL for credentials
        /// issued by this backend.
        /// </summary>
        [Input("defaultLeaseTtlSeconds")]
        public Input<int>? DefaultLeaseTtlSeconds { get; set; }

        /// <summary>
        /// A human-friendly description for this backend.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If set, opts out of mount migration on path updates.
        /// See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        /// </summary>
        [Input("disableRemount")]
        public Input<bool>? DisableRemount { get; set; }

        /// <summary>
        /// The maximum TTL that can be requested
        /// for credentials issued by this backend.
        /// </summary>
        [Input("maxLeaseTtlSeconds")]
        public Input<int>? MaxLeaseTtlSeconds { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// Specifies the RabbitMQ management administrator password.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Specifies a password policy to use when creating dynamic credentials. Defaults to generating an alphanumeric password if not set.
        /// </summary>
        [Input("passwordPolicy")]
        public Input<string>? PasswordPolicy { get; set; }

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `rabbitmq`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        [Input("username")]
        private Input<string>? _username;

        /// <summary>
        /// Specifies the RabbitMQ management administrator username.
        /// </summary>
        public Input<string>? Username
        {
            get => _username;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _username = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Template describing how dynamic usernames are generated.
        /// </summary>
        [Input("usernameTemplate")]
        public Input<string>? UsernameTemplate { get; set; }

        /// <summary>
        /// Specifies whether to verify connection URI, username, and password.
        /// Defaults to `true`.
        /// </summary>
        [Input("verifyConnection")]
        public Input<bool>? VerifyConnection { get; set; }

        public SecretBackendState()
        {
        }
        public static new SecretBackendState Empty => new SecretBackendState();
    }
}
