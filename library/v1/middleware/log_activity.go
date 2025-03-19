package middleware

//
//import (
//	"encoding/json"
//	"fmt"
//	"log"
//	"net/http"
//	"regexp"
//	"time"
//
//	"github.com/dev-digitalproject/go-simpp-auth-core/internal/adapters/core/entity"
//	"github.com/dev-digitalproject/go-simpp-auth-core/internal/ports/framework/secondary/repository"
//)
//
//type ResponseTimeRecord struct {
//	Endpoint       string  `json:"endpoint"`
//	ResponseTimeMs float64 `json:"response_time_ms"`
//	JsonResponse   string  `json:"json_response"`
//}
//
//func LogActivityMiddleware(repo repository.LogActivity) func(handler http.Handler) http.Handler {
//	return func(next http.Handler) http.Handler {
//		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//			path := r.URL.Path
//			startTime := time.Now()
//
//			rec := &responseWriterWrapper{ResponseWriter: w}
//
//			next.ServeHTTP(rec, r)
//
//			duration := time.Since(startTime)
//			var durationString string
//			if duration >= time.Second {
//				durationString = fmt.Sprintf("%.2f s", duration.Seconds())
//			} else {
//				durationString = fmt.Sprintf("%d ms", duration.Milliseconds())
//			}
//
//			// var userId string
//			// userId := extractToken();
//
//			var userID string
//			if userID == "" {
//				userID = "unknown" // Jika tidak ada user_id, bisa menggunakan nilai default
//			}
//
//			var responseJSON map[string]interface{}
//			if err := json.Unmarshal(rec.responseBody, &responseJSON); err != nil {
//				log.Println("Error unmarshaling response:", err)
//				return
//			}
//
//			responseJSONBytes, err := json.Marshal(responseJSON)
//			if err != nil {
//				log.Println("Error marshaling response:", err)
//				return
//			}
//
//			re := regexp.MustCompile(`^/api/v1/`)
//			path = re.ReplaceAllString(path, "/") // Menghilangkan api/v1/
//
//			// Menghapus bagian dinamis yang berupa angka atau apapun setelah '/'
//			// re = regexp.MustCompile(`/[^/]+/[^/]+$`) // Menangani bagian setelah '/users/1/reject' atau '/risalah-rapat/approve'
//			// path = re.ReplaceAllString(path, "")     // H
//
//			repo.Create(r.Context(), &entity.LogActivity{
//				Event:     path,
//				Method:    r.Method,
//				PathURL:   r.URL.Path,
//				IPAddr:    r.RemoteAddr,
//				UserAgent: r.UserAgent(),
//				Code:      fmt.Sprintf("%d", rec.statusCode),
//				Duration:  durationString,
//				Metadata:  responseJSONBytes,
//			})
//		})
//	}
//}
//
//type responseWriterWrapper struct {
//	http.ResponseWriter
//	statusCode   int
//	responseBody []byte
//}
//
//func (rw *responseWriterWrapper) WriteHeader(statusCode int) {
//	rw.statusCode = statusCode
//	rw.ResponseWriter.WriteHeader(statusCode)
//}
//
//func (rw *responseWriterWrapper) Write(p []byte) (n int, err error) {
//	rw.responseBody = p
//	return rw.ResponseWriter.Write(p)
//}
