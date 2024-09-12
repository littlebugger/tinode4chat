package entity

import "errors"

var ErrInvalidUserEntry = errors.New("invalid user entry")
var ErrCryptoFailed = errors.New("failed to hash password")
var ErrUserNotFound = errors.New("user not found")
var ErrInvalidCredentials = errors.New("invalid credentials")
var ErrInsufficientPermissions = errors.New("insufficient permissions")
var ErrUserAlreadyExists = errors.New("user already exists")
var ErrMessageTooLong = errors.New("message too long")
var ErrChatRoomNotFound = errors.New("chat room not found")
var ErrChatRoomAlreadyExists = errors.New("chat room already exists")
var ErrFailedToCreateChatRoom = errors.New("failed to create chat room")
var ErrFailedToJoinChatRoom = errors.New("failed to join chat room")
var ErrFailedToLeaveChatRoom = errors.New("failed to leave chat room")
var ErrFailedToGetChatRoomMembers = errors.New("failed to get chat room members")
var ErrFailedToGetChatRoomMessages = errors.New("failed to get chat room messages")
var ErrFailedToSendMessageToChatRoom = errors.New("failed to send message to chat room")
var ErrFailedToSendMessage = errors.New("failed to send message")
var ErrSessionExpired = errors.New("session has expired")
var ErrJWTFailed = errors.New("JWT token missing or invalid")
var ErrDbFailed = errors.New("database operation failed")
