package helper

const SimpleJsonConfig = `{
	"steps": [
		{
			"action": "command",
			"command": "npm install"
		},
		{
			"action": "command",
			"command": "composer install"
		}
	],
	"hosts": [
		{
			"hostname": "some-host",
			"ssh_user": "ducky",
			"ssh_password": "secret123",
			"path": "/some/path"
		}
	]
}`
