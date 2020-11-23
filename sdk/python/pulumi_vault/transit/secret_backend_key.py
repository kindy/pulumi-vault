# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['SecretBackendKey']


class SecretBackendKey(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_plaintext_backup: Optional[pulumi.Input[bool]] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 convergent_encryption: Optional[pulumi.Input[bool]] = None,
                 deletion_allowed: Optional[pulumi.Input[bool]] = None,
                 derived: Optional[pulumi.Input[bool]] = None,
                 exportable: Optional[pulumi.Input[bool]] = None,
                 min_decryption_version: Optional[pulumi.Input[int]] = None,
                 min_encryption_version: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates an Encryption Keyring on a Transit Secret Backend for Vault.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        transit = vault.Mount("transit",
            default_lease_ttl_seconds=3600,
            description="Example description",
            max_lease_ttl_seconds=86400,
            path="transit",
            type="transit")
        key = vault.transit.SecretBackendKey("key", backend=transit.path)
        ```

        ## Import

        Transit secret backend keys can be imported using the `path`, e.g.

        ```sh
         $ pulumi import vault:transit/secretBackendKey:SecretBackendKey key transit/keys/my_key
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_plaintext_backup: Enables taking backup of entire keyring in the plaintext format. Once set, this cannot be disabled.
               * Refer to Vault API documentation on key backups for more information: [Backup Key](https://www.vaultproject.io/api-docs/secret/transit#backup-key)
        :param pulumi.Input[str] backend: The path the transit secret backend is mounted at, with no leading or trailing `/`s.
        :param pulumi.Input[bool] convergent_encryption: Whether or not to support convergent encryption, where the same plaintext creates the same ciphertext. This requires `derived` to be set to `true`.
        :param pulumi.Input[bool] deletion_allowed: Specifies if the key is allowed to be deleted.
        :param pulumi.Input[bool] derived: Specifies if key derivation is to be used. If enabled, all encrypt/decrypt requests to this key must provide a context which is used for key derivation.
        :param pulumi.Input[bool] exportable: Enables keys to be exportable. This allows for all valid private keys in the keyring to be exported. Once set, this cannot be disabled.
        :param pulumi.Input[int] min_decryption_version: Minimum key version to use for decryption.
        :param pulumi.Input[int] min_encryption_version: Minimum key version to use for encryption
        :param pulumi.Input[str] name: The name to identify this key within the backend. Must be unique within the backend.
        :param pulumi.Input[str] type: Specifies the type of key to create. The currently-supported types are: `aes128-gcm96`, `aes256-gcm96` (default), `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `ecdsa-p384`, `ecdsa-p521`, `rsa-2048`, `rsa-3072` and `rsa-4096`. 
               * Refer to the Vault documentation on transit key types for more information: [Key Types](https://www.vaultproject.io/docs/secrets/transit#key-types)
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

            __props__['allow_plaintext_backup'] = allow_plaintext_backup
            if backend is None:
                raise TypeError("Missing required property 'backend'")
            __props__['backend'] = backend
            __props__['convergent_encryption'] = convergent_encryption
            __props__['deletion_allowed'] = deletion_allowed
            __props__['derived'] = derived
            __props__['exportable'] = exportable
            __props__['min_decryption_version'] = min_decryption_version
            __props__['min_encryption_version'] = min_encryption_version
            __props__['name'] = name
            __props__['type'] = type
            __props__['keys'] = None
            __props__['latest_version'] = None
            __props__['min_available_version'] = None
            __props__['supports_decryption'] = None
            __props__['supports_derivation'] = None
            __props__['supports_encryption'] = None
            __props__['supports_signing'] = None
        super(SecretBackendKey, __self__).__init__(
            'vault:transit/secretBackendKey:SecretBackendKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_plaintext_backup: Optional[pulumi.Input[bool]] = None,
            backend: Optional[pulumi.Input[str]] = None,
            convergent_encryption: Optional[pulumi.Input[bool]] = None,
            deletion_allowed: Optional[pulumi.Input[bool]] = None,
            derived: Optional[pulumi.Input[bool]] = None,
            exportable: Optional[pulumi.Input[bool]] = None,
            keys: Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, Any]]]]] = None,
            latest_version: Optional[pulumi.Input[int]] = None,
            min_available_version: Optional[pulumi.Input[int]] = None,
            min_decryption_version: Optional[pulumi.Input[int]] = None,
            min_encryption_version: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            supports_decryption: Optional[pulumi.Input[bool]] = None,
            supports_derivation: Optional[pulumi.Input[bool]] = None,
            supports_encryption: Optional[pulumi.Input[bool]] = None,
            supports_signing: Optional[pulumi.Input[bool]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'SecretBackendKey':
        """
        Get an existing SecretBackendKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_plaintext_backup: Enables taking backup of entire keyring in the plaintext format. Once set, this cannot be disabled.
               * Refer to Vault API documentation on key backups for more information: [Backup Key](https://www.vaultproject.io/api-docs/secret/transit#backup-key)
        :param pulumi.Input[str] backend: The path the transit secret backend is mounted at, with no leading or trailing `/`s.
        :param pulumi.Input[bool] convergent_encryption: Whether or not to support convergent encryption, where the same plaintext creates the same ciphertext. This requires `derived` to be set to `true`.
        :param pulumi.Input[bool] deletion_allowed: Specifies if the key is allowed to be deleted.
        :param pulumi.Input[bool] derived: Specifies if key derivation is to be used. If enabled, all encrypt/decrypt requests to this key must provide a context which is used for key derivation.
        :param pulumi.Input[bool] exportable: Enables keys to be exportable. This allows for all valid private keys in the keyring to be exported. Once set, this cannot be disabled.
        :param pulumi.Input[Sequence[pulumi.Input[Mapping[str, Any]]]] keys: List of key versions in the keyring. This attribute is zero-indexed and will contain a map of values depending on the `type` of the encryption key.
               * for key types `aes128-gcm96`, `aes256-gcm96` and `chacha20-poly1305`, each key version will be a map of a single value `id` which is just a hash of the key's metadata.
               * for key types `ed25519`, `ecdsa-p256`, `ecdsa-p384`, `ecdsa-p521`, `rsa-2048`, `rsa-3072` and `rsa-4096`, each key version will be a map of the following:
        :param pulumi.Input[int] latest_version: Latest key version available. This value is 1-indexed, so if `latest_version` is `1`, then the key's information can be referenced from `keys` by selecting element `0`
        :param pulumi.Input[int] min_available_version: Minimum key version available for use. If keys have been archived by increasing `min_decryption_version`, this attribute will reflect that change.
        :param pulumi.Input[int] min_decryption_version: Minimum key version to use for decryption.
        :param pulumi.Input[int] min_encryption_version: Minimum key version to use for encryption
        :param pulumi.Input[str] name: The name to identify this key within the backend. Must be unique within the backend.
        :param pulumi.Input[bool] supports_decryption: Whether or not the key supports decryption, based on key type.
        :param pulumi.Input[bool] supports_derivation: Whether or not the key supports derivation, based on key type.
        :param pulumi.Input[bool] supports_encryption: Whether or not the key supports encryption, based on key type.
        :param pulumi.Input[bool] supports_signing: Whether or not the key supports signing, based on key type.
        :param pulumi.Input[str] type: Specifies the type of key to create. The currently-supported types are: `aes128-gcm96`, `aes256-gcm96` (default), `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `ecdsa-p384`, `ecdsa-p521`, `rsa-2048`, `rsa-3072` and `rsa-4096`. 
               * Refer to the Vault documentation on transit key types for more information: [Key Types](https://www.vaultproject.io/docs/secrets/transit#key-types)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allow_plaintext_backup"] = allow_plaintext_backup
        __props__["backend"] = backend
        __props__["convergent_encryption"] = convergent_encryption
        __props__["deletion_allowed"] = deletion_allowed
        __props__["derived"] = derived
        __props__["exportable"] = exportable
        __props__["keys"] = keys
        __props__["latest_version"] = latest_version
        __props__["min_available_version"] = min_available_version
        __props__["min_decryption_version"] = min_decryption_version
        __props__["min_encryption_version"] = min_encryption_version
        __props__["name"] = name
        __props__["supports_decryption"] = supports_decryption
        __props__["supports_derivation"] = supports_derivation
        __props__["supports_encryption"] = supports_encryption
        __props__["supports_signing"] = supports_signing
        __props__["type"] = type
        return SecretBackendKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowPlaintextBackup")
    def allow_plaintext_backup(self) -> pulumi.Output[Optional[bool]]:
        """
        Enables taking backup of entire keyring in the plaintext format. Once set, this cannot be disabled.
        * Refer to Vault API documentation on key backups for more information: [Backup Key](https://www.vaultproject.io/api-docs/secret/transit#backup-key)
        """
        return pulumi.get(self, "allow_plaintext_backup")

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[str]:
        """
        The path the transit secret backend is mounted at, with no leading or trailing `/`s.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="convergentEncryption")
    def convergent_encryption(self) -> pulumi.Output[Optional[bool]]:
        """
        Whether or not to support convergent encryption, where the same plaintext creates the same ciphertext. This requires `derived` to be set to `true`.
        """
        return pulumi.get(self, "convergent_encryption")

    @property
    @pulumi.getter(name="deletionAllowed")
    def deletion_allowed(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies if the key is allowed to be deleted.
        """
        return pulumi.get(self, "deletion_allowed")

    @property
    @pulumi.getter
    def derived(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies if key derivation is to be used. If enabled, all encrypt/decrypt requests to this key must provide a context which is used for key derivation.
        """
        return pulumi.get(self, "derived")

    @property
    @pulumi.getter
    def exportable(self) -> pulumi.Output[Optional[bool]]:
        """
        Enables keys to be exportable. This allows for all valid private keys in the keyring to be exported. Once set, this cannot be disabled.
        """
        return pulumi.get(self, "exportable")

    @property
    @pulumi.getter
    def keys(self) -> pulumi.Output[Sequence[Mapping[str, Any]]]:
        """
        List of key versions in the keyring. This attribute is zero-indexed and will contain a map of values depending on the `type` of the encryption key.
        * for key types `aes128-gcm96`, `aes256-gcm96` and `chacha20-poly1305`, each key version will be a map of a single value `id` which is just a hash of the key's metadata.
        * for key types `ed25519`, `ecdsa-p256`, `ecdsa-p384`, `ecdsa-p521`, `rsa-2048`, `rsa-3072` and `rsa-4096`, each key version will be a map of the following:
        """
        return pulumi.get(self, "keys")

    @property
    @pulumi.getter(name="latestVersion")
    def latest_version(self) -> pulumi.Output[int]:
        """
        Latest key version available. This value is 1-indexed, so if `latest_version` is `1`, then the key's information can be referenced from `keys` by selecting element `0`
        """
        return pulumi.get(self, "latest_version")

    @property
    @pulumi.getter(name="minAvailableVersion")
    def min_available_version(self) -> pulumi.Output[int]:
        """
        Minimum key version available for use. If keys have been archived by increasing `min_decryption_version`, this attribute will reflect that change.
        """
        return pulumi.get(self, "min_available_version")

    @property
    @pulumi.getter(name="minDecryptionVersion")
    def min_decryption_version(self) -> pulumi.Output[Optional[int]]:
        """
        Minimum key version to use for decryption.
        """
        return pulumi.get(self, "min_decryption_version")

    @property
    @pulumi.getter(name="minEncryptionVersion")
    def min_encryption_version(self) -> pulumi.Output[Optional[int]]:
        """
        Minimum key version to use for encryption
        """
        return pulumi.get(self, "min_encryption_version")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name to identify this key within the backend. Must be unique within the backend.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="supportsDecryption")
    def supports_decryption(self) -> pulumi.Output[bool]:
        """
        Whether or not the key supports decryption, based on key type.
        """
        return pulumi.get(self, "supports_decryption")

    @property
    @pulumi.getter(name="supportsDerivation")
    def supports_derivation(self) -> pulumi.Output[bool]:
        """
        Whether or not the key supports derivation, based on key type.
        """
        return pulumi.get(self, "supports_derivation")

    @property
    @pulumi.getter(name="supportsEncryption")
    def supports_encryption(self) -> pulumi.Output[bool]:
        """
        Whether or not the key supports encryption, based on key type.
        """
        return pulumi.get(self, "supports_encryption")

    @property
    @pulumi.getter(name="supportsSigning")
    def supports_signing(self) -> pulumi.Output[bool]:
        """
        Whether or not the key supports signing, based on key type.
        """
        return pulumi.get(self, "supports_signing")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the type of key to create. The currently-supported types are: `aes128-gcm96`, `aes256-gcm96` (default), `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `ecdsa-p384`, `ecdsa-p521`, `rsa-2048`, `rsa-3072` and `rsa-4096`. 
        * Refer to the Vault documentation on transit key types for more information: [Key Types](https://www.vaultproject.io/docs/secrets/transit#key-types)
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

