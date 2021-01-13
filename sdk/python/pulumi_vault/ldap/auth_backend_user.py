# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['AuthBackendUser']


class AuthBackendUser(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a resource to create a user in an [LDAP auth backend within Vault](https://www.vaultproject.io/docs/auth/ldap.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        ldap = vault.ldap.AuthBackend("ldap",
            discoverdn=False,
            groupdn="OU=Groups,DC=example,DC=org",
            groupfilter="(&(objectClass=group)(member:1.2.840.113556.1.4.1941:={{.UserDN}}))",
            path="ldap",
            upndomain="EXAMPLE.ORG",
            url="ldaps://dc-01.example.org",
            userattr="sAMAccountName",
            userdn="OU=Users,OU=Accounts,DC=example,DC=org")
        user = vault.ldap.AuthBackendUser("user",
            backend=ldap.path,
            policies=[
                "dba",
                "sysops",
            ],
            username="test-user")
        ```

        ## Import

        LDAP authentication backend users can be imported using the `path`, e.g.

        ```sh
         $ pulumi import vault:ldap/authBackendUser:AuthBackendUser foo auth/ldap/users/foo
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: Path to the authentication backend
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: Override LDAP groups which should be granted to user
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: Policies which should be granted to user
        :param pulumi.Input[str] username: The LDAP username
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
            __props__['groups'] = groups
            __props__['policies'] = policies
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__['username'] = username
        super(AuthBackendUser, __self__).__init__(
            'vault:ldap/authBackendUser:AuthBackendUser',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend: Optional[pulumi.Input[str]] = None,
            groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'AuthBackendUser':
        """
        Get an existing AuthBackendUser resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: Path to the authentication backend
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: Override LDAP groups which should be granted to user
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: Policies which should be granted to user
        :param pulumi.Input[str] username: The LDAP username
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backend"] = backend
        __props__["groups"] = groups
        __props__["policies"] = policies
        __props__["username"] = username
        return AuthBackendUser(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[Optional[str]]:
        """
        Path to the authentication backend
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter
    def groups(self) -> pulumi.Output[Sequence[str]]:
        """
        Override LDAP groups which should be granted to user
        """
        return pulumi.get(self, "groups")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Sequence[str]]:
        """
        Policies which should be granted to user
        """
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        The LDAP username
        """
        return pulumi.get(self, "username")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

