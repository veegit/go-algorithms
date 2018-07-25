package test

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	ds "github.com/veegit/go-algorithms/datastructures"
	dp "github.com/veegit/go-algorithms/dynprog"
)

func TestMaxSum(t *testing.T) {
	array := []int{2, -3, 5, 7, -11, 13}
	a := dp.MaxSum(array)
	t.Log(a)
}

func TestEncode(t *testing.T) {
	array := dp.EncodeInt(3)
	t.Log(array)
}

func TestLinkedList(t *testing.T) {
	list := ds.NewList()
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Display()
	var n = list.Find(2)
	t.Log(n.Prev.Value)
}

func TestTrie(t *testing.T) {
	trie := ds.NewTrie()
	fd, _ := os.Open("/usr/share/dict/words")
	defer fd.Close()
	scanner := bufio.NewScanner(fd)
	for scanner.Scan() {
		trie.Insert(scanner.Text())
	}
	a := trie.Autocomplete("dinanath")
	fmt.Printf("+%v", a)
}

func TestBinaryTree(t *testing.T) {
	bt := ds.NewBinaryTree(1)
	bt.AddLeft(bt.Root, 2)
	r1 := bt.AddRight(bt.Root, 3)
	bt.AddLeft(r1, 4)
	bt.Infix()
	fmt.Printf("\n%+v", bt.Find(3))
}
