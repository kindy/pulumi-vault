// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./authBackend";
export * from "./authBackendGroup";
export * from "./authBackendUser";

// Import resources to register:
import { AuthBackend } from "./authBackend";
import { AuthBackendGroup } from "./authBackendGroup";
import { AuthBackendUser } from "./authBackendUser";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "vault:ldap/authBackend:AuthBackend":
                return new AuthBackend(name, <any>undefined, { urn })
            case "vault:ldap/authBackendGroup:AuthBackendGroup":
                return new AuthBackendGroup(name, <any>undefined, { urn })
            case "vault:ldap/authBackendUser:AuthBackendUser":
                return new AuthBackendUser(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("vault", "ldap/authBackend", _module)
pulumi.runtime.registerResourceModule("vault", "ldap/authBackendGroup", _module)
pulumi.runtime.registerResourceModule("vault", "ldap/authBackendUser", _module)
