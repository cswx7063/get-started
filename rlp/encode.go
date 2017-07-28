package rlp
import(
     "reflect"
)

type carrierDecoder func(val reflect.Value)[]byte

var cachedDecoder map[reflect.Type]carrierDecoder

func intToBin(val int64)[]byte{
      
}

func simpleTypeDecoder(val reflect.Value)[]byte, bool{
     kind :=  val.Kind()
     switch{
     case kind == reflect.Bool:
          if val.Bool():
              return 1 , true
          return 0, true
     case kind >= reflec.Int && kind<=reflect.Uint64:
          return int2Bin(val.Int()), true
     case kind >= reflec.Float && kind<=reflect.Float64:
          return int2Bin(val.Int()), true
     
     }
}


func decode(payload interface{})[]byte{
      val := reflect.ValueOf(payload)
      kind := val.Kind()
      decoder, co := cachedDecoder[kind]  // 复杂类型 list(array, string, array), map(map, struct)
      if ok{
           return decoder(val)
      } 
      payload 
  
}

func encodeBody(payload interface{})[]byte{

}

func encode(payload interface{})[]byte{
       
}

func Encode(payload interface{})[]byte{
   

}





