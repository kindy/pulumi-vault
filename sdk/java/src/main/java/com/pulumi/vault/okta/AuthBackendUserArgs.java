// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.okta;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthBackendUserArgs extends com.pulumi.resources.ResourceArgs {

    public static final AuthBackendUserArgs Empty = new AuthBackendUserArgs();

    /**
     * List of Okta groups to associate with this user
     * 
     */
    @Import(name="groups")
    private @Nullable Output<List<String>> groups;

    /**
     * @return List of Okta groups to associate with this user
     * 
     */
    public Optional<Output<List<String>>> groups() {
        return Optional.ofNullable(this.groups);
    }

    /**
     * The path where the Okta auth backend is mounted
     * 
     */
    @Import(name="path", required=true)
    private Output<String> path;

    /**
     * @return The path where the Okta auth backend is mounted
     * 
     */
    public Output<String> path() {
        return this.path;
    }

    /**
     * List of Vault policies to associate with this user
     * 
     */
    @Import(name="policies")
    private @Nullable Output<List<String>> policies;

    /**
     * @return List of Vault policies to associate with this user
     * 
     */
    public Optional<Output<List<String>>> policies() {
        return Optional.ofNullable(this.policies);
    }

    /**
     * Name of the user within Okta
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return Name of the user within Okta
     * 
     */
    public Output<String> username() {
        return this.username;
    }

    private AuthBackendUserArgs() {}

    private AuthBackendUserArgs(AuthBackendUserArgs $) {
        this.groups = $.groups;
        this.path = $.path;
        this.policies = $.policies;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthBackendUserArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthBackendUserArgs $;

        public Builder() {
            $ = new AuthBackendUserArgs();
        }

        public Builder(AuthBackendUserArgs defaults) {
            $ = new AuthBackendUserArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param groups List of Okta groups to associate with this user
         * 
         * @return builder
         * 
         */
        public Builder groups(@Nullable Output<List<String>> groups) {
            $.groups = groups;
            return this;
        }

        /**
         * @param groups List of Okta groups to associate with this user
         * 
         * @return builder
         * 
         */
        public Builder groups(List<String> groups) {
            return groups(Output.of(groups));
        }

        /**
         * @param groups List of Okta groups to associate with this user
         * 
         * @return builder
         * 
         */
        public Builder groups(String... groups) {
            return groups(List.of(groups));
        }

        /**
         * @param path The path where the Okta auth backend is mounted
         * 
         * @return builder
         * 
         */
        public Builder path(Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The path where the Okta auth backend is mounted
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param policies List of Vault policies to associate with this user
         * 
         * @return builder
         * 
         */
        public Builder policies(@Nullable Output<List<String>> policies) {
            $.policies = policies;
            return this;
        }

        /**
         * @param policies List of Vault policies to associate with this user
         * 
         * @return builder
         * 
         */
        public Builder policies(List<String> policies) {
            return policies(Output.of(policies));
        }

        /**
         * @param policies List of Vault policies to associate with this user
         * 
         * @return builder
         * 
         */
        public Builder policies(String... policies) {
            return policies(List.of(policies));
        }

        /**
         * @param username Name of the user within Okta
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username Name of the user within Okta
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public AuthBackendUserArgs build() {
            $.path = Objects.requireNonNull($.path, "expected parameter 'path' to be non-null");
            $.username = Objects.requireNonNull($.username, "expected parameter 'username' to be non-null");
            return $;
        }
    }

}