// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package user

import (
	v1 "github.com/fyzercmd/myGoServer/pkg/api/apiserver/v1"
	"github.com/fyzercmd/myGoServer/pkg/core"
	"github.com/fyzercmd/myGoServer/pkg/errors"
	metav1 "github.com/fyzercmd/myGoServer/pkg/meta/v1"
	"github.com/gin-gonic/gin"

	"github.com/fyzercmd/myGoServer/internal/pkg/code"
	"github.com/fyzercmd/myGoServer/pkg/log"
)

// Create add new user to the storage.
func (u *UserController) Create(c *gin.Context) {
	log.L(c).Info("user create function called.")

	var r v1.User

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	if errs := r.Validate(); len(errs) != 0 {
		core.WriteResponse(c, errors.WithCode(code.ErrValidation, errs.ToAggregate().Error()), nil)

		return
	}

	// Insert the user to the storage.
	if err := u.srv.Users().Create(c, &r, metav1.CreateOptions{}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, r)
}
