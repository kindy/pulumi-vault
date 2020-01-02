// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Identity
{
    /// <summary>
    /// Creates an Identity Group for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.
    /// 
    /// A group can contain multiple entities as its members. A group can also have subgroups. Policies set on the group is granted to all members of the group. During request time, when the token's entity ID is being evaluated for the policies that it has access to; along with the policies on the entity itself, policies that are inherited due to group memberships are also granted.
    /// 
    /// &gt; This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/identity_group.html.markdown.
    /// </summary>
    public partial class Group : Pulumi.CustomResource
    {
        /// <summary>
        /// `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `vault.identity.GroupPolicies` to manage policies for this group in a decoupled manner.
        /// </summary>
        [Output("externalPolicies")]
        public Output<bool?> ExternalPolicies { get; private set; } = null!;

        /// <summary>
        /// A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
        /// </summary>
        [Output("memberEntityIds")]
        public Output<ImmutableArray<string>> MemberEntityIds { get; private set; } = null!;

        /// <summary>
        /// A list of Group IDs to be assigned as group members.
        /// </summary>
        [Output("memberGroupIds")]
        public Output<ImmutableArray<string>> MemberGroupIds { get; private set; } = null!;

        /// <summary>
        /// A Map of additional metadata to associate with the group.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>?> Metadata { get; private set; } = null!;

        /// <summary>
        /// Name of the identity group to create.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A list of policies to apply to the group.
        /// </summary>
        [Output("policies")]
        public Output<ImmutableArray<string>> Policies { get; private set; } = null!;

        /// <summary>
        /// Type of the group, internal or external. Defaults to `internal`.
        /// </summary>
        [Output("type")]
        public Output<string?> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Group resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Group(string name, GroupArgs? args = null, CustomResourceOptions? options = null)
            : base("vault:identity/group:Group", name, args ?? ResourceArgs.Empty, MakeResourceOptions(options, ""))
        {
        }

        private Group(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
            : base("vault:identity/group:Group", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Group resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Group Get(string name, Input<string> id, GroupState? state = null, CustomResourceOptions? options = null)
        {
            return new Group(name, id, state, options);
        }
    }

    public sealed class GroupArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `vault.identity.GroupPolicies` to manage policies for this group in a decoupled manner.
        /// </summary>
        [Input("externalPolicies")]
        public Input<bool>? ExternalPolicies { get; set; }

        [Input("memberEntityIds")]
        private InputList<string>? _memberEntityIds;

        /// <summary>
        /// A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
        /// </summary>
        public InputList<string> MemberEntityIds
        {
            get => _memberEntityIds ?? (_memberEntityIds = new InputList<string>());
            set => _memberEntityIds = value;
        }

        [Input("memberGroupIds")]
        private InputList<string>? _memberGroupIds;

        /// <summary>
        /// A list of Group IDs to be assigned as group members.
        /// </summary>
        public InputList<string> MemberGroupIds
        {
            get => _memberGroupIds ?? (_memberGroupIds = new InputList<string>());
            set => _memberGroupIds = value;
        }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// A Map of additional metadata to associate with the group.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// Name of the identity group to create.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// A list of policies to apply to the group.
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        /// <summary>
        /// Type of the group, internal or external. Defaults to `internal`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GroupArgs()
        {
        }
    }

    public sealed class GroupState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `vault.identity.GroupPolicies` to manage policies for this group in a decoupled manner.
        /// </summary>
        [Input("externalPolicies")]
        public Input<bool>? ExternalPolicies { get; set; }

        [Input("memberEntityIds")]
        private InputList<string>? _memberEntityIds;

        /// <summary>
        /// A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
        /// </summary>
        public InputList<string> MemberEntityIds
        {
            get => _memberEntityIds ?? (_memberEntityIds = new InputList<string>());
            set => _memberEntityIds = value;
        }

        [Input("memberGroupIds")]
        private InputList<string>? _memberGroupIds;

        /// <summary>
        /// A list of Group IDs to be assigned as group members.
        /// </summary>
        public InputList<string> MemberGroupIds
        {
            get => _memberGroupIds ?? (_memberGroupIds = new InputList<string>());
            set => _memberGroupIds = value;
        }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// A Map of additional metadata to associate with the group.
        /// </summary>
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        /// <summary>
        /// Name of the identity group to create.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policies")]
        private InputList<string>? _policies;

        /// <summary>
        /// A list of policies to apply to the group.
        /// </summary>
        public InputList<string> Policies
        {
            get => _policies ?? (_policies = new InputList<string>());
            set => _policies = value;
        }

        /// <summary>
        /// Type of the group, internal or external. Defaults to `internal`.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GroupState()
        {
        }
    }
}
