// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag at
// 2020-03-11 20:24:10.71013 +0900 KST m=+0.035090125

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
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "account email",
                        "name": "email",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "account service provider",
                        "name": "provider",
                        "in": "query"
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
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "parameters": [
                    {
                        "description": "Update Account data",
                        "name": "UpdateAccount",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/body.UpdateAccount"
                        }
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
            "post": {
                "description": "create account",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Account"
                ],
                "parameters": [
                    {
                        "description": "Create Account data",
                        "name": "CreateAccount",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/body.CreateAccount"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Account"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "delete account",
                "tags": [
                    "Account"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Account"
                        }
                    }
                }
            }
        },
        "/files": {
            "post": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "create file",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "file usage",
                        "name": "usage",
                        "in": "formData",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "Profile image file",
                        "name": "image",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "201": {}
                }
            }
        },
        "/files/{id}": {
            "get": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Files"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "file Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.File"
                        }
                    }
                }
            }
        },
        "/profiles": {
            "get": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profiles"
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Profile"
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
                "description": "update profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profiles"
                ],
                "parameters": [
                    {
                        "description": "update profile data",
                        "name": "UpdateProfile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/body.UpdateProfile"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Profile"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "AccessToken": []
                    }
                ],
                "description": "create profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profiles"
                ],
                "parameters": [
                    {
                        "description": "Create Profile data",
                        "name": "CreateProfile",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "object",
                            "$ref": "#/definitions/body.CreateProfile"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Profile"
                        }
                    }
                }
            }
        },
        "/profiles/{id}": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Profiles"
                ],
                "parameters": [
                    {
                        "type": "string",
                        "description": "profile id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Profile"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "body.CreateAccount": {
            "type": "object",
            "required": [
                "email",
                "provider"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "test@gmail.com"
                },
                "fcmToken": {
                    "type": "string",
                    "example": "fcmToken"
                },
                "gender": {
                    "type": "string",
                    "example": "male"
                },
                "interestedField": {
                    "type": "string",
                    "example": "develop"
                },
                "interestedFieldDetail": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "web",
                        "server"
                    ]
                },
                "password": {
                    "type": "string",
                    "example": "password"
                },
                "provider": {
                    "type": "string",
                    "example": "gmail"
                },
                "socialId": {
                    "type": "string",
                    "example": "socialId"
                }
            }
        },
        "body.CreateProfile": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string",
                    "example": "test@gmail.com"
                },
                "gender": {
                    "type": "string",
                    "example": "male"
                },
                "interestedField": {
                    "type": "string",
                    "example": "develop"
                },
                "interestedFieldDetail": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "web",
                        "server"
                    ]
                }
            }
        },
        "body.UpdateAccount": {
            "type": "object",
            "properties": {
                "fcmToken": {
                    "type": "string",
                    "example": "fcmToken"
                },
                "password": {
                    "type": "string",
                    "example": "password"
                }
            }
        },
        "body.UpdateProfile": {
            "type": "object",
            "properties": {
                "fileId": {
                    "type": "string",
                    "example": "fileId"
                },
                "interestedField": {
                    "type": "string",
                    "example": "develop"
                },
                "interestedFieldDetail": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "web",
                        "server"
                    ]
                }
            }
        },
        "model.Account": {
            "type": "object",
            "properties": {
                "accessToken": {
                    "type": "string",
                    "example": "accesstoken"
                },
                "createdAt": {
                    "type": "string",
                    "example": "2019-12-23 12:27:37"
                },
                "email": {
                    "type": "string",
                    "example": "test@gmail.com"
                },
                "id": {
                    "type": "string",
                    "example": "accountId"
                },
                "provider": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string",
                    "example": "2019-12-23 12:27:37"
                }
            }
        },
        "model.File": {
            "type": "object",
            "properties": {
                "accountId": {
                    "type": "string",
                    "example": "389df385-ccaa-49c1-aee2-698ba1191857"
                },
                "createdAt": {
                    "type": "string",
                    "example": "2019-12-23 12:27:37"
                },
                "id": {
                    "type": "string",
                    "example": "389df385-ccaa-49c1-aee2-698ba1191857"
                },
                "imageUrl": {
                    "type": "string",
                    "example": "profile.image_url.com"
                },
                "usage": {
                    "type": "string",
                    "example": "profile"
                }
            }
        },
        "model.Profile": {
            "type": "object",
            "properties": {
                "accountId": {
                    "type": "string",
                    "example": "accountId"
                },
                "createdAt": {
                    "type": "string",
                    "example": "2019-12-23 12:27:37"
                },
                "gender": {
                    "type": "string"
                },
                "id": {
                    "type": "string",
                    "example": "profileId"
                },
                "imageUrl": {
                    "type": "string",
                    "example": "profile.image_url.com"
                },
                "interestedField": {
                    "type": "string",
                    "example": "develop"
                },
                "interestedFieldDetail": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    },
                    "example": [
                        "web",
                        "server"
                    ]
                },
                "updatedAt": {
                    "type": "string",
                    "example": "2019-12-23 12:27:37"
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
