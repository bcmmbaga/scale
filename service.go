package scale

type (
	AccountStore interface {
	}
	MessageStore interface {
	}
	HistoryStore interface {
	}
)

type Service struct {
	AccountStore AccountStore
	MessageStore MessageStore
	HistoryStore HistoryStore
}

func NewService(accountStore AccountStore, messageStore MessageStore, historyStore HistoryStore) *Service {
	return &Service{AccountStore: accountStore, MessageStore: messageStore, HistoryStore: historyStore}
}
