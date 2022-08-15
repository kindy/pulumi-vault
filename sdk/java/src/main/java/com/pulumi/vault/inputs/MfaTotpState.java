// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class MfaTotpState extends com.pulumi.resources.ResourceArgs {

    public static final MfaTotpState Empty = new MfaTotpState();

    /**
     * `(string)` - Specifies the hashing algorithm used to generate the TOTP code.
     * Options include `SHA1`, `SHA256` and `SHA512`
     * 
     */
    @Import(name="algorithm")
    private @Nullable Output<String> algorithm;

    /**
     * @return `(string)` - Specifies the hashing algorithm used to generate the TOTP code.
     * Options include `SHA1`, `SHA256` and `SHA512`
     * 
     */
    public Optional<Output<String>> algorithm() {
        return Optional.ofNullable(this.algorithm);
    }

    /**
     * `(int)` - The number of digits in the generated TOTP token.
     * This value can either be 6 or 8.
     * 
     */
    @Import(name="digits")
    private @Nullable Output<Integer> digits;

    /**
     * @return `(int)` - The number of digits in the generated TOTP token.
     * This value can either be 6 or 8.
     * 
     */
    public Optional<Output<Integer>> digits() {
        return Optional.ofNullable(this.digits);
    }

    /**
     * `(string: &lt;required&gt;)` - The name of the key&#39;s issuing organization.
     * 
     */
    @Import(name="issuer")
    private @Nullable Output<String> issuer;

    /**
     * @return `(string: &lt;required&gt;)` - The name of the key&#39;s issuing organization.
     * 
     */
    public Optional<Output<String>> issuer() {
        return Optional.ofNullable(this.issuer);
    }

    /**
     * `(int)` - Specifies the size in bytes of the generated key.
     * 
     */
    @Import(name="keySize")
    private @Nullable Output<Integer> keySize;

    /**
     * @return `(int)` - Specifies the size in bytes of the generated key.
     * 
     */
    public Optional<Output<Integer>> keySize() {
        return Optional.ofNullable(this.keySize);
    }

    /**
     * `(string: &lt;required&gt;)` – Name of the MFA method.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return `(string: &lt;required&gt;)` – Name of the MFA method.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * `(int)` - The length of time used to generate a counter for the TOTP token calculation.
     * 
     */
    @Import(name="period")
    private @Nullable Output<Integer> period;

    /**
     * @return `(int)` - The length of time used to generate a counter for the TOTP token calculation.
     * 
     */
    public Optional<Output<Integer>> period() {
        return Optional.ofNullable(this.period);
    }

    /**
     * `(int)` - The pixel size of the generated square QR code.
     * 
     */
    @Import(name="qrSize")
    private @Nullable Output<Integer> qrSize;

    /**
     * @return `(int)` - The pixel size of the generated square QR code.
     * 
     */
    public Optional<Output<Integer>> qrSize() {
        return Optional.ofNullable(this.qrSize);
    }

    /**
     * `(int)` - The number of delay periods that are allowed when validating a TOTP token.
     * This value can either be 0 or 1.
     * 
     */
    @Import(name="skew")
    private @Nullable Output<Integer> skew;

    /**
     * @return `(int)` - The number of delay periods that are allowed when validating a TOTP token.
     * This value can either be 0 or 1.
     * 
     */
    public Optional<Output<Integer>> skew() {
        return Optional.ofNullable(this.skew);
    }

    private MfaTotpState() {}

    private MfaTotpState(MfaTotpState $) {
        this.algorithm = $.algorithm;
        this.digits = $.digits;
        this.issuer = $.issuer;
        this.keySize = $.keySize;
        this.name = $.name;
        this.period = $.period;
        this.qrSize = $.qrSize;
        this.skew = $.skew;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(MfaTotpState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private MfaTotpState $;

        public Builder() {
            $ = new MfaTotpState();
        }

        public Builder(MfaTotpState defaults) {
            $ = new MfaTotpState(Objects.requireNonNull(defaults));
        }

        /**
         * @param algorithm `(string)` - Specifies the hashing algorithm used to generate the TOTP code.
         * Options include `SHA1`, `SHA256` and `SHA512`
         * 
         * @return builder
         * 
         */
        public Builder algorithm(@Nullable Output<String> algorithm) {
            $.algorithm = algorithm;
            return this;
        }

        /**
         * @param algorithm `(string)` - Specifies the hashing algorithm used to generate the TOTP code.
         * Options include `SHA1`, `SHA256` and `SHA512`
         * 
         * @return builder
         * 
         */
        public Builder algorithm(String algorithm) {
            return algorithm(Output.of(algorithm));
        }

        /**
         * @param digits `(int)` - The number of digits in the generated TOTP token.
         * This value can either be 6 or 8.
         * 
         * @return builder
         * 
         */
        public Builder digits(@Nullable Output<Integer> digits) {
            $.digits = digits;
            return this;
        }

        /**
         * @param digits `(int)` - The number of digits in the generated TOTP token.
         * This value can either be 6 or 8.
         * 
         * @return builder
         * 
         */
        public Builder digits(Integer digits) {
            return digits(Output.of(digits));
        }

        /**
         * @param issuer `(string: &lt;required&gt;)` - The name of the key&#39;s issuing organization.
         * 
         * @return builder
         * 
         */
        public Builder issuer(@Nullable Output<String> issuer) {
            $.issuer = issuer;
            return this;
        }

        /**
         * @param issuer `(string: &lt;required&gt;)` - The name of the key&#39;s issuing organization.
         * 
         * @return builder
         * 
         */
        public Builder issuer(String issuer) {
            return issuer(Output.of(issuer));
        }

        /**
         * @param keySize `(int)` - Specifies the size in bytes of the generated key.
         * 
         * @return builder
         * 
         */
        public Builder keySize(@Nullable Output<Integer> keySize) {
            $.keySize = keySize;
            return this;
        }

        /**
         * @param keySize `(int)` - Specifies the size in bytes of the generated key.
         * 
         * @return builder
         * 
         */
        public Builder keySize(Integer keySize) {
            return keySize(Output.of(keySize));
        }

        /**
         * @param name `(string: &lt;required&gt;)` – Name of the MFA method.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name `(string: &lt;required&gt;)` – Name of the MFA method.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param period `(int)` - The length of time used to generate a counter for the TOTP token calculation.
         * 
         * @return builder
         * 
         */
        public Builder period(@Nullable Output<Integer> period) {
            $.period = period;
            return this;
        }

        /**
         * @param period `(int)` - The length of time used to generate a counter for the TOTP token calculation.
         * 
         * @return builder
         * 
         */
        public Builder period(Integer period) {
            return period(Output.of(period));
        }

        /**
         * @param qrSize `(int)` - The pixel size of the generated square QR code.
         * 
         * @return builder
         * 
         */
        public Builder qrSize(@Nullable Output<Integer> qrSize) {
            $.qrSize = qrSize;
            return this;
        }

        /**
         * @param qrSize `(int)` - The pixel size of the generated square QR code.
         * 
         * @return builder
         * 
         */
        public Builder qrSize(Integer qrSize) {
            return qrSize(Output.of(qrSize));
        }

        /**
         * @param skew `(int)` - The number of delay periods that are allowed when validating a TOTP token.
         * This value can either be 0 or 1.
         * 
         * @return builder
         * 
         */
        public Builder skew(@Nullable Output<Integer> skew) {
            $.skew = skew;
            return this;
        }

        /**
         * @param skew `(int)` - The number of delay periods that are allowed when validating a TOTP token.
         * This value can either be 0 or 1.
         * 
         * @return builder
         * 
         */
        public Builder skew(Integer skew) {
            return skew(Output.of(skew));
        }

        public MfaTotpState build() {
            return $;
        }
    }

}
