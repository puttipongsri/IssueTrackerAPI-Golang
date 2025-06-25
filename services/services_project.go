package services

import (
	"IssueTrackerAPI/database"
	m "IssueTrackerAPI/models"
	"strings"
)

func CreateProject(name, description string, employeeIDs []uint, imagePath string) error {
	db := database.DBConn
	project := m.Project{
		Name:        name,
		Description: description,
		Image:       imagePath,
	}

	if err := db.Create(&project).Error; err != nil {
		return err
	}

	for _, empID := range employeeIDs {
		projectEmployee := m.ProjectEmployee{
			ProjectID:  project.ID,
			EmployeeID: empID, 
		}
		if err := db.Create(&projectEmployee).Error; err != nil {
			return err
		}
	}

	return nil
}
func GetProjects() ([]m.Project, error) {
	db := database.DBConn
	var projects []m.Project
	db.Select("id", "name", "description", "image").Find(&projects)
	for i, p := range projects {
		projects[i].Image = strings.Replace(p.Image, "./fileimg/", "/images/", 1)
	}
	return projects, nil
}
func GetProjectByID(ID uint) (*m.Project, error) {
	db := database.DBConn
	var project m.Project
	if err := db.Where("id = ?", ID).First(&project).Error; err != nil {
		return nil, err
	}
	project.Image = strings.Replace(project.Image, "./fileimg/", "/images/", 1)
	return &project, nil
}
func UpdateProject(id uint, name, description, imagePath string) error {
	db := database.DBConn
	var project m.Project

	if err := db.First(&project, id).Error; err != nil {
		return err
	}

	project.Name = name
	project.Description = description
	if imagePath != "" {
		project.Image = imagePath
	}

	return db.Save(&project).Error
}
func DeleteProject(id uint) error {
	db := database.DBConn
	var project m.Project
	var projectEmployee m.ProjectEmployee
	if err := db.First(&project, id).Error; err != nil {
		return err
	}
	if err := db.Where("project_id = ?", id).Delete(&projectEmployee).Error; err != nil {
		return err
	}
	return db.Delete(&project).Error
}
