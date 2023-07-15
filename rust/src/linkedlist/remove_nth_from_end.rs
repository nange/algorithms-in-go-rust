use super::ListNode;

pub fn remove_nth_from_end(head: Option<Box<ListNode>>, n: i32) -> Option<Box<ListNode>> {
    let mut dummy_clone = Some(Box::new(ListNode {
        val: 0,
        next: head.clone(),
    }));
    let mut dummy = Some(Box::new(ListNode { val: 0, next: head }));

    let mut first = &mut dummy_clone;
    let mut second = &mut dummy;

    // 第一步，让first节点先向前移动n+1步
    for i in 0..n + 1 {
        first = &mut first.as_mut().unwrap().next;
        if first.is_none() && i < n {
            return dummy.unwrap().next;
        }
    }

    // 第二步，让first和second节点同时向前移动，直到first节点到达链表的末尾
    while first.is_some() {
        first = &mut first.as_mut().unwrap().next;
        second = &mut second.as_mut().unwrap().next;
    }

    // 第三步，删除second节点的下一个节点
    second.as_mut().unwrap().next = second.as_mut().unwrap().next.as_mut().unwrap().next.take();

    dummy.unwrap().next
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_remove_nth_from_end() {
        let ret1 = remove_nth_from_end(None, 1);
        assert!(ret1.is_none());

        let lists = Some(Box::new(ListNode {
            val: 1,
            next: Some(Box::new(ListNode {
                val: 4,
                next: Some(Box::new(ListNode { val: 5, next: None })),
            })),
        }));
        let ret2 = remove_nth_from_end(lists, 1);
        assert_eq!(
            ret2,
            Some(Box::new(ListNode {
                val: 1,
                next: Some(Box::new(ListNode { val: 4, next: None }))
            }))
        );

        let lists = Some(Box::new(ListNode {
            val: 1,
            next: Some(Box::new(ListNode {
                val: 4,
                next: Some(Box::new(ListNode { val: 5, next: None })),
            })),
        }));
        let ret2 = remove_nth_from_end(lists, 2);
        assert_eq!(
            ret2,
            Some(Box::new(ListNode {
                val: 1,
                next: Some(Box::new(ListNode { val: 5, next: None }))
            }))
        );

        let lists = Some(Box::new(ListNode {
            val: 1,
            next: Some(Box::new(ListNode {
                val: 4,
                next: Some(Box::new(ListNode { val: 5, next: None })),
            })),
        }));
        let ret2 = remove_nth_from_end(lists, 3);
        assert_eq!(
            ret2,
            Some(Box::new(ListNode {
                val: 4,
                next: Some(Box::new(ListNode { val: 5, next: None }))
            }))
        );
    }
}
