use super::ListNode;

pub fn partition(head: Option<Box<ListNode>>, x: i32) -> Option<Box<ListNode>> {
    // 左节点用于存储小于x的值
    let mut dummy_left = Some(Box::new(ListNode { val: 0, next: None }));
    // 右节点用于存储大于x的值
    let mut dummy_right = Some(Box::new(ListNode { val: 0, next: None }));

    // p1, p2引用分别用生成新的左右链表节点
    let mut p1 = &mut dummy_left;
    let mut p2 = &mut dummy_right;

    // p用于遍历head链表
    let mut p = head;

    while p.is_some() {
        // 断开新生成的p1,p2引用和原链表的连接关系
        let tmp = p.as_mut().unwrap().next.take();
        if p.as_ref().unwrap().val < x {
            // 设置左链表的next为当前p节点
            p1.as_mut().unwrap().next = p;
            // 移动p1引用到最新位置
            p1 = &mut p1.as_mut().unwrap().next;
        } else {
            p2.as_mut().unwrap().next = p;
            p2 = &mut p2.as_mut().unwrap().next;
        }
        p = tmp;
    }

    p1.as_mut().unwrap().next = dummy_right.as_mut().unwrap().next.take();

    dummy_left.unwrap().next
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_partition() {
        let ret1 = partition(None, 0);
        assert!(ret1.is_none());

        let ret2 = partition(
            Some(Box::new(ListNode {
                val: 1,
                next: Some(Box::new(ListNode {
                    val: 4,
                    next: Some(Box::new(ListNode {
                        val: 3,
                        next: Some(Box::new(ListNode {
                            val: 2,
                            next: Some(Box::new(ListNode {
                                val: 5,
                                next: Some(Box::new(ListNode { val: 2, next: None })),
                            })),
                        })),
                    })),
                })),
            })),
            3,
        );
        assert_eq!(
            ret2,
            Some(Box::new(ListNode {
                val: 1,
                next: Some(Box::new(ListNode {
                    val: 2,
                    next: Some(Box::new(ListNode {
                        val: 2,
                        next: Some(Box::new(ListNode {
                            val: 4,
                            next: Some(Box::new(ListNode {
                                val: 3,
                                next: Some(Box::new(ListNode { val: 5, next: None })),
                            })),
                        })),
                    })),
                })),
            }))
        )
    }
}
