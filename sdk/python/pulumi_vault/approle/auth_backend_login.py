# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['AuthBackendLogin']


class AuthBackendLogin(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 role_id: Optional[pulumi.Input[str]] = None,
                 secret_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Logs into Vault using the AppRole auth backend. See the [Vault
        documentation](https://www.vaultproject.io/docs/auth/approle) for more
        information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        approle = vault.AuthBackend("approle", type="approle")
        example = vault.app_role.AuthBackendRole("example",
            backend=approle.path,
            policies=[
                "default",
                "dev",
                "prod",
            ],
            role_name="test-role")
        id = vault.app_role.AuthBackendRoleSecretID("id",
            backend=approle.path,
            role_name=example.role_name)
        login = vault.app_role.AuthBackendLogin("login",
            backend=approle.path,
            role_id=example.role_id,
            secret_id=id.secret_id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The unique path of the Vault backend to log in with.
        :param pulumi.Input[str] role_id: The ID of the role to log in with.
        :param pulumi.Input[str] secret_id: The secret ID of the role to log in with. Required
               unless `bind_secret_id` is set to false on the role.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['backend'] = backend
            if role_id is None and not opts.urn:
                raise TypeError("Missing required property 'role_id'")
            __props__['role_id'] = role_id
            __props__['secret_id'] = secret_id
            __props__['accessor'] = None
            __props__['client_token'] = None
            __props__['lease_duration'] = None
            __props__['lease_started'] = None
            __props__['metadata'] = None
            __props__['policies'] = None
            __props__['renewable'] = None
        super(AuthBackendLogin, __self__).__init__(
            'vault:appRole/authBackendLogin:AuthBackendLogin',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accessor: Optional[pulumi.Input[str]] = None,
            backend: Optional[pulumi.Input[str]] = None,
            client_token: Optional[pulumi.Input[str]] = None,
            lease_duration: Optional[pulumi.Input[int]] = None,
            lease_started: Optional[pulumi.Input[str]] = None,
            metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            renewable: Optional[pulumi.Input[bool]] = None,
            role_id: Optional[pulumi.Input[str]] = None,
            secret_id: Optional[pulumi.Input[str]] = None) -> 'AuthBackendLogin':
        """
        Get an existing AuthBackendLogin resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accessor: The accessor for the token.
        :param pulumi.Input[str] backend: The unique path of the Vault backend to log in with.
        :param pulumi.Input[str] client_token: The Vault token created.
        :param pulumi.Input[int] lease_duration: How long the token is valid for, in seconds.
        :param pulumi.Input[str] lease_started: The date and time the lease started, in RFC 3339 format.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: The metadata associated with the token.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: A list of policies applied to the token.
        :param pulumi.Input[bool] renewable: Whether the token is renewable or not.
        :param pulumi.Input[str] role_id: The ID of the role to log in with.
        :param pulumi.Input[str] secret_id: The secret ID of the role to log in with. Required
               unless `bind_secret_id` is set to false on the role.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["accessor"] = accessor
        __props__["backend"] = backend
        __props__["client_token"] = client_token
        __props__["lease_duration"] = lease_duration
        __props__["lease_started"] = lease_started
        __props__["metadata"] = metadata
        __props__["policies"] = policies
        __props__["renewable"] = renewable
        __props__["role_id"] = role_id
        __props__["secret_id"] = secret_id
        return AuthBackendLogin(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def accessor(self) -> pulumi.Output[str]:
        """
        The accessor for the token.
        """
        return pulumi.get(self, "accessor")

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[Optional[str]]:
        """
        The unique path of the Vault backend to log in with.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="clientToken")
    def client_token(self) -> pulumi.Output[str]:
        """
        The Vault token created.
        """
        return pulumi.get(self, "client_token")

    @property
    @pulumi.getter(name="leaseDuration")
    def lease_duration(self) -> pulumi.Output[int]:
        """
        How long the token is valid for, in seconds.
        """
        return pulumi.get(self, "lease_duration")

    @property
    @pulumi.getter(name="leaseStarted")
    def lease_started(self) -> pulumi.Output[str]:
        """
        The date and time the lease started, in RFC 3339 format.
        """
        return pulumi.get(self, "lease_started")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Mapping[str, str]]:
        """
        The metadata associated with the token.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of policies applied to the token.
        """
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter
    def renewable(self) -> pulumi.Output[bool]:
        """
        Whether the token is renewable or not.
        """
        return pulumi.get(self, "renewable")

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Output[str]:
        """
        The ID of the role to log in with.
        """
        return pulumi.get(self, "role_id")

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> pulumi.Output[Optional[str]]:
        """
        The secret ID of the role to log in with. Required
        unless `bind_secret_id` is set to false on the role.
        """
        return pulumi.get(self, "secret_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

