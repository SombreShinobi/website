// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.648
package layout

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import "time"
import "fmt"

func Base(partial bool) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		if partial {
			templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html> <html lang=\"en\"><head><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><link href=\"/dist/tailwind.css\" rel=\"stylesheet\"><script src=\"https://unpkg.com/htmx.org@1.9.5\" integrity=\"sha384-xcuj3WpfgjlKF+FXhSQFQ0ZNr39ln+hwjN3npfM9VBnUskLolQAcN80McRIVOPuO\" crossorigin=\"anonymous\"></script><title></title></head><body class=\"min-h-screen bg-gradient-to-br from-slate-800 to-slate-950 px-12 pt-12 font-thin text-gray-50\"><div class=\"min-h-[calc(100vh-96px)] flex flex-col border-y border-gray-50\"><nav class=\"sticky top-0 backdrop-blur-sm\"><ul class=\"container flex items-center p-4 select-none\"><li hx-get=\"/home\" hx-push-url=\"true\" hx-target=\"#content\" hx-swap=\"innerHTML transition:true\" class=\"mx-1.5 sm:mx-6 border-transparent border-b hover:border-gray-50 cursor-pointer\">Home</li><li hx-get=\"/blog\" hx-push-url=\"true\" hx-target=\"#content\" hx-swap=\"innerHTML transition:true\" class=\"mx-1.5 sm:mx-6 border-transparent border-b hover:border-gray-50 cursor-pointer\">Blog</li><li hx-get=\"/about\" hx-push-url=\"true\" hx-target=\"#content\" hx-swap=\"innerHTML transition:true\" class=\"mx-1.5 sm:mx-6 border-transparent border-b hover:border-gray-50 cursor-pointer\">About</li></ul></nav><main class=\"flex self-center items-center flex-col mx-auto p-4 text-left\"><div id=\"content\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			templ_7745c5c3_Err = templ_7745c5c3_Var1.Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></main></div><span class=\"text-right block w-full py-3 pr-3\">&copy; Sibin Sedlan ")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var2 string
			templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(fmt.Sprint(time.Now().Year()))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `views/layout/base.templ`, Line: 64, Col: 103}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</span></body></html>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
