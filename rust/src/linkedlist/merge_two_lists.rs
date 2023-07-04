/// 合并两个有序链表：https://leetcode.cn/problems/merge-two-sorted-lists/

/// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

pub fn merge_two_lists(
    list1: Option<Box<ListNode>>,
    list2: Option<Box<ListNode>>,
) -> Option<Box<ListNode>> {
    if list1.is_none() {
        return list2;
    }
    if list2.is_none() {
        return list1;
    }

    let mut dummy = ListNode::new(-1);
    let mut p = &mut dummy;

    let mut p1 = list1.unwrap();
    let mut p2 = list2.unwrap();

    loop {}

    todo!()
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_list_empty() {
        let ret1 = merge_two_lists(None, None);
        assert!(ret1.is_none());

        let ret2 = merge_two_lists(None, Some(Box::new(ListNode::new(1))));
        assert_eq!(ret2, Some(Box::new(ListNode::new(1))));
    }
}
