package _41_Flatten_Nested_List_Iterator

type NestedInteger struct{}

func (this NestedInteger) IsInteger() bool {
	return false
}

func (this NestedInteger) GetInteger() int {
	return 0
}

func (n *NestedInteger) SetInteger(value int) {

}

func (this *NestedInteger) Add(elem NestedInteger) {}

func (this NestedInteger) GetList() []*NestedInteger {
	return nil
}
