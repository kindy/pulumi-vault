// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretsMountRedisElasticachArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretsMountRedisElasticachArgs Empty = new SecretsMountRedisElasticachArgs();

    /**
     * A list of roles that are allowed to use this
     * connection.
     * 
     */
    @Import(name="allowedRoles")
    private @Nullable Output<List<String>> allowedRoles;

    /**
     * @return A list of roles that are allowed to use this
     * connection.
     * 
     */
    public Optional<Output<List<String>>> allowedRoles() {
        return Optional.ofNullable(this.allowedRoles);
    }

    /**
     * A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
     * 
     * Supported list of database secrets engines that can be configured:
     * 
     */
    @Import(name="data")
    private @Nullable Output<Map<String,Object>> data;

    /**
     * @return A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
     * 
     * Supported list of database secrets engines that can be configured:
     * 
     */
    public Optional<Output<Map<String,Object>>> data() {
        return Optional.ofNullable(this.data);
    }

    @Import(name="name", required=true)
    private Output<String> name;

    public Output<String> name() {
        return this.name;
    }

    /**
     * The root credential password used in the connection URL.
     * 
     */
    @Import(name="password")
    private @Nullable Output<String> password;

    /**
     * @return The root credential password used in the connection URL.
     * 
     */
    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    /**
     * Specifies the name of the plugin to use.
     * 
     */
    @Import(name="pluginName")
    private @Nullable Output<String> pluginName;

    /**
     * @return Specifies the name of the plugin to use.
     * 
     */
    public Optional<Output<String>> pluginName() {
        return Optional.ofNullable(this.pluginName);
    }

    /**
     * The AWS region where the ElastiCache cluster is hosted.
     * If omitted the plugin tries to infer the region from the environment.
     * 
     */
    @Import(name="region")
    private @Nullable Output<String> region;

    /**
     * @return The AWS region where the ElastiCache cluster is hosted.
     * If omitted the plugin tries to infer the region from the environment.
     * 
     */
    public Optional<Output<String>> region() {
        return Optional.ofNullable(this.region);
    }

    /**
     * A list of database statements to be executed to rotate the root user&#39;s credentials.
     * 
     */
    @Import(name="rootRotationStatements")
    private @Nullable Output<List<String>> rootRotationStatements;

    /**
     * @return A list of database statements to be executed to rotate the root user&#39;s credentials.
     * 
     */
    public Optional<Output<List<String>>> rootRotationStatements() {
        return Optional.ofNullable(this.rootRotationStatements);
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
     * The root credential username used in the connection URL.
     * 
     */
    @Import(name="username")
    private @Nullable Output<String> username;

    /**
     * @return The root credential username used in the connection URL.
     * 
     */
    public Optional<Output<String>> username() {
        return Optional.ofNullable(this.username);
    }

    /**
     * Whether the connection should be verified on
     * initial configuration or not.
     * 
     */
    @Import(name="verifyConnection")
    private @Nullable Output<Boolean> verifyConnection;

    /**
     * @return Whether the connection should be verified on
     * initial configuration or not.
     * 
     */
    public Optional<Output<Boolean>> verifyConnection() {
        return Optional.ofNullable(this.verifyConnection);
    }

    private SecretsMountRedisElasticachArgs() {}

    private SecretsMountRedisElasticachArgs(SecretsMountRedisElasticachArgs $) {
        this.allowedRoles = $.allowedRoles;
        this.data = $.data;
        this.name = $.name;
        this.password = $.password;
        this.pluginName = $.pluginName;
        this.region = $.region;
        this.rootRotationStatements = $.rootRotationStatements;
        this.url = $.url;
        this.username = $.username;
        this.verifyConnection = $.verifyConnection;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretsMountRedisElasticachArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretsMountRedisElasticachArgs $;

        public Builder() {
            $ = new SecretsMountRedisElasticachArgs();
        }

        public Builder(SecretsMountRedisElasticachArgs defaults) {
            $ = new SecretsMountRedisElasticachArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowedRoles A list of roles that are allowed to use this
         * connection.
         * 
         * @return builder
         * 
         */
        public Builder allowedRoles(@Nullable Output<List<String>> allowedRoles) {
            $.allowedRoles = allowedRoles;
            return this;
        }

        /**
         * @param allowedRoles A list of roles that are allowed to use this
         * connection.
         * 
         * @return builder
         * 
         */
        public Builder allowedRoles(List<String> allowedRoles) {
            return allowedRoles(Output.of(allowedRoles));
        }

        /**
         * @param allowedRoles A list of roles that are allowed to use this
         * connection.
         * 
         * @return builder
         * 
         */
        public Builder allowedRoles(String... allowedRoles) {
            return allowedRoles(List.of(allowedRoles));
        }

        /**
         * @param data A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
         * 
         * Supported list of database secrets engines that can be configured:
         * 
         * @return builder
         * 
         */
        public Builder data(@Nullable Output<Map<String,Object>> data) {
            $.data = data;
            return this;
        }

        /**
         * @param data A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
         * 
         * Supported list of database secrets engines that can be configured:
         * 
         * @return builder
         * 
         */
        public Builder data(Map<String,Object> data) {
            return data(Output.of(data));
        }

        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param password The root credential password used in the connection URL.
         * 
         * @return builder
         * 
         */
        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        /**
         * @param password The root credential password used in the connection URL.
         * 
         * @return builder
         * 
         */
        public Builder password(String password) {
            return password(Output.of(password));
        }

        /**
         * @param pluginName Specifies the name of the plugin to use.
         * 
         * @return builder
         * 
         */
        public Builder pluginName(@Nullable Output<String> pluginName) {
            $.pluginName = pluginName;
            return this;
        }

        /**
         * @param pluginName Specifies the name of the plugin to use.
         * 
         * @return builder
         * 
         */
        public Builder pluginName(String pluginName) {
            return pluginName(Output.of(pluginName));
        }

        /**
         * @param region The AWS region where the ElastiCache cluster is hosted.
         * If omitted the plugin tries to infer the region from the environment.
         * 
         * @return builder
         * 
         */
        public Builder region(@Nullable Output<String> region) {
            $.region = region;
            return this;
        }

        /**
         * @param region The AWS region where the ElastiCache cluster is hosted.
         * If omitted the plugin tries to infer the region from the environment.
         * 
         * @return builder
         * 
         */
        public Builder region(String region) {
            return region(Output.of(region));
        }

        /**
         * @param rootRotationStatements A list of database statements to be executed to rotate the root user&#39;s credentials.
         * 
         * @return builder
         * 
         */
        public Builder rootRotationStatements(@Nullable Output<List<String>> rootRotationStatements) {
            $.rootRotationStatements = rootRotationStatements;
            return this;
        }

        /**
         * @param rootRotationStatements A list of database statements to be executed to rotate the root user&#39;s credentials.
         * 
         * @return builder
         * 
         */
        public Builder rootRotationStatements(List<String> rootRotationStatements) {
            return rootRotationStatements(Output.of(rootRotationStatements));
        }

        /**
         * @param rootRotationStatements A list of database statements to be executed to rotate the root user&#39;s credentials.
         * 
         * @return builder
         * 
         */
        public Builder rootRotationStatements(String... rootRotationStatements) {
            return rootRotationStatements(List.of(rootRotationStatements));
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
         * @param username The root credential username used in the connection URL.
         * 
         * @return builder
         * 
         */
        public Builder username(@Nullable Output<String> username) {
            $.username = username;
            return this;
        }

        /**
         * @param username The root credential username used in the connection URL.
         * 
         * @return builder
         * 
         */
        public Builder username(String username) {
            return username(Output.of(username));
        }

        /**
         * @param verifyConnection Whether the connection should be verified on
         * initial configuration or not.
         * 
         * @return builder
         * 
         */
        public Builder verifyConnection(@Nullable Output<Boolean> verifyConnection) {
            $.verifyConnection = verifyConnection;
            return this;
        }

        /**
         * @param verifyConnection Whether the connection should be verified on
         * initial configuration or not.
         * 
         * @return builder
         * 
         */
        public Builder verifyConnection(Boolean verifyConnection) {
            return verifyConnection(Output.of(verifyConnection));
        }

        public SecretsMountRedisElasticachArgs build() {
            $.name = Objects.requireNonNull($.name, "expected parameter 'name' to be non-null");
            $.url = Objects.requireNonNull($.url, "expected parameter 'url' to be non-null");
            return $;
        }
    }

}
