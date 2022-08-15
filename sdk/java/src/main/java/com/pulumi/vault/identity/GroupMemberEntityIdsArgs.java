// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GroupMemberEntityIdsArgs extends com.pulumi.resources.ResourceArgs {

    public static final GroupMemberEntityIdsArgs Empty = new GroupMemberEntityIdsArgs();

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
    @Import(name="groupId", required=true)
    private Output<String> groupId;

    /**
     * @return Group ID to assign member entities to.
     * 
     */
    public Output<String> groupId() {
        return this.groupId;
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

    private GroupMemberEntityIdsArgs() {}

    private GroupMemberEntityIdsArgs(GroupMemberEntityIdsArgs $) {
        this.exclusive = $.exclusive;
        this.groupId = $.groupId;
        this.memberEntityIds = $.memberEntityIds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GroupMemberEntityIdsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GroupMemberEntityIdsArgs $;

        public Builder() {
            $ = new GroupMemberEntityIdsArgs();
        }

        public Builder(GroupMemberEntityIdsArgs defaults) {
            $ = new GroupMemberEntityIdsArgs(Objects.requireNonNull(defaults));
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
        public Builder groupId(Output<String> groupId) {
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

        public GroupMemberEntityIdsArgs build() {
            $.groupId = Objects.requireNonNull($.groupId, "expected parameter 'groupId' to be non-null");
            return $;
        }
    }

}
