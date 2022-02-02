// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Creates an Encryption Keyring on a Transit Secret Backend for Vault.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const transit = new vault.Mount("transit", {
 *     path: "transit",
 *     type: "transit",
 *     description: "Example description",
 *     defaultLeaseTtlSeconds: 3600,
 *     maxLeaseTtlSeconds: 86400,
 * });
 * const key = new vault.transit.SecretBackendKey("key", {backend: transit.path});
 * ```
 *
 * ## Import
 *
 * Transit secret backend keys can be imported using the `path`, e.g.
 *
 * ```sh
 *  $ pulumi import vault:transit/secretBackendKey:SecretBackendKey key transit/keys/my_key
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
     * @param opts Optional settings to control the behavior of the CustomResource.
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
     * * for key types `aes128-gcm96`, `aes256-gcm96` and `chacha20-poly1305`, each key version will be a map of a single value `id` which is just a hash of the key's metadata.
     * * for key types `ed25519`, `ecdsa-p256`, `ecdsa-p384`, `ecdsa-p521`, `rsa-2048`, `rsa-3072` and `rsa-4096`, each key version will be a map of the following:
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
     * Specifies the type of key to create. The currently-supported types are: `aes128-gcm96`, `aes256-gcm96` (default), `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `ecdsa-p384`, `ecdsa-p521`, `rsa-2048`, `rsa-3072` and `rsa-4096`. 
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
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretBackendKeyState | undefined;
            resourceInputs["allowPlaintextBackup"] = state ? state.allowPlaintextBackup : undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["convergentEncryption"] = state ? state.convergentEncryption : undefined;
            resourceInputs["deletionAllowed"] = state ? state.deletionAllowed : undefined;
            resourceInputs["derived"] = state ? state.derived : undefined;
            resourceInputs["exportable"] = state ? state.exportable : undefined;
            resourceInputs["keys"] = state ? state.keys : undefined;
            resourceInputs["latestVersion"] = state ? state.latestVersion : undefined;
            resourceInputs["minAvailableVersion"] = state ? state.minAvailableVersion : undefined;
            resourceInputs["minDecryptionVersion"] = state ? state.minDecryptionVersion : undefined;
            resourceInputs["minEncryptionVersion"] = state ? state.minEncryptionVersion : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["supportsDecryption"] = state ? state.supportsDecryption : undefined;
            resourceInputs["supportsDerivation"] = state ? state.supportsDerivation : undefined;
            resourceInputs["supportsEncryption"] = state ? state.supportsEncryption : undefined;
            resourceInputs["supportsSigning"] = state ? state.supportsSigning : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as SecretBackendKeyArgs | undefined;
            if ((!args || args.backend === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backend'");
            }
            resourceInputs["allowPlaintextBackup"] = args ? args.allowPlaintextBackup : undefined;
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["convergentEncryption"] = args ? args.convergentEncryption : undefined;
            resourceInputs["deletionAllowed"] = args ? args.deletionAllowed : undefined;
            resourceInputs["derived"] = args ? args.derived : undefined;
            resourceInputs["exportable"] = args ? args.exportable : undefined;
            resourceInputs["minDecryptionVersion"] = args ? args.minDecryptionVersion : undefined;
            resourceInputs["minEncryptionVersion"] = args ? args.minEncryptionVersion : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["keys"] = undefined /*out*/;
            resourceInputs["latestVersion"] = undefined /*out*/;
            resourceInputs["minAvailableVersion"] = undefined /*out*/;
            resourceInputs["supportsDecryption"] = undefined /*out*/;
            resourceInputs["supportsDerivation"] = undefined /*out*/;
            resourceInputs["supportsEncryption"] = undefined /*out*/;
            resourceInputs["supportsSigning"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecretBackendKey.__pulumiType, name, resourceInputs, opts);
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
    allowPlaintextBackup?: pulumi.Input<boolean>;
    /**
     * The path the transit secret backend is mounted at, with no leading or trailing `/`s.
     */
    backend?: pulumi.Input<string>;
    /**
     * Whether or not to support convergent encryption, where the same plaintext creates the same ciphertext. This requires `derived` to be set to `true`.
     */
    convergentEncryption?: pulumi.Input<boolean>;
    /**
     * Specifies if the key is allowed to be deleted.
     */
    deletionAllowed?: pulumi.Input<boolean>;
    /**
     * Specifies if key derivation is to be used. If enabled, all encrypt/decrypt requests to this key must provide a context which is used for key derivation.
     */
    derived?: pulumi.Input<boolean>;
    /**
     * Enables keys to be exportable. This allows for all valid private keys in the keyring to be exported. Once set, this cannot be disabled.
     */
    exportable?: pulumi.Input<boolean>;
    /**
     * List of key versions in the keyring. This attribute is zero-indexed and will contain a map of values depending on the `type` of the encryption key.
     * * for key types `aes128-gcm96`, `aes256-gcm96` and `chacha20-poly1305`, each key version will be a map of a single value `id` which is just a hash of the key's metadata.
     * * for key types `ed25519`, `ecdsa-p256`, `ecdsa-p384`, `ecdsa-p521`, `rsa-2048`, `rsa-3072` and `rsa-4096`, each key version will be a map of the following:
     */
    keys?: pulumi.Input<pulumi.Input<{[key: string]: any}>[]>;
    /**
     * Latest key version available. This value is 1-indexed, so if `latestVersion` is `1`, then the key's information can be referenced from `keys` by selecting element `0`
     */
    latestVersion?: pulumi.Input<number>;
    /**
     * Minimum key version available for use. If keys have been archived by increasing `minDecryptionVersion`, this attribute will reflect that change.
     */
    minAvailableVersion?: pulumi.Input<number>;
    /**
     * Minimum key version to use for decryption.
     */
    minDecryptionVersion?: pulumi.Input<number>;
    /**
     * Minimum key version to use for encryption
     */
    minEncryptionVersion?: pulumi.Input<number>;
    /**
     * The name to identify this key within the backend. Must be unique within the backend.
     */
    name?: pulumi.Input<string>;
    /**
     * Whether or not the key supports decryption, based on key type.
     */
    supportsDecryption?: pulumi.Input<boolean>;
    /**
     * Whether or not the key supports derivation, based on key type.
     */
    supportsDerivation?: pulumi.Input<boolean>;
    /**
     * Whether or not the key supports encryption, based on key type.
     */
    supportsEncryption?: pulumi.Input<boolean>;
    /**
     * Whether or not the key supports signing, based on key type.
     */
    supportsSigning?: pulumi.Input<boolean>;
    /**
     * Specifies the type of key to create. The currently-supported types are: `aes128-gcm96`, `aes256-gcm96` (default), `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `ecdsa-p384`, `ecdsa-p521`, `rsa-2048`, `rsa-3072` and `rsa-4096`. 
     * * Refer to the Vault documentation on transit key types for more information: [Key Types](https://www.vaultproject.io/docs/secrets/transit#key-types)
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretBackendKey resource.
 */
export interface SecretBackendKeyArgs {
    /**
     * Enables taking backup of entire keyring in the plaintext format. Once set, this cannot be disabled.
     * * Refer to Vault API documentation on key backups for more information: [Backup Key](https://www.vaultproject.io/api-docs/secret/transit#backup-key)
     */
    allowPlaintextBackup?: pulumi.Input<boolean>;
    /**
     * The path the transit secret backend is mounted at, with no leading or trailing `/`s.
     */
    backend: pulumi.Input<string>;
    /**
     * Whether or not to support convergent encryption, where the same plaintext creates the same ciphertext. This requires `derived` to be set to `true`.
     */
    convergentEncryption?: pulumi.Input<boolean>;
    /**
     * Specifies if the key is allowed to be deleted.
     */
    deletionAllowed?: pulumi.Input<boolean>;
    /**
     * Specifies if key derivation is to be used. If enabled, all encrypt/decrypt requests to this key must provide a context which is used for key derivation.
     */
    derived?: pulumi.Input<boolean>;
    /**
     * Enables keys to be exportable. This allows for all valid private keys in the keyring to be exported. Once set, this cannot be disabled.
     */
    exportable?: pulumi.Input<boolean>;
    /**
     * Minimum key version to use for decryption.
     */
    minDecryptionVersion?: pulumi.Input<number>;
    /**
     * Minimum key version to use for encryption
     */
    minEncryptionVersion?: pulumi.Input<number>;
    /**
     * The name to identify this key within the backend. Must be unique within the backend.
     */
    name?: pulumi.Input<string>;
    /**
     * Specifies the type of key to create. The currently-supported types are: `aes128-gcm96`, `aes256-gcm96` (default), `chacha20-poly1305`, `ed25519`, `ecdsa-p256`, `ecdsa-p384`, `ecdsa-p521`, `rsa-2048`, `rsa-3072` and `rsa-4096`. 
     * * Refer to the Vault documentation on transit key types for more information: [Key Types](https://www.vaultproject.io/docs/secrets/transit#key-types)
     */
    type?: pulumi.Input<string>;
}
