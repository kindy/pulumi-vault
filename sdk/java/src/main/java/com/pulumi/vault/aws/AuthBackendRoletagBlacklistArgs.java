// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.aws;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthBackendRoletagBlacklistArgs extends com.pulumi.resources.ResourceArgs {

    public static final AuthBackendRoletagBlacklistArgs Empty = new AuthBackendRoletagBlacklistArgs();

    /**
     * The path the AWS auth backend being configured was
     * mounted at.
     * 
     */
    @Import(name="backend", required=true)
    private Output<String> backend;

    /**
     * @return The path the AWS auth backend being configured was
     * mounted at.
     * 
     */
    public Output<String> backend() {
        return this.backend;
    }

    /**
     * If set to true, disables the periodic
     * tidying of the roletag blacklist entries. Defaults to false.
     * 
     */
    @Import(name="disablePeriodicTidy")
    private @Nullable Output<Boolean> disablePeriodicTidy;

    /**
     * @return If set to true, disables the periodic
     * tidying of the roletag blacklist entries. Defaults to false.
     * 
     */
    public Optional<Output<Boolean>> disablePeriodicTidy() {
        return Optional.ofNullable(this.disablePeriodicTidy);
    }

    /**
     * The amount of extra time that must have passed
     * beyond the roletag expiration, before it is removed from the backend storage.
     * Defaults to 259,200 seconds, or 72 hours.
     * 
     */
    @Import(name="safetyBuffer")
    private @Nullable Output<Integer> safetyBuffer;

    /**
     * @return The amount of extra time that must have passed
     * beyond the roletag expiration, before it is removed from the backend storage.
     * Defaults to 259,200 seconds, or 72 hours.
     * 
     */
    public Optional<Output<Integer>> safetyBuffer() {
        return Optional.ofNullable(this.safetyBuffer);
    }

    private AuthBackendRoletagBlacklistArgs() {}

    private AuthBackendRoletagBlacklistArgs(AuthBackendRoletagBlacklistArgs $) {
        this.backend = $.backend;
        this.disablePeriodicTidy = $.disablePeriodicTidy;
        this.safetyBuffer = $.safetyBuffer;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthBackendRoletagBlacklistArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthBackendRoletagBlacklistArgs $;

        public Builder() {
            $ = new AuthBackendRoletagBlacklistArgs();
        }

        public Builder(AuthBackendRoletagBlacklistArgs defaults) {
            $ = new AuthBackendRoletagBlacklistArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend The path the AWS auth backend being configured was
         * mounted at.
         * 
         * @return builder
         * 
         */
        public Builder backend(Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend The path the AWS auth backend being configured was
         * mounted at.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param disablePeriodicTidy If set to true, disables the periodic
         * tidying of the roletag blacklist entries. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder disablePeriodicTidy(@Nullable Output<Boolean> disablePeriodicTidy) {
            $.disablePeriodicTidy = disablePeriodicTidy;
            return this;
        }

        /**
         * @param disablePeriodicTidy If set to true, disables the periodic
         * tidying of the roletag blacklist entries. Defaults to false.
         * 
         * @return builder
         * 
         */
        public Builder disablePeriodicTidy(Boolean disablePeriodicTidy) {
            return disablePeriodicTidy(Output.of(disablePeriodicTidy));
        }

        /**
         * @param safetyBuffer The amount of extra time that must have passed
         * beyond the roletag expiration, before it is removed from the backend storage.
         * Defaults to 259,200 seconds, or 72 hours.
         * 
         * @return builder
         * 
         */
        public Builder safetyBuffer(@Nullable Output<Integer> safetyBuffer) {
            $.safetyBuffer = safetyBuffer;
            return this;
        }

        /**
         * @param safetyBuffer The amount of extra time that must have passed
         * beyond the roletag expiration, before it is removed from the backend storage.
         * Defaults to 259,200 seconds, or 72 hours.
         * 
         * @return builder
         * 
         */
        public Builder safetyBuffer(Integer safetyBuffer) {
            return safetyBuffer(Output.of(safetyBuffer));
        }

        public AuthBackendRoletagBlacklistArgs build() {
            $.backend = Objects.requireNonNull($.backend, "expected parameter 'backend' to be non-null");
            return $;
        }
    }

}