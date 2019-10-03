// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

/**
 * > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/database_secret_backend_role.html.markdown.
 */
export class SecretBackendRole extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackendRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendRoleState, opts?: pulumi.CustomResourceOptions): SecretBackendRole {
        return new SecretBackendRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:database/secretBackendRole:SecretBackendRole';

    /**
     * Returns true if the given object is an instance of SecretBackendRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackendRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackendRole.__pulumiType;
    }

    /**
     * The unique name of the Vault mount to configure.
     */
    public readonly backend!: pulumi.Output<string>;
    /**
     * The database statements to execute when
     * creating a user.
     */
    public readonly creationStatements!: pulumi.Output<string[]>;
    /**
     * The unique name of the database connection to use for
     * the role.
     */
    public readonly dbName!: pulumi.Output<string>;
    /**
     * The default number of seconds for leases for this
     * role.
     */
    public readonly defaultTtl!: pulumi.Output<number | undefined>;
    /**
     * The maximum number of seconds for leases for this
     * role.
     */
    public readonly maxTtl!: pulumi.Output<number | undefined>;
    /**
     * A unique name to give the role.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The database statements to execute when
     * renewing a user.
     */
    public readonly renewStatements!: pulumi.Output<string[] | undefined>;
    /**
     * The database statements to execute when
     * revoking a user.
     */
    public readonly revocationStatements!: pulumi.Output<string[] | undefined>;
    /**
     * The database statements to execute when
     * rolling back creation due to an error.
     */
    public readonly rollbackStatements!: pulumi.Output<string[] | undefined>;

    /**
     * Create a SecretBackendRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretBackendRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendRoleArgs | SecretBackendRoleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as SecretBackendRoleState | undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["creationStatements"] = state ? state.creationStatements : undefined;
            inputs["dbName"] = state ? state.dbName : undefined;
            inputs["defaultTtl"] = state ? state.defaultTtl : undefined;
            inputs["maxTtl"] = state ? state.maxTtl : undefined;
            inputs["name"] = state ? state.name : undefined;
            inputs["renewStatements"] = state ? state.renewStatements : undefined;
            inputs["revocationStatements"] = state ? state.revocationStatements : undefined;
            inputs["rollbackStatements"] = state ? state.rollbackStatements : undefined;
        } else {
            const args = argsOrState as SecretBackendRoleArgs | undefined;
            if (!args || args.backend === undefined) {
                throw new Error("Missing required property 'backend'");
            }
            if (!args || args.creationStatements === undefined) {
                throw new Error("Missing required property 'creationStatements'");
            }
            if (!args || args.dbName === undefined) {
                throw new Error("Missing required property 'dbName'");
            }
            inputs["backend"] = args ? args.backend : undefined;
            inputs["creationStatements"] = args ? args.creationStatements : undefined;
            inputs["dbName"] = args ? args.dbName : undefined;
            inputs["defaultTtl"] = args ? args.defaultTtl : undefined;
            inputs["maxTtl"] = args ? args.maxTtl : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["renewStatements"] = args ? args.renewStatements : undefined;
            inputs["revocationStatements"] = args ? args.revocationStatements : undefined;
            inputs["rollbackStatements"] = args ? args.rollbackStatements : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(SecretBackendRole.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackendRole resources.
 */
export interface SecretBackendRoleState {
    /**
     * The unique name of the Vault mount to configure.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * The database statements to execute when
     * creating a user.
     */
    readonly creationStatements?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The unique name of the database connection to use for
     * the role.
     */
    readonly dbName?: pulumi.Input<string>;
    /**
     * The default number of seconds for leases for this
     * role.
     */
    readonly defaultTtl?: pulumi.Input<number>;
    /**
     * The maximum number of seconds for leases for this
     * role.
     */
    readonly maxTtl?: pulumi.Input<number>;
    /**
     * A unique name to give the role.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The database statements to execute when
     * renewing a user.
     */
    readonly renewStatements?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The database statements to execute when
     * revoking a user.
     */
    readonly revocationStatements?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The database statements to execute when
     * rolling back creation due to an error.
     */
    readonly rollbackStatements?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a SecretBackendRole resource.
 */
export interface SecretBackendRoleArgs {
    /**
     * The unique name of the Vault mount to configure.
     */
    readonly backend: pulumi.Input<string>;
    /**
     * The database statements to execute when
     * creating a user.
     */
    readonly creationStatements: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The unique name of the database connection to use for
     * the role.
     */
    readonly dbName: pulumi.Input<string>;
    /**
     * The default number of seconds for leases for this
     * role.
     */
    readonly defaultTtl?: pulumi.Input<number>;
    /**
     * The maximum number of seconds for leases for this
     * role.
     */
    readonly maxTtl?: pulumi.Input<number>;
    /**
     * A unique name to give the role.
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The database statements to execute when
     * renewing a user.
     */
    readonly renewStatements?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The database statements to execute when
     * revoking a user.
     */
    readonly revocationStatements?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The database statements to execute when
     * rolling back creation due to an error.
     */
    readonly rollbackStatements?: pulumi.Input<pulumi.Input<string>[]>;
}
