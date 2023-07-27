mod detect_cycle;
mod get_intersection_node;
mod merge_k_lists;
mod merge_two_lists;
mod middle_node;
mod partition;
mod remove_nth_from_end;

pub use detect_cycle::detect_cycle;
pub use get_intersection_node::get_intersection_node;
pub use merge_k_lists::merge_k_lists;
pub use merge_two_lists::merge_two_lists;
pub use middle_node::middle_node;
pub use partition::partition;
pub use remove_nth_from_end::remove_nth_from_end;

use std::{cell::RefCell, cmp::Ord, rc::Rc};

/// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl Ord for ListNode {
    fn cmp(&self, other: &Self) -> std::cmp::Ordering {
        self.val.cmp(&other.val)
    }
}

impl PartialOrd for ListNode {
    fn partial_cmp(&self, other: &Self) -> Option<std::cmp::Ordering> {
        self.val.partial_cmp(&other.val)
    }
}

#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode2 {
    pub val: i32,
    pub next: Option<Rc<RefCell<ListNode2>>>,
}
