package RON

// bit-separator conversions generated by ./sep2bits.pl on
// Thu Oct  5 01:11:45 +05 2017
// from
// commit 8ae1ed4a88861196211a70cc8d8db602455d23d3
// Author: Victor Grishchenko <victor.grishchenko@gmail.com>
// Date:   Wed Oct 4 17:11:44 2017 +0500
//
//     big spring cleaning - uuids are back

var PUNCT, BITS [128]int8

const SPEC_PUNCT = "*#@:"
const (
	SPEC_TYPE_SEP   = '*'
	SPEC_OBJECT_SEP = '#'
	SPEC_EVENT_SEP  = '@'
	SPEC_REF_SEP    = ':'
)
const (
	SPEC_TYPE   = 0
	SPEC_OBJECT = 1
	SPEC_EVENT  = 2
	SPEC_REF    = 3
)

func specSep2Bits(sep byte) uint {
	switch sep {
	case SPEC_TYPE_SEP:
		return SPEC_TYPE
	case SPEC_OBJECT_SEP:
		return SPEC_OBJECT
	case SPEC_EVENT_SEP:
		return SPEC_EVENT
	case SPEC_REF_SEP:
		return SPEC_REF
	default:
		panic("invalid spec separator")
	}
}
func specBits2Sep(bits uint) byte {
	switch bits {
	case SPEC_TYPE:
		return SPEC_TYPE_SEP
	case SPEC_OBJECT:
		return SPEC_OBJECT_SEP
	case SPEC_EVENT:
		return SPEC_EVENT_SEP
	case SPEC_REF:
		return SPEC_REF_SEP
	default:
		panic("invalid spec bits")
	}
}

const UUID_PUNCT = "$%+-"
const (
	UUID_NAME_SEP    = '$'
	UUID_HASH_SEP    = '%'
	UUID_EVENT_SEP   = '+'
	UUID_DERIVED_SEP = '-'
)
const (
	UUID_NAME    = 0
	UUID_HASH    = 1
	UUID_EVENT   = 2
	UUID_DERIVED = 3
)

func uuidSep2Bits(sep byte) uint {
	switch sep {
	case UUID_NAME_SEP:
		return UUID_NAME
	case UUID_HASH_SEP:
		return UUID_HASH
	case UUID_EVENT_SEP:
		return UUID_EVENT
	case UUID_DERIVED_SEP:
		return UUID_DERIVED
	default:
		panic("invalid uuid separator")
	}
}
func uuidBits2Sep(bits uint) byte {
	switch bits {
	case UUID_NAME:
		return UUID_NAME_SEP
	case UUID_HASH:
		return UUID_HASH_SEP
	case UUID_EVENT:
		return UUID_EVENT_SEP
	case UUID_DERIVED:
		return UUID_DERIVED_SEP
	default:
		panic("invalid uuid bits")
	}
}

const ATOM_PUNCT = ">='^"
const (
	ATOM_UUID_SEP   = '>'
	ATOM_INT_SEP    = '='
	ATOM_STRING_SEP = '\''
	ATOM_FLOAT_SEP  = '^'
)
const (
	ATOM_UUID   = 0
	ATOM_INT    = 1
	ATOM_STRING = 2
	ATOM_FLOAT  = 3
)

func atomSep2Bits(sep byte) uint {
	switch sep {
	case ATOM_UUID_SEP:
		return ATOM_UUID
	case ATOM_INT_SEP:
		return ATOM_INT
	case ATOM_STRING_SEP:
		return ATOM_STRING
	case ATOM_FLOAT_SEP:
		return ATOM_FLOAT
	default:
		panic("invalid atom separator")
	}
}
func atomBits2Sep(bits uint) byte {
	switch bits {
	case ATOM_UUID:
		return ATOM_UUID_SEP
	case ATOM_INT:
		return ATOM_INT_SEP
	case ATOM_STRING:
		return ATOM_STRING_SEP
	case ATOM_FLOAT:
		return ATOM_FLOAT_SEP
	default:
		panic("invalid atom bits")
	}
}

const TERM_PUNCT = ";,!?"
const (
	TERM_RAW_SEP     = ';'
	TERM_REDUCED_SEP = ','
	TERM_HEADER_SEP  = '!'
	TERM_QUERY_SEP   = '?'
)
const (
	TERM_RAW     = 0
	TERM_REDUCED = 1
	TERM_HEADER  = 2
	TERM_QUERY   = 3
)

func termSep2Bits(sep byte) uint {
	switch sep {
	case TERM_RAW_SEP:
		return TERM_RAW
	case TERM_REDUCED_SEP:
		return TERM_REDUCED
	case TERM_HEADER_SEP:
		return TERM_HEADER
	case TERM_QUERY_SEP:
		return TERM_QUERY
	default:
		panic("invalid term separator")
	}
}
func termBits2Sep(bits uint) byte {
	switch bits {
	case TERM_RAW:
		return TERM_RAW_SEP
	case TERM_REDUCED:
		return TERM_REDUCED_SEP
	case TERM_HEADER:
		return TERM_HEADER_SEP
	case TERM_QUERY:
		return TERM_QUERY_SEP
	default:
		panic("invalid term bits")
	}
}

const REDEF_PUNCT = "`\\|/"
const (
	REDEF_PREV_SEP   = '`'
	REDEF_OBJECT_SEP = '\\'
	REDEF_EVENT_SEP  = '|'
	REDEF_REF_SEP    = '/'
)
const (
	REDEF_PREV   = 0
	REDEF_OBJECT = 1
	REDEF_EVENT  = 2
	REDEF_REF    = 3
)

func redefSep2Bits(sep byte) uint {
	switch sep {
	case REDEF_PREV_SEP:
		return REDEF_PREV
	case REDEF_OBJECT_SEP:
		return REDEF_OBJECT
	case REDEF_EVENT_SEP:
		return REDEF_EVENT
	case REDEF_REF_SEP:
		return REDEF_REF
	default:
		panic("invalid redef separator")
	}
}
func redefBits2Sep(bits uint) byte {
	switch bits {
	case REDEF_PREV:
		return REDEF_PREV_SEP
	case REDEF_OBJECT:
		return REDEF_OBJECT_SEP
	case REDEF_EVENT:
		return REDEF_EVENT_SEP
	case REDEF_REF:
		return REDEF_REF_SEP
	default:
		panic("invalid redef bits")
	}
}

const PREFIX_PUNCT = "([{}])"
const (
	PREFIX_PRE4_SEP = '('
	PREFIX_PRE5_SEP = '['
	PREFIX_PRE6_SEP = '{'
	PREFIX_PRE7_SEP = '}'
	PREFIX_PRE8_SEP = ']'
	PREFIX_PRE9_SEP = ')'
)
const (
	PREFIX_PRE4 = 0
	PREFIX_PRE5 = 1
	PREFIX_PRE6 = 2
	PREFIX_PRE7 = 3
	PREFIX_PRE8 = 4
	PREFIX_PRE9 = 5
)

func prefixSep2Bits(sep byte) uint {
	switch sep {
	case PREFIX_PRE4_SEP:
		return PREFIX_PRE4
	case PREFIX_PRE5_SEP:
		return PREFIX_PRE5
	case PREFIX_PRE6_SEP:
		return PREFIX_PRE6
	case PREFIX_PRE7_SEP:
		return PREFIX_PRE7
	case PREFIX_PRE8_SEP:
		return PREFIX_PRE8
	case PREFIX_PRE9_SEP:
		return PREFIX_PRE9
	default:
		panic("invalid prefix separator")
	}
}
func prefixBits2Sep(bits uint) byte {
	switch bits {
	case PREFIX_PRE4:
		return PREFIX_PRE4_SEP
	case PREFIX_PRE5:
		return PREFIX_PRE5_SEP
	case PREFIX_PRE6:
		return PREFIX_PRE6_SEP
	case PREFIX_PRE7:
		return PREFIX_PRE7_SEP
	case PREFIX_PRE8:
		return PREFIX_PRE8_SEP
	case PREFIX_PRE9:
		return PREFIX_PRE9_SEP
	default:
		panic("invalid prefix bits")
	}
}

const BASE_PUNCT = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ_abcdefghijklmnopqrstuvwxyz~"
const (
	BASE_0_SEP  = '0'
	BASE_1_SEP  = '1'
	BASE_2_SEP  = '2'
	BASE_3_SEP  = '3'
	BASE_4_SEP  = '4'
	BASE_5_SEP  = '5'
	BASE_6_SEP  = '6'
	BASE_7_SEP  = '7'
	BASE_8_SEP  = '8'
	BASE_9_SEP  = '9'
	BASE_10_SEP = 'A'
	BASE_11_SEP = 'B'
	BASE_12_SEP = 'C'
	BASE_13_SEP = 'D'
	BASE_14_SEP = 'E'
	BASE_15_SEP = 'F'
	BASE_16_SEP = 'G'
	BASE_17_SEP = 'H'
	BASE_18_SEP = 'I'
	BASE_19_SEP = 'J'
	BASE_20_SEP = 'K'
	BASE_21_SEP = 'L'
	BASE_22_SEP = 'M'
	BASE_23_SEP = 'N'
	BASE_24_SEP = 'O'
	BASE_25_SEP = 'P'
	BASE_26_SEP = 'Q'
	BASE_27_SEP = 'R'
	BASE_28_SEP = 'S'
	BASE_29_SEP = 'T'
	BASE_30_SEP = 'U'
	BASE_31_SEP = 'V'
	BASE_32_SEP = 'W'
	BASE_33_SEP = 'X'
	BASE_34_SEP = 'Y'
	BASE_35_SEP = 'Z'
	BASE_36_SEP = '_'
	BASE_37_SEP = 'a'
	BASE_38_SEP = 'b'
	BASE_39_SEP = 'c'
	BASE_40_SEP = 'd'
	BASE_41_SEP = 'e'
	BASE_42_SEP = 'f'
	BASE_43_SEP = 'g'
	BASE_44_SEP = 'h'
	BASE_45_SEP = 'i'
	BASE_46_SEP = 'j'
	BASE_47_SEP = 'k'
	BASE_48_SEP = 'l'
	BASE_49_SEP = 'm'
	BASE_50_SEP = 'n'
	BASE_51_SEP = 'o'
	BASE_52_SEP = 'p'
	BASE_53_SEP = 'q'
	BASE_54_SEP = 'r'
	BASE_55_SEP = 's'
	BASE_56_SEP = 't'
	BASE_57_SEP = 'u'
	BASE_58_SEP = 'v'
	BASE_59_SEP = 'w'
	BASE_60_SEP = 'x'
	BASE_61_SEP = 'y'
	BASE_62_SEP = 'z'
	BASE_63_SEP = '~'
)
const (
	BASE_0  = 0
	BASE_1  = 1
	BASE_2  = 2
	BASE_3  = 3
	BASE_4  = 4
	BASE_5  = 5
	BASE_6  = 6
	BASE_7  = 7
	BASE_8  = 8
	BASE_9  = 9
	BASE_10 = 10
	BASE_11 = 11
	BASE_12 = 12
	BASE_13 = 13
	BASE_14 = 14
	BASE_15 = 15
	BASE_16 = 16
	BASE_17 = 17
	BASE_18 = 18
	BASE_19 = 19
	BASE_20 = 20
	BASE_21 = 21
	BASE_22 = 22
	BASE_23 = 23
	BASE_24 = 24
	BASE_25 = 25
	BASE_26 = 26
	BASE_27 = 27
	BASE_28 = 28
	BASE_29 = 29
	BASE_30 = 30
	BASE_31 = 31
	BASE_32 = 32
	BASE_33 = 33
	BASE_34 = 34
	BASE_35 = 35
	BASE_36 = 36
	BASE_37 = 37
	BASE_38 = 38
	BASE_39 = 39
	BASE_40 = 40
	BASE_41 = 41
	BASE_42 = 42
	BASE_43 = 43
	BASE_44 = 44
	BASE_45 = 45
	BASE_46 = 46
	BASE_47 = 47
	BASE_48 = 48
	BASE_49 = 49
	BASE_50 = 50
	BASE_51 = 51
	BASE_52 = 52
	BASE_53 = 53
	BASE_54 = 54
	BASE_55 = 55
	BASE_56 = 56
	BASE_57 = 57
	BASE_58 = 58
	BASE_59 = 59
	BASE_60 = 60
	BASE_61 = 61
	BASE_62 = 62
	BASE_63 = 63
)

func baseSep2Bits(sep byte) uint {
	switch sep {
	case BASE_0_SEP:
		return BASE_0
	case BASE_1_SEP:
		return BASE_1
	case BASE_2_SEP:
		return BASE_2
	case BASE_3_SEP:
		return BASE_3
	case BASE_4_SEP:
		return BASE_4
	case BASE_5_SEP:
		return BASE_5
	case BASE_6_SEP:
		return BASE_6
	case BASE_7_SEP:
		return BASE_7
	case BASE_8_SEP:
		return BASE_8
	case BASE_9_SEP:
		return BASE_9
	case BASE_10_SEP:
		return BASE_10
	case BASE_11_SEP:
		return BASE_11
	case BASE_12_SEP:
		return BASE_12
	case BASE_13_SEP:
		return BASE_13
	case BASE_14_SEP:
		return BASE_14
	case BASE_15_SEP:
		return BASE_15
	case BASE_16_SEP:
		return BASE_16
	case BASE_17_SEP:
		return BASE_17
	case BASE_18_SEP:
		return BASE_18
	case BASE_19_SEP:
		return BASE_19
	case BASE_20_SEP:
		return BASE_20
	case BASE_21_SEP:
		return BASE_21
	case BASE_22_SEP:
		return BASE_22
	case BASE_23_SEP:
		return BASE_23
	case BASE_24_SEP:
		return BASE_24
	case BASE_25_SEP:
		return BASE_25
	case BASE_26_SEP:
		return BASE_26
	case BASE_27_SEP:
		return BASE_27
	case BASE_28_SEP:
		return BASE_28
	case BASE_29_SEP:
		return BASE_29
	case BASE_30_SEP:
		return BASE_30
	case BASE_31_SEP:
		return BASE_31
	case BASE_32_SEP:
		return BASE_32
	case BASE_33_SEP:
		return BASE_33
	case BASE_34_SEP:
		return BASE_34
	case BASE_35_SEP:
		return BASE_35
	case BASE_36_SEP:
		return BASE_36
	case BASE_37_SEP:
		return BASE_37
	case BASE_38_SEP:
		return BASE_38
	case BASE_39_SEP:
		return BASE_39
	case BASE_40_SEP:
		return BASE_40
	case BASE_41_SEP:
		return BASE_41
	case BASE_42_SEP:
		return BASE_42
	case BASE_43_SEP:
		return BASE_43
	case BASE_44_SEP:
		return BASE_44
	case BASE_45_SEP:
		return BASE_45
	case BASE_46_SEP:
		return BASE_46
	case BASE_47_SEP:
		return BASE_47
	case BASE_48_SEP:
		return BASE_48
	case BASE_49_SEP:
		return BASE_49
	case BASE_50_SEP:
		return BASE_50
	case BASE_51_SEP:
		return BASE_51
	case BASE_52_SEP:
		return BASE_52
	case BASE_53_SEP:
		return BASE_53
	case BASE_54_SEP:
		return BASE_54
	case BASE_55_SEP:
		return BASE_55
	case BASE_56_SEP:
		return BASE_56
	case BASE_57_SEP:
		return BASE_57
	case BASE_58_SEP:
		return BASE_58
	case BASE_59_SEP:
		return BASE_59
	case BASE_60_SEP:
		return BASE_60
	case BASE_61_SEP:
		return BASE_61
	case BASE_62_SEP:
		return BASE_62
	case BASE_63_SEP:
		return BASE_63
	default:
		panic("invalid base separator")
	}
}
func baseBits2Sep(bits uint) byte {
	switch bits {
	case BASE_0:
		return BASE_0_SEP
	case BASE_1:
		return BASE_1_SEP
	case BASE_2:
		return BASE_2_SEP
	case BASE_3:
		return BASE_3_SEP
	case BASE_4:
		return BASE_4_SEP
	case BASE_5:
		return BASE_5_SEP
	case BASE_6:
		return BASE_6_SEP
	case BASE_7:
		return BASE_7_SEP
	case BASE_8:
		return BASE_8_SEP
	case BASE_9:
		return BASE_9_SEP
	case BASE_10:
		return BASE_10_SEP
	case BASE_11:
		return BASE_11_SEP
	case BASE_12:
		return BASE_12_SEP
	case BASE_13:
		return BASE_13_SEP
	case BASE_14:
		return BASE_14_SEP
	case BASE_15:
		return BASE_15_SEP
	case BASE_16:
		return BASE_16_SEP
	case BASE_17:
		return BASE_17_SEP
	case BASE_18:
		return BASE_18_SEP
	case BASE_19:
		return BASE_19_SEP
	case BASE_20:
		return BASE_20_SEP
	case BASE_21:
		return BASE_21_SEP
	case BASE_22:
		return BASE_22_SEP
	case BASE_23:
		return BASE_23_SEP
	case BASE_24:
		return BASE_24_SEP
	case BASE_25:
		return BASE_25_SEP
	case BASE_26:
		return BASE_26_SEP
	case BASE_27:
		return BASE_27_SEP
	case BASE_28:
		return BASE_28_SEP
	case BASE_29:
		return BASE_29_SEP
	case BASE_30:
		return BASE_30_SEP
	case BASE_31:
		return BASE_31_SEP
	case BASE_32:
		return BASE_32_SEP
	case BASE_33:
		return BASE_33_SEP
	case BASE_34:
		return BASE_34_SEP
	case BASE_35:
		return BASE_35_SEP
	case BASE_36:
		return BASE_36_SEP
	case BASE_37:
		return BASE_37_SEP
	case BASE_38:
		return BASE_38_SEP
	case BASE_39:
		return BASE_39_SEP
	case BASE_40:
		return BASE_40_SEP
	case BASE_41:
		return BASE_41_SEP
	case BASE_42:
		return BASE_42_SEP
	case BASE_43:
		return BASE_43_SEP
	case BASE_44:
		return BASE_44_SEP
	case BASE_45:
		return BASE_45_SEP
	case BASE_46:
		return BASE_46_SEP
	case BASE_47:
		return BASE_47_SEP
	case BASE_48:
		return BASE_48_SEP
	case BASE_49:
		return BASE_49_SEP
	case BASE_50:
		return BASE_50_SEP
	case BASE_51:
		return BASE_51_SEP
	case BASE_52:
		return BASE_52_SEP
	case BASE_53:
		return BASE_53_SEP
	case BASE_54:
		return BASE_54_SEP
	case BASE_55:
		return BASE_55_SEP
	case BASE_56:
		return BASE_56_SEP
	case BASE_57:
		return BASE_57_SEP
	case BASE_58:
		return BASE_58_SEP
	case BASE_59:
		return BASE_59_SEP
	case BASE_60:
		return BASE_60_SEP
	case BASE_61:
		return BASE_61_SEP
	case BASE_62:
		return BASE_62_SEP
	case BASE_63:
		return BASE_63_SEP
	default:
		panic("invalid base bits")
	}
}

const FRAME_PUNCT = "."
const (
	FRAME_TERM_SEP = '.'
)
const (
	FRAME_TERM = 0
)

func frameSep2Bits(sep byte) uint {
	switch sep {
	case FRAME_TERM_SEP:
		return FRAME_TERM
	default:
		panic("invalid frame separator")
	}
}
func frameBits2Sep(bits uint) byte {
	switch bits {
	case FRAME_TERM:
		return FRAME_TERM_SEP
	default:
		panic("invalid frame bits")
	}
}
