# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities, _tables

__all__ = [
    'GetNomadAccessTokenResult',
    'AwaitableGetNomadAccessTokenResult',
    'get_nomad_access_token',
]

@pulumi.output_type
class GetNomadAccessTokenResult:
    """
    A collection of values returned by getNomadAccessToken.
    """
    def __init__(__self__, accessor_id=None, backend=None, id=None, role=None, secret_id=None):
        if accessor_id and not isinstance(accessor_id, str):
            raise TypeError("Expected argument 'accessor_id' to be a str")
        pulumi.set(__self__, "accessor_id", accessor_id)
        if backend and not isinstance(backend, str):
            raise TypeError("Expected argument 'backend' to be a str")
        pulumi.set(__self__, "backend", backend)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if role and not isinstance(role, str):
            raise TypeError("Expected argument 'role' to be a str")
        pulumi.set(__self__, "role", role)
        if secret_id and not isinstance(secret_id, str):
            raise TypeError("Expected argument 'secret_id' to be a str")
        pulumi.set(__self__, "secret_id", secret_id)

    @property
    @pulumi.getter(name="accessorId")
    def accessor_id(self) -> str:
        """
        The public identifier for a specific token. It can be used 
        to look up information about a token or to revoke a token.
        """
        return pulumi.get(self, "accessor_id")

    @property
    @pulumi.getter
    def backend(self) -> str:
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def role(self) -> str:
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> str:
        """
        The token to be used when making requests to Nomad and should be kept private.
        """
        return pulumi.get(self, "secret_id")


class AwaitableGetNomadAccessTokenResult(GetNomadAccessTokenResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNomadAccessTokenResult(
            accessor_id=self.accessor_id,
            backend=self.backend,
            id=self.id,
            role=self.role,
            secret_id=self.secret_id)


def get_nomad_access_token(backend: Optional[str] = None,
                           role: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNomadAccessTokenResult:
    """
    Use this data source to access information about an existing resource.

    :param str backend: The path to the Nomad secret backend to
           read credentials from, with no leading or trailing `/`s.
    :param str role: The name of the Nomad secret backend role to generate
           a token for, with no leading or trailing `/`s.
    """
    __args__ = dict()
    __args__['backend'] = backend
    __args__['role'] = role
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vault:index/getNomadAccessToken:getNomadAccessToken', __args__, opts=opts, typ=GetNomadAccessTokenResult).value

    return AwaitableGetNomadAccessTokenResult(
        accessor_id=__ret__.accessor_id,
        backend=__ret__.backend,
        id=__ret__.id,
        role=__ret__.role,
        secret_id=__ret__.secret_id)
