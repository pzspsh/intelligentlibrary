/*
@File   : deserializebinarytrees.go
@Author : pan
@Time   : 2023-05-23 10:14:56
*/
package main

import (
	"strconv"
	"strings"
)

func main() {
	var codec = Codec{}
	var tresnode = &TreeNode{}
	codec.serialize(tresnode)
	codec.deserialize("")
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 序列化和反序列化二叉搜索树
type Codec struct {
	res []string
}

func Constructor() Codec {
	return Codec{}
}

// 序列化
func (c *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "," + c.serialize(root.Left) + "," + c.serialize(root.Right)
}

// 反序列化
func (c *Codec) deserialize(data string) *TreeNode {
	c.res = strings.Split(data, ",")
	return c.dfsDeserialize()
}

func (c *Codec) dfsDeserialize() *TreeNode {
	node := c.res[0]
	c.res = c.res[1:]
	if node == "#" {
		return nil
	}
	value, _ := strconv.Atoi(node)
	return &TreeNode{
		Val:   value,
		Left:  c.dfsDeserialize(),
		Right: c.dfsDeserialize(),
	}
}
