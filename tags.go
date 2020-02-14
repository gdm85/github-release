package main

import (
	"fmt"
)

const (
	TAGS_URI = "/repos/%s/%s/tags"
)

type Tag struct {
	Name       string `json:"name"`
	Commit     Commit `json:"commit"`
	ZipBallUrl string `json:"zipball_url"`
	TarBallUrl string `json:"tarball_url"`
}

func (t *Tag) String() string {
	return t.Name + " (commit: " + t.Commit.Url + ")"
}

// Tags gets the tags associated with a repository.
func (cmd *CommandParams) Tags() ([]Tag, error) {
	var tags []Tag
	return tags, cmd.Get(fmt.Sprintf(TAGS_URI, cmd.User, cmd.Repo), &tags)
}
