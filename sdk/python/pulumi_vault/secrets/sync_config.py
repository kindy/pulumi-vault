# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SyncConfigArgs', 'SyncConfig']

@pulumi.input_type
class SyncConfigArgs:
    def __init__(__self__, *,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 queue_capacity: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a SyncConfig resource.
        :param pulumi.Input[bool] disabled: Disables the syncing process between Vault and external destinations. Defaults to `false`.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               This resource can only be configured in the root namespace.
               *Available only for Vault Enterprise*.
        :param pulumi.Input[int] queue_capacity: Maximum number of pending sync operations allowed on the queue. Defaults to `1000000`.
        """
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if queue_capacity is not None:
            pulumi.set(__self__, "queue_capacity", queue_capacity)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Disables the syncing process between Vault and external destinations. Defaults to `false`.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        The namespace to provision the resource in.
        This resource can only be configured in the root namespace.
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="queueCapacity")
    def queue_capacity(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of pending sync operations allowed on the queue. Defaults to `1000000`.
        """
        return pulumi.get(self, "queue_capacity")

    @queue_capacity.setter
    def queue_capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "queue_capacity", value)


@pulumi.input_type
class _SyncConfigState:
    def __init__(__self__, *,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 queue_capacity: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering SyncConfig resources.
        :param pulumi.Input[bool] disabled: Disables the syncing process between Vault and external destinations. Defaults to `false`.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               This resource can only be configured in the root namespace.
               *Available only for Vault Enterprise*.
        :param pulumi.Input[int] queue_capacity: Maximum number of pending sync operations allowed on the queue. Defaults to `1000000`.
        """
        if disabled is not None:
            pulumi.set(__self__, "disabled", disabled)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if queue_capacity is not None:
            pulumi.set(__self__, "queue_capacity", queue_capacity)

    @property
    @pulumi.getter
    def disabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Disables the syncing process between Vault and external destinations. Defaults to `false`.
        """
        return pulumi.get(self, "disabled")

    @disabled.setter
    def disabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disabled", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        The namespace to provision the resource in.
        This resource can only be configured in the root namespace.
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="queueCapacity")
    def queue_capacity(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of pending sync operations allowed on the queue. Defaults to `1000000`.
        """
        return pulumi.get(self, "queue_capacity")

    @queue_capacity.setter
    def queue_capacity(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "queue_capacity", value)


class SyncConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 queue_capacity: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        Configures the secret sync global config.
        The config is global and can only be managed in the root namespace.

        > **Important** The config is global so the secrets.SyncConfig resource must not be defined
        multiple times for the same Vault server. If multiple definition exists, the last one applied will be
        effective.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        global_config = vault.secrets.SyncConfig("globalConfig",
            disabled=True,
            queue_capacity=500000)
        ```

        ## Import

        ```sh
         $ pulumi import vault:secrets/syncConfig:SyncConfig config global_config
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disabled: Disables the syncing process between Vault and external destinations. Defaults to `false`.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               This resource can only be configured in the root namespace.
               *Available only for Vault Enterprise*.
        :param pulumi.Input[int] queue_capacity: Maximum number of pending sync operations allowed on the queue. Defaults to `1000000`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SyncConfigArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configures the secret sync global config.
        The config is global and can only be managed in the root namespace.

        > **Important** The config is global so the secrets.SyncConfig resource must not be defined
        multiple times for the same Vault server. If multiple definition exists, the last one applied will be
        effective.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        global_config = vault.secrets.SyncConfig("globalConfig",
            disabled=True,
            queue_capacity=500000)
        ```

        ## Import

        ```sh
         $ pulumi import vault:secrets/syncConfig:SyncConfig config global_config
        ```

        :param str resource_name: The name of the resource.
        :param SyncConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SyncConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 queue_capacity: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SyncConfigArgs.__new__(SyncConfigArgs)

            __props__.__dict__["disabled"] = disabled
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["queue_capacity"] = queue_capacity
        super(SyncConfig, __self__).__init__(
            'vault:secrets/syncConfig:SyncConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            queue_capacity: Optional[pulumi.Input[int]] = None) -> 'SyncConfig':
        """
        Get an existing SyncConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disabled: Disables the syncing process between Vault and external destinations. Defaults to `false`.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               This resource can only be configured in the root namespace.
               *Available only for Vault Enterprise*.
        :param pulumi.Input[int] queue_capacity: Maximum number of pending sync operations allowed on the queue. Defaults to `1000000`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SyncConfigState.__new__(_SyncConfigState)

        __props__.__dict__["disabled"] = disabled
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["queue_capacity"] = queue_capacity
        return SyncConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Disables the syncing process between Vault and external destinations. Defaults to `false`.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[str]]:
        """
        The namespace to provision the resource in.
        This resource can only be configured in the root namespace.
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="queueCapacity")
    def queue_capacity(self) -> pulumi.Output[Optional[int]]:
        """
        Maximum number of pending sync operations allowed on the queue. Defaults to `1000000`.
        """
        return pulumi.get(self, "queue_capacity")

