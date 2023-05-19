/*
@File   : codec.go
@Author : pan
@Time   : 2023-05-19 13:22:19
*/
package main

import (
	"strconv"
	"strings"
)

func main() {

}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉树的序列化与反序列化
type Codec struct {
	Res []string
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (c *Codec) Serialize(root *TreeNode) string {
	if root == nil {
		return "#"
	}
	return strconv.Itoa(root.Val) + "," + c.Serialize(root.Left) + "," + c.Serialize(root.Right)
}

// Deserializes your encoded data to tree.
func (c *Codec) Deserialize(data string) *TreeNode {
	c.Res = strings.Split(data, ",")
	return c.DfsDeserialize()
}

func (c *Codec) DfsDeserialize() *TreeNode {
	node := c.Res[0]
	c.Res = c.Res[1:]
	if node == "#" {
		return nil
	}
	value, _ := strconv.Atoi(node)
	return &TreeNode{
		Val:   value,
		Left:  c.DfsDeserialize(),
		Right: c.DfsDeserialize(),
	}
}
