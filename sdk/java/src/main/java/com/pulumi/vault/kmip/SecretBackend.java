// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.kmip;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.kmip.SecretBackendArgs;
import com.pulumi.vault.kmip.inputs.SecretBackendState;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages KMIP Secret backends in a Vault server. This feature requires
 * Vault Enterprise. See the [Vault documentation](https://www.vaultproject.io/docs/secrets/kmip)
 * for more information.
 * 
 * ## Example Usage
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.kmip.SecretBackend;
 * import com.pulumi.vault.kmip.SecretBackendArgs;
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
 *         var default_ = new SecretBackend(&#34;default&#34;, SecretBackendArgs.builder()        
 *             .defaultTlsClientKeyBits(4096)
 *             .defaultTlsClientKeyType(&#34;rsa&#34;)
 *             .defaultTlsClientTtl(86400)
 *             .description(&#34;Vault KMIP backend&#34;)
 *             .listenAddrs(            
 *                 &#34;127.0.0.1:5696&#34;,
 *                 &#34;127.0.0.1:8080&#34;)
 *             .path(&#34;kmip&#34;)
 *             .tlsCaKeyBits(4096)
 *             .tlsCaKeyType(&#34;rsa&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * KMIP Secret backend can be imported using the `path`, e.g.
 * 
 * ```sh
 *  $ pulumi import vault:kmip/secretBackend:SecretBackend default kmip
 * ```
 * 
 */
@ResourceType(type="vault:kmip/secretBackend:SecretBackend")
public class SecretBackend extends com.pulumi.resources.CustomResource {
    /**
     * Client certificate key bits, valid values depend on key type.
     * 
     */
    @Export(name="defaultTlsClientKeyBits", type=Integer.class, parameters={})
    private Output<Integer> defaultTlsClientKeyBits;

    /**
     * @return Client certificate key bits, valid values depend on key type.
     * 
     */
    public Output<Integer> defaultTlsClientKeyBits() {
        return this.defaultTlsClientKeyBits;
    }
    /**
     * Client certificate key type, `rsa` or `ec`.
     * 
     */
    @Export(name="defaultTlsClientKeyType", type=String.class, parameters={})
    private Output<String> defaultTlsClientKeyType;

    /**
     * @return Client certificate key type, `rsa` or `ec`.
     * 
     */
    public Output<String> defaultTlsClientKeyType() {
        return this.defaultTlsClientKeyType;
    }
    /**
     * Client certificate TTL in seconds
     * 
     */
    @Export(name="defaultTlsClientTtl", type=Integer.class, parameters={})
    private Output<Integer> defaultTlsClientTtl;

    /**
     * @return Client certificate TTL in seconds
     * 
     */
    public Output<Integer> defaultTlsClientTtl() {
        return this.defaultTlsClientTtl;
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
     * Addresses the KMIP server should listen on (`host:port`).
     * 
     */
    @Export(name="listenAddrs", type=List.class, parameters={String.class})
    private Output<List<String>> listenAddrs;

    /**
     * @return Addresses the KMIP server should listen on (`host:port`).
     * 
     */
    public Output<List<String>> listenAddrs() {
        return this.listenAddrs;
    }
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `kmip`.
     * 
     */
    @Export(name="path", type=String.class, parameters={})
    private Output<String> path;

    /**
     * @return The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `kmip`.
     * 
     */
    public Output<String> path() {
        return this.path;
    }
    /**
     * Hostnames to include in the server&#39;s TLS certificate as SAN DNS names. The first will be used as the common name (CN).
     * 
     */
    @Export(name="serverHostnames", type=List.class, parameters={String.class})
    private Output<List<String>> serverHostnames;

    /**
     * @return Hostnames to include in the server&#39;s TLS certificate as SAN DNS names. The first will be used as the common name (CN).
     * 
     */
    public Output<List<String>> serverHostnames() {
        return this.serverHostnames;
    }
    /**
     * IPs to include in the server&#39;s TLS certificate as SAN IP addresses.
     * 
     */
    @Export(name="serverIps", type=List.class, parameters={String.class})
    private Output<List<String>> serverIps;

    /**
     * @return IPs to include in the server&#39;s TLS certificate as SAN IP addresses.
     * 
     */
    public Output<List<String>> serverIps() {
        return this.serverIps;
    }
    /**
     * CA key bits, valid values depend on key type.
     * 
     */
    @Export(name="tlsCaKeyBits", type=Integer.class, parameters={})
    private Output<Integer> tlsCaKeyBits;

    /**
     * @return CA key bits, valid values depend on key type.
     * 
     */
    public Output<Integer> tlsCaKeyBits() {
        return this.tlsCaKeyBits;
    }
    /**
     * CA key type, rsa or ec.
     * 
     */
    @Export(name="tlsCaKeyType", type=String.class, parameters={})
    private Output<String> tlsCaKeyType;

    /**
     * @return CA key type, rsa or ec.
     * 
     */
    public Output<String> tlsCaKeyType() {
        return this.tlsCaKeyType;
    }
    /**
     * Minimum TLS version to accept.
     * 
     */
    @Export(name="tlsMinVersion", type=String.class, parameters={})
    private Output<String> tlsMinVersion;

    /**
     * @return Minimum TLS version to accept.
     * 
     */
    public Output<String> tlsMinVersion() {
        return this.tlsMinVersion;
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
    public SecretBackend(String name, SecretBackendArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public SecretBackend(String name, SecretBackendArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:kmip/secretBackend:SecretBackend", name, args == null ? SecretBackendArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecretBackend(String name, Output<String> id, @Nullable SecretBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:kmip/secretBackend:SecretBackend", name, state, makeResourceOptions(options, id));
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