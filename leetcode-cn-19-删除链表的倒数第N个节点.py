class Solution:
    def removeNthFromEnd(self, head, n):
        first_node = head
        second_node = head.next
        third_node = None
        if not second_node is None and not second_node.next is None:
            third_node = second_node.next
        index = -1
        length = 0
        cur = head
        while not cur is None:
            if index < n:
                index = index + 1
            else:
                if not first_node is None:
                    first_node = first_node.next
                if not second_node is None:
                    second_node = second_node.next
                third_node = None
                if not third_node is None:
                    third_node = third_node.next
            cur = cur.next
            length = length + 1
        if length == n:
            head = second_node
        elif not first_node is None and not second_node is None:
            first_node.next = second_node.next
        elif not first_node is None:
            first_node.next = None
        return head


class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

nums = [1,2,3,4,5]
n = 2
cur = None
head = cur
for i in range(len(nums) - 1, -1, -1):
    head = ListNode(nums[i])
    head.next = cur
    cur = head

solution = Solution()
ans = solution.removeNthFromEnd(head, 2)
while not ans is None:
    print(ans.val)
    ans = ans.next
