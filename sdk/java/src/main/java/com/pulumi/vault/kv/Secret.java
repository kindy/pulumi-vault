// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.kv;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.kv.SecretArgs;
import com.pulumi.vault.kv.inputs.SecretState;
import java.lang.Object;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Writes a KV-V1 secret to a given path in Vault.
 * 
 * For more information on Vault&#39;s KV-V1 secret backend
 * [see here](https://www.vaultproject.io/docs/secrets/kv/kv-v1).
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.Mount;
 * import com.pulumi.vault.MountArgs;
 * import com.pulumi.vault.kv.Secret;
 * import com.pulumi.vault.kv.SecretArgs;
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
 *         var kvv1 = new Mount(&#34;kvv1&#34;, MountArgs.builder()        
 *             .path(&#34;kvv1&#34;)
 *             .type(&#34;kv&#34;)
 *             .options(Map.of(&#34;version&#34;, &#34;1&#34;))
 *             .description(&#34;KV Version 1 secret engine mount&#34;)
 *             .build());
 * 
 *         var secret = new Secret(&#34;secret&#34;, SecretArgs.builder()        
 *             .path(kvv1.path().applyValue(path -&gt; String.format(&#34;%s/secret&#34;, path)))
 *             .dataJson(serializeJson(
 *                 jsonObject(
 *                     jsonProperty(&#34;zip&#34;, &#34;zap&#34;),
 *                     jsonProperty(&#34;foo&#34;, &#34;bar&#34;)
 *                 )))
 *             .build());
 * 
 *     }
 * }
 * ```
 * ## Required Vault Capabilities
 * 
 * Use of this resource requires the `create` or `update` capability
 * (depending on whether the resource already exists) on the given path,
 * the `delete` capability if the resource is removed from configuration,
 * and the `read` capability for drift detection (by default).
 * 
 * ## Import
 * 
 * KV-V1 secrets can be imported using the `path`, e.g.
 * 
 * ```sh
 *  $ pulumi import vault:kv/secret:Secret secret kvv1/secret
 * ```
 * 
 */
@ResourceType(type="vault:kv/secret:Secret")
public class Secret extends com.pulumi.resources.CustomResource {
    /**
     * A mapping whose keys are the top-level data keys returned from
     * Vault and whose values are the corresponding values. This map can only
     * represent string data, so any non-string values returned from Vault are
     * serialized as JSON.
     * 
     */
    @Export(name="data", type=Map.class, parameters={String.class, Object.class})
    private Output<Map<String,Object>> data;

    /**
     * @return A mapping whose keys are the top-level data keys returned from
     * Vault and whose values are the corresponding values. This map can only
     * represent string data, so any non-string values returned from Vault are
     * serialized as JSON.
     * 
     */
    public Output<Map<String,Object>> data() {
        return this.data;
    }
    /**
     * JSON-encoded string that will be
     * written as the secret data at the given path.
     * 
     */
    @Export(name="dataJson", type=String.class, parameters={})
    private Output<String> dataJson;

    /**
     * @return JSON-encoded string that will be
     * written as the secret data at the given path.
     * 
     */
    public Output<String> dataJson() {
        return this.dataJson;
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
     * Full path of the KV-V1 secret.
     * 
     */
    @Export(name="path", type=String.class, parameters={})
    private Output<String> path;

    /**
     * @return Full path of the KV-V1 secret.
     * 
     */
    public Output<String> path() {
        return this.path;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Secret(String name) {
        this(name, SecretArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Secret(String name, SecretArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Secret(String name, SecretArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:kv/secret:Secret", name, args == null ? SecretArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Secret(String name, Output<String> id, @Nullable SecretState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:kv/secret:Secret", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "data",
                "dataJson"
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
    public static Secret get(String name, Output<String> id, @Nullable SecretState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Secret(name, id, state, options);
    }
}