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
        "/contagem-pessoas": {
            "get": {
                "description": "find one person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "people"
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/pessoas": {
            "get": {
                "description": "find one person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "people"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "search",
                        "name": "t",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_patrickchagastavares_rinha-backend_internal_entities.Person"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_patrickchagastavares_rinha-backend_internal_entities.HttpErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/github_com_patrickchagastavares_rinha-backend_internal_entities.HttpErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "post": {
                "description": "Create one person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "people"
                ],
                "parameters": [
                    {
                        "description": "create new person",
                        "name": "house",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/github_com_patrickchagastavares_rinha-backend_internal_entities.PersonRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/github_com_patrickchagastavares_rinha-backend_internal_entities.Person"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_patrickchagastavares_rinha-backend_internal_entities.HttpErr"
                        }
                    },
                    "422": {
                        "description": "Unprocessable Entity",
                        "schema": {
                            "$ref": "#/definitions/github_com_patrickchagastavares_rinha-backend_internal_entities.HttpErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/pessoas/:id": {
            "get": {
                "description": "find one person",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "people"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "person id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_patrickchagastavares_rinha-backend_internal_entities.Person"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/github_com_patrickchagastavares_rinha-backend_internal_entities.HttpErr"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/github_com_patrickchagastavares_rinha-backend_internal_entities.HttpErr"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "github_com_patrickchagastavares_rinha-backend_internal_entities.HttpErr": {
            "type": "object",
            "properties": {
                "http_code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "github_com_patrickchagastavares_rinha-backend_internal_entities.Person": {
            "type": "object",
            "properties": {
                "apelido": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "nascimento": {
                    "type": "string"
                },
                "nome": {
                    "type": "string"
                },
                "stack": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "github_com_patrickchagastavares_rinha-backend_internal_entities.PersonRequest": {
            "type": "object",
            "required": [
                "apelido",
                "nascimento",
                "nome"
            ],
            "properties": {
                "apelido": {
                    "type": "string",
                    "maxLength": 32
                },
                "nascimento": {
                    "type": "string"
                },
                "nome": {
                    "type": "string",
                    "maxLength": 100
                },
                "stack": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
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
