// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

/**
 * This data source supports the "/transform/encode/{role_name}" Vault endpoint.
 *
 * It encodes the provided value using a named role.
 */
export function getEncode(args: GetEncodeArgs, opts?: pulumi.InvokeOptions): Promise<GetEncodeResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("vault:transform/getEncode:getEncode", {
        "batchInputs": args.batchInputs,
        "batchResults": args.batchResults,
        "encodedValue": args.encodedValue,
        "path": args.path,
        "roleName": args.roleName,
        "transformation": args.transformation,
        "tweak": args.tweak,
        "value": args.value,
    }, opts);
}

/**
 * A collection of arguments for invoking getEncode.
 */
export interface GetEncodeArgs {
    /**
     * Specifies a list of items to be encoded in a single batch. If this parameter is set, the parameters 'value', 'transformation' and 'tweak' will be ignored. Each batch item within the list can specify these parameters instead.
     */
    readonly batchInputs?: {[key: string]: any}[];
    /**
     * The result of encoding a batch.
     */
    readonly batchResults?: {[key: string]: any}[];
    /**
     * The result of encoding a value.
     */
    readonly encodedValue?: string;
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
     * The value in which to encode.
     */
    readonly value?: string;
}

/**
 * A collection of values returned by getEncode.
 */
export interface GetEncodeResult {
    readonly batchInputs?: {[key: string]: any}[];
    readonly batchResults: {[key: string]: any}[];
    readonly encodedValue: string;
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
