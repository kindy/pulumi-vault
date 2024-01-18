// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.gcp;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.vault.gcp.inputs.AuthBackendCustomEndpointArgs;
import com.pulumi.vault.gcp.inputs.AuthBackendTuneArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthBackendArgs extends com.pulumi.resources.ResourceArgs {

    public static final AuthBackendArgs Empty = new AuthBackendArgs();

    /**
     * The clients email associated with the credentials
     * 
     */
    @Import(name="clientEmail")
    private @Nullable Output<String> clientEmail;

    /**
     * @return The clients email associated with the credentials
     * 
     */
    public Optional<Output<String>> clientEmail() {
        return Optional.ofNullable(this.clientEmail);
    }

    /**
     * The Client ID of the credentials
     * 
     */
    @Import(name="clientId")
    private @Nullable Output<String> clientId;

    /**
     * @return The Client ID of the credentials
     * 
     */
    public Optional<Output<String>> clientId() {
        return Optional.ofNullable(this.clientId);
    }

    /**
     * A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
     * 
     */
    @Import(name="credentials")
    private @Nullable Output<String> credentials;

    /**
     * @return A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
     * 
     */
    public Optional<Output<String>> credentials() {
        return Optional.ofNullable(this.credentials);
    }

    /**
     * Specifies overrides to
     * [service endpoints](https://cloud.google.com/apis/design/glossary#api_service_endpoint)
     * used when making API requests. This allows specific requests made during authentication
     * to target alternative service endpoints for use in [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
     * environments. Requires Vault 1.11+.
     * 
     * Overrides are set at the subdomain level using the following keys:
     * 
     */
    @Import(name="customEndpoint")
    private @Nullable Output<AuthBackendCustomEndpointArgs> customEndpoint;

    /**
     * @return Specifies overrides to
     * [service endpoints](https://cloud.google.com/apis/design/glossary#api_service_endpoint)
     * used when making API requests. This allows specific requests made during authentication
     * to target alternative service endpoints for use in [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
     * environments. Requires Vault 1.11+.
     * 
     * Overrides are set at the subdomain level using the following keys:
     * 
     */
    public Optional<Output<AuthBackendCustomEndpointArgs>> customEndpoint() {
        return Optional.ofNullable(this.customEndpoint);
    }

    /**
     * A description of the auth method.
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return A description of the auth method.
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     * 
     */
    @Import(name="disableRemount")
    private @Nullable Output<Boolean> disableRemount;

    /**
     * @return If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     * 
     */
    public Optional<Output<Boolean>> disableRemount() {
        return Optional.ofNullable(this.disableRemount);
    }

    /**
     * Specifies if the auth method is local only.
     * 
     */
    @Import(name="local")
    private @Nullable Output<Boolean> local;

    /**
     * @return Specifies if the auth method is local only.
     * 
     */
    public Optional<Output<Boolean>> local() {
        return Optional.ofNullable(this.local);
    }

    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * The path to mount the auth method — this defaults to &#39;gcp&#39;.
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return The path to mount the auth method — this defaults to &#39;gcp&#39;.
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * The ID of the private key from the credentials
     * 
     */
    @Import(name="privateKeyId")
    private @Nullable Output<String> privateKeyId;

    /**
     * @return The ID of the private key from the credentials
     * 
     */
    public Optional<Output<String>> privateKeyId() {
        return Optional.ofNullable(this.privateKeyId);
    }

    /**
     * The GCP Project ID
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The GCP Project ID
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
    }

    /**
     * Extra configuration block. Structure is documented below.
     * 
     * The `tune` block is used to tune the auth backend:
     * 
     */
    @Import(name="tune")
    private @Nullable Output<AuthBackendTuneArgs> tune;

    /**
     * @return Extra configuration block. Structure is documented below.
     * 
     * The `tune` block is used to tune the auth backend:
     * 
     */
    public Optional<Output<AuthBackendTuneArgs>> tune() {
        return Optional.ofNullable(this.tune);
    }

    private AuthBackendArgs() {}

    private AuthBackendArgs(AuthBackendArgs $) {
        this.clientEmail = $.clientEmail;
        this.clientId = $.clientId;
        this.credentials = $.credentials;
        this.customEndpoint = $.customEndpoint;
        this.description = $.description;
        this.disableRemount = $.disableRemount;
        this.local = $.local;
        this.namespace = $.namespace;
        this.path = $.path;
        this.privateKeyId = $.privateKeyId;
        this.projectId = $.projectId;
        this.tune = $.tune;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthBackendArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthBackendArgs $;

        public Builder() {
            $ = new AuthBackendArgs();
        }

        public Builder(AuthBackendArgs defaults) {
            $ = new AuthBackendArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param clientEmail The clients email associated with the credentials
         * 
         * @return builder
         * 
         */
        public Builder clientEmail(@Nullable Output<String> clientEmail) {
            $.clientEmail = clientEmail;
            return this;
        }

        /**
         * @param clientEmail The clients email associated with the credentials
         * 
         * @return builder
         * 
         */
        public Builder clientEmail(String clientEmail) {
            return clientEmail(Output.of(clientEmail));
        }

        /**
         * @param clientId The Client ID of the credentials
         * 
         * @return builder
         * 
         */
        public Builder clientId(@Nullable Output<String> clientId) {
            $.clientId = clientId;
            return this;
        }

        /**
         * @param clientId The Client ID of the credentials
         * 
         * @return builder
         * 
         */
        public Builder clientId(String clientId) {
            return clientId(Output.of(clientId));
        }

        /**
         * @param credentials A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
         * 
         * @return builder
         * 
         */
        public Builder credentials(@Nullable Output<String> credentials) {
            $.credentials = credentials;
            return this;
        }

        /**
         * @param credentials A JSON string containing the contents of a GCP credentials file. If this value is empty, Vault will try to use Application Default Credentials from the machine on which the Vault server is running.
         * 
         * @return builder
         * 
         */
        public Builder credentials(String credentials) {
            return credentials(Output.of(credentials));
        }

        /**
         * @param customEndpoint Specifies overrides to
         * [service endpoints](https://cloud.google.com/apis/design/glossary#api_service_endpoint)
         * used when making API requests. This allows specific requests made during authentication
         * to target alternative service endpoints for use in [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
         * environments. Requires Vault 1.11+.
         * 
         * Overrides are set at the subdomain level using the following keys:
         * 
         * @return builder
         * 
         */
        public Builder customEndpoint(@Nullable Output<AuthBackendCustomEndpointArgs> customEndpoint) {
            $.customEndpoint = customEndpoint;
            return this;
        }

        /**
         * @param customEndpoint Specifies overrides to
         * [service endpoints](https://cloud.google.com/apis/design/glossary#api_service_endpoint)
         * used when making API requests. This allows specific requests made during authentication
         * to target alternative service endpoints for use in [Private Google Access](https://cloud.google.com/vpc/docs/configure-private-google-access)
         * environments. Requires Vault 1.11+.
         * 
         * Overrides are set at the subdomain level using the following keys:
         * 
         * @return builder
         * 
         */
        public Builder customEndpoint(AuthBackendCustomEndpointArgs customEndpoint) {
            return customEndpoint(Output.of(customEndpoint));
        }

        /**
         * @param description A description of the auth method.
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description A description of the auth method.
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        /**
         * @param disableRemount If set, opts out of mount migration on path updates.
         * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
         * 
         * @return builder
         * 
         */
        public Builder disableRemount(@Nullable Output<Boolean> disableRemount) {
            $.disableRemount = disableRemount;
            return this;
        }

        /**
         * @param disableRemount If set, opts out of mount migration on path updates.
         * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
         * 
         * @return builder
         * 
         */
        public Builder disableRemount(Boolean disableRemount) {
            return disableRemount(Output.of(disableRemount));
        }

        /**
         * @param local Specifies if the auth method is local only.
         * 
         * @return builder
         * 
         */
        public Builder local(@Nullable Output<Boolean> local) {
            $.local = local;
            return this;
        }

        /**
         * @param local Specifies if the auth method is local only.
         * 
         * @return builder
         * 
         */
        public Builder local(Boolean local) {
            return local(Output.of(local));
        }

        /**
         * @param namespace The namespace to provision the resource in.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param namespace The namespace to provision the resource in.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param path The path to mount the auth method — this defaults to &#39;gcp&#39;.
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The path to mount the auth method — this defaults to &#39;gcp&#39;.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param privateKeyId The ID of the private key from the credentials
         * 
         * @return builder
         * 
         */
        public Builder privateKeyId(@Nullable Output<String> privateKeyId) {
            $.privateKeyId = privateKeyId;
            return this;
        }

        /**
         * @param privateKeyId The ID of the private key from the credentials
         * 
         * @return builder
         * 
         */
        public Builder privateKeyId(String privateKeyId) {
            return privateKeyId(Output.of(privateKeyId));
        }

        /**
         * @param projectId The GCP Project ID
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The GCP Project ID
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
        }

        /**
         * @param tune Extra configuration block. Structure is documented below.
         * 
         * The `tune` block is used to tune the auth backend:
         * 
         * @return builder
         * 
         */
        public Builder tune(@Nullable Output<AuthBackendTuneArgs> tune) {
            $.tune = tune;
            return this;
        }

        /**
         * @param tune Extra configuration block. Structure is documented below.
         * 
         * The `tune` block is used to tune the auth backend:
         * 
         * @return builder
         * 
         */
        public Builder tune(AuthBackendTuneArgs tune) {
            return tune(Output.of(tune));
        }

        public AuthBackendArgs build() {
            return $;
        }
    }

}
