// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Okta.Inputs
{

    public sealed class AuthBackendGroupGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the group within the Okta
        /// </summary>
        [Input("groupName", required: true)]
        public Input<string> GroupName { get; set; } = null!;

        [Input("policies", required: true)]
        private InputList<string>? _policies;

        /// <summary>
        /// List of Vault policies to associate with this user
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        public AuthBackendGroupGetArgs()
        {
        }
        public static new AuthBackendGroupGetArgs Empty => new AuthBackendGroupGetArgs();
    }
}
