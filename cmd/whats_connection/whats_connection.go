package whats_connection

import (
	"fmt"
	"github.com/Rhymen/go-whatsapp"
	"github.com/jsdaniell/whats-cli/utils/whats_utils"
	"github.com/spf13/cobra"
	"os"
	"time"
)

// WhatsConnect handles other command, customize it!.
var WhatsConnect = &cobra.Command{
	Use:   "connect",
	Short: "Connect or restore connection to whatsapp on CLI",
	Run: func(cmd *cobra.Command, args []string) {
		// run your different command
		//create new WhatsApp connection

		wac, err := whatsapp.NewConn(5 * time.Second)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error creating connection: %v\n", err)

		}

		err = whats_utils.Login(wac, true)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error logging in: %v\n", err)

		}

		<-time.After(3 * time.Second)
	},
}
