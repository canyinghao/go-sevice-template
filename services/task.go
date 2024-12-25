package services

import (
	"errors"

	"github.com/canyinghao/go-sevice-template/model"
	"go.uber.org/zap"
)

func GetTaskOne(ids string) (*model.Task, error) {

	row := db.QueryRow("select * from task where id=$1", ids)
	var id int
	var name string
	var description string
	var priority string
	err := row.Scan(&id, &name, &description, &priority)
	if err != nil {
		return nil, errors.New("没有数据")
	}
	task := model.Task{Id: id, Name: name, Description: description, Priority: priority}

	return &task, nil
}

func GetTaskAll() ([]model.Task, error) {

	rows, err := db.Query("select * from task limit 20")

	if err != nil {
		zap.L().Error("查询数据失败", zap.Error(err))
		return nil, err
	}
	defer rows.Close()
	var list []model.Task

	for rows.Next() {
		var id int
		var name string
		var description string
		var priority string
		rows.Scan(&id, &name, &description, &priority)
		task := model.Task{Id: id, Name: name, Description: description, Priority: priority}
		list = append(list, task)
	}

	return list, nil
}
