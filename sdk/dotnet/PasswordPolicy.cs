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
    /// Provides a resource to manage Password Policies
    /// 
    /// **Note** this feature is available only Vault 1.5+
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var alphanumeric = new Vault.PasswordPolicy("alphanumeric", new Vault.PasswordPolicyArgs
    ///         {
    ///             Policy = @"    length = 20
    ///     rule ""charset"" {
    ///       charset = ""abcdefghijklmnopqrstuvwxyz0123456789""
    ///     }
    ///   
    /// ",
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Password policies can be imported using the `name`, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import vault:index/passwordPolicy:PasswordPolicy alphanumeric alphanumeric
    /// ```
    /// </summary>
    [VaultResourceType("vault:index/passwordPolicy:PasswordPolicy")]
    public partial class PasswordPolicy : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the password policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// String containing a password policy.
        /// </summary>
        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;


        /// <summary>
        /// Create a PasswordPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public PasswordPolicy(string name, PasswordPolicyArgs args, CustomResourceOptions? options = null)
            : base("vault:index/passwordPolicy:PasswordPolicy", name, args ?? new PasswordPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private PasswordPolicy(string name, Input<string> id, PasswordPolicyState? state = null, CustomResourceOptions? options = null)
            : base("vault:index/passwordPolicy:PasswordPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing PasswordPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static PasswordPolicy Get(string name, Input<string> id, PasswordPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new PasswordPolicy(name, id, state, options);
        }
    }

    public sealed class PasswordPolicyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the password policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// String containing a password policy.
        /// </summary>
        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        public PasswordPolicyArgs()
        {
        }
    }

    public sealed class PasswordPolicyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the password policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// String containing a password policy.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        public PasswordPolicyState()
        {
        }
    }
}
