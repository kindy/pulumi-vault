// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/d/identity_entity.html.markdown.
 */
export function getEntity(args?: GetEntityArgs, opts?: pulumi.InvokeOptions): Promise<GetEntityResult> & GetEntityResult {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    const promise: Promise<GetEntityResult> = pulumi.runtime.invoke("vault:identity/getEntity:getEntity", {
        "aliasId": args.aliasId,
        "aliasMountAccessor": args.aliasMountAccessor,
        "aliasName": args.aliasName,
        "entityId": args.entityId,
        "entityName": args.entityName,
    }, opts);

    return pulumi.utils.liftProperties(promise, opts);
}

/**
 * A collection of arguments for invoking getEntity.
 */
export interface GetEntityArgs {
    /**
     * ID of the alias.
     */
    readonly aliasId?: string;
    /**
     * Accessor of the mount to which the alias belongs to.
     * This should be supplied in conjunction with `aliasName`.
     */
    readonly aliasMountAccessor?: string;
    /**
     * Name of the alias. This should be supplied in conjunction with
     * `aliasMountAccessor`.
     */
    readonly aliasName?: string;
    /**
     * ID of the entity.
     */
    readonly entityId?: string;
    /**
     * Name of the entity.
     */
    readonly entityName?: string;
}

/**
 * A collection of values returned by getEntity.
 */
export interface GetEntityResult {
    readonly aliasId: string;
    readonly aliasMountAccessor: string;
    readonly aliasName: string;
    /**
     * A list of entity alias. Structure is documented below.
     */
    readonly aliases: outputs.identity.GetEntityAlias[];
    /**
     * Creation time of the Alias
     */
    readonly creationTime: string;
    /**
     * A string containing the full data payload retrieved from
     * Vault, serialized in JSON format.
     */
    readonly dataJson: string;
    /**
     * List of Group IDs of which the entity is directly a member of
     */
    readonly directGroupIds: string[];
    /**
     * Whether the entity is disabled
     */
    readonly disabled: boolean;
    readonly entityId: string;
    readonly entityName: string;
    /**
     * List of all Group IDs of which the entity is a member of
     */
    readonly groupIds: string[];
    /**
     * List of all Group IDs of which the entity is a member of transitively
     */
    readonly inheritedGroupIds: string[];
    /**
     * Last update time of the alias
     */
    readonly lastUpdateTime: string;
    /**
     * Other entity IDs which is merged with this entity
     */
    readonly mergedEntityIds: string[];
    /**
     * Arbitrary metadata
     */
    readonly metadata: {[key: string]: any};
    /**
     * Namespace of which the entity is part of
     */
    readonly namespaceId: string;
    /**
     * List of policies attached to the entity
     */
    readonly policies: string[];
    /**
     * id is the provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
}
