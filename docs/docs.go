// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/demo/v1/getUserInfo": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "获取用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "cookie",
                        "name": "cookie",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.GetUserInfoRes"
                        }
                    }
                }
            }
        },
        "/demo/v1/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "entity",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.LoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.LoginRes"
                        }
                    }
                }
            }
        },
        "/demo/v1/register": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "注册或更新用户信息",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "entity",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.RegisterOrUpdateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {}
                }
            }
        },
        "/demo/v1/test": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户"
                ],
                "summary": "测试一下",
                "parameters": [
                    {
                        "description": "请求参数",
                        "name": "entity",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user.DemoTestReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/user.DemoTestResp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "user.DemoTestReq": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "user.DemoTestResp": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                }
            }
        },
        "user.GetUserInfoRes": {
            "type": "object",
            "properties": {
                "nickname": {
                    "type": "string"
                },
                "passport": {
                    "type": "string"
                },
                "password": {
                    "type": "integer"
                }
            }
        },
        "user.LoginReq": {
            "type": "object",
            "properties": {
                "passport": {
                    "type": "string"
                },
                "password": {
                    "type": "integer"
                }
            }
        },
        "user.LoginRes": {
            "type": "object",
            "properties": {
                "sessionId": {
                    "type": "string"
                }
            }
        },
        "user.RegisterOrUpdateReq": {
            "type": "object",
            "properties": {
                "nickname": {
                    "type": "string"
                },
                "passport": {
                    "type": "string"
                },
                "password": {
                    "type": "integer"
                }
            }
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "gd-demo",
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
