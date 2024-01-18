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
    'AuthBackendTuneArgs',
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
class AuthBackendTuneArgs:
    def __init__(__self__, *,
                 allowed_response_headers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 audit_non_hmac_request_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 audit_non_hmac_response_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 default_lease_ttl: Optional[pulumi.Input[str]] = None,
                 listing_visibility: Optional[pulumi.Input[str]] = None,
                 max_lease_ttl: Optional[pulumi.Input[str]] = None,
                 passthrough_request_headers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 token_type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_response_headers: List of headers to whitelist and allowing
               a plugin to include them in the response.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] audit_non_hmac_request_keys: Specifies the list of keys that will
               not be HMAC'd by audit devices in the request data object.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] audit_non_hmac_response_keys: Specifies the list of keys that will
               not be HMAC'd by audit devices in the response data object.
        :param pulumi.Input[str] default_lease_ttl: Specifies the default time-to-live.
               If set, this overrides the global default.
               Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
        :param pulumi.Input[str] listing_visibility: Specifies whether to show this mount in
               the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
        :param pulumi.Input[str] max_lease_ttl: Specifies the maximum time-to-live.
               If set, this overrides the global default.
               Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
        :param pulumi.Input[Sequence[pulumi.Input[str]]] passthrough_request_headers: List of headers to whitelist and
               pass from the request to the backend.
        :param pulumi.Input[str] token_type: Specifies the type of tokens that should be returned by
               the mount. Valid values are "default-service", "default-batch", "service", "batch".
               
               
               For more details on the usage of each argument consult the [Vault GCP API documentation](https://www.vaultproject.io/api-docs/auth/gcp#configure).
        """
        if allowed_response_headers is not None:
            pulumi.set(__self__, "allowed_response_headers", allowed_response_headers)
        if audit_non_hmac_request_keys is not None:
            pulumi.set(__self__, "audit_non_hmac_request_keys", audit_non_hmac_request_keys)
        if audit_non_hmac_response_keys is not None:
            pulumi.set(__self__, "audit_non_hmac_response_keys", audit_non_hmac_response_keys)
        if default_lease_ttl is not None:
            pulumi.set(__self__, "default_lease_ttl", default_lease_ttl)
        if listing_visibility is not None:
            pulumi.set(__self__, "listing_visibility", listing_visibility)
        if max_lease_ttl is not None:
            pulumi.set(__self__, "max_lease_ttl", max_lease_ttl)
        if passthrough_request_headers is not None:
            pulumi.set(__self__, "passthrough_request_headers", passthrough_request_headers)
        if token_type is not None:
            pulumi.set(__self__, "token_type", token_type)

    @property
    @pulumi.getter(name="allowedResponseHeaders")
    def allowed_response_headers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of headers to whitelist and allowing
        a plugin to include them in the response.
        """
        return pulumi.get(self, "allowed_response_headers")

    @allowed_response_headers.setter
    def allowed_response_headers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_response_headers", value)

    @property
    @pulumi.getter(name="auditNonHmacRequestKeys")
    def audit_non_hmac_request_keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the list of keys that will
        not be HMAC'd by audit devices in the request data object.
        """
        return pulumi.get(self, "audit_non_hmac_request_keys")

    @audit_non_hmac_request_keys.setter
    def audit_non_hmac_request_keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "audit_non_hmac_request_keys", value)

    @property
    @pulumi.getter(name="auditNonHmacResponseKeys")
    def audit_non_hmac_response_keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the list of keys that will
        not be HMAC'd by audit devices in the response data object.
        """
        return pulumi.get(self, "audit_non_hmac_response_keys")

    @audit_non_hmac_response_keys.setter
    def audit_non_hmac_response_keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "audit_non_hmac_response_keys", value)

    @property
    @pulumi.getter(name="defaultLeaseTtl")
    def default_lease_ttl(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the default time-to-live.
        If set, this overrides the global default.
        Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
        """
        return pulumi.get(self, "default_lease_ttl")

    @default_lease_ttl.setter
    def default_lease_ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_lease_ttl", value)

    @property
    @pulumi.getter(name="listingVisibility")
    def listing_visibility(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether to show this mount in
        the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
        """
        return pulumi.get(self, "listing_visibility")

    @listing_visibility.setter
    def listing_visibility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "listing_visibility", value)

    @property
    @pulumi.getter(name="maxLeaseTtl")
    def max_lease_ttl(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the maximum time-to-live.
        If set, this overrides the global default.
        Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
        """
        return pulumi.get(self, "max_lease_ttl")

    @max_lease_ttl.setter
    def max_lease_ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "max_lease_ttl", value)

    @property
    @pulumi.getter(name="passthroughRequestHeaders")
    def passthrough_request_headers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of headers to whitelist and
        pass from the request to the backend.
        """
        return pulumi.get(self, "passthrough_request_headers")

    @passthrough_request_headers.setter
    def passthrough_request_headers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "passthrough_request_headers", value)

    @property
    @pulumi.getter(name="tokenType")
    def token_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the type of tokens that should be returned by
        the mount. Valid values are "default-service", "default-batch", "service", "batch".


        For more details on the usage of each argument consult the [Vault GCP API documentation](https://www.vaultproject.io/api-docs/auth/gcp#configure).
        """
        return pulumi.get(self, "token_type")

    @token_type.setter
    def token_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token_type", value)


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


