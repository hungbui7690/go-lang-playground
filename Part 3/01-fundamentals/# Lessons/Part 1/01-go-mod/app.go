/*
	Organizing Code with Packages
	- https://codesandbox.io/
	- file name: hello-world.go or hello_world.go
	- go -> use "" or `` -> not ''
	- must have at least a single package in a go program
		-> can have 1 or more packages

	🚀 go run app.go second.go
	🌲 go run .


\\\\\\\\\\\\\\\\\\\\\\\\\\\

	The Important Of The Name _main
	- package main
	- we can use other names as well
	- main -> entry point -> where the app should start


\\\\\\\\\\\\\\\\\\\\\\\\\\\

	Understanding Go Modules & Building Go Programs
	- if we run "go build" -> cannot find main module
	- a module contains multiple packages

	++++++++++++++++
	- create <go.mod> file -> create module
		# "go mod init playground"
				-> can be a name
		# "go mod init github.com/yourorg/first-app"
			"go mod init example.com/first-app"
				-> or a path

	++++++++++++++++
	- after creating go.mod file -> we can run "go build"
		-> create <first-app.exe>

	++++++++++++++++
	- to run:
		-> ./first-app


\\\\\\\\\\\\\\\\\\\\\\\\\\\\

	The _main_ Function Is Important
	- we just can create main() in main package
	- only one main() function is allowed in a package


*/

package main 

func main() {
	sayHi()	// no need to import -> same package
}