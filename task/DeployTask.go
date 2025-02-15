// Copyright 2022 The Goploy Authors. All rights reserved.
// Use of this source code is governed by a GPLv3-style
// license that can be found in the LICENSE file.

package task

import (
	"container/list"
	"github.com/zhenorzz/goploy/config"
	"github.com/zhenorzz/goploy/service"
	"github.com/zhenorzz/goploy/ws"
	"sync"
	"sync/atomic"
	"time"
)

var deployList = list.New()
var deployTick = time.Tick(time.Millisecond)

func startDeployTask() {
	atomic.AddInt32(&counter, 1)
	var deployingNumber int32
	var wg sync.WaitGroup
	go func() {
		for {
			select {
			case <-deployTick:
				if atomic.LoadInt32(&deployingNumber) < config.Toml.APP.DeployLimit {
					atomic.AddInt32(&deployingNumber, 1)
					deployElem := deployList.Front()
					if deployElem != nil {
						wg.Add(1)
						go func() {
							deployList.Remove(deployElem).(service.Gsync).Exec()
							atomic.AddInt32(&deployingNumber, -1)
							wg.Done()
						}()
					} else {
						atomic.AddInt32(&deployingNumber, -1)
					}
				}
			case <-stop:
				wg.Wait()
				atomic.AddInt32(&counter, -1)
				return
			}
		}
	}()
}

func AddDeployTask(gSync service.Gsync) {
	ws.GetHub().Data <- &ws.Data{
		Type:    ws.TypeProject,
		Message: ws.ProjectMessage{ProjectID: gSync.Project.ID, ProjectName: gSync.Project.Name, State: ws.TaskWaiting, Message: "Task waiting"},
	}
	deployList.PushBack(gSync)
}
