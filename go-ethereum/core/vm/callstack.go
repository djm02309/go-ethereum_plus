package vm

import (
	"fmt"
	//"github.com/ethereum/go-ethereum/common"
)


// FunctionNode represents node use for stack
type FunctionNode struct {
	FunctionName    string
	ContractName    string
	ContractAddress []byte
	FunctionOrder    string
	FunctionID       []byte
}

// Request represents printing method
func (n *FunctionNode) String() string {
	return fmt.Sprint(n.FunctionName, n.ContractName, string(n.ContractAddress), n.FunctionOrder, string(n.FunctionID))
}

// NewCallStack represents configuring new stack
func NewCallStack() *CallStack {
	return &CallStack{}
}

// CallStack represents structure for stack
type CallStack struct {
	functions []*FunctionNode
	count     int
}

// Push represents put nodes in stack
func (s *CallStack) Push(n *FunctionNode) {
	s.functions = append(s.functions[:s.count], n)
	s.count++
	fmt.Printf("In Stack function signature: %x\n",n.FunctionID);
	//fmt.Printf("contract address: %s\n",n.ContractAddress);
}

// Pop represents remove nodes in stack
func (s *CallStack) Pop() *FunctionNode {
	if s.count == 0 {
		fmt.Println("nothing to pop")
		return nil
	}
	s.count--

	fmt.Printf("Stack pop-> function signature: %x\n",s.functions[s.count].FunctionID);
	//fmt.Printf("Contract Address: %s\n",s.functions[s.count].ContractAddress);
	return s.functions[s.count]
}
