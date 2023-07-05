/// Ref：https://leetcode.cn/problems/merge-two-sorted-lists/

/// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

pub fn merge_two_lists(
    list1: Option<Box<ListNode>>,
    list2: Option<Box<ListNode>>,
) -> Option<Box<ListNode>> {
    let mut dummy_head = Some(Box::new(ListNode { val: 0, next: None }));
    let mut tail = &mut dummy_head;

    let mut l1 = list1;
    let mut l2 = list2;
    while l1.is_some() && l2.is_some() {
        if l1.as_ref().unwrap().val < l2.as_ref().unwrap().val {
            let next = l1.as_mut().unwrap().next.take();
            tail.as_mut().unwrap().next = l1;
            l1 = next;
        } else {
            let next = l2.as_mut().unwrap().next.take();
            tail.as_mut().unwrap().next = l2;
            l2 = next;
        }
        tail = &mut tail.as_mut().unwrap().next;
    }

    if l1.is_some() {
        tail.as_mut().unwrap().next = l1;
    } else {
        tail.as_mut().unwrap().next = l2;
    }

    dummy_head.unwrap().next
}

#[cfg(test)]
mod tests {
    use crate::linkedlist::merge_two_lists;

    use super::*;

    #[test]
    fn test_list_empty() {
        let ret1 = merge_two_lists(None, None);
        assert!(ret1.is_none());

        let ret2 = merge_two_lists(None, Some(Box::new(ListNode { val: 1, next: None })));
        assert_eq!(ret2, Some(Box::new(ListNode { val: 1, next: None })));

        let ret3 = merge_two_lists(
            Some(Box::new(ListNode { val: 1, next: None })),
            Some(Box::new(ListNode { val: 2, next: None })),
        );
        assert_eq!(
            ret3,
            Some(Box::new(ListNode {
                val: 1,
                next: Some(Box::new(ListNode { val: 2, next: None }))
            }))
        )
    }
}
