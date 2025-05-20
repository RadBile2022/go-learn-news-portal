package convert

import (
	"fmt"
	"github.com/google/uuid"
	"math/big"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

func Param(r *http.Request, key string) string {
	return r.URL.Query().Get(key)
}

func Time(t *time.Time) *string {
	if t != nil {
		formatted := t.Format(time.DateTime)
		return &formatted
	}

	return nil
}

func GenerateOTP() int64 {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a 4-digit OTP
	otp := rand.Int63n(9000) + 1000 // Generates a number between 1000 and 9999
	return otp
}

func IDUint(r *http.Request) uint {
	ID, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
	if err != nil {
		return 0
	}
	return uint(ID)
}

func PathValueIDUintChi(r *http.Request) uint {
	ID, err := strconv.ParseUint(r.PathValue("id"), 10, 32)
	if err != nil {
		return 0
	}
	return uint(ID)
}

func PathValueIDInt64Chi(r *http.Request) int64 {
	ID, err := strconv.ParseInt(r.PathValue("id"), 10, 32)
	if err != nil {
		return 0
	}
	return int64(ID)
}

const base62 = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func GenerateString(length int64) string {
	result := make([]byte, length)
	rand.Seed(time.Now().UnixNano())
	for i := range result {
		result[i] = base62[rand.Intn(len(base62))]
	}
	return string(result)
}

func GenerateFileName() string {
	now := time.Now()
	dateNow := now.Format("20060102150405.000")
	uuidStr := uuid.New().String() // misal: "b8a2e7e6-3a4d-4a0d-a1f8-05b87d7b916e"
	// ambil bagian depan saja kalau ingin pendek:
	uuidShort := uuidStr[:8]
	//randomString := GenerateString(7)
	result := fmt.Sprintf("%s-%s", dateNow, uuidShort)
	return result
}

func GenerateFileNameWithSuffix(suffix string) string {
	filename := GenerateFileName()
	result := fmt.Sprintf("%s-%s", filename, suffix)
	return result
}

func GenerateFileNameWithExt(userId int64, ext string) string {
	filename := GenerateFileName()
	result := fmt.Sprintf("%d/%s.%s", userId, filename, ext)
	return result
}

func TruncateFloat64(f float64, unit float64) float64 {
	bf := big.NewFloat(0).SetPrec(100).SetFloat64(f)
	bu := big.NewFloat(0).SetPrec(100).SetFloat64(unit)

	bf.Quo(bf, bu)

	// Truncate:
	i := big.NewInt(0)
	bf.Int(i)
	bf.SetInt(i)

	f, _ = bf.Mul(bf, bu).Float64()
	return f
}

func UintToStr(v uint) string {
	return fmt.Sprintf("%d", v)
}

func Float64ToStr(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func PointerStr(v string) *string {
	return &v
}

func FloatFixed2(v float64) float64 {
	str := fmt.Sprintf("%.2f", v)
	r, _ := strconv.ParseFloat(str, 64)
	return r
}
