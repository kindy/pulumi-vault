# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class AuthBackendGroup(pulumi.CustomResource):
    backend: pulumi.Output[str]
    """
    Path to the authentication backend
    """
    groupname: pulumi.Output[str]
    """
    The LDAP groupname
    """
    policies: pulumi.Output[list]
    """
    Policies which should be granted to members of the group
    """
    def __init__(__self__, resource_name, opts=None, backend=None, groupname=None, policies=None, __props__=None, __name__=None, __opts__=None):
        """
        Provides a resource to create a group in an [LDAP auth backend within Vault](https://www.vaultproject.io/docs/auth/ldap.html).

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
        group = vault.ldap.AuthBackendGroup("group",
            backend=ldap.path,
            groupname="dba",
            policies=["dba"])
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: Path to the authentication backend
        :param pulumi.Input[str] groupname: The LDAP groupname
        :param pulumi.Input[list] policies: Policies which should be granted to members of the group
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['backend'] = backend
            if groupname is None:
                raise TypeError("Missing required property 'groupname'")
            __props__['groupname'] = groupname
            __props__['policies'] = policies
        super(AuthBackendGroup, __self__).__init__(
            'vault:ldap/authBackendGroup:AuthBackendGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, backend=None, groupname=None, policies=None):
        """
        Get an existing AuthBackendGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: Path to the authentication backend
        :param pulumi.Input[str] groupname: The LDAP groupname
        :param pulumi.Input[list] policies: Policies which should be granted to members of the group
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backend"] = backend
        __props__["groupname"] = groupname
        __props__["policies"] = policies
        return AuthBackendGroup(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

