// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault
{
    /// <summary>
    /// ## Import
    /// 
    /// Nomad secret backend can be imported using the `backend`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import vault:index/nomadSecretBackend:NomadSecretBackend nomad nomad
    /// ```
    /// </summary>
    public partial class NomadSecretBackend : Pulumi.CustomResource
    {
        /// <summary>
        /// Specifies the address of the Nomad instance, provided
        /// as "protocol://host:port" like "http://127.0.0.1:4646".
        /// </summary>
        [Output("address")]
        public Output<string?> Address { get; private set; } = null!;

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `nomad`.
        /// </summary>
        [Output("backend")]
        public Output<string?> Backend { get; private set; } = null!;

        /// <summary>
        /// CA certificate to use when verifying the Nomad server certificate, must be
        /// x509 PEM encoded.
        /// </summary>
        [Output("caCert")]
        public Output<string?> CaCert { get; private set; } = null!;

        /// <summary>
        /// Client certificate to provide to the Nomad server, must be x509 PEM encoded.
        /// </summary>
        [Output("clientCert")]
        public Output<string?> ClientCert { get; private set; } = null!;

        /// <summary>
        /// Client certificate key to provide to the Nomad server, must be x509 PEM encoded.
        /// </summary>
        [Output("clientKey")]
        public Output<string?> ClientKey { get; private set; } = null!;

        /// <summary>
        /// Default lease duration for secrets in seconds.
        /// </summary>
        [Output("defaultLeaseTtlSeconds")]
        public Output<int> DefaultLeaseTtlSeconds { get; private set; } = null!;

        /// <summary>
        /// Human-friendly description of the mount for the Active Directory backend.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Mark the secrets engine as local-only. Local engines are not replicated or removed by
        /// replication.Tolerance duration to use when checking the last rotation time.
        /// </summary>
        [Output("local")]
        public Output<bool?> Local { get; private set; } = null!;

        /// <summary>
        /// Maximum possible lease duration for secrets in seconds.
        /// </summary>
        [Output("maxLeaseTtlSeconds")]
        public Output<int> MaxLeaseTtlSeconds { get; private set; } = null!;

        /// <summary>
        /// Specifies the maximum length to use for the name of the Nomad token
        /// generated with Generate Credential. If omitted, 0 is used and ignored, defaulting to the max value allowed
        /// by the Nomad version.
        /// </summary>
        [Output("maxTokenNameLength")]
        public Output<int> MaxTokenNameLength { get; private set; } = null!;

        /// <summary>
        /// Maximum possible lease duration for secrets in seconds.
        /// </summary>
        [Output("maxTtl")]
        public Output<int> MaxTtl { get; private set; } = null!;

        /// <summary>
        /// Specifies the Nomad Management token to use.
        /// </summary>
        [Output("token")]
        public Output<string?> Token { get; private set; } = null!;

        /// <summary>
        /// Specifies the ttl of the lease for the generated token.
        /// </summary>
        [Output("ttl")]
        public Output<int> Ttl { get; private set; } = null!;


        /// <summary>
        /// Create a NomadSecretBackend resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NomadSecretBackend(string name, NomadSecretBackendArgs? args = null, CustomResourceOptions? options = null)
            : base("vault:index/nomadSecretBackend:NomadSecretBackend", name, args ?? new NomadSecretBackendArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NomadSecretBackend(string name, Input<string> id, NomadSecretBackendState? state = null, CustomResourceOptions? options = null)
            : base("vault:index/nomadSecretBackend:NomadSecretBackend", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NomadSecretBackend resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NomadSecretBackend Get(string name, Input<string> id, NomadSecretBackendState? state = null, CustomResourceOptions? options = null)
        {
            return new NomadSecretBackend(name, id, state, options);
        }
    }

    public sealed class NomadSecretBackendArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the address of the Nomad instance, provided
        /// as "protocol://host:port" like "http://127.0.0.1:4646".
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `nomad`.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// CA certificate to use when verifying the Nomad server certificate, must be
        /// x509 PEM encoded.
        /// </summary>
        [Input("caCert")]
        public Input<string>? CaCert { get; set; }

        /// <summary>
        /// Client certificate to provide to the Nomad server, must be x509 PEM encoded.
        /// </summary>
        [Input("clientCert")]
        public Input<string>? ClientCert { get; set; }

        /// <summary>
        /// Client certificate key to provide to the Nomad server, must be x509 PEM encoded.
        /// </summary>
        [Input("clientKey")]
        public Input<string>? ClientKey { get; set; }

        /// <summary>
        /// Default lease duration for secrets in seconds.
        /// </summary>
        [Input("defaultLeaseTtlSeconds")]
        public Input<int>? DefaultLeaseTtlSeconds { get; set; }

        /// <summary>
        /// Human-friendly description of the mount for the Active Directory backend.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Mark the secrets engine as local-only. Local engines are not replicated or removed by
        /// replication.Tolerance duration to use when checking the last rotation time.
        /// </summary>
        [Input("local")]
        public Input<bool>? Local { get; set; }

        /// <summary>
        /// Maximum possible lease duration for secrets in seconds.
        /// </summary>
        [Input("maxLeaseTtlSeconds")]
        public Input<int>? MaxLeaseTtlSeconds { get; set; }

        /// <summary>
        /// Specifies the maximum length to use for the name of the Nomad token
        /// generated with Generate Credential. If omitted, 0 is used and ignored, defaulting to the max value allowed
        /// by the Nomad version.
        /// </summary>
        [Input("maxTokenNameLength")]
        public Input<int>? MaxTokenNameLength { get; set; }

        /// <summary>
        /// Maximum possible lease duration for secrets in seconds.
        /// </summary>
        [Input("maxTtl")]
        public Input<int>? MaxTtl { get; set; }

        /// <summary>
        /// Specifies the Nomad Management token to use.
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        /// <summary>
        /// Specifies the ttl of the lease for the generated token.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        public NomadSecretBackendArgs()
        {
        }
    }

    public sealed class NomadSecretBackendState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies the address of the Nomad instance, provided
        /// as "protocol://host:port" like "http://127.0.0.1:4646".
        /// </summary>
        [Input("address")]
        public Input<string>? Address { get; set; }

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `nomad`.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// CA certificate to use when verifying the Nomad server certificate, must be
        /// x509 PEM encoded.
        /// </summary>
        [Input("caCert")]
        public Input<string>? CaCert { get; set; }

        /// <summary>
        /// Client certificate to provide to the Nomad server, must be x509 PEM encoded.
        /// </summary>
        [Input("clientCert")]
        public Input<string>? ClientCert { get; set; }

        /// <summary>
        /// Client certificate key to provide to the Nomad server, must be x509 PEM encoded.
        /// </summary>
        [Input("clientKey")]
        public Input<string>? ClientKey { get; set; }

        /// <summary>
        /// Default lease duration for secrets in seconds.
        /// </summary>
        [Input("defaultLeaseTtlSeconds")]
        public Input<int>? DefaultLeaseTtlSeconds { get; set; }

        /// <summary>
        /// Human-friendly description of the mount for the Active Directory backend.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Mark the secrets engine as local-only. Local engines are not replicated or removed by
        /// replication.Tolerance duration to use when checking the last rotation time.
        /// </summary>
        [Input("local")]
        public Input<bool>? Local { get; set; }

        /// <summary>
        /// Maximum possible lease duration for secrets in seconds.
        /// </summary>
        [Input("maxLeaseTtlSeconds")]
        public Input<int>? MaxLeaseTtlSeconds { get; set; }

        /// <summary>
        /// Specifies the maximum length to use for the name of the Nomad token
        /// generated with Generate Credential. If omitted, 0 is used and ignored, defaulting to the max value allowed
        /// by the Nomad version.
        /// </summary>
        [Input("maxTokenNameLength")]
        public Input<int>? MaxTokenNameLength { get; set; }

        /// <summary>
        /// Maximum possible lease duration for secrets in seconds.
        /// </summary>
        [Input("maxTtl")]
        public Input<int>? MaxTtl { get; set; }

        /// <summary>
        /// Specifies the Nomad Management token to use.
        /// </summary>
        [Input("token")]
        public Input<string>? Token { get; set; }

        /// <summary>
        /// Specifies the ttl of the lease for the generated token.
        /// </summary>
        [Input("ttl")]
        public Input<int>? Ttl { get; set; }

        public NomadSecretBackendState()
        {
        }
    }
}