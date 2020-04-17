package controllers

import "net/http"

type User struct {
}

func Users() *User {
	return &User{}
}

func (u *User) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get users..."))
}

func (u *User) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create user..."))
}

func (u *User) GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Get user..."))
}

func (u *User) UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Update user..."))
}

func (u *User) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Delete user..."))
}
