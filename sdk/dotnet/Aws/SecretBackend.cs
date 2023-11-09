// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Aws
{
    /// <summary>
    /// ## Import
    /// 
    /// AWS secret backends can be imported using the `path`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import vault:aws/secretBackend:SecretBackend aws aws
    /// ```
    /// </summary>
    [VaultResourceType("vault:aws/secretBackend:SecretBackend")]
    public partial class SecretBackend : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The AWS Access Key ID this backend should use to
        /// issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
        /// </summary>
        [Output("accessKey")]
        public Output<string?> AccessKey { get; private set; } = null!;

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
        /// Specifies a custom HTTP IAM endpoint to use.
        /// </summary>
        [Output("iamEndpoint")]
        public Output<string?> IamEndpoint { get; private set; } = null!;

        /// <summary>
        /// Specifies whether the secrets mount will be marked as local. Local mounts are not replicated to performance replicas.
        /// </summary>
        [Output("local")]
        public Output<bool?> Local { get; private set; } = null!;

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
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `aws`.
        /// </summary>
        [Output("path")]
        public Output<string?> Path { get; private set; } = null!;

        /// <summary>
        /// The AWS region to make API calls against. Defaults to us-east-1.
        /// </summary>
        [Output("region")]
        public Output<string> Region { get; private set; } = null!;

        /// <summary>
        /// The AWS Secret Access Key to use when generating new credentials.
        /// </summary>
        [Output("secretKey")]
        public Output<string?> SecretKey { get; private set; } = null!;

        /// <summary>
        /// Specifies a custom HTTP STS endpoint to use.
        /// </summary>
        [Output("stsEndpoint")]
        public Output<string?> StsEndpoint { get; private set; } = null!;

        /// <summary>
        /// Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template:
        /// </summary>
        [Output("usernameTemplate")]
        public Output<string> UsernameTemplate { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackend resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackend(string name, SecretBackendArgs? args = null, CustomResourceOptions? options = null)
            : base("vault:aws/secretBackend:SecretBackend", name, args ?? new SecretBackendArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretBackend(string name, Input<string> id, SecretBackendState? state = null, CustomResourceOptions? options = null)
            : base("vault:aws/secretBackend:SecretBackend", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                AdditionalSecretOutputs =
                {
                    "accessKey",
                    "secretKey",
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
        [Input("accessKey")]
        private Input<string>? _accessKey;

        /// <summary>
        /// The AWS Access Key ID this backend should use to
        /// issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
        /// </summary>
        public Input<string>? AccessKey
        {
            get => _accessKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

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
        /// Specifies a custom HTTP IAM endpoint to use.
        /// </summary>
        [Input("iamEndpoint")]
        public Input<string>? IamEndpoint { get; set; }

        /// <summary>
        /// Specifies whether the secrets mount will be marked as local. Local mounts are not replicated to performance replicas.
        /// </summary>
        [Input("local")]
        public Input<bool>? Local { get; set; }

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

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `aws`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The AWS region to make API calls against. Defaults to us-east-1.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("secretKey")]
        private Input<string>? _secretKey;

        /// <summary>
        /// The AWS Secret Access Key to use when generating new credentials.
        /// </summary>
        public Input<string>? SecretKey
        {
            get => _secretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Specifies a custom HTTP STS endpoint to use.
        /// </summary>
        [Input("stsEndpoint")]
        public Input<string>? StsEndpoint { get; set; }

        /// <summary>
        /// Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template:
        /// </summary>
        [Input("usernameTemplate")]
        public Input<string>? UsernameTemplate { get; set; }

        public SecretBackendArgs()
        {
        }
        public static new SecretBackendArgs Empty => new SecretBackendArgs();
    }

    public sealed class SecretBackendState : global::Pulumi.ResourceArgs
    {
        [Input("accessKey")]
        private Input<string>? _accessKey;

        /// <summary>
        /// The AWS Access Key ID this backend should use to
        /// issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
        /// </summary>
        public Input<string>? AccessKey
        {
            get => _accessKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _accessKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

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
        /// Specifies a custom HTTP IAM endpoint to use.
        /// </summary>
        [Input("iamEndpoint")]
        public Input<string>? IamEndpoint { get; set; }

        /// <summary>
        /// Specifies whether the secrets mount will be marked as local. Local mounts are not replicated to performance replicas.
        /// </summary>
        [Input("local")]
        public Input<bool>? Local { get; set; }

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

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `aws`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The AWS region to make API calls against. Defaults to us-east-1.
        /// </summary>
        [Input("region")]
        public Input<string>? Region { get; set; }

        [Input("secretKey")]
        private Input<string>? _secretKey;

        /// <summary>
        /// The AWS Secret Access Key to use when generating new credentials.
        /// </summary>
        public Input<string>? SecretKey
        {
            get => _secretKey;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _secretKey = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// Specifies a custom HTTP STS endpoint to use.
        /// </summary>
        [Input("stsEndpoint")]
        public Input<string>? StsEndpoint { get; set; }

        /// <summary>
        /// Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template:
        /// </summary>
        [Input("usernameTemplate")]
        public Input<string>? UsernameTemplate { get; set; }

        public SecretBackendState()
        {
        }
        public static new SecretBackendState Empty => new SecretBackendState();
    }
}
