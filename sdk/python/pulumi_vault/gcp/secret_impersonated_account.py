# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SecretImpersonatedAccountArgs', 'SecretImpersonatedAccount']

@pulumi.input_type
class SecretImpersonatedAccountArgs:
    def __init__(__self__, *,
                 backend: pulumi.Input[str],
                 impersonated_account: pulumi.Input[str],
                 service_account_email: pulumi.Input[str],
                 namespace: Optional[pulumi.Input[str]] = None,
                 token_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a SecretImpersonatedAccount resource.
        :param pulumi.Input[str] backend: Path where the GCP Secrets Engine is mounted
        :param pulumi.Input[str] impersonated_account: Name of the Impersonated Account to create
        :param pulumi.Input[str] service_account_email: Email of the GCP service account to impersonate.
        :param pulumi.Input[str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[Sequence[pulumi.Input[str]]] token_scopes: List of OAuth scopes to assign to access tokens generated under this impersonated account.
        """
        pulumi.set(__self__, "backend", backend)
        pulumi.set(__self__, "impersonated_account", impersonated_account)
        pulumi.set(__self__, "service_account_email", service_account_email)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if token_scopes is not None:
            pulumi.set(__self__, "token_scopes", token_scopes)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Input[str]:
        """
        Path where the GCP Secrets Engine is mounted
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: pulumi.Input[str]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="impersonatedAccount")
    def impersonated_account(self) -> pulumi.Input[str]:
        """
        Name of the Impersonated Account to create
        """
        return pulumi.get(self, "impersonated_account")

    @impersonated_account.setter
    def impersonated_account(self, value: pulumi.Input[str]):
        pulumi.set(self, "impersonated_account", value)

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> pulumi.Input[str]:
        """
        Email of the GCP service account to impersonate.
        """
        return pulumi.get(self, "service_account_email")

    @service_account_email.setter
    def service_account_email(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_account_email", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Target namespace. (requires Enterprise)
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="tokenScopes")
    def token_scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of OAuth scopes to assign to access tokens generated under this impersonated account.
        """
        return pulumi.get(self, "token_scopes")

    @token_scopes.setter
    def token_scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "token_scopes", value)


@pulumi.input_type
class _SecretImpersonatedAccountState:
    def __init__(__self__, *,
                 backend: Optional[pulumi.Input[str]] = None,
                 impersonated_account: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 service_account_email: Optional[pulumi.Input[str]] = None,
                 service_account_project: Optional[pulumi.Input[str]] = None,
                 token_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering SecretImpersonatedAccount resources.
        :param pulumi.Input[str] backend: Path where the GCP Secrets Engine is mounted
        :param pulumi.Input[str] impersonated_account: Name of the Impersonated Account to create
        :param pulumi.Input[str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[str] service_account_email: Email of the GCP service account to impersonate.
        :param pulumi.Input[str] service_account_project: Project the service account belongs to.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] token_scopes: List of OAuth scopes to assign to access tokens generated under this impersonated account.
        """
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if impersonated_account is not None:
            pulumi.set(__self__, "impersonated_account", impersonated_account)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if service_account_email is not None:
            pulumi.set(__self__, "service_account_email", service_account_email)
        if service_account_project is not None:
            pulumi.set(__self__, "service_account_project", service_account_project)
        if token_scopes is not None:
            pulumi.set(__self__, "token_scopes", token_scopes)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[str]]:
        """
        Path where the GCP Secrets Engine is mounted
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="impersonatedAccount")
    def impersonated_account(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Impersonated Account to create
        """
        return pulumi.get(self, "impersonated_account")

    @impersonated_account.setter
    def impersonated_account(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "impersonated_account", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Target namespace. (requires Enterprise)
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> Optional[pulumi.Input[str]]:
        """
        Email of the GCP service account to impersonate.
        """
        return pulumi.get(self, "service_account_email")

    @service_account_email.setter
    def service_account_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account_email", value)

    @property
    @pulumi.getter(name="serviceAccountProject")
    def service_account_project(self) -> Optional[pulumi.Input[str]]:
        """
        Project the service account belongs to.
        """
        return pulumi.get(self, "service_account_project")

    @service_account_project.setter
    def service_account_project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account_project", value)

    @property
    @pulumi.getter(name="tokenScopes")
    def token_scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of OAuth scopes to assign to access tokens generated under this impersonated account.
        """
        return pulumi.get(self, "token_scopes")

    @token_scopes.setter
    def token_scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "token_scopes", value)


class SecretImpersonatedAccount(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 impersonated_account: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 service_account_email: Optional[pulumi.Input[str]] = None,
                 token_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Creates a Impersonated Account in the [GCP Secrets Engine](https://www.vaultproject.io/docs/secrets/gcp/index.html) for Vault.

        Each [impersonated account](https://www.vaultproject.io/docs/secrets/gcp/index.html#impersonated-accounts) is tied to a separately managed
        Service Account.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_gcp as gcp
        import pulumi_vault as vault

        this = gcp.service_account.Account("this", account_id="my-awesome-account")
        gcp = vault.gcp.SecretBackend("gcp",
            path="gcp",
            credentials=(lambda path: open(path).read())("credentials.json"))
        impersonated_account = vault.gcp.SecretImpersonatedAccount("impersonatedAccount",
            backend=gcp.path,
            impersonated_account="this",
            service_account_email=this.email,
            token_scopes=["https://www.googleapis.com/auth/cloud-platform"])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        A impersonated account can be imported using its Vault Path. For example, referencing the example above,

        ```sh
        $ pulumi import vault:gcp/secretImpersonatedAccount:SecretImpersonatedAccount impersonated_account gcp/impersonated-account/project_viewer
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: Path where the GCP Secrets Engine is mounted
        :param pulumi.Input[str] impersonated_account: Name of the Impersonated Account to create
        :param pulumi.Input[str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[str] service_account_email: Email of the GCP service account to impersonate.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] token_scopes: List of OAuth scopes to assign to access tokens generated under this impersonated account.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretImpersonatedAccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a Impersonated Account in the [GCP Secrets Engine](https://www.vaultproject.io/docs/secrets/gcp/index.html) for Vault.

        Each [impersonated account](https://www.vaultproject.io/docs/secrets/gcp/index.html#impersonated-accounts) is tied to a separately managed
        Service Account.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_gcp as gcp
        import pulumi_vault as vault

        this = gcp.service_account.Account("this", account_id="my-awesome-account")
        gcp = vault.gcp.SecretBackend("gcp",
            path="gcp",
            credentials=(lambda path: open(path).read())("credentials.json"))
        impersonated_account = vault.gcp.SecretImpersonatedAccount("impersonatedAccount",
            backend=gcp.path,
            impersonated_account="this",
            service_account_email=this.email,
            token_scopes=["https://www.googleapis.com/auth/cloud-platform"])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        A impersonated account can be imported using its Vault Path. For example, referencing the example above,

        ```sh
        $ pulumi import vault:gcp/secretImpersonatedAccount:SecretImpersonatedAccount impersonated_account gcp/impersonated-account/project_viewer
        ```

        :param str resource_name: The name of the resource.
        :param SecretImpersonatedAccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretImpersonatedAccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 impersonated_account: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 service_account_email: Optional[pulumi.Input[str]] = None,
                 token_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretImpersonatedAccountArgs.__new__(SecretImpersonatedAccountArgs)

            if backend is None and not opts.urn:
                raise TypeError("Missing required property 'backend'")
            __props__.__dict__["backend"] = backend
            if impersonated_account is None and not opts.urn:
                raise TypeError("Missing required property 'impersonated_account'")
            __props__.__dict__["impersonated_account"] = impersonated_account
            __props__.__dict__["namespace"] = namespace
            if service_account_email is None and not opts.urn:
                raise TypeError("Missing required property 'service_account_email'")
            __props__.__dict__["service_account_email"] = service_account_email
            __props__.__dict__["token_scopes"] = token_scopes
            __props__.__dict__["service_account_project"] = None
        super(SecretImpersonatedAccount, __self__).__init__(
            'vault:gcp/secretImpersonatedAccount:SecretImpersonatedAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend: Optional[pulumi.Input[str]] = None,
            impersonated_account: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            service_account_email: Optional[pulumi.Input[str]] = None,
            service_account_project: Optional[pulumi.Input[str]] = None,
            token_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'SecretImpersonatedAccount':
        """
        Get an existing SecretImpersonatedAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: Path where the GCP Secrets Engine is mounted
        :param pulumi.Input[str] impersonated_account: Name of the Impersonated Account to create
        :param pulumi.Input[str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[str] service_account_email: Email of the GCP service account to impersonate.
        :param pulumi.Input[str] service_account_project: Project the service account belongs to.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] token_scopes: List of OAuth scopes to assign to access tokens generated under this impersonated account.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretImpersonatedAccountState.__new__(_SecretImpersonatedAccountState)

        __props__.__dict__["backend"] = backend
        __props__.__dict__["impersonated_account"] = impersonated_account
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["service_account_email"] = service_account_email
        __props__.__dict__["service_account_project"] = service_account_project
        __props__.__dict__["token_scopes"] = token_scopes
        return SecretImpersonatedAccount(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[str]:
        """
        Path where the GCP Secrets Engine is mounted
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="impersonatedAccount")
    def impersonated_account(self) -> pulumi.Output[str]:
        """
        Name of the Impersonated Account to create
        """
        return pulumi.get(self, "impersonated_account")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[str]]:
        """
        Target namespace. (requires Enterprise)
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> pulumi.Output[str]:
        """
        Email of the GCP service account to impersonate.
        """
        return pulumi.get(self, "service_account_email")

    @property
    @pulumi.getter(name="serviceAccountProject")
    def service_account_project(self) -> pulumi.Output[str]:
        """
        Project the service account belongs to.
        """
        return pulumi.get(self, "service_account_project")

    @property
    @pulumi.getter(name="tokenScopes")
    def token_scopes(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of OAuth scopes to assign to access tokens generated under this impersonated account.
        """
        return pulumi.get(self, "token_scopes")

