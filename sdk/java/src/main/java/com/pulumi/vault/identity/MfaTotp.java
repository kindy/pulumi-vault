// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.identity.MfaTotpArgs;
import com.pulumi.vault.identity.inputs.MfaTotpState;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Resource for configuring the totp MFA method.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.identity.MfaTotp;
 * import com.pulumi.vault.identity.MfaTotpArgs;
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
 *         var example = new MfaTotp(&#34;example&#34;, MfaTotpArgs.builder()        
 *             .issuer(&#34;issuer1&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Resource can be imported using its `uuid` field, e.g.
 * 
 * ```sh
 * $ pulumi import vault:identity/mfaTotp:MfaTotp example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec
 * ```
 * 
 */
@ResourceType(type="vault:identity/mfaTotp:MfaTotp")
public class MfaTotp extends com.pulumi.resources.CustomResource {
    /**
     * Specifies the hashing algorithm used to generate the TOTP code. Options include SHA1, SHA256, SHA512.
     * 
     */
    @Export(name="algorithm", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> algorithm;

    /**
     * @return Specifies the hashing algorithm used to generate the TOTP code. Options include SHA1, SHA256, SHA512.
     * 
     */
    public Output<Optional<String>> algorithm() {
        return Codegen.optional(this.algorithm);
    }
    /**
     * The number of digits in the generated TOTP token. This value can either be 6 or 8
     * 
     */
    @Export(name="digits", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> digits;

    /**
     * @return The number of digits in the generated TOTP token. This value can either be 6 or 8
     * 
     */
    public Output<Optional<Integer>> digits() {
        return Codegen.optional(this.digits);
    }
    /**
     * The name of the key&#39;s issuing organization.
     * 
     */
    @Export(name="issuer", refs={String.class}, tree="[0]")
    private Output<String> issuer;

    /**
     * @return The name of the key&#39;s issuing organization.
     * 
     */
    public Output<String> issuer() {
        return this.issuer;
    }
    /**
     * Specifies the size in bytes of the generated key.
     * 
     */
    @Export(name="keySize", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> keySize;

    /**
     * @return Specifies the size in bytes of the generated key.
     * 
     */
    public Output<Optional<Integer>> keySize() {
        return Codegen.optional(this.keySize);
    }
    /**
     * The maximum number of consecutive failed validation attempts allowed.
     * 
     */
    @Export(name="maxValidationAttempts", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxValidationAttempts;

    /**
     * @return The maximum number of consecutive failed validation attempts allowed.
     * 
     */
    public Output<Optional<Integer>> maxValidationAttempts() {
        return Codegen.optional(this.maxValidationAttempts);
    }
    /**
     * Method ID.
     * 
     */
    @Export(name="methodId", refs={String.class}, tree="[0]")
    private Output<String> methodId;

    /**
     * @return Method ID.
     * 
     */
    public Output<String> methodId() {
        return this.methodId;
    }
    /**
     * Mount accessor.
     * 
     */
    @Export(name="mountAccessor", refs={String.class}, tree="[0]")
    private Output<String> mountAccessor;

    /**
     * @return Mount accessor.
     * 
     */
    public Output<String> mountAccessor() {
        return this.mountAccessor;
    }
    /**
     * Method name.
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Method name.
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * Target namespace. (requires Enterprise)
     * 
     */
    @Export(name="namespace", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> namespace;

    /**
     * @return Target namespace. (requires Enterprise)
     * 
     */
    public Output<Optional<String>> namespace() {
        return Codegen.optional(this.namespace);
    }
    /**
     * Method&#39;s namespace ID.
     * 
     */
    @Export(name="namespaceId", refs={String.class}, tree="[0]")
    private Output<String> namespaceId;

    /**
     * @return Method&#39;s namespace ID.
     * 
     */
    public Output<String> namespaceId() {
        return this.namespaceId;
    }
    /**
     * Method&#39;s namespace path.
     * 
     */
    @Export(name="namespacePath", refs={String.class}, tree="[0]")
    private Output<String> namespacePath;

    /**
     * @return Method&#39;s namespace path.
     * 
     */
    public Output<String> namespacePath() {
        return this.namespacePath;
    }
    /**
     * The length of time in seconds used to generate a counter for the TOTP token calculation.
     * 
     */
    @Export(name="period", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> period;

    /**
     * @return The length of time in seconds used to generate a counter for the TOTP token calculation.
     * 
     */
    public Output<Optional<Integer>> period() {
        return Codegen.optional(this.period);
    }
    /**
     * The pixel size of the generated square QR code.
     * 
     */
    @Export(name="qrSize", refs={Integer.class}, tree="[0]")
    private Output<Integer> qrSize;

    /**
     * @return The pixel size of the generated square QR code.
     * 
     */
    public Output<Integer> qrSize() {
        return this.qrSize;
    }
    /**
     * The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1.
     * 
     */
    @Export(name="skew", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> skew;

    /**
     * @return The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1.
     * 
     */
    public Output<Optional<Integer>> skew() {
        return Codegen.optional(this.skew);
    }
    /**
     * MFA type.
     * 
     */
    @Export(name="type", refs={String.class}, tree="[0]")
    private Output<String> type;

    /**
     * @return MFA type.
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * Resource UUID.
     * 
     */
    @Export(name="uuid", refs={String.class}, tree="[0]")
    private Output<String> uuid;

    /**
     * @return Resource UUID.
     * 
     */
    public Output<String> uuid() {
        return this.uuid;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public MfaTotp(String name) {
        this(name, MfaTotpArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public MfaTotp(String name, MfaTotpArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public MfaTotp(String name, MfaTotpArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:identity/mfaTotp:MfaTotp", name, args == null ? MfaTotpArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private MfaTotp(String name, Output<String> id, @Nullable MfaTotpState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:identity/mfaTotp:MfaTotp", name, state, makeResourceOptions(options, id));
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
    public static MfaTotp get(String name, Output<String> id, @Nullable MfaTotpState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new MfaTotp(name, id, state, options);
    }
}
