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
        "/blocks/{networkCode}/{hash}": {
            "get": {
                "description": "Get a block along with the first ten transactions",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Blocks"
                ],
                "summary": "Get a block",
                "operationId": "get-block",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The acronym of the network you're querying required",
                        "name": "networkCode",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The blockhash or height (number) on the network you're querying",
                        "name": "hash",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.BlockResponse"
                        }
                    }
                }
            }
        },
        "/transactions/{networkCode}/{transactionId}": {
            "get": {
                "description": "Get a transaction",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Get a transaction",
                "operationId": "get-transaction",
                "parameters": [
                    {
                        "type": "string",
                        "description": "The acronym of the network you're querying required",
                        "name": "networkCode",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "The transaction hash (id) on the network you're querying",
                        "name": "transactionId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/presenter.TransactionResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "presenter.BlockResponse": {
            "type": "object",
            "properties": {
                "blockNumber": {
                    "description": "The height of the block in the blockchain, or its number",
                    "type": "integer"
                },
                "dateTime": {
                    "description": "The time at which this block was mined by the miner",
                    "type": "string"
                },
                "networkCode": {
                    "description": "The acronym of the network",
                    "type": "string"
                },
                "nextBlockhash": {
                    "description": "The block hash of the next block in the blockchain. NextBlockhash=null if this is the last block in the blockchain",
                    "type": "string"
                },
                "previousBlockhash": {
                    "description": "The block hash of the previous block in the blockchain",
                    "type": "string"
                },
                "size": {
                    "description": "The size of the block in bytes",
                    "type": "integer"
                },
                "transactions": {
                    "description": "The array of ids of all transactions in this block, starting with the newly generated coins (only the first 10)",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/presenter.TransactionResponse"
                    }
                }
            }
        },
        "presenter.TransactionResponse": {
            "type": "object",
            "properties": {
                "dateTime": {
                    "description": "The time at which this transaction received by SoChain, or was mined by the miner",
                    "type": "string"
                },
                "fee": {
                    "description": "The fee paid to the miner",
                    "type": "number"
                },
                "sentValue": {
                    "description": "The total value of all coins sent in this transaction",
                    "type": "number"
                },
                "transactionId": {
                    "description": "The transaction id",
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
	Host:        "",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Blockchain API",
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
