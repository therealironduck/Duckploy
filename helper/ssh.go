package helper

import "github.com/melbahja/goph"

type SshClient interface {
	Run(command string) (output []byte, err error)
}

var GetPasswordClient = func(username, password, hostname string) (SshClient, error) {
	callback, _ := goph.DefaultKnownHosts()

	// if err != nil {
	// 	return nil, nil // TODO
	// }

	return goph.NewConn(&goph.Config{
		User:     username,
		Addr:     hostname,
		Port:     2222,
		Auth:     goph.Password(password),
		Timeout:  goph.DefaultTimeout,
		Callback: callback,
	})
}

// MOCKED SSH CLIENT
type FakeSshClient struct {
	Commands []string
}

func (f *FakeSshClient) Run(command string) (output []byte, err error) {
	f.Commands = append(f.Commands, command)
	return []byte("output"), nil
}
