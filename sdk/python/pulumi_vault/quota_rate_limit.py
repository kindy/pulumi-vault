# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables

__all__ = ['QuotaRateLimit']


class QuotaRateLimit(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 rate: Optional[pulumi.Input[float]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manage rate limit quotas which enforce API rate limiting using a token bucket algorithm.
        A rate limit quota can be created at the root level or defined on a namespace or mount by
        specifying a path when creating the quota.

        See [Vault's Documentation](https://www.vaultproject.io/docs/concepts/resource-quotas) for more
        information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        global_ = vault.QuotaRateLimit("global",
            path="",
            rate=100)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of the rate limit quota
        :param pulumi.Input[str] path: Path of the mount or namespace to apply the quota. A blank path configures a
               global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
               `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
               Updating this field on an existing quota can have "moving" effects. For example, updating
               `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
               a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
        :param pulumi.Input[float] rate: The maximum number of requests at any given second to be allowed by the quota
               rule. The `rate` must be positive.
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

            __props__['name'] = name
            __props__['path'] = path
            if rate is None:
                raise TypeError("Missing required property 'rate'")
            __props__['rate'] = rate
        super(QuotaRateLimit, __self__).__init__(
            'vault:index/quotaRateLimit:QuotaRateLimit',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            path: Optional[pulumi.Input[str]] = None,
            rate: Optional[pulumi.Input[float]] = None) -> 'QuotaRateLimit':
        """
        Get an existing QuotaRateLimit resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of the rate limit quota
        :param pulumi.Input[str] path: Path of the mount or namespace to apply the quota. A blank path configures a
               global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
               `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
               Updating this field on an existing quota can have "moving" effects. For example, updating
               `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
               a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
        :param pulumi.Input[float] rate: The maximum number of requests at any given second to be allowed by the quota
               rule. The `rate` must be positive.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["name"] = name
        __props__["path"] = path
        __props__["rate"] = rate
        return QuotaRateLimit(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the rate limit quota
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[Optional[str]]:
        """
        Path of the mount or namespace to apply the quota. A blank path configures a
        global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
        `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
        Updating this field on an existing quota can have "moving" effects. For example, updating
        `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
        a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def rate(self) -> pulumi.Output[float]:
        """
        The maximum number of requests at any given second to be allowed by the quota
        rule. The `rate` must be positive.
        """
        return pulumi.get(self, "rate")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

