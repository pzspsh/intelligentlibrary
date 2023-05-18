# -*- encoding: utf-8 -*-
'''
@File   : towadd.py
@Time   : 2023-05-17 10:07:53
@Author : pan
'''
class ListNode:
    def __init__(self, x, next=None):
        self.val = x
        self.next = next  # 让next默认为None,方便后期传值。

class Solution:
    def addTwoNumbers(self, l1: ListNode, l2: ListNode) -> ListNode:
        f = ListNode(0)  # 定义输出链表的第一个节点
        p = f   # p类似于指针，代表输出链表的第一个节点
        carry = 0  # 进位
        while l1 or l2:  # 如果l1或者l2为空，则循环结束
            x = l1.val if l1 else 0  # 如果l1不为空则取l1.val否则取0
            y = l2.val if l2 else 0
            res = x + y + carry  # 与进位一起相加的结果
            carry = res // 10  # 检查是否有进位
            p.next = ListNode(res % 10)  # 这个时候p实际是输出链表的第一个节点，它指向下一个节点的地址
            p = p.next  # 节点进位，这时p指向的是上一行定义的节点
            l1 = l1.next if l1 else 0  # 如果l1不为空则取l1.next否则取0
            l2 = l2.next if l2 else 0
        if carry > 0:  # 如果到最后还有进位，那么还要新建个节点记录进位的值
            p.next = ListNode(1)
        return f.next  # 返回的链表一定要从定义的输出链表f的第二个节点开始


if __name__ == '__main__':
    a = ListNode(2, ListNode(4, ListNode(3, )))  # 链表a
    b = ListNode(5, ListNode(6, ListNode(4, )))  # 链表b

    obj = Solution()
    Node= obj.addTwoNumbers(a, b)
    while Node:
        print(Node.val)
        Node = Node.next