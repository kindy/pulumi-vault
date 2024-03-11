// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vault.Database
{
    /// <summary>
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
    ///     var db = new Vault.Mount("db", new()
    ///     {
    ///         Path = "postgres",
    ///         Type = "database",
    ///     });
    /// 
    ///     var postgres = new Vault.Database.SecretBackendConnection("postgres", new()
    ///     {
    ///         Backend = db.Path,
    ///         AllowedRoles = new[]
    ///         {
    ///             "dev",
    ///             "prod",
    ///         },
    ///         Postgresql = new Vault.Database.Inputs.SecretBackendConnectionPostgresqlArgs
    ///         {
    ///             ConnectionUrl = "postgres://username:password@host:port/database",
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// &lt;!--End PulumiCodeChooser --&gt;
    /// 
    /// ## Import
    /// 
    /// Database secret backend connections can be imported using the `backend`, `/config/`, and the `name` e.g.
    /// 
    /// ```sh
    /// $ pulumi import vault:database/secretBackendConnection:SecretBackendConnection example postgres/config/postgres
    /// ```
    /// </summary>
    [VaultResourceType("vault:database/secretBackendConnection:SecretBackendConnection")]
    public partial class SecretBackendConnection : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A list of roles that are allowed to use this
        /// connection.
        /// </summary>
        [Output("allowedRoles")]
        public Output<ImmutableArray<string>> AllowedRoles { get; private set; } = null!;

        /// <summary>
        /// The unique name of the Vault mount to configure.
        /// </summary>
        [Output("backend")]
        public Output<string> Backend { get; private set; } = null!;

        /// <summary>
        /// A nested block containing configuration options for Cassandra connections.
        /// </summary>
        [Output("cassandra")]
        public Output<Outputs.SecretBackendConnectionCassandra?> Cassandra { get; private set; } = null!;

        /// <summary>
        /// A nested block containing configuration options for Couchbase connections.
        /// </summary>
        [Output("couchbase")]
        public Output<Outputs.SecretBackendConnectionCouchbase?> Couchbase { get; private set; } = null!;

        /// <summary>
        /// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
        /// </summary>
        [Output("data")]
        public Output<ImmutableDictionary<string, object>?> Data { get; private set; } = null!;

        /// <summary>
        /// A nested block containing configuration options for Elasticsearch connections.
        /// </summary>
        [Output("elasticsearch")]
        public Output<Outputs.SecretBackendConnectionElasticsearch?> Elasticsearch { get; private set; } = null!;

        /// <summary>
        /// A nested block containing configuration options for SAP HanaDB connections.
        /// </summary>
        [Output("hana")]
        public Output<Outputs.SecretBackendConnectionHana?> Hana { get; private set; } = null!;

        /// <summary>
        /// A nested block containing configuration options for InfluxDB connections.
        /// </summary>
        [Output("influxdb")]
        public Output<Outputs.SecretBackendConnectionInfluxdb?> Influxdb { get; private set; } = null!;

        /// <summary>
        /// A nested block containing configuration options for MongoDB connections.
        /// </summary>
        [Output("mongodb")]
        public Output<Outputs.SecretBackendConnectionMongodb?> Mongodb { get; private set; } = null!;

        /// <summary>
        /// A nested block containing configuration options for MongoDB Atlas connections.
        /// </summary>
        [Output("mongodbatlas")]
        public Output<Outputs.SecretBackendConnectionMongodbatlas?> Mongodbatlas { get; private set; } = null!;

        /// <summary>
        /// A nested block containing configuration options for MSSQL connections.
        /// </summary>
        [Output("mssql")]
        public Output<Outputs.SecretBackendConnectionMssql?> Mssql { get; private set; } = null!;

        /// <summary>
        /// A nested block containing configuration options for MySQL connections.
        /// </summary>
        [Output("mysql")]
        public Output<Outputs.SecretBackendConnectionMysql?> Mysql { get; private set; } = null!;

        /// <summary>
        /// A nested block containing configuration options for Aurora MySQL connections.
        /// </summary>
        [Output("mysqlAurora")]
        public Output<Outputs.SecretBackendConnectionMysqlAurora?> MysqlAurora { get; private set; } = null!;

        /// <summary>
        /// A nested block containing configuration options for legacy MySQL connections.
        /// </summary>
        [Output("mysqlLegacy")]
        public Output<Outputs.SecretBackendConnectionMysqlLegacy?> MysqlLegacy { get; private set; } = null!;

        /// <summary>
        /// A nested block containing configuration options for RDS MySQL connections.
        /// </summary>
        [Output("mysqlRds")]
        public Output<Outputs.SecretBackendConnectionMysqlRds?> MysqlRds { get; private set; } = null!;

        /// <summary>
        /// A unique name to give the database connection.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured namespace.
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Output("namespace")]
        public Output<string?> Namespace { get; private set; } = null!;

        /// <summary>
        /// A nested block containing configuration options for Oracle connections.
        /// </summary>
        [Output("oracle")]
        public Output<Outputs.SecretBackendConnectionOracle?> Oracle { get; private set; } = null!;

        /// <summary>
        /// Specifies the name of the plugin to use.
        /// </summary>
        [Output("pluginName")]
        public Output<string> PluginName { get; private set; } = null!;

        /// <summary>
        /// A nested block containing configuration options for PostgreSQL connections.
        /// </summary>
        [Output("postgresql")]
        public Output<Outputs.SecretBackendConnectionPostgresql?> Postgresql { get; private set; } = null!;

        /// <summary>
        /// A nested block containing configuration options for Redis connections.
        /// </summary>
        [Output("redis")]
        public Output<Outputs.SecretBackendConnectionRedis?> Redis { get; private set; } = null!;

        /// <summary>
        /// A nested block containing configuration options for Redis ElastiCache connections.
        /// 
        /// Exactly one of the nested blocks of configuration options must be supplied.
        /// </summary>
        [Output("redisElasticache")]
        public Output<Outputs.SecretBackendConnectionRedisElasticache?> RedisElasticache { get; private set; } = null!;

        /// <summary>
        /// Connection parameters for the redshift-database-plugin plugin.
        /// </summary>
        [Output("redshift")]
        public Output<Outputs.SecretBackendConnectionRedshift?> Redshift { get; private set; } = null!;

        /// <summary>
        /// A list of database statements to be executed to rotate the root user's credentials.
        /// </summary>
        [Output("rootRotationStatements")]
        public Output<ImmutableArray<string>> RootRotationStatements { get; private set; } = null!;

        /// <summary>
        /// A nested block containing configuration options for Snowflake connections.
        /// </summary>
        [Output("snowflake")]
        public Output<Outputs.SecretBackendConnectionSnowflake?> Snowflake { get; private set; } = null!;

        /// <summary>
        /// Whether the connection should be verified on
        /// initial configuration or not.
        /// </summary>
        [Output("verifyConnection")]
        public Output<bool?> VerifyConnection { get; private set; } = null!;


        /// <summary>
        /// Create a SecretBackendConnection resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public SecretBackendConnection(string name, SecretBackendConnectionArgs args, CustomResourceOptions? options = null)
            : base("vault:database/secretBackendConnection:SecretBackendConnection", name, args ?? new SecretBackendConnectionArgs(), MakeResourceOptions(options, ""))
        {
        }

        private SecretBackendConnection(string name, Input<string> id, SecretBackendConnectionState? state = null, CustomResourceOptions? options = null)
            : base("vault:database/secretBackendConnection:SecretBackendConnection", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing SecretBackendConnection resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static SecretBackendConnection Get(string name, Input<string> id, SecretBackendConnectionState? state = null, CustomResourceOptions? options = null)
        {
            return new SecretBackendConnection(name, id, state, options);
        }
    }

    public sealed class SecretBackendConnectionArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedRoles")]
        private InputList<string>? _allowedRoles;

        /// <summary>
        /// A list of roles that are allowed to use this
        /// connection.
        /// </summary>
        public InputList<string> AllowedRoles
        {
            get => _allowedRoles ?? (_allowedRoles = new InputList<string>());
            set => _allowedRoles = value;
        }

        /// <summary>
        /// The unique name of the Vault mount to configure.
        /// </summary>
        [Input("backend", required: true)]
        public Input<string> Backend { get; set; } = null!;

        /// <summary>
        /// A nested block containing configuration options for Cassandra connections.
        /// </summary>
        [Input("cassandra")]
        public Input<Inputs.SecretBackendConnectionCassandraArgs>? Cassandra { get; set; }

        /// <summary>
        /// A nested block containing configuration options for Couchbase connections.
        /// </summary>
        [Input("couchbase")]
        public Input<Inputs.SecretBackendConnectionCouchbaseArgs>? Couchbase { get; set; }

        [Input("data")]
        private InputMap<object>? _data;

        /// <summary>
        /// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
        /// </summary>
        public InputMap<object> Data
        {
            get => _data ?? (_data = new InputMap<object>());
            set => _data = value;
        }

        /// <summary>
        /// A nested block containing configuration options for Elasticsearch connections.
        /// </summary>
        [Input("elasticsearch")]
        public Input<Inputs.SecretBackendConnectionElasticsearchArgs>? Elasticsearch { get; set; }

        /// <summary>
        /// A nested block containing configuration options for SAP HanaDB connections.
        /// </summary>
        [Input("hana")]
        public Input<Inputs.SecretBackendConnectionHanaArgs>? Hana { get; set; }

        /// <summary>
        /// A nested block containing configuration options for InfluxDB connections.
        /// </summary>
        [Input("influxdb")]
        public Input<Inputs.SecretBackendConnectionInfluxdbArgs>? Influxdb { get; set; }

        /// <summary>
        /// A nested block containing configuration options for MongoDB connections.
        /// </summary>
        [Input("mongodb")]
        public Input<Inputs.SecretBackendConnectionMongodbArgs>? Mongodb { get; set; }

        /// <summary>
        /// A nested block containing configuration options for MongoDB Atlas connections.
        /// </summary>
        [Input("mongodbatlas")]
        public Input<Inputs.SecretBackendConnectionMongodbatlasArgs>? Mongodbatlas { get; set; }

        /// <summary>
        /// A nested block containing configuration options for MSSQL connections.
        /// </summary>
        [Input("mssql")]
        public Input<Inputs.SecretBackendConnectionMssqlArgs>? Mssql { get; set; }

        /// <summary>
        /// A nested block containing configuration options for MySQL connections.
        /// </summary>
        [Input("mysql")]
        public Input<Inputs.SecretBackendConnectionMysqlArgs>? Mysql { get; set; }

        /// <summary>
        /// A nested block containing configuration options for Aurora MySQL connections.
        /// </summary>
        [Input("mysqlAurora")]
        public Input<Inputs.SecretBackendConnectionMysqlAuroraArgs>? MysqlAurora { get; set; }

        /// <summary>
        /// A nested block containing configuration options for legacy MySQL connections.
        /// </summary>
        [Input("mysqlLegacy")]
        public Input<Inputs.SecretBackendConnectionMysqlLegacyArgs>? MysqlLegacy { get; set; }

        /// <summary>
        /// A nested block containing configuration options for RDS MySQL connections.
        /// </summary>
        [Input("mysqlRds")]
        public Input<Inputs.SecretBackendConnectionMysqlRdsArgs>? MysqlRds { get; set; }

        /// <summary>
        /// A unique name to give the database connection.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured namespace.
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// A nested block containing configuration options for Oracle connections.
        /// </summary>
        [Input("oracle")]
        public Input<Inputs.SecretBackendConnectionOracleArgs>? Oracle { get; set; }

        /// <summary>
        /// Specifies the name of the plugin to use.
        /// </summary>
        [Input("pluginName")]
        public Input<string>? PluginName { get; set; }

        /// <summary>
        /// A nested block containing configuration options for PostgreSQL connections.
        /// </summary>
        [Input("postgresql")]
        public Input<Inputs.SecretBackendConnectionPostgresqlArgs>? Postgresql { get; set; }

        /// <summary>
        /// A nested block containing configuration options for Redis connections.
        /// </summary>
        [Input("redis")]
        public Input<Inputs.SecretBackendConnectionRedisArgs>? Redis { get; set; }

        /// <summary>
        /// A nested block containing configuration options for Redis ElastiCache connections.
        /// 
        /// Exactly one of the nested blocks of configuration options must be supplied.
        /// </summary>
        [Input("redisElasticache")]
        public Input<Inputs.SecretBackendConnectionRedisElasticacheArgs>? RedisElasticache { get; set; }

        /// <summary>
        /// Connection parameters for the redshift-database-plugin plugin.
        /// </summary>
        [Input("redshift")]
        public Input<Inputs.SecretBackendConnectionRedshiftArgs>? Redshift { get; set; }

        [Input("rootRotationStatements")]
        private InputList<string>? _rootRotationStatements;

        /// <summary>
        /// A list of database statements to be executed to rotate the root user's credentials.
        /// </summary>
        public InputList<string> RootRotationStatements
        {
            get => _rootRotationStatements ?? (_rootRotationStatements = new InputList<string>());
            set => _rootRotationStatements = value;
        }

        /// <summary>
        /// A nested block containing configuration options for Snowflake connections.
        /// </summary>
        [Input("snowflake")]
        public Input<Inputs.SecretBackendConnectionSnowflakeArgs>? Snowflake { get; set; }

        /// <summary>
        /// Whether the connection should be verified on
        /// initial configuration or not.
        /// </summary>
        [Input("verifyConnection")]
        public Input<bool>? VerifyConnection { get; set; }

        public SecretBackendConnectionArgs()
        {
        }
        public static new SecretBackendConnectionArgs Empty => new SecretBackendConnectionArgs();
    }

    public sealed class SecretBackendConnectionState : global::Pulumi.ResourceArgs
    {
        [Input("allowedRoles")]
        private InputList<string>? _allowedRoles;

        /// <summary>
        /// A list of roles that are allowed to use this
        /// connection.
        /// </summary>
        public InputList<string> AllowedRoles
        {
            get => _allowedRoles ?? (_allowedRoles = new InputList<string>());
            set => _allowedRoles = value;
        }

        /// <summary>
        /// The unique name of the Vault mount to configure.
        /// </summary>
        [Input("backend")]
        public Input<string>? Backend { get; set; }

        /// <summary>
        /// A nested block containing configuration options for Cassandra connections.
        /// </summary>
        [Input("cassandra")]
        public Input<Inputs.SecretBackendConnectionCassandraGetArgs>? Cassandra { get; set; }

        /// <summary>
        /// A nested block containing configuration options for Couchbase connections.
        /// </summary>
        [Input("couchbase")]
        public Input<Inputs.SecretBackendConnectionCouchbaseGetArgs>? Couchbase { get; set; }

        [Input("data")]
        private InputMap<object>? _data;

        /// <summary>
        /// A map of sensitive data to pass to the endpoint. Useful for templated connection strings.
        /// </summary>
        public InputMap<object> Data
        {
            get => _data ?? (_data = new InputMap<object>());
            set => _data = value;
        }

        /// <summary>
        /// A nested block containing configuration options for Elasticsearch connections.
        /// </summary>
        [Input("elasticsearch")]
        public Input<Inputs.SecretBackendConnectionElasticsearchGetArgs>? Elasticsearch { get; set; }

        /// <summary>
        /// A nested block containing configuration options for SAP HanaDB connections.
        /// </summary>
        [Input("hana")]
        public Input<Inputs.SecretBackendConnectionHanaGetArgs>? Hana { get; set; }

        /// <summary>
        /// A nested block containing configuration options for InfluxDB connections.
        /// </summary>
        [Input("influxdb")]
        public Input<Inputs.SecretBackendConnectionInfluxdbGetArgs>? Influxdb { get; set; }

        /// <summary>
        /// A nested block containing configuration options for MongoDB connections.
        /// </summary>
        [Input("mongodb")]
        public Input<Inputs.SecretBackendConnectionMongodbGetArgs>? Mongodb { get; set; }

        /// <summary>
        /// A nested block containing configuration options for MongoDB Atlas connections.
        /// </summary>
        [Input("mongodbatlas")]
        public Input<Inputs.SecretBackendConnectionMongodbatlasGetArgs>? Mongodbatlas { get; set; }

        /// <summary>
        /// A nested block containing configuration options for MSSQL connections.
        /// </summary>
        [Input("mssql")]
        public Input<Inputs.SecretBackendConnectionMssqlGetArgs>? Mssql { get; set; }

        /// <summary>
        /// A nested block containing configuration options for MySQL connections.
        /// </summary>
        [Input("mysql")]
        public Input<Inputs.SecretBackendConnectionMysqlGetArgs>? Mysql { get; set; }

        /// <summary>
        /// A nested block containing configuration options for Aurora MySQL connections.
        /// </summary>
        [Input("mysqlAurora")]
        public Input<Inputs.SecretBackendConnectionMysqlAuroraGetArgs>? MysqlAurora { get; set; }

        /// <summary>
        /// A nested block containing configuration options for legacy MySQL connections.
        /// </summary>
        [Input("mysqlLegacy")]
        public Input<Inputs.SecretBackendConnectionMysqlLegacyGetArgs>? MysqlLegacy { get; set; }

        /// <summary>
        /// A nested block containing configuration options for RDS MySQL connections.
        /// </summary>
        [Input("mysqlRds")]
        public Input<Inputs.SecretBackendConnectionMysqlRdsGetArgs>? MysqlRds { get; set; }

        /// <summary>
        /// A unique name to give the database connection.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The namespace to provision the resource in.
        /// The value should not contain leading or trailing forward slashes.
        /// The `namespace` is always relative to the provider's configured namespace.
        /// *Available only for Vault Enterprise*.
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// A nested block containing configuration options for Oracle connections.
        /// </summary>
        [Input("oracle")]
        public Input<Inputs.SecretBackendConnectionOracleGetArgs>? Oracle { get; set; }

        /// <summary>
        /// Specifies the name of the plugin to use.
        /// </summary>
        [Input("pluginName")]
        public Input<string>? PluginName { get; set; }

        /// <summary>
        /// A nested block containing configuration options for PostgreSQL connections.
        /// </summary>
        [Input("postgresql")]
        public Input<Inputs.SecretBackendConnectionPostgresqlGetArgs>? Postgresql { get; set; }

        /// <summary>
        /// A nested block containing configuration options for Redis connections.
        /// </summary>
        [Input("redis")]
        public Input<Inputs.SecretBackendConnectionRedisGetArgs>? Redis { get; set; }

        /// <summary>
        /// A nested block containing configuration options for Redis ElastiCache connections.
        /// 
        /// Exactly one of the nested blocks of configuration options must be supplied.
        /// </summary>
        [Input("redisElasticache")]
        public Input<Inputs.SecretBackendConnectionRedisElasticacheGetArgs>? RedisElasticache { get; set; }

        /// <summary>
        /// Connection parameters for the redshift-database-plugin plugin.
        /// </summary>
        [Input("redshift")]
        public Input<Inputs.SecretBackendConnectionRedshiftGetArgs>? Redshift { get; set; }

        [Input("rootRotationStatements")]
        private InputList<string>? _rootRotationStatements;

        /// <summary>
        /// A list of database statements to be executed to rotate the root user's credentials.
        /// </summary>
        public InputList<string> RootRotationStatements
        {
            get => _rootRotationStatements ?? (_rootRotationStatements = new InputList<string>());
            set => _rootRotationStatements = value;
        }

        /// <summary>
        /// A nested block containing configuration options for Snowflake connections.
        /// </summary>
        [Input("snowflake")]
        public Input<Inputs.SecretBackendConnectionSnowflakeGetArgs>? Snowflake { get; set; }

        /// <summary>
        /// Whether the connection should be verified on
        /// initial configuration or not.
        /// </summary>
        [Input("verifyConnection")]
        public Input<bool>? VerifyConnection { get; set; }

        public SecretBackendConnectionState()
        {
        }
        public static new SecretBackendConnectionState Empty => new SecretBackendConnectionState();
    }
}
