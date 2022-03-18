# whatsapp-cli

Connect and send messages as a whatsapp client directly from command line, **this is only for experimental use**.

## Installing

##### Install with NPM

`npm i -g whats-cli`

##### Install with SourceCode

`git clone https://github.com/jsdaniell/whats-cli.git`

`cd whats-cli`

`go install`

`go build`

`./whats-cli connect-qr`

##### Install with release

Download the release file of your system type on the [Releases Page](https://github.com/jsdaniell/whats-cli/releases).

`cd {downloadedRelease}`

`./whats-cli connect-qr`

## Testing

##### Get the QRCode on command line

`whats-cli connect-qr "sessionID"`

##### Get the string of QRCode on command line

`whats-cli connect "sessionID"`

##### Send message

`whats-cli send "number" "message" "sessionID"`

#### Dealing with bad restored sessions:

Sometimes if you want to reset the stored session you can erase the `whatsappSession.gob` file, what's generated locally when logging a session.

`rm whatsappSession-{sessionID}.gob`



