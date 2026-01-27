package interfaces

import "context"

type INotificationService interface {
	SendRegisterNotification(ctx context.Context, userID int64, email string)
}
