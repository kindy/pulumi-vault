// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretBackendConnectionRedisElasticacheArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretBackendConnectionRedisElasticacheArgs Empty = new SecretBackendConnectionRedisElasticacheArgs();

    /**
     * The password to authenticate with.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return The password to authenticate with.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * The region where the ElastiCache cluster is hosted. If omitted Vault tries to infer from the environment instead.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The region where the ElastiCache cluster is hosted. If omitted Vault tries to infer from the environment instead.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * The URL for Elasticsearch&#39;s API. https requires certificate
     * by trusted CA if used.
     * 
     */
    @Import(name="url", required=true)
    private Output<String> url;

    /**
     * @return The URL for Elasticsearch&#39;s API. https requires certificate
     * by trusted CA if used.
     * 
     */
    public Output<String> url() {
        return this.url;
    }

    /**
     * The username to authenticate with.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return The username to authenticate with.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    private SecretBackendConnectionRedisElasticacheArgs() {}

    private SecretBackendConnectionRedisElasticacheArgs(SecretBackendConnectionRedisElasticacheArgs $) {
        this.password = $.password;
        this.region = $.region;
        this.url = $.url;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretBackendConnectionRedisElasticacheArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretBackendConnectionRedisElasticacheArgs $;

        public Builder() {
            $ = new SecretBackendConnectionRedisElasticacheArgs();
        }

        public Builder(SecretBackendConnectionRedisElasticacheArgs defaults) {
            $ = new SecretBackendConnectionRedisElasticacheArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param password The password to authenticate with.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The password to authenticate with.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param region The region where the ElastiCache cluster is hosted. If omitted Vault tries to infer from the environment instead.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The region where the ElastiCache cluster is hosted. If omitted Vault tries to infer from the environment instead.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param url The URL for Elasticsearch&#39;s API. https requires certificate
         * by trusted CA if used.
         * 
         * @return builder
         * 
         */
        public Builder url(Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url The URL for Elasticsearch&#39;s API. https requires certificate
         * by trusted CA if used.
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        /**
         * @param username The username to authenticate with.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The username to authenticate with.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        public SecretBackendConnectionRedisElasticacheArgs build() {
            $.url = Objects.requireNonNull($.url, "expected parameter 'url' to be non-null");
            return $;
        }
    }

}
