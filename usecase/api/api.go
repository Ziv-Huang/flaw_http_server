package api

import (
	"flaw_http_server/domain/request"

	"github.com/gofiber/fiber/v2"
)

func Hello(c *fiber.Ctx) error {
	c.Status(200)
	return c.SendString("Hello, World 👋!")
}

func Heartbeat(c *fiber.Ctx) error {
	c.Status(200)
	return c.SendString("OK")
}

func GetDeviceInfo(c *fiber.Ctx) error {
	// TODO
	var result request.DeviceInfo
	result.IsBusy = true
	result.IsAlive = true
	c.Status(200)
	return c.JSON(result)
}

func CreateAJob(c *fiber.Ctx) error {
	// TODO
	var result request.Task
	result.TaskID = "1"
	result.JobStatus = request.Running
	result.Msg = "OK"
	result.Code = 200
	c.Status(200)
	return c.JSON(result)
}

func DeleteAJob(c *fiber.Ctx) error {
	// TODO
	var result request.Task
	result.TaskID = c.Params("id")
	result.JobStatus = request.Killed
	result.Msg = "OK"
	result.Code = 200
	c.Status(200)
	return c.JSON(result)
}

func GetResultByID(c *fiber.Ctx) error {
	// TODO
	var result request.Result
	result.TaskID = c.Params("id")
	result.JobStatus = request.Running
	result.Msg = "OK"
	result.Code = 200
	result.AnalysisResult = "OK"
	c.Status(200)
	return c.JSON(result)
}
