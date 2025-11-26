package constants

import "os"

const (
	ConfigFolderPathExtension                       = "info.richardjecooke.PeriodicTasks"
	ConfigPathExtensionAndFileName                  = ConfigFolderPathExtension + "/config.json"
	DefaultDataPathExtensionAndFileName             = ConfigFolderPathExtension + "/tasks.json"
	Permission_RWX_RX_RX                os.FileMode = 0755
	TaskFileVersion                                 = 1
)
