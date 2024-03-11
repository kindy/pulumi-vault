// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Kmip
{
    /// <summary>
    /// Manages KMIP Secret roles in a Vault server. This feature requires
    /// Vault Enterprise. See the [Vault documentation](https://www.vaultproject.io/docs/secrets/kmip)
    /// for more information.
    /// 
    /// ## Example Usage
    /// 
    /// &lt;!--Start PulumiCodeChooser --&gt;
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vault = Pulumi.Vault;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var @default = new Vault.Kmip.SecretBackend("default", new()
    ///     {
    ///         Path = "kmip",
    ///         Description = "Vault KMIP backend",
    ///     });
    /// 
    ///     var dev = new Vault.Kmip.SecretScope("dev", new()
    ///     {
    ///         Path = @default.Path,
    ///         Scope = "dev",
    ///         Force = true,
    ///     });
    /// 
    ///     var admin = new Vault.Kmip.SecretRole("admin", new()
    ///     {
    ///         Path = dev.Path,
    ///         Scope = dev.Scope,
    ///         Role = "admin",
    ///         TlsClientKeyType = "ec",
    ///         TlsClientKeyBits = 256,
    ///         OperationActivate = true,
    ///         OperationGet = true,
    ///         OperationGetAttributes = true,
    ///         OperationCreate = true,
    ///         OperationDestroy = true,
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// KMIP Secret role can be imported using the `path`, e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:kmip/secretRole:SecretRole admin kmip
    /// ```
    /// </summary>
    [VaultResourceType("vault:kmip/secretRole:SecretRole")]
    public partial class SecretRole : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// Grant permission to use the KMIP Activate operation.
        /// </summary>
        [Output("operationActivate")]
        public Output<bool> OperationActivate { get; private set; } = null!;

        /// <summary>
        /// Grant permission to use the KMIP Add Attribute operation.
        /// </summary>
        [Output("operationAddAttribute")]
        public Output<bool> OperationAddAttribute { get; private set; } = null!;

        /// <summary>
        /// Grant all permissions to this role. May not be specified with any other `operation_*` params.
        /// </summary>
        [Output("operationAll")]
        public Output<bool> OperationAll { get; private set; } = null!;

        /// <summary>
        /// Grant permission to use the KMIP Create operation.
        /// </summary>
        [Output("operationCreate")]
        public Output<bool> OperationCreate { get; private set; } = null!;

        /// <summary>
        /// Grant permission to use the KMIP Destroy operation.
        /// </summary>
        [Output("operationDestroy")]
        public Output<bool> OperationDestroy { get; private set; } = null!;

        /// <summary>
        /// Grant permission to use the KMIP Discover Version operation.
        /// </summary>
        [Output("operationDiscoverVersions")]
        public Output<bool> OperationDiscoverVersions { get; private set; } = null!;

        /// <summary>
        /// Grant permission to use the KMIP Get operation.
        /// </summary>
        [Output("operationGet")]
        public Output<bool> OperationGet { get; private set; } = null!;

        /// <summary>
        /// Grant permission to use the KMIP Get Atrribute List operation.
        /// </summary>
        [Output("operationGetAttributeList")]
        public Output<bool> OperationGetAttributeList { get; private set; } = null!;

        /// <summary>
        /// Grant permission to use the KMIP Get Atrributes operation.
        /// </summary>
        [Output("operationGetAttributes")]
        public Output<bool> OperationGetAttributes { get; private set; } = null!;

        /// <summary>
        /// Grant permission to use the KMIP Get Locate operation.
        /// </summary>
        [Output("operationLocate")]
        public Output<bool> OperationLocate { get; private set; } = null!;

        /// <summary>
        /// Remove all permissions from this role. May not be specified with any other `operation_*` params.
        /// </summary>
        [Output("operationNone")]
        public Output<bool> OperationNone { get; private set; } = null!;

        /// <summary>
        /// Grant permission to use the KMIP Register operation.
        /// </summary>
        [Output("operationRegister")]
        public Output<bool> OperationRegister { get; private set; } = null!;

        /// <summary>
        /// Grant permission to use the KMIP Rekey operation.
        /// </summary>
        [Output("operationRekey")]
        public Output<bool> OperationRekey { get; private set; } = null!;

        /// <summary>
        /// Grant permission to use the KMIP Revoke operation.
        /// </summary>
        [Output("operationRevoke")]
        public Output<bool> OperationRevoke { get; private set; } = null!;

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `kmip`.
        /// </summary>
        [Output("path")]
        public Output<string> Path { get; private set; } = null!;

        /// <summary>
        /// Name of the role.
        /// </summary>
        [Output("role")]
        public Output<string> Role { get; private set; } = null!;

        /// <summary>
        /// Name of the scope.
        /// </summary>
        [Output("scope")]
        public Output<string> Scope { get; private set; } = null!;

        /// <summary>
        /// Client certificate key bits, valid values depend on key type.
        /// </summary>
        [Output("tlsClientKeyBits")]
        public Output<int?> TlsClientKeyBits { get; private set; } = null!;

        /// <summary>
        /// Client certificate key type, `rsa` or `ec`.
        /// </summary>
        [Output("tlsClientKeyType")]
        public Output<string?> TlsClientKeyType { get; private set; } = null!;

        /// <summary>
        /// Client certificate TTL in seconds.
        /// </summary>
        [Output("tlsClientTtl")]
        public Output<int?> TlsClientTtl { get; private set; } = null!;


        /// <summary>
        /// Create a SecretRole resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretRole(string name, SecretRoleArgs args, CustomResourceOptions? options = null)
            : base("vault:kmip/secretRole:SecretRole", name, args ?? new SecretRoleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretRole(string name, Input<string> id, SecretRoleState? state = null, CustomResourceOptions? options = null)
            : base("vault:kmip/secretRole:SecretRole", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing SecretRole resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretRole Get(string name, Input<string> id, SecretRoleState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretRole(name, id, state, options);
        }
    }

    public sealed class SecretRoleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Grant permission to use the KMIP Activate operation.
        /// </summary>
        [Input("operationActivate")]
        public Input<bool>? OperationActivate { get; set; }

        /// <summary>
        /// Grant permission to use the KMIP Add Attribute operation.
        /// </summary>
        [Input("operationAddAttribute")]
        public Input<bool>? OperationAddAttribute { get; set; }

        /// <summary>
        /// Grant all permissions to this role. May not be specified with any other `operation_*` params.
        /// </summary>
        [Input("operationAll")]
        public Input<bool>? OperationAll { get; set; }

        /// <summary>
        /// Grant permission to use the KMIP Create operation.
        /// </summary>
        [Input("operationCreate")]
        public Input<bool>? OperationCreate { get; set; }

        /// <summary>
        /// Grant permission to use the KMIP Destroy operation.
        /// </summary>
        [Input("operationDestroy")]
        public Input<bool>? OperationDestroy { get; set; }

        /// <summary>
        /// Grant permission to use the KMIP Discover Version operation.
        /// </summary>
        [Input("operationDiscoverVersions")]
        public Input<bool>? OperationDiscoverVersions { get; set; }

        /// <summary>
        /// Grant permission to use the KMIP Get operation.
        /// </summary>
        [Input("operationGet")]
        public Input<bool>? OperationGet { get; set; }

        /// <summary>
        /// Grant permission to use the KMIP Get Atrribute List operation.
        /// </summary>
        [Input("operationGetAttributeList")]
        public Input<bool>? OperationGetAttributeList { get; set; }

        /// <summary>
        /// Grant permission to use the KMIP Get Atrributes operation.
        /// </summary>
        [Input("operationGetAttributes")]
        public Input<bool>? OperationGetAttributes { get; set; }

        /// <summary>
        /// Grant permission to use the KMIP Get Locate operation.
        /// </summary>
        [Input("operationLocate")]
        public Input<bool>? OperationLocate { get; set; }

        /// <summary>
        /// Remove all permissions from this role. May not be specified with any other `operation_*` params.
        /// </summary>
        [Input("operationNone")]
        public Input<bool>? OperationNone { get; set; }

        /// <summary>
        /// Grant permission to use the KMIP Register operation.
        /// </summary>
        [Input("operationRegister")]
        public Input<bool>? OperationRegister { get; set; }

        /// <summary>
        /// Grant permission to use the KMIP Rekey operation.
        /// </summary>
        [Input("operationRekey")]
        public Input<bool>? OperationRekey { get; set; }

        /// <summary>
        /// Grant permission to use the KMIP Revoke operation.
        /// </summary>
        [Input("operationRevoke")]
        public Input<bool>? OperationRevoke { get; set; }

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `kmip`.
        /// </summary>
        [Input("path", required: true)]
        public Input<string> Path { get; set; } = null!;

        /// <summary>
        /// Name of the role.
        /// </summary>
        [Input("role", required: true)]
        public Input<string> Role { get; set; } = null!;

        /// <summary>
        /// Name of the scope.
        /// </summary>
        [Input("scope", required: true)]
        public Input<string> Scope { get; set; } = null!;

        /// <summary>
        /// Client certificate key bits, valid values depend on key type.
        /// </summary>
        [Input("tlsClientKeyBits")]
        public Input<int>? TlsClientKeyBits { get; set; }

        /// <summary>
        /// Client certificate key type, `rsa` or `ec`.
        /// </summary>
        [Input("tlsClientKeyType")]
        public Input<string>? TlsClientKeyType { get; set; }

        /// <summary>
        /// Client certificate TTL in seconds.
        /// </summary>
        [Input("tlsClientTtl")]
        public Input<int>? TlsClientTtl { get; set; }

        public SecretRoleArgs()
        {
        }
        public static new SecretRoleArgs Empty => new SecretRoleArgs();
    }

    public sealed class SecretRoleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Grant permission to use the KMIP Activate operation.
        /// </summary>
        [Input("operationActivate")]
        public Input<bool>? OperationActivate { get; set; }

        /// <summary>
        /// Grant permission to use the KMIP Add Attribute operation.
        /// </summary>
        [Input("operationAddAttribute")]
        public Input<bool>? OperationAddAttribute { get; set; }

        /// <summary>
        /// Grant all permissions to this role. May not be specified with any other `operation_*` params.
        /// </summary>
        [Input("operationAll")]
        public Input<bool>? OperationAll { get; set; }

        /// <summary>
        /// Grant permission to use the KMIP Create operation.
        /// </summary>
        [Input("operationCreate")]
        public Input<bool>? OperationCreate { get; set; }

        /// <summary>
        /// Grant permission to use the KMIP Destroy operation.
        /// </summary>
        [Input("operationDestroy")]
        public Input<bool>? OperationDestroy { get; set; }

        /// <summary>
        /// Grant permission to use the KMIP Discover Version operation.
        /// </summary>
        [Input("operationDiscoverVersions")]
        public Input<bool>? OperationDiscoverVersions { get; set; }

        /// <summary>
        /// Grant permission to use the KMIP Get operation.
        /// </summary>
        [Input("operationGet")]
        public Input<bool>? OperationGet { get; set; }

        /// <summary>
        /// Grant permission to use the KMIP Get Atrribute List operation.
        /// </summary>
        [Input("operationGetAttributeList")]
        public Input<bool>? OperationGetAttributeList { get; set; }

        /// <summary>
        /// Grant permission to use the KMIP Get Atrributes operation.
        /// </summary>
        [Input("operationGetAttributes")]
        public Input<bool>? OperationGetAttributes { get; set; }

        /// <summary>
        /// Grant permission to use the KMIP Get Locate operation.
        /// </summary>
        [Input("operationLocate")]
        public Input<bool>? OperationLocate { get; set; }

        /// <summary>
        /// Remove all permissions from this role. May not be specified with any other `operation_*` params.
        /// </summary>
        [Input("operationNone")]
        public Input<bool>? OperationNone { get; set; }

        /// <summary>
        /// Grant permission to use the KMIP Register operation.
        /// </summary>
        [Input("operationRegister")]
        public Input<bool>? OperationRegister { get; set; }

        /// <summary>
        /// Grant permission to use the KMIP Rekey operation.
        /// </summary>
        [Input("operationRekey")]
        public Input<bool>? OperationRekey { get; set; }

        /// <summary>
        /// Grant permission to use the KMIP Revoke operation.
        /// </summary>
        [Input("operationRevoke")]
        public Input<bool>? OperationRevoke { get; set; }

        /// <summary>
        /// The unique path this backend should be mounted at. Must
        /// not begin or end with a `/`. Defaults to `kmip`.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// Name of the role.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        /// <summary>
        /// Name of the scope.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// Client certificate key bits, valid values depend on key type.
        /// </summary>
        [Input("tlsClientKeyBits")]
        public Input<int>? TlsClientKeyBits { get; set; }

        /// <summary>
        /// Client certificate key type, `rsa` or `ec`.
        /// </summary>
        [Input("tlsClientKeyType")]
        public Input<string>? TlsClientKeyType { get; set; }

        /// <summary>
        /// Client certificate TTL in seconds.
        /// </summary>
        [Input("tlsClientTtl")]
        public Input<int>? TlsClientTtl { get; set; }

        public SecretRoleState()
        {
        }
        public static new SecretRoleState Empty => new SecretRoleState();
    }
}
