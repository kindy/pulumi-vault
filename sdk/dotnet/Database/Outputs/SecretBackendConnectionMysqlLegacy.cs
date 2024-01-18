// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Database.Outputs
{

    [OutputType]
    public sealed class SecretBackendConnectionMysqlLegacy
    {
        /// <summary>
        /// Enable IAM authentication to a Google Cloud instance when set to `gcp_iam`
        /// </summary>
        public readonly string? AuthType;
        /// <summary>
        /// A URL containing connection information. See
        /// the [Vault
        /// docs](https://www.vaultproject.io/api-docs/secret/databases/mongodb.html#sample-payload)
        /// for an example.
        /// </summary>
        public readonly string? ConnectionUrl;
        /// <summary>
        /// The maximum number of seconds to keep
        /// a connection alive for.
        /// </summary>
        public readonly int? MaxConnectionLifetime;
        /// <summary>
        /// The maximum number of idle connections to
        /// maintain.
        /// </summary>
        public readonly int? MaxIdleConnections;
        /// <summary>
        /// The maximum number of open connections to
        /// use.
        /// </summary>
        public readonly int? MaxOpenConnections;
        /// <summary>
        /// The password to authenticate with.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// JSON encoding of an IAM access key. Requires `auth_type` to be `gcp_iam`.
        /// </summary>
        public readonly string? ServiceAccountJson;
        /// <summary>
        /// x509 CA file for validating the certificate presented by the MySQL server. Must be PEM encoded.
        /// </summary>
        public readonly string? TlsCa;
        /// <summary>
        /// x509 certificate for connecting to the database. This must be a PEM encoded version of the private key and the certificate combined.
        /// </summary>
        public readonly string? TlsCertificateKey;
        /// <summary>
        /// The username to authenticate with.
        /// </summary>
        public readonly string? Username;
        /// <summary>
        /// Template describing how dynamic usernames are generated.
        /// </summary>
        public readonly string? UsernameTemplate;

        [OutputConstructor]
        private SecretBackendConnectionMysqlLegacy(
            string? authType,

            string? connectionUrl,

            int? maxConnectionLifetime,

            int? maxIdleConnections,

            int? maxOpenConnections,

            string? password,

            string? serviceAccountJson,

            string? tlsCa,

            string? tlsCertificateKey,

            string? username,

            string? usernameTemplate)
        {
            AuthType = authType;
            ConnectionUrl = connectionUrl;
            MaxConnectionLifetime = maxConnectionLifetime;
            MaxIdleConnections = maxIdleConnections;
            MaxOpenConnections = maxOpenConnections;
            Password = password;
            ServiceAccountJson = serviceAccountJson;
            TlsCa = tlsCa;
            TlsCertificateKey = tlsCertificateKey;
            Username = username;
            UsernameTemplate = usernameTemplate;
        }
    }
}
