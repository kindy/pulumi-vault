// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.appRole;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthBackendLoginArgs extends com.pulumi.resources.ResourceArgs {

    public static final AuthBackendLoginArgs Empty = new AuthBackendLoginArgs();

    /**
     * The unique path of the Vault backend to log in with.
     * 
     */
    @Import(name="backend")
    private @Nullable Output<String> backend;

    /**
     * @return The unique path of the Vault backend to log in with.
     * 
     */
    public Optional<Output<String>> backend() {
        return Optional.ofNullable(this.backend);
    }

    /**
     * The ID of the role to log in with.
     * 
     */
    @Import(name="roleId", required=true)
    private Output<String> roleId;

    /**
     * @return The ID of the role to log in with.
     * 
     */
    public Output<String> roleId() {
        return this.roleId;
    }

    /**
     * The secret ID of the role to log in with. Required
     * unless `bind_secret_id` is set to false on the role.
     * 
     */
    @Import(name="secretId")
    private @Nullable Output<String> secretId;

    /**
     * @return The secret ID of the role to log in with. Required
     * unless `bind_secret_id` is set to false on the role.
     * 
     */
    public Optional<Output<String>> secretId() {
        return Optional.ofNullable(this.secretId);
    }

    private AuthBackendLoginArgs() {}

    private AuthBackendLoginArgs(AuthBackendLoginArgs $) {
        this.backend = $.backend;
        this.roleId = $.roleId;
        this.secretId = $.secretId;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthBackendLoginArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthBackendLoginArgs $;

        public Builder() {
            $ = new AuthBackendLoginArgs();
        }

        public Builder(AuthBackendLoginArgs defaults) {
            $ = new AuthBackendLoginArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend The unique path of the Vault backend to log in with.
         * 
         * @return builder
         * 
         */
        public Builder backend(@Nullable Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend The unique path of the Vault backend to log in with.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param roleId The ID of the role to log in with.
         * 
         * @return builder
         * 
         */
        public Builder roleId(Output<String> roleId) {
            $.roleId = roleId;
            return this;
        }

        /**
         * @param roleId The ID of the role to log in with.
         * 
         * @return builder
         * 
         */
        public Builder roleId(String roleId) {
            return roleId(Output.of(roleId));
        }

        /**
         * @param secretId The secret ID of the role to log in with. Required
         * unless `bind_secret_id` is set to false on the role.
         * 
         * @return builder
         * 
         */
        public Builder secretId(@Nullable Output<String> secretId) {
            $.secretId = secretId;
            return this;
        }

        /**
         * @param secretId The secret ID of the role to log in with. Required
         * unless `bind_secret_id` is set to false on the role.
         * 
         * @return builder
         * 
         */
        public Builder secretId(String secretId) {
            return secretId(Output.of(secretId));
        }

        public AuthBackendLoginArgs build() {
            $.roleId = Objects.requireNonNull($.roleId, "expected parameter 'roleId' to be non-null");
            return $;
        }
    }

}