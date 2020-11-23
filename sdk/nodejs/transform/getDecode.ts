// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source supports the "/transform/decode/{role_name}" Vault endpoint.
 *
 * It decodes the provided value using a named role.
 */
export function getDecode(args: GetDecodeArgs, opts?: pulumi.InvokeOptions): Promise<GetDecodeResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("vault:transform/getDecode:getDecode", {
        "batchInputs": args.batchInputs,
        "batchResults": args.batchResults,
        "decodedValue": args.decodedValue,
        "path": args.path,
        "roleName": args.roleName,
        "transformation": args.transformation,
        "tweak": args.tweak,
        "value": args.value,
    }, opts);
}

/**
 * A collection of arguments for invoking getDecode.
 */
export interface GetDecodeArgs {
    /**
     * Specifies a list of items to be decoded in a single batch. If this parameter is set, the top-level parameters 'value', 'transformation' and 'tweak' will be ignored. Each batch item within the list can specify these parameters instead.
     */
    readonly batchInputs?: {[key: string]: any}[];
    /**
     * The result of decoding a batch.
     */
    readonly batchResults?: {[key: string]: any}[];
    /**
     * The result of decoding a value.
     */
    readonly decodedValue?: string;
    /**
     * Path to where the back-end is mounted within Vault.
     */
    readonly path: string;
    /**
     * The name of the role.
     */
    readonly roleName: string;
    /**
     * The transformation to perform. If no value is provided and the role contains a single transformation, this value will be inferred from the role.
     */
    readonly transformation?: string;
    /**
     * The tweak value to use. Only applicable for FPE transformations
     */
    readonly tweak?: string;
    /**
     * The value in which to decode.
     */
    readonly value?: string;
}

/**
 * A collection of values returned by getDecode.
 */
export interface GetDecodeResult {
    readonly batchInputs?: {[key: string]: any}[];
    readonly batchResults: {[key: string]: any}[];
    readonly decodedValue: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly path: string;
    readonly roleName: string;
    readonly transformation?: string;
    readonly tweak?: string;
    readonly value?: string;
}
