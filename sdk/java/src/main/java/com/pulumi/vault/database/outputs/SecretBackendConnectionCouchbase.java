// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SecretBackendConnectionCouchbase {
    /**
     * @return Required if `tls` is `true`. Specifies the certificate authority of the Couchbase server, as a PEM certificate that has been base64 encoded.
     * 
     */
    private final @Nullable String base64Pem;
    /**
     * @return Required for Couchbase versions prior to 6.5.0. This is only used to verify vault&#39;s connection to the server.
     * 
     */
    private final @Nullable String bucketName;
    /**
     * @return A set of Couchbase URIs to connect to. Must use `couchbases://` scheme if `tls` is `true`.
     * 
     */
    private final List<String> hosts;
    /**
     * @return Whether to skip verification of the server
     * certificate when using TLS.
     * 
     */
    private final @Nullable Boolean insecureTls;
    /**
     * @return The root credential password used in the connection URL.
     * 
     */
    private final String password;
    /**
     * @return Whether to use TLS when connecting to Cassandra.
     * 
     */
    private final @Nullable Boolean tls;
    /**
     * @return The root credential username used in the connection URL.
     * 
     */
    private final String username;
    /**
     * @return - [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated.
     * 
     */
    private final @Nullable String usernameTemplate;

    @CustomType.Constructor
    private SecretBackendConnectionCouchbase(
        @CustomType.Parameter("base64Pem") @Nullable String base64Pem,
        @CustomType.Parameter("bucketName") @Nullable String bucketName,
        @CustomType.Parameter("hosts") List<String> hosts,
        @CustomType.Parameter("insecureTls") @Nullable Boolean insecureTls,
        @CustomType.Parameter("password") String password,
        @CustomType.Parameter("tls") @Nullable Boolean tls,
        @CustomType.Parameter("username") String username,
        @CustomType.Parameter("usernameTemplate") @Nullable String usernameTemplate) {
        this.base64Pem = base64Pem;
        this.bucketName = bucketName;
        this.hosts = hosts;
        this.insecureTls = insecureTls;
        this.password = password;
        this.tls = tls;
        this.username = username;
        this.usernameTemplate = usernameTemplate;
    }

    /**
     * @return Required if `tls` is `true`. Specifies the certificate authority of the Couchbase server, as a PEM certificate that has been base64 encoded.
     * 
     */
    public Optional<String> base64Pem() {
        return Optional.ofNullable(this.base64Pem);
    }
    /**
     * @return Required for Couchbase versions prior to 6.5.0. This is only used to verify vault&#39;s connection to the server.
     * 
     */
    public Optional<String> bucketName() {
        return Optional.ofNullable(this.bucketName);
    }
    /**
     * @return A set of Couchbase URIs to connect to. Must use `couchbases://` scheme if `tls` is `true`.
     * 
     */
    public List<String> hosts() {
        return this.hosts;
    }
    /**
     * @return Whether to skip verification of the server
     * certificate when using TLS.
     * 
     */
    public Optional<Boolean> insecureTls() {
        return Optional.ofNullable(this.insecureTls);
    }
    /**
     * @return The root credential password used in the connection URL.
     * 
     */
    public String password() {
        return this.password;
    }
    /**
     * @return Whether to use TLS when connecting to Cassandra.
     * 
     */
    public Optional<Boolean> tls() {
        return Optional.ofNullable(this.tls);
    }
    /**
     * @return The root credential username used in the connection URL.
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

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SecretBackendConnectionCouchbase defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String base64Pem;
        private @Nullable String bucketName;
        private List<String> hosts;
        private @Nullable Boolean insecureTls;
        private String password;
        private @Nullable Boolean tls;
        private String username;
        private @Nullable String usernameTemplate;

        public Builder() {
    	      // Empty
        }

        public Builder(SecretBackendConnectionCouchbase defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.base64Pem = defaults.base64Pem;
    	      this.bucketName = defaults.bucketName;
    	      this.hosts = defaults.hosts;
    	      this.insecureTls = defaults.insecureTls;
    	      this.password = defaults.password;
    	      this.tls = defaults.tls;
    	      this.username = defaults.username;
    	      this.usernameTemplate = defaults.usernameTemplate;
        }

        public Builder base64Pem(@Nullable String base64Pem) {
            this.base64Pem = base64Pem;
            return this;
        }
        public Builder bucketName(@Nullable String bucketName) {
            this.bucketName = bucketName;
            return this;
        }
        public Builder hosts(List<String> hosts) {
            this.hosts = Objects.requireNonNull(hosts);
            return this;
        }
        public Builder hosts(String... hosts) {
            return hosts(List.of(hosts));
        }
        public Builder insecureTls(@Nullable Boolean insecureTls) {
            this.insecureTls = insecureTls;
            return this;
        }
        public Builder password(String password) {
            this.password = Objects.requireNonNull(password);
            return this;
        }
        public Builder tls(@Nullable Boolean tls) {
            this.tls = tls;
            return this;
        }
        public Builder username(String username) {
            this.username = Objects.requireNonNull(username);
            return this;
        }
        public Builder usernameTemplate(@Nullable String usernameTemplate) {
            this.usernameTemplate = usernameTemplate;
            return this;
        }        public SecretBackendConnectionCouchbase build() {
            return new SecretBackendConnectionCouchbase(base64Pem, bucketName, hosts, insecureTls, password, tls, username, usernameTemplate);
        }
    }
}