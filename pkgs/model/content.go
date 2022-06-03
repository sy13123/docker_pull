package model

type m1 struct {
	Config   string
	RepoTags []string
	Layers   []string
}

var content []m1

func Contentvar() []m1 {
	content = append(content, m1{
		Config:   "",
		RepoTags: []string{},
		Layers:   []string{},
	})
	return content
}