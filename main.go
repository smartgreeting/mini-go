/*
 * @Author: lihuan
 * @Date: 2024-08-26 20:15:09
 * @LastEditors: lihuan
 * @LastEditTime: 2024-09-20 20:40:56
 * @Email: 17719495105@163.com
 */
package main

import (
	"fmt"

	"github.com/smartgreeting/mini-go/handler"
	"github.com/smartgreeting/mini-go/svc"
	"github.com/smartgreeting/mini-go/utils"
)

func main() {
	cfg, err := utils.InitConf("./config/config.yaml")
	if err != nil {
		panic(err)
	}
	svcCtx := svc.NewSvcContext(cfg)
	r := handler.SetupRouter(svcCtx)
	r.Run(fmt.Sprintf("%s:%d", cfg.Application.Address, cfg.Application.Port))

}
