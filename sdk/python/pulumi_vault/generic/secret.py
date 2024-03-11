# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SecretArgs', 'Secret']

@pulumi.input_type
class SecretArgs:
    def __init__(__self__, *,
                 data_json: pulumi.Input[str],
                 path: pulumi.Input[str],
                 delete_all_versions: Optional[pulumi.Input[bool]] = None,
                 disable_read: Optional[pulumi.Input[bool]] = None,
                 namespace: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Secret resource.
        :param pulumi.Input[str] data_json: String containing a JSON-encoded object that will be
               written as the secret data at the given path.
        :param pulumi.Input[str] path: The full logical path at which to write the given data.
               To write data into the "generic" secret backend mounted in Vault by default,
               this should be prefixed with `secret/`. Writing to other backends with this
               resource is possible; consult each backend's documentation to see which
               endpoints support the `PUT` and `DELETE` methods.
        :param pulumi.Input[bool] delete_all_versions: true/false.  Only applicable for kv-v2 stores.
               If set to `true`, permanently deletes all versions for
               the specified key. The default behavior is to only delete the latest version of the
               secret.
        :param pulumi.Input[bool] disable_read: true/false. Set this to true if your vault
               authentication is not able to read the data. Setting this to `true` will
               break drift detection. Defaults to false.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        """
        pulumi.set(__self__, "data_json", data_json)
        pulumi.set(__self__, "path", path)
        if delete_all_versions is not None:
            pulumi.set(__self__, "delete_all_versions", delete_all_versions)
        if disable_read is not None:
            pulumi.set(__self__, "disable_read", disable_read)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter(name="dataJson")
    def data_json(self) -> pulumi.Input[str]:
        """
        String containing a JSON-encoded object that will be
        written as the secret data at the given path.
        """
        return pulumi.get(self, "data_json")

    @data_json.setter
    def data_json(self, value: pulumi.Input[str]):
        pulumi.set(self, "data_json", value)

    @property
    @pulumi.getter
    def path(self) -> pulumi.Input[str]:
        """
        The full logical path at which to write the given data.
        To write data into the "generic" secret backend mounted in Vault by default,
        this should be prefixed with `secret/`. Writing to other backends with this
        resource is possible; consult each backend's documentation to see which
        endpoints support the `PUT` and `DELETE` methods.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: pulumi.Input[str]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter(name="deleteAllVersions")
    def delete_all_versions(self) -> Optional[pulumi.Input[bool]]:
        """
        true/false.  Only applicable for kv-v2 stores.
        If set to `true`, permanently deletes all versions for
        the specified key. The default behavior is to only delete the latest version of the
        secret.
        """
        return pulumi.get(self, "delete_all_versions")

    @delete_all_versions.setter
    def delete_all_versions(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_all_versions", value)

    @property
    @pulumi.getter(name="disableRead")
    def disable_read(self) -> Optional[pulumi.Input[bool]]:
        """
        true/false. Set this to true if your vault
        authentication is not able to read the data. Setting this to `true` will
        break drift detection. Defaults to false.
        """
        return pulumi.get(self, "disable_read")

    @disable_read.setter
    def disable_read(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_read", value)

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
class _SecretState:
    def __init__(__self__, *,
                 data: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 data_json: Optional[pulumi.Input[str]] = None,
                 delete_all_versions: Optional[pulumi.Input[bool]] = None,
                 disable_read: Optional[pulumi.Input[bool]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Secret resources.
        :param pulumi.Input[Mapping[str, Any]] data: A mapping whose keys are the top-level data keys returned from
               Vault and whose values are the corresponding values. This map can only
               represent string data, so any non-string values returned from Vault are
               serialized as JSON.
        :param pulumi.Input[str] data_json: String containing a JSON-encoded object that will be
               written as the secret data at the given path.
        :param pulumi.Input[bool] delete_all_versions: true/false.  Only applicable for kv-v2 stores.
               If set to `true`, permanently deletes all versions for
               the specified key. The default behavior is to only delete the latest version of the
               secret.
        :param pulumi.Input[bool] disable_read: true/false. Set this to true if your vault
               authentication is not able to read the data. Setting this to `true` will
               break drift detection. Defaults to false.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: The full logical path at which to write the given data.
               To write data into the "generic" secret backend mounted in Vault by default,
               this should be prefixed with `secret/`. Writing to other backends with this
               resource is possible; consult each backend's documentation to see which
               endpoints support the `PUT` and `DELETE` methods.
        """
        if data is not None:
            pulumi.set(__self__, "data", data)
        if data_json is not None:
            pulumi.set(__self__, "data_json", data_json)
        if delete_all_versions is not None:
            pulumi.set(__self__, "delete_all_versions", delete_all_versions)
        if disable_read is not None:
            pulumi.set(__self__, "disable_read", disable_read)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if path is not None:
            pulumi.set(__self__, "path", path)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        A mapping whose keys are the top-level data keys returned from
        Vault and whose values are the corresponding values. This map can only
        represent string data, so any non-string values returned from Vault are
        serialized as JSON.
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter(name="dataJson")
    def data_json(self) -> Optional[pulumi.Input[str]]:
        """
        String containing a JSON-encoded object that will be
        written as the secret data at the given path.
        """
        return pulumi.get(self, "data_json")

    @data_json.setter
    def data_json(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data_json", value)

    @property
    @pulumi.getter(name="deleteAllVersions")
    def delete_all_versions(self) -> Optional[pulumi.Input[bool]]:
        """
        true/false.  Only applicable for kv-v2 stores.
        If set to `true`, permanently deletes all versions for
        the specified key. The default behavior is to only delete the latest version of the
        secret.
        """
        return pulumi.get(self, "delete_all_versions")

    @delete_all_versions.setter
    def delete_all_versions(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_all_versions", value)

    @property
    @pulumi.getter(name="disableRead")
    def disable_read(self) -> Optional[pulumi.Input[bool]]:
        """
        true/false. Set this to true if your vault
        authentication is not able to read the data. Setting this to `true` will
        break drift detection. Defaults to false.
        """
        return pulumi.get(self, "disable_read")

    @disable_read.setter
    def disable_read(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_read", value)

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
        The full logical path at which to write the given data.
        To write data into the "generic" secret backend mounted in Vault by default,
        this should be prefixed with `secret/`. Writing to other backends with this
        resource is possible; consult each backend's documentation to see which
        endpoints support the `PUT` and `DELETE` methods.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)


class Secret(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_json: Optional[pulumi.Input[str]] = None,
                 delete_all_versions: Optional[pulumi.Input[bool]] = None,
                 disable_read: Optional[pulumi.Input[bool]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Import

        Generic secrets can be imported using the `path`, e.g.

        ```sh
        $ pulumi import vault:generic/secret:Secret example secret/foo
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_json: String containing a JSON-encoded object that will be
               written as the secret data at the given path.
        :param pulumi.Input[bool] delete_all_versions: true/false.  Only applicable for kv-v2 stores.
               If set to `true`, permanently deletes all versions for
               the specified key. The default behavior is to only delete the latest version of the
               secret.
        :param pulumi.Input[bool] disable_read: true/false. Set this to true if your vault
               authentication is not able to read the data. Setting this to `true` will
               break drift detection. Defaults to false.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: The full logical path at which to write the given data.
               To write data into the "generic" secret backend mounted in Vault by default,
               this should be prefixed with `secret/`. Writing to other backends with this
               resource is possible; consult each backend's documentation to see which
               endpoints support the `PUT` and `DELETE` methods.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Generic secrets can be imported using the `path`, e.g.

        ```sh
        $ pulumi import vault:generic/secret:Secret example secret/foo
        ```

        :param str resource_name: The name of the resource.
        :param SecretArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_json: Optional[pulumi.Input[str]] = None,
                 delete_all_versions: Optional[pulumi.Input[bool]] = None,
                 disable_read: Optional[pulumi.Input[bool]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretArgs.__new__(SecretArgs)

            if data_json is None and not opts.urn:
                raise TypeError("Missing required property 'data_json'")
            __props__.__dict__["data_json"] = None if data_json is None else pulumi.Output.secret(data_json)
            __props__.__dict__["delete_all_versions"] = delete_all_versions
            __props__.__dict__["disable_read"] = disable_read
            __props__.__dict__["namespace"] = namespace
            if path is None and not opts.urn:
                raise TypeError("Missing required property 'path'")
            __props__.__dict__["path"] = path
            __props__.__dict__["data"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["data", "dataJson"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Secret, __self__).__init__(
            'vault:generic/secret:Secret',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            data: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            data_json: Optional[pulumi.Input[str]] = None,
            delete_all_versions: Optional[pulumi.Input[bool]] = None,
            disable_read: Optional[pulumi.Input[bool]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            path: Optional[pulumi.Input[str]] = None) -> 'Secret':
        """
        Get an existing Secret resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] data: A mapping whose keys are the top-level data keys returned from
               Vault and whose values are the corresponding values. This map can only
               represent string data, so any non-string values returned from Vault are
               serialized as JSON.
        :param pulumi.Input[str] data_json: String containing a JSON-encoded object that will be
               written as the secret data at the given path.
        :param pulumi.Input[bool] delete_all_versions: true/false.  Only applicable for kv-v2 stores.
               If set to `true`, permanently deletes all versions for
               the specified key. The default behavior is to only delete the latest version of the
               secret.
        :param pulumi.Input[bool] disable_read: true/false. Set this to true if your vault
               authentication is not able to read the data. Setting this to `true` will
               break drift detection. Defaults to false.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: The full logical path at which to write the given data.
               To write data into the "generic" secret backend mounted in Vault by default,
               this should be prefixed with `secret/`. Writing to other backends with this
               resource is possible; consult each backend's documentation to see which
               endpoints support the `PUT` and `DELETE` methods.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretState.__new__(_SecretState)

        __props__.__dict__["data"] = data
        __props__.__dict__["data_json"] = data_json
        __props__.__dict__["delete_all_versions"] = delete_all_versions
        __props__.__dict__["disable_read"] = disable_read
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["path"] = path
        return Secret(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def data(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        A mapping whose keys are the top-level data keys returned from
        Vault and whose values are the corresponding values. This map can only
        represent string data, so any non-string values returned from Vault are
        serialized as JSON.
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter(name="dataJson")
    def data_json(self) -> pulumi.Output[str]:
        """
        String containing a JSON-encoded object that will be
        written as the secret data at the given path.
        """
        return pulumi.get(self, "data_json")

    @property
    @pulumi.getter(name="deleteAllVersions")
    def delete_all_versions(self) -> pulumi.Output[Optional[bool]]:
        """
        true/false.  Only applicable for kv-v2 stores.
        If set to `true`, permanently deletes all versions for
        the specified key. The default behavior is to only delete the latest version of the
        secret.
        """
        return pulumi.get(self, "delete_all_versions")

    @property
    @pulumi.getter(name="disableRead")
    def disable_read(self) -> pulumi.Output[Optional[bool]]:
        """
        true/false. Set this to true if your vault
        authentication is not able to read the data. Setting this to `true` will
        break drift detection. Defaults to false.
        """
        return pulumi.get(self, "disable_read")

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
        The full logical path at which to write the given data.
        To write data into the "generic" secret backend mounted in Vault by default,
        this should be prefixed with `secret/`. Writing to other backends with this
        resource is possible; consult each backend's documentation to see which
        endpoints support the `PUT` and `DELETE` methods.
        """
        return pulumi.get(self, "path")

