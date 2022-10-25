// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.appRole;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.appRole.AuthBackendRoleSecretIdArgs;
import com.pulumi.vault.appRole.inputs.AuthBackendRoleSecretIdState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an AppRole auth backend SecretID in a Vault server. See the [Vault
 * documentation](https://www.vaultproject.io/docs/auth/approle) for more
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
 * import com.pulumi.vault.appRole.AuthBackendRole;
 * import com.pulumi.vault.appRole.AuthBackendRoleArgs;
 * import com.pulumi.vault.appRole.AuthBackendRoleSecretId;
 * import com.pulumi.vault.appRole.AuthBackendRoleSecretIdArgs;
 * import static com.pulumi.codegen.internal.Serialization.*;
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
 *         var approle = new AuthBackend(&#34;approle&#34;, AuthBackendArgs.builder()        
 *             .type(&#34;approle&#34;)
 *             .build());
 * 
 *         var example = new AuthBackendRole(&#34;example&#34;, AuthBackendRoleArgs.builder()        
 *             .backend(approle.path())
 *             .roleName(&#34;test-role&#34;)
 *             .tokenPolicies(            
 *                 &#34;default&#34;,
 *                 &#34;dev&#34;,
 *                 &#34;prod&#34;)
 *             .build());
 * 
 *         var id = new AuthBackendRoleSecretId(&#34;id&#34;, AuthBackendRoleSecretIdArgs.builder()        
 *             .backend(approle.path())
 *             .roleName(example.roleName())
 *             .metadata(serializeJson(
 *                 jsonObject(
 *                     jsonProperty(&#34;hello&#34;, &#34;world&#34;)
 *                 )))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="vault:appRole/authBackendRoleSecretId:AuthBackendRoleSecretId")
public class AuthBackendRoleSecretId extends com.pulumi.resources.CustomResource {
    /**
     * The unique ID for this SecretID that can be safely logged.
     * 
     */
    @Export(name="accessor", type=String.class, parameters={})
    private Output<String> accessor;

    /**
     * @return The unique ID for this SecretID that can be safely logged.
     * 
     */
    public Output<String> accessor() {
        return this.accessor;
    }
    /**
     * Unique name of the auth backend to configure.
     * 
     */
    @Export(name="backend", type=String.class, parameters={})
    private Output</* @Nullable */ String> backend;

    /**
     * @return Unique name of the auth backend to configure.
     * 
     */
    public Output<Optional<String>> backend() {
        return Codegen.optional(this.backend);
    }
    /**
     * If set, specifies blocks of IP addresses which can
     * perform the login operation using this SecretID.
     * 
     */
    @Export(name="cidrLists", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> cidrLists;

    /**
     * @return If set, specifies blocks of IP addresses which can
     * perform the login operation using this SecretID.
     * 
     */
    public Output<Optional<List<String>>> cidrLists() {
        return Codegen.optional(this.cidrLists);
    }
    /**
     * A JSON-encoded string containing metadata in
     * key-value pairs to be set on tokens issued with this SecretID.
     * 
     */
    @Export(name="metadata", type=String.class, parameters={})
    private Output</* @Nullable */ String> metadata;

    /**
     * @return A JSON-encoded string containing metadata in
     * key-value pairs to be set on tokens issued with this SecretID.
     * 
     */
    public Output<Optional<String>> metadata() {
        return Codegen.optional(this.metadata);
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
     * The name of the role to create the SecretID for.
     * 
     */
    @Export(name="roleName", type=String.class, parameters={})
    private Output<String> roleName;

    /**
     * @return The name of the role to create the SecretID for.
     * 
     */
    public Output<String> roleName() {
        return this.roleName;
    }
    /**
     * The SecretID to be created. If set, uses &#34;Push&#34;
     * mode.  Defaults to Vault auto-generating SecretIDs.
     * 
     */
    @Export(name="secretId", type=String.class, parameters={})
    private Output<String> secretId;

    /**
     * @return The SecretID to be created. If set, uses &#34;Push&#34;
     * mode.  Defaults to Vault auto-generating SecretIDs.
     * 
     */
    public Output<String> secretId() {
        return this.secretId;
    }
    /**
     * Set to `true` to use the wrapped secret-id accessor as the resource ID.
     * If `false` (default value), a fresh secret ID will be regenerated whenever the wrapping token is expired or
     * invalidated through unwrapping.
     * 
     */
    @Export(name="withWrappedAccessor", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> withWrappedAccessor;

    /**
     * @return Set to `true` to use the wrapped secret-id accessor as the resource ID.
     * If `false` (default value), a fresh secret ID will be regenerated whenever the wrapping token is expired or
     * invalidated through unwrapping.
     * 
     */
    public Output<Optional<Boolean>> withWrappedAccessor() {
        return Codegen.optional(this.withWrappedAccessor);
    }
    /**
     * The unique ID for the response-wrapped SecretID that can
     * be safely logged.
     * 
     */
    @Export(name="wrappingAccessor", type=String.class, parameters={})
    private Output<String> wrappingAccessor;

    /**
     * @return The unique ID for the response-wrapped SecretID that can
     * be safely logged.
     * 
     */
    public Output<String> wrappingAccessor() {
        return this.wrappingAccessor;
    }
    /**
     * The token used to retrieve a response-wrapped SecretID.
     * 
     */
    @Export(name="wrappingToken", type=String.class, parameters={})
    private Output<String> wrappingToken;

    /**
     * @return The token used to retrieve a response-wrapped SecretID.
     * 
     */
    public Output<String> wrappingToken() {
        return this.wrappingToken;
    }
    /**
     * If set, the SecretID response will be
     * [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
     * and available for the duration specified. Only a single unwrapping of the
     * token is allowed.
     * 
     */
    @Export(name="wrappingTtl", type=String.class, parameters={})
    private Output</* @Nullable */ String> wrappingTtl;

    /**
     * @return If set, the SecretID response will be
     * [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
     * and available for the duration specified. Only a single unwrapping of the
     * token is allowed.
     * 
     */
    public Output<Optional<String>> wrappingTtl() {
        return Codegen.optional(this.wrappingTtl);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuthBackendRoleSecretId(String name) {
        this(name, AuthBackendRoleSecretIdArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuthBackendRoleSecretId(String name, AuthBackendRoleSecretIdArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuthBackendRoleSecretId(String name, AuthBackendRoleSecretIdArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:appRole/authBackendRoleSecretId:AuthBackendRoleSecretId", name, args == null ? AuthBackendRoleSecretIdArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AuthBackendRoleSecretId(String name, Output<String> id, @Nullable AuthBackendRoleSecretIdState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:appRole/authBackendRoleSecretId:AuthBackendRoleSecretId", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "secretId",
                "wrappingToken"
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
    public static AuthBackendRoleSecretId get(String name, Output<String> id, @Nullable AuthBackendRoleSecretIdState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuthBackendRoleSecretId(name, id, state, options);
    }
}