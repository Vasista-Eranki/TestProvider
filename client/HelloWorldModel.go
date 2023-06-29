package client

type HelloWorldModel struct {
	Name string `json:"name"`
}

func (p HelloWorldModel) SetName(name string) string {

	p.Name = name
	print("Set new name " + p.Name)
	return p.Name
}

func (p HelloWorldModel) UpdateName(name string) string {
	oldname := p.Name
	p.Name = name
	print("updated Name from " + oldname + " to new name " + p.Name)
	return p.Name
}

func (p HelloWorldModel) GetName() string {
	return p.Name
}

func (p HelloWorldModel) RemoveName() {
	p.SetName("")
}

func GetNewModel() *HelloWorldModel {
	return &HelloWorldModel{}
}