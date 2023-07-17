# File: inl_test.go

inl_test.go文件是用于测试Go编译器中的函数内联功能（function inlining）。函数内联是指编译器在编译时将函数调用替换为函数体的复制，以避免函数调用带来的开销并提高程序执行效率。该文件包含一系列测试函数，这些函数包括具有不同类型参数和返回值的常规函数和方法，以及使用内联优化的函数和方法。测试函数使用不同的编译选项进行编译，并通过比较函数执行时间和内存使用情况来比较函数内联和非内联版本的性能差异。这些测试可以帮助开发人员优化代码以提高程序性能，同时也确保Go编译器内联功能的正确性和稳定性。

## Functions:

### TestIntendedInlining





### collectInlCands





### TestIssue56044




