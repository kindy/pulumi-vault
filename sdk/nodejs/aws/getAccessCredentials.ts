// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";
import * as utilities from "../utilities";

export function getAccessCredentials(args: GetAccessCredentialsArgs, opts?: pulumi.InvokeOptions): Promise<GetAccessCredentialsResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("vault:aws/getAccessCredentials:getAccessCredentials", {
        "backend": args.backend,
        "region": args.region,
        "role": args.role,
        "roleArn": args.roleArn,
        "type": args.type,
    }, opts);
}

/**
 * A collection of arguments for invoking getAccessCredentials.
 */
export interface GetAccessCredentialsArgs {
    /**
     * The path to the AWS secret backend to
     * read credentials from, with no leading or trailing `/`s.
     */
    readonly backend: string;
    readonly region?: string;
    /**
     * The name of the AWS secret backend role to read
     * credentials from, with no leading or trailing `/`s.
     */
    readonly role: string;
    /**
     * The specific AWS ARN to use
     * from the configured role. If the role does not have multiple ARNs, this does
     * not need to be specified.
     */
    readonly roleArn?: string;
    /**
     * The type of credentials to read. Defaults
     * to `"creds"`, which just returns an AWS Access Key ID and Secret
     * Key. Can also be set to `"sts"`, which will return a security token
     * in addition to the keys.
     */
    readonly type?: string;
}

/**
 * A collection of values returned by getAccessCredentials.
 */
export interface GetAccessCredentialsResult {
    /**
     * The AWS Access Key ID returned by Vault.
     */
    readonly accessKey: string;
    readonly backend: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The duration of the secret lease, in seconds relative
     * to the time the data was requested. Once this time has passed any plan
     * generated with this data may fail to apply.
     */
    readonly leaseDuration: number;
    /**
     * The lease identifier assigned by Vault.
     */
    readonly leaseId: string;
    readonly leaseRenewable: boolean;
    readonly leaseStartTime: string;
    readonly region?: string;
    readonly role: string;
    readonly roleArn?: string;
    /**
     * The AWS Secret Key returned by Vault.
     */
    readonly secretKey: string;
    /**
     * The STS token returned by Vault, if any.
     */
    readonly securityToken: string;
    readonly type?: string;
}
