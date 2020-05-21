// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Immutable;

namespace Pulumi.Vault
{
    public static class Config
    {
        private static readonly Pulumi.Config __config = new Pulumi.Config("vault");
        /// <summary>
        /// If true, adds the value of the `address` argument to the Terraform process environment.
        /// </summary>
        public static string? AddAddressToEnv { get; set; } = __config.Get("addAddressToEnv");

        /// <summary>
        /// URL of the root of the target Vault server.
        /// </summary>
        public static string? Address { get; set; } = __config.Get("address") ?? Utilities.GetEnv("VAULT_ADDR");

        /// <summary>
        /// Login to vault with an existing auth method using auth/&lt;mount&gt;/login
        /// </summary>
        public static ImmutableArray<Pulumi.Vault.Config.Types.AuthLogins> AuthLogins { get; set; } = __config.GetObject<ImmutableArray<Pulumi.Vault.Config.Types.AuthLogins>>("authLogins");

        /// <summary>
        /// Path to directory containing CA certificate files to validate the server's certificate.
        /// </summary>
        public static string? CaCertDir { get; set; } = __config.Get("caCertDir") ?? Utilities.GetEnv("VAULT_CAPATH");

        /// <summary>
        /// Path to a CA certificate file to validate the server's certificate.
        /// </summary>
        public static string? CaCertFile { get; set; } = __config.Get("caCertFile") ?? Utilities.GetEnv("VAULT_CACERT");

        /// <summary>
        /// Client authentication credentials.
        /// </summary>
        public static ImmutableArray<Pulumi.Vault.Config.Types.ClientAuths> ClientAuths { get; set; } = __config.GetObject<ImmutableArray<Pulumi.Vault.Config.Types.ClientAuths>>("clientAuths");

        /// <summary>
        /// The headers to send with each Vault request.
        /// </summary>
        public static ImmutableArray<Pulumi.Vault.Config.Types.Headers> Headers { get; set; } = __config.GetObject<ImmutableArray<Pulumi.Vault.Config.Types.Headers>>("headers");

        /// <summary>
        /// Maximum TTL for secret leases requested by this provider
        /// </summary>
        public static int? MaxLeaseTtlSeconds { get; set; } = __config.GetInt32("maxLeaseTtlSeconds") ?? Utilities.GetEnvInt32("TERRAFORM_VAULT_MAX_TTL") ?? 1200;

        /// <summary>
        /// Maximum number of retries when a 5xx error code is encountered.
        /// </summary>
        public static int? MaxRetries { get; set; } = __config.GetInt32("maxRetries") ?? Utilities.GetEnvInt32("VAULT_MAX_RETRIES") ?? 2;

        /// <summary>
        /// The namespace to use. Available only for Vault Enterprise
        /// </summary>
        public static string? Namespace { get; set; } = __config.Get("namespace") ?? Utilities.GetEnv("VAULT_NAMESPACE");

        /// <summary>
        /// Set this to true only if the target Vault server is an insecure development instance.
        /// </summary>
        public static bool? SkipTlsVerify { get; set; } = __config.GetBoolean("skipTlsVerify") ?? Utilities.GetEnvBoolean("VAULT_SKIP_VERIFY");

        /// <summary>
        /// Token to use to authenticate to Vault.
        /// </summary>
        public static string? Token { get; set; } = __config.Get("token") ?? Utilities.GetEnv("VAULT_TOKEN");

        /// <summary>
        /// Token name to use for creating the Vault child token.
        /// </summary>
        public static string? TokenName { get; set; } = __config.Get("tokenName") ?? Utilities.GetEnv("VAULT_TOKEN_NAME");

        public static class Types
        {

             public class AuthLogins
             {
                public string? Namespace { get; set; } = null!;
                public ImmutableDictionary<string, string>? Parameters { get; set; } = null!;
                public string Path { get; set; }
            }

             public class ClientAuths
             {
                public string CertFile { get; set; }
                public string KeyFile { get; set; }
            }

             public class Headers
             {
                public string Name { get; set; }
                public string Value { get; set; }
            }
        }
    }
}
