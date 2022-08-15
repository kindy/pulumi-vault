# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['MfaOktaArgs', 'MfaOkta']

@pulumi.input_type
class MfaOktaArgs:
    def __init__(__self__, *,
                 api_token: pulumi.Input[str],
                 mount_accessor: pulumi.Input[str],
                 org_name: pulumi.Input[str],
                 base_url: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 primary_email: Optional[pulumi.Input[bool]] = None,
                 username_format: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MfaOkta resource.
        :param pulumi.Input[str] api_token: `(string: <required>)` - Okta API key.
        :param pulumi.Input[str] mount_accessor: `(string: <required>)` - The mount to tie this method to for use in automatic mappings. 
               The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
        :param pulumi.Input[str] org_name: `(string: <required>)` - Name of the organization to be used in the Okta API.
        :param pulumi.Input[str] base_url: `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`, 
               `oktapreview.com`, and `okta-emea.com`.
        :param pulumi.Input[str] name: `(string: <required>)` – Name of the MFA method.
        :param pulumi.Input[bool] primary_email: `(string: <required>)` - If set to true, the username will only match the 
               primary email for the account.
        :param pulumi.Input[str] username_format: `(string)` - A format string for mapping Identity names to MFA method names. 
               Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`.
               If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
               - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
               - entity.name: The name configured for the Entity
               - alias.metadata.`<key>`: The value of the Alias's metadata parameter
               - entity.metadata.`<key>`: The value of the Entity's metadata parameter
        """
        pulumi.set(__self__, "api_token", api_token)
        pulumi.set(__self__, "mount_accessor", mount_accessor)
        pulumi.set(__self__, "org_name", org_name)
        if base_url is not None:
            pulumi.set(__self__, "base_url", base_url)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if primary_email is not None:
            pulumi.set(__self__, "primary_email", primary_email)
        if username_format is not None:
            pulumi.set(__self__, "username_format", username_format)

    @property
    @pulumi.getter(name="apiToken")
    def api_token(self) -> pulumi.Input[str]:
        """
        `(string: <required>)` - Okta API key.
        """
        return pulumi.get(self, "api_token")

    @api_token.setter
    def api_token(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_token", value)

    @property
    @pulumi.getter(name="mountAccessor")
    def mount_accessor(self) -> pulumi.Input[str]:
        """
        `(string: <required>)` - The mount to tie this method to for use in automatic mappings. 
        The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
        """
        return pulumi.get(self, "mount_accessor")

    @mount_accessor.setter
    def mount_accessor(self, value: pulumi.Input[str]):
        pulumi.set(self, "mount_accessor", value)

    @property
    @pulumi.getter(name="orgName")
    def org_name(self) -> pulumi.Input[str]:
        """
        `(string: <required>)` - Name of the organization to be used in the Okta API.
        """
        return pulumi.get(self, "org_name")

    @org_name.setter
    def org_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "org_name", value)

    @property
    @pulumi.getter(name="baseUrl")
    def base_url(self) -> Optional[pulumi.Input[str]]:
        """
        `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`, 
        `oktapreview.com`, and `okta-emea.com`.
        """
        return pulumi.get(self, "base_url")

    @base_url.setter
    def base_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "base_url", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        `(string: <required>)` – Name of the MFA method.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="primaryEmail")
    def primary_email(self) -> Optional[pulumi.Input[bool]]:
        """
        `(string: <required>)` - If set to true, the username will only match the 
        primary email for the account.
        """
        return pulumi.get(self, "primary_email")

    @primary_email.setter
    def primary_email(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "primary_email", value)

    @property
    @pulumi.getter(name="usernameFormat")
    def username_format(self) -> Optional[pulumi.Input[str]]:
        """
        `(string)` - A format string for mapping Identity names to MFA method names. 
        Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`.
        If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
        - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
        - entity.name: The name configured for the Entity
        - alias.metadata.`<key>`: The value of the Alias's metadata parameter
        - entity.metadata.`<key>`: The value of the Entity's metadata parameter
        """
        return pulumi.get(self, "username_format")

    @username_format.setter
    def username_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username_format", value)


@pulumi.input_type
class _MfaOktaState:
    def __init__(__self__, *,
                 api_token: Optional[pulumi.Input[str]] = None,
                 base_url: Optional[pulumi.Input[str]] = None,
                 mount_accessor: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_name: Optional[pulumi.Input[str]] = None,
                 primary_email: Optional[pulumi.Input[bool]] = None,
                 username_format: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MfaOkta resources.
        :param pulumi.Input[str] api_token: `(string: <required>)` - Okta API key.
        :param pulumi.Input[str] base_url: `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`, 
               `oktapreview.com`, and `okta-emea.com`.
        :param pulumi.Input[str] mount_accessor: `(string: <required>)` - The mount to tie this method to for use in automatic mappings. 
               The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
        :param pulumi.Input[str] name: `(string: <required>)` – Name of the MFA method.
        :param pulumi.Input[str] org_name: `(string: <required>)` - Name of the organization to be used in the Okta API.
        :param pulumi.Input[bool] primary_email: `(string: <required>)` - If set to true, the username will only match the 
               primary email for the account.
        :param pulumi.Input[str] username_format: `(string)` - A format string for mapping Identity names to MFA method names. 
               Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`.
               If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
               - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
               - entity.name: The name configured for the Entity
               - alias.metadata.`<key>`: The value of the Alias's metadata parameter
               - entity.metadata.`<key>`: The value of the Entity's metadata parameter
        """
        if api_token is not None:
            pulumi.set(__self__, "api_token", api_token)
        if base_url is not None:
            pulumi.set(__self__, "base_url", base_url)
        if mount_accessor is not None:
            pulumi.set(__self__, "mount_accessor", mount_accessor)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_name is not None:
            pulumi.set(__self__, "org_name", org_name)
        if primary_email is not None:
            pulumi.set(__self__, "primary_email", primary_email)
        if username_format is not None:
            pulumi.set(__self__, "username_format", username_format)

    @property
    @pulumi.getter(name="apiToken")
    def api_token(self) -> Optional[pulumi.Input[str]]:
        """
        `(string: <required>)` - Okta API key.
        """
        return pulumi.get(self, "api_token")

    @api_token.setter
    def api_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_token", value)

    @property
    @pulumi.getter(name="baseUrl")
    def base_url(self) -> Optional[pulumi.Input[str]]:
        """
        `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`, 
        `oktapreview.com`, and `okta-emea.com`.
        """
        return pulumi.get(self, "base_url")

    @base_url.setter
    def base_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "base_url", value)

    @property
    @pulumi.getter(name="mountAccessor")
    def mount_accessor(self) -> Optional[pulumi.Input[str]]:
        """
        `(string: <required>)` - The mount to tie this method to for use in automatic mappings. 
        The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
        """
        return pulumi.get(self, "mount_accessor")

    @mount_accessor.setter
    def mount_accessor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mount_accessor", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        `(string: <required>)` – Name of the MFA method.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="orgName")
    def org_name(self) -> Optional[pulumi.Input[str]]:
        """
        `(string: <required>)` - Name of the organization to be used in the Okta API.
        """
        return pulumi.get(self, "org_name")

    @org_name.setter
    def org_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_name", value)

    @property
    @pulumi.getter(name="primaryEmail")
    def primary_email(self) -> Optional[pulumi.Input[bool]]:
        """
        `(string: <required>)` - If set to true, the username will only match the 
        primary email for the account.
        """
        return pulumi.get(self, "primary_email")

    @primary_email.setter
    def primary_email(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "primary_email", value)

    @property
    @pulumi.getter(name="usernameFormat")
    def username_format(self) -> Optional[pulumi.Input[str]]:
        """
        `(string)` - A format string for mapping Identity names to MFA method names. 
        Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`.
        If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
        - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
        - entity.name: The name configured for the Entity
        - alias.metadata.`<key>`: The value of the Alias's metadata parameter
        - entity.metadata.`<key>`: The value of the Entity's metadata parameter
        """
        return pulumi.get(self, "username_format")

    @username_format.setter
    def username_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username_format", value)


class MfaOkta(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_token: Optional[pulumi.Input[str]] = None,
                 base_url: Optional[pulumi.Input[str]] = None,
                 mount_accessor: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_name: Optional[pulumi.Input[str]] = None,
                 primary_email: Optional[pulumi.Input[bool]] = None,
                 username_format: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage [Okta MFA](https://www.vaultproject.io/docs/enterprise/mfa/mfa-okta).

        **Note** this feature is available only with Vault Enterprise.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        userpass = vault.AuthBackend("userpass",
            type="userpass",
            path="userpass")
        my_okta = vault.MfaOkta("myOkta",
            mount_accessor=userpass.accessor,
            username_format="user@example.com",
            org_name="hashicorp",
            api_token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")
        ```

        ## Import

        Mounts can be imported using the `path`, e.g.

        ```sh
         $ pulumi import vault:index/mfaOkta:MfaOkta my_okta my_okta
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_token: `(string: <required>)` - Okta API key.
        :param pulumi.Input[str] base_url: `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`, 
               `oktapreview.com`, and `okta-emea.com`.
        :param pulumi.Input[str] mount_accessor: `(string: <required>)` - The mount to tie this method to for use in automatic mappings. 
               The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
        :param pulumi.Input[str] name: `(string: <required>)` – Name of the MFA method.
        :param pulumi.Input[str] org_name: `(string: <required>)` - Name of the organization to be used in the Okta API.
        :param pulumi.Input[bool] primary_email: `(string: <required>)` - If set to true, the username will only match the 
               primary email for the account.
        :param pulumi.Input[str] username_format: `(string)` - A format string for mapping Identity names to MFA method names. 
               Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`.
               If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
               - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
               - entity.name: The name configured for the Entity
               - alias.metadata.`<key>`: The value of the Alias's metadata parameter
               - entity.metadata.`<key>`: The value of the Entity's metadata parameter
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MfaOktaArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage [Okta MFA](https://www.vaultproject.io/docs/enterprise/mfa/mfa-okta).

        **Note** this feature is available only with Vault Enterprise.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        userpass = vault.AuthBackend("userpass",
            type="userpass",
            path="userpass")
        my_okta = vault.MfaOkta("myOkta",
            mount_accessor=userpass.accessor,
            username_format="user@example.com",
            org_name="hashicorp",
            api_token="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9")
        ```

        ## Import

        Mounts can be imported using the `path`, e.g.

        ```sh
         $ pulumi import vault:index/mfaOkta:MfaOkta my_okta my_okta
        ```

        :param str resource_name: The name of the resource.
        :param MfaOktaArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MfaOktaArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_token: Optional[pulumi.Input[str]] = None,
                 base_url: Optional[pulumi.Input[str]] = None,
                 mount_accessor: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_name: Optional[pulumi.Input[str]] = None,
                 primary_email: Optional[pulumi.Input[bool]] = None,
                 username_format: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MfaOktaArgs.__new__(MfaOktaArgs)

            if api_token is None and not opts.urn:
                raise TypeError("Missing required property 'api_token'")
            __props__.__dict__["api_token"] = api_token
            __props__.__dict__["base_url"] = base_url
            if mount_accessor is None and not opts.urn:
                raise TypeError("Missing required property 'mount_accessor'")
            __props__.__dict__["mount_accessor"] = mount_accessor
            __props__.__dict__["name"] = name
            if org_name is None and not opts.urn:
                raise TypeError("Missing required property 'org_name'")
            __props__.__dict__["org_name"] = org_name
            __props__.__dict__["primary_email"] = primary_email
            __props__.__dict__["username_format"] = username_format
        super(MfaOkta, __self__).__init__(
            'vault:index/mfaOkta:MfaOkta',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_token: Optional[pulumi.Input[str]] = None,
            base_url: Optional[pulumi.Input[str]] = None,
            mount_accessor: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_name: Optional[pulumi.Input[str]] = None,
            primary_email: Optional[pulumi.Input[bool]] = None,
            username_format: Optional[pulumi.Input[str]] = None) -> 'MfaOkta':
        """
        Get an existing MfaOkta resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_token: `(string: <required>)` - Okta API key.
        :param pulumi.Input[str] base_url: `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`, 
               `oktapreview.com`, and `okta-emea.com`.
        :param pulumi.Input[str] mount_accessor: `(string: <required>)` - The mount to tie this method to for use in automatic mappings. 
               The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
        :param pulumi.Input[str] name: `(string: <required>)` – Name of the MFA method.
        :param pulumi.Input[str] org_name: `(string: <required>)` - Name of the organization to be used in the Okta API.
        :param pulumi.Input[bool] primary_email: `(string: <required>)` - If set to true, the username will only match the 
               primary email for the account.
        :param pulumi.Input[str] username_format: `(string)` - A format string for mapping Identity names to MFA method names. 
               Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`.
               If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
               - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
               - entity.name: The name configured for the Entity
               - alias.metadata.`<key>`: The value of the Alias's metadata parameter
               - entity.metadata.`<key>`: The value of the Entity's metadata parameter
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MfaOktaState.__new__(_MfaOktaState)

        __props__.__dict__["api_token"] = api_token
        __props__.__dict__["base_url"] = base_url
        __props__.__dict__["mount_accessor"] = mount_accessor
        __props__.__dict__["name"] = name
        __props__.__dict__["org_name"] = org_name
        __props__.__dict__["primary_email"] = primary_email
        __props__.__dict__["username_format"] = username_format
        return MfaOkta(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiToken")
    def api_token(self) -> pulumi.Output[str]:
        """
        `(string: <required>)` - Okta API key.
        """
        return pulumi.get(self, "api_token")

    @property
    @pulumi.getter(name="baseUrl")
    def base_url(self) -> pulumi.Output[Optional[str]]:
        """
        `(string)` - If set, will be used as the base domain for API requests. Examples are `okta.com`, 
        `oktapreview.com`, and `okta-emea.com`.
        """
        return pulumi.get(self, "base_url")

    @property
    @pulumi.getter(name="mountAccessor")
    def mount_accessor(self) -> pulumi.Output[str]:
        """
        `(string: <required>)` - The mount to tie this method to for use in automatic mappings. 
        The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
        """
        return pulumi.get(self, "mount_accessor")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        `(string: <required>)` – Name of the MFA method.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgName")
    def org_name(self) -> pulumi.Output[str]:
        """
        `(string: <required>)` - Name of the organization to be used in the Okta API.
        """
        return pulumi.get(self, "org_name")

    @property
    @pulumi.getter(name="primaryEmail")
    def primary_email(self) -> pulumi.Output[Optional[bool]]:
        """
        `(string: <required>)` - If set to true, the username will only match the 
        primary email for the account.
        """
        return pulumi.get(self, "primary_email")

    @property
    @pulumi.getter(name="usernameFormat")
    def username_format(self) -> pulumi.Output[Optional[str]]:
        """
        `(string)` - A format string for mapping Identity names to MFA method names. 
        Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`.
        If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
        - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
        - entity.name: The name configured for the Entity
        - alias.metadata.`<key>`: The value of the Alias's metadata parameter
        - entity.metadata.`<key>`: The value of the Entity's metadata parameter
        """
        return pulumi.get(self, "username_format")

