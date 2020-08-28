module github.com/pulumi/pulumi-vault/provider/v2

go 1.14

require (
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.5.2
	github.com/pulumi/pulumi/sdk/v2 v2.5.1-0.20200630091945-bb358c4d2173
	github.com/terraform-providers/terraform-provider-vault v1.9.1-0.20200827211630-a5872f788451
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
