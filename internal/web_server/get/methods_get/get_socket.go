package methods_get

import (
	"encoding/json"
	"log"
	"net/http"
	"text/template"

	"projects/doc/doc_service/internal/web_server/sessions"
	"projects/doc/doc_service/internal/web_server/socket"

	"github.com/gin-gonic/gin"

	"github.com/gorilla/websocket"
)

var (
	UpGrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

type TSocket struct {
	tmp *template.Template
}

func newSocket(in *TInGetPage) IGetPage {
	t := &TSocket{
		tmp: in.Tmp,
	}

	return t
}

func (t *TSocket) GetPath() string {
	return "/api/ws"
}

func (t *TSocket) GetContext(c *gin.Context) {
	var err error
	ses := sessions.Ses.GetSes(c)
	if ses == nil {
		return
	}

	if !ses.Authorization() {
		return
	}

	defer func() {
		if r := recover(); r != nil {
			log.Printf("NewGet(): ошибка сокета, err=%v", r)
		}
	}()

	conn := ses.GetConn()
	// ses, err := sessions

	UpGrader.CheckOrigin = func(req *http.Request) bool {
		// allow all connections by default
		return true
	}

	conn.Conn, err = UpGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Printf("[ERROR] TSocket.GetContext(): чтение сессии, err=%v", err)
	}

	constructors, err := socket.NewMethodsSocket()
	if err != nil {
		log.Printf("[ERROR] TSocket.GetContext(): получаю конвеиер, err=%v", err)
	}
	for {
		mType, msg, err := conn.ReadMessage()
		if ce, ok := err.(*websocket.CloseError); ok {
			close := false

			switch ce.Code {
			case websocket.CloseGoingAway:
				close = true
			default:
				log.Printf("[ERROR] TSocket.GetContext(): чтение сообщения, err=%v", err)
				close = true
			}

			if close {
				break
			}
		}
		if msg == nil {
			continue
		}

		var dataWs interface{}
		err = json.Unmarshal(msg, &dataWs)
		if err != nil {
			log.Panicln(err)
			break
		}

		d, ok := dataWs.(map[string]interface{})
		if !ok {
			log.Println("[ERROR] Ошибка приведения dataWs")
			continue
		}

		tp, ok := d["tp"].(string)
		if !ok {
			continue
		}

		cmd, ok := d["cmd"].(string)
		if !ok {
			continue
		}

		var f socket.ISocket

		fn, ok := constructors.Get(tp)
		if !ok {
			log.Println("[WARNING] GetPage(): Конструктор типа: ", tp, " не задан")
			continue
		}

		f, err = fn(&socket.TSocketValue{
			Data: d,
			Ses:  ses,
		})
		if err != nil {
			log.Println("[WARNING] GetPage(): Не удалось создать конструктор типа: ", tp, ", err=%w", err)
			continue
		}

		switch cmd {
		case "Start":
			err = f.Start()
			if err != nil {
				log.Println("[ERROR] Ошибка конструктора, err=", err)
			}
		case "Stop":
			err = f.Stop()
			if err != nil {
				log.Println("[ERROR] Ошибка остановки конструктора, err=", err)
			}
		}

		resp, err := f.Response()
		if err != nil {
			log.Println("[ERROR] RunUserSocket(): чтение ответа, err=%w", err)
		}

		if resp != nil {
			d = resp
		}

		// Out data
		wrWs, err := json.Marshal(&d)
		if err != nil {
			log.Println("[ERROR] RunUserSocket(): error marshal json", err)
		}

		// sesRec.Conn.Lock()
		if err = conn.WriteMessage(mType, wrWs); err != nil {
			// sesRec.Conn.Unlock()
			break
		}
		// sesRec.Conn.Unlock()

	}

	err = conn.Close()
	if err != nil {
		log.Println("[ERROR] TSocket.GetContext(): error marshal json", err)
	}
}

func init() {
	err := constructors.Add("NewSocket", newSocket)
	if err != nil {
		log.Printf("NewSocket(): не удалось добавить в конструктор")
	}
}
