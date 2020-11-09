# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['SecretBackendRootSignIntermediate']


class SecretBackendRootSignIntermediate(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alt_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 common_name: Optional[pulumi.Input[str]] = None,
                 country: Optional[pulumi.Input[str]] = None,
                 csr: Optional[pulumi.Input[str]] = None,
                 exclude_cn_from_sans: Optional[pulumi.Input[bool]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 ip_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 locality: Optional[pulumi.Input[str]] = None,
                 max_path_length: Optional[pulumi.Input[int]] = None,
                 organization: Optional[pulumi.Input[str]] = None,
                 other_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ou: Optional[pulumi.Input[str]] = None,
                 permitted_dns_domains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 postal_code: Optional[pulumi.Input[str]] = None,
                 province: Optional[pulumi.Input[str]] = None,
                 street_address: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 uri_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 use_csr_values: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates an PKI certificate.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        root = vault.pki_secret.SecretBackendRootSignIntermediate("root",
            backend=vault_pki_secret_backend["root"]["path"],
            csr=vault_pki_secret_backend_intermediate_cert_request["intermediate"]["csr"],
            common_name="Intermediate CA",
            exclude_cn_from_sans=True,
            ou="My OU",
            organization="My organization",
            opts=ResourceOptions(depends_on=["vault_pki_secret_backend_intermediate_cert_request.intermediate"]))
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] alt_names: List of alternative names
        :param pulumi.Input[str] backend: The PKI secret backend the resource belongs to.
        :param pulumi.Input[str] common_name: CN of intermediate to create
        :param pulumi.Input[str] country: The country
        :param pulumi.Input[str] csr: The CSR
        :param pulumi.Input[bool] exclude_cn_from_sans: Flag to exclude CN from SANs
        :param pulumi.Input[str] format: The format of data
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_sans: List of alternative IPs
        :param pulumi.Input[str] locality: The locality
        :param pulumi.Input[int] max_path_length: The maximum path length to encode in the generated certificate
        :param pulumi.Input[str] organization: The organization
        :param pulumi.Input[Sequence[pulumi.Input[str]]] other_sans: List of other SANs
        :param pulumi.Input[str] ou: The organization unit
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permitted_dns_domains: List of domains for which certificates are allowed to be issued
        :param pulumi.Input[str] postal_code: The postal code
        :param pulumi.Input[str] province: The province
        :param pulumi.Input[str] street_address: The street address
        :param pulumi.Input[str] ttl: Time to live
        :param pulumi.Input[Sequence[pulumi.Input[str]]] uri_sans: List of alternative URIs
        :param pulumi.Input[bool] use_csr_values: Preserve CSR values
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

            __props__['alt_names'] = alt_names
            if backend is None:
                raise TypeError("Missing required property 'backend'")
            __props__['backend'] = backend
            if common_name is None:
                raise TypeError("Missing required property 'common_name'")
            __props__['common_name'] = common_name
            __props__['country'] = country
            if csr is None:
                raise TypeError("Missing required property 'csr'")
            __props__['csr'] = csr
            __props__['exclude_cn_from_sans'] = exclude_cn_from_sans
            __props__['format'] = format
            __props__['ip_sans'] = ip_sans
            __props__['locality'] = locality
            __props__['max_path_length'] = max_path_length
            __props__['organization'] = organization
            __props__['other_sans'] = other_sans
            __props__['ou'] = ou
            __props__['permitted_dns_domains'] = permitted_dns_domains
            __props__['postal_code'] = postal_code
            __props__['province'] = province
            __props__['street_address'] = street_address
            __props__['ttl'] = ttl
            __props__['uri_sans'] = uri_sans
            __props__['use_csr_values'] = use_csr_values
            __props__['ca_chain'] = None
            __props__['certificate'] = None
            __props__['issuing_ca'] = None
            __props__['serial'] = None
        super(SecretBackendRootSignIntermediate, __self__).__init__(
            'vault:pkiSecret/secretBackendRootSignIntermediate:SecretBackendRootSignIntermediate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alt_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            backend: Optional[pulumi.Input[str]] = None,
            ca_chain: Optional[pulumi.Input[str]] = None,
            certificate: Optional[pulumi.Input[str]] = None,
            common_name: Optional[pulumi.Input[str]] = None,
            country: Optional[pulumi.Input[str]] = None,
            csr: Optional[pulumi.Input[str]] = None,
            exclude_cn_from_sans: Optional[pulumi.Input[bool]] = None,
            format: Optional[pulumi.Input[str]] = None,
            ip_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            issuing_ca: Optional[pulumi.Input[str]] = None,
            locality: Optional[pulumi.Input[str]] = None,
            max_path_length: Optional[pulumi.Input[int]] = None,
            organization: Optional[pulumi.Input[str]] = None,
            other_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            ou: Optional[pulumi.Input[str]] = None,
            permitted_dns_domains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            postal_code: Optional[pulumi.Input[str]] = None,
            province: Optional[pulumi.Input[str]] = None,
            serial: Optional[pulumi.Input[str]] = None,
            street_address: Optional[pulumi.Input[str]] = None,
            ttl: Optional[pulumi.Input[str]] = None,
            uri_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            use_csr_values: Optional[pulumi.Input[bool]] = None) -> 'SecretBackendRootSignIntermediate':
        """
        Get an existing SecretBackendRootSignIntermediate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] alt_names: List of alternative names
        :param pulumi.Input[str] backend: The PKI secret backend the resource belongs to.
        :param pulumi.Input[str] ca_chain: The CA chain
        :param pulumi.Input[str] certificate: The certificate
        :param pulumi.Input[str] common_name: CN of intermediate to create
        :param pulumi.Input[str] country: The country
        :param pulumi.Input[str] csr: The CSR
        :param pulumi.Input[bool] exclude_cn_from_sans: Flag to exclude CN from SANs
        :param pulumi.Input[str] format: The format of data
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_sans: List of alternative IPs
        :param pulumi.Input[str] issuing_ca: The issuing CA
        :param pulumi.Input[str] locality: The locality
        :param pulumi.Input[int] max_path_length: The maximum path length to encode in the generated certificate
        :param pulumi.Input[str] organization: The organization
        :param pulumi.Input[Sequence[pulumi.Input[str]]] other_sans: List of other SANs
        :param pulumi.Input[str] ou: The organization unit
        :param pulumi.Input[Sequence[pulumi.Input[str]]] permitted_dns_domains: List of domains for which certificates are allowed to be issued
        :param pulumi.Input[str] postal_code: The postal code
        :param pulumi.Input[str] province: The province
        :param pulumi.Input[str] serial: The serial
        :param pulumi.Input[str] street_address: The street address
        :param pulumi.Input[str] ttl: Time to live
        :param pulumi.Input[Sequence[pulumi.Input[str]]] uri_sans: List of alternative URIs
        :param pulumi.Input[bool] use_csr_values: Preserve CSR values
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["alt_names"] = alt_names
        __props__["backend"] = backend
        __props__["ca_chain"] = ca_chain
        __props__["certificate"] = certificate
        __props__["common_name"] = common_name
        __props__["country"] = country
        __props__["csr"] = csr
        __props__["exclude_cn_from_sans"] = exclude_cn_from_sans
        __props__["format"] = format
        __props__["ip_sans"] = ip_sans
        __props__["issuing_ca"] = issuing_ca
        __props__["locality"] = locality
        __props__["max_path_length"] = max_path_length
        __props__["organization"] = organization
        __props__["other_sans"] = other_sans
        __props__["ou"] = ou
        __props__["permitted_dns_domains"] = permitted_dns_domains
        __props__["postal_code"] = postal_code
        __props__["province"] = province
        __props__["serial"] = serial
        __props__["street_address"] = street_address
        __props__["ttl"] = ttl
        __props__["uri_sans"] = uri_sans
        __props__["use_csr_values"] = use_csr_values
        return SecretBackendRootSignIntermediate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="altNames")
    def alt_names(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of alternative names
        """
        return pulumi.get(self, "alt_names")

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[str]:
        """
        The PKI secret backend the resource belongs to.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="caChain")
    def ca_chain(self) -> pulumi.Output[str]:
        """
        The CA chain
        """
        return pulumi.get(self, "ca_chain")

    @property
    @pulumi.getter
    def certificate(self) -> pulumi.Output[str]:
        """
        The certificate
        """
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter(name="commonName")
    def common_name(self) -> pulumi.Output[str]:
        """
        CN of intermediate to create
        """
        return pulumi.get(self, "common_name")

    @property
    @pulumi.getter
    def country(self) -> pulumi.Output[Optional[str]]:
        """
        The country
        """
        return pulumi.get(self, "country")

    @property
    @pulumi.getter
    def csr(self) -> pulumi.Output[str]:
        """
        The CSR
        """
        return pulumi.get(self, "csr")

    @property
    @pulumi.getter(name="excludeCnFromSans")
    def exclude_cn_from_sans(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to exclude CN from SANs
        """
        return pulumi.get(self, "exclude_cn_from_sans")

    @property
    @pulumi.getter
    def format(self) -> pulumi.Output[Optional[str]]:
        """
        The format of data
        """
        return pulumi.get(self, "format")

    @property
    @pulumi.getter(name="ipSans")
    def ip_sans(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of alternative IPs
        """
        return pulumi.get(self, "ip_sans")

    @property
    @pulumi.getter(name="issuingCa")
    def issuing_ca(self) -> pulumi.Output[str]:
        """
        The issuing CA
        """
        return pulumi.get(self, "issuing_ca")

    @property
    @pulumi.getter
    def locality(self) -> pulumi.Output[Optional[str]]:
        """
        The locality
        """
        return pulumi.get(self, "locality")

    @property
    @pulumi.getter(name="maxPathLength")
    def max_path_length(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum path length to encode in the generated certificate
        """
        return pulumi.get(self, "max_path_length")

    @property
    @pulumi.getter
    def organization(self) -> pulumi.Output[Optional[str]]:
        """
        The organization
        """
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter(name="otherSans")
    def other_sans(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of other SANs
        """
        return pulumi.get(self, "other_sans")

    @property
    @pulumi.getter
    def ou(self) -> pulumi.Output[Optional[str]]:
        """
        The organization unit
        """
        return pulumi.get(self, "ou")

    @property
    @pulumi.getter(name="permittedDnsDomains")
    def permitted_dns_domains(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of domains for which certificates are allowed to be issued
        """
        return pulumi.get(self, "permitted_dns_domains")

    @property
    @pulumi.getter(name="postalCode")
    def postal_code(self) -> pulumi.Output[Optional[str]]:
        """
        The postal code
        """
        return pulumi.get(self, "postal_code")

    @property
    @pulumi.getter
    def province(self) -> pulumi.Output[Optional[str]]:
        """
        The province
        """
        return pulumi.get(self, "province")

    @property
    @pulumi.getter
    def serial(self) -> pulumi.Output[str]:
        """
        The serial
        """
        return pulumi.get(self, "serial")

    @property
    @pulumi.getter(name="streetAddress")
    def street_address(self) -> pulumi.Output[Optional[str]]:
        """
        The street address
        """
        return pulumi.get(self, "street_address")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[Optional[str]]:
        """
        Time to live
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter(name="uriSans")
    def uri_sans(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of alternative URIs
        """
        return pulumi.get(self, "uri_sans")

    @property
    @pulumi.getter(name="useCsrValues")
    def use_csr_values(self) -> pulumi.Output[Optional[bool]]:
        """
        Preserve CSR values
        """
        return pulumi.get(self, "use_csr_values")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
