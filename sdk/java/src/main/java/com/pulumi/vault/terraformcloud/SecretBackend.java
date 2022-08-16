// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.terraformcloud;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.terraformcloud.SecretBackendArgs;
import com.pulumi.vault.terraformcloud.inputs.SecretBackendState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.terraformcloud.SecretBackend;
 * import com.pulumi.vault.terraformcloud.SecretBackendArgs;
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
 *         var test = new SecretBackend(&#34;test&#34;, SecretBackendArgs.builder()        
 *             .backend(&#34;terraform&#34;)
 *             .description(&#34;Manages the Terraform Cloud backend&#34;)
 *             .token(&#34;V0idfhi2iksSDU234ucdbi2nidsi...&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Terraform Cloud secret backends can be imported using the `backend`, e.g.
 * 
 * ```sh
 *  $ pulumi import vault:terraformcloud/secretBackend:SecretBackend example terraform
 * ```
 * 
 */
@ResourceType(type="vault:terraformcloud/secretBackend:SecretBackend")
public class SecretBackend extends com.pulumi.resources.CustomResource {
    /**
     * Specifies the address of the Terraform Cloud instance, provided as &#34;host:port&#34; like &#34;127.0.0.1:8500&#34;.
     * 
     */
    @Export(name="address", type=String.class, parameters={})
    private Output</* @Nullable */ String> address;

    /**
     * @return Specifies the address of the Terraform Cloud instance, provided as &#34;host:port&#34; like &#34;127.0.0.1:8500&#34;.
     * 
     */
    public Output<Optional<String>> address() {
        return Codegen.optional(this.address);
    }
    /**
     * Unique name of the Vault Terraform Cloud mount to configure
     * 
     */
    @Export(name="backend", type=String.class, parameters={})
    private Output</* @Nullable */ String> backend;

    /**
     * @return Unique name of the Vault Terraform Cloud mount to configure
     * 
     */
    public Output<Optional<String>> backend() {
        return Codegen.optional(this.backend);
    }
    /**
     * Specifies the base path for the Terraform Cloud or Enterprise API.
     * 
     */
    @Export(name="basePath", type=String.class, parameters={})
    private Output</* @Nullable */ String> basePath;

    /**
     * @return Specifies the base path for the Terraform Cloud or Enterprise API.
     * 
     */
    public Output<Optional<String>> basePath() {
        return Codegen.optional(this.basePath);
    }
    /**
     * The default TTL for credentials issued by this backend.
     * 
     */
    @Export(name="defaultLeaseTtlSeconds", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> defaultLeaseTtlSeconds;

    /**
     * @return The default TTL for credentials issued by this backend.
     * 
     */
    public Output<Optional<Integer>> defaultLeaseTtlSeconds() {
        return Codegen.optional(this.defaultLeaseTtlSeconds);
    }
    /**
     * A human-friendly description for this backend.
     * 
     */
    @Export(name="description", type=String.class, parameters={})
    private Output</* @Nullable */ String> description;

    /**
     * @return A human-friendly description for this backend.
     * 
     */
    public Output<Optional<String>> description() {
        return Codegen.optional(this.description);
    }
    /**
     * The maximum TTL that can be requested
     * for credentials issued by this backend.
     * 
     */
    @Export(name="maxLeaseTtlSeconds", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> maxLeaseTtlSeconds;

    /**
     * @return The maximum TTL that can be requested
     * for credentials issued by this backend.
     * 
     */
    public Output<Optional<Integer>> maxLeaseTtlSeconds() {
        return Codegen.optional(this.maxLeaseTtlSeconds);
    }
    /**
     * Specifies the Terraform Cloud access token to use.
     * 
     */
    @Export(name="token", type=String.class, parameters={})
    private Output</* @Nullable */ String> token;

    /**
     * @return Specifies the Terraform Cloud access token to use.
     * 
     */
    public Output<Optional<String>> token() {
        return Codegen.optional(this.token);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretBackend(String name) {
        this(name, SecretBackendArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretBackend(String name, @Nullable SecretBackendArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretBackend(String name, @Nullable SecretBackendArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:terraformcloud/secretBackend:SecretBackend", name, args == null ? SecretBackendArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecretBackend(String name, Output<String> id, @Nullable SecretBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:terraformcloud/secretBackend:SecretBackend", name, state, makeResourceOptions(options, id));
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
    public static SecretBackend get(String name, Output<String> id, @Nullable SecretBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretBackend(name, id, state, options);
    }
}