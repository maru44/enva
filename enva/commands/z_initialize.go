package commands

import "github.com/maru44/enva/service/api/pkg/domain"

func init() {
	Commands["pull"] = func() domain.ICommandInteractor {
		return &pull{}
	}

	Commands["get"] = func() domain.ICommandInteractor {
		return &get{}
	}

	Commands["add"] = func() domain.ICommandInteractor {
		return &add{}
	}

	Commands["edit"] = func() domain.ICommandInteractor {
		return &edit{}
	}

	Commands["delete"] = func() domain.ICommandInteractor {
		return &delete{}
	}

	Commands["diff"] = func() domain.ICommandInteractor {
		return &diff{}
	}

	Commands["init"] = func() domain.ICommandInteractor {
		return &initialize{}
	}

	Commands["set"] = func() domain.ICommandInteractor {
		return &set{}
	}

	Commands["help"] = func() domain.ICommandInteractor {
		return &help{}
	}

	Commands["version"] = func() domain.ICommandInteractor {
		return &version{}
	}
}
