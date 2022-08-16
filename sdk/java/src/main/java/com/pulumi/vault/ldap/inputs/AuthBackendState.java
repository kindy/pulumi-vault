// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.ldap.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthBackendState extends com.pulumi.resources.ResourceArgs {

    public static final AuthBackendState Empty = new AuthBackendState();

    /**
     * The accessor for this auth mount.
     * 
     */
    @Import(name="accessor")
    private @Nullable Output<String> accessor;

    /**
     * @return The accessor for this auth mount.
     * 
     */
    public Optional<Output<String>> accessor() {
        return Optional.ofNullable(this.accessor);
    }

    /**
     * DN of object to bind when performing user search
     * 
     */
    @Import(name="binddn")
    private @Nullable Output<String> binddn;

    /**
     * @return DN of object to bind when performing user search
     * 
     */
    public Optional<Output<String>> binddn() {
        return Optional.ofNullable(this.binddn);
    }

    /**
     * Password to use with `binddn` when performing user search
     * 
     */
    @Import(name="bindpass")
    private @Nullable Output<String> bindpass;

    /**
     * @return Password to use with `binddn` when performing user search
     * 
     */
    public Optional<Output<String>> bindpass() {
        return Optional.ofNullable(this.bindpass);
    }

    /**
     * Control case senstivity of objects fetched from LDAP, this is used for object matching in vault
     * 
     */
    @Import(name="caseSensitiveNames")
    private @Nullable Output<Boolean> caseSensitiveNames;

    /**
     * @return Control case senstivity of objects fetched from LDAP, this is used for object matching in vault
     * 
     */
    public Optional<Output<Boolean>> caseSensitiveNames() {
        return Optional.ofNullable(this.caseSensitiveNames);
    }

    /**
     * Trusted CA to validate TLS certificate
     * 
     */
    @Import(name="certificate")
    private @Nullable Output<String> certificate;

    /**
     * @return Trusted CA to validate TLS certificate
     * 
     */
    public Optional<Output<String>> certificate() {
        return Optional.ofNullable(this.certificate);
    }

    @Import(name="clientTlsCert")
    private @Nullable Output<String> clientTlsCert;

    public Optional<Output<String>> clientTlsCert() {
        return Optional.ofNullable(this.clientTlsCert);
    }

    @Import(name="clientTlsKey")
    private @Nullable Output<String> clientTlsKey;

    public Optional<Output<String>> clientTlsKey() {
        return Optional.ofNullable(this.clientTlsKey);
    }

    @Import(name="denyNullBind")
    private @Nullable Output<Boolean> denyNullBind;

    public Optional<Output<Boolean>> denyNullBind() {
        return Optional.ofNullable(this.denyNullBind);
    }

    /**
     * Description for the LDAP auth backend mount
     * 
     */
    @Import(name="description")
    private @Nullable Output<String> description;

    /**
     * @return Description for the LDAP auth backend mount
     * 
     */
    public Optional<Output<String>> description() {
        return Optional.ofNullable(this.description);
    }

    @Import(name="discoverdn")
    private @Nullable Output<Boolean> discoverdn;

    public Optional<Output<Boolean>> discoverdn() {
        return Optional.ofNullable(this.discoverdn);
    }

    /**
     * LDAP attribute to follow on objects returned by groupfilter
     * 
     */
    @Import(name="groupattr")
    private @Nullable Output<String> groupattr;

    /**
     * @return LDAP attribute to follow on objects returned by groupfilter
     * 
     */
    public Optional<Output<String>> groupattr() {
        return Optional.ofNullable(this.groupattr);
    }

    /**
     * Base DN under which to perform group search
     * 
     */
    @Import(name="groupdn")
    private @Nullable Output<String> groupdn;

    /**
     * @return Base DN under which to perform group search
     * 
     */
    public Optional<Output<String>> groupdn() {
        return Optional.ofNullable(this.groupdn);
    }

    /**
     * Go template used to construct group membership query
     * 
     */
    @Import(name="groupfilter")
    private @Nullable Output<String> groupfilter;

    /**
     * @return Go template used to construct group membership query
     * 
     */
    public Optional<Output<String>> groupfilter() {
        return Optional.ofNullable(this.groupfilter);
    }

    /**
     * Control whether or TLS certificates must be validated
     * 
     */
    @Import(name="insecureTls")
    private @Nullable Output<Boolean> insecureTls;

    /**
     * @return Control whether or TLS certificates must be validated
     * 
     */
    public Optional<Output<Boolean>> insecureTls() {
        return Optional.ofNullable(this.insecureTls);
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
     * Path to mount the LDAP auth backend under
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return Path to mount the LDAP auth backend under
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * Control use of TLS when conecting to LDAP
     * 
     */
    @Import(name="starttls")
    private @Nullable Output<Boolean> starttls;

    /**
     * @return Control use of TLS when conecting to LDAP
     * 
     */
    public Optional<Output<Boolean>> starttls() {
        return Optional.ofNullable(this.starttls);
    }

    /**
     * Maximum acceptable version of TLS
     * 
     */
    @Import(name="tlsMaxVersion")
    private @Nullable Output<String> tlsMaxVersion;

    /**
     * @return Maximum acceptable version of TLS
     * 
     */
    public Optional<Output<String>> tlsMaxVersion() {
        return Optional.ofNullable(this.tlsMaxVersion);
    }

    /**
     * Minimum acceptable version of TLS
     * 
     */
    @Import(name="tlsMinVersion")
    private @Nullable Output<String> tlsMinVersion;

    /**
     * @return Minimum acceptable version of TLS
     * 
     */
    public Optional<Output<String>> tlsMinVersion() {
        return Optional.ofNullable(this.tlsMinVersion);
    }

    /**
     * List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     * 
     */
    @Import(name="tokenBoundCidrs")
    private @Nullable Output<List<String>> tokenBoundCidrs;

    /**
     * @return List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     * 
     */
    public Optional<Output<List<String>>> tokenBoundCidrs() {
        return Optional.ofNullable(this.tokenBoundCidrs);
    }

    /**
     * If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `token_ttl` and
     * `token_max_ttl` would otherwise allow a renewal.
     * 
     */
    @Import(name="tokenExplicitMaxTtl")
    private @Nullable Output<Integer> tokenExplicitMaxTtl;

    /**
     * @return If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `token_ttl` and
     * `token_max_ttl` would otherwise allow a renewal.
     * 
     */
    public Optional<Output<Integer>> tokenExplicitMaxTtl() {
        return Optional.ofNullable(this.tokenExplicitMaxTtl);
    }

    /**
     * The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     * 
     */
    @Import(name="tokenMaxTtl")
    private @Nullable Output<Integer> tokenMaxTtl;

    /**
     * @return The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     * 
     */
    public Optional<Output<Integer>> tokenMaxTtl() {
        return Optional.ofNullable(this.tokenMaxTtl);
    }

    /**
     * If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     * 
     */
    @Import(name="tokenNoDefaultPolicy")
    private @Nullable Output<Boolean> tokenNoDefaultPolicy;

    /**
     * @return If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     * 
     */
    public Optional<Output<Boolean>> tokenNoDefaultPolicy() {
        return Optional.ofNullable(this.tokenNoDefaultPolicy);
    }

    /**
     * The [maximum number](https://www.vaultproject.io/api-docs/ldap#token_num_uses)
     * of times a generated token may be used (within its lifetime); 0 means unlimited.
     * 
     */
    @Import(name="tokenNumUses")
    private @Nullable Output<Integer> tokenNumUses;

    /**
     * @return The [maximum number](https://www.vaultproject.io/api-docs/ldap#token_num_uses)
     * of times a generated token may be used (within its lifetime); 0 means unlimited.
     * 
     */
    public Optional<Output<Integer>> tokenNumUses() {
        return Optional.ofNullable(this.tokenNumUses);
    }

    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token&#39;s TTL will be set to the
     * value of this field. Specified in seconds.
     * 
     */
    @Import(name="tokenPeriod")
    private @Nullable Output<Integer> tokenPeriod;

    /**
     * @return If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token&#39;s TTL will be set to the
     * value of this field. Specified in seconds.
     * 
     */
    public Optional<Output<Integer>> tokenPeriod() {
        return Optional.ofNullable(this.tokenPeriod);
    }

    /**
     * List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     * 
     */
    @Import(name="tokenPolicies")
    private @Nullable Output<List<String>> tokenPolicies;

    /**
     * @return List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     * 
     */
    public Optional<Output<List<String>>> tokenPolicies() {
        return Optional.ofNullable(this.tokenPolicies);
    }

    /**
     * The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     * 
     */
    @Import(name="tokenTtl")
    private @Nullable Output<Integer> tokenTtl;

    /**
     * @return The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     * 
     */
    public Optional<Output<Integer>> tokenTtl() {
        return Optional.ofNullable(this.tokenTtl);
    }

    /**
     * The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount&#39;s tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     * 
     */
    @Import(name="tokenType")
    private @Nullable Output<String> tokenType;

    /**
     * @return The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount&#39;s tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     * 
     */
    public Optional<Output<String>> tokenType() {
        return Optional.ofNullable(this.tokenType);
    }

    /**
     * The userPrincipalDomain used to construct UPN string
     * 
     */
    @Import(name="upndomain")
    private @Nullable Output<String> upndomain;

    /**
     * @return The userPrincipalDomain used to construct UPN string
     * 
     */
    public Optional<Output<String>> upndomain() {
        return Optional.ofNullable(this.upndomain);
    }

    /**
     * The URL of the LDAP server
     * 
     */
    @Import(name="url")
    private @Nullable Output<String> url;

    /**
     * @return The URL of the LDAP server
     * 
     */
    public Optional<Output<String>> url() {
        return Optional.ofNullable(this.url);
    }

    /**
     * Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships
     * 
     */
    @Import(name="useTokenGroups")
    private @Nullable Output<Boolean> useTokenGroups;

    /**
     * @return Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships
     * 
     */
    public Optional<Output<Boolean>> useTokenGroups() {
        return Optional.ofNullable(this.useTokenGroups);
    }

    /**
     * Attribute on user object matching username passed in
     * 
     */
    @Import(name="userattr")
    private @Nullable Output<String> userattr;

    /**
     * @return Attribute on user object matching username passed in
     * 
     */
    public Optional<Output<String>> userattr() {
        return Optional.ofNullable(this.userattr);
    }

    /**
     * Base DN under which to perform user search
     * 
     */
    @Import(name="userdn")
    private @Nullable Output<String> userdn;

    /**
     * @return Base DN under which to perform user search
     * 
     */
    public Optional<Output<String>> userdn() {
        return Optional.ofNullable(this.userdn);
    }

    /**
     * LDAP user search filter
     * 
     */
    @Import(name="userfilter")
    private @Nullable Output<String> userfilter;

    /**
     * @return LDAP user search filter
     * 
     */
    public Optional<Output<String>> userfilter() {
        return Optional.ofNullable(this.userfilter);
    }

    private AuthBackendState() {}

    private AuthBackendState(AuthBackendState $) {
        this.accessor = $.accessor;
        this.binddn = $.binddn;
        this.bindpass = $.bindpass;
        this.caseSensitiveNames = $.caseSensitiveNames;
        this.certificate = $.certificate;
        this.clientTlsCert = $.clientTlsCert;
        this.clientTlsKey = $.clientTlsKey;
        this.denyNullBind = $.denyNullBind;
        this.description = $.description;
        this.discoverdn = $.discoverdn;
        this.groupattr = $.groupattr;
        this.groupdn = $.groupdn;
        this.groupfilter = $.groupfilter;
        this.insecureTls = $.insecureTls;
        this.local = $.local;
        this.path = $.path;
        this.starttls = $.starttls;
        this.tlsMaxVersion = $.tlsMaxVersion;
        this.tlsMinVersion = $.tlsMinVersion;
        this.tokenBoundCidrs = $.tokenBoundCidrs;
        this.tokenExplicitMaxTtl = $.tokenExplicitMaxTtl;
        this.tokenMaxTtl = $.tokenMaxTtl;
        this.tokenNoDefaultPolicy = $.tokenNoDefaultPolicy;
        this.tokenNumUses = $.tokenNumUses;
        this.tokenPeriod = $.tokenPeriod;
        this.tokenPolicies = $.tokenPolicies;
        this.tokenTtl = $.tokenTtl;
        this.tokenType = $.tokenType;
        this.upndomain = $.upndomain;
        this.url = $.url;
        this.useTokenGroups = $.useTokenGroups;
        this.userattr = $.userattr;
        this.userdn = $.userdn;
        this.userfilter = $.userfilter;
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
         * @param accessor The accessor for this auth mount.
         * 
         * @return builder
         * 
         */
        public Builder accessor(@Nullable Output<String> accessor) {
            $.accessor = accessor;
            return this;
        }

        /**
         * @param accessor The accessor for this auth mount.
         * 
         * @return builder
         * 
         */
        public Builder accessor(String accessor) {
            return accessor(Output.of(accessor));
        }

        /**
         * @param binddn DN of object to bind when performing user search
         * 
         * @return builder
         * 
         */
        public Builder binddn(@Nullable Output<String> binddn) {
            $.binddn = binddn;
            return this;
        }

        /**
         * @param binddn DN of object to bind when performing user search
         * 
         * @return builder
         * 
         */
        public Builder binddn(String binddn) {
            return binddn(Output.of(binddn));
        }

        /**
         * @param bindpass Password to use with `binddn` when performing user search
         * 
         * @return builder
         * 
         */
        public Builder bindpass(@Nullable Output<String> bindpass) {
            $.bindpass = bindpass;
            return this;
        }

        /**
         * @param bindpass Password to use with `binddn` when performing user search
         * 
         * @return builder
         * 
         */
        public Builder bindpass(String bindpass) {
            return bindpass(Output.of(bindpass));
        }

        /**
         * @param caseSensitiveNames Control case senstivity of objects fetched from LDAP, this is used for object matching in vault
         * 
         * @return builder
         * 
         */
        public Builder caseSensitiveNames(@Nullable Output<Boolean> caseSensitiveNames) {
            $.caseSensitiveNames = caseSensitiveNames;
            return this;
        }

        /**
         * @param caseSensitiveNames Control case senstivity of objects fetched from LDAP, this is used for object matching in vault
         * 
         * @return builder
         * 
         */
        public Builder caseSensitiveNames(Boolean caseSensitiveNames) {
            return caseSensitiveNames(Output.of(caseSensitiveNames));
        }

        /**
         * @param certificate Trusted CA to validate TLS certificate
         * 
         * @return builder
         * 
         */
        public Builder certificate(@Nullable Output<String> certificate) {
            $.certificate = certificate;
            return this;
        }

        /**
         * @param certificate Trusted CA to validate TLS certificate
         * 
         * @return builder
         * 
         */
        public Builder certificate(String certificate) {
            return certificate(Output.of(certificate));
        }

        public Builder clientTlsCert(@Nullable Output<String> clientTlsCert) {
            $.clientTlsCert = clientTlsCert;
            return this;
        }

        public Builder clientTlsCert(String clientTlsCert) {
            return clientTlsCert(Output.of(clientTlsCert));
        }

        public Builder clientTlsKey(@Nullable Output<String> clientTlsKey) {
            $.clientTlsKey = clientTlsKey;
            return this;
        }

        public Builder clientTlsKey(String clientTlsKey) {
            return clientTlsKey(Output.of(clientTlsKey));
        }

        public Builder denyNullBind(@Nullable Output<Boolean> denyNullBind) {
            $.denyNullBind = denyNullBind;
            return this;
        }

        public Builder denyNullBind(Boolean denyNullBind) {
            return denyNullBind(Output.of(denyNullBind));
        }

        /**
         * @param description Description for the LDAP auth backend mount
         * 
         * @return builder
         * 
         */
        public Builder description(@Nullable Output<String> description) {
            $.description = description;
            return this;
        }

        /**
         * @param description Description for the LDAP auth backend mount
         * 
         * @return builder
         * 
         */
        public Builder description(String description) {
            return description(Output.of(description));
        }

        public Builder discoverdn(@Nullable Output<Boolean> discoverdn) {
            $.discoverdn = discoverdn;
            return this;
        }

        public Builder discoverdn(Boolean discoverdn) {
            return discoverdn(Output.of(discoverdn));
        }

        /**
         * @param groupattr LDAP attribute to follow on objects returned by groupfilter
         * 
         * @return builder
         * 
         */
        public Builder groupattr(@Nullable Output<String> groupattr) {
            $.groupattr = groupattr;
            return this;
        }

        /**
         * @param groupattr LDAP attribute to follow on objects returned by groupfilter
         * 
         * @return builder
         * 
         */
        public Builder groupattr(String groupattr) {
            return groupattr(Output.of(groupattr));
        }

        /**
         * @param groupdn Base DN under which to perform group search
         * 
         * @return builder
         * 
         */
        public Builder groupdn(@Nullable Output<String> groupdn) {
            $.groupdn = groupdn;
            return this;
        }

        /**
         * @param groupdn Base DN under which to perform group search
         * 
         * @return builder
         * 
         */
        public Builder groupdn(String groupdn) {
            return groupdn(Output.of(groupdn));
        }

        /**
         * @param groupfilter Go template used to construct group membership query
         * 
         * @return builder
         * 
         */
        public Builder groupfilter(@Nullable Output<String> groupfilter) {
            $.groupfilter = groupfilter;
            return this;
        }

        /**
         * @param groupfilter Go template used to construct group membership query
         * 
         * @return builder
         * 
         */
        public Builder groupfilter(String groupfilter) {
            return groupfilter(Output.of(groupfilter));
        }

        /**
         * @param insecureTls Control whether or TLS certificates must be validated
         * 
         * @return builder
         * 
         */
        public Builder insecureTls(@Nullable Output<Boolean> insecureTls) {
            $.insecureTls = insecureTls;
            return this;
        }

        /**
         * @param insecureTls Control whether or TLS certificates must be validated
         * 
         * @return builder
         * 
         */
        public Builder insecureTls(Boolean insecureTls) {
            return insecureTls(Output.of(insecureTls));
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
         * @param path Path to mount the LDAP auth backend under
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path Path to mount the LDAP auth backend under
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param starttls Control use of TLS when conecting to LDAP
         * 
         * @return builder
         * 
         */
        public Builder starttls(@Nullable Output<Boolean> starttls) {
            $.starttls = starttls;
            return this;
        }

        /**
         * @param starttls Control use of TLS when conecting to LDAP
         * 
         * @return builder
         * 
         */
        public Builder starttls(Boolean starttls) {
            return starttls(Output.of(starttls));
        }

        /**
         * @param tlsMaxVersion Maximum acceptable version of TLS
         * 
         * @return builder
         * 
         */
        public Builder tlsMaxVersion(@Nullable Output<String> tlsMaxVersion) {
            $.tlsMaxVersion = tlsMaxVersion;
            return this;
        }

        /**
         * @param tlsMaxVersion Maximum acceptable version of TLS
         * 
         * @return builder
         * 
         */
        public Builder tlsMaxVersion(String tlsMaxVersion) {
            return tlsMaxVersion(Output.of(tlsMaxVersion));
        }

        /**
         * @param tlsMinVersion Minimum acceptable version of TLS
         * 
         * @return builder
         * 
         */
        public Builder tlsMinVersion(@Nullable Output<String> tlsMinVersion) {
            $.tlsMinVersion = tlsMinVersion;
            return this;
        }

        /**
         * @param tlsMinVersion Minimum acceptable version of TLS
         * 
         * @return builder
         * 
         */
        public Builder tlsMinVersion(String tlsMinVersion) {
            return tlsMinVersion(Output.of(tlsMinVersion));
        }

        /**
         * @param tokenBoundCidrs List of CIDR blocks; if set, specifies blocks of IP
         * addresses which can authenticate successfully, and ties the resulting token to these blocks
         * as well.
         * 
         * @return builder
         * 
         */
        public Builder tokenBoundCidrs(@Nullable Output<List<String>> tokenBoundCidrs) {
            $.tokenBoundCidrs = tokenBoundCidrs;
            return this;
        }

        /**
         * @param tokenBoundCidrs List of CIDR blocks; if set, specifies blocks of IP
         * addresses which can authenticate successfully, and ties the resulting token to these blocks
         * as well.
         * 
         * @return builder
         * 
         */
        public Builder tokenBoundCidrs(List<String> tokenBoundCidrs) {
            return tokenBoundCidrs(Output.of(tokenBoundCidrs));
        }

        /**
         * @param tokenBoundCidrs List of CIDR blocks; if set, specifies blocks of IP
         * addresses which can authenticate successfully, and ties the resulting token to these blocks
         * as well.
         * 
         * @return builder
         * 
         */
        public Builder tokenBoundCidrs(String... tokenBoundCidrs) {
            return tokenBoundCidrs(List.of(tokenBoundCidrs));
        }

        /**
         * @param tokenExplicitMaxTtl If set, will encode an
         * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
         * onto the token in number of seconds. This is a hard cap even if `token_ttl` and
         * `token_max_ttl` would otherwise allow a renewal.
         * 
         * @return builder
         * 
         */
        public Builder tokenExplicitMaxTtl(@Nullable Output<Integer> tokenExplicitMaxTtl) {
            $.tokenExplicitMaxTtl = tokenExplicitMaxTtl;
            return this;
        }

        /**
         * @param tokenExplicitMaxTtl If set, will encode an
         * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
         * onto the token in number of seconds. This is a hard cap even if `token_ttl` and
         * `token_max_ttl` would otherwise allow a renewal.
         * 
         * @return builder
         * 
         */
        public Builder tokenExplicitMaxTtl(Integer tokenExplicitMaxTtl) {
            return tokenExplicitMaxTtl(Output.of(tokenExplicitMaxTtl));
        }

        /**
         * @param tokenMaxTtl The maximum lifetime for generated tokens in number of seconds.
         * Its current value will be referenced at renewal time.
         * 
         * @return builder
         * 
         */
        public Builder tokenMaxTtl(@Nullable Output<Integer> tokenMaxTtl) {
            $.tokenMaxTtl = tokenMaxTtl;
            return this;
        }

        /**
         * @param tokenMaxTtl The maximum lifetime for generated tokens in number of seconds.
         * Its current value will be referenced at renewal time.
         * 
         * @return builder
         * 
         */
        public Builder tokenMaxTtl(Integer tokenMaxTtl) {
            return tokenMaxTtl(Output.of(tokenMaxTtl));
        }

        /**
         * @param tokenNoDefaultPolicy If set, the default policy will not be set on
         * generated tokens; otherwise it will be added to the policies set in token_policies.
         * 
         * @return builder
         * 
         */
        public Builder tokenNoDefaultPolicy(@Nullable Output<Boolean> tokenNoDefaultPolicy) {
            $.tokenNoDefaultPolicy = tokenNoDefaultPolicy;
            return this;
        }

        /**
         * @param tokenNoDefaultPolicy If set, the default policy will not be set on
         * generated tokens; otherwise it will be added to the policies set in token_policies.
         * 
         * @return builder
         * 
         */
        public Builder tokenNoDefaultPolicy(Boolean tokenNoDefaultPolicy) {
            return tokenNoDefaultPolicy(Output.of(tokenNoDefaultPolicy));
        }

        /**
         * @param tokenNumUses The [maximum number](https://www.vaultproject.io/api-docs/ldap#token_num_uses)
         * of times a generated token may be used (within its lifetime); 0 means unlimited.
         * 
         * @return builder
         * 
         */
        public Builder tokenNumUses(@Nullable Output<Integer> tokenNumUses) {
            $.tokenNumUses = tokenNumUses;
            return this;
        }

        /**
         * @param tokenNumUses The [maximum number](https://www.vaultproject.io/api-docs/ldap#token_num_uses)
         * of times a generated token may be used (within its lifetime); 0 means unlimited.
         * 
         * @return builder
         * 
         */
        public Builder tokenNumUses(Integer tokenNumUses) {
            return tokenNumUses(Output.of(tokenNumUses));
        }

        /**
         * @param tokenPeriod If set, indicates that the
         * token generated using this role should never expire. The token should be renewed within the
         * duration specified by this value. At each renewal, the token&#39;s TTL will be set to the
         * value of this field. Specified in seconds.
         * 
         * @return builder
         * 
         */
        public Builder tokenPeriod(@Nullable Output<Integer> tokenPeriod) {
            $.tokenPeriod = tokenPeriod;
            return this;
        }

        /**
         * @param tokenPeriod If set, indicates that the
         * token generated using this role should never expire. The token should be renewed within the
         * duration specified by this value. At each renewal, the token&#39;s TTL will be set to the
         * value of this field. Specified in seconds.
         * 
         * @return builder
         * 
         */
        public Builder tokenPeriod(Integer tokenPeriod) {
            return tokenPeriod(Output.of(tokenPeriod));
        }

        /**
         * @param tokenPolicies List of policies to encode onto generated tokens. Depending
         * on the auth method, this list may be supplemented by user/group/other values.
         * 
         * @return builder
         * 
         */
        public Builder tokenPolicies(@Nullable Output<List<String>> tokenPolicies) {
            $.tokenPolicies = tokenPolicies;
            return this;
        }

        /**
         * @param tokenPolicies List of policies to encode onto generated tokens. Depending
         * on the auth method, this list may be supplemented by user/group/other values.
         * 
         * @return builder
         * 
         */
        public Builder tokenPolicies(List<String> tokenPolicies) {
            return tokenPolicies(Output.of(tokenPolicies));
        }

        /**
         * @param tokenPolicies List of policies to encode onto generated tokens. Depending
         * on the auth method, this list may be supplemented by user/group/other values.
         * 
         * @return builder
         * 
         */
        public Builder tokenPolicies(String... tokenPolicies) {
            return tokenPolicies(List.of(tokenPolicies));
        }

        /**
         * @param tokenTtl The incremental lifetime for generated tokens in number of seconds.
         * Its current value will be referenced at renewal time.
         * 
         * @return builder
         * 
         */
        public Builder tokenTtl(@Nullable Output<Integer> tokenTtl) {
            $.tokenTtl = tokenTtl;
            return this;
        }

        /**
         * @param tokenTtl The incremental lifetime for generated tokens in number of seconds.
         * Its current value will be referenced at renewal time.
         * 
         * @return builder
         * 
         */
        public Builder tokenTtl(Integer tokenTtl) {
            return tokenTtl(Output.of(tokenTtl));
        }

        /**
         * @param tokenType The type of token that should be generated. Can be `service`,
         * `batch`, or `default` to use the mount&#39;s tuned default (which unless changed will be
         * `service` tokens). For token store roles, there are two additional possibilities:
         * `default-service` and `default-batch` which specify the type to return unless the client
         * requests a different type at generation time.
         * 
         * @return builder
         * 
         */
        public Builder tokenType(@Nullable Output<String> tokenType) {
            $.tokenType = tokenType;
            return this;
        }

        /**
         * @param tokenType The type of token that should be generated. Can be `service`,
         * `batch`, or `default` to use the mount&#39;s tuned default (which unless changed will be
         * `service` tokens). For token store roles, there are two additional possibilities:
         * `default-service` and `default-batch` which specify the type to return unless the client
         * requests a different type at generation time.
         * 
         * @return builder
         * 
         */
        public Builder tokenType(String tokenType) {
            return tokenType(Output.of(tokenType));
        }

        /**
         * @param upndomain The userPrincipalDomain used to construct UPN string
         * 
         * @return builder
         * 
         */
        public Builder upndomain(@Nullable Output<String> upndomain) {
            $.upndomain = upndomain;
            return this;
        }

        /**
         * @param upndomain The userPrincipalDomain used to construct UPN string
         * 
         * @return builder
         * 
         */
        public Builder upndomain(String upndomain) {
            return upndomain(Output.of(upndomain));
        }

        /**
         * @param url The URL of the LDAP server
         * 
         * @return builder
         * 
         */
        public Builder url(@Nullable Output<String> url) {
            $.url = url;
            return this;
        }

        /**
         * @param url The URL of the LDAP server
         * 
         * @return builder
         * 
         */
        public Builder url(String url) {
            return url(Output.of(url));
        }

        /**
         * @param useTokenGroups Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships
         * 
         * @return builder
         * 
         */
        public Builder useTokenGroups(@Nullable Output<Boolean> useTokenGroups) {
            $.useTokenGroups = useTokenGroups;
            return this;
        }

        /**
         * @param useTokenGroups Use the Active Directory tokenGroups constructed attribute of the user to find the group memberships
         * 
         * @return builder
         * 
         */
        public Builder useTokenGroups(Boolean useTokenGroups) {
            return useTokenGroups(Output.of(useTokenGroups));
        }

        /**
         * @param userattr Attribute on user object matching username passed in
         * 
         * @return builder
         * 
         */
        public Builder userattr(@Nullable Output<String> userattr) {
            $.userattr = userattr;
            return this;
        }

        /**
         * @param userattr Attribute on user object matching username passed in
         * 
         * @return builder
         * 
         */
        public Builder userattr(String userattr) {
            return userattr(Output.of(userattr));
        }

        /**
         * @param userdn Base DN under which to perform user search
         * 
         * @return builder
         * 
         */
        public Builder userdn(@Nullable Output<String> userdn) {
            $.userdn = userdn;
            return this;
        }

        /**
         * @param userdn Base DN under which to perform user search
         * 
         * @return builder
         * 
         */
        public Builder userdn(String userdn) {
            return userdn(Output.of(userdn));
        }

        /**
         * @param userfilter LDAP user search filter
         * 
         * @return builder
         * 
         */
        public Builder userfilter(@Nullable Output<String> userfilter) {
            $.userfilter = userfilter;
            return this;
        }

        /**
         * @param userfilter LDAP user search filter
         * 
         * @return builder
         * 
         */
        public Builder userfilter(String userfilter) {
            return userfilter(Output.of(userfilter));
        }

        public AuthBackendState build() {
            return $;
        }
    }

}