// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * AD secret backend can be imported using the `backend`, e.g.
 *
 * ```sh
 *  $ pulumi import vault:ad/secretBackend:SecretBackend ad ad
 * ```
 */
export class SecretBackend extends pulumi.CustomResource {
    /**
     * Get an existing SecretBackend resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SecretBackendState, opts?: pulumi.CustomResourceOptions): SecretBackend {
        return new SecretBackend(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:ad/secretBackend:SecretBackend';

    /**
     * Returns true if the given object is an instance of SecretBackend.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SecretBackend {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SecretBackend.__pulumiType;
    }

    /**
     * Use anonymous binds when performing LDAP group searches
     * (if true the initial credentials will still be used for the initial connection test).
     */
    public readonly anonymousGroupSearch!: pulumi.Output<boolean | undefined>;
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `ad`.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * Distinguished name of object to bind when performing user and group search.
     */
    public readonly binddn!: pulumi.Output<string>;
    /**
     * Password to use along with binddn when performing user search.
     */
    public readonly bindpass!: pulumi.Output<string>;
    /**
     * If set, user and group names assigned to policies within the
     * backend will be case sensitive. Otherwise, names will be normalized to lower case.
     */
    public readonly caseSensitiveNames!: pulumi.Output<boolean | undefined>;
    /**
     * CA certificate to use when verifying LDAP server certificate, must be
     * x509 PEM encoded.
     */
    public readonly certificate!: pulumi.Output<string | undefined>;
    /**
     * Client certificate to provide to the LDAP server, must be x509 PEM encoded.
     */
    public readonly clientTlsCert!: pulumi.Output<string | undefined>;
    /**
     * Client certificate key to provide to the LDAP server, must be x509 PEM encoded.
     */
    public readonly clientTlsKey!: pulumi.Output<string | undefined>;
    /**
     * Default lease duration for secrets in seconds.
     */
    public readonly defaultLeaseTtlSeconds!: pulumi.Output<number>;
    /**
     * Denies an unauthenticated LDAP bind request if the user's password is empty;
     * defaults to true.
     */
    public readonly denyNullBind!: pulumi.Output<boolean | undefined>;
    /**
     * Human-friendly description of the mount for the Active Directory backend.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Use anonymous bind to discover the bind Distinguished Name of a user.
     */
    public readonly discoverdn!: pulumi.Output<boolean | undefined>;
    /**
     * Text to insert the password into, ex. "customPrefix{{PASSWORD}}customSuffix". This
     * setting is deprecated and should instead use `passwordPolicy`.
     *
     * @deprecated Formatter is deprecated and password_policy should be used with Vault >= 1.5.
     */
    public readonly formatter!: pulumi.Output<string>;
    /**
     * LDAP attribute to follow on objects returned by <groupfilter> in order to enumerate
     * user group membership. Examples: `cn` or `memberOf`, etc. Defaults to `cn`.
     */
    public readonly groupattr!: pulumi.Output<string | undefined>;
    /**
     * LDAP search base to use for group membership search (eg: ou=Groups,dc=example,dc=org).
     */
    public readonly groupdn!: pulumi.Output<string | undefined>;
    /**
     * Go template for querying group membership of user (optional) The template can access
     * the following context variables: UserDN, Username. Defaults to `(|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))`
     */
    public readonly groupfilter!: pulumi.Output<string | undefined>;
    /**
     * Skip LDAP server SSL Certificate verification. This is not recommended for production.
     * Defaults to `false`.
     */
    public readonly insecureTls!: pulumi.Output<boolean | undefined>;
    /**
     * The number of seconds after a Vault rotation where, if Active Directory
     * shows a later rotation, it should be considered out-of-band
     */
    public readonly lastRotationTolerance!: pulumi.Output<number>;
    /**
     * The desired length of passwords that Vault generates. This
     * setting is deprecated and should instead use `passwordPolicy`.
     *
     * @deprecated Length is deprecated and password_policy should be used with Vault >= 1.5.
     */
    public readonly length!: pulumi.Output<number>;
    /**
     * Mark the secrets engine as local-only. Local engines are not replicated or removed by
     * replication.Tolerance duration to use when checking the last rotation time.
     */
    public readonly local!: pulumi.Output<boolean | undefined>;
    /**
     * Maximum possible lease duration for secrets in seconds.
     */
    public readonly maxLeaseTtlSeconds!: pulumi.Output<number>;
    /**
     * In seconds, the maximum password time-to-live.
     */
    public readonly maxTtl!: pulumi.Output<number>;
    /**
     * Name of the password policy to use to generate passwords.
     */
    public readonly passwordPolicy!: pulumi.Output<string | undefined>;
    /**
     * Timeout, in seconds, for the connection when making requests against the server
     * before returning back an error.
     */
    public readonly requestTimeout!: pulumi.Output<number | undefined>;
    /**
     * Issue a StartTLS command after establishing unencrypted connection.
     */
    public readonly starttls!: pulumi.Output<boolean>;
    /**
     * Maximum TLS version to use. Accepted values are `tls10`, `tls11`,
     * `tls12` or `tls13`. Defaults to `tls12`.
     */
    public readonly tlsMaxVersion!: pulumi.Output<string>;
    /**
     * Minimum TLS version to use. Accepted values are `tls10`, `tls11`,
     * `tls12` or `tls13`. Defaults to `tls12`.
     */
    public readonly tlsMinVersion!: pulumi.Output<string>;
    /**
     * In seconds, the default password time-to-live.
     */
    public readonly ttl!: pulumi.Output<number>;
    /**
     * Enables userPrincipalDomain login with [username]@UPNDomain.
     */
    public readonly upndomain!: pulumi.Output<string>;
    /**
     * LDAP URL to connect to. Multiple URLs can be specified by concatenating
     * them with commas; they will be tried in-order. Defaults to `ldap://127.0.0.1`.
     */
    public readonly url!: pulumi.Output<string | undefined>;
    /**
     * In Vault 1.1.1 a fix for handling group CN values of
     * different cases unfortunately introduced a regression that could cause previously defined groups
     * to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for
     * matching group CNs will be used. This is only needed in some upgrade scenarios for backwards
     * compatibility. It is enabled by default if the config is upgraded but disabled by default on
     * new configurations.
     */
    public readonly usePre111GroupCnBehavior!: pulumi.Output<boolean>;
    /**
     * If true, use the Active Directory tokenGroups constructed attribute of the
     * user to find the group memberships. This will find all security groups including nested ones.
     */
    public readonly useTokenGroups!: pulumi.Output<boolean | undefined>;
    /**
     * Attribute used when searching users. Defaults to `cn`.
     */
    public readonly userattr!: pulumi.Output<string | undefined>;
    /**
     * LDAP domain to use for users (eg: ou=People,dc=example,dc=org)`.
     */
    public readonly userdn!: pulumi.Output<string | undefined>;

    /**
     * Create a SecretBackend resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SecretBackendArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SecretBackendArgs | SecretBackendState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SecretBackendState | undefined;
            inputs["anonymousGroupSearch"] = state ? state.anonymousGroupSearch : undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["binddn"] = state ? state.binddn : undefined;
            inputs["bindpass"] = state ? state.bindpass : undefined;
            inputs["caseSensitiveNames"] = state ? state.caseSensitiveNames : undefined;
            inputs["certificate"] = state ? state.certificate : undefined;
            inputs["clientTlsCert"] = state ? state.clientTlsCert : undefined;
            inputs["clientTlsKey"] = state ? state.clientTlsKey : undefined;
            inputs["defaultLeaseTtlSeconds"] = state ? state.defaultLeaseTtlSeconds : undefined;
            inputs["denyNullBind"] = state ? state.denyNullBind : undefined;
            inputs["description"] = state ? state.description : undefined;
            inputs["discoverdn"] = state ? state.discoverdn : undefined;
            inputs["formatter"] = state ? state.formatter : undefined;
            inputs["groupattr"] = state ? state.groupattr : undefined;
            inputs["groupdn"] = state ? state.groupdn : undefined;
            inputs["groupfilter"] = state ? state.groupfilter : undefined;
            inputs["insecureTls"] = state ? state.insecureTls : undefined;
            inputs["lastRotationTolerance"] = state ? state.lastRotationTolerance : undefined;
            inputs["length"] = state ? state.length : undefined;
            inputs["local"] = state ? state.local : undefined;
            inputs["maxLeaseTtlSeconds"] = state ? state.maxLeaseTtlSeconds : undefined;
            inputs["maxTtl"] = state ? state.maxTtl : undefined;
            inputs["passwordPolicy"] = state ? state.passwordPolicy : undefined;
            inputs["requestTimeout"] = state ? state.requestTimeout : undefined;
            inputs["starttls"] = state ? state.starttls : undefined;
            inputs["tlsMaxVersion"] = state ? state.tlsMaxVersion : undefined;
            inputs["tlsMinVersion"] = state ? state.tlsMinVersion : undefined;
            inputs["ttl"] = state ? state.ttl : undefined;
            inputs["upndomain"] = state ? state.upndomain : undefined;
            inputs["url"] = state ? state.url : undefined;
            inputs["usePre111GroupCnBehavior"] = state ? state.usePre111GroupCnBehavior : undefined;
            inputs["useTokenGroups"] = state ? state.useTokenGroups : undefined;
            inputs["userattr"] = state ? state.userattr : undefined;
            inputs["userdn"] = state ? state.userdn : undefined;
        } else {
            const args = argsOrState as SecretBackendArgs | undefined;
            if ((!args || args.binddn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'binddn'");
            }
            if ((!args || args.bindpass === undefined) && !opts.urn) {
                throw new Error("Missing required property 'bindpass'");
            }
            inputs["anonymousGroupSearch"] = args ? args.anonymousGroupSearch : undefined;
            inputs["backend"] = args ? args.backend : undefined;
            inputs["binddn"] = args ? args.binddn : undefined;
            inputs["bindpass"] = args ? args.bindpass : undefined;
            inputs["caseSensitiveNames"] = args ? args.caseSensitiveNames : undefined;
            inputs["certificate"] = args ? args.certificate : undefined;
            inputs["clientTlsCert"] = args ? args.clientTlsCert : undefined;
            inputs["clientTlsKey"] = args ? args.clientTlsKey : undefined;
            inputs["defaultLeaseTtlSeconds"] = args ? args.defaultLeaseTtlSeconds : undefined;
            inputs["denyNullBind"] = args ? args.denyNullBind : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["discoverdn"] = args ? args.discoverdn : undefined;
            inputs["formatter"] = args ? args.formatter : undefined;
            inputs["groupattr"] = args ? args.groupattr : undefined;
            inputs["groupdn"] = args ? args.groupdn : undefined;
            inputs["groupfilter"] = args ? args.groupfilter : undefined;
            inputs["insecureTls"] = args ? args.insecureTls : undefined;
            inputs["lastRotationTolerance"] = args ? args.lastRotationTolerance : undefined;
            inputs["length"] = args ? args.length : undefined;
            inputs["local"] = args ? args.local : undefined;
            inputs["maxLeaseTtlSeconds"] = args ? args.maxLeaseTtlSeconds : undefined;
            inputs["maxTtl"] = args ? args.maxTtl : undefined;
            inputs["passwordPolicy"] = args ? args.passwordPolicy : undefined;
            inputs["requestTimeout"] = args ? args.requestTimeout : undefined;
            inputs["starttls"] = args ? args.starttls : undefined;
            inputs["tlsMaxVersion"] = args ? args.tlsMaxVersion : undefined;
            inputs["tlsMinVersion"] = args ? args.tlsMinVersion : undefined;
            inputs["ttl"] = args ? args.ttl : undefined;
            inputs["upndomain"] = args ? args.upndomain : undefined;
            inputs["url"] = args ? args.url : undefined;
            inputs["usePre111GroupCnBehavior"] = args ? args.usePre111GroupCnBehavior : undefined;
            inputs["useTokenGroups"] = args ? args.useTokenGroups : undefined;
            inputs["userattr"] = args ? args.userattr : undefined;
            inputs["userdn"] = args ? args.userdn : undefined;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(SecretBackend.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SecretBackend resources.
 */
export interface SecretBackendState {
    /**
     * Use anonymous binds when performing LDAP group searches
     * (if true the initial credentials will still be used for the initial connection test).
     */
    readonly anonymousGroupSearch?: pulumi.Input<boolean>;
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `ad`.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * Distinguished name of object to bind when performing user and group search.
     */
    readonly binddn?: pulumi.Input<string>;
    /**
     * Password to use along with binddn when performing user search.
     */
    readonly bindpass?: pulumi.Input<string>;
    /**
     * If set, user and group names assigned to policies within the
     * backend will be case sensitive. Otherwise, names will be normalized to lower case.
     */
    readonly caseSensitiveNames?: pulumi.Input<boolean>;
    /**
     * CA certificate to use when verifying LDAP server certificate, must be
     * x509 PEM encoded.
     */
    readonly certificate?: pulumi.Input<string>;
    /**
     * Client certificate to provide to the LDAP server, must be x509 PEM encoded.
     */
    readonly clientTlsCert?: pulumi.Input<string>;
    /**
     * Client certificate key to provide to the LDAP server, must be x509 PEM encoded.
     */
    readonly clientTlsKey?: pulumi.Input<string>;
    /**
     * Default lease duration for secrets in seconds.
     */
    readonly defaultLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * Denies an unauthenticated LDAP bind request if the user's password is empty;
     * defaults to true.
     */
    readonly denyNullBind?: pulumi.Input<boolean>;
    /**
     * Human-friendly description of the mount for the Active Directory backend.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Use anonymous bind to discover the bind Distinguished Name of a user.
     */
    readonly discoverdn?: pulumi.Input<boolean>;
    /**
     * Text to insert the password into, ex. "customPrefix{{PASSWORD}}customSuffix". This
     * setting is deprecated and should instead use `passwordPolicy`.
     *
     * @deprecated Formatter is deprecated and password_policy should be used with Vault >= 1.5.
     */
    readonly formatter?: pulumi.Input<string>;
    /**
     * LDAP attribute to follow on objects returned by <groupfilter> in order to enumerate
     * user group membership. Examples: `cn` or `memberOf`, etc. Defaults to `cn`.
     */
    readonly groupattr?: pulumi.Input<string>;
    /**
     * LDAP search base to use for group membership search (eg: ou=Groups,dc=example,dc=org).
     */
    readonly groupdn?: pulumi.Input<string>;
    /**
     * Go template for querying group membership of user (optional) The template can access
     * the following context variables: UserDN, Username. Defaults to `(|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))`
     */
    readonly groupfilter?: pulumi.Input<string>;
    /**
     * Skip LDAP server SSL Certificate verification. This is not recommended for production.
     * Defaults to `false`.
     */
    readonly insecureTls?: pulumi.Input<boolean>;
    /**
     * The number of seconds after a Vault rotation where, if Active Directory
     * shows a later rotation, it should be considered out-of-band
     */
    readonly lastRotationTolerance?: pulumi.Input<number>;
    /**
     * The desired length of passwords that Vault generates. This
     * setting is deprecated and should instead use `passwordPolicy`.
     *
     * @deprecated Length is deprecated and password_policy should be used with Vault >= 1.5.
     */
    readonly length?: pulumi.Input<number>;
    /**
     * Mark the secrets engine as local-only. Local engines are not replicated or removed by
     * replication.Tolerance duration to use when checking the last rotation time.
     */
    readonly local?: pulumi.Input<boolean>;
    /**
     * Maximum possible lease duration for secrets in seconds.
     */
    readonly maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * In seconds, the maximum password time-to-live.
     */
    readonly maxTtl?: pulumi.Input<number>;
    /**
     * Name of the password policy to use to generate passwords.
     */
    readonly passwordPolicy?: pulumi.Input<string>;
    /**
     * Timeout, in seconds, for the connection when making requests against the server
     * before returning back an error.
     */
    readonly requestTimeout?: pulumi.Input<number>;
    /**
     * Issue a StartTLS command after establishing unencrypted connection.
     */
    readonly starttls?: pulumi.Input<boolean>;
    /**
     * Maximum TLS version to use. Accepted values are `tls10`, `tls11`,
     * `tls12` or `tls13`. Defaults to `tls12`.
     */
    readonly tlsMaxVersion?: pulumi.Input<string>;
    /**
     * Minimum TLS version to use. Accepted values are `tls10`, `tls11`,
     * `tls12` or `tls13`. Defaults to `tls12`.
     */
    readonly tlsMinVersion?: pulumi.Input<string>;
    /**
     * In seconds, the default password time-to-live.
     */
    readonly ttl?: pulumi.Input<number>;
    /**
     * Enables userPrincipalDomain login with [username]@UPNDomain.
     */
    readonly upndomain?: pulumi.Input<string>;
    /**
     * LDAP URL to connect to. Multiple URLs can be specified by concatenating
     * them with commas; they will be tried in-order. Defaults to `ldap://127.0.0.1`.
     */
    readonly url?: pulumi.Input<string>;
    /**
     * In Vault 1.1.1 a fix for handling group CN values of
     * different cases unfortunately introduced a regression that could cause previously defined groups
     * to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for
     * matching group CNs will be used. This is only needed in some upgrade scenarios for backwards
     * compatibility. It is enabled by default if the config is upgraded but disabled by default on
     * new configurations.
     */
    readonly usePre111GroupCnBehavior?: pulumi.Input<boolean>;
    /**
     * If true, use the Active Directory tokenGroups constructed attribute of the
     * user to find the group memberships. This will find all security groups including nested ones.
     */
    readonly useTokenGroups?: pulumi.Input<boolean>;
    /**
     * Attribute used when searching users. Defaults to `cn`.
     */
    readonly userattr?: pulumi.Input<string>;
    /**
     * LDAP domain to use for users (eg: ou=People,dc=example,dc=org)`.
     */
    readonly userdn?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a SecretBackend resource.
 */
export interface SecretBackendArgs {
    /**
     * Use anonymous binds when performing LDAP group searches
     * (if true the initial credentials will still be used for the initial connection test).
     */
    readonly anonymousGroupSearch?: pulumi.Input<boolean>;
    /**
     * The unique path this backend should be mounted at. Must
     * not begin or end with a `/`. Defaults to `ad`.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * Distinguished name of object to bind when performing user and group search.
     */
    readonly binddn: pulumi.Input<string>;
    /**
     * Password to use along with binddn when performing user search.
     */
    readonly bindpass: pulumi.Input<string>;
    /**
     * If set, user and group names assigned to policies within the
     * backend will be case sensitive. Otherwise, names will be normalized to lower case.
     */
    readonly caseSensitiveNames?: pulumi.Input<boolean>;
    /**
     * CA certificate to use when verifying LDAP server certificate, must be
     * x509 PEM encoded.
     */
    readonly certificate?: pulumi.Input<string>;
    /**
     * Client certificate to provide to the LDAP server, must be x509 PEM encoded.
     */
    readonly clientTlsCert?: pulumi.Input<string>;
    /**
     * Client certificate key to provide to the LDAP server, must be x509 PEM encoded.
     */
    readonly clientTlsKey?: pulumi.Input<string>;
    /**
     * Default lease duration for secrets in seconds.
     */
    readonly defaultLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * Denies an unauthenticated LDAP bind request if the user's password is empty;
     * defaults to true.
     */
    readonly denyNullBind?: pulumi.Input<boolean>;
    /**
     * Human-friendly description of the mount for the Active Directory backend.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * Use anonymous bind to discover the bind Distinguished Name of a user.
     */
    readonly discoverdn?: pulumi.Input<boolean>;
    /**
     * Text to insert the password into, ex. "customPrefix{{PASSWORD}}customSuffix". This
     * setting is deprecated and should instead use `passwordPolicy`.
     *
     * @deprecated Formatter is deprecated and password_policy should be used with Vault >= 1.5.
     */
    readonly formatter?: pulumi.Input<string>;
    /**
     * LDAP attribute to follow on objects returned by <groupfilter> in order to enumerate
     * user group membership. Examples: `cn` or `memberOf`, etc. Defaults to `cn`.
     */
    readonly groupattr?: pulumi.Input<string>;
    /**
     * LDAP search base to use for group membership search (eg: ou=Groups,dc=example,dc=org).
     */
    readonly groupdn?: pulumi.Input<string>;
    /**
     * Go template for querying group membership of user (optional) The template can access
     * the following context variables: UserDN, Username. Defaults to `(|(memberUid={{.Username}})(member={{.UserDN}})(uniqueMember={{.UserDN}}))`
     */
    readonly groupfilter?: pulumi.Input<string>;
    /**
     * Skip LDAP server SSL Certificate verification. This is not recommended for production.
     * Defaults to `false`.
     */
    readonly insecureTls?: pulumi.Input<boolean>;
    /**
     * The number of seconds after a Vault rotation where, if Active Directory
     * shows a later rotation, it should be considered out-of-band
     */
    readonly lastRotationTolerance?: pulumi.Input<number>;
    /**
     * The desired length of passwords that Vault generates. This
     * setting is deprecated and should instead use `passwordPolicy`.
     *
     * @deprecated Length is deprecated and password_policy should be used with Vault >= 1.5.
     */
    readonly length?: pulumi.Input<number>;
    /**
     * Mark the secrets engine as local-only. Local engines are not replicated or removed by
     * replication.Tolerance duration to use when checking the last rotation time.
     */
    readonly local?: pulumi.Input<boolean>;
    /**
     * Maximum possible lease duration for secrets in seconds.
     */
    readonly maxLeaseTtlSeconds?: pulumi.Input<number>;
    /**
     * In seconds, the maximum password time-to-live.
     */
    readonly maxTtl?: pulumi.Input<number>;
    /**
     * Name of the password policy to use to generate passwords.
     */
    readonly passwordPolicy?: pulumi.Input<string>;
    /**
     * Timeout, in seconds, for the connection when making requests against the server
     * before returning back an error.
     */
    readonly requestTimeout?: pulumi.Input<number>;
    /**
     * Issue a StartTLS command after establishing unencrypted connection.
     */
    readonly starttls?: pulumi.Input<boolean>;
    /**
     * Maximum TLS version to use. Accepted values are `tls10`, `tls11`,
     * `tls12` or `tls13`. Defaults to `tls12`.
     */
    readonly tlsMaxVersion?: pulumi.Input<string>;
    /**
     * Minimum TLS version to use. Accepted values are `tls10`, `tls11`,
     * `tls12` or `tls13`. Defaults to `tls12`.
     */
    readonly tlsMinVersion?: pulumi.Input<string>;
    /**
     * In seconds, the default password time-to-live.
     */
    readonly ttl?: pulumi.Input<number>;
    /**
     * Enables userPrincipalDomain login with [username]@UPNDomain.
     */
    readonly upndomain?: pulumi.Input<string>;
    /**
     * LDAP URL to connect to. Multiple URLs can be specified by concatenating
     * them with commas; they will be tried in-order. Defaults to `ldap://127.0.0.1`.
     */
    readonly url?: pulumi.Input<string>;
    /**
     * In Vault 1.1.1 a fix for handling group CN values of
     * different cases unfortunately introduced a regression that could cause previously defined groups
     * to not be found due to a change in the resulting name. If set true, the pre-1.1.1 behavior for
     * matching group CNs will be used. This is only needed in some upgrade scenarios for backwards
     * compatibility. It is enabled by default if the config is upgraded but disabled by default on
     * new configurations.
     */
    readonly usePre111GroupCnBehavior?: pulumi.Input<boolean>;
    /**
     * If true, use the Active Directory tokenGroups constructed attribute of the
     * user to find the group memberships. This will find all security groups including nested ones.
     */
    readonly useTokenGroups?: pulumi.Input<boolean>;
    /**
     * Attribute used when searching users. Defaults to `cn`.
     */
    readonly userattr?: pulumi.Input<string>;
    /**
     * LDAP domain to use for users (eg: ou=People,dc=example,dc=org)`.
     */
    readonly userdn?: pulumi.Input<string>;
}
