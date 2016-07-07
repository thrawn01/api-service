// Copyright 2016 Derrick J. Wippler. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package model

import (
	"github.com/howler-chat/api-service/validate"
	"github.com/howler-chat/api-service/validate/field"
	"golang.org/x/net/context"
)

// A Message represents a point in time message generated by the client and attached to a channel.
type Message struct {
	Id        string `json:"id"`
	ChannelId string `json:"channelId"`
	Text      string `json:"text"`
}

// After marshaling from JSON, call this method to validate the object is intact
func (self *Message) Validate(ctx context.Context) error {
	if err := validate.IsValidId(self.ChannelId); err != nil {
		validate.Fail(ctx, err.Error(), field.NewPath("channelId"))
	}
	if err := validate.IsMessageText(self.Text); err != nil {
		validate.Fail(ctx, err.Error(), field.NewPath("channelId"))
	}
	return nil
}

// Modify the model before create
/*func (self *Message) PreCreate(msg *Message) error {
	return nil
}

// Modify the model before update
func (self *Message) PreUpdate(msg *Message) error {
	return nil
}

// Scrub the model of sensitive data before serializing to JSON
func (self *Message) Sanitize(msg *Message) error {
	return nil
}*/

// A MessageGetRequest represents a request by the client to retrieve a specific message
type GetMessageRequest struct {
	MessageId string `json:"messageId"`
	ChannelId string `json:"channelId"`
}

// After marshaling from JSON, call this method to validate the object is intact
func (self *GetMessageRequest) Validate(ctx context.Context) error {
	if err := validate.IsValidId(self.MessageId); err != nil {
		validate.Fail(ctx, err.Error(), field.NewPath("messageId"))
	}

	if err := validate.IsValidId(self.ChannelId); err != nil {
		validate.Fail(ctx, err.Error(), field.NewPath("channelId"))
	}
	return nil
}

// A MessageListRequest represents a request by the client to retrieve a list of messages usually associated with a
// channel
type ListMessageRequest struct {
	ChannelId string `json:"channelId"`
}

// After marshaling from JSON, call this method to validate the object is intact
func (self *ListMessageRequest) Validate(ctx context.Context) error {
	if err := validate.IsValidId(self.ChannelId); err != nil {
		validate.Fail(ctx, err.Error(), field.NewPath("channelId"))
	}
	return nil
}
