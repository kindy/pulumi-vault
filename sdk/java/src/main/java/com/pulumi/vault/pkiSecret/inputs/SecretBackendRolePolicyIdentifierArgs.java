// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretBackendRolePolicyIdentifierArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretBackendRolePolicyIdentifierArgs Empty = new SecretBackendRolePolicyIdentifierArgs();

    /**
     * The URL of the CPS for the policy identifier
     * 
     * Example usage:
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.vault.Mount;
     * import com.pulumi.vault.MountArgs;
     * import com.pulumi.vault.pkiSecret.SecretBackendRole;
     * import com.pulumi.vault.pkiSecret.SecretBackendRoleArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         var pki = new Mount(&#34;pki&#34;, MountArgs.builder()        
     *             .path(&#34;pki&#34;)
     *             .type(&#34;pki&#34;)
     *             .defaultLeaseTtlSeconds(3600)
     *             .maxLeaseTtlSeconds(86400)
     *             .build());
     * 
     *         var role = new SecretBackendRole(&#34;role&#34;, SecretBackendRoleArgs.builder()        
     *             .backend(pki.path())
     *             .ttl(3600)
     *             .allowIpSans(true)
     *             .keyType(&#34;rsa&#34;)
     *             .keyBits(4096)
     *             .allowedDomains(            
     *                 &#34;example.com&#34;,
     *                 &#34;my.domain&#34;)
     *             .allowSubdomains(true)
     *             .policyIdentifiers(            
     *                 Map.ofEntries(
     *                     Map.entry(&#34;oid&#34;, &#34;1.3.6.1.4.1.7.8&#34;),
     *                     Map.entry(&#34;notice&#34;, &#34;I am a user Notice&#34;)
     *                 ),
     *                 Map.ofEntries(
     *                     Map.entry(&#34;oid&#34;, &#34;1.3.6.1.4.1.44947.1.2.4&#34;),
     *                     Map.entry(&#34;cps&#34;, &#34;https://example.com&#34;)
     *                 ))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    @Import(name="cps")
    private @Nullable Output<String> cps;

    /**
     * @return The URL of the CPS for the policy identifier
     * 
     * Example usage:
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.vault.Mount;
     * import com.pulumi.vault.MountArgs;
     * import com.pulumi.vault.pkiSecret.SecretBackendRole;
     * import com.pulumi.vault.pkiSecret.SecretBackendRoleArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App {
     *     public static void main(String[] args) {
     *         Pulumi.run(App::stack);
     *     }
     * 
     *     public static void stack(Context ctx) {
     *         var pki = new Mount(&#34;pki&#34;, MountArgs.builder()        
     *             .path(&#34;pki&#34;)
     *             .type(&#34;pki&#34;)
     *             .defaultLeaseTtlSeconds(3600)
     *             .maxLeaseTtlSeconds(86400)
     *             .build());
     * 
     *         var role = new SecretBackendRole(&#34;role&#34;, SecretBackendRoleArgs.builder()        
     *             .backend(pki.path())
     *             .ttl(3600)
     *             .allowIpSans(true)
     *             .keyType(&#34;rsa&#34;)
     *             .keyBits(4096)
     *             .allowedDomains(            
     *                 &#34;example.com&#34;,
     *                 &#34;my.domain&#34;)
     *             .allowSubdomains(true)
     *             .policyIdentifiers(            
     *                 Map.ofEntries(
     *                     Map.entry(&#34;oid&#34;, &#34;1.3.6.1.4.1.7.8&#34;),
     *                     Map.entry(&#34;notice&#34;, &#34;I am a user Notice&#34;)
     *                 ),
     *                 Map.ofEntries(
     *                     Map.entry(&#34;oid&#34;, &#34;1.3.6.1.4.1.44947.1.2.4&#34;),
     *                     Map.entry(&#34;cps&#34;, &#34;https://example.com&#34;)
     *                 ))
     *             .build());
     * 
     *     }
     * }
     * ```
     * 
     */
    public Optional<Output<String>> cps() {
        return Optional.ofNullable(this.cps);
    }

    /**
     * A notice for the policy identifier
     * 
     */
    @Import(name="notice")
    private @Nullable Output<String> notice;

    /**
     * @return A notice for the policy identifier
     * 
     */
    public Optional<Output<String>> notice() {
        return Optional.ofNullable(this.notice);
    }

    /**
     * The OID for the policy identifier
     * 
     */
    @Import(name="oid", required=true)
    private Output<String> oid;

    /**
     * @return The OID for the policy identifier
     * 
     */
    public Output<String> oid() {
        return this.oid;
    }

    private SecretBackendRolePolicyIdentifierArgs() {}

    private SecretBackendRolePolicyIdentifierArgs(SecretBackendRolePolicyIdentifierArgs $) {
        this.cps = $.cps;
        this.notice = $.notice;
        this.oid = $.oid;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretBackendRolePolicyIdentifierArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretBackendRolePolicyIdentifierArgs $;

        public Builder() {
            $ = new SecretBackendRolePolicyIdentifierArgs();
        }

        public Builder(SecretBackendRolePolicyIdentifierArgs defaults) {
            $ = new SecretBackendRolePolicyIdentifierArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param cps The URL of the CPS for the policy identifier
         * 
         * Example usage:
         * ```java
         * package generated_program;
         * 
         * import com.pulumi.Context;
         * import com.pulumi.Pulumi;
         * import com.pulumi.core.Output;
         * import com.pulumi.vault.Mount;
         * import com.pulumi.vault.MountArgs;
         * import com.pulumi.vault.pkiSecret.SecretBackendRole;
         * import com.pulumi.vault.pkiSecret.SecretBackendRoleArgs;
         * import java.util.List;
         * import java.util.ArrayList;
         * import java.util.Map;
         * import java.io.File;
         * import java.nio.file.Files;
         * import java.nio.file.Paths;
         * 
         * public class App {
         *     public static void main(String[] args) {
         *         Pulumi.run(App::stack);
         *     }
         * 
         *     public static void stack(Context ctx) {
         *         var pki = new Mount(&#34;pki&#34;, MountArgs.builder()        
         *             .path(&#34;pki&#34;)
         *             .type(&#34;pki&#34;)
         *             .defaultLeaseTtlSeconds(3600)
         *             .maxLeaseTtlSeconds(86400)
         *             .build());
         * 
         *         var role = new SecretBackendRole(&#34;role&#34;, SecretBackendRoleArgs.builder()        
         *             .backend(pki.path())
         *             .ttl(3600)
         *             .allowIpSans(true)
         *             .keyType(&#34;rsa&#34;)
         *             .keyBits(4096)
         *             .allowedDomains(            
         *                 &#34;example.com&#34;,
         *                 &#34;my.domain&#34;)
         *             .allowSubdomains(true)
         *             .policyIdentifiers(            
         *                 Map.ofEntries(
         *                     Map.entry(&#34;oid&#34;, &#34;1.3.6.1.4.1.7.8&#34;),
         *                     Map.entry(&#34;notice&#34;, &#34;I am a user Notice&#34;)
         *                 ),
         *                 Map.ofEntries(
         *                     Map.entry(&#34;oid&#34;, &#34;1.3.6.1.4.1.44947.1.2.4&#34;),
         *                     Map.entry(&#34;cps&#34;, &#34;https://example.com&#34;)
         *                 ))
         *             .build());
         * 
         *     }
         * }
         * ```
         * 
         * @return builder
         * 
         */
        public Builder cps(@Nullable Output<String> cps) {
            $.cps = cps;
            return this;
        }

        /**
         * @param cps The URL of the CPS for the policy identifier
         * 
         * Example usage:
         * ```java
         * package generated_program;
         * 
         * import com.pulumi.Context;
         * import com.pulumi.Pulumi;
         * import com.pulumi.core.Output;
         * import com.pulumi.vault.Mount;
         * import com.pulumi.vault.MountArgs;
         * import com.pulumi.vault.pkiSecret.SecretBackendRole;
         * import com.pulumi.vault.pkiSecret.SecretBackendRoleArgs;
         * import java.util.List;
         * import java.util.ArrayList;
         * import java.util.Map;
         * import java.io.File;
         * import java.nio.file.Files;
         * import java.nio.file.Paths;
         * 
         * public class App {
         *     public static void main(String[] args) {
         *         Pulumi.run(App::stack);
         *     }
         * 
         *     public static void stack(Context ctx) {
         *         var pki = new Mount(&#34;pki&#34;, MountArgs.builder()        
         *             .path(&#34;pki&#34;)
         *             .type(&#34;pki&#34;)
         *             .defaultLeaseTtlSeconds(3600)
         *             .maxLeaseTtlSeconds(86400)
         *             .build());
         * 
         *         var role = new SecretBackendRole(&#34;role&#34;, SecretBackendRoleArgs.builder()        
         *             .backend(pki.path())
         *             .ttl(3600)
         *             .allowIpSans(true)
         *             .keyType(&#34;rsa&#34;)
         *             .keyBits(4096)
         *             .allowedDomains(            
         *                 &#34;example.com&#34;,
         *                 &#34;my.domain&#34;)
         *             .allowSubdomains(true)
         *             .policyIdentifiers(            
         *                 Map.ofEntries(
         *                     Map.entry(&#34;oid&#34;, &#34;1.3.6.1.4.1.7.8&#34;),
         *                     Map.entry(&#34;notice&#34;, &#34;I am a user Notice&#34;)
         *                 ),
         *                 Map.ofEntries(
         *                     Map.entry(&#34;oid&#34;, &#34;1.3.6.1.4.1.44947.1.2.4&#34;),
         *                     Map.entry(&#34;cps&#34;, &#34;https://example.com&#34;)
         *                 ))
         *             .build());
         * 
         *     }
         * }
         * ```
         * 
         * @return builder
         * 
         */
        public Builder cps(String cps) {
            return cps(Output.of(cps));
        }

        /**
         * @param notice A notice for the policy identifier
         * 
         * @return builder
         * 
         */
        public Builder notice(@Nullable Output<String> notice) {
            $.notice = notice;
            return this;
        }

        /**
         * @param notice A notice for the policy identifier
         * 
         * @return builder
         * 
         */
        public Builder notice(String notice) {
            return notice(Output.of(notice));
        }

        /**
         * @param oid The OID for the policy identifier
         * 
         * @return builder
         * 
         */
        public Builder oid(Output<String> oid) {
            $.oid = oid;
            return this;
        }

        /**
         * @param oid The OID for the policy identifier
         * 
         * @return builder
         * 
         */
        public Builder oid(String oid) {
            return oid(Output.of(oid));
        }

        public SecretBackendRolePolicyIdentifierArgs build() {
            $.oid = Objects.requireNonNull($.oid, "expected parameter 'oid' to be non-null");
            return $;
        }
    }

}
