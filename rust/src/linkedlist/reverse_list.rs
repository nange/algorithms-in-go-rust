use super::ListNode;

pub fn reverse_list(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    // 定义helper函数，帮助递归状态的保存。
    // 因为递归的时候除了需要head参数外，还需要一个状态就是原始链表当前节点的前一个节点，
    // 因为我们要把当前节点的next节点设置为前一个节点
    // 这是由于Rust的所有权生命周期规则限制造成的，如果不这样做，则需要使用RefCell和Rc来维护和更新状态
    fn helper(head: Option<Box<ListNode>>, prev: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
        match head {
            Some(mut node) => {
                let next = node.next.take();
                node.next = prev;
                helper(next, Some(node))
            }
            None => prev,
        }
    }

    helper(head, None) // head没有前一个节点，所以是None
}

#[allow(dead_code)]
fn reverse_list_with_loop(head: Option<Box<ListNode>>) -> Option<Box<ListNode>> {
    let mut head = head;
    let mut prev: Option<Box<ListNode>> = None;

    while let Some(mut node) = head {
        head = node.next.take();
        node.next = prev;
        prev = Some(node);
    }

    prev
}

#[cfg(test)]
mod tests {
    use super::super::ListNode;
    use super::*;

    #[test]
    fn test_reverse_list() {
        let ret = reverse_list(None);
        assert!(ret.is_none());

        let head = Some(Box::new(ListNode {
            val: 1,
            next: Some(Box::new(ListNode {
                val: 2,
                next: Some(Box::new(ListNode { val: 3, next: None })),
            })),
        }));
        let expect = Some(Box::new(ListNode {
            val: 3,
            next: Some(Box::new(ListNode {
                val: 2,
                next: Some(Box::new(ListNode { val: 1, next: None })),
            })),
        }));
        let ret = reverse_list(head.clone());
        assert_eq!(ret, expect.clone());

        let ret = reverse_list_with_loop(head.clone());
        assert_eq!(ret, expect.clone());
    }
}
