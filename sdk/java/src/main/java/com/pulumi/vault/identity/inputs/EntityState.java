// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EntityState extends com.pulumi.resources.ResourceArgs {

    public static final EntityState Empty = new EntityState();

    /**
     * True/false Is this entity currently disabled. Defaults to `false`
     * 
     */
    @Import(name="disabled")
    private @Nullable Output<Boolean> disabled;

    /**
     * @return True/false Is this entity currently disabled. Defaults to `false`
     * 
     */
    public Optional<Output<Boolean>> disabled() {
        return Optional.ofNullable(this.disabled);
    }

    /**
     * `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `vault.identity.EntityPolicies` to manage policies for this entity in a decoupled manner.
     * 
     */
    @Import(name="externalPolicies")
    private @Nullable Output<Boolean> externalPolicies;

    /**
     * @return `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `vault.identity.EntityPolicies` to manage policies for this entity in a decoupled manner.
     * 
     */
    public Optional<Output<Boolean>> externalPolicies() {
        return Optional.ofNullable(this.externalPolicies);
    }

    /**
     * A Map of additional metadata to associate with the user.
     * 
     */
    @Import(name="metadata")
    private @Nullable Output<Map<String,String>> metadata;

    /**
     * @return A Map of additional metadata to associate with the user.
     * 
     */
    public Optional<Output<Map<String,String>>> metadata() {
        return Optional.ofNullable(this.metadata);
    }

    /**
     * Name of the identity entity to create.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the identity entity to create.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * A list of policies to apply to the entity.
     * 
     */
    @Import(name="policies")
    private @Nullable Output<List<String>> policies;

    /**
     * @return A list of policies to apply to the entity.
     * 
     */
    public Optional<Output<List<String>>> policies() {
        return Optional.ofNullable(this.policies);
    }

    private EntityState() {}

    private EntityState(EntityState $) {
        this.disabled = $.disabled;
        this.externalPolicies = $.externalPolicies;
        this.metadata = $.metadata;
        this.name = $.name;
        this.policies = $.policies;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EntityState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EntityState $;

        public Builder() {
            $ = new EntityState();
        }

        public Builder(EntityState defaults) {
            $ = new EntityState(Objects.requireNonNull(defaults));
        }

        /**
         * @param disabled True/false Is this entity currently disabled. Defaults to `false`
         * 
         * @return builder
         * 
         */
        public Builder disabled(@Nullable Output<Boolean> disabled) {
            $.disabled = disabled;
            return this;
        }

        /**
         * @param disabled True/false Is this entity currently disabled. Defaults to `false`
         * 
         * @return builder
         * 
         */
        public Builder disabled(Boolean disabled) {
            return disabled(Output.of(disabled));
        }

        /**
         * @param externalPolicies `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `vault.identity.EntityPolicies` to manage policies for this entity in a decoupled manner.
         * 
         * @return builder
         * 
         */
        public Builder externalPolicies(@Nullable Output<Boolean> externalPolicies) {
            $.externalPolicies = externalPolicies;
            return this;
        }

        /**
         * @param externalPolicies `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `vault.identity.EntityPolicies` to manage policies for this entity in a decoupled manner.
         * 
         * @return builder
         * 
         */
        public Builder externalPolicies(Boolean externalPolicies) {
            return externalPolicies(Output.of(externalPolicies));
        }

        /**
         * @param metadata A Map of additional metadata to associate with the user.
         * 
         * @return builder
         * 
         */
        public Builder metadata(@Nullable Output<Map<String,String>> metadata) {
            $.metadata = metadata;
            return this;
        }

        /**
         * @param metadata A Map of additional metadata to associate with the user.
         * 
         * @return builder
         * 
         */
        public Builder metadata(Map<String,String> metadata) {
            return metadata(Output.of(metadata));
        }

        /**
         * @param name Name of the identity entity to create.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the identity entity to create.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param policies A list of policies to apply to the entity.
         * 
         * @return builder
         * 
         */
        public Builder policies(@Nullable Output<List<String>> policies) {
            $.policies = policies;
            return this;
        }

        /**
         * @param policies A list of policies to apply to the entity.
         * 
         * @return builder
         * 
         */
        public Builder policies(List<String> policies) {
            return policies(Output.of(policies));
        }

        /**
         * @param policies A list of policies to apply to the entity.
         * 
         * @return builder
         * 
         */
        public Builder policies(String... policies) {
            return policies(List.of(policies));
        }

        public EntityState build() {
            return $;
        }
    }

}
