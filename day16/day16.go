package day16

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/ImmaculatePine/adventofcode2021/utils"
)

type Packet interface {
	Version() uint64
	TypeID() uint64
	IsOperator() bool
	Subpackets() []Packet
	Value() uint64
}

func NewPacket(bits string) (Packet, string, error) {
	version, err := binToDec(bits[0:3])
	if err != nil {
		return nil, bits, err
	}

	typeID, err := binToDec(bits[3:6])
	if err != nil {
		return nil, bits, err
	}

	bits = bits[6:]

	if typeID == 4 {
		return NewValuePacket(version, typeID, bits)
	} else {
		return NewOperatorPacket(version, typeID, bits)
	}
}

type ValuePacket struct {
	version uint64
	typeID  uint64
	value   uint64
}

func NewValuePacket(version uint64, typeID uint64, bits string) (*ValuePacket, string, error) {
	var values []string
	for {
		chunk := bits[0:5]
		values = append(values, chunk[1:5])
		bits = bits[5:]
		if chunk[0] == '0' {
			break
		}
	}

	value, err := binToDec(strings.Join(values, ""))
	if err != nil {
		return nil, bits, err
	}

	packet := &ValuePacket{
		version,
		typeID,
		value,
	}

	return packet, bits, nil
}

func (p *ValuePacket) Version() uint64 {
	return p.version
}

func (p *ValuePacket) TypeID() uint64 {
	return p.typeID
}

func (_ *ValuePacket) IsOperator() bool {
	return false
}

func (_ *ValuePacket) Subpackets() []Packet {
	return nil
}

func (p *ValuePacket) Value() uint64 {
	return p.value
}

type OperatorPacket struct {
	version   uint64
	typeID    uint64
	mode      string
	modeParam uint64
	data      string
	packets   []Packet
}

func NewOperatorPacket(version uint64, typeID uint64, bits string) (*OperatorPacket, string, error) {
	mode := bits[:1]
	var modeParam uint64
	var data string
	var rest string
	var err error
	switch mode {
	case "0":
		modeParam, err = binToDec(bits[1:16])
		if err != nil {
			return nil, bits, err
		}
		data = bits[16 : 16+modeParam]
		rest = bits[16+modeParam:]
	case "1":
		modeParam, err = binToDec(bits[1:12])
		if err != nil {
			return nil, bits, err
		}
		data = bits[12:]
	default:
		return nil, bits, fmt.Errorf("unexpected operator mode %s", mode)
	}

	var packets []Packet
	if mode == "0" {
		for {
			sub, subrest, err := NewPacket(data)
			if err != nil {
				return nil, bits, err
			}
			packets = append(packets, sub)
			data = subrest
			if data == "" {
				break
			}
		}
	} else {
		for i := 0; i < int(modeParam); i++ {
			sub, subrest, err := NewPacket(data)
			if err != nil {
				return nil, bits, err
			}
			packets = append(packets, sub)
			data = subrest
			rest = subrest
		}
	}

	packet := &OperatorPacket{
		version,
		typeID,
		mode,
		modeParam,
		data,
		packets,
	}

	return packet, rest, nil
}

func (p *OperatorPacket) Version() uint64 {
	return p.version
}

func (p *OperatorPacket) TypeID() uint64 {
	return p.typeID
}

func (_ *OperatorPacket) IsOperator() bool {
	return true
}

func (p *OperatorPacket) Subpackets() []Packet {
	return p.packets
}

func (p *OperatorPacket) Value() uint64 {
	switch p.typeID {
	// Sum
	case 0:
		var sum uint64
		for _, sub := range p.packets {
			sum += sub.Value()
		}
		return sum
	// Product
	case 1:
		product := uint64(1)
		for _, sub := range p.packets {
			product *= sub.Value()
		}
		return product
	// Min
	case 2:
		min := p.packets[0].Value()
		for _, sub := range p.packets {
			val := sub.Value()
			if val < min {
				min = val
			}
		}
		return min
	// Max
	case 3:
		max := p.packets[0].Value()
		for _, sub := range p.packets {
			val := sub.Value()
			if val > max {
				max = val
			}
		}
		return max
	// GT
	case 5:
		if p.packets[0].Value() > p.packets[1].Value() {
			return 1
		} else {
			return 0
		}
	// LT
	case 6:
		if p.packets[0].Value() < p.packets[1].Value() {
			return 1
		} else {
			return 0
		}
	// EQ
	case 7:
		if p.packets[0].Value() == p.packets[1].Value() {
			return 1
		} else {
			return 0
		}
	default:
		return 0
	}
}

func hexToBin(hex string) (string, error) {
	ui, err := strconv.ParseUint(hex, 16, 64)
	if err != nil {
		return "", err
	}

	binary := fmt.Sprintf("%04b", ui)
	padding := len(hex)*4 - len(binary)
	for i := 0; i < padding; i++ {
		binary = "0" + binary
	}

	return binary, nil
}

func binToDec(bin string) (uint64, error) {
	dec, err := strconv.ParseUint(bin, 2, 64)
	if err != nil {
		return 0, err
	}

	return dec, nil
}

func Task1() error {
	root, err := readInput()
	if err != nil {
		return err
	}
	result := sumVersions(root, 0)
	fmt.Println(result)
	return nil
}

func Task2() error {
	root, err := readInput()
	if err != nil {
		return err
	}
	fmt.Println(root.Value())
	return nil
}

func sumVersions(packet Packet, acc uint64) uint64 {
	acc += packet.Version()
	if packet.IsOperator() {
		for _, sub := range packet.Subpackets() {
			acc += sumVersions(sub, 0)
		}
	}
	return acc
}

func readInput() (Packet, error) {
	lines, err := utils.ReadInputStrings("./day16/input.txt")
	if err != nil {
		return nil, err
	}

	var bits string
	for _, hex := range lines[0] {
		bin, err := hexToBin(string(hex))
		if err != nil {
			return nil, err
		}
		bits += bin
	}

	root, _, err := NewPacket(bits)
	return root, err
}
