package main

import "fmt"

//请你设计并实现一个满足 LRU (最近最少使用) 缓存 约束的数据结构。
//
// 实现 LRUCache 类：
//
//
//
//
// LRUCache(int capacity) 以 正整数 作为容量 capacity 初始化 LRU 缓存
// int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
// void put(int key, int value) 如果关键字 key 已经存在，则变更其数据值 value ；如果不存在，则向缓存中插入该组
//key-value 。如果插入操作导致关键字数量超过 capacity ，则应该 逐出 最久未使用的关键字。
//
//
// 函数 get 和 put 必须以 O(1) 的平均时间复杂度运行。
//
//
//
//
//
// 示例：
//
//
//输入
//["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
//[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
//输出
//[null, null, null, 1, null, -1, null, -1, 3, 4]
//
//解释
//LRUCache lRUCache = new LRUCache(2);
//lRUCache.put(1, 1); // 缓存是 {1=1}
//lRUCache.put(2, 2); // 缓存是 {1=1, 2=2}
//lRUCache.get(1);    // 返回 1
//lRUCache.put(3, 3); // 该操作会使得关键字 2 作废，缓存是 {1=1, 3=3}
//lRUCache.get(2);    // 返回 -1 (未找到)
//lRUCache.put(4, 4); // 该操作会使得关键字 1 作废，缓存是 {4=4, 3=3}
//lRUCache.get(1);    // 返回 -1 (未找到)
//lRUCache.get(3);    // 返回 3
//lRUCache.get(4);    // 返回 4
//
//
//
//
// 提示：
//
//
// 1 <= capacity <= 3000
// 0 <= key <= 10000
// 0 <= value <= 10⁵
// 最多调用 2 * 10⁵ 次 get 和 put
//
// Related Topics 设计 哈希表 链表 双向链表 👍 2204 👎 0

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	no146Print("expire: 1, value: %d", obj.Get(1))
	obj.Put(3, 3)
	no146Print("expire: -1, value: %d", obj.Get(2))
	obj.Put(4, 4)
	no146Print("expire: -1, value: %d", obj.Get(1))
	no146Print("expire: 3, value: %d", obj.Get(3))
	no146Print("expire: 4, value: %d", obj.Get(4))
}

func no146Print(format string, params ...interface{}) {
	println(fmt.Sprintf(format, params...))
}

//leetcode submit region begin(Prohibit modification and deletion)
type entry struct {
	key, value int
	pre, next  *entry
}

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*entry
	head, tail *entry
}

func initEntry(key, value int) *entry {
	return &entry{key: key, value: value}
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    make(map[int]*entry),
		head:     initEntry(0, 0),
		tail:     initEntry(0, 0),
	}
	lru.head.next = lru.tail
	lru.tail.pre = lru.head
	return lru
}

func (this *LRUCache) Get(key int) int {
	e, ok := this.cache[key]
	if !ok {
		return -1
	}
	this.remove(e)
	this.addToHead(e)
	return e.value
}

func (this *LRUCache) Put(key int, value int) {
	if e, ok := this.cache[key]; ok {
		this.remove(e)
	}
	for this.size > this.capacity-1 {
		e := this.removeFromTail()
		delete(this.cache, e.key)
	}
	e := initEntry(key, value)
	this.cache[key] = e
	this.addToHead(e)

}

func (this *LRUCache) addToHead(e *entry) {
	e.pre = this.head
	e.next = this.head.next
	this.head.next.pre = e
	this.head.next = e
	this.size++
}

func (this *LRUCache) remove(e *entry) {
	e.pre.next = e.next
	e.next.pre = e.pre
	this.size--
}

func (this *LRUCache) removeFromTail() *entry {
	e := this.tail.pre
	this.remove(e)
	return e
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//leetcode submit region end(Prohibit modification and deletion)
