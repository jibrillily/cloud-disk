package logic

import (
	"cloud_disk/core/define"
	"cloud_disk/core/helper"
	"cloud_disk/core/models"
	"context"
	"errors"
	"time"

	"cloud_disk/core/internal/svc"
	"cloud_disk/core/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MailCodeSendRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailCodeSendRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) MailCodeSendRegisterLogic {
	return MailCodeSendRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailCodeSendRegisterLogic) MailCodeSendRegister(req types.MailCodeSendRequest) (resp *types.MailCodeSendReply, err error) {
	//邮箱未被注册
	count, err := l.svcCtx.Engine.Where("email = ?", req.Email).Count(new(models.UserBasic))
	if err != nil {
		return
	}
	if count > 0 {
		err = errors.New("该邮箱已被注册")
		return
	}

	//获取验证码
	code := helper.RandCode()
	//存储验证码在redis中
	l.svcCtx.RDB.Set(l.ctx, req.Email, code, time.Second*time.Duration(define.CodeExpire))
	//发送验证码
	err = helper.MailCodeSend(req.Email, code)

	return
}
