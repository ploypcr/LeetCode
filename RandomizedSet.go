package gopractice

import "math/rand"

type RandomizedSet struct {
	val []int
}

func Constructor() RandomizedSet {
	var newRandomized RandomizedSet
	return newRandomized
}

func (this *RandomizedSet) Insert(val int) bool {
	if this.val == nil {
		this.val = make([]int,0)
	}
	for _, v := range this.val {
		if val == v {
			return false
		}
	}
	this.val = append(this.val, val)
	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	for i, v := range this.val {
		if v == val {
			this.val = append(this.val[:i], this.val[i+1:]...)
			return true
		}

	}
	return false
}

func (this *RandomizedSet) GetRandom() int {
	n := rand.Intn(len(this.val))
	return this.val[n]
}