package ZuoYeBang

/*
字典树 from 作业帮
*/

type TireTree struct {
	children [26]*TireTree
	isEnd    bool
}

func NewTire() *TireTree {
	return &TireTree{}
}

func (t *TireTree) Insert(s string) {
	node := t
	for _, ch := range s {
		num := ch - 'a'
		if node.children[num] == nil {
			node.children[num] = NewTire()
		}
		node = node.children[num]
	}
	node.isEnd = true
}

func (t *TireTree) search(str string) bool {
	if t == nil {
		return false
	}
	node := t
	for _, s := range str {
		num := s - 'a'
		if node.children[num] == nil {
			return false
		}
		node = node.children[num]
	}
	return node.isEnd == true
}
