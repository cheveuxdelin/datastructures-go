package bst

type node[T any] struct {
	value T
	left  *node[T]
	right *node[T]
}

type BST[T any] struct {
	Compare func(n1 T, n2 T) bool
	Root    *node[T]
	Size    int
}

func (b *BST[T]) Push(x T) {
	var newNode *node[T] = &node[T]{x, nil, nil}
	if b.Size == 0 {
		b.Root = newNode
	} else {
		var aux *node[T] = b.Root

		for {
			if b.Compare(aux.value, newNode.value) {
				if aux.left == nil {
					aux.left = newNode
					break
				} else {
					aux = aux.left
				}
			} else {
				if aux.right == nil {
					aux.right = newNode
					break
				} else {
					aux = aux.right
				}
			}
		}
	}
	b.Size++
}
