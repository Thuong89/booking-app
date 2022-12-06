# This is about Go overal level and simple example

# Overview
* Every executable use in go file should come from a package.
* import package come from go standard library (see official Go website).
* 1 program can have many go files, file main.go (always specify 'package main' at the beginning line).
* The folder containing program should be initialize using `go mod init foldername` .
* To run the program `go run main.go`

# Program file
* example see file **main.go**
* what the example covers:
** Minimum struction of a main go program (package, import, code)
** using print
** using function

# Clean code
** no unused package
** no unused variable

# Learning code
* `var variablename = somevalue`
* `const variablename = some value`  means the variable not allow to change.
* define variable `var name datatype`.  Read more about datatype.
* short writing to define variable `name := somevalue` .
*  print variable or combination of two strings: always have automatically added a space at the end.
*  `Printf` to print formatted data. `%v` is a placeholder. Others format can be found in list in Go documentation under fmt package.
* `Scan` function to ask input value to variable by user.
* package time can be used to add waiting time.
* `//` to comment
* Array:index start 0
* for loop syntax ```for initialization; condition; update {
  statement(s)
}```
* syntax `++` means increment

# My reflection
Go syntax is quite similar to Java. Obviously the language is young ( invented 2007 by Google, on top of C C++ and Java). 

# Source: 
**TechWorldwithNana https://www.youtube.com/watch?v=yyUHQIec83I&ab_channel=TechWorldwithNana
** https://www.w3schools.com/
