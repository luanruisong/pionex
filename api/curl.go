package api

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/valyala/fasthttp"
)

const maxBodyForLog = 4096 // 最大打印的 body 长度，避免日志爆炸
// 将 fasthttp.Request 转为 curl 命令

func curlHook(req *fasthttp.Request) string {
	method := strings.ToUpper(string(req.Header.Method()))
	if method == "" {
		method = "GET"
	}
	url := string(req.URI().FullURI())
	if url == "" {
		url = req.URI().String()
	}

	// 首行
	first := fmt.Sprintf("curl --location --request %s %s", method, shellQuote(url))

	// 头部
	var lines []string
	compressed := false
	req.Header.VisitAll(func(k, v []byte) {
		key := strings.ToLower(string(k))
		switch key {
		case "host", "content-length":
			return
		case "accept-encoding":
			// 用 --compressed 代替显式的 Accept-Encoding 头
			compressed = true
			return
		}
		lines = append(lines, fmt.Sprintf("--header %s", shellQuote(fmt.Sprintf("%s: %s", string(k), string(v)))))
	})
	if compressed {
		lines = append(lines, "--compressed")
	}

	// Body
	body := req.Body()
	ct := strings.ToLower(string(req.Header.ContentType()))
	note := ""
	if len(body) > 0 {
		if len(body) > maxBodyForLog {
			note = fmt.Sprintf("# NOTE: body truncated to %d of %d bytes", maxBodyForLog, len(body))
			body = body[:maxBodyForLog]
		}
		flag := "--data-binary"
		if isTextualCT(ct) && isPrintableText(body) {
			flag = "--data-raw"
			lines = append(lines, fmt.Sprintf("%s %s", flag, shellQuote(string(body))))
		} else if isPrintableText(body) {
			lines = append(lines, fmt.Sprintf("%s %s", flag, shellQuote(string(body))))
		} else {
			// 二进制体就不直接塞进命令里了，避免污染终端
			if note == "" {
				note = fmt.Sprintf("# NOTE: binary body (%d bytes) not included", len(req.Body()))
			}
		}
	}

	// 组装成多行（最后一行不带续行反斜杠）
	if len(lines) == 0 {
		if note != "" {
			return first + "\n" + note
		}
		return first
	}
	return first + " \\\n" + strings.Join(lines, " \\\n") + func() string {
		if note != "" {
			return "\n" + note
		}
		return ""
	}()
}

// 单引号 shell 转义
func shellQuote(s string) string {
	if s == "" {
		return "''"
	}
	return "'" + strings.ReplaceAll(s, "'", `'"'"'`) + "'"
}

// 简单判断是否文本类 Content-Type
func isTextualCT(ct string) bool {
	if ct == "" {
		return true
	}
	if strings.HasPrefix(ct, "text/") {
		return true
	}
	for _, kw := range []string{
		"json", "xml", "x-www-form-urlencoded", "javascript", "graphql",
	} {
		if strings.Contains(ct, kw) {
			return true
		}
	}
	return false
}

// 判定是否可打印文本
func isPrintableText(b []byte) bool {
	if !utf8.Valid(b) {
		return false
	}
	for _, c := range b {
		if c == '\n' || c == '\r' || c == '\t' {
			continue
		}
		if c < 0x20 {
			return false
		}
	}
	return true
}
