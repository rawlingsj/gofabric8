/*
Copyright (C) 2016 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package sshutil

import (
	"bytes"
	"fmt"
	"io"
	"path/filepath"
	"sync"

	"github.com/docker/machine/libmachine/drivers"
	machinessh "github.com/docker/machine/libmachine/ssh"
	"golang.org/x/crypto/ssh"
)

// SSHSession provides methods for running commands on a host.
type SSHSession interface {
	Close() error
	StdinPipe() (io.WriteCloser, error)
	Run(cmd string) error
	Wait() error
}

// NewSSHClient returns an SSH client object for running commands.
func NewSSHClient(d drivers.Driver) (*ssh.Client, error) {
	h, err := newSSHHost(d)
	if err != nil {
		return nil, err

	}
	auth := &machinessh.Auth{}
	if h.SSHKeyPath != "" {
		auth.Keys = []string{h.SSHKeyPath}
	}
	config, err := machinessh.NewNativeConfig(h.Username, auth)
	if err != nil {
		return nil, err
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", h.IP, h.Port), &config)
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Transfer uses an SSH session to copy a file to the remote machine.
func Transfer(data []byte, remotedir, filename string, perm string, c *ssh.Client) error {
	// Delete the old file first. This makes sure permissions get reset.
	deleteCmd := fmt.Sprintf("sudo rm -f %s", filepath.Join(remotedir, filename))
	mkdirCmd := fmt.Sprintf("sudo mkdir -p %s", remotedir)
	for _, cmd := range []string{deleteCmd, mkdirCmd} {
		if err := RunCommand(c, cmd); err != nil {
			return err
		}
	}

	s, err := c.NewSession()
	if err != nil {
		return err
	}

	w, err := s.StdinPipe()
	if err != nil {
		return err
	}
	// The scpcmd below *should not* return until all data is copied and the
	// StdinPipe is closed. But let's use a WaitGroup to make it expicit.
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer w.Close()
		header := fmt.Sprintf("C%s %d %s\n", perm, len(data), filename)
		fmt.Fprint(w, header)
		reader := bytes.NewReader(data)
		io.Copy(w, reader)
		fmt.Fprint(w, "\x00")
	}()
	scpcmd := fmt.Sprintf("sudo /usr/local/bin/scp -t %s", remotedir)
	if err := s.Run(scpcmd); err != nil {
		return err
	}
	wg.Wait()

	return nil
}

func RunCommand(c *ssh.Client, cmd string) error {
	s, err := c.NewSession()
	defer s.Close()
	if err != nil {
		return err
	}

	return s.Run(cmd)
}

type sshHost struct {
	IP         string
	Port       int
	SSHKeyPath string
	Username   string
}

func newSSHHost(d drivers.Driver) (*sshHost, error) {

	ip, err := d.GetSSHHostname()
	if err != nil {
		return nil, err
	}
	port, err := d.GetSSHPort()
	if err != nil {
		return nil, err
	}
	return &sshHost{
		IP:         ip,
		Port:       port,
		SSHKeyPath: d.GetSSHKeyPath(),
		Username:   d.GetSSHUsername(),
	}, nil
}
