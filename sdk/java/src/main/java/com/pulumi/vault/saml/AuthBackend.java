// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.saml;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.saml.AuthBackendArgs;
import com.pulumi.vault.saml.inputs.AuthBackendState;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages a SAML Auth mount in a Vault server. See the [Vault
 * documentation](https://www.vaultproject.io/docs/auth/saml/) for more
 * information.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.saml.AuthBackend;
 * import com.pulumi.vault.saml.AuthBackendArgs;
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
 *         var test = new AuthBackend(&#34;test&#34;, AuthBackendArgs.builder()        
 *             .acsUrls(&#34;https://my.vault.primary/v1/auth/saml/callback&#34;)
 *             .defaultRole(&#34;admin&#34;)
 *             .entityId(&#34;https://my.vault/v1/auth/saml&#34;)
 *             .idpMetadataUrl(&#34;https://company.okta.com/app/abc123eb9xnIfzlaf697/sso/saml/metadata&#34;)
 *             .path(&#34;saml&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * SAML authentication mounts can be imported using the `path`, e.g.
 * 
 * ```sh
 *  $ pulumi import vault:saml/authBackend:AuthBackend example saml
 * ```
 * 
 */
@ResourceType(type="vault:saml/authBackend:AuthBackend")
public class AuthBackend extends com.pulumi.resources.CustomResource {
    /**
     * The well-formatted URLs of your Assertion Consumer Service (ACS)
     * that should receive a response from the identity provider.
     * 
     */
    @Export(name="acsUrls", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> acsUrls;

    /**
     * @return The well-formatted URLs of your Assertion Consumer Service (ACS)
     * that should receive a response from the identity provider.
     * 
     */
    public Output<List<String>> acsUrls() {
        return this.acsUrls;
    }
    /**
     * The role to use if no role is provided during login.
     * 
     */
    @Export(name="defaultRole", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> defaultRole;

    /**
     * @return The role to use if no role is provided during login.
     * 
     */
    public Output<Optional<String>> defaultRole() {
        return Codegen.optional(this.defaultRole);
    }
    /**
     * If set to `true`, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     * 
     */
    @Export(name="disableRemount", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableRemount;

    /**
     * @return If set to `true`, opts out of mount migration on path updates.
     * See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
     * 
     */
    public Output<Optional<Boolean>> disableRemount() {
        return Codegen.optional(this.disableRemount);
    }
    /**
     * The entity ID of the SAML authentication service provider.
     * 
     */
    @Export(name="entityId", refs={String.class}, tree="[0]")
    private Output<String> entityId;

    /**
     * @return The entity ID of the SAML authentication service provider.
     * 
     */
    public Output<String> entityId() {
        return this.entityId;
    }
    /**
     * The PEM encoded certificate of the identity provider. Mutually exclusive
     * with `idp_metadata_url`.
     * 
     */
    @Export(name="idpCert", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> idpCert;

    /**
     * @return The PEM encoded certificate of the identity provider. Mutually exclusive
     * with `idp_metadata_url`.
     * 
     */
    public Output<Optional<String>> idpCert() {
        return Codegen.optional(this.idpCert);
    }
    /**
     * The entity ID of the identity provider. Mutually exclusive with
     * `idp_metadata_url`.
     * 
     */
    @Export(name="idpEntityId", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> idpEntityId;

    /**
     * @return The entity ID of the identity provider. Mutually exclusive with
     * `idp_metadata_url`.
     * 
     */
    public Output<Optional<String>> idpEntityId() {
        return Codegen.optional(this.idpEntityId);
    }
    /**
     * The metadata URL of the identity provider.
     * 
     */
    @Export(name="idpMetadataUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> idpMetadataUrl;

    /**
     * @return The metadata URL of the identity provider.
     * 
     */
    public Output<Optional<String>> idpMetadataUrl() {
        return Codegen.optional(this.idpMetadataUrl);
    }
    /**
     * The SSO URL of the identity provider. Mutually exclusive with
     * `idp_metadata_url`.
     * 
     */
    @Export(name="idpSsoUrl", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> idpSsoUrl;

    /**
     * @return The SSO URL of the identity provider. Mutually exclusive with
     * `idp_metadata_url`.
     * 
     */
    public Output<Optional<String>> idpSsoUrl() {
        return Codegen.optional(this.idpSsoUrl);
    }
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Export(name="namespace", refs={String.class}, tree="[0]")
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
     * Path where the auth backend will be mounted. Defaults to `auth/saml`
     * if not specified.
     * 
     */
    @Export(name="path", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> path;

    /**
     * @return Path where the auth backend will be mounted. Defaults to `auth/saml`
     * if not specified.
     * 
     */
    public Output<Optional<String>> path() {
        return Codegen.optional(this.path);
    }
    /**
     * If set to `true`, logs additional, potentially sensitive
     * information during the SAML exchange according to the current logging level. Not
     * recommended for production.
     * 
     */
    @Export(name="verboseLogging", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> verboseLogging;

    /**
     * @return If set to `true`, logs additional, potentially sensitive
     * information during the SAML exchange according to the current logging level. Not
     * recommended for production.
     * 
     */
    public Output<Boolean> verboseLogging() {
        return this.verboseLogging;
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
        super("vault:saml/authBackend:AuthBackend", name, args == null ? AuthBackendArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AuthBackend(String name, Output<String> id, @Nullable AuthBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:saml/authBackend:AuthBackend", name, state, makeResourceOptions(options, id));
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