use std::{cell::RefCell, rc::Rc};

use super::ListNode2;

pub fn get_intersection_node(
    head1: Option<Rc<RefCell<ListNode2>>>,
    head2: Option<Rc<RefCell<ListNode2>>>,
) -> Option<Rc<RefCell<ListNode2>>> {
    let mut p1 = head1.clone();
    let mut p2 = head2.clone();
    let mut len_a = 0;
    let mut len_b = 0;

    while p1.is_some() {
        p1 = p1.unwrap().borrow().next.clone();
        len_a += 1;
    }
    while p2.is_some() {
        p2 = p2.unwrap().borrow().next.clone();
        len_b += 1;
    }
    if len_a == 0 || len_b == 0 {
        return None;
    }

    p1 = head1.clone();
    p2 = head2.clone();
    if len_a >= len_b {
        // A 比 B 长，则A先走 lenA - lenB 步
        for _ in 0..(len_a - len_b) {
            p1 = p1.unwrap().borrow().next.clone();
        }
    } else {
        // B 比 A 长，则B先走 lenB - lenA 步
        for _ in 0..(len_b - len_a) {
            p2 = p2.unwrap().borrow().next.clone();
        }
    }

    // 再一起往后走，直到相遇
    while !Rc::ptr_eq(p1.as_ref().unwrap(), p2.as_ref().unwrap()) {
        p1 = p1.unwrap().borrow().next.clone();
        p2 = p2.unwrap().borrow().next.clone();
    }

    p1
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_get_intersection_node() {
        let ret = get_intersection_node(None, None);
        assert!(ret.is_none());

        let head1 = Some(Rc::new(RefCell::new(ListNode2 { val: 1, next: None })));
        let head2 = Some(Rc::new(RefCell::new(ListNode2 {
            val: 10,
            next: None,
        })));
        let node1 = Some(Rc::new(RefCell::new(ListNode2 { val: 2, next: None })));
        let node2 = Some(Rc::new(RefCell::new(ListNode2 {
            val: 11,
            next: None,
        })));
        let node_inte = Some(Rc::new(RefCell::new(ListNode2 {
            val: 3,
            next: Some(Rc::new(RefCell::new(ListNode2 { val: 4, next: None }))),
        })));

        head1.as_ref().unwrap().borrow_mut().next = node1.clone();
        head2.as_ref().unwrap().borrow_mut().next = node2.clone();
        node1.as_ref().unwrap().borrow_mut().next = node_inte.clone();
        node2.as_ref().unwrap().borrow_mut().next = node_inte.clone();

        let ret = get_intersection_node(head1, head2);
        assert_eq!(ret, node_inte.clone());
    }
}
