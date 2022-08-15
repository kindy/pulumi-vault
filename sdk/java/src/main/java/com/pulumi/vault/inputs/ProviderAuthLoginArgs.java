// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderAuthLoginArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderAuthLoginArgs Empty = new ProviderAuthLoginArgs();

    @Import(name="method")
    private @Nullable Output<String> method;

    public Optional<Output<String>> method() {
        return Optional.ofNullable(this.method);
    }

    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    @Import(name="parameters")
    private @Nullable Output<Map<String,String>> parameters;

    public Optional<Output<Map<String,String>>> parameters() {
        return Optional.ofNullable(this.parameters);
    }

    @Import(name="path", required=true)
    private Output<String> path;

    public Output<String> path() {
        return this.path;
    }

    private ProviderAuthLoginArgs() {}

    private ProviderAuthLoginArgs(ProviderAuthLoginArgs $) {
        this.method = $.method;
        this.namespace = $.namespace;
        this.parameters = $.parameters;
        this.path = $.path;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderAuthLoginArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderAuthLoginArgs $;

        public Builder() {
            $ = new ProviderAuthLoginArgs();
        }

        public Builder(ProviderAuthLoginArgs defaults) {
            $ = new ProviderAuthLoginArgs(Objects.requireNonNull(defaults));
        }

        public Builder method(@Nullable Output<String> method) {
            $.method = method;
            return this;
        }

        public Builder method(String method) {
            return method(Output.of(method));
        }

        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        public Builder parameters(@Nullable Output<Map<String,String>> parameters) {
            $.parameters = parameters;
            return this;
        }

        public Builder parameters(Map<String,String> parameters) {
            return parameters(Output.of(parameters));
        }

        public Builder path(Output<String> path) {
            $.path = path;
            return this;
        }

        public Builder path(String path) {
            return path(Output.of(path));
        }

        public ProviderAuthLoginArgs build() {
            $.path = Objects.requireNonNull($.path, "expected parameter 'path' to be non-null");
            return $;
        }
    }

}
