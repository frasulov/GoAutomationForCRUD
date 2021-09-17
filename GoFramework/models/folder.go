package models

import (
	"os"
)

type Folder struct {
	Name string
	Location string
}

func NewFolder(name, location string) *Folder {
	return &Folder{
		Name: name,
		Location: location,
	}
}

func (f * Folder) CreateFolder() error {
	pre, err := os.Getwd()
	if err != nil{
		return err
	}
	if err := os.Chdir("./result/"); err != nil {
		return err
	}
	if f.Location != "." || f.Location != "./"{
		if err := os.Chdir(f.Location); err != nil {
			os.Chdir(pre)
			return err
		}
	}
	if err := os.Chdir(f.Name); err == nil{
		return os.Chdir(pre)
	}
	if err:=os.Mkdir(f.Name, 0755);err!=nil{
		os.Chdir(pre)
		return err
	}
	return os.Chdir(pre)
}
