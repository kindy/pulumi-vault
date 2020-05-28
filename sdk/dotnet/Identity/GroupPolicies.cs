// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Identity
{
    /// <summary>
    /// Manages policies for an Identity Group for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.
    /// 
    /// ## Example Usage
    /// 
    /// ### Exclusive Policies
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @internal = new Vault.Identity.Group("internal", new Vault.Identity.GroupArgs
    ///         {
    ///             Type = "internal",
    ///             ExternalPolicies = true,
    ///             Metadata = 
    ///             {
    ///                 { "version", "2" },
    ///             },
    ///         });
    ///         var policies = new Vault.Identity.GroupPolicies("policies", new Vault.Identity.GroupPoliciesArgs
    ///         {
    ///             Policies = 
    ///             {
    ///                 "default",
    ///                 "test",
    ///             },
    ///             Exclusive = true,
    ///             GroupId = @internal.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ### Non-exclusive Policies
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var @internal = new Vault.Identity.Group("internal", new Vault.Identity.GroupArgs
    ///         {
    ///             Type = "internal",
    ///             ExternalPolicies = true,
    ///             Metadata = 
    ///             {
    ///                 { "version", "2" },
    ///             },
    ///         });
    ///         var @default = new Vault.Identity.GroupPolicies("default", new Vault.Identity.GroupPoliciesArgs
    ///         {
    ///             Policies = 
    ///             {
    ///                 "default",
    ///                 "test",
    ///             },
    ///             Exclusive = false,
    ///             GroupId = @internal.Id,
    ///         });
    ///         var others = new Vault.Identity.GroupPolicies("others", new Vault.Identity.GroupPoliciesArgs
    ///         {
    ///             Policies = 
    ///             {
    ///                 "others",
    ///             },
    ///             Exclusive = false,
    ///             GroupId = @internal.Id,
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    public partial class GroupPolicies : Pulumi.CustomResource
    {
        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Output("exclusive")]
        public Output<bool?> Exclusive { get; private set; } = null!;

        /// <summary>
        /// Group ID to assign policies to.
        /// </summary>
        [Output("groupId")]
        public Output<string> GroupId { get; private set; } = null!;

        /// <summary>
        /// The name of the group that are assigned the policies.
        /// </summary>
        [Output("groupName")]
        public Output<string> GroupName { get; private set; } = null!;

        /// <summary>
        /// List of policies to assign to the group
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<string>> Policies { get; private set; } = null!;


        /// <summary>
        /// Create a GroupPolicies resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public GroupPolicies(string name, GroupPoliciesArgs args, CustomResourceOptions? options = null)
            : base("vault:identity/groupPolicies:GroupPolicies", name, args ?? new GroupPoliciesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private GroupPolicies(string name, Input<string> id, GroupPoliciesState? state = null, CustomResourceOptions? options = null)
            : base("vault:identity/groupPolicies:GroupPolicies", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing GroupPolicies resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static GroupPolicies Get(string name, Input<string> id, GroupPoliciesState? state = null, CustomResourceOptions? options = null)
        {
            return new GroupPolicies(name, id, state, options);
        }
    }

    public sealed class GroupPoliciesArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Input("exclusive")]
        public Input<bool>? Exclusive { get; set; }

        /// <summary>
        /// Group ID to assign policies to.
        /// </summary>
        [Input("groupId", required: true)]
        public Input<string> GroupId { get; set; } = null!;

        [Input("policies", required: true)]
        private InputList<string>? _policies;

        /// <summary>
        /// List of policies to assign to the group
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        public GroupPoliciesArgs()
        {
        }
    }

    public sealed class GroupPoliciesState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defaults to `true`.
        /// </summary>
        [Input("exclusive")]
        public Input<bool>? Exclusive { get; set; }

        /// <summary>
        /// Group ID to assign policies to.
        /// </summary>
        [Input("groupId")]
        public Input<string>? GroupId { get; set; }

        /// <summary>
        /// The name of the group that are assigned the policies.
        /// </summary>
        [Input("groupName")]
        public Input<string>? GroupName { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// List of policies to assign to the group
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        public GroupPoliciesState()
        {
        }
    }
}
