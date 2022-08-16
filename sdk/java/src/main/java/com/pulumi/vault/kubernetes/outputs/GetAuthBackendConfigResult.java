// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.kubernetes.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetAuthBackendConfigResult {
    private final @Nullable String backend;
    private final Boolean disableIssValidation;
    private final Boolean disableLocalCaJwt;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private final String id;
    /**
     * @return Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
     * 
     */
    private final String issuer;
    /**
     * @return PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
     * 
     */
    private final String kubernetesCaCert;
    /**
     * @return Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
     * 
     */
    private final String kubernetesHost;
    /**
     * @return Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
     * 
     */
    private final List<String> pemKeys;

    @CustomType.Constructor
    private GetAuthBackendConfigResult(
        @CustomType.Parameter("backend") @Nullable String backend,
        @CustomType.Parameter("disableIssValidation") Boolean disableIssValidation,
        @CustomType.Parameter("disableLocalCaJwt") Boolean disableLocalCaJwt,
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("issuer") String issuer,
        @CustomType.Parameter("kubernetesCaCert") String kubernetesCaCert,
        @CustomType.Parameter("kubernetesHost") String kubernetesHost,
        @CustomType.Parameter("pemKeys") List<String> pemKeys) {
        this.backend = backend;
        this.disableIssValidation = disableIssValidation;
        this.disableLocalCaJwt = disableLocalCaJwt;
        this.id = id;
        this.issuer = issuer;
        this.kubernetesCaCert = kubernetesCaCert;
        this.kubernetesHost = kubernetesHost;
        this.pemKeys = pemKeys;
    }

    public Optional<String> backend() {
        return Optional.ofNullable(this.backend);
    }
    public Boolean disableIssValidation() {
        return this.disableIssValidation;
    }
    public Boolean disableLocalCaJwt() {
        return this.disableLocalCaJwt;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
     * 
     */
    public String issuer() {
        return this.issuer;
    }
    /**
     * @return PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
     * 
     */
    public String kubernetesCaCert() {
        return this.kubernetesCaCert;
    }
    /**
     * @return Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
     * 
     */
    public String kubernetesHost() {
        return this.kubernetesHost;
    }
    /**
     * @return Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
     * 
     */
    public List<String> pemKeys() {
        return this.pemKeys;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAuthBackendConfigResult defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private @Nullable String backend;
        private Boolean disableIssValidation;
        private Boolean disableLocalCaJwt;
        private String id;
        private String issuer;
        private String kubernetesCaCert;
        private String kubernetesHost;
        private List<String> pemKeys;

        public Builder() {
    	      // Empty
        }

        public Builder(GetAuthBackendConfigResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.backend = defaults.backend;
    	      this.disableIssValidation = defaults.disableIssValidation;
    	      this.disableLocalCaJwt = defaults.disableLocalCaJwt;
    	      this.id = defaults.id;
    	      this.issuer = defaults.issuer;
    	      this.kubernetesCaCert = defaults.kubernetesCaCert;
    	      this.kubernetesHost = defaults.kubernetesHost;
    	      this.pemKeys = defaults.pemKeys;
        }

        public Builder backend(@Nullable String backend) {
            this.backend = backend;
            return this;
        }
        public Builder disableIssValidation(Boolean disableIssValidation) {
            this.disableIssValidation = Objects.requireNonNull(disableIssValidation);
            return this;
        }
        public Builder disableLocalCaJwt(Boolean disableLocalCaJwt) {
            this.disableLocalCaJwt = Objects.requireNonNull(disableLocalCaJwt);
            return this;
        }
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder issuer(String issuer) {
            this.issuer = Objects.requireNonNull(issuer);
            return this;
        }
        public Builder kubernetesCaCert(String kubernetesCaCert) {
            this.kubernetesCaCert = Objects.requireNonNull(kubernetesCaCert);
            return this;
        }
        public Builder kubernetesHost(String kubernetesHost) {
            this.kubernetesHost = Objects.requireNonNull(kubernetesHost);
            return this;
        }
        public Builder pemKeys(List<String> pemKeys) {
            this.pemKeys = Objects.requireNonNull(pemKeys);
            return this;
        }
        public Builder pemKeys(String... pemKeys) {
            return pemKeys(List.of(pemKeys));
        }        public GetAuthBackendConfigResult build() {
            return new GetAuthBackendConfigResult(backend, disableIssValidation, disableLocalCaJwt, id, issuer, kubernetesCaCert, kubernetesHost, pemKeys);
        }
    }
}