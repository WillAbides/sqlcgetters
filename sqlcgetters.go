// sqlcgengetters generates getters for files output by sqlc.
// example: sqlcgengetters $SQLC_OUT/*.sql.go $SQLC_OUT/models.go > $SQLC_OUT/getters.go

package sqlcgetters

import (
	"bytes"
	"embed"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io"
	"io/fs"
	"sort"
	"text/template"

	"golang.org/x/tools/imports"
)

//go:embed *.tmpl
var templates embed.FS

// GenerateGetters generates getters for the files in sources.
func GenerateGetters(w io.Writer, fsys fs.FS, sources ...string) (int64, error) {
	if len(sources) == 0 {
		return 0, fmt.Errorf("no sources provided")
	}
	data, err := buildTemplateData(fsys, sources)
	if err != nil {
		return 0, fmt.Errorf("could not parse go source: %v", err)
	}
	tmpl, err := template.New("getters.go.tmpl").ParseFS(templates, "*")
	if err != nil {
		return 0, err
	}
	var buffer bytes.Buffer
	err = tmpl.Execute(&buffer, data)
	if err != nil {
		return 0, err
	}
	src, err := format.Source(buffer.Bytes())
	if err != nil {
		return 0, err
	}
	src, err = imports.Process("", src, nil)
	if err != nil {
		return 0, err
	}
	return io.Copy(w, bytes.NewReader(src))
}

func buildTemplateData(fsys fs.FS, sourceFiles []string) (*tmplData, error) {
	var td tmplData
	fset := token.NewFileSet()
	for _, sourceFile := range sourceFiles {
		b, err := fs.ReadFile(fsys, sourceFile)
		if err != nil {
			return nil, err
		}
		file, err := parser.ParseFile(fset, sourceFile, b, 0)
		if err != nil {
			return nil, err
		}
		if td.PackageName != "" && td.PackageName != file.Name.Name {
			return nil, fmt.Errorf("package name mismatch: %s != %s", td.PackageName, file.Name.Name)
		}
		td.PackageName = file.Name.Name
		for _, decl := range file.Decls {
			err = setGetters(&td, fset, decl)
			if err != nil {
				return nil, err
			}
		}
		for _, imp := range file.Imports {
			impName := ""
			if imp.Name != nil {
				impName = imp.Name.Name
			}
			td.imports = append(td.imports, fmt.Sprintf("%s %s", impName, imp.Path.Value))
		}
	}
	return &td, nil
}

func setGetters(td *tmplData, fset *token.FileSet, decl ast.Decl) error {
	d, ok := decl.(*ast.GenDecl)
	if !ok {
		return nil
	}
	for _, spec := range d.Specs {
		var ts *ast.TypeSpec
		ts, ok = spec.(*ast.TypeSpec)
		if !ok {
			continue
		}
		if !ts.Name.IsExported() {
			continue
		}
		structName := ts.Name.Name
		var st *ast.StructType
		st, ok = ts.Type.(*ast.StructType)
		if !ok {
			continue
		}
		for _, field := range st.Fields.List {
			if len(field.Names) == 0 || !field.Names[0].IsExported() {
				continue
			}
			var typeBuf bytes.Buffer
			err := printer.Fprint(&typeBuf, fset, field.Type)
			if err != nil {
				return err
			}
			td.getters = append(td.getters, getter{
				Type:      typeBuf.String(),
				FieldName: field.Names[0].Name,
				Receiver:  structName,
			})
		}
	}
	return nil
}

type getter struct {
	FieldName  string
	MethodName string
	Receiver   string
	Type       string
}

type tmplData struct {
	PackageName string
	imports     []string
	getters     []getter
}

// Imports returns sorted and deduplicated imports
func (td *tmplData) Imports() []string {
	imps := make(map[string]struct{})
	for _, imp := range td.imports {
		imps[imp] = struct{}{}
	}
	var sorted []string
	for imp := range imps {
		sorted = append(sorted, imp)
	}
	sort.Strings(sorted)
	return sorted
}

func (td *tmplData) Getters() []getter {
	getterKey := func(r, m string) string {
		return fmt.Sprintf("%s.%s", r, m)
	}
	getters := make([]getter, len(td.getters))
	copy(getters, td.getters)
	takenNames := make(map[string]bool, len(getters))
	for _, g := range getters {
		takenNames[getterKey(g.Receiver, g.FieldName)] = true
	}
	for i := range getters {
		getters[i].MethodName = fmt.Sprintf("Get%s", getters[i].FieldName)
		for takenNames[getterKey(getters[i].Receiver, getters[i].MethodName)] {
			getters[i].MethodName += "_"
		}
		takenNames[getterKey(getters[i].Receiver, getters[i].MethodName)] = true
	}
	return getters
}
