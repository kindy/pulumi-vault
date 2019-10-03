# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class SecretBackendRole(pulumi.CustomResource):
    backend: pulumi.Output[str]
    """
    The path the RabbitMQ secret backend is mounted at,
    with no leading or trailing `/`s.
    """
    name: pulumi.Output[str]
    """
    The name to identify this role within the backend.
    Must be unique within the backend.
    """
    tags: pulumi.Output[str]
    """
    Specifies a comma-separated RabbitMQ management tags.
    """
    vhosts: pulumi.Output[list]
    def __init__(__self__, resource_name, opts=None, backend=None, name=None, tags=None, vhosts=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a SecretBackendRole resource with the given unique name, props, and options.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The path the RabbitMQ secret backend is mounted at,
               with no leading or trailing `/`s.
        :param pulumi.Input[str] name: The name to identify this role within the backend.
               Must be unique within the backend.
        :param pulumi.Input[str] tags: Specifies a comma-separated RabbitMQ management tags.
        
        The **vhosts** object supports the following:
        
          * `configure` (`pulumi.Input[str]`)
          * `host` (`pulumi.Input[str]`)
          * `read` (`pulumi.Input[str]`)
          * `write` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/rabbitmq_secret_backend_role.html.markdown.
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

            if backend is None:
                raise TypeError("Missing required property 'backend'")
            __props__['backend'] = backend
            __props__['name'] = name
            __props__['tags'] = tags
            __props__['vhosts'] = vhosts
        super(SecretBackendRole, __self__).__init__(
            'vault:rabbitMq/secretBackendRole:SecretBackendRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, backend=None, name=None, tags=None, vhosts=None):
        """
        Get an existing SecretBackendRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The path the RabbitMQ secret backend is mounted at,
               with no leading or trailing `/`s.
        :param pulumi.Input[str] name: The name to identify this role within the backend.
               Must be unique within the backend.
        :param pulumi.Input[str] tags: Specifies a comma-separated RabbitMQ management tags.
        
        The **vhosts** object supports the following:
        
          * `configure` (`pulumi.Input[str]`)
          * `host` (`pulumi.Input[str]`)
          * `read` (`pulumi.Input[str]`)
          * `write` (`pulumi.Input[str]`)

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/rabbitmq_secret_backend_role.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["backend"] = backend
        __props__["name"] = name
        __props__["tags"] = tags
        __props__["vhosts"] = vhosts
        return SecretBackendRole(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

