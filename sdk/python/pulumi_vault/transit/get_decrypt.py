# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = [
    'GetDecryptResult',
    'AwaitableGetDecryptResult',
    'get_decrypt',
]

@pulumi.output_type
class GetDecryptResult:
    """
    A collection of values returned by getDecrypt.
    """
    def __init__(__self__, backend=None, ciphertext=None, context=None, id=None, key=None, plaintext=None):
        if backend and not isinstance(backend, str):
            raise TypeError("Expected argument 'backend' to be a str")
        pulumi.set(__self__, "backend", backend)
        if ciphertext and not isinstance(ciphertext, str):
            raise TypeError("Expected argument 'ciphertext' to be a str")
        pulumi.set(__self__, "ciphertext", ciphertext)
        if context and not isinstance(context, str):
            raise TypeError("Expected argument 'context' to be a str")
        pulumi.set(__self__, "context", context)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if key and not isinstance(key, str):
            raise TypeError("Expected argument 'key' to be a str")
        pulumi.set(__self__, "key", key)
        if plaintext and not isinstance(plaintext, str):
            raise TypeError("Expected argument 'plaintext' to be a str")
        pulumi.set(__self__, "plaintext", plaintext)

    @property
    @pulumi.getter
    def backend(self) -> str:
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter
    def ciphertext(self) -> str:
        return pulumi.get(self, "ciphertext")

    @property
    @pulumi.getter
    def context(self) -> Optional[str]:
        return pulumi.get(self, "context")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def plaintext(self) -> str:
        """
        Decrypted plaintext returned from Vault
        """
        return pulumi.get(self, "plaintext")


class AwaitableGetDecryptResult(GetDecryptResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDecryptResult(
            backend=self.backend,
            ciphertext=self.ciphertext,
            context=self.context,
            id=self.id,
            key=self.key,
            plaintext=self.plaintext)


def get_decrypt(backend: Optional[str] = None,
                ciphertext: Optional[str] = None,
                context: Optional[str] = None,
                key: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDecryptResult:
    """
    This is a data source which can be used to decrypt ciphertext using a Vault Transit key.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vault as vault

    test = vault.transit.get_decrypt(backend="transit",
        ciphertext="vault:v1:S3GtnJ5GUNCWV+/pdL9+g1Feu/nzAv+RlmTmE91Tu0rBkeIU8MEb2nSspC/1IQ==",
        key="test")
    ```


    :param str backend: The path the transit secret backend is mounted at, with no leading or trailing `/`.
    :param str ciphertext: Ciphertext to be decoded.
    :param str context: Context for key derivation. This is required if key derivation is enabled for this key.
    :param str key: Specifies the name of the transit key to decrypt against.
    """
    __args__ = dict()
    __args__['backend'] = backend
    __args__['ciphertext'] = ciphertext
    __args__['context'] = context
    __args__['key'] = key
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vault:transit/getDecrypt:getDecrypt', __args__, opts=opts, typ=GetDecryptResult).value

    return AwaitableGetDecryptResult(
        backend=__ret__.backend,
        ciphertext=__ret__.ciphertext,
        context=__ret__.context,
        id=__ret__.id,
        key=__ret__.key,
        plaintext=__ret__.plaintext)
