package sayhello

import (
	"context"
	"terraform-provider-helloworld/client"

	//"helloworld/client"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

type sayHelloResource struct {
	client *client.HelloWorldClient
}

func NewSayHelloResource() resource.Resource {
	return &sayHelloResource{}
}

func (*sayHelloResource) Schema(ctx context.Context, request resource.SchemaRequest, response *resource.SchemaResponse) {
	response.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"name" : schema.StringAttribute{
				Required : true,
				Description: "Name of the person for whom to tell Hello World!!!",
			},
		},
	}
}

func (r *sayHelloResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	
	r.client = req.ProviderData.(*client.HelloWorldClient)
}
func (r *sayHelloResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_say_hello"
}

func (r *sayHelloResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	//Read the existing data from plan
	var givenName string
	req.Plan.Get(ctx, &givenName)

	createResponseData := r.client.SetName(givenName)
	print(createResponseData)
	resp.State.Set(ctx, givenName)
}

func (r *sayHelloResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	givenName := r.client.GetName()
	resp.State.Set(ctx, &givenName)
}

func (r *sayHelloResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	//Read the existing data from plan
	var givenName string
	req.Plan.Get(ctx, &givenName)

	updateResponseData := r.client.SetName(givenName)
	print(updateResponseData)
	resp.State.Set(ctx, givenName)
}

func (r *sayHelloResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.client.RemoveName()
	resp.State.RemoveResource(ctx)
}