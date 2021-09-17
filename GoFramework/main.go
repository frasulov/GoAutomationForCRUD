package main

import (
	"GoFramework/config"
	"fmt"
)

func main() {
	config.Init()
	fmt.Println(config.Configuration.Libraries)
	config.GenerateFoldersFromConfiguration()
	//fmt.Println(config.Configuration)
	//newFolder := models.NewFolder(config.Configuration.Folder.Name, config.Configuration.Folder.Location)
	//newFolder.CreateFolder()
	//fmt.Println(config.Configuration.Project[0].App)
	//fmt.Println(config.Configuration.Project[0].App[0])
	//fmt.Println(config.Configuration.Project[0].App[0].Folder)
	//fmt.Println(config.Configuration.Project[0].App[0].Folder[0].Name)
	//fmt.Println(config.Configuration.Project[0].App[0].Folder[0].Files)
}
