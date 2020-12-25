package main

// func xadd(val *int32, delta int32) (new int32) {
// 	for {
// 		v := *val
// 		if Cas(val, v, v+delta) {
// 			return v + delta
// 		}
// 	}
// 	panic("unreached")
// }

//go:noescape
func Cas(ptr *uint32, old, new uint32) bool

//go:noescape
func Cas64(addr *uint64, old, new uint64) bool

func main() {

}
