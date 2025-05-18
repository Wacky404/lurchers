package handler

import (
    "log"
)

func EventImg(event string, callback func(string)) {
    log.Printf()
    callback(event)
}
