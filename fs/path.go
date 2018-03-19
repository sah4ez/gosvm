package fs

//ExistsGoProject check basePaht + project exists in $GOPATH/src/
func ExistsGoProject(basePath, project string) bool {
	return Exists(PathToProject(basePath, project))
}

func PathToProject(basePath, project string) string {
	return gopath + "/src/" + basePath + "/" + project
}
