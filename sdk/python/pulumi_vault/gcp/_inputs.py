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
    'AuthBackendCustomEndpointArgs',
    'SecretRolesetBindingArgs',
    'SecretStaticAccountBindingArgs',
]

@pulumi.input_type
class AuthBackendCustomEndpointArgs:
    def __init__(__self__, *,
                 api: Optional[pulumi.Input[str]] = None,
                 compute: Optional[pulumi.Input[str]] = None,
                 crm: Optional[pulumi.Input[str]] = None,
                 iam: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] api: Replaces the service endpoint used in API requests to `https://www.googleapis.com`.
        :param pulumi.Input[str] compute: Replaces the service endpoint used in API requests to `https://compute.googleapis.com`.
               
               The endpoint value provided for a given key has the form of `scheme://host:port`.
               The `scheme://` and `:port` portions of the endpoint value are optional.
               
               For more details on the usage of each argument consult the [Vault GCP API documentation](https://www.vaultproject.io/api-docs/auth/gcp#configure).
        :param pulumi.Input[str] crm: Replaces the service endpoint used in API requests to `https://cloudresourcemanager.googleapis.com`.
        :param pulumi.Input[str] iam: Replaces the service endpoint used in API requests to `https://iam.googleapis.com`.
        """
        if api is not None:
            pulumi.set(__self__, "api", api)
        if compute is not None:
            pulumi.set(__self__, "compute", compute)
        if crm is not None:
            pulumi.set(__self__, "crm", crm)
        if iam is not None:
            pulumi.set(__self__, "iam", iam)

    @property
    @pulumi.getter
    def api(self) -> Optional[pulumi.Input[str]]:
        """
        Replaces the service endpoint used in API requests to `https://www.googleapis.com`.
        """
        return pulumi.get(self, "api")

    @api.setter
    def api(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api", value)

    @property
    @pulumi.getter
    def compute(self) -> Optional[pulumi.Input[str]]:
        """
        Replaces the service endpoint used in API requests to `https://compute.googleapis.com`.

        The endpoint value provided for a given key has the form of `scheme://host:port`.
        The `scheme://` and `:port` portions of the endpoint value are optional.

        For more details on the usage of each argument consult the [Vault GCP API documentation](https://www.vaultproject.io/api-docs/auth/gcp#configure).
        """
        return pulumi.get(self, "compute")

    @compute.setter
    def compute(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "compute", value)

    @property
    @pulumi.getter
    def crm(self) -> Optional[pulumi.Input[str]]:
        """
        Replaces the service endpoint used in API requests to `https://cloudresourcemanager.googleapis.com`.
        """
        return pulumi.get(self, "crm")

    @crm.setter
    def crm(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "crm", value)

    @property
    @pulumi.getter
    def iam(self) -> Optional[pulumi.Input[str]]:
        """
        Replaces the service endpoint used in API requests to `https://iam.googleapis.com`.
        """
        return pulumi.get(self, "iam")

    @iam.setter
    def iam(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iam", value)


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


@pulumi.input_type
class SecretStaticAccountBindingArgs:
    def __init__(__self__, *,
                 resource: pulumi.Input[str],
                 roles: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        :param pulumi.Input[str] resource: Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#bindings).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource.
        """
        pulumi.set(__self__, "resource", resource)
        pulumi.set(__self__, "roles", roles)

    @property
    @pulumi.getter
    def resource(self) -> pulumi.Input[str]:
        """
        Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#bindings).
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


