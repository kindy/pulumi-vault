# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SecretBackendCrlConfigArgs', 'SecretBackendCrlConfig']

@pulumi.input_type
class SecretBackendCrlConfigArgs:
    def __init__(__self__, *,
                 backend: pulumi.Input[str],
                 auto_rebuild: Optional[pulumi.Input[bool]] = None,
                 auto_rebuild_grace_period: Optional[pulumi.Input[str]] = None,
                 cross_cluster_revocation: Optional[pulumi.Input[bool]] = None,
                 delta_rebuild_interval: Optional[pulumi.Input[str]] = None,
                 disable: Optional[pulumi.Input[bool]] = None,
                 enable_delta: Optional[pulumi.Input[bool]] = None,
                 expiry: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 ocsp_disable: Optional[pulumi.Input[bool]] = None,
                 ocsp_expiry: Optional[pulumi.Input[str]] = None,
                 unified_crl: Optional[pulumi.Input[bool]] = None,
                 unified_crl_on_existing_paths: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a SecretBackendCrlConfig resource.
        :param pulumi.Input[str] backend: The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        :param pulumi.Input[bool] auto_rebuild: Enables periodic rebuilding of the CRL upon expiry. **Vault 1.12+**
        :param pulumi.Input[str] auto_rebuild_grace_period: Grace period before CRL expiry to attempt rebuild of CRL. **Vault 1.12+**
        :param pulumi.Input[bool] cross_cluster_revocation: Enable cross-cluster revocation request queues. **Vault 1.13+**
        :param pulumi.Input[str] delta_rebuild_interval: Interval to check for new revocations on, to regenerate the delta CRL.
        :param pulumi.Input[bool] disable: Disables or enables CRL building.
        :param pulumi.Input[bool] enable_delta: Enables building of delta CRLs with up-to-date revocation information, 
               augmenting the last complete CRL.  **Vault 1.12+**
        :param pulumi.Input[str] expiry: Specifies the time until expiration.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[bool] ocsp_disable: Disables the OCSP responder in Vault. **Vault 1.12+**
        :param pulumi.Input[str] ocsp_expiry: The amount of time an OCSP response can be cached for, useful for OCSP stapling 
               refresh durations. **Vault 1.12+**
        :param pulumi.Input[bool] unified_crl: Enables unified CRL and OCSP building. **Vault 1.13+**
        :param pulumi.Input[bool] unified_crl_on_existing_paths: Enables serving the unified CRL and OCSP on the existing, previously
               cluster-local paths. **Vault 1.13+**
        """
        pulumi.set(__self__, "backend", backend)
        if auto_rebuild is not None:
            pulumi.set(__self__, "auto_rebuild", auto_rebuild)
        if auto_rebuild_grace_period is not None:
            pulumi.set(__self__, "auto_rebuild_grace_period", auto_rebuild_grace_period)
        if cross_cluster_revocation is not None:
            pulumi.set(__self__, "cross_cluster_revocation", cross_cluster_revocation)
        if delta_rebuild_interval is not None:
            pulumi.set(__self__, "delta_rebuild_interval", delta_rebuild_interval)
        if disable is not None:
            pulumi.set(__self__, "disable", disable)
        if enable_delta is not None:
            pulumi.set(__self__, "enable_delta", enable_delta)
        if expiry is not None:
            pulumi.set(__self__, "expiry", expiry)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if ocsp_disable is not None:
            pulumi.set(__self__, "ocsp_disable", ocsp_disable)
        if ocsp_expiry is not None:
            pulumi.set(__self__, "ocsp_expiry", ocsp_expiry)
        if unified_crl is not None:
            pulumi.set(__self__, "unified_crl", unified_crl)
        if unified_crl_on_existing_paths is not None:
            pulumi.set(__self__, "unified_crl_on_existing_paths", unified_crl_on_existing_paths)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Input[str]:
        """
        The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: pulumi.Input[str]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="autoRebuild")
    def auto_rebuild(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables periodic rebuilding of the CRL upon expiry. **Vault 1.12+**
        """
        return pulumi.get(self, "auto_rebuild")

    @auto_rebuild.setter
    def auto_rebuild(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_rebuild", value)

    @property
    @pulumi.getter(name="autoRebuildGracePeriod")
    def auto_rebuild_grace_period(self) -> Optional[pulumi.Input[str]]:
        """
        Grace period before CRL expiry to attempt rebuild of CRL. **Vault 1.12+**
        """
        return pulumi.get(self, "auto_rebuild_grace_period")

    @auto_rebuild_grace_period.setter
    def auto_rebuild_grace_period(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_rebuild_grace_period", value)

    @property
    @pulumi.getter(name="crossClusterRevocation")
    def cross_cluster_revocation(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable cross-cluster revocation request queues. **Vault 1.13+**
        """
        return pulumi.get(self, "cross_cluster_revocation")

    @cross_cluster_revocation.setter
    def cross_cluster_revocation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "cross_cluster_revocation", value)

    @property
    @pulumi.getter(name="deltaRebuildInterval")
    def delta_rebuild_interval(self) -> Optional[pulumi.Input[str]]:
        """
        Interval to check for new revocations on, to regenerate the delta CRL.
        """
        return pulumi.get(self, "delta_rebuild_interval")

    @delta_rebuild_interval.setter
    def delta_rebuild_interval(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delta_rebuild_interval", value)

    @property
    @pulumi.getter
    def disable(self) -> Optional[pulumi.Input[bool]]:
        """
        Disables or enables CRL building.
        """
        return pulumi.get(self, "disable")

    @disable.setter
    def disable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable", value)

    @property
    @pulumi.getter(name="enableDelta")
    def enable_delta(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables building of delta CRLs with up-to-date revocation information, 
        augmenting the last complete CRL.  **Vault 1.12+**
        """
        return pulumi.get(self, "enable_delta")

    @enable_delta.setter
    def enable_delta(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_delta", value)

    @property
    @pulumi.getter
    def expiry(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the time until expiration.
        """
        return pulumi.get(self, "expiry")

    @expiry.setter
    def expiry(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiry", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="ocspDisable")
    def ocsp_disable(self) -> Optional[pulumi.Input[bool]]:
        """
        Disables the OCSP responder in Vault. **Vault 1.12+**
        """
        return pulumi.get(self, "ocsp_disable")

    @ocsp_disable.setter
    def ocsp_disable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ocsp_disable", value)

    @property
    @pulumi.getter(name="ocspExpiry")
    def ocsp_expiry(self) -> Optional[pulumi.Input[str]]:
        """
        The amount of time an OCSP response can be cached for, useful for OCSP stapling 
        refresh durations. **Vault 1.12+**
        """
        return pulumi.get(self, "ocsp_expiry")

    @ocsp_expiry.setter
    def ocsp_expiry(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ocsp_expiry", value)

    @property
    @pulumi.getter(name="unifiedCrl")
    def unified_crl(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables unified CRL and OCSP building. **Vault 1.13+**
        """
        return pulumi.get(self, "unified_crl")

    @unified_crl.setter
    def unified_crl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "unified_crl", value)

    @property
    @pulumi.getter(name="unifiedCrlOnExistingPaths")
    def unified_crl_on_existing_paths(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables serving the unified CRL and OCSP on the existing, previously
        cluster-local paths. **Vault 1.13+**
        """
        return pulumi.get(self, "unified_crl_on_existing_paths")

    @unified_crl_on_existing_paths.setter
    def unified_crl_on_existing_paths(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "unified_crl_on_existing_paths", value)


@pulumi.input_type
class _SecretBackendCrlConfigState:
    def __init__(__self__, *,
                 auto_rebuild: Optional[pulumi.Input[bool]] = None,
                 auto_rebuild_grace_period: Optional[pulumi.Input[str]] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 cross_cluster_revocation: Optional[pulumi.Input[bool]] = None,
                 delta_rebuild_interval: Optional[pulumi.Input[str]] = None,
                 disable: Optional[pulumi.Input[bool]] = None,
                 enable_delta: Optional[pulumi.Input[bool]] = None,
                 expiry: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 ocsp_disable: Optional[pulumi.Input[bool]] = None,
                 ocsp_expiry: Optional[pulumi.Input[str]] = None,
                 unified_crl: Optional[pulumi.Input[bool]] = None,
                 unified_crl_on_existing_paths: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering SecretBackendCrlConfig resources.
        :param pulumi.Input[bool] auto_rebuild: Enables periodic rebuilding of the CRL upon expiry. **Vault 1.12+**
        :param pulumi.Input[str] auto_rebuild_grace_period: Grace period before CRL expiry to attempt rebuild of CRL. **Vault 1.12+**
        :param pulumi.Input[str] backend: The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        :param pulumi.Input[bool] cross_cluster_revocation: Enable cross-cluster revocation request queues. **Vault 1.13+**
        :param pulumi.Input[str] delta_rebuild_interval: Interval to check for new revocations on, to regenerate the delta CRL.
        :param pulumi.Input[bool] disable: Disables or enables CRL building.
        :param pulumi.Input[bool] enable_delta: Enables building of delta CRLs with up-to-date revocation information, 
               augmenting the last complete CRL.  **Vault 1.12+**
        :param pulumi.Input[str] expiry: Specifies the time until expiration.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[bool] ocsp_disable: Disables the OCSP responder in Vault. **Vault 1.12+**
        :param pulumi.Input[str] ocsp_expiry: The amount of time an OCSP response can be cached for, useful for OCSP stapling 
               refresh durations. **Vault 1.12+**
        :param pulumi.Input[bool] unified_crl: Enables unified CRL and OCSP building. **Vault 1.13+**
        :param pulumi.Input[bool] unified_crl_on_existing_paths: Enables serving the unified CRL and OCSP on the existing, previously
               cluster-local paths. **Vault 1.13+**
        """
        if auto_rebuild is not None:
            pulumi.set(__self__, "auto_rebuild", auto_rebuild)
        if auto_rebuild_grace_period is not None:
            pulumi.set(__self__, "auto_rebuild_grace_period", auto_rebuild_grace_period)
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if cross_cluster_revocation is not None:
            pulumi.set(__self__, "cross_cluster_revocation", cross_cluster_revocation)
        if delta_rebuild_interval is not None:
            pulumi.set(__self__, "delta_rebuild_interval", delta_rebuild_interval)
        if disable is not None:
            pulumi.set(__self__, "disable", disable)
        if enable_delta is not None:
            pulumi.set(__self__, "enable_delta", enable_delta)
        if expiry is not None:
            pulumi.set(__self__, "expiry", expiry)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if ocsp_disable is not None:
            pulumi.set(__self__, "ocsp_disable", ocsp_disable)
        if ocsp_expiry is not None:
            pulumi.set(__self__, "ocsp_expiry", ocsp_expiry)
        if unified_crl is not None:
            pulumi.set(__self__, "unified_crl", unified_crl)
        if unified_crl_on_existing_paths is not None:
            pulumi.set(__self__, "unified_crl_on_existing_paths", unified_crl_on_existing_paths)

    @property
    @pulumi.getter(name="autoRebuild")
    def auto_rebuild(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables periodic rebuilding of the CRL upon expiry. **Vault 1.12+**
        """
        return pulumi.get(self, "auto_rebuild")

    @auto_rebuild.setter
    def auto_rebuild(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_rebuild", value)

    @property
    @pulumi.getter(name="autoRebuildGracePeriod")
    def auto_rebuild_grace_period(self) -> Optional[pulumi.Input[str]]:
        """
        Grace period before CRL expiry to attempt rebuild of CRL. **Vault 1.12+**
        """
        return pulumi.get(self, "auto_rebuild_grace_period")

    @auto_rebuild_grace_period.setter
    def auto_rebuild_grace_period(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "auto_rebuild_grace_period", value)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[str]]:
        """
        The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="crossClusterRevocation")
    def cross_cluster_revocation(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable cross-cluster revocation request queues. **Vault 1.13+**
        """
        return pulumi.get(self, "cross_cluster_revocation")

    @cross_cluster_revocation.setter
    def cross_cluster_revocation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "cross_cluster_revocation", value)

    @property
    @pulumi.getter(name="deltaRebuildInterval")
    def delta_rebuild_interval(self) -> Optional[pulumi.Input[str]]:
        """
        Interval to check for new revocations on, to regenerate the delta CRL.
        """
        return pulumi.get(self, "delta_rebuild_interval")

    @delta_rebuild_interval.setter
    def delta_rebuild_interval(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "delta_rebuild_interval", value)

    @property
    @pulumi.getter
    def disable(self) -> Optional[pulumi.Input[bool]]:
        """
        Disables or enables CRL building.
        """
        return pulumi.get(self, "disable")

    @disable.setter
    def disable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable", value)

    @property
    @pulumi.getter(name="enableDelta")
    def enable_delta(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables building of delta CRLs with up-to-date revocation information, 
        augmenting the last complete CRL.  **Vault 1.12+**
        """
        return pulumi.get(self, "enable_delta")

    @enable_delta.setter
    def enable_delta(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_delta", value)

    @property
    @pulumi.getter
    def expiry(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the time until expiration.
        """
        return pulumi.get(self, "expiry")

    @expiry.setter
    def expiry(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "expiry", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="ocspDisable")
    def ocsp_disable(self) -> Optional[pulumi.Input[bool]]:
        """
        Disables the OCSP responder in Vault. **Vault 1.12+**
        """
        return pulumi.get(self, "ocsp_disable")

    @ocsp_disable.setter
    def ocsp_disable(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ocsp_disable", value)

    @property
    @pulumi.getter(name="ocspExpiry")
    def ocsp_expiry(self) -> Optional[pulumi.Input[str]]:
        """
        The amount of time an OCSP response can be cached for, useful for OCSP stapling 
        refresh durations. **Vault 1.12+**
        """
        return pulumi.get(self, "ocsp_expiry")

    @ocsp_expiry.setter
    def ocsp_expiry(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ocsp_expiry", value)

    @property
    @pulumi.getter(name="unifiedCrl")
    def unified_crl(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables unified CRL and OCSP building. **Vault 1.13+**
        """
        return pulumi.get(self, "unified_crl")

    @unified_crl.setter
    def unified_crl(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "unified_crl", value)

    @property
    @pulumi.getter(name="unifiedCrlOnExistingPaths")
    def unified_crl_on_existing_paths(self) -> Optional[pulumi.Input[bool]]:
        """
        Enables serving the unified CRL and OCSP on the existing, previously
        cluster-local paths. **Vault 1.13+**
        """
        return pulumi.get(self, "unified_crl_on_existing_paths")

    @unified_crl_on_existing_paths.setter
    def unified_crl_on_existing_paths(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "unified_crl_on_existing_paths", value)


class SecretBackendCrlConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_rebuild: Optional[pulumi.Input[bool]] = None,
                 auto_rebuild_grace_period: Optional[pulumi.Input[str]] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 cross_cluster_revocation: Optional[pulumi.Input[bool]] = None,
                 delta_rebuild_interval: Optional[pulumi.Input[str]] = None,
                 disable: Optional[pulumi.Input[bool]] = None,
                 enable_delta: Optional[pulumi.Input[bool]] = None,
                 expiry: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 ocsp_disable: Optional[pulumi.Input[bool]] = None,
                 ocsp_expiry: Optional[pulumi.Input[str]] = None,
                 unified_crl: Optional[pulumi.Input[bool]] = None,
                 unified_crl_on_existing_paths: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Allows setting the duration for which the generated CRL should be marked valid. If the CRL is disabled, it will return a signed but zero-length CRL for any request. If enabled, it will re-build the CRL.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        pki = vault.Mount("pki",
            path="%s",
            type="pki",
            default_lease_ttl_seconds=3600,
            max_lease_ttl_seconds=86400)
        crl_config = vault.pki_secret.SecretBackendCrlConfig("crlConfig",
            backend=pki.path,
            expiry="72h",
            disable=False)
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_rebuild: Enables periodic rebuilding of the CRL upon expiry. **Vault 1.12+**
        :param pulumi.Input[str] auto_rebuild_grace_period: Grace period before CRL expiry to attempt rebuild of CRL. **Vault 1.12+**
        :param pulumi.Input[str] backend: The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        :param pulumi.Input[bool] cross_cluster_revocation: Enable cross-cluster revocation request queues. **Vault 1.13+**
        :param pulumi.Input[str] delta_rebuild_interval: Interval to check for new revocations on, to regenerate the delta CRL.
        :param pulumi.Input[bool] disable: Disables or enables CRL building.
        :param pulumi.Input[bool] enable_delta: Enables building of delta CRLs with up-to-date revocation information, 
               augmenting the last complete CRL.  **Vault 1.12+**
        :param pulumi.Input[str] expiry: Specifies the time until expiration.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[bool] ocsp_disable: Disables the OCSP responder in Vault. **Vault 1.12+**
        :param pulumi.Input[str] ocsp_expiry: The amount of time an OCSP response can be cached for, useful for OCSP stapling 
               refresh durations. **Vault 1.12+**
        :param pulumi.Input[bool] unified_crl: Enables unified CRL and OCSP building. **Vault 1.13+**
        :param pulumi.Input[bool] unified_crl_on_existing_paths: Enables serving the unified CRL and OCSP on the existing, previously
               cluster-local paths. **Vault 1.13+**
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretBackendCrlConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows setting the duration for which the generated CRL should be marked valid. If the CRL is disabled, it will return a signed but zero-length CRL for any request. If enabled, it will re-build the CRL.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        pki = vault.Mount("pki",
            path="%s",
            type="pki",
            default_lease_ttl_seconds=3600,
            max_lease_ttl_seconds=86400)
        crl_config = vault.pki_secret.SecretBackendCrlConfig("crlConfig",
            backend=pki.path,
            expiry="72h",
            disable=False)
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param SecretBackendCrlConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretBackendCrlConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auto_rebuild: Optional[pulumi.Input[bool]] = None,
                 auto_rebuild_grace_period: Optional[pulumi.Input[str]] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 cross_cluster_revocation: Optional[pulumi.Input[bool]] = None,
                 delta_rebuild_interval: Optional[pulumi.Input[str]] = None,
                 disable: Optional[pulumi.Input[bool]] = None,
                 enable_delta: Optional[pulumi.Input[bool]] = None,
                 expiry: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 ocsp_disable: Optional[pulumi.Input[bool]] = None,
                 ocsp_expiry: Optional[pulumi.Input[str]] = None,
                 unified_crl: Optional[pulumi.Input[bool]] = None,
                 unified_crl_on_existing_paths: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretBackendCrlConfigArgs.__new__(SecretBackendCrlConfigArgs)

            __props__.__dict__["auto_rebuild"] = auto_rebuild
            __props__.__dict__["auto_rebuild_grace_period"] = auto_rebuild_grace_period
            if backend is None and not opts.urn:
                raise TypeError("Missing required property 'backend'")
            __props__.__dict__["backend"] = backend
            __props__.__dict__["cross_cluster_revocation"] = cross_cluster_revocation
            __props__.__dict__["delta_rebuild_interval"] = delta_rebuild_interval
            __props__.__dict__["disable"] = disable
            __props__.__dict__["enable_delta"] = enable_delta
            __props__.__dict__["expiry"] = expiry
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["ocsp_disable"] = ocsp_disable
            __props__.__dict__["ocsp_expiry"] = ocsp_expiry
            __props__.__dict__["unified_crl"] = unified_crl
            __props__.__dict__["unified_crl_on_existing_paths"] = unified_crl_on_existing_paths
        super(SecretBackendCrlConfig, __self__).__init__(
            'vault:pkiSecret/secretBackendCrlConfig:SecretBackendCrlConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auto_rebuild: Optional[pulumi.Input[bool]] = None,
            auto_rebuild_grace_period: Optional[pulumi.Input[str]] = None,
            backend: Optional[pulumi.Input[str]] = None,
            cross_cluster_revocation: Optional[pulumi.Input[bool]] = None,
            delta_rebuild_interval: Optional[pulumi.Input[str]] = None,
            disable: Optional[pulumi.Input[bool]] = None,
            enable_delta: Optional[pulumi.Input[bool]] = None,
            expiry: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            ocsp_disable: Optional[pulumi.Input[bool]] = None,
            ocsp_expiry: Optional[pulumi.Input[str]] = None,
            unified_crl: Optional[pulumi.Input[bool]] = None,
            unified_crl_on_existing_paths: Optional[pulumi.Input[bool]] = None) -> 'SecretBackendCrlConfig':
        """
        Get an existing SecretBackendCrlConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] auto_rebuild: Enables periodic rebuilding of the CRL upon expiry. **Vault 1.12+**
        :param pulumi.Input[str] auto_rebuild_grace_period: Grace period before CRL expiry to attempt rebuild of CRL. **Vault 1.12+**
        :param pulumi.Input[str] backend: The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        :param pulumi.Input[bool] cross_cluster_revocation: Enable cross-cluster revocation request queues. **Vault 1.13+**
        :param pulumi.Input[str] delta_rebuild_interval: Interval to check for new revocations on, to regenerate the delta CRL.
        :param pulumi.Input[bool] disable: Disables or enables CRL building.
        :param pulumi.Input[bool] enable_delta: Enables building of delta CRLs with up-to-date revocation information, 
               augmenting the last complete CRL.  **Vault 1.12+**
        :param pulumi.Input[str] expiry: Specifies the time until expiration.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[bool] ocsp_disable: Disables the OCSP responder in Vault. **Vault 1.12+**
        :param pulumi.Input[str] ocsp_expiry: The amount of time an OCSP response can be cached for, useful for OCSP stapling 
               refresh durations. **Vault 1.12+**
        :param pulumi.Input[bool] unified_crl: Enables unified CRL and OCSP building. **Vault 1.13+**
        :param pulumi.Input[bool] unified_crl_on_existing_paths: Enables serving the unified CRL and OCSP on the existing, previously
               cluster-local paths. **Vault 1.13+**
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretBackendCrlConfigState.__new__(_SecretBackendCrlConfigState)

        __props__.__dict__["auto_rebuild"] = auto_rebuild
        __props__.__dict__["auto_rebuild_grace_period"] = auto_rebuild_grace_period
        __props__.__dict__["backend"] = backend
        __props__.__dict__["cross_cluster_revocation"] = cross_cluster_revocation
        __props__.__dict__["delta_rebuild_interval"] = delta_rebuild_interval
        __props__.__dict__["disable"] = disable
        __props__.__dict__["enable_delta"] = enable_delta
        __props__.__dict__["expiry"] = expiry
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["ocsp_disable"] = ocsp_disable
        __props__.__dict__["ocsp_expiry"] = ocsp_expiry
        __props__.__dict__["unified_crl"] = unified_crl
        __props__.__dict__["unified_crl_on_existing_paths"] = unified_crl_on_existing_paths
        return SecretBackendCrlConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="autoRebuild")
    def auto_rebuild(self) -> pulumi.Output[Optional[bool]]:
        """
        Enables periodic rebuilding of the CRL upon expiry. **Vault 1.12+**
        """
        return pulumi.get(self, "auto_rebuild")

    @property
    @pulumi.getter(name="autoRebuildGracePeriod")
    def auto_rebuild_grace_period(self) -> pulumi.Output[str]:
        """
        Grace period before CRL expiry to attempt rebuild of CRL. **Vault 1.12+**
        """
        return pulumi.get(self, "auto_rebuild_grace_period")

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[str]:
        """
        The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="crossClusterRevocation")
    def cross_cluster_revocation(self) -> pulumi.Output[bool]:
        """
        Enable cross-cluster revocation request queues. **Vault 1.13+**
        """
        return pulumi.get(self, "cross_cluster_revocation")

    @property
    @pulumi.getter(name="deltaRebuildInterval")
    def delta_rebuild_interval(self) -> pulumi.Output[str]:
        """
        Interval to check for new revocations on, to regenerate the delta CRL.
        """
        return pulumi.get(self, "delta_rebuild_interval")

    @property
    @pulumi.getter
    def disable(self) -> pulumi.Output[Optional[bool]]:
        """
        Disables or enables CRL building.
        """
        return pulumi.get(self, "disable")

    @property
    @pulumi.getter(name="enableDelta")
    def enable_delta(self) -> pulumi.Output[Optional[bool]]:
        """
        Enables building of delta CRLs with up-to-date revocation information, 
        augmenting the last complete CRL.  **Vault 1.12+**
        """
        return pulumi.get(self, "enable_delta")

    @property
    @pulumi.getter
    def expiry(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the time until expiration.
        """
        return pulumi.get(self, "expiry")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="ocspDisable")
    def ocsp_disable(self) -> pulumi.Output[Optional[bool]]:
        """
        Disables the OCSP responder in Vault. **Vault 1.12+**
        """
        return pulumi.get(self, "ocsp_disable")

    @property
    @pulumi.getter(name="ocspExpiry")
    def ocsp_expiry(self) -> pulumi.Output[str]:
        """
        The amount of time an OCSP response can be cached for, useful for OCSP stapling 
        refresh durations. **Vault 1.12+**
        """
        return pulumi.get(self, "ocsp_expiry")

    @property
    @pulumi.getter(name="unifiedCrl")
    def unified_crl(self) -> pulumi.Output[bool]:
        """
        Enables unified CRL and OCSP building. **Vault 1.13+**
        """
        return pulumi.get(self, "unified_crl")

    @property
    @pulumi.getter(name="unifiedCrlOnExistingPaths")
    def unified_crl_on_existing_paths(self) -> pulumi.Output[bool]:
        """
        Enables serving the unified CRL and OCSP on the existing, previously
        cluster-local paths. **Vault 1.13+**
        """
        return pulumi.get(self, "unified_crl_on_existing_paths")

