// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.github;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.github.AuthBackendArgs;
import com.pulumi.vault.github.inputs.AuthBackendState;
import com.pulumi.vault.github.outputs.AuthBackendTune;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a GitHub Auth mount in a Vault server. See the [Vault
 * documentation](https://www.vaultproject.io/docs/auth/github/) for more
 * information.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.github.AuthBackend;
 * import com.pulumi.vault.github.AuthBackendArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var example = new AuthBackend(&#34;example&#34;, AuthBackendArgs.builder()        
 *             .organization(&#34;myorg&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * GitHub authentication mounts can be imported using the `path`, e.g.
 * 
 * ```sh
 *  $ pulumi import vault:github/authBackend:AuthBackend example github
 * ```
 * 
 */
@ResourceType(type="vault:github/authBackend:AuthBackend")
public class AuthBackend extends com.pulumi.resources.CustomResource {
    /**
     * The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
     * 
     */
    @Export(name="accessor", type=String.class, parameters={})
    private Output<String> accessor;

    /**
     * @return The mount accessor related to the auth mount. It is useful for integration with [Identity Secrets Engine](https://www.vaultproject.io/docs/secrets/identity/index.html).
     * 
     */
    public Output<String> accessor() {
        return this.accessor;
    }
    /**
     * The API endpoint to use. Useful if you
     * are running GitHub Enterprise or an API-compatible authentication server.
     * 
     */
    @Export(name="baseUrl", type=String.class, parameters={})
    private Output</* @Nullable */ String> baseUrl;

    /**
     * @return The API endpoint to use. Useful if you
     * are running GitHub Enterprise or an API-compatible authentication server.
     * 
     */
    public Output<Optional<String>> baseUrl() {
        return Codegen.optional(this.baseUrl);
    }
    /**
     * Specifies the description of the mount.
     * This overrides the current stored value, if any.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return Specifies the description of the mount.
     * This overrides the current stored value, if any.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The organization configured users must be part of.
     * 
     */
    @Export(name="organization", type=String.class, parameters={})
    private Output<String> organization;

    /**
     * @return The organization configured users must be part of.
     * 
     */
    public Output<String> organization() {
        return this.organization;
    }
    /**
     * The ID of the organization users must be part of.
     * Vault will attempt to fetch and set this value if it is not provided. (Vault 1.10+)
     * 
     */
    @Export(name="organizationId", type=Integer.class, parameters={})
    private Output<Integer> organizationId;

    /**
     * @return The ID of the organization users must be part of.
     * Vault will attempt to fetch and set this value if it is not provided. (Vault 1.10+)
     * 
     */
    public Output<Integer> organizationId() {
        return this.organizationId;
    }
    /**
     * Path where the auth backend is mounted. Defaults to `auth/github`
     * if not specified.
     * 
     */
    @Export(name="path", type=String.class, parameters={})
    private Output</* @Nullable */ String> path;

    /**
     * @return Path where the auth backend is mounted. Defaults to `auth/github`
     * if not specified.
     * 
     */
    public Output<Optional<String>> path() {
        return Codegen.optional(this.path);
    }
    /**
     * (Optional) List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     * 
     */
    @Export(name="tokenBoundCidrs", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> tokenBoundCidrs;

    /**
     * @return (Optional) List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     * 
     */
    public Output<Optional<List<String>>> tokenBoundCidrs() {
        return Codegen.optional(this.tokenBoundCidrs);
    }
    /**
     * (Optional) If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `token_ttl` and
     * `token_max_ttl` would otherwise allow a renewal.
     * 
     */
    @Export(name="tokenExplicitMaxTtl", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> tokenExplicitMaxTtl;

    /**
     * @return (Optional) If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `token_ttl` and
     * `token_max_ttl` would otherwise allow a renewal.
     * 
     */
    public Output<Optional<Integer>> tokenExplicitMaxTtl() {
        return Codegen.optional(this.tokenExplicitMaxTtl);
    }
    /**
     * (Optional) The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     * 
     */
    @Export(name="tokenMaxTtl", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> tokenMaxTtl;

    /**
     * @return (Optional) The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     * 
     */
    public Output<Optional<Integer>> tokenMaxTtl() {
        return Codegen.optional(this.tokenMaxTtl);
    }
    /**
     * (Optional) If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     * 
     */
    @Export(name="tokenNoDefaultPolicy", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> tokenNoDefaultPolicy;

    /**
     * @return (Optional) If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     * 
     */
    public Output<Optional<Boolean>> tokenNoDefaultPolicy() {
        return Codegen.optional(this.tokenNoDefaultPolicy);
    }
    /**
     * (Optional) The [maximum number](https://www.vaultproject.io/api-docs/github#token_num_uses)
     * of times a generated token may be used (within its lifetime); 0 means unlimited.
     * 
     */
    @Export(name="tokenNumUses", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> tokenNumUses;

    /**
     * @return (Optional) The [maximum number](https://www.vaultproject.io/api-docs/github#token_num_uses)
     * of times a generated token may be used (within its lifetime); 0 means unlimited.
     * 
     */
    public Output<Optional<Integer>> tokenNumUses() {
        return Codegen.optional(this.tokenNumUses);
    }
    /**
     * (Optional) If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token&#39;s TTL will be set to the
     * value of this field. Specified in seconds.
     * 
     */
    @Export(name="tokenPeriod", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> tokenPeriod;

    /**
     * @return (Optional) If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token&#39;s TTL will be set to the
     * value of this field. Specified in seconds.
     * 
     */
    public Output<Optional<Integer>> tokenPeriod() {
        return Codegen.optional(this.tokenPeriod);
    }
    /**
     * (Optional) List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     * 
     */
    @Export(name="tokenPolicies", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> tokenPolicies;

    /**
     * @return (Optional) List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     * 
     */
    public Output<Optional<List<String>>> tokenPolicies() {
        return Codegen.optional(this.tokenPolicies);
    }
    /**
     * (Optional) The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     * 
     */
    @Export(name="tokenTtl", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> tokenTtl;

    /**
     * @return (Optional) The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     * 
     */
    public Output<Optional<Integer>> tokenTtl() {
        return Codegen.optional(this.tokenTtl);
    }
    /**
     * Specifies the type of tokens that should be returned by
     * the mount. Valid values are &#34;default-service&#34;, &#34;default-batch&#34;, &#34;service&#34;, &#34;batch&#34;.
     * 
     */
    @Export(name="tokenType", type=String.class, parameters={})
    private Output</* @Nullable */ String> tokenType;

    /**
     * @return Specifies the type of tokens that should be returned by
     * the mount. Valid values are &#34;default-service&#34;, &#34;default-batch&#34;, &#34;service&#34;, &#34;batch&#34;.
     * 
     */
    public Output<Optional<String>> tokenType() {
        return Codegen.optional(this.tokenType);
    }
    /**
     * Extra configuration block. Structure is documented below.
     * 
     */
    @Export(name="tune", type=AuthBackendTune.class, parameters={})
    private Output<AuthBackendTune> tune;

    /**
     * @return Extra configuration block. Structure is documented below.
     * 
     */
    public Output<AuthBackendTune> tune() {
        return this.tune;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuthBackend(String name) {
        this(name, AuthBackendArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuthBackend(String name, AuthBackendArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuthBackend(String name, AuthBackendArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:github/authBackend:AuthBackend", name, args == null ? AuthBackendArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AuthBackend(String name, Output<String> id, @Nullable AuthBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:github/authBackend:AuthBackend", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static AuthBackend get(String name, Output<String> id, @Nullable AuthBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuthBackend(name, id, state, options);
    }
}
