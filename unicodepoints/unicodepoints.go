package unicodepoints

import "fmt"

func CodePointToBytes(codepoint uint32) []byte {
	if codepoint < 0x80 {
		return []byte{byte(codepoint)}
	} else if codepoint < 0x0800 { // א = U+05D0 = U+0b0101_1101_0000
		const u8BitsToRepresentCodePoint = 11
		var leadingU8Byte byte = 0b110 << 5
		var topFiveCPBytes = byte(codepoint >> (u8BitsToRepresentCodePoint - 5)) // 0b101110
		leadingU8Byte |= topFiveCPBytes                                          // 0b11101110

		var nextU8Byte byte = 0b10 << 6
		var bottomSixCPBits = byte(codepoint & 0b11_1111) // 0b1_0000
		nextU8Byte |= bottomSixCPBits
		return []byte{leadingU8Byte, nextU8Byte}
	} else if codepoint < 0x1_0000 { // € = U+20AC = U+0b0010_0000_1010_1100.
		const u8BitsToRepresentCodePoint = 16
		var leadingU8Byte byte = 0b1110 << 4
		var topFourCPBits = byte(codepoint >> (u8BitsToRepresentCodePoint - 4)) // 0010
		leadingU8Byte |= topFourCPBits

		var secondU8Byte byte = 0b10 << 6
		var nextSixCPBits byte = byte((codepoint & (0b11_1111 << 6)) >> 6)
		secondU8Byte |= nextSixCPBits

		var thirdU8Byte byte = 0b10 << 6
		nextSixCPBits = byte(codepoint & 0b11_1111)
		thirdU8Byte |= nextSixCPBits

		return []byte{leadingU8Byte, secondU8Byte, thirdU8Byte}
	} else if codepoint <= 0x10_FFFF { // 🥳 = U+1F973 = U+0b0001_1111_1001_0111_0011
		const u8BitsToRepresentCodePoint = 21
		var leadingU8Byte byte = 0b11110 << 3
		var topThreeCPBits = byte(codepoint >> (u8BitsToRepresentCodePoint - 3)) // 0010
		leadingU8Byte |= topThreeCPBits

		var secondU8Byte byte = 0b10 << 6
		var nextSixCPBits byte = byte((codepoint & (0b11_1111 << 12)) >> 12)
		secondU8Byte |= nextSixCPBits

		var thirdU8Byte byte = 0b10 << 6
		nextSixCPBits = byte((codepoint & (0b11_1111 << 6)) >> 6)
		thirdU8Byte |= nextSixCPBits

		var fourthU8Byte byte = 0b10 << 6
		nextSixCPBits = byte(codepoint & (0b11_1111))
		fourthU8Byte |= nextSixCPBits

		return []byte{leadingU8Byte, secondU8Byte, thirdU8Byte, fourthU8Byte}
	}
	panic(fmt.Sprintf("invalid codepoint: %d", codepoint))
}

func CodePointToString(codepoint uint32) string {
	return string(CodePointToBytes(codepoint))
}

func codePointToBytesCopilot(codePoint uint32) []byte {
	var bytes []byte
	if codePoint <= 0x7F {
		bytes = append(bytes, byte(codePoint))
	} else if codePoint <= 0x7FF {
		bytes = append(bytes, byte(0xC0|codePoint>>6))
		bytes = append(bytes, byte(0x80|codePoint&0x3F))
	} else if codePoint <= 0xFFFF {
		bytes = append(bytes, byte(0xE0|codePoint>>12))
		bytes = append(bytes, byte(0x80|codePoint>>6&0x3F))
		bytes = append(bytes, byte(0x80|codePoint&0x3F))
	} else if codePoint <= 0x10FFFF {
		bytes = append(bytes, byte(0xF0|codePoint>>18))
		bytes = append(bytes, byte(0x80|codePoint>>12&0x3F))
		bytes = append(bytes, byte(0x80|codePoint>>6&0x3F))
		bytes = append(bytes, byte(0x80|codePoint&0x3F))
	} else {
		panic(fmt.Sprintf("Code point %d is not a valid Unicode code point", codePoint))
	}
	return bytes
}
