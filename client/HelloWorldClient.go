package client

type HelloWorldClient struct {
	Name string
}

func (p HelloWorldClient) SetName(name string) string {
	p.Name = name
	return "Set new name " + p.Name
}

func (p HelloWorldClient) UpdateName(name string) string {
	oldname := p.Name
	p.Name = name

	return "updated Name from " + oldname + " to new name " + p.Name
}

func (p HelloWorldClient) GetName() string {
	return p.Name
}

func (p HelloWorldClient) SayName() string {
	return "Hello " + p.GetName() + ". Happy learning!!"
}

func (p HelloWorldClient) RemoveName() {
	p.Name = ""
}

func GetNewClient() *HelloWorldClient {
	return &HelloWorldClient{}
}