package pathresolver

import "strings"

type PathResolver struct {
	path []string
}

func (p *PathResolver) Resolve(subPath string) string {
	elements := strings.Split(subPath, "/")

	currentPath := new(PathResolver)
	currentPath.path = p.path
	for _, element := range elements {
		if element == "." {
			continue
		} else if element == ".." {
			currentPath.RemoveLast()
		} else {
			currentPath.Add(element)
		}
	}

	return currentPath.CurrentPath()
}

func (p *PathResolver) GoTo(path string) {
	if len(path) == 0 {
		p.Clear()
	} else {
		newPath := p.Resolve(strings.TrimPrefix(path, "/"))
		p.path = strings.Split(newPath, "/")
	}
}

func (p *PathResolver) Add(subPath string) {
	p.path = append(p.path, subPath)
}

func (p *PathResolver) Clear() {
	p.path = []string{}
}

func (p *PathResolver) RemoveLast() {
	if len(p.path) > 0 {
		p.path = p.path[:len(p.path) - 1]
	}
}

func (p *PathResolver) CurrentPath() string {
	if len(p.path) == 0 {
		return "/"
	}
	return normalize(strings.Join(p.path, "/"))
}

func normalize(path string) string {
	if strings.HasPrefix(path, "/") == false {
		path = "/" + path
	}
	if strings.HasSuffix(path, "/") {
		path = path[:len(path) - 1]
	}
	return path
}
