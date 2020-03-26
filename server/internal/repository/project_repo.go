package repository

import (
	"fmt"

	"github.com/kradnoel/teka/internal/persistance"
	"github.com/kradnoel/teka/internal/types"
	"github.com/kradnoel/teka/internal/utils"
)

func CreateProject(project types.Project) bool {

	persistance := persistance.GetPersistance()

	err := persistance.Model(&types.Project{}).Create(project).Error

	if err != nil {
		return false
	}

	return true
}

func DeleteProject(id string) bool {

	persistance := persistance.GetPersistance()
	tasks := make([]types.Task, 0)

	// Start transactiong
	tx := persistance.Begin()

	// Search for all task associated with the project
	err := tx.Model(&types.Task{}).Where("project_id = ?", utils.ParseToUUID(id)).Find(&tasks).Error

	if err != nil {
		tx.Rollback()
		return false
	}

	// Delete all subtasks associated with the tasks of the project
	for _, task := range tasks {
		err = tx.Model(&types.Subtask{}).Where("task_id = ?", task.Guid).Delete(&types.Subtask{}).Error
		if err != nil {
			tx.Rollback()
			return false
		}
	}

	// Delete all the tasks associated with the project
	err = tx.Model(&types.Task{}).Where("project_id = ?", utils.ParseToUUID(id)).Delete(&types.Task{}).Error

	if err != nil {
		tx.Rollback()
		return false
	}

	// Delete the project
	err = tx.Model(&types.Project{}).Where("guid = ?", utils.ParseToUUID(id)).Delete(&types.Project{}).Error

	if err != nil {
		tx.Rollback()
		return false
	}

	/*
		err := persistance.Model(&types.Project{}).Where("guid = ?", utils.ParseToUUID(id)).Delete(&types.Project{})

		if err != nil {
			return false
		}
	*/

	tx.Commit()
	return true

}

func GetProjects() []types.Project {

	persistance := persistance.GetPersistance()

	projects := make([]types.Project, 0)

	rows, err := persistance.Model(&types.Project{}).Rows()
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var project types.Project
		persistance.ScanRows(rows, &project)

		tasks := GetTasks(project.Guid.String())

		for _, task := range tasks {
			project.TimeSpent = project.TimeSpent + task.TimeSpent
		}

		projects = append(projects, project)
	}

	return projects

}

func GetProject(id string) *types.Project {
	persistance := persistance.GetPersistance()

	var project types.Project
	tasks := make([]types.Task, 0)
	//var spend_hours float64

	err := persistance.Model(&types.Project{}).Where("guid = ?", utils.ParseToUUID(id)).First(&project).Error
	// defer rows.Close()

	if err != nil {
		fmt.Println(err)
		return nil
	}

	tasks = GetTasks(id)
	for _, task := range tasks {
		project.TimeSpent = project.TimeSpent + task.TimeSpent
	}

	/*

		for rows.Next() {
			var task types.Task
			persistance.ScanRows(rows, &task)
			tasks = append(tasks, task)
		}*/

	return &project
}

func Seed() {
	persistance := persistance.GetPersistance()

	projects := []types.Project{
		types.Project{
			Guid:        utils.ParseToUUID("c889b8b7-e134-450f-a9c0-0ce2e468d80e"),
			Title:       "CambioMz",
			Description: "Unnoficial BCI Exchange API",
			Completed:   true,
		},
		types.Project{
			Guid:        utils.GenerateUUID(),
			Title:       "XRTLink",
			Description: "Shortlink in express + go",
			Completed:   true,
		},
		types.Project{
			Guid:        utils.GenerateUUID(),
			Title:       "Biblioteka",
			Description: "Document manager in express + go",
			Completed:   false,
		},
	}

	//project1 := types.Project{Guid: utils.GenerateUUID(), Title: "CambioMz", Description: "Unnoficial BCI Exchange API"}
	//project2 := types.Project{Guid: utils.GenerateUUID(), Title: "XRTLink", Description: "Shortlink in express + go"}
	//project3 := types.Project{Guid: utils.GenerateUUID(), Title: "Biblioteka", Description: "Document manager in express + go"}
	//persistance.Model(&types.Project{}).Create(project1)

	for _, value := range projects {
		persistance.Model(&types.Project{}).Create(value)
	}

}

func GetProjectStatus() types.ProjectStatus {

	var total_projects = 0
	var completed_projects = 0
	var projects_time = make([]float64, 0)

	/*rows, err := persistance.Model(&types.Project{}).Rows()
	defer rows.Close()

	if err != nil {
		fmt.Println(err)
	}

	for rows.Next() {
		var project types.Project
		persistance.ScanRows(rows, &project)

		if project.Completed == true {
			completed_projects = completed_projects + 1
		}
		total_projects = total_projects + 1
		fmt.Println("project time = ", project.TimeSpent)
		projects_time = append(projects_time, project.TimeSpent)
	}*/

	projects := GetProjects()

	for _, project := range projects {

		if project.Completed == true {
			completed_projects += 1
		}
		total_projects += 1
		fmt.Println("project time = ", project.TimeSpent)
		projects_time = append(projects_time, project.TimeSpent)
	}

	status := types.ProjectStatus{
		Total:       total_projects,
		AverageTime: utils.ListAverageValue(projects_time),
		Completed:   completed_projects,
	}

	return status
}
