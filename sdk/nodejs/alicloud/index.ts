// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AuthBackendRoleArgs, AuthBackendRoleState } from "./authBackendRole";
export type AuthBackendRole = import("./authBackendRole").AuthBackendRole;
export const AuthBackendRole: typeof import("./authBackendRole").AuthBackendRole = null as any;
utilities.lazyLoad(exports, ["AuthBackendRole"], () => require("./authBackendRole"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "vault:alicloud/authBackendRole:AuthBackendRole":
                return new AuthBackendRole(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("vault", "alicloud/authBackendRole", _module)
