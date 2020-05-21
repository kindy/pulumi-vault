// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates an Encryption Keyring on a Transit Secret Backend for Vault.
 *
 * ## Example Usage
 *
 *
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const transit = new vault.Mount("transit", {
 *     defaultLeaseTtlSeconds: 3600,
 *     description: "Example description",
 *     maxLeaseTtlSeconds: 86400,
 *     path: "transit",
 *     type: "transit",
 * });
 * const key = new vault.transit.SecretBackendKey("key", {
 *     backend: transit.path,
 * });
 * ```
 */
export class SecretBackendKey extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackendKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendKeyState, opts?: pulumi.CustomResourceOptions): SecretBackendKey {
        return new SecretBackendKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:transit/secretBackendKey:SecretBackendKey';

    /**
     * Returns true if the given object is an instance of SecretBackendKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackendKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackendKey.__pulumiType;
    }

    /**
     * Enables taking backup of entire keyring in the plaintext format. Once set, this cannot be disabled.
     * * Refer to Vault API documentation on key backups for more information: [Backup Key](https://www.vaultproject.io/api-docs/secret/transit#backup-key)
     */
    public readonly allowPlaintextBackup!: pulumi.Output<boolean | undefined>;
    /**
     * The path the transit secret backend is mounted at, with no leading or trailing `/`s.
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * Whether or not to support convergent encryption, where the same plaintext creates the same ciphertext. This requires `derived` to be set to `true`.
     */
    public readonly convergentEncryption!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies if the key is allowed to be deleted.
     */
    public readonly deletionAllowed!: pulumi.Output<boolean | undefined>;
    /**
     * Specifies if key derivation is to be used. If enabled, all encrypt/decrypt requests to this key must provide a context which is used for key derivation.
     */
    public readonly derived!: pulumi.Output<boolean | undefined>;
    /**
     * Enables keys to be exportable. This allows for all valid private keys in the keyring to be exported. Once set, this cannot be disabled.
     */
    public readonly exportable!: pulumi.Output<boolean | undefined>;
    /**
     * List of key versions in the keyring. This attribute is zero-indexed and will contain a map of values depending on the `type` of the encryption key.
     * * for key types `aes256-gcm96` and `chacha20-poly1305`, each key version will be a map of a single value `id` which is just a hash of the key's metadata.
     * * for key types `ed25519`, `ecdsa-p256`, `rsa-2048` and `rsa-4096`, each key version will be a map of the following:
     */
    public /*out*/ readonly keys!: pulumi.Output<{[key: string]: any}[]>;
    /**
     * Latest key version available. This value is 1-indexed, so if `latestVersion` is `1`, then the key's information can be referenced from `keys` by selecting element `0`
     */
    public /*out*/ readonly latestVersion!: pulumi.Output<number>;
    /**
     * Minimum key version available for use. If keys have been archived by increasing `minDecryptionVersion`, this attribute will reflect that change.
     */
    public /*out*/ readonly minAvailableVersion!: pulumi.Output<number>;
    /**
     * Minimum key version to use for decryption.
     */
    public readonly minDecryptionVersion!: pulumi.Output<number | undefined>;
    /**
     * Minimum key version to use for encryption
     */
    public readonly minEncryptionVersion!: pulumi.Output<number | undefined>;
    /**
     * The name to identify this key within the backend. Must be unique within the backend.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Whether or not the key supports decryption, based on key type.
     */
    public /*out*/ readonly supportsDecryption!: pulumi.Output<boolean>;
    /**
     * Whether or not the key supports derivation, based on key type.
     */
    public /*out*/ readonly supportsDerivation!: pulumi.Output<boolean>;
    /**
     * Whether or not the key supports encryption, based on key type.
     */
    public /*out*/ readonly supportsEncryption!: pulumi.Output<boolean>;
    /**
     * Whether or not the key supports signing, based on key type.
     */
    public /*out*/ readonly supportsSigning!: pulumi.Output<boolean>;
    /**
     * Specifies the type of key to create. The currently-supported types are: `aes256-gcm96` (default), `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `rsa-2048` and `rsa-4096`. 
     * * Refer to the Vault documentation on transit key types for more information: [Key Types](https://www.vaultproject.io/docs/secrets/transit#key-types)
     */
    public readonly type!: pulumi.Output<string | undefined>;

    /**
     * Create a SecretBackendKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretBackendKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendKeyArgs | SecretBackendKeyState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecretBackendKeyState | undefined;
            inputs["allowPlaintextBackup"] = state ? state.allowPlaintextBackup : undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["convergentEncryption"] = state ? state.convergentEncryption : undefined;
            inputs["deletionAllowed"] = state ? state.deletionAllowed : undefined;
            inputs["derived"] = state ? state.derived : undefined;
            inputs["exportable"] = state ? state.exportable : undefined;
            inputs["keys"] = state ? state.keys : undefined;
            inputs["latestVersion"] = state ? state.latestVersion : undefined;
            inputs["minAvailableVersion"] = state ? state.minAvailableVersion : undefined;
            inputs["minDecryptionVersion"] = state ? state.minDecryptionVersion : undefined;
            inputs["minEncryptionVersion"] = state ? state.minEncryptionVersion : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["supportsDecryption"] = state ? state.supportsDecryption : undefined;
            inputs["supportsDerivation"] = state ? state.supportsDerivation : undefined;
            inputs["supportsEncryption"] = state ? state.supportsEncryption : undefined;
            inputs["supportsSigning"] = state ? state.supportsSigning : undefined;
            inputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as SecretBackendKeyArgs | undefined;
            if (!args || args.backend === undefined) {
                throw new Error("Missing required property 'backend'");
            }
            inputs["allowPlaintextBackup"] = args ? args.allowPlaintextBackup : undefined;
            inputs["backend"] = args ? args.backend : undefined;
            inputs["convergentEncryption"] = args ? args.convergentEncryption : undefined;
            inputs["deletionAllowed"] = args ? args.deletionAllowed : undefined;
            inputs["derived"] = args ? args.derived : undefined;
            inputs["exportable"] = args ? args.exportable : undefined;
            inputs["minDecryptionVersion"] = args ? args.minDecryptionVersion : undefined;
            inputs["minEncryptionVersion"] = args ? args.minEncryptionVersion : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["type"] = args ? args.type : undefined;
            inputs["keys"] = undefined /*out*/;
            inputs["latestVersion"] = undefined /*out*/;
            inputs["minAvailableVersion"] = undefined /*out*/;
            inputs["supportsDecryption"] = undefined /*out*/;
            inputs["supportsDerivation"] = undefined /*out*/;
            inputs["supportsEncryption"] = undefined /*out*/;
            inputs["supportsSigning"] = undefined /*out*/;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SecretBackendKey.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackendKey resources.
 */
export interface SecretBackendKeyState {
    /**
     * Enables taking backup of entire keyring in the plaintext format. Once set, this cannot be disabled.
     * * Refer to Vault API documentation on key backups for more information: [Backup Key](https://www.vaultproject.io/api-docs/secret/transit#backup-key)
     */
    readonly allowPlaintextBackup?: pulumi.Input<boolean>;
    /**
     * The path the transit secret backend is mounted at, with no leading or trailing `/`s.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * Whether or not to support convergent encryption, where the same plaintext creates the same ciphertext. This requires `derived` to be set to `true`.
     */
    readonly convergentEncryption?: pulumi.Input<boolean>;
    /**
     * Specifies if the key is allowed to be deleted.
     */
    readonly deletionAllowed?: pulumi.Input<boolean>;
    /**
     * Specifies if key derivation is to be used. If enabled, all encrypt/decrypt requests to this key must provide a context which is used for key derivation.
     */
    readonly derived?: pulumi.Input<boolean>;
    /**
     * Enables keys to be exportable. This allows for all valid private keys in the keyring to be exported. Once set, this cannot be disabled.
     */
    readonly exportable?: pulumi.Input<boolean>;
    /**
     * List of key versions in the keyring. This attribute is zero-indexed and will contain a map of values depending on the `type` of the encryption key.
     * * for key types `aes256-gcm96` and `chacha20-poly1305`, each key version will be a map of a single value `id` which is just a hash of the key's metadata.
     * * for key types `ed25519`, `ecdsa-p256`, `rsa-2048` and `rsa-4096`, each key version will be a map of the following:
     */
    readonly keys?: pulumi.Input<pulumi.Input<{[key: string]: any}>[]>;
    /**
     * Latest key version available. This value is 1-indexed, so if `latestVersion` is `1`, then the key's information can be referenced from `keys` by selecting element `0`
     */
    readonly latestVersion?: pulumi.Input<number>;
    /**
     * Minimum key version available for use. If keys have been archived by increasing `minDecryptionVersion`, this attribute will reflect that change.
     */
    readonly minAvailableVersion?: pulumi.Input<number>;
    /**
     * Minimum key version to use for decryption.
     */
    readonly minDecryptionVersion?: pulumi.Input<number>;
    /**
     * Minimum key version to use for encryption
     */
    readonly minEncryptionVersion?: pulumi.Input<number>;
    /**
     * The name to identify this key within the backend. Must be unique within the backend.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Whether or not the key supports decryption, based on key type.
     */
    readonly supportsDecryption?: pulumi.Input<boolean>;
    /**
     * Whether or not the key supports derivation, based on key type.
     */
    readonly supportsDerivation?: pulumi.Input<boolean>;
    /**
     * Whether or not the key supports encryption, based on key type.
     */
    readonly supportsEncryption?: pulumi.Input<boolean>;
    /**
     * Whether or not the key supports signing, based on key type.
     */
    readonly supportsSigning?: pulumi.Input<boolean>;
    /**
     * Specifies the type of key to create. The currently-supported types are: `aes256-gcm96` (default), `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `rsa-2048` and `rsa-4096`. 
     * * Refer to the Vault documentation on transit key types for more information: [Key Types](https://www.vaultproject.io/docs/secrets/transit#key-types)
     */
    readonly type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretBackendKey resource.
 */
export interface SecretBackendKeyArgs {
    /**
     * Enables taking backup of entire keyring in the plaintext format. Once set, this cannot be disabled.
     * * Refer to Vault API documentation on key backups for more information: [Backup Key](https://www.vaultproject.io/api-docs/secret/transit#backup-key)
     */
    readonly allowPlaintextBackup?: pulumi.Input<boolean>;
    /**
     * The path the transit secret backend is mounted at, with no leading or trailing `/`s.
     */
    readonly backend: pulumi.Input<string>;
    /**
     * Whether or not to support convergent encryption, where the same plaintext creates the same ciphertext. This requires `derived` to be set to `true`.
     */
    readonly convergentEncryption?: pulumi.Input<boolean>;
    /**
     * Specifies if the key is allowed to be deleted.
     */
    readonly deletionAllowed?: pulumi.Input<boolean>;
    /**
     * Specifies if key derivation is to be used. If enabled, all encrypt/decrypt requests to this key must provide a context which is used for key derivation.
     */
    readonly derived?: pulumi.Input<boolean>;
    /**
     * Enables keys to be exportable. This allows for all valid private keys in the keyring to be exported. Once set, this cannot be disabled.
     */
    readonly exportable?: pulumi.Input<boolean>;
    /**
     * Minimum key version to use for decryption.
     */
    readonly minDecryptionVersion?: pulumi.Input<number>;
    /**
     * Minimum key version to use for encryption
     */
    readonly minEncryptionVersion?: pulumi.Input<number>;
    /**
     * The name to identify this key within the backend. Must be unique within the backend.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * Specifies the type of key to create. The currently-supported types are: `aes256-gcm96` (default), `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `rsa-2048` and `rsa-4096`. 
     * * Refer to the Vault documentation on transit key types for more information: [Key Types](https://www.vaultproject.io/docs/secrets/transit#key-types)
     */
    readonly type?: pulumi.Input<string>;
}
