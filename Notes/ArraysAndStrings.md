# Arrays and Strings

## Hash Tables

1. Compute a hash code
    * returns int or long
    * two diffrent keys could have the same hash code

2. Map hash code to an index
    * index = hash code % array_size

3. At the index there is a linked list of keys and values
    * if the key is already in the hash table, then update the value
    * if the key is not in the hash table, then add the key and value to the linked list

To retrive a value from the hash table, we need to compute the hash code of the key, then use the hash code to find the index of the linked list. Then we need to traverse the linked list to find the key.

Complexity:
* Time: O(n)
* Space: O(n)

Sample:
```c#
class HashTable {
    private int arraySize;
    private LinkedList<KeyValuePair>[] array;

    public HashTable(int arraySize) {
        this.arraySize = arraySize;
        array = new LinkedList[arraySize];
    }

    public void put(int key, int value) {
        int hashCode = key % arraySize;
        if (array[hashCode] == null) {
            array[hashCode] = new LinkedList<KeyValuePair>();
        }
        for (KeyValuePair pair : array[hashCode]) {
            if (pair.key == key) {
                pair.value = value;
                return;
            }
        }
        array[hashCode].add(new KeyValuePair(key, value));
    }

    public int get(int key) {
        int hashCode = key % arraySize;
        if (array[hashCode] == null) {
            return -1;
        }
        for (KeyValuePair pair : array[hashCode]) {
            if (pair.key == key) {
                return pair.value;
            }
        }
        return -1;
    }
}
```


### Alternative:

We can use balanced binary search trees to store the keys and values. With complexity O(log n), we can find the key in the tree.
* Using less space than hash table
* We can itarate through the tree to find the key in order

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