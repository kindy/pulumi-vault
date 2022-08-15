// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Double;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class QuotaRateLimitState extends com.pulumi.resources.ResourceArgs {

    public static final QuotaRateLimitState Empty = new QuotaRateLimitState();

    /**
     * Name of the rate limit quota
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the rate limit quota
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * Path of the mount or namespace to apply the quota. A blank path configures a
     * global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
     * `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
     * Updating this field on an existing quota can have &#34;moving&#34; effects. For example, updating
     * `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
     * a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return Path of the mount or namespace to apply the quota. A blank path configures a
     * global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
     * `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
     * Updating this field on an existing quota can have &#34;moving&#34; effects. For example, updating
     * `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
     * a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * The maximum number of requests at any given second to be allowed by the quota
     * rule. The `rate` must be positive.
     * 
     */
    @Import(name="rate")
    private @Nullable Output<Double> rate;

    /**
     * @return The maximum number of requests at any given second to be allowed by the quota
     * rule. The `rate` must be positive.
     * 
     */
    public Optional<Output<Double>> rate() {
        return Optional.ofNullable(this.rate);
    }

    private QuotaRateLimitState() {}

    private QuotaRateLimitState(QuotaRateLimitState $) {
        this.name = $.name;
        this.path = $.path;
        this.rate = $.rate;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(QuotaRateLimitState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private QuotaRateLimitState $;

        public Builder() {
            $ = new QuotaRateLimitState();
        }

        public Builder(QuotaRateLimitState defaults) {
            $ = new QuotaRateLimitState(Objects.requireNonNull(defaults));
        }

        /**
         * @param name Name of the rate limit quota
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the rate limit quota
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param path Path of the mount or namespace to apply the quota. A blank path configures a
         * global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
         * `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
         * Updating this field on an existing quota can have &#34;moving&#34; effects. For example, updating
         * `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
         * a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path Path of the mount or namespace to apply the quota. A blank path configures a
         * global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
         * `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
         * Updating this field on an existing quota can have &#34;moving&#34; effects. For example, updating
         * `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
         * a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param rate The maximum number of requests at any given second to be allowed by the quota
         * rule. The `rate` must be positive.
         * 
         * @return builder
         * 
         */
        public Builder rate(@Nullable Output<Double> rate) {
            $.rate = rate;
            return this;
        }

        /**
         * @param rate The maximum number of requests at any given second to be allowed by the quota
         * rule. The `rate` must be positive.
         * 
         * @return builder
         * 
         */
        public Builder rate(Double rate) {
            return rate(Output.of(rate));
        }

        public QuotaRateLimitState build() {
            return $;
        }
    }

}
