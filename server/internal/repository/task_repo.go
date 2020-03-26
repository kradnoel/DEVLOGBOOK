package repository

import (
	"github.com/kradnoel/teka/internal/persistance"
	"github.com/kradnoel/teka/internal/types"
	"github.com/kradnoel/teka/internal/utils"
)

func CreateTask(task types.Task) bool {

	persistance := persistance.GetPersistance()

	err := persistance.Model(&types.Task{}).Create(task).Error

	if err != nil {
		return false
	}

	return true
}

func DeleteTask(id string) bool {

	persistance := persistance.GetPersistance()

	// Start transactiong
	tx := persistance.Begin()

	// Delete all subtasks associated with the task
	err := tx.Model(&types.Subtask{}).Where("task_id = ?", utils.ParseToUUID(id)).Delete(&types.Subtask{}).Error
	if err != nil {
		tx.Rollback()
		return false
	}

	// Delete the task
	err = tx.Model(&types.Task{}).Where("guid = ?", utils.ParseToUUID(id)).Delete(&types.Task{}).Error

	if err != nil {
		tx.Rollback()
		return false
	}

	tx.Commit()
	return true

}

func GetTasks(id string) []types.Task {
	persistance := persistance.GetPersistance()
	tasks := make([]types.Task, 0)
	tasks1 := make([]types.Task, 0)
	//subtasks := make([]types.Subtask, 0)

	err := persistance.Model(&types.Task{}).Where("project_id = ?", utils.ParseToUUID(id)).Find(&tasks).Error

	//	err := persistance.Table("tasks").Where("project_id = ?", utils.ParseToUUID(id)).Find(&tasks).Error
	//defer rows.Close()

	if err != nil {
		return []types.Task{}
	}

	for _, task := range tasks {
		var spend_hours float64

		subtasks := GetSubTasks(task.Guid.String())

		for _, subtask := range subtasks {
			result, err := utils.BetweenDateToHour(subtask.From, subtask.To)

			if err != nil {
				return []types.Task{}
			}

			spend_hours = spend_hours + result
		}

		task.TimeSpent, _ = utils.CustomFloat(spend_hours)

		tasks1 = append(tasks1, task)

		//fmt.Printf("spend-hours = %f \n", task.SpendHours)
	}

	tasks = tasks1

	/*

		for rows.Next() {
			var task types.Task
			persistance.ScanRows(rows, &task)
			tasks = append(tasks, task)
		}*/

	return tasks

}

func SeedTasks() {
	persistance := persistance.GetPersistance()

	tasks := []types.Task{
		types.Task{Guid: utils.ParseToUUID("8c53e53e-0aa1-4465-8572-7aa1124284ce"), Description: "Unnoficial BCI Exchange API", ProjectId: utils.ParseToUUID("c889b8b7-e134-450f-a9c0-0ce2e468d80e")},
		types.Task{Guid: utils.GenerateUUID(), Description: "Shortlink in express + go", ProjectId: utils.ParseToUUID("c889b8b7-e134-450f-a9c0-0ce2e468d80e")},
		types.Task{Guid: utils.GenerateUUID(), Description: "Document manager in express + go", ProjectId: utils.ParseToUUID("c889b8b7-e134-450f-a9c0-0ce2e468d80e")},
	}

	//project1 := types.Project{Guid: utils.GenerateUUID(), Title: "CambioMz", Description: "Unnoficial BCI Exchange API"}
	//project2 := types.Project{Guid: utils.GenerateUUID(), Title: "XRTLink", Description: "Shortlink in express + go"}
	//project3 := types.Project{Guid: utils.GenerateUUID(), Title: "Biblioteka", Description: "Document manager in express + go"}
	//persistance.Model(&types.Project{}).Create(project1)

	for _, value := range tasks {
		persistance.Model(&types.Task{}).Create(value)
	}

}
