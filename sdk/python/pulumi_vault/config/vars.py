# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

__config__ = pulumi.Config('vault')

address = __config__.get('address')
"""
URL of the root of the target Vault server.
"""

ca_cert_dir = __config__.get('caCertDir')
"""
Path to directory containing CA certificate files to validate the server's certificate.
"""

ca_cert_file = __config__.get('caCertFile')
"""
Path to a CA certificate file to validate the server's certificate.
"""

client_auths = __config__.get('clientAuths')
"""
Client authentication credentials.
"""

max_lease_ttl_seconds = __config__.get('maxLeaseTtlSeconds')
"""
Maximum TTL for secret leases requested by this provider
"""

max_retries = __config__.get('maxRetries')
"""
Maximum number of retries when a 5xx error code is encountered.
"""

namespace = __config__.get('namespace')
"""
The namespace to use. Available only for Vault Enterprise
"""

skip_tls_verify = __config__.get('skipTlsVerify')
"""
Set this to true only if the target Vault server is an insecure development instance.
"""

token = __config__.get('token')
"""
Token to use to authenticate to Vault.
"""

