package sayhello


type HelloWorldStruct struct {
	Name string `tfsdk:"name"`
}

func GetNewStruct() *HelloWorldStruct {
	return &HelloWorldStruct{}
}