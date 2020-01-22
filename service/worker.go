package service

import (
	"awesomeProject6/common"
	"awesomeProject6/dao"
	"fmt"
	"strconv"
	"time"
)

func getLoadByTask(task dao.Task, frequency string) float64 {
	if task.Type == dao.SingleTask {
		return 0.3
	}else if task.Type == dao.PeriodicTask {
		return getLoadByFrequency(frequency)
	}else {
		panic("任务类型不存在")
	}
}


func getLoadByFrequency(frequency string) float64 {
	ns, err := time.ParseDuration(frequency)
	if err != nil {
		panic(err)
	}
	secondsStr := fmt.Sprintf("%0.f\n", ns.Seconds(), ns)
	secondsInt, err := strconv.Atoi(secondsStr)
	if err != nil {
		panic("Frequency 参数不合法")
	}
	load := fmt.Sprintf("%.8f", 1 / secondsInt)
	value, err := strconv.ParseFloat(load, 64)
	if err != nil{
		panic(err)
	}
	return value
}

func getLowLoadWorker(site dao.Site)  {

	worker := dao.Worker{}
	common.GormPool.First(&worker)
}