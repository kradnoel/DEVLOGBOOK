package repository

import (
	"fmt"
	"time"

	"github.com/kradnoel/teka/internal/persistance"
	"github.com/kradnoel/teka/internal/types"
	"github.com/kradnoel/teka/internal/utils"
)

func CreateSubTask(subtask types.Subtask) bool {

	persistance := persistance.GetPersistance()

	err := persistance.Model(&types.Subtask{}).Create(subtask).Error

	if err != nil {
		return false
	}

	return true
}

func DeleteSubTask(id string) bool {

	persistance := persistance.GetPersistance()

	// Delete subtask
	err := persistance.Model(&types.Subtask{}).Where("guid = ?", utils.ParseToUUID(id)).Delete(&types.Subtask{}).Error

	if err != nil {
		return false
	}

	return true
}

func GetSubTasks(id string) []types.Subtask {
	persistance := persistance.GetPersistance()

	subtasks := make([]types.Subtask, 0)

	err := persistance.Model(&types.Subtask{}).Where("task_id = ?", utils.ParseToUUID(id)).Find(&subtasks).Error
	//defer rows.Close()

	if err != nil {
		fmt.Println(err)
	}

	/*

		for rows.Next() {
			var task types.Task
			persistance.ScanRows(rows, &task)
			tasks = append(tasks, task)
		}*/

	return subtasks

}

func SeedSubTasks() {
	persistance := persistance.GetPersistance()

	subtasks := []types.Subtask{
		types.Subtask{Guid: utils.GenerateUUID(), From: time.Now(), To: time.Now().Add(time.Minute * 10), TaskId: utils.ParseToUUID("8c53e53e-0aa1-4465-8572-7aa1124284ce")},
		types.Subtask{Guid: utils.GenerateUUID(), From: time.Now(), To: time.Now().Add(time.Minute * 10), TaskId: utils.ParseToUUID("8c53e53e-0aa1-4465-8572-7aa1124284ce")},
		types.Subtask{Guid: utils.GenerateUUID(), From: time.Now(), To: time.Now().Add(time.Minute * 10), TaskId: utils.ParseToUUID("8c53e53e-0aa1-4465-8572-7aa1124284ce")},
		types.Subtask{Guid: utils.GenerateUUID(), From: time.Now(), To: time.Now().Add(time.Minute * 10), TaskId: utils.ParseToUUID("8c53e53e-0aa1-4465-8572-7aa1124284ce")},
		/*types.Task{Guid: utils.GenerateUUID(), Description: "Unnoficial BCI Exchange API", ProjectId: "c889b8b7-e134-450f-a9c0-0ce2e468d80e"},
		types.Task{Guid: utils.GenerateUUID(), Description: "Shortlink in express + go", ProjectId: "c889b8b7-e134-450f-a9c0-0ce2e468d80e"},
		types.Task{Guid: utils.GenerateUUID(), Description: "Document manager in express + go", ProjectId: "c889b8b7-e134-450f-a9c0-0ce2e468d80e"},*/
	}

	//project1 := types.Project{Guid: utils.GenerateUUID(), Title: "CambioMz", Description: "Unnoficial BCI Exchange API"}
	//project2 := types.Project{Guid: utils.GenerateUUID(), Title: "XRTLink", Description: "Shortlink in express + go"}
	//project3 := types.Project{Guid: utils.GenerateUUID(), Title: "Biblioteka", Description: "Document manager in express + go"}
	//persistance.Model(&types.Project{}).Create(project1)

	for _, value := range subtasks {
		persistance.Model(&types.Subtask{}).Create(value)
	}

}
