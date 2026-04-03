// ===----------------------------------------------------------------------===//
//
//                         BusTub
//
// trie.go
//
// Identification: primer/trie/trie.go
//
// Copyright (c) 2015-2025, Carnegie Mellon University Database Group
//
// ===----------------------------------------------------------------------===//

// Package trie implements a copy-on-write persistent trie data structure.
// All operations on a Trie do NOT modify the trie itself. They reuse existing
// nodes as much as possible and create new nodes to represent the new trie.
package trie

// TrieNode is a node in the Trie.
type TrieNode struct {
	// children maps the next character in the key to the next TrieNode.
	// You MUST store children information in this structure.
	children map[byte]*TrieNode

	// isValueNode indicates if this node is the terminal node for some key.
	isValueNode bool

	// value holds the value if this is a value node. The concrete type is
	// stored as interface{} to support generics (since Go generics differ from
	// C++ templates in this educational context we keep it flexible).
	value interface{}
}

// NewTrieNode creates a TrieNode with no children.
func NewTrieNode() *TrieNode {
	return &TrieNode{
		children: make(map[byte]*TrieNode),
	}
}

// Clone returns a shallow copy of this TrieNode (children map is copied by
// reference-per-entry, which is safe because we never mutate existing nodes).
func (n *TrieNode) Clone() *TrieNode {
	newChildren := make(map[byte]*TrieNode, len(n.children))
	for k, v := range n.children {
		newChildren[k] = v
	}
	return &TrieNode{
		children:    newChildren,
		isValueNode: n.isValueNode,
		value:       n.value,
	}
}

// Trie is a data structure that maps strings to values. All operations on a
// Trie should NOT modify the trie itself. It should reuse existing nodes as
// much as possible and create new nodes to represent the new trie.
//
// You are NOT allowed to use mutable state to bypass the immutability contract.
type Trie struct {
	// root is the root of the trie. nil means an empty trie.
	root *TrieNode
}

// NewTrie creates an empty Trie.
func NewTrie() *Trie {
	return &Trie{}
}

// GetRoot returns the root of the trie. Should only be used in tests.
func (t *Trie) GetRoot() *TrieNode {
	return t.root
}

// Get returns the value associated with the given key, or (nil, false) if:
//  1. The key is not in the trie.
//  2. The key is in the trie but the stored value's type does not match T.
//
// TODO(P0): Add implementation.
//
// You should walk through the trie to find the node corresponding to the key.
// If the node doesn't exist, return nil, false.
// After you find the node, you should use a type assertion to cast it to T.
// If the assertion fails, return nil, false.
// Otherwise, return the value and true.
func Get[T any](t *Trie, key string) (*T, bool) {
	// TODO(P0): Trie.Get is not implemented.
	// Stub so the package compiles; return "not found".
	return nil, false
}

// Put inserts a new key-value pair into the trie. If the key already exists,
// the value is overwritten. Returns the new trie (the original is unchanged).
//
// Note: T might be a non-copyable type. Always prefer moving/pointers.
//
// TODO(P0): Add implementation.
//
// You should walk through the trie and create new nodes if necessary.
// If the node corresponding to the key already exists, you should create a new
// TrieNode that is a value node with the new value. Make sure to clone nodes
// on the path from root to the target node so you don't modify the original.
func Put[T any](t *Trie, key string, value T) *Trie {
	// TODO(P0): Trie.Put is not implemented.
	// Stub: keep immutability by returning the original trie unchanged.
	// (When you implement, you should return a new trie.)
	return t
}

// Remove removes the key from the trie. If the key does not exist, returns the
// original trie unchanged. Otherwise, returns the new trie.
//
// TODO(P0): Add implementation.
//
// You should walk through the trie and remove nodes if necessary.
// If the node doesn't contain a value anymore, convert it to a plain TrieNode.
// If a node has no children anymore, remove it entirely.
func Remove(t *Trie, key string) *Trie {
	// TODO(P0): Trie.Remove is not implemented.
	// Stub: return the original trie unchanged.
	return t
}
