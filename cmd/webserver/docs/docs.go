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
        "contact": {
            "name": "API Support",
            "url": "http://servicedesk.tattelecom.ttc",
            "email": "servicedesk@tattelecom.ru"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/": {
            "get": {
                "produces": [
                    "text/html"
                ],
                "summary": "Текущая версия API",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "body"
                        }
                    }
                }
            }
        },
        "/changeclientcard/": {
            "post": {
                "description": "Передача данных по делам из Jeffit во внешнюю систему",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Сохраняет данные по делам в биллинговой системе",
                "parameters": [
                    {
                        "description": "name search by q",
                        "name": "q",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.TDelo"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.TData_resp"
                        }
                    }
                }
            }
        },
        "/home/": {
            "post": {
                "produces": [
                    "application/json"
                ],
                "summary": "Текущая версия API",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.TVersion"
                        }
                    }
                }
            }
        },
        "/test/": {
            "post": {
                "description": "получить строку по идентификатору",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Retrieves user based on given ID",
                "parameters": [
                    {
                        "description": "name search by q",
                        "name": "q",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.TData_req"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/models.TData_resp"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.TAssignee": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.TClaimApplicant": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.TClaimRecipient": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.TClient": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "representative": {
                    "$ref": "#/definitions/models.TRepresentatives"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "models.TData_req": {
            "type": "object",
            "properties": {
                "user": {
                    "type": "string"
                }
            }
        },
        "models.TData_resp": {
            "type": "object",
            "properties": {
                "status": {
                    "type": "integer"
                }
            }
        },
        "models.TDecisions": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "models.TDelo": {
            "type": "object",
            "properties": {
                "assigneEmail": {
                    "type": "string"
                },
                "assigneName": {
                    "type": "string"
                },
                "assignee": {
                    "$ref": "#/definitions/models.TAssignee"
                },
                "claimApplicant": {
                    "$ref": "#/definitions/models.TClaimApplicant"
                },
                "claimRecipient": {
                    "$ref": "#/definitions/models.TClaimRecipient"
                },
                "client": {
                    "$ref": "#/definitions/models.TClient"
                },
                "ctime": {
                    "description": "дата создания дела \u003cДата\u003e",
                    "type": "string"
                },
                "dateOfReceivedClaim": {
                    "description": "дата отправки претензии \u003cДата\u003e,",
                    "type": "string"
                },
                "dateOfSendClaim": {
                    "type": "string"
                },
                "decisions": {
                    "$ref": "#/definitions/models.TDecisions"
                },
                "declarantEmail": {
                    "description": "email заявителя \u003cСтроковое значение\u003e - устаревшее поле, сохранённое для совместимости, его надо игнорировать и использовать объект client,",
                    "type": "string"
                },
                "declarantName": {
                    "description": "имя заявителя \u003cСтроковое значение\u003e - устаревшее поле, сохранённое для совместимости, его надо игнорировать и использовать объект client,",
                    "type": "string"
                },
                "dueDate": {
                    "description": "ближайший норматив \u003cДата и время\u003e,",
                    "type": "string"
                },
                "fields": {
                    "type": "string"
                },
                "id": {
                    "description": "внешний id дела \u003cСтроковое значение\u003e",
                    "type": "string"
                },
                "issueDescription": {
                    "description": "описание дела  \u003cСтроковое значение\u003e",
                    "type": "string"
                },
                "issueSubject": {
                    "description": "название дела \u003cСтроковое значение\u003e",
                    "type": "string"
                },
                "links": {
                    "$ref": "#/definitions/models.TLinks"
                },
                "serviceId": {
                    "description": "внешний id категории \u003cСтроковое значение\u003e",
                    "type": "string"
                },
                "serviceName": {
                    "description": "название категории \u003cСтроковое значение\u003e",
                    "type": "string"
                }
            }
        },
        "models.TLinks": {
            "type": "object",
            "properties": {
                "link": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.TRepresentatives": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.TVersion": {
            "type": "object",
            "properties": {
                "buildTime": {
                    "type": "string"
                },
                "commit": {
                    "type": "string"
                },
                "release": {
                    "type": "string"
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
	Host:        "localhost:5000",
	BasePath:    "/billing/api/v1",
	Schemes:     []string{},
	Title:       "Swagger IRBiS API.",
	Description: "wbfywefqybfeyuf",
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
