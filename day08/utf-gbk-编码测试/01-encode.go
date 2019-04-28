package main
import "fmt"
import "github.com/axgle/mahonia"
func main(){
	enc:=mahonia.NewEncoder("gbk")
	//converts a  string from UTF-8 to gbk encoding.
	fmt.Println(enc.ConvertString("hello,世界"))
	fmt.Println("hell 世界")
}
