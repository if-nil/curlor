package main

import (
	"errors"
	"log"
	"net/url"
	"strings"
)

type Parameter map[string]*Param

func (p *Parameter) Get(key string) (*Param, bool) {
	param, ok := (*p)[key]
	return param, ok
}

func (p *Parameter) MustGet(key string) *Param {
	return (*p)[key]
}

func (p *Parameter) Protocol() string {
	param, _ := p.Get("url")
	u, err := url.Parse(param.StringVal())
	if err != nil {
		return ""
	}
	return u.Scheme
}

type ParamType int

const (
	argNone ParamType = iota
	argBool
	argString
	argFilename
)

type Param struct {
	letter    string
	lname     string
	desc      ParamType
	valBool   bool
	valString string
}

func (p *Param) SetVal(argv []string) ([]string, error) {
	switch p.desc {
	case argBool, argNone:
		p.valBool = true
	case argString, argFilename:
		if len(argv) == 0 {
			return argv, errors.New("missing argument")
		}
		p.valString = argv[0]
		argv = argv[1:]
	}
	return argv, nil
}

func (p *Param) StringVal() string {
	return p.valString
}

func (p *Param) BoolVal() bool {
	return p.valBool
}

var (
	longNameParameter = &Parameter{
		"url":                        {"*@", "url", argString, false, ""},
		"dns-ipv4-addr":              {"*4", "dns-ipv4-addr", argString, false, ""},
		"dns-ipv6-addr":              {"*6", "dns-ipv6-addr", argString, false, ""},
		"random-file":                {"*a", "random-file", argFilename, false, ""},
		"egd-file":                   {"*b", "egd-file", argString, false, ""},
		"oauth2-bearer":              {"*B", "oauth2-bearer", argString, false, ""},
		"connect-timeout":            {"*c", "connect-timeout", argString, false, ""},
		"doh-url":                    {"*C", "doh-url", argString, false, ""},
		"ciphers":                    {"*d", "ciphers", argString, false, ""},
		"dns-interface":              {"*D", "dns-interface", argString, false, ""},
		"disable-epsv":               {"*e", "disable-epsv", argBool, false, ""},
		"disallow-username-in-url":   {"*f", "disallow-username-in-url", argBool, false, ""},
		"epsv":                       {"*E", "epsv", argBool, false, ""},
		"dns-servers":                {"*F", "dns-servers", argString, false, ""},
		"trace":                      {"*g", "trace", argFilename, false, ""},
		"npn":                        {"*G", "npn", argBool, false, ""},
		"trace-ascii":                {"*h", "trace-ascii", argFilename, false, ""},
		"alpn":                       {"*H", "alpn", argBool, false, ""},
		"limit-rate":                 {"*i", "limit-rate", argString, false, ""},
		"rate":                       {"*I", "rate", argString, false, ""},
		"compressed":                 {"*j", "compressed", argBool, false, ""},
		"tr-encoding":                {"*J", "tr-encoding", argBool, false, ""},
		"digest":                     {"*k", "digest", argBool, false, ""},
		"negotiate":                  {"*l", "negotiate", argBool, false, ""},
		"ntlm":                       {"*m", "ntlm", argBool, false, ""},
		"ntlm-wb":                    {"*M", "ntlm-wb", argBool, false, ""},
		"basic":                      {"*n", "basic", argBool, false, ""},
		"anyauth":                    {"*o", "anyauth", argBool, false, ""},
		"wdebug":                     {"*p", "wdebug", argBool, false, ""},
		"ftp-create-dirs":            {"*q", "ftp-create-dirs", argBool, false, ""},
		"create-dirs":                {"*r", "create-dirs", argBool, false, ""},
		"create-file-mode":           {"*R", "create-file-mode", argString, false, ""},
		"max-redirs":                 {"*s", "max-redirs", argString, false, ""},
		"proxy-ntlm":                 {"*desc", "proxy-ntlm", argBool, false, ""},
		"crlf":                       {"*u", "crlf", argBool, false, ""},
		"stderr":                     {"*v", "stderr", argFilename, false, ""},
		"aws-sigv4":                  {"*V", "aws-sigv4", argString, false, ""},
		"interface":                  {"*w", "interface", argString, false, ""},
		"krb":                        {"*x", "krb", argString, false, ""},
		"krb4":                       {"*x", "krb4", argString, false, ""},
		"haproxy-protocol":           {"*X", "haproxy-protocol", argBool, false, ""},
		"max-filesize":               {"*y", "max-filesize", argString, false, ""},
		"disable-eprt":               {"*z", "disable-eprt", argBool, false, ""},
		"eprt":                       {"*Z", "eprt", argBool, false, ""},
		"xattr":                      {"*~", "xattr", argBool, false, ""},
		"ftp-ssl":                    {"$a", "ftp-ssl", argBool, false, ""},
		"ssl":                        {"$a", "ssl", argBool, false, ""},
		"ftp-pasv":                   {"$b", "ftp-pasv", argBool, false, ""},
		"socks5":                     {"$c", "socks5", argString, false, ""},
		"tcp-nodelay":                {"$d", "tcp-nodelay", argBool, false, ""},
		"proxy-digest":               {"$e", "proxy-digest", argBool, false, ""},
		"proxy-basic":                {"$f", "proxy-basic", argBool, false, ""},
		"retry":                      {"$g", "retry", argString, false, ""},
		"retry-connrefused":          {"$V", "retry-connrefused", argBool, false, ""},
		"retry-delay":                {"$h", "retry-delay", argString, false, ""},
		"retry-max-time":             {"$i", "retry-max-time", argString, false, ""},
		"proxy-negotiate":            {"$k", "proxy-negotiate", argBool, false, ""},
		"form-escape":                {"$l", "form-escape", argBool, false, ""},
		"ftp-account":                {"$m", "ftp-account", argString, false, ""},
		"proxy-anyauth":              {"$n", "proxy-anyauth", argBool, false, ""},
		"trace-time":                 {"$o", "trace-time", argBool, false, ""},
		"ignore-content-length":      {"$p", "ignore-content-length", argBool, false, ""},
		"ftp-skip-pasv-ip":           {"$q", "ftp-skip-pasv-ip", argBool, false, ""},
		"ftp-method":                 {"$r", "ftp-method", argString, false, ""},
		"local-port":                 {"$s", "local-port", argString, false, ""},
		"socks4":                     {"$desc", "socks4", argString, false, ""},
		"socks4a":                    {"$T", "socks4a", argString, false, ""},
		"ftp-alternative-to-user":    {"$u", "ftp-alternative-to-user", argString, false, ""},
		"ftp-ssl-reqd":               {"$v", "ftp-ssl-reqd", argBool, false, ""},
		"ssl-reqd":                   {"$v", "ssl-reqd", argBool, false, ""},
		"sessionid":                  {"$w", "sessionid", argBool, false, ""},
		"ftp-ssl-control":            {"$x", "ftp-ssl-control", argBool, false, ""},
		"ftp-ssl-ccc":                {"$y", "ftp-ssl-ccc", argBool, false, ""},
		"ftp-ssl-ccc-mode":           {"$j", "ftp-ssl-ccc-mode", argString, false, ""},
		"libcurl":                    {"$z", "libcurl", argString, false, ""},
		"raw":                        {"$#", "raw", argBool, false, ""},
		"post301":                    {"$0", "post301", argBool, false, ""},
		"keepalive":                  {"$1", "keepalive", argBool, false, ""},
		"socks5-hostname":            {"$2", "socks5-hostname", argString, false, ""},
		"keepalive-time":             {"$3", "keepalive-time", argString, false, ""},
		"post302":                    {"$4", "post302", argBool, false, ""},
		"noproxy":                    {"$5", "noproxy", argString, false, ""},
		"socks5-gssapi-nec":          {"$7", "socks5-gssapi-nec", argBool, false, ""},
		"proxy1.0":                   {"$8", "proxy1.0", argString, false, ""},
		"tftp-blksize":               {"$9", "tftp-blksize", argString, false, ""},
		"mail-from":                  {"$A", "mail-from", argString, false, ""},
		"mail-rcpt":                  {"$B", "mail-rcpt", argString, false, ""},
		"ftp-pret":                   {"$C", "ftp-pret", argBool, false, ""},
		"proto":                      {"$D", "proto", argString, false, ""},
		"proto-redir":                {"$E", "proto-redir", argString, false, ""},
		"resolve":                    {"$F", "resolve", argString, false, ""},
		"delegation":                 {"$G", "delegation", argString, false, ""},
		"mail-auth":                  {"$H", "mail-auth", argString, false, ""},
		"post303":                    {"$I", "post303", argBool, false, ""},
		"metalink":                   {"$J", "metalink", argBool, false, ""},
		"sasl-authzid":               {"$6", "sasl-authzid", argString, false, ""},
		"sasl-ir":                    {"$K", "sasl-ir", argBool, false, ""},
		"test-event":                 {"$L", "test-event", argBool, false, ""},
		"unix-socket":                {"$M", "unix-socket", argFilename, false, ""},
		"path-as-is":                 {"$N", "path-as-is", argBool, false, ""},
		"socks5-gssapi-service":      {"$O", "socks5-gssapi-service", argString, false, ""},
		"proxy-service-name":         {"$O", "proxy-service-name", argString, false, ""},
		"service-name":               {"$P", "service-name", argString, false, ""},
		"proto-default":              {"$Q", "proto-default", argString, false, ""},
		"expect100-timeout":          {"$R", "expect100-timeout", argString, false, ""},
		"tftp-no-options":            {"$S", "tftp-no-options", argBool, false, ""},
		"connect-to":                 {"$U", "connect-to", argString, false, ""},
		"abstract-unix-socket":       {"$W", "abstract-unix-socket", argFilename, false, ""},
		"tls-max":                    {"$X", "tls-max", argString, false, ""},
		"suppress-connect-headers":   {"$Y", "suppress-connect-headers", argBool, false, ""},
		"compressed-ssh":             {"$Z", "compressed-ssh", argBool, false, ""},
		"happy-eyeballs-timeout-ms":  {"$~", "happy-eyeballs-timeout-ms", argString, false, ""},
		"retry-all-errors":           {"$!", "retry-all-errors", argBool, false, ""},
		"http1.0":                    {"0", "http1.0", argNone, false, ""},
		"http1.1":                    {"01", "http1.1", argNone, false, ""},
		"http2":                      {"02", "http2", argNone, false, ""},
		"http2-prior-knowledge":      {"03", "http2-prior-knowledge", argNone, false, ""},
		"http3":                      {"04", "http3", argNone, false, ""},
		"http3-only":                 {"05", "http3-only", argNone, false, ""},
		"http0.9":                    {"09", "http0.9", argBool, false, ""},
		"tlsv1":                      {"1", "tlsv1", argNone, false, ""},
		"tlsv1.0":                    {"10", "tlsv1.0", argNone, false, ""},
		"tlsv1.1":                    {"11", "tlsv1.1", argNone, false, ""},
		"tlsv1.2":                    {"12", "tlsv1.2", argNone, false, ""},
		"tlsv1.3":                    {"13", "tlsv1.3", argNone, false, ""},
		"tls13-ciphers":              {"1A", "tls13-ciphers", argString, false, ""},
		"proxy-tls13-ciphers":        {"1B", "proxy-tls13-ciphers", argString, false, ""},
		"sslv2":                      {"2", "sslv2", argNone, false, ""},
		"sslv3":                      {"3", "sslv3", argNone, false, ""},
		"ipv4":                       {"4", "ipv4", argNone, false, ""},
		"ipv6":                       {"6", "ipv6", argNone, false, ""},
		"append":                     {"a", "append", argBool, false, ""},
		"user-agent":                 {"A", "user-agent", argString, false, ""},
		"cookie":                     {"b", "cookie", argString, false, ""},
		"alt-svc":                    {"ba", "alt-svc", argString, false, ""},
		"hsts":                       {"bb", "hsts", argString, false, ""},
		"use-ascii":                  {"B", "use-ascii", argBool, false, ""},
		"cookie-jar":                 {"c", "cookie-jar", argString, false, ""},
		"continue-at":                {"C", "continue-at", argString, false, ""},
		"data":                       {"d", "data", argString, false, ""},
		"data-raw":                   {"dr", "data-raw", argString, false, ""},
		"data-ascii":                 {"da", "data-ascii", argString, false, ""},
		"data-binary":                {"db", "data-binary", argString, false, ""},
		"data-urlencode":             {"de", "data-urlencode", argString, false, ""},
		"json":                       {"df", "json", argString, false, ""},
		"url-query":                  {"dg", "url-query", argString, false, ""},
		"dump-header":                {"D", "dump-header", argFilename, false, ""},
		"referer":                    {"e", "referer", argString, false, ""},
		"cert":                       {"E", "cert", argFilename, false, ""},
		"cacert":                     {"Ea", "cacert", argFilename, false, ""},
		"cert-type":                  {"Eb", "cert-type", argString, false, ""},
		"key":                        {"Ec", "key", argFilename, false, ""},
		"key-type":                   {"Ed", "key-type", argString, false, ""},
		"pass":                       {"Ee", "pass", argString, false, ""},
		"engine":                     {"Ef", "engine", argString, false, ""},
		"capath":                     {"Eg", "capath", argFilename, false, ""},
		"pubkey":                     {"Eh", "pubkey", argString, false, ""},
		"hostpubmd5":                 {"Ei", "hostpubmd5", argString, false, ""},
		"hostpubsha256":              {"EF", "hostpubsha256", argString, false, ""},
		"crlfile":                    {"Ej", "crlfile", argFilename, false, ""},
		"tlsuser":                    {"Ek", "tlsuser", argString, false, ""},
		"tlspassword":                {"El", "tlspassword", argString, false, ""},
		"tlsauthtype":                {"Em", "tlsauthtype", argString, false, ""},
		"ssl-allow-beast":            {"En", "ssl-allow-beast", argBool, false, ""},
		"ssl-auto-client-cert":       {"Eo", "ssl-auto-client-cert", argBool, false, ""},
		"proxy-ssl-auto-client-cert": {"EO", "proxy-ssl-auto-client-cert", argBool, false, ""},
		"pinnedpubkey":               {"Ep", "pinnedpubkey", argString, false, ""},
		"proxy-pinnedpubkey":         {"EP", "proxy-pinnedpubkey", argString, false, ""},
		"cert-status":                {"Eq", "cert-status", argBool, false, ""},
		"doh-cert-status":            {"EQ", "doh-cert-status", argBool, false, ""},
		"false-start":                {"Er", "false-start", argBool, false, ""},
		"ssl-no-revoke":              {"Es", "ssl-no-revoke", argBool, false, ""},
		"ssl-revoke-best-effort":     {"ES", "ssl-revoke-best-effort", argBool, false, ""},
		"tcp-fastopen":               {"Et", "tcp-fastopen", argBool, false, ""},
		"proxy-tlsuser":              {"Eu", "proxy-tlsuser", argString, false, ""},
		"proxy-tlspassword":          {"Ev", "proxy-tlspassword", argString, false, ""},
		"proxy-tlsauthtype":          {"Ew", "proxy-tlsauthtype", argString, false, ""},
		"proxy-cert":                 {"Ex", "proxy-cert", argFilename, false, ""},
		"proxy-cert-type":            {"Ey", "proxy-cert-type", argString, false, ""},
		"proxy-key":                  {"Ez", "proxy-key", argFilename, false, ""},
		"proxy-key-type":             {"E0", "proxy-key-type", argString, false, ""},
		"proxy-pass":                 {"E1", "proxy-pass", argString, false, ""},
		"proxy-ciphers":              {"E2", "proxy-ciphers", argString, false, ""},
		"proxy-crlfile":              {"E3", "proxy-crlfile", argFilename, false, ""},
		"proxy-ssl-allow-beast":      {"E4", "proxy-ssl-allow-beast", argBool, false, ""},
		"login-options":              {"E5", "login-options", argString, false, ""},
		"proxy-cacert":               {"E6", "proxy-cacert", argFilename, false, ""},
		"proxy-capath":               {"E7", "proxy-capath", argFilename, false, ""},
		"proxy-insecure":             {"E8", "proxy-insecure", argBool, false, ""},
		"proxy-tlsv1":                {"E9", "proxy-tlsv1", argNone, false, ""},
		"socks5-basic":               {"EA", "socks5-basic", argBool, false, ""},
		"socks5-gssapi":              {"EB", "socks5-gssapi", argBool, false, ""},
		"etag-save":                  {"EC", "etag-save", argFilename, false, ""},
		"etag-compare":               {"ED", "etag-compare", argFilename, false, ""},
		"curves":                     {"EE", "curves", argString, false, ""},
		"fail":                       {"f", "fail", argBool, false, ""},
		"fail-early":                 {"fa", "fail-early", argBool, false, ""},
		"styled-output":              {"fb", "styled-output", argBool, false, ""},
		"mail-rcpt-allowfails":       {"fc", "mail-rcpt-allowfails", argBool, false, ""},
		"fail-with-body":             {"fd", "fail-with-body", argBool, false, ""},
		"remove-on-error":            {"fe", "remove-on-error", argBool, false, ""},
		"form":                       {"F", "form", argString, false, ""},
		"form-string":                {"Fs", "form-string", argString, false, ""},
		"globoff":                    {"g", "globoff", argBool, false, ""},
		"get":                        {"G", "get", argBool, false, ""},
		"request-target":             {"Ga", "request-target", argString, false, ""},
		"help":                       {"h", "help", argBool, false, ""},
		"header":                     {"H", "header", argString, false, ""},
		"proxy-header":               {"Hp", "proxy-header", argString, false, ""},
		"include":                    {"i", "include", argBool, false, ""},
		"head":                       {"I", "head", argBool, false, ""},
		"junk-session-cookies":       {"j", "junk-session-cookies", argBool, false, ""},
		"remote-header-name":         {"J", "remote-header-name", argBool, false, ""},
		"insecure":                   {"k", "insecure", argBool, false, ""},
		"doh-insecure":               {"kd", "doh-insecure", argBool, false, ""},
		"config":                     {"K", "config", argFilename, false, ""},
		"list-only":                  {"l", "list-only", argBool, false, ""},
		"location":                   {"L", "location", argBool, false, ""},
		"location-trusted":           {"Lt", "location-trusted", argBool, false, ""},
		"max-time":                   {"m", "max-time", argString, false, ""},
		"manual":                     {"M", "manual", argBool, false, ""},
		"netrc":                      {"n", "netrc", argBool, false, ""},
		"netrc-optional":             {"no", "netrc-optional", argBool, false, ""},
		"netrc-file":                 {"ne", "netrc-file", argFilename, false, ""},
		"buffer":                     {"N", "buffer", argBool, false, ""},
		"output":                     {"o", "output", argFilename, false, ""},
		"remote-name":                {"O", "remote-name", argBool, false, ""},
		"remote-name-all":            {"Oa", "remote-name-all", argBool, false, ""},
		"output-dir":                 {"Ob", "output-dir", argString, false, ""},
		"clobber":                    {"Oc", "clobber", argBool, false, ""},
		"proxytunnel":                {"p", "proxytunnel", argBool, false, ""},
		"ftp-port":                   {"P", "ftp-port", argString, false, ""},
		"disable":                    {"q", "disable", argBool, false, ""},
		"quote":                      {"Q", "quote", argString, false, ""},
		"range":                      {"r", "range", argString, false, ""},
		"remote-time":                {"R", "remote-time", argBool, false, ""},
		"silent":                     {"s", "silent", argBool, false, ""},
		"show-error":                 {"S", "show-error", argBool, false, ""},
		"telnet-option":              {"desc", "telnet-option", argString, false, ""},
		"upload-file":                {"T", "upload-file", argFilename, false, ""},
		"user":                       {"u", "user", argString, false, ""},
		"proxy-user":                 {"U", "proxy-user", argString, false, ""},
		"verbose":                    {"v", "verbose", argBool, false, ""},
		"version":                    {"V", "version", argBool, false, ""},
		"write-out":                  {"w", "write-out", argString, false, ""},
		"proxy":                      {"x", "proxy", argString, false, ""},
		"preproxy":                   {"xa", "preproxy", argString, false, ""},
		"request":                    {"X", "request", argString, false, ""},
		"speed-limit":                {"Y", "speed-limit", argString, false, ""},
		"speed-time":                 {"y", "speed-time", argString, false, ""},
		"time-cond":                  {"z", "time-cond", argString, false, ""},
		"parallel":                   {"Z", "parallel", argBool, false, ""},
		"parallel-max":               {"Zb", "parallel-max", argString, false, ""},
		"parallel-immediate":         {"Zc", "parallel-immediate", argBool, false, ""},
		"progress-bar":               {"#", "progress-bar", argBool, false, ""},
		"progress-meter":             {"#m", "progress-meter", argBool, false, ""},
		"next":                       {":", "next", argNone, false, ""},
	}
)

func getParameter(flag string, argv []string) ([]string, error) {
	if strings.HasPrefix(flag, "--") || !strings.HasPrefix(flag, "-") {
		/* this should be a long name */
		word := strings.TrimPrefix(flag, "--")
		if strings.HasPrefix(word, "no-") {
			/* disable this option but ignore the "no-" part when looking for it */
			word = strings.TrimPrefix(word, "no-")
			param, ok := longNameParameter.Get(word)
			if ok {
				param.valBool = false
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
			argv = append([]string{flag[2:]}, argv...)
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

func ParseArgs(argv []string) (Parameter, error) {
	var err error
	for len(argv) > 0 {
		if strings.HasPrefix(argv[0], "-") {
			if len(argv[0]) == 1 {
				argv, err = getParameter(argv[0], nil)
			} else {
				argv, err = getParameter(argv[0], argv[1:])
			}
			if err != nil {
				return nil, err
			}
		} else {
			/* Just add the URL please */
			argv, err = getParameter("--url", argv)
			if err != nil {
				return nil, err
			}
		}
	}
	return *longNameParameter, err
}
