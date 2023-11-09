// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package github

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
)

// Manages policy mappings for Github Teams authenticated via Github. See the [Vault
// documentation](https://www.vaultproject.io/docs/auth/github/) for more
// information.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/pulumi/pulumi-vault/sdk/v5/go/vault/github"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			example, err := github.NewAuthBackend(ctx, "example", &github.AuthBackendArgs{
//				Organization: pulumi.String("myorg"),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = github.NewTeam(ctx, "tfDevs", &github.TeamArgs{
//				Backend: example.ID(),
//				Team:    pulumi.String("terraform-developers"),
//				Policies: pulumi.StringArray{
//					pulumi.String("developer"),
//					pulumi.String("read-only"),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// Github team mappings can be imported using the `path`, e.g.
//
// ```sh
//
//	$ pulumi import vault:github/team:Team tf_devs auth/github/map/teams/terraform-developers
//
// ```
type Team struct {
	pulumi.CustomResourceState

	// Path where the github auth backend is mounted. Defaults to `github`
	// if not specified.
	Backend pulumi.StringPtrOutput `pulumi:"backend"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrOutput `pulumi:"namespace"`
	// An array of strings specifying the policies to be set on tokens
	// issued using this role.
	Policies pulumi.StringArrayOutput `pulumi:"policies"`
	// GitHub team name in "slugified" format.
	Team pulumi.StringOutput `pulumi:"team"`
}

// NewTeam registers a new resource with the given unique name, arguments, and options.
func NewTeam(ctx *pulumi.Context,
	name string, args *TeamArgs, opts ...pulumi.ResourceOption) (*Team, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Team == nil {
		return nil, errors.New("invalid value for required argument 'Team'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Team
	err := ctx.RegisterResource("vault:github/team:Team", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTeam gets an existing Team resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTeam(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TeamState, opts ...pulumi.ResourceOption) (*Team, error) {
	var resource Team
	err := ctx.ReadResource("vault:github/team:Team", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Team resources.
type teamState struct {
	// Path where the github auth backend is mounted. Defaults to `github`
	// if not specified.
	Backend *string `pulumi:"backend"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// An array of strings specifying the policies to be set on tokens
	// issued using this role.
	Policies []string `pulumi:"policies"`
	// GitHub team name in "slugified" format.
	Team *string `pulumi:"team"`
}

type TeamState struct {
	// Path where the github auth backend is mounted. Defaults to `github`
	// if not specified.
	Backend pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// An array of strings specifying the policies to be set on tokens
	// issued using this role.
	Policies pulumi.StringArrayInput
	// GitHub team name in "slugified" format.
	Team pulumi.StringPtrInput
}

func (TeamState) ElementType() reflect.Type {
	return reflect.TypeOf((*teamState)(nil)).Elem()
}

type teamArgs struct {
	// Path where the github auth backend is mounted. Defaults to `github`
	// if not specified.
	Backend *string `pulumi:"backend"`
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace *string `pulumi:"namespace"`
	// An array of strings specifying the policies to be set on tokens
	// issued using this role.
	Policies []string `pulumi:"policies"`
	// GitHub team name in "slugified" format.
	Team string `pulumi:"team"`
}

// The set of arguments for constructing a Team resource.
type TeamArgs struct {
	// Path where the github auth backend is mounted. Defaults to `github`
	// if not specified.
	Backend pulumi.StringPtrInput
	// The namespace to provision the resource in.
	// The value should not contain leading or trailing forward slashes.
	// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
	// *Available only for Vault Enterprise*.
	Namespace pulumi.StringPtrInput
	// An array of strings specifying the policies to be set on tokens
	// issued using this role.
	Policies pulumi.StringArrayInput
	// GitHub team name in "slugified" format.
	Team pulumi.StringInput
}

func (TeamArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*teamArgs)(nil)).Elem()
}

type TeamInput interface {
	pulumi.Input

	ToTeamOutput() TeamOutput
	ToTeamOutputWithContext(ctx context.Context) TeamOutput
}

func (*Team) ElementType() reflect.Type {
	return reflect.TypeOf((**Team)(nil)).Elem()
}

func (i *Team) ToTeamOutput() TeamOutput {
	return i.ToTeamOutputWithContext(context.Background())
}

func (i *Team) ToTeamOutputWithContext(ctx context.Context) TeamOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamOutput)
}

func (i *Team) ToOutput(ctx context.Context) pulumix.Output[*Team] {
	return pulumix.Output[*Team]{
		OutputState: i.ToTeamOutputWithContext(ctx).OutputState,
	}
}

// TeamArrayInput is an input type that accepts TeamArray and TeamArrayOutput values.
// You can construct a concrete instance of `TeamArrayInput` via:
//
//	TeamArray{ TeamArgs{...} }
type TeamArrayInput interface {
	pulumi.Input

	ToTeamArrayOutput() TeamArrayOutput
	ToTeamArrayOutputWithContext(context.Context) TeamArrayOutput
}

type TeamArray []TeamInput

func (TeamArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Team)(nil)).Elem()
}

func (i TeamArray) ToTeamArrayOutput() TeamArrayOutput {
	return i.ToTeamArrayOutputWithContext(context.Background())
}

func (i TeamArray) ToTeamArrayOutputWithContext(ctx context.Context) TeamArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamArrayOutput)
}

func (i TeamArray) ToOutput(ctx context.Context) pulumix.Output[[]*Team] {
	return pulumix.Output[[]*Team]{
		OutputState: i.ToTeamArrayOutputWithContext(ctx).OutputState,
	}
}

// TeamMapInput is an input type that accepts TeamMap and TeamMapOutput values.
// You can construct a concrete instance of `TeamMapInput` via:
//
//	TeamMap{ "key": TeamArgs{...} }
type TeamMapInput interface {
	pulumi.Input

	ToTeamMapOutput() TeamMapOutput
	ToTeamMapOutputWithContext(context.Context) TeamMapOutput
}

type TeamMap map[string]TeamInput

func (TeamMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Team)(nil)).Elem()
}

func (i TeamMap) ToTeamMapOutput() TeamMapOutput {
	return i.ToTeamMapOutputWithContext(context.Background())
}

func (i TeamMap) ToTeamMapOutputWithContext(ctx context.Context) TeamMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamMapOutput)
}

func (i TeamMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*Team] {
	return pulumix.Output[map[string]*Team]{
		OutputState: i.ToTeamMapOutputWithContext(ctx).OutputState,
	}
}

type TeamOutput struct{ *pulumi.OutputState }

func (TeamOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Team)(nil)).Elem()
}

func (o TeamOutput) ToTeamOutput() TeamOutput {
	return o
}

func (o TeamOutput) ToTeamOutputWithContext(ctx context.Context) TeamOutput {
	return o
}

func (o TeamOutput) ToOutput(ctx context.Context) pulumix.Output[*Team] {
	return pulumix.Output[*Team]{
		OutputState: o.OutputState,
	}
}

// Path where the github auth backend is mounted. Defaults to `github`
// if not specified.
func (o TeamOutput) Backend() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.StringPtrOutput { return v.Backend }).(pulumi.StringPtrOutput)
}

// The namespace to provision the resource in.
// The value should not contain leading or trailing forward slashes.
// The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
// *Available only for Vault Enterprise*.
func (o TeamOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Team) pulumi.StringPtrOutput { return v.Namespace }).(pulumi.StringPtrOutput)
}

// An array of strings specifying the policies to be set on tokens
// issued using this role.
func (o TeamOutput) Policies() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Team) pulumi.StringArrayOutput { return v.Policies }).(pulumi.StringArrayOutput)
}

// GitHub team name in "slugified" format.
func (o TeamOutput) Team() pulumi.StringOutput {
	return o.ApplyT(func(v *Team) pulumi.StringOutput { return v.Team }).(pulumi.StringOutput)
}

type TeamArrayOutput struct{ *pulumi.OutputState }

func (TeamArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Team)(nil)).Elem()
}

func (o TeamArrayOutput) ToTeamArrayOutput() TeamArrayOutput {
	return o
}

func (o TeamArrayOutput) ToTeamArrayOutputWithContext(ctx context.Context) TeamArrayOutput {
	return o
}

func (o TeamArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*Team] {
	return pulumix.Output[[]*Team]{
		OutputState: o.OutputState,
	}
}

func (o TeamArrayOutput) Index(i pulumi.IntInput) TeamOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Team {
		return vs[0].([]*Team)[vs[1].(int)]
	}).(TeamOutput)
}

type TeamMapOutput struct{ *pulumi.OutputState }

func (TeamMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Team)(nil)).Elem()
}

func (o TeamMapOutput) ToTeamMapOutput() TeamMapOutput {
	return o
}

func (o TeamMapOutput) ToTeamMapOutputWithContext(ctx context.Context) TeamMapOutput {
	return o
}

func (o TeamMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*Team] {
	return pulumix.Output[map[string]*Team]{
		OutputState: o.OutputState,
	}
}

func (o TeamMapOutput) MapIndex(k pulumi.StringInput) TeamOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Team {
		return vs[0].(map[string]*Team)[vs[1].(string)]
	}).(TeamOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TeamInput)(nil)).Elem(), &Team{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamArrayInput)(nil)).Elem(), TeamArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TeamMapInput)(nil)).Elem(), TeamMap{})
	pulumi.RegisterOutputType(TeamOutput{})
	pulumi.RegisterOutputType(TeamArrayOutput{})
	pulumi.RegisterOutputType(TeamMapOutput{})
}
