# -*- encoding: utf-8 -*-
'''
@File   : towadd.py
@Time   : 2023-05-17 10:07:53
@Author : pan
'''
def addTwoNumbers(l1, l2):
    carry = 0
    # dummy head
    head = curr = ListNode(0)
    while l1 or l2:
        val = carry
        if l1:
            val += l1.val
            l1 = l1.next
        if l2:
            val += l2.val
            l2 = l2.next
        curr.next = ListNode(val % 10)
        curr = curr.next
        carry = int(val / 10)
    if carry > 0:
        curr.next = ListNode(carry)
    return head.next

class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


if __name__ == "__main__":
    pass