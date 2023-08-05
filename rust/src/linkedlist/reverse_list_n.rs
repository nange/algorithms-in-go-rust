use super::ListNode2;
use std::cell::RefCell;
use std::rc::Rc;

pub fn reverse_list_n(
    head: Option<Rc<RefCell<ListNode2>>>,
    n: i32,
) -> Option<Rc<RefCell<ListNode2>>> {
    let (last, _) = helper(head, n);

    last
}

fn helper(
    node: Option<Rc<RefCell<ListNode2>>>,
    n: i32,
) -> (
    Option<Rc<RefCell<ListNode2>>>,
    Option<Rc<RefCell<ListNode2>>>,
) {
    if n <= 0 || (node.is_none() || node.as_ref().unwrap().borrow().next.is_none()) {
        // 边界条件
        return (node.clone(), node.clone());
    }

    if n == 1 {
        // 递归退出条件
        return (node.clone(), node.as_ref().unwrap().borrow().next.clone());
    }

    let (last, successor) = helper(node.as_ref().unwrap().borrow().next.clone(), n - 1);
    node.as_ref()
        .unwrap()
        .borrow()
        .next
        .as_ref()
        .unwrap()
        .borrow_mut()
        .next = node.clone();
    node.as_ref().unwrap().borrow_mut().next = successor.clone();

    (last, successor)
}

#[cfg(test)]
mod tests {
    use super::super::ListNode2;
    use super::*;

    #[test]
    fn test_reverse_list_n() {
        let ret = reverse_list_n(None, 0);
        assert!(ret.is_none());

        let ret = reverse_list_n(
            Some(Rc::new(RefCell::new(ListNode2 {
                val: 1,
                next: Some(Rc::new(RefCell::new(ListNode2 {
                    val: 2,
                    next: Some(Rc::new(RefCell::new(ListNode2 {
                        val: 3,
                        next: Some(Rc::new(RefCell::new(ListNode2 { val: 4, next: None }))),
                    }))),
                }))),
            }))),
            2,
        );
        assert_eq!(
            ret,
            Some(Rc::new(RefCell::new(ListNode2 {
                val: 2,
                next: Some(Rc::new(RefCell::new(ListNode2 {
                    val: 1,
                    next: Some(Rc::new(RefCell::new(ListNode2 {
                        val: 3,
                        next: Some(Rc::new(RefCell::new(ListNode2 { val: 4, next: None }))),
                    }))),
                }))),
            })))
        )
    }
}
