
package swaggerui
//This file is generated automatically. Do not try to edit it manually.

var resourceListingJson = `{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://localhost:8081/swagger-ui",
    "apis": [
        {
            "path": "/accounts",
            "description": "account"
        }
    ],
    "info": {
        "title": "yakiniku blog",
        "description": "yakiniku blog sample"
    }
}`
var apiDescriptionsJson = map[string]string{"accounts":`{
    "apiVersion": "1.0.0",
    "swaggerVersion": "1.2",
    "basePath": "http://localhost:8081",
    "resourcePath": "/accounts",
    "apis": [
        {
            "path": "/accounts",
            "description": "アカウント一覧取得API",
            "operations": [
                {
                    "httpMethod": "GET",
                    "nickname": "アカウント一覧取得",
                    "type": "array",
                    "items": {
                        "$ref": "github.com.y-ogura.yakiniku.account.Client"
                    },
                    "summary": "アカウント一覧取得API",
                    "responseMessages": [
                        {
                            "code": 200,
                            "message": "アカウント",
                            "responseType": "array",
                            "responseModel": "github.com.y-ogura.yakiniku.account.Client"
                        },
                        {
                            "code": 500,
                            "message": "internal server error",
                            "responseType": "object",
                            "responseModel": "error"
                        }
                    ]
                }
            ]
        }
    ],
    "models": {
        "github.com.satori.go.uuid.UUID": {
            "id": "github.com.satori.go.uuid.UUID",
            "properties": null
        },
        "github.com.y-ogura.yakiniku.account.Client": {
            "id": "github.com.y-ogura.yakiniku.account.Client",
            "properties": {
                "email": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "id": {
                    "type": "github.com.satori.go.uuid.UUID",
                    "description": "",
                    "items": {},
                    "format": ""
                },
                "name": {
                    "type": "string",
                    "description": "",
                    "items": {},
                    "format": ""
                }
            }
        }
    }
}`,}
