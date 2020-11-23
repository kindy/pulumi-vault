# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['SecretRoleset']


class SecretRoleset(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 bindings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretRolesetBindingArgs']]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 roleset: Optional[pulumi.Input[str]] = None,
                 secret_type: Optional[pulumi.Input[str]] = None,
                 token_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a Roleset in the [GCP Secrets Engine](https://www.vaultproject.io/docs/secrets/gcp/index.html) for Vault.

        Each Roleset is [tied](https://www.vaultproject.io/docs/secrets/gcp/index.html#service-accounts-are-tied-to-rolesets) to a Service Account, and can have one or more [bindings](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings) associated with it.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        project = "my-awesome-project"
        gcp = vault.gcp.SecretBackend("gcp",
            credentials=(lambda path: open(path).read())("credentials.json"),
            path="gcp")
        roleset = vault.gcp.SecretRoleset("roleset",
            backend=gcp.path,
            bindings=[vault.gcp.SecretRolesetBindingArgs(
                resource=f"//cloudresourcemanager.googleapis.com/projects/{project}",
                roles=["roles/viewer"],
            )],
            project=project,
            roleset="project_viewer",
            secret_type="access_token",
            token_scopes=["https://www.googleapis.com/auth/cloud-platform"])
        ```

        ## Import

        A roleset can be imported using its Vault Path. For example, referencing the example above,

        ```sh
         $ pulumi import vault:gcp/secretRoleset:SecretRoleset roleset gcp/roleset/project_viewer
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: Path where the GCP Secrets Engine is mounted
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretRolesetBindingArgs']]]] bindings: Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
        :param pulumi.Input[str] project: Name of the GCP project that this roleset's service account will belong to.
        :param pulumi.Input[str] roleset: Name of the Roleset to create
        :param pulumi.Input[str] secret_type: Type of secret generated for this role set. Accepted values: `access_token`, `service_account_key`. Defaults to `access_token`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] token_scopes: List of OAuth scopes to assign to `access_token` secrets generated under this role set (`access_token` role sets only).
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

            if backend is None:
                raise TypeError("Missing required property 'backend'")
            __props__['backend'] = backend
            if bindings is None:
                raise TypeError("Missing required property 'bindings'")
            __props__['bindings'] = bindings
            if project is None:
                raise TypeError("Missing required property 'project'")
            __props__['project'] = project
            if roleset is None:
                raise TypeError("Missing required property 'roleset'")
            __props__['roleset'] = roleset
            __props__['secret_type'] = secret_type
            __props__['token_scopes'] = token_scopes
            __props__['service_account_email'] = None
        super(SecretRoleset, __self__).__init__(
            'vault:gcp/secretRoleset:SecretRoleset',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend: Optional[pulumi.Input[str]] = None,
            bindings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretRolesetBindingArgs']]]]] = None,
            project: Optional[pulumi.Input[str]] = None,
            roleset: Optional[pulumi.Input[str]] = None,
            secret_type: Optional[pulumi.Input[str]] = None,
            service_account_email: Optional[pulumi.Input[str]] = None,
            token_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'SecretRoleset':
        """
        Get an existing SecretRoleset resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: Path where the GCP Secrets Engine is mounted
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretRolesetBindingArgs']]]] bindings: Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
        :param pulumi.Input[str] project: Name of the GCP project that this roleset's service account will belong to.
        :param pulumi.Input[str] roleset: Name of the Roleset to create
        :param pulumi.Input[str] secret_type: Type of secret generated for this role set. Accepted values: `access_token`, `service_account_key`. Defaults to `access_token`.
        :param pulumi.Input[str] service_account_email: Email of the service account created by Vault for this Roleset
        :param pulumi.Input[Sequence[pulumi.Input[str]]] token_scopes: List of OAuth scopes to assign to `access_token` secrets generated under this role set (`access_token` role sets only).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backend"] = backend
        __props__["bindings"] = bindings
        __props__["project"] = project
        __props__["roleset"] = roleset
        __props__["secret_type"] = secret_type
        __props__["service_account_email"] = service_account_email
        __props__["token_scopes"] = token_scopes
        return SecretRoleset(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[str]:
        """
        Path where the GCP Secrets Engine is mounted
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter
    def bindings(self) -> pulumi.Output[Sequence['outputs.SecretRolesetBinding']]:
        """
        Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
        """
        return pulumi.get(self, "bindings")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        Name of the GCP project that this roleset's service account will belong to.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def roleset(self) -> pulumi.Output[str]:
        """
        Name of the Roleset to create
        """
        return pulumi.get(self, "roleset")

    @property
    @pulumi.getter(name="secretType")
    def secret_type(self) -> pulumi.Output[str]:
        """
        Type of secret generated for this role set. Accepted values: `access_token`, `service_account_key`. Defaults to `access_token`.
        """
        return pulumi.get(self, "secret_type")

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> pulumi.Output[str]:
        """
        Email of the service account created by Vault for this Roleset
        """
        return pulumi.get(self, "service_account_email")

    @property
    @pulumi.getter(name="tokenScopes")
    def token_scopes(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of OAuth scopes to assign to `access_token` secrets generated under this role set (`access_token` role sets only).
        """
        return pulumi.get(self, "token_scopes")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

