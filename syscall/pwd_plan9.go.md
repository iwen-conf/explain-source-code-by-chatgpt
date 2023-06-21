# File: pwd_plan9.go

pwd_plan9.go是Go语言标准库中syscall包的一个子文件，主要用于实现在Plan 9操作系统中获取和设置工作目录的相关功能。Plan 9是一种由贝尔实验室开发的分布式操作系统，使用一种称为“文件是命名的管道”的文件系统来组织文件和目录，工作目录是用户当前工作所在目录。pwd_plan9.go文件中包含了与工作目录相关的系统调用函数的封装实现，包括getwd(), chdir(), getcwd()等，用于获取、更改和查询当前工作目录路径。这些函数的底层实现是通过系统调用方式调用Plan 9操作系统内核中的相关函数。通过调用这些函数，Go语言的程序员能够方便地在Plan 9操作系统中管理工作目录，从而更好地控制系统的执行流程。




---

### Var:

### wdmu

在 "pwd_plan9.go" 文件中， "wdmu" 是一个类型为 "sync.Mutex" 的变量。它用于保护进程或线程安全地在 Plan 9 中访问工作目录。

在 Plan 9 中，工作目录是一个全局变量，多个进程或线程可能同时访问和更改它。为了确保访问和更改的安全性，使用互斥锁来限制对该变量的访问。因此，在 "pwd_plan9.go" 文件中， "wdmu" 用于同步对工作目录的访问。

当任何进程或线程需要更改工作目录时，它必须先获得互斥锁。一旦获得锁，它就可以访问和更改工作目录，然后释放锁以供其他进程或线程使用。这样，就可以避免多个进程或线程同时更改工作目录引起的问题。



### wdSet

在go/src/syscall/pwd_plan9.go文件中，wdSet是一个全局变量，其类型为*int，并且用于存储当前工作目录的文件描述符。该变量在多个函数中使用，如:

1. Getwd()函数使用wdSet变量来检索当前工作目录。

2. Chdir()函数使用wdSet变量来更改当前工作目录。

3. Mkdir()函数使用wdSet变量来确保新创建的目录包含正确的文件描述符。

4. 标准库"os"包中的其他相关函数也使用wdSet变量来管理当前工作目录。

总体来说，wdSet变量是用于跟踪当前工作目录的文件描述符，以便其他函数可以更轻松地处理文件和目录路径。因此，在执行与文件路径相关的操作时，此变量起着非常重要的作用。



### wdStr

在go/src/syscall中pwd_plan9.go文件中，wdStr这个变量是一个字符串，表示当前目录的路径。它的作用是为了在Plan 9操作系统中获取当前工作目录的路径信息。

通常情况下，我们可以使用os.Getwd()函数来获取当前工作目录的路径。但在Plan 9操作系统中，os.Getwd()函数并不适用，因为Plan 9操作系统没有Unix和Windows中的“当前工作目录”概念。相反，Plan 9操作系统中使用的是名为“工作目录”的文件夹来表示当前工作目录。为了正确获取当前工作目录的路径信息，我们需要使用Plan 9操作系统所提供的函数和变量。

在pwd_plan9.go文件中，我们可以看到有一个叫做getwd()的函数，它的作用是获取当前工作目录的路径信息，并将其赋值给wdStr变量。然后，如果我们需要获取当前工作目录的路径信息，就可以直接使用wdStr变量来表示。因此，wdStr变量在Plan 9操作系统中是非常重要的，它能够帮助我们快速准确地获取当前工作目录的路径信息。



## Functions:

### Fixwd

在 Unix 系统中，每个进程都有一个工作目录（working directory），通常表示为当前目录。当进程打开一个文件时，操作系统会先在该进程的当前目录中查找该文件。因此，如果程序需要访问特定路径下的文件，必须先确保该路径的正确性。

在 Plan 9 操作系统中，每个进程有一个 PWDFD 文件描述符，指向其当前目录。Fixwd 函数通过读取该文件描述符并解析路径，将目录路径信息与实际磁盘路径信息进行匹配，从而修复当前目录。Fixwd 函数的具体作用包括：

1. 更新当前进程的本地工作目录
2. 递归解析当前目录和其父目录
3. 处理目录中的符号链接
4. 对通用的 Unix 路径进行转换，以匹配 Plan 9 的路径规范

Fixwd 函数的实现是通过系统调用来进行的，这些调用包括 open、lstat、readlink、chdir 等。由于 Fixwd 函数执行的系统调用涉及到磁盘 I/O，因此其效率会比较低，通常用于初始化阶段或调试目的。



### fixwdLocked

fixwdLocked函数是一个私有函数，用于修复当前工作目录，以确保它始终是绝对路径。在Plan 9操作系统中，当前工作目录可以是相对或绝对路径，但在Go编程语言中，它必须始终是绝对路径。

该函数首先检查当前工作目录是否是绝对路径，如果不是，则将其转换为绝对路径。接着，它检查它是否与之前的保存路径相同，如果不同，则更新保存路径。

fixwdLocked函数的作用是确保当前工作目录始终是绝对路径，并在需要时更新保存路径，以便在使用chdir()操作时可以正确回到之前的目录。



### fixwd

func fixwd(path string) (string, error) {
    path = clean(path)
    if len(path) == 0 {
        return "", ErrBadPath
    }

    // absolute paths are returned as-is
    if isAbs(path) {
        return path, nil
    }

    // relative paths are interpreted relative to cwd
    cwd, err := Getwd()
    if err != nil {
        return "", err
    }
    return Join(cwd, path), nil
}

该函数的作用是修正路径（path），使其变成绝对路径。如果路径是相对路径，它将根据当前目录（cwd）解释相对路径，并返回解释后的绝对路径。

该函数首先会将路径进行规范化（调用clean函数），去掉多余的`/`和`.`等字符，确保路径的正确性。如果路径长度为0，则返回错误“ErrBadPath”。

如果路径是绝对路径（调用isAbs函数判断是否是绝对路径），则直接返回该路径。

如果路径是相对路径，则调用Getwd获取当前目录（cwd），然后用Join将当前目录和相对路径拼接起来，形成绝对路径，并返回。如果获取当前目录失败，则返回错误。



### getwd

getwd函数在Plan 9操作系统中用来获取当前工作目录的路径。由于Plan 9操作系统采用的是分层虚拟文件系统，所以当前工作目录可能会与传统操作系统中的期望值不同。getwd函数通过调用plan9.Syscall来实现，具体过程如下：

1. 调用系统调用“fstat”获取当前进程的状态信息。
2. 使用状态信息中的qid字段获取当前进程所在目录的目录标识符。
3. 调用系统调用“stat”获取当前所在目录的状态信息。
4. 如果当前所在目录是一个软链接，则根据链接地址继续调用“stat”获取实际目录的状态信息。
5. 将获取到的目录路径保存到缓存中，以便下次需要时可以直接获取。

通过getwd函数获取到的目录路径可能还包含一些特殊字符，如“/./”和“/../”，需要进行相应的处理才能得到正确的路径。此外，由于Plan 9操作系统支持动态挂载文件系统，所以getwd函数获取的当前工作目录可能会随着文件系统的挂载点的变化而发生改变。



### Getwd

Getwd是一个系统调用函数，在Plan 9操作系统中用于获取当前工作目录的路径。当一个进程启动后，它会在一个默认的目录中开始运行，而Getwd函数能够返回当前工作目录的路径，以便进程可以更方便地查找和操作文件。

这个函数的作用非常简单，就是返回当前工作目录路径。Getwd函数会调用Plan 9操作系统提供的一些特定命令和函数，来查找当前的工作目录，并返回它所在的路径。它可以帮助程序员更方便地查找、读取和操作文件，尤其在相对路径中使用。

需要注意的是，这个函数是针对Plan 9系统提供的，如果想要在其他操作系统中获取当前工作目录，需要使用其他操作系统提供的相应函数。例如在Linux和Unix中，可以使用getcwd函数，而在Windows中，可以使用GetCurrentDirectory函数。



### Chdir

Chdir是一个函数，用于改变当前工作目录。具体来说，它会将当前工作目录设置为指定的目录。在Unix系统中，工作目录是指操作系统当前对于特定进程而言的当前目录，该进程随后对文件操作（如读取和写入文件）都将基于此目录。

在pwd_plan9.go文件中，Chdir函数是为Plan 9操作系统（一种类Unix操作系统）实现的。它使用系统调用chdir来更改当前工作目录。在Go中，syscall软件包提供了对操作系统底层接口的访问，因此Chdir函数使用了syscall包中的相关函数。

Chdir函数的作用在于让程序员能够方便地更改当前工作目录，以便于对于文件系统中某个特定文件或目录的访问。它是Unix系统中常用的标准函数之一，也在Plan 9系统中广泛使用。

在使用Chdir函数时，需要传入目标目录的路径。例如，调用Chdir("/home/user/Documents")将更改当前工作目录为/user/Documents。如果操作成功，则返回nil，否则返回错误信息。


