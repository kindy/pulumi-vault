// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class OidcState extends com.pulumi.resources.ResourceArgs {

    public static final OidcState Empty = new OidcState();

    /**
     * Issuer URL to be used in the iss claim of the token. If not set, Vault&#39;s
     * `api_addr` will be used. The issuer is a case sensitive URL using the https scheme that contains
     * scheme, host, and optionally, port number and path components, but no query or fragment
     * components.
     * 
     */
    @Import(name="issuer")
    private @Nullable Output<String> issuer;

    /**
     * @return Issuer URL to be used in the iss claim of the token. If not set, Vault&#39;s
     * `api_addr` will be used. The issuer is a case sensitive URL using the https scheme that contains
     * scheme, host, and optionally, port number and path components, but no query or fragment
     * components.
     * 
     */
    public Optional<Output<String>> issuer() {
        return Optional.ofNullable(this.issuer);
    }

    private OidcState() {}

    private OidcState(OidcState $) {
        this.issuer = $.issuer;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(OidcState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private OidcState $;

        public Builder() {
            $ = new OidcState();
        }

        public Builder(OidcState defaults) {
            $ = new OidcState(Objects.requireNonNull(defaults));
        }

        /**
         * @param issuer Issuer URL to be used in the iss claim of the token. If not set, Vault&#39;s
         * `api_addr` will be used. The issuer is a case sensitive URL using the https scheme that contains
         * scheme, host, and optionally, port number and path components, but no query or fragment
         * components.
         * 
         * @return builder
         * 
         */
        public Builder issuer(@Nullable Output<String> issuer) {
            $.issuer = issuer;
            return this;
        }

        /**
         * @param issuer Issuer URL to be used in the iss claim of the token. If not set, Vault&#39;s
         * `api_addr` will be used. The issuer is a case sensitive URL using the https scheme that contains
         * scheme, host, and optionally, port number and path components, but no query or fragment
         * components.
         * 
         * @return builder
         * 
         */
        public Builder issuer(String issuer) {
            return issuer(Output.of(issuer));
        }

        public OidcState build() {
            return $;
        }
    }

}