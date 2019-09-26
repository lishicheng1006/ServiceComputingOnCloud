package main

import (
	"testing"

	. "gopkg.in/check.v1"
)

//继承testing的方法，可以直接使用go test命令运行
func Test(t *testing.T) { TestingT(t) }

//创建测试套件结构体
type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestInit(c *C) {
	input := []Node{{2}, {9}, {1}, {3}, {10}, {6}}
	wantOutput := []Node{{1}, {3}, {2}, {9}, {10}, {6}}
	Init(input)
	c.Assert(input, DeepEquals, wantOutput)
}

func (s *MySuite) TestDown(c *C) {
	input := []Node{{10}, {3}, {2}, {9}, {1}, {6}}
	wantOutput := []Node{{2}, {3}, {6}, {9}, {1}, {10}}
	down(input, 0, len(input))
	c.Assert(input, DeepEquals, wantOutput)
}

func (s *MySuite) TestUp(c *C) {
	input := []Node{{1}, {3}, {6}, {9}, {10}, {7}, {4}}
	wantOutput := []Node{{1}, {3}, {4}, {9}, {10}, {7}, {6}}
	up(input, len(input)-1)
	c.Assert(input, DeepEquals, wantOutput)
}

func (s *MySuite) TestPop(c *C) {
	input := []Node{{1}, {3}, {4}, {9}, {10}, {7}, {6}}
	wantOutput := []Node{{3}, {6}, {4}, {9}, {10}, {7}}
	wantReturn := Node{1}
	actualRes, actualNodes := Pop(input)
	c.Assert(actualRes, Equals, wantReturn)
	c.Assert(actualNodes, DeepEquals, wantOutput)
}

func (s *MySuite) TestPush(c *C) {
	input := []Node{{3}, {6}, {4}, {9}, {10}, {7}}
	wantOutput := []Node{{1}, {6}, {3}, {9}, {10}, {7}, {4}}
	c.Assert(Push(Node{1}, input), DeepEquals, wantOutput)
}

func (s *MySuite) TestRemove(c *C) {
	input := []Node{{1}, {6}, {3}, {9}, {10}, {7}, {4}}
	wantOutput := []Node{{3}, {6}, {4}, {9}, {10}, {7}}
	c.Assert(Remove(input, Node{1}), DeepEquals, wantOutput)
}
