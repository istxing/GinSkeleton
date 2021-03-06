package Users

import (
	"GinSkeleton/App/Global/Consts"
	"GinSkeleton/App/Http/Controller/Admin"
	"GinSkeleton/App/Http/Validator/Core/DaTaTransfer"
	"GinSkeleton/App/Utils/Response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Destroy struct {
	Id float64 `form:"id"  json:"id" binding:"required,min=1"`
}

// 验证器语法，参见 Register.go文件，有详细说明

func (d *Destroy) CheckParams(context *gin.Context) {

	if err := context.ShouldBind(d); err != nil {
		errs := gin.H{
			"tips": "UserDestroy参数校验失败，参数校验失败，请检查id(>=1)",
			"err":  err.Error(),
		}
		Response.ReturnJson(context, http.StatusBadRequest, Consts.Validator_ParamsCheck_Fail_Code, Consts.Validator_ParamsCheck_Fail_Msg, errs)
		return
	}

	//  该函数主要是将绑定的数据以 键=>值 形式直接传递给下一步（控制器）
	extraAddBindDataContext := DaTaTransfer.DataAddContext(d, Consts.Validator_Prefix, context)
	if extraAddBindDataContext == nil {
		Response.ReturnJson(context, http.StatusInternalServerError, Consts.Server_Occurred_Error_Code, Consts.Server_Occurred_Error_Msg+",UserShow表单参数验证器json化失败..", "")
		return
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&Admin.Users{}).Destroy(extraAddBindDataContext)
	}
}
