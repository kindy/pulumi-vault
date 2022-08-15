// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SecretsMountPostgresql {
    /**
     * @return A list of roles that are allowed to use this
     * connection.
     * 
     */
    private final @Nullable List<String> allowedRoles;
    /**
     * @return A URL containing connection information.\
     * See [Vault docs](https://www.vaultproject.io/api-docs/secret/databases/snowflake#sample-payload)
     * 
     */
    private final @Nullable String connectionUrl;
    /**
     * @return A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
     * 
     */
    private final @Nullable Map<String,Object> data;
    /**
     * @return The maximum number of seconds to keep
     * a connection alive for.
     * 
     */
    private final @Nullable Integer maxConnectionLifetime;
    /**
     * @return The maximum number of idle connections to
     * maintain.
     * 
     */
    private final @Nullable Integer maxIdleConnections;
    /**
     * @return The maximum number of open connections to
     * use.
     * 
     */
    private final @Nullable Integer maxOpenConnections;
    private final String name;
    /**
     * @return The password to be used in the connection.
     * 
     */
    private final @Nullable String password;
    /**
     * @return Specifies the name of the plugin to use.
     * 
     */
    private final @Nullable String pluginName;
    /**
     * @return A list of database statements to be executed to rotate the root user&#39;s credentials.
     * 
     */
    private final @Nullable List<String> rootRotationStatements;
    /**
     * @return The username to be used in the connection (the account admin level).
     * 
     */
    private final @Nullable String username;
    /**
     * @return - [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated.
     * 
     */
    private final @Nullable String usernameTemplate;
    /**
     * @return Whether the connection should be verified on
     * initial configuration or not.
     * 
     */
    private final @Nullable Boolean verifyConnection;

    @CustomType.Constructor
    private SecretsMountPostgresql(
        @CustomType.Parameter("allowedRoles") @Nullable List<String> allowedRoles,
        @CustomType.Parameter("connectionUrl") @Nullable String connectionUrl,
        @CustomType.Parameter("data") @Nullable Map<String,Object> data,
        @CustomType.Parameter("maxConnectionLifetime") @Nullable Integer maxConnectionLifetime,
        @CustomType.Parameter("maxIdleConnections") @Nullable Integer maxIdleConnections,
        @CustomType.Parameter("maxOpenConnections") @Nullable Integer maxOpenConnections,
        @CustomType.Parameter("name") String name,
        @CustomType.Parameter("password") @Nullable String password,
        @CustomType.Parameter("pluginName") @Nullable String pluginName,
        @CustomType.Parameter("rootRotationStatements") @Nullable List<String> rootRotationStatements,
        @CustomType.Parameter("username") @Nullable String username,
        @CustomType.Parameter("usernameTemplate") @Nullable String usernameTemplate,
        @CustomType.Parameter("verifyConnection") @Nullable Boolean verifyConnection) {
        this.allowedRoles = allowedRoles;
        this.connectionUrl = connectionUrl;
        this.data = data;
        this.maxConnectionLifetime = maxConnectionLifetime;
        this.maxIdleConnections = maxIdleConnections;
        this.maxOpenConnections = maxOpenConnections;
        this.name = name;
        this.password = password;
        this.pluginName = pluginName;
        this.rootRotationStatements = rootRotationStatements;
        this.username = username;
        this.usernameTemplate = usernameTemplate;
        this.verifyConnection = verifyConnection;
    }

    /**
     * @return A list of roles that are allowed to use this
     * connection.
     * 
     */
    public List<String> allowedRoles() {
        return this.allowedRoles == null ? List.of() : this.allowedRoles;
    }
    /**
     * @return A URL containing connection information.\
     * See [Vault docs](https://www.vaultproject.io/api-docs/secret/databases/snowflake#sample-payload)
     * 
     */
    public Optional<String> connectionUrl() {
        return Optional.ofNullable(this.connectionUrl);
    }
    /**
     * @return A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
     * 
     */
    public Map<String,Object> data() {
        return this.data == null ? Map.of() : this.data;
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
    public String name() {
        return this.name;
    }
    /**
     * @return The password to be used in the connection.
     * 
     */
    public Optional<String> password() {
        return Optional.ofNullable(this.password);
    }
    /**
     * @return Specifies the name of the plugin to use.
     * 
     */
    public Optional<String> pluginName() {
        return Optional.ofNullable(this.pluginName);
    }
    /**
     * @return A list of database statements to be executed to rotate the root user&#39;s credentials.
     * 
     */
    public List<String> rootRotationStatements() {
        return this.rootRotationStatements == null ? List.of() : this.rootRotationStatements;
    }
    /**
     * @return The username to be used in the connection (the account admin level).
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
    /**
     * @return Whether the connection should be verified on
     * initial configuration or not.
     * 
     */
    public Optional<Boolean> verifyConnection() {
        return Optional.ofNullable(this.verifyConnection);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SecretsMountPostgresql defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable List<String> allowedRoles;
        private @Nullable String connectionUrl;
        private @Nullable Map<String,Object> data;
        private @Nullable Integer maxConnectionLifetime;
        private @Nullable Integer maxIdleConnections;
        private @Nullable Integer maxOpenConnections;
        private String name;
        private @Nullable String password;
        private @Nullable String pluginName;
        private @Nullable List<String> rootRotationStatements;
        private @Nullable String username;
        private @Nullable String usernameTemplate;
        private @Nullable Boolean verifyConnection;

        public Builder() {
    	      // Empty
        }

        public Builder(SecretsMountPostgresql defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowedRoles = defaults.allowedRoles;
    	      this.connectionUrl = defaults.connectionUrl;
    	      this.data = defaults.data;
    	      this.maxConnectionLifetime = defaults.maxConnectionLifetime;
    	      this.maxIdleConnections = defaults.maxIdleConnections;
    	      this.maxOpenConnections = defaults.maxOpenConnections;
    	      this.name = defaults.name;
    	      this.password = defaults.password;
    	      this.pluginName = defaults.pluginName;
    	      this.rootRotationStatements = defaults.rootRotationStatements;
    	      this.username = defaults.username;
    	      this.usernameTemplate = defaults.usernameTemplate;
    	      this.verifyConnection = defaults.verifyConnection;
        }

        public Builder allowedRoles(@Nullable List<String> allowedRoles) {
            this.allowedRoles = allowedRoles;
            return this;
        }
        public Builder allowedRoles(String... allowedRoles) {
            return allowedRoles(List.of(allowedRoles));
        }
        public Builder connectionUrl(@Nullable String connectionUrl) {
            this.connectionUrl = connectionUrl;
            return this;
        }
        public Builder data(@Nullable Map<String,Object> data) {
            this.data = data;
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
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public Builder password(@Nullable String password) {
            this.password = password;
            return this;
        }
        public Builder pluginName(@Nullable String pluginName) {
            this.pluginName = pluginName;
            return this;
        }
        public Builder rootRotationStatements(@Nullable List<String> rootRotationStatements) {
            this.rootRotationStatements = rootRotationStatements;
            return this;
        }
        public Builder rootRotationStatements(String... rootRotationStatements) {
            return rootRotationStatements(List.of(rootRotationStatements));
        }
        public Builder username(@Nullable String username) {
            this.username = username;
            return this;
        }
        public Builder usernameTemplate(@Nullable String usernameTemplate) {
            this.usernameTemplate = usernameTemplate;
            return this;
        }
        public Builder verifyConnection(@Nullable Boolean verifyConnection) {
            this.verifyConnection = verifyConnection;
            return this;
        }        public SecretsMountPostgresql build() {
            return new SecretsMountPostgresql(allowedRoles, connectionUrl, data, maxConnectionLifetime, maxIdleConnections, maxOpenConnections, name, password, pluginName, rootRotationStatements, username, usernameTemplate, verifyConnection);
        }
    }
}
