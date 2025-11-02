package leetcode
// LFU
type DListNode struct {
    Key int
    Value int
    Freq int
    Prev *DListNode
    Next *DListNode
}

type LFUCache struct {
    Capacity int
    Size int
    MinFreq int
    KeyToNode map[int]*DListNode
    FreqToDList map[int]*DListNode
}


func Constructor(capacity int) LFUCache {
    return LFUCache{
        Capacity: capacity,
        Size: 0,
        MinFreq: 0,
        KeyToNode: make(map[int]*DListNode),
        FreqToDList: make(map[int]*DListNode),
    }
}


func (this *LFUCache) Get(key int) int {
	node, ok := this.KeyToNode[key]
	if !ok {
		return -1
	}
	this.removeNode(node)
	dummy := this.FreqToDList[node.Freq]
	if dummy.Next == dummy {
		// 说明当前频率的链表为空
		delete(this.FreqToDList, node.Freq)
		if node.Freq == this.MinFreq {
			this.MinFreq++
			// 说明当前频率的链表为空，且当前频率就是最小频率，需要更新最小频率
		}
	}
	node.Freq++
	this.pustNode(node.Freq, node)
	return node.Value
}


func (this *LFUCache) Put(key int, value int)  {
    node, ok := this.KeyToNode[key]
    if ok {
		this.removeNode(node)
		dummy := this.FreqToDList[node.Freq]
		if dummy.Next == dummy {
			// 说明当前频率的链表为空
			delete(this.FreqToDList, node.Freq)
			if node.Freq == this.MinFreq {
				this.MinFreq++
				// 说明当前频率的链表为空，且当前频率就是最小频率，需要更新最小频率
			}
		}
		node.Freq++
		this.pustNode(node.Freq, node)
		return
	}
	if this.Size >= this.Capacity {
		// 说明缓存已满，需要删除最小频率的链表的最后一个节点
		dummy := this.FreqToDList[this.MinFreq]
		last := dummy.Prev
		this.removeNode(last)
		delete(this.KeyToNode, last.Key)
		this.Size--
	}
	this.Size++
	this.MinFreq = 1
	this.pustNode(1, &DListNode{
		Key: key,
		Value: value,
		Freq: 1,
	})
	this.KeyToNode[key] = node
}


func (this *LFUCache)removeNode(node *DListNode) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LFUCache)pustNode(freq int, node *DListNode) {
	dummy, ok := this.FreqToDList[freq]
	if !ok {
		dummy = &DListNode{
			Freq: freq,
		}
		dummy.Next = dummy
		dummy.Prev = dummy
		this.FreqToDList[freq] = dummy
	}
	node.Next = dummy.Next
	node.Prev = dummy
	dummy.Next.Prev = node
	dummy.Next = node
}