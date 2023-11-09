// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SecretBackendConnectionOracle {
    /**
     * @return A URL containing connection information. See
     * the [Vault
     * docs](https://www.vaultproject.io/api-docs/secret/databases/mongodb.html#sample-payload)
     * for an example.
     * 
     */
    private @Nullable String connectionUrl;
    /**
     * @return The maximum number of seconds to keep
     * a connection alive for.
     * 
     */
    private @Nullable Integer maxConnectionLifetime;
    /**
     * @return The maximum number of idle connections to
     * maintain.
     * 
     */
    private @Nullable Integer maxIdleConnections;
    /**
     * @return The maximum number of open connections to
     * use.
     * 
     */
    private @Nullable Integer maxOpenConnections;
    /**
     * @return The password to authenticate with.
     * 
     */
    private @Nullable String password;
    /**
     * @return The username to authenticate with.
     * 
     */
    private @Nullable String username;
    /**
     * @return Template describing how dynamic usernames are generated.
     * 
     */
    private @Nullable String usernameTemplate;

    private SecretBackendConnectionOracle() {}
    /**
     * @return A URL containing connection information. See
     * the [Vault
     * docs](https://www.vaultproject.io/api-docs/secret/databases/mongodb.html#sample-payload)
     * for an example.
     * 
     */
    public Optional<String> connectionUrl() {
        return Optional.ofNullable(this.connectionUrl);
    }
    /**
     * @return The maximum number of seconds to keep
     * a connection alive for.
     * 
     */
    public Optional<Integer> maxConnectionLifetime() {
        return Optional.ofNullable(this.maxConnectionLifetime);
    }
    /**
     * @return The maximum number of idle connections to
     * maintain.
     * 
     */
    public Optional<Integer> maxIdleConnections() {
        return Optional.ofNullable(this.maxIdleConnections);
    }
    /**
     * @return The maximum number of open connections to
     * use.
     * 
     */
    public Optional<Integer> maxOpenConnections() {
        return Optional.ofNullable(this.maxOpenConnections);
    }
    /**
     * @return The password to authenticate with.
     * 
     */
    public Optional<String> password() {
        return Optional.ofNullable(this.password);
    }
    /**
     * @return The username to authenticate with.
     * 
     */
    public Optional<String> username() {
        return Optional.ofNullable(this.username);
    }
    /**
     * @return Template describing how dynamic usernames are generated.
     * 
     */
    public Optional<String> usernameTemplate() {
        return Optional.ofNullable(this.usernameTemplate);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SecretBackendConnectionOracle defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String connectionUrl;
        private @Nullable Integer maxConnectionLifetime;
        private @Nullable Integer maxIdleConnections;
        private @Nullable Integer maxOpenConnections;
        private @Nullable String password;
        private @Nullable String username;
        private @Nullable String usernameTemplate;
        public Builder() {}
        public Builder(SecretBackendConnectionOracle defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.connectionUrl = defaults.connectionUrl;
    	      this.maxConnectionLifetime = defaults.maxConnectionLifetime;
    	      this.maxIdleConnections = defaults.maxIdleConnections;
    	      this.maxOpenConnections = defaults.maxOpenConnections;
    	      this.password = defaults.password;
    	      this.username = defaults.username;
    	      this.usernameTemplate = defaults.usernameTemplate;
        }

        @CustomType.Setter
        public Builder connectionUrl(@Nullable String connectionUrl) {
            this.connectionUrl = connectionUrl;
            return this;
        }
        @CustomType.Setter
        public Builder maxConnectionLifetime(@Nullable Integer maxConnectionLifetime) {
            this.maxConnectionLifetime = maxConnectionLifetime;
            return this;
        }
        @CustomType.Setter
        public Builder maxIdleConnections(@Nullable Integer maxIdleConnections) {
            this.maxIdleConnections = maxIdleConnections;
            return this;
        }
        @CustomType.Setter
        public Builder maxOpenConnections(@Nullable Integer maxOpenConnections) {
            this.maxOpenConnections = maxOpenConnections;
            return this;
        }
        @CustomType.Setter
        public Builder password(@Nullable String password) {
            this.password = password;
            return this;
        }
        @CustomType.Setter
        public Builder username(@Nullable String username) {
            this.username = username;
            return this;
        }
        @CustomType.Setter
        public Builder usernameTemplate(@Nullable String usernameTemplate) {
            this.usernameTemplate = usernameTemplate;
            return this;
        }
        public SecretBackendConnectionOracle build() {
            final var o = new SecretBackendConnectionOracle();
            o.connectionUrl = connectionUrl;
            o.maxConnectionLifetime = maxConnectionLifetime;
            o.maxIdleConnections = maxIdleConnections;
            o.maxOpenConnections = maxOpenConnections;
            o.password = password;
            o.username = username;
            o.usernameTemplate = usernameTemplate;
            return o;
        }
    }
}
