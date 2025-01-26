package wsjtxudp

import (
	"encoding/binary"
	"fmt"
	"math"

	"time"

	"github.com/leemcloughlin/jdn"
)

func (p *WSJTXParser) parseQuint8() (uint8, error) {
	if p.cursor >= len(p.raw) {
		return 0, fmt.Errorf("Cannot parse Quint8, not enough bytes left in buffer.")
	}

	parsedInt8 := p.raw[p.cursor]
	p.cursor++

	return parsedInt8, nil
}

func (p *WSJTXParser) parseQint32() (int32, error) {
	parsedUnsigned, err := p.parseQuint32()

	return int32(parsedUnsigned), err
}

func (p *WSJTXParser) parseQuint32() (uint32, error) {
	end := p.cursor + 4
	if end >= len(p.raw) {
		return 0, fmt.Errorf("Cannot parse Quint32, not enough bytes left in buffer.")
	}

	parsedInt32 := binary.BigEndian.Uint32(p.raw[p.cursor:end])
	p.cursor = end

	return parsedInt32, nil
}

func (p *WSJTXParser) parseQuint64() (uint64, error) {
	end := p.cursor + 8
	if end >= len(p.raw) {
		return 0, fmt.Errorf("Cannot parse Quint64, not enough bytes left in buffer.")
	}

	parsedInt64 := binary.BigEndian.Uint64(p.raw[p.cursor:end])
	p.cursor = end

	return parsedInt64, nil
}

func (p *WSJTXParser) parseBool() (bool, error) {
	if p.cursor >= len(p.raw) {
		return false, fmt.Errorf("Cannot parse boolean, not enough bytes left in buffer.")
	}

	parsedBool := p.raw[p.cursor] != 0
	p.cursor++

	return parsedBool, nil
}

func (p *WSJTXParser) parseFloat64() (float64, error) {
	end := p.cursor + 8
	if end >= len(p.raw) {
		return 0, fmt.Errorf("Cannot parse float, not enough bytes left in buffer.")
	}

	floatData := binary.BigEndian.Uint64(p.raw[p.cursor:end])
	parsedFloat := math.Float64frombits(floatData)

	p.cursor = end

	return parsedFloat, nil
}

func (p *WSJTXParser) parseQTime() (uint32, error) {
	return p.parseQuint32()
}

func (p *WSJTXParser) parseQDateTime() (time.Time, error) {
	parsedTime := time.Now()

	// See note about QDateTime vs standard uint32 from the official documentation.
	// https://sourceforge.net/p/wsjt/wsjtx/ci/master/tree/Network/NetworkMessage.hpp#l31
	julianDay, err := p.parseQuint64()
	if err != nil {
		return parsedTime, err
	}
	year, month, day := jdn.FromNumber(int(julianDay))

	msSinceMidnight, err := p.parseQuint32()
	if err != nil {
		return parsedTime, err
	}

	hours := int(msSinceMidnight / 1000 / 60 / 60)
	minutes := int((msSinceMidnight / 1000 / 60) % 60)
	seconds := int((msSinceMidnight / 1000) % 60)
	milliseconds := int(msSinceMidnight % 1000)

	timeSpec, err := p.parseQuint8()
	if err != nil {
		return parsedTime, err
	}

	switch int(timeSpec) {
	case 0:
		return time.Date(year, month, day, hours, minutes, seconds, milliseconds, time.Local), nil
	case 1:
		return time.Date(year, month, day, hours, minutes, seconds, milliseconds, time.UTC), nil
	default:
		return parsedTime, fmt.Errorf("Unsupported timespec: %d", timeSpec)
	}
}

/*
 * Type utf8  is a  utf-8 byte  string formatted  as a  QByteArray for
 * serialization purposes  (currently a quint32 size  followed by size
 * bytes, no terminator is present or counted).
 */
func (p *WSJTXParser) parseUtf8() (string, error) {

	end := p.cursor + 4
	if end >= len(p.raw) {
		return "", fmt.Errorf("Cannot parse Utf8, not enough bytes left in buffer to parse the string length.")
	}

	stringSize := binary.BigEndian.Uint32(p.raw[p.cursor:end])

	p.cursor = end
	end = p.cursor + int(stringSize)
	if end >= len(p.raw) {
		return "", fmt.Errorf("Cannot parse Utf8, not enough bytes left in buffer to parse the string.")
	}
	parsedString := string(p.raw[p.cursor:end])

	p.cursor = end
	return parsedString, nil
}
