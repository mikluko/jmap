package sieve

import "github.com/mikluko/jmap"

const URI jmap.URI = "urn:ietf:params:jmap:sieve"

const SieveScriptEvent jmap.EventType = "SieveScript"

func init() {
	jmap.RegisterCapability(&Capability{})
}

type Capability struct {
	// Implementation is the name of the Sieve implementation
	Implementation string `json:"implementation"`

	// MaxSizeScriptName is the maximum length of the script name in octets
	MaxSizeScriptName uint64 `json:"maxSizeScriptName"`

	// MaxSizeScript is the maximum size of a script in octets, or nil for no limit
	MaxSizeScript *uint64 `json:"maxSizeScript"`

	// MaxNumberScripts is the maximum number of scripts, or nil for no limit
	MaxNumberScripts *uint64 `json:"maxNumberScripts"`

	// MaxNumberRedirects is the maximum number of redirect actions in a script, or nil for no limit
	MaxNumberRedirects *uint64 `json:"maxNumberRedirects"`

	// SieveExtensions lists the Sieve extensions supported by the server
	SieveExtensions []string `json:"sieveExtensions"`

	// NotificationMethods lists the notification URI schemes supported, or nil if not supported
	NotificationMethods []string `json:"notificationMethods"`

	// ExternalLists lists the external list URI schemes supported, or nil if not supported
	ExternalLists []string `json:"externalLists"`
}

func (c *Capability) URI() jmap.URI       { return URI }
func (c *Capability) New() jmap.Capability { return &Capability{} }
