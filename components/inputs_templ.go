// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

func inputGroup() templ.CSSClass {
	templ_7745c5c3_CSSBuilder := templruntime.GetBuilder()
	templ_7745c5c3_CSSBuilder.WriteString(`display:flex;`)
	templ_7745c5c3_CSSBuilder.WriteString(`justify-content:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`align-items:center;`)
	templ_7745c5c3_CSSBuilder.WriteString(`gap:10px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`width:100%;`)
	templ_7745c5c3_CSSID := templ.CSSID(`inputGroup`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func label() templ.CSSClass {
	templ_7745c5c3_CSSBuilder := templruntime.GetBuilder()
	templ_7745c5c3_CSSBuilder.WriteString(`font-size:16px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`text-align:right;`)
	templ_7745c5c3_CSSBuilder.WriteString(`width:10%;`)
	templ_7745c5c3_CSSID := templ.CSSID(`label`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

func input() templ.CSSClass {
	templ_7745c5c3_CSSBuilder := templruntime.GetBuilder()
	templ_7745c5c3_CSSBuilder.WriteString(`margin:0.5rem 0.5rem;`)
	templ_7745c5c3_CSSBuilder.WriteString(`padding:6px 8px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`font-size:16px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`width:20%;`)
	templ_7745c5c3_CSSBuilder.WriteString(`background-color:#ddd;`)
	templ_7745c5c3_CSSBuilder.WriteString(`border-radius:10px;`)
	templ_7745c5c3_CSSBuilder.WriteString(`border:none;`)
	templ_7745c5c3_CSSBuilder.WriteString(`box-shadow:none;`)
	templ_7745c5c3_CSSID := templ.CSSID(`input`, templ_7745c5c3_CSSBuilder.String())
	return templ.ComponentCSSClass{
		ID:    templ_7745c5c3_CSSID,
		Class: templ.SafeCSS(`.` + templ_7745c5c3_CSSID + `{` + templ_7745c5c3_CSSBuilder.String() + `}`),
	}
}

var _ = templruntime.GeneratedTemplate
