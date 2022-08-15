// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.tokenauth;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.tokenauth.AuthBackendRoleArgs;
import com.pulumi.vault.tokenauth.inputs.AuthBackendRoleState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages Token auth backend role in a Vault server. See the [Vault
 * documentation](https://www.vaultproject.io/docs/auth/token.html) for more
 * information.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.tokenauth.AuthBackendRole;
 * import com.pulumi.vault.tokenauth.AuthBackendRoleArgs;
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
 *         var example = new AuthBackendRole(&#34;example&#34;, AuthBackendRoleArgs.builder()        
 *             .allowedEntityAliases(&#34;test_entity&#34;)
 *             .allowedPolicies(            
 *                 &#34;dev&#34;,
 *                 &#34;test&#34;)
 *             .disallowedPolicies(&#34;default&#34;)
 *             .orphan(true)
 *             .pathSuffix(&#34;path-suffix&#34;)
 *             .renewable(true)
 *             .roleName(&#34;my-role&#34;)
 *             .tokenExplicitMaxTtl(&#34;115200&#34;)
 *             .tokenPeriod(&#34;86400&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Token auth backend roles can be imported with `auth/token/roles/` followed by the `role_name`, e.g.
 * 
 * ```sh
 *  $ pulumi import vault:tokenauth/authBackendRole:AuthBackendRole example auth/token/roles/my-role
 * ```
 * 
 */
@ResourceType(type="vault:tokenauth/authBackendRole:AuthBackendRole")
public class AuthBackendRole extends com.pulumi.resources.CustomResource {
    /**
     * List of allowed entity aliases.
     * 
     */
    @Export(name="allowedEntityAliases", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> allowedEntityAliases;

    /**
     * @return List of allowed entity aliases.
     * 
     */
    public Output<Optional<List<String>>> allowedEntityAliases() {
        return Codegen.optional(this.allowedEntityAliases);
    }
    /**
     * List of allowed policies for given role.
     * 
     */
    @Export(name="allowedPolicies", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> allowedPolicies;

    /**
     * @return List of allowed policies for given role.
     * 
     */
    public Output<Optional<List<String>>> allowedPolicies() {
        return Codegen.optional(this.allowedPolicies);
    }
    /**
     * Set of allowed policies with glob match for given role.
     * 
     */
    @Export(name="allowedPoliciesGlobs", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> allowedPoliciesGlobs;

    /**
     * @return Set of allowed policies with glob match for given role.
     * 
     */
    public Output<Optional<List<String>>> allowedPoliciesGlobs() {
        return Codegen.optional(this.allowedPoliciesGlobs);
    }
    /**
     * List of disallowed policies for given role.
     * 
     */
    @Export(name="disallowedPolicies", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> disallowedPolicies;

    /**
     * @return List of disallowed policies for given role.
     * 
     */
    public Output<Optional<List<String>>> disallowedPolicies() {
        return Codegen.optional(this.disallowedPolicies);
    }
    /**
     * Set of disallowed policies with glob match for given role.
     * 
     */
    @Export(name="disallowedPoliciesGlobs", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> disallowedPoliciesGlobs;

    /**
     * @return Set of disallowed policies with glob match for given role.
     * 
     */
    public Output<Optional<List<String>>> disallowedPoliciesGlobs() {
        return Codegen.optional(this.disallowedPoliciesGlobs);
    }
    /**
     * If true, tokens created against this policy will be orphan tokens.
     * 
     */
    @Export(name="orphan", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> orphan;

    /**
     * @return If true, tokens created against this policy will be orphan tokens.
     * 
     */
    public Output<Optional<Boolean>> orphan() {
        return Codegen.optional(this.orphan);
    }
    /**
     * Tokens created against this role will have the given suffix as part of their path in addition to the role name.
     * 
     */
    @Export(name="pathSuffix", type=String.class, parameters={})
    private Output</* @Nullable */ String> pathSuffix;

    /**
     * @return Tokens created against this role will have the given suffix as part of their path in addition to the role name.
     * 
     */
    public Output<Optional<String>> pathSuffix() {
        return Codegen.optional(this.pathSuffix);
    }
    /**
     * Whether to disable the ability of the token to be renewed past its initial TTL.
     * 
     */
    @Export(name="renewable", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> renewable;

    /**
     * @return Whether to disable the ability of the token to be renewed past its initial TTL.
     * 
     */
    public Output<Optional<Boolean>> renewable() {
        return Codegen.optional(this.renewable);
    }
    /**
     * The name of the role.
     * 
     */
    @Export(name="roleName", type=String.class, parameters={})
    private Output<String> roleName;

    /**
     * @return The name of the role.
     * 
     */
    public Output<String> roleName() {
        return this.roleName;
    }
    /**
     * List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     * 
     */
    @Export(name="tokenBoundCidrs", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> tokenBoundCidrs;

    /**
     * @return List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     * 
     */
    public Output<Optional<List<String>>> tokenBoundCidrs() {
        return Codegen.optional(this.tokenBoundCidrs);
    }
    /**
     * If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `token_ttl` and
     * `token_max_ttl` would otherwise allow a renewal.
     * 
     */
    @Export(name="tokenExplicitMaxTtl", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> tokenExplicitMaxTtl;

    /**
     * @return If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `token_ttl` and
     * `token_max_ttl` would otherwise allow a renewal.
     * 
     */
    public Output<Optional<Integer>> tokenExplicitMaxTtl() {
        return Codegen.optional(this.tokenExplicitMaxTtl);
    }
    /**
     * The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     * 
     */
    @Export(name="tokenMaxTtl", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> tokenMaxTtl;

    /**
     * @return The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     * 
     */
    public Output<Optional<Integer>> tokenMaxTtl() {
        return Codegen.optional(this.tokenMaxTtl);
    }
    /**
     * If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     * 
     */
    @Export(name="tokenNoDefaultPolicy", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> tokenNoDefaultPolicy;

    /**
     * @return If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     * 
     */
    public Output<Optional<Boolean>> tokenNoDefaultPolicy() {
        return Codegen.optional(this.tokenNoDefaultPolicy);
    }
    /**
     * The [maximum number](https://www.vaultproject.io/api-docs/token#token_num_uses)
     * of times a generated token may be used (within its lifetime); 0 means unlimited.
     * 
     */
    @Export(name="tokenNumUses", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> tokenNumUses;

    /**
     * @return The [maximum number](https://www.vaultproject.io/api-docs/token#token_num_uses)
     * of times a generated token may be used (within its lifetime); 0 means unlimited.
     * 
     */
    public Output<Optional<Integer>> tokenNumUses() {
        return Codegen.optional(this.tokenNumUses);
    }
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token&#39;s TTL will be set to the
     * value of this field. Specified in seconds.
     * 
     */
    @Export(name="tokenPeriod", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> tokenPeriod;

    /**
     * @return If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token&#39;s TTL will be set to the
     * value of this field. Specified in seconds.
     * 
     */
    public Output<Optional<Integer>> tokenPeriod() {
        return Codegen.optional(this.tokenPeriod);
    }
    /**
     * Generated Token&#39;s Policies
     * 
     */
    @Export(name="tokenPolicies", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> tokenPolicies;

    /**
     * @return Generated Token&#39;s Policies
     * 
     */
    public Output<Optional<List<String>>> tokenPolicies() {
        return Codegen.optional(this.tokenPolicies);
    }
    /**
     * The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     * 
     */
    @Export(name="tokenTtl", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> tokenTtl;

    /**
     * @return The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     * 
     */
    public Output<Optional<Integer>> tokenTtl() {
        return Codegen.optional(this.tokenTtl);
    }
    /**
     * The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount&#39;s tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     * 
     */
    @Export(name="tokenType", type=String.class, parameters={})
    private Output</* @Nullable */ String> tokenType;

    /**
     * @return The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount&#39;s tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     * 
     */
    public Output<Optional<String>> tokenType() {
        return Codegen.optional(this.tokenType);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuthBackendRole(String name) {
        this(name, AuthBackendRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuthBackendRole(String name, AuthBackendRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuthBackendRole(String name, AuthBackendRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:tokenauth/authBackendRole:AuthBackendRole", name, args == null ? AuthBackendRoleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AuthBackendRole(String name, Output<String> id, @Nullable AuthBackendRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:tokenauth/authBackendRole:AuthBackendRole", name, state, makeResourceOptions(options, id));
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
    public static AuthBackendRole get(String name, Output<String> id, @Nullable AuthBackendRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuthBackendRole(name, id, state, options);
    }
}
