# Arrays and Strings

## ArrayList and Resiable Arrays

* Dynamic resize
* ArrayList resizes itself while still providing constant time access to the elements O(1)
* Resizable array is a fixed size array that can grow and shrink
* When the array is full, it doubles in size O(n)

Complexity:
* Time: O(1)
* Space: O(n)

Sample:
```c#
class  ArrayList {
    private int size;
    private int capacity;
    private int[] array;

    public ArrayList() {
        size = 0;
        capacity = 1;
        array = new int[capacity];
    }

    public void add(int value) {
        if (size == capacity) {
            capacity *= 2;
            int[] newArray = new int[capacity];
            for (int i = 0; i < size; i++) {
                newArray[i] = array[i];
            }
            array = newArray;
        }
        array[size++] = value;
    }

    public int get(int index) {
        if (index < 0 || index >= size) {
            throw new IndexOutOfBoundsException();
        }
        return array[index];
    }
}
```

## StringBuilder

* Creates a resizable array of all the strings, copying them back to string only when necessary.

Complexity:
* Time: O(n)
* Space: O(n)

Sample:
```c#
class StringBuilder {
    private int size;
    private int capacity;
    private char[] array;

    public StringBuilder() {
        size = 0;
        capacity = 1;
        array = new char[capacity];
    }

    public void append(char c) {
        if (size == capacity) {
            capacity *= 2;
            char[] newArray = new char[capacity];
            for (int i = 0; i < size; i++) {
                newArray[i] = array[i];
            }
            array = newArray;
        }
        array[size++] = c;
    }

    public String toString() {
        return new String(array, 0, size);
    }
}
```
