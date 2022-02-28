package domain

//go:generate go run github.com/golang/mock/mockgen -destination mock$GOPACKAGE/gen_mockgen.go -package mock$GOPACKAGE . IKvInteractor,ICliKvInteractor,IOrgInteractor,IProjectInteractor,ISmtpInteractor,IUserInteractor
