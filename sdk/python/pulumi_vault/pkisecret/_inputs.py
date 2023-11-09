# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'SecretBackendRolePolicyIdentifierArgs',
]

@pulumi.input_type
class SecretBackendRolePolicyIdentifierArgs:
    def __init__(__self__, *,
                 oid: pulumi.Input[str],
                 cps: Optional[pulumi.Input[str]] = None,
                 notice: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] oid: The OID for the policy identifier
        :param pulumi.Input[str] cps: The URL of the CPS for the policy identifier
               
               Example usage:
        :param pulumi.Input[str] notice: A notice for the policy identifier
        """
        pulumi.set(__self__, "oid", oid)
        if cps is not None:
            pulumi.set(__self__, "cps", cps)
        if notice is not None:
            pulumi.set(__self__, "notice", notice)

    @property
    @pulumi.getter
    def oid(self) -> pulumi.Input[str]:
        """
        The OID for the policy identifier
        """
        return pulumi.get(self, "oid")

    @oid.setter
    def oid(self, value: pulumi.Input[str]):
        pulumi.set(self, "oid", value)

    @property
    @pulumi.getter
    def cps(self) -> Optional[pulumi.Input[str]]:
        """
        The URL of the CPS for the policy identifier

        Example usage:
        """
        return pulumi.get(self, "cps")

    @cps.setter
    def cps(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cps", value)

    @property
    @pulumi.getter
    def notice(self) -> Optional[pulumi.Input[str]]:
        """
        A notice for the policy identifier
        """
        return pulumi.get(self, "notice")

    @notice.setter
    def notice(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "notice", value)


