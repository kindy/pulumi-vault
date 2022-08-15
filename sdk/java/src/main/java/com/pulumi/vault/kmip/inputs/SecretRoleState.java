// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.kmip.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretRoleState extends com.pulumi.resources.ResourceArgs {

    public static final SecretRoleState Empty = new SecretRoleState();

    /**
     * Grant permission to use the KMIP Activate operation.
     * 
     */
    @Import(name="operationActivate")
    private @Nullable Output<Boolean> operationActivate;

    /**
     * @return Grant permission to use the KMIP Activate operation.
     * 
     */
    public Optional<Output<Boolean>> operationActivate() {
        return Optional.ofNullable(this.operationActivate);
    }

    /**
     * Grant permission to use the KMIP Add Attribute operation.
     * 
     */
    @Import(name="operationAddAttribute")
    private @Nullable Output<Boolean> operationAddAttribute;

    /**
     * @return Grant permission to use the KMIP Add Attribute operation.
     * 
     */
    public Optional<Output<Boolean>> operationAddAttribute() {
        return Optional.ofNullable(this.operationAddAttribute);
    }

    /**
     * Grant all permissions to this role. May not be specified with any other `operation_*` params.
     * 
     */
    @Import(name="operationAll")
    private @Nullable Output<Boolean> operationAll;

    /**
     * @return Grant all permissions to this role. May not be specified with any other `operation_*` params.
     * 
     */
    public Optional<Output<Boolean>> operationAll() {
        return Optional.ofNullable(this.operationAll);
    }

    /**
     * Grant permission to use the KMIP Create operation.
     * 
     */
    @Import(name="operationCreate")
    private @Nullable Output<Boolean> operationCreate;

    /**
     * @return Grant permission to use the KMIP Create operation.
     * 
     */
    public Optional<Output<Boolean>> operationCreate() {
        return Optional.ofNullable(this.operationCreate);
    }

    /**
     * Grant permission to use the KMIP Destroy operation.
     * 
     */
    @Import(name="operationDestroy")
    private @Nullable Output<Boolean> operationDestroy;

    /**
     * @return Grant permission to use the KMIP Destroy operation.
     * 
     */
    public Optional<Output<Boolean>> operationDestroy() {
        return Optional.ofNullable(this.operationDestroy);
    }

    /**
     * Grant permission to use the KMIP Discover Version operation.
     * 
     */
    @Import(name="operationDiscoverVersions")
    private @Nullable Output<Boolean> operationDiscoverVersions;

    /**
     * @return Grant permission to use the KMIP Discover Version operation.
     * 
     */
    public Optional<Output<Boolean>> operationDiscoverVersions() {
        return Optional.ofNullable(this.operationDiscoverVersions);
    }

    /**
     * Grant permission to use the KMIP Get operation.
     * 
     */
    @Import(name="operationGet")
    private @Nullable Output<Boolean> operationGet;

    /**
     * @return Grant permission to use the KMIP Get operation.
     * 
     */
    public Optional<Output<Boolean>> operationGet() {
        return Optional.ofNullable(this.operationGet);
    }

    /**
     * Grant permission to use the KMIP Get Atrribute List operation.
     * 
     */
    @Import(name="operationGetAttributeList")
    private @Nullable Output<Boolean> operationGetAttributeList;

    /**
     * @return Grant permission to use the KMIP Get Atrribute List operation.
     * 
     */
    public Optional<Output<Boolean>> operationGetAttributeList() {
        return Optional.ofNullable(this.operationGetAttributeList);
    }

    /**
     * Grant permission to use the KMIP Get Atrributes operation.
     * 
     */
    @Import(name="operationGetAttributes")
    private @Nullable Output<Boolean> operationGetAttributes;

    /**
     * @return Grant permission to use the KMIP Get Atrributes operation.
     * 
     */
    public Optional<Output<Boolean>> operationGetAttributes() {
        return Optional.ofNullable(this.operationGetAttributes);
    }

    /**
     * Grant permission to use the KMIP Get Locate operation.
     * 
     */
    @Import(name="operationLocate")
    private @Nullable Output<Boolean> operationLocate;

    /**
     * @return Grant permission to use the KMIP Get Locate operation.
     * 
     */
    public Optional<Output<Boolean>> operationLocate() {
        return Optional.ofNullable(this.operationLocate);
    }

    /**
     * Remove all permissions from this role. May not be specified with any other `operation_*` params.
     * 
     */
    @Import(name="operationNone")
    private @Nullable Output<Boolean> operationNone;

    /**
     * @return Remove all permissions from this role. May not be specified with any other `operation_*` params.
     * 
     */
    public Optional<Output<Boolean>> operationNone() {
        return Optional.ofNullable(this.operationNone);
    }

    /**
     * Grant permission to use the KMIP Register operation.
     * 
     */
    @Import(name="operationRegister")
    private @Nullable Output<Boolean> operationRegister;

    /**
     * @return Grant permission to use the KMIP Register operation.
     * 
     */
    public Optional<Output<Boolean>> operationRegister() {
        return Optional.ofNullable(this.operationRegister);
    }

    /**
     * Grant permission to use the KMIP Rekey operation.
     * 
     */
    @Import(name="operationRekey")
    private @Nullable Output<Boolean> operationRekey;

    /**
     * @return Grant permission to use the KMIP Rekey operation.
     * 
     */
    public Optional<Output<Boolean>> operationRekey() {
        return Optional.ofNullable(this.operationRekey);
    }

    /**
     * Grant permission to use the KMIP Revoke operation.
     * 
     */
    @Import(name="operationRevoke")
    private @Nullable Output<Boolean> operationRevoke;

    /**
     * @return Grant permission to use the KMIP Revoke operation.
     * 
     */
    public Optional<Output<Boolean>> operationRevoke() {
        return Optional.ofNullable(this.operationRevoke);
    }

    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `kmip`.
     * 
     */
    @Import(name="path")
    private @Nullable Output<String> path;

    /**
     * @return The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `kmip`.
     * 
     */
    public Optional<Output<String>> path() {
        return Optional.ofNullable(this.path);
    }

    /**
     * Name of the role.
     * 
     */
    @Import(name="role")
    private @Nullable Output<String> role;

    /**
     * @return Name of the role.
     * 
     */
    public Optional<Output<String>> role() {
        return Optional.ofNullable(this.role);
    }

    /**
     * Name of the scope.
     * 
     */
    @Import(name="scope")
    private @Nullable Output<String> scope;

    /**
     * @return Name of the scope.
     * 
     */
    public Optional<Output<String>> scope() {
        return Optional.ofNullable(this.scope);
    }

    /**
     * Client certificate key bits, valid values depend on key type.
     * 
     */
    @Import(name="tlsClientKeyBits")
    private @Nullable Output<Integer> tlsClientKeyBits;

    /**
     * @return Client certificate key bits, valid values depend on key type.
     * 
     */
    public Optional<Output<Integer>> tlsClientKeyBits() {
        return Optional.ofNullable(this.tlsClientKeyBits);
    }

    /**
     * Client certificate key type, `rsa` or `ec`.
     * 
     */
    @Import(name="tlsClientKeyType")
    private @Nullable Output<String> tlsClientKeyType;

    /**
     * @return Client certificate key type, `rsa` or `ec`.
     * 
     */
    public Optional<Output<String>> tlsClientKeyType() {
        return Optional.ofNullable(this.tlsClientKeyType);
    }

    /**
     * Client certificate TTL in seconds.
     * 
     */
    @Import(name="tlsClientTtl")
    private @Nullable Output<Integer> tlsClientTtl;

    /**
     * @return Client certificate TTL in seconds.
     * 
     */
    public Optional<Output<Integer>> tlsClientTtl() {
        return Optional.ofNullable(this.tlsClientTtl);
    }

    private SecretRoleState() {}

    private SecretRoleState(SecretRoleState $) {
        this.operationActivate = $.operationActivate;
        this.operationAddAttribute = $.operationAddAttribute;
        this.operationAll = $.operationAll;
        this.operationCreate = $.operationCreate;
        this.operationDestroy = $.operationDestroy;
        this.operationDiscoverVersions = $.operationDiscoverVersions;
        this.operationGet = $.operationGet;
        this.operationGetAttributeList = $.operationGetAttributeList;
        this.operationGetAttributes = $.operationGetAttributes;
        this.operationLocate = $.operationLocate;
        this.operationNone = $.operationNone;
        this.operationRegister = $.operationRegister;
        this.operationRekey = $.operationRekey;
        this.operationRevoke = $.operationRevoke;
        this.path = $.path;
        this.role = $.role;
        this.scope = $.scope;
        this.tlsClientKeyBits = $.tlsClientKeyBits;
        this.tlsClientKeyType = $.tlsClientKeyType;
        this.tlsClientTtl = $.tlsClientTtl;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretRoleState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretRoleState $;

        public Builder() {
            $ = new SecretRoleState();
        }

        public Builder(SecretRoleState defaults) {
            $ = new SecretRoleState(Objects.requireNonNull(defaults));
        }

        /**
         * @param operationActivate Grant permission to use the KMIP Activate operation.
         * 
         * @return builder
         * 
         */
        public Builder operationActivate(@Nullable Output<Boolean> operationActivate) {
            $.operationActivate = operationActivate;
            return this;
        }

        /**
         * @param operationActivate Grant permission to use the KMIP Activate operation.
         * 
         * @return builder
         * 
         */
        public Builder operationActivate(Boolean operationActivate) {
            return operationActivate(Output.of(operationActivate));
        }

        /**
         * @param operationAddAttribute Grant permission to use the KMIP Add Attribute operation.
         * 
         * @return builder
         * 
         */
        public Builder operationAddAttribute(@Nullable Output<Boolean> operationAddAttribute) {
            $.operationAddAttribute = operationAddAttribute;
            return this;
        }

        /**
         * @param operationAddAttribute Grant permission to use the KMIP Add Attribute operation.
         * 
         * @return builder
         * 
         */
        public Builder operationAddAttribute(Boolean operationAddAttribute) {
            return operationAddAttribute(Output.of(operationAddAttribute));
        }

        /**
         * @param operationAll Grant all permissions to this role. May not be specified with any other `operation_*` params.
         * 
         * @return builder
         * 
         */
        public Builder operationAll(@Nullable Output<Boolean> operationAll) {
            $.operationAll = operationAll;
            return this;
        }

        /**
         * @param operationAll Grant all permissions to this role. May not be specified with any other `operation_*` params.
         * 
         * @return builder
         * 
         */
        public Builder operationAll(Boolean operationAll) {
            return operationAll(Output.of(operationAll));
        }

        /**
         * @param operationCreate Grant permission to use the KMIP Create operation.
         * 
         * @return builder
         * 
         */
        public Builder operationCreate(@Nullable Output<Boolean> operationCreate) {
            $.operationCreate = operationCreate;
            return this;
        }

        /**
         * @param operationCreate Grant permission to use the KMIP Create operation.
         * 
         * @return builder
         * 
         */
        public Builder operationCreate(Boolean operationCreate) {
            return operationCreate(Output.of(operationCreate));
        }

        /**
         * @param operationDestroy Grant permission to use the KMIP Destroy operation.
         * 
         * @return builder
         * 
         */
        public Builder operationDestroy(@Nullable Output<Boolean> operationDestroy) {
            $.operationDestroy = operationDestroy;
            return this;
        }

        /**
         * @param operationDestroy Grant permission to use the KMIP Destroy operation.
         * 
         * @return builder
         * 
         */
        public Builder operationDestroy(Boolean operationDestroy) {
            return operationDestroy(Output.of(operationDestroy));
        }

        /**
         * @param operationDiscoverVersions Grant permission to use the KMIP Discover Version operation.
         * 
         * @return builder
         * 
         */
        public Builder operationDiscoverVersions(@Nullable Output<Boolean> operationDiscoverVersions) {
            $.operationDiscoverVersions = operationDiscoverVersions;
            return this;
        }

        /**
         * @param operationDiscoverVersions Grant permission to use the KMIP Discover Version operation.
         * 
         * @return builder
         * 
         */
        public Builder operationDiscoverVersions(Boolean operationDiscoverVersions) {
            return operationDiscoverVersions(Output.of(operationDiscoverVersions));
        }

        /**
         * @param operationGet Grant permission to use the KMIP Get operation.
         * 
         * @return builder
         * 
         */
        public Builder operationGet(@Nullable Output<Boolean> operationGet) {
            $.operationGet = operationGet;
            return this;
        }

        /**
         * @param operationGet Grant permission to use the KMIP Get operation.
         * 
         * @return builder
         * 
         */
        public Builder operationGet(Boolean operationGet) {
            return operationGet(Output.of(operationGet));
        }

        /**
         * @param operationGetAttributeList Grant permission to use the KMIP Get Atrribute List operation.
         * 
         * @return builder
         * 
         */
        public Builder operationGetAttributeList(@Nullable Output<Boolean> operationGetAttributeList) {
            $.operationGetAttributeList = operationGetAttributeList;
            return this;
        }

        /**
         * @param operationGetAttributeList Grant permission to use the KMIP Get Atrribute List operation.
         * 
         * @return builder
         * 
         */
        public Builder operationGetAttributeList(Boolean operationGetAttributeList) {
            return operationGetAttributeList(Output.of(operationGetAttributeList));
        }

        /**
         * @param operationGetAttributes Grant permission to use the KMIP Get Atrributes operation.
         * 
         * @return builder
         * 
         */
        public Builder operationGetAttributes(@Nullable Output<Boolean> operationGetAttributes) {
            $.operationGetAttributes = operationGetAttributes;
            return this;
        }

        /**
         * @param operationGetAttributes Grant permission to use the KMIP Get Atrributes operation.
         * 
         * @return builder
         * 
         */
        public Builder operationGetAttributes(Boolean operationGetAttributes) {
            return operationGetAttributes(Output.of(operationGetAttributes));
        }

        /**
         * @param operationLocate Grant permission to use the KMIP Get Locate operation.
         * 
         * @return builder
         * 
         */
        public Builder operationLocate(@Nullable Output<Boolean> operationLocate) {
            $.operationLocate = operationLocate;
            return this;
        }

        /**
         * @param operationLocate Grant permission to use the KMIP Get Locate operation.
         * 
         * @return builder
         * 
         */
        public Builder operationLocate(Boolean operationLocate) {
            return operationLocate(Output.of(operationLocate));
        }

        /**
         * @param operationNone Remove all permissions from this role. May not be specified with any other `operation_*` params.
         * 
         * @return builder
         * 
         */
        public Builder operationNone(@Nullable Output<Boolean> operationNone) {
            $.operationNone = operationNone;
            return this;
        }

        /**
         * @param operationNone Remove all permissions from this role. May not be specified with any other `operation_*` params.
         * 
         * @return builder
         * 
         */
        public Builder operationNone(Boolean operationNone) {
            return operationNone(Output.of(operationNone));
        }

        /**
         * @param operationRegister Grant permission to use the KMIP Register operation.
         * 
         * @return builder
         * 
         */
        public Builder operationRegister(@Nullable Output<Boolean> operationRegister) {
            $.operationRegister = operationRegister;
            return this;
        }

        /**
         * @param operationRegister Grant permission to use the KMIP Register operation.
         * 
         * @return builder
         * 
         */
        public Builder operationRegister(Boolean operationRegister) {
            return operationRegister(Output.of(operationRegister));
        }

        /**
         * @param operationRekey Grant permission to use the KMIP Rekey operation.
         * 
         * @return builder
         * 
         */
        public Builder operationRekey(@Nullable Output<Boolean> operationRekey) {
            $.operationRekey = operationRekey;
            return this;
        }

        /**
         * @param operationRekey Grant permission to use the KMIP Rekey operation.
         * 
         * @return builder
         * 
         */
        public Builder operationRekey(Boolean operationRekey) {
            return operationRekey(Output.of(operationRekey));
        }

        /**
         * @param operationRevoke Grant permission to use the KMIP Revoke operation.
         * 
         * @return builder
         * 
         */
        public Builder operationRevoke(@Nullable Output<Boolean> operationRevoke) {
            $.operationRevoke = operationRevoke;
            return this;
        }

        /**
         * @param operationRevoke Grant permission to use the KMIP Revoke operation.
         * 
         * @return builder
         * 
         */
        public Builder operationRevoke(Boolean operationRevoke) {
            return operationRevoke(Output.of(operationRevoke));
        }

        /**
         * @param path The unique path this backend should be mounted at. Must
         * not begin or end with a `/`. Defaults to `kmip`.
         * 
         * @return builder
         * 
         */
        public Builder path(@Nullable Output<String> path) {
            $.path = path;
            return this;
        }

        /**
         * @param path The unique path this backend should be mounted at. Must
         * not begin or end with a `/`. Defaults to `kmip`.
         * 
         * @return builder
         * 
         */
        public Builder path(String path) {
            return path(Output.of(path));
        }

        /**
         * @param role Name of the role.
         * 
         * @return builder
         * 
         */
        public Builder role(@Nullable Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role Name of the role.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        /**
         * @param scope Name of the scope.
         * 
         * @return builder
         * 
         */
        public Builder scope(@Nullable Output<String> scope) {
            $.scope = scope;
            return this;
        }

        /**
         * @param scope Name of the scope.
         * 
         * @return builder
         * 
         */
        public Builder scope(String scope) {
            return scope(Output.of(scope));
        }

        /**
         * @param tlsClientKeyBits Client certificate key bits, valid values depend on key type.
         * 
         * @return builder
         * 
         */
        public Builder tlsClientKeyBits(@Nullable Output<Integer> tlsClientKeyBits) {
            $.tlsClientKeyBits = tlsClientKeyBits;
            return this;
        }

        /**
         * @param tlsClientKeyBits Client certificate key bits, valid values depend on key type.
         * 
         * @return builder
         * 
         */
        public Builder tlsClientKeyBits(Integer tlsClientKeyBits) {
            return tlsClientKeyBits(Output.of(tlsClientKeyBits));
        }

        /**
         * @param tlsClientKeyType Client certificate key type, `rsa` or `ec`.
         * 
         * @return builder
         * 
         */
        public Builder tlsClientKeyType(@Nullable Output<String> tlsClientKeyType) {
            $.tlsClientKeyType = tlsClientKeyType;
            return this;
        }

        /**
         * @param tlsClientKeyType Client certificate key type, `rsa` or `ec`.
         * 
         * @return builder
         * 
         */
        public Builder tlsClientKeyType(String tlsClientKeyType) {
            return tlsClientKeyType(Output.of(tlsClientKeyType));
        }

        /**
         * @param tlsClientTtl Client certificate TTL in seconds.
         * 
         * @return builder
         * 
         */
        public Builder tlsClientTtl(@Nullable Output<Integer> tlsClientTtl) {
            $.tlsClientTtl = tlsClientTtl;
            return this;
        }

        /**
         * @param tlsClientTtl Client certificate TTL in seconds.
         * 
         * @return builder
         * 
         */
        public Builder tlsClientTtl(Integer tlsClientTtl) {
            return tlsClientTtl(Output.of(tlsClientTtl));
        }

        public SecretRoleState build() {
            return $;
        }
    }

}
