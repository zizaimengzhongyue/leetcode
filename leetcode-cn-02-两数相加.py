class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

class Solution:

    def reverse(self,l):
        before = l
        cur = l.next
        before.next = None
        while cur:
            next = cur.next
            cur.next = before
            
            before = cur
            cur = next
        return before

    def addTwoNumbers(self,l1,l2):        
        """
        l1 = self.reverse(l1)
        l2 = self.reverse(l2)
        """

        l3 = ListNode(0)
        cur = l3
        temp = 0
        while l1 or l2:
            l1_val = l1.val if l1 else 0
            l2_val = l2.val if l2 else 0
            cur.val = l1_val + l2_val + temp
            temp = cur.val // 10
            cur.val = cur.val % 10
            l1 = l1.next if l1 else None
            l2 = l2.next if l2 else None
            if l1 or l2:
                cur.next = ListNode(temp)
                cur = cur.next
        if temp > 0:
            cur.next = ListNode(temp)
            cur = cur.next

        return l3

l1_1 = ListNode(2)
l1_2 = ListNode(4)
l1_3 = ListNode(3)
l1_1.next = l1_2
l1_2.next = l1_3

l2_1 = ListNode(5)
l2_2 = ListNode(6)
l2_3 = ListNode(4)
l2_1.next = l2_2
l2_2.next = l2_3
"""
l1_1 = ListNode(0)
l2_1 = ListNode(0)
"""

solution = Solution()
l = solution.addTwoNumbers(l1_1,l2_1)
while l:
    print(l.val)
    l = l.next
