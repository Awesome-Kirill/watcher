// Code generated by swaggo/swag. DO NOT EDIT.

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
        "/admin/stat": {
            "get": {
                "description": "Return most fasts site",
                "produces": [
                    "application/json"
                ],
                "summary": "Return statistic",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.GetMinResponse"
                        }
                    }
                }
            }
        },
        "/stat/max": {
            "get": {
                "description": "Return most slow site",
                "produces": [
                    "application/json"
                ],
                "summary": "Return most slow site",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.GetMaxResponse"
                        }
                    }
                }
            }
        },
        "/stat/min": {
            "get": {
                "description": "Return most fasts site",
                "produces": [
                    "application/json"
                ],
                "summary": "Return most fasts site",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.GetMinResponse"
                        }
                    }
                }
            }
        },
        "/stat/{id}/site": {
            "get": {
                "description": "Return most fasts site",
                "produces": [
                    "application/json"
                ],
                "summary": "Return most fasts site",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Site ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.GetSiteStatResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.GetMaxResponse": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "seconds": {
                    "type": "number"
                }
            }
        },
        "handler.GetMinResponse": {
            "type": "object",
            "properties": {
                "name": {
                    "type": "string"
                },
                "seconds": {
                    "type": "number"
                }
            }
        },
        "handler.GetSiteStatResponse": {
            "type": "object",
            "properties": {
                "IsAlive": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "seconds": {
                    "type": "number"
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