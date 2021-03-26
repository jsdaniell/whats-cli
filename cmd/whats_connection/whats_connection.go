package whats_connection

import (
	"errors"
	"fmt"
	"github.com/Rhymen/go-whatsapp"
	"github.com/jsdaniell/whats-cli/utils/json_utils"
	"github.com/jsdaniell/whats-cli/utils/whats_utils"
	"github.com/spf13/cobra"
	"time"
)

// WhatsConnect handles other command, customize it!.
var WhatsConnect = &cobra.Command{
	Use:   "connect",
	Short: "Connect or restore connection to whatsapp on CLI",
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var sessionID = ""

		if len(args) > 0 {
			sessionID = args[0]
		}

		wac, err := whatsapp.NewConn(20 * time.Second)
		if err != nil {
			return errors.New("error creating connection: %v\n" + err.Error())

		}

		err = whats_utils.NewLogin(wac, true, sessionID)
		if err != nil {
			return errors.New("error logging in: %v\n" + err.Error())
		}

		return nil
	},
}

// WhatsConnectQR handles other command, customize it!.
var WhatsConnectQR = &cobra.Command{
	Use:   "connect-qr",
	Args: cobra.MaximumNArgs(1),
	Short: "Connect or restore connection to whatsapp on CLI showing the QRCode",
	RunE: func(cmd *cobra.Command, args []string) error {
		var sessionID = ""

		if len(args) > 0 {
			sessionID = args[0]
		}

		wac, err := whatsapp.NewConn(20 * time.Second)
		if err != nil {
			return errors.New("error creating connection: %v\n" + err.Error())
		}

		err = whats_utils.NewLogin(wac, false, sessionID)
		if err != nil {
			return errors.New("error logging in: %v\n" + err.Error())
		}

		return nil
	},
}

// WhatsConnectQR handles other command, customize it!.
var WhatsVersion = &cobra.Command{
	Use:   "version",
	Short: "Connect or restore connection to whatsapp on CLI showing the QRCode",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(json_utils.GettingJSONVersion())
	},
}


// WhatsDisconnect handles other command, customize it!.
var WhatsDisconnect = &cobra.Command{
	Use:   "disconnect",
	Short: "Connect or restore connection to whatsapp on CLI showing the QRCode",
	Args: cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var sessionID = ""

		if len(args) > 0 {
			sessionID = args[0]
		}

		err := whats_utils.Disconnect(sessionID)
		if err != nil {
			return errors.New("error logging in: %v\n" + err.Error())

		}

		return nil
	},
}
