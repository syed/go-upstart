package upstart

import (
	"fmt"
	"github.com/godbus/dbus"
)

type Job struct {
	Conn   *dbus.Conn
	Object *dbus.Object
	Path   dbus.ObjectPath
}

func NewJob(job_path dbus.ObjectPath, conn *dbus.Conn) (*Job, error) {

	obj := conn.Object(
		"com.ubuntu.Upstart",
		job_path,
	)

	return &Job{
		Conn:   conn,
		Object: obj,
		Path:   job_path,
	}, nil
}

func (job *Job) GetAllInstances() (*[]Instance, error) {

	var ipaths []dbus.ObjectPath

	err := job.Object.Call(
		"com.ubuntu.Upstart0_6.Job.GetAllInstances",
		0,
	).Store(&ipaths)

	if err != nil {
		return nil, err
	}

	out := make([]Instance, len(ipaths))

	for i, path := range out {
		out[i] = NewInstance(path, job.Conn)
	}

	return out, nil

}

func (job *Job) Start() {
	var env []string
	var wait bool

	wait = true

	job.Object.Call(
		"com.ubuntu.Upstart0_6.Job.Start",
		0,
		env,
		wait,
	)

}

func (job *Job) Stop() {

	var env []string
	var wait bool

	wait = true

	job.Object.Call(
		"com.ubuntu.Upstart0_6.Job.Stop",
		0,
		env,
		wait,
	)

}

func (job *Job) Restart() {
	var env []string
	var wait bool

	wait = true

	job.Object.Call(
		"com.ubuntu.Upstart0_6.Job.Restart",
		0,
		env,
		wait,
	)
}
