# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class GetDecryptResult:
    """
    A collection of values returned by getDecrypt.
    """
    def __init__(__self__, backend=None, ciphertext=None, context=None, id=None, key=None, plaintext=None):
        if backend and not isinstance(backend, str):
            raise TypeError("Expected argument 'backend' to be a str")
        __self__.backend = backend
        if ciphertext and not isinstance(ciphertext, str):
            raise TypeError("Expected argument 'ciphertext' to be a str")
        __self__.ciphertext = ciphertext
        if context and not isinstance(context, str):
            raise TypeError("Expected argument 'context' to be a str")
        __self__.context = context
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        __self__.id = id
        """
        The provider-assigned unique ID for this managed resource.
        """
        if key and not isinstance(key, str):
            raise TypeError("Expected argument 'key' to be a str")
        __self__.key = key
        if plaintext and not isinstance(plaintext, str):
            raise TypeError("Expected argument 'plaintext' to be a str")
        __self__.plaintext = plaintext
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

def get_decrypt(backend=None,ciphertext=None,context=None,key=None,opts=None):
    """
    Use this data source to access information about an existing resource.
    """
    __args__ = dict()


    __args__['backend'] = backend
    __args__['ciphertext'] = ciphertext
    __args__['context'] = context
    __args__['key'] = key
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vault:transit/getDecrypt:getDecrypt', __args__, opts=opts).value

    return AwaitableGetDecryptResult(
        backend=__ret__.get('backend'),
        ciphertext=__ret__.get('ciphertext'),
        context=__ret__.get('context'),
        id=__ret__.get('id'),
        key=__ret__.get('key'),
        plaintext=__ret__.get('plaintext'))
