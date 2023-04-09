package curlcolor

import (
	"errors"
	"log"
	"strings"
)

type CurlParameter map[string]*CurlParam

func (p *CurlParameter) GetBool(key string) bool {
	param, ok := p.Get(key)
	if !ok {
		return false
	}
	return param.boolValue
}

func (p *CurlParameter) GetString(key string) string {
	param, ok := p.Get(key)
	if !ok {
		return ""
	}
	return param.stringValue
}

func (p *CurlParameter) Get(key string) (*CurlParam, bool) {
	param, ok := (*p)[key]
	return param, ok
}

type ArgType int

const (
	ArgNone ArgType = iota
	ArgBool
	ArgString
	ArgFilename
)

type CurlParam struct {
	letter      string
	lname       string
	desc        ArgType
	boolValue   bool
	stringValue string
}

func (p *CurlParam) SetVal(argv []string) ([]string, error) {
	switch p.desc {
	case ArgBool, ArgNone:
		p.boolValue = true
	case ArgString, ArgFilename:
		if len(argv) == 0 {
			return argv, errors.New("missing argument")
		}
		p.stringValue = argv[0]
		argv = argv[1:]
	}
	return argv, nil
}

var (
	longNameParameter = &CurlParameter{
		"url":                        {"*@", "url", ArgString, false, ""},
		"dns-ipv4-addr":              {"*4", "dns-ipv4-addr", ArgString, false, ""},
		"dns-ipv6-addr":              {"*6", "dns-ipv6-addr", ArgString, false, ""},
		"random-file":                {"*a", "random-file", ArgFilename, false, ""},
		"egd-file":                   {"*b", "egd-file", ArgString, false, ""},
		"oauth2-bearer":              {"*B", "oauth2-bearer", ArgString, false, ""},
		"connect-timeout":            {"*c", "connect-timeout", ArgString, false, ""},
		"doh-url":                    {"*C", "doh-url", ArgString, false, ""},
		"ciphers":                    {"*d", "ciphers", ArgString, false, ""},
		"dns-interface":              {"*D", "dns-interface", ArgString, false, ""},
		"disable-epsv":               {"*e", "disable-epsv", ArgBool, false, ""},
		"disallow-username-in-url":   {"*f", "disallow-username-in-url", ArgBool, false, ""},
		"epsv":                       {"*E", "epsv", ArgBool, false, ""},
		"dns-servers":                {"*F", "dns-servers", ArgString, false, ""},
		"trace":                      {"*g", "trace", ArgFilename, false, ""},
		"npn":                        {"*G", "npn", ArgBool, false, ""},
		"trace-ascii":                {"*h", "trace-ascii", ArgFilename, false, ""},
		"alpn":                       {"*H", "alpn", ArgBool, false, ""},
		"limit-rate":                 {"*i", "limit-rate", ArgString, false, ""},
		"rate":                       {"*I", "rate", ArgString, false, ""},
		"compressed":                 {"*j", "compressed", ArgBool, false, ""},
		"tr-encoding":                {"*J", "tr-encoding", ArgBool, false, ""},
		"digest":                     {"*k", "digest", ArgBool, false, ""},
		"negotiate":                  {"*l", "negotiate", ArgBool, false, ""},
		"ntlm":                       {"*m", "ntlm", ArgBool, false, ""},
		"ntlm-wb":                    {"*M", "ntlm-wb", ArgBool, false, ""},
		"basic":                      {"*n", "basic", ArgBool, false, ""},
		"anyauth":                    {"*o", "anyauth", ArgBool, false, ""},
		"wdebug":                     {"*p", "wdebug", ArgBool, false, ""},
		"ftp-create-dirs":            {"*q", "ftp-create-dirs", ArgBool, false, ""},
		"create-dirs":                {"*r", "create-dirs", ArgBool, false, ""},
		"create-file-mode":           {"*R", "create-file-mode", ArgString, false, ""},
		"max-redirs":                 {"*s", "max-redirs", ArgString, false, ""},
		"proxy-ntlm":                 {"*desc", "proxy-ntlm", ArgBool, false, ""},
		"crlf":                       {"*u", "crlf", ArgBool, false, ""},
		"stderr":                     {"*v", "stderr", ArgFilename, false, ""},
		"aws-sigv4":                  {"*V", "aws-sigv4", ArgString, false, ""},
		"interface":                  {"*w", "interface", ArgString, false, ""},
		"krb":                        {"*x", "krb", ArgString, false, ""},
		"krb4":                       {"*x", "krb4", ArgString, false, ""},
		"haproxy-protocol":           {"*X", "haproxy-protocol", ArgBool, false, ""},
		"max-filesize":               {"*y", "max-filesize", ArgString, false, ""},
		"disable-eprt":               {"*z", "disable-eprt", ArgBool, false, ""},
		"eprt":                       {"*Z", "eprt", ArgBool, false, ""},
		"xattr":                      {"*~", "xattr", ArgBool, false, ""},
		"ftp-ssl":                    {"$a", "ftp-ssl", ArgBool, false, ""},
		"ssl":                        {"$a", "ssl", ArgBool, false, ""},
		"ftp-pasv":                   {"$b", "ftp-pasv", ArgBool, false, ""},
		"socks5":                     {"$c", "socks5", ArgString, false, ""},
		"tcp-nodelay":                {"$d", "tcp-nodelay", ArgBool, false, ""},
		"proxy-digest":               {"$e", "proxy-digest", ArgBool, false, ""},
		"proxy-basic":                {"$f", "proxy-basic", ArgBool, false, ""},
		"retry":                      {"$g", "retry", ArgString, false, ""},
		"retry-connrefused":          {"$V", "retry-connrefused", ArgBool, false, ""},
		"retry-delay":                {"$h", "retry-delay", ArgString, false, ""},
		"retry-max-time":             {"$i", "retry-max-time", ArgString, false, ""},
		"proxy-negotiate":            {"$k", "proxy-negotiate", ArgBool, false, ""},
		"form-escape":                {"$l", "form-escape", ArgBool, false, ""},
		"ftp-account":                {"$m", "ftp-account", ArgString, false, ""},
		"proxy-anyauth":              {"$n", "proxy-anyauth", ArgBool, false, ""},
		"trace-time":                 {"$o", "trace-time", ArgBool, false, ""},
		"ignore-content-length":      {"$p", "ignore-content-length", ArgBool, false, ""},
		"ftp-skip-pasv-ip":           {"$q", "ftp-skip-pasv-ip", ArgBool, false, ""},
		"ftp-method":                 {"$r", "ftp-method", ArgString, false, ""},
		"local-port":                 {"$s", "local-port", ArgString, false, ""},
		"socks4":                     {"$desc", "socks4", ArgString, false, ""},
		"socks4a":                    {"$T", "socks4a", ArgString, false, ""},
		"ftp-alternative-to-user":    {"$u", "ftp-alternative-to-user", ArgString, false, ""},
		"ftp-ssl-reqd":               {"$v", "ftp-ssl-reqd", ArgBool, false, ""},
		"ssl-reqd":                   {"$v", "ssl-reqd", ArgBool, false, ""},
		"sessionid":                  {"$w", "sessionid", ArgBool, false, ""},
		"ftp-ssl-control":            {"$x", "ftp-ssl-control", ArgBool, false, ""},
		"ftp-ssl-ccc":                {"$y", "ftp-ssl-ccc", ArgBool, false, ""},
		"ftp-ssl-ccc-mode":           {"$j", "ftp-ssl-ccc-mode", ArgString, false, ""},
		"libcurl":                    {"$z", "libcurl", ArgString, false, ""},
		"raw":                        {"$#", "raw", ArgBool, false, ""},
		"post301":                    {"$0", "post301", ArgBool, false, ""},
		"keepalive":                  {"$1", "keepalive", ArgBool, false, ""},
		"socks5-hostname":            {"$2", "socks5-hostname", ArgString, false, ""},
		"keepalive-time":             {"$3", "keepalive-time", ArgString, false, ""},
		"post302":                    {"$4", "post302", ArgBool, false, ""},
		"noproxy":                    {"$5", "noproxy", ArgString, false, ""},
		"socks5-gssapi-nec":          {"$7", "socks5-gssapi-nec", ArgBool, false, ""},
		"proxy1.0":                   {"$8", "proxy1.0", ArgString, false, ""},
		"tftp-blksize":               {"$9", "tftp-blksize", ArgString, false, ""},
		"mail-from":                  {"$A", "mail-from", ArgString, false, ""},
		"mail-rcpt":                  {"$B", "mail-rcpt", ArgString, false, ""},
		"ftp-pret":                   {"$C", "ftp-pret", ArgBool, false, ""},
		"proto":                      {"$D", "proto", ArgString, false, ""},
		"proto-redir":                {"$E", "proto-redir", ArgString, false, ""},
		"resolve":                    {"$F", "resolve", ArgString, false, ""},
		"delegation":                 {"$G", "delegation", ArgString, false, ""},
		"mail-auth":                  {"$H", "mail-auth", ArgString, false, ""},
		"post303":                    {"$I", "post303", ArgBool, false, ""},
		"metalink":                   {"$J", "metalink", ArgBool, false, ""},
		"sasl-authzid":               {"$6", "sasl-authzid", ArgString, false, ""},
		"sasl-ir":                    {"$K", "sasl-ir", ArgBool, false, ""},
		"test-event":                 {"$L", "test-event", ArgBool, false, ""},
		"unix-socket":                {"$M", "unix-socket", ArgFilename, false, ""},
		"path-as-is":                 {"$N", "path-as-is", ArgBool, false, ""},
		"socks5-gssapi-service":      {"$O", "socks5-gssapi-service", ArgString, false, ""},
		"proxy-service-name":         {"$O", "proxy-service-name", ArgString, false, ""},
		"service-name":               {"$P", "service-name", ArgString, false, ""},
		"proto-default":              {"$Q", "proto-default", ArgString, false, ""},
		"expect100-timeout":          {"$R", "expect100-timeout", ArgString, false, ""},
		"tftp-no-options":            {"$S", "tftp-no-options", ArgBool, false, ""},
		"connect-to":                 {"$U", "connect-to", ArgString, false, ""},
		"abstract-unix-socket":       {"$W", "abstract-unix-socket", ArgFilename, false, ""},
		"tls-max":                    {"$X", "tls-max", ArgString, false, ""},
		"suppress-connect-headers":   {"$Y", "suppress-connect-headers", ArgBool, false, ""},
		"compressed-ssh":             {"$Z", "compressed-ssh", ArgBool, false, ""},
		"happy-eyeballs-timeout-ms":  {"$~", "happy-eyeballs-timeout-ms", ArgString, false, ""},
		"retry-all-errors":           {"$!", "retry-all-errors", ArgBool, false, ""},
		"http1.0":                    {"0", "http1.0", ArgNone, false, ""},
		"http1.1":                    {"01", "http1.1", ArgNone, false, ""},
		"http2":                      {"02", "http2", ArgNone, false, ""},
		"http2-prior-knowledge":      {"03", "http2-prior-knowledge", ArgNone, false, ""},
		"http3":                      {"04", "http3", ArgNone, false, ""},
		"http3-only":                 {"05", "http3-only", ArgNone, false, ""},
		"http0.9":                    {"09", "http0.9", ArgBool, false, ""},
		"tlsv1":                      {"1", "tlsv1", ArgNone, false, ""},
		"tlsv1.0":                    {"10", "tlsv1.0", ArgNone, false, ""},
		"tlsv1.1":                    {"11", "tlsv1.1", ArgNone, false, ""},
		"tlsv1.2":                    {"12", "tlsv1.2", ArgNone, false, ""},
		"tlsv1.3":                    {"13", "tlsv1.3", ArgNone, false, ""},
		"tls13-ciphers":              {"1A", "tls13-ciphers", ArgString, false, ""},
		"proxy-tls13-ciphers":        {"1B", "proxy-tls13-ciphers", ArgString, false, ""},
		"sslv2":                      {"2", "sslv2", ArgNone, false, ""},
		"sslv3":                      {"3", "sslv3", ArgNone, false, ""},
		"ipv4":                       {"4", "ipv4", ArgNone, false, ""},
		"ipv6":                       {"6", "ipv6", ArgNone, false, ""},
		"append":                     {"a", "append", ArgBool, false, ""},
		"user-agent":                 {"A", "user-agent", ArgString, false, ""},
		"cookie":                     {"b", "cookie", ArgString, false, ""},
		"alt-svc":                    {"ba", "alt-svc", ArgString, false, ""},
		"hsts":                       {"bb", "hsts", ArgString, false, ""},
		"use-ascii":                  {"B", "use-ascii", ArgBool, false, ""},
		"cookie-jar":                 {"c", "cookie-jar", ArgString, false, ""},
		"continue-at":                {"C", "continue-at", ArgString, false, ""},
		"data":                       {"d", "data", ArgString, false, ""},
		"data-raw":                   {"dr", "data-raw", ArgString, false, ""},
		"data-ascii":                 {"da", "data-ascii", ArgString, false, ""},
		"data-binary":                {"db", "data-binary", ArgString, false, ""},
		"data-urlencode":             {"de", "data-urlencode", ArgString, false, ""},
		"json":                       {"df", "json", ArgString, false, ""},
		"url-query":                  {"dg", "url-query", ArgString, false, ""},
		"dump-header":                {"D", "dump-header", ArgFilename, false, ""},
		"referer":                    {"e", "referer", ArgString, false, ""},
		"cert":                       {"E", "cert", ArgFilename, false, ""},
		"cacert":                     {"Ea", "cacert", ArgFilename, false, ""},
		"cert-type":                  {"Eb", "cert-type", ArgString, false, ""},
		"key":                        {"Ec", "key", ArgFilename, false, ""},
		"key-type":                   {"Ed", "key-type", ArgString, false, ""},
		"pass":                       {"Ee", "pass", ArgString, false, ""},
		"engine":                     {"Ef", "engine", ArgString, false, ""},
		"capath":                     {"Eg", "capath", ArgFilename, false, ""},
		"pubkey":                     {"Eh", "pubkey", ArgString, false, ""},
		"hostpubmd5":                 {"Ei", "hostpubmd5", ArgString, false, ""},
		"hostpubsha256":              {"EF", "hostpubsha256", ArgString, false, ""},
		"crlfile":                    {"Ej", "crlfile", ArgFilename, false, ""},
		"tlsuser":                    {"Ek", "tlsuser", ArgString, false, ""},
		"tlspassword":                {"El", "tlspassword", ArgString, false, ""},
		"tlsauthtype":                {"Em", "tlsauthtype", ArgString, false, ""},
		"ssl-allow-beast":            {"En", "ssl-allow-beast", ArgBool, false, ""},
		"ssl-auto-client-cert":       {"Eo", "ssl-auto-client-cert", ArgBool, false, ""},
		"proxy-ssl-auto-client-cert": {"EO", "proxy-ssl-auto-client-cert", ArgBool, false, ""},
		"pinnedpubkey":               {"Ep", "pinnedpubkey", ArgString, false, ""},
		"proxy-pinnedpubkey":         {"EP", "proxy-pinnedpubkey", ArgString, false, ""},
		"cert-status":                {"Eq", "cert-status", ArgBool, false, ""},
		"doh-cert-status":            {"EQ", "doh-cert-status", ArgBool, false, ""},
		"false-start":                {"Er", "false-start", ArgBool, false, ""},
		"ssl-no-revoke":              {"Es", "ssl-no-revoke", ArgBool, false, ""},
		"ssl-revoke-best-effort":     {"ES", "ssl-revoke-best-effort", ArgBool, false, ""},
		"tcp-fastopen":               {"Et", "tcp-fastopen", ArgBool, false, ""},
		"proxy-tlsuser":              {"Eu", "proxy-tlsuser", ArgString, false, ""},
		"proxy-tlspassword":          {"Ev", "proxy-tlspassword", ArgString, false, ""},
		"proxy-tlsauthtype":          {"Ew", "proxy-tlsauthtype", ArgString, false, ""},
		"proxy-cert":                 {"Ex", "proxy-cert", ArgFilename, false, ""},
		"proxy-cert-type":            {"Ey", "proxy-cert-type", ArgString, false, ""},
		"proxy-key":                  {"Ez", "proxy-key", ArgFilename, false, ""},
		"proxy-key-type":             {"E0", "proxy-key-type", ArgString, false, ""},
		"proxy-pass":                 {"E1", "proxy-pass", ArgString, false, ""},
		"proxy-ciphers":              {"E2", "proxy-ciphers", ArgString, false, ""},
		"proxy-crlfile":              {"E3", "proxy-crlfile", ArgFilename, false, ""},
		"proxy-ssl-allow-beast":      {"E4", "proxy-ssl-allow-beast", ArgBool, false, ""},
		"login-options":              {"E5", "login-options", ArgString, false, ""},
		"proxy-cacert":               {"E6", "proxy-cacert", ArgFilename, false, ""},
		"proxy-capath":               {"E7", "proxy-capath", ArgFilename, false, ""},
		"proxy-insecure":             {"E8", "proxy-insecure", ArgBool, false, ""},
		"proxy-tlsv1":                {"E9", "proxy-tlsv1", ArgNone, false, ""},
		"socks5-basic":               {"EA", "socks5-basic", ArgBool, false, ""},
		"socks5-gssapi":              {"EB", "socks5-gssapi", ArgBool, false, ""},
		"etag-save":                  {"EC", "etag-save", ArgFilename, false, ""},
		"etag-compare":               {"ED", "etag-compare", ArgFilename, false, ""},
		"curves":                     {"EE", "curves", ArgString, false, ""},
		"fail":                       {"f", "fail", ArgBool, false, ""},
		"fail-early":                 {"fa", "fail-early", ArgBool, false, ""},
		"styled-output":              {"fb", "styled-output", ArgBool, false, ""},
		"mail-rcpt-allowfails":       {"fc", "mail-rcpt-allowfails", ArgBool, false, ""},
		"fail-with-body":             {"fd", "fail-with-body", ArgBool, false, ""},
		"remove-on-error":            {"fe", "remove-on-error", ArgBool, false, ""},
		"form":                       {"F", "form", ArgString, false, ""},
		"form-string":                {"Fs", "form-string", ArgString, false, ""},
		"globoff":                    {"g", "globoff", ArgBool, false, ""},
		"get":                        {"G", "get", ArgBool, false, ""},
		"request-target":             {"Ga", "request-target", ArgString, false, ""},
		"help":                       {"h", "help", ArgBool, false, ""},
		"header":                     {"H", "header", ArgString, false, ""},
		"proxy-header":               {"Hp", "proxy-header", ArgString, false, ""},
		"include":                    {"i", "include", ArgBool, false, ""},
		"head":                       {"I", "head", ArgBool, false, ""},
		"junk-session-cookies":       {"j", "junk-session-cookies", ArgBool, false, ""},
		"remote-header-name":         {"J", "remote-header-name", ArgBool, false, ""},
		"insecure":                   {"k", "insecure", ArgBool, false, ""},
		"doh-insecure":               {"kd", "doh-insecure", ArgBool, false, ""},
		"config":                     {"K", "config", ArgFilename, false, ""},
		"list-only":                  {"l", "list-only", ArgBool, false, ""},
		"location":                   {"L", "location", ArgBool, false, ""},
		"location-trusted":           {"Lt", "location-trusted", ArgBool, false, ""},
		"max-time":                   {"m", "max-time", ArgString, false, ""},
		"manual":                     {"M", "manual", ArgBool, false, ""},
		"netrc":                      {"n", "netrc", ArgBool, false, ""},
		"netrc-optional":             {"no", "netrc-optional", ArgBool, false, ""},
		"netrc-file":                 {"ne", "netrc-file", ArgFilename, false, ""},
		"buffer":                     {"N", "buffer", ArgBool, false, ""},
		"output":                     {"o", "output", ArgFilename, false, ""},
		"remote-name":                {"O", "remote-name", ArgBool, false, ""},
		"remote-name-all":            {"Oa", "remote-name-all", ArgBool, false, ""},
		"output-dir":                 {"Ob", "output-dir", ArgString, false, ""},
		"clobber":                    {"Oc", "clobber", ArgBool, false, ""},
		"proxytunnel":                {"p", "proxytunnel", ArgBool, false, ""},
		"ftp-port":                   {"P", "ftp-port", ArgString, false, ""},
		"disable":                    {"q", "disable", ArgBool, false, ""},
		"quote":                      {"Q", "quote", ArgString, false, ""},
		"range":                      {"r", "range", ArgString, false, ""},
		"remote-time":                {"R", "remote-time", ArgBool, false, ""},
		"silent":                     {"s", "silent", ArgBool, false, ""},
		"show-error":                 {"S", "show-error", ArgBool, false, ""},
		"telnet-option":              {"desc", "telnet-option", ArgString, false, ""},
		"upload-file":                {"T", "upload-file", ArgFilename, false, ""},
		"user":                       {"u", "user", ArgString, false, ""},
		"proxy-user":                 {"U", "proxy-user", ArgString, false, ""},
		"verbose":                    {"v", "verbose", ArgBool, false, ""},
		"version":                    {"V", "version", ArgBool, false, ""},
		"write-out":                  {"w", "write-out", ArgString, false, ""},
		"proxy":                      {"x", "proxy", ArgString, false, ""},
		"preproxy":                   {"xa", "preproxy", ArgString, false, ""},
		"request":                    {"X", "request", ArgString, false, ""},
		"speed-limit":                {"Y", "speed-limit", ArgString, false, ""},
		"speed-time":                 {"y", "speed-time", ArgString, false, ""},
		"time-cond":                  {"z", "time-cond", ArgString, false, ""},
		"parallel":                   {"Z", "parallel", ArgBool, false, ""},
		"parallel-max":               {"Zb", "parallel-max", ArgString, false, ""},
		"parallel-immediate":         {"Zc", "parallel-immediate", ArgBool, false, ""},
		"progress-bar":               {"#", "progress-bar", ArgBool, false, ""},
		"progress-meter":             {"#m", "progress-meter", ArgBool, false, ""},
		"next":                       {":", "next", ArgNone, false, ""},
	}
)

func setParameter(flag string, argv []string) ([]string, error) {
	if strings.HasPrefix(flag, "--") || !strings.HasPrefix(flag, "-") {
		/* this should be a long name */
		word := strings.TrimPrefix(flag, "--")
		if strings.HasPrefix(word, "no-") {
			/* disable this option but ignore the "no-" part when looking for it */
			word = strings.TrimPrefix(word, "no-")
			param, ok := longNameParameter.Get(word)
			if ok {
				param.boolValue = false
			} else {
				log.Println("[warning] unknown parameter: --" + word)
			}
			return argv, nil
		}
		param, ok := longNameParameter.Get(word)
		if !ok {
			log.Println("[warning] unknown parameter: --" + word)
			return argv, nil
		}
		return param.SetVal(argv)
	} else {
		/* prefixed with one dash */
		var word string
		if len(flag) < 2 {
			return argv, errors.New("invalid parameter: " + flag)
		} else if len(flag) > 2 {
			/* this is the actual extra parameter */
			argv = append([]string{"-" + flag[2:]}, argv...)
		}
		word = string(strings.TrimPrefix(flag, "-")[0])
		var err error
		for _, param := range *longNameParameter {
			if param.letter == word {
				argv, err = param.SetVal(argv)
				break
			}
		}
		return argv, err
	}
}

func ParseArgs(argv []string) (CurlParameter, error) {
	var err error
	for len(argv) > 0 {
		if strings.HasPrefix(argv[0], "-") {
			if len(argv[0]) == 1 {
				argv, err = setParameter(argv[0], nil)
			} else {
				argv, err = setParameter(argv[0], argv[1:])
			}
			if err != nil {
				return nil, err
			}
		} else {
			/* Just add the URL please */
			argv, err = setParameter("--url", argv)
			if err != nil {
				return nil, err
			}
		}
	}
	return *longNameParameter, err
}
