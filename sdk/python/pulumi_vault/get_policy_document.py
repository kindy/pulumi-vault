# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class GetPolicyDocumentResult:
    """
    A collection of values returned by getPolicyDocument.
    """
    def __init__(__self__, hcl=None, id=None, rules=None):
        if hcl and not isinstance(hcl, str):
            raise TypeError("Expected argument 'hcl' to be a str")
        __self__.hcl = hcl
        """
        The above arguments serialized as a standard Vault HCL policy document.
        """
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        __self__.rules = rules
class AwaitableGetPolicyDocumentResult(GetPolicyDocumentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPolicyDocumentResult(
            hcl=self.hcl,
            id=self.id,
            rules=self.rules)

def get_policy_document(rules=None,opts=None):
    """
    This is a data source which can be used to construct a HCL representation of an Vault policy document, for use with resources which expect policy documents, such as the `.Policy` resource.

    ## Example Usage



    ```python
    import pulumi
    import pulumi_vault as vault

    example_policy_document = vault.get_policy_document(rules=[{
        "capabilities": [
            "create",
            "read",
            "update",
            "delete",
            "list",
        ],
        "description": "allow all on secrets",
        "path": "secret/*",
    }])
    example_policy = vault.Policy("examplePolicy", policy=example_policy_document.hcl)
    ```




    The **rules** object supports the following:

      * `allowedParameters` (`list`) - Whitelists a list of keys and values that are permitted on the given path. See Parameters below.
        * `key` (`str`) - name of permitted or denied parameter.
        * `values` (`list`) - list of values what are permitted or denied by policy rule.

      * `capabilities` (`list`) - A list of capabilities that this rule apply to `path`. For example, ["read", "write"].
      * `deniedParameters` (`list`) - Blacklists a list of parameter and values. Any values specified here take precedence over `allowed_parameter`. See Parameters below.
        * `key` (`str`) - name of permitted or denied parameter.
        * `values` (`list`) - list of values what are permitted or denied by policy rule.

      * `description` (`str`) - Description of the rule. Will be added as a commend to rendered rule.
      * `maxWrappingTtl` (`str`) - The maximum allowed TTL that clients can specify for a wrapped response.
      * `minWrappingTtl` (`str`) - The minimum allowed TTL that clients can specify for a wrapped response.
      * `path` (`str`) - A path in Vault that this rule applies to.
      * `requiredParameters` (`list`) - A list of parameters that must be specified.
    """
    __args__ = dict()


    __args__['rules'] = rules
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vault:index/getPolicyDocument:getPolicyDocument', __args__, opts=opts).value

    return AwaitableGetPolicyDocumentResult(
        hcl=__ret__.get('hcl'),
        id=__ret__.get('id'),
        rules=__ret__.get('rules'))
