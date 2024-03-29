// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/task": {
            "post": {
                "description": "Adds a new task with the provided text.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Create a new task",
                "parameters": [
                    {
                        "description": "Task creation request body",
                        "name": "requestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/middleware.CreateTaskForm"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Successfully created task with the provided text, returns task which created.",
                        "schema": {
                            "$ref": "#/definitions/middleware.TaskResponse"
                        }
                    }
                }
            }
        },
        "/task/{id}": {
            "put": {
                "description": "Update the task with the specified ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Update a task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Fields and values to update",
                        "name": "task",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/middleware.UpdateTaskForm"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task updated successfully",
                        "schema": {
                            "$ref": "#/definitions/middleware.TaskResponse"
                        }
                    },
                    "400": {
                        "description": "Missing ID or incorrect fields",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "404": {
                        "description": "Task of ID not found",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "Deletes a task by its ID.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Delete a task",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID of the task to delete",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Task deleted successfully"
                    },
                    "400": {
                        "description": "Missing ID"
                    },
                    "403": {
                        "description": "Task of ID not found"
                    }
                }
            }
        },
        "/tasks": {
            "get": {
                "description": "Retrieves a list of tasks.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "List all tasks",
                "responses": {
                    "200": {
                        "description": "A list of tasks",
                        "schema": {
                            "$ref": "#/definitions/middleware.ListTaskResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "middleware.CreateTaskForm": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string"
                }
            }
        },
        "middleware.ListTaskResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.Task"
                    }
                }
            }
        },
        "middleware.TaskResponse": {
            "type": "object",
            "properties": {
                "result": {
                    "$ref": "#/definitions/models.Task"
                }
            }
        },
        "middleware.UpdateTaskForm": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "models.Task": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "status": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
