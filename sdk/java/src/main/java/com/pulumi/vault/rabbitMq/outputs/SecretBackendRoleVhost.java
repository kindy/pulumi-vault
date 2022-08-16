// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.rabbitMq.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class SecretBackendRoleVhost {
    private final String configure;
    private final String host;
    private final String read;
    private final String write;

    @CustomType.Constructor
    private SecretBackendRoleVhost(
        @CustomType.Parameter("configure") String configure,
        @CustomType.Parameter("host") String host,
        @CustomType.Parameter("read") String read,
        @CustomType.Parameter("write") String write) {
        this.configure = configure;
        this.host = host;
        this.read = read;
        this.write = write;
    }

    public String configure() {
        return this.configure;
    }
    public String host() {
        return this.host;
    }
    public String read() {
        return this.read;
    }
    public String write() {
        return this.write;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SecretBackendRoleVhost defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String configure;
        private String host;
        private String read;
        private String write;

        public Builder() {
    	      // Empty
        }

        public Builder(SecretBackendRoleVhost defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.configure = defaults.configure;
    	      this.host = defaults.host;
    	      this.read = defaults.read;
    	      this.write = defaults.write;
        }

        public Builder configure(String configure) {
            this.configure = Objects.requireNonNull(configure);
            return this;
        }
        public Builder host(String host) {
            this.host = Objects.requireNonNull(host);
            return this;
        }
        public Builder read(String read) {
            this.read = Objects.requireNonNull(read);
            return this;
        }
        public Builder write(String write) {
            this.write = Objects.requireNonNull(write);
            return this;
        }        public SecretBackendRoleVhost build() {
            return new SecretBackendRoleVhost(configure, host, read, write);
        }
    }
}