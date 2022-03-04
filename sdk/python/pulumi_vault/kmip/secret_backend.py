# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SecretBackendArgs', 'SecretBackend']

@pulumi.input_type
class SecretBackendArgs:
    def __init__(__self__, *,
                 path: pulumi.Input[str],
                 default_tls_client_key_bits: Optional[pulumi.Input[int]] = None,
                 default_tls_client_key_type: Optional[pulumi.Input[str]] = None,
                 default_tls_client_ttl: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 listen_addrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 server_hostnames: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 server_ips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tls_ca_key_bits: Optional[pulumi.Input[int]] = None,
                 tls_ca_key_type: Optional[pulumi.Input[str]] = None,
                 tls_min_version: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SecretBackend resource.
        :param pulumi.Input[str] path: The unique path this backend should be mounted at. Must
               not begin or end with a `/`. Defaults to `kmip`.
        :param pulumi.Input[int] default_tls_client_key_bits: Client certificate key bits, valid values depend on key type.
        :param pulumi.Input[str] default_tls_client_key_type: Client certificate key type, `rsa` or `ec`.
        :param pulumi.Input[int] default_tls_client_ttl: Client certificate TTL in seconds
        :param pulumi.Input[str] description: A human-friendly description for this backend.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] listen_addrs: Addresses the KMIP server should listen on (`host:port`).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] server_hostnames: Hostnames to include in the server's TLS certificate as SAN DNS names. The first will be used as the common name (CN).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] server_ips: IPs to include in the server's TLS certificate as SAN IP addresses.
        :param pulumi.Input[int] tls_ca_key_bits: CA key bits, valid values depend on key type.
        :param pulumi.Input[str] tls_ca_key_type: CA key type, rsa or ec.
        :param pulumi.Input[str] tls_min_version: Minimum TLS version to accept.
        """
        pulumi.set(__self__, "path", path)
        if default_tls_client_key_bits is not None:
            pulumi.set(__self__, "default_tls_client_key_bits", default_tls_client_key_bits)
        if default_tls_client_key_type is not None:
            pulumi.set(__self__, "default_tls_client_key_type", default_tls_client_key_type)
        if default_tls_client_ttl is not None:
            pulumi.set(__self__, "default_tls_client_ttl", default_tls_client_ttl)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if listen_addrs is not None:
            pulumi.set(__self__, "listen_addrs", listen_addrs)
        if server_hostnames is not None:
            pulumi.set(__self__, "server_hostnames", server_hostnames)
        if server_ips is not None:
            pulumi.set(__self__, "server_ips", server_ips)
        if tls_ca_key_bits is not None:
            pulumi.set(__self__, "tls_ca_key_bits", tls_ca_key_bits)
        if tls_ca_key_type is not None:
            pulumi.set(__self__, "tls_ca_key_type", tls_ca_key_type)
        if tls_min_version is not None:
            pulumi.set(__self__, "tls_min_version", tls_min_version)

    @property
    @pulumi.getter
    def path(self) -> pulumi.Input[str]:
        """
        The unique path this backend should be mounted at. Must
        not begin or end with a `/`. Defaults to `kmip`.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: pulumi.Input[str]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter(name="defaultTlsClientKeyBits")
    def default_tls_client_key_bits(self) -> Optional[pulumi.Input[int]]:
        """
        Client certificate key bits, valid values depend on key type.
        """
        return pulumi.get(self, "default_tls_client_key_bits")

    @default_tls_client_key_bits.setter
    def default_tls_client_key_bits(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_tls_client_key_bits", value)

    @property
    @pulumi.getter(name="defaultTlsClientKeyType")
    def default_tls_client_key_type(self) -> Optional[pulumi.Input[str]]:
        """
        Client certificate key type, `rsa` or `ec`.
        """
        return pulumi.get(self, "default_tls_client_key_type")

    @default_tls_client_key_type.setter
    def default_tls_client_key_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_tls_client_key_type", value)

    @property
    @pulumi.getter(name="defaultTlsClientTtl")
    def default_tls_client_ttl(self) -> Optional[pulumi.Input[int]]:
        """
        Client certificate TTL in seconds
        """
        return pulumi.get(self, "default_tls_client_ttl")

    @default_tls_client_ttl.setter
    def default_tls_client_ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_tls_client_ttl", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A human-friendly description for this backend.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="listenAddrs")
    def listen_addrs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Addresses the KMIP server should listen on (`host:port`).
        """
        return pulumi.get(self, "listen_addrs")

    @listen_addrs.setter
    def listen_addrs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "listen_addrs", value)

    @property
    @pulumi.getter(name="serverHostnames")
    def server_hostnames(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Hostnames to include in the server's TLS certificate as SAN DNS names. The first will be used as the common name (CN).
        """
        return pulumi.get(self, "server_hostnames")

    @server_hostnames.setter
    def server_hostnames(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "server_hostnames", value)

    @property
    @pulumi.getter(name="serverIps")
    def server_ips(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        IPs to include in the server's TLS certificate as SAN IP addresses.
        """
        return pulumi.get(self, "server_ips")

    @server_ips.setter
    def server_ips(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "server_ips", value)

    @property
    @pulumi.getter(name="tlsCaKeyBits")
    def tls_ca_key_bits(self) -> Optional[pulumi.Input[int]]:
        """
        CA key bits, valid values depend on key type.
        """
        return pulumi.get(self, "tls_ca_key_bits")

    @tls_ca_key_bits.setter
    def tls_ca_key_bits(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "tls_ca_key_bits", value)

    @property
    @pulumi.getter(name="tlsCaKeyType")
    def tls_ca_key_type(self) -> Optional[pulumi.Input[str]]:
        """
        CA key type, rsa or ec.
        """
        return pulumi.get(self, "tls_ca_key_type")

    @tls_ca_key_type.setter
    def tls_ca_key_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_ca_key_type", value)

    @property
    @pulumi.getter(name="tlsMinVersion")
    def tls_min_version(self) -> Optional[pulumi.Input[str]]:
        """
        Minimum TLS version to accept.
        """
        return pulumi.get(self, "tls_min_version")

    @tls_min_version.setter
    def tls_min_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_min_version", value)


@pulumi.input_type
class _SecretBackendState:
    def __init__(__self__, *,
                 default_tls_client_key_bits: Optional[pulumi.Input[int]] = None,
                 default_tls_client_key_type: Optional[pulumi.Input[str]] = None,
                 default_tls_client_ttl: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 listen_addrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 server_hostnames: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 server_ips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tls_ca_key_bits: Optional[pulumi.Input[int]] = None,
                 tls_ca_key_type: Optional[pulumi.Input[str]] = None,
                 tls_min_version: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SecretBackend resources.
        :param pulumi.Input[int] default_tls_client_key_bits: Client certificate key bits, valid values depend on key type.
        :param pulumi.Input[str] default_tls_client_key_type: Client certificate key type, `rsa` or `ec`.
        :param pulumi.Input[int] default_tls_client_ttl: Client certificate TTL in seconds
        :param pulumi.Input[str] description: A human-friendly description for this backend.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] listen_addrs: Addresses the KMIP server should listen on (`host:port`).
        :param pulumi.Input[str] path: The unique path this backend should be mounted at. Must
               not begin or end with a `/`. Defaults to `kmip`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] server_hostnames: Hostnames to include in the server's TLS certificate as SAN DNS names. The first will be used as the common name (CN).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] server_ips: IPs to include in the server's TLS certificate as SAN IP addresses.
        :param pulumi.Input[int] tls_ca_key_bits: CA key bits, valid values depend on key type.
        :param pulumi.Input[str] tls_ca_key_type: CA key type, rsa or ec.
        :param pulumi.Input[str] tls_min_version: Minimum TLS version to accept.
        """
        if default_tls_client_key_bits is not None:
            pulumi.set(__self__, "default_tls_client_key_bits", default_tls_client_key_bits)
        if default_tls_client_key_type is not None:
            pulumi.set(__self__, "default_tls_client_key_type", default_tls_client_key_type)
        if default_tls_client_ttl is not None:
            pulumi.set(__self__, "default_tls_client_ttl", default_tls_client_ttl)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if listen_addrs is not None:
            pulumi.set(__self__, "listen_addrs", listen_addrs)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if server_hostnames is not None:
            pulumi.set(__self__, "server_hostnames", server_hostnames)
        if server_ips is not None:
            pulumi.set(__self__, "server_ips", server_ips)
        if tls_ca_key_bits is not None:
            pulumi.set(__self__, "tls_ca_key_bits", tls_ca_key_bits)
        if tls_ca_key_type is not None:
            pulumi.set(__self__, "tls_ca_key_type", tls_ca_key_type)
        if tls_min_version is not None:
            pulumi.set(__self__, "tls_min_version", tls_min_version)

    @property
    @pulumi.getter(name="defaultTlsClientKeyBits")
    def default_tls_client_key_bits(self) -> Optional[pulumi.Input[int]]:
        """
        Client certificate key bits, valid values depend on key type.
        """
        return pulumi.get(self, "default_tls_client_key_bits")

    @default_tls_client_key_bits.setter
    def default_tls_client_key_bits(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_tls_client_key_bits", value)

    @property
    @pulumi.getter(name="defaultTlsClientKeyType")
    def default_tls_client_key_type(self) -> Optional[pulumi.Input[str]]:
        """
        Client certificate key type, `rsa` or `ec`.
        """
        return pulumi.get(self, "default_tls_client_key_type")

    @default_tls_client_key_type.setter
    def default_tls_client_key_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_tls_client_key_type", value)

    @property
    @pulumi.getter(name="defaultTlsClientTtl")
    def default_tls_client_ttl(self) -> Optional[pulumi.Input[int]]:
        """
        Client certificate TTL in seconds
        """
        return pulumi.get(self, "default_tls_client_ttl")

    @default_tls_client_ttl.setter
    def default_tls_client_ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_tls_client_ttl", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A human-friendly description for this backend.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="listenAddrs")
    def listen_addrs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Addresses the KMIP server should listen on (`host:port`).
        """
        return pulumi.get(self, "listen_addrs")

    @listen_addrs.setter
    def listen_addrs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "listen_addrs", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        The unique path this backend should be mounted at. Must
        not begin or end with a `/`. Defaults to `kmip`.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter(name="serverHostnames")
    def server_hostnames(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Hostnames to include in the server's TLS certificate as SAN DNS names. The first will be used as the common name (CN).
        """
        return pulumi.get(self, "server_hostnames")

    @server_hostnames.setter
    def server_hostnames(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "server_hostnames", value)

    @property
    @pulumi.getter(name="serverIps")
    def server_ips(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        IPs to include in the server's TLS certificate as SAN IP addresses.
        """
        return pulumi.get(self, "server_ips")

    @server_ips.setter
    def server_ips(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "server_ips", value)

    @property
    @pulumi.getter(name="tlsCaKeyBits")
    def tls_ca_key_bits(self) -> Optional[pulumi.Input[int]]:
        """
        CA key bits, valid values depend on key type.
        """
        return pulumi.get(self, "tls_ca_key_bits")

    @tls_ca_key_bits.setter
    def tls_ca_key_bits(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "tls_ca_key_bits", value)

    @property
    @pulumi.getter(name="tlsCaKeyType")
    def tls_ca_key_type(self) -> Optional[pulumi.Input[str]]:
        """
        CA key type, rsa or ec.
        """
        return pulumi.get(self, "tls_ca_key_type")

    @tls_ca_key_type.setter
    def tls_ca_key_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_ca_key_type", value)

    @property
    @pulumi.getter(name="tlsMinVersion")
    def tls_min_version(self) -> Optional[pulumi.Input[str]]:
        """
        Minimum TLS version to accept.
        """
        return pulumi.get(self, "tls_min_version")

    @tls_min_version.setter
    def tls_min_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tls_min_version", value)


class SecretBackend(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_tls_client_key_bits: Optional[pulumi.Input[int]] = None,
                 default_tls_client_key_type: Optional[pulumi.Input[str]] = None,
                 default_tls_client_ttl: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 listen_addrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 server_hostnames: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 server_ips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tls_ca_key_bits: Optional[pulumi.Input[int]] = None,
                 tls_ca_key_type: Optional[pulumi.Input[str]] = None,
                 tls_min_version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages KMIP Secret backends in a Vault server. This feature requires
        Vault Enterprise. See the [Vault documentation](https://www.vaultproject.io/docs/secrets/kmip)
        for more information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        default = vault.kmip.SecretBackend("default",
            default_tls_client_key_bits=4096,
            default_tls_client_key_type="rsa",
            default_tls_client_ttl=86400,
            description="Vault KMIP backend",
            listen_addrs=[
                "127.0.0.1:5696",
                "127.0.0.1:8080",
            ],
            path="kmip",
            tls_ca_key_bits=4096,
            tls_ca_key_type="rsa")
        ```

        ## Import

        KMIP Secret backend can be imported using the `path`, e.g.

        ```sh
         $ pulumi import vault:kmip/secretBackend:SecretBackend default kmip
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] default_tls_client_key_bits: Client certificate key bits, valid values depend on key type.
        :param pulumi.Input[str] default_tls_client_key_type: Client certificate key type, `rsa` or `ec`.
        :param pulumi.Input[int] default_tls_client_ttl: Client certificate TTL in seconds
        :param pulumi.Input[str] description: A human-friendly description for this backend.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] listen_addrs: Addresses the KMIP server should listen on (`host:port`).
        :param pulumi.Input[str] path: The unique path this backend should be mounted at. Must
               not begin or end with a `/`. Defaults to `kmip`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] server_hostnames: Hostnames to include in the server's TLS certificate as SAN DNS names. The first will be used as the common name (CN).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] server_ips: IPs to include in the server's TLS certificate as SAN IP addresses.
        :param pulumi.Input[int] tls_ca_key_bits: CA key bits, valid values depend on key type.
        :param pulumi.Input[str] tls_ca_key_type: CA key type, rsa or ec.
        :param pulumi.Input[str] tls_min_version: Minimum TLS version to accept.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretBackendArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages KMIP Secret backends in a Vault server. This feature requires
        Vault Enterprise. See the [Vault documentation](https://www.vaultproject.io/docs/secrets/kmip)
        for more information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        default = vault.kmip.SecretBackend("default",
            default_tls_client_key_bits=4096,
            default_tls_client_key_type="rsa",
            default_tls_client_ttl=86400,
            description="Vault KMIP backend",
            listen_addrs=[
                "127.0.0.1:5696",
                "127.0.0.1:8080",
            ],
            path="kmip",
            tls_ca_key_bits=4096,
            tls_ca_key_type="rsa")
        ```

        ## Import

        KMIP Secret backend can be imported using the `path`, e.g.

        ```sh
         $ pulumi import vault:kmip/secretBackend:SecretBackend default kmip
        ```

        :param str resource_name: The name of the resource.
        :param SecretBackendArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretBackendArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_tls_client_key_bits: Optional[pulumi.Input[int]] = None,
                 default_tls_client_key_type: Optional[pulumi.Input[str]] = None,
                 default_tls_client_ttl: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 listen_addrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 server_hostnames: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 server_ips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tls_ca_key_bits: Optional[pulumi.Input[int]] = None,
                 tls_ca_key_type: Optional[pulumi.Input[str]] = None,
                 tls_min_version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretBackendArgs.__new__(SecretBackendArgs)

            __props__.__dict__["default_tls_client_key_bits"] = default_tls_client_key_bits
            __props__.__dict__["default_tls_client_key_type"] = default_tls_client_key_type
            __props__.__dict__["default_tls_client_ttl"] = default_tls_client_ttl
            __props__.__dict__["description"] = description
            __props__.__dict__["listen_addrs"] = listen_addrs
            if path is None and not opts.urn:
                raise TypeError("Missing required property 'path'")
            __props__.__dict__["path"] = path
            __props__.__dict__["server_hostnames"] = server_hostnames
            __props__.__dict__["server_ips"] = server_ips
            __props__.__dict__["tls_ca_key_bits"] = tls_ca_key_bits
            __props__.__dict__["tls_ca_key_type"] = tls_ca_key_type
            __props__.__dict__["tls_min_version"] = tls_min_version
        super(SecretBackend, __self__).__init__(
            'vault:kmip/secretBackend:SecretBackend',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            default_tls_client_key_bits: Optional[pulumi.Input[int]] = None,
            default_tls_client_key_type: Optional[pulumi.Input[str]] = None,
            default_tls_client_ttl: Optional[pulumi.Input[int]] = None,
            description: Optional[pulumi.Input[str]] = None,
            listen_addrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            path: Optional[pulumi.Input[str]] = None,
            server_hostnames: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            server_ips: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tls_ca_key_bits: Optional[pulumi.Input[int]] = None,
            tls_ca_key_type: Optional[pulumi.Input[str]] = None,
            tls_min_version: Optional[pulumi.Input[str]] = None) -> 'SecretBackend':
        """
        Get an existing SecretBackend resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] default_tls_client_key_bits: Client certificate key bits, valid values depend on key type.
        :param pulumi.Input[str] default_tls_client_key_type: Client certificate key type, `rsa` or `ec`.
        :param pulumi.Input[int] default_tls_client_ttl: Client certificate TTL in seconds
        :param pulumi.Input[str] description: A human-friendly description for this backend.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] listen_addrs: Addresses the KMIP server should listen on (`host:port`).
        :param pulumi.Input[str] path: The unique path this backend should be mounted at. Must
               not begin or end with a `/`. Defaults to `kmip`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] server_hostnames: Hostnames to include in the server's TLS certificate as SAN DNS names. The first will be used as the common name (CN).
        :param pulumi.Input[Sequence[pulumi.Input[str]]] server_ips: IPs to include in the server's TLS certificate as SAN IP addresses.
        :param pulumi.Input[int] tls_ca_key_bits: CA key bits, valid values depend on key type.
        :param pulumi.Input[str] tls_ca_key_type: CA key type, rsa or ec.
        :param pulumi.Input[str] tls_min_version: Minimum TLS version to accept.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretBackendState.__new__(_SecretBackendState)

        __props__.__dict__["default_tls_client_key_bits"] = default_tls_client_key_bits
        __props__.__dict__["default_tls_client_key_type"] = default_tls_client_key_type
        __props__.__dict__["default_tls_client_ttl"] = default_tls_client_ttl
        __props__.__dict__["description"] = description
        __props__.__dict__["listen_addrs"] = listen_addrs
        __props__.__dict__["path"] = path
        __props__.__dict__["server_hostnames"] = server_hostnames
        __props__.__dict__["server_ips"] = server_ips
        __props__.__dict__["tls_ca_key_bits"] = tls_ca_key_bits
        __props__.__dict__["tls_ca_key_type"] = tls_ca_key_type
        __props__.__dict__["tls_min_version"] = tls_min_version
        return SecretBackend(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="defaultTlsClientKeyBits")
    def default_tls_client_key_bits(self) -> pulumi.Output[int]:
        """
        Client certificate key bits, valid values depend on key type.
        """
        return pulumi.get(self, "default_tls_client_key_bits")

    @property
    @pulumi.getter(name="defaultTlsClientKeyType")
    def default_tls_client_key_type(self) -> pulumi.Output[str]:
        """
        Client certificate key type, `rsa` or `ec`.
        """
        return pulumi.get(self, "default_tls_client_key_type")

    @property
    @pulumi.getter(name="defaultTlsClientTtl")
    def default_tls_client_ttl(self) -> pulumi.Output[int]:
        """
        Client certificate TTL in seconds
        """
        return pulumi.get(self, "default_tls_client_ttl")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A human-friendly description for this backend.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="listenAddrs")
    def listen_addrs(self) -> pulumi.Output[Sequence[str]]:
        """
        Addresses the KMIP server should listen on (`host:port`).
        """
        return pulumi.get(self, "listen_addrs")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[str]:
        """
        The unique path this backend should be mounted at. Must
        not begin or end with a `/`. Defaults to `kmip`.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="serverHostnames")
    def server_hostnames(self) -> pulumi.Output[Sequence[str]]:
        """
        Hostnames to include in the server's TLS certificate as SAN DNS names. The first will be used as the common name (CN).
        """
        return pulumi.get(self, "server_hostnames")

    @property
    @pulumi.getter(name="serverIps")
    def server_ips(self) -> pulumi.Output[Sequence[str]]:
        """
        IPs to include in the server's TLS certificate as SAN IP addresses.
        """
        return pulumi.get(self, "server_ips")

    @property
    @pulumi.getter(name="tlsCaKeyBits")
    def tls_ca_key_bits(self) -> pulumi.Output[int]:
        """
        CA key bits, valid values depend on key type.
        """
        return pulumi.get(self, "tls_ca_key_bits")

    @property
    @pulumi.getter(name="tlsCaKeyType")
    def tls_ca_key_type(self) -> pulumi.Output[str]:
        """
        CA key type, rsa or ec.
        """
        return pulumi.get(self, "tls_ca_key_type")

    @property
    @pulumi.getter(name="tlsMinVersion")
    def tls_min_version(self) -> pulumi.Output[str]:
        """
        Minimum TLS version to accept.
        """
        return pulumi.get(self, "tls_min_version")
