package entity

import "github.com/lailiseptiandi/go-store-api/models"

type UserFormatter struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email,omitempty"`
	Address string `json:"address"`
	Roles   string `json:"roles,omitempty"`
	Token   string `json:"token,omitempty"`
}

func FormatUser(user models.User, Token string) UserFormatter {
	formatter := UserFormatter{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Roles:   user.Roles,
		Token:   Token,
	}

	return formatter
}

type FormateDetailUser struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Roles   string `json:"roles"`
}

func FormatDetailUser(user models.User) FormateDetailUser {
	formatter := FormateDetailUser{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Address: user.Address,
		Roles:   user.Roles,
	}
	return formatter
}

func FormatListUser(users []models.User) []UserFormatter {
	formatter := make([]UserFormatter, 0)
	for _, u := range users {
		userFormat := UserFormatter{
			ID:      u.ID,
			Name:    u.Name,
			Address: u.Address,
		}
		formatter = append(formatter, userFormat)
	}
	return formatter
}

// category

type CategoryFormater struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
}

func FormatCategory(category models.Category) CategoryFormater {
	formater := CategoryFormater{
		ID:   category.ID,
		Name: category.Name,
		Slug: category.Slug,
	}

	return formater
}

func FormatterCategory(categories []models.Category) []CategoryFormater {

	categoryFormatter := []CategoryFormater{}
	for _, category := range categories {
		categoriesFormatter := FormatCategory(category)
		categoryFormatter = append(categoryFormatter, categoriesFormatter)
	}

	return categoryFormatter
}
