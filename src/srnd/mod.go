//
// mod.go
// post moderation
//
package srnd

import (
  "fmt"
  "net/http"
  "strings"
  "time"
)


// interface for moderation ui
type ModUI interface {

  // channel for daemon to poll for nntp articles from the mod ui
  MessageChan() chan *NNTPMessage

  // check if this key is allowed to access
  // return true if it can otherwise false
  CheckKey(privkey string) (bool, error)

  // an http handler function
  ServeHTTP(wr http.ResponseWriter, r *http.Request)
  
}

// moderation engine
type Moderation struct {
  // channel to send commands down line by line after they are authenticated
  feed chan string
  daemon *NNTPDaemon
}


func (self Moderation) Init(d *NNTPDaemon) {
  self.daemon = d
}

// TODO: implement
func (self Moderation) AllowPubkey(key string) bool {
  return false
}


type ModEvent interface {
  // turn it into a string for putting into an article
  String() string
  // what type of mod event
  Action() string
  // what reason for the event
  Reason() string
  // what is the event acting on
  Target() string
  // scope of the event, regex of newsgroup
  Scope() string
  // when this mod event expires, unix nano
  Expires() int64
}

type simpleModEvent string

func (self simpleModEvent) String() string {
  return string(self)
}

func (self simpleModEvent) Action() string {
  return strings.Split(string(self), " ")[0]
}

func (self simpleModEvent) Reason() string {
  return ""
}

func (self simpleModEvent) Target() string {
  return strings.Split(string(self), " ")[1]
}

func (self simpleModEvent) Scope() string {
  // TODO: hard coded
  return "overchan.*"
}

func (self simpleModEvent) Expires() int64 {
  // no expiration
  return -1
}

// moderation message
type ModMessage struct {
  Date time.Time
  Events []ModEvent
}

func (self ModMessage) String() string {
  body := "Content-Type: text/plain; charset=UTF-8\n"
  body += fmt.Sprintf("Date: %s\n", self.Date.Format(time.RFC1123Z))
  for _, ev := range(self.Events) {
    body += ev.String()
  }
  return body
}



func ParseModEvent(line string) ModEvent {
  return simpleModEvent(line)
}
