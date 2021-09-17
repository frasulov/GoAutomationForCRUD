package config

import (
	"GoFramework/models"
	"archive/zip"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var Configuration Configurations
func Init() {
	viper.SetConfigName("config.yml")
	viper.AddConfigPath("./config")
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil{
		fmt.Printf("Error reading config file: %v", err.Error())
	}

	err := viper.Unmarshal(&Configuration)
	if err != nil{
		fmt.Printf("Unable to decode into struct %v", err.Error())
	}
}

func CreateDBConnection(appName string) error{
	modelFile := models.NewFile("db.go", Configuration.Project.Name+"/"+appName+"/config")
	if err := modelFile.CreateFile(); err != nil{
		fmt.Println(err)
		return err
	}

	pre, err := os.Getwd()
	if err != nil{
		return err
	}
	data, err := os.ReadFile("./data/db.txt")
	if err := os.WriteFile("./result/"+Configuration.Project.Name+"/"+appName+"/config/db.go", data, 0644); err != nil{
		fmt.Println(err)
		os.Chdir(pre)
		return err
	}
	return os.Chdir(pre)

}

func CreateModels(array []ModelArray, appName string) error{
	for _, model := range array {
		modelFile := models.NewFile(model.Model.Name+".go", Configuration.Project.Name+"/"+appName+"/models")
		if err := modelFile.CreateFile(); err != nil{
			return err
		}
		text := GenerateModelFileText(model.Model)
		if err := modelFile.WriteDataToFile(text, Configuration.Project.Name+"/"+appName+"/models"); err != nil{
			return err
		}
		text, err := GenerateRepoFileText(model.Model, appName)
		if err != nil{
			return err
		}
		repoFile := models.NewFile(model.Model.Name+"Repository.go", Configuration.Project.Name+"/"+appName+"/repository")
		if err := repoFile.CreateFile(); err != nil{
			return err
		}
		if err := repoFile.WriteDataToFile(text, "."); err != nil{
			return err
		}


		text, err = GenerateServiceFileText(model.Model, appName)
		if err != nil{
			return err
		}
		serviceFile := models.NewFile(model.Model.Name+"Service.go", Configuration.Project.Name+"/"+appName+"/service")
		if err := serviceFile.CreateFile(); err != nil{
			return err
		}
		if err := serviceFile.WriteDataToFile(text, "."); err != nil{
			return err
		}

		text, err = GenerateControllerFileText(model.Model, appName)
		if err != nil{
			return err
		}
		controllerFile := models.NewFile(model.Model.Name+"Controller.go", Configuration.Project.Name+"/"+appName+"/controller")
		if err := controllerFile.CreateFile(); err != nil{
			return err
		}
		if err := controllerFile.WriteDataToFile(text, "."); err != nil{
			return err
		}

	}
	return nil
}

func GenerateControllerFileText(model Model, appName string) ([]byte, error){
	data, err := os.ReadFile("./data/controller.txt")
	if err != nil{
		return nil, err
	}
	dataString := fmt.Sprintf(string(data), Configuration.Project.Name, appName, Configuration.Project.Name, appName, Configuration.Project.Name, appName, Configuration.Project.Name, appName,
		strings.ToLower(model.Name), model.Name, Configuration.Database.Host, Configuration.Database.User, Configuration.Database.Password, Configuration.Database.Dbname, Configuration.Database.Sslmode, Configuration.Database.Timezone, Configuration.Database.Port, strings.ToLower(model.Name),model.Name,strings.ToLower(model.Name),strings.ToLower(model.Name), model.Name, strings.ToLower(model.Name), model.Name,strings.ToLower(model.Name) ,model.Name, model.Name)
	return []byte(dataString), nil
}


func GenerateServiceFileText(model Model, appName string) ([]byte, error){
	data, err := os.ReadFile("./data/service.txt")
	if err != nil{
		return nil, err
	}
	dataString := fmt.Sprintf(string(data), Configuration.Project.Name, appName,Configuration.Project.Name, appName, model.Name, model.Name, model.Name,
		strings.ToLower(model.Name), model.Name, model.Name, model.Name, strings.ToLower(model.Name), strings.ToLower(model.Name), model.Name, model.Name, model.Name, strings.ToLower(model.Name), model.Name, model.Name, strings.ToLower(model.Name))
	return []byte(dataString), nil
}

func GenerateRepoFileText(model Model, appName string) ([]byte, error) {
	data, err := os.ReadFile("./data/repository.txt")
	if err != nil{
		return nil, err
	}
	dataString := fmt.Sprintf(string(data),Configuration.Project.Name,appName, model.Name, model.Name, model.Name, model.Name, strings.ToLower(model.Name), model.Name, strings.ToLower(model.Name), model.Name, strings.ToLower(model.Name), model.Name, model.Name,model.Name, strings.ToLower(model.Name), model.Name, strings.ToLower(model.Name), strings.ToLower(model.Name), strings.ToLower(model.Name), strings.ToLower(model.Name), model.Name,
		strings.ToLower(model.Name), model.Name, strings.ToLower(model.Name), strings.ToLower(model.Name), strings.ToLower(model.Name), model.Name, model.Name, model.Name, strings.ToLower(model.Name), model.Name, strings.ToLower(model.Name), strings.ToLower(model.Name), strings.ToLower(model.Name),
		strings.ToLower(model.Name), model.Name, model.Name, strings.ToLower(model.Name), model.Name)
	return []byte(dataString), nil
}

func GenerateModelFileText(model Model) []byte {
	data := "package models\n\n"
	middle := fmt.Sprintf("type %s struct {\n", model.Name)

	for _, field := range model.Fields {
		if strings.Contains(field.Column.Type, "time.") {
			data += "import (\n\t\"time\"\n\n)\n\n"
		}
		fieldText := fmt.Sprintf("\t%s\t\t%s\t\t`json:\"%s\"`\n", field.Column.Name, field.Column.Type, ToSnakeCase(field.Column.Name))
		middle += fieldText
	}
	middle += "}"
	return []byte(data+middle)
}

func GenerateFoldersFromConfiguration() error {
	if err:= CreateProject(Configuration.Project.Name); err != nil{
		return err
	}
	CreateMainFunctionFile(Configuration.Project.Name)
	for _, apps := range Configuration.Project.Apps {
		appFolder := models.NewFolder(apps.App.Name, Configuration.Project.Name)
		if err := appFolder.CreateFolder(); err != nil {
			return err
		}
		for _, folders := range apps.App.Folders {
			partFolder := models.NewFolder(folders.Folder.Name, Configuration.Project.Name+"/"+apps.App.Name)
			if err := partFolder.CreateFolder(); err != nil {
				return err
			}
				for _, goFiles := range folders.Folder.Files {
					file := models.NewFile(goFiles.File.Name, Configuration.Project.Name+"/"+apps.App.Name + "/" + folders.Folder.Name)
					if err := file.CreateFile(); err!=nil{
						return err
					}
					if err := file.WriteDataToFile([]byte("package "+folders.Folder.Name),  Configuration.Project.Name+"/"+apps.App.Name + "/" + folders.Folder.Name); err != nil{
						fmt.Println(err)
						return err
					}

				}
		}
		CreateModels(apps.App.Models, apps.App.Name)
		CreateDBConnection(apps.App.Name)
	}
	//InstallLibraries()
	CreateZipFile()
	return nil

}

func CreateMainFunctionFile(projectName string) error {
	pre, err := os.Getwd()
	if err != nil{
		return err
	}
	data, err := os.ReadFile("./data/main.txt")
	server := models.NewFile("server.go", projectName)
	server.CreateFile()
	if err := os.Chdir("./result/"+projectName+"/"); err != nil{
		os.Chdir(pre)
		return nil
	}
	dataString := fmt.Sprintf(string(data), Configuration.Server.Port)
	if err := os.WriteFile(server.Name, []byte(dataString), 0644); err != nil{
		os.Chdir(pre)
		return err
	}
	//if err := exec.Command("sudo","go", "mod", "init", projectName).Run(); err!=nil{
	//	os.Chdir(pre)
	//	return err
	//}
	return os.Chdir(pre)
}

func InstallLibraries() error{
	pre, err := os.Getwd()
	if err != nil{
		return err
	}
	if err := os.Chdir("./result/"+Configuration.Project.Name); err!=nil{
		os.Chdir(pre)
		return err
	}

	if err := exec.Command("sudo","go", "get", Configuration.Libraries[0].Library).Run(); err!=nil{
		os.Chdir(pre)
		return err
	}

	for i:=1; i<len(Configuration.Libraries); i++ {
		if err := exec.Command("go", "get", Configuration.Libraries[i].Library).Run(); err != nil{
			os.Chdir(pre)
			return err
		}
	}
	return os.Chdir(pre)

}

func CreateProject(fileName string) error{
	project := models.NewFolder(fileName, "./")
	return project.CreateFolder()
}

func CreateZipFile() error {
	archive, err := os.Create("./"+ToSnakeCase(Configuration.Project.Name)+".zip")
	if err != nil{
		return err
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)
	addFiles(zipWriter, "./result/", "")
	zipWriter.Close()
	os.RemoveAll("./result/"+Configuration.Project.Name)
	return nil
}

func addFiles(w *zip.Writer, basePath, baseInZip string) {
	// Open the Directory
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		fmt.Println(basePath + file.Name())
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(basePath + file.Name())
			if err != nil {
				fmt.Println(err)
			}

			// Add some files to the archive.
			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				fmt.Println(err)
			}
			_, err = f.Write(dat)
			if err != nil {
				fmt.Println(err)
			}
		} else if file.IsDir() {

			// Recurse
			newBase := basePath + file.Name() + "/"
			fmt.Println("Recursing and Adding SubDir: " + file.Name())
			fmt.Println("Recursing and Adding SubDir: " + newBase)

			addFiles(w, newBase, baseInZip  + file.Name() + "/")
		}
	}
}

var matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
var matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake  = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}