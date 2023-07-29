package core

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

type bvhNode struct {
	left, right Object
	box         *aabb
}

func createBoxComparator(axis int, time0, time1 Double) func(a, b Object) int {
	return func(a, b Object) int {
		okA, boxA := a.boundingBox(time0, time1)
		okB, boxB := b.boundingBox(time0, time1)
		if !okA || !okB {
			panic("no bounding box in bvh_node constructor")
		}

		debug("comparing %v %v", boxA.min, boxB.min)
		diff := boxA.min.e(axis) - boxB.min.e(axis)
		if diff < 0.0 {
			return -1
		} else if diff > 0.0 {
			return 1
		}
		return 0
	}
}

func NewBvhNode(objects ObjectList, time0, time1 Double) *bvhNode {
	debug("NewBvhNode %v", objects)

	axis := rand.Int() % 3
	comparator := createBoxComparator(axis, time0, time1)

	var left Object
	var right Object

	span := len(objects)
	if span == 0 {
		panic("span 0")
	} else if span == 1 {
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
		// log := fmt.Sprintf("axis: %d ", axis)
		// for _, v := range objects {
		// 	_, box := v.BoundingBox(time0, time1)
		// 	log += fmt.Sprintf(" %0.2f,%0.2f,%0.2f ", box.Min.x, box.Min.y, box.Min.z)
		// }
		// info(log)

		mid := span / 2
		left = NewBvhNode(objects[:mid], time0, time1)
		right = NewBvhNode(objects[mid:], time0, time1)
	}

	leftOk, boxLeft := left.boundingBox(time0, time1)
	rightOk, boxRight := right.boundingBox(time0, time1)
	if !leftOk || !rightOk {
		panic("no bounding box in bvh_node constructor")
	}

	box := newSurroundingBox(boxLeft, boxRight)
	return &bvhNode{left: left, right: right, box: box}
}

func (b *bvhNode) hit(r *Ray, tMin, tMax Double) (bool, *hitRecord) {
	if !b.box.hit(r, tMin, tMax) {
		return false, nil
	}

	hitLeft, hrLeft := b.left.hit(r, tMin, tMax)
	if hitLeft {
		tMax = hrLeft.t
	}

	if hitRight, hrRight := b.right.hit(r, tMin, tMax); hitRight {
		return hitRight, hrRight
	}

	return hitLeft, hrLeft
}

func (b *bvhNode) boundingBox(t0, t1 Double) (bool, *aabb) {
	return true, b.box
}

func (b *bvhNode) String() string {
	return b.Show(0)
}

func (b *bvhNode) Show(indent int) string {
	buf := fmt.Sprintf("%s [[%0.2f,%0.2f,%0.2f], [%0.2f,%0.2f,%0.2f]],\n",
		strings.Repeat(" ", indent),
		b.box.min.x, b.box.min.y, b.box.min.z,
		b.box.max.x, b.box.max.y, b.box.max.z)
	if node, ok := b.left.(*bvhNode); ok {
		buf += node.Show(indent + 1)
	}
	if node, ok := b.right.(*bvhNode); ok {
		buf += node.Show(indent + 1)
	}
	return buf
}
