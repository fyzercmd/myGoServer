// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package policy

import (
	v1 "github.com/fyzercmd/myGoServer/pkg/api/apiserver/v1"
	"github.com/fyzercmd/myGoServer/pkg/core"
	"github.com/fyzercmd/myGoServer/pkg/errors"
	metav1 "github.com/fyzercmd/myGoServer/pkg/meta/v1"
	"github.com/gin-gonic/gin"

	"github.com/fyzercmd/myGoServer/internal/pkg/code"
	"github.com/fyzercmd/myGoServer/internal/pkg/middleware"
	"github.com/fyzercmd/myGoServer/pkg/log"
)

// Update updates policy by the policy identifier.
func (p *PolicyController) Update(c *gin.Context) {
	log.L(c).Info("update policy function called.")

	var r v1.Policy
	if err := c.ShouldBindJSON(&r); err != nil {
		core.WriteResponse(c, errors.WithCode(code.ErrBind, err.Error()), nil)

		return
	}

	pol, err := p.srv.Policies().Get(c, c.GetString(middleware.UsernameKey), c.Param("name"), metav1.GetOptions{})
	if err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	// only update policy string
	pol.Policy = r.Policy
	pol.Extend = r.Extend

	if errs := pol.Validate(); len(errs) != 0 {
		core.WriteResponse(c, errors.WithCode(code.ErrValidation, errs.ToAggregate().Error()), nil)

		return
	}

	if err := p.srv.Policies().Update(c, pol, metav1.UpdateOptions{}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, pol)
}
