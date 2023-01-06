package list

import "fmt"

type Ring struct {
	Next, Prev *Ring       //前驱和后继
	Value      interface{} //数据
}

func (r *Ring) init() *Ring {
	r.Next = r
	r.Prev = r
	return r
}

func initRing() {
	r := new(Ring)
	r.init()
}

// todo:创建 N 个节点的循环链表
func New(n int) *Ring {
	if n <= 0 {
		return nil
	}
	r := new(Ring)
	r.Value = 1
	p := r
	for i := 1; i < n; i++ {
		p.Next = &Ring{
			Prev:  p,
			Value: i + 1,
		}
		p = p.Next
	}
	p.Next = r
	r.Prev = p
	return r
}

func (r *Ring) PrintRing() {
	p := r
	fmt.Println(r.Value)
	p = r.Next
	for r != p {
		fmt.Println(p.Value)
		p = p.Next
	}
}
