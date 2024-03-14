package config

type TagmanagerPrefixes struct {
	Client    string                    `yaml:"client"`
	Folder    string                    `yaml:"folder"`
	Tags      TagmangerTagPrefixes      `yaml:"tags"`
	Triggers  TagmangerTriggerPrefixes  `yaml:"triggers"`
	Variables TagmangerVariablePrefixes `yaml:"variables"`
}

func (p TagmanagerPrefixes) ClientName(name string) string {
	return p.Client + name
}

func (p TagmanagerPrefixes) FolderName(name string) string {
	return p.Folder + name
}
