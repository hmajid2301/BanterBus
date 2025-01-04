// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.793
package layouts

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import (
	"github.com/a-h/templ"
	templruntime "github.com/a-h/templ/runtime"

	"gitlab.com/hmajid2301/banterbus/internal/views/components"
)

func Base(languages map[string]string) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><meta name=\"apple-mobile-web-app-title\" content=\"Banter Bus\"><title>Banter Bus</title><link rel=\"stylesheet\" href=\"/static/css/styles.css\"><link rel=\"icon\" type=\"image/png\" href=\"/static/images/favicon-48x48.png\" sizes=\"48x48\"><link rel=\"icon\" type=\"image/svg+xml\" href=\"/static/images/favicon.svg\"><link rel=\"shortcut icon\" href=\"/static/images/favicon.ico\"><link rel=\"apple-touch-icon\" sizes=\"180x180\" href=\"/static/images/apple-touch-icon.png\"><link rel=\"manifest\" href=\"/static/site.webmanifest\"></head><body>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Toast().Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-full min-h-screen text-lg bg-center bg-no-repeat bg-cover bg-mantle bg-background\" hx-ext=\"ws\" ws-connect=\"/ws\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = components.Header(languages).Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section class=\"flex flex-col justify-center items-center min-h-screen\"><div class=\"py-10 px-20 max-w-3xl rounded-xl bg-surface0\"><div class=\"flex flex-col justify-center items-center\"><div class=\"flex flex-col items-center space-y-10\"><h1 class=\"text-8xl tracking-tighter text-center text-text font-header text-shadow-custom\">Banter Bus</h1><div id=\"page\" class=\"mt-5 w-full font-main\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div></div></div></section><div id=\"error\"></div></div></body><script src=\"/static/js/htmx.min.js\"></script><script src=\"/static/js/htmx.ws.js\"></script><script src=\"/static/js/alpine.min.js\"></script><script defer>\n            htmx.on(\"htmx:wsBeforeMessage\", (evt) => {\n                try {\n                    const {message, type} = JSON.parse(event.detail.message);\n                    window.toast(message, type);\n                } catch (_) {}\n            });\n        </script></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
