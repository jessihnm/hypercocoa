package api

import (
	"net/http"
	"strings"

	C "code.cerinuts.io/uni/hypercocoagateway/api/commands"
)

const AppName, VersionMajor, VersionMinor, VersionPatch, VersionBuild string = "hypercocoagateway", "0", "0", "1", "d"
const FullVersion string = AppName + VersionMajor + "." + VersionMinor + "." + VersionPatch + VersionBuild

//Config

type Config struct {
	Port     string
	SslCert  string
	SslKey   string
	SslPort  string
	Loglevel int
}

//Comparser (short for command parser) handles the specific requests

type Comparser struct {
	Conf *Config
}

// Init the Command Parser. Nothing needs to be done here at the moment
func (comp *Comparser) Init() {
}

// GetVersion returns the gateway version
func (comp *Comparser) GetVersion() string {
	vString := FullVersion
	return vString
}

// GetAssets returns all assets
func (comp *Comparser) GetAssets(w http.ResponseWriter, r *http.Request) string {
	if r.Method == http.MethodGet {
		res, err := C.GetAllAssets()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return err.Error()
		}
		w.Header().Add("Content-Type", "application/json")
		return res
	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return ""
	}
}

// GetAsset returns a single asset
func (comp *Comparser) GetAsset(w http.ResponseWriter, r *http.Request) string {
	idSplit := strings.Split(r.RequestURI, "/")
	id := idSplit[len(idSplit)-1]
	j, err := C.GetAsset(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return ""
	}
	w.Header().Add("Content-Type", "application/json")
	return j
}

// AddAsset adds an asset to the chain
func (comp *Comparser) AddAsset(w http.ResponseWriter, r *http.Request) string {
	j, err := C.AddAsset(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return ""
	}
	return j
}

// UpdateAsset updates an asset on the chain
func (comp *Comparser) UpdateAsset(w http.ResponseWriter, r *http.Request) string {
	j, err := C.UpdateAsset(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return ""
	}
	return j
}

// DeleteAsset deletes an asset from the chain
func (comp *Comparser) DeleteAsset(w http.ResponseWriter, r *http.Request) string {
	idSplit := strings.Split(r.RequestURI, "/")
	id := idSplit[len(idSplit)-1]
	j, err := C.DeleteAsset(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return ""
	}
	return j
}
