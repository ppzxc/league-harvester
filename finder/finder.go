package finder

import (
	"bytes"
	"context"
	"errors"
	log "github.com/sirupsen/logrus"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Option struct {
	UseWmic bool
	Timeout time.Duration
}

type Finder interface {
	Execute() (*Result, error)
}

type Result struct {
	Pid   int
	Port  int
	Token string
}

type command struct {
	ctx    context.Context
	option Option
}

func (c *command) Execute() (*Result, error) {
	ctx, cancelFunc := context.WithTimeout(c.ctx, c.option.Timeout)
	path, err := exec.LookPath("powershell.exe")
	if err != nil {
		cancelFunc()
		return nil, err
	}
	var cmd *exec.Cmd
	if c.option.UseWmic {
		cmd = exec.CommandContext(ctx, "wmic", "process", "where", "caption='LeagueClientUx.exe'", "get", "commandLine")
	} else {
		cmd = exec.CommandContext(ctx, path, "-NoProfile", "-NonInteractive", "Get-CimInstance", "-Query", c.getQuery(), "|", "Select-Object -ExpandProperty", "CommandLine")
	}

	var stdoutBuffer, stderrBuffer bytes.Buffer
	cmd.Stdout = &stdoutBuffer
	cmd.Stderr = &stderrBuffer

	if err := cmd.Run(); err != nil {
		cancelFunc()
		return nil, err
	}
	stdout := stdoutBuffer.String()
	//stderr := stdoutBuffer.String()

	if len(strings.TrimSpace(stdout)) == 0 {
		cancelFunc()
		return nil, errors.New("not found: League Of Legends Client")
	}

	stdout = strings.ReplaceAll(stdout, "\r\n", " ")
	log.Debug(stdout)
	port, err := c.getPort(stdout)
	if err != nil {
		cancelFunc()
		return nil, err
	}
	authToken, err := c.getAuthToken(stdout)
	if err != nil {
		cancelFunc()
		return nil, err
	}
	pid, err := c.getPid(stdout)
	if err != nil {
		cancelFunc()
		return nil, err
	}
	cancelFunc()
	return &Result{Pid: pid, Port: port, Token: authToken}, nil
}

func (c *command) getAuthToken(stdout string) (string, error) {
	tokenRegex, err := regexp.Compile("--remoting-auth-token=(.+?)\"")
	if err != nil {
		return "", err
	}
	tokenPrefix := tokenRegex.FindString(stdout)
	tokenPostfix := strings.Replace(tokenPrefix, "--remoting-auth-token=", "", -1)
	authToken := strings.Replace(tokenPostfix, "\"", "", -1)
	return authToken, nil
}

func (c *command) getPort(stdout string) (int, error) {
	r, err := regexp.Compile("--app-port=([0-9]+)")
	if err != nil {
		return 0, err
	}
	findAppPort := r.FindString(stdout)
	port, _ := strconv.Atoi(strings.Replace(findAppPort, "--app-port=", "", -1))
	return port, nil
}

func (c *command) getQuery() string {
	return "\"SELECT * FROM Win32_Process WHERE name LIKE 'LeagueClientUx.exe'\""
}

func (c *command) getPid(stdout string) (int, error) {
	r, err := regexp.Compile("--app-pid=([0-9]+)")
	if err != nil {
		return 0, err
	}
	findAppPort := r.FindString(stdout)
	pid, _ := strconv.Atoi(strings.Replace(findAppPort, "--app-pid=", "", -1))
	return pid, nil
}

func New(ctx context.Context, option Option) Finder {
	return &command{
		ctx:    ctx,
		option: option,
	}
}
