# DataStructures-And-Algorithms-With-Four-Language

四种语言（Go，Javascript，Java，C）实现常用数据结构和算法。

所有代码都经过具体测试运行，具体环境为：
- 操作系统：ubuntu14.04 64位
- Go语言版本： Go 1.4
- Javascript执行环境： Nodejs 0.11.14

- Java版本： Java 1.8.0_25
- C语言编译器： gcc 4.8.2


## 目录
1. [列表](#user-content-列表) (完成)
2. [栈](#user-content-栈) (完成)
3. [队列](#user-content-队列) (进行中...)
4. [链表](#user-content-链表) (待处理...)
5. [字典](#user-content-字典) (待处理...)
6. [散列](#user-content-散列) (待处理...)
7. [集合](#user-content-集合) (待处理...) (待处理...)
8. [二叉树和二叉查找树](#user-content-二叉树和二叉查找树) (待处理...)
9. [图和图算法](#user-content-图和图算法) (待处理...)
10. [排序算法](#user-content-排序算法) (待处理...)
11. [检索算法](#user-content-检索算法) (待处理...)


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
