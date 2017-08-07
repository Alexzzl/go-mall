package response

import (
    //"encoding/json"
)

func SuccRes() map[string]interface{} {
    return map[string]interface{} {
        "ok": true,
    }
}

func FailRes(message interface {}) map[string]interface{} {
    return map[string]interface{} {
        "ok": false,
        "message": message,
    }
}
