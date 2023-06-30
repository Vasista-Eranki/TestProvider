package client

import (
	"io/ioutil"
	"log"
	"os"
)

var (
	filePath = "d:\\helloWorld.txt"
)

type HelloWorldModel struct {
	Name string `json:"name"`
}

func (p HelloWorldModel) SetName(name string) string {
	ioutil.WriteFile(filePath, []byte(name), 0644)
	p.Name = name
	print("Set new name " + p.Name)
	return p.Name
}

func (p HelloWorldModel) UpdateName(name string) string {
	oldname := p.Name
	p.Name = name
	ioutil.WriteFile(filePath, []byte(name), 0644)
	print("updated Name from " + oldname + " to new name " + p.Name)
	return p.Name
}

func (p HelloWorldModel) GetName() string {
	name, _ := ioutil.ReadFile(filePath)
	return string(name)
}

func (p HelloWorldModel) RemoveName() {
	//os.Remove(filePath)
	e := os.Remove(filePath)
	if e != nil {
			log.Fatal(e)
	}
}

func GetNewModel() *HelloWorldModel {
	file, _ := os.Create(filePath )
	defer file.Close()
	return &HelloWorldModel{}
}