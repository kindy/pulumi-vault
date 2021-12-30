// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Database.Inputs
{

    public sealed class SecretBackendConnectionInfluxdbArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of seconds to use as a connection
        /// timeout.
        /// </summary>
        [Input("connectTimeout")]
        public Input<int>? ConnectTimeout { get; set; }

        /// <summary>
        /// The host to connect to.
        /// </summary>
        [Input("host", required: true)]
        public Input<string> Host { get; set; } = null!;

        /// <summary>
        /// Whether to skip verification of the server
        /// certificate when using TLS.
        /// </summary>
        [Input("insecureTls")]
        public Input<bool>? InsecureTls { get; set; }

        /// <summary>
        /// The password to be used in the connection.
        /// </summary>
        [Input("password", required: true)]
        public Input<string> Password { get; set; } = null!;

        /// <summary>
        /// Concatenated PEM blocks configuring the certificate
        /// chain.
        /// </summary>
        [Input("pemBundle")]
        public Input<string>? PemBundle { get; set; }

        /// <summary>
        /// A JSON structure configuring the certificate chain.
        /// </summary>
        [Input("pemJson")]
        public Input<string>? PemJson { get; set; }

        /// <summary>
        /// The default port to connect to if no port is specified as
        /// part of the host.
        /// </summary>
        [Input("port")]
        public Input<int>? Port { get; set; }

        /// <summary>
        /// Whether to use TLS when connecting to Cassandra.
        /// </summary>
        [Input("tls")]
        public Input<bool>? Tls { get; set; }

        /// <summary>
        /// The username to be used in the connection (the account admin level).
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        /// <summary>
        /// - [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated.
        /// </summary>
        [Input("usernameTemplate")]
        public Input<string>? UsernameTemplate { get; set; }

        public SecretBackendConnectionInfluxdbArgs()
        {
        }
    }
}
