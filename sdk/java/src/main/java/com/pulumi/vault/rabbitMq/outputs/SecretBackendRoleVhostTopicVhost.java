// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.rabbitMq.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.String;
import java.util.Objects;

@CustomType
public final class SecretBackendRoleVhostTopicVhost {
    private final String read;
    private final String topic;
    private final String write;

    @CustomType.Constructor
    private SecretBackendRoleVhostTopicVhost(
        @CustomType.Parameter("read") String read,
        @CustomType.Parameter("topic") String topic,
        @CustomType.Parameter("write") String write) {
        this.read = read;
        this.topic = topic;
        this.write = write;
    }

    public String read() {
        return this.read;
    }
    public String topic() {
        return this.topic;
    }
    public String write() {
        return this.write;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SecretBackendRoleVhostTopicVhost defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private String read;
        private String topic;
        private String write;

        public Builder() {
    	      // Empty
        }

        public Builder(SecretBackendRoleVhostTopicVhost defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.read = defaults.read;
    	      this.topic = defaults.topic;
    	      this.write = defaults.write;
        }

        public Builder read(String read) {
            this.read = Objects.requireNonNull(read);
            return this;
        }
        public Builder topic(String topic) {
            this.topic = Objects.requireNonNull(topic);
            return this;
        }
        public Builder write(String write) {
            this.write = Objects.requireNonNull(write);
            return this;
        }        public SecretBackendRoleVhostTopicVhost build() {
            return new SecretBackendRoleVhostTopicVhost(read, topic, write);
        }
    }
}
