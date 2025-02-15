// Copyright 2022 The Goploy Authors. All rights reserved.
// Use of this source code is governed by a GPLv3-style
// license that can be found in the LICENSE file.

package middleware

import (
	"encoding/json"
	"errors"
	"github.com/zhenorzz/goploy/core"
	"github.com/zhenorzz/goploy/model"
	"github.com/zhenorzz/goploy/permission"
)

// HasProjectPermission check the user has publish auth
func HasProjectPermission(gp *core.Goploy) error {
	if _, ok := gp.Namespace.PermissionIDs[permission.GetAllDeployList]; ok {
		return nil
	}
	type ReqData struct {
		ProjectID int64 `json:"projectId"`
	}
	var reqData ReqData
	if err := json.Unmarshal(gp.Body, &reqData); err != nil {
		return err
	}

	_, err := model.Project{ID: reqData.ProjectID, UserID: gp.UserInfo.ID}.GetUserProjectData()
	if err != nil {
		return errors.New("no permission")
	}
	return nil
}

// FilterEvent check the webhook event has publish auth
func FilterEvent(gp *core.Goploy) error {
	if XGitHubEvent := gp.Request.Header.Get("X-GitHub-Event"); len(XGitHubEvent) != 0 && XGitHubEvent == "push" {
		return nil
	} else if XGitLabEvent := gp.Request.Header.Get("X-Gitlab-Event"); len(XGitLabEvent) != 0 && XGitLabEvent == "Push Hook" {
		return nil
	} else if XGiteeEvent := gp.Request.Header.Get("X-Gitee-Event"); len(XGiteeEvent) != 0 && XGiteeEvent == "Push Hook" {
		return nil
	} else if XSVNEvent := gp.Request.Header.Get("X-SVN-Event"); len(XSVNEvent) != 0 && XSVNEvent == "push" {
		return nil
	} else {
		return errors.New("only receive push event")
	}
}
