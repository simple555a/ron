package RON

func unzipPrefixSeparator(input []byte) (prefix uint8, length int) {
	var i = ABC[input[0]]
	if i <= -10 {
		prefix = uint8(-i-10) + 4
		length = 1
	}
	return
}

func UnzipBase64(input []byte, number *uint64) int {

	// TODO migrate to Ragel

	var l int = 10
	if l > len(input) {
		l = len(input)
	}
	var i = 0
	var res uint64
	for ; i < l; i++ {
		code := ABC[input[i]]
		if code < 0 {
			break
		}
		res <<= 6
		res |= uint64(code)
	}
	if number != nil {
		*number = res
	}
	return i
}

func (t UUID) Equal(b UUID) bool {
	return t.Value == b.Value && t.Origin == b.Origin
}

func CommonPrefix(value, context uint64) uint {
	// TODO use math.bits
	var xor = value ^ context
	if xor >= 1<<(6*6) {
		return 0
	}
	if xor == 0 {
		return 10
	}
	if xor >= 1<<(3*6) { // 456
		if xor >= 1<<(5*6) {
			return 4
		} else if xor >= 1<<(4*6) {
			return 5
		} else {
			return 6
		}
	} else { // 789
		if xor >= 1<<(2*6) {
			return 7
		} else if xor >= 1<<(1*6) {
			return 8
		} else {
			return 9
		}
	}
}

func ZeroTail(value *uint64) (tail uint) {
	if *value&((1<<30)-1) == 0 {
		tail += 5
		*value >>= 30
	}
	if *value&((1<<18)-1) == 0 {
		tail += 3
		*value >>= 18
	}
	if *value&((1<<12)-1) == 0 {
		tail += 2
		*value >>= 12
	}
	if tail < 10 && *value&((1<<6)-1) == 0 {
		tail += 1
		*value >>= 6
	}
	return
}

const prefix_mask uint64 = 0xffffff << 36

func ZipUUIDString(uuid, context UUID) string {
	var ret = make([]byte, 21, 21)
	len := FormatZippedUUID(ret, uuid, context)
	return string(ret[0:len])
}

func (uuid UUID) String() string {
	return ZipUUIDString(uuid, ZERO_UUID)
}

func (a UUID) LessThan(b UUID) bool {
	if a.Value == b.Value {
		if a.Sign == b.Sign {
			return a.Origin < b.Origin
		} else {
			return a.Sign < b.Sign
		}
	} else {
		return a.Value < b.Value
	}
}

func FormatTrimmedInt(output []byte, value uint64) int {
	if value == 0 {
		output[0] = '0'
		return 1
	}
	l := 10
	if value&((1<<24)-1) == 0 {
		value >>= 24
		l -= 4
	}
	for value&63 == 0 {
		value >>= 6
		l--
	}
	k := l
	for k > 0 {
		k--
		output[k] = base64[value&63]
		value >>= 6
	}
	return l
}

func FormatInt(output []byte, value uint64) int {
	l := FormatTrimmedInt(output, value)
	for l < 10 {
		output[l] = '0'
		l++
	}
	return l
}

func FormatZippedInt(output []byte, value, context uint64) int {
	var prefix uint = CommonPrefix(value, context)
	var off int
	if prefix < 4 {
		off += FormatTrimmedInt(output[off:], value)
	} else {
		if prefix == 10 {
			return 0
		}
		output[0] = PREFIX_PUNCT[prefix-4]
		off++
		value = (value << (prefix * 6)) & PREFIX10
		if value != 0 {
			off += FormatTrimmedInt(output[off:], value)
		}
	}
	return off
}

func FormatUUID(output []byte, uuid UUID) int {
	l := FormatTrimmedInt(output, uuid.Value)
	output[l] = uuid.Sign
	l++
	l += FormatTrimmedInt(output[l:], uuid.Origin)
	return l
}

func FormatZippedUUID(output []byte, uuid UUID, context UUID) int {

	if uuid == context && uuid != ZERO_UUID { // FIXME options
		return 0
	}
	off := FormatZippedInt(output, uuid.Value, context.Value)
	if uuid.Sign == NAME_UUID_SEP && uuid.Origin == 0 {
		return off
	}
	if uuid.Value == context.Value || uuid.Sign != context.Sign ||
		(uuid.Origin&prefix_mask) != (context.Origin&prefix_mask) ||
		(uuid.Origin == context.Origin && ABC[output[0]]>=0) { // FIXME this if
		output[off] = uuid.Sign
		off++
	}
	if uuid.Origin != context.Origin {
		off += FormatZippedInt(output[off:], uuid.Origin, context.Origin)
	}
	return off
}

func FormatSpec(output []byte, op *Op) int {
	var off int
	// expand to 88+values
	for t := 0; t < 4; t++ {
		output[off] = SPEC_PUNCT[t]
		off++
		off += FormatUUID(output[off:], op.GetUUID(t))
	}
	return off
}

func FormatZippedSpec(output []byte, op *Op, context *Op) int {
	var off int
	// expand to 88+values
	for t := 0; t < 4; t++ {
		if op.GetUUID(t) == context.GetUUID(t) {
			continue
		}
		output[off] = SPEC_PUNCT[t]
		off++
		off += FormatZippedUUID(output[off:], op.GetUUID(t), context.GetUUID(t))
	}
	return off
}

// optimize for close values
// context==nil is valid
func FormatOp(output []byte, op *Op, context *Op) int {
	off := FormatZippedSpec(output, op, context)
	from := op.AtomOffsets[0]
	copy(output[off:], op.Body[from:])
	off += len(op.Body) - from
	return off
}

func (op *Op) String() string {
	buf := make([]byte, op.AtomOffsets[op.AtomCount-1]+100) // FIXME!!!
	l := FormatOp(buf, op, &ZERO_OP)
	return string(buf[:l])
}

func (frame *Frame) String() string {
	return string(frame.Body)
}

func (frame *Frame) AppendOp(op *Op) {
	var l int
	var uuids [11 * 2 * 4]byte
	if !frame.last.isZero() || len(frame.Body)==0 {
		l = FormatZippedSpec(uuids[:], op, &frame.last)
	} else {
		l = FormatSpec(uuids[:], op)
	}
	frame.Body = append(frame.Body, uuids[:l]...)
	frame.Body = append(frame.Body, op.Body[op.AtomOffsets[0]:]...)
	frame.last = *op
}

func (frame *Frame) Append(t, o, e, l UUID, atoms []byte) {
	var parsed Op
	off := XParseOp(atoms, &parsed, ZERO_OP)
	if off <= 0 {
		off = XParseOp([]byte("'parse error'"), &parsed, ZERO_OP)
	}
	parsed.uuids = [4]UUID{t,o,e,l}
	frame.AppendOp(&parsed)
}

func (frame *Frame) AppendRange(i, j Iterator) {
	if i.offset >= j.offset {
		return
	}
	frame.AppendOp(&i.Op)
	// if error => plus1 is closed
	frame.Body = append(frame.Body, j.Rest()...)
	frame.last = j.Op
}

func (frame *Frame) AppendAll(i Iterator) {
	frame.AppendRange(i, i.frame.End())
}

func (frame *Frame) AppendFrame(second Frame) {
	frame.AppendRange(second.Begin(), second.End())
}

//func (frame *Frame) AppendEnd() {
// no explicit end marker for now
//}

func (frame *Frame) AppendError(comment string) {

}

func (frame *Frame) Clone() Frame { // TODO size hint
	return Frame{}
}
