package mw

import (
    "net/http"
    "strings"
)

// CORS ตอบ header อนุญาตเฉพาะ origin ที่อยู่ในรายชื่อ ไม่ใช้ * เพราะใช้กับ Authorization ไม่ได้
func CORS(allowed []string, next http.Handler) http.Handler {
    ok := map[string]bool{}
    for _, o := range allowed {
        if o = strings.TrimSpace(o); o != "" {
            ok[o] = true
        }
    }
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Vary", "Origin")
        if origin := r.Header.Get("Origin"); ok[origin] {
            w.Header().Set("Access-Control-Allow-Origin", origin)
            w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
            w.Header().Set("Access-Control-Allow-Headers", "Authorization, Content-Type")
        }
        // preflight ตอบ 204 ทันที ไม่ส่งต่อให้ handler
        if r.Method == http.MethodOptions {
            w.WriteHeader(http.StatusNoContent)
            return
        }
        next.ServeHTTP(w, r)
    })
}
