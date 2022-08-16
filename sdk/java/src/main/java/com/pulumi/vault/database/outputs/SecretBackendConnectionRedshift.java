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
public final class SecretBackendConnectionRedshift {
    /**
     * @return Specifies the Redshift DSN. See
     * the [Vault
     * docs](https://www.vaultproject.io/api-docs/secret/databases/redshift#sample-payload)
     * for an example.
     * 
     */
    private final @Nullable String connectionUrl;
    /**
     * @return The maximum amount of time a connection may be reused.
     * 
     */
    private final @Nullable Integer maxConnectionLifetime;
    /**
     * @return The maximum number of idle connections to
     * the database.
     * 
     */
    private final @Nullable Integer maxIdleConnections;
    /**
     * @return The maximum number of open connections to
     * the database.
     * 
     */
    private final @Nullable Integer maxOpenConnections;
    /**
     * @return The root credential password used in the connection URL.
     * 
     */
    private final @Nullable String password;
    /**
     * @return The root credential username used in the connection URL.
     * 
     */
    private final @Nullable String username;
    /**
     * @return - [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated.
     * 
     */
    private final @Nullable String usernameTemplate;

    @CustomType.Constructor
    private SecretBackendConnectionRedshift(
        @CustomType.Parameter("connectionUrl") @Nullable String connectionUrl,
        @CustomType.Parameter("maxConnectionLifetime") @Nullable Integer maxConnectionLifetime,
        @CustomType.Parameter("maxIdleConnections") @Nullable Integer maxIdleConnections,
        @CustomType.Parameter("maxOpenConnections") @Nullable Integer maxOpenConnections,
        @CustomType.Parameter("password") @Nullable String password,
        @CustomType.Parameter("username") @Nullable String username,
        @CustomType.Parameter("usernameTemplate") @Nullable String usernameTemplate) {
        this.connectionUrl = connectionUrl;
        this.maxConnectionLifetime = maxConnectionLifetime;
        this.maxIdleConnections = maxIdleConnections;
        this.maxOpenConnections = maxOpenConnections;
        this.password = password;
        this.username = username;
        this.usernameTemplate = usernameTemplate;
    }

    /**
     * @return Specifies the Redshift DSN. See
     * the [Vault
     * docs](https://www.vaultproject.io/api-docs/secret/databases/redshift#sample-payload)
     * for an example.
     * 
     */
    public Optional<String> connectionUrl() {
        return Optional.ofNullable(this.connectionUrl);
    }
    /**
     * @return The maximum amount of time a connection may be reused.
     * 
     */
    public Optional<Integer> maxConnectionLifetime() {
        return Optional.ofNullable(this.maxConnectionLifetime);
    }
    /**
     * @return The maximum number of idle connections to
     * the database.
     * 
     */
    public Optional<Integer> maxIdleConnections() {
        return Optional.ofNullable(this.maxIdleConnections);
    }
    /**
     * @return The maximum number of open connections to
     * the database.
     * 
     */
    public Optional<Integer> maxOpenConnections() {
        return Optional.ofNullable(this.maxOpenConnections);
    }
    /**
     * @return The root credential password used in the connection URL.
     * 
     */
    public Optional<String> password() {
        return Optional.ofNullable(this.password);
    }
    /**
     * @return The root credential username used in the connection URL.
     * 
     */
    public Optional<String> username() {
        return Optional.ofNullable(this.username);
    }
    /**
     * @return - [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated.
     * 
     */
    public Optional<String> usernameTemplate() {
        return Optional.ofNullable(this.usernameTemplate);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SecretBackendConnectionRedshift defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String connectionUrl;
        private @Nullable Integer maxConnectionLifetime;
        private @Nullable Integer maxIdleConnections;
        private @Nullable Integer maxOpenConnections;
        private @Nullable String password;
        private @Nullable String username;
        private @Nullable String usernameTemplate;

        public Builder() {
    	      // Empty
        }

        public Builder(SecretBackendConnectionRedshift defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.connectionUrl = defaults.connectionUrl;
    	      this.maxConnectionLifetime = defaults.maxConnectionLifetime;
    	      this.maxIdleConnections = defaults.maxIdleConnections;
    	      this.maxOpenConnections = defaults.maxOpenConnections;
    	      this.password = defaults.password;
    	      this.username = defaults.username;
    	      this.usernameTemplate = defaults.usernameTemplate;
        }

        public Builder connectionUrl(@Nullable String connectionUrl) {
            this.connectionUrl = connectionUrl;
            return this;
        }
        public Builder maxConnectionLifetime(@Nullable Integer maxConnectionLifetime) {
            this.maxConnectionLifetime = maxConnectionLifetime;
            return this;
        }
        public Builder maxIdleConnections(@Nullable Integer maxIdleConnections) {
            this.maxIdleConnections = maxIdleConnections;
            return this;
        }
        public Builder maxOpenConnections(@Nullable Integer maxOpenConnections) {
            this.maxOpenConnections = maxOpenConnections;
            return this;
        }
        public Builder password(@Nullable String password) {
            this.password = password;
            return this;
        }
        public Builder username(@Nullable String username) {
            this.username = username;
            return this;
        }
        public Builder usernameTemplate(@Nullable String usernameTemplate) {
            this.usernameTemplate = usernameTemplate;
            return this;
        }        public SecretBackendConnectionRedshift build() {
            return new SecretBackendConnectionRedshift(connectionUrl, maxConnectionLifetime, maxIdleConnections, maxOpenConnections, password, username, usernameTemplate);
        }
    }
}