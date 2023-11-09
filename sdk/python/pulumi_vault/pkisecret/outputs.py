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
    'SecretBackendRolePolicyIdentifier',
]

@pulumi.output_type
class SecretBackendRolePolicyIdentifier(dict):
    def __init__(__self__, *,
                 oid: str,
                 cps: Optional[str] = None,
                 notice: Optional[str] = None):
        """
        :param str oid: The OID for the policy identifier
        :param str cps: The URL of the CPS for the policy identifier
               
               Example usage:
        :param str notice: A notice for the policy identifier
        """
        pulumi.set(__self__, "oid", oid)
        if cps is not None:
            pulumi.set(__self__, "cps", cps)
        if notice is not None:
            pulumi.set(__self__, "notice", notice)

    @property
    @pulumi.getter
    def oid(self) -> str:
        """
        The OID for the policy identifier
        """
        return pulumi.get(self, "oid")

    @property
    @pulumi.getter
    def cps(self) -> Optional[str]:
        """
        The URL of the CPS for the policy identifier

        Example usage:
        """
        return pulumi.get(self, "cps")

    @property
    @pulumi.getter
    def notice(self) -> Optional[str]:
        """
        A notice for the policy identifier
        """
        return pulumi.get(self, "notice")

