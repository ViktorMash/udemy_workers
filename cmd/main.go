package main

import (
	"fmt"
	"sync"
	"time"
	"workers/user"
)

const (
	workersCount = 100
	usersCount   = 100
)

func main() {

	startTime := time.Now()

	jobsNewUser := make(chan int, usersCount)          // Канал для задания ID пользователей
	resultsNewUser := make(chan user.User, usersCount) // Канал для результатов

	var wg sync.WaitGroup

	users := make([]user.User, usersCount) // создаем срез с длиной равной usersCount, юзеров раскладываем по индексу, для сортировки

	// создаем ожидающие воркеры (запущенные горутины) они ничего не будут делать, пока в канал не поступят данные
	startWorkersNewUser(users, workersCount, jobsNewUser, &wg)

	// отправляем задания (ID пользователей в канал jobs) и уже ожидающие воркеры (запущенные горутины) начинают их обрабатывать
	sendJobs(usersCount, jobsNewUser)

	wg.Wait() // Ожидаем завершения всех воркеров

	close(resultsNewUser) // Закрываем канал результатов после завершения всех воркеров

	// === Вывод результатов ===
	fmt.Printf("\n\nELAPSED TIME: %0.2f сек", time.Since(startTime).Seconds())

	for _, u := range users {
		fmt.Printf("\nUser ID: %d, Email: %s", u.Id, u.Email)
	}
}

// startWorkers запускает воркеров
func startWorkersNewUser(users []user.User, workersCount int, jobsNewUser <-chan int, wg *sync.WaitGroup) {
	for i := 0; i < workersCount; i++ {
		wg.Add(1) // Увеличиваем счётчик горутин
		go workerNewUser(users, jobsNewUser, wg)
	}
}

// sendJobs отправляет задания (ID юзеров) в канал
func sendJobs(usersCount int, jobsNewUser chan<- int) {
	for i := 1; i <= usersCount; i++ {
		jobsNewUser <- i
	}
	close(jobsNewUser) // Закрываем jobs, чтобы воркеры завершили свою работу
}

// workerNewUser читает задания из канала jobs, обрабатывает их и создает коного пользователя + файл
func workerNewUser(users []user.User, jobsNewUser <-chan int, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счётчик горутин в WaitGroup при завершении, даже в случае ошибки

	for userID := range jobsNewUser {
		user.GenerateUser(users, userID) // Создаем юзера в слайсе
		u := users[userID-1]             // Получаем только что созданного юзера
		user.SaveUserInfo(u)             // Сохраняем данные в файл
	}
}

// w/o workers ELAPSED TIME: 111.51 сек
// 100 workers ELAPSED TIME: 1.49 сек
