# lib

[![Go Reference](https://pkg.go.dev/badge/github.com/limitcool/lib.svg)](https://pkg.go.dev/github.com/limitcool/lib)
[![Go Report Card](https://goreportcard.com/badge/github.com/limitcool/lib)](https://goreportcard.com/report/github.com/limitcool/lib)

This is a Go language library that provides some useful functions to facilitate rapid development.

[中文自述在这里.](README.zh_cn.md)

## Installation

To use this library, you need to install it in your project. You can install it by running the following command:

```bash
go get github.com/limitcool/lib
```

## Function List

### Debugging and Environment

- `SetDebugMode(debugFunction func())`: Set the debug mode based on the environment variable and execute a debug function.

### Slice Manipulation

- `InSlice(element T, Slice []T) bool`: Check if an element exists in a slice.
- `Unique(slice []T) []T`: Remove duplicate elements from a slice.
- `Push(slice []T, element T) []T`: Add an element to the end of a slice.
- `Pop(slice []T) ([]T, T)`: Remove and return the last element from a slice.
- `PopFront(slice []T) ([]T, T)`: Remove and return the first element from a slice.
- `Exclude(slice []T, element T) []T`: Return a slice without the specified element.
- `Reverse(s []T) []T`: Reverse the order of elements in a slice.

### Map Manipulation

- `Keys(m map[K]V) []K`: Return a slice containing all the keys in a map.

### Type Conversion

- `IntToBool(a T) bool`: Convert an integer to a boolean value.
- `ConvertAnyToInterface(anyValue *any.Any) (interface{}, error)`: Convert Google's protobuf `any.Any` type to Go's `interface{}` type.
- `ConvertInterfaceToAny(v interface{}) (*any.Any, error)`: Convert Go's `interface{}` type to Google's protobuf `any.Any` type.
- `MustConvertInterfaceToAny(v interface{}) *any.Any`: Similar to `ConvertInterfaceToAny`, but ignores errors.
- `MustConvertAnyToInterface(anyValue *any.Any) interface{}`: Similar to `ConvertAnyToInterface`, but ignores errors.

### File and Folder Operations

- `FileExists(filename string) bool`: Check if a file exists.
- `FolderExists(foldername string) bool`: Check if a folder exists.
- `FilesInFolder(dir, filename string) ([]string, error)`: Return a list of file paths containing the specified file in the given directory.
- `ReadFile(filename string) (string, error)`: Read the content of a file and return it.
- `CopyDir(src, dst, skip string) error`: Copy a directory from the source path to the destination path, skipping files with a specific suffix.
- `CopyDirHasSuffix(src, dst, suffix string) error`: Copy files with a specific suffix from the source directory to the destination directory, renaming the copied files.
- `CopyFile(src, dst string) error`: Copy a single file from the source path to the destination path.
- `ParentDir(p string) string`: Return the parent directory of the given path.
- `BaseDir(p string) string`: Return the base directory of the given path.
- `ParentBaseDir(p string) string`: Return the parent base directory of the given path.
- `CompressDir(dir string) error`: Compress the specified directory into a zip file.

### Hashing and Encryption

- `GetSha256Hash(str string) string`: Calculate the SHA-256 hash value of the given string and return it.

### Web

- `GetPage(c *gin.Context, DefaultPageSize int)`: Calculate the offset based on the page number parameter from the request and return it.
- `ParseToken(token, secret string)`: Parse and validate the given token, and return the claims contained in it along with any potential errors.

### MongoDB Specific

- `MustMongoSlice2StringSlice(i interface{}) []string`: Convert a MongoDB slice type to a Go string slice type, ignoring errors.
- `MongoSlice2StringSlice(i interface{}) ([]string, error)`: Convert a MongoDB slice type to a Go string slice type, returning errors.

## Usage Example

```go
package main

import (
    "fmt"
    "github.com/limitcool/lib"
)

func main() {
    slice := []int{1, 2, 3, 4, 3}
    uniqueSlice := lib.Unique(slice)
    fmt.Println(uniqueSlice) // Output: [1 2 3 4]
}
```