package main


import "fmt"


// for sorting by key.
type ByKey []KV

// for sorting by key.
func (a ByKey) Len() int           { return len(a) }
func (a ByKey) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByKey) Less(i, j int) bool { return a[i].Key < a[j].Key }


func main(){

	values:= []string{"1","2"}
	fmt.Println(len(values))
	kv = KV{'shiki',2}
	fmt.Println(kv)



}


type KV struct {
	k string
        v int
      }
