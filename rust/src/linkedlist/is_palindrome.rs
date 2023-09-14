use std::cell::RefCell;
use std::rc::Rc;

use super::ListNode2;

pub fn is_palindrome(head: Option<Rc<RefCell<ListNode2>>>) -> bool {
    if head.is_none() {
        return false;
    }
    if head.as_ref().unwrap().borrow().next.is_none() {
        return true;
    }
    // 1. 找到链表中点
    let mut slow = head.clone();
    let mut fast = head.clone();
    while !fast.is_none() && !fast.as_ref().unwrap().borrow().next.is_none() {
        let next = slow.as_ref().unwrap().borrow().next.clone();
        slow = next;
        let next2 = fast.as_ref().unwrap().borrow().next.as_ref().unwrap().borrow().next.clone();
        fast = next2;
    }
    let middle = slow;

    // 2. 反转中点之后的链表
    let mut pre = None;
    let mut curr = middle.clone();
    while !curr.is_none() {
        let next = curr.as_ref().unwrap().borrow().next.clone();
        curr.as_ref().unwrap().borrow_mut().next = pre.take();
        pre = curr.clone();
        curr = next;
    }


    // 3. 比较新节点`new`和老节点`head`是否一致
    let mut new = pre;
    let mut old = head.clone();
    while !new.is_none() {
        if new.as_ref().unwrap().borrow().val != old.as_ref().unwrap().borrow().val {
            return false;
        }
        let new_next = new.as_ref().unwrap().borrow().next.clone();
        let old_next = old.as_ref().unwrap().borrow().next.clone();
        new = new_next;
        old = old_next;
    }

    true
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_is_palindrome_v1() {
        let res = is_palindrome(None);
        assert!(!res);

        let res = is_palindrome(Some(Rc::new(RefCell::new(ListNode2::new(0)))));
        assert!(res);

        let head = Some(Rc::new(RefCell::new(ListNode2 {
            val: 1,
            next: Some(Rc::new(RefCell::new(ListNode2 {
                val: 1,
                next: None,
            }))),
        })));
        let res = is_palindrome(head);
        assert!(res);

        let head = Some(Rc::new(RefCell::new(ListNode2 {
            val: 1,
            next: Some(Rc::new(RefCell::new(ListNode2 {
                val: 2,
                next: Some(Rc::new(RefCell::new(ListNode2 {
                    val: 1,
                    next: None,
                }))),
            }))),
        })));
        let res = is_palindrome(head);
        assert!(res);

        let head = Some(Rc::new(RefCell::new(ListNode2 {
            val: 1,
            next: Some(Rc::new(RefCell::new(ListNode2 {
                val: 2,
                next: Some(Rc::new(RefCell::new(ListNode2 {
                    val: 2,
                    next: Some(Rc::new(RefCell::new(ListNode2 {
                        val: 1,
                        next: None,
                    }))),
                }))),
            }))),
        })));
        let res = is_palindrome(head);
        assert!(res);

        let head = Some(Rc::new(RefCell::new(ListNode2 {
            val: 1,
            next: Some(Rc::new(RefCell::new(ListNode2 {
                val: 2,
                next: Some(Rc::new(RefCell::new(ListNode2 {
                    val: 2,
                    next: Some(Rc::new(RefCell::new(ListNode2 {
                        val: 1,
                        next: Some(Rc::new(RefCell::new(ListNode2 {
                            val: 1,
                            next: None,
                        }))),
                    }))),
                }))),
            }))),
        })));
        let res = is_palindrome(head);
        assert!(!res);
    }
}