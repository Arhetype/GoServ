// Package docs Code generated by swaggo/swag. DO NOT EDIT
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
        "/api/v1/admins/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieves admin status and greets the admin or informs if the user is not an admin.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Get admin information",
                "responses": {
                    "200": {
                        "description": "Greeting message",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Something went wrong.",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/login": {
            "post": {
                "description": "Authenticate user and return a JWT token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Sign in user",
                "parameters": [
                    {
                        "description": "User login and password",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.SignInBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "token",
                        "schema": {
                            "$ref": "#/definitions/model.SignInResponse"
                        }
                    },
                    "400": {
                        "description": "incorrect request structure",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "couldn't create the token",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/auth/register": {
            "post": {
                "description": "Create a new user account and return the user ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "Register a new user",
                "parameters": [
                    {
                        "description": "User login and password",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.RegisterBody"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "account id",
                        "schema": {
                            "$ref": "#/definitions/model.RegisterResponse"
                        }
                    },
                    "400": {
                        "description": "incorrect request structure",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "couldn't create the user",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/users/": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Retrieve a list of users with optional filters.",
                "tags": [
                    "songs"
                ],
                "summary": "Get users",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Number of users per page",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "Filter by age",
                        "name": "age",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Filter by name",
                        "name": "name",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "List of users",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.User"
                            }
                        }
                    },
                    "400": {
                        "description": "Invalid query parameters",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "post": {
                "description": "This function accepts a JSON object of the user and adds it to the system.",
                "tags": [
                    "songs"
                ],
                "summary": "Add a new user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.AddUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ID of the added user",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Incorrect fields in request",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/v1/users/{id}": {
            "get": {
                "description": "Retrieve a user based on the provided user ID.",
                "tags": [
                    "songs"
                ],
                "summary": "Get user by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User details",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Invalid user ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "put": {
                "description": "This function updates user information based on the provided ID and data.",
                "tags": [
                    "songs"
                ],
                "summary": "Update an existing user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated user data",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.UpdateUserSwagger"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Updated user information",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "400": {
                        "description": "Incorrect fields or invalid user ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "description": "This function deletes a user based on the provided ID.",
                "tags": [
                    "songs"
                ],
                "summary": "Delete an existing user",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Indicates if the user was successfully deleted",
                        "schema": {
                            "type": "boolean"
                        }
                    },
                    "400": {
                        "description": "Invalid user ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.AddUser": {
            "type": "object",
            "required": [
                "age",
                "name"
            ],
            "properties": {
                "age": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.RegisterBody": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.RegisterResponse": {
            "type": "object",
            "properties": {
                "account_id": {
                    "type": "integer"
                }
            }
        },
        "model.SignInBody": {
            "type": "object",
            "properties": {
                "login": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.SignInResponse": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                }
            }
        },
        "model.UpdateUserSwagger": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "age": {
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Users API",
	Description:      "This is an API server for working with users",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
