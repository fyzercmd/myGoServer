// Copyright 2020 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package policy

import (
	"github.com/fyzercmd/myGoServer/pkg/core"
	metav1 "github.com/fyzercmd/myGoServer/pkg/meta/v1"
	"github.com/gin-gonic/gin"

	"github.com/fyzercmd/myGoServer/internal/pkg/middleware"
	"github.com/fyzercmd/myGoServer/pkg/log"
)

// Delete deletes the policy by the policy identifier.
func (p *PolicyController) Delete(c *gin.Context) {
	log.L(c).Info("delete policy function called.")

	if err := p.srv.Policies().Delete(c, c.GetString(middleware.UsernameKey), c.Param("name"),
		metav1.DeleteOptions{}); err != nil {
		core.WriteResponse(c, err, nil)

		return
	}

	core.WriteResponse(c, nil, nil)
}
