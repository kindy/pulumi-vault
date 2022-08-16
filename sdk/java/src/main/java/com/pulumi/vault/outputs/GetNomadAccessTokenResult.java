// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class GetNomadAccessTokenResult {
    /**
     * @return The public identifier for a specific token. It can be used
     * to look up information about a token or to revoke a token.
     * 
     */
    private final String accessorId;
    private final String backend;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private final String id;
    private final String role;
    /**
     * @return The token to be used when making requests to Nomad and should be kept private.
     * 
     */
    private final String secretId;

    @CustomType.Constructor
    private GetNomadAccessTokenResult(
        @CustomType.Parameter("accessorId") String accessorId,
        @CustomType.Parameter("backend") String backend,
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("role") String role,
        @CustomType.Parameter("secretId") String secretId) {
        this.accessorId = accessorId;
        this.backend = backend;
        this.id = id;
        this.role = role;
        this.secretId = secretId;
    }

    /**
     * @return The public identifier for a specific token. It can be used
     * to look up information about a token or to revoke a token.
     * 
     */
    public String accessorId() {
        return this.accessorId;
    }
    public String backend() {
        return this.backend;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String role() {
        return this.role;
    }
    /**
     * @return The token to be used when making requests to Nomad and should be kept private.
     * 
     */
    public String secretId() {
        return this.secretId;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetNomadAccessTokenResult defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String accessorId;
        private String backend;
        private String id;
        private String role;
        private String secretId;

        public Builder() {
    	      // Empty
        }

        public Builder(GetNomadAccessTokenResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessorId = defaults.accessorId;
    	      this.backend = defaults.backend;
    	      this.id = defaults.id;
    	      this.role = defaults.role;
    	      this.secretId = defaults.secretId;
        }

        public Builder accessorId(String accessorId) {
            this.accessorId = Objects.requireNonNull(accessorId);
            return this;
        }
        public Builder backend(String backend) {
            this.backend = Objects.requireNonNull(backend);
            return this;
        }
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder role(String role) {
            this.role = Objects.requireNonNull(role);
            return this;
        }
        public Builder secretId(String secretId) {
            this.secretId = Objects.requireNonNull(secretId);
            return this;
        }        public GetNomadAccessTokenResult build() {
            return new GetNomadAccessTokenResult(accessorId, backend, id, role, secretId);
        }
    }
}