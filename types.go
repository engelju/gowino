package main

import (
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUser(name, email string) *User {
	return &User{
		ID:    rand.Intn(1000),
		Name:  name,
		Email: email,
	}
}

func NewUserWithId(id int, name, email string) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}

func (u *User) Validate() error {
	if u.Name == "" {
		return fmt.Errorf("name is required")
	}
	if u.Email == "" {
		return fmt.Errorf("email is required")
	}
	return nil
}

type Wine struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Vintage int    `json:"vintage"`
	Winery  string `json:"winery"`
	Region  string `json:"region"`
	Country string `json:"country"`
}

func NewWine(name, winery, region, country string, vintage int) *Wine {
	return &Wine{
		ID:      rand.Intn(1000),
		Name:    name,
		Vintage: vintage,
		Winery:  winery,
		Region:  region,
		Country: country,
	}
}

func NewineWithId(id int, name, winery, region, country string, vintage int) *Wine {
	return &Wine{
		ID:      id,
		Name:    name,
		Vintage: vintage,
		Winery:  winery,
		Region:  region,
		Country: country,
	}
}

type DrinkEntry struct {
	ID          int       `json:"id"`
	WineID      int       `json:"wine_id"`
	DateDrank   time.Time `json:"date_drank"`
	DrankWith   string    `json:"drank_with,omitempty"`
	DrankAt     string    `json:"drank_at,omitempty"`
	Rating      int       `json:"rating"`
	Description string    `json:"description,omitempty"`
}

func NewDrinkEntry(wineID int, dateDrank time.Time, drankWith string, drankAt string, rating int, description string) *DrinkEntry {
	return &DrinkEntry{
		ID:          rand.Intn(1000),
		WineID:      wineID,
		DateDrank:   dateDrank,
		DrankWith:   drankWith,
		DrankAt:     drankAt,
		Rating:      rating,
		Description: description,
	}
}
