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
    /// Resource for configuring MFA login-enforcement
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleMfaDuo = new Vault.Identity.MfaDuo("exampleMfaDuo", new()
    ///     {
    ///         SecretKey = "secret-key",
    ///         IntegrationKey = "int-key",
    ///         ApiHostname = "foo.baz",
    ///         PushInfo = "push-info",
    ///     });
    /// 
    ///     var exampleMfaLoginEnforcement = new Vault.Identity.MfaLoginEnforcement("exampleMfaLoginEnforcement", new()
    ///     {
    ///         MfaMethodIds = new[]
    ///         {
    ///             exampleMfaDuo.MethodId,
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// Resource can be imported using its `name` field, e.g.
    /// 
    /// ```sh
    ///  $ pulumi import vault:identity/mfaLoginEnforcement:MfaLoginEnforcement example default
    /// ```
    /// </summary>
    [VaultResourceType("vault:identity/mfaLoginEnforcement:MfaLoginEnforcement")]
    public partial class MfaLoginEnforcement : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Set of auth method accessor IDs.
        /// </summary>
        [Output("authMethodAccessors")]
        public Output<ImmutableArray<string>> AuthMethodAccessors { get; private set; } = null!;

        /// <summary>
        /// Set of auth method types.
        /// </summary>
        [Output("authMethodTypes")]
        public Output<ImmutableArray<string>> AuthMethodTypes { get; private set; } = null!;

        /// <summary>
        /// Set of identity entity IDs.
        /// </summary>
        [Output("identityEntityIds")]
        public Output<ImmutableArray<string>> IdentityEntityIds { get; private set; } = null!;

        /// <summary>
        /// Set of identity group IDs.
        /// </summary>
        [Output("identityGroupIds")]
        public Output<ImmutableArray<string>> IdentityGroupIds { get; private set; } = null!;

        /// <summary>
        /// Set of MFA method UUIDs.
        /// </summary>
        [Output("mfaMethodIds")]
        public Output<ImmutableArray<string>> MfaMethodIds { get; private set; } = null!;

        /// <summary>
        /// Login enforcement name.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Target namespace. (requires Enterprise)
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Method's namespace ID.
        /// </summary>
        [Output("namespaceId")]
        public Output<string> NamespaceId { get; private set; } = null!;

        /// <summary>
        /// Method's namespace path.
        /// </summary>
        [Output("namespacePath")]
        public Output<string> NamespacePath { get; private set; } = null!;

        /// <summary>
        /// Resource UUID.
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;


        /// <summary>
        /// Create a MfaLoginEnforcement resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public MfaLoginEnforcement(string name, MfaLoginEnforcementArgs args, CustomResourceOptions? options = null)
            : base("vault:identity/mfaLoginEnforcement:MfaLoginEnforcement", name, args ?? new MfaLoginEnforcementArgs(), MakeResourceOptions(options, ""))
        {
        }

        private MfaLoginEnforcement(string name, Input<string> id, MfaLoginEnforcementState? state = null, CustomResourceOptions? options = null)
            : base("vault:identity/mfaLoginEnforcement:MfaLoginEnforcement", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing MfaLoginEnforcement resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static MfaLoginEnforcement Get(string name, Input<string> id, MfaLoginEnforcementState? state = null, CustomResourceOptions? options = null)
        {
            return new MfaLoginEnforcement(name, id, state, options);
        }
    }

    public sealed class MfaLoginEnforcementArgs : global::Pulumi.ResourceArgs
    {
        [Input("authMethodAccessors")]
        private InputList<string>? _authMethodAccessors;

        /// <summary>
        /// Set of auth method accessor IDs.
        /// </summary>
        public InputList<string> AuthMethodAccessors
        {
            get => _authMethodAccessors ?? (_authMethodAccessors = new InputList<string>());
            set => _authMethodAccessors = value;
        }

        [Input("authMethodTypes")]
        private InputList<string>? _authMethodTypes;

        /// <summary>
        /// Set of auth method types.
        /// </summary>
        public InputList<string> AuthMethodTypes
        {
            get => _authMethodTypes ?? (_authMethodTypes = new InputList<string>());
            set => _authMethodTypes = value;
        }

        [Input("identityEntityIds")]
        private InputList<string>? _identityEntityIds;

        /// <summary>
        /// Set of identity entity IDs.
        /// </summary>
        public InputList<string> IdentityEntityIds
        {
            get => _identityEntityIds ?? (_identityEntityIds = new InputList<string>());
            set => _identityEntityIds = value;
        }

        [Input("identityGroupIds")]
        private InputList<string>? _identityGroupIds;

        /// <summary>
        /// Set of identity group IDs.
        /// </summary>
        public InputList<string> IdentityGroupIds
        {
            get => _identityGroupIds ?? (_identityGroupIds = new InputList<string>());
            set => _identityGroupIds = value;
        }

        [Input("mfaMethodIds", required: true)]
        private InputList<string>? _mfaMethodIds;

        /// <summary>
        /// Set of MFA method UUIDs.
        /// </summary>
        public InputList<string> MfaMethodIds
        {
            get => _mfaMethodIds ?? (_mfaMethodIds = new InputList<string>());
            set => _mfaMethodIds = value;
        }

        /// <summary>
        /// Login enforcement name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Target namespace. (requires Enterprise)
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        public MfaLoginEnforcementArgs()
        {
        }
        public static new MfaLoginEnforcementArgs Empty => new MfaLoginEnforcementArgs();
    }

    public sealed class MfaLoginEnforcementState : global::Pulumi.ResourceArgs
    {
        [Input("authMethodAccessors")]
        private InputList<string>? _authMethodAccessors;

        /// <summary>
        /// Set of auth method accessor IDs.
        /// </summary>
        public InputList<string> AuthMethodAccessors
        {
            get => _authMethodAccessors ?? (_authMethodAccessors = new InputList<string>());
            set => _authMethodAccessors = value;
        }

        [Input("authMethodTypes")]
        private InputList<string>? _authMethodTypes;

        /// <summary>
        /// Set of auth method types.
        /// </summary>
        public InputList<string> AuthMethodTypes
        {
            get => _authMethodTypes ?? (_authMethodTypes = new InputList<string>());
            set => _authMethodTypes = value;
        }

        [Input("identityEntityIds")]
        private InputList<string>? _identityEntityIds;

        /// <summary>
        /// Set of identity entity IDs.
        /// </summary>
        public InputList<string> IdentityEntityIds
        {
            get => _identityEntityIds ?? (_identityEntityIds = new InputList<string>());
            set => _identityEntityIds = value;
        }

        [Input("identityGroupIds")]
        private InputList<string>? _identityGroupIds;

        /// <summary>
        /// Set of identity group IDs.
        /// </summary>
        public InputList<string> IdentityGroupIds
        {
            get => _identityGroupIds ?? (_identityGroupIds = new InputList<string>());
            set => _identityGroupIds = value;
        }

        [Input("mfaMethodIds")]
        private InputList<string>? _mfaMethodIds;

        /// <summary>
        /// Set of MFA method UUIDs.
        /// </summary>
        public InputList<string> MfaMethodIds
        {
            get => _mfaMethodIds ?? (_mfaMethodIds = new InputList<string>());
            set => _mfaMethodIds = value;
        }

        /// <summary>
        /// Login enforcement name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Target namespace. (requires Enterprise)
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Method's namespace ID.
        /// </summary>
        [Input("namespaceId")]
        public Input<string>? NamespaceId { get; set; }

        /// <summary>
        /// Method's namespace path.
        /// </summary>
        [Input("namespacePath")]
        public Input<string>? NamespacePath { get; set; }

        /// <summary>
        /// Resource UUID.
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        public MfaLoginEnforcementState()
        {
        }
        public static new MfaLoginEnforcementState Empty => new MfaLoginEnforcementState();
    }
}
