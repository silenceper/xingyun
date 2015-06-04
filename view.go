package xingyun

import(
    "html/template"
    "strings"
    "path"
    "bytes"
)

type View struct{
    Dir string
    FuncMap template.FuncMap
}

func (v *View) getTemplateInstance(tpl []string) (*template.Template, error) {
    //key := strings.Join(tpl, "-")
    var (
        t    *template.Template
        e    error
        file []string = make([]string, len(tpl))
    )
    for i, tp := range tpl {
        file[i] = path.Join(v.Dir, tp)
    }
    t = template.New(path.Base(tpl[0]))
    t.Funcs(v.FuncMap)
    t, e = t.ParseFiles(file...)
    if e != nil {
        return nil, e
    }
    return t, nil

}

func (v *View) Render(tpl string, data map[string]interface{}) ([]byte, error) {
    t, e := v.getTemplateInstance(strings.Split(tpl, ","))
    if e != nil {
        return nil, e
    }
    var buf bytes.Buffer
    e = t.Execute(&buf, data)
    if e != nil {
        return nil, e
    }
    return buf.Bytes(), nil
}

func NewView(dir string) *View {
    v := new(View)
    v.Dir = dir
    v.FuncMap = make(template.FuncMap)
        
    //default function
    v.FuncMap["Html"] = func(str string) template.HTML {
        return template.HTML(str)
    }
    return v
}
