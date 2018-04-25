/**
User:       wliangde
CreateTime: 2018/4/25 下午9:29
**/
package design_pattern

import "testing"

func TestIterator(t *testing.T) {
	ptBookShelf := NewBookShelf()
	ptBookShelf.AddBook(&Book{"c++", 20})
	ptBookShelf.AddBook(&Book{"golang", 21})
	ptBookShelf.AddBook(&Book{"php", 22})

	it := ptBookShelf.Iterator()

	for it.HasNext() {
		ptBook := it.Next().(*Book)
		t.Log(ptBook)
	}
}
