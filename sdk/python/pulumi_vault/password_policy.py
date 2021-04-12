# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities, _tables

__all__ = ['PasswordPolicyArgs', 'PasswordPolicy']

@pulumi.input_type
class PasswordPolicyArgs:
    def __init__(__self__, *,
                 policy: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PasswordPolicy resource.
        :param pulumi.Input[str] policy: String containing a password policy.
        :param pulumi.Input[str] name: The name of the password policy.
        """
        pulumi.set(__self__, "policy", policy)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Input[str]:
        """
        String containing a password policy.
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the password policy.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class PasswordPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a resource to manage Password Policies

        **Note** this feature is available only Vault 1.5+

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        alphanumeric = vault.PasswordPolicy("alphanumeric", policy=\"\"\"length = 20
        rule "charset" {
        	  charset = "abcdefghijklmnopqrstuvwxyz0123456789"
        }

        \"\"\")
        ```

        ## Import

        Password policies can be imported using the `name`, e.g.

        ```sh
         $ pulumi import vault:index/passwordPolicy:PasswordPolicy alphanumeric alphanumeric
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the password policy.
        :param pulumi.Input[str] policy: String containing a password policy.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PasswordPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage Password Policies

        **Note** this feature is available only Vault 1.5+

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        alphanumeric = vault.PasswordPolicy("alphanumeric", policy=\"\"\"length = 20
        rule "charset" {
        	  charset = "abcdefghijklmnopqrstuvwxyz0123456789"
        }

        \"\"\")
        ```

        ## Import

        Password policies can be imported using the `name`, e.g.

        ```sh
         $ pulumi import vault:index/passwordPolicy:PasswordPolicy alphanumeric alphanumeric
        ```

        :param str resource_name: The name of the resource.
        :param PasswordPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PasswordPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            __props__['name'] = name
            if policy is None and not opts.urn:
                raise TypeError("Missing required property 'policy'")
            __props__['policy'] = policy
        super(PasswordPolicy, __self__).__init__(
            'vault:index/passwordPolicy:PasswordPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            policy: Optional[pulumi.Input[str]] = None) -> 'PasswordPolicy':
        """
        Get an existing PasswordPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the password policy.
        :param pulumi.Input[str] policy: String containing a password policy.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["name"] = name
        __props__["policy"] = policy
        return PasswordPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the password policy.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[str]:
        """
        String containing a password policy.
        """
        return pulumi.get(self, "policy")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

