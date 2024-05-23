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
        "/exchange": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Calculate"
                ],
                "summary": "Calculate event data by Id",
                "parameters": [
                    {
                        "description": "Exchange Request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.ExchangeRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "request"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponce"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.ErrorResponce"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handlers.ErrorResponce": {
            "type": "object",
            "properties": {
                "error": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "handlers.ExchangeRequest": {
            "type": "object",
            "properties": {
                "amount": {
                    "type": "integer"
                },
                "banknotes": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Exchanger API",
	Description:      "Exchanger app server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	// LeftDelim:        "{{",
	// RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
