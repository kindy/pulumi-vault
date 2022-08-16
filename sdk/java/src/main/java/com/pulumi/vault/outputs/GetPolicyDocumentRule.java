// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.vault.outputs.GetPolicyDocumentRuleAllowedParameter;
import com.pulumi.vault.outputs.GetPolicyDocumentRuleDeniedParameter;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetPolicyDocumentRule {
    /**
     * @return Whitelists a list of keys and values that are permitted on the given path. See Parameters below.
     * 
     */
    private final @Nullable List<GetPolicyDocumentRuleAllowedParameter> allowedParameters;
    /**
     * @return A list of capabilities that this rule apply to `path`. For example, [&#34;read&#34;, &#34;write&#34;].
     * 
     */
    private final List<String> capabilities;
    /**
     * @return Blacklists a list of parameter and values. Any values specified here take precedence over `allowed_parameter`. See Parameters below.
     * 
     */
    private final @Nullable List<GetPolicyDocumentRuleDeniedParameter> deniedParameters;
    /**
     * @return Description of the rule. Will be added as a comment to rendered rule.
     * 
     */
    private final @Nullable String description;
    /**
     * @return The maximum allowed TTL that clients can specify for a wrapped response.
     * 
     */
    private final @Nullable String maxWrappingTtl;
    /**
     * @return The minimum allowed TTL that clients can specify for a wrapped response.
     * 
     */
    private final @Nullable String minWrappingTtl;
    /**
     * @return A path in Vault that this rule applies to.
     * 
     */
    private final String path;
    /**
     * @return A list of parameters that must be specified.
     * 
     */
    private final @Nullable List<String> requiredParameters;

    @CustomType.Constructor
    private GetPolicyDocumentRule(
        @CustomType.Parameter("allowedParameters") @Nullable List<GetPolicyDocumentRuleAllowedParameter> allowedParameters,
        @CustomType.Parameter("capabilities") List<String> capabilities,
        @CustomType.Parameter("deniedParameters") @Nullable List<GetPolicyDocumentRuleDeniedParameter> deniedParameters,
        @CustomType.Parameter("description") @Nullable String description,
        @CustomType.Parameter("maxWrappingTtl") @Nullable String maxWrappingTtl,
        @CustomType.Parameter("minWrappingTtl") @Nullable String minWrappingTtl,
        @CustomType.Parameter("path") String path,
        @CustomType.Parameter("requiredParameters") @Nullable List<String> requiredParameters) {
        this.allowedParameters = allowedParameters;
        this.capabilities = capabilities;
        this.deniedParameters = deniedParameters;
        this.description = description;
        this.maxWrappingTtl = maxWrappingTtl;
        this.minWrappingTtl = minWrappingTtl;
        this.path = path;
        this.requiredParameters = requiredParameters;
    }

    /**
     * @return Whitelists a list of keys and values that are permitted on the given path. See Parameters below.
     * 
     */
    public List<GetPolicyDocumentRuleAllowedParameter> allowedParameters() {
        return this.allowedParameters == null ? List.of() : this.allowedParameters;
    }
    /**
     * @return A list of capabilities that this rule apply to `path`. For example, [&#34;read&#34;, &#34;write&#34;].
     * 
     */
    public List<String> capabilities() {
        return this.capabilities;
    }
    /**
     * @return Blacklists a list of parameter and values. Any values specified here take precedence over `allowed_parameter`. See Parameters below.
     * 
     */
    public List<GetPolicyDocumentRuleDeniedParameter> deniedParameters() {
        return this.deniedParameters == null ? List.of() : this.deniedParameters;
    }
    /**
     * @return Description of the rule. Will be added as a comment to rendered rule.
     * 
     */
    public Optional<String> description() {
        return Optional.ofNullable(this.description);
    }
    /**
     * @return The maximum allowed TTL that clients can specify for a wrapped response.
     * 
     */
    public Optional<String> maxWrappingTtl() {
        return Optional.ofNullable(this.maxWrappingTtl);
    }
    /**
     * @return The minimum allowed TTL that clients can specify for a wrapped response.
     * 
     */
    public Optional<String> minWrappingTtl() {
        return Optional.ofNullable(this.minWrappingTtl);
    }
    /**
     * @return A path in Vault that this rule applies to.
     * 
     */
    public String path() {
        return this.path;
    }
    /**
     * @return A list of parameters that must be specified.
     * 
     */
    public List<String> requiredParameters() {
        return this.requiredParameters == null ? List.of() : this.requiredParameters;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetPolicyDocumentRule defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable List<GetPolicyDocumentRuleAllowedParameter> allowedParameters;
        private List<String> capabilities;
        private @Nullable List<GetPolicyDocumentRuleDeniedParameter> deniedParameters;
        private @Nullable String description;
        private @Nullable String maxWrappingTtl;
        private @Nullable String minWrappingTtl;
        private String path;
        private @Nullable List<String> requiredParameters;

        public Builder() {
    	      // Empty
        }

        public Builder(GetPolicyDocumentRule defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.allowedParameters = defaults.allowedParameters;
    	      this.capabilities = defaults.capabilities;
    	      this.deniedParameters = defaults.deniedParameters;
    	      this.description = defaults.description;
    	      this.maxWrappingTtl = defaults.maxWrappingTtl;
    	      this.minWrappingTtl = defaults.minWrappingTtl;
    	      this.path = defaults.path;
    	      this.requiredParameters = defaults.requiredParameters;
        }

        public Builder allowedParameters(@Nullable List<GetPolicyDocumentRuleAllowedParameter> allowedParameters) {
            this.allowedParameters = allowedParameters;
            return this;
        }
        public Builder allowedParameters(GetPolicyDocumentRuleAllowedParameter... allowedParameters) {
            return allowedParameters(List.of(allowedParameters));
        }
        public Builder capabilities(List<String> capabilities) {
            this.capabilities = Objects.requireNonNull(capabilities);
            return this;
        }
        public Builder capabilities(String... capabilities) {
            return capabilities(List.of(capabilities));
        }
        public Builder deniedParameters(@Nullable List<GetPolicyDocumentRuleDeniedParameter> deniedParameters) {
            this.deniedParameters = deniedParameters;
            return this;
        }
        public Builder deniedParameters(GetPolicyDocumentRuleDeniedParameter... deniedParameters) {
            return deniedParameters(List.of(deniedParameters));
        }
        public Builder description(@Nullable String description) {
            this.description = description;
            return this;
        }
        public Builder maxWrappingTtl(@Nullable String maxWrappingTtl) {
            this.maxWrappingTtl = maxWrappingTtl;
            return this;
        }
        public Builder minWrappingTtl(@Nullable String minWrappingTtl) {
            this.minWrappingTtl = minWrappingTtl;
            return this;
        }
        public Builder path(String path) {
            this.path = Objects.requireNonNull(path);
            return this;
        }
        public Builder requiredParameters(@Nullable List<String> requiredParameters) {
            this.requiredParameters = requiredParameters;
            return this;
        }
        public Builder requiredParameters(String... requiredParameters) {
            return requiredParameters(List.of(requiredParameters));
        }        public GetPolicyDocumentRule build() {
            return new GetPolicyDocumentRule(allowedParameters, capabilities, deniedParameters, description, maxWrappingTtl, minWrappingTtl, path, requiredParameters);
        }
    }
}