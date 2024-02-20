// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.secrets;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.secrets.SyncGcpDestinationArgs;
import com.pulumi.vault.secrets.inputs.SyncGcpDestinationState;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
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
 * import com.pulumi.vault.secrets.SyncGcpDestination;
 * import com.pulumi.vault.secrets.SyncGcpDestinationArgs;
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
 *         var gcp = new SyncGcpDestination(&#34;gcp&#34;, SyncGcpDestinationArgs.builder()        
 *             .credentials(Files.readString(Paths.get(var_.credentials_file())))
 *             .secretNameTemplate(&#34;vault_{{ .MountAccessor | lowercase }}_{{ .SecretPath | lowercase }}&#34;)
 *             .customTags(Map.of(&#34;foo&#34;, &#34;bar&#34;))
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * GCP Secrets sync destinations can be imported using the `name`, e.g.
 * 
 * ```sh
 *  $ pulumi import vault:secrets/syncGcpDestination:SyncGcpDestination gcp gcp-dest
 * ```
 * 
 */
@ResourceType(type="vault:secrets/syncGcpDestination:SyncGcpDestination")
public class SyncGcpDestination extends com.pulumi.resources.CustomResource {
    /**
     * JSON-encoded credentials to use to connect to GCP.
     * Can be omitted and directly provided to Vault using the `GOOGLE_APPLICATION_CREDENTIALS` environment
     * variable.
     * 
     */
    @Export(name="credentials", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> credentials;

    /**
     * @return JSON-encoded credentials to use to connect to GCP.
     * Can be omitted and directly provided to Vault using the `GOOGLE_APPLICATION_CREDENTIALS` environment
     * variable.
     * 
     */
    public Output<Optional<String>> credentials() {
        return Codegen.optional(this.credentials);
    }
    /**
     * Custom tags to set on the secret managed at the destination.
     * 
     */
    @Export(name="customTags", refs={Map.class,String.class,Object.class}, tree="[0,1,2]")
    private Output</* @Nullable */ Map<String,Object>> customTags;

    /**
     * @return Custom tags to set on the secret managed at the destination.
     * 
     */
    public Output<Optional<Map<String,Object>>> customTags() {
        return Codegen.optional(this.customTags);
    }
    /**
     * Unique name of the GCP destination.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Unique name of the GCP destination.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * 
     */
    @Export(name="namespace", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * 
     */
    public Output<Optional<String>> namespace() {
        return Codegen.optional(this.namespace);
    }
    /**
     * Template describing how to generate external secret names.
     * Supports a subset of the Go Template syntax.
     * 
     */
    @Export(name="secretNameTemplate", refs={String.class}, tree="[0]")
    private Output<String> secretNameTemplate;

    /**
     * @return Template describing how to generate external secret names.
     * Supports a subset of the Go Template syntax.
     * 
     */
    public Output<String> secretNameTemplate() {
        return this.secretNameTemplate;
    }
    /**
     * The type of the secrets destination (`gcp-sm`).
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return The type of the secrets destination (`gcp-sm`).
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SyncGcpDestination(String name) {
        this(name, SyncGcpDestinationArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SyncGcpDestination(String name, @Nullable SyncGcpDestinationArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SyncGcpDestination(String name, @Nullable SyncGcpDestinationArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:secrets/syncGcpDestination:SyncGcpDestination", name, args == null ? SyncGcpDestinationArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SyncGcpDestination(String name, Output<String> id, @Nullable SyncGcpDestinationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:secrets/syncGcpDestination:SyncGcpDestination", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "credentials"
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
    public static SyncGcpDestination get(String name, Output<String> id, @Nullable SyncGcpDestinationState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SyncGcpDestination(name, id, state, options);
    }
}
