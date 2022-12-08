package spotify

func Get_Scopes() []string {
	return []string{
		"user-read-playback-state",
		"user-modify-playback-state",
		"user-read-currently-playing",
		"streaming",
		"app-remote-control",
		"user-read-email",
		"user-read-private",
		"playlist-read-private",
		"playlist-read-collaborative",
		"playlist-modify-public",
		"playlist-modify-private",
		"user-library-read",
		"user-library-modify",
		"user-top-read",
		"user-read-playback-position",
		"user-read-recently-played",
		"user-follow-read",
		"user-follow-modify",
	}
}
