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

1. Basic structure taken from the standards;
2. Layout of the main file;
3. File `go.mod` with project name;
4. Makefile with parameters:
   1. build: builds the application;
   2. run: start the app in dev mode;
   3. test: rub all tests;
   4. lint: checks your project for design errors;

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