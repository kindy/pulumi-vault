// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;


public final class ProviderClientAuthArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderClientAuthArgs Empty = new ProviderClientAuthArgs();

    @Import(name="certFile", required=true)
    private Output<String> certFile;

    public Output<String> certFile() {
        return this.certFile;
    }

    @Import(name="keyFile", required=true)
    private Output<String> keyFile;

    public Output<String> keyFile() {
        return this.keyFile;
    }

    private ProviderClientAuthArgs() {}

    private ProviderClientAuthArgs(ProviderClientAuthArgs $) {
        this.certFile = $.certFile;
        this.keyFile = $.keyFile;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderClientAuthArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderClientAuthArgs $;

        public Builder() {
            $ = new ProviderClientAuthArgs();
        }

        public Builder(ProviderClientAuthArgs defaults) {
            $ = new ProviderClientAuthArgs(Objects.requireNonNull(defaults));
        }

        public Builder certFile(Output<String> certFile) {
            $.certFile = certFile;
            return this;
        }

        public Builder certFile(String certFile) {
            return certFile(Output.of(certFile));
        }

        public Builder keyFile(Output<String> keyFile) {
            $.keyFile = keyFile;
            return this;
        }

        public Builder keyFile(String keyFile) {
            return keyFile(Output.of(keyFile));
        }

        public ProviderClientAuthArgs build() {
            $.certFile = Objects.requireNonNull($.certFile, "expected parameter 'certFile' to be non-null");
            $.keyFile = Objects.requireNonNull($.keyFile, "expected parameter 'keyFile' to be non-null");
            return $;
        }
    }

}
