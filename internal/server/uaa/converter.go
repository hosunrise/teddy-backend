package uaa

import (
	"github.com/golang/protobuf/ptypes"
	"teddy-backend/internal/models"
	"teddy-backend/internal/proto/uaa"
)

func copyFromAccountToPBAccount(acc *models.Account, pbacc *uaa.Account) error {
	if acc == nil || pbacc == nil {
		return nil
	}
	pbacc.Uid = acc.UID
	pbacc.Username = acc.Username
	pbacc.Email = acc.Email
	pbacc.Password = acc.Password
	pbacc.Locked = acc.Locked
	pbacc.CredentialsExpired = acc.CredentialsExpired
	pbacc.Roles = acc.Roles
	pbacc.OauthUIDs = acc.OAuthUIds
	pbacc.LastSignInIP = acc.LastSignInIP

	tmp, err := ptypes.TimestampProto(acc.CreateDate)
	if err != nil {
		return err
	}
	pbacc.CreateDate = tmp

	tmp, err = ptypes.TimestampProto(acc.UpdateDate)
	if err != nil {
		return err
	}
	pbacc.UpdateDate = tmp

	tmp, err = ptypes.TimestampProto(acc.LastSignInTime)
	if err != nil {
		return err
	}
	pbacc.LastSignInTime = tmp
	return nil
}
