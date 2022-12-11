package gosca

func CreateScaffolding(n string) *Scaffolding {
	return &Scaffolding{
		Scaffolding: Dir{
			Name: n,
		},
	}
}

func CreateGolangScaffolding(n string) *Scaffolding {
	return &Scaffolding{
		Scaffolding: Dir{
			Name:  n,
			Files: []string{"main.go", "go.mod", "readme.md"},
			Dirs: []Dir{
				{
					Name: "bin",
				},
				{
					Name: "cmd",
				},
				{
					Name: "pkg",
				},
				{
					Name: "test",
				},
			},
		},
	}
}

func CreatePythonScaffolding(n string) *Scaffolding {
	return &Scaffolding{
		Scaffolding: Dir{
			Name:  n,
			Files: []string{"dependency.txt", "readme.md"},
			Dirs: []Dir{
				{
					Name: "src",
					Dirs: []Dir{
						{
							Name:  n,
							Files: []string{"__init__.py"},
						},
					},
				},
				{
					Name: "test",
				},
			},
		},
	}
}
