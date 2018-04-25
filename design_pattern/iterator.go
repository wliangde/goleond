/**
User:       wliangde
CreateTime: 2018/4/24 下午9:49
**/
package design_pattern

/**
迭代器模式
介绍：

目的：
对于使用层，封装集合的实现。
集合底层变化时不会改动遍历代码
*/

type Book struct {
	Name  string
	Price uint32
}

type BookShelf struct {
	slcBook []*Book
}

func NewBookShelf() *BookShelf {
	return &BookShelf{
		slcBook: make([]*Book, 0),
	}
}

func (this *BookShelf) AddBook(ptBook *Book) {
	this.slcBook = append(this.slcBook, ptBook)
}

func (this *BookShelf) Iterator() Iterator {
	return NewBookIterator(this)
}

/**
Iterator
*/
type Iterator interface {
	HasNext() bool
	Next() interface{}
}

type BookIterator struct {
	dwIndex     int
	ptBookShelf *BookShelf
}

func NewBookIterator(ptBookShelf *BookShelf) *BookIterator {
	return &BookIterator{
		dwIndex:     0,
		ptBookShelf: ptBookShelf,
	}
}

func (this *BookIterator) HasNext() bool {
	if this.dwIndex >= len(this.ptBookShelf.slcBook) {
		return false
	}
	return true
}

func (this *BookIterator) Next() interface{} {
	i := this.ptBookShelf.slcBook[this.dwIndex]
	this.dwIndex++
	return i
}
