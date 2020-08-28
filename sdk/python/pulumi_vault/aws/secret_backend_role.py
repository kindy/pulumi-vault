# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables


class SecretBackendRole(pulumi.CustomResource):
    backend: pulumi.Output[str]
    """
    The path the AWS secret backend is mounted at,
    with no leading or trailing `/`s.
    """
    credential_type: pulumi.Output[str]
    """
    Specifies the type of credential to be used when
    retrieving credentials from the role. Must be one of `iam_user`, `assumed_role`, or
    `federation_token`.
    """
    default_sts_ttl: pulumi.Output[float]
    """
    The default TTL in seconds for STS credentials.
    When a TTL is not specified when STS credentials are requested,
    and a default TTL is specified on the role,
    then this default TTL will be used. Valid only when `credential_type` is one of
    `assumed_role` or `federation_token`.
    """
    iam_groups: pulumi.Output[list]
    """
    A list of IAM group names. IAM users generated
    against this vault role will be added to these IAM Groups. For a credential
    type of `assumed_role` or `federation_token`, the policies sent to the
    corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the
    policies from each group in `iam_groups` combined with the `policy_document`
    and `policy_arns` parameters.
    """
    max_sts_ttl: pulumi.Output[float]
    """
    The max allowed TTL in seconds for STS credentials
    (credentials TTL are capped to `max_sts_ttl`). Valid only when `credential_type` is
    one of `assumed_role` or `federation_token`.
    """
    name: pulumi.Output[str]
    """
    The name to identify this role within the backend.
    Must be unique within the backend.
    """
    policy_arns: pulumi.Output[list]
    """
    Specifies a list of AWS managed policy ARNs. The
    behavior depends on the credential type. With `iam_user`, the policies will be
    attached to IAM users when they are requested. With `assumed_role` and
    `federation_token`, the policy ARNs will act as a filter on what the credentials
    can do, similar to `policy_document`. When `credential_type` is `iam_user` or
    `federation_token`, at least one of `policy_document` or `policy_arns` must
    be specified.
    """
    policy_document: pulumi.Output[str]
    """
    The IAM policy document for the role. The
    behavior depends on the credential type. With `iam_user`, the policy document
    will be attached to the IAM user generated and augment the permissions the IAM
    user has. With `assumed_role` and `federation_token`, the policy document will
    act as a filter on what the credentials can do, similar to `policy_arns`.
    """
    role_arns: pulumi.Output[list]
    """
    Specifies the ARNs of the AWS roles this Vault role
    is allowed to assume. Required when `credential_type` is `assumed_role` and
    prohibited otherwise.
    """
    def __init__(__self__, resource_name, opts=None, backend=None, credential_type=None, default_sts_ttl=None, iam_groups=None, max_sts_ttl=None, name=None, policy_arns=None, policy_document=None, role_arns=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a SecretBackendRole resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The path the AWS secret backend is mounted at,
               with no leading or trailing `/`s.
        :param pulumi.Input[str] credential_type: Specifies the type of credential to be used when
               retrieving credentials from the role. Must be one of `iam_user`, `assumed_role`, or
               `federation_token`.
        :param pulumi.Input[float] default_sts_ttl: The default TTL in seconds for STS credentials.
               When a TTL is not specified when STS credentials are requested,
               and a default TTL is specified on the role,
               then this default TTL will be used. Valid only when `credential_type` is one of
               `assumed_role` or `federation_token`.
        :param pulumi.Input[list] iam_groups: A list of IAM group names. IAM users generated
               against this vault role will be added to these IAM Groups. For a credential
               type of `assumed_role` or `federation_token`, the policies sent to the
               corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the
               policies from each group in `iam_groups` combined with the `policy_document`
               and `policy_arns` parameters.
        :param pulumi.Input[float] max_sts_ttl: The max allowed TTL in seconds for STS credentials
               (credentials TTL are capped to `max_sts_ttl`). Valid only when `credential_type` is
               one of `assumed_role` or `federation_token`.
        :param pulumi.Input[str] name: The name to identify this role within the backend.
               Must be unique within the backend.
        :param pulumi.Input[list] policy_arns: Specifies a list of AWS managed policy ARNs. The
               behavior depends on the credential type. With `iam_user`, the policies will be
               attached to IAM users when they are requested. With `assumed_role` and
               `federation_token`, the policy ARNs will act as a filter on what the credentials
               can do, similar to `policy_document`. When `credential_type` is `iam_user` or
               `federation_token`, at least one of `policy_document` or `policy_arns` must
               be specified.
        :param pulumi.Input[str] policy_document: The IAM policy document for the role. The
               behavior depends on the credential type. With `iam_user`, the policy document
               will be attached to the IAM user generated and augment the permissions the IAM
               user has. With `assumed_role` and `federation_token`, the policy document will
               act as a filter on what the credentials can do, similar to `policy_arns`.
        :param pulumi.Input[list] role_arns: Specifies the ARNs of the AWS roles this Vault role
               is allowed to assume. Required when `credential_type` is `assumed_role` and
               prohibited otherwise.
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if backend is None:
                raise TypeError("Missing required property 'backend'")
            __props__['backend'] = backend
            if credential_type is None:
                raise TypeError("Missing required property 'credential_type'")
            __props__['credential_type'] = credential_type
            __props__['default_sts_ttl'] = default_sts_ttl
            __props__['iam_groups'] = iam_groups
            __props__['max_sts_ttl'] = max_sts_ttl
            __props__['name'] = name
            __props__['policy_arns'] = policy_arns
            __props__['policy_document'] = policy_document
            __props__['role_arns'] = role_arns
        super(SecretBackendRole, __self__).__init__(
            'vault:aws/secretBackendRole:SecretBackendRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, backend=None, credential_type=None, default_sts_ttl=None, iam_groups=None, max_sts_ttl=None, name=None, policy_arns=None, policy_document=None, role_arns=None):
        """
        Get an existing SecretBackendRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The path the AWS secret backend is mounted at,
               with no leading or trailing `/`s.
        :param pulumi.Input[str] credential_type: Specifies the type of credential to be used when
               retrieving credentials from the role. Must be one of `iam_user`, `assumed_role`, or
               `federation_token`.
        :param pulumi.Input[float] default_sts_ttl: The default TTL in seconds for STS credentials.
               When a TTL is not specified when STS credentials are requested,
               and a default TTL is specified on the role,
               then this default TTL will be used. Valid only when `credential_type` is one of
               `assumed_role` or `federation_token`.
        :param pulumi.Input[list] iam_groups: A list of IAM group names. IAM users generated
               against this vault role will be added to these IAM Groups. For a credential
               type of `assumed_role` or `federation_token`, the policies sent to the
               corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the
               policies from each group in `iam_groups` combined with the `policy_document`
               and `policy_arns` parameters.
        :param pulumi.Input[float] max_sts_ttl: The max allowed TTL in seconds for STS credentials
               (credentials TTL are capped to `max_sts_ttl`). Valid only when `credential_type` is
               one of `assumed_role` or `federation_token`.
        :param pulumi.Input[str] name: The name to identify this role within the backend.
               Must be unique within the backend.
        :param pulumi.Input[list] policy_arns: Specifies a list of AWS managed policy ARNs. The
               behavior depends on the credential type. With `iam_user`, the policies will be
               attached to IAM users when they are requested. With `assumed_role` and
               `federation_token`, the policy ARNs will act as a filter on what the credentials
               can do, similar to `policy_document`. When `credential_type` is `iam_user` or
               `federation_token`, at least one of `policy_document` or `policy_arns` must
               be specified.
        :param pulumi.Input[str] policy_document: The IAM policy document for the role. The
               behavior depends on the credential type. With `iam_user`, the policy document
               will be attached to the IAM user generated and augment the permissions the IAM
               user has. With `assumed_role` and `federation_token`, the policy document will
               act as a filter on what the credentials can do, similar to `policy_arns`.
        :param pulumi.Input[list] role_arns: Specifies the ARNs of the AWS roles this Vault role
               is allowed to assume. Required when `credential_type` is `assumed_role` and
               prohibited otherwise.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backend"] = backend
        __props__["credential_type"] = credential_type
        __props__["default_sts_ttl"] = default_sts_ttl
        __props__["iam_groups"] = iam_groups
        __props__["max_sts_ttl"] = max_sts_ttl
        __props__["name"] = name
        __props__["policy_arns"] = policy_arns
        __props__["policy_document"] = policy_document
        __props__["role_arns"] = role_arns
        return SecretBackendRole(resource_name, opts=opts, __props__=__props__)

    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
