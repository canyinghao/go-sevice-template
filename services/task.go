package services

import (
	"go.uber.org/zap"
)

type Task struct {
	Id          int
	Name        string
	Description string
	Priority    string
}

func GetTaskOne() Task {

	rows, err := db.Query("select * from task where id=1")
	if err != nil {
		zap.L().Error("查询数据失败", zap.Error(err))
	}
	defer rows.Close()
	var list [1]Task
	for rows.Next() {
		var id int
		var name string
		var description string
		var priority string
		rows.Scan(&id, &name, &description, &priority)
		task := Task{Id: id, Name: name, Description: description, Priority: priority}
		list[0] = task
	}

	return list[0]
}
