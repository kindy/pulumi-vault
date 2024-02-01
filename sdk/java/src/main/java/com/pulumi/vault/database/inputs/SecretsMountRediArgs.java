// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.database.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretsMountRediArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretsMountRediArgs Empty = new SecretsMountRediArgs();

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
     * The path to a PEM-encoded CA cert file to use to verify the Elasticsearch server&#39;s identity.
     * 
     */
    @Import(name="caCert")
    private @Nullable Output<String> caCert;

    /**
     * @return The path to a PEM-encoded CA cert file to use to verify the Elasticsearch server&#39;s identity.
     * 
     */
    public Optional<Output<String>> caCert() {
        return Optional.ofNullable(this.caCert);
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

    /**
     * The host to connect to.
     * 
     */
    @Import(name="host", required=true)
    private Output<String> host;

    /**
     * @return The host to connect to.
     * 
     */
    public Output<String> host() {
        return this.host;
    }

    /**
     * Whether to skip verification of the server
     * certificate when using TLS.
     * 
     */
    @Import(name="insecureTls")
    private @Nullable Output<Boolean> insecureTls;

    /**
     * @return Whether to skip verification of the server
     * certificate when using TLS.
     * 
     */
    public Optional<Output<Boolean>> insecureTls() {
        return Optional.ofNullable(this.insecureTls);
    }

    /**
     * Name of the database connection.
     * 
     */
    @Import(name="name", required=true)
    private Output<String> name;

    /**
     * @return Name of the database connection.
     * 
     */
    public Output<String> name() {
        return this.name;
    }

    /**
     * The root credential password used in the connection URL.
     * 
     */
    @Import(name="password", required=true)
    private Output<String> password;

    /**
     * @return The root credential password used in the connection URL.
     * 
     */
    public Output<String> password() {
        return this.password;
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
     * The default port to connect to if no port is specified as
     * part of the host.
     * 
     */
    @Import(name="port")
    private @Nullable Output<Integer> port;

    /**
     * @return The default port to connect to if no port is specified as
     * part of the host.
     * 
     */
    public Optional<Output<Integer>> port() {
        return Optional.ofNullable(this.port);
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
     * Whether to use TLS when connecting to Cassandra.
     * 
     */
    @Import(name="tls")
    private @Nullable Output<Boolean> tls;

    /**
     * @return Whether to use TLS when connecting to Cassandra.
     * 
     */
    public Optional<Output<Boolean>> tls() {
        return Optional.ofNullable(this.tls);
    }

    /**
     * The root credential username used in the connection URL.
     * 
     */
    @Import(name="username", required=true)
    private Output<String> username;

    /**
     * @return The root credential username used in the connection URL.
     * 
     */
    public Output<String> username() {
        return this.username;
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

    private SecretsMountRediArgs() {}

    private SecretsMountRediArgs(SecretsMountRediArgs $) {
        this.allowedRoles = $.allowedRoles;
        this.caCert = $.caCert;
        this.data = $.data;
        this.host = $.host;
        this.insecureTls = $.insecureTls;
        this.name = $.name;
        this.password = $.password;
        this.pluginName = $.pluginName;
        this.port = $.port;
        this.rootRotationStatements = $.rootRotationStatements;
        this.tls = $.tls;
        this.username = $.username;
        this.verifyConnection = $.verifyConnection;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretsMountRediArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretsMountRediArgs $;

        public Builder() {
            $ = new SecretsMountRediArgs();
        }

        public Builder(SecretsMountRediArgs defaults) {
            $ = new SecretsMountRediArgs(Objects.requireNonNull(defaults));
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
         * @param caCert The path to a PEM-encoded CA cert file to use to verify the Elasticsearch server&#39;s identity.
         * 
         * @return builder
         * 
         */
        public Builder caCert(@Nullable Output<String> caCert) {
            $.caCert = caCert;
            return this;
        }

        /**
         * @param caCert The path to a PEM-encoded CA cert file to use to verify the Elasticsearch server&#39;s identity.
         * 
         * @return builder
         * 
         */
        public Builder caCert(String caCert) {
            return caCert(Output.of(caCert));
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

        /**
         * @param host The host to connect to.
         * 
         * @return builder
         * 
         */
        public Builder host(Output<String> host) {
            $.host = host;
            return this;
        }

        /**
         * @param host The host to connect to.
         * 
         * @return builder
         * 
         */
        public Builder host(String host) {
            return host(Output.of(host));
        }

        /**
         * @param insecureTls Whether to skip verification of the server
         * certificate when using TLS.
         * 
         * @return builder
         * 
         */
        public Builder insecureTls(@Nullable Output<Boolean> insecureTls) {
            $.insecureTls = insecureTls;
            return this;
        }

        /**
         * @param insecureTls Whether to skip verification of the server
         * certificate when using TLS.
         * 
         * @return builder
         * 
         */
        public Builder insecureTls(Boolean insecureTls) {
            return insecureTls(Output.of(insecureTls));
        }

        /**
         * @param name Name of the database connection.
         * 
         * @return builder
         * 
         */
        public Builder name(Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the database connection.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param password The root credential password used in the connection URL.
         * 
         * @return builder
         * 
         */
        public Builder password(Output<String> password) {
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
         * @param port The default port to connect to if no port is specified as
         * part of the host.
         * 
         * @return builder
         * 
         */
        public Builder port(@Nullable Output<Integer> port) {
            $.port = port;
            return this;
        }

        /**
         * @param port The default port to connect to if no port is specified as
         * part of the host.
         * 
         * @return builder
         * 
         */
        public Builder port(Integer port) {
            return port(Output.of(port));
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
         * @param tls Whether to use TLS when connecting to Cassandra.
         * 
         * @return builder
         * 
         */
        public Builder tls(@Nullable Output<Boolean> tls) {
            $.tls = tls;
            return this;
        }

        /**
         * @param tls Whether to use TLS when connecting to Cassandra.
         * 
         * @return builder
         * 
         */
        public Builder tls(Boolean tls) {
            return tls(Output.of(tls));
        }

        /**
         * @param username The root credential username used in the connection URL.
         * 
         * @return builder
         * 
         */
        public Builder username(Output<String> username) {
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

        public SecretsMountRediArgs build() {
            if ($.host == null) {
                throw new MissingRequiredPropertyException("SecretsMountRediArgs", "host");
            }
            if ($.name == null) {
                throw new MissingRequiredPropertyException("SecretsMountRediArgs", "name");
            }
            if ($.password == null) {
                throw new MissingRequiredPropertyException("SecretsMountRediArgs", "password");
            }
            if ($.username == null) {
                throw new MissingRequiredPropertyException("SecretsMountRediArgs", "username");
            }
            return $;
        }
    }

}
