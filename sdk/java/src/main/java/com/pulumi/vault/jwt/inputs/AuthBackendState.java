// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.jwt.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.vault.jwt.inputs.AuthBackendTuneArgs;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthBackendState extends com.pulumi.resources.ResourceArgs {

    public static final AuthBackendState Empty = new AuthBackendState();

    /**
     * The accessor for this auth method
     * 
     */
    @Import(name="accessor")
    private @Nullable Output<String> accessor;

    /**
     * @return The accessor for this auth method
     * 
     */
    public Optional<Output<String>> accessor() {
        return Optional.ofNullable(this.accessor);
    }

    /**
     * The value against which to match the iss claim in a JWT
     * 
     */
    @Import(name="boundIssuer")
    private @Nullable Output<String> boundIssuer;

    /**
     * @return The value against which to match the iss claim in a JWT
     * 
     */
    public Optional<Output<String>> boundIssuer() {
        return Optional.ofNullable(this.boundIssuer);
    }

    /**
     * The default role to use if none is provided during login
     * 
     */
    @Import(name="defaultRole")
    private @Nullable Output<String> defaultRole;

    /**
     * @return The default role to use if none is provided during login
     * 
     */
    public Optional<Output<String>> defaultRole() {
        return Optional.ofNullable(this.defaultRole);
    }

    /**
     * The description of the auth backend
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return The description of the auth backend
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
     * The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
     * 
     */
    @Import(name="jwksCaPem")
    private @Nullable Output<String> jwksCaPem;

    /**
     * @return The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
     * 
     */
    public Optional<Output<String>> jwksCaPem() {
        return Optional.ofNullable(this.jwksCaPem);
    }

    /**
     * JWKS URL to use to authenticate signatures. Cannot be used with &#34;oidc_discovery_url&#34; or &#34;jwt_validation_pubkeys&#34;.
     * 
     */
    @Import(name="jwksUrl")
    private @Nullable Output<String> jwksUrl;

    /**
     * @return JWKS URL to use to authenticate signatures. Cannot be used with &#34;oidc_discovery_url&#34; or &#34;jwt_validation_pubkeys&#34;.
     * 
     */
    public Optional<Output<String>> jwksUrl() {
        return Optional.ofNullable(this.jwksUrl);
    }

    /**
     * A list of supported signing algorithms. Vault 1.1.0 defaults to [RS256] but future or past versions of Vault may differ
     * 
     */
    @Import(name="jwtSupportedAlgs")
    private @Nullable Output<List<String>> jwtSupportedAlgs;

    /**
     * @return A list of supported signing algorithms. Vault 1.1.0 defaults to [RS256] but future or past versions of Vault may differ
     * 
     */
    public Optional<Output<List<String>>> jwtSupportedAlgs() {
        return Optional.ofNullable(this.jwtSupportedAlgs);
    }

    /**
     * A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used in combination with `oidc_discovery_url`
     * 
     */
    @Import(name="jwtValidationPubkeys")
    private @Nullable Output<List<String>> jwtValidationPubkeys;

    /**
     * @return A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used in combination with `oidc_discovery_url`
     * 
     */
    public Optional<Output<List<String>>> jwtValidationPubkeys() {
        return Optional.ofNullable(this.jwtValidationPubkeys);
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
     * Pass namespace in the OIDC state parameter instead of as a separate query parameter. With this setting, the allowed redirect URL(s) in Vault and on the provider side should not contain a namespace query parameter. This means only one redirect URL entry needs to be maintained on the OIDC provider side for all vault namespaces that will be authenticating against it. Defaults to true for new configs
     * 
     * * tune - (Optional) Extra configuration block. Structure is documented below.
     * 
     * The `tune` block is used to tune the auth backend:
     * 
     */
    @Import(name="namespaceInState")
    private @Nullable Output<Boolean> namespaceInState;

    /**
     * @return Pass namespace in the OIDC state parameter instead of as a separate query parameter. With this setting, the allowed redirect URL(s) in Vault and on the provider side should not contain a namespace query parameter. This means only one redirect URL entry needs to be maintained on the OIDC provider side for all vault namespaces that will be authenticating against it. Defaults to true for new configs
     * 
     * * tune - (Optional) Extra configuration block. Structure is documented below.
     * 
     * The `tune` block is used to tune the auth backend:
     * 
     */
    public Optional<Output<Boolean>> namespaceInState() {
        return Optional.ofNullable(this.namespaceInState);
    }

    /**
     * Client ID used for OIDC backends
     * 
     */
    @Import(name="oidcClientId")
    private @Nullable Output<String> oidcClientId;

    /**
     * @return Client ID used for OIDC backends
     * 
     */
    public Optional<Output<String>> oidcClientId() {
        return Optional.ofNullable(this.oidcClientId);
    }

    /**
     * Client Secret used for OIDC backends
     * 
     */
    @Import(name="oidcClientSecret")
    private @Nullable Output<String> oidcClientSecret;

    /**
     * @return Client Secret used for OIDC backends
     * 
     */
    public Optional<Output<String>> oidcClientSecret() {
        return Optional.ofNullable(this.oidcClientSecret);
    }

    /**
     * The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
     * 
     */
    @Import(name="oidcDiscoveryCaPem")
    private @Nullable Output<String> oidcDiscoveryCaPem;

    /**
     * @return The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
     * 
     */
    public Optional<Output<String>> oidcDiscoveryCaPem() {
        return Optional.ofNullable(this.oidcDiscoveryCaPem);
    }

    /**
     * The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwt_validation_pubkeys`
     * 
     */
    @Import(name="oidcDiscoveryUrl")
    private @Nullable Output<String> oidcDiscoveryUrl;

    /**
     * @return The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwt_validation_pubkeys`
     * 
     */
    public Optional<Output<String>> oidcDiscoveryUrl() {
        return Optional.ofNullable(this.oidcDiscoveryUrl);
    }

    /**
     * The response mode to be used in the OAuth2 request. Allowed values are `query` and `form_post`. Defaults to `query`. If using Vault namespaces, and `oidc_response_mode` is `form_post`, then `namespace_in_state` should be set to `false`.
     * 
     */
    @Import(name="oidcResponseMode")
    private @Nullable Output<String> oidcResponseMode;

    /**
     * @return The response mode to be used in the OAuth2 request. Allowed values are `query` and `form_post`. Defaults to `query`. If using Vault namespaces, and `oidc_response_mode` is `form_post`, then `namespace_in_state` should be set to `false`.
     * 
     */
    public Optional<Output<String>> oidcResponseMode() {
        return Optional.ofNullable(this.oidcResponseMode);
    }

    /**
     * List of response types to request. Allowed values are &#39;code&#39; and &#39;id_token&#39;. Defaults to `[&#34;code&#34;]`. Note: `id_token` may only be used if `oidc_response_mode` is set to `form_post`.
     * 
     */
    @Import(name="oidcResponseTypes")
    private @Nullable Output<List<String>> oidcResponseTypes;

    /**
     * @return List of response types to request. Allowed values are &#39;code&#39; and &#39;id_token&#39;. Defaults to `[&#34;code&#34;]`. Note: `id_token` may only be used if `oidc_response_mode` is set to `form_post`.
     * 
     */
    public Optional<Output<List<String>>> oidcResponseTypes() {
        return Optional.ofNullable(this.oidcResponseTypes);
    }

    /**
     * Path to mount the JWT/OIDC auth backend
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return Path to mount the JWT/OIDC auth backend
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * Provider specific handling configuration. All values may be strings, and the provider will convert to the appropriate type when configuring Vault.
     * 
     */
    @Import(name="providerConfig")
    private @Nullable Output<Map<String,String>> providerConfig;

    /**
     * @return Provider specific handling configuration. All values may be strings, and the provider will convert to the appropriate type when configuring Vault.
     * 
     */
    public Optional<Output<Map<String,String>>> providerConfig() {
        return Optional.ofNullable(this.providerConfig);
    }

    @Import(name="tune")
    private @Nullable Output<AuthBackendTuneArgs> tune;

    public Optional<Output<AuthBackendTuneArgs>> tune() {
        return Optional.ofNullable(this.tune);
    }

    /**
     * Type of auth backend. Should be one of `jwt` or `oidc`. Default - `jwt`
     * 
     */
    @Import(name="type")
    private @Nullable Output<String> type;

    /**
     * @return Type of auth backend. Should be one of `jwt` or `oidc`. Default - `jwt`
     * 
     */
    public Optional<Output<String>> type() {
        return Optional.ofNullable(this.type);
    }

    private AuthBackendState() {}

    private AuthBackendState(AuthBackendState $) {
        this.accessor = $.accessor;
        this.boundIssuer = $.boundIssuer;
        this.defaultRole = $.defaultRole;
        this.description = $.description;
        this.disableRemount = $.disableRemount;
        this.jwksCaPem = $.jwksCaPem;
        this.jwksUrl = $.jwksUrl;
        this.jwtSupportedAlgs = $.jwtSupportedAlgs;
        this.jwtValidationPubkeys = $.jwtValidationPubkeys;
        this.local = $.local;
        this.namespace = $.namespace;
        this.namespaceInState = $.namespaceInState;
        this.oidcClientId = $.oidcClientId;
        this.oidcClientSecret = $.oidcClientSecret;
        this.oidcDiscoveryCaPem = $.oidcDiscoveryCaPem;
        this.oidcDiscoveryUrl = $.oidcDiscoveryUrl;
        this.oidcResponseMode = $.oidcResponseMode;
        this.oidcResponseTypes = $.oidcResponseTypes;
        this.path = $.path;
        this.providerConfig = $.providerConfig;
        this.tune = $.tune;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthBackendState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthBackendState $;

        public Builder() {
            $ = new AuthBackendState();
        }

        public Builder(AuthBackendState defaults) {
            $ = new AuthBackendState(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessor The accessor for this auth method
         * 
         * @return builder
         * 
         */
        public Builder accessor(@Nullable Output<String> accessor) {
            $.accessor = accessor;
            return this;
        }

        /**
         * @param accessor The accessor for this auth method
         * 
         * @return builder
         * 
         */
        public Builder accessor(String accessor) {
            return accessor(Output.of(accessor));
        }

        /**
         * @param boundIssuer The value against which to match the iss claim in a JWT
         * 
         * @return builder
         * 
         */
        public Builder boundIssuer(@Nullable Output<String> boundIssuer) {
            $.boundIssuer = boundIssuer;
            return this;
        }

        /**
         * @param boundIssuer The value against which to match the iss claim in a JWT
         * 
         * @return builder
         * 
         */
        public Builder boundIssuer(String boundIssuer) {
            return boundIssuer(Output.of(boundIssuer));
        }

        /**
         * @param defaultRole The default role to use if none is provided during login
         * 
         * @return builder
         * 
         */
        public Builder defaultRole(@Nullable Output<String> defaultRole) {
            $.defaultRole = defaultRole;
            return this;
        }

        /**
         * @param defaultRole The default role to use if none is provided during login
         * 
         * @return builder
         * 
         */
        public Builder defaultRole(String defaultRole) {
            return defaultRole(Output.of(defaultRole));
        }

        /**
         * @param description The description of the auth backend
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description The description of the auth backend
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
         * @param jwksCaPem The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
         * 
         * @return builder
         * 
         */
        public Builder jwksCaPem(@Nullable Output<String> jwksCaPem) {
            $.jwksCaPem = jwksCaPem;
            return this;
        }

        /**
         * @param jwksCaPem The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
         * 
         * @return builder
         * 
         */
        public Builder jwksCaPem(String jwksCaPem) {
            return jwksCaPem(Output.of(jwksCaPem));
        }

        /**
         * @param jwksUrl JWKS URL to use to authenticate signatures. Cannot be used with &#34;oidc_discovery_url&#34; or &#34;jwt_validation_pubkeys&#34;.
         * 
         * @return builder
         * 
         */
        public Builder jwksUrl(@Nullable Output<String> jwksUrl) {
            $.jwksUrl = jwksUrl;
            return this;
        }

        /**
         * @param jwksUrl JWKS URL to use to authenticate signatures. Cannot be used with &#34;oidc_discovery_url&#34; or &#34;jwt_validation_pubkeys&#34;.
         * 
         * @return builder
         * 
         */
        public Builder jwksUrl(String jwksUrl) {
            return jwksUrl(Output.of(jwksUrl));
        }

        /**
         * @param jwtSupportedAlgs A list of supported signing algorithms. Vault 1.1.0 defaults to [RS256] but future or past versions of Vault may differ
         * 
         * @return builder
         * 
         */
        public Builder jwtSupportedAlgs(@Nullable Output<List<String>> jwtSupportedAlgs) {
            $.jwtSupportedAlgs = jwtSupportedAlgs;
            return this;
        }

        /**
         * @param jwtSupportedAlgs A list of supported signing algorithms. Vault 1.1.0 defaults to [RS256] but future or past versions of Vault may differ
         * 
         * @return builder
         * 
         */
        public Builder jwtSupportedAlgs(List<String> jwtSupportedAlgs) {
            return jwtSupportedAlgs(Output.of(jwtSupportedAlgs));
        }

        /**
         * @param jwtSupportedAlgs A list of supported signing algorithms. Vault 1.1.0 defaults to [RS256] but future or past versions of Vault may differ
         * 
         * @return builder
         * 
         */
        public Builder jwtSupportedAlgs(String... jwtSupportedAlgs) {
            return jwtSupportedAlgs(List.of(jwtSupportedAlgs));
        }

        /**
         * @param jwtValidationPubkeys A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used in combination with `oidc_discovery_url`
         * 
         * @return builder
         * 
         */
        public Builder jwtValidationPubkeys(@Nullable Output<List<String>> jwtValidationPubkeys) {
            $.jwtValidationPubkeys = jwtValidationPubkeys;
            return this;
        }

        /**
         * @param jwtValidationPubkeys A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used in combination with `oidc_discovery_url`
         * 
         * @return builder
         * 
         */
        public Builder jwtValidationPubkeys(List<String> jwtValidationPubkeys) {
            return jwtValidationPubkeys(Output.of(jwtValidationPubkeys));
        }

        /**
         * @param jwtValidationPubkeys A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used in combination with `oidc_discovery_url`
         * 
         * @return builder
         * 
         */
        public Builder jwtValidationPubkeys(String... jwtValidationPubkeys) {
            return jwtValidationPubkeys(List.of(jwtValidationPubkeys));
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
         * @param namespaceInState Pass namespace in the OIDC state parameter instead of as a separate query parameter. With this setting, the allowed redirect URL(s) in Vault and on the provider side should not contain a namespace query parameter. This means only one redirect URL entry needs to be maintained on the OIDC provider side for all vault namespaces that will be authenticating against it. Defaults to true for new configs
         * 
         * * tune - (Optional) Extra configuration block. Structure is documented below.
         * 
         * The `tune` block is used to tune the auth backend:
         * 
         * @return builder
         * 
         */
        public Builder namespaceInState(@Nullable Output<Boolean> namespaceInState) {
            $.namespaceInState = namespaceInState;
            return this;
        }

        /**
         * @param namespaceInState Pass namespace in the OIDC state parameter instead of as a separate query parameter. With this setting, the allowed redirect URL(s) in Vault and on the provider side should not contain a namespace query parameter. This means only one redirect URL entry needs to be maintained on the OIDC provider side for all vault namespaces that will be authenticating against it. Defaults to true for new configs
         * 
         * * tune - (Optional) Extra configuration block. Structure is documented below.
         * 
         * The `tune` block is used to tune the auth backend:
         * 
         * @return builder
         * 
         */
        public Builder namespaceInState(Boolean namespaceInState) {
            return namespaceInState(Output.of(namespaceInState));
        }

        /**
         * @param oidcClientId Client ID used for OIDC backends
         * 
         * @return builder
         * 
         */
        public Builder oidcClientId(@Nullable Output<String> oidcClientId) {
            $.oidcClientId = oidcClientId;
            return this;
        }

        /**
         * @param oidcClientId Client ID used for OIDC backends
         * 
         * @return builder
         * 
         */
        public Builder oidcClientId(String oidcClientId) {
            return oidcClientId(Output.of(oidcClientId));
        }

        /**
         * @param oidcClientSecret Client Secret used for OIDC backends
         * 
         * @return builder
         * 
         */
        public Builder oidcClientSecret(@Nullable Output<String> oidcClientSecret) {
            $.oidcClientSecret = oidcClientSecret;
            return this;
        }

        /**
         * @param oidcClientSecret Client Secret used for OIDC backends
         * 
         * @return builder
         * 
         */
        public Builder oidcClientSecret(String oidcClientSecret) {
            return oidcClientSecret(Output.of(oidcClientSecret));
        }

        /**
         * @param oidcDiscoveryCaPem The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
         * 
         * @return builder
         * 
         */
        public Builder oidcDiscoveryCaPem(@Nullable Output<String> oidcDiscoveryCaPem) {
            $.oidcDiscoveryCaPem = oidcDiscoveryCaPem;
            return this;
        }

        /**
         * @param oidcDiscoveryCaPem The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used
         * 
         * @return builder
         * 
         */
        public Builder oidcDiscoveryCaPem(String oidcDiscoveryCaPem) {
            return oidcDiscoveryCaPem(Output.of(oidcDiscoveryCaPem));
        }

        /**
         * @param oidcDiscoveryUrl The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwt_validation_pubkeys`
         * 
         * @return builder
         * 
         */
        public Builder oidcDiscoveryUrl(@Nullable Output<String> oidcDiscoveryUrl) {
            $.oidcDiscoveryUrl = oidcDiscoveryUrl;
            return this;
        }

        /**
         * @param oidcDiscoveryUrl The OIDC Discovery URL, without any .well-known component (base path). Cannot be used in combination with `jwt_validation_pubkeys`
         * 
         * @return builder
         * 
         */
        public Builder oidcDiscoveryUrl(String oidcDiscoveryUrl) {
            return oidcDiscoveryUrl(Output.of(oidcDiscoveryUrl));
        }

        /**
         * @param oidcResponseMode The response mode to be used in the OAuth2 request. Allowed values are `query` and `form_post`. Defaults to `query`. If using Vault namespaces, and `oidc_response_mode` is `form_post`, then `namespace_in_state` should be set to `false`.
         * 
         * @return builder
         * 
         */
        public Builder oidcResponseMode(@Nullable Output<String> oidcResponseMode) {
            $.oidcResponseMode = oidcResponseMode;
            return this;
        }

        /**
         * @param oidcResponseMode The response mode to be used in the OAuth2 request. Allowed values are `query` and `form_post`. Defaults to `query`. If using Vault namespaces, and `oidc_response_mode` is `form_post`, then `namespace_in_state` should be set to `false`.
         * 
         * @return builder
         * 
         */
        public Builder oidcResponseMode(String oidcResponseMode) {
            return oidcResponseMode(Output.of(oidcResponseMode));
        }

        /**
         * @param oidcResponseTypes List of response types to request. Allowed values are &#39;code&#39; and &#39;id_token&#39;. Defaults to `[&#34;code&#34;]`. Note: `id_token` may only be used if `oidc_response_mode` is set to `form_post`.
         * 
         * @return builder
         * 
         */
        public Builder oidcResponseTypes(@Nullable Output<List<String>> oidcResponseTypes) {
            $.oidcResponseTypes = oidcResponseTypes;
            return this;
        }

        /**
         * @param oidcResponseTypes List of response types to request. Allowed values are &#39;code&#39; and &#39;id_token&#39;. Defaults to `[&#34;code&#34;]`. Note: `id_token` may only be used if `oidc_response_mode` is set to `form_post`.
         * 
         * @return builder
         * 
         */
        public Builder oidcResponseTypes(List<String> oidcResponseTypes) {
            return oidcResponseTypes(Output.of(oidcResponseTypes));
        }

        /**
         * @param oidcResponseTypes List of response types to request. Allowed values are &#39;code&#39; and &#39;id_token&#39;. Defaults to `[&#34;code&#34;]`. Note: `id_token` may only be used if `oidc_response_mode` is set to `form_post`.
         * 
         * @return builder
         * 
         */
        public Builder oidcResponseTypes(String... oidcResponseTypes) {
            return oidcResponseTypes(List.of(oidcResponseTypes));
        }

        /**
         * @param path Path to mount the JWT/OIDC auth backend
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path Path to mount the JWT/OIDC auth backend
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param providerConfig Provider specific handling configuration. All values may be strings, and the provider will convert to the appropriate type when configuring Vault.
         * 
         * @return builder
         * 
         */
        public Builder providerConfig(@Nullable Output<Map<String,String>> providerConfig) {
            $.providerConfig = providerConfig;
            return this;
        }

        /**
         * @param providerConfig Provider specific handling configuration. All values may be strings, and the provider will convert to the appropriate type when configuring Vault.
         * 
         * @return builder
         * 
         */
        public Builder providerConfig(Map<String,String> providerConfig) {
            return providerConfig(Output.of(providerConfig));
        }

        public Builder tune(@Nullable Output<AuthBackendTuneArgs> tune) {
            $.tune = tune;
            return this;
        }

        public Builder tune(AuthBackendTuneArgs tune) {
            return tune(Output.of(tune));
        }

        /**
         * @param type Type of auth backend. Should be one of `jwt` or `oidc`. Default - `jwt`
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type Type of auth backend. Should be one of `jwt` or `oidc`. Default - `jwt`
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public AuthBackendState build() {
            return $;
        }
    }

}
