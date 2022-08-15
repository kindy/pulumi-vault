// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.gcp.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class SecretStaticAccountBindingArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretStaticAccountBindingArgs Empty = new SecretStaticAccountBindingArgs();

    /**
     * Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#bindings).
     * 
     */
    @Import(name="resource", required=true)
    private Output<String> resource;

    /**
     * @return Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#bindings).
     * 
     */
    public Output<String> resource() {
        return this.resource;
    }

    /**
     * List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource.
     * 
     */
    @Import(name="roles", required=true)
    private Output<List<String>> roles;

    /**
     * @return List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource.
     * 
     */
    public Output<List<String>> roles() {
        return this.roles;
    }

    private SecretStaticAccountBindingArgs() {}

    private SecretStaticAccountBindingArgs(SecretStaticAccountBindingArgs $) {
        this.resource = $.resource;
        this.roles = $.roles;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretStaticAccountBindingArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretStaticAccountBindingArgs $;

        public Builder() {
            $ = new SecretStaticAccountBindingArgs();
        }

        public Builder(SecretStaticAccountBindingArgs defaults) {
            $ = new SecretStaticAccountBindingArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param resource Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#bindings).
         * 
         * @return builder
         * 
         */
        public Builder resource(Output<String> resource) {
            $.resource = resource;
            return this;
        }

        /**
         * @param resource Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#bindings).
         * 
         * @return builder
         * 
         */
        public Builder resource(String resource) {
            return resource(Output.of(resource));
        }

        /**
         * @param roles List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource.
         * 
         * @return builder
         * 
         */
        public Builder roles(Output<List<String>> roles) {
            $.roles = roles;
            return this;
        }

        /**
         * @param roles List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource.
         * 
         * @return builder
         * 
         */
        public Builder roles(List<String> roles) {
            return roles(Output.of(roles));
        }

        /**
         * @param roles List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource.
         * 
         * @return builder
         * 
         */
        public Builder roles(String... roles) {
            return roles(List.of(roles));
        }

        public SecretStaticAccountBindingArgs build() {
            $.resource = Objects.requireNonNull($.resource, "expected parameter 'resource' to be non-null");
            $.roles = Objects.requireNonNull($.roles, "expected parameter 'roles' to be non-null");
            return $;
        }
    }

}
