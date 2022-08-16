// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.azure.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class BackendRoleAzureRole {
    private final @Nullable String roleId;
    private final String roleName;
    private final String scope;

    @CustomType.Constructor
    private BackendRoleAzureRole(
        @CustomType.Parameter("roleId") @Nullable String roleId,
        @CustomType.Parameter("roleName") String roleName,
        @CustomType.Parameter("scope") String scope) {
        this.roleId = roleId;
        this.roleName = roleName;
        this.scope = scope;
    }

    public Optional<String> roleId() {
        return Optional.ofNullable(this.roleId);
    }
    public String roleName() {
        return this.roleName;
    }
    public String scope() {
        return this.scope;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(BackendRoleAzureRole defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String roleId;
        private String roleName;
        private String scope;

        public Builder() {
    	      // Empty
        }

        public Builder(BackendRoleAzureRole defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.roleId = defaults.roleId;
    	      this.roleName = defaults.roleName;
    	      this.scope = defaults.scope;
        }

        public Builder roleId(@Nullable String roleId) {
            this.roleId = roleId;
            return this;
        }
        public Builder roleName(String roleName) {
            this.roleName = Objects.requireNonNull(roleName);
            return this;
        }
        public Builder scope(String scope) {
            this.scope = Objects.requireNonNull(scope);
            return this;
        }        public BackendRoleAzureRole build() {
            return new BackendRoleAzureRole(roleId, roleName, scope);
        }
    }
}