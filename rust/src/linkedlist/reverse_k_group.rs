use std::{cell::RefCell, rc::Rc};

use super::{ListNode, ListNode2};

pub fn reverse_k_group(head: Option<Box<ListNode>>, k: i32) -> Option<Box<ListNode>> {
    if k <= 1 {
        return head;
    }

    // 1. 从head开始，确认其后至少包含k个节点，否则直接返回head
    let mut head = head;
    let mut curr = &mut head;
    for _ in 0..k - 1 {
        if curr.is_none() || curr.as_ref().unwrap().next.is_none() {
            return head;
        }
        curr = &mut curr.as_mut().unwrap().next;
    }

    // 2. 反转前k个节点
    let mut prev: Option<Box<ListNode>> = None;
    let mut curr = head;
    for _ in 0..k {
        let next = curr.as_mut().unwrap().next.take();
        curr.as_mut().unwrap().next = prev.take();
        prev = curr;
        curr = next;
    }

    // 3. 找到新的末尾节点，并设置其next为后续新递归结果的头节点
    // 注意：这种算法明显不是最优的，因为每次都需要在这里循环找到尾节点，更好的方式可能是使用ListNode2这样的结构
    let mut tail = &mut prev;
    for i in 0..k - 1 {
        tail = &mut tail.as_mut().unwrap().next;
        if i == k - 2 {
            break;
        }
    }
    tail.as_mut().unwrap().next = reverse_k_group(curr, k);

    prev
}

/// 另一种通过`Rc`, `RefCell`实现的版本
pub fn reverse_k_group2(
    head: Option<Rc<RefCell<ListNode2>>>,
    k: i32,
) -> Option<Rc<RefCell<ListNode2>>> {
    if k <= 1 {
        return head;
    }

    // 1. 从head开始，确认其后至少包含k个节点，否则直接返回head
    let mut curr = head.clone();
    for _ in 0..k - 1 {
        if curr.is_none() || curr.as_ref().unwrap().borrow().next.is_none() {
            return head.clone();
        }
        let next = curr.as_deref().unwrap().borrow().next.clone();
        curr = next;
    }

    // 2. 反转前k个节点
    let mut curr = head.clone();
    let mut prev = None;
    for _ in 0..k {
        let next = curr.as_ref().unwrap().borrow_mut().next.take();
        curr.as_ref().unwrap().borrow_mut().next = prev.take();
        prev = curr;
        curr = next;
    }

    // 3. 将翻转后的尾节点的next指向后续递归翻转后的头节点
    let new_tail = head.clone();
    new_tail.as_ref().unwrap().borrow_mut().next = reverse_k_group2(curr, k);

    // 返回翻转后的头节点
    prev
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_reverse_k_group() {
        let ret = reverse_k_group(None, 2);
        assert!(ret.is_none());

        let head = Some(Box::new(ListNode {
            val: 1,
            next: Some(Box::new(ListNode {
                val: 2,
                next: Some(Box::new(ListNode {
                    val: 3,
                    next: Some(Box::new(ListNode {
                        val: 4,
                        next: Some(Box::new(ListNode { val: 5, next: None })),
                    })),
                })),
            })),
        }));

        let ret = reverse_k_group(head, 2);
        assert_eq!(
            ret,
            Some(Box::new(ListNode {
                val: 2,
                next: Some(Box::new(ListNode {
                    val: 1,
                    next: Some(Box::new(ListNode {
                        val: 4,
                        next: Some(Box::new(ListNode {
                            val: 3,
                            next: Some(Box::new(ListNode { val: 5, next: None })),
                        })),
                    })),
                })),
            }))
        );
    }

    #[test]
    fn test_reverse_k_group2() {
        let ret = reverse_k_group2(None, 2);
        assert!(ret.is_none());

        let head = Some(Rc::new(RefCell::new(ListNode2 {
            val: 1,
            next: Some(Rc::new(RefCell::new(ListNode2 {
                val: 2,
                next: Some(Rc::new(RefCell::new(ListNode2 {
                    val: 3,
                    next: Some(Rc::new(RefCell::new(ListNode2 {
                        val: 4,
                        next: Some(Rc::new(RefCell::new(ListNode2 { val: 5, next: None }))),
                    }))),
                }))),
            }))),
        })));

        let ret = reverse_k_group2(head, 2);
        assert_eq!(
            ret,
            Some(Rc::new(RefCell::new(ListNode2 {
                val: 2,
                next: Some(Rc::new(RefCell::new(ListNode2 {
                    val: 1,
                    next: Some(Rc::new(RefCell::new(ListNode2 {
                        val: 4,
                        next: Some(Rc::new(RefCell::new(ListNode2 {
                            val: 3,
                            next: Some(Rc::new(RefCell::new(ListNode2 { val: 5, next: None }))),
                        }))),
                    }))),
                }))),
            })))
        );
    }
}
