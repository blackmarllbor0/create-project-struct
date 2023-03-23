# create-project-struct

### You no longer have to think about creating the structure of your application yourself!

### About

This is a cli tool will create the base dirs and files of your app for you. 
The base structure is taken from this 
[repository](https://github.com/golang-standards/project-layout).

### How it works?

You must have Golang & Make installed to run the build.
Clone the repository on your locale machine, open a terminal in the current
dir and run the `go mod tidy` command.
Then in a terminal with the same dir, run the command 
`sudo make exec`.
Once finished, you can use the program by calling it in the
terminal with the command `cps`.

### Launch parameters

It's very simple! To create a project in the current 
directory, simply call `cps .`, and to create project
in a new directory, simply write what you want your project 
to be called. For example `cps my-app`.

### What does this program?

`cps` creates a root project structure, with ready-made 
templates. It also creates a main package with the template
output **Hello!** and a **go.mod** file with the project 
name and your locale language version.

### What is expected next?

Expected in upcoming updates:

1. Adding configuration for **golangci-lint**
2. Creating s local git repository.

### Conclusion

If you have any suggestions for changed or improving the
current project or would like to help | participate, you
can contact me at:

1. 3100194@gmail.com
2. t.me/blackmarllbor0

If you liked the project or found it useful, give it a star!
See you soon!