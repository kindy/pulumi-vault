// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.okta;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.okta.AuthBackendArgs;
import com.pulumi.vault.okta.inputs.AuthBackendState;
import com.pulumi.vault.okta.outputs.AuthBackendGroup;
import com.pulumi.vault.okta.outputs.AuthBackendUser;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource for managing an
 * [Okta auth backend within Vault](https://www.vaultproject.io/docs/auth/okta.html).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.okta.AuthBackend;
 * import com.pulumi.vault.okta.AuthBackendArgs;
 * import com.pulumi.vault.okta.inputs.AuthBackendGroupArgs;
 * import com.pulumi.vault.okta.inputs.AuthBackendUserArgs;
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
 *             .description(&#34;Demonstration of the Terraform Okta auth backend&#34;)
 *             .groups(AuthBackendGroupArgs.builder()
 *                 .groupName(&#34;foo&#34;)
 *                 .policies(                
 *                     &#34;one&#34;,
 *                     &#34;two&#34;)
 *                 .build())
 *             .organization(&#34;example&#34;)
 *             .token(&#34;something that should be kept secret&#34;)
 *             .users(AuthBackendUserArgs.builder()
 *                 .groups(&#34;foo&#34;)
 *                 .username(&#34;bar&#34;)
 *                 .build())
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Okta authentication backends can be imported using its `path`, e.g.
 * 
 * ```sh
 *  $ pulumi import vault:okta/authBackend:AuthBackend example okta
 * ```
 * 
 */
@ResourceType(type="vault:okta/authBackend:AuthBackend")
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
     * The Okta url. Examples: oktapreview.com, okta.com
     * 
     */
    @Export(name="baseUrl", type=String.class, parameters={})
    private Output</* @Nullable */ String> baseUrl;

    /**
     * @return The Okta url. Examples: oktapreview.com, okta.com
     * 
     */
    public Output<Optional<String>> baseUrl() {
        return Codegen.optional(this.baseUrl);
    }
    /**
     * When true, requests by Okta for a MFA check will be bypassed. This also disallows certain status checks on the account, such as whether the password is expired.
     * 
     */
    @Export(name="bypassOktaMfa", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> bypassOktaMfa;

    /**
     * @return When true, requests by Okta for a MFA check will be bypassed. This also disallows certain status checks on the account, such as whether the password is expired.
     * 
     */
    public Output<Optional<Boolean>> bypassOktaMfa() {
        return Codegen.optional(this.bypassOktaMfa);
    }
    /**
     * The description of the auth backend
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return The description of the auth backend
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     * 
     */
    @Export(name="disableRemount", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> disableRemount;

    /**
     * @return If set, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     * 
     */
    public Output<Optional<Boolean>> disableRemount() {
        return Codegen.optional(this.disableRemount);
    }
    /**
     * Associate Okta groups with policies within Vault.
     * See below for more details.
     * 
     */
    @Export(name="groups", type=List.class, parameters={AuthBackendGroup.class})
    private Output<List<AuthBackendGroup>> groups;

    /**
     * @return Associate Okta groups with policies within Vault.
     * See below for more details.
     * 
     */
    public Output<List<AuthBackendGroup>> groups() {
        return this.groups;
    }
    /**
     * Maximum duration after which authentication will be expired
     * [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
     * 
     */
    @Export(name="maxTtl", type=String.class, parameters={})
    private Output</* @Nullable */ String> maxTtl;

    /**
     * @return Maximum duration after which authentication will be expired
     * [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
     * 
     */
    public Output<Optional<String>> maxTtl() {
        return Codegen.optional(this.maxTtl);
    }
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Export(name="namespace", type=String.class, parameters={})
    private Output</* @Nullable */ String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Output<Optional<String>> namespace() {
        return Codegen.optional(this.namespace);
    }
    /**
     * The Okta organization. This will be the first part of the url `https://XXX.okta.com`
     * 
     */
    @Export(name="organization", type=String.class, parameters={})
    private Output<String> organization;

    /**
     * @return The Okta organization. This will be the first part of the url `https://XXX.okta.com`
     * 
     */
    public Output<String> organization() {
        return this.organization;
    }
    /**
     * Path to mount the Okta auth backend. Default to path `okta`.
     * 
     */
    @Export(name="path", type=String.class, parameters={})
    private Output</* @Nullable */ String> path;

    /**
     * @return Path to mount the Okta auth backend. Default to path `okta`.
     * 
     */
    public Output<Optional<String>> path() {
        return Codegen.optional(this.path);
    }
    /**
     * The Okta API token. This is required to query Okta for user group membership.
     * If this is not supplied only locally configured groups will be enabled.
     * 
     */
    @Export(name="token", type=String.class, parameters={})
    private Output</* @Nullable */ String> token;

    /**
     * @return The Okta API token. This is required to query Okta for user group membership.
     * If this is not supplied only locally configured groups will be enabled.
     * 
     */
    public Output<Optional<String>> token() {
        return Codegen.optional(this.token);
    }
    /**
     * Duration after which authentication will be expired.
     * [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
     * 
     */
    @Export(name="ttl", type=String.class, parameters={})
    private Output</* @Nullable */ String> ttl;

    /**
     * @return Duration after which authentication will be expired.
     * [See the documentation for info on valid duration formats](https://golang.org/pkg/time/#ParseDuration).
     * 
     */
    public Output<Optional<String>> ttl() {
        return Codegen.optional(this.ttl);
    }
    /**
     * Associate Okta users with groups or policies within Vault.
     * See below for more details.
     * 
     */
    @Export(name="users", type=List.class, parameters={AuthBackendUser.class})
    private Output<List<AuthBackendUser>> users;

    /**
     * @return Associate Okta users with groups or policies within Vault.
     * See below for more details.
     * 
     */
    public Output<List<AuthBackendUser>> users() {
        return this.users;
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
        super("vault:okta/authBackend:AuthBackend", name, args == null ? AuthBackendArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AuthBackend(String name, Output<String> id, @Nullable AuthBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:okta/authBackend:AuthBackend", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "token"
            ))
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
