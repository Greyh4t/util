package executil

import (
	"bytes"
	"context"
	"io"
	"os/exec"
	"github.com/greyh4t/util/encoding/converter"
)

type Cmd struct {
	cmd    *exec.Cmd
	stdout io.ReadCloser
}

// CommandContext 创建一个新的Cmd
func CommandContext(ctx context.Context, name string, args ...string) *Cmd {
	cmd := exec.CommandContext(ctx, name, args...)
	return &Cmd{cmd: cmd}
}

// Run 执行命令，等待程序退出，返回stdout和stderr
func (c *Cmd) Run() ([]byte, []byte, error) {
	var o, e bytes.Buffer
	c.cmd.Stdout = &o
	c.cmd.Stderr = &e

	err := c.cmd.Run()
	if err != nil {
		return nil, nil, err
	}

	return converter.GBKToUTF8(o.Bytes()), converter.GBKToUTF8(e.Bytes()), nil
}

// Start 执行命令，非阻塞，返回stdout对应的reader
func (c *Cmd) Start() (io.Reader, error) {
	stdout, err := c.cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	c.stdout = stdout

	c.cmd.Stderr = &prefixSuffixSaver{N: 32 << 10}

	err = c.cmd.Start()
	if err != nil {
		stdout.Close()
		return nil, err
	}

	return stdout, nil
}

// Wait 等待程序退出，并返回是否有错误，与Start组合使用
func (c *Cmd) Wait() error {
	err := c.cmd.Wait()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			ee.Stderr = c.cmd.Stderr.(*prefixSuffixSaver).Bytes()
		}
	}

	if c.stdout != nil {
		c.stdout.Close()
	}

	return err
}
