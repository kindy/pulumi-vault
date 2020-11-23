# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['AuthBackendIdentityWhitelist']


class AuthBackendIdentityWhitelist(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 disable_periodic_tidy: Optional[pulumi.Input[bool]] = None,
                 safety_buffer: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Configures the periodic tidying operation of the whitelisted identity entries.

        For more information, see the
        [Vault docs](https://www.vaultproject.io/api-docs/auth/aws#configure-identity-whitelist-tidy-operation).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        example_auth_backend = vault.AuthBackend("exampleAuthBackend", type="aws")
        example_auth_backend_identity_whitelist = vault.aws.AuthBackendIdentityWhitelist("exampleAuthBackendIdentityWhitelist",
            backend=example_auth_backend.path,
            safety_buffer=3600)
        ```

        ## Import

        AWS auth backend identity whitelists can be imported using `auth/`, the `backend` path, and `/config/tidy/identity-whitelist` e.g.

        ```sh
         $ pulumi import vault:aws/authBackendIdentityWhitelist:AuthBackendIdentityWhitelist example auth/aws/config/tidy/identity-whitelist
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The path of the AWS backend being configured.
        :param pulumi.Input[bool] disable_periodic_tidy: If set to true, disables the periodic
               tidying of the identity-whitelist entries.
        :param pulumi.Input[int] safety_buffer: The amount of extra time, in minutes, that must
               have passed beyond the roletag expiration, before it is removed from the
               backend storage.
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
            __props__['disable_periodic_tidy'] = disable_periodic_tidy
            __props__['safety_buffer'] = safety_buffer
        super(AuthBackendIdentityWhitelist, __self__).__init__(
            'vault:aws/authBackendIdentityWhitelist:AuthBackendIdentityWhitelist',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend: Optional[pulumi.Input[str]] = None,
            disable_periodic_tidy: Optional[pulumi.Input[bool]] = None,
            safety_buffer: Optional[pulumi.Input[int]] = None) -> 'AuthBackendIdentityWhitelist':
        """
        Get an existing AuthBackendIdentityWhitelist resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The path of the AWS backend being configured.
        :param pulumi.Input[bool] disable_periodic_tidy: If set to true, disables the periodic
               tidying of the identity-whitelist entries.
        :param pulumi.Input[int] safety_buffer: The amount of extra time, in minutes, that must
               have passed beyond the roletag expiration, before it is removed from the
               backend storage.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backend"] = backend
        __props__["disable_periodic_tidy"] = disable_periodic_tidy
        __props__["safety_buffer"] = safety_buffer
        return AuthBackendIdentityWhitelist(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[Optional[str]]:
        """
        The path of the AWS backend being configured.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="disablePeriodicTidy")
    def disable_periodic_tidy(self) -> pulumi.Output[Optional[bool]]:
        """
        If set to true, disables the periodic
        tidying of the identity-whitelist entries.
        """
        return pulumi.get(self, "disable_periodic_tidy")

    @property
    @pulumi.getter(name="safetyBuffer")
    def safety_buffer(self) -> pulumi.Output[Optional[int]]:
        """
        The amount of extra time, in minutes, that must
        have passed beyond the roletag expiration, before it is removed from the
        backend storage.
        """
        return pulumi.get(self, "safety_buffer")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

