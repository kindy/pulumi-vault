// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

export class SecretLibrary extends pulumi.CustomResource {
    /**
     * Get an existing SecretLibrary resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretLibraryState, opts?: pulumi.CustomResourceOptions): SecretLibrary {
        return new SecretLibrary(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:ad/secretLibrary:SecretLibrary';

    /**
     * Returns true if the given object is an instance of SecretLibrary.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretLibrary {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretLibrary.__pulumiType;
    }

    /**
     * The mount path for the AD backend.
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * Disable enforcing that service accounts must be checked in by the entity or client token that checked them out.
     */
    public readonly disableCheckInEnforcement!: pulumi.Output<boolean | undefined>;
    /**
     * The maximum amount of time, in seconds, a check-out last with renewal before Vault automatically checks it back in.
     */
    public readonly maxTtl!: pulumi.Output<number>;
    /**
     * The name of the set of service accounts.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The names of all the service accounts that can be checked out from this set. These service accounts must already exist
     * in Active Directory.
     */
    public readonly serviceAccountNames!: pulumi.Output<string[]>;
    /**
     * The amount of time, in seconds, a single check-out lasts before Vault automatically checks it back in.
     */
    public readonly ttl!: pulumi.Output<number>;

    /**
     * Create a SecretLibrary resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretLibraryArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretLibraryArgs | SecretLibraryState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretLibraryState | undefined;
            resourceInputs["backend"] = state ? state.backend : undefined;
            resourceInputs["disableCheckInEnforcement"] = state ? state.disableCheckInEnforcement : undefined;
            resourceInputs["maxTtl"] = state ? state.maxTtl : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["serviceAccountNames"] = state ? state.serviceAccountNames : undefined;
            resourceInputs["ttl"] = state ? state.ttl : undefined;
        } else {
            const args = argsOrState as SecretLibraryArgs | undefined;
            if ((!args || args.backend === undefined) && !opts.urn) {
                throw new Error("Missing required property 'backend'");
            }
            if ((!args || args.serviceAccountNames === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceAccountNames'");
            }
            resourceInputs["backend"] = args ? args.backend : undefined;
            resourceInputs["disableCheckInEnforcement"] = args ? args.disableCheckInEnforcement : undefined;
            resourceInputs["maxTtl"] = args ? args.maxTtl : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["serviceAccountNames"] = args ? args.serviceAccountNames : undefined;
            resourceInputs["ttl"] = args ? args.ttl : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(SecretLibrary.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretLibrary resources.
 */
export interface SecretLibraryState {
    /**
     * The mount path for the AD backend.
     */
    backend?: pulumi.Input<string>;
    /**
     * Disable enforcing that service accounts must be checked in by the entity or client token that checked them out.
     */
    disableCheckInEnforcement?: pulumi.Input<boolean>;
    /**
     * The maximum amount of time, in seconds, a check-out last with renewal before Vault automatically checks it back in.
     */
    maxTtl?: pulumi.Input<number>;
    /**
     * The name of the set of service accounts.
     */
    name?: pulumi.Input<string>;
    /**
     * The names of all the service accounts that can be checked out from this set. These service accounts must already exist
     * in Active Directory.
     */
    serviceAccountNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The amount of time, in seconds, a single check-out lasts before Vault automatically checks it back in.
     */
    ttl?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a SecretLibrary resource.
 */
export interface SecretLibraryArgs {
    /**
     * The mount path for the AD backend.
     */
    backend: pulumi.Input<string>;
    /**
     * Disable enforcing that service accounts must be checked in by the entity or client token that checked them out.
     */
    disableCheckInEnforcement?: pulumi.Input<boolean>;
    /**
     * The maximum amount of time, in seconds, a check-out last with renewal before Vault automatically checks it back in.
     */
    maxTtl?: pulumi.Input<number>;
    /**
     * The name of the set of service accounts.
     */
    name?: pulumi.Input<string>;
    /**
     * The names of all the service accounts that can be checked out from this set. These service accounts must already exist
     * in Active Directory.
     */
    serviceAccountNames: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The amount of time, in seconds, a single check-out lasts before Vault automatically checks it back in.
     */
    ttl?: pulumi.Input<number>;
}
