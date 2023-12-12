// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.config.inputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class AuthLoginTokenFile {
    private String filename;
    private @Nullable String namespace;
    private @Nullable Boolean useRootNamespace;

    private AuthLoginTokenFile() {}
    public String filename() {
        return this.filename;
    }
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }
    public Optional<Boolean> useRootNamespace() {
        return Optional.ofNullable(this.useRootNamespace);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AuthLoginTokenFile defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String filename;
        private @Nullable String namespace;
        private @Nullable Boolean useRootNamespace;
        public Builder() {}
        public Builder(AuthLoginTokenFile defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.filename = defaults.filename;
    	      this.namespace = defaults.namespace;
    	      this.useRootNamespace = defaults.useRootNamespace;
        }

        @CustomType.Setter
        public Builder filename(String filename) {
            this.filename = Objects.requireNonNull(filename);
            return this;
        }
        @CustomType.Setter
        public Builder namespace(@Nullable String namespace) {
            this.namespace = namespace;
            return this;
        }
        @CustomType.Setter
        public Builder useRootNamespace(@Nullable Boolean useRootNamespace) {
            this.useRootNamespace = useRootNamespace;
            return this;
        }
        public AuthLoginTokenFile build() {
            final var _resultValue = new AuthLoginTokenFile();
            _resultValue.filename = filename;
            _resultValue.namespace = namespace;
            _resultValue.useRootNamespace = useRootNamespace;
            return _resultValue;
        }
    }
}
