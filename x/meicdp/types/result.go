package types

import (
	"encoding/binary"
	"errors"
)

type BorshDecoder struct {
	data   []byte
	offset uint32
}

func NewBorshDecoder(data []byte) BorshDecoder {
	return BorshDecoder{
		data:   data,
		offset: 0,
	}
}

func (decoder *BorshDecoder) Finished() bool {
	return decoder.offset == uint32(len(decoder.data))
}

func (decoder *BorshDecoder) DecodeString() (string, error) {
	length, err := decoder.DecodeU32()
	if err != nil {
		return "", err
	}
	if uint32(len(decoder.data)) < decoder.offset+length {
		return "", errors.New("Borsh: out of range")
	}
	val := string(decoder.data[decoder.offset : decoder.offset+uint32(length)])
	decoder.offset += length
	return val, nil
}

func (decoder *BorshDecoder) DecodeU8() (uint8, error) {
	if uint32(len(decoder.data)) < decoder.offset+1 {
		return 0, errors.New("Borsh: out of range")
	}
	val := uint8(decoder.data[decoder.offset])
	decoder.offset++
	return val, nil
}

func (decoder *BorshDecoder) DecodeU32() (uint32, error) {
	if uint32(len(decoder.data)) < decoder.offset+4 {
		return 0, errors.New("Borsh: out of range")
	}
	val := binary.LittleEndian.Uint32(decoder.data[decoder.offset : decoder.offset+4])
	decoder.offset += 4
	return val, nil
}

func (decoder *BorshDecoder) DecodeU64() (uint64, error) {
	if uint32(len(decoder.data)) < decoder.offset+8 {
		return 0, errors.New("Borsh: out of range")
	}
	val := binary.LittleEndian.Uint64(decoder.data[decoder.offset : decoder.offset+8])
	decoder.offset += 8
	return val, nil
}

type Result struct {
	Px uint64
}

func DecodeResult(data []byte) (Result, error) {
	decoder := NewBorshDecoder(data)

	px, err := decoder.DecodeU64()
	if err != nil {
		return Result{}, err
	}

	if !decoder.Finished() {
		return Result{}, errors.New("Borsh: bytes left when decode result")
	}

	return Result{
		Px: px,
	}, nil
}
