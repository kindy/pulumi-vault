// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetAuthBackendsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetAuthBackendsPlainArgs Empty = new GetAuthBackendsPlainArgs();

    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable String namespace;

    /**
     * @return The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * The name of the auth method type. Allows filtering of backends returned by type.
     * 
     */
    @Import(name="type")
    private @Nullable String type;

    /**
     * @return The name of the auth method type. Allows filtering of backends returned by type.
     * 
     */
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    private GetAuthBackendsPlainArgs() {}

    private GetAuthBackendsPlainArgs(GetAuthBackendsPlainArgs $) {
        this.namespace = $.namespace;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetAuthBackendsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetAuthBackendsPlainArgs $;

        public Builder() {
            $ = new GetAuthBackendsPlainArgs();
        }

        public Builder(GetAuthBackendsPlainArgs defaults) {
            $ = new GetAuthBackendsPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param namespace The namespace of the target resource.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(@Nullable String namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param type The name of the auth method type. Allows filtering of backends returned by type.
         * 
         * @return builder
         * 
         */
        public Builder type(@Nullable String type) {
            $.type = type;
            return this;
        }

        public GetAuthBackendsPlainArgs build() {
            return $;
        }
    }

}
