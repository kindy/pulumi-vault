# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SecretBackendCaArgs', 'SecretBackendCa']

@pulumi.input_type
class SecretBackendCaArgs:
    def __init__(__self__, *,
                 backend: Optional[pulumi.Input[str]] = None,
                 generate_signing_key: Optional[pulumi.Input[bool]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SecretBackendCa resource.
        :param pulumi.Input[str] backend: The path where the SSH secret backend is mounted. Defaults to 'ssh'
        :param pulumi.Input[bool] generate_signing_key: Whether Vault should generate the signing key pair internally. Defaults to true
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] private_key: Private key part the SSH CA key pair; required if generate_signing_key is false.
        :param pulumi.Input[str] public_key: The public key part the SSH CA key pair; required if generate_signing_key is false.
        """
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if generate_signing_key is not None:
            pulumi.set(__self__, "generate_signing_key", generate_signing_key)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if private_key is not None:
            pulumi.set(__self__, "private_key", private_key)
        if public_key is not None:
            pulumi.set(__self__, "public_key", public_key)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[str]]:
        """
        The path where the SSH secret backend is mounted. Defaults to 'ssh'
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="generateSigningKey")
    def generate_signing_key(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether Vault should generate the signing key pair internally. Defaults to true
        """
        return pulumi.get(self, "generate_signing_key")

    @generate_signing_key.setter
    def generate_signing_key(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "generate_signing_key", value)

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
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[pulumi.Input[str]]:
        """
        Private key part the SSH CA key pair; required if generate_signing_key is false.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> Optional[pulumi.Input[str]]:
        """
        The public key part the SSH CA key pair; required if generate_signing_key is false.
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key", value)


@pulumi.input_type
class _SecretBackendCaState:
    def __init__(__self__, *,
                 backend: Optional[pulumi.Input[str]] = None,
                 generate_signing_key: Optional[pulumi.Input[bool]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SecretBackendCa resources.
        :param pulumi.Input[str] backend: The path where the SSH secret backend is mounted. Defaults to 'ssh'
        :param pulumi.Input[bool] generate_signing_key: Whether Vault should generate the signing key pair internally. Defaults to true
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] private_key: Private key part the SSH CA key pair; required if generate_signing_key is false.
        :param pulumi.Input[str] public_key: The public key part the SSH CA key pair; required if generate_signing_key is false.
        """
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if generate_signing_key is not None:
            pulumi.set(__self__, "generate_signing_key", generate_signing_key)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if private_key is not None:
            pulumi.set(__self__, "private_key", private_key)
        if public_key is not None:
            pulumi.set(__self__, "public_key", public_key)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[str]]:
        """
        The path where the SSH secret backend is mounted. Defaults to 'ssh'
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="generateSigningKey")
    def generate_signing_key(self) -> Optional[pulumi.Input[bool]]:
        """
        Whether Vault should generate the signing key pair internally. Defaults to true
        """
        return pulumi.get(self, "generate_signing_key")

    @generate_signing_key.setter
    def generate_signing_key(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "generate_signing_key", value)

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
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[pulumi.Input[str]]:
        """
        Private key part the SSH CA key pair; required if generate_signing_key is false.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_key", value)

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> Optional[pulumi.Input[str]]:
        """
        The public key part the SSH CA key pair; required if generate_signing_key is false.
        """
        return pulumi.get(self, "public_key")

    @public_key.setter
    def public_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_key", value)


class SecretBackendCa(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 generate_signing_key: Optional[pulumi.Input[bool]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 private_key: Optional[pulumi.Input[str]] = None,
                 public_key: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage CA information in an SSH secret backend
        [SSH secret backend within Vault](https://www.vaultproject.io/docs/secrets/ssh/index.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.Mount("example", type="ssh")
        foo = vault.ssh.SecretBackendCa("foo", backend=example.path)
        ```

        ## Import

        SSH secret backend CAs can be imported using the `path`, e.g.

        ```sh
         $ pulumi import vault:ssh/secretBackendCa:SecretBackendCa foo ssh
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The path where the SSH secret backend is mounted. Defaults to 'ssh'
        :param pulumi.Input[bool] generate_signing_key: Whether Vault should generate the signing key pair internally. Defaults to true
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] private_key: Private key part the SSH CA key pair; required if generate_signing_key is false.
        :param pulumi.Input[str] public_key: The public key part the SSH CA key pair; required if generate_signing_key is false.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SecretBackendCaArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage CA information in an SSH secret backend
        [SSH secret backend within Vault](https://www.vaultproject.io/docs/secrets/ssh/index.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.Mount("example", type="ssh")
        foo = vault.ssh.SecretBackendCa("foo", backend=example.path)
        ```

        ## Import

        SSH secret backend CAs can be imported using the `path`, e.g.

        ```sh
         $ pulumi import vault:ssh/secretBackendCa:SecretBackendCa foo ssh
        ```

        :param str resource_name: The name of the resource.
        :param SecretBackendCaArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretBackendCaArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 generate_signing_key: Optional[pulumi.Input[bool]] = None,
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
            __props__ = SecretBackendCaArgs.__new__(SecretBackendCaArgs)

            __props__.__dict__["backend"] = backend
            __props__.__dict__["generate_signing_key"] = generate_signing_key
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["private_key"] = None if private_key is None else pulumi.Output.secret(private_key)
            __props__.__dict__["public_key"] = public_key
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["privateKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SecretBackendCa, __self__).__init__(
            'vault:ssh/secretBackendCa:SecretBackendCa',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend: Optional[pulumi.Input[str]] = None,
            generate_signing_key: Optional[pulumi.Input[bool]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            private_key: Optional[pulumi.Input[str]] = None,
            public_key: Optional[pulumi.Input[str]] = None) -> 'SecretBackendCa':
        """
        Get an existing SecretBackendCa resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The path where the SSH secret backend is mounted. Defaults to 'ssh'
        :param pulumi.Input[bool] generate_signing_key: Whether Vault should generate the signing key pair internally. Defaults to true
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] private_key: Private key part the SSH CA key pair; required if generate_signing_key is false.
        :param pulumi.Input[str] public_key: The public key part the SSH CA key pair; required if generate_signing_key is false.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretBackendCaState.__new__(_SecretBackendCaState)

        __props__.__dict__["backend"] = backend
        __props__.__dict__["generate_signing_key"] = generate_signing_key
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["private_key"] = private_key
        __props__.__dict__["public_key"] = public_key
        return SecretBackendCa(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[Optional[str]]:
        """
        The path where the SSH secret backend is mounted. Defaults to 'ssh'
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="generateSigningKey")
    def generate_signing_key(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether Vault should generate the signing key pair internally. Defaults to true
        """
        return pulumi.get(self, "generate_signing_key")

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
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Output[str]:
        """
        Private key part the SSH CA key pair; required if generate_signing_key is false.
        """
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter(name="publicKey")
    def public_key(self) -> pulumi.Output[str]:
        """
        The public key part the SSH CA key pair; required if generate_signing_key is false.
        """
        return pulumi.get(self, "public_key")

