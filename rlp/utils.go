package rlp

import (
    "encoding/binary"
    "math"
    "fmt"
)

func removeVacant(bytes []byte)[]byte{
    for idx, _ := range bytes{
        if bytes[idx] != 0{
             return bytes[idx:]       
        }
    }
    return bytes
}

func Uint64ToByte(val uint64) []byte {
    bytes := make([]byte, 8)
    binary.LittleEndian.PutUint64(bytes, val)
    return removeVacant(bytes)
}

func ByteToUint64(bytes []byte)uint64{
    return binary.LittleEndian.Uint64(bytes) 
}

func Uint32ToByte(val uint32) []byte {
    bytes := make([]byte, 4)
    binary.LittleEndian.PutUint32(bytes, val)
    return removeVacant(bytes)
}

func ByteToUint32(bytes []byte)uint32{
    return binary.LittleEndian.Uint32(bytes) 
}

func Float32ToByte(float float32) []byte {
    bits := math.Float32bits(float)
    bytes := make([]byte, 4)
    binary.LittleEndian.PutUint32(bytes, bits)
    return removeVacant(bytes)
}

func ByteToFloat32(bytes []byte) float32 {
    if len(bytes) != 4{
         all := make([]byte, 4-len(bytes))
         all = append(all, bytes...)
         bytes = all
    }
    bits := binary.LittleEndian.Uint32(bytes)
    return math.Float32frombits(bits)
}

func Float64ToByte(float float64) []byte {
    bits := math.Float64bits(float)
    bytes := make([]byte, 8)
    binary.LittleEndian.PutUint64(bytes, bits)
    return removeVacant(bytes)
}

func ByteToFloat64(bytes []byte) float64 {
    if len(bytes) != 8{
         all := make([]byte, 8-len(bytes))
         all = append(all, bytes...)
         bytes = all
    }
    bits := binary.LittleEndian.Uint64(bytes)
    return math.Float64frombits(bits)
}






