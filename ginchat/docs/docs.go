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
        "/index": {
            "get": {
                "tags": [
                    "首页"
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
        "/user/createUser": {
            "get": {
                "tags": [
                    "用户模块"
                ],
                "summary": "新增用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "二次密码",
                        "name": "repassword",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/deleteUser": {
            "get": {
                "tags": [
                    "用户模块"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID",
                        "name": "id",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/findUserByNameAndPsw": {
            "post": {
                "tags": [
                    "用户模块"
                ],
                "summary": "验证账号和密码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\",\"data}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/getUserList": {
            "get": {
                "tags": [
                    "用户模块"
                ],
                "summary": "用户列表",
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/user/updateUser": {
            "post": {
                "tags": [
                    "用户模块"
                ],
                "summary": "修改用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id",
                        "name": "id",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "用户名",
                        "name": "name",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "密码",
                        "name": "password",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "二次密码",
                        "name": "rePassword",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "电话号码",
                        "name": "phone",
                        "in": "formData"
                    },
                    {
                        "type": "string",
                        "description": "邮箱",
                        "name": "email",
                        "in": "formData"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "code\",\"message\"}",
                        "schema": {
                            "type": "string"
                        }
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
