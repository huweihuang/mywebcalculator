package data

type StackNode struct {
	Data interface{}
	next *StackNode
}

type LinkStack struct {
	top   *StackNode
	Count int
}

//栈初始化
func (this *LinkStack) Init() {
	this.top = nil
	this.Count = 0
}

//入栈
func (this *LinkStack) Push(data interface{}) {
	var node *StackNode = new(StackNode)
	node.Data = data
	node.next = this.top
	this.top = node
	this.Count++
}

//出栈
func (this *LinkStack) Pop() interface{} {
	if this.top == nil {
		return nil
	}
	returnData := this.top.Data
	this.top = this.top.next
	this.Count--
	return returnData
}

//获取栈顶元素
func (this *LinkStack) LookTop() interface{} {
	if this.top == nil {
		return nil
	}
	return this.top.Data
}
