// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package secret

import (
	"github.com/AlekSi/pointer"
	v1 "github.com/fyzercmd/myGoServer/pkg/api/apiserver/v1"
	"github.com/fyzercmd/myGoServer/pkg/core"
	"github.com/fyzercmd/myGoServer/pkg/errors"
	metav1 "github.com/fyzercmd/myGoServer/pkg/meta/v1"
	"github.com/gin-gonic/gin"

	"github.com/fyzercmd/myGoServer/internal/pkg/code"
	"github.com/fyzercmd/myGoServer/internal/pkg/middleware"
	"github.com/fyzercmd/myGoServer/pkg/log"
)

const maxSecretCount = 10

// Create add new secret key pairs to the storage.
func (s *SecretController) Create(c *gin.Context) {
	log.L(c).Info("create secret function called.")

	var r v1.Secret

	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	if errs := r.Validate(); len(errs) != 0 {
		core.WriteResponse(c, errors.WithCode(code.ErrValidation, errs.ToAggregate().Error()), nil)

		return
	}

	username := c.GetString(middleware.UsernameKey)

	secrets, err := s.srv.Secrets().List(c, username, metav1.ListOptions{
		Offset: pointer.ToInt64(0),
		Limit:  pointer.ToInt64(-1),
	})
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	if secrets.TotalCount >= maxSecretCount {
		core.WriteResponse(c, errors.WithCode(code.ErrReachMaxCount, "secret count: %d", secrets.TotalCount), nil)

		return
	}

	// must reassign username
	r.Username = username

	if err := s.srv.Secrets().Create(c, &r, metav1.CreateOptions{}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, r)
}
