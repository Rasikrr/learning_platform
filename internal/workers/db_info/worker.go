// nolint
package db_info

//import (
//	"context"
//	usersR "github.com/Rasikrr/learning_platform/internal/repositories/users"
//	"github.com/Rasikrr/learning_platform/internal/workers"
//	"log"
//	"time"
//)
//
//const (
//	workerTimeout = time.Second * 10
//)
//
//type Worker struct {
//	userRepo usersR.Repository
//}
//
//func NewWorker(userRepo usersR.Repository) workers.Worker {
//	return &Worker{
//		userRepo: userRepo,
//	}
//}
//
//func (w *Worker) Run(ctx context.Context) error {
//	ticker := time.NewTicker(workerTimeout)
//	for {
//		select {
//		case <-ctx.Done():
//			log.Println("DB info worker stopped")
//			return nil
//		case <-ticker.C:
//			users, err := w.userRepo.GetAll(ctx)
//			if err != nil {
//				log.Printf("error while logging users: %v", err)
//				continue
//			}
//			for _, user := range users {
//				log.Printf("user info: %+v", user)
//			}
//			log.Println("user count: ", len(users))
//		}
//	}
// nolint
//}
