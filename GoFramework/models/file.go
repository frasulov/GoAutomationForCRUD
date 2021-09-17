package models

import (
	"fmt"
	"os"
)

type File struct {
	Name string
	Location string
}


func NewFile(name, location string) *File {
	return &File{
		Name: name,
		Location: location,
	}
}

func (f *File) CreateFile() error{
	pre, err := os.Getwd()
	if err != nil{
		return err
	}
	if err := os.Chdir("./result"); err!=nil{
		os.Chdir(pre)
		return err
	}
	if f.Location != "." || f.Location != "./"{
		if err := os.Chdir(f.Location); err != nil{
			os.Chdir(pre)
			return err
		}
	}
	if _, err := os.Create(f.Name); err != nil{
		return err
	}
	return os.Chdir(pre)
}

func (f* File) WriteDataToFile(text []byte, location string) error{
	fmt.Println(f.Location, f.Name)
	pre, err := os.Getwd()
	if err != nil{
		return err
	}
	if err := os.Chdir("./result"); err!=nil{
		os.Chdir(pre)
		return err
	}
	if f.Location != "." || f.Location != "./"{
		if err := os.Chdir(f.Location); err != nil{
			os.Chdir(pre)
			return err
		}
	}
	if err:= os.WriteFile(f.Name, text, 0644); err != nil{
		os.Chdir(pre)
		return err
	}
	return os.Chdir(pre)
}