package main

import (
	"fmt"
)

type Phone interface {
	call()
}
type AndroidPhone struct {
	name string
	price float64
}
type Iphone struct {
	name string
	price float64
}

func (a AndroidPhone) call(){
	fmt.Println(a.name+"手机正在通话")
}
func (a AndroidPhone) vedio(){
	fmt.Println(a.name+"正在播放视频")

}

func (p Iphone) call(){
	fmt.Println(p.name+"手机正在通话......")
}


func main() {
	var huawei Phone = AndroidPhone{"华为", 3895.99}
	huawei.call()

	huaweiNew,ok:= huawei.(AndroidPhone)
	huaweiNew.vedio()
	fmt.Println(ok)

	switch t := huawei.(type) {
		case AndroidPhone:
			fmt.Println("androidPhone",t)
		case Iphone:
			fmt.Println("iphone",t)
	}

	iPhone := Iphone{"苹果", 3895.99}
	iPhone.call()
}
