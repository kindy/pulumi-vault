// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.azure;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthBackendRoleArgs extends com.pulumi.resources.ResourceArgs {

    public static final AuthBackendRoleArgs Empty = new AuthBackendRoleArgs();

    /**
     * Unique name of the auth backend to configure.
     * 
     */
    @Import(name="backend")
    private @Nullable Output<String> backend;

    /**
     * @return Unique name of the auth backend to configure.
     * 
     */
    public Optional<Output<String>> backend() {
        return Optional.ofNullable(this.backend);
    }

    /**
     * If set, defines a constraint on the groups
     * that can perform the login operation that they should be using the group
     * ID specified by this field.
     * 
     */
    @Import(name="boundGroupIds")
    private @Nullable Output<List<String>> boundGroupIds;

    /**
     * @return If set, defines a constraint on the groups
     * that can perform the login operation that they should be using the group
     * ID specified by this field.
     * 
     */
    public Optional<Output<List<String>>> boundGroupIds() {
        return Optional.ofNullable(this.boundGroupIds);
    }

    /**
     * If set, defines a constraint on the virtual machines
     * that can perform the login operation that the location in their identity
     * document must match the one specified by this field.
     * 
     */
    @Import(name="boundLocations")
    private @Nullable Output<List<String>> boundLocations;

    /**
     * @return If set, defines a constraint on the virtual machines
     * that can perform the login operation that the location in their identity
     * document must match the one specified by this field.
     * 
     */
    public Optional<Output<List<String>>> boundLocations() {
        return Optional.ofNullable(this.boundLocations);
    }

    /**
     * If set, defines a constraint on the virtual
     * machines that can perform the login operation that they be associated with
     * the resource group that matches the value specified by this field.
     * 
     */
    @Import(name="boundResourceGroups")
    private @Nullable Output<List<String>> boundResourceGroups;

    /**
     * @return If set, defines a constraint on the virtual
     * machines that can perform the login operation that they be associated with
     * the resource group that matches the value specified by this field.
     * 
     */
    public Optional<Output<List<String>>> boundResourceGroups() {
        return Optional.ofNullable(this.boundResourceGroups);
    }

    /**
     * If set, defines a constraint on the virtual
     * machines that can perform the login operation that they must match the scale set
     * specified by this field.
     * 
     */
    @Import(name="boundScaleSets")
    private @Nullable Output<List<String>> boundScaleSets;

    /**
     * @return If set, defines a constraint on the virtual
     * machines that can perform the login operation that they must match the scale set
     * specified by this field.
     * 
     */
    public Optional<Output<List<String>>> boundScaleSets() {
        return Optional.ofNullable(this.boundScaleSets);
    }

    /**
     * If set, defines a constraint on the
     * service principals that can perform the login operation that they should be possess
     * the ids specified by this field.
     * 
     */
    @Import(name="boundServicePrincipalIds")
    private @Nullable Output<List<String>> boundServicePrincipalIds;

    /**
     * @return If set, defines a constraint on the
     * service principals that can perform the login operation that they should be possess
     * the ids specified by this field.
     * 
     */
    public Optional<Output<List<String>>> boundServicePrincipalIds() {
        return Optional.ofNullable(this.boundServicePrincipalIds);
    }

    /**
     * If set, defines a constraint on the subscriptions
     * that can perform the login operation to ones which  matches the value specified by this
     * field.
     * 
     */
    @Import(name="boundSubscriptionIds")
    private @Nullable Output<List<String>> boundSubscriptionIds;

    /**
     * @return If set, defines a constraint on the subscriptions
     * that can perform the login operation to ones which  matches the value specified by this
     * field.
     * 
     */
    public Optional<Output<List<String>>> boundSubscriptionIds() {
        return Optional.ofNullable(this.boundSubscriptionIds);
    }

    /**
     * The name of the role.
     * 
     */
    @Import(name="role", required=true)
    private Output<String> role;

    /**
     * @return The name of the role.
     * 
     */
    public Output<String> role() {
        return this.role;
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
     * The [maximum number](https://www.vaultproject.io/api-docs/azure#token_num_uses)
     * of times a generated token may be used (within its lifetime); 0 means unlimited.
     * 
     */
    @Import(name="tokenNumUses")
    private @Nullable Output<Integer> tokenNumUses;

    /**
     * @return The [maximum number](https://www.vaultproject.io/api-docs/azure#token_num_uses)
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

    private AuthBackendRoleArgs() {}

    private AuthBackendRoleArgs(AuthBackendRoleArgs $) {
        this.backend = $.backend;
        this.boundGroupIds = $.boundGroupIds;
        this.boundLocations = $.boundLocations;
        this.boundResourceGroups = $.boundResourceGroups;
        this.boundScaleSets = $.boundScaleSets;
        this.boundServicePrincipalIds = $.boundServicePrincipalIds;
        this.boundSubscriptionIds = $.boundSubscriptionIds;
        this.role = $.role;
        this.tokenBoundCidrs = $.tokenBoundCidrs;
        this.tokenExplicitMaxTtl = $.tokenExplicitMaxTtl;
        this.tokenMaxTtl = $.tokenMaxTtl;
        this.tokenNoDefaultPolicy = $.tokenNoDefaultPolicy;
        this.tokenNumUses = $.tokenNumUses;
        this.tokenPeriod = $.tokenPeriod;
        this.tokenPolicies = $.tokenPolicies;
        this.tokenTtl = $.tokenTtl;
        this.tokenType = $.tokenType;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthBackendRoleArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthBackendRoleArgs $;

        public Builder() {
            $ = new AuthBackendRoleArgs();
        }

        public Builder(AuthBackendRoleArgs defaults) {
            $ = new AuthBackendRoleArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend Unique name of the auth backend to configure.
         * 
         * @return builder
         * 
         */
        public Builder backend(@Nullable Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend Unique name of the auth backend to configure.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param boundGroupIds If set, defines a constraint on the groups
         * that can perform the login operation that they should be using the group
         * ID specified by this field.
         * 
         * @return builder
         * 
         */
        public Builder boundGroupIds(@Nullable Output<List<String>> boundGroupIds) {
            $.boundGroupIds = boundGroupIds;
            return this;
        }

        /**
         * @param boundGroupIds If set, defines a constraint on the groups
         * that can perform the login operation that they should be using the group
         * ID specified by this field.
         * 
         * @return builder
         * 
         */
        public Builder boundGroupIds(List<String> boundGroupIds) {
            return boundGroupIds(Output.of(boundGroupIds));
        }

        /**
         * @param boundGroupIds If set, defines a constraint on the groups
         * that can perform the login operation that they should be using the group
         * ID specified by this field.
         * 
         * @return builder
         * 
         */
        public Builder boundGroupIds(String... boundGroupIds) {
            return boundGroupIds(List.of(boundGroupIds));
        }

        /**
         * @param boundLocations If set, defines a constraint on the virtual machines
         * that can perform the login operation that the location in their identity
         * document must match the one specified by this field.
         * 
         * @return builder
         * 
         */
        public Builder boundLocations(@Nullable Output<List<String>> boundLocations) {
            $.boundLocations = boundLocations;
            return this;
        }

        /**
         * @param boundLocations If set, defines a constraint on the virtual machines
         * that can perform the login operation that the location in their identity
         * document must match the one specified by this field.
         * 
         * @return builder
         * 
         */
        public Builder boundLocations(List<String> boundLocations) {
            return boundLocations(Output.of(boundLocations));
        }

        /**
         * @param boundLocations If set, defines a constraint on the virtual machines
         * that can perform the login operation that the location in their identity
         * document must match the one specified by this field.
         * 
         * @return builder
         * 
         */
        public Builder boundLocations(String... boundLocations) {
            return boundLocations(List.of(boundLocations));
        }

        /**
         * @param boundResourceGroups If set, defines a constraint on the virtual
         * machines that can perform the login operation that they be associated with
         * the resource group that matches the value specified by this field.
         * 
         * @return builder
         * 
         */
        public Builder boundResourceGroups(@Nullable Output<List<String>> boundResourceGroups) {
            $.boundResourceGroups = boundResourceGroups;
            return this;
        }

        /**
         * @param boundResourceGroups If set, defines a constraint on the virtual
         * machines that can perform the login operation that they be associated with
         * the resource group that matches the value specified by this field.
         * 
         * @return builder
         * 
         */
        public Builder boundResourceGroups(List<String> boundResourceGroups) {
            return boundResourceGroups(Output.of(boundResourceGroups));
        }

        /**
         * @param boundResourceGroups If set, defines a constraint on the virtual
         * machines that can perform the login operation that they be associated with
         * the resource group that matches the value specified by this field.
         * 
         * @return builder
         * 
         */
        public Builder boundResourceGroups(String... boundResourceGroups) {
            return boundResourceGroups(List.of(boundResourceGroups));
        }

        /**
         * @param boundScaleSets If set, defines a constraint on the virtual
         * machines that can perform the login operation that they must match the scale set
         * specified by this field.
         * 
         * @return builder
         * 
         */
        public Builder boundScaleSets(@Nullable Output<List<String>> boundScaleSets) {
            $.boundScaleSets = boundScaleSets;
            return this;
        }

        /**
         * @param boundScaleSets If set, defines a constraint on the virtual
         * machines that can perform the login operation that they must match the scale set
         * specified by this field.
         * 
         * @return builder
         * 
         */
        public Builder boundScaleSets(List<String> boundScaleSets) {
            return boundScaleSets(Output.of(boundScaleSets));
        }

        /**
         * @param boundScaleSets If set, defines a constraint on the virtual
         * machines that can perform the login operation that they must match the scale set
         * specified by this field.
         * 
         * @return builder
         * 
         */
        public Builder boundScaleSets(String... boundScaleSets) {
            return boundScaleSets(List.of(boundScaleSets));
        }

        /**
         * @param boundServicePrincipalIds If set, defines a constraint on the
         * service principals that can perform the login operation that they should be possess
         * the ids specified by this field.
         * 
         * @return builder
         * 
         */
        public Builder boundServicePrincipalIds(@Nullable Output<List<String>> boundServicePrincipalIds) {
            $.boundServicePrincipalIds = boundServicePrincipalIds;
            return this;
        }

        /**
         * @param boundServicePrincipalIds If set, defines a constraint on the
         * service principals that can perform the login operation that they should be possess
         * the ids specified by this field.
         * 
         * @return builder
         * 
         */
        public Builder boundServicePrincipalIds(List<String> boundServicePrincipalIds) {
            return boundServicePrincipalIds(Output.of(boundServicePrincipalIds));
        }

        /**
         * @param boundServicePrincipalIds If set, defines a constraint on the
         * service principals that can perform the login operation that they should be possess
         * the ids specified by this field.
         * 
         * @return builder
         * 
         */
        public Builder boundServicePrincipalIds(String... boundServicePrincipalIds) {
            return boundServicePrincipalIds(List.of(boundServicePrincipalIds));
        }

        /**
         * @param boundSubscriptionIds If set, defines a constraint on the subscriptions
         * that can perform the login operation to ones which  matches the value specified by this
         * field.
         * 
         * @return builder
         * 
         */
        public Builder boundSubscriptionIds(@Nullable Output<List<String>> boundSubscriptionIds) {
            $.boundSubscriptionIds = boundSubscriptionIds;
            return this;
        }

        /**
         * @param boundSubscriptionIds If set, defines a constraint on the subscriptions
         * that can perform the login operation to ones which  matches the value specified by this
         * field.
         * 
         * @return builder
         * 
         */
        public Builder boundSubscriptionIds(List<String> boundSubscriptionIds) {
            return boundSubscriptionIds(Output.of(boundSubscriptionIds));
        }

        /**
         * @param boundSubscriptionIds If set, defines a constraint on the subscriptions
         * that can perform the login operation to ones which  matches the value specified by this
         * field.
         * 
         * @return builder
         * 
         */
        public Builder boundSubscriptionIds(String... boundSubscriptionIds) {
            return boundSubscriptionIds(List.of(boundSubscriptionIds));
        }

        /**
         * @param role The name of the role.
         * 
         * @return builder
         * 
         */
        public Builder role(Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role The name of the role.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
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
         * @param tokenNumUses The [maximum number](https://www.vaultproject.io/api-docs/azure#token_num_uses)
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
         * @param tokenNumUses The [maximum number](https://www.vaultproject.io/api-docs/azure#token_num_uses)
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

        public AuthBackendRoleArgs build() {
            $.role = Objects.requireNonNull($.role, "expected parameter 'role' to be non-null");
            return $;
        }
    }

}
