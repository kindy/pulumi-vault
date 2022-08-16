// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.transit.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetEncryptResult {
    private final String backend;
    /**
     * @return Encrypted ciphertext returned from Vault
     * 
     */
    private final String ciphertext;
    private final @Nullable String context;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private final String id;
    private final String key;
    private final @Nullable Integer keyVersion;
    private final String plaintext;

    @CustomType.Constructor
    private GetEncryptResult(
        @CustomType.Parameter("backend") String backend,
        @CustomType.Parameter("ciphertext") String ciphertext,
        @CustomType.Parameter("context") @Nullable String context,
        @CustomType.Parameter("id") String id,
        @CustomType.Parameter("key") String key,
        @CustomType.Parameter("keyVersion") @Nullable Integer keyVersion,
        @CustomType.Parameter("plaintext") String plaintext) {
        this.backend = backend;
        this.ciphertext = ciphertext;
        this.context = context;
        this.id = id;
        this.key = key;
        this.keyVersion = keyVersion;
        this.plaintext = plaintext;
    }

    public String backend() {
        return this.backend;
    }
    /**
     * @return Encrypted ciphertext returned from Vault
     * 
     */
    public String ciphertext() {
        return this.ciphertext;
    }
    public Optional<String> context() {
        return Optional.ofNullable(this.context);
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    public String key() {
        return this.key;
    }
    public Optional<Integer> keyVersion() {
        return Optional.ofNullable(this.keyVersion);
    }
    public String plaintext() {
        return this.plaintext;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetEncryptResult defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String backend;
        private String ciphertext;
        private @Nullable String context;
        private String id;
        private String key;
        private @Nullable Integer keyVersion;
        private String plaintext;

        public Builder() {
    	      // Empty
        }

        public Builder(GetEncryptResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.backend = defaults.backend;
    	      this.ciphertext = defaults.ciphertext;
    	      this.context = defaults.context;
    	      this.id = defaults.id;
    	      this.key = defaults.key;
    	      this.keyVersion = defaults.keyVersion;
    	      this.plaintext = defaults.plaintext;
        }

        public Builder backend(String backend) {
            this.backend = Objects.requireNonNull(backend);
            return this;
        }
        public Builder ciphertext(String ciphertext) {
            this.ciphertext = Objects.requireNonNull(ciphertext);
            return this;
        }
        public Builder context(@Nullable String context) {
            this.context = context;
            return this;
        }
        public Builder id(String id) {
            this.id = Objects.requireNonNull(id);
            return this;
        }
        public Builder key(String key) {
            this.key = Objects.requireNonNull(key);
            return this;
        }
        public Builder keyVersion(@Nullable Integer keyVersion) {
            this.keyVersion = keyVersion;
            return this;
        }
        public Builder plaintext(String plaintext) {
            this.plaintext = Objects.requireNonNull(plaintext);
            return this;
        }        public GetEncryptResult build() {
            return new GetEncryptResult(backend, ciphertext, context, id, key, keyVersion, plaintext);
        }
    }
}