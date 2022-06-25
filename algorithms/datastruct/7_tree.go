package datastruct

import (
	"fmt"
)

/**
除了链表存储外，还有
*/

type node4 struct {
	v     int
	left  *node4
	right *node4
}

func (t *node4) find(data int) bool {
	for t != nil {
		if t.v == data {
			return true
		}
		if t.v > data {
			if t.left == nil {
				return false
			}
			t = t.left
			continue
		}
		if t.v < data {
			if t.right == nil {
				return false
			}
			t = t.right
			continue
		}
	}
	return false
}

func (t *node4) insert(data int) {
	//如果存在如何插入?
	for t != nil {
		if t.v > data {
			if t.left == nil {
				t.left = &node4{v: data}
				return
			}
			t = t.left
			continue
		}
		if t.v < data {
			if t.right == nil {
				t.right = &node4{v: data}
				return
			}
			t = t.right
			continue
		}
	}
}
func (t *node4) delete(data int) bool {
	//第一种情况是，如果要删除的节点没有子节点，我们只需要直接将父节点中，指向要删除节点的指针置为 null。比如图中的删除节点 55。
	//
	// 第二种情况是，如果要删除的节点只有一个子节点（只有左子节点或者右子节点），我们只需要更新父节点中，指向要删除节点的指针，让它指向要删除节点的子节点就可以了。比如图中的删除节点 13。
	//
	// 第三种情况是，如果要删除的节点有两个子节点，这就比较复杂了。我们需要找到这个节点的右子树中的最小节点，把它替换到要删除的节点上。然后再删除掉这个最小节点，因为最小节点肯定没有左子节点（如果有左子结点，那就不是最小节点了），所以，我们可以应用上面两条规则来删除这个最小节点。比如图中的删除节点 18。
	//1 节点存在
	p := t
	var pp *node4 //要删除的节点的父节点
	for p != nil {
		if p.v == data {
			break
		}
		if p.v > data {
			pp = p
			p = p.left
			continue
		}
		if p.v < data {
			pp = p
			p = p.right
			continue
		}
	}
	if p == nil {
		return false
	}
	//2 删除的节点有两个子节点
	if p.right != nil && p.left != nil {
		minP := p.right
		minPP := p
		for minP.left != nil {
			minPP = minP
			minP = minP.left
		}
		p.v = minP.v //赋值，而不是真的删除了
		p = minP     //下面就变成了删除minP，此时就变成了第3步：删除的节点是叶子节点或者仅有一个子节点
		pp = minPP
	}

	//3 删除的节点是叶子节点或者仅有一个子节点

	var child *node4
	if p.left != nil { //获取子节点，也就是说子节点下的所有点了
		child = p.left
	} else if p.right != nil {
		child = p.right
	}
	//这里将子节点放到父节点下面（其中会做一些判断是放到左边还是右边）
	if pp == nil { //这里其实不会是空的。
		t = child
	} else if pp.left == p {
		pp.left = child
	} else {
		pp.right = child
	}
	return true
}

func (t *node4) inOrderTraversal() {
	if t == nil {
		return
	}
	t.left.inOrderTraversal()
	fmt.Println(t.v)
	t.right.inOrderTraversal()
}

//root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1
/**
如果节点 X 存储在数组中下标为 i 的位置，下标为 2 * i 的位置存储的就是左子节点，
	下标为 2 * i + 1 的位置存储的就是右子节点。反过来，下标为 i/2 的位置存储就是它的父节点。
	通过这种方式，我们只要知道根节点存储的位置（一般情况下，为了方便计算子节点，根节点会存储在下标为 1 的位置），
	这样就可以通过下标计算，把整棵树都串起来。
*/
