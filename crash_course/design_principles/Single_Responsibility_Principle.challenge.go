package main

import "fmt"

type Appconfig struct {
	appName string
	version int
}

type AppconfigDisplay struct {}

func (AppconfigDisplay) display_function(a Appconfig)(string,int){
	return a.appName , a.version
}

func main() {
	app:=Appconfig{appName: "hello", version:1}
	app_display_class := AppconfigDisplay{}
	

	display_name, display_version := app_display_class.display_function(app)

	fmt.Println(display_name,display_version)


}