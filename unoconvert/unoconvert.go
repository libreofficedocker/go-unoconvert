package unoconvert

import (
	"context"
	"fmt"
	"log"
	"os/exec"
	"time"
)

var (
	DefaultContextTimeout = 0 * time.Minute
)

var (
	ContextTimeout = DefaultContextTimeout
)

func New() *Unoconvert {
	return &Unoconvert{}
}

func Default() *Unoconvert {
	return &Unoconvert{
		Host:       "127.0.0.1",
		Port:       "2002",
		Executable: "unoconvert",
	}
}

type Unoconvert struct {
	Host       string
	Port       string
	Executable string
}

func (u *Unoconvert) SetExecutable(executable string) {
	u.Executable = executable
}

func (u *Unoconvert) SetHost(interf string) {
	u.Host = interf
}

func (u *Unoconvert) SetPort(port string) {
	u.Port = port
}

func (u *Unoconvert) SetContextTimeout(timeout time.Duration) {
	ContextTimeout = timeout
}

func (u *Unoconvert) Create(infile string, outfile string, opts ...UnoconvertOption) *exec.Cmd {
	var args = []string{}

	connections := []string{
		fmt.Sprintf("--interface=%s", u.Host),
		fmt.Sprintf("--port=%s", u.Port),
	}

	files := []string{infile, outfile}

	args = append(connections, files...)

	for _, opt := range opts {
		args = append(args, fmt.Sprintf("%s=%s", opt.Key, opt.Value))
	}

	log.Printf("Command: %s %s", u.Executable, args)
	cmd := exec.Command(u.Executable, args...)
	return cmd
}

func (u *Unoconvert) CreateContext(ctx context.Context, infile string, outfile string, opts ...UnoconvertOption) *exec.Cmd {
	ctx, cancel := context.WithTimeout(ctx, ContextTimeout)
	defer cancel()

	var args = []string{}

	connections := []string{
		fmt.Sprintf("--interface=%s", u.Host),
		fmt.Sprintf("--port=%s", u.Port),
	}

	files := []string{infile, outfile}

	args = append(connections, files...)

	for _, opt := range opts {
		args = append(args, fmt.Sprintf("%s=%s", opt.Key, opt.Value))
	}

	log.Printf("Command: %s %s", u.Executable, args)
	cmd := exec.CommandContext(ctx, u.Executable, args...)
	return cmd
}
