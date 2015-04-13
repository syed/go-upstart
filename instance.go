package upstart

import (
	"fmt"
	"github.com/godbus/dbus"
)

type Instance struct {
	Object *dbus.Object
	Path   dbus.ObjectPath
}

func NewInstance(instance_path dbus.ObjectPath, conn *dbus.Conn) (*Instance, error) {

	obj := conn.Object(
		"com.ubuntu.Upstart",
		instance_path,
	)

	return &Job{
		Conn:   conn,
		Object: obj,
		Path:   instance_path,
	}, nil
}

func (instance *Instance) Start() {
	var env []string
	var wait bool

	wait = true

	instance.Object.Call(
		"com.ubuntu.Upstart0_6.Instance.Start",
		0,
		env,
		wait,
	)

}

func (instance *Instance) Stop() {

	var env []string
	var wait bool

	wait = true

	instance.Object.Call(
		"com.ubuntu.Upstart0_6.Instance.Stop",
		0,
		env,
		wait,
	)

}

func (instance *Instance) Restart() {
	var env []string
	var wait bool

	wait = true

	job.Object.Call(
		"com.ubuntu.Upstart0_6.Instance.Restart",
		0,
		env,
		wait,
	)
}
