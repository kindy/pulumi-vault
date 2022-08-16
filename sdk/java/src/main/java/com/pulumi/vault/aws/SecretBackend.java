// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.aws;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.aws.SecretBackendArgs;
import com.pulumi.vault.aws.inputs.SecretBackendState;
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
 * import com.pulumi.vault.aws.SecretBackend;
 * import com.pulumi.vault.aws.SecretBackendArgs;
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
 *         var aws = new SecretBackend(&#34;aws&#34;, SecretBackendArgs.builder()        
 *             .accessKey(&#34;AKIA.....&#34;)
 *             .secretKey(&#34;AWS secret key&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * 
 * ## Import
 * 
 * AWS secret backends can be imported using the `path`, e.g.
 * 
 * ```sh
 *  $ pulumi import vault:aws/secretBackend:SecretBackend aws aws
 * ```
 * 
 */
@ResourceType(type="vault:aws/secretBackend:SecretBackend")
public class SecretBackend extends com.pulumi.resources.CustomResource {
    /**
     * The AWS Access Key ID this backend should use to
     * issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
     * 
     */
    @Export(name="accessKey", type=String.class, parameters={})
    private Output</* @Nullable */ String> accessKey;

    /**
     * @return The AWS Access Key ID this backend should use to
     * issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
     * 
     */
    public Output<Optional<String>> accessKey() {
        return Codegen.optional(this.accessKey);
    }
    /**
     * The default TTL for credentials
     * issued by this backend.
     * 
     */
    @Export(name="defaultLeaseTtlSeconds", type=Integer.class, parameters={})
    private Output<Integer> defaultLeaseTtlSeconds;

    /**
     * @return The default TTL for credentials
     * issued by this backend.
     * 
     */
    public Output<Integer> defaultLeaseTtlSeconds() {
        return this.defaultLeaseTtlSeconds;
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
     * Specifies a custom HTTP IAM endpoint to use.
     * 
     */
    @Export(name="iamEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> iamEndpoint;

    /**
     * @return Specifies a custom HTTP IAM endpoint to use.
     * 
     */
    public Output<Optional<String>> iamEndpoint() {
        return Codegen.optional(this.iamEndpoint);
    }
    /**
     * The maximum TTL that can be requested
     * for credentials issued by this backend.
     * 
     */
    @Export(name="maxLeaseTtlSeconds", type=Integer.class, parameters={})
    private Output<Integer> maxLeaseTtlSeconds;

    /**
     * @return The maximum TTL that can be requested
     * for credentials issued by this backend.
     * 
     */
    public Output<Integer> maxLeaseTtlSeconds() {
        return this.maxLeaseTtlSeconds;
    }
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `aws`.
     * 
     */
    @Export(name="path", type=String.class, parameters={})
    private Output</* @Nullable */ String> path;

    /**
     * @return The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `aws`.
     * 
     */
    public Output<Optional<String>> path() {
        return Codegen.optional(this.path);
    }
    /**
     * The AWS region for API calls. Defaults to `us-east-1`.
     * 
     */
    @Export(name="region", type=String.class, parameters={})
    private Output<String> region;

    /**
     * @return The AWS region for API calls. Defaults to `us-east-1`.
     * 
     */
    public Output<String> region() {
        return this.region;
    }
    /**
     * The AWS Secret Key this backend should use to
     * issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
     * 
     */
    @Export(name="secretKey", type=String.class, parameters={})
    private Output</* @Nullable */ String> secretKey;

    /**
     * @return The AWS Secret Key this backend should use to
     * issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
     * 
     */
    public Output<Optional<String>> secretKey() {
        return Codegen.optional(this.secretKey);
    }
    /**
     * Specifies a custom HTTP STS endpoint to use.
     * 
     */
    @Export(name="stsEndpoint", type=String.class, parameters={})
    private Output</* @Nullable */ String> stsEndpoint;

    /**
     * @return Specifies a custom HTTP STS endpoint to use.
     * 
     */
    public Output<Optional<String>> stsEndpoint() {
        return Codegen.optional(this.stsEndpoint);
    }
    /**
     * Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template:
     * 
     */
    @Export(name="usernameTemplate", type=String.class, parameters={})
    private Output<String> usernameTemplate;

    /**
     * @return Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template:
     * 
     */
    public Output<String> usernameTemplate() {
        return this.usernameTemplate;
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
        super("vault:aws/secretBackend:SecretBackend", name, args == null ? SecretBackendArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private SecretBackend(String name, Output<String> id, @Nullable SecretBackendState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:aws/secretBackend:SecretBackend", name, state, makeResourceOptions(options, id));
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