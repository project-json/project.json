package main

//import (
//	"github.com/wplib/box-cli/box"
//	"fmt"
//	"github.com/pkg/errors"
//)
//
//func foo() error {
//	return nil
//}
//func bar() (int,error) {
//	return 1,nil
//}
//func baz() (int,string,error) {
//	return 1,"abc",errors.New("baz!")
//}
//func somethingElseIsWrong() bool  {
//	return true
//}
//func main()  {
//
//	var err error
//	for {
//		err = foo()
//		if err!=nil {
//			break
//		}
//		num,err := bar()
//		if err!=nil {
//			break
//		}
//		num,str,err := baz()
//		if err!=nil {
//			break
//		}
//		if somethingElseIsWrong() {
//			err= errors.New("Something else is wrong!")
//			break
//		}
//		fmt.Printf("Num: %v and Str: %v", num, str )
//	}
//	if err != nil {
//		fmt.Printf("Oh, the HORROR!" )
//	}
//
//
//
//
//
//
//	box.Execute()
//}
