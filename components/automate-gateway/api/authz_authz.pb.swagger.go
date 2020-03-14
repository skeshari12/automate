package api

func init() {
	Swagger.Add("authz_authz", `{
  "swagger": "2.0",
  "info": {
    "title": "components/automate-gateway/api/authz/authz.proto",
    "version": "version not set"
  },
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/auth/introspect": {
      "get": {
        "operationId": "IntrospectAll",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.authz.response.IntrospectResp"
            }
          }
        },
        "tags": [
          "Authorization"
        ]
      },
      "post": {
        "operationId": "Introspect",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.authz.response.IntrospectResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.authz.request.IntrospectReq"
            }
          }
        ],
        "tags": [
          "Authorization"
        ]
      }
    },
    "/auth/introspect_some": {
      "post": {
        "operationId": "IntrospectSome",
        "responses": {
          "200": {
            "description": "A successful response.",
            "schema": {
              "$ref": "#/definitions/chef.automate.api.authz.response.IntrospectResp"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/chef.automate.api.authz.request.IntrospectSomeReq"
            }
          }
        ],
        "tags": [
          "Authorization"
        ]
      }
    }
  },
  "definitions": {
    "chef.automate.api.authz.request.IntrospectReq": {
      "type": "object",
      "properties": {
        "path": {
          "type": "string"
        },
        "parameters": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "chef.automate.api.authz.request.IntrospectSomeReq": {
      "type": "object",
      "properties": {
        "paths": {
          "type": "array",
          "items": {
            "type": "string"
          }
        }
      }
    },
    "chef.automate.api.authz.response.IntrospectResp": {
      "type": "object",
      "properties": {
        "endpoints": {
          "type": "object",
          "additionalProperties": {
            "$ref": "#/definitions/chef.automate.api.authz.response.MethodsAllowed"
          }
        }
      }
    },
    "chef.automate.api.authz.response.MethodsAllowed": {
      "type": "object",
      "properties": {
        "get": {
          "type": "boolean",
          "format": "boolean"
        },
        "put": {
          "type": "boolean",
          "format": "boolean"
        },
        "post": {
          "type": "boolean",
          "format": "boolean"
        },
        "delete": {
          "type": "boolean",
          "format": "boolean"
        },
        "patch": {
          "type": "boolean",
          "format": "boolean"
        }
      }
    }
  }
}
`)
}