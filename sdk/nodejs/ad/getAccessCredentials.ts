// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getAccessCredentials(args: GetAccessCredentialsArgs, opts?: pulumi.InvokeOptions): Promise<GetAccessCredentialsResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("vault:ad/getAccessCredentials:getAccessCredentials", {
        "backend": args.backend,
        "role": args.role,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccessCredentials.
 */
export interface GetAccessCredentialsArgs {
    /**
     * The path to the AD secret backend to
     * read credentials from, with no leading or trailing `/`s.
     */
    backend: string;
    /**
     * The name of the AD secret backend role to read
     * credentials from, with no leading or trailing `/`s.
     */
    role: string;
}

/**
 * A collection of values returned by getAccessCredentials.
 */
export interface GetAccessCredentialsResult {
    readonly backend: string;
    /**
     * The current set password on the Active Directory service account.
     */
    readonly currentPassword: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The current set password on the Active Directory service account, provided because AD is eventually consistent.
     */
    readonly lastPassword: string;
    readonly role: string;
    /**
     * The Active Directory service account username.
     */
    readonly username: string;
}

export function getAccessCredentialsOutput(args: GetAccessCredentialsOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetAccessCredentialsResult> {
    return pulumi.output(args).apply(a => getAccessCredentials(a, opts))
}

/**
 * A collection of arguments for invoking getAccessCredentials.
 */
export interface GetAccessCredentialsOutputArgs {
    /**
     * The path to the AD secret backend to
     * read credentials from, with no leading or trailing `/`s.
     */
    backend: pulumi.Input<string>;
    /**
     * The name of the AD secret backend role to read
     * credentials from, with no leading or trailing `/`s.
     */
    role: pulumi.Input<string>;
}
