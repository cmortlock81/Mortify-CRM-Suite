package auth
func Redact(token string) string { if len(token)<8 { return "***" }; return token[:4]+"..."+token[len(token)-4:] }
