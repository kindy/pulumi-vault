// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export function getBackendIssuer(args: GetBackendIssuerArgs, opts?: pulumi.InvokeOptions): Promise<GetBackendIssuerResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vault:pkiSecret/getBackendIssuer:getBackendIssuer", {
        "backend": args.backend,
        "issuerRef": args.issuerRef,
        "namespace": args.namespace,
    }, opts);
}

/**
 * A collection of arguments for invoking getBackendIssuer.
 */
export interface GetBackendIssuerArgs {
    /**
     * The path to the PKI secret backend to
     * read the issuer from, with no leading or trailing `/`s.
     */
    backend: string;
    /**
     * Reference to an existing issuer.
     */
    issuerRef: string;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: string;
}

/**
 * A collection of values returned by getBackendIssuer.
 */
export interface GetBackendIssuerResult {
    readonly backend: string;
    /**
     * The CA chain as a list of format specific certificates.
     */
    readonly caChains: string[];
    /**
     * Certificate associated with this issuer.
     */
    readonly certificate: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * ID of the issuer.
     */
    readonly issuerId: string;
    /**
     * Name of the issuer.
     */
    readonly issuerName: string;
    readonly issuerRef: string;
    /**
     * ID of the key used by the issuer.
     */
    readonly keyId: string;
    /**
     * Behavior of a leaf's NotAfter field during issuance.
     */
    readonly leafNotAfterBehavior: string;
    /**
     * Chain of issuer references to build this issuer's computed 
     * CAChain field from, when non-empty.
     */
    readonly manualChains: string[];
    readonly namespace?: string;
    /**
     * Allowed usages for this issuer.
     */
    readonly usage: string;
}
export function getBackendIssuerOutput(args: GetBackendIssuerOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetBackendIssuerResult> {
    return pulumi.output(args).apply((a: any) => getBackendIssuer(a, opts))
}

/**
 * A collection of arguments for invoking getBackendIssuer.
 */
export interface GetBackendIssuerOutputArgs {
    /**
     * The path to the PKI secret backend to
     * read the issuer from, with no leading or trailing `/`s.
     */
    backend: pulumi.Input<string>;
    /**
     * Reference to an existing issuer.
     */
    issuerRef: pulumi.Input<string>;
    /**
     * The namespace of the target resource.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
     * *Available only for Vault Enterprise*.
     */
    namespace?: pulumi.Input<string>;
}
