// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./authBackendLogin";
export * from "./authBackendRole";
export * from "./authBackendRoleSecretID";
export * from "./getAuthBackendRoleId";

// Import resources to register:
import { AuthBackendLogin } from "./authBackendLogin";
import { AuthBackendRole } from "./authBackendRole";
import { AuthBackendRoleSecretID } from "./authBackendRoleSecretID";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "vault:appRole/authBackendLogin:AuthBackendLogin":
                return new AuthBackendLogin(name, <any>undefined, { urn })
            case "vault:appRole/authBackendRole:AuthBackendRole":
                return new AuthBackendRole(name, <any>undefined, { urn })
            case "vault:appRole/authBackendRoleSecretID:AuthBackendRoleSecretID":
                return new AuthBackendRoleSecretID(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("vault", "appRole/authBackendLogin", _module)
pulumi.runtime.registerResourceModule("vault", "appRole/authBackendRole", _module)
pulumi.runtime.registerResourceModule("vault", "appRole/authBackendRoleSecretID", _module)
