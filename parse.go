package RON

import (
	"github.com/pkg/errors"
	"strconv"
)

func (op Op) ParseInt(pos int) (i int64, err error) { // FIXME no error
	if pos > op.Count || op.Types[pos] != '=' {
		err = errors.New("no int at the pos")
		return
	}
	var from, till int
	from = op.Offsets[pos] + 1
	if pos < 7 {
		till = op.Offsets[pos+1]
	} else {
		till = len(op.Body)
	}
	str := string(op.Body[from:till])
	i, err = strconv.ParseInt(str, 10, 64)
	if err != nil {
		err = errors.Wrap(err, "unparseable int atom")
	}
	return
}

func ParseAtoms (body []byte) Atoms {
	var parsed Op
	off := XParseOp(body, &parsed, ZERO_OP)
	if off <= 0 {
		off = XParseOp([]byte("'parse error'"), &parsed, ZERO_OP)
	}
	return parsed.Atoms
}

func (op Op) ParseFloat(pos int) (ret float64, err error) {
	var from, till int
	from = op.Offsets[pos] + 1 // FIXME refac
	if pos+1 < op.Count {
		till = op.Offsets[pos+1]
	} else {
		till = len(op.Body)
	}
	str := string(op.Body[from:till])
	ret, err = strconv.ParseFloat(str, 64)
	if err != nil {
		err = errors.Wrap(err, "unparseable float atom")
	}
	return
}

func ParseUUIDString(uuid string) (ret UUID, err error) {
	ret, l := ParseUUID([]byte(uuid), ZERO_UUID)
	if l <= 0 {
		err = errors.New("invalid UUID string")
	}
	return
}

func ParseUUID(data []byte, context UUID) (uuid UUID, length int) {
	uuid = context
	length = XParseUUID(data, &uuid)
	return
}

func ParseOp(data []byte, context Op) (op Op, length int) {
	op = context
	length = XParseOp(data, &op, context)
	return
}

func ParseFrame(data []byte) (ret Frame) {
	ret.Body = data
	return
}

func Parse(str string) (Frame, error) {
	ret := Frame{Body: []byte(str)}
	_ = ret.Begin() // FIXME iterator - errors
	return ret, nil
}

func UUIDSep2Sign(char byte) uint64 {
	switch char {
	case NAME_UUID_SEP:
		return NAME_SIGN
	case HASH_UUID_SEP:
		return HASH_SIGN
	case EVENT_UUID_SEP:
		return EVENT_SIGN
	case DERIVED_EVENT_SEP:
		return DERIVED_EVENT_SIGN
	default:
		panic("not an UUID separator")
	}
}
