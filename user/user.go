package user

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
	"workers/logs"
	"workers/utils"
)

const userLogsFolder = "users_logs"

type User struct {
	Id    int
	Email string
	Logs  []logs.Log
}

func (u User) GetActivityInfo() string {
	output := fmt.Sprintf("UID: %d; Email: %s; \nActivity log:\n", u.Id, u.Email)
	for index, item := range u.Logs {
		output += fmt.Sprintf("%d. [%s] at %s\n", index, item.Action, item.Timestamp.Format(time.RFC3339))
	}
	return output
}

// GenerateUsers записывает юзера в следующий пустой элемент массива
func GenerateUser(users []User, userID int) []User {

	index := userID - 1
	users[index] = User{
		Id:    userID,
		Email: fmt.Sprintf("user%d@company.com", userID),
		Logs:  logs.GenerateLogs(rand.Intn(1000)),
	}

	fmt.Printf("\nUser %d generated", userID)
	time.Sleep(100 * time.Millisecond)

	return users
}

func SaveUserInfo(user User) {
	fmt.Printf("\nWriting file for userID %d", user.Id)

	utils.CreateFolder(userLogsFolder)

	filename := fmt.Sprintf("%s/uid%d.txt", userLogsFolder, user.Id)

	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString(user.GetActivityInfo())
	time.Sleep(time.Second)
}
