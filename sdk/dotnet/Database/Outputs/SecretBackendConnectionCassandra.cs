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
    public sealed class SecretBackendConnectionCassandra
    {
        /// <summary>
        /// The number of seconds to use as a connection
        /// timeout.
        /// </summary>
        public readonly int? ConnectTimeout;
        /// <summary>
        /// The hosts to connect to.
        /// </summary>
        public readonly ImmutableArray<string> Hosts;
        /// <summary>
        /// Whether to skip verification of the server
        /// certificate when using TLS.
        /// </summary>
        public readonly bool? InsecureTls;
        /// <summary>
        /// The password to authenticate with.
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// Concatenated PEM blocks configuring the certificate
        /// chain.
        /// </summary>
        public readonly string? PemBundle;
        /// <summary>
        /// A JSON structure configuring the certificate chain.
        /// </summary>
        public readonly string? PemJson;
        /// <summary>
        /// The default port to connect to if no port is specified as
        /// part of the host.
        /// </summary>
        public readonly int? Port;
        /// <summary>
        /// The CQL protocol version to use.
        /// </summary>
        public readonly int? ProtocolVersion;
        /// <summary>
        /// Whether to use TLS when connecting to Cassandra.
        /// </summary>
        public readonly bool? Tls;
        /// <summary>
        /// The username to authenticate with.
        /// </summary>
        public readonly string? Username;

        [OutputConstructor]
        private SecretBackendConnectionCassandra(
            int? connectTimeout,

            ImmutableArray<string> hosts,

            bool? insecureTls,

            string? password,

            string? pemBundle,

            string? pemJson,

            int? port,

            int? protocolVersion,

            bool? tls,

            string? username)
        {
            ConnectTimeout = connectTimeout;
            Hosts = hosts;
            InsecureTls = insecureTls;
            Password = password;
            PemBundle = pemBundle;
            PemJson = pemJson;
            Port = port;
            ProtocolVersion = protocolVersion;
            Tls = tls;
            Username = username;
        }
    }
}
