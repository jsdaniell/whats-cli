package whats_sender

import (
	"fmt"
	"github.com/Rhymen/go-whatsapp"
	"github.com/davecgh/go-spew/spew"
	"github.com/jsdaniell/whats-cli/utils/whats_utils"
	"github.com/spf13/cobra"
	"os"
	"time"
)

// OtherCommand handles other command, customize it!.
var WhatsSendMessage = &cobra.Command{
	Use:   "send [number-to-send] [message-text] [sessionID]",
	Short: "Send a whatsapp message (the prefix +55 is fixed)",
	Args:  cobra.MinimumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var sessionID = ""

		if len(args) > 2 {
			sessionID = args[3]
		}

		wac, err := whatsapp.NewConn(20 * time.Second)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error creating connection: %v\n", err)
		}

		err = whats_utils.Login(wac, true, sessionID)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error logging in: %v\n", err)
		}

		<-time.After(3 * time.Second)

		spew.Dump(args)

		msg := whatsapp.TextMessage{
			Info: whatsapp.MessageInfo{
				RemoteJid: args[0] + "@s.whatsapp.net",
			},
			ContextInfo: whatsapp.ContextInfo{},
			Text:        args[1],
		}

		msgId, err := wac.Send(msg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error sending message: %v", err)
			os.Exit(1)
		} else {
			fmt.Println("Message Sent -> ID : " + msgId)
		}
	},
}
