package provider

import (
	"context"
	//"helloworld/client"
	//"helloworld/sayhello"

	"terraform-provider-helloworld/client"
	"terraform-provider-helloworld/sayhello"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

type pfprovider struct{
	client *client.HelloWorldModel
}


func (p *pfprovider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "helloworld"
}

func (p *pfprovider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	cli := client.GetNewModel()
	print(cli.Name)
	p.client = cli
	resp.ResourceData = cli
	resp.DataSourceData = cli
}

func (p *pfprovider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		sayhello.NewSayHelloResource,
	}
}

func (p *pfprovider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}


func (p *pfprovider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{}
}


func New() provider.Provider {
	return &pfprovider{}
}
