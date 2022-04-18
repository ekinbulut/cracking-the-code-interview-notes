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

## StringBuilder

* Creates a resizable array of all the strings, copying them back to string only when necessary.

Complexity:
* Time: O(n)
* Space: O(n)