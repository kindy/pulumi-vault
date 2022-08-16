// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.NamespaceArgs;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.inputs.NamespaceState;
import java.lang.String;
import javax.annotation.Nullable;

/**
 * Provides a resource to manage [Namespaces](https://www.vaultproject.io/docs/enterprise/namespaces/index.html).
 * 
 * **Note** this feature is available only with Vault Enterprise.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.Namespace;
 * import com.pulumi.vault.NamespaceArgs;
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
 *         var ns1 = new Namespace(&#34;ns1&#34;, NamespaceArgs.builder()        
 *             .path(&#34;ns1&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * Namespaces can be imported using its `name` as accessor id
 * 
 * ```sh
 *  $ pulumi import vault:index/namespace:Namespace example &lt;name&gt;
 * ```
 * 
 *  If the declared resource is imported and intends to support namespaces using a provider alias, then the name is relative to the namespace path. provider &#34;vault&#34; {
 * 
 * # Configuration options
 * 
 *  namespace = &#34;example&#34;
 * 
 *  alias
 * 
 *  = &#34;example&#34; } resource vault_namespace &#34;example2&#34; {
 * 
 *  provider = vault.example }
 * 
 * ```sh
 *  $ pulumi import vault:index/namespace:Namespace example2 example2
 * ```
 * 
 *  $ terraform state show vault_namespace.example2 # vault_namespace.example2 resource &#34;vault_namespace&#34; &#34;example2&#34; {
 * 
 *  id
 * 
 *  = &#34;example/example2/&#34;
 * 
 *  namespace_id = &lt;known after import&gt;
 * 
 *  path
 * 
 *  = &#34;example2&#34; }
 * 
 */
@ResourceType(type="vault:index/namespace:Namespace")
public class Namespace extends com.pulumi.resources.CustomResource {
    /**
     * ID of the namepsace.
     * 
     */
    @Export(name="namespaceId", type=String.class, parameters={})
    private Output<String> namespaceId;

    /**
     * @return ID of the namepsace.
     * 
     */
    public Output<String> namespaceId() {
        return this.namespaceId;
    }
    /**
     * The path of the namespace. Must not have a trailing `/`
     * 
     */
    @Export(name="path", type=String.class, parameters={})
    private Output<String> path;

    /**
     * @return The path of the namespace. Must not have a trailing `/`
     * 
     */
    public Output<String> path() {
        return this.path;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Namespace(String name) {
        this(name, NamespaceArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Namespace(String name, NamespaceArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Namespace(String name, NamespaceArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/namespace:Namespace", name, args == null ? NamespaceArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Namespace(String name, Output<String> id, @Nullable NamespaceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/namespace:Namespace", name, state, makeResourceOptions(options, id));
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
    public static Namespace get(String name, Output<String> id, @Nullable NamespaceState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Namespace(name, id, state, options);
    }
}