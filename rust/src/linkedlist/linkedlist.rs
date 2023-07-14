use std::cmp::Ord;

/// Definition for singly-linked list.
#[derive(PartialEq, PartialOrd, Ord, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}
