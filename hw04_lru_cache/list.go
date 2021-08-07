package hw04_lru_cache //nolint:golint,stylecheck

type List interface {
	Len() int                          // длина списка
	Front() *ListItem                  // первый элемент списка
	Back() *ListItem                   // последний элемент списка
	PushFront(v interface{}) *ListItem // добавить значение в начало
	PushBack(v interface{}) *ListItem  // добавить значение в конец
	Remove(i *ListItem)                // удалить элемент
	MoveToFront(i *ListItem)           // переместить элемент в начало
}

type ListItem struct {
	Value interface{} // значение
	Next  *ListItem   // следующий элемент
	Prev  *ListItem   // предыдущий элемент
}

type list struct {
	len   int
	front *ListItem
	back  *ListItem
	//	ListItems []*ListItem
}

func NewList() List {
	newList := new(list)
	newList.len = 0
	newList.front = nil
	newList.back = nil
	//	newList.ListItems = make([]*ListItem, 1)
	return newList
}

func (list *list) Len() int {
	return list.len
}

func (list *list) Front() *ListItem {
	return list.front
}

func (list *list) Back() *ListItem {
	return list.back
}

func (list *list) PushFront(v interface{}) *ListItem {
	if list.len == 0 || list.front == nil {
		list.front = &ListItem{Value: v, Next: nil, Prev: nil}
		list.back = list.front
		//fmt.Println("This is 1st list.front:", list.front, "	This his prev and next:", list.front.Prev, list.front.Next)

	} else {
		//fmt.Println("Prepare new list.front")
		//fmt.Println("This was list.front:", list.front, "	This his prev and next:", list.front.Prev, list.front.Next)
		newFront := &ListItem{Value: v, Next: list.front, Prev: nil}
		//fmt.Println("This is gonna be new list.front:", newFront, "	This his prev and next:", newFront.Prev, newFront.Next)
		list.front.Prev = newFront
		list.front = newFront
		//fmt.Println("This is new list.front:", list.front, "	This his prev and next:", list.front.Prev, list.front.Next)
		//fmt.Println("This is old list.front:", list.front.Next, "	This his prev and next:", list.front.Next.Prev, list.front.Next.Next)
	}
	list.len++
	return list.front
}

func (list *list) PushBack(v interface{}) *ListItem {
	if list.len == 0 || list.back == nil {
		list.front = &ListItem{Value: v, Next: nil, Prev: nil}
		list.back = list.front
	} else {
		newBack := &ListItem{Value: v, Next: nil, Prev: list.back}
		list.back.Next = newBack
		list.back = newBack
	}
	list.len++
	return list.back
}

func (list *list) Remove(i *ListItem) {
	if nil == i {
		return
	} else {
		if i.Next != nil { //если i не был концом
			i.Next.Prev = i.Prev
		} else {
			list.back = i.Prev
		}
		if i.Prev != nil { //если i не был началом
			i.Prev.Next = i.Next
		} else {
			list.front = i.Next
		}
		i.Next = nil
		i.Prev = nil
		list.len--
	}

}

func (list *list) MoveToFront(i *ListItem) {
	if (i.Next == nil && i.Prev == nil) || (nil == i) || list.front == i {
		return
	} else {
		list.Remove(i)
		list.PushFront(i.Value)
	}
}
