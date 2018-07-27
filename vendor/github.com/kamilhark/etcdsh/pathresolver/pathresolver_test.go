package pathresolver

import "testing"
import "github.com/kamilhark/goassert"

func TestAddToPathSingleDirectory(t *testing.T) {
	pathResolver := new(PathResolver)

	pathResolver.Add("foo")
	currentPath := pathResolver.CurrentPath()

	goassert.AssertThat(t, currentPath).IsEqual("/foo")
}

func TestAddToPathManyDirectoriesAsOnce(t *testing.T) {
	pathResolver := new(PathResolver)

	pathResolver.Add("foo/bar")
	currentPath := pathResolver.CurrentPath()

	goassert.AssertThat(t, currentPath).IsEqual("/foo/bar")
}

func TestResolveRelativePathWhenCurrentIsRoot(t *testing.T) {
	pathResolver := new(PathResolver)

	resolved := pathResolver.Resolve("foo/bar")

	goassert.AssertThat(t, resolved).IsEqual("/foo/bar")
}

func TestResolveRelativePathWhenCurrentIsNotRoot(t *testing.T) {
	pathResolver := new(PathResolver)
	pathResolver.path = append(pathResolver.path, "foo")

	resolved := pathResolver.Resolve("bar")

	goassert.AssertThat(t, resolved).IsEqual("/foo/bar")
}

func TestResolvePathStartingFormTwoDots(t *testing.T) {
	pathResolver := new(PathResolver)
	pathResolver.path = append(pathResolver.path, "foo")

	resolved := pathResolver.Resolve("..")

	goassert.AssertThat(t, resolved).IsEqual("/")
}

func TestResolvePathMostComplicetedCAse(t *testing.T) {
	pathResolver := new(PathResolver)
	pathResolver.path = append(pathResolver.path, "foo")
	pathResolver.path = append(pathResolver.path, "bar")
	pathResolver.path = append(pathResolver.path, "jar")

	resolved := pathResolver.Resolve("../../bar")

	goassert.AssertThat(t, resolved).IsEqual("/foo/bar")
}
