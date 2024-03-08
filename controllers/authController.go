package controllers

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/bahati-hakizimana/e-learning-backend/database"
	"github.com/bahati-hakizimana/e-learning-backend/models"
	"github.com/gofiber/fiber/v2"
)

//validate Email function

func validateEmail(email string) bool{
	Re:= regexp.MustCompile(`[a-z0-9._%+\-]+@[a-z0-9._%+\-]+\.[a-z0-9._%+\-]`)
	return Re.MatchString(email)
}



func Register(c *fiber.Ctx) error{
	var data map[string] interface{}
    var userData models.User
	if err:=c.BodyParser(&data);err!=nil{
		fmt.Println("Un able to parse body")
	}

	//check if passwsord is validated

	if len(data["password"].(string))<=6{
		c.Status(400)
		return c.JSON(fiber.Map{
			"message" :"Password must be greater than 6 charactror",
		})
	}

	if !validateEmail(strings.TrimSpace(data["email"].(string))){
		c.Status(400)
		return c.JSON(fiber.Map{
			"message" :"Invalid Email address",
		})
	}

	//check if Email already exist in database

	database.DB.Where("email=?", strings.TrimSpace(data["email"].(string))).First(&userData)
	if userData.Id!=0{
		c.Status(400)
		return c.JSON(fiber.Map{
			"message" :"Email already Exist",
		})

	}

	user:=models.User{
		FirstName: data["first_name"].(string),
		LastName: data["last_name"].(string),
		Phone: data["phone"].(string),
		Email: strings.TrimSpace(data["email"].(string)),
	}

	user.SetPassword(data["password"].(string))
	err:=database.DB.Create(&user)

	if err!= nil{
		log.Println(err)
	}
	c.Status(200)
		return c.JSON(fiber.Map{
			"user":user,
			"message" :"Account created successful",
		})

		//info@easyandpossible.org 
}

func Login(c *fiber.Ctx) error {

	var  data map[string]string
	if err:=c.BodyParser(&data);err!=nil{
		fmt.Println("Un able to parser body")
	}

	var user models.User
	database.DB.Where("email=?", data["email"]).First(&user)

	if user.Id==0{
		c.Status(404)
		return c.JSON(fiber.Map{
			"message" : "Email doesn't exist, kindly create an account",
		})
	}

}
