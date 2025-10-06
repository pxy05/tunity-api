package sqlite

// THIS FILE SHOULD ONLY BE USED TO INITIALIZE THE DATABASE NOTHING IN THIS FILE SHOULD BE USED ELSEWHERE.

import (
	"log"
	"time"
	"tunity-api/tools/structures"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)



var db *gorm.DB

func initDB() {
	var err error
	db, err = gorm.Open(sqlite.Open("tunity_test.db"), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return
	}
}

func StartDB() {
	initDB()
	db.AutoMigrate(&structures.User{}, &structures.Application{}, &structures.Interview{})
	seedDatabase()
}

func seedDatabase() {
	var userCount int64
	db.Model(&structures.User{}).Count(&userCount)
	if userCount > 0 {
		return
	}

	users := []structures.User{
		{
			Username:  "john_doe",
			Email:     "john.doe@email.com",
			Password:  "hashed_password_1",
			FirstName: "John",
			LastName:  "Doe",
		},
		{
			Username:  "jane_smith",
			Email:     "jane.smith@email.com",
			Password:  "hashed_password_2",
			FirstName: "Jane",
			LastName:  "Smith",
		},
		{
			Username:  "mike_johnson",
			Email:     "mike.johnson@email.com",
			Password:  "hashed_password_3",
			FirstName: "Mike",
			LastName:  "Johnson",
		},
		{
			Username:  "sarah_wilson",
			Email:     "sarah.wilson@email.com",
			Password:  "hashed_password_4",
			FirstName: "Sarah",
			LastName:  "Wilson",
		},
		{
			Username:  "alex_brown",
			Email:     "alex.brown@email.com",
			Password:  "hashed_password_5",
			FirstName: "Alex",
			LastName:  "Brown",
		},
	}

	for i := range users {
		db.Create(&users[i])
	}

	now := time.Now()
	companies := []string{"Google", "Microsoft", "Apple", "Amazon", "Meta", "Netflix", "Tesla", "Uber", "Airbnb", "Spotify", "Twitter", "LinkedIn", "Adobe", "Salesforce", "Oracle", "IBM", "Intel", "NVIDIA", "AMD", "Cisco", "VMware", "Atlassian", "Slack", "Zoom", "Shopify"}
	jobTitles := []string{"Software Engineer", "Senior Software Engineer", "Frontend Developer", "Backend Developer", "Full Stack Developer", "DevOps Engineer", "Data Scientist", "Product Manager", "UX Designer", "UI Designer", "QA Engineer", "Security Engineer", "Mobile Developer", "Cloud Engineer", "Machine Learning Engineer"}
	locations := []string{"San Francisco, CA", "New York, NY", "Seattle, WA", "Austin, TX", "Boston, MA", "Los Angeles, CA", "Chicago, IL", "Denver, CO", "Remote", "Hybrid"}
	statuses := []string{"applied", "interview", "offer", "rejected", "offer_accepted"}

	applicationID := 1
	for _, user := range users {
		for i := 0; i < 5; i++ {
			appliedDate := now.AddDate(0, 0, -30+i*5)
			deadline := appliedDate.AddDate(0, 0, 14) 
			
			application := structures.Application{
				ID:                      string(rune(applicationID)),
				UserID:                  user.ID,
				Company:                 companies[applicationID%len(companies)],
				Status:                  statuses[applicationID%len(statuses)],
				Job_Title:               jobTitles[applicationID%len(jobTitles)],
				Job_Location:            locations[applicationID%len(locations)],
				Application_URL:         &[]string{"https://company.com/jobs/" + string(rune(applicationID))}[0],
				Application_Applied_Date: &appliedDate,
				Application_Deadline:    &deadline,
			}

			if applicationID%3 == 0 {
				notes := "Great company culture and benefits"
				application.Job_Notes = &notes
			}

			if application.Status == "rejected" {
				rejectedDate := appliedDate.AddDate(0, 0, 7)
				application.Application_Rejected_Date = &rejectedDate
			}

			if application.Status == "rejected" && applicationID%2 == 0 {
				failedDate := appliedDate.AddDate(0, 0, 10)
				application.Interview_Failed_Date = &failedDate
			}

			db.Create(&application)
			applicationID++
		}
	}
}

