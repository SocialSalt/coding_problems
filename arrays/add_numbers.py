# Definition for singly-linked list.
class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

    def __str__(self):
        if self.next is not None:
            final = self.next.__str__()
        else:
            final = ""
        return str(self.val) + final


def addTwoNumbers(l1, l2):
    carry = 0
    root = ListNode()
    current = root
    nxt = True

    while nxt or carry:
        if l1 is None and l2 is None:
            nxt = False
            continue
        try:
            l1val = l1.val
            l1 = l1.next
        except:
            l1val = 0
        try:
            l2val = l2.val
            l2 = l2.next
        except:
            l2val = 0

        add = l1val + l2val + carry
        carry = add >= 10
        add = add % 10
        node = ListNode(add, None)
        current.next = node
        current = node

    return root.next


l1 = ListNode(2, ListNode(4, ListNode(3)))
l2 = ListNode(5, ListNode(6, ListNode(4)))

print(addTwoNumbers(l1, l2))
