// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class EntityAliasState extends com.pulumi.resources.ResourceArgs {

    public static final EntityAliasState Empty = new EntityAliasState();

    /**
     * Entity ID to which this alias belongs to.
     * 
     */
    @Import(name="canonicalId")
    private @Nullable Output<String> canonicalId;

    /**
     * @return Entity ID to which this alias belongs to.
     * 
     */
    public Optional<Output<String>> canonicalId() {
        return Optional.ofNullable(this.canonicalId);
    }

    /**
     * Custom metadata to be associated with this alias.
     * 
     */
    @Import(name="customMetadata")
    private @Nullable Output<Map<String,String>> customMetadata;

    /**
     * @return Custom metadata to be associated with this alias.
     * 
     */
    public Optional<Output<Map<String,String>>> customMetadata() {
        return Optional.ofNullable(this.customMetadata);
    }

    /**
     * Accessor of the mount to which the alias should belong to.
     * 
     */
    @Import(name="mountAccessor")
    private @Nullable Output<String> mountAccessor;

    /**
     * @return Accessor of the mount to which the alias should belong to.
     * 
     */
    public Optional<Output<String>> mountAccessor() {
        return Optional.ofNullable(this.mountAccessor);
    }

    /**
     * Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    private EntityAliasState() {}

    private EntityAliasState(EntityAliasState $) {
        this.canonicalId = $.canonicalId;
        this.customMetadata = $.customMetadata;
        this.mountAccessor = $.mountAccessor;
        this.name = $.name;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(EntityAliasState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private EntityAliasState $;

        public Builder() {
            $ = new EntityAliasState();
        }

        public Builder(EntityAliasState defaults) {
            $ = new EntityAliasState(Objects.requireNonNull(defaults));
        }

        /**
         * @param canonicalId Entity ID to which this alias belongs to.
         * 
         * @return builder
         * 
         */
        public Builder canonicalId(@Nullable Output<String> canonicalId) {
            $.canonicalId = canonicalId;
            return this;
        }

        /**
         * @param canonicalId Entity ID to which this alias belongs to.
         * 
         * @return builder
         * 
         */
        public Builder canonicalId(String canonicalId) {
            return canonicalId(Output.of(canonicalId));
        }

        /**
         * @param customMetadata Custom metadata to be associated with this alias.
         * 
         * @return builder
         * 
         */
        public Builder customMetadata(@Nullable Output<Map<String,String>> customMetadata) {
            $.customMetadata = customMetadata;
            return this;
        }

        /**
         * @param customMetadata Custom metadata to be associated with this alias.
         * 
         * @return builder
         * 
         */
        public Builder customMetadata(Map<String,String> customMetadata) {
            return customMetadata(Output.of(customMetadata));
        }

        /**
         * @param mountAccessor Accessor of the mount to which the alias should belong to.
         * 
         * @return builder
         * 
         */
        public Builder mountAccessor(@Nullable Output<String> mountAccessor) {
            $.mountAccessor = mountAccessor;
            return this;
        }

        /**
         * @param mountAccessor Accessor of the mount to which the alias should belong to.
         * 
         * @return builder
         * 
         */
        public Builder mountAccessor(String mountAccessor) {
            return mountAccessor(Output.of(mountAccessor));
        }

        /**
         * @param name Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Name of the alias. Name should be the identifier of the client in the authentication source. For example, if the alias belongs to userpass backend, the name should be a valid username within userpass backend. If alias belongs to GitHub, it should be the GitHub username.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        public EntityAliasState build() {
            return $;
        }
    }

}
