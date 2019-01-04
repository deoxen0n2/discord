# Discord

A command line discord client. Currently support only sending a message through webhook. PRs are welcome.

## Usage

It reads a message through STDIN and sends it through provided webhook URL.

```sh
echo "Test `message` with **markdown** syntax" | discord --webhook-url='https://discordapp.com/api/webhooks/1234/AbCD
```

## Build

```sh
go build -o discord cmd/*.go
```

## LICENSE

MIT
