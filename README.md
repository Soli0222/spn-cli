# SpotifyNowPlaying (spn-cli)

SpotifyNowPlaying, or `spn-cli`, is a command-line tool that allows you to easily share the song you're currently playing on Spotify to various social media platforms. It automates the process, making it quick and efficient to keep your followers updated on your musical tastes.

## Features

- **Login to Spotify**: Authenticate your Spotify account.
- **Get Currently Playing Song**: Display the song you're currently listening to, along with the artist and track URL.
- **Share to Misskey**: Post your currently playing Spotify track to Misskey.
- **Share to Twitter**: Tweet your currently playing Spotify track.
- **Retrieve Misskey Profile**: Fetch your Misskey user profile information.
- **Store Misskey Credentials**: Save your Misskey hostname and token for future use.
- **Version Information**: Display the version of the CLI tool.

## Installation

You can run the program using the pre-built binary without needing to build the source code.

1. Download the latest binary for your OS from the [releases page](https://github.com/Soli0222/spn-cli/releases).
2. Extract the downloaded binary and make it executable (if necessary).

   ```bash
   chmod +x spn-cli
   ```

3. Move it to a directory in your PATH or use it directly as a command.

## Usage

After installing, you can use the following commands:

```bash
spn-cli [command]
```

### Available Commands

- **`completion`**: Generate autocompletion script for your shell.
- **`getMisskey`**: Retrieve your Misskey user profile information.
- **`login`**: Login to Spotify.
- **`misskey`**: Share the currently playing Spotify track to Misskey.
- **`nowplaying`**: Get the currently playing song with artist name and track URL.
- **`setMisskey`**: Set Misskey hostname and token for authentication.
- **`tweet`**: Share the currently playing Spotify track to Twitter.
- **`version`**: Displays the current version of the SpotifyNowPlaying CLI tool.

## Examples

- **Login to Spotify**:
  ```bash
  spn-cli login
  ```

- **Get the currently playing song**:
  ```bash
  spn-cli nowplaying
  ```

- **Login to Misskey**:
  ```bash
  spn-cli setMisskey -H example.tld -T MISSKEY_API_TOKEN
  ```

- **Get Misskey Account**:
  ```bash
  spn-cli getMisskey
  ```

- **Share the current track to Misskey**:
  ```bash
  spn-cli misskey
  ```

- **Tweet the current track**:
  ```bash
  spn-cli tweet
  ```

## License

This project is licensed under the MIT License.
