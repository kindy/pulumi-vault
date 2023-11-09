// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderAuthLoginAzureArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderAuthLoginAzureArgs Empty = new ProviderAuthLoginAzureArgs();

    @Import(name="clientId")
    private @Nullable Output<String> clientId;

    public Optional<Output<String>> clientId() {
        return Optional.ofNullable(this.clientId);
    }

    @Import(name="jwt")
    private @Nullable Output<String> jwt;

    public Optional<Output<String>> jwt() {
        return Optional.ofNullable(this.jwt);
    }

    @Import(name="mount")
    private @Nullable Output<String> mount;

    public Optional<Output<String>> mount() {
        return Optional.ofNullable(this.mount);
    }

    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    @Import(name="resourceGroupName", required=true)
    private Output<String> resourceGroupName;

    public Output<String> resourceGroupName() {
        return this.resourceGroupName;
    }

    @Import(name="role", required=true)
    private Output<String> role;

    public Output<String> role() {
        return this.role;
    }

    @Import(name="scope")
    private @Nullable Output<String> scope;

    public Optional<Output<String>> scope() {
        return Optional.ofNullable(this.scope);
    }

    @Import(name="subscriptionId", required=true)
    private Output<String> subscriptionId;

    public Output<String> subscriptionId() {
        return this.subscriptionId;
    }

    @Import(name="tenantId")
    private @Nullable Output<String> tenantId;

    public Optional<Output<String>> tenantId() {
        return Optional.ofNullable(this.tenantId);
    }

    @Import(name="useRootNamespace")
    private @Nullable Output<Boolean> useRootNamespace;

    public Optional<Output<Boolean>> useRootNamespace() {
        return Optional.ofNullable(this.useRootNamespace);
    }

    @Import(name="vmName")
    private @Nullable Output<String> vmName;

    public Optional<Output<String>> vmName() {
        return Optional.ofNullable(this.vmName);
    }

    @Import(name="vmssName")
    private @Nullable Output<String> vmssName;

    public Optional<Output<String>> vmssName() {
        return Optional.ofNullable(this.vmssName);
    }

    private ProviderAuthLoginAzureArgs() {}

    private ProviderAuthLoginAzureArgs(ProviderAuthLoginAzureArgs $) {
        this.clientId = $.clientId;
        this.jwt = $.jwt;
        this.mount = $.mount;
        this.namespace = $.namespace;
        this.resourceGroupName = $.resourceGroupName;
        this.role = $.role;
        this.scope = $.scope;
        this.subscriptionId = $.subscriptionId;
        this.tenantId = $.tenantId;
        this.useRootNamespace = $.useRootNamespace;
        this.vmName = $.vmName;
        this.vmssName = $.vmssName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderAuthLoginAzureArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderAuthLoginAzureArgs $;

        public Builder() {
            $ = new ProviderAuthLoginAzureArgs();
        }

        public Builder(ProviderAuthLoginAzureArgs defaults) {
            $ = new ProviderAuthLoginAzureArgs(Objects.requireNonNull(defaults));
        }

        public Builder clientId(@Nullable Output<String> clientId) {
            $.clientId = clientId;
            return this;
        }

        public Builder clientId(String clientId) {
            return clientId(Output.of(clientId));
        }

        public Builder jwt(@Nullable Output<String> jwt) {
            $.jwt = jwt;
            return this;
        }

        public Builder jwt(String jwt) {
            return jwt(Output.of(jwt));
        }

        public Builder mount(@Nullable Output<String> mount) {
            $.mount = mount;
            return this;
        }

        public Builder mount(String mount) {
            return mount(Output.of(mount));
        }

        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        public Builder resourceGroupName(Output<String> resourceGroupName) {
            $.resourceGroupName = resourceGroupName;
            return this;
        }

        public Builder resourceGroupName(String resourceGroupName) {
            return resourceGroupName(Output.of(resourceGroupName));
        }

        public Builder role(Output<String> role) {
            $.role = role;
            return this;
        }

        public Builder role(String role) {
            return role(Output.of(role));
        }

        public Builder scope(@Nullable Output<String> scope) {
            $.scope = scope;
            return this;
        }

        public Builder scope(String scope) {
            return scope(Output.of(scope));
        }

        public Builder subscriptionId(Output<String> subscriptionId) {
            $.subscriptionId = subscriptionId;
            return this;
        }

        public Builder subscriptionId(String subscriptionId) {
            return subscriptionId(Output.of(subscriptionId));
        }

        public Builder tenantId(@Nullable Output<String> tenantId) {
            $.tenantId = tenantId;
            return this;
        }

        public Builder tenantId(String tenantId) {
            return tenantId(Output.of(tenantId));
        }

        public Builder useRootNamespace(@Nullable Output<Boolean> useRootNamespace) {
            $.useRootNamespace = useRootNamespace;
            return this;
        }

        public Builder useRootNamespace(Boolean useRootNamespace) {
            return useRootNamespace(Output.of(useRootNamespace));
        }

        public Builder vmName(@Nullable Output<String> vmName) {
            $.vmName = vmName;
            return this;
        }

        public Builder vmName(String vmName) {
            return vmName(Output.of(vmName));
        }

        public Builder vmssName(@Nullable Output<String> vmssName) {
            $.vmssName = vmssName;
            return this;
        }

        public Builder vmssName(String vmssName) {
            return vmssName(Output.of(vmssName));
        }

        public ProviderAuthLoginAzureArgs build() {
            $.resourceGroupName = Objects.requireNonNull($.resourceGroupName, "expected parameter 'resourceGroupName' to be non-null");
            $.role = Objects.requireNonNull($.role, "expected parameter 'role' to be non-null");
            $.subscriptionId = Objects.requireNonNull($.subscriptionId, "expected parameter 'subscriptionId' to be non-null");
            return $;
        }
    }

}
