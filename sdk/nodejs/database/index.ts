// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./secretBackendConnection";
export * from "./secretBackendRole";
export * from "./secretBackendStaticRole";

// Import resources to register:
import { SecretBackendConnection } from "./secretBackendConnection";
import { SecretBackendRole } from "./secretBackendRole";
import { SecretBackendStaticRole } from "./secretBackendStaticRole";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "vault:database/secretBackendConnection:SecretBackendConnection":
                return new SecretBackendConnection(name, <any>undefined, { urn })
            case "vault:database/secretBackendRole:SecretBackendRole":
                return new SecretBackendRole(name, <any>undefined, { urn })
            case "vault:database/secretBackendStaticRole:SecretBackendStaticRole":
                return new SecretBackendStaticRole(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("vault", "database/secretBackendConnection", _module)
pulumi.runtime.registerResourceModule("vault", "database/secretBackendRole", _module)
pulumi.runtime.registerResourceModule("vault", "database/secretBackendStaticRole", _module)
