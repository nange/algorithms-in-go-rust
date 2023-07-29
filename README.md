# The Algorithms in go and rust language

用Go、Rust语言实现常用数据结构和算法。用于学习目的，算法效率不一定最优。

所有代码都经过具体测试运行，具体环境为：

- 操作系统：Ubuntu-latest 64位
- Go语言版本： Go 1.20
- Rust语言版本： Rust 1.70

## 目录

### 链表相关

#### 双指针技巧

此技巧可用于解决一下一些问题：

1. 合并两个有序链表
2. 链表的分解
3. 合并 k 个有序链表
4. 寻找单链表的倒数第 k 个节点
5. 寻找单链表的中点
6. 判断单链表是否包含环并找出环起点
7. 判断两个单链表是否相交并找出交点

##### 合并两个有序链表

Ref: <https://leetcode.cn/problems/merge-two-sorted-lists/>

[Go实现](/go/linkedlist/merge_two_lists.go)、[Rust实现](/rust/src/linkedlist/merge_two_lists.rs)

代码中用到了`p1`, `p2`两个指针用于比较两个链表节点的推进，`p`指针用于生成新链表节点的推进。在这个过程中没有数据的复制，只有对指针指向的改变，效率很高。

代码中还用到一个链表的算法题中是很常见的「虚拟头结点」技巧，也就是 `dummy` 节点。如果不使用 dummy 虚拟节点，代码会复杂一些，需要额外处理指针 `p` 为空的情况。
> 什么时候需要用虚拟头结点？当需要创造一条新链表的时候，可以使用虚拟头结点简化边界情况的处理。比如说，把两条有序链表合并成一条新的有序链表，是不是要创造一条新链表？再比如想把一条链表分解成两条链表，是不是也在创造新链表？这些情况都可以使用虚拟头结点简化边界情况的处理。
>
> 这个问题还有另外一种解法，就是利用递归，可以去掉两个指针，因为在递归栈中可以自动保存这两个状态。参考上面两个文件中的递归实现。通常来说，递归实现的代码都会更简单，但空间复杂度更高。

##### 单链表的分解

Ref: <https://leetcode.cn/problems/partition-list/description/>

[Go实现](/go/linkedlist/partition.go)、[Rust实现](/rust/src/linkedlist/partition.rs)

我们可以把原链表分成两个小链表，一个链表中的元素大小都小于 x，另一个链表中的元素都大于等于 x，最后再把这两条链表接到一起，就得到了题目想要的结果。

##### 合并K个有序链表

Ref: <https://leetcode.cn/problems/merge-k-sorted-lists/>

[Go实现](/go/linkedlist/merge_k_lists.go)、[Rust实现](/rust/src/linkedlist/merge_k_lists.rs)

合并 k 个有序链表的逻辑类似合并两个有序链表，难点在于，如何快速得到 k 个节点中的最小节点，接到结果链表上？

通常的一种想法是可以把链表头放到数组里，可以写一个循环判断，每次找出一个最小值，这样的时间复杂度是`O(N*K)`，N是所有链表的节点数量，K 是链表个数。

不过还有一种更高效的方式，可以把K个链表头放入 **优先级队列（二叉堆）** 中，放入一个最小堆中，每次就能方便的获得K个节点中的最小值，算法复杂度是：`O(N*Log(K))`。

##### 单链表倒数第K个节点

Ref: <https://leetcode.cn/problems/remove-nth-node-from-end-of-list/>

[Go实现](/go/linkedlist/remove_nth_from_end.go)、[Rust实现](/rust/src/linkedlist/remove_nth_from_end.rs)

要找到链表倒数第k个节点，一般的想法是得知道链表的长度`n`，再用`n-k+1`就得到倒数第k个节点，但是这样需要遍历两次链表，如何在只遍历一次的情况下找到呢？

同样使用双指针技巧，先让一个指针先走`k`步，然后第二个指针开始和第一个指针一起走，第一个指针走完这个整个链表，第二个指针的位置刚好就在`n-k+1`的位置上。

> 另一种方法是使用递归，先一次性递归到链表结尾，递归出栈过程中对计数器加一，当计数器达到k时，此时的节点就是要被删除的节点。

##### 单链表中点

Ref: <https://leetcode.cn/problems/middle-of-the-linked-list/>

[Go实现](/go/linkedlist/middle_node.go)、[Rust实现](/rust/src/linkedlist/middle_node.rs)

这个和上面一题类似，只是这次是找中点。同样使用快慢指针技巧，每当慢指针 slow 前进一步，快指针 fast 就前进两步，这样，当 fast 走到链表末尾时，slow 就指向了链表中点。

##### 判断链表是否包含环，并找出环起点

Ref: <https://leetcode.cn/problems/linked-list-cycle-ii/>

[Go实现](/go/linkedlist/detect_cycle.go)、[Rust实现](/rust/src/linkedlist/detect_cycle.rs)

首先我们得知道链表是否有环，实现方法和上面一题类似，通过快慢指针，慢指针走一步，快指针走两步，如果最终两个指针走到了一起，则证明有环，如果快指针为null，则证明无环。

为什么有环快指针就一定会碰到慢指针？我们随便构造一个有环的链表，就会发现，每次走两步，每经过一次有环处，下一次快指针经过的节点都和上一次不同，所以快慢指针早晚一定会遇到。

如何找到环的起点呢？关键在于当快慢指针相遇时，让其中任一个指针指向头节点，然后让它俩以相同速度前进，再次相遇时所在的节点位置就是环开始的位置。

##### 两个链表是否相交

Ref: <https://leetcode.cn/problems/intersection-of-two-linked-lists/description/>

[Go实现](/go/linkedlist/get_intersection_node.go)、[Rust实现](/rust/src/linkedlist/get_intersection_node.rs)

这题最简单的思路是给每个节点构造HashSet，然后通过HashSet判断是否存在相同节点，但这样空间复杂度较高。

想要通过O(1)的空间复杂度完成就要充分观察相交链表的特点：我们可以通过让两个指针遍历两个链表，想办法让两个指针在相交处相遇，这样就找到了相交节点。

方法一：先分别遍历两个链表，计算出链表长度，将两个长度相减即得到长链表比短链表多出来的长度N，然后让长链表指针先走N步，然后两个指针再一起往前走，两个指针相同了，则找到相交节点。

方法二：分别定义两个指针，用于遍历两个链表，指针1遍历完链表1，则将指针1切换到链表2，对指针2做同样的逻辑，最终当两个指针相同时，则就是相交链表节点。这种方法技巧性更强，不容易想到。

#### 递归魔法

##### 反转链表

Ref: <https://leetcode.cn/problems/reverse-linked-list/>

[Go实现](/go/linkedlist/reverse_list.go)、[Rust实现](/rust/src/linkedlist/reverse_list.rs)

这个题目明显可以使用迭代方法，但需要小心处理临时对象的保存和边界条件的处理。
更简单的方式是使用递归法，通过递归的方式找到最后一个节点，之后从倒数第二个节点开始逐个反转链表。

## 参考资料

本项目内容主要参考：[labuladong 的算法小抄](https://labuladong.github.io/algo/)
