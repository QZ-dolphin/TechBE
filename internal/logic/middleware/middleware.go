package middleware

import (
	"TechBE/internal/service"
	"sync"
	"time"

	"github.com/gogf/gf/v2/net/ghttp"
	"golang.org/x/time/rate"
)

type sMiddleware struct{}

func init() {
	service.RegisterMiddleware(New())
}

func New() service.IMiddleware {
	return &sMiddleware{}
}

func (s *sMiddleware) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s *sMiddleware) Response(r *ghttp.Request) {
	r.Middleware.Next()
	res := r.GetHandlerResponse()
	r.Response.WriteJson(res)
}

// 1. 全局限流器（限制整个服务的请求速率）
var globalLimiter = rate.NewLimiter(5, 10) // 每秒 5 个请求，最大突发 10 个

// 2. IP 限流器（基于 IP 进行限流）
var ipLimiters = make(map[string]*rate.Limiter)
var mu sync.Mutex

// 获取或创建 IP 限流器
func getIPLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := ipLimiters[ip]
	if !exists {
		limiter = rate.NewLimiter(1, 5) // 每秒 1 个请求，突发 5 个
		ipLimiters[ip] = limiter

		// 定期清理旧的 IP 限流器，防止内存泄漏
		go func(ip string) {
			time.Sleep(10 * time.Minute)
			mu.Lock()
			delete(ipLimiters, ip)
			mu.Unlock()
		}(ip)
	}

	return limiter
}

// 限流中间件
func (s *sMiddleware) RateLimit(r *ghttp.Request) {
	// 获取客户端 IP（GoFrame 提供的获取方式）
	clientIP := r.GetClientIp()

	// 全局限流检查
	if !globalLimiter.Allow() {
		r.Response.WriteStatus(429, "Too Many Requests (Global Limit)")
		return
	}

	// IP 级别限流
	ipLimiter := getIPLimiter(clientIP)
	if !ipLimiter.Allow() {
		r.Response.WriteStatus(429, "Too Many Requests (IP Limit)")
		return
	}

	// 继续处理请求
	r.Middleware.Next()
}
