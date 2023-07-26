use super::ListNode;

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

#[allow(dead_code)]
fn merge_two_lists_with_recursion(
    list1: Option<Box<ListNode>>,
    list2: Option<Box<ListNode>>,
) -> Option<Box<ListNode>> {
    match (list1, list2) {
        (Some(node1), Some(node2)) => {
            if node1.val <= node2.val {
                Some(Box::new(ListNode {
                    val: node1.val,
                    next: merge_two_lists_with_recursion(node1.next, Some(node2)),
                }))
            } else {
                Some(Box::new(ListNode {
                    val: node2.val,
                    next: merge_two_lists_with_recursion(Some(node1), node2.next),
                }))
            }
        }
        (Some(node1), None) => Some(node1),
        (None, Some(node2)) => Some(node2),
        (None, None) => None,
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_list_empty() {
        let list_fn: [fn(
            list1: Option<Box<ListNode>>,
            list2: Option<Box<ListNode>>,
        ) -> Option<Box<ListNode>>; 2] = [merge_two_lists, merge_two_lists_with_recursion];

        for f in list_fn {
            let ret1 = f(None, None);
            assert!(ret1.is_none());

            let ret2 = f(None, Some(Box::new(ListNode { val: 1, next: None })));
            assert_eq!(ret2, Some(Box::new(ListNode { val: 1, next: None })));

            let ret3 = f(
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
}
