package dsl

//func execTemplate(tmpl *template.Template, val interface{}) string {
//	buf := bytes.Buffer{}
//	err := tmpl.Execute(&buf, val)
//	if err != nil {
//		panic(err)
//	}
//	return buf.String()
//}
//
//var moduleTmpl = template.Must(template.New("module").Parse(`
//{{range .Decls}}
//  {{.}}
//{{end}}
//`))
//
//func (m Module) String() string {
//	return execTemplate(moduleTmpl, m)
//}
//
//var declTmpl = template.Must(template.New("module").Parse(`
//{{range .Decls}}
//  {{.}}
//{{end}}
//`))
//
//func (m Decl) String() string {
//	return execTemplate(moduleTmpl, m)
//}
