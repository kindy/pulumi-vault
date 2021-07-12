// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Generic
{
    public static class GetSecret
    {
        public static Task<GetSecretResult> InvokeAsync(GetSecretArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetSecretResult>("vault:generic/getSecret:getSecret", args ?? new GetSecretArgs(), options.WithVersion());
    }


    public sealed class GetSecretArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The full logical path from which to request data.
        /// To read data from the "generic" secret backend mounted in Vault by
        /// default, this should be prefixed with `secret/`. Reading from other backends
        /// with this data source is possible; consult each backend's documentation
        /// to see which endpoints support the `GET` method.
        /// </summary>
        [Input("path", required: true)]
        public string Path { get; set; } = null!;

        /// <summary>
        /// The version of the secret to read. This is used by the
        /// Vault KV secrets engine - version 2 to indicate which version of the secret
        /// to read.
        /// </summary>
        [Input("version")]
        public int? Version { get; set; }

        public GetSecretArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetSecretResult
    {
        /// <summary>
        /// A mapping whose keys are the top-level data keys returned from
        /// Vault and whose values are the corresponding values. This map can only
        /// represent string data, so any non-string values returned from Vault are
        /// serialized as JSON.
        /// </summary>
        public readonly ImmutableDictionary<string, object> Data;
        /// <summary>
        /// A string containing the full data payload retrieved from
        /// Vault, serialized in JSON format.
        /// </summary>
        public readonly string DataJson;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The duration of the secret lease, in seconds relative
        /// to the time the data was requested. Once this time has passed any plan
        /// generated with this data may fail to apply.
        /// </summary>
        public readonly int LeaseDuration;
        /// <summary>
        /// The lease identifier assigned by Vault, if any.
        /// </summary>
        public readonly string LeaseId;
        public readonly bool LeaseRenewable;
        public readonly string LeaseStartTime;
        public readonly string Path;
        public readonly int? Version;

        [OutputConstructor]
        private GetSecretResult(
            ImmutableDictionary<string, object> data,

            string dataJson,

            string id,

            int leaseDuration,

            string leaseId,

            bool leaseRenewable,

            string leaseStartTime,

            string path,

            int? version)
        {
            Data = data;
            DataJson = dataJson;
            Id = id;
            LeaseDuration = leaseDuration;
            LeaseId = leaseId;
            LeaseRenewable = leaseRenewable;
            LeaseStartTime = leaseStartTime;
            Path = path;
            Version = version;
        }
    }
}
