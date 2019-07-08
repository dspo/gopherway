package production_and_consumer

import (
	"fmt"
)

func ProductionAndConsumerModel() {
	productionLines := 5 //生产线的条数
	consumersCnt := 1 //消费者个数
	tasks := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m"}
	taskChan := make(chan string, len(tasks))
	productionSigns := make(chan bool, productionLines)
	consumerSigns := make(chan bool, consumersCnt)

	for i := 0; i < len(tasks); i++{
		taskChan <- tasks[i]
	}

	resChan := make(chan string, 1000) //成品库

	fmt.Println("开启所有生产线")
	//开启所有生产线
	for i := 0; i < productionLines; i++ {
		//生产线：拿到原材料，加工，放到成品库
		go func() {
			for {
				if len(taskChan) == 0 {
					//没有生产任务时，发送停产信号，并停产
					productionSigns <- true
					return
				}
				taskElement := <-taskChan  //拿到原材料
				result := func(ele string) string{return taskElement + taskElement}(taskElement) //某个加工原材料的函数
				fmt.Println("拿到", taskElement, "生产生了==>", result)
				resChan <- result  //放到成品库
			}
		}()
	}

	fmt.Println("一个生产者进行消费")
	//一个生产者进行消费
	go func() {
		for {
			//如果工厂已经停止生产，且成品库里也没有产品了，就停止消费
			if len(productionSigns) == cap(productionSigns) && len(resChan) == 0 {
				consumerSigns <- true
				return
			}
			result := <- resChan
			fmt.Println("消费了：", result)
		}
	}()

	for {
		if len(consumerSigns) == cap(consumerSigns) {
			break
		}
	}
	close(productionSigns)
	close(consumerSigns)
	close(taskChan)
	close(resChan)

	fmt.Println("生产完毕，消费完毕！")
}
