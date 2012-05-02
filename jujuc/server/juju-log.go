package server

import (
	"errors"
	"launchpad.net/gnuflag"
	"launchpad.net/juju/go/cmd"
	"launchpad.net/juju/go/log"
	"strings"
)

// JujuLogCommand implements the juju-log command.
type JujuLogCommand struct {
	ctx     *Context
	Message string
	Debug   bool
}

func (c *JujuLogCommand) checkContext() error {
	return nil
}

func (c *JujuLogCommand) Info() *cmd.Info {
	return &cmd.Info{"juju-log", "<message>", "write a message to the juju log", ""}
}

func (c *JujuLogCommand) Init(f *gnuflag.FlagSet, args []string) error {
	f.BoolVar(&c.Debug, "debug", false, "log at debug level")
	if err := f.Parse(true, args); err != nil {
		return err
	}
	args = f.Args()
	if args == nil {
		return errors.New("no message specified")
	}
	c.Message = strings.Join(args, " ")
	return nil
}

func (c *JujuLogCommand) Run(_ *cmd.Context) error {
	s := []string{}
	if c.ctx.LocalUnitName != "" {
		s = append(s, c.ctx.LocalUnitName)
	}
	if c.ctx.RelationName != "" {
		s = append(s, c.ctx.RelationName)
	}
	msg := c.Message
	if len(s) > 0 {
		msg = strings.Join(s, " ") + ": " + msg
	}
	if c.Debug {
		log.Debugf("%s", msg)
	} else {
		log.Printf("%s", msg)
	}
	return nil
}
