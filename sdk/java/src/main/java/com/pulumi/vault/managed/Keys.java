// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.managed;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.managed.KeysArgs;
import com.pulumi.vault.managed.inputs.KeysState;
import com.pulumi.vault.managed.outputs.KeysAw;
import com.pulumi.vault.managed.outputs.KeysAzure;
import com.pulumi.vault.managed.outputs.KeysPkc;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * ## Import
 * 
 * Mounts can be imported using the `id` of `default`, e.g.
 * 
 * ```sh
 *  $ pulumi import vault:managed/keys:Keys keys default
 * ```
 * 
 */
@ResourceType(type="vault:managed/keys:Keys")
public class Keys extends com.pulumi.resources.CustomResource {
    /**
     * Configuration block for AWS Managed Keys
     * 
     */
    @Export(name="aws", type=List.class, parameters={KeysAw.class})
    private Output</* @Nullable */ List<KeysAw>> aws;

    /**
     * @return Configuration block for AWS Managed Keys
     * 
     */
    public Output<Optional<List<KeysAw>>> aws() {
        return Codegen.optional(this.aws);
    }
    /**
     * Configuration block for Azure Managed Keys
     * 
     */
    @Export(name="azures", type=List.class, parameters={KeysAzure.class})
    private Output</* @Nullable */ List<KeysAzure>> azures;

    /**
     * @return Configuration block for Azure Managed Keys
     * 
     */
    public Output<Optional<List<KeysAzure>>> azures() {
        return Codegen.optional(this.azures);
    }
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured namespace.
     * *Available only for Vault Enterprise*.
     * 
     */
    @Export(name="namespace", type=String.class, parameters={})
    private Output</* @Nullable */ String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured namespace.
     * *Available only for Vault Enterprise*.
     * 
     */
    public Output<Optional<String>> namespace() {
        return Codegen.optional(this.namespace);
    }
    /**
     * Configuration block for PKCS Managed Keys
     * 
     */
    @Export(name="pkcs", type=List.class, parameters={KeysPkc.class})
    private Output</* @Nullable */ List<KeysPkc>> pkcs;

    /**
     * @return Configuration block for PKCS Managed Keys
     * 
     */
    public Output<Optional<List<KeysPkc>>> pkcs() {
        return Codegen.optional(this.pkcs);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public Keys(String name) {
        this(name, KeysArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public Keys(String name, @Nullable KeysArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public Keys(String name, @Nullable KeysArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:managed/keys:Keys", name, args == null ? KeysArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private Keys(String name, Output<String> id, @Nullable KeysState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:managed/keys:Keys", name, state, makeResourceOptions(options, id));
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
    public static Keys get(String name, Output<String> id, @Nullable KeysState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new Keys(name, id, state, options);
    }
}