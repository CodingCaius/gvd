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
        "/ap1/users_info": {
            "get": {
                "description": "用户信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/user_api.UserInfoResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/image": {
            "post": {
                "description": "上传图片",
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片管理"
                ],
                "summary": "上传图片",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "file",
                        "description": "文件上传",
                        "name": "image",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/api/images": {
            "get": {
                "description": "图片列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片管理"
                ],
                "summary": "图片列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "模糊匹配的关键字",
                        "name": "key",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/res.ListResponse-image_api_ImageListResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "delete": {
                "description": "删除图片",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "图片管理"
                ],
                "summary": "删除图片",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/api/login": {
            "post": {
                "description": "用户登录",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user_api.UserLoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/api/logs": {
            "get": {
                "description": "日志列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "日志管理"
                ],
                "summary": "日志列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "模糊匹配的关键字",
                        "name": "key",
                        "in": "query"
                    },
                    {
                        "enum": [
                            1,
                            2,
                            3
                        ],
                        "type": "integer",
                        "x-enum-varnames": [
                            "Info",
                            "Warning",
                            "Error"
                        ],
                        "description": "日志查询的等级",
                        "name": "level",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "enum": [
                            1,
                            2,
                            3
                        ],
                        "type": "integer",
                        "x-enum-varnames": [
                            "LoginType",
                            "ActionType",
                            "RuntimeType"
                        ],
                        "description": "日志的类型   1 登录日志  2 操作日志  3 运行日志",
                        "name": "type",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/res.ListResponse-log_stash_LogModel"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            }
        },
        "/api/roles": {
            "get": {
                "description": "角色列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色管理"
                ],
                "summary": "角色列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "模糊匹配的关键字",
                        "name": "key",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/res.ListResponse-role_api_RoleListResponse"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "description": "更新角色",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色管理"
                ],
                "summary": "更新角色",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/role_api.RoleCreateRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "创建角色",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "角色管理"
                ],
                "summary": "创建角色",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/role_api.RoleCreateRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/api/site": {
            "get": {
                "description": "站点配置查询",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "站点配置"
                ],
                "summary": "站点配置查询",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/config.Site"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "description": "站点配置更新",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "站点配置"
                ],
                "summary": "站点配置更新",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/api/users": {
            "get": {
                "description": "用户列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户列表",
                "parameters": [
                    {
                        "type": "string",
                        "description": "模糊匹配的关键字",
                        "name": "key",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "allOf": [
                                {
                                    "$ref": "#/definitions/res.Response"
                                },
                                {
                                    "type": "object",
                                    "properties": {
                                        "data": {
                                            "$ref": "#/definitions/res.ListResponse-models_UserModel"
                                        }
                                    }
                                }
                            ]
                        }
                    }
                }
            },
            "put": {
                "description": "管理员更新用户信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "管理员更新用户信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user_api.UserUpdateRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            },
            "post": {
                "description": "创建用户，只能管理员创建",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "创建用户",
                "parameters": [
                    {
                        "description": "参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user_api.UserCreateRequest"
                        }
                    },
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            },
            "delete": {
                "description": "删除用户",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "删除用户",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/api/users_info": {
            "put": {
                "description": "用户更新自己的信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户更新自己的信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user_api.UserUpdateInfoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        },
        "/api/users_password": {
            "put": {
                "description": "用户修改密码",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "用户管理"
                ],
                "summary": "用户修改密码",
                "parameters": [
                    {
                        "type": "string",
                        "description": "token",
                        "name": "token",
                        "in": "header",
                        "required": true
                    },
                    {
                        "description": "参数",
                        "name": "data",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/user_api.UserUpdatePasswordRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/res.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "config.Site": {
            "type": "object",
            "properties": {
                "abstract": {
                    "description": "网站简介",
                    "type": "string"
                },
                "content": {
                    "description": "内容",
                    "type": "string"
                },
                "footer": {
                    "description": "尾部信息",
                    "type": "string"
                },
                "href": {
                    "description": "点击go的跳转链接",
                    "type": "string"
                },
                "icon": {
                    "description": "首页的图标",
                    "type": "string"
                },
                "iconHref": {
                    "description": "图标链接",
                    "type": "string"
                },
                "title": {
                    "description": "网站名称",
                    "type": "string"
                }
            }
        },
        "image_api.ImageListResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "fileName": {
                    "type": "string"
                },
                "hash": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "nickName": {
                    "type": "string"
                },
                "path": {
                    "description": "update/xx.png",
                    "type": "string"
                },
                "size": {
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userID": {
                    "type": "integer"
                },
                "webPath": {
                    "type": "string"
                }
            }
        },
        "log_stash.Level": {
            "type": "integer",
            "enum": [
                1,
                2,
                3
            ],
            "x-enum-varnames": [
                "Info",
                "Warning",
                "Error"
            ]
        },
        "log_stash.LogModel": {
            "type": "object",
            "properties": {
                "addr": {
                    "description": "地址",
                    "type": "string"
                },
                "content": {
                    "description": "详情",
                    "type": "string"
                },
                "createdAt": {
                    "description": "添加时间",
                    "type": "string"
                },
                "id": {
                    "description": "主键id",
                    "type": "integer"
                },
                "ip": {
                    "description": "ip",
                    "type": "string"
                },
                "level": {
                    "description": "等级",
                    "allOf": [
                        {
                            "$ref": "#/definitions/log_stash.Level"
                        }
                    ]
                },
                "serviceName": {
                    "description": "服务名称",
                    "type": "string"
                },
                "status": {
                    "description": "登录状态",
                    "type": "boolean"
                },
                "title": {
                    "description": "标题",
                    "type": "string"
                },
                "type": {
                    "description": "日志的类型  1 登录 2 操作 3 运行",
                    "allOf": [
                        {
                            "$ref": "#/definitions/log_stash.LogType"
                        }
                    ]
                },
                "updatedAt": {
                    "description": "更新时间",
                    "type": "string"
                },
                "userID": {
                    "description": "用户id",
                    "type": "integer"
                },
                "userName": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "log_stash.LogType": {
            "type": "integer",
            "enum": [
                1,
                2,
                3
            ],
            "x-enum-varnames": [
                "LoginType",
                "ActionType",
                "RuntimeType"
            ]
        },
        "models.UserModel": {
            "type": "object",
            "properties": {
                "addr": {
                    "description": "地址",
                    "type": "string"
                },
                "avatar": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "ip": {
                    "type": "string"
                },
                "lastLogin": {
                    "description": "指明两个表的外键关系\nGORM 将会在查询 UserModel 表时自动进行联接（JOIN），并根据 RoleID 字段的值来获取关联的角色信息",
                    "type": "string"
                },
                "nickName": {
                    "type": "string"
                },
                "roleID": {
                    "description": "用户对应的角色",
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                }
            }
        },
        "res.Code": {
            "type": "integer",
            "enum": [
                0,
                7,
                9
            ],
            "x-enum-comments": {
                "ErrorCode": "系统错误",
                "ValidCode": "校验错误"
            },
            "x-enum-varnames": [
                "SUCCESS",
                "ErrorCode",
                "ValidCode"
            ]
        },
        "res.ListResponse-image_api_ImageListResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/image_api.ImageListResponse"
                    }
                }
            }
        },
        "res.ListResponse-log_stash_LogModel": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/log_stash.LogModel"
                    }
                }
            }
        },
        "res.ListResponse-models_UserModel": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/models.UserModel"
                    }
                }
            }
        },
        "res.ListResponse-role_api_RoleListResponse": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "list": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/role_api.RoleListResponse"
                    }
                }
            }
        },
        "res.Response": {
            "type": "object",
            "properties": {
                "code": {
                    "$ref": "#/definitions/res.Code"
                },
                "data": {},
                "msg": {
                    "type": "string"
                }
            }
        },
        "role_api.RoleCreateRequest": {
            "type": "object",
            "required": [
                "title"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "pwd": {
                    "type": "string"
                },
                "title": {
                    "type": "string",
                    "maxLength": 16,
                    "minLength": 2
                }
            }
        },
        "role_api.RoleListResponse": {
            "type": "object",
            "properties": {
                "createdAt": {
                    "type": "string"
                },
                "docCount": {
                    "description": "角色拥有的文档数",
                    "type": "integer"
                },
                "id": {
                    "type": "integer"
                },
                "isSystem": {
                    "description": "是否是系统角色",
                    "type": "boolean"
                },
                "pwd": {
                    "description": "角色密码",
                    "type": "string"
                },
                "title": {
                    "description": "角色的名称",
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userCount": {
                    "description": "角色下的用户数",
                    "type": "integer"
                }
            }
        },
        "user_api.UserCreateRequest": {
            "type": "object",
            "required": [
                "password",
                "roleID",
                "userName"
            ],
            "properties": {
                "nickName": {
                    "description": "昵称",
                    "type": "string"
                },
                "password": {
                    "description": "密码",
                    "type": "string"
                },
                "roleID": {
                    "description": "角色id",
                    "type": "integer"
                },
                "userName": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "user_api.UserInfoResponse": {
            "type": "object",
            "properties": {
                "addr": {
                    "description": "地址",
                    "type": "string"
                },
                "avatar": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "ip": {
                    "type": "string"
                },
                "lastLogin": {
                    "description": "指明两个表的外键关系\nGORM 将会在查询 UserModel 表时自动进行联接（JOIN），并根据 RoleID 字段的值来获取关联的角色信息",
                    "type": "string"
                },
                "nickName": {
                    "type": "string"
                },
                "role": {
                    "type": "string"
                },
                "roleID": {
                    "description": "用户对应的角色",
                    "type": "integer"
                },
                "updatedAt": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "user_api.UserLoginRequest": {
            "type": "object",
            "required": [
                "password",
                "userName"
            ],
            "properties": {
                "password": {
                    "type": "string"
                },
                "userName": {
                    "type": "string"
                }
            }
        },
        "user_api.UserUpdateInfoRequest": {
            "type": "object",
            "properties": {
                "avatar": {
                    "type": "string"
                },
                "nickName": {
                    "type": "string"
                }
            }
        },
        "user_api.UserUpdatePasswordRequest": {
            "type": "object",
            "required": [
                "oldPwd",
                "password"
            ],
            "properties": {
                "oldPwd": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "user_api.UserUpdateRequest": {
            "type": "object",
            "required": [
                "id",
                "password",
                "roleID"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "nickName": {
                    "type": "string"
                },
                "password": {
                    "description": "bingding字段 定义字段在数据绑定（比如 HTTP 请求参数绑定）时的行为",
                    "type": "string"
                },
                "roleID": {
                    "description": "角色ID",
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "101.43.78.114:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "文档项目api文档",
	Description:      "API文档",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
