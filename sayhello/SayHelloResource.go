package sayhello

import (
	"context"
	"terraform-provider-helloworld/client"

	//"helloworld/client"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

type sayHelloResource struct {
	client *client.HelloWorldModel
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
	
	r.client = req.ProviderData.(*client.HelloWorldModel)
}
func (r *sayHelloResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_say_hello"
}

func (r *sayHelloResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	//Read the existing data from plan
	var givenName HelloWorldStruct
	req.Plan.Get(ctx, &givenName)
	tflog.Info(ctx, ">>>>>>>>>>>>>>>>>>>NAME: " + givenName.Name)
	createResponseData := r.client.SetName(givenName.Name)

	toState := r.FlattenModel(createResponseData)
	print(createResponseData)
	resp.State.Set(ctx, toState)
}

func (r *sayHelloResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var givenName HelloWorldStruct
	req.State.Get(ctx, &givenName)
	
	tflog.Info(ctx, ">>>>>>>>>>>>>>>>>>> [READ] NAME: " + givenName.Name)

	readResponse  := r.FlattenModel(givenName.Name)
	resp.State.Set(ctx, &readResponse)
}

func (r *sayHelloResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	//Read the existing data from plan
	var givenName HelloWorldStruct
	req.Plan.Get(ctx, &givenName)

	updateResponseData := r.client.SetName(givenName.Name)

	toState := r.FlattenModel(updateResponseData)
	print(updateResponseData)

	resp.State.Set(ctx, toState)
}

func (r *sayHelloResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	r.client.RemoveName()
	resp.State.RemoveResource(ctx)
}


func (r *sayHelloResource) FlattenModel(name string) *HelloWorldStruct {
	givenName  := GetNewStruct()
	givenName.Name = name
	return givenName
}