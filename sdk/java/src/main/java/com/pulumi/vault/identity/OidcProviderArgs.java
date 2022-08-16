// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OidcProviderArgs extends com.pulumi.resources.ResourceArgs {

    public static final OidcProviderArgs Empty = new OidcProviderArgs();

    /**
     * The client IDs that are permitted to use the provider.
     * If empty, no clients are allowed. If `*`, all clients are allowed.
     * 
     */
    @Import(name="allowedClientIds")
    private @Nullable Output<List<String>> allowedClientIds;

    /**
     * @return The client IDs that are permitted to use the provider.
     * If empty, no clients are allowed. If `*`, all clients are allowed.
     * 
     */
    public Optional<Output<List<String>>> allowedClientIds() {
        return Optional.ofNullable(this.allowedClientIds);
    }

    /**
     * Set to true if the issuer endpoint uses HTTPS.
     * 
     */
    @Import(name="httpsEnabled")
    private @Nullable Output<Boolean> httpsEnabled;

    /**
     * @return Set to true if the issuer endpoint uses HTTPS.
     * 
     */
    public Optional<Output<Boolean>> httpsEnabled() {
        return Optional.ofNullable(this.httpsEnabled);
    }

    /**
     * The host for the issuer. Can be either host or host:port.
     * 
     */
    @Import(name="issuerHost")
    private @Nullable Output<String> issuerHost;

    /**
     * @return The host for the issuer. Can be either host or host:port.
     * 
     */
    public Optional<Output<String>> issuerHost() {
        return Optional.ofNullable(this.issuerHost);
    }

    /**
     * The name of the provider.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the provider.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The scopes available for requesting on the provider.
     * 
     */
    @Import(name="scopesSupporteds")
    private @Nullable Output<List<String>> scopesSupporteds;

    /**
     * @return The scopes available for requesting on the provider.
     * 
     */
    public Optional<Output<List<String>>> scopesSupporteds() {
        return Optional.ofNullable(this.scopesSupporteds);
    }

    private OidcProviderArgs() {}

    private OidcProviderArgs(OidcProviderArgs $) {
        this.allowedClientIds = $.allowedClientIds;
        this.httpsEnabled = $.httpsEnabled;
        this.issuerHost = $.issuerHost;
        this.name = $.name;
        this.scopesSupporteds = $.scopesSupporteds;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OidcProviderArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OidcProviderArgs $;

        public Builder() {
            $ = new OidcProviderArgs();
        }

        public Builder(OidcProviderArgs defaults) {
            $ = new OidcProviderArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param allowedClientIds The client IDs that are permitted to use the provider.
         * If empty, no clients are allowed. If `*`, all clients are allowed.
         * 
         * @return builder
         * 
         */
        public Builder allowedClientIds(@Nullable Output<List<String>> allowedClientIds) {
            $.allowedClientIds = allowedClientIds;
            return this;
        }

        /**
         * @param allowedClientIds The client IDs that are permitted to use the provider.
         * If empty, no clients are allowed. If `*`, all clients are allowed.
         * 
         * @return builder
         * 
         */
        public Builder allowedClientIds(List<String> allowedClientIds) {
            return allowedClientIds(Output.of(allowedClientIds));
        }

        /**
         * @param allowedClientIds The client IDs that are permitted to use the provider.
         * If empty, no clients are allowed. If `*`, all clients are allowed.
         * 
         * @return builder
         * 
         */
        public Builder allowedClientIds(String... allowedClientIds) {
            return allowedClientIds(List.of(allowedClientIds));
        }

        /**
         * @param httpsEnabled Set to true if the issuer endpoint uses HTTPS.
         * 
         * @return builder
         * 
         */
        public Builder httpsEnabled(@Nullable Output<Boolean> httpsEnabled) {
            $.httpsEnabled = httpsEnabled;
            return this;
        }

        /**
         * @param httpsEnabled Set to true if the issuer endpoint uses HTTPS.
         * 
         * @return builder
         * 
         */
        public Builder httpsEnabled(Boolean httpsEnabled) {
            return httpsEnabled(Output.of(httpsEnabled));
        }

        /**
         * @param issuerHost The host for the issuer. Can be either host or host:port.
         * 
         * @return builder
         * 
         */
        public Builder issuerHost(@Nullable Output<String> issuerHost) {
            $.issuerHost = issuerHost;
            return this;
        }

        /**
         * @param issuerHost The host for the issuer. Can be either host or host:port.
         * 
         * @return builder
         * 
         */
        public Builder issuerHost(String issuerHost) {
            return issuerHost(Output.of(issuerHost));
        }

        /**
         * @param name The name of the provider.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the provider.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param scopesSupporteds The scopes available for requesting on the provider.
         * 
         * @return builder
         * 
         */
        public Builder scopesSupporteds(@Nullable Output<List<String>> scopesSupporteds) {
            $.scopesSupporteds = scopesSupporteds;
            return this;
        }

        /**
         * @param scopesSupporteds The scopes available for requesting on the provider.
         * 
         * @return builder
         * 
         */
        public Builder scopesSupporteds(List<String> scopesSupporteds) {
            return scopesSupporteds(Output.of(scopesSupporteds));
        }

        /**
         * @param scopesSupporteds The scopes available for requesting on the provider.
         * 
         * @return builder
         * 
         */
        public Builder scopesSupporteds(String... scopesSupporteds) {
            return scopesSupporteds(List.of(scopesSupporteds));
        }

        public OidcProviderArgs build() {
            return $;
        }
    }

}