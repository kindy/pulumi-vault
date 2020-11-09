# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'SecretRolesetBindingArgs',
]

@pulumi.input_type
class SecretRolesetBindingArgs:
    def __init__(__self__, *,
                 resource: pulumi.Input[str],
                 roles: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        :param pulumi.Input[str] resource: Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource.
        """
        pulumi.set(__self__, "resource", resource)
        pulumi.set(__self__, "roles", roles)

    @property
    @pulumi.getter
    def resource(self) -> pulumi.Input[str]:
        """
        Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings).
        """
        return pulumi.get(self, "resource")

    @resource.setter
    def resource(self, value: pulumi.Input[str]):
        pulumi.set(self, "resource", value)

    @property
    @pulumi.getter
    def roles(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource.
        """
        return pulumi.get(self, "roles")

    @roles.setter
    def roles(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "roles", value)

