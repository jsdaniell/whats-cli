package whats_messages

import (
	"errors"
	"fmt"
	"github.com/Rhymen/go-whatsapp"
	"github.com/jsdaniell/whats-cli/utils/whats_utils"
	"github.com/spf13/cobra"
	"time"
)

// OtherCommand handles other command, customize it!.
var WhatsSendMessage = &cobra.Command{
	Use:   "send",
	Short: "send [number-to-send] [message-text] [sessionID]",
	Args:  cobra.MinimumNArgs(3),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) < 3 {
			return errors.New("you have to inform thr arguments: send [number-to-send] [message-text] [sessionID]")
		}

		var sessionID = ""

		if len(args) > 2 {
			sessionID = args[2]
		}

		wac, err := whatsapp.NewConn(20 * time.Second)
		if err != nil {
			return errors.New("error creating connection: %v\n" + err.Error())
		}

		err = whats_utils.Login(wac, sessionID)
		if err != nil {
			return errors.New("error logging in: %v\n" + err.Error())
		}

		msg := whatsapp.TextMessage{
			Info: whatsapp.MessageInfo{
				RemoteJid: args[0] + "@s.whatsapp.net",
			},
			ContextInfo: whatsapp.ContextInfo{},
			Text:        args[1],
		}

		msgId, err := wac.Send(msg)
		if err != nil {
			return errors.New("error sending message: %v" + err.Error())
		} else {
			fmt.Println("Message Sent -> ID : " + msgId)
		}

		return nil
	},
}
