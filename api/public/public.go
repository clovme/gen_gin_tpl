package public

import (
	"embed"
)

//go:embed all:email
var EmailFS embed.FS

//go:embed web/templates/**/*.html
var TemplatesFS embed.FS

//go:embed web/static
var StaticFS embed.FS

//go:embed all:initweb
var InitWebFS embed.FS

//go:embed favicon.ico
var Favicon embed.FS

//go:embed rsa/public.pem
var PublicPEM []byte

//go:embed rsa/private.pem
var PrivatePEM []byte
