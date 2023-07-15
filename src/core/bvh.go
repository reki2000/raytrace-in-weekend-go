package core

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

type BvhNode struct {
	left, right Object
	box         *Aabb
}

func createBoxComparator(axis int, time0, time1 double) func(a, b Object) int {
	return func(a, b Object) int {
		log("comparing %v %v", a, b)
		okA, boxA := a.BoundingBox(time0, time1)
		okB, boxB := b.BoundingBox(time0, time1)
		if !okA || !okB {
			panic("no bounding box in bvh_node constructor")
		}

		diff := boxA.Min.E(axis) - boxB.Min.E(axis)
		if diff < 0.0 {
			return -1
		} else if diff > 0.0 {
			return 1
		}
		return 0
	}
}

func NewBvhNode(objects ObjectList, time0, time1 double) *BvhNode {
	log("NewBvhNode %v", objects)

	axis := rand.Int() % 3
	comparator := createBoxComparator(axis, time0, time1)

	var left Object
	var right Object

	span := len(objects)
	if span == 1 {
		left = objects[0]
		right = objects[0]
	} else if span == 2 {
		if comparator(objects[0], objects[1]) < 0 {
			left = objects[0]
			right = objects[1]
		} else {
			left = objects[1]
			right = objects[0]
		}
	} else {
		sort.Slice(objects, func(i, j int) bool { return comparator(objects[i], objects[j]) < 0 })

		mid := span / 2
		left = NewBvhNode(objects[:mid], time0, time1)
		right = NewBvhNode(objects[mid:], time0, time1)
	}

	leftOk, boxLeft := left.BoundingBox(time0, time1)
	rightOk, boxRight := right.BoundingBox(time0, time1)
	if !leftOk || !rightOk {
		panic("no bounding box in bvh_node constructor")
	}

	box := NewSurroundingBox(boxLeft, boxRight)
	return &BvhNode{left: left, right: right, box: box}
}

func (b *BvhNode) Hit(r *Ray, tMin, tMax double) (bool, *HitRecord) {
	if !b.box.Hit(r, tMin, tMax) {
		return false, nil
	}

	hitLeft, hrLeft := b.left.Hit(r, tMin, tMax)
	if hitLeft {
		tMax = hrLeft.T
	}

	if hitRight, hrRight := b.right.Hit(r, tMin, tMax); hitRight {
		return hitRight, hrRight
	}

	return hitLeft, hrLeft
}

func (b *BvhNode) BoundingBox(t0, t1 double) (bool, *Aabb) {
	return true, b.box
}

func (b *BvhNode) String() string {
	return b.Show(0)
}

func (b *BvhNode) Show(indent int) string {
	buf := fmt.Sprintf("%s (%0.2f,%0.2f,%0.2f)-(%0.2f,%0.2f,%0.2f)\n",
		strings.Repeat(" ", indent),
		b.box.Min.X, b.box.Min.Y, b.box.Min.Z,
		b.box.Max.X, b.box.Max.Y, b.box.Max.Z)
	if node, ok := b.left.(*BvhNode); ok {
		buf += node.Show(indent + 1)
	}
	if node, ok := b.right.(*BvhNode); ok {
		buf += node.Show(indent + 1)
	}
	return buf
}
