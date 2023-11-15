package monitor

import (
	"log"
	"os"
	"projects/doc/doc_service/internal/web_server/sessions"
	"projects/doc/doc_service/pkg/types"
	"sync"
	"time"
)

var monitor types.IMonitor

type tStore struct {
	sync.Mutex
	data map[string]int64
}
type tMonitor struct {
	store tStore
	tick  chan string
	next  chan bool
}

func New() (types.IMonitor, error) {
	t := &tMonitor{
		store: tStore{
			data: make(map[string]int64),
		},
		tick: make(chan string, 1000),
		next: make(chan bool, 100),
	}

	t.run()

	return t, nil
}
func (t *tMonitor) run() {
	go func() {
		for {
			time.Sleep(1 * time.Second)
			t.next <- true
		}
	}()
	go func() {

		for {
			select {
			case key := <-t.tick:
				t.store.Lock()
				val, ok := t.store.data[key]
				if !ok {
					t.store.data[key] = 1
				}
				t.store.data[key] = val + 1
				t.store.Unlock()
			case <-t.next:
				sessions.Ses.RangeSes(func(ses sessions.ISession) {
					defer func() {
						if r := recover(); r != nil {
							log.Printf("NewGet(): критическая ошибка, err=%v", r)
						}
					}()
					if ses.GetCurrentPage() != "/gui/" {
						return
					}
					ses.SendMessage(map[string]interface{}{
						"tp":     "ChartTick",
						"charts": t.store.data,
					})
				})

				t.store.Lock()
				for key := range t.store.data {
					t.store.data[key] = 0
				}
				t.store.Unlock()
			}
		}
	}()
}

func Monitor() types.IMonitor {
	return monitor
}

func (t *tMonitor) Add(key string) {
	if t.tick == nil {
		return
	}
	t.tick <- key
}

func init() {
	var err error
	monitor, err = New()
	if err != nil {
		log.Printf("monitor.init(): создание объекта, err = %v", err)
		os.Exit(1)
	}
}
