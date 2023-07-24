# fs

Outputs strings with different styles

## Author

* [@smslash](https://github.com/smslash)

## Objectives

You must follow the same [instructions](https://01.alem.school/git/root/public/src/branch/master/subjects/ascii-art/README.md) as in the first subject but the second argument must be the name of the template. I know some templates may be hard to read, just do not obsess about it.

## Instructions

- Your project must be written in **Go**.

- The code must respect the [good practices](https://01.alem.school/git/root/public/src/branch/master/subjects/good-practices/README.md).

- It is recommended to have **test files** for [unit testing](https://go.dev/doc/tutorial/add-a-test).

- You can see all about the **banners** [here](https://01.alem.school/git/root/public/src/branch/master/subjects/ascii-art/).

- The usage must respect this format `go run . [STRING] [BANNER]`, any other formats must return the following usage message:

```bash
Usage: go run . [STRING] [BANNER]

EX: go run . something standard
```

If there are other `ascii-art` optional projects implemented, the program should accept other correctly formatted `[OPTION]` and/or `[BANNER]`.

Additionally, the program must still be able to run with a single `[STRING]` argument.

## Allowed packages

- Only the [standard Go](https://pkg.go.dev/std) packages are allowed
This project will help you learn about :

- The Go file system(**fs**) API

- Data manipulation