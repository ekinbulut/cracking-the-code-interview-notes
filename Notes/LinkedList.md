# Linked List

### Notes

* The benefit of a linked list is that it can add and remove items from the begining of the list in constant time.
* A linked list does not provide constant time access to the elements.

### The Runner Technique
* Means that you iterate through the list and keep track of the previous node.

### Recursive Problem Solving

* If you are having trouble solving a linked list problem, you may want to look at the recursive solution.

Sample:
```c#
class LinkedList {
    private Node head;
    private Node tail;
    private int size;

    public LinkedList() {
        head = null;
        tail = null;
        size = 0;
    }

    public void add(string value) {
        Node node = new Node(value);
        if (head == null) {
            head = node;
            tail = node;
        } else {
            tail.next = node;
            tail = node;
        }
        size++;
    }

    public string get(int index) {
        if (index < 0 || index >= size) {
            throw new IndexOutOfBoundsException();
        }
        Node node = head;
        for (int i = 0; i < index; i++) {
            node = node.next;
        }
        return node.value;
    }

    public void delete(Node node) {
        if (node == head) {
            head = head.next;
        } else {
            Node current = head;
            while (current.next != node) {
                current = current.next;
            }
            current.next = node.next;
        }
        size--;
    }
}
```