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
    private @Nullable List<String> allowedRoles;
    /**
     * @return The path to a PEM-encoded CA cert file to use to verify the Elasticsearch server&#39;s identity.
     * 
     */
    private @Nullable String caCert;
    /**
     * @return The path to a directory of PEM-encoded CA cert files to use to verify the Elasticsearch server&#39;s identity.
     * 
     */
    private @Nullable String caPath;
    /**
     * @return The path to the certificate for the Elasticsearch client to present for communication.
     * 
     */
    private @Nullable String clientCert;
    /**
     * @return The path to the key for the Elasticsearch client to use for communication.
     * 
     */
    private @Nullable String clientKey;
    /**
     * @return A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
     * 
     * Supported list of database secrets engines that can be configured:
     * 
     */
    private @Nullable Map<String,Object> data;
    /**
     * @return Whether to disable certificate verification.
     * 
     */
    private @Nullable Boolean insecure;
    private String name;
    /**
     * @return The root credential password used in the connection URL.
     * 
     */
    private String password;
    /**
     * @return Specifies the name of the plugin to use.
     * 
     */
    private @Nullable String pluginName;
    /**
     * @return A list of database statements to be executed to rotate the root user&#39;s credentials.
     * 
     */
    private @Nullable List<String> rootRotationStatements;
    /**
     * @return This, if set, is used to set the SNI host when connecting via TLS.
     * 
     */
    private @Nullable String tlsServerName;
    /**
     * @return The URL for Elasticsearch&#39;s API. https requires certificate
     * by trusted CA if used.
     * 
     */
    private String url;
    /**
     * @return The root credential username used in the connection URL.
     * 
     */
    private String username;
    /**
     * @return [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated.
     * 
     */
    private @Nullable String usernameTemplate;
    /**
     * @return Whether the connection should be verified on
     * initial configuration or not.
     * 
     */
    private @Nullable Boolean verifyConnection;

    private SecretsMountElasticsearch() {}
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
     * Supported list of database secrets engines that can be configured:
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
     * @return The root credential password used in the connection URL.
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
     * @return The root credential username used in the connection URL.
     * 
     */
    public String username() {
        return this.username;
    }
    /**
     * @return [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated.
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
    @CustomType.Builder
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
        public Builder() {}
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

        @CustomType.Setter
        public Builder allowedRoles(@Nullable List<String> allowedRoles) {
            this.allowedRoles = allowedRoles;
            return this;
        }
        public Builder allowedRoles(String... allowedRoles) {
            return allowedRoles(List.of(allowedRoles));
        }
        @CustomType.Setter
        public Builder caCert(@Nullable String caCert) {
            this.caCert = caCert;
            return this;
        }
        @CustomType.Setter
        public Builder caPath(@Nullable String caPath) {
            this.caPath = caPath;
            return this;
        }
        @CustomType.Setter
        public Builder clientCert(@Nullable String clientCert) {
            this.clientCert = clientCert;
            return this;
        }
        @CustomType.Setter
        public Builder clientKey(@Nullable String clientKey) {
            this.clientKey = clientKey;
            return this;
        }
        @CustomType.Setter
        public Builder data(@Nullable Map<String,Object> data) {
            this.data = data;
            return this;
        }
        @CustomType.Setter
        public Builder insecure(@Nullable Boolean insecure) {
            this.insecure = insecure;
            return this;
        }
        @CustomType.Setter
        public Builder name(String name) {
            this.name = Objects.requireNonNull(name);
            return this;
        }
        @CustomType.Setter
        public Builder password(String password) {
            this.password = Objects.requireNonNull(password);
            return this;
        }
        @CustomType.Setter
        public Builder pluginName(@Nullable String pluginName) {
            this.pluginName = pluginName;
            return this;
        }
        @CustomType.Setter
        public Builder rootRotationStatements(@Nullable List<String> rootRotationStatements) {
            this.rootRotationStatements = rootRotationStatements;
            return this;
        }
        public Builder rootRotationStatements(String... rootRotationStatements) {
            return rootRotationStatements(List.of(rootRotationStatements));
        }
        @CustomType.Setter
        public Builder tlsServerName(@Nullable String tlsServerName) {
            this.tlsServerName = tlsServerName;
            return this;
        }
        @CustomType.Setter
        public Builder url(String url) {
            this.url = Objects.requireNonNull(url);
            return this;
        }
        @CustomType.Setter
        public Builder username(String username) {
            this.username = Objects.requireNonNull(username);
            return this;
        }
        @CustomType.Setter
        public Builder usernameTemplate(@Nullable String usernameTemplate) {
            this.usernameTemplate = usernameTemplate;
            return this;
        }
        @CustomType.Setter
        public Builder verifyConnection(@Nullable Boolean verifyConnection) {
            this.verifyConnection = verifyConnection;
            return this;
        }
        public SecretsMountElasticsearch build() {
            final var o = new SecretsMountElasticsearch();
            o.allowedRoles = allowedRoles;
            o.caCert = caCert;
            o.caPath = caPath;
            o.clientCert = clientCert;
            o.clientKey = clientKey;
            o.data = data;
            o.insecure = insecure;
            o.name = name;
            o.password = password;
            o.pluginName = pluginName;
            o.rootRotationStatements = rootRotationStatements;
            o.tlsServerName = tlsServerName;
            o.url = url;
            o.username = username;
            o.usernameTemplate = usernameTemplate;
            o.verifyConnection = verifyConnection;
            return o;
        }
    }
}
