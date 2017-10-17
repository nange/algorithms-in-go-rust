# DataStructures-And-Algorithms-With-Four-Language

四种语言（Go，Javascript，Java，C）实现常用数据结构和算法。

所有代码都经过具体测试运行，具体环境为：
- 操作系统：ubuntu16.04 64位
- Go语言版本： Go 1.8.3
- Javascript执行环境： Nodejs 6.11.3
- Java版本： Java 1.8.0_25
- C语言编译器： gcc 4.8.2


## 目录
1. [列表](#user-content-列表) (完成)
2. [栈](#user-content-栈) (完成)
3. [队列](#user-content-队列) (完成)
4. [链表](#user-content-链表) (golang,javascript 完成...)
5. [字典](#user-content-字典) (golang 完成)
6. [二叉树和二叉查找树](#user-content-二叉树和二叉查找树) (golang 完成)
7. [图和图算法](#user-content-图和图算法) (待处理...)
8. [排序算法](#user-content-排序算法) (待处理...)
9. [检索算法](#user-content-检索算法) (待处理...)


### 列表
列表是非常常见的一种数据结构，比如日常所见的购物清单、待办事项等等。
它提供了对列表数据的一系列操作，比如：添加、删除、修改、遍历等功能。
当我们把这样的具体问题抽象成用列表去解决的时候，往往可以简化问题。

具体实现： [Go](go/arraylist) 、 [Javascript](javascript/arraylist) 、
[Java](java/arraylist) 、 [C](c/arraylist)


### 栈
栈也是一种非常常见的数据结构,在计算机的世界里,在计算机的世界里,
栈是一种很高效的数据结构,因为数据只能在栈顶添加或者删除.
因此栈也被称为一种后入先出的数据结构.栈的使用遍布程序语言实现的方方面面,
从表达式求值到函数调用.

具体实现： [Go](go/stack) 、 [Javascript](javascript/stack) 、
[Java](java/stack) 、 [C](c/stack)


### 队列
队列是一种前进先出的数据结构. 在日常生活中非常常见:比如去银行排队办理业务.
在计算机中也极其常见, 很多情况下,当有大量任务需要完成时, 就会把任务暂时加入到
任务队列中, 执行一个删除一个,继续执行下一个任务.

具体实现: [Go](go/queue) 、[Javascript](javascript/queue) 、
[Java](java/queue) 、 [C](c/queue)


### 链表
有时候数组不一定是最佳的组织数据的数据结构，因为数组通常都是固定大小的，当数据填满时，
再加入新元素就变得很困难。在数组中，添加和删除元素也很麻烦，因为要移动数组中的其他元素。
因此如果需要频繁的添加或者删除元素，可以考虑使用链表组织数据。

具体实现: [Go](go/linkedlist) 、[Javascript](javascript/linkedlist) 、
[Java](java/linkedlist) 、 [C](c/linkedlist)


### 字典
字典类型(map)，是一种KV结构，通过key能够找到对应的value，map类型的特点是查询非常快(O(1))，
原理是通过一个hash函数能快速的找到对应的key的位置，在实际编程中，map类型几乎无处不在。

具体实现: [Go](go/hashmap) 、[Javascript](javascript/hashmap) 、
[Java](java/hashmap) 、 [C](c/hashmap)


### 二叉树和二叉查找(搜索)树
二叉树可以看作是链表的二维扩展，链表的一个节点只有一个next指针，但是二叉树的一个节点可以有两个
next指针(通常被称作left,right)，二维扩展后可以解决一维链表的一些问题，比如查找慢等。 
二叉查找树可以把查找，添加，删除元素的时间复杂度降到O(logN)。 

具体实现: [Go](go/binarytree) 、[Javascript](javascript/binarytree) 、
[Java](java/binarytree) 、 [C](c/binarytree)
