package rtech

import (
	"time"

	arraylist "github.com/emirpasic/gods/lists/arraylist"
	m "github.com/gabereiser/rtech/math"
)

type Translatable interface {
	GetPosition() m.Vector3
}
type Rotatable interface {
	GetRotation() m.Quaternion
}
type Scalable interface {
	GetScale() m.Vector3
}
type Updatable interface {
	Update(gameTime time.Duration)
}
type Renderable interface {
	Render()
}
type SceneNode interface {
	GetParent() SceneNode3D
	GetChildren() *arraylist.List
	AddChild(node SceneNode3D)
	DeleteChild(node SceneNode3D)
	Destroy()
}

type SceneNode3D interface {
	SceneNode
	Translatable
	Rotatable
	Scalable
	Updatable
	Renderable
}

type Scene interface {
	GetRootNode() SceneNode3D
	SetRootNode(node SceneNode3D)
	CreateNode() SceneNode3D
	DeleteNode(node SceneNode3D)
	Update(gameTime time.Duration)
	DestroyAll()
	Destroy()
}

type RScene struct {
	root *RSceneNode
}

func (s *RScene) GetRootNode() SceneNode3D {
	return s.root
}
func (s *RScene) SetRootNode(node *RSceneNode) {
	s.root = node
}
func (s *RScene) CreateNode() *RSceneNode {
	return &RSceneNode{
		position: m.Vector3{0, 0, 0},
		rotation: m.Quaternion{V: m.Vector3{0, 0, 0}, W: 1},
		scale:    m.Vector3{1, 1, 1},
		parent:   nil,
		children: arraylist.New(),
	}
}
func (s *RScene) DeleteNode(node SceneNode3D) {
	s.root.DeleteChild(node)
}
func (s *RScene) Update(gameTime time.Duration) {
	s.root.Update(gameTime)
}

func (s *RScene) DestroyAll() {
	s.root.Destroy()
	s.root = nil
}

type RSceneNode struct {
	position m.Vector3
	rotation m.Quaternion
	scale    m.Vector3
	parent   SceneNode3D
	children *arraylist.List
}

func (n *RSceneNode) GetPosition() m.Vector3 {
	return n.position
}
func (n *RSceneNode) GetRotation() m.Quaternion {
	return n.rotation
}
func (n *RSceneNode) GetScale() m.Vector3 {
	return n.scale
}
func (n *RSceneNode) GetParent() SceneNode3D {
	return n.parent
}
func (n *RSceneNode) GetChildren() *arraylist.List {
	return n.children
}
func (n *RSceneNode) AddChild(node SceneNode3D) {
	n.children.Add(node)
}
func (n *RSceneNode) DeleteChild(node SceneNode3D) {
	idx := n.children.IndexOf(node)
	if idx > 0 {
		n.children.Remove(idx)
	} else {
		for _, child := range n.children.Values() {
			n := child.(SceneNode3D)
			n.DeleteChild(node)
		}
	}
}
func (n *RSceneNode) Destroy() {
	for _, child := range n.children.Values() {
		n := child.(SceneNode3D)
		n.Destroy()
	}
}
func (n *RSceneNode) Update(gameTime time.Duration) {
	for _, child := range n.children.Values() {
		n := child.(SceneNode3D)
		n.Update(gameTime)
	}
}
func (n *RSceneNode) Render() {
	for _, child := range n.children.Values() {
		n := child.(SceneNode3D)
		n.Render()
	}
}
