use super::ListNode;
use std::cmp::Reverse;
use std::collections::BinaryHeap;

pub fn merge_k_lists(lists: Vec<Option<Box<ListNode>>>) -> Option<Box<ListNode>> {
    // 将链表头节点加入堆
    let mut heap = BinaryHeap::new();
    for node in lists.into_iter().flatten() {
        heap.push(Reverse(node));
    }

    // 虚拟头结点
    let mut dummy = Box::new(ListNode { val: 0, next: None });
    // 指向虚拟头节点尾部节点，用于设置新节点
    let mut tail = &mut dummy;

    // 从堆中逐个取出节点，设置到新链表上，
    // 并把取出节点的next节点重新加入堆中
    while let Some(Reverse(mut node)) = heap.pop() {
        if let Some(next) = node.next.take() {
            heap.push(Reverse(next));
        }
        tail.next = Some(node);
        tail = tail.next.as_mut().unwrap();
    }

    dummy.next
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_merge_k_lists() {
        let ret1 = merge_k_lists(Vec::new());
        assert!(ret1.is_none());

        let lists = vec![
            Some(Box::new(ListNode {
                val: 1,
                next: Some(Box::new(ListNode {
                    val: 4,
                    next: Some(Box::new(ListNode { val: 5, next: None })),
                })),
            })),
            Some(Box::new(ListNode {
                val: 1,
                next: Some(Box::new(ListNode {
                    val: 3,
                    next: Some(Box::new(ListNode { val: 4, next: None })),
                })),
            })),
            Some(Box::new(ListNode {
                val: 2,
                next: Some(Box::new(ListNode { val: 6, next: None })),
            })),
        ];
        let ret2 = merge_k_lists(lists);
        assert_eq!(
            ret2,
            Some(Box::new(ListNode {
                val: 1,
                next: Some(Box::new(ListNode {
                    val: 1,
                    next: Some(Box::new(ListNode {
                        val: 2,
                        next: Some(Box::new(ListNode {
                            val: 3,
                            next: Some(Box::new(ListNode {
                                val: 4,
                                next: Some(Box::new(ListNode {
                                    val: 4,
                                    next: Some(Box::new(ListNode {
                                        val: 5,
                                        next: Some(Box::new(ListNode { val: 6, next: None }))
                                    }))
                                }))
                            }))
                        }))
                    }))
                }))
            }))
        );
    }
}
