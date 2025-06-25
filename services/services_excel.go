package services

import (
	"IssueTrackerAPI/database"
	m "IssueTrackerAPI/models"
	"strings"

	"github.com/xuri/excelize/v2"
)

func ExportProjectExcel(projectID uint) (*excelize.File, error) {
	db := database.DBConn

	// 1. หาโปรเจกต์
	var project m.Project
	if err := db.First(&project, projectID).Error; err != nil {
		return nil, err
	}

	// 2. หา project_employees ตาม projectID
	var projectEmployees []m.ProjectEmployee
	if err := db.Where("project_id = ?", projectID).Find(&projectEmployees).Error; err != nil {
		return nil, err
	}

	// 3. ดึงรายชื่อพนักงานจาก employee_id
	var employeeNames []string
	for _, pe := range projectEmployees {
		var emp m.Employee
		if err := db.First(&emp, pe.EmployeeID).Error; err == nil {
			employeeNames = append(employeeNames, emp.Name)
		}
	}

	// 4. สร้าง Excel
	f := excelize.NewFile()
	sheet := "Sheet1"

	f.SetCellValue(sheet, "A1", "Project Name")
	f.SetCellValue(sheet, "B1", "Description")
	f.SetCellValue(sheet, "C1", "Employees")

	f.SetCellValue(sheet, "A2", project.Name)
	f.SetCellValue(sheet, "B2", project.Description)
	f.SetCellValue(sheet, "C2", strings.Join(employeeNames, ", "))

	return f, nil
}
