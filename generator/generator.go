package generator

import (
	"embed"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/swiftcarrot/queryx/inflect"
	"github.com/swiftcarrot/queryx/schema"
)

type Generator struct {
	Adapter   string
	template  *template.Template
	Templates map[string]*template.Template
}

// load template files from go embed
func (g *Generator) LoadTemplates(src embed.FS, adapter string) error {
	t := template.New("templates").Funcs(inflect.TemplateFunctions)

	if err := fs.WalkDir(src, "templates", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return fmt.Errorf("error walking directory %s: %w", path, err)
		}
		if d.IsDir() {
			return nil
		}

		templateName := strings.TrimPrefix(path, "templates")
		buf, err := src.ReadFile(path)
		if err != nil {
			return fmt.Errorf("error reading file '%s': %w", path, err)
		}
		tmpl := t.New(templateName)
		_, err = tmpl.Parse(string(buf))
		if err != nil {
			return fmt.Errorf("parsing template '%s': %w", path, err)
		}

		return nil
	}); err != nil {
		return err
	}

	g.template = t

	return nil
}

func (g *Generator) CreateFile(f string, tpl *template.Template, data interface{}) error {
	fmt.Println("Created", f)

	dir := filepath.Dir(f)
	if dir != "" {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}

	wr, err := os.Create(f)
	if err != nil {
		return err
	}

	if err := tpl.Execute(wr, data); err != nil {
		return err
	}

	return nil
}

func stringInSlice(str string, list []string) bool {
	for _, v := range list {
		if v == str {
			return true
		}
	}
	return false
}

func (g *Generator) Generate(schema *schema.Schema) error {
	database := schema.Databases[0]
	dir := database.Name
	created := []string{}
	err := checkRelationship(database)
	if err != nil {
		return err
	}

	for _, tpl := range g.template.Templates() {
		name := tpl.Name()
		if name == "templates" || !strings.Contains(name, ".") { // TODO: ignore templates defined in templates
			continue
		}

		if strings.Contains(name, "[model]") {
			for _, model := range database.Models {
				n := strings.ReplaceAll(name, "[model]", inflect.Snake(model.Name))
				f := path.Join(dir, strings.TrimSuffix(n, "tmpl"))
				created = append(created, f)

				data := map[string]interface{}{
					"packageName": dir,
					"client":      database,
					"model":       model,
				}
				if err := g.CreateFile(f, tpl, data); err != nil {
					return err
				}
			}
		} else {
			f := path.Join(dir, strings.TrimSuffix(name, "tmpl"))
			created = append(created, f)
			data := map[string]interface{}{
				"packageName": dir,
				"client":      database,
			}
			if err := g.CreateFile(f, tpl, data); err != nil {
				return err
			}
		}
	}

	deleted := []string{}
	files, err := readDir(dir)

	if err != nil {
		return err
	}
	for _, f := range files {
		if !stringInSlice(f, created) {
			deleted = append(deleted, f)
		}
	}
	for _, f := range deleted {
		os.Remove(f)
		fmt.Println("Deleted", f)
	}

	return nil
}

func readDir(dir string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && info.Name() == "migrations" {
			return filepath.SkipDir
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func checkRelationship(d *schema.Database) error {
	modeMap := make(map[string]struct{})
	for i := 0; i < len(d.Models); i++ {
		modeMap[d.Models[i].Name] = struct{}{}
	}
	for i := 0; i < len(d.Models); i++ {
		for j := 0; j < len(d.Models[i].HasOne); j++ {
			if len(d.Models[i].HasOne) > 0 {
				if _, ok := modeMap[inflect.Pascal(inflect.Singular(d.Models[i].HasOne[j].Name))]; !ok {
					return fmt.Errorf(fmt.Sprintf("the model of %s do not exist", d.Models[i].HasOne[j].Name))
				}
			}

		}
		for k := 0; k < len(d.Models[i].HasMany); k++ {
			if len(d.Models[i].HasMany) > 0 {
				if _, ok := modeMap[inflect.Pascal(inflect.Singular(d.Models[i].HasMany[k].Name))]; !ok {
					return fmt.Errorf(fmt.Sprintf("the model of %s do not exist", d.Models[i].HasMany[k].Name))
				}
			}

		}

		for h := 0; h < len(d.Models[i].BelongsTo); h++ {
			if len(d.Models[i].BelongsTo) > 0 {
				if _, ok := modeMap[inflect.Pascal(inflect.Singular(d.Models[i].BelongsTo[h].Name))]; !ok {
					return fmt.Errorf(fmt.Sprintf("the model of %s do not exist", d.Models[i].BelongsTo[h].Name))
				}
			}
		}
	}

	return nil
}
