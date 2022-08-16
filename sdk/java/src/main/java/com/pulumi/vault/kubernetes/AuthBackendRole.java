// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.kubernetes;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.kubernetes.AuthBackendRoleArgs;
import com.pulumi.vault.kubernetes.inputs.AuthBackendRoleState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an Kubernetes auth backend role in a Vault server. See the [Vault
 * documentation](https://www.vaultproject.io/docs/auth/kubernetes.html) for more
 * information.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.AuthBackend;
 * import com.pulumi.vault.AuthBackendArgs;
 * import com.pulumi.vault.kubernetes.AuthBackendRole;
 * import com.pulumi.vault.kubernetes.AuthBackendRoleArgs;
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
 *         var kubernetes = new AuthBackend(&#34;kubernetes&#34;, AuthBackendArgs.builder()        
 *             .type(&#34;kubernetes&#34;)
 *             .build());
 * 
 *         var example = new AuthBackendRole(&#34;example&#34;, AuthBackendRoleArgs.builder()        
 *             .backend(kubernetes.path())
 *             .roleName(&#34;example-role&#34;)
 *             .boundServiceAccountNames(&#34;example&#34;)
 *             .boundServiceAccountNamespaces(&#34;example&#34;)
 *             .tokenTtl(3600)
 *             .tokenPolicies(            
 *                 &#34;default&#34;,
 *                 &#34;dev&#34;,
 *                 &#34;prod&#34;)
 *             .audience(&#34;vault&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Kubernetes auth backend role can be imported using the `path`, e.g.
 * 
 * ```sh
 *  $ pulumi import vault:kubernetes/authBackendRole:AuthBackendRole foo auth/kubernetes/role/foo
 * ```
 * 
 */
@ResourceType(type="vault:kubernetes/authBackendRole:AuthBackendRole")
public class AuthBackendRole extends com.pulumi.resources.CustomResource {
    /**
     * Configures how identity aliases are generated.
     * Valid choices are: `serviceaccount_uid`, `serviceaccount_name`. (vault-1.9+)
     * 
     */
    @Export(name="aliasNameSource", type=String.class, parameters={})
    private Output<String> aliasNameSource;

    /**
     * @return Configures how identity aliases are generated.
     * Valid choices are: `serviceaccount_uid`, `serviceaccount_name`. (vault-1.9+)
     * 
     */
    public Output<String> aliasNameSource() {
        return this.aliasNameSource;
    }
    /**
     * Audience claim to verify in the JWT.
     * 
     */
    @Export(name="audience", type=String.class, parameters={})
    private Output</* @Nullable */ String> audience;

    /**
     * @return Audience claim to verify in the JWT.
     * 
     */
    public Output<Optional<String>> audience() {
        return Codegen.optional(this.audience);
    }
    /**
     * Unique name of the kubernetes backend to configure.
     * 
     */
    @Export(name="backend", type=String.class, parameters={})
    private Output</* @Nullable */ String> backend;

    /**
     * @return Unique name of the kubernetes backend to configure.
     * 
     */
    public Output<Optional<String>> backend() {
        return Codegen.optional(this.backend);
    }
    /**
     * List of service account names able to access this role. If set to `[&#34;*&#34;]` all names are allowed, both this and bound_service_account_namespaces can not be &#34;*&#34;.
     * 
     */
    @Export(name="boundServiceAccountNames", type=List.class, parameters={String.class})
    private Output<List<String>> boundServiceAccountNames;

    /**
     * @return List of service account names able to access this role. If set to `[&#34;*&#34;]` all names are allowed, both this and bound_service_account_namespaces can not be &#34;*&#34;.
     * 
     */
    public Output<List<String>> boundServiceAccountNames() {
        return this.boundServiceAccountNames;
    }
    /**
     * List of namespaces allowed to access this role. If set to `[&#34;*&#34;]` all namespaces are allowed, both this and bound_service_account_names can not be set to &#34;*&#34;.
     * 
     */
    @Export(name="boundServiceAccountNamespaces", type=List.class, parameters={String.class})
    private Output<List<String>> boundServiceAccountNamespaces;

    /**
     * @return List of namespaces allowed to access this role. If set to `[&#34;*&#34;]` all namespaces are allowed, both this and bound_service_account_names can not be set to &#34;*&#34;.
     * 
     */
    public Output<List<String>> boundServiceAccountNamespaces() {
        return this.boundServiceAccountNamespaces;
    }
    /**
     * Name of the role.
     * 
     */
    @Export(name="roleName", type=String.class, parameters={})
    private Output<String> roleName;

    /**
     * @return Name of the role.
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
     * The [maximum number](https://www.vaultproject.io/api-docs/kubernetes#token_num_uses)
     * of times a generated token may be used (within its lifetime); 0 means unlimited.
     * 
     */
    @Export(name="tokenNumUses", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> tokenNumUses;

    /**
     * @return The [maximum number](https://www.vaultproject.io/api-docs/kubernetes#token_num_uses)
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
     * List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     * 
     */
    @Export(name="tokenPolicies", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> tokenPolicies;

    /**
     * @return List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     * 
     */
    public Output<Optional<List<String>>> tokenPolicies() {
        return Codegen.optional(this.tokenPolicies);
    }
    /**
     * The initial ttl of the token to generate in seconds
     * 
     */
    @Export(name="tokenTtl", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> tokenTtl;

    /**
     * @return The initial ttl of the token to generate in seconds
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
        super("vault:kubernetes/authBackendRole:AuthBackendRole", name, args == null ? AuthBackendRoleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AuthBackendRole(String name, Output<String> id, @Nullable AuthBackendRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:kubernetes/authBackendRole:AuthBackendRole", name, state, makeResourceOptions(options, id));
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