package main

//import (
//	"fmt"
//	"github.com/wplib/go-lib/project"
//	"os"
//)
//
//func main() {
//	wd, _ := os.Getwd()
//	println(wd)
//
//	defer func() {
//		if r := recover(); r != nil {
//			fmt.Println("ERROR:", r)
//		}
//	}()
//	pf := stack.NewProjectFile()
//	pf.LoadJSON()
//	for i, c := range pf.GetComponents() {
//		fmt.Printf("\n[%d] %-22v %v", i, c.GetType()+":", c.GetLocation())
//	}
//}
