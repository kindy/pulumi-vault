# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from . import _utilities, _tables

__all__ = ['Token']


class Token(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 explicit_max_ttl: Optional[pulumi.Input[str]] = None,
                 no_default_policy: Optional[pulumi.Input[bool]] = None,
                 no_parent: Optional[pulumi.Input[bool]] = None,
                 num_uses: Optional[pulumi.Input[float]] = None,
                 period: Optional[pulumi.Input[str]] = None,
                 pgp_key: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 renew_increment: Optional[pulumi.Input[float]] = None,
                 renew_min_lease: Optional[pulumi.Input[float]] = None,
                 renewable: Optional[pulumi.Input[bool]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 wrapping_ttl: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Token resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] display_name: String containing the token display name
        :param pulumi.Input[str] explicit_max_ttl: The explicit max TTL of this token
        :param pulumi.Input[bool] no_default_policy: Flag to not attach the default policy to this token
        :param pulumi.Input[bool] no_parent: Flag to create a token without parent
        :param pulumi.Input[float] num_uses: The number of allowed uses of this token
        :param pulumi.Input[str] period: The period of this token
        :param pulumi.Input[str] pgp_key: The PGP key (base64 encoded) to encrypt the token.
        :param pulumi.Input[List[pulumi.Input[str]]] policies: List of policies to attach to this token
        :param pulumi.Input[float] renew_increment: The renew increment
        :param pulumi.Input[float] renew_min_lease: The minimal lease to renew this token
        :param pulumi.Input[bool] renewable: Flag to allow to renew this token
        :param pulumi.Input[str] role_name: The token role name
        :param pulumi.Input[str] ttl: The TTL period of this token
        :param pulumi.Input[str] wrapping_ttl: The TTL period of the wrapped token.
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

            __props__['display_name'] = display_name
            __props__['explicit_max_ttl'] = explicit_max_ttl
            __props__['no_default_policy'] = no_default_policy
            __props__['no_parent'] = no_parent
            __props__['num_uses'] = num_uses
            __props__['period'] = period
            __props__['pgp_key'] = pgp_key
            __props__['policies'] = policies
            __props__['renew_increment'] = renew_increment
            __props__['renew_min_lease'] = renew_min_lease
            __props__['renewable'] = renewable
            __props__['role_name'] = role_name
            __props__['ttl'] = ttl
            __props__['wrapping_ttl'] = wrapping_ttl
            __props__['client_token'] = None
            __props__['encrypted_client_token'] = None
            __props__['lease_duration'] = None
            __props__['lease_started'] = None
            __props__['wrapped_token'] = None
            __props__['wrapping_accessor'] = None
        super(Token, __self__).__init__(
            'vault:index/token:Token',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            client_token: Optional[pulumi.Input[str]] = None,
            display_name: Optional[pulumi.Input[str]] = None,
            encrypted_client_token: Optional[pulumi.Input[str]] = None,
            explicit_max_ttl: Optional[pulumi.Input[str]] = None,
            lease_duration: Optional[pulumi.Input[float]] = None,
            lease_started: Optional[pulumi.Input[str]] = None,
            no_default_policy: Optional[pulumi.Input[bool]] = None,
            no_parent: Optional[pulumi.Input[bool]] = None,
            num_uses: Optional[pulumi.Input[float]] = None,
            period: Optional[pulumi.Input[str]] = None,
            pgp_key: Optional[pulumi.Input[str]] = None,
            policies: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            renew_increment: Optional[pulumi.Input[float]] = None,
            renew_min_lease: Optional[pulumi.Input[float]] = None,
            renewable: Optional[pulumi.Input[bool]] = None,
            role_name: Optional[pulumi.Input[str]] = None,
            ttl: Optional[pulumi.Input[str]] = None,
            wrapped_token: Optional[pulumi.Input[str]] = None,
            wrapping_accessor: Optional[pulumi.Input[str]] = None,
            wrapping_ttl: Optional[pulumi.Input[str]] = None) -> 'Token':
        """
        Get an existing Token resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_token: String containing the client token if stored in present file
        :param pulumi.Input[str] display_name: String containing the token display name
        :param pulumi.Input[str] encrypted_client_token: String containing the client token encrypted with the given `pgp_key` if stored in present file
        :param pulumi.Input[str] explicit_max_ttl: The explicit max TTL of this token
        :param pulumi.Input[float] lease_duration: String containing the token lease duration if present in state file
        :param pulumi.Input[str] lease_started: String containing the token lease started time if present in state file
        :param pulumi.Input[bool] no_default_policy: Flag to not attach the default policy to this token
        :param pulumi.Input[bool] no_parent: Flag to create a token without parent
        :param pulumi.Input[float] num_uses: The number of allowed uses of this token
        :param pulumi.Input[str] period: The period of this token
        :param pulumi.Input[str] pgp_key: The PGP key (base64 encoded) to encrypt the token.
        :param pulumi.Input[List[pulumi.Input[str]]] policies: List of policies to attach to this token
        :param pulumi.Input[float] renew_increment: The renew increment
        :param pulumi.Input[float] renew_min_lease: The minimal lease to renew this token
        :param pulumi.Input[bool] renewable: Flag to allow to renew this token
        :param pulumi.Input[str] role_name: The token role name
        :param pulumi.Input[str] ttl: The TTL period of this token
        :param pulumi.Input[str] wrapped_token: The client wrapped token.
        :param pulumi.Input[str] wrapping_accessor: The client wrapping accessor.
        :param pulumi.Input[str] wrapping_ttl: The TTL period of the wrapped token.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["client_token"] = client_token
        __props__["display_name"] = display_name
        __props__["encrypted_client_token"] = encrypted_client_token
        __props__["explicit_max_ttl"] = explicit_max_ttl
        __props__["lease_duration"] = lease_duration
        __props__["lease_started"] = lease_started
        __props__["no_default_policy"] = no_default_policy
        __props__["no_parent"] = no_parent
        __props__["num_uses"] = num_uses
        __props__["period"] = period
        __props__["pgp_key"] = pgp_key
        __props__["policies"] = policies
        __props__["renew_increment"] = renew_increment
        __props__["renew_min_lease"] = renew_min_lease
        __props__["renewable"] = renewable
        __props__["role_name"] = role_name
        __props__["ttl"] = ttl
        __props__["wrapped_token"] = wrapped_token
        __props__["wrapping_accessor"] = wrapping_accessor
        __props__["wrapping_ttl"] = wrapping_ttl
        return Token(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientToken")
    def client_token(self) -> pulumi.Output[str]:
        """
        String containing the client token if stored in present file
        """
        return pulumi.get(self, "client_token")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[Optional[str]]:
        """
        String containing the token display name
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="encryptedClientToken")
    def encrypted_client_token(self) -> pulumi.Output[str]:
        """
        String containing the client token encrypted with the given `pgp_key` if stored in present file
        """
        return pulumi.get(self, "encrypted_client_token")

    @property
    @pulumi.getter(name="explicitMaxTtl")
    def explicit_max_ttl(self) -> pulumi.Output[Optional[str]]:
        """
        The explicit max TTL of this token
        """
        return pulumi.get(self, "explicit_max_ttl")

    @property
    @pulumi.getter(name="leaseDuration")
    def lease_duration(self) -> pulumi.Output[float]:
        """
        String containing the token lease duration if present in state file
        """
        return pulumi.get(self, "lease_duration")

    @property
    @pulumi.getter(name="leaseStarted")
    def lease_started(self) -> pulumi.Output[str]:
        """
        String containing the token lease started time if present in state file
        """
        return pulumi.get(self, "lease_started")

    @property
    @pulumi.getter(name="noDefaultPolicy")
    def no_default_policy(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to not attach the default policy to this token
        """
        return pulumi.get(self, "no_default_policy")

    @property
    @pulumi.getter(name="noParent")
    def no_parent(self) -> pulumi.Output[bool]:
        """
        Flag to create a token without parent
        """
        return pulumi.get(self, "no_parent")

    @property
    @pulumi.getter(name="numUses")
    def num_uses(self) -> pulumi.Output[float]:
        """
        The number of allowed uses of this token
        """
        return pulumi.get(self, "num_uses")

    @property
    @pulumi.getter
    def period(self) -> pulumi.Output[Optional[str]]:
        """
        The period of this token
        """
        return pulumi.get(self, "period")

    @property
    @pulumi.getter(name="pgpKey")
    def pgp_key(self) -> pulumi.Output[Optional[str]]:
        """
        The PGP key (base64 encoded) to encrypt the token.
        """
        return pulumi.get(self, "pgp_key")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Optional[List[str]]]:
        """
        List of policies to attach to this token
        """
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter(name="renewIncrement")
    def renew_increment(self) -> pulumi.Output[Optional[float]]:
        """
        The renew increment
        """
        return pulumi.get(self, "renew_increment")

    @property
    @pulumi.getter(name="renewMinLease")
    def renew_min_lease(self) -> pulumi.Output[Optional[float]]:
        """
        The minimal lease to renew this token
        """
        return pulumi.get(self, "renew_min_lease")

    @property
    @pulumi.getter
    def renewable(self) -> pulumi.Output[bool]:
        """
        Flag to allow to renew this token
        """
        return pulumi.get(self, "renewable")

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> pulumi.Output[Optional[str]]:
        """
        The token role name
        """
        return pulumi.get(self, "role_name")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[Optional[str]]:
        """
        The TTL period of this token
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter(name="wrappedToken")
    def wrapped_token(self) -> pulumi.Output[str]:
        """
        The client wrapped token.
        """
        return pulumi.get(self, "wrapped_token")

    @property
    @pulumi.getter(name="wrappingAccessor")
    def wrapping_accessor(self) -> pulumi.Output[str]:
        """
        The client wrapping accessor.
        """
        return pulumi.get(self, "wrapping_accessor")

    @property
    @pulumi.getter(name="wrappingTtl")
    def wrapping_ttl(self) -> pulumi.Output[Optional[str]]:
        """
        The TTL period of the wrapped token.
        """
        return pulumi.get(self, "wrapping_ttl")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

