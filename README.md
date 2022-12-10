# go-tify 🎵

A simple CLI tool which let's you control your Spotify player from the terminal, built using Go 💙

<img src="https://i.imgur.com/qQ665ZY.png" width="500">

## Usage / Development

Since this is a pet project, I don't plan to release it, this was just a fun way to learn Go. However, if you want to use it, you can clone the repo and run it locally.

### Prerequisites

- Go
- Spotify Premium Account
- A registered Spotify developer app

### Setting up the Spotify app

- Create a Spotify app in the [developer dashboard](https://developer.spotify.com/dashboard/applications), call it anything you like.
- Edit the app and add `http://localhost:420` as a valid redirect URI
- Copy the **Client ID** and **Client Secret** into your `.env` file.

### Installation

- Clone the [repository](https://github.com/aacevski/go-tify)
- Copy the `.env.example` file to `.env` and fill in the values
- Run `go build` to build the binary
- Run `go install` to install the binary
- Run `go-tify` to start the CLI

### Shortcuts

- `Tab` to switch between the table and search input
- `J` to move down the table
- `K` to move up the table
- `Enter` to select a track
- `g` to go to the top of the table
- `G` to go to the bottom of the table
- `Ctrl + C` to quit the app
