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
    binary.BigEndian.PutUint64(bytes, val)
    return removeVacant(bytes)
}

func ByteToUint64(bytes []byte)uint64{
    if len(bytes) != 8{
         all := make([]byte, 8-len(bytes))
         all = append(all, bytes...)
         bytes = all
    }
    return binary.BigEndian.Uint64(bytes) 
}

func Uint32ToByte(val uint32) []byte {
    bytes := make([]byte, 4)
    binary.BigEndian.PutUint32(bytes, val)
    return removeVacant(bytes)
}

func ByteToUint32(bytes []byte)uint32{
    if len(bytes) != 4{
         all := make([]byte, 4-len(bytes))
         all = append(all, bytes...)
         bytes = all
    }
    return binary.BigEndian.Uint32(bytes) 
}

func Float32ToByte(val float32) []byte {
    bits := math.Float32bits(val)
    bytes := make([]byte, 4)
    binary.BigEndian.PutUint32(bytes, bits)
    return removeVacant(bytes)
}

func ByteToFloat32(bytes []byte) float32{
    if len(bytes) != 4{
         all := make([]byte, 4-len(bytes))
         all = append(all, bytes...)
         bytes = all
    }
    bits := binary.BigEndian.Uint32(bytes)
    return math.Float32frombits(bits)
}

func Float64ToByte(float float64) []byte {
    bits := math.Float64bits(float)
    bytes := make([]byte, 8)
    binary.BigEndian.PutUint64(bytes, bits)
    return removeVacant(bytes)
}

func ByteToFloat64(bytes []byte) float64 {
    if len(bytes) != 8{
         all := make([]byte, 8-len(bytes))
         all = append(all, bytes...)
         bytes = all
    }
    bits := binary.BigEndian.Uint64(bytes)
    return math.Float64frombits(bits)
}

func numEncode(val interface{})[]byte{
    bytes := make([]byte, 2) 
    v := 0
    signed := func (){
         bytes[1] = 1
        if v >=0{
            bytes[1] = 0
        }
    }
    switch v := val.(type){
    case int32:
        signed() 
        iBytes := Uint32ToByte(uint32(v)) 
        bytes[0] = 0x00 + byte(len(iBytes))    
        return append(bytes, iBytes...) 
    case float32:
        signed() 
        iBytes := Float32ToByte(v) 
        bytes[0] = 0x04 + byte(len(iBytes))    
        return append(bytes, iBytes...) 
    case int64:
        signed() 
        iBytes := Uint64ToByte(uint64(v)) 
        bytes[0] = 0x08 + byte(len(iBytes))    
        return append(bytes, iBytes...) 
    case float64:
        signed() 
        iBytes := Float64ToByte(float64(v)) 
        bytes[0] = 0x20 + byte(len(iBytes))    
        return append(bytes, iBytes...) 
    } 
    return nil 
}

func numDecode(payload []byte)interface{}{
    head := payload[0]
    switch{
    case head<=0x04:
         length := payload[0] - 0x00
         data := payload[2:2+length]
         v := int32(ByteToUint32(data))  
         if payload[1] == 1{
              return 0-v
         }     
         return v
    case head<=0x08:
         length := payload[0] - 0x04
         data := payload[2:2+length]
         v := ByteToFloat32(data)  
         if payload[1] == 1{
              return 0-v
         }     
         return v
    case head<=0x1f:
         length := payload[0] - 0x08
         data := payload[2:2+length]
         v := ByteToUint64(data)  
         if payload[1] == 1{
              return 0-v
         }     
         return v
    case head<=0x28:
         length := payload[0] - 0x20
         data := payload[2:2+length]
         v := ByteToFloat64(data)  
         if payload[1] == 1{
              return 0-v
         }     
         return v
    } 
    return nil
}

