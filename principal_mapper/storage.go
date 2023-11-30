package principal_mapper

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/sheikhrachel/PMapper-go/api_common/errUtils"
)

// GetDefaultGraphPath returns a path to a given account or organization by the provided string.
func GetDefaultGraphPath(accountOrOrg string) (string, error) {
	storageRoot, err := GetStorageRoot()
	if errUtils.HandleError(err) {
		return "", err
	}
	return filepath.Join(storageRoot, accountOrOrg), nil
}

type OperatingSystem string

const (
	Windows OperatingSystem = "windows"
	MacOS   OperatingSystem = "darwin"
	Linux   OperatingSystem = "linux"
	FreeBSD OperatingSystem = "freebsd"
	OpenBSD OperatingSystem = "openbsd"
)

// String returns the string representation of the OperatingSystem.
func (os OperatingSystem) String() string {
	return string(os)
}

/*
GetStorageRoot locates and returns a path to the storage root, depending on OS.
	If the path does not exist yet, it is created.
*/
func GetStorageRoot() (path string, err error) {
	pMapperEnvVar, envVarExists := os.LookupEnv("PMAPPER_STORAGE")
	if envVarExists {
		return pMapperEnvVar, nil
	}
	path, err = getStoragePath()
	if errUtils.HandleError(err) {
		return "", err
	}
	if _, err = os.Stat(path); os.IsNotExist(err) {
		if err = os.MkdirAll(path, 0700); errUtils.HandleError(err) {
			return "", err
		}
	}
	return path, nil
}

// getStoragePath returns the path to the storage root, depending on OS.
func getStoragePath() (path string, err error) {
	switch runtime.GOOS {
	case Windows.String():
		appDataDir, appDataExists := os.LookupEnv("APPDATA")
		if !appDataExists {
			return "", errUtils.ErrAppDataNotSet
		}
		path = filepath.Join(appDataDir, "principal_mapper")
	case MacOS.String():
		homeDir, _ := os.UserHomeDir()
		path = filepath.Join(homeDir, "Library", "Application Support", "com.sheikhrachel.principal_mapper")
	case Linux.String(), FreeBSD.String(), OpenBSD.String():
		appDataDir, appDataExists := os.LookupEnv("XDG_DATA_HOME")
		if !appDataExists {
			homeDir, _ := os.UserHomeDir()
			appDataDir = filepath.Join(homeDir, ".local", "share")
		}
		path = filepath.Join(appDataDir, "principal_mapper")
	}
	return path, nil
}
