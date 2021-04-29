package findingpeople

import (
	"github.com/gofiber/fiber"
	"strconv"
)

type person struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Job  string `json:"job"`
}

var people = []*person{
	{
		Id:  1,
		Name: "John",
		Job: "doctor",
	},
	{
		Id: 2,
		Name: "Jane",
		Job: "doctor",
	},
	{
		Id: 3,
		Name: "Rafa",
		Job: "lawyer",
	},
	{
		Id: 4,
		Name: "John",
		Job: "lawyer",
	},
}

func FindOne(ctx *fiber.Ctx) {
	paramsId := ctx.Params("input")
	id, err := strconv.Atoi(paramsId)

	var paramsName string
	if err != nil {
		paramsName = ctx.Params("input")
	}

	for _, person := range people{
		if id == person.Id || paramsName == person.Name {
			ctx.Status(fiber.StatusOK).JSON(person)
			return
		}
	}

	ctx.Status(fiber.StatusNotFound)
}

func FindWithTwo(ctx *fiber.Ctx) {
	paramsJob := ctx.Params("job")
	paramsName := ctx.Params("name")

	for _, person :=  range people {
		if person.Job == paramsJob && person.Name == paramsName {
			ctx.Status(fiber.StatusOK).JSON(person)
			return
		}
	}

	ctx.Status(fiber.StatusNotFound)
}

func ShowAll(ctx *fiber.Ctx) {
	ctx.Status(fiber.StatusOK).JSON(people)
}

func CreteNew(ctx *fiber.Ctx) {
	type request struct {
		Name string `json:"name"`
		Job string `json:"job"`
	}

	var body request

	err:= ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse json",
		})
		return
	}

	newPerson := &person{
		Id: len(people)+1,
		Name: body.Name,
		Job: body.Job,
	}

	people = append(people, newPerson)

	ctx.Status(fiber.StatusCreated).JSON(newPerson)
}