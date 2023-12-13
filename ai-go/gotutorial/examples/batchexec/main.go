/*
@File   : main.go
@Author : pan
@Time   : 2023-12-13 10:19:15
*/
package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"
)

// 批量执行标准模板
/*
需求
一个接口调用时，接收到一个列表，十个元素，需要并发执行十个任务，每个任务都要返回执行的结果和异常，
然后对返回的结果装填到一个切片列表里，统一返回结果。
*/

type Order struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

// 需要记录原始顺序的时候，定义个带编号的结构体
type OrderWithSeq struct {
	Seq       int
	OrderItem Order
}

// 重写相关排序类型
type BySeq []OrderWithSeq

func (a BySeq) Len() int {
	return len(a)
}
func (a BySeq) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a BySeq) Less(i, j int) bool {
	return a[i].Seq < a[j].Seq
}

func main() {
	taskNum := 10
	orderCh := make(chan OrderWithSeq, taskNum) //接收带序号的结构体
	errCh := make(chan error, taskNum)          //接收返回的异常
	wg := sync.WaitGroup{}
	//在执行任务时，加入序号
	for i := 0; i < taskNum; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
				//协程内单独捕捉异常
				if r := recover(); r != nil {
					err := fmt.Errorf("system panic:%v", r)
					errCh <- err //此处将panic信息转为err返回，也可以按需求和异常等级进行处理
					return
				}
			}()
			//组装返回结果
			res := Order{
				Name: "num: " + strconv.Itoa(i),
				Id:   i,
			}
			orderCh <- OrderWithSeq{
				Seq:       i, //带上i这个序号
				OrderItem: res,
			}
		}()
		wg.Wait()
		//接收信息，也按带序号的结构体进行组装
		orderSeqList := make([]OrderWithSeq, taskNum)
		timeoutTime := time.Second * 3
		taskTimer := time.NewTimer(timeoutTime)
		for i := 0; i < taskNum; i++ {
			select {
			case order, ok := <-orderCh: //接收orderCh
				if ok {
					orderSeqList = append(orderSeqList, order)
				}
			case err := <-errCh: //接收errCh
				if err != nil {
					fmt.Println(err)
				}
				return
			case <-taskTimer.C: //处理超时
				err := errors.New("task timeout")
				fmt.Println(err)
				return
			default:
				fmt.Println("done")
			}
			taskTimer.Reset(timeoutTime)
		}
		close(orderCh)
		close(errCh)
		//按原始顺序进行排序
		sort.Sort(BySeq(orderSeqList))
	}
	fmt.Println("##########################################")
}
