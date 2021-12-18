package commands

import (
	"context"

	"github.com/maru44/enva/service/api/internal/interface/database"
	"github.com/maru44/enva/service/api/internal/interface/password"
	"github.com/maru44/enva/service/api/internal/usecase"
	"github.com/maru44/enva/service/api/pkg/domain"
)

type kvList struct {
	kIn domain.IKvInteractor
	cIn domain.ICliUserInteractor
}

func init() {
	Commands["get"] = func(sql database.ISqlHandler, pass password.IPassword) domain.ICommandInteractor {
		return &kvList{
			kIn: usecase.NewKvInteractor(
				&database.KvRepository{ISqlHandler: sql},
			),
			cIn: usecase.NewCliUserInteractor(
				&database.CliUserRepository{
					ISqlHandler: sql,
					IPassword:   pass,
				},
			),
		}
	}
}

func (c *kvList) Run(ctx context.Context, opts ...string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// settings, err := readSettings()
	// if err != nil {
	// 	return err
	// }

	// create input

	//

	return nil
}
