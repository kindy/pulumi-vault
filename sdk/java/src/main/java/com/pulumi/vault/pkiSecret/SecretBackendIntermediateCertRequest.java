// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.pkiSecret.SecretBackendIntermediateCertRequestArgs;
import com.pulumi.vault.pkiSecret.inputs.SecretBackendIntermediateCertRequestState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
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
 * import com.pulumi.vault.pkiSecret.SecretBackendIntermediateCertRequest;
 * import com.pulumi.vault.pkiSecret.SecretBackendIntermediateCertRequestArgs;
 * import com.pulumi.resources.CustomResourceOptions;
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
 *         var test = new SecretBackendIntermediateCertRequest(&#34;test&#34;, SecretBackendIntermediateCertRequestArgs.builder()        
 *             .backend(vault_mount.pki().path())
 *             .type(&#34;internal&#34;)
 *             .commonName(&#34;app.my.domain&#34;)
 *             .build(), CustomResourceOptions.builder()
 *                 .dependsOn(vault_mount.pki())
 *                 .build());
 * 
 *     }
 * }
 * ```
 * 
 */
@ResourceType(type="vault:pkiSecret/secretBackendIntermediateCertRequest:SecretBackendIntermediateCertRequest")
public class SecretBackendIntermediateCertRequest extends com.pulumi.resources.CustomResource {
    /**
     * Adds a Basic Constraints extension with &#39;CA: true&#39;.
     * Only needed as a workaround in some compatibility scenarios with Active Directory
     * Certificate Services
     * 
     */
    @Export(name="addBasicConstraints", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> addBasicConstraints;

    /**
     * @return Adds a Basic Constraints extension with &#39;CA: true&#39;.
     * Only needed as a workaround in some compatibility scenarios with Active Directory
     * Certificate Services
     * 
     */
    public Output<Optional<Boolean>> addBasicConstraints() {
        return Codegen.optional(this.addBasicConstraints);
    }
    /**
     * List of alternative names
     * 
     */
    @Export(name="altNames", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> altNames;

    /**
     * @return List of alternative names
     * 
     */
    public Output<Optional<List<String>>> altNames() {
        return Codegen.optional(this.altNames);
    }
    /**
     * The PKI secret backend the resource belongs to.
     * 
     */
    @Export(name="backend", type=String.class, parameters={})
    private Output<String> backend;

    /**
     * @return The PKI secret backend the resource belongs to.
     * 
     */
    public Output<String> backend() {
        return this.backend;
    }
    /**
     * CN of intermediate to create
     * 
     */
    @Export(name="commonName", type=String.class, parameters={})
    private Output<String> commonName;

    /**
     * @return CN of intermediate to create
     * 
     */
    public Output<String> commonName() {
        return this.commonName;
    }
    /**
     * The country
     * 
     */
    @Export(name="country", type=String.class, parameters={})
    private Output</* @Nullable */ String> country;

    /**
     * @return The country
     * 
     */
    public Output<Optional<String>> country() {
        return Codegen.optional(this.country);
    }
    /**
     * The CSR
     * 
     */
    @Export(name="csr", type=String.class, parameters={})
    private Output<String> csr;

    /**
     * @return The CSR
     * 
     */
    public Output<String> csr() {
        return this.csr;
    }
    /**
     * Flag to exclude CN from SANs
     * 
     */
    @Export(name="excludeCnFromSans", type=Boolean.class, parameters={})
    private Output</* @Nullable */ Boolean> excludeCnFromSans;

    /**
     * @return Flag to exclude CN from SANs
     * 
     */
    public Output<Optional<Boolean>> excludeCnFromSans() {
        return Codegen.optional(this.excludeCnFromSans);
    }
    /**
     * The format of data
     * 
     */
    @Export(name="format", type=String.class, parameters={})
    private Output</* @Nullable */ String> format;

    /**
     * @return The format of data
     * 
     */
    public Output<Optional<String>> format() {
        return Codegen.optional(this.format);
    }
    /**
     * List of alternative IPs
     * 
     */
    @Export(name="ipSans", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> ipSans;

    /**
     * @return List of alternative IPs
     * 
     */
    public Output<Optional<List<String>>> ipSans() {
        return Codegen.optional(this.ipSans);
    }
    /**
     * The number of bits to use
     * 
     */
    @Export(name="keyBits", type=Integer.class, parameters={})
    private Output</* @Nullable */ Integer> keyBits;

    /**
     * @return The number of bits to use
     * 
     */
    public Output<Optional<Integer>> keyBits() {
        return Codegen.optional(this.keyBits);
    }
    /**
     * The desired key type
     * 
     */
    @Export(name="keyType", type=String.class, parameters={})
    private Output</* @Nullable */ String> keyType;

    /**
     * @return The desired key type
     * 
     */
    public Output<Optional<String>> keyType() {
        return Codegen.optional(this.keyType);
    }
    /**
     * The locality
     * 
     */
    @Export(name="locality", type=String.class, parameters={})
    private Output</* @Nullable */ String> locality;

    /**
     * @return The locality
     * 
     */
    public Output<Optional<String>> locality() {
        return Codegen.optional(this.locality);
    }
    /**
     * The ID of the previously configured managed key. This field is
     * required if `type` is `kms` and it conflicts with `managed_key_name`
     * 
     */
    @Export(name="managedKeyId", type=String.class, parameters={})
    private Output</* @Nullable */ String> managedKeyId;

    /**
     * @return The ID of the previously configured managed key. This field is
     * required if `type` is `kms` and it conflicts with `managed_key_name`
     * 
     */
    public Output<Optional<String>> managedKeyId() {
        return Codegen.optional(this.managedKeyId);
    }
    /**
     * The name of the previously configured managed key. This field is
     * required if `type` is `kms`  and it conflicts with `managed_key_id`
     * 
     */
    @Export(name="managedKeyName", type=String.class, parameters={})
    private Output</* @Nullable */ String> managedKeyName;

    /**
     * @return The name of the previously configured managed key. This field is
     * required if `type` is `kms`  and it conflicts with `managed_key_id`
     * 
     */
    public Output<Optional<String>> managedKeyName() {
        return Codegen.optional(this.managedKeyName);
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
     * The organization
     * 
     */
    @Export(name="organization", type=String.class, parameters={})
    private Output</* @Nullable */ String> organization;

    /**
     * @return The organization
     * 
     */
    public Output<Optional<String>> organization() {
        return Codegen.optional(this.organization);
    }
    /**
     * List of other SANs
     * 
     */
    @Export(name="otherSans", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> otherSans;

    /**
     * @return List of other SANs
     * 
     */
    public Output<Optional<List<String>>> otherSans() {
        return Codegen.optional(this.otherSans);
    }
    /**
     * The organization unit
     * 
     */
    @Export(name="ou", type=String.class, parameters={})
    private Output</* @Nullable */ String> ou;

    /**
     * @return The organization unit
     * 
     */
    public Output<Optional<String>> ou() {
        return Codegen.optional(this.ou);
    }
    /**
     * The postal code
     * 
     */
    @Export(name="postalCode", type=String.class, parameters={})
    private Output</* @Nullable */ String> postalCode;

    /**
     * @return The postal code
     * 
     */
    public Output<Optional<String>> postalCode() {
        return Codegen.optional(this.postalCode);
    }
    /**
     * The private key
     * 
     */
    @Export(name="privateKey", type=String.class, parameters={})
    private Output<String> privateKey;

    /**
     * @return The private key
     * 
     */
    public Output<String> privateKey() {
        return this.privateKey;
    }
    /**
     * The private key format
     * 
     */
    @Export(name="privateKeyFormat", type=String.class, parameters={})
    private Output</* @Nullable */ String> privateKeyFormat;

    /**
     * @return The private key format
     * 
     */
    public Output<Optional<String>> privateKeyFormat() {
        return Codegen.optional(this.privateKeyFormat);
    }
    /**
     * The private key type
     * 
     */
    @Export(name="privateKeyType", type=String.class, parameters={})
    private Output<String> privateKeyType;

    /**
     * @return The private key type
     * 
     */
    public Output<String> privateKeyType() {
        return this.privateKeyType;
    }
    /**
     * The province
     * 
     */
    @Export(name="province", type=String.class, parameters={})
    private Output</* @Nullable */ String> province;

    /**
     * @return The province
     * 
     */
    public Output<Optional<String>> province() {
        return Codegen.optional(this.province);
    }
    /**
     * The street address
     * 
     */
    @Export(name="streetAddress", type=String.class, parameters={})
    private Output</* @Nullable */ String> streetAddress;

    /**
     * @return The street address
     * 
     */
    public Output<Optional<String>> streetAddress() {
        return Codegen.optional(this.streetAddress);
    }
    /**
     * Type of intermediate to create. Must be either \&#34;exported\&#34; or \&#34;internal\&#34;
     * or \&#34;kms\&#34;
     * 
     */
    @Export(name="type", type=String.class, parameters={})
    private Output<String> type;

    /**
     * @return Type of intermediate to create. Must be either \&#34;exported\&#34; or \&#34;internal\&#34;
     * or \&#34;kms\&#34;
     * 
     */
    public Output<String> type() {
        return this.type;
    }
    /**
     * List of alternative URIs
     * 
     */
    @Export(name="uriSans", type=List.class, parameters={String.class})
    private Output</* @Nullable */ List<String>> uriSans;

    /**
     * @return List of alternative URIs
     * 
     */
    public Output<Optional<List<String>>> uriSans() {
        return Codegen.optional(this.uriSans);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public SecretBackendIntermediateCertRequest(String name) {
        this(name, SecretBackendIntermediateCertRequestArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public SecretBackendIntermediateCertRequest(String name, SecretBackendIntermediateCertRequestArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretBackendIntermediateCertRequest(String name, SecretBackendIntermediateCertRequestArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:pkiSecret/secretBackendIntermediateCertRequest:SecretBackendIntermediateCertRequest", name, args == null ? SecretBackendIntermediateCertRequestArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecretBackendIntermediateCertRequest(String name, Output<String> id, @Nullable SecretBackendIntermediateCertRequestState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:pkiSecret/secretBackendIntermediateCertRequest:SecretBackendIntermediateCertRequest", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .additionalSecretOutputs(List.of(
                "privateKey"
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
    public static SecretBackendIntermediateCertRequest get(String name, Output<String> id, @Nullable SecretBackendIntermediateCertRequestState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new SecretBackendIntermediateCertRequest(name, id, state, options);
    }
}
