package worker

import (
	"fmt"

	"github.com/surfycui/iec104"
)

//Task 数据处理任务
func Task(data *iec104.APDU) {
	//TODO 自定义数据处理
	println("do task")
	for _, signal := range data.Signals {
		fmt.Printf("%v %v %f %v\n", signal.TypeID, signal.Address, signal.Value, signal.Ts)
	}
}
