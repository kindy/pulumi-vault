# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class AuthBackendRole(pulumi.CustomResource):
    allowed_redirect_uris: pulumi.Output[list]
    """
    The list of allowed values for redirect_uri during OIDC logins.
    Required for OIDC roles
    """
    backend: pulumi.Output[str]
    """
    The unique name of the auth backend to configure.
    Defaults to `jwt`.
    """
    bound_audiences: pulumi.Output[list]
    """
    List of `aud` claims to match
    against. Any match is sufficient.
    """
    bound_cidrs: pulumi.Output[list]
    """
    If set, a list of
    CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
    """
    bound_claims: pulumi.Output[dict]
    """
    If set, a map of claims/values to match against.
    The expected value may be a single string or a list of strings.
    """
    bound_subject: pulumi.Output[str]
    """
    If set, requires that the `sub` claim matches
    this value.
    """
    claim_mappings: pulumi.Output[dict]
    """
    If set, a map of claims (keys) to be copied
    to specified metadata fields (values).
    """
    groups_claim: pulumi.Output[str]
    """
    The claim to use to uniquely identify
    the set of groups to which the user belongs; this will be used as the names
    for the Identity group aliases created due to a successful login. The claim
    value must be a list of strings.
    """
    groups_claim_delimiter_pattern: pulumi.Output[str]
    """
    (Optional; Deprecated. This field has been
    removed since Vault 1.1. If the groups claim is not at the top level, it can
    now be specified as a [JSONPointer](https://tools.ietf.org/html/rfc6901).)
    A pattern of delimiters
    used to allow the groups_claim to live outside of the top-level JWT structure.
    For instance, a groups_claim of meta/user.name/groups with this field
    set to // will expect nested structures named meta, user.name, and groups.
    If this field was set to /./ the groups information would expect to be
    via nested structures of meta, user, name, and groups.
    """
    max_ttl: pulumi.Output[float]
    """
    The maximum allowed lifetime of tokens
    issued using this role, provided as a number of seconds.
    """
    num_uses: pulumi.Output[float]
    """
    If set, puts a use-count
    limitation on the issued token.
    """
    oidc_scopes: pulumi.Output[list]
    """
    If set, a list of OIDC scopes to be used with an OIDC role.
    The standard scope "openid" is automatically included and need not be specified.
    """
    period: pulumi.Output[float]
    """
    If set, indicates that the
    token generated using this role should never expire. The token should be renewed within the
    duration specified by this value. At each renewal, the token's TTL will be set to the
    value of this field. The maximum allowed lifetime of token issued using this
    role. Specified as a number of seconds.
    """
    policies: pulumi.Output[list]
    """
    An array of strings
    specifying the policies to be set on tokens issued using this role.
    """
    role_name: pulumi.Output[str]
    """
    The name of the role.
    """
    role_type: pulumi.Output[str]
    """
    Type of role, either "oidc" (default) or "jwt".
    """
    token_bound_cidrs: pulumi.Output[list]
    """
    List of CIDR blocks; if set, specifies blocks of IP
    addresses which can authenticate successfully, and ties the resulting token to these blocks
    as well.
    """
    token_explicit_max_ttl: pulumi.Output[float]
    """
    If set, will encode an
    [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
    onto the token in number of seconds. This is a hard cap even if `token_ttl` and
    `token_max_ttl` would otherwise allow a renewal.
    """
    token_max_ttl: pulumi.Output[float]
    """
    The maximum lifetime for generated tokens in number of seconds.
    Its current value will be referenced at renewal time.
    """
    token_no_default_policy: pulumi.Output[bool]
    """
    If set, the default policy will not be set on
    generated tokens; otherwise it will be added to the policies set in token_policies.
    """
    token_num_uses: pulumi.Output[float]
    """
    The
    [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
    if any, in number of seconds to set on the token.
    """
    token_period: pulumi.Output[float]
    token_policies: pulumi.Output[list]
    """
    List of policies to encode onto generated tokens. Depending
    on the auth method, this list may be supplemented by user/group/other values.
    """
    token_ttl: pulumi.Output[float]
    """
    The incremental lifetime for generated tokens in number of seconds.
    Its current value will be referenced at renewal time.
    """
    token_type: pulumi.Output[str]
    """
    The type of token that should be generated. Can be `service`,
    `batch`, or `default` to use the mount's tuned default (which unless changed will be
    `service` tokens). For token store roles, there are two additional possibilities:
    `default-service` and `default-batch` which specify the type to return unless the client
    requests a different type at generation time.
    """
    ttl: pulumi.Output[float]
    """
    The TTL period of tokens issued
    using this role, provided as a number of seconds.
    """
    user_claim: pulumi.Output[str]
    """
    The claim to use to uniquely identify
    the user; this will be used as the name for the Identity entity alias created
    due to a successful login.
    """
    def __init__(__self__, resource_name, opts=None, allowed_redirect_uris=None, backend=None, bound_audiences=None, bound_cidrs=None, bound_claims=None, bound_subject=None, claim_mappings=None, groups_claim=None, groups_claim_delimiter_pattern=None, max_ttl=None, num_uses=None, oidc_scopes=None, period=None, policies=None, role_name=None, role_type=None, token_bound_cidrs=None, token_explicit_max_ttl=None, token_max_ttl=None, token_no_default_policy=None, token_num_uses=None, token_period=None, token_policies=None, token_ttl=None, token_type=None, ttl=None, user_claim=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages an JWT/OIDC auth backend role in a Vault server. See the [Vault
        documentation](https://www.vaultproject.io/docs/auth/jwt.html) for more
        information.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] allowed_redirect_uris: The list of allowed values for redirect_uri during OIDC logins.
               Required for OIDC roles
        :param pulumi.Input[str] backend: The unique name of the auth backend to configure.
               Defaults to `jwt`.
        :param pulumi.Input[list] bound_audiences: List of `aud` claims to match
               against. Any match is sufficient.
        :param pulumi.Input[list] bound_cidrs: If set, a list of
               CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
        :param pulumi.Input[dict] bound_claims: If set, a map of claims/values to match against.
               The expected value may be a single string or a list of strings.
        :param pulumi.Input[str] bound_subject: If set, requires that the `sub` claim matches
               this value.
        :param pulumi.Input[dict] claim_mappings: If set, a map of claims (keys) to be copied
               to specified metadata fields (values).
        :param pulumi.Input[str] groups_claim: The claim to use to uniquely identify
               the set of groups to which the user belongs; this will be used as the names
               for the Identity group aliases created due to a successful login. The claim
               value must be a list of strings.
        :param pulumi.Input[str] groups_claim_delimiter_pattern: (Optional; Deprecated. This field has been
               removed since Vault 1.1. If the groups claim is not at the top level, it can
               now be specified as a [JSONPointer](https://tools.ietf.org/html/rfc6901).)
               A pattern of delimiters
               used to allow the groups_claim to live outside of the top-level JWT structure.
               For instance, a groups_claim of meta/user.name/groups with this field
               set to // will expect nested structures named meta, user.name, and groups.
               If this field was set to /./ the groups information would expect to be
               via nested structures of meta, user, name, and groups.
        :param pulumi.Input[float] max_ttl: The maximum allowed lifetime of tokens
               issued using this role, provided as a number of seconds.
        :param pulumi.Input[float] num_uses: If set, puts a use-count
               limitation on the issued token.
        :param pulumi.Input[list] oidc_scopes: If set, a list of OIDC scopes to be used with an OIDC role.
               The standard scope "openid" is automatically included and need not be specified.
        :param pulumi.Input[float] period: If set, indicates that the
               token generated using this role should never expire. The token should be renewed within the
               duration specified by this value. At each renewal, the token's TTL will be set to the
               value of this field. The maximum allowed lifetime of token issued using this
               role. Specified as a number of seconds.
        :param pulumi.Input[list] policies: An array of strings
               specifying the policies to be set on tokens issued using this role.
        :param pulumi.Input[str] role_name: The name of the role.
        :param pulumi.Input[str] role_type: Type of role, either "oidc" (default) or "jwt".
        :param pulumi.Input[list] token_bound_cidrs: List of CIDR blocks; if set, specifies blocks of IP
               addresses which can authenticate successfully, and ties the resulting token to these blocks
               as well.
        :param pulumi.Input[float] token_explicit_max_ttl: If set, will encode an
               [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
               onto the token in number of seconds. This is a hard cap even if `token_ttl` and
               `token_max_ttl` would otherwise allow a renewal.
        :param pulumi.Input[float] token_max_ttl: The maximum lifetime for generated tokens in number of seconds.
               Its current value will be referenced at renewal time.
        :param pulumi.Input[bool] token_no_default_policy: If set, the default policy will not be set on
               generated tokens; otherwise it will be added to the policies set in token_policies.
        :param pulumi.Input[float] token_num_uses: The
               [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
               if any, in number of seconds to set on the token.
        :param pulumi.Input[list] token_policies: List of policies to encode onto generated tokens. Depending
               on the auth method, this list may be supplemented by user/group/other values.
        :param pulumi.Input[float] token_ttl: The incremental lifetime for generated tokens in number of seconds.
               Its current value will be referenced at renewal time.
        :param pulumi.Input[str] token_type: The type of token that should be generated. Can be `service`,
               `batch`, or `default` to use the mount's tuned default (which unless changed will be
               `service` tokens). For token store roles, there are two additional possibilities:
               `default-service` and `default-batch` which specify the type to return unless the client
               requests a different type at generation time.
        :param pulumi.Input[float] ttl: The TTL period of tokens issued
               using this role, provided as a number of seconds.
        :param pulumi.Input[str] user_claim: The claim to use to uniquely identify
               the user; this will be used as the name for the Identity entity alias created
               due to a successful login.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/jwt_auth_backend_role.html.markdown.
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

            __props__['allowed_redirect_uris'] = allowed_redirect_uris
            __props__['backend'] = backend
            if bound_audiences is None:
                raise TypeError("Missing required property 'bound_audiences'")
            __props__['bound_audiences'] = bound_audiences
            __props__['bound_cidrs'] = bound_cidrs
            __props__['bound_claims'] = bound_claims
            __props__['bound_subject'] = bound_subject
            __props__['claim_mappings'] = claim_mappings
            __props__['groups_claim'] = groups_claim
            __props__['groups_claim_delimiter_pattern'] = groups_claim_delimiter_pattern
            __props__['max_ttl'] = max_ttl
            __props__['num_uses'] = num_uses
            __props__['oidc_scopes'] = oidc_scopes
            __props__['period'] = period
            __props__['policies'] = policies
            if role_name is None:
                raise TypeError("Missing required property 'role_name'")
            __props__['role_name'] = role_name
            __props__['role_type'] = role_type
            __props__['token_bound_cidrs'] = token_bound_cidrs
            __props__['token_explicit_max_ttl'] = token_explicit_max_ttl
            __props__['token_max_ttl'] = token_max_ttl
            __props__['token_no_default_policy'] = token_no_default_policy
            __props__['token_num_uses'] = token_num_uses
            __props__['token_period'] = token_period
            __props__['token_policies'] = token_policies
            __props__['token_ttl'] = token_ttl
            __props__['token_type'] = token_type
            __props__['ttl'] = ttl
            if user_claim is None:
                raise TypeError("Missing required property 'user_claim'")
            __props__['user_claim'] = user_claim
        super(AuthBackendRole, __self__).__init__(
            'vault:jwt/authBackendRole:AuthBackendRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, allowed_redirect_uris=None, backend=None, bound_audiences=None, bound_cidrs=None, bound_claims=None, bound_subject=None, claim_mappings=None, groups_claim=None, groups_claim_delimiter_pattern=None, max_ttl=None, num_uses=None, oidc_scopes=None, period=None, policies=None, role_name=None, role_type=None, token_bound_cidrs=None, token_explicit_max_ttl=None, token_max_ttl=None, token_no_default_policy=None, token_num_uses=None, token_period=None, token_policies=None, token_ttl=None, token_type=None, ttl=None, user_claim=None):
        """
        Get an existing AuthBackendRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[list] allowed_redirect_uris: The list of allowed values for redirect_uri during OIDC logins.
               Required for OIDC roles
        :param pulumi.Input[str] backend: The unique name of the auth backend to configure.
               Defaults to `jwt`.
        :param pulumi.Input[list] bound_audiences: List of `aud` claims to match
               against. Any match is sufficient.
        :param pulumi.Input[list] bound_cidrs: If set, a list of
               CIDRs valid as the source address for login requests. This value is also encoded into any resulting token.
        :param pulumi.Input[dict] bound_claims: If set, a map of claims/values to match against.
               The expected value may be a single string or a list of strings.
        :param pulumi.Input[str] bound_subject: If set, requires that the `sub` claim matches
               this value.
        :param pulumi.Input[dict] claim_mappings: If set, a map of claims (keys) to be copied
               to specified metadata fields (values).
        :param pulumi.Input[str] groups_claim: The claim to use to uniquely identify
               the set of groups to which the user belongs; this will be used as the names
               for the Identity group aliases created due to a successful login. The claim
               value must be a list of strings.
        :param pulumi.Input[str] groups_claim_delimiter_pattern: (Optional; Deprecated. This field has been
               removed since Vault 1.1. If the groups claim is not at the top level, it can
               now be specified as a [JSONPointer](https://tools.ietf.org/html/rfc6901).)
               A pattern of delimiters
               used to allow the groups_claim to live outside of the top-level JWT structure.
               For instance, a groups_claim of meta/user.name/groups with this field
               set to // will expect nested structures named meta, user.name, and groups.
               If this field was set to /./ the groups information would expect to be
               via nested structures of meta, user, name, and groups.
        :param pulumi.Input[float] max_ttl: The maximum allowed lifetime of tokens
               issued using this role, provided as a number of seconds.
        :param pulumi.Input[float] num_uses: If set, puts a use-count
               limitation on the issued token.
        :param pulumi.Input[list] oidc_scopes: If set, a list of OIDC scopes to be used with an OIDC role.
               The standard scope "openid" is automatically included and need not be specified.
        :param pulumi.Input[float] period: If set, indicates that the
               token generated using this role should never expire. The token should be renewed within the
               duration specified by this value. At each renewal, the token's TTL will be set to the
               value of this field. The maximum allowed lifetime of token issued using this
               role. Specified as a number of seconds.
        :param pulumi.Input[list] policies: An array of strings
               specifying the policies to be set on tokens issued using this role.
        :param pulumi.Input[str] role_name: The name of the role.
        :param pulumi.Input[str] role_type: Type of role, either "oidc" (default) or "jwt".
        :param pulumi.Input[list] token_bound_cidrs: List of CIDR blocks; if set, specifies blocks of IP
               addresses which can authenticate successfully, and ties the resulting token to these blocks
               as well.
        :param pulumi.Input[float] token_explicit_max_ttl: If set, will encode an
               [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
               onto the token in number of seconds. This is a hard cap even if `token_ttl` and
               `token_max_ttl` would otherwise allow a renewal.
        :param pulumi.Input[float] token_max_ttl: The maximum lifetime for generated tokens in number of seconds.
               Its current value will be referenced at renewal time.
        :param pulumi.Input[bool] token_no_default_policy: If set, the default policy will not be set on
               generated tokens; otherwise it will be added to the policies set in token_policies.
        :param pulumi.Input[float] token_num_uses: The
               [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
               if any, in number of seconds to set on the token.
        :param pulumi.Input[list] token_policies: List of policies to encode onto generated tokens. Depending
               on the auth method, this list may be supplemented by user/group/other values.
        :param pulumi.Input[float] token_ttl: The incremental lifetime for generated tokens in number of seconds.
               Its current value will be referenced at renewal time.
        :param pulumi.Input[str] token_type: The type of token that should be generated. Can be `service`,
               `batch`, or `default` to use the mount's tuned default (which unless changed will be
               `service` tokens). For token store roles, there are two additional possibilities:
               `default-service` and `default-batch` which specify the type to return unless the client
               requests a different type at generation time.
        :param pulumi.Input[float] ttl: The TTL period of tokens issued
               using this role, provided as a number of seconds.
        :param pulumi.Input[str] user_claim: The claim to use to uniquely identify
               the user; this will be used as the name for the Identity entity alias created
               due to a successful login.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/jwt_auth_backend_role.html.markdown.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()
        __props__["allowed_redirect_uris"] = allowed_redirect_uris
        __props__["backend"] = backend
        __props__["bound_audiences"] = bound_audiences
        __props__["bound_cidrs"] = bound_cidrs
        __props__["bound_claims"] = bound_claims
        __props__["bound_subject"] = bound_subject
        __props__["claim_mappings"] = claim_mappings
        __props__["groups_claim"] = groups_claim
        __props__["groups_claim_delimiter_pattern"] = groups_claim_delimiter_pattern
        __props__["max_ttl"] = max_ttl
        __props__["num_uses"] = num_uses
        __props__["oidc_scopes"] = oidc_scopes
        __props__["period"] = period
        __props__["policies"] = policies
        __props__["role_name"] = role_name
        __props__["role_type"] = role_type
        __props__["token_bound_cidrs"] = token_bound_cidrs
        __props__["token_explicit_max_ttl"] = token_explicit_max_ttl
        __props__["token_max_ttl"] = token_max_ttl
        __props__["token_no_default_policy"] = token_no_default_policy
        __props__["token_num_uses"] = token_num_uses
        __props__["token_period"] = token_period
        __props__["token_policies"] = token_policies
        __props__["token_ttl"] = token_ttl
        __props__["token_type"] = token_type
        __props__["ttl"] = ttl
        __props__["user_claim"] = user_claim
        return AuthBackendRole(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

