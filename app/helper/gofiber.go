package helper

import (
	"github.com/geraldie900/todo-app/config/validator"
	"github.com/gofiber/fiber/v2"
)

// BodyParseAndValidateStruct parse request body validates the input struct
func BodyParseAndValidateStruct(ctx *fiber.Ctx, payload interface{}) error {
	var err error
	// parse using gofiber ctx.BodyParse
	if err = ctx.BodyParser(payload); err != nil {
		return err
	}

	// validate struct
	if err = validator.StructValidator(payload); err != nil {
		return err
	}

	return err
}
