package controller

import (
	"awesomeProject6/common"
	"awesomeProject6/dao"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
	"strings"
)

func TaskAdd(c *gin.Context)  {
	taskType, err := strconv.Atoi(c.PostForm("type"))
	frequency := c.PostForm("Frequency")
	sitesArray := c.PostFormArray("sites[]")
	sites := strings.Replace(strings.Trim(fmt.Sprint(sitesArray), "[]"), " ", ",", -1)
	if err != nil {
		panic(err)
	}
	task := dao.Task{
		Type:      taskType,
		Frequency: frequency,
		Sites:     sites,
		Status:    dao.TaskRunning,
	}
	common.GormPool.Create(&task)
	taskInfo, err := json.Marshal(c.Request.PostForm)
	if err != nil {
		log.Fatalf("Json marshaling failed：%s", err)
	}
	taskData := dao.TaskData{
		TaskID: task.ID,
		Data: string(taskInfo),
	}
	common.GormPool.Create(&taskData)
	data := map[string]interface{} {
		"task_id": task.ID,
	}

	c.JSON(200, gin.H{
		"state": 1,
		"msg": "success",
		"data": data,
	})
}

func Dispatch(task dao.Task, data dao.TaskData){
	sites := strings.Split(task.Sites, ",")
	for siteID := range sites{
		site := dao.Site{}
		err := common.GormPool.Where("id = ?", siteID).First(&site).Error
		if err != nil {
			panic(err)
		}
		if site.ID < 1 {
			panic("site: " + string(siteID) + " 不存在")
		}
		go dispatchWorkerToSite(task, data, site)
	}
}

func dispatchWorkerToSite(task dao.Task, data dao.TaskData, site dao.Site)  {

}
