package leetcode

/*
题目：LRU缓存
请你设计并实现一个满足 LRU (最近最少使用) 缓存 约束的数据结构。
实现 LRUCache 类：
LRUCache(int capacity) 以 正整数 作为容量capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value)如果关键字key 已经存在，则变更其数据值value ；如果不存在，则向缓存中插入该组key-value 。如果插入操作导致关键字数量超过capacity ，则应该 逐出 最久未使用的关键字。

*/

type LRUCache struct {
	size       int
	capacity   int
	cache      map[int]*DLinkList
	head, tail *DLinkList // 双向链表头尾
}

type DLinkList struct {
	Key, Value int
	Pre, Next  *DLinkList
}

func initDLinkList(key, val int) *DLinkList {
	return &DLinkList{
		Key:   key,
		Value: val,
		Pre:   nil,
		Next:  nil,
	}
}

func ConstructorLRU(capacity int) LRUCache {
	l := LRUCache{
		size:     0,
		capacity: capacity,
		cache:    map[int]*DLinkList{},
		head:     initDLinkList(0, 0),
		tail:     initDLinkList(0, 0),
	}
	l.head.Next = l.tail
	l.tail.Pre = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.Value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.cache[key]; !ok {
		node := initDLinkList(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.Key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.Value = value
		this.moveToHead(node)
	}
}

func (l *LRUCache) addToHead(node *DLinkList) {
	node.Pre = l.head
	node.Next = l.head.Next
	l.head.Next.Pre = node
	l.head.Next = node
}

func (d *LRUCache) removeNode(node *DLinkList) {
	node.Pre.Next = node.Next
	node.Next.Pre = node.Pre
}

func (d *LRUCache) moveToHead(node *DLinkList) {
	d.removeNode(node)
	d.addToHead(node)
}

func (d *LRUCache) removeTail() *DLinkList {
	node := d.tail.Pre
	d.removeNode(node)
	return node
}
