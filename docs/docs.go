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
        "/jenkins/jkDingTalk": {
            "get": {
                "description": "构建环境;构建人,构建的项目名,构建的分支名,构建的URL,构建ID,执行时间,状态",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "jenkins相关接口"
                ],
                "summary": "jenkins构建钉钉通知接口",
                "parameters": [
                    {
                        "type": "string",
                        "name": "branch",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "name": "id",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "required为必须",
                        "name": "name",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "name": "url",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controller.ResponseData"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controller.ResCode": {
            "type": "integer",
            "enum": [
                1000,
                1001,
                1002
            ],
            "x-enum-varnames": [
                "CodeSuccess",
                "CodeInvalidParam",
                "CodeServerBusy"
            ]
        },
        "controller.ResponseData": {
            "type": "object",
            "properties": {
                "code": {
                    "$ref": "#/definitions/controller.ResCode"
                },
                "data": {
                    "description": "omitempty 没有值就不展示"
                },
                "msg": {}
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
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}