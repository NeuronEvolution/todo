// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Todo Private API",
    "contact": {
      "name": "mars"
    },
    "version": "v1"
  },
  "basePath": "/api-private/v1/todo",
  "paths": {
    "/": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "GetTodoList",
        "parameters": [
          {
            "type": "string",
            "name": "friendID",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/TodoItem"
              }
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "AddTodo",
        "parameters": [
          {
            "name": "todoItem",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/TodoItem"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/categoryNames": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "GetCategoryNameList",
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "type": "array",
              "items": {
                "type": "string"
              }
            }
          }
        }
      }
    },
    "/friends": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "GetFriendsList",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "pageSize",
            "in": "query"
          },
          {
            "type": "string",
            "name": "pageToken",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/FriendInfoList"
            }
          }
        }
      }
    },
    "/friends/{friendID}": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "GetFriend",
        "parameters": [
          {
            "type": "string",
            "name": "friendID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/FriendInfo"
            }
          }
        }
      }
    },
    "/listByCategory": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "GetTodoListByCategory",
        "parameters": [
          {
            "type": "string",
            "name": "friendID",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/TodoItemGroup"
              }
            }
          }
        }
      }
    },
    "/userProfile": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "GetUserProfile",
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/UserProfile"
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "UpdateUserProfile",
        "parameters": [
          {
            "name": "userProfile",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/UserProfile"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok"
          }
        }
      }
    },
    "/userProfile/todoVisibility": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "UpdateUserProfileTodoVisibility",
        "parameters": [
          {
            "name": "visibility",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/TodoVisibility"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok"
          }
        }
      }
    },
    "/userProfile/userName": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "UpdateUserProfileUserName",
        "parameters": [
          {
            "type": "string",
            "name": "userName",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "ok"
          }
        }
      }
    },
    "/{todoId}": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "GetTodo",
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/TodoItem"
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "UpdateTodo",
        "parameters": [
          {
            "name": "todoItem",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/TodoItem"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok"
          }
        }
      },
      "delete": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "RemoveTodo",
        "responses": {
          "200": {
            "description": "ok"
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "todoId",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "FriendInfo": {
      "type": "object",
      "properties": {
        "todoCount": {
          "type": "integer",
          "format": "int64"
        },
        "todoVisibility": {
          "$ref": "#/definitions/TodoVisibility"
        },
        "userID": {
          "type": "string"
        },
        "userName": {
          "type": "string"
        }
      }
    },
    "FriendInfoList": {
      "type": "object",
      "properties": {
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/FriendInfo"
          }
        },
        "nextPageToken": {
          "type": "string"
        }
      }
    },
    "TodoItem": {
      "type": "object",
      "properties": {
        "category": {
          "type": "string"
        },
        "desc": {
          "type": "string"
        },
        "priority": {
          "type": "integer",
          "format": "int32"
        },
        "status": {
          "$ref": "#/definitions/TodoStatus"
        },
        "title": {
          "type": "string"
        },
        "todoId": {
          "type": "string"
        },
        "userID": {
          "type": "string"
        }
      }
    },
    "TodoItemGroup": {
      "type": "object",
      "properties": {
        "category": {
          "type": "string"
        },
        "todoItemList": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/TodoItem"
          }
        }
      }
    },
    "TodoStatus": {
      "type": "string",
      "enum": [
        "ongoing",
        "completed",
        "discard"
      ]
    },
    "TodoVisibility": {
      "type": "string",
      "enum": [
        "private",
        "public",
        "friend"
      ]
    },
    "UserProfile": {
      "type": "object",
      "properties": {
        "todoVisibility": {
          "$ref": "#/definitions/TodoVisibility"
        },
        "userID": {
          "type": "string"
        },
        "userName": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "Basic": {
      "type": "basic"
    },
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "title": "Todo Private API",
    "contact": {
      "name": "mars"
    },
    "version": "v1"
  },
  "basePath": "/api-private/v1/todo",
  "paths": {
    "/": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "GetTodoList",
        "parameters": [
          {
            "type": "string",
            "name": "friendID",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/TodoItem"
              }
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "AddTodo",
        "parameters": [
          {
            "name": "todoItem",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/TodoItem"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "type": "string"
            }
          }
        }
      }
    },
    "/categoryNames": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "GetCategoryNameList",
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "type": "array",
              "items": {
                "type": "string"
              }
            }
          }
        }
      }
    },
    "/friends": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "GetFriendsList",
        "parameters": [
          {
            "type": "integer",
            "format": "int64",
            "name": "pageSize",
            "in": "query"
          },
          {
            "type": "string",
            "name": "pageToken",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/FriendInfoList"
            }
          }
        }
      }
    },
    "/friends/{friendID}": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "GetFriend",
        "parameters": [
          {
            "type": "string",
            "name": "friendID",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/FriendInfo"
            }
          }
        }
      }
    },
    "/listByCategory": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "GetTodoListByCategory",
        "parameters": [
          {
            "type": "string",
            "name": "friendID",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/TodoItemGroup"
              }
            }
          }
        }
      }
    },
    "/userProfile": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "GetUserProfile",
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/UserProfile"
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "UpdateUserProfile",
        "parameters": [
          {
            "name": "userProfile",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/UserProfile"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok"
          }
        }
      }
    },
    "/userProfile/todoVisibility": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "UpdateUserProfileTodoVisibility",
        "parameters": [
          {
            "name": "visibility",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/TodoVisibility"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok"
          }
        }
      }
    },
    "/userProfile/userName": {
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "UpdateUserProfileUserName",
        "parameters": [
          {
            "type": "string",
            "name": "userName",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "ok"
          }
        }
      }
    },
    "/{todoId}": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "GetTodo",
        "responses": {
          "200": {
            "description": "ok",
            "schema": {
              "$ref": "#/definitions/TodoItem"
            }
          }
        }
      },
      "post": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "UpdateTodo",
        "parameters": [
          {
            "name": "todoItem",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/TodoItem"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "ok"
          }
        }
      },
      "delete": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "operationId": "RemoveTodo",
        "responses": {
          "200": {
            "description": "ok"
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "todoId",
          "in": "path",
          "required": true
        }
      ]
    }
  },
  "definitions": {
    "FriendInfo": {
      "type": "object",
      "properties": {
        "todoCount": {
          "type": "integer",
          "format": "int64"
        },
        "todoVisibility": {
          "$ref": "#/definitions/TodoVisibility"
        },
        "userID": {
          "type": "string"
        },
        "userName": {
          "type": "string"
        }
      }
    },
    "FriendInfoList": {
      "type": "object",
      "properties": {
        "items": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/FriendInfo"
          }
        },
        "nextPageToken": {
          "type": "string"
        }
      }
    },
    "TodoItem": {
      "type": "object",
      "properties": {
        "category": {
          "type": "string"
        },
        "desc": {
          "type": "string"
        },
        "priority": {
          "type": "integer",
          "format": "int32"
        },
        "status": {
          "$ref": "#/definitions/TodoStatus"
        },
        "title": {
          "type": "string"
        },
        "todoId": {
          "type": "string"
        },
        "userID": {
          "type": "string"
        }
      }
    },
    "TodoItemGroup": {
      "type": "object",
      "properties": {
        "category": {
          "type": "string"
        },
        "todoItemList": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/TodoItem"
          }
        }
      }
    },
    "TodoStatus": {
      "type": "string",
      "enum": [
        "ongoing",
        "completed",
        "discard"
      ]
    },
    "TodoVisibility": {
      "type": "string",
      "enum": [
        "private",
        "public",
        "friend"
      ]
    },
    "UserProfile": {
      "type": "object",
      "properties": {
        "todoVisibility": {
          "$ref": "#/definitions/TodoVisibility"
        },
        "userID": {
          "type": "string"
        },
        "userName": {
          "type": "string"
        }
      }
    }
  },
  "securityDefinitions": {
    "Basic": {
      "type": "basic"
    },
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
}
