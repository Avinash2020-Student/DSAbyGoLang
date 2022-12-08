package Trees

import("fmt")

//data is given as string
type Treenode struct{
	data string
	children make([]string, 3)
}

func (t Treenode) str(level=0 int) string{
	ret := " " * level + data + "\n"
	for child in t.children{
		ret += child.str(level+1)
	}
	return ret
}

func (t Treenode) addchild(Treenode){
	t.children.append(Treenode)
}

//create root nodea nd  children and print out the tree
func basictree(){
	tree := Treenode("rootnode", {}) //tree name ==> tree
	leftNodeOne := Treenode("leftNodeOne", {})
	rightNodeOne := Treenode("rightNodeOne", {})
	tree.addchild(leftNodeOne)
	tree.addchild(rightNodeOne)
	fmt.Println(tree)
}