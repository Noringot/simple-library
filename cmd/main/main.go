package main

import (
	"github.com/Noringotq/simple-grpc/internal/adapters/repository"
	"github.com/Noringotq/simple-grpc/internal/config"
	"github.com/Noringotq/simple-grpc/internal/usecase"
	"github.com/Noringotq/simple-grpc/pkg/lib"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// Получение экземпляра конфига
	cfg := config.GetConfig()
	dbAddr := cfg.MySQL.Host + ":" + cfg.MySQL.Port

	// Создание репозитория для взаимодействия с БД
	rep := repository.NewLibRepository(cfg.MySQL.Username, cfg.MySQL.Password, dbAddr, cfg.MySQL.Database)

	// Создание usecase, в котором находится бизнесс-логика и взаимодействие с репозиторием
	usc := usecase.NewLibStorage(rep)

	// Инициализация и запуск grpc сервера
	lis, err := net.Listen(cfg.Listen.Net, cfg.Listen.Address)
	if err != nil {
		log.Fatalf("cannot start server: %s", err)
	}

	s := grpc.NewServer()

	lib.RegisterLibServer(s, usc)
	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("Cannot serve grpc server: %s", err)
	}
}
