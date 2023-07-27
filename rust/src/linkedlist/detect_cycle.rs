use std::{cell::RefCell, rc::Rc};

use super::ListNode2;

pub fn detect_cycle(head: Option<Rc<RefCell<ListNode2>>>) -> Option<Rc<RefCell<ListNode2>>> {
    let mut slow = head.clone();
    let mut fast = head.clone();

    while fast.is_some() && fast.as_ref().unwrap().borrow().next.is_some() {
        slow = slow.unwrap().borrow().next.clone();
        fast = fast
            .unwrap()
            .borrow()
            .next
            .as_ref()
            .unwrap()
            .borrow()
            .next
            .clone();

        if Rc::ptr_eq(slow.as_ref().unwrap(), fast.as_ref().unwrap()) {
            let mut head = head;
            while !Rc::ptr_eq(head.as_ref().unwrap(), slow.as_ref().unwrap()) {
                slow = slow.unwrap().borrow().next.clone();
                head = head.unwrap().borrow().next.clone();
            }
            return head;
        }
    }

    None
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_detect_cycle() {
        let ret = detect_cycle(None);
        assert!(ret.is_none());

        let head = Some(Rc::new(RefCell::new(ListNode2 { val: 1, next: None })));
        let node1 = Some(Rc::new(RefCell::new(ListNode2 { val: 2, next: None })));
        let node2 = Some(Rc::new(RefCell::new(ListNode2 { val: 3, next: None })));
        let node3 = Some(Rc::new(RefCell::new(ListNode2 { val: 4, next: None })));
        head.as_ref().unwrap().borrow_mut().next = node1.clone();
        node1.as_ref().unwrap().borrow_mut().next = node2.clone();
        node2.as_ref().unwrap().borrow_mut().next = node3.clone();
        node3.as_ref().unwrap().borrow_mut().next = node2.clone();

        let ret = detect_cycle(head);
        assert_eq!(ret, node2);
    }
}
