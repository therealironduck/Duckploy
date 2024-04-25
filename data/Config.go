package data

type Host struct {
	Hostname    string `json:"hostname"`
	SshUser     string `json:"ssh_user"`
	SshPassword string `json:"ssh_password"`
	Path        string `json:"path"`
}

type Config struct {
	Hosts []Host `json:"hosts"`
}
