package email

import "github.com/mikluko/jmap"

const SMIMEVerify jmap.URI = "urn:ietf:params:jmap:smimeverify"

type smimeVerify struct{}

func (s *smimeVerify) URI() jmap.URI { return SMIMEVerify }

func (s *smimeVerify) New() jmap.Capability { return &smimeVerify{} }
