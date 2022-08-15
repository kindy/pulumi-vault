// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Ssh.Inputs
{

    public sealed class SecretBackendRoleAllowedUserKeyConfigGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("lengths", required: true)]
        private InputList<int>? _lengths;

        /// <summary>
        /// A list of allowed key lengths as integers. 
        /// For key types that do not support setting the length a value of `[0]` should be used.
        /// Setting multiple lengths is only supported on Vault 1.10+. For prior releases `length`
        /// must be set to a single element list.
        /// </summary>
        public InputList<int> Lengths
        {
            get => _lengths ?? (_lengths = new InputList<int>());
            set => _lengths = value;
        }

        /// <summary>
        /// The SSH public key type.  
        /// *Supported key types are:*
        /// `rsa`, `ecdsa`, `ec`, `dsa`, `ed25519`, `ssh-rsa`, `ssh-dss`, `ssh-ed25519`,
        /// `ecdsa-sha2-nistp256`, `ecdsa-sha2-nistp384`, `ecdsa-sha2-nistp521`
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public SecretBackendRoleAllowedUserKeyConfigGetArgs()
        {
        }
        public static new SecretBackendRoleAllowedUserKeyConfigGetArgs Empty => new SecretBackendRoleAllowedUserKeyConfigGetArgs();
    }
}
