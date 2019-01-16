package strcon
import (
    "bytes"
    "unicode"
)

func SwapCase(str string)string {
    buf:=&bytes.buffer{}
    
    for _,r:=range(str){
	if unicode.isUpper(r){
	    buf.WriteRune(unicode.ToLower(r))
	} else {
	    buf.WriteRune(unicode.ToUpper(r))
	}
    }
    return buf.String()
}