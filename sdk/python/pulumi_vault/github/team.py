# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TeamArgs', 'Team']

@pulumi.input_type
class TeamArgs:
    def __init__(__self__, *,
                 team: pulumi.Input[str],
                 backend: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a Team resource.
        :param pulumi.Input[str] team: GitHub team name in "slugified" format.
        :param pulumi.Input[str] backend: Path where the github auth backend is mounted. Defaults to `github`
               if not specified.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: An array of strings specifying the policies to be set on tokens
               issued using this role.
        """
        pulumi.set(__self__, "team", team)
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)

    @property
    @pulumi.getter
    def team(self) -> pulumi.Input[str]:
        """
        GitHub team name in "slugified" format.
        """
        return pulumi.get(self, "team")

    @team.setter
    def team(self, value: pulumi.Input[str]):
        pulumi.set(self, "team", value)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[str]]:
        """
        Path where the github auth backend is mounted. Defaults to `github`
        if not specified.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of strings specifying the policies to be set on tokens
        issued using this role.
        """
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "policies", value)


@pulumi.input_type
class _TeamState:
    def __init__(__self__, *,
                 backend: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 team: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Team resources.
        :param pulumi.Input[str] backend: Path where the github auth backend is mounted. Defaults to `github`
               if not specified.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: An array of strings specifying the policies to be set on tokens
               issued using this role.
        :param pulumi.Input[str] team: GitHub team name in "slugified" format.
        """
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)
        if team is not None:
            pulumi.set(__self__, "team", team)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[str]]:
        """
        Path where the github auth backend is mounted. Defaults to `github`
        if not specified.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of strings specifying the policies to be set on tokens
        issued using this role.
        """
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "policies", value)

    @property
    @pulumi.getter
    def team(self) -> Optional[pulumi.Input[str]]:
        """
        GitHub team name in "slugified" format.
        """
        return pulumi.get(self, "team")

    @team.setter
    def team(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "team", value)


class Team(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 team: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages policy mappings for Github Teams authenticated via Github. See the [Vault
        documentation](https://www.vaultproject.io/docs/auth/github/) for more
        information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.github.AuthBackend("example", organization="myorg")
        tf_devs = vault.github.Team("tfDevs",
            backend=example.id,
            team="terraform-developers",
            policies=[
                "developer",
                "read-only",
            ])
        ```

        ## Import

        Github team mappings can be imported using the `path`, e.g.

        ```sh
         $ pulumi import vault:github/team:Team tf_devs auth/github/map/teams/terraform-developers
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: Path where the github auth backend is mounted. Defaults to `github`
               if not specified.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: An array of strings specifying the policies to be set on tokens
               issued using this role.
        :param pulumi.Input[str] team: GitHub team name in "slugified" format.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TeamArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages policy mappings for Github Teams authenticated via Github. See the [Vault
        documentation](https://www.vaultproject.io/docs/auth/github/) for more
        information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.github.AuthBackend("example", organization="myorg")
        tf_devs = vault.github.Team("tfDevs",
            backend=example.id,
            team="terraform-developers",
            policies=[
                "developer",
                "read-only",
            ])
        ```

        ## Import

        Github team mappings can be imported using the `path`, e.g.

        ```sh
         $ pulumi import vault:github/team:Team tf_devs auth/github/map/teams/terraform-developers
        ```

        :param str resource_name: The name of the resource.
        :param TeamArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TeamArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 team: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TeamArgs.__new__(TeamArgs)

            __props__.__dict__["backend"] = backend
            __props__.__dict__["policies"] = policies
            if team is None and not opts.urn:
                raise TypeError("Missing required property 'team'")
            __props__.__dict__["team"] = team
        super(Team, __self__).__init__(
            'vault:github/team:Team',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend: Optional[pulumi.Input[str]] = None,
            policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            team: Optional[pulumi.Input[str]] = None) -> 'Team':
        """
        Get an existing Team resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: Path where the github auth backend is mounted. Defaults to `github`
               if not specified.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: An array of strings specifying the policies to be set on tokens
               issued using this role.
        :param pulumi.Input[str] team: GitHub team name in "slugified" format.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TeamState.__new__(_TeamState)

        __props__.__dict__["backend"] = backend
        __props__.__dict__["policies"] = policies
        __props__.__dict__["team"] = team
        return Team(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[Optional[str]]:
        """
        Path where the github auth backend is mounted. Defaults to `github`
        if not specified.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        An array of strings specifying the policies to be set on tokens
        issued using this role.
        """
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter
    def team(self) -> pulumi.Output[str]:
        """
        GitHub team name in "slugified" format.
        """
        return pulumi.get(self, "team")

