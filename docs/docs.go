// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "https://github.com/dbalucas",
            "url": "https://github.com/dbalucas",
            "email": "https://github.com/dbalucas"
        },
        "license": {
            "url": "https://github.com/dbalucas/go_crm-backend_udacity/blob/main/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/customers": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "slice"
                        }
                    }
                }
            },
            "put": {
                "description": "From POST request update an existing customer by its uuid if exists",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "Update an existing customer by its uuid",
                "parameters": [
                    {
                        "description": "update new Customer",
                        "name": "customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Customer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Customer"
                        }
                    }
                }
            },
            "post": {
                "description": "From POST request create new Customer if not exists",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "Add a new customer",
                "parameters": [
                    {
                        "description": "Add new Customer",
                        "name": "customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Customer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/main.Customer"
                        }
                    }
                }
            },
            "delete": {
                "description": "remove a single Customer by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customer"
                ],
                "summary": "Update an existing customer by its uuid",
                "parameters": [
                    {
                        "description": "delete new Customer",
                        "name": "customer",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/main.Customer"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/customers/{id}": {
            "get": {
                "description": "Retrieve single customer by ID in /customer/{id}",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "customers"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "customer"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "main.Customer": {
            "type": "object",
            "properties": {
                "contacted": {
                    "type": "boolean"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "phone": {
                    "type": "integer"
                },
                "role": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:3000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "CRM - GOlang API Documentation",
	Description:      "This is a sample crm-server and contains a final project of a udacity course",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
