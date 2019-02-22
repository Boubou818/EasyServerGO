package main

import (
	"html/template"
	"log"
	"net/http"
)

func main()  {

	http.HandleFunc("/", mainPage) // первый обработчик
	http.HandleFunc("/users", usersPage) // второй обработчик

	port := ":9090"
	println("Server listen on port", port)
	err := http.ListenAndServe(port, nil)
	if err !=nil {
		log.Fatal("ListenAndServe", err)
	}
}

type User struct { //Json для сервака, тут нет вопросов. Тут сделали тип
	FirstName string `json:"first_name"`
	LastName string	`json:"last_name"`
	IsFired bool
}

func mainPage(w http.ResponseWriter, r *http.Request)  {
	//user := User{"Шерлас", "Хорн"}
	//js, _ := json.Marshal(user)// преварщаем юзера в Json. Этот тип берет Json и конвертирует его в строку. _ это место ошибки, но мы предпологаем, что ошибки у нас нет, и ее мы обрабатывать не будем

	tmpl, err := template.ParseFiles("static/index.html")
	if err !=nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if err := tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	//w.Write(js) //передача текста JSON
}

func usersPage(w http.ResponseWriter, r *http.Request)  {
	users := []User{User{"Шерлас", "Хорн", false}, User{"Миша","Рэйдж", true}}
	//js, _ := json.Marshal(user)// преварщаем юзера в Json. Этот тип берет Json и конвертирует его в строку. _ это место ошибки, но мы предпологаем, что ошибки у нас нет, и ее мы обрабатывать не будем

	tmpl, err := template.ParseFiles("static/users.html")
	if err !=nil {
		http.Error(w, err.Error(), 400)
		return
	}

	if err := tmpl.Execute(w, users); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	//w.Write(js) //передача текста JSON
}