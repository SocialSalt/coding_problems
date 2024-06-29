import queue
from typing import Optional, TypeVar

T = TypeVar("T")


class Node[T]:
    def __init__(self, value: T, left=None, right=None):
        self.left: Node | None = left
        self.right: Node | None = right
        self.value: T = value

    def __str__(self):
        return f"{self.value}"


class Stack:
    def __init__(self, items: Optional[list] = None):
        if items:
            self._value: list = items
        else:
            self._value: list = []

    def put(self, item):
        self._value.append(item)

    def get(self):
        return self._value.pop()

    def empty(self) -> bool:
        return len(self._value) == 0


def find_node(root, value):
    q = Stack()  # queue.Queue()

    q.put(root)
    while not q.empty():
        node = q.get()
        if node.value == value:
            return node
        if node.left:
            q.put(node.left)
        if node.right:
            q.put(node.right)
    raise ValueError(f"value: {value} not found in tree")


def main():
    tree: Node[int] = Node(
        10, Node(5, Node(2), Node(8)), Node(30, Node(90), Node(56, None, Node(0)))
    )
    try:
        print(find_node(tree, 6))
    except:
        print("6 not found")
    print(find_node(tree, 56))
    print(find_node(tree, 5))
    print(find_node(tree, 10))


main()
