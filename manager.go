package upstart

import (
	"github.com/godbus/dbus"
)

type Manager struct {
	Conn   *dbus.Conn
	Object *dbus.Object
}

func NewManager(bus_type string) (*Manager, error) {
	var err error
	var conn *dbus.Conn

	switch bus_type {
	case "system":
		conn, err = dbus.SystemBus()
	case "session":
		conn, err = dbus.SessionBus()
	}

	if err != nil {
		return nil, err
	}

	obj := conn.Object("com.ubuntu.Upstart", "/com/ubuntu/Upstart")

	return &Manager{
		Conn:   conn,
		Object: obj,
	}, nil

}

func (manager *Manager) GetJobByName(name string) (*Job, error) {

	var job_path dbus.ObjectPath

	err := manager.Object.Call(
		"com.ubuntu.Upstart0_6.GetJobByName",
		0,
		name,
	).Store(&job_path)

	if err != nil {
		return nil, err
	}
	
	job := NewJob(
	return &job_path, nil

}

func (manager *Manager) GetAllJobs() (*[]dbus.ObjectPath, error) {

	var jobs []dbus.ObjectPath

	call := manager.Object.Call(
		"com.ubuntu.Upstart0_6.GetAllJobs",
		0,
	)

	err := call.Store(&jobs)
	if err != nil {
		return nil, err
	}

	return &jobs, nil
}

func (manager *Manager) EmitEvent(name string, env []string, wait bool) {

	manager.Object.Call(
		"com.ubuntu.Upstart0_6.EmitEvent",
		0,
		name,
		env,
		wait,
	)
}
