package entity

import "errors"

var ErrInvalidUserEntry = errors.New("invalid user entry")
var ErrCryptoFailed = errors.New("failed to hash password")
var ErrUserNotFound = errors.New("user not found")
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrUserAlreadyExists = errors.New("user already exists")
var ErrChatRoomNotFound = errors.New("chat room not found")
var ErrDbFailed = errors.New("database operation failed")
var ErrInvalidRoomName = errors.New("invalid room name")
var ErrChatRoomAlreadyExists = errors.New("chat room already exists")
var ErrUserNotInChatRoom = errors.New("user not in chat room")
var ErrInvalidUserID = errors.New("invalid user id")
var ErrUnauthorized = errors.New("unauthorized")
