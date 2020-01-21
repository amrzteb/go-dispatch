package dao

import (
	"awesomeProject6/common"
	"time"
)

const TaskStop = 0
const TaskRunning = 1

const SingleTask  = 0
const PeriodicTask  = 1

type Task struct {
	ID        int  `json:"id" orm:"column(id);auto"`
	Type      int
	Frequency string
	Sites     string
	Status    int `gorm:"default:1"`
	createAt  time.Time
}

func(f *Task) Find(id int) (*Task, error) {
	var task Task
	if err := common.GormPool.Where("id = ?", id).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}

func (f *Task) Create(task Task) (int, error) {

	if err := common.GormPool.Create(&task).Error; err != nil {
		return 0, err
	}
	return task.ID, nil

}