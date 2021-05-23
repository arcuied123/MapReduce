
package main

import "fmt"
import "os"
import "io/ioutil"
import "strings"
import "unicode"




func main() {
	file,_ := os.Open("pg-being_ernest.txt")
	content,_ := ioutil.ReadAll(file)
	//fmt.Println(file)
//	fmt.Println(content)
	c:= string(content)
	f:= func(b rune) bool {
		   return !unicode.IsLetter(b) }
	words:=strings.FieldsFunc(c,f)
//	fmt.Println(words)
	kva := []KV{}
	res :=  []KV{}
	for _,v:= range words{
		kv_e:=KV{v,1}
		kva = append(kva,kv_e)
		iscontain := isContain(v,res)
		if iscontain == false{
			res = append(res,KV{v,0})
		}

	}
//	fmt.Println(res)
	for _,v := range kva{
		for index,j :=  range res{
			if v.k == j.k{
				res[index].v = res[index].v+1
				break
			}
		}

	}
	fmt.Println(res)
}


func isContain(target string, kv_list []KV) bool{
	for _,v := range kv_list{
		if target == v.k{
			return true
		}
	}
	return false

}


type KV struct {
	k string
	v int

}


