// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package secret

import (
	"github.com/fyzercmd/myGoServer/pkg/core"
	"github.com/fyzercmd/myGoServer/pkg/errors"
	metav1 "github.com/fyzercmd/myGoServer/pkg/meta/v1"
	"github.com/gin-gonic/gin"

	"github.com/fyzercmd/myGoServer/internal/pkg/code"
	"github.com/fyzercmd/myGoServer/internal/pkg/middleware"
	"github.com/fyzercmd/myGoServer/pkg/log"
)

// List list all the secrets.
func (s *SecretController) List(c *gin.Context) {
	log.L(c).Info("list secret function called.")
	var r metav1.ListOptions
	if err := c.ShouldBindQuery(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	secrets, err := s.srv.Secrets().List(c, c.GetString(middleware.UsernameKey), r)
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, secrets)
}
