package algorithm385

import (
	"strconv"
)

/**
 * // This is the interface that allows for creating nested lists.
 * // You should not implement it, or speculate about its implementation
 * type NestedInteger struct {
 * }
 *
 * // Return true if this NestedInteger holds a single integer, rather than a nested list.
 * func (n NestedInteger) IsInteger() bool {}
 *
 * // Return the single integer that this NestedInteger holds, if it holds a single integer
 * // The result is undefined if this NestedInteger holds a nested list
 * // So before calling this method, you should have a check
 * func (n NestedInteger) GetInteger() int {}
 *
 * // Set this NestedInteger to hold a single integer.
 * func (n *NestedInteger) SetInteger(value int) {}
 *
 * // Set this NestedInteger to hold a nested list and adds a nested integer to it.
 * func (n *NestedInteger) Add(elem NestedInteger) {}
 *
 * // Return the nested list that this NestedInteger holds, if it holds a nested list
 * // The list length is zero if this NestedInteger holds a single integer
 * // You can access NestedInteger's List element directly if you want to modify it
 * func (n NestedInteger) GetList() []*NestedInteger {}
 */
type NestedInteger struct{}

func (n NestedInteger) IsInteger() bool {
	return false
}

func (n NestedInteger) GetInteger() int {
	return 0
}

func (n *NestedInteger) SetInteger(value int) {}

func (n *NestedInteger) Add(elem NestedInteger) {}

func (n NestedInteger) GetList() []*NestedInteger {
	return nil
}

func deserialize(s string) *NestedInteger {
	ni := &NestedInteger{}
	if len(s) == 0 {
		return ni
	}
	if s[0] != '[' {
		v, _ := strconv.Atoi(s)
		ni.SetInteger(v)
		return ni
	}
	s = s[1 : len(s)-1]
	for i, deep, start := 0, 0, 0; i < len(s); i++ {
		if s[i] == '[' {
			deep++
			continue
		}
		if s[i] == ']' {
			deep--
		}
		if deep == 0 {
			if s[i] == ',' {
				ni.Add(*deserialize(s[start:i]))
				start = i + 1
			} else if i == len(s)-1 {
				ni.Add(*deserialize(s[start : i+1]))
			}
		}
	}
	return ni
}
