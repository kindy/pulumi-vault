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
public final class SecretBackendConnectionMysql {
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
     * @return x509 CA file for validating the certificate presented by the MySQL server. Must be PEM encoded.
     * 
     */
    private final @Nullable String tlsCa;
    /**
     * @return x509 certificate for connecting to the database. This must be a PEM encoded version of the private key and the certificate combined.
     * 
     */
    private final @Nullable String tlsCertificateKey;
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
    private SecretBackendConnectionMysql(
        @CustomType.Parameter("connectionUrl") @Nullable String connectionUrl,
        @CustomType.Parameter("maxConnectionLifetime") @Nullable Integer maxConnectionLifetime,
        @CustomType.Parameter("maxIdleConnections") @Nullable Integer maxIdleConnections,
        @CustomType.Parameter("maxOpenConnections") @Nullable Integer maxOpenConnections,
        @CustomType.Parameter("password") @Nullable String password,
        @CustomType.Parameter("tlsCa") @Nullable String tlsCa,
        @CustomType.Parameter("tlsCertificateKey") @Nullable String tlsCertificateKey,
        @CustomType.Parameter("username") @Nullable String username,
        @CustomType.Parameter("usernameTemplate") @Nullable String usernameTemplate) {
        this.connectionUrl = connectionUrl;
        this.maxConnectionLifetime = maxConnectionLifetime;
        this.maxIdleConnections = maxIdleConnections;
        this.maxOpenConnections = maxOpenConnections;
        this.password = password;
        this.tlsCa = tlsCa;
        this.tlsCertificateKey = tlsCertificateKey;
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
     * @return x509 CA file for validating the certificate presented by the MySQL server. Must be PEM encoded.
     * 
     */
    public Optional<String> tlsCa() {
        return Optional.ofNullable(this.tlsCa);
    }
    /**
     * @return x509 certificate for connecting to the database. This must be a PEM encoded version of the private key and the certificate combined.
     * 
     */
    public Optional<String> tlsCertificateKey() {
        return Optional.ofNullable(this.tlsCertificateKey);
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

    public static Builder builder(SecretBackendConnectionMysql defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String connectionUrl;
        private @Nullable Integer maxConnectionLifetime;
        private @Nullable Integer maxIdleConnections;
        private @Nullable Integer maxOpenConnections;
        private @Nullable String password;
        private @Nullable String tlsCa;
        private @Nullable String tlsCertificateKey;
        private @Nullable String username;
        private @Nullable String usernameTemplate;

        public Builder() {
    	      // Empty
        }

        public Builder(SecretBackendConnectionMysql defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.connectionUrl = defaults.connectionUrl;
    	      this.maxConnectionLifetime = defaults.maxConnectionLifetime;
    	      this.maxIdleConnections = defaults.maxIdleConnections;
    	      this.maxOpenConnections = defaults.maxOpenConnections;
    	      this.password = defaults.password;
    	      this.tlsCa = defaults.tlsCa;
    	      this.tlsCertificateKey = defaults.tlsCertificateKey;
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
        public Builder tlsCa(@Nullable String tlsCa) {
            this.tlsCa = tlsCa;
            return this;
        }
        public Builder tlsCertificateKey(@Nullable String tlsCertificateKey) {
            this.tlsCertificateKey = tlsCertificateKey;
            return this;
        }
        public Builder username(@Nullable String username) {
            this.username = username;
            return this;
        }
        public Builder usernameTemplate(@Nullable String usernameTemplate) {
            this.usernameTemplate = usernameTemplate;
            return this;
        }        public SecretBackendConnectionMysql build() {
            return new SecretBackendConnectionMysql(connectionUrl, maxConnectionLifetime, maxIdleConnections, maxOpenConnections, password, tlsCa, tlsCertificateKey, username, usernameTemplate);
        }
    }
}