// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupMemberEntityIdsState extends com.pulumi.resources.ResourceArgs {

    public static final GroupMemberEntityIdsState Empty = new GroupMemberEntityIdsState();

    /**
     * Defaults to `true`.
     * 
     */
    @Import(name="exclusive")
    private @Nullable Output<Boolean> exclusive;

    /**
     * @return Defaults to `true`.
     * 
     */
    public Optional<Output<Boolean>> exclusive() {
        return Optional.ofNullable(this.exclusive);
    }

    /**
     * Group ID to assign member entities to.
     * 
     */
    @Import(name="groupId")
    private @Nullable Output<String> groupId;

    /**
     * @return Group ID to assign member entities to.
     * 
     */
    public Optional<Output<String>> groupId() {
        return Optional.ofNullable(this.groupId);
    }

    /**
     * The name of the group that are assigned the member entities.\
     * *Deprecated: The value for group_name may not always be accurate*
     * *use* `data.vault_identity_group.*.group_name`, *or* `vault_identity_group.*.group_name` *instead.*
     * 
     * @deprecated
     * The value for group_name may not always be accurate,
     * use &#34;data.vault_identity_group.*.group_name&#34;, &#34;vault_identity_group.*.group_name&#34; instead
     * 
     */
    @Deprecated /* The value for group_name may not always be accurate, 
use ""data.vault_identity_group.*.group_name"", ""vault_identity_group.*.group_name"" instead */
    @Import(name="groupName")
    private @Nullable Output<String> groupName;

    /**
     * @return The name of the group that are assigned the member entities.\
     * *Deprecated: The value for group_name may not always be accurate*
     * *use* `data.vault_identity_group.*.group_name`, *or* `vault_identity_group.*.group_name` *instead.*
     * 
     * @deprecated
     * The value for group_name may not always be accurate,
     * use &#34;data.vault_identity_group.*.group_name&#34;, &#34;vault_identity_group.*.group_name&#34; instead
     * 
     */
    @Deprecated /* The value for group_name may not always be accurate, 
use ""data.vault_identity_group.*.group_name"", ""vault_identity_group.*.group_name"" instead */
    public Optional<Output<String>> groupName() {
        return Optional.ofNullable(this.groupName);
    }

    /**
     * List of member entities that belong to the group
     * 
     */
    @Import(name="memberEntityIds")
    private @Nullable Output<List<String>> memberEntityIds;

    /**
     * @return List of member entities that belong to the group
     * 
     */
    public Optional<Output<List<String>>> memberEntityIds() {
        return Optional.ofNullable(this.memberEntityIds);
    }

    private GroupMemberEntityIdsState() {}

    private GroupMemberEntityIdsState(GroupMemberEntityIdsState $) {
        this.exclusive = $.exclusive;
        this.groupId = $.groupId;
        this.groupName = $.groupName;
        this.memberEntityIds = $.memberEntityIds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupMemberEntityIdsState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupMemberEntityIdsState $;

        public Builder() {
            $ = new GroupMemberEntityIdsState();
        }

        public Builder(GroupMemberEntityIdsState defaults) {
            $ = new GroupMemberEntityIdsState(Objects.requireNonNull(defaults));
        }

        /**
         * @param exclusive Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder exclusive(@Nullable Output<Boolean> exclusive) {
            $.exclusive = exclusive;
            return this;
        }

        /**
         * @param exclusive Defaults to `true`.
         * 
         * @return builder
         * 
         */
        public Builder exclusive(Boolean exclusive) {
            return exclusive(Output.of(exclusive));
        }

        /**
         * @param groupId Group ID to assign member entities to.
         * 
         * @return builder
         * 
         */
        public Builder groupId(@Nullable Output<String> groupId) {
            $.groupId = groupId;
            return this;
        }

        /**
         * @param groupId Group ID to assign member entities to.
         * 
         * @return builder
         * 
         */
        public Builder groupId(String groupId) {
            return groupId(Output.of(groupId));
        }

        /**
         * @param groupName The name of the group that are assigned the member entities.\
         * *Deprecated: The value for group_name may not always be accurate*
         * *use* `data.vault_identity_group.*.group_name`, *or* `vault_identity_group.*.group_name` *instead.*
         * 
         * @return builder
         * 
         * @deprecated
         * The value for group_name may not always be accurate,
         * use &#34;data.vault_identity_group.*.group_name&#34;, &#34;vault_identity_group.*.group_name&#34; instead
         * 
         */
        @Deprecated /* The value for group_name may not always be accurate, 
use ""data.vault_identity_group.*.group_name"", ""vault_identity_group.*.group_name"" instead */
        public Builder groupName(@Nullable Output<String> groupName) {
            $.groupName = groupName;
            return this;
        }

        /**
         * @param groupName The name of the group that are assigned the member entities.\
         * *Deprecated: The value for group_name may not always be accurate*
         * *use* `data.vault_identity_group.*.group_name`, *or* `vault_identity_group.*.group_name` *instead.*
         * 
         * @return builder
         * 
         * @deprecated
         * The value for group_name may not always be accurate,
         * use &#34;data.vault_identity_group.*.group_name&#34;, &#34;vault_identity_group.*.group_name&#34; instead
         * 
         */
        @Deprecated /* The value for group_name may not always be accurate, 
use ""data.vault_identity_group.*.group_name"", ""vault_identity_group.*.group_name"" instead */
        public Builder groupName(String groupName) {
            return groupName(Output.of(groupName));
        }

        /**
         * @param memberEntityIds List of member entities that belong to the group
         * 
         * @return builder
         * 
         */
        public Builder memberEntityIds(@Nullable Output<List<String>> memberEntityIds) {
            $.memberEntityIds = memberEntityIds;
            return this;
        }

        /**
         * @param memberEntityIds List of member entities that belong to the group
         * 
         * @return builder
         * 
         */
        public Builder memberEntityIds(List<String> memberEntityIds) {
            return memberEntityIds(Output.of(memberEntityIds));
        }

        /**
         * @param memberEntityIds List of member entities that belong to the group
         * 
         * @return builder
         * 
         */
        public Builder memberEntityIds(String... memberEntityIds) {
            return memberEntityIds(List.of(memberEntityIds));
        }

        public GroupMemberEntityIdsState build() {
            return $;
        }
    }

}