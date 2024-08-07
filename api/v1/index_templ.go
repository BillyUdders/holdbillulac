// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.747
package v1

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func indexPage(idx IndexPage) templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!doctype html><html lang=\"en\"><head><!--  HTML5 responsive viewport and title  --><meta charset=\"UTF-8\"><meta name=\"viewport\" content=\"width=device-width, initial-scale=1.0\"><title>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 string
		templ_7745c5c3_Var2, templ_7745c5c3_Err = templ.JoinStringErrs(idx.Title)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `api/v1/index.templ`, Line: 10, Col: 26}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var2))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</title><!-- Scripts --><script src=\"https://unpkg.com/htmx.org@2.0.1\" integrity=\"sha384-QWGpdj554B4ETpJJC9z+ZHJcA/i59TyjxEPXiiUgN2WmTyV5OEZWCD6gQhgkdpB/\" crossorigin=\"anonymous\"></script><script src=\"https://unpkg.com/htmx-ext-json-enc@2.0.0/json-enc.js\"></script><script>\n            document.addEventListener(\"DOMContentLoaded\", function (){\n                const form = document.getElementById(\"player-form\");\n                const submitButton = document.getElementById(\"submit\");\n\n                function checkFormValidity() {\n                    submitButton.disabled = !form.checkValidity();\n                }\n\n                form.addEventListener(\"input\", checkFormValidity);\n                checkFormValidity();\n            });\n        </script><script>\n            /*to prevent Firefox FOUC, this must be here*/\n            let FF_FOUC_FIX;\n        </script><!-- Linked icons and stylesheets --><link rel=\"icon\" type=\"image/x-icon\" href=\"/static/favicon.ico\"><link rel=\"stylesheet\" href=\"/static/styles.css\"><link rel=\"stylesheet\" href=\"https://fonts.googleapis.com/css?family=Marcellus+SC\"><link rel=\"stylesheet\" href=\"https://fonts.googleapis.com/css?family=Roboto\"></head><body><header><div><h1 class=\"title\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var3 string
		templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(idx.Description)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `api/v1/index.templ`, Line: 45, Col: 47}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</h1><a class=\"subtitle\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(idx.Tagline)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `api/v1/index.templ`, Line: 46, Col: 45}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a></div><nav class=\"navbar\" hx-get=\"/nav\" hx-trigger=\"load\"><div class=\"dropdown\"><button class=\"dropbtn\">Dropdown 2 <i class=\"fa fa-caret-down\"></i></button><div class=\"dropdown-content\"><a href=\"#\">Link 4</a> <a href=\"#\">Link 5</a> <a href=\"#\">Link 6</a></div></div><div class=\"dropdown\"><button class=\"dropbtn\">Dropdown 3 <i class=\"fa fa-caret-down\"></i></button><div class=\"dropdown-content\"><a href=\"#\">Link 4</a> <a href=\"#\">Link 5</a> <a href=\"#\">Link 6</a></div></div><div class=\"dropdown\"><button class=\"dropbtn\">Dropdown 4</button><div class=\"dropdown-content\"><a href=\"#\">Link 4</a> <a href=\"#\">Link 5</a> <a href=\"#\">Link 6</a></div></div><div class=\"dropdown\"><button class=\"dropbtn\">Dropdown 5 <i class=\"fa fa-caret-down\"></i></button><div class=\"dropdown-content\"><a href=\"#\">Link 4</a> <a href=\"#\">Link 5</a> <a href=\"#\">Link 6</a></div></div></nav></header><form id=\"player-form\" class=\"player-form\" hx-post=\"/player\" hx-trigger=\"submit\" hx-target=\"#table-body\" hx-swap=\"beforeend\" hx-ext=\"json-enc\"><label for=\"name\">Name:</label> <input type=\"text\" id=\"name\" name=\"name\" required><br><br><label for=\"age\">Age:</label> <input type=\"number\" id=\"age\" name=\"age\" inputmode=\"numeric\" required><br><br><label for=\"mmr\">MMR:</label> <input type=\"number\" id=\"mmr\" name=\"mmr\" inputmode=\"numeric\" required><br><br><button type=\"submit\" id=\"submit\">Submit</button></form><table class=\"table\"><thead><tr><th>ID</th><th>Name</th><th>Age</th><th>MMR</th><th>Actions</th></tr></thead> <tbody id=\"table-body\" hx-get=\"/player\" hx-trigger=\"load\"></tbody></table></body></html>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}