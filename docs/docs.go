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
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Gidhan Bagus Algary",
            "email": "gidhanbagusalgary@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/login": {
            "post": {
                "description": "Login user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Login user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/users.LoginUserInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    }
                }
            }
        },
        "/orders": {
            "get": {
                "description": "Get all orders",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Show all orders",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "Create new order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Create new order",
                "parameters": [
                    {
                        "description": "Order",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/orders.SaveOrderInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    }
                }
            }
        },
        "/orders/person/{order_id}": {
            "get": {
                "description": "Get order person by order id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Show order person by order id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "order_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    }
                }
            }
        },
        "/orders/{order_id}": {
            "get": {
                "description": "Get orders by user id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Show orders by user id",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "order_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    }
                }
            },
            "put": {
                "description": "Update order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Update order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "order_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Order",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/orders.UpdateOrderInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete order",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "orders"
                ],
                "summary": "Delete order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "order_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    }
                }
            }
        },
        "/register": {
            "post": {
                "description": "Register new user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Register new user",
                "parameters": [
                    {
                        "description": "User",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/users.RegisterUserInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/helper.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "helper.Meta": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "helper.Response": {
            "type": "object",
            "properties": {
                "data": {},
                "meta": {
                    "$ref": "#/definitions/helper.Meta"
                }
            }
        },
        "orders.SaveItemInput": {
            "type": "object",
            "required": [
                "description",
                "itemCode",
                "quantity"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "itemCode": {
                    "type": "string"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "orders.SaveOrderInput": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/orders.SaveItemInput"
                    }
                }
            }
        },
        "orders.UpdateItemInput": {
            "type": "object",
            "required": [
                "description",
                "itemCode",
                "lineItemId",
                "quantity"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "itemCode": {
                    "type": "string"
                },
                "lineItemId": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "orders.UpdateOrderInput": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/orders.UpdateItemInput"
                    }
                }
            }
        },
        "users.LoginUserInput": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "users.RegisterUserInput": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "127.0.0.1:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{"http"},
	Title:            "Go Hacktiv8 Assignment 2 API Documentation",
	Description:      "This is a sample server for a store.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
