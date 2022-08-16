// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.okta.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class AuthBackendUser {
    /**
     * @return List of Okta groups to associate with this user
     * 
     */
    private final List<String> groups;
    /**
     * @return List of Vault policies to associate with this user
     * 
     */
    private final @Nullable List<String> policies;
    /**
     * @return Name of the user within Okta
     * 
     */
    private final String username;

    @CustomType.Constructor
    private AuthBackendUser(
        @CustomType.Parameter("groups") List<String> groups,
        @CustomType.Parameter("policies") @Nullable List<String> policies,
        @CustomType.Parameter("username") String username) {
        this.groups = groups;
        this.policies = policies;
        this.username = username;
    }

    /**
     * @return List of Okta groups to associate with this user
     * 
     */
    public List<String> groups() {
        return this.groups;
    }
    /**
     * @return List of Vault policies to associate with this user
     * 
     */
    public List<String> policies() {
        return this.policies == null ? List.of() : this.policies;
    }
    /**
     * @return Name of the user within Okta
     * 
     */
    public String username() {
        return this.username;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AuthBackendUser defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private List<String> groups;
        private @Nullable List<String> policies;
        private String username;

        public Builder() {
    	      // Empty
        }

        public Builder(AuthBackendUser defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.groups = defaults.groups;
    	      this.policies = defaults.policies;
    	      this.username = defaults.username;
        }

        public Builder groups(List<String> groups) {
            this.groups = Objects.requireNonNull(groups);
            return this;
        }
        public Builder groups(String... groups) {
            return groups(List.of(groups));
        }
        public Builder policies(@Nullable List<String> policies) {
            this.policies = policies;
            return this;
        }
        public Builder policies(String... policies) {
            return policies(List.of(policies));
        }
        public Builder username(String username) {
            this.username = Objects.requireNonNull(username);
            return this;
        }        public AuthBackendUser build() {
            return new AuthBackendUser(groups, policies, username);
        }
    }
}