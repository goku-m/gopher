package main

import (
	"slices"
	"strings"
)

type Teacher struct {
	Name           string
	Subjects       []string
	Certifications *Certification
}

func NewTeacher(name string, subjects []string, certs []string) *Teacher {
	certification := NewCertification(certs...)
	teacher := &Teacher{
		Name:           name,
		Subjects:       subjects,
		Certifications: certification,
	}

	return teacher
}

func (t *Teacher) GetCerts() []string {
	return t.Certifications.GetCerts()
}

func (t *Teacher) GetSubjects() []string {
	return t.Subjects
}

type Certification struct {
	Wassce bool
	Bsc    bool
	Msc    bool
	Phd    bool
}

func NewCertification(c ...string) *Certification {
	return &Certification{
		Wassce: check(c, "wassce"),
		Bsc:    check(c, "bsc"),
		Msc:    check(c, "msc"),
		Phd:    check(c, "phd"),
	}
}

func (c *Certification) GetCerts() []string {
	var certs []string
	if c.Bsc {
		certs = append(certs, "Bsc")
	}
	if c.Msc {
		certs = append(certs, "Msc")
	}
	if c.Phd {
		certs = append(certs, "Phd")
	}
	if c.Wassce {
		certs = append(certs, "Wassce")
	}

	return certs
}

func check(cert []string, ct string) bool {
	for i, v := range cert {
		cert[i] = strings.ToLower(v)
	}
	return slices.Contains(cert, ct)
}
