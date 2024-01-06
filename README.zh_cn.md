# lib

这是一个Go语言库，包含了一些实用的函数，方便您在开发中快速使用。

## 安装

要使用这个库，您需要先将其安装到您的项目中。可以通过以下命令安装：

```bash
go get github.com/limitcool/lib
```

## 函数列表

- `SetDebugMode(debugFunction func())`: 用于设置调试模式的函数。它接受一个无参数的函数 `debugFunction` 作为参数，并根据环境变量 `DEBUG_MODE` 的值来决定是否执行该函数。
- `InSlice(element T, Slice []T) bool`: 检查一个元素是否存在于切片中。
- `Unique(slice []T) []T`: 返回一个去重后的切片。
- `Push(slice []T, element T) []T`: 将一个元素添加到切片末尾。
- `Pop(slice []T) []T`: 删除切片末尾的元素并返回。
- `PopFront(slice []T) []T`: 删除切片开头的元素并返回。
- `Exclude(slice []T, element T) []T`: 返回一个不包含指定元素的切片。
- `Reverse(s []T) []T`: 反转切片中的元素顺序。
- `Keys(m map[K]V) []K`: 返回一个映射中的所有键的切片。
- `IntToBool(a T) bool`: 将一个整数转换为布尔值。
- `FileExists(filename string) bool`: 检查文件是否存在。
- `FolderExists(foldername string) bool`: 检查文件夹是否存在。
- `FilesInFolder(dir, filename string) ([]string, error)`: 返回给定目录下包含指定文件的文件路径列表。
- `ReadFile(filename string) (string, error)`: 读取给定文件的内容并返回。
- `CopyDir(src, dst, skip string) error`: 从源目录复制目录到目标目录，跳过具有特定后缀的文件。
- `CopyDirHasSuffix(src, dst, suffix string) error`: 从源目录复制具有特定后缀的文件到目标目录，并重命名复制的文件。
- `CopyFile(src, dst string) error`: 复制单个文件从源路径到目标路径。`
- `ParentDir(p string) string`: 返回给定路径的父目录。
- `BaseDir(p string) string`: 返回给定路径的基本目录。
- `ParentBaseDir(p string) string`: 返回给定路径的父目录的基本目录。
- `CompressDir(dir string) error`: 将指定目录压缩为zip文件。
- `ConvertAnyToInterface(anyValue *any.Any) (interface{}, error)`: 此函数将Google的protobuf的`any.Any`类型转换为Go的`interface{}`类型。它首先将`any.Any`类型解码为字节切片，然后使用`json.Unmarshal`将字节切片转换为`interface{}`类型。
- `ConvertInterfaceToAny(v interface{}) (*any.Any, error)`: 此函数将Go的`interface{}`类型转换为Google的protobuf的`any.Any`类型。它首先将`interface{}`类型编码为字节切片，然后使用`anypb.MarshalFrom`将字节切片转换为`any.Any`类型。
- `MustConvertInterfaceToAny(v interface{}) *any.Any`: 此函数与`ConvertInterfaceToAny`功能相似，不同之处在于它忽略了可能的错误并返回一个未初始化的`any.Any`对象。
- `MustConvertAnyToInterface(anyValue *any.Any) interface{}`: 此函数与`ConvertAnyToInterface`功能相似，不同之处在于它忽略了可能的错误并返回一个未初始化的`interface{}`对象。
- `MustMongoSlice2StringSlice(i interface{}) []string`: 此函数将MongoDB的切片类型转换为Go的字符串切片类型。如果输入的接口类型是`primitive.A`类型，它会遍历该切片并将每个元素转换为字符串，然后将这些字符串添加到新的字符串切片中并返回。如果输入的接口类型不是`primitive.A`类型，它将尝试使用`cast.ToStringSliceE`函数将输入转换为字符串切片。如果转换失败，它将返回一个空的字符串切片。
- `MongoSlice2StringSlice(i interface{}) ([]string, error)`: 此函数与`MustMongoSlice2StringSlice`功能相似，但它返回一个字符串切片和一个错误。如果输入的接口类型是`primitive.A`类型，它会遍历该切片并将每个元素转换为字符串，然后将这些字符串添加到新的字符串切片中并返回。如果输入的接口类型不是`primitive.A`类型，它将尝试使用`cast.ToStringSliceE`函数将输入转换为字符串切片。如果转换失败，它将返回一个空的字符串切片和一个错误。
- `GetSha256Hash(str string) string`: 计算给定字符串的SHA-256哈希值并返回。
- `GetPage(c *gin.Context, DefaultPageSize int)`: 根据请求中的页码参数计算偏移量并返回。

- `ParseToken(token, secret string)`: 解析和校验给定的token，并返回其包含的声明和可能的错误。

## 使用示例

```go
package main

import (
    "fmt"
    "github.com/limitcool/lib" // 请替换为您的实际地址
)

func main() {
    slice := []int{1, 2, 3, 4, 3}
    uniqueSlice := lib.Unique(slice)
    fmt.Println(uniqueSlice) // 输出: [1 2 3 4]
}
