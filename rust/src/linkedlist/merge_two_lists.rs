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
    let mut dummy = Some(Box::new(ListNode { val: 0, next: None }));
    let mut p = &mut dummy;

    let mut p1 = list1;
    let mut p2 = list2;

    while p1.is_some() && p2.is_some() {
        if p1.as_ref().unwrap().val <= p2.as_ref().unwrap().val {
            let next = p1.as_mut().unwrap().next.take();
            p.as_mut().unwrap().next = p1;
            p1 = next;
        } else {
            let next = p2.as_mut().unwrap().next.take();
            p.as_mut().unwrap().next = p2;
            p2 = next;
        }

        p = &mut p.as_mut().unwrap().next;
    }

    if p1.is_some() {
        p.as_mut().unwrap().next = p1;
    }
    if p2.is_some() {
        p.as_mut().unwrap().next = p2;
    }

    dummy.unwrap().next
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
            Some(Box::new(ListNode {
                val: 1,
                next: Some(Box::new(ListNode { val: 3, next: None })),
            })),
            Some(Box::new(ListNode { val: 2, next: None })),
        );
        assert_eq!(
            ret3,
            Some(Box::new(ListNode {
                val: 1,
                next: Some(Box::new(ListNode {
                    val: 2,
                    next: Some(Box::new(ListNode { val: 3, next: None }))
                }))
            }))
        )
    }
}
