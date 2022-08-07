package proError

import (
	"github.com/pkg/errors"
	"github.com/xtclalala/ScanNetWeb/tools"
)

var (
	SUCCESS = 200
	ERROR   = 500

	// Controller Error
	ParamResolveFault = 1001

	// task crud Error
	UsernameAndPasswdError = 2001
	FindTaskError          = 2102
	FindRoleError          = 2103
	FindPermissionError    = 2104
	FindMenuError          = 2105
	SearchTaskResultError  = 2106
	SearchTaskError        = 2107

	UpdateRoleError      = 2201
	UpdateRoleMenusError = 2202
	UpdateRolePerError   = 2203
	UpdateTaskError      = 2204
	UpdateOrgBaseError   = 2205
	FindFileError        = 2206

	CreateRoleError         = 2301
	CreateTaskError         = 2302
	CreateOrganizationError = 2303

	DeleteUserError         = 2401
	DeleteRoleError         = 2402
	DeleteTaskError         = 2403
	DeleteOrganizationError = 2404
	DeleteMenuError         = 2405

	RunTaskError = 2501

	UploadFileError = 2604
	// Token Error
	TokenExpired      = 3001
	TokenNotValid     = 3002
	TokenInvalid      = 3003
	TokenMalformed    = 3004
	ClaimParseFailed  = 3005
	TokenCreateFailed = 3006
	NOTOKEN           = 3007

	FileReadType = 4001
	FileSave     = 4002

	// task Error
	TaskCantRevise = 4003
)

var codeMsg = map[int]string{
	SUCCESS: "操作成功",
	ERROR:   "操作失败",

	ParamResolveFault: "解析参数失败",

	UsernameAndPasswdError: "sys.user.usernameAndPasswdError",
	FindTaskError:          "sys.user.findOrgError",
	FindRoleError:          "sys.user.findRoleError",
	FindPermissionError:    "sys.user.findPermissionError",
	FindMenuError:          "sys.user.findMenuError",
	SearchTaskResultError:  "sys.user.findUserError",
	SearchTaskError:        "搜索任务失败",

	UpdateRoleError:      "sys.user.updateRoleError",
	UpdateRoleMenusError: "sys.user.updateRoleMenusError",
	UpdateRolePerError:   "sys.user.updateRolePerError",
	UpdateTaskError:      "修改任务失败",
	UpdateOrgBaseError:   "sys.user.updateOrgBaseError",
	FindFileError:        "sys.user.updateMenuBaseError",

	CreateRoleError:         "sys.user.createRoleError",
	CreateTaskError:         "创建任务失败",
	CreateOrganizationError: "sys.user.createOrganizationError",

	DeleteUserError:         "sys.user.deleteUserError",
	DeleteRoleError:         "sys.user.deleteRoleError",
	DeleteTaskError:         "删除任务失败",
	DeleteOrganizationError: "sys.user.deleteOrganizationError",
	DeleteMenuError:         "sys.user.deleteMenuError",

	RunTaskError: "启动任务失败",

	UploadFileError: "上传文件失败",

	NOTOKEN:           "sys.token.noToken",
	TokenExpired:      "sys.token.expired",
	TokenInvalid:      "sys.token.invalid",
	TokenMalformed:    "sys.token.malformed",
	TokenNotValid:     "sys.token.notValid",
	ClaimParseFailed:  "sys.claim.parseError",
	TokenCreateFailed: "sys.token.createFailed",

	FileReadType: "sys.file.read",
	FileSave:     "sys.file.save",

	TaskCantRevise: "任务已经不能改变",
}

func GetErrorMessage(code int) string {
	return codeMsg[code]
}

func Wrap(err error, s string, args ...any) error {
	return errors.Wrapf(err, tools.Out(s), args)
}

func WrapOrNil(err error, s string, args ...any) error {
	if err != nil {
		return errors.Wrapf(err, tools.Out(s), args)
	}
	return nil
}
