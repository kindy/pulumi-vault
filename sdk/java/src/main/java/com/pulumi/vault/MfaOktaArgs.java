// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MfaOktaArgs extends com.pulumi.resources.ResourceArgs {

    public static final MfaOktaArgs Empty = new MfaOktaArgs();

    /**
     * `(string: &lt;required&gt;)` - Okta API key.
     * 
     */
    @Import(name="apiToken", required=true)
    private Output<String> apiToken;

    /**
     * @return `(string: &lt;required&gt;)` - Okta API key.
     * 
     */
    public Output<String> apiToken() {
        return this.apiToken;
    }

    /**
     * `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`,
     * `oktapreview.com`, and `okta-emea.com`.
     * 
     */
    @Import(name="baseUrl")
    private @Nullable Output<String> baseUrl;

    /**
     * @return `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`,
     * `oktapreview.com`, and `okta-emea.com`.
     * 
     */
    public Optional<Output<String>> baseUrl() {
        return Optional.ofNullable(this.baseUrl);
    }

    /**
     * `(string: &lt;required&gt;)` - The mount to tie this method to for use in automatic mappings.
     * The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
     * 
     */
    @Import(name="mountAccessor", required=true)
    private Output<String> mountAccessor;

    /**
     * @return `(string: &lt;required&gt;)` - The mount to tie this method to for use in automatic mappings.
     * The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
     * 
     */
    public Output<String> mountAccessor() {
        return this.mountAccessor;
    }

    /**
     * `(string: &lt;required&gt;)` – Name of the MFA method.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return `(string: &lt;required&gt;)` – Name of the MFA method.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * `(string: &lt;required&gt;)` - Name of the organization to be used in the Okta API.
     * 
     */
    @Import(name="orgName", required=true)
    private Output<String> orgName;

    /**
     * @return `(string: &lt;required&gt;)` - Name of the organization to be used in the Okta API.
     * 
     */
    public Output<String> orgName() {
        return this.orgName;
    }

    /**
     * `(string: &lt;required&gt;)` - If set to true, the username will only match the
     * primary email for the account.
     * 
     */
    @Import(name="primaryEmail")
    private @Nullable Output<Boolean> primaryEmail;

    /**
     * @return `(string: &lt;required&gt;)` - If set to true, the username will only match the
     * primary email for the account.
     * 
     */
    public Optional<Output<Boolean>> primaryEmail() {
        return Optional.ofNullable(this.primaryEmail);
    }

    /**
     * `(string)` - A format string for mapping Identity names to MFA method names.
     * Values to substitute should be placed in `{{}}`. For example, `&#34;{{alias.name}}@example.com&#34;`.
     * If blank, the Alias&#39;s Name field will be used as-is. Currently-supported mappings:
     * - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
     * - entity.name: The name configured for the Entity
     * - alias.metadata.`&lt;key&gt;`: The value of the Alias&#39;s metadata parameter
     * - entity.metadata.`&lt;key&gt;`: The value of the Entity&#39;s metadata parameter
     * 
     */
    @Import(name="usernameFormat")
    private @Nullable Output<String> usernameFormat;

    /**
     * @return `(string)` - A format string for mapping Identity names to MFA method names.
     * Values to substitute should be placed in `{{}}`. For example, `&#34;{{alias.name}}@example.com&#34;`.
     * If blank, the Alias&#39;s Name field will be used as-is. Currently-supported mappings:
     * - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
     * - entity.name: The name configured for the Entity
     * - alias.metadata.`&lt;key&gt;`: The value of the Alias&#39;s metadata parameter
     * - entity.metadata.`&lt;key&gt;`: The value of the Entity&#39;s metadata parameter
     * 
     */
    public Optional<Output<String>> usernameFormat() {
        return Optional.ofNullable(this.usernameFormat);
    }

    private MfaOktaArgs() {}

    private MfaOktaArgs(MfaOktaArgs $) {
        this.apiToken = $.apiToken;
        this.baseUrl = $.baseUrl;
        this.mountAccessor = $.mountAccessor;
        this.name = $.name;
        this.orgName = $.orgName;
        this.primaryEmail = $.primaryEmail;
        this.usernameFormat = $.usernameFormat;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MfaOktaArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MfaOktaArgs $;

        public Builder() {
            $ = new MfaOktaArgs();
        }

        public Builder(MfaOktaArgs defaults) {
            $ = new MfaOktaArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param apiToken `(string: &lt;required&gt;)` - Okta API key.
         * 
         * @return builder
         * 
         */
        public Builder apiToken(Output<String> apiToken) {
            $.apiToken = apiToken;
            return this;
        }

        /**
         * @param apiToken `(string: &lt;required&gt;)` - Okta API key.
         * 
         * @return builder
         * 
         */
        public Builder apiToken(String apiToken) {
            return apiToken(Output.of(apiToken));
        }

        /**
         * @param baseUrl `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`,
         * `oktapreview.com`, and `okta-emea.com`.
         * 
         * @return builder
         * 
         */
        public Builder baseUrl(@Nullable Output<String> baseUrl) {
            $.baseUrl = baseUrl;
            return this;
        }

        /**
         * @param baseUrl `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`,
         * `oktapreview.com`, and `okta-emea.com`.
         * 
         * @return builder
         * 
         */
        public Builder baseUrl(String baseUrl) {
            return baseUrl(Output.of(baseUrl));
        }

        /**
         * @param mountAccessor `(string: &lt;required&gt;)` - The mount to tie this method to for use in automatic mappings.
         * The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
         * 
         * @return builder
         * 
         */
        public Builder mountAccessor(Output<String> mountAccessor) {
            $.mountAccessor = mountAccessor;
            return this;
        }

        /**
         * @param mountAccessor `(string: &lt;required&gt;)` - The mount to tie this method to for use in automatic mappings.
         * The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
         * 
         * @return builder
         * 
         */
        public Builder mountAccessor(String mountAccessor) {
            return mountAccessor(Output.of(mountAccessor));
        }

        /**
         * @param name `(string: &lt;required&gt;)` – Name of the MFA method.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name `(string: &lt;required&gt;)` – Name of the MFA method.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param orgName `(string: &lt;required&gt;)` - Name of the organization to be used in the Okta API.
         * 
         * @return builder
         * 
         */
        public Builder orgName(Output<String> orgName) {
            $.orgName = orgName;
            return this;
        }

        /**
         * @param orgName `(string: &lt;required&gt;)` - Name of the organization to be used in the Okta API.
         * 
         * @return builder
         * 
         */
        public Builder orgName(String orgName) {
            return orgName(Output.of(orgName));
        }

        /**
         * @param primaryEmail `(string: &lt;required&gt;)` - If set to true, the username will only match the
         * primary email for the account.
         * 
         * @return builder
         * 
         */
        public Builder primaryEmail(@Nullable Output<Boolean> primaryEmail) {
            $.primaryEmail = primaryEmail;
            return this;
        }

        /**
         * @param primaryEmail `(string: &lt;required&gt;)` - If set to true, the username will only match the
         * primary email for the account.
         * 
         * @return builder
         * 
         */
        public Builder primaryEmail(Boolean primaryEmail) {
            return primaryEmail(Output.of(primaryEmail));
        }

        /**
         * @param usernameFormat `(string)` - A format string for mapping Identity names to MFA method names.
         * Values to substitute should be placed in `{{}}`. For example, `&#34;{{alias.name}}@example.com&#34;`.
         * If blank, the Alias&#39;s Name field will be used as-is. Currently-supported mappings:
         * - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
         * - entity.name: The name configured for the Entity
         * - alias.metadata.`&lt;key&gt;`: The value of the Alias&#39;s metadata parameter
         * - entity.metadata.`&lt;key&gt;`: The value of the Entity&#39;s metadata parameter
         * 
         * @return builder
         * 
         */
        public Builder usernameFormat(@Nullable Output<String> usernameFormat) {
            $.usernameFormat = usernameFormat;
            return this;
        }

        /**
         * @param usernameFormat `(string)` - A format string for mapping Identity names to MFA method names.
         * Values to substitute should be placed in `{{}}`. For example, `&#34;{{alias.name}}@example.com&#34;`.
         * If blank, the Alias&#39;s Name field will be used as-is. Currently-supported mappings:
         * - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
         * - entity.name: The name configured for the Entity
         * - alias.metadata.`&lt;key&gt;`: The value of the Alias&#39;s metadata parameter
         * - entity.metadata.`&lt;key&gt;`: The value of the Entity&#39;s metadata parameter
         * 
         * @return builder
         * 
         */
        public Builder usernameFormat(String usernameFormat) {
            return usernameFormat(Output.of(usernameFormat));
        }

        public MfaOktaArgs build() {
            $.apiToken = Objects.requireNonNull($.apiToken, "expected parameter 'apiToken' to be non-null");
            $.mountAccessor = Objects.requireNonNull($.mountAccessor, "expected parameter 'mountAccessor' to be non-null");
            $.orgName = Objects.requireNonNull($.orgName, "expected parameter 'orgName' to be non-null");
            return $;
        }
    }

}