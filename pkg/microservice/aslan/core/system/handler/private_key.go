/*
Copyright 2021 The KodeRover Authors.

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

package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	commonmodels "github.com/koderover/zadig/pkg/microservice/aslan/core/common/repository/models"
	"github.com/koderover/zadig/pkg/microservice/aslan/core/system/service"
	internalhandler "github.com/koderover/zadig/pkg/shared/handler"
	e "github.com/koderover/zadig/pkg/tool/errors"
	"github.com/koderover/zadig/pkg/tool/log"
	"github.com/koderover/zadig/pkg/types/permission"
)

func ListPrivateKeys(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()

	ctx.Resp, ctx.Err = service.ListPrivateKeys(ctx.Logger)
}

func GetPrivateKey(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()

	ctx.Resp, ctx.Err = service.GetPrivateKey(c.Param("id"), ctx.Logger)
}

func CreatePrivateKey(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()

	args := new(commonmodels.PrivateKey)
	data, err := c.GetRawData()
	if err != nil {
		log.Errorf("CreatePrivateKey c.GetRawData() err : %v", err)
	}
	if err = json.Unmarshal(data, args); err != nil {
		log.Errorf("CreatePrivateKey json.Unmarshal err : %v", err)
	}
	internalhandler.InsertOperationLog(c, ctx.Username, "", "新增", "资源管理-主机管理", fmt.Sprintf("hostName:%s ip:%s", args.Name, args.IP), permission.SuperUserUUID, string(data), ctx.Logger)

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	if err := c.ShouldBindWith(&args, binding.JSON); err != nil {
		ctx.Err = e.ErrInvalidParam.AddDesc("invalid PrivateKey args")
		return
	}
	args.UpdateBy = ctx.Username

	ctx.Err = service.CreatePrivateKey(args, ctx.Logger)
}

func UpdatePrivateKey(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()

	args := new(commonmodels.PrivateKey)
	data, err := c.GetRawData()
	if err != nil {
		log.Errorf("UpdatePrivateKey c.GetRawData() err : %v", err)
	}
	if err = json.Unmarshal(data, args); err != nil {
		log.Errorf("UpdatePrivateKey json.Unmarshal err : %v", err)
	}
	internalhandler.InsertOperationLog(c, ctx.Username, "", "更新", "资源管理-主机管理", fmt.Sprintf("hostName:%s ip:%s", args.Name, args.IP), permission.SuperUserUUID, string(data), ctx.Logger)

	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	if err := c.ShouldBindWith(&args, binding.JSON); err != nil {
		ctx.Err = e.ErrInvalidParam.AddDesc("invalid PrivateKey args")
		return
	}
	args.UpdateBy = ctx.Username

	ctx.Err = service.UpdatePrivateKey(c.Param("id"), args, ctx.Logger)
}

func DeletePrivateKey(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()

	internalhandler.InsertOperationLog(c, ctx.Username, "", "删除", "资源管理-主机管理", fmt.Sprintf("id:%s", c.Param("id")), permission.SuperUserUUID, "", ctx.Logger)
	ctx.Err = service.DeletePrivateKey(c.Param("id"), ctx.Logger)
}
