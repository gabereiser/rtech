package rtech

import (
	"time"

	arraylist "github.com/emirpasic/gods/lists/arraylist"
	mgl "github.com/go-gl/mathgl/mgl64"
)

type Translatable interface {
	GetPosition() mgl.Vec3
}
type Rotatable interface {
	GetRotation() mgl.Quat
}
type Scalable interface {
	GetScale() mgl.Vec3
}
type Updatable interface {
	Update(gameTime time.Duration)
}
type Renderable interface {
	Render()
}
type SceneNode interface {
	GetParent() SceneNode
	GetChildren() *arraylist.List
	AddChild(node SceneNode)
	DeleteChild(node SceneNode)
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
	GetRootNode() SceneNode
	SetRootNode(node SceneNode)
	CreateNode() SceneNode
	DeleteNode(node SceneNode)
	Update()
	DestroyAll()
	Destroy()
}

type RScene struct {
	root *RSceneNode
}

func (s *RScene) GetRootNode() SceneNode {
	return s.root
}
func (s *RScene) SetRootNode(node *RSceneNode) {
	s.root = node
}
func (s *RScene) CreateNode() *RSceneNode {
	return &RSceneNode{
		position: mgl.Vec3{0, 0, 0},
		rotation: mgl.Quat{V: mgl.Vec3{0, 0, 0}, W: 1},
		scale:    mgl.Vec3{1, 1, 1},
		parent:   nil,
		children: arraylist.New(),
	}
}
func (s *RScene) DeleteNode(node SceneNode) {
	s.root.DeleteChild(node)
}
func (s *RScene) Update() {

}

func (s *RScene) DestroyAll() {
	s.root.Destroy()
	s.root = nil
}

type RSceneNode struct {
	position mgl.Vec3
	rotation mgl.Quat
	scale    mgl.Vec3
	parent   SceneNode
	children *arraylist.List
}

func (n *RSceneNode) GetPosition() mgl.Vec3 {
	return n.position
}
func (n *RSceneNode) GetRotation() mgl.Quat {
	return n.rotation
}
func (n *RSceneNode) GetScale() mgl.Vec3 {
	return n.scale
}
func (n *RSceneNode) GetParent() SceneNode {
	return n.parent
}
func (n *RSceneNode) GetChildren() *arraylist.List {
	return n.children
}
func (n *RSceneNode) AddChild(node SceneNode) {
	n.children.Add(node)
}
func (n *RSceneNode) DeleteChild(node SceneNode) {
	idx := n.children.IndexOf(node)
	if idx > 0 {
		n.children.Remove(idx)
	} else {
		for _, child := range n.children.Values() {
			n := child.(SceneNode)
			n.DeleteChild(node)
		}
	}
}
func (n *RSceneNode) Destroy() {
	for _, child := range n.children.Values() {
		n := child.(SceneNode)
		n.Destroy()
	}
}
