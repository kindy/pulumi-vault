# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SecretBackendArgs', 'SecretBackend']

@pulumi.input_type
class SecretBackendArgs:
    def __init__(__self__, *,
                 mount: pulumi.Input[str],
                 private_key: pulumi.Input[str],
                 public_key: pulumi.Input[str],
                 namespace: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SecretBackend resource.
        :param pulumi.Input[str] mount: Path where the MongoDB Atlas Secrets Engine is mounted.
        :param pulumi.Input[str] private_key: Specifies the Private API Key used to authenticate with the MongoDB Atlas API.
        :param pulumi.Input[str] public_key: Specifies the Public API Key used to authenticate with the MongoDB Atlas API.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        """
        pulumi.set(__self__, "mount", mount)
        pulumi.set(__self__, "private_key", private_key)
        pulumi.set(__self__, "public_key", public_key)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def mount(self) -> pulumi.Input[str]:
        """
        Path where the MongoDB Atlas Secrets Engine is mounted.
        """
        return pulumi.get(self, "mount")

    @mount.setter
    def mount(self, value: pulumi.Input[str]):
        pulumi.set(self, "mount", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Input[str]:
        """
        Specifies the Private API Key used to authenticate with the MongoDB Atlas API.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Input[str]:
        """
        Specifies the Public API Key used to authenticate with the MongoDB Atlas API.
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "public_key", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)


@pulumi.input_type
class _SecretBackendState:
    def __init__(__self__, *,
                 mount: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SecretBackend resources.
        :param pulumi.Input[str] mount: Path where the MongoDB Atlas Secrets Engine is mounted.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: Path where MongoDB Atlas configuration is located
        :param pulumi.Input[str] private_key: Specifies the Private API Key used to authenticate with the MongoDB Atlas API.
        :param pulumi.Input[str] public_key: Specifies the Public API Key used to authenticate with the MongoDB Atlas API.
        """
        if mount is not None:
            pulumi.set(__self__, "mount", mount)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if private_key is not None:
            pulumi.set(__self__, "private_key", private_key)
        if public_key is not None:
            pulumi.set(__self__, "public_key", public_key)

    @property
    @pulumi.getter
    def mount(self) -> Optional[pulumi.Input[str]]:
        """
        Path where the MongoDB Atlas Secrets Engine is mounted.
        """
        return pulumi.get(self, "mount")

    @mount.setter
    def mount(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mount", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        Path where MongoDB Atlas configuration is located
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the Private API Key used to authenticate with the MongoDB Atlas API.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the Public API Key used to authenticate with the MongoDB Atlas API.
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key", value)


class SecretBackend(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 mount: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        mongo = vault.Mount("mongo",
            description="MongoDB Atlas secret engine mount",
            path="mongodbatlas",
            type="mongodbatlas")
        config = vault.mongodbatlas.SecretBackend("config",
            mount="vault_mount.mongo.path",
            private_key="privateKey",
            public_key="publicKey")
        ```

        ## Import

        MongoDB Atlas secret backends can be imported using the `${mount}/config`, e.g.

        ```sh
         $ pulumi import vault:mongodbatlas/secretBackend:SecretBackend config mongodbatlas/config
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] mount: Path where the MongoDB Atlas Secrets Engine is mounted.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] private_key: Specifies the Private API Key used to authenticate with the MongoDB Atlas API.
        :param pulumi.Input[str] public_key: Specifies the Public API Key used to authenticate with the MongoDB Atlas API.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretBackendArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        mongo = vault.Mount("mongo",
            description="MongoDB Atlas secret engine mount",
            path="mongodbatlas",
            type="mongodbatlas")
        config = vault.mongodbatlas.SecretBackend("config",
            mount="vault_mount.mongo.path",
            private_key="privateKey",
            public_key="publicKey")
        ```

        ## Import

        MongoDB Atlas secret backends can be imported using the `${mount}/config`, e.g.

        ```sh
         $ pulumi import vault:mongodbatlas/secretBackend:SecretBackend config mongodbatlas/config
        ```

        :param str resource_name: The name of the resource.
        :param SecretBackendArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretBackendArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 mount: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretBackendArgs.__new__(SecretBackendArgs)

            if mount is None and not opts.urn:
                raise TypeError("Missing required property 'mount'")
            __props__.__dict__["mount"] = mount
            __props__.__dict__["namespace"] = namespace
            if private_key is None and not opts.urn:
                raise TypeError("Missing required property 'private_key'")
            __props__.__dict__["private_key"] = private_key
            if public_key is None and not opts.urn:
                raise TypeError("Missing required property 'public_key'")
            __props__.__dict__["public_key"] = public_key
            __props__.__dict__["path"] = None
        super(SecretBackend, __self__).__init__(
            'vault:mongodbatlas/secretBackend:SecretBackend',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            mount: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            path: Optional[pulumi.Input[str]] = None,
            private_key: Optional[pulumi.Input[str]] = None,
            public_key: Optional[pulumi.Input[str]] = None) -> 'SecretBackend':
        """
        Get an existing SecretBackend resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] mount: Path where the MongoDB Atlas Secrets Engine is mounted.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: Path where MongoDB Atlas configuration is located
        :param pulumi.Input[str] private_key: Specifies the Private API Key used to authenticate with the MongoDB Atlas API.
        :param pulumi.Input[str] public_key: Specifies the Public API Key used to authenticate with the MongoDB Atlas API.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretBackendState.__new__(_SecretBackendState)

        __props__.__dict__["mount"] = mount
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["path"] = path
        __props__.__dict__["private_key"] = private_key
        __props__.__dict__["public_key"] = public_key
        return SecretBackend(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def mount(self) -> pulumi.Output[str]:
        """
        Path where the MongoDB Atlas Secrets Engine is mounted.
        """
        return pulumi.get(self, "mount")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[str]:
        """
        Path where MongoDB Atlas configuration is located
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Output[str]:
        """
        Specifies the Private API Key used to authenticate with the MongoDB Atlas API.
        """
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Output[str]:
        """
        Specifies the Public API Key used to authenticate with the MongoDB Atlas API.
        """
        return pulumi.get(self, "public_key")

