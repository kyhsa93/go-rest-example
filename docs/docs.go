// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-02-02 14:09:38.94477 +0900 KST m=+0.051240111

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "license": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/accounts": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "account email",
                        "name": "email",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "account service provider",
                        "name": "provider",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "account password (email provider only)",
                        "name": "password",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "account social_id",
                        "name": "social_id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Token"
                        }
                    }
                }
            },
            "post": {
                "description": "create account",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "account email address",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "login service provider",
                        "name": "provider",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user's gender male of female",
                        "name": "gender",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Profile image file",
                        "name": "image",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "socialId when use social login",
                        "name": "social_id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "need if don't use social login",
                        "name": "password",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "201": {}
                }
            }
        },
        "/accounts/{id}": {
            "get": {
                "security": [
                    {
                        "AccessToken": []
                    },
                    {
                        "RefreshToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "account id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Account"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "update account",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Accounts"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "accountId",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "account email address",
                        "name": "email",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "login service provider",
                        "name": "provider",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "user's gender male of female",
                        "name": "gender",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Profile image file",
                        "name": "image",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "socialId when use social login",
                        "name": "social_id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "need if don't use social login",
                        "name": "password",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {}
                }
            },
            "delete": {
                "security": [
                    {
                        "AccessToken": []
                    },
                    {
                        "RefreshToken": []
                    }
                ],
                "description": "delete account by id",
                "tags": [
                    "Accounts"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "account Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        }
    },
    "definitions": {
        "model.Account": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string",
                    "example": "2019-12-23 12:27:37"
                },
                "email": {
                    "type": "string",
                    "example": "test@gmail.com"
                },
                "gender": {
                    "type": "string",
                    "example": "male"
                },
                "id": {
                    "type": "string",
                    "example": "389df385-ccaa-49c1-aee2-698ba1191857"
                },
                "image_url": {
                    "type": "string",
                    "example": "profile.image_url.com"
                },
                "provider": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string",
                    "example": "2019-12-23 12:27:37"
                }
            }
        },
        "model.Token": {
            "type": "object",
            "properties": {
                "access": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "AccessToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
