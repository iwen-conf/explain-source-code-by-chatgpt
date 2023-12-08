# File: tools/go/callgraph/vta/testdata/src/ranges.go

ranges.go是一个用于测试的文件，用于展示Go语言中使用range操作符的一些特性和行为。

在该文件中定义了两个结构体I和B，分别代表了一个整数类型和一个布尔类型。

接下来定义了四个函数：

1. `func Foo()`：该函数使用range操作符迭代切片s并打印每个元素的值。它还使用了两个匿名函数，分别使用了变量`val1`和`val2`，并在for循环开始前和结束后打印它们的值。

2. `func Bar()`：该函数使用range操作符迭代数组a，并使用匿名函数打印每个元素的值。

3. `func Baz()`：该函数使用range操作符迭代map m，并使用匿名函数打印每个键值对的键和值。

4. `func Quux()`：该函数使用range操作符迭代字符串s，并使用匿名函数打印每个字符的值。

这些函数展示了在使用range操作符时，可以访问的变量和迭代的对象的不同类型。
