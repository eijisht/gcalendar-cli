# gcal-cli

A small command line tool for managing Google Calendar events from the terminal.

## Why

The Google Calendar web UI is slow when you need to add many events at once, for example every exam date at the start of a semester. This tool makes creating, listing, and deleting events quick from the command line. It was built as a project to learn Go.

## Setup

1. Create a Google Cloud project and enable the Google Calendar API.
2. Create an OAuth client ID of type "Desktop app".
3. Download the credentials and save the file as `credentials.json` in the project root. See `credentials_example.json` for the expected shape.
4. Build the binary:

   ```
   make build
   ```

5. Authenticate:

   ```
   ./gcal init
   ```

   This starts a local server, prints a URL for you to authorize in the browser, and saves the token to `token.json`.

## Usage

List upcoming events:

```
./gcal read [-n calendar] [-c count] [-d days]
```

Create an event:

```
./gcal create -s "Summary" [-desc "Description"] -start "2006-01-02 15:04" -end "2006-01-02 15:04"
```

Delete an event by ID (IDs are shown by read and create):

```
./gcal remove <eventID>
```

Other commands:

```
./gcal reset    Delete the cached token
./gcal help     Show usage
```

## Development

```
make build   Build the gcal binary
make test    Run the tests
make vet     Run go vet
make check   Run vet and tests
```
