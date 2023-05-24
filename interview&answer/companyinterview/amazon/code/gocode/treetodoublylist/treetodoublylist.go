/*
@File   : treetodoublylist.go
@Author : pan
@Time   : 2023-05-24 10:52:30
*/
package main

func main() {

}

type Node struct {
	Val         int
	Left, Right *Node
}

// 二叉搜索树与双向链表
func TreeToDoublyList(root *Node) *Node {
	if root == nil {
		return nil
	}
	var dfs func(node *Node) (head, tail *Node)
	dfs = func(node *Node) (head, tail *Node) {
		if node == nil {
			return
		}
		//递归,左子树
		lHead, lTail := dfs(node.Left)
		if lHead != nil {
			//如果左子树不为空,头结点为左子树的头节点.并拼接当前节点到左子树的尾节点
			head = lHead
			lTail.Right = node
			node.Left = lTail
		} else {
			//左子树为空,头结点为当前节点
			head = node
		}
		//递归,右子树
		rHead, rTail := dfs(node.Right)
		if rTail != nil {
			//如果右子树不为空,尾节点为右子树的尾节点.并拼接当前节点到右子树的头结点
			tail = rTail
			node.Right = rHead
			rHead.Left = node
		} else {
			//右子树为空,尾节点为当前节点
			tail = node
		}
		return
	}
	head, tail := dfs(root)
	//最后将返回的头尾节点拼接成环
	tail.Right = head
	head.Left = tail
	return head
}

/*
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var leftLast *TreeNode

//递归
func convert(root *TreeNode) *TreeNode{
   if root == nil {
      return nil
   }
   if root.Left == nil && root.Right == nil {
      leftLast = root
      return root
   }
   //左子树
   left := convert(root.Left)
   if left != nil {
      leftLast.Right = root
      root.Left = leftLast
   }
   leftLast = root
   //右子树
   right := convert(root.Right)
   if right != nil {
      right.Left = root
      root.Right = right
   }
   //?
   if left != nil {
      return left
   } else {
      return root
   }
}
*/
