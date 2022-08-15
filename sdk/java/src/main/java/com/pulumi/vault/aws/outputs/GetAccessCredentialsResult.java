// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.aws.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetAccessCredentialsResult {
    /**
     * @return The AWS Access Key ID returned by Vault.
     * 
     */
    private final String accessKey;
    private final String backend;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private final String id;
    /**
     * @return The duration of the secret lease, in seconds relative
     * to the time the data was requested. Once this time has passed any plan
     * generated with this data may fail to apply.
     * 
     */
    private final Integer leaseDuration;
    /**
     * @return The lease identifier assigned by Vault.
     * 
     */
    private final String leaseId;
    private final Boolean leaseRenewable;
    private final String leaseStartTime;
    private final @Nullable String region;
    private final String role;
    private final @Nullable String roleArn;
    /**
     * @return The AWS Secret Key returned by Vault.
     * 
     */
    private final String secretKey;
    /**
     * @return The STS token returned by Vault, if any.
     * 
     */
    private final String securityToken;
    private final @Nullable String ttl;
    private final @Nullable String type;

    @CustomType.Constructor
    private GetAccessCredentialsResult(
        @CustomType.Parameter("accessKey") String accessKey,
        @CustomType.Parameter("backend") String backend,
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("leaseDuration") Integer leaseDuration,
        @CustomType.Parameter("leaseId") String leaseId,
        @CustomType.Parameter("leaseRenewable") Boolean leaseRenewable,
        @CustomType.Parameter("leaseStartTime") String leaseStartTime,
        @CustomType.Parameter("region") @Nullable String region,
        @CustomType.Parameter("role") String role,
        @CustomType.Parameter("roleArn") @Nullable String roleArn,
        @CustomType.Parameter("secretKey") String secretKey,
        @CustomType.Parameter("securityToken") String securityToken,
        @CustomType.Parameter("ttl") @Nullable String ttl,
        @CustomType.Parameter("type") @Nullable String type) {
        this.accessKey = accessKey;
        this.backend = backend;
        this.id = id;
        this.leaseDuration = leaseDuration;
        this.leaseId = leaseId;
        this.leaseRenewable = leaseRenewable;
        this.leaseStartTime = leaseStartTime;
        this.region = region;
        this.role = role;
        this.roleArn = roleArn;
        this.secretKey = secretKey;
        this.securityToken = securityToken;
        this.ttl = ttl;
        this.type = type;
    }

    /**
     * @return The AWS Access Key ID returned by Vault.
     * 
     */
    public String accessKey() {
        return this.accessKey;
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
    /**
     * @return The duration of the secret lease, in seconds relative
     * to the time the data was requested. Once this time has passed any plan
     * generated with this data may fail to apply.
     * 
     */
    public Integer leaseDuration() {
        return this.leaseDuration;
    }
    /**
     * @return The lease identifier assigned by Vault.
     * 
     */
    public String leaseId() {
        return this.leaseId;
    }
    public Boolean leaseRenewable() {
        return this.leaseRenewable;
    }
    public String leaseStartTime() {
        return this.leaseStartTime;
    }
    public Optional<String> region() {
        return Optional.ofNullable(this.region);
    }
    public String role() {
        return this.role;
    }
    public Optional<String> roleArn() {
        return Optional.ofNullable(this.roleArn);
    }
    /**
     * @return The AWS Secret Key returned by Vault.
     * 
     */
    public String secretKey() {
        return this.secretKey;
    }
    /**
     * @return The STS token returned by Vault, if any.
     * 
     */
    public String securityToken() {
        return this.securityToken;
    }
    public Optional<String> ttl() {
        return Optional.ofNullable(this.ttl);
    }
    public Optional<String> type() {
        return Optional.ofNullable(this.type);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetAccessCredentialsResult defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String accessKey;
        private String backend;
        private String id;
        private Integer leaseDuration;
        private String leaseId;
        private Boolean leaseRenewable;
        private String leaseStartTime;
        private @Nullable String region;
        private String role;
        private @Nullable String roleArn;
        private String secretKey;
        private String securityToken;
        private @Nullable String ttl;
        private @Nullable String type;

        public Builder() {
    	      // Empty
        }

        public Builder(GetAccessCredentialsResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.accessKey = defaults.accessKey;
    	      this.backend = defaults.backend;
    	      this.id = defaults.id;
    	      this.leaseDuration = defaults.leaseDuration;
    	      this.leaseId = defaults.leaseId;
    	      this.leaseRenewable = defaults.leaseRenewable;
    	      this.leaseStartTime = defaults.leaseStartTime;
    	      this.region = defaults.region;
    	      this.role = defaults.role;
    	      this.roleArn = defaults.roleArn;
    	      this.secretKey = defaults.secretKey;
    	      this.securityToken = defaults.securityToken;
    	      this.ttl = defaults.ttl;
    	      this.type = defaults.type;
        }

        public Builder accessKey(String accessKey) {
            this.accessKey = Objects.requireNonNull(accessKey);
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
        public Builder leaseDuration(Integer leaseDuration) {
            this.leaseDuration = Objects.requireNonNull(leaseDuration);
            return this;
        }
        public Builder leaseId(String leaseId) {
            this.leaseId = Objects.requireNonNull(leaseId);
            return this;
        }
        public Builder leaseRenewable(Boolean leaseRenewable) {
            this.leaseRenewable = Objects.requireNonNull(leaseRenewable);
            return this;
        }
        public Builder leaseStartTime(String leaseStartTime) {
            this.leaseStartTime = Objects.requireNonNull(leaseStartTime);
            return this;
        }
        public Builder region(@Nullable String region) {
            this.region = region;
            return this;
        }
        public Builder role(String role) {
            this.role = Objects.requireNonNull(role);
            return this;
        }
        public Builder roleArn(@Nullable String roleArn) {
            this.roleArn = roleArn;
            return this;
        }
        public Builder secretKey(String secretKey) {
            this.secretKey = Objects.requireNonNull(secretKey);
            return this;
        }
        public Builder securityToken(String securityToken) {
            this.securityToken = Objects.requireNonNull(securityToken);
            return this;
        }
        public Builder ttl(@Nullable String ttl) {
            this.ttl = ttl;
            return this;
        }
        public Builder type(@Nullable String type) {
            this.type = type;
            return this;
        }        public GetAccessCredentialsResult build() {
            return new GetAccessCredentialsResult(accessKey, backend, id, leaseDuration, leaseId, leaseRenewable, leaseStartTime, region, role, roleArn, secretKey, securityToken, ttl, type);
        }
    }
}
