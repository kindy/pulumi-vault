# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['AuthBackendConfig']


class AuthBackendConfig(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 environment: Optional[pulumi.Input[str]] = None,
                 resource: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## Import

        Azure auth backends can be imported using `auth/`, the `backend` path, and `/config` e.g.

        ```sh
         $ pulumi import vault:azure/authBackendConfig:AuthBackendConfig example auth/azure/config
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The path the Azure auth backend being configured was
               mounted at.  Defaults to `azure`.
        :param pulumi.Input[str] client_id: The client id for credentials to query the Azure APIs.
               Currently read permissions to query compute resources are required.
        :param pulumi.Input[str] client_secret: The client secret for credentials to query the
               Azure APIs.
        :param pulumi.Input[str] environment: The Azure cloud environment. Valid values:
               AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
               AzureGermanCloud.  Defaults to `AzurePublicCloud`.
        :param pulumi.Input[str] resource: The configured URL for the application registered in
               Azure Active Directory.
        :param pulumi.Input[str] tenant_id: The tenant id for the Azure Active Directory
               organization.
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
            __props__['client_id'] = client_id
            __props__['client_secret'] = client_secret
            __props__['environment'] = environment
            if resource is None:
                raise TypeError("Missing required property 'resource'")
            __props__['resource'] = resource
            if tenant_id is None:
                raise TypeError("Missing required property 'tenant_id'")
            __props__['tenant_id'] = tenant_id
        super(AuthBackendConfig, __self__).__init__(
            'vault:azure/authBackendConfig:AuthBackendConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend: Optional[pulumi.Input[str]] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            client_secret: Optional[pulumi.Input[str]] = None,
            environment: Optional[pulumi.Input[str]] = None,
            resource: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None) -> 'AuthBackendConfig':
        """
        Get an existing AuthBackendConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The path the Azure auth backend being configured was
               mounted at.  Defaults to `azure`.
        :param pulumi.Input[str] client_id: The client id for credentials to query the Azure APIs.
               Currently read permissions to query compute resources are required.
        :param pulumi.Input[str] client_secret: The client secret for credentials to query the
               Azure APIs.
        :param pulumi.Input[str] environment: The Azure cloud environment. Valid values:
               AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
               AzureGermanCloud.  Defaults to `AzurePublicCloud`.
        :param pulumi.Input[str] resource: The configured URL for the application registered in
               Azure Active Directory.
        :param pulumi.Input[str] tenant_id: The tenant id for the Azure Active Directory
               organization.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backend"] = backend
        __props__["client_id"] = client_id
        __props__["client_secret"] = client_secret
        __props__["environment"] = environment
        __props__["resource"] = resource
        __props__["tenant_id"] = tenant_id
        return AuthBackendConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[Optional[str]]:
        """
        The path the Azure auth backend being configured was
        mounted at.  Defaults to `azure`.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[Optional[str]]:
        """
        The client id for credentials to query the Azure APIs.
        Currently read permissions to query compute resources are required.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> pulumi.Output[Optional[str]]:
        """
        The client secret for credentials to query the
        Azure APIs.
        """
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output[Optional[str]]:
        """
        The Azure cloud environment. Valid values:
        AzurePublicCloud, AzureUSGovernmentCloud, AzureChinaCloud,
        AzureGermanCloud.  Defaults to `AzurePublicCloud`.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter
    def resource(self) -> pulumi.Output[str]:
        """
        The configured URL for the application registered in
        Azure Active Directory.
        """
        return pulumi.get(self, "resource")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[str]:
        """
        The tenant id for the Azure Active Directory
        organization.
        """
        return pulumi.get(self, "tenant_id")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

