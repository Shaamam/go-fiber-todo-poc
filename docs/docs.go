// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `
{
    "swagger": "2.0",
    "info": {
        "description": "API for managing Todo items",
        "title": "Todo API",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Your Name",
            "email": "your.email@example.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "1.0"
    },
    "host": "localhost:8080",
    "basePath": "/api/v1",
    "paths": {
        "/todos": {
            "get": {
                "description": "Retrieve all todos",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "Get all todos",
                "responses": {
                    "200": {
                        "description": "A list of todos",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/Todo"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "Create a todo",
                "parameters": [
                    {
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/TodoInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "The created todo",
                        "schema": {
                            "$ref": "#/definitions/Todo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/todos/{id}": {
            "put": {
                "description": "Update an existing todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "Update a todo",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "description": "The ID of the todo to update"
                    },
                    {
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/TodoInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "The updated todo",
                        "schema": {
                            "$ref": "#/definitions/Todo"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a todo",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "todos"
                ],
                "summary": "Delete a todo",
                "parameters": [
                    {
                        "name": "id",
                        "in": "path",
                        "required": true,
                        "type": "string",
                        "description": "The ID of the todo to delete"
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No content"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "ErrorResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "Todo": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "completed": {
                    "type": "boolean"
                }
            }
        },
        "TodoInput": {
            "type": "object",
            "properties": {
                "title": {
                    "type": "string"
                },
                "completed": {
                    "type": "boolean"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}
`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Todo API",
	Description:      "API for managing Todo items",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
