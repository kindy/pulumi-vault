// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./audit";
export * from "./authBackend";
export * from "./certAuthBackendRole";
export * from "./egpPolicy";
export * from "./getAuthBackend";
export * from "./getNomadAccessToken";
export * from "./getPolicyDocument";
export * from "./mfaDuo";
export * from "./mount";
export * from "./namespace";
export * from "./nomadSecretBackend";
export * from "./nomadSecretRole";
export * from "./passwordPolicy";
export * from "./policy";
export * from "./provider";
export * from "./quotaLeaseCount";
export * from "./quotaRateLimit";
export * from "./rgpPolicy";
export * from "./token";

// Export sub-modules:
import * as ad from "./ad";
import * as alicloud from "./alicloud";
import * as approle from "./approle";
import * as aws from "./aws";
import * as azure from "./azure";
import * as config from "./config";
import * as consul from "./consul";
import * as database from "./database";
import * as gcp from "./gcp";
import * as generic from "./generic";
import * as github from "./github";
import * as identity from "./identity";
import * as jwt from "./jwt";
import * as kubernetes from "./kubernetes";
import * as ldap from "./ldap";
import * as okta from "./okta";
import * as pkisecret from "./pkisecret";
import * as rabbitmq from "./rabbitmq";
import * as ssh from "./ssh";
import * as tokenauth from "./tokenauth";
import * as transform from "./transform";
import * as transit from "./transit";
import * as types from "./types";

export {
    ad,
    alicloud,
    approle,
    aws,
    azure,
    config,
    consul,
    database,
    gcp,
    generic,
    github,
    identity,
    jwt,
    kubernetes,
    ldap,
    okta,
    pkisecret,
    rabbitmq,
    ssh,
    tokenauth,
    transform,
    transit,
    types,
};

// Import resources to register:
import { Audit } from "./audit";
import { AuthBackend } from "./authBackend";
import { CertAuthBackendRole } from "./certAuthBackendRole";
import { EgpPolicy } from "./egpPolicy";
import { MfaDuo } from "./mfaDuo";
import { Mount } from "./mount";
import { Namespace } from "./namespace";
import { NomadSecretBackend } from "./nomadSecretBackend";
import { NomadSecretRole } from "./nomadSecretRole";
import { PasswordPolicy } from "./passwordPolicy";
import { Policy } from "./policy";
import { QuotaLeaseCount } from "./quotaLeaseCount";
import { QuotaRateLimit } from "./quotaRateLimit";
import { RgpPolicy } from "./rgpPolicy";
import { Token } from "./token";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "vault:index/audit:Audit":
                return new Audit(name, <any>undefined, { urn })
            case "vault:index/authBackend:AuthBackend":
                return new AuthBackend(name, <any>undefined, { urn })
            case "vault:index/certAuthBackendRole:CertAuthBackendRole":
                return new CertAuthBackendRole(name, <any>undefined, { urn })
            case "vault:index/egpPolicy:EgpPolicy":
                return new EgpPolicy(name, <any>undefined, { urn })
            case "vault:index/mfaDuo:MfaDuo":
                return new MfaDuo(name, <any>undefined, { urn })
            case "vault:index/mount:Mount":
                return new Mount(name, <any>undefined, { urn })
            case "vault:index/namespace:Namespace":
                return new Namespace(name, <any>undefined, { urn })
            case "vault:index/nomadSecretBackend:NomadSecretBackend":
                return new NomadSecretBackend(name, <any>undefined, { urn })
            case "vault:index/nomadSecretRole:NomadSecretRole":
                return new NomadSecretRole(name, <any>undefined, { urn })
            case "vault:index/passwordPolicy:PasswordPolicy":
                return new PasswordPolicy(name, <any>undefined, { urn })
            case "vault:index/policy:Policy":
                return new Policy(name, <any>undefined, { urn })
            case "vault:index/quotaLeaseCount:QuotaLeaseCount":
                return new QuotaLeaseCount(name, <any>undefined, { urn })
            case "vault:index/quotaRateLimit:QuotaRateLimit":
                return new QuotaRateLimit(name, <any>undefined, { urn })
            case "vault:index/rgpPolicy:RgpPolicy":
                return new RgpPolicy(name, <any>undefined, { urn })
            case "vault:index/token:Token":
                return new Token(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("vault", "index/audit", _module)
pulumi.runtime.registerResourceModule("vault", "index/authBackend", _module)
pulumi.runtime.registerResourceModule("vault", "index/certAuthBackendRole", _module)
pulumi.runtime.registerResourceModule("vault", "index/egpPolicy", _module)
pulumi.runtime.registerResourceModule("vault", "index/mfaDuo", _module)
pulumi.runtime.registerResourceModule("vault", "index/mount", _module)
pulumi.runtime.registerResourceModule("vault", "index/namespace", _module)
pulumi.runtime.registerResourceModule("vault", "index/nomadSecretBackend", _module)
pulumi.runtime.registerResourceModule("vault", "index/nomadSecretRole", _module)
pulumi.runtime.registerResourceModule("vault", "index/passwordPolicy", _module)
pulumi.runtime.registerResourceModule("vault", "index/policy", _module)
pulumi.runtime.registerResourceModule("vault", "index/quotaLeaseCount", _module)
pulumi.runtime.registerResourceModule("vault", "index/quotaRateLimit", _module)
pulumi.runtime.registerResourceModule("vault", "index/rgpPolicy", _module)
pulumi.runtime.registerResourceModule("vault", "index/token", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("vault", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:vault") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
