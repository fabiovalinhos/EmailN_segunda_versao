package main

import (
	"emailn/internal/domain/campaign"

	"github.com/go-playground/validator/v10"
)

func main() {
	contacts := []campaign.Contact{{Email: "Henti"}, {Email: "Henti2"}}
	campaign := campaign.Campaign{Contacts: contacts}
	validate := validator.New()
	err := validate.Struct(campaign)
	if err == nil {
		println("Nenhum erro encontrado")
	} else {
		validationErrors := err.(validator.ValidationErrors)
		for _, v := range validationErrors {

			switch v.Tag() {
			case "required":
				println("O campo " + v.StructField() + " é obrigatório")
			case "email":
				println("O campo " + v.StructField() + " deve ser um email válido")
			case "min":
				println("O campo " + v.StructField() + " deve ter no mínimo " + v.Param() + " caracteres")
			case "max":
				println("O campo " + v.StructField() + " deve ter no máximo " + v.Param() + " caracteres")
			}
		}
	}
}
