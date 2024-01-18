// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Gcp.Outputs
{

    [OutputType]
    public sealed class AuthBackendCustomEndpoint
    {
        /// <summary>
        /// Replaces the service endpoint used in API requests to `https://www.googleapis.com`.
        /// </summary>
        public readonly string? Api;
        /// <summary>
        /// Replaces the service endpoint used in API requests to `https://compute.googleapis.com`.
        /// 
        /// The endpoint value provided for a given key has the form of `scheme://host:port`.
        /// The `scheme://` and `:port` portions of the endpoint value are optional.
        /// </summary>
        public readonly string? Compute;
        /// <summary>
        /// Replaces the service endpoint used in API requests to `https://cloudresourcemanager.googleapis.com`.
        /// </summary>
        public readonly string? Crm;
        /// <summary>
        /// Replaces the service endpoint used in API requests to `https://iam.googleapis.com`.
        /// </summary>
        public readonly string? Iam;

        [OutputConstructor]
        private AuthBackendCustomEndpoint(
            string? api,

            string? compute,

            string? crm,

            string? iam)
        {
            Api = api;
            Compute = compute;
            Crm = crm;
            Iam = iam;
        }
    }
}
