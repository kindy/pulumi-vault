// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.generic.inputs;

import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetSecretPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetSecretPlainArgs Empty = new GetSecretPlainArgs();

    /**
     * The full logical path from which to request data.
     * To read data from the &#34;generic&#34; secret backend mounted in Vault by
     * default, this should be prefixed with `secret/`. Reading from other backends
     * with this data source is possible; consult each backend&#39;s documentation
     * to see which endpoints support the `GET` method.
     * 
     */
    @Import(name="path", required=true)
    private String path;

    /**
     * @return The full logical path from which to request data.
     * To read data from the &#34;generic&#34; secret backend mounted in Vault by
     * default, this should be prefixed with `secret/`. Reading from other backends
     * with this data source is possible; consult each backend&#39;s documentation
     * to see which endpoints support the `GET` method.
     * 
     */
    public String path() {
        return this.path;
    }

    /**
     * The version of the secret to read. This is used by the
     * Vault KV secrets engine - version 2 to indicate which version of the secret
     * to read.
     * 
     */
    @Import(name="version")
    private @Nullable Integer version;

    /**
     * @return The version of the secret to read. This is used by the
     * Vault KV secrets engine - version 2 to indicate which version of the secret
     * to read.
     * 
     */
    public Optional<Integer> version() {
        return Optional.ofNullable(this.version);
    }

    @Import(name="withLeaseStartTime")
    private @Nullable Boolean withLeaseStartTime;

    public Optional<Boolean> withLeaseStartTime() {
        return Optional.ofNullable(this.withLeaseStartTime);
    }

    private GetSecretPlainArgs() {}

    private GetSecretPlainArgs(GetSecretPlainArgs $) {
        this.path = $.path;
        this.version = $.version;
        this.withLeaseStartTime = $.withLeaseStartTime;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetSecretPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetSecretPlainArgs $;

        public Builder() {
            $ = new GetSecretPlainArgs();
        }

        public Builder(GetSecretPlainArgs defaults) {
            $ = new GetSecretPlainArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param path The full logical path from which to request data.
         * To read data from the &#34;generic&#34; secret backend mounted in Vault by
         * default, this should be prefixed with `secret/`. Reading from other backends
         * with this data source is possible; consult each backend&#39;s documentation
         * to see which endpoints support the `GET` method.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            $.path = path;
            return this;
        }

        /**
         * @param version The version of the secret to read. This is used by the
         * Vault KV secrets engine - version 2 to indicate which version of the secret
         * to read.
         * 
         * @return builder
         * 
         */
        public Builder version(@Nullable Integer version) {
            $.version = version;
            return this;
        }

        public Builder withLeaseStartTime(@Nullable Boolean withLeaseStartTime) {
            $.withLeaseStartTime = withLeaseStartTime;
            return this;
        }

        public GetSecretPlainArgs build() {
            $.path = Objects.requireNonNull($.path, "expected parameter 'path' to be non-null");
            return $;
        }
    }

}