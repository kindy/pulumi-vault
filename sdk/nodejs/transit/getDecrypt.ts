// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

export function getDecrypt(args: GetDecryptArgs, opts?: pulumi.InvokeOptions): Promise<GetDecryptResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("vault:transit/getDecrypt:getDecrypt", {
        "backend": args.backend,
        "ciphertext": args.ciphertext,
        "context": args.context,
        "key": args.key,
    }, opts);
}

/**
 * A collection of arguments for invoking getDecrypt.
 */
export interface GetDecryptArgs {
    readonly backend: string;
    readonly ciphertext: string;
    readonly context?: string;
    readonly key: string;
}

/**
 * A collection of values returned by getDecrypt.
 */
export interface GetDecryptResult {
    readonly backend: string;
    readonly ciphertext: string;
    readonly context?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly key: string;
    readonly plaintext: string;
}
