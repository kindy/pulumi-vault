// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SecretsMountElasticsearch {
    /**
     * @return A list of roles that are allowed to use this
     * connection.
     * 
     */
    private final @Nullable List<String> allowedRoles;
    /**
     * @return The path to a PEM-encoded CA cert file to use to verify the Elasticsearch server&#39;s identity.
     * 
     */
    private final @Nullable String caCert;
    /**
     * @return The path to a directory of PEM-encoded CA cert files to use to verify the Elasticsearch server&#39;s identity.
     * 
     */
    private final @Nullable String caPath;
    /**
     * @return The path to the certificate for the Elasticsearch client to present for communication.
     * 
     */
    private final @Nullable String clientCert;
    /**
     * @return The path to the key for the Elasticsearch client to use for communication.
     * 
     */
    private final @Nullable String clientKey;
    /**
     * @return A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
     * 
     */
    private final @Nullable Map<String,Object> data;
    /**
     * @return Whether to disable certificate verification.
     * 
     */
    private final @Nullable Boolean insecure;
    private final String name;
    /**
     * @return The password to be used in the connection.
     * 
     */
    private final String password;
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
     * @return This, if set, is used to set the SNI host when connecting via TLS.
     * 
     */
    private final @Nullable String tlsServerName;
    /**
     * @return The URL for Elasticsearch&#39;s API. https requires certificate
     * by trusted CA if used.
     * 
     */
    private final String url;
    /**
     * @return The username to be used in the connection (the account admin level).
     * 
     */
    private final String username;
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
    private SecretsMountElasticsearch(
        @CustomType.Parameter("allowedRoles") @Nullable List<String> allowedRoles,
        @CustomType.Parameter("caCert") @Nullable String caCert,
        @CustomType.Parameter("caPath") @Nullable String caPath,
        @CustomType.Parameter("clientCert") @Nullable String clientCert,
        @CustomType.Parameter("clientKey") @Nullable String clientKey,
        @CustomType.Parameter("data") @Nullable Map<String,Object> data,
        @CustomType.Parameter("insecure") @Nullable Boolean insecure,
        @CustomType.Parameter("name") String name,
        @CustomType.Parameter("password") String password,
        @CustomType.Parameter("pluginName") @Nullable String pluginName,
        @CustomType.Parameter("rootRotationStatements") @Nullable List<String> rootRotationStatements,
        @CustomType.Parameter("tlsServerName") @Nullable String tlsServerName,
        @CustomType.Parameter("url") String url,
        @CustomType.Parameter("username") String username,
        @CustomType.Parameter("usernameTemplate") @Nullable String usernameTemplate,
        @CustomType.Parameter("verifyConnection") @Nullable Boolean verifyConnection) {
        this.allowedRoles = allowedRoles;
        this.caCert = caCert;
        this.caPath = caPath;
        this.clientCert = clientCert;
        this.clientKey = clientKey;
        this.data = data;
        this.insecure = insecure;
        this.name = name;
        this.password = password;
        this.pluginName = pluginName;
        this.rootRotationStatements = rootRotationStatements;
        this.tlsServerName = tlsServerName;
        this.url = url;
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
     * @return The path to a PEM-encoded CA cert file to use to verify the Elasticsearch server&#39;s identity.
     * 
     */
    public Optional<String> caCert() {
        return Optional.ofNullable(this.caCert);
    }
    /**
     * @return The path to a directory of PEM-encoded CA cert files to use to verify the Elasticsearch server&#39;s identity.
     * 
     */
    public Optional<String> caPath() {
        return Optional.ofNullable(this.caPath);
    }
    /**
     * @return The path to the certificate for the Elasticsearch client to present for communication.
     * 
     */
    public Optional<String> clientCert() {
        return Optional.ofNullable(this.clientCert);
    }
    /**
     * @return The path to the key for the Elasticsearch client to use for communication.
     * 
     */
    public Optional<String> clientKey() {
        return Optional.ofNullable(this.clientKey);
    }
    /**
     * @return A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
     * 
     */
    public Map<String,Object> data() {
        return this.data == null ? Map.of() : this.data;
    }
    /**
     * @return Whether to disable certificate verification.
     * 
     */
    public Optional<Boolean> insecure() {
        return Optional.ofNullable(this.insecure);
    }
    public String name() {
        return this.name;
    }
    /**
     * @return The password to be used in the connection.
     * 
     */
    public String password() {
        return this.password;
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
     * @return This, if set, is used to set the SNI host when connecting via TLS.
     * 
     */
    public Optional<String> tlsServerName() {
        return Optional.ofNullable(this.tlsServerName);
    }
    /**
     * @return The URL for Elasticsearch&#39;s API. https requires certificate
     * by trusted CA if used.
     * 
     */
    public String url() {
        return this.url;
    }
    /**
     * @return The username to be used in the connection (the account admin level).
     * 
     */
    public String username() {
        return this.username;
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

    public static Builder builder(SecretsMountElasticsearch defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable List<String> allowedRoles;
        private @Nullable String caCert;
        private @Nullable String caPath;
        private @Nullable String clientCert;
        private @Nullable String clientKey;
        private @Nullable Map<String,Object> data;
        private @Nullable Boolean insecure;
        private String name;
        private String password;
        private @Nullable String pluginName;
        private @Nullable List<String> rootRotationStatements;
        private @Nullable String tlsServerName;
        private String url;
        private String username;
        private @Nullable String usernameTemplate;
        private @Nullable Boolean verifyConnection;

        public Builder() {
    	      // Empty
        }

        public Builder(SecretsMountElasticsearch defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowedRoles = defaults.allowedRoles;
    	      this.caCert = defaults.caCert;
    	      this.caPath = defaults.caPath;
    	      this.clientCert = defaults.clientCert;
    	      this.clientKey = defaults.clientKey;
    	      this.data = defaults.data;
    	      this.insecure = defaults.insecure;
    	      this.name = defaults.name;
    	      this.password = defaults.password;
    	      this.pluginName = defaults.pluginName;
    	      this.rootRotationStatements = defaults.rootRotationStatements;
    	      this.tlsServerName = defaults.tlsServerName;
    	      this.url = defaults.url;
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
        public Builder caCert(@Nullable String caCert) {
            this.caCert = caCert;
            return this;
        }
        public Builder caPath(@Nullable String caPath) {
            this.caPath = caPath;
            return this;
        }
        public Builder clientCert(@Nullable String clientCert) {
            this.clientCert = clientCert;
            return this;
        }
        public Builder clientKey(@Nullable String clientKey) {
            this.clientKey = clientKey;
            return this;
        }
        public Builder data(@Nullable Map<String,Object> data) {
            this.data = data;
            return this;
        }
        public Builder insecure(@Nullable Boolean insecure) {
            this.insecure = insecure;
            return this;
        }
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        public Builder password(String password) {
            this.password = Objects.requireNonNull(password);
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
        public Builder tlsServerName(@Nullable String tlsServerName) {
            this.tlsServerName = tlsServerName;
            return this;
        }
        public Builder url(String url) {
            this.url = Objects.requireNonNull(url);
            return this;
        }
        public Builder username(String username) {
            this.username = Objects.requireNonNull(username);
            return this;
        }
        public Builder usernameTemplate(@Nullable String usernameTemplate) {
            this.usernameTemplate = usernameTemplate;
            return this;
        }
        public Builder verifyConnection(@Nullable Boolean verifyConnection) {
            this.verifyConnection = verifyConnection;
            return this;
        }        public SecretsMountElasticsearch build() {
            return new SecretsMountElasticsearch(allowedRoles, caCert, caPath, clientCert, clientKey, data, insecure, name, password, pluginName, rootRotationStatements, tlsServerName, url, username, usernameTemplate, verifyConnection);
        }
    }
}