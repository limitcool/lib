# lib

[![Go Reference](https://pkg.go.dev/badge/github.com/limitcool/lib.svg)](https://pkg.go.dev/github.com/limitcool/lib)

This is a Go language library that provides some useful functions to facilitate rapid development.

[中文自述在这里.](README.zh_cn.md)

## Installation

To use this library, you need to install it in your project. You can install it by running the following command:

```bash
go get github.com/limitcool/lib
```

## Function List

- `SetDebugMode(debugFunction func())`: A function to set the debug mode. It accepts a function `debugFunction` without arguments as a parameter, and based on the value of the environment variable `DEBUG_MODE`, it decides whether to execute the function.
- `InSlice(element T, Slice []T) bool`: Checks if an element exists in a slice.
- `Unique(slice []T) []T`: Returns a slice with duplicate elements removed.
- `Push(slice []T, element T) []T`: Adds an element to the end of a slice.
- `Pop(slice []T) []T`: Removes the last element from a slice and returns it.
- `PopFront(slice []T) []T`: Removes the first element from a slice and returns it.
- `Exclude(slice []T, element T) []T`: Returns a slice that does not contain the specified element.
- `Reverse(s []T) []T`: Reverses the order of elements in a slice.
- `Keys(m map[K]V) []K`: Returns a slice containing all the keys in a map.
- `IntToBool(a T) bool`: Converts an integer to a boolean value.
- `FileExists(filename string) bool`: Checks if a file exists.
- `FolderExists(foldername string) bool`: Checks if a folder exists.
- `FilesInFolder(dir, filename string) ([]string, error)`: Returns a list of file paths containing the specified file in the given directory.
- `ReadFile(filename string) (string, error)`: Reads the content of a file and returns it.
- `CopyDir(src, dst, skip string) error`: Copies a directory from the source path to the destination path, skipping files with a specific suffix.
- `CopyDirHasSuffix(src, dst, suffix string) error`: Copies files with a specific suffix from the source directory to the destination directory, renaming the copied files.
- `CopyFile(src, dst string) error`: Copies a single file from the source path to the destination path.
- `ParentDir(p string) string`: Returns the parent directory of the given path.
- `BaseDir(p string) string`: Returns the base directory of the given path.
- `ParentBaseDir(p string) string`: Returns the parent base directory of the given path.
- `CompressDir(dir string) error`: Compresses the specified directory into a zip file.
- `ConvertAnyToInterface(anyValue *any.Any) (interface{}, error)`: This function converts a Google's protobuf's `any.Any` type into a Go's `interface{}` type. It first decodes the `any.Any` type into a byte slice and then uses `json.Unmarshal` to convert the byte slice into an `interface{}` type.
- `ConvertInterfaceToAny(v interface{}) (*any.Any, error)`: This function converts a Go's `interface{}` type into Google's protobuf's `any.Any` type. It first encodes the `interface{}` type into a byte slice and then uses `anypb.MarshalFrom` to convert the byte slice into an `any.Any` type.
- `MustConvertInterfaceToAny(v interface{}) *any.Any`: This function is similar to `ConvertInterfaceToAny`, but it ignores any potential errors and returns an uninitialized `any.Any` object instead.
- `MustConvertAnyToInterface(anyValue *any.Any) interface{}`: This function is similar to `ConvertAnyToInterface`, but it ignores any potential errors and returns an uninitialized `interface{}` object instead.

- `MustMongoSlice2StringSlice(i interface{}) []string`: This function converts a MongoDB slice type to a Go string slice type. If the input interface type is of type `primitive.A`, it iterates over the slice and converts each element to a string, then adds these strings to a new string slice and returns it. If the input interface type is not `primitive.A`, it attempts to convert the input to a string slice using the `cast.ToStringSliceE` function. If the conversion fails, it returns an empty string slice.
- `MongoSlice2StringSlice(i interface{}) ([]string, error)`: This function is similar to `MustMongoSlice2StringSlice`, but it returns a string slice and an error. If the input interface type is of type `primitive.A`, it iterates over the slice and converts each element to a string, then adds these strings to a new string slice and returns it. If the input interface type is not `primitive.A`, it attempts to convert the input to a string slice using the `cast.ToStringSliceE` function. If the conversion fails, it returns an empty string slice and an error.
- `GetSha256Hash(str string) string`: This function calculates the SHA-256 hash value of the given string and returns it.
- `GetPage(c *gin.Context, DefaultPageSize int)`: This function calculates the offset based on the page number parameter from the request and returns it.
- `ParseToken(token, secret string)`: This function parses and validates the given token, and returns the claims contained in it along with any potential errors.

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