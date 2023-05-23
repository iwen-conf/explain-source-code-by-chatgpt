# File: mranges.go

mranges.go这个文件包含了Go语言运行时的内存分配器中与memory range有关的核心代码。它的作用是管理不同内存范围的分配和释放，为程序提供高效的内存分配与回收功能。

在Go语言的内存分配器中，mranges.go主要实现了两个功能：

1. 规划和分配内存范围（Memory Range）
内存范围是Go内存分配器管理的内存区域，一个memory range可以包含多个arena，一个arena又包含多个span，多个span可以被用来存储某种特定大小的对象，arena和span在内存中的划分是由mranges.go实现的。

2. 管理内存范围（Memory Range）
内存范围的管理包括扩展和收缩内存范围的大小和位置、初始化和清理内存范围等。

在函数级别上，mranges.go中实现了以下几个函数：

1. newMemoryRange() - 用于申请一个新的memory range
2. prepareForSweep(m*pageAlloc, gv*growableBitmap) - 将所有能够清除的span的对象都清除掉
3. isSweepDone(uint32) - 判断清除工作是否完成
4. makeBudget(slice<mrangeBudget>) - 分配并初始化内存范围的预算
5. scavengeContext.gcScanRange() - 用于在垃圾回收之前扫描内存范围
6. scavengeContextClass - 定义申请内存和垃圾回收的上下文类别、对象大小和堆大小等

综上所述，mranges.go是Go语言运行时的内存分配器核心代码之一，它实现了内存范围的管理和分配、垃圾回收等功能，为程序提供高效的内存分配和回收服务。




---

### Var:

### minOffAddr

minOffAddr是runtime中的一个全局变量，它用于描述堆中的mcache的最小地址偏移量。在go语言中，每个goroutine都有自己的mcache，它用于分配小对象的内存。mcache由一系列的mcentral和mspan组成，每个mcentral对应一种大小的对象，而每个mspan则控制这些对象的分配和回收。

当进行内存分配时，堆会为当前的goroutine分配一个mcache，这个mcache中包含的mspan的地址是连续的，所以它们的地址偏移量必须是递增的。minOffAddr保存的就是当前mcache中地址偏移量的最小值。它会在mcache被创建时初始化为一个大正整数，表示此时还没有mspan被分配给这个mcache。每次分配mspan后，就会将minOffAddr更新为该mspan的地址偏移量，以便下一次分配可以继续在这个偏移量的基础上进行。

minOffAddr的作用是确保每个mcache中mspan的地址偏移量是连续的，从而使得内存分配和回收更加高效。同时，它也被用于判断当前mcache中是否还有足够的空间可以分配新的mspan。如果minOffAddr已经超出了mcache中所有可用地址，则说明当前的mcache已经无法再分配新的mspan，需要重新分配一个新的mcache。



### maxOffAddr

maxOffAddr是一个全局变量，它存储了最大的偏移地址，在应用程序的虚拟地址空间中，这相当于最大的堆地址。在Go语言运行时的内存管理中，使用这个变量来帮助确定堆的界限。这个界限是堆的最大範圍，超出这个界限的内存会被系统认为是非法的，如果程序在这个区域中尝试分配内存，就会导致一个崩溃。 

具体来说，在分配一个新的Arena时，Go语言运行时会计算出 Arena 的地址范围，然后将其与 maxOffAddr 进行比较。如果 Arena 的地址范围超出了maxOffAddr，那么这个 Arena 就不能用于分配内存，因为它将与堆的后面的系统分配区重叠。因此，maxOffAddr 可以看作是堆的最大大小的一个标记，Go语言运行时可以利用它来防止堆的大小越界。

总体来说，在Go语言运行时内存管理中，maxOffAddr 作为一个堆的界限标记，用于确保程序不会意外地分配超过可用内存，从而保证了系统的稳定性和安全性。






---

### Structs:

### addrRange

在 Go 语言的垃圾回收器中，addrRange 结构体用于表示一段内存地址范围。它包含了该内存范围的起始地址和结束地址，以及该内存范围的状态信息。

addrRange 由 mranges.go 中的 rangeTab 数组存储。在垃圾回收器扫描整个进程内存时，会根据 rangeTab 数组中的 addrRange 信息来遍历整个内存空间。垃圾回收器会根据 addrRange 的状态信息来判断该内存范围中的对象是否需要进行回收。如果内存范围中的对象已经没有被引用，则垃圾回收器会将其标记为可回收的。

除了用于垃圾回收器的扫描操作之外，addrRange 还可以用于调试和分析内存问题。程序员可以通过 addrRange 来查找内存中的数据是否被正确地分配和释放，以及分配和释放的顺序是否正确。

总之，addrRange 结构体提供了一种可靠和高效的方式来管理和操作内存地址范围，是实现垃圾回收等内存管理操作的重要工具。



### offAddr

在Go语言的runtime中，mranges.go这个文件中的offAddr结构体是用于表示堆内存分配器中的空闲内存块的起始位置和结束位置的结构体。它由两个字段组成：off表示空闲块在堆内存中的偏移地址，addr表示该空闲块的起始地址。

offAddr结构体一般被存储在mheap结构体的free结构体中，free结构体是用于存储空闲内存块链表的结构体。每个free结构体都有一个链表，该链表由多个offAddr结构体组成，每个offAddr结构体表示一个空闲内存块。

当程序需要分配一块内存时，内存分配器会在free链表中查找一个足够大的空闲内存块，并将其从链表中删除。如果内存分配器需要回收一块内存，则会将其转换成一个offAddr结构体，并将其插入到free链表中，以便被后续的内存分配所使用。

因此，offAddr结构体在内存分配器中起着关键的作用，它们是分配和回收内存的基本单位。同时，由于内存分配器需要频繁地创建和删除offAddr结构体，因此该结构体的实现必须非常高效，以保证整个内存分配过程的性能。



### atomicOffAddr

在go/src/runtime中，mranges.go文件中的atomicOffAddr结构体是用于维护映射的一些元数据的。

具体来说，atomicOffAddr结构体中维护了一个地址指针，该指针指向一个uint32类型的值。这个值的含义是，映射的键使用另一个结构体offAddr作为索引时，该键所对应的值在元素存储切片中的偏移量。

映射中的键值对信息存储在一个切片中。如果需要查找某个键的值，就需要知道该键在切片中的偏移量。而该偏移量可以通过offAddr结构体计算得到。因此，使用atomicOffAddr结构体记录每个键的偏移量，可以实现高效的查找。

在进行添加、删除、查找操作时，可以使用原子操作保证并发安全。

总之，atomicOffAddr结构体在实现映射时扮演着非常重要的角色，可以有效维护切片中元素的偏移量信息，提升查找效率和并发安全性。



### addrRanges

addrRanges结构体定义在mranges.go中，它表示一组内存地址范围。addrRanges被用于记录可用的虚拟内存地址范围，用于在堆上分配新的对象。

具体来说，addrRanges结构体包含以下字段：

- start：结构体表示的地址范围的开始地址；
- end：结构体表示的地址范围的结束地址；
- next: 指向下一个 addrRanges 结构体对象的指针。

其中，start和end字段被用于描述虚拟内存范围的起始地址和终止地址。next指针被用于构造一个地址范围链表，在堆上分配新对象时使用。

在runtime中，addrRanges被用于记录当前可用的虚拟内存地址范围，并在堆上分配新对象时从中选择一个地址范围。当一个地址范围被用于分配新对象后，它将从链表中移除或者被分裂成两个地址范围，并且可能会变得不连续或非常小。

总之，addrRanges结构体是runtime用于管理虚拟内存的一种数据结构，它可以记录可用的虚拟内存地址范围，并在需要时选择一个地址范围用于堆上的新对象分配。



## Functions:

### makeAddrRange

makeAddrRange函数负责将一段内存地址范围加入到运行时的内部映射中，即将其转换为可用的mheap可分配地址范围。该函数主要被以下场景调用：

1. 分配内存空间时（例如mheap.alloc函数），需要获取可分配的内存地址范围。

2. 内存回收时（例如mheap.cacheSpan函数），需要将被回收的内存地址范围返回给运行时，以便重新分配。

该函数的实现步骤如下：

1. 获取当前操作系统页的大小。

2. 计算内存地址范围的起始位置和长度，将起始位置和长度对齐到操作系统页的大小的倍数。

3. 将起始位置和长度转换为连续的span对象，并将其加入到运行时的heap中。

4. 如果该span的状态为MSpanFree（空闲），则将其挂入对应的heap空闲链表中。

5. 如果该span的状态为MSpanInUse（已分配），则将其挂入对应的heap分配链表中。

6. 更新heap的相关属性，例如heap.alloc中记录已经分配的空间大小。

通过makeAddrRange函数，运行时可以动态管理内存地址范围，并保证内存分配和回收的正确性和性能。



### size

在Go语言中，mranges.go文件中的size函数是用来计算内存块的大小的。这个函数接收一个类型为*mspan的参数，以及一个uintptr类型的addr参数，返回一个uintptr类型的值。

具体来说，mranges.go中的size函数会根据mspan的内存布局，计算addr指向的内存块在这个mspan中占据的字节数。具体的计算规则是：

1. 如果addr指向的内存块是一个span的头部，返回span的大小。

2. 如果addr指向的内存块是一个对象，返回对象所在的span的大小。

3. 否则，返回0。

这个函数的作用很重要，因为在Go语言中，内存分配和垃圾回收都会涉及到内存块的大小计算。例如，当我们需要调用malloc函数申请内存时，就需要计算需要申请的内存块的大小，并为其分配足够的内存空间。而在垃圾回收过程中，也需要计算对象所占的内存块的大小，以便于确定是否需要回收这个对象。

总之，size函数在Go语言的内存管理中扮演着重要的角色，通过这个函数的计算，我们可以更加高效、准确地管理内存。



### contains

在Go语言中，垃圾回收器（GC）是由runtime包提供的。其中，mranges.go文件包含了一些关于内存区域的操作函数，包括contains函数。

contains函数的作用是判断一个地址是否在一个内存区间之内。它接收两个参数：base和size。其中，base是内存区间的起始地址，size是内存区间的长度。如果给定的地址在该内存区间内，返回true；否则返回false。

这个函数在GC过程中非常重要，因为GC需要遍历程序的堆内存，寻找不再使用的对象以进行回收。当GC遍历一个内存区间时，需要判断当前遍历到的对象是否在该内存区间内。如果contains函数返回true，说明该对象在内存区间内，可以进行回收。

总之，contains函数是垃圾回收器中非常重要的一个函数，它用于判断一个地址是否在一个内存区间内，帮助GC进行内存回收。



### subtract

`subtract`函数是用来从`xspan`中减去`yspan`的函数。这两个参数都是`mspan`类型，代表的是一段连续的内存区域。`subtract`函数会将`yspan`切割成多个部分，将这些部分从`xspan`中删去。具体来说，`subtract`函数会检查每个`yspan`中的连续页是否在`xspan`中，如果是，则将其从`xspan`中删除。如果`yspan`包含强制隔离的页，则将删除该页并将其添加到被强制隔离的页列表中。最终，`subtract`会返回`xspan`，该`xspan`已经减去了`yspan`。 

在 Go 的运行时中，`mranges.go`这个文件实现了用于跨度（`mspan`）管理的许多内存管理算法。内存管理是 Go 程序中非常重要的组件之一，因为它影响了程序的运行效率和可靠性。`subtract`函数的作用是减去一段内存区域，用于管理跨度中的内存分配和释放。该函数的实现是为了确保跨度中的内存使用率最大化，以便提高程序的性能和效率。



### takeFromFront

takeFromFront函数的作用是从链表的头部中取出一段连续的可用内存块。这个函数被用于管理内存分配器中的空闲区域列表，以便在运行时动态地为程序分配内存块。

该函数的实现过程如下：

首先，通过调用atomic.Loaduintptr函数获取当前链表头的信息，并将其存储在p变量中。

接下来，通过p & ~uintptr(blockSize-1)的计算方法，将p和blockSize进行对齐，以确保取出的内存块大小为blockSize的整数倍。

然后，通过atomic.Casuintptr函数将链表的头指针向后移动blockSize个字节，并确保已完成此操作。

最后，返回p作为取出的内存块的首地址。

总体来说，takeFromFront函数的作用是帮助内存分配器管理可用的内存，确保程序始终可以分配到连续的内存块。



### takeFromBack

takeFromBack函数的作用是从mspan的后面分配连续的一页（或更多）给mcache。

在Golang中，内存空间是以固定大小的块（一页）进行分配。每个mspan代表了一系列连续的、大小固定的内存页面。mcache则是mcentral的私有缓存，包含了一些已经被分配的对象，用于在没有足够内存分配时尝试减少一次内存分配的时间。

在takeFromBack函数中，它会先从mspan的尾部开始检查，如果mspan的末页还没有被翻转，则翻转该页，并返回该页面的地址。如果已经翻转，则检查前一页是否已经被翻转。如果没有，则翻转前一页，并返回该页面的地址。如果所有页面都已经被翻转，则返回nil。

这个函数的目的是为了高效地分配连续内存页给mcache，避免频繁地分配单页内存导致内存碎片化，从而影响程序的性能。



### removeGreaterEqual

在go语言中，运行时系统（runtime）的mranges.go文件中的removeGreaterEqual()函数用于在traceback记录中删除所有大于或等于指定地址的范围。在Go的堆栈跟踪过程中，每个运行时Goroutine的状态都被放入一个数据结构中，称为traceback记录。这个结构包括调用堆栈、当前程序计数器（PC）和Goroutine的状态信息等。此外，堆栈跟踪记录还包括一系列范围信息，这些范围代表一组连续的内存区域。

removeGreaterEqual()函数的作用是删除堆栈跟踪记录中所有大于或等于给定地址的范围。这个函数被用于从堆栈跟踪记录中删除已释放的内存区域信息，以避免出现错误的内存访问。例如，在goroutine完成时，它将无法再次访问堆栈跟踪记录，因此所有范围信息都应该被清除。

具体实现方案是，在traceback记录中，会按照内存地址的大小顺序进行排序，然后使用二分查找来查找第一个大于或等于给定地址的范围。接下来，它将删除剩余的所有范围信息，并将记录截断以匹配最后一个范围的结束地址。

因此，removeGreaterEqual()函数的作用是确保在堆栈跟踪中只包含尚未释放的内存范围，以避免出现错误或不正确的跟踪信息。



### add

在Go语言的运行时中，mranges.go文件中的add函数用于将新的内存范围添加到一系列已有的内存范围中。这些内存范围都是用于管理堆内存的。每个堆内存范围都包含一组连续的内存地址，并且在执行垃圾回收等操作时需要对其进行管理。

add函数首先会检查是否存在与新范围重叠的现有范围，如果存在，它会将它们合并成单个范围。如果不存在重叠，则将新范围添加到已有范围列表中。add函数还会更新堆内存使用情况的统计信息。

具体来说，add函数会引用heap、mheap和mspan结构体中的一些字段。堆是一个存储所有分配的内存的区域。在堆中，每个分配的内存都被存储在连续的区域中。mheap存储有关堆的元数据，例如与可用空间相关的信息。mspan结构描述了heap中的一块内存。add函数还涉及mspan和mcache结构体中的一些字段，mcache是每个goroutine私有的一组可重用对象的缓存，包括用于分配和释放内存的对象。

总的来说，add函数是一种重要的内存管理功能，它确保了所有的堆内存范围都被管理，并且能够有效地进行垃圾回收等操作。



### sub

sub函数是用来处理两个内存块之间的相对偏移量的。在runtime中，每个分配的对象都被分配在一段连续的内存块中，这个内存块被称为mheap。当需要释放一些内存时， runtime会将释放的内存块与之前分配的内存块进行合并，使得垃圾回收时可以更有效地回收内存。

sub函数用于计算两个内存块之间的相对偏移量。这个偏移量将用于对指向内存块的指针进行纠正，以确保垃圾回收不会将不在内存块中的指针误认为是指向该内存块中的对象。

sub函数的具体实现包括对传入参数进行一些简单的检查，以确保传入的指针确实指向mheap内存块中的对象。然后，它根据传入的指针计算出相对偏移量，并返回这个偏移量作为结果。

总之，sub函数是runtime中非常重要的一个函数，它确保了垃圾回收可以更有效地回收不再使用的内存，从而提高了程序的整体性能。



### diff

在 Go 语言中，mranges.go 是用于管理内存分配的文件。其中，diff() 函数用于计算两个内存区域之间的差异。

具体来说，diff() 函数接受两个指针参数和一个长度参数。它会比较这两个指针所指向的内存区域，找出它们之间的差异并返回一个指向该差异区域的指针。

diff() 函数的实现利用了 Go 语言中 unsafe 包提供的指针运算功能，以及内存区域中元素的大小和对齐方式等信息。它在内存分配和垃圾回收等相关过程中发挥重要作用，用于检测内存区域之间的重叠和相互影响等情况，以确保内存的正确分配和使用。

总之，diff() 函数是 Go 语言中用于计算内存区域差异的重要工具函数，对于内存管理和调试等方面具有重要意义。



### lessThan

函数lessThan主要用于比较两个地址是否小于。该函数是在runtime中mranges.go文件中使用的一个辅助函数。该函数可以接受两个参数，分别为x和y，它会比较地址x和y的大小，如果x的地址小于y的地址，则返回true，否则返回false。

在mranges.go文件中，lessThan主要用于排序ContiguousRange数组。ContiguousRange在runtime中表示一个连续的内存区域。这个数组用于保存Go程序中的栈内存和堆内存段，因此对数组中的元素进行排序是必须的。在排序中，通过调用lessThan函数来比较两个内存段的位置，然后根据比较结果进行排序。

另外，从其函数名lessThan可以看出，该函数被设计用于在内存分配时判断两个地址的大小，以确定内存申请时所需内存块的大小。因为在操作系统中，新分配的内存一般都是从高地址向低地址分配的，因此lessThan函数的实现需要注意这一点，避免出现内存覆盖等问题。

在总体上，lessThan函数的作用是为了确保内存段在ContiguousRange数组中的顺序正确性，并测量内存应分配到哪里。



### lessEqual

func lessEqual(x, y *mspan) bool

该函数在runtime的mranges.go文件中定义，它被用于对mspan结构体进行排序，用于选择合适的mspan去满足用户的内存申请，其中x和y是两个要比较的mspan指针。

函数的返回值为布尔值，如果x的尺寸小于或者等于y的尺寸，则返回true；否则返回false。

这个函数的作用是为了帮助runtime将内存池中空闲的mspan排列成一个链表，使得它们可以按照尺寸从小到大排列，以供用户的内存申请使用。



### equal

equal函数是用于比较两个range对象是否相等的函数。一个range对象代表了一段连续的内存空间，包含了起始地址、大小、以及对应的span（span是runtime中管理内存的基本单位，大小为2的幂次方，通常为8KB到512KB之间）。当一个goroutine申请一段内存时，runtime会寻找一个足够大的span，并将其分割成适当大小的range。在释放内存时，相邻的range可能会合并成一个更大的range，这个时候就需要使用equal函数来判断两个range是否相等。

该函数的实现通过比较两个range的起始地址、大小以及span指针来判断它们是否相等。如果两个range对象的属性都相同，则认为它们是相等的。如果两个range的span是不同的，但是它们的起始地址和大小都相同，则认为它们也是相等的。这种情况可以出现在range合并的过程中。

equal函数主要在mranges.go和mspan.go两个文件中使用，用于管理和操作内存分配的数据结构。



### addr

在Go语言中，内存分配是由runtime包负责的，mranges.go这个文件是runtime包中的一个文件，其中的addr函数主要是用于初始化和管理小对象的内存区域。

具体来说，addr函数的作用如下：

1. 初始化小对象的内存区域

在程序启动时，addr函数会通过调用sysAlloc函数从操作系统中申请一段内存空间，然后将这段内存空间转换为mheap结构体，并将其保存到runtime包的全局变量mheap中。该内存空间会被划分为多个大小相等的小对象区域，以便进行内存分配。

2. 管理小对象的内存空间

在程序运行期间，addr函数还会负责管理小对象的内存空间。具体来说，它会调用releaseSpans函数来管理对象所在的span（即内存块）。如果一个span中的所有对象都已被释放，addr函数会将该span插入到空闲span链表中，以便下次分配使用。

此外，addr函数还会处理小对象的对齐问题。由于小对象的大小不一定是系统字对齐（通常为8字节）的整数倍，因此需要对小对象进行对齐操作。addr函数会通过调用alloc_m和spanalloc_m函数来进行内存分配，并保证分配出来的对象的地址是系统字对齐的。

总之，addr函数是Go语言中内存分配的重要组成部分，通过初始化和管理小对象的内存空间，它提高了内存分配的效率。



### Clear

Clear函数是在运行时内存管理中用于清除所有的内存范围（mrange）的函数。内存范围是对一段内存区域的抽象，每个mheap可包含多个内存范围。

Clear函数主要有以下作用：

1. 释放所有内存范围占用的虚拟内存地址，这些地址将被还回给系统，以便让其他程序使用。

2. 将所有内存范围的状态重新设置为未分配，即可以重新使用。

3. 清除所有内存范围对应的链表，以便重新分配内存时从头开始查找可用内存范围。

4. 清除所有内存范围对应的统计信息，以便重新统计内存的使用情况。

5. 清除所有内存范围对应的调试信息，以便重新诊断内存管理的问题。

Clear函数通常在操作系统重新分配内存或运行时内存管理器初始化时使用。在操作系统重新分配内存时，系统需要清除之前的内存范围并释放虚拟内存地址；在初始化时，由于所有内存范围已经不再使用，所以需要将它们清空，以便重新分配内存。



### StoreMin

StoreMin是一个函数，在Go语言的运行时库runtime中的mranges.go文件中定义。它的作用是为给定的mspan（在堆上分配内存的区域）分配最小的大小，以便确保任何未使用的内存都可以被最小的大小所占据。这样可以保证使用内存时的高效性。

具体来说，该函数的实现方式如下：

- 首先它会检查给定的mspan是否被标记为“正在使用中”或“已释放”，如果是则跳过。
- 然后它会计算可以容得下的最小大小（minSize），这个大小取决于mspan的大小和它是否被标记为缓存mspan（cache mspan）。对于非缓存mspan，minSize为mspan的大小减去minPhysPageSize（最小物理页面大小），以保证还可以容纳一个最小物理页面大小的分配。对于缓存mspan，minSize直接被设置为minPhysPageSize（因为它已经足够小）。
- 接下来，StoreMin会使用minSize来调整mspan的大小，使其至少包含minSize大小的内存。
- 最后，它将更新mspan的大小和可用字节数。如果mspan是缓存的，则将它放入缓存列表中。

通过这些步骤，StoreMin保证了每个mspan中都至少有minSize大小的内存，从而提高了内存的利用效率。



### StoreUnmark

StoreUnmark是一个函数，位于Go语言运行时包中runtime/mranges.go文件中，主要作用是将GC标记位从Span中清除。具体来说，当GC需要将某个对象标记为不可回收时，它会设置一个标记位。然后这个标记位被储存在Span的状态字中。

然而，当对象的垃圾回收状态改变时，GC需要将标记清除，在这个过程中，StoreUnmark就派上用场了。在运行过程中，当Span的状态从被标记为可回收变为不可回收时，StoreUnmark会被调用，将Span中的垃圾回收标记位清除。

这个函数的作用是非常重要的，因为它确保了GC所需的内存和系统资源得以充分利用。换言之，它确保了运行时系统内部状态的一致性，并使垃圾回收过程更加高效。

总之，StoreUnmark函数是Go语言运行时系统中的一个重要函数，它的作用是清空Span的垃圾回收标记位，以确保系统状态的一致性和垃圾回收的高效性。



### StoreMarked

StoreMarked函数是Go语言运行时系统中用于将被标记的堆对象存储到标记循环标记的活动页中的函数。它的作用是在垃圾回收期间，将被标记的堆对象放入活动页，以便在下一轮循环中进行标记和处理。

具体而言，StoreMarked函数接受一个参数mpage，表示当前正在处理的页。该函数遍历该页中的所有对象，并将被标记的对象添加到mpage的对象列表中。在下一次循环期间，处理器将处理这个对象列表，并对其中的对象进行标记和处理。

需要注意的是，StoreMarked函数是在标记循环标记的过程中被调用的，而标记循环标记是一种并发垃圾回收算法，它通过将标记阶段和扫描阶段分别并行运行，以减少垃圾回收造成的停顿时间。因此，在StoreMarked函数中，需要考虑并发访问和修改对象的情况，从而实现线程安全的对象存储。

总之，StoreMarked函数是Go语言运行时系统中用于将被标记的堆对象存储到标记循环标记的活动页中的函数，它是实现并发垃圾回收的关键之一。



### Load

mranges.go文件中的Load函数主要用于在程序启动时，将堆上的内存块（mheap）分配给对应的M或G，以支持后续的并发运行。具体来说，该函数会遍历整个内存块链表（heap arena），为每个内存块分配所需的M或G，并将它们与对应的内存块进行绑定。这样，当后续有goroutine需要执行时，就可以直接从对应的M或G获取已经分配好的内存块，节省了内存申请和释放的开销，提高了程序的性能。

Load函数的具体实现细节比较复杂，它会涉及到对内存块的地址计算、堆栈指针的设置、内存块的绑定、内存页数的计算等操作。但总体来说，它的作用就是为每个内存块找到一个合适的M或G，并将它们绑定在一起，以支持后续的goroutine并发执行。这也是Go语言可以高效支持大规模多线程并发的关键之一。



### init

在go/src/runtime中，mranges.go文件中的init func的作用是初始化全局的mheap和mspan缓存。mheap用于管理heap中的内存，mspan用于管理每个span（一块连续的内存）的元信息。

具体来说，init func首先会初始化mcentral队列，每个队列代表特定大小对象的分配器。接着会初始化mheap中的free列表，这个列表用于储存能够重新分配的span。然后，init func还会初始化mspan缓存，为每个span分配对应的mspan结构体，并将其标记为未被使用状态。最后，init func还会设置背景mcache分配器所需的初始化参数。

总之，init func在运行时初始化了全局的mheap和mspan缓存，为后续的内存分配和管理提供了必要的基础。



### findSucc

findSucc是在runtime中mranges.go文件中定义的函数，它的作用是在一组指向区间的指针范围中查找指定的位置，返回第一个指向区间的指针。该区间可以用存储在mranges数组中的start、end指针列表来表示，这些指针指向一段内存区域的起始和结束位置。

具体地说，findSucc函数采用二分搜索来查找指定的位置，以便快速找到第一个指向区间的指针。它使用比较函数来比较指向区间的指针和查找位置，以便找到第一个大于或等于查找位置的指针。该函数的时间复杂度为O(log n)。

该函数在runtime中的一些核心模块中使用，如在垃圾回收器中，它可以使用它来查找需要扫描的内存区域。这些区域可能会被垃圾回收器扫描，以找到不再被需要的对象。因此，findSucc函数是一个非常重要的函数，在runtime中有广泛的应用。



### findAddrGreaterEqual

在 Go 语言的运行时(runtime)系统中，每个 Goroutine（Go 协程）都有一个在堆空间中分配（alloc）内存的“本地缓存”（local cache），用于提高内存分配和使用的效率。这个本地缓存就是由 mranges.go 文件中的 findAddrGreaterEqual 函数来管理的。

findAddrGreaterEqual 函数的作用是在本地缓存中查找一个合适的内存块，以便分配新的内存。它首先尝试在本地缓存中的链表中哈希查找。如果找到了一个空闲内存块，则将其从链表中移除并返回该内存块的地址；否则，它会尝试从堆分配更多的内存，并将其添加到本地缓存链表中。

findAddrGreaterEqual 函数的另一个作用是保证本地缓存中的内存块是连续的。因为本地缓存中的内存块是在堆上分配的，所以它们可能不是连续的。如果本地缓存中的内存块被分散在堆上，那么在分配大块内存时，需要从不同的局部缓存中进行分配，这会导致性能下降。因此，在分配内存时，findAddrGreaterEqual 函数会对本地缓存中的内存块进行合并，以确保它们是连续的。

总之，findAddrGreaterEqual 函数是 Go 语言运行时系统中的关键组件之一，它管理着 Goroutine 的本地缓存，并且在内存分配和释放时起到了至关重要的作用。



### contains

在 Go 语言中，mranges.go 文件中的 contains 函数是一个用于检查一个内存块是否包含另一个内存块的函数。该函数通常用于内存分配和释放。

在内存分配时，需要检查当前的可用内存块是否足够容纳所需的内存。如果当前的可用内存块足够，那么剩余的部分就可以继续使用。因此，contains 函数需要检查当前的可用内存块是否包含了所需的内存块，以便确定可用内存块是否足够容纳所需的内存。

另外，在内存释放时，需要将被释放的内存块标记为空闲状态，以便后续的内存分配能够重新使用它。但是，在标记为空闲之前，需要先检查该内存块是否属于当前的可用内存块。如果该内存块不属于当前的可用内存块，那么就可能会出现内存错误和崩溃。因此，contains 函数也被用于内存释放以确保释放的内存块属于当前的可用内存块。

总的来说，contains 函数在内存分配和释放中扮演了重要的角色，它保证了内存的正确使用和减少了内存错误的可能性。



### add

在go语言的运行时(runtime)中，mranges.go文件中的add函数的作用是向mheap结构体中增加一个新的内存片段。具体的实现过程如下：

1. 首先判断是否有足够的空间容纳所请求的内存片段，如果没有则调用mheap中的grow函数进行扩容。

2. 调用mheap中的alloc函数分配新的内存片段，并将其添加到mheap中。

3. 更新mheap中的相关信息，如mheap_.free和mheap_.busy等参数，以及推进mstats中的相关信息。

4. 返回新分配的内存片段的地址和大小。

add函数的主要作用是管理mheap中的内存片段，保证内存的分配和回收机制正常运作，从而维护整个go程序的内存管理。在并发情况下，add函数也能够确保线程安全并处理竞争条件，保证程序的正确性和稳定性。



### removeLast

removeLast这个func的作用是从某个内存range的列表中移除最后一个内存range。在runtime中，内存range是由一段内存地址区域和对应的元数据（如大小、类型等）组成的。

具体而言，removeLast的实现是对mheap中的mspan的操作。mspan是一个跨度(span)的缩写，其表示了一段连续的内存区间。在某些情况下，比如垃圾回收过程中，需要将某个mspan中的内存回收掉，同时将这个mspan从对应内存range的列表中移除。这就需要用到removeLast函数。

removeLast函数首先获取对应内存range的列表中的最后一个mspan，然后会对其进行一些预处理（如清零状态标志），再将其从列表中移除。最后，函数将移除的mspan返回，以供之后的操作使用。

需要注意的是，removeLast函数只能在单线程情况下被调用，因为它会直接修改mheap中的全局变量。如果在多线程情况下调用removeLast，可能会导致数据竞争和不可预测的结果。



### removeGreaterEqual

这个函数的作用是从一组范围中删除所有包含指定地址或高于指定地址的范围。在runtime中，Goroutine、堆和栈等对象都由范围表示，这些范围最初被分配并加入到范围列表中，当它们释放时，它们被从列表中删除。removeGreaterEqual函数为范围列表提供了删除包含指定地址的范围的功能，以及删除高于指定地址的范围的功能。

该函数的实现使用了一些辅助函数，如findrange、splitrange、insert_range、newobject和mergestack，这些函数都有其特定的功能，可以协同完成该函数的任务。例如，findrange函数用于查找包含指定地址的范围，如果找到了该范围，则返回该范围的指针；否则，返回nil。splitrange函数用于将范围分成两个，其中一个包含指定地址，另一个包含高于指定地址的部分，并返回一个指向新范围列表的指针。此外，该函数还确保范围列表中的所有范围按地址顺序排列。

总之，removeGreaterEqual函数是runtime中对范围列表进行操作的一个关键函数，它为底层的内存管理提供了重要的支持。



### cloneInto

在Go语言中，内存管理是由runtime包来实现的。mranges.go文件中的cloneInto函数是用来复制一个mheap区间到另一个mheap区间的函数。

mheap是Go语言中的一个堆结构，用于管理内存。cloneInto函数的功能是将一个mheap区间中的所有内存块复制到一个新的mheap区间中。这个函数通常用于在不同的堆区间之间进行内存迁移或者合并。

具体来说，cloneInto函数首先会计算需要复制的内存的总大小，然后通过分配内存来创建一个新的heap对象。接下来，将原heap对象中的每个内存块都遍历一遍，并将其逐个拷贝到新的heap对象中。最后，将原heap对象的meta数据和指针数据也复制到新的heap对象中。这样就完成了一个heap对象到另一个heap对象的复制过程。

总之，cloneInto函数是Go语言内存管理系统中的一个重要组成部分，用于完成heap区间的内存迁移或合并。


