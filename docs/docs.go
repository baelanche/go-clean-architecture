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
        "/api/article": {
            "post": {
                "description": "Article 을 생성합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Article"
                ],
                "summary": "Article 생성",
                "operationId": "create-article",
                "parameters": [
                    {
                        "description": "새로운 Article 정보",
                        "name": "article",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/controller.ArticleSwag"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/controller.ArticleSwag"
                        }
                    }
                }
            },
            "delete": {
                "description": "Article 을 제거합니다.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Article"
                ],
                "summary": "Article 제거",
                "operationId": "delete-article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "삭제할 Article 제목",
                        "name": "title",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/api/article/{title}": {
            "get": {
                "description": "단일 Article 을 조회합니다.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Article"
                ],
                "summary": "단일 Article 조회",
                "operationId": "get-article",
                "parameters": [
                    {
                        "type": "string",
                        "description": "조회할 Article 제목",
                        "name": "title",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ArticleSwag"
                        }
                    }
                }
            }
        },
        "/api/articles": {
            "get": {
                "description": "Article 리스트를 조회합니다.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Article"
                ],
                "summary": "Article 리스트 조회",
                "operationId": "get-articles",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/controller.ArticleSwag"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.ArticleSwag": {
            "type": "object",
            "properties": {
                "body": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "localhost:8080",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "go-clean-architecture",
	Description:      "go-clean-architecture",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
