package config

// Directories to create during init
var Directories = []string{
	".exfs",
	".exfs/.index",
	"exfs",
	"_exfs",
	"exfs_views",
	"exfs_removed",
}

// CSVFile represents a CSV file with its schema
type CSVFile struct {
	Path   string
	Schema string
}

// CSVFiles to create during init
var CSVFiles = []CSVFile{
	{Path: ".exfs/files.csv", Schema: "id,hash,ext,spawn,raw_hash,created_at,updated_at\n"},
	{Path: ".exfs/views.csv", Schema: "id,name,created_at\n"},
	{Path: ".exfs/relations.csv", Schema: "id,type,name,view_id,parent_id,file_id\n"},
	{Path: ".exfs/config.csv", Schema: "key,value\n"},
}
