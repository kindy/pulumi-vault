// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretBackendRootSignIntermediateArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretBackendRootSignIntermediateArgs Empty = new SecretBackendRootSignIntermediateArgs();

    /**
     * List of alternative names
     * 
     */
    @Import(name="altNames")
    private @Nullable Output<List<String>> altNames;

    /**
     * @return List of alternative names
     * 
     */
    public Optional<Output<List<String>>> altNames() {
        return Optional.ofNullable(this.altNames);
    }

    /**
     * The PKI secret backend the resource belongs to.
     * 
     */
    @Import(name="backend", required=true)
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
    @Import(name="commonName", required=true)
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
    @Import(name="country")
    private @Nullable Output<String> country;

    /**
     * @return The country
     * 
     */
    public Optional<Output<String>> country() {
        return Optional.ofNullable(this.country);
    }

    /**
     * The CSR
     * 
     */
    @Import(name="csr", required=true)
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
    @Import(name="excludeCnFromSans")
    private @Nullable Output<Boolean> excludeCnFromSans;

    /**
     * @return Flag to exclude CN from SANs
     * 
     */
    public Optional<Output<Boolean>> excludeCnFromSans() {
        return Optional.ofNullable(this.excludeCnFromSans);
    }

    /**
     * The format of data
     * 
     */
    @Import(name="format")
    private @Nullable Output<String> format;

    /**
     * @return The format of data
     * 
     */
    public Optional<Output<String>> format() {
        return Optional.ofNullable(this.format);
    }

    /**
     * List of alternative IPs
     * 
     */
    @Import(name="ipSans")
    private @Nullable Output<List<String>> ipSans;

    /**
     * @return List of alternative IPs
     * 
     */
    public Optional<Output<List<String>>> ipSans() {
        return Optional.ofNullable(this.ipSans);
    }

    /**
     * The locality
     * 
     */
    @Import(name="locality")
    private @Nullable Output<String> locality;

    /**
     * @return The locality
     * 
     */
    public Optional<Output<String>> locality() {
        return Optional.ofNullable(this.locality);
    }

    /**
     * The maximum path length to encode in the generated certificate
     * 
     */
    @Import(name="maxPathLength")
    private @Nullable Output<Integer> maxPathLength;

    /**
     * @return The maximum path length to encode in the generated certificate
     * 
     */
    public Optional<Output<Integer>> maxPathLength() {
        return Optional.ofNullable(this.maxPathLength);
    }

    /**
     * The organization
     * 
     */
    @Import(name="organization")
    private @Nullable Output<String> organization;

    /**
     * @return The organization
     * 
     */
    public Optional<Output<String>> organization() {
        return Optional.ofNullable(this.organization);
    }

    /**
     * List of other SANs
     * 
     */
    @Import(name="otherSans")
    private @Nullable Output<List<String>> otherSans;

    /**
     * @return List of other SANs
     * 
     */
    public Optional<Output<List<String>>> otherSans() {
        return Optional.ofNullable(this.otherSans);
    }

    /**
     * The organization unit
     * 
     */
    @Import(name="ou")
    private @Nullable Output<String> ou;

    /**
     * @return The organization unit
     * 
     */
    public Optional<Output<String>> ou() {
        return Optional.ofNullable(this.ou);
    }

    /**
     * List of domains for which certificates are allowed to be issued
     * 
     */
    @Import(name="permittedDnsDomains")
    private @Nullable Output<List<String>> permittedDnsDomains;

    /**
     * @return List of domains for which certificates are allowed to be issued
     * 
     */
    public Optional<Output<List<String>>> permittedDnsDomains() {
        return Optional.ofNullable(this.permittedDnsDomains);
    }

    /**
     * The postal code
     * 
     */
    @Import(name="postalCode")
    private @Nullable Output<String> postalCode;

    /**
     * @return The postal code
     * 
     */
    public Optional<Output<String>> postalCode() {
        return Optional.ofNullable(this.postalCode);
    }

    /**
     * The province
     * 
     */
    @Import(name="province")
    private @Nullable Output<String> province;

    /**
     * @return The province
     * 
     */
    public Optional<Output<String>> province() {
        return Optional.ofNullable(this.province);
    }

    /**
     * If set to `true`, the certificate will be revoked on resource destruction.
     * 
     */
    @Import(name="revoke")
    private @Nullable Output<Boolean> revoke;

    /**
     * @return If set to `true`, the certificate will be revoked on resource destruction.
     * 
     */
    public Optional<Output<Boolean>> revoke() {
        return Optional.ofNullable(this.revoke);
    }

    /**
     * The street address
     * 
     */
    @Import(name="streetAddress")
    private @Nullable Output<String> streetAddress;

    /**
     * @return The street address
     * 
     */
    public Optional<Output<String>> streetAddress() {
        return Optional.ofNullable(this.streetAddress);
    }

    /**
     * Time to live
     * 
     */
    @Import(name="ttl")
    private @Nullable Output<String> ttl;

    /**
     * @return Time to live
     * 
     */
    public Optional<Output<String>> ttl() {
        return Optional.ofNullable(this.ttl);
    }

    /**
     * List of alternative URIs
     * 
     */
    @Import(name="uriSans")
    private @Nullable Output<List<String>> uriSans;

    /**
     * @return List of alternative URIs
     * 
     */
    public Optional<Output<List<String>>> uriSans() {
        return Optional.ofNullable(this.uriSans);
    }

    /**
     * Preserve CSR values
     * 
     */
    @Import(name="useCsrValues")
    private @Nullable Output<Boolean> useCsrValues;

    /**
     * @return Preserve CSR values
     * 
     */
    public Optional<Output<Boolean>> useCsrValues() {
        return Optional.ofNullable(this.useCsrValues);
    }

    private SecretBackendRootSignIntermediateArgs() {}

    private SecretBackendRootSignIntermediateArgs(SecretBackendRootSignIntermediateArgs $) {
        this.altNames = $.altNames;
        this.backend = $.backend;
        this.commonName = $.commonName;
        this.country = $.country;
        this.csr = $.csr;
        this.excludeCnFromSans = $.excludeCnFromSans;
        this.format = $.format;
        this.ipSans = $.ipSans;
        this.locality = $.locality;
        this.maxPathLength = $.maxPathLength;
        this.organization = $.organization;
        this.otherSans = $.otherSans;
        this.ou = $.ou;
        this.permittedDnsDomains = $.permittedDnsDomains;
        this.postalCode = $.postalCode;
        this.province = $.province;
        this.revoke = $.revoke;
        this.streetAddress = $.streetAddress;
        this.ttl = $.ttl;
        this.uriSans = $.uriSans;
        this.useCsrValues = $.useCsrValues;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretBackendRootSignIntermediateArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretBackendRootSignIntermediateArgs $;

        public Builder() {
            $ = new SecretBackendRootSignIntermediateArgs();
        }

        public Builder(SecretBackendRootSignIntermediateArgs defaults) {
            $ = new SecretBackendRootSignIntermediateArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param altNames List of alternative names
         * 
         * @return builder
         * 
         */
        public Builder altNames(@Nullable Output<List<String>> altNames) {
            $.altNames = altNames;
            return this;
        }

        /**
         * @param altNames List of alternative names
         * 
         * @return builder
         * 
         */
        public Builder altNames(List<String> altNames) {
            return altNames(Output.of(altNames));
        }

        /**
         * @param altNames List of alternative names
         * 
         * @return builder
         * 
         */
        public Builder altNames(String... altNames) {
            return altNames(List.of(altNames));
        }

        /**
         * @param backend The PKI secret backend the resource belongs to.
         * 
         * @return builder
         * 
         */
        public Builder backend(Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend The PKI secret backend the resource belongs to.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param commonName CN of intermediate to create
         * 
         * @return builder
         * 
         */
        public Builder commonName(Output<String> commonName) {
            $.commonName = commonName;
            return this;
        }

        /**
         * @param commonName CN of intermediate to create
         * 
         * @return builder
         * 
         */
        public Builder commonName(String commonName) {
            return commonName(Output.of(commonName));
        }

        /**
         * @param country The country
         * 
         * @return builder
         * 
         */
        public Builder country(@Nullable Output<String> country) {
            $.country = country;
            return this;
        }

        /**
         * @param country The country
         * 
         * @return builder
         * 
         */
        public Builder country(String country) {
            return country(Output.of(country));
        }

        /**
         * @param csr The CSR
         * 
         * @return builder
         * 
         */
        public Builder csr(Output<String> csr) {
            $.csr = csr;
            return this;
        }

        /**
         * @param csr The CSR
         * 
         * @return builder
         * 
         */
        public Builder csr(String csr) {
            return csr(Output.of(csr));
        }

        /**
         * @param excludeCnFromSans Flag to exclude CN from SANs
         * 
         * @return builder
         * 
         */
        public Builder excludeCnFromSans(@Nullable Output<Boolean> excludeCnFromSans) {
            $.excludeCnFromSans = excludeCnFromSans;
            return this;
        }

        /**
         * @param excludeCnFromSans Flag to exclude CN from SANs
         * 
         * @return builder
         * 
         */
        public Builder excludeCnFromSans(Boolean excludeCnFromSans) {
            return excludeCnFromSans(Output.of(excludeCnFromSans));
        }

        /**
         * @param format The format of data
         * 
         * @return builder
         * 
         */
        public Builder format(@Nullable Output<String> format) {
            $.format = format;
            return this;
        }

        /**
         * @param format The format of data
         * 
         * @return builder
         * 
         */
        public Builder format(String format) {
            return format(Output.of(format));
        }

        /**
         * @param ipSans List of alternative IPs
         * 
         * @return builder
         * 
         */
        public Builder ipSans(@Nullable Output<List<String>> ipSans) {
            $.ipSans = ipSans;
            return this;
        }

        /**
         * @param ipSans List of alternative IPs
         * 
         * @return builder
         * 
         */
        public Builder ipSans(List<String> ipSans) {
            return ipSans(Output.of(ipSans));
        }

        /**
         * @param ipSans List of alternative IPs
         * 
         * @return builder
         * 
         */
        public Builder ipSans(String... ipSans) {
            return ipSans(List.of(ipSans));
        }

        /**
         * @param locality The locality
         * 
         * @return builder
         * 
         */
        public Builder locality(@Nullable Output<String> locality) {
            $.locality = locality;
            return this;
        }

        /**
         * @param locality The locality
         * 
         * @return builder
         * 
         */
        public Builder locality(String locality) {
            return locality(Output.of(locality));
        }

        /**
         * @param maxPathLength The maximum path length to encode in the generated certificate
         * 
         * @return builder
         * 
         */
        public Builder maxPathLength(@Nullable Output<Integer> maxPathLength) {
            $.maxPathLength = maxPathLength;
            return this;
        }

        /**
         * @param maxPathLength The maximum path length to encode in the generated certificate
         * 
         * @return builder
         * 
         */
        public Builder maxPathLength(Integer maxPathLength) {
            return maxPathLength(Output.of(maxPathLength));
        }

        /**
         * @param organization The organization
         * 
         * @return builder
         * 
         */
        public Builder organization(@Nullable Output<String> organization) {
            $.organization = organization;
            return this;
        }

        /**
         * @param organization The organization
         * 
         * @return builder
         * 
         */
        public Builder organization(String organization) {
            return organization(Output.of(organization));
        }

        /**
         * @param otherSans List of other SANs
         * 
         * @return builder
         * 
         */
        public Builder otherSans(@Nullable Output<List<String>> otherSans) {
            $.otherSans = otherSans;
            return this;
        }

        /**
         * @param otherSans List of other SANs
         * 
         * @return builder
         * 
         */
        public Builder otherSans(List<String> otherSans) {
            return otherSans(Output.of(otherSans));
        }

        /**
         * @param otherSans List of other SANs
         * 
         * @return builder
         * 
         */
        public Builder otherSans(String... otherSans) {
            return otherSans(List.of(otherSans));
        }

        /**
         * @param ou The organization unit
         * 
         * @return builder
         * 
         */
        public Builder ou(@Nullable Output<String> ou) {
            $.ou = ou;
            return this;
        }

        /**
         * @param ou The organization unit
         * 
         * @return builder
         * 
         */
        public Builder ou(String ou) {
            return ou(Output.of(ou));
        }

        /**
         * @param permittedDnsDomains List of domains for which certificates are allowed to be issued
         * 
         * @return builder
         * 
         */
        public Builder permittedDnsDomains(@Nullable Output<List<String>> permittedDnsDomains) {
            $.permittedDnsDomains = permittedDnsDomains;
            return this;
        }

        /**
         * @param permittedDnsDomains List of domains for which certificates are allowed to be issued
         * 
         * @return builder
         * 
         */
        public Builder permittedDnsDomains(List<String> permittedDnsDomains) {
            return permittedDnsDomains(Output.of(permittedDnsDomains));
        }

        /**
         * @param permittedDnsDomains List of domains for which certificates are allowed to be issued
         * 
         * @return builder
         * 
         */
        public Builder permittedDnsDomains(String... permittedDnsDomains) {
            return permittedDnsDomains(List.of(permittedDnsDomains));
        }

        /**
         * @param postalCode The postal code
         * 
         * @return builder
         * 
         */
        public Builder postalCode(@Nullable Output<String> postalCode) {
            $.postalCode = postalCode;
            return this;
        }

        /**
         * @param postalCode The postal code
         * 
         * @return builder
         * 
         */
        public Builder postalCode(String postalCode) {
            return postalCode(Output.of(postalCode));
        }

        /**
         * @param province The province
         * 
         * @return builder
         * 
         */
        public Builder province(@Nullable Output<String> province) {
            $.province = province;
            return this;
        }

        /**
         * @param province The province
         * 
         * @return builder
         * 
         */
        public Builder province(String province) {
            return province(Output.of(province));
        }

        /**
         * @param revoke If set to `true`, the certificate will be revoked on resource destruction.
         * 
         * @return builder
         * 
         */
        public Builder revoke(@Nullable Output<Boolean> revoke) {
            $.revoke = revoke;
            return this;
        }

        /**
         * @param revoke If set to `true`, the certificate will be revoked on resource destruction.
         * 
         * @return builder
         * 
         */
        public Builder revoke(Boolean revoke) {
            return revoke(Output.of(revoke));
        }

        /**
         * @param streetAddress The street address
         * 
         * @return builder
         * 
         */
        public Builder streetAddress(@Nullable Output<String> streetAddress) {
            $.streetAddress = streetAddress;
            return this;
        }

        /**
         * @param streetAddress The street address
         * 
         * @return builder
         * 
         */
        public Builder streetAddress(String streetAddress) {
            return streetAddress(Output.of(streetAddress));
        }

        /**
         * @param ttl Time to live
         * 
         * @return builder
         * 
         */
        public Builder ttl(@Nullable Output<String> ttl) {
            $.ttl = ttl;
            return this;
        }

        /**
         * @param ttl Time to live
         * 
         * @return builder
         * 
         */
        public Builder ttl(String ttl) {
            return ttl(Output.of(ttl));
        }

        /**
         * @param uriSans List of alternative URIs
         * 
         * @return builder
         * 
         */
        public Builder uriSans(@Nullable Output<List<String>> uriSans) {
            $.uriSans = uriSans;
            return this;
        }

        /**
         * @param uriSans List of alternative URIs
         * 
         * @return builder
         * 
         */
        public Builder uriSans(List<String> uriSans) {
            return uriSans(Output.of(uriSans));
        }

        /**
         * @param uriSans List of alternative URIs
         * 
         * @return builder
         * 
         */
        public Builder uriSans(String... uriSans) {
            return uriSans(List.of(uriSans));
        }

        /**
         * @param useCsrValues Preserve CSR values
         * 
         * @return builder
         * 
         */
        public Builder useCsrValues(@Nullable Output<Boolean> useCsrValues) {
            $.useCsrValues = useCsrValues;
            return this;
        }

        /**
         * @param useCsrValues Preserve CSR values
         * 
         * @return builder
         * 
         */
        public Builder useCsrValues(Boolean useCsrValues) {
            return useCsrValues(Output.of(useCsrValues));
        }

        public SecretBackendRootSignIntermediateArgs build() {
            $.backend = Objects.requireNonNull($.backend, "expected parameter 'backend' to be non-null");
            $.commonName = Objects.requireNonNull($.commonName, "expected parameter 'commonName' to be non-null");
            $.csr = Objects.requireNonNull($.csr, "expected parameter 'csr' to be non-null");
            return $;
        }
    }

}
