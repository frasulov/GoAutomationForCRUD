package config

type Configurations struct {
	Project Project
	Server Server
	Database Database
	Libraries []Library
}

type Database struct {
	Host string
	User string
	Password string
	Dbname string
	Sslmode string
	Timezone string
	Port string
}

type Server struct {
	Port string
}

type Library struct {
	Library string
}

type Project struct {
	Name string
	Apps []AppArray
}

type AppArray struct {
	App App
}

type App struct {
	Name string
	Models []ModelArray
	Folders []FolderArray
}

type ModelArray struct {
	Model Model
}

type Model struct {
	Name string
	Fields []FieldsArray
}

type FieldsArray struct {
	Column Column
}

type Column struct {
	Name string
	Type string
}

type FolderArray struct {
	Folder Folder
}

type Folder struct {
	Name string
	Location string
	Files []FileArray
}

type FileArray struct {
	File File
}

type File struct {
	Name string
	Empty bool
}

//
//func default_tag(f File) string {
//
//	// TypeOf returns type of
//	// interface value passed to it
//	typ := reflect.TypeOf(f)
//
//	// checking if null string
//	if f.Empty {
//		f, _ := typ.FieldByName("Empty")
//
//		f.Name = f.Tag.Get("default")
//	}
//
//	return fmt.Sprintf("%s", f.Empty)
//}