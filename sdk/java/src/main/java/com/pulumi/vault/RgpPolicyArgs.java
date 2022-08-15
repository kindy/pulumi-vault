// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RgpPolicyArgs extends com.pulumi.resources.ResourceArgs {

    public static final RgpPolicyArgs Empty = new RgpPolicyArgs();

    /**
     * Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
     * 
     */
    @Import(name="enforcementLevel", required=true)
    private Output<String> enforcementLevel;

    /**
     * @return Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
     * 
     */
    public Output<String> enforcementLevel() {
        return this.enforcementLevel;
    }

    /**
     * The name of the policy
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name of the policy
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * String containing a Sentinel policy
     * 
     */
    @Import(name="policy", required=true)
    private Output<String> policy;

    /**
     * @return String containing a Sentinel policy
     * 
     */
    public Output<String> policy() {
        return this.policy;
    }

    private RgpPolicyArgs() {}

    private RgpPolicyArgs(RgpPolicyArgs $) {
        this.enforcementLevel = $.enforcementLevel;
        this.name = $.name;
        this.policy = $.policy;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RgpPolicyArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RgpPolicyArgs $;

        public Builder() {
            $ = new RgpPolicyArgs();
        }

        public Builder(RgpPolicyArgs defaults) {
            $ = new RgpPolicyArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param enforcementLevel Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
         * 
         * @return builder
         * 
         */
        public Builder enforcementLevel(Output<String> enforcementLevel) {
            $.enforcementLevel = enforcementLevel;
            return this;
        }

        /**
         * @param enforcementLevel Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
         * 
         * @return builder
         * 
         */
        public Builder enforcementLevel(String enforcementLevel) {
            return enforcementLevel(Output.of(enforcementLevel));
        }

        /**
         * @param name The name of the policy
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name of the policy
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param policy String containing a Sentinel policy
         * 
         * @return builder
         * 
         */
        public Builder policy(Output<String> policy) {
            $.policy = policy;
            return this;
        }

        /**
         * @param policy String containing a Sentinel policy
         * 
         * @return builder
         * 
         */
        public Builder policy(String policy) {
            return policy(Output.of(policy));
        }

        public RgpPolicyArgs build() {
            $.enforcementLevel = Objects.requireNonNull($.enforcementLevel, "expected parameter 'enforcementLevel' to be non-null");
            $.policy = Objects.requireNonNull($.policy, "expected parameter 'policy' to be non-null");
            return $;
        }
    }

}
