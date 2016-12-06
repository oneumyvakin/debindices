// Parser for Binary Packages Indices files(Packages.gz, Packages.bz2, Packages) of Debian/Ubuntu repository
package debindices

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

// Supported fields
var Fields = []string{
	"Package",
	"Priority",
	"Section",
	"Installed-Size",
	"Maintainer",
	"Architecture",
	"Version",
	"Depends",
	"Filename",
	"Size",
	"MD5sum",
	"SHA1",
	"SHA256",
}

// Package contains parsed data
type Package struct {
	Package       string
	Priority      string
	Section       string
	InstalledSize int64
	Maintainer    string
	Architecture  string
	Version       string
	Depends       string
	Filename      string
	Size          int64
	MD5sum        string
	SHA1          string
	SHA256        string
}

// Parse takes plain-text reader and returns map of Packages
// with keys equal to field name defined by fieldKey
// Parse fails on duplicate key if failOnDup is true
func Parse(r io.Reader, fieldKey string, failOnDup bool) (packages map[string]Package, err error) {
	if !strings.Contains(strings.Join(Fields, "|"), fieldKey) {
		err = errors.New(fmt.Sprintf("unknown field can't be used as map key: %s", fieldKey))
		return
	}

	packages = make(map[string]Package)
	regList := make(map[string]*regexp.Regexp)
	for _, field := range Fields {
		regList[field] = regexp.MustCompile(fmt.Sprintf("(?i)^%s: (.+)$", field))
	}

	s := bufio.NewScanner(r)
	p := Package{}
	key := ""
	for s.Scan() {
		curStr := s.Text()
		if curStr == "" {
			if failOnDup {
				if _, ok := packages[key]; ok {
					err = errors.New(fmt.Sprintf("duplicate package for field key %s: %s", fieldKey, key))
					return
				}
			}
			packages[key] = p

			p = Package{}
			key = ""
		}
		for field, reg := range regList {
			match := reg.FindAllStringSubmatch(curStr, -1)
			if len(match) == 0 {
				continue
			}
			val := match[0][1]

			if field == fieldKey {
				key = val
			}

			switch field {
			case "Package":
				p.setPackage(val)
				break
			case "Priority":
				p.setPriority(val)
				break
			case "Section":
				p.setSection(val)
				break
			case "Installed-Size":
				var valInt int
				valInt, err = strconv.Atoi(val)
				if err != nil {
					return
				}
				p.setInstalledSize(int64(valInt))
				break
			case "Maintainer":
				p.setMaintainer(val)
				break
			case "Architecture":
				p.setArchitecture(val)
				break
			case "Version":
				p.setVersion(val)
				break
			case "Depends":
				p.setDepends(val)
				break
			case "Filename":
				p.setFilename(val)
				break
			case "Size":
				var valInt int
				valInt, err = strconv.Atoi(val)
				if err != nil {
					return
				}
				p.setSize(int64(valInt))
				break
			case "MD5sum":
				p.setMD5sum(val)
				break
			case "SHA1":
				p.setSHA1(val)
				break
			case "SHA256":
				p.setSHA256(val)
				break
			}
		}

	}
	return
}

func (p *Package) setPackage(name string) {
	p.Package = name
}

func (p *Package) setPriority(priority string) {
	p.Priority = priority
}

func (p *Package) setSection(section string) {
	p.Section = section
}

func (p *Package) setInstalledSize(size int64) {
	p.InstalledSize = size
}

func (p *Package) setMaintainer(maintainer string) {
	p.Maintainer = maintainer
}

func (p *Package) setArchitecture(architecture string) {
	p.Architecture = architecture
}

func (p *Package) setVersion(version string) {
	p.Version = version
}

func (p *Package) setDepends(depends string) {
	p.Depends = depends
}

func (p *Package) setFilename(filename string) {
	p.Filename = filename
}

func (p *Package) setSize(size int64) {
	p.Size = size
}

func (p *Package) setMD5sum(md5sum string) {
	p.MD5sum = md5sum
}

func (p *Package) setSHA1(sha1 string) {
	p.SHA1 = sha1
}

func (p *Package) setSHA256(sha256 string) {
	p.SHA256 = sha256
}