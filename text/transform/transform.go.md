# File: text/width/transform.go

在Go语言的text包中，text/width/transform.go文件的作用是提供一些用于文本宽度转换的函数和数据结构。

foldTransform、narrowTransform和wideTransform是用于执行文本宽度转换的结构体。

- foldTransform结构体用于将文本进行换行，根据提供的宽度边界将文本折叠成多行。
- narrowTransform结构体用于将文本中的全角字符（中文字符和全角符号）转换为半角字符，即将字符的宽度从2个字符变为1个字符。
- wideTransform结构体用于将文本中的半角字符转换为全角字符，即将字符的宽度从1个字符变为2个字符。

Span函数用于执行foldTransform结构体的文本宽度转换操作。它接受一个宽度边界和文本，并返回一个迭代器，用于依次返回转换后的折叠行。

Transform函数用于执行narrowTransform和wideTransform结构体的文本宽度转换操作。它接受一个将要被转换的字符串，并返回转换后的字符串。

这些函数和结构体一起提供了一些方便的方法来处理文本的宽度转换，尤其是在处理需要将文本格式化为特定宽度的场景中非常有用。
