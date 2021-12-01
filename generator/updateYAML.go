package main

// UpdateYAML holds the input update data from an update.yaml file
type UpdateYAML struct {
	VersionCode int `yaml:"versionCode"`
	VersionName string `yaml:"versionName"`

	OSRequirement string `yaml:"osRequirement"`
	ProtocolRequirement string `yaml:"protocolRequirement"`

	URL string `yaml:"url"`
	ExternalDownload bool `yaml:"externalDownload"`
}