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
        "/category": {
            "get": {
                "description": "Get detailed information of all data category and product based on category",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "Get All Data Category",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.Category"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseErrorNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseErrorInternalServer"
                        }
                    }
                }
            }
        },
        "/category/input": {
            "post": {
                "description": "Add a new category to the system",
                "consumes": [
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Category"
                ],
                "summary": "post data Category",
                "parameters": [
                    {
                        "description": "Category Input",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.CategoryInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseErrorBadRequest"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseErrorInternalServer"
                        }
                    }
                }
            }
        },
        "/product": {
            "get": {
                "description": "Get detailed information of all data Product and detailProduct by Id Product",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "Get All Data Product",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.Product"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseErrorNotFound"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseErrorInternalServer"
                        }
                    }
                }
            }
        },
        "/product/input": {
            "post": {
                "description": "Add a new Product to the system",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Product"
                ],
                "summary": "post data Product",
                "parameters": [
                    {
                        "description": "Product Input",
                        "name": "category",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ProductInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseSuccess"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseErrorBadRequest"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseErrorInternalServer"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.CategoryInput": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                }
            }
        },
        "controller.ProductInput": {
            "type": "object",
            "properties": {
                "categoryId": {
                    "type": "integer"
                },
                "hargaProduct": {
                    "type": "number"
                },
                "image": {
                    "type": "string"
                },
                "namaProduct": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "controller.ResponseErrorBadRequest": {
            "type": "object",
            "properties": {
                "errorFields": {
                    "type": "array",
                    "items": {
                        "type": "object",
                        "additionalProperties": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "controller.ResponseErrorInternalServer": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "controller.ResponseErrorNotFound": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "controller.ResponseSuccess": {
            "type": "object",
            "properties": {
                "Success": {
                    "type": "string"
                }
            }
        },
        "dto.Category": {
            "type": "object",
            "required": [
                "category"
            ],
            "properties": {
                "category": {
                    "type": "string",
                    "maxLength": 100
                },
                "createdAt": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "product": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.Product"
                    }
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "dto.Product": {
            "type": "object",
            "required": [
                "categoryId",
                "hargaProduct",
                "image",
                "namaProduct",
                "type"
            ],
            "properties": {
                "categoryId": {
                    "type": "integer"
                },
                "created_at": {
                    "type": "string"
                },
                "hargaProduct": {
                    "type": "number"
                },
                "id": {
                    "type": "integer"
                },
                "image": {
                    "type": "string"
                },
                "namaProduct": {
                    "type": "string",
                    "maxLength": 100
                },
                "type": {
                    "type": "string",
                    "maxLength": 100
                },
                "updated_at": {
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
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
