# Coding Interviews

Algorithms, data structures, Leetcode style problems and just coding-interview style katas.


## To Do

### **1. Dynamic Array**

```go
// Basic CRUD + utility operations
func (a *MyDynamicArray[T]) Add(val T)                  // append at end
func (a *MyDynamicArray[T]) Insert(index int, val T)    // insert at specific index
func (a *MyDynamicArray[T]) Remove(index int) T         // remove element at index
func (a *MyDynamicArray[T]) Get(index int) T
func (a *MyDynamicArray[T]) Set(index int, val T)
func (a *MyDynamicArray[T]) Size() int
func (a *MyDynamicArray[T]) Capacity() int
func (a *MyDynamicArray[T]) Clear()
func (a *MyDynamicArray[T]) Contains(val T) bool
func (a *MyDynamicArray[T]) IndexOf(val T) int
func (a *MyDynamicArray[T]) ToSlice() []T
```

---

### **2. Stack**

```go
func (s *Stack[T]) Push(val T)
func (s *Stack[T]) Pop() T
func (s *Stack[T]) Peek() T
func (s *Stack[T]) IsEmpty() bool
func (s *Stack[T]) Size() int
func (s *Stack[T]) Clear()
```

---

### **3. Queue**

```go
func (q *Queue[T]) Enqueue(val T)
func (q *Queue[T]) Dequeue() T
func (q *Queue[T]) Peek() T
func (q *Queue[T]) IsEmpty() bool
func (q *Queue[T]) Size() int
func (q *Queue[T]) Clear()
```

---

### **4. Deque (Double-Ended Queue)**

```go
func (d *Deque[T]) AddFront(val T)
func (d *Deque[T]) AddBack(val T)
func (d *Deque[T]) RemoveFront() T
func (d *Deque[T]) RemoveBack() T
func (d *Deque[T]) PeekFront() T
func (d *Deque[T]) PeekBack() T
func (d *Deque[T]) IsEmpty() bool
func (d *Deque[T]) Size() int
func (d *Deque[T]) Clear()
```

---

### **5. Linked List (Singly / Doubly)**

```go
func (l *LinkedList[T]) AddFront(val T)
func (l *LinkedList[T]) AddBack(val T)
func (l *LinkedList[T]) InsertAt(index int, val T)
func (l *LinkedList[T]) RemoveAt(index int) T
func (l *LinkedList[T]) RemoveVal(val T) bool
func (l *LinkedList[T]) Get(index int) T
func (l *LinkedList[T]) Set(index int, val T)
func (l *LinkedList[T]) Size() int
func (l *LinkedList[T]) IsEmpty() bool
func (l *LinkedList[T]) Contains(val T) bool
func (l *LinkedList[T]) ToSlice() []T
```

---

### **6. Hash Map**

```go
func (m *HashMap[K, V]) Put(key K, val V)
func (m *HashMap[K, V]) Get(key K) (V, bool)
func (m *HashMap[K, V]) Remove(key K)
func (m *HashMap[K, V]) ContainsKey(key K) bool
func (m *HashMap[K, V]) Keys() []K
func (m *HashMap[K, V]) Values() []V
func (m *HashMap[K, V]) Size() int
func (m *HashMap[K, V]) Clear()
```

---

### **7. Set**

```go
func (s *Set[T]) Add(val T)
func (s *Set[T]) Remove(val T)
func (s *Set[T]) Contains(val T) bool
func (s *Set[T]) Size() int
func (s *Set[T]) Clear()
func (s *Set[T]) Union(other *Set[T]) *Set[T]
func (s *Set[T]) Intersection(other *Set[T]) *Set[T]
func (s *Set[T]) Difference(other *Set[T]) *Set[T]
```

---

### **8. Binary Tree**

```go
func (t *BinaryTree[T]) Insert(val T)
func (t *BinaryTree[T]) Delete(val T)
func (t *BinaryTree[T]) Find(val T) *TreeNode[T]
func (t *BinaryTree[T]) Inorder() []T
func (t *BinaryTree[T]) Preorder() []T
func (t *BinaryTree[T]) Postorder() []T
func (t *BinaryTree[T]) Size() int
func (t *BinaryTree[T]) Height() int
func (t *BinaryTree[T]) Clear()
```

---

### **9. Binary Search Tree (BST)**

```go
func (bst *BST[T]) Insert(val T)
func (bst *BST[T]) Delete(val T)
func (bst *BST[T]) Find(val T) *TreeNode[T]
func (bst *BST[T]) Min() T
func (bst *BST[T]) Max() T
func (bst *BST[T]) Inorder() []T
func (bst *BST[T]) Predecessor(val T) *TreeNode[T]
func (bst *BST[T]) Successor(val T) *TreeNode[T]
func (bst *BST[T]) Size() int
func (bst *BST[T]) Height() int
func (bst *BST[T]) Clear()
```

---

### **10. Heap (Min / Max)**

```go
func (h *Heap[T]) Insert(val T)
func (h *Heap[T]) Extract() T        // Min or Max
func (h *Heap[T]) Peek() T
func (h *Heap[T]) Size() int
func (h *Heap[T]) Clear()
func (h *Heap[T]) Heapify()          // Optional: build heap from slice
```

---

### **11. Graph (Adjacency List)**

```go
func (g *Graph[T]) AddVertex(v T)
func (g *Graph[T]) RemoveVertex(v T)
func (g *Graph[T]) AddEdge(u, v T)
func (g *Graph[T]) RemoveEdge(u, v T)
func (g *Graph[T]) Neighbors(v T) []T
func (g *Graph[T]) BFS(start T) []T
func (g *Graph[T]) DFS(start T) []T
func (g *Graph[T]) HasPath(u, v T) bool
func (g *Graph[T]) Clear()
```

---

### **12. Trie**

```go
func (t *Trie) Insert(word string)
func (t *Trie) Search(word string) bool
func (t *Trie) StartsWith(prefix string) bool
func (t *Trie) Delete(word string) bool
func (t *Trie) Clear()
```

---

### **13. Priority Queue**

```go
func (pq *PriorityQueue[T]) Insert(val T, priority int)
func (pq *PriorityQueue[T]) Extract() (T, int)    // returns val + priority
func (pq *PriorityQueue[T]) Peek() (T, int)
func (pq *PriorityQueue[T]) Size() int
func (pq *PriorityQueue[T]) Clear()
```

---


### **14.  Array / Slice Operations**

```go
func MaxSlidingWindow(nums []int, k int) []int
func MinSlidingWindow(nums []int, k int) []int
func PrefixSum(nums []int) []int
func BinarySearch(nums []int, target int) int
func Kadane(nums []int) int                          // Max contiguous sum
func Reverse(nums []int)
func Rotate(nums []int, k int)
func Partition(nums []int, pivot int)
func MergeIntervals(intervals [][]int) [][]int
func Subarrays(nums []int) [][]int
func TwoPointersSum(nums []int, target int) (int,int)
func CountFrequency(nums []int) map[int]int
```

