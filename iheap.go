package ron

//import "github.com/gritzko/RON"

// FrameHeap is an iterator heap - gives the minimum available element
// at every step. Useful for merge sort like algorithms.
type FrameHeap struct {
	// Most of the time, a heap has 2 elements, optimize for that.
	// Sometimes, it can get millions of elements, ensure that is O(NlogN)
	iters               []*Frame
	primary, secondary  int
	prim_desc, sec_desc bool
}

// sort modes, e.g. PRIM_EVENT|PRIM_DESC|SEC_LOCATION
const (
	PRIM_DESC     = 4
	PRIM_TYPE     = SPEC_TYPE
	PRIM_OBJECT   = SPEC_OBJECT
	PRIM_EVENT    = SPEC_EVENT
	PRIM_LOCATION = SPEC_REF
	SEC_DESC      = 32
	SEC_TYPE      = SPEC_TYPE << 3
	SEC_OBJECT    = SPEC_OBJECT << 3
	SEC_EVENT     = SPEC_EVENT << 3
	SEC_LOCATION  = SPEC_REF << 3
)

func MakeFrameHeap(mode, size int) (ret FrameHeap) {
	ret.iters = make([]*Frame, 1, size+1)
	ret.prim_desc = (mode & PRIM_DESC) != 0
	ret.sec_desc = (mode & SEC_DESC) != 0
	ret.primary = mode & 3
	ret.secondary = (mode >> 3) & 3
	return
}

func (h FrameHeap) less(i, j int) bool {
	c := Compare(h.iters[i].uuids[h.primary], h.iters[j].uuids[h.primary])
	if c == 0 {
		c = Compare(h.iters[i].uuids[h.secondary], h.iters[j].uuids[h.secondary])
		if h.sec_desc {
			c = -c
		}
	} else if h.prim_desc {
		c = -c
	}
	//fmt.Printf("CMP %s %s %d\n", h.iters[i].String(), h.iters[j].String(), c)
	return c < 0
}

func (h *FrameHeap) sink(i int) {
	to := i
	j := i << 1
	if j < len(h.iters) && h.less(j, i) {
		to = j
	}
	j++
	if j < len(h.iters) && h.less(j, to) {
		to = j
	}
	if to != i {
		h.swap(i, to)
		h.sink(to)
	}
}

func (h *FrameHeap) raise(i int) {
	j := i >> 1
	if j > 0 && h.less(i, j) {
		h.swap(i, j)
		if j > 1 {
			h.raise(j)
		}
	}
}

func (h FrameHeap) Len() int { return len(h.iters) - 1 }

func (h FrameHeap) swap(i, j int) {
	//fmt.Printf("SWAP %d %d\n", i, j)
	h.iters[i], h.iters[j] = h.iters[j], h.iters[i]
}

func (h *FrameHeap) Put(i *Frame) {
	if !i.IsEmpty() {
		at := len(h.iters)
		h.iters = append(h.iters, i)
		h.raise(at)
	}
}

func (h *FrameHeap) Op() (op *Op) {
	if len(h.iters) > 1 {
		op = &h.iters[1].Op
	} else {
		op = &ZERO_OP
	}
	return
}

func (h *FrameHeap) remove(i int) {
	h.iters[i] = h.iters[len(h.iters)-1]
	h.iters = h.iters[:len(h.iters)-1]
	h.sink(i)
}

func (h *FrameHeap) next(i int) {
	h.iters[i].Next()
	if h.iters[i].IsEmpty() {
		h.remove(i)
	} else {
		h.sink(i)
	}
}

func (h *FrameHeap) Next() (op *Op) {
	h.next(1)
	return h.Op()
}

func (h *FrameHeap) nexteq(i int, uuid UUID) {
	if h.iters[i].uuids[h.primary] == uuid {
		j := i << 1
		if j < len(h.iters) {
			if j+1 < len(h.iters) { // FIXME rightmost first!
				h.nexteq(j+1, uuid)
			}
			h.nexteq(j, uuid)
		}
		h.next(i)
		for i < len(h.iters) && h.iters[i].uuids[h.primary] == uuid {
			h.next(i) // FIXME this fix (recheck after removal)
		}
	}
}

func (h *FrameHeap) NextPrim() (op *Op) {
	if !h.IsEmpty() {
		event := h.iters[1].uuids[h.primary]
		h.nexteq(1, event)
	}
	return h.Op()
}

func (h *FrameHeap) PutFrame(frame Frame) {
	h.Put(&frame)
}

func (h *FrameHeap) IsEmpty() bool {
	return len(h.iters) == 1
}

func (h *FrameHeap) Frame() Frame {
	cur := MakeFrame(128)
	for !h.IsEmpty() {
		cur.AppendOp(*h.Op())
		h.Next()
	}
	return cur.Close()
}

func (h *FrameHeap) Clear() {
	h.iters = h.iters[:1]
}
