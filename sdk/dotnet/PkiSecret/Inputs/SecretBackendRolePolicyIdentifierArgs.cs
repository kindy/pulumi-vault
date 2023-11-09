// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.PkiSecret.Inputs
{

    public sealed class SecretBackendRolePolicyIdentifierArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The URL of the CPS for the policy identifier
        /// 
        /// Example usage:
        /// </summary>
        [Input("cps")]
        public Input<string>? Cps { get; set; }

        /// <summary>
        /// A notice for the policy identifier
        /// </summary>
        [Input("notice")]
        public Input<string>? Notice { get; set; }

        /// <summary>
        /// The OID for the policy identifier
        /// </summary>
        [Input("oid", required: true)]
        public Input<string> Oid { get; set; } = null!;

        public SecretBackendRolePolicyIdentifierArgs()
        {
        }
        public static new SecretBackendRolePolicyIdentifierArgs Empty => new SecretBackendRolePolicyIdentifierArgs();
    }
}
