/*
Copyright 2025 Nscale.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package trivy

import (
	"time"
)

// Severity is used to parse the value from a report into a programmatic value that can be used for comparisons later.
type Severity string

const (
	UNKNOWN  Severity = "UNKNOWN"
	LOW      Severity = "LOW"
	MEDIUM   Severity = "MEDIUM"
	HIGH     Severity = "HIGH"
	CRITICAL Severity = "CRITICAL"
)

//// CheckSeverity compares two severities to see if a threshold has been met. IE: is sev: HIGH >= check: MEDIUM.
//func CheckSeverity(sev, threshold Severity) bool {
//	var sevValue, thresholdValue int
//	sevValue = parseSeverity(sev)
//	thresholdValue = parseSeverity(threshold)
//
//	return sevValue >= thresholdValue
//}

// ValidSeverity confirms that the supplied value is a valid severity value.
func ValidSeverity(val Severity) bool {
	switch val {
	case UNKNOWN:
		return true
	case LOW:
		return true
	case MEDIUM:
		return true
	case HIGH:
		return true
	case CRITICAL:
		return true
	}

	return false
}

// ParseSeverity takes a Severity and returns everything from that severity value upwards as a string slice
func ParseSeverity(val Severity) []string {
	severityList := []string{"UNKNOWN", "LOW", "MEDIUM", "HIGH", "CRITICAL"}
	index := -1
	for i, v := range severityList {
		if Severity(v) == val {
			index = i
		}
	}

	return severityList[index:]
}

type ScanFailedReport struct {
	VulnerabilityID  string `json:"VulnerabilityID"`
	Description      string `json:"Description"`
	PkgName          string `json:"PkgName"`
	InstalledVersion string `json:"InstalledVersion"`
	Severity         string `json:"Severity"`
	Cvss             CVSS   `json:"CVSS"`
	FixedVersion     string `json:"FixedVersion"`
}

// Report and all its sub-structs is used to unmarshal the json reports into a usable format.
type Report struct {
	Name          string `json:"name"`
	ShortName     string `json:"short_name"`
	SchemaVersion int    `json:"SchemaVersion"`
	ArtifactName  string `json:"ArtifactName"`
	ArtifactType  string `json:"ArtifactType"`
	Metadata      struct {
		Os struct {
			Family string `json:"Family"`
			Name   string `json:"Name"`
		} `json:"OS"`
		ImageConfig struct {
			Architecture string    `json:"architecture"`
			Created      time.Time `json:"created"`
			Os           string    `json:"os"`
			Rootfs       struct {
				Type    string `json:"type"`
				DiffIds any    `json:"diff_ids"`
			} `json:"rootfs"`
			Config struct {
			} `json:"config"`
		} `json:"ImageConfig"`
	} `json:"Metadata"`
	Results []struct {
		Target          string            `json:"Target"`
		Class           string            `json:"Class"`
		Type            string            `json:"Type"`
		Vulnerabilities []Vulnerabilities `json:"Vulnerabilities"`
		Secrets         []Secrets         `json:"Secrets"`
	} `json:"Results"`
}

// CVSS stores all the score data from different sources within the Trivy report.
type CVSS struct {
	Ghsa   *Score `json:"ghsa"`
	Nvd    *Score `json:"nvd"`
	Redhat *Score `json:"redhat"`
}

// Score contains the score values and vectors from a Trivy report.
type Score struct {
	V2Vector string  `json:"V2Vector"`
	V3Vector string  `json:"V3Vector"`
	V2Score  float64 `json:"V2Score"`
	V3Score  float64 `json:"V3Score"`
}

// Vulnerabilities contains the vulnerability information from a Trivy report.
type Vulnerabilities struct {
	VulnerabilityID  string `json:"VulnerabilityID"`
	PkgID            string `json:"PkgID"`
	PkgName          string `json:"PkgName"`
	InstalledVersion string `json:"InstalledVersion"`
	Layer            struct {
		Digest string `json:"Digest"`
		DiffID string `json:"DiffID"`
	} `json:"layer"`
	SeveritySource string `json:"SeveritySource"`
	PrimaryURL     string `json:"PrimaryURL"`
	DataSource     struct {
		ID   string `json:"ID"`
		Name string `json:"Name"`
		URL  string `json:"URL"`
	} `json:"DataSource"`
	Title            string    `json:"Title"`
	Description      string    `json:"Description"`
	Severity         string    `json:"Severity"`
	CweIDs           []string  `json:"CweIDs"`
	Cvss             CVSS      `json:"CVSS"`
	References       []string  `json:"References"`
	PublishedDate    time.Time `json:"PublishedDate"`
	LastModifiedDate time.Time `json:"LastModifiedDate"`
	FixedVersion     string    `json:"FixedVersion"`
}

// Secrets contains the secret information from a Trivy report.
type Secrets struct {
	RuleID    string `json:"RuleID"`
	Category  string `json:"Category"`
	Severity  string `json:"Severity"`
	Title     string `json:"Title"`
	StartLine int    `json:"StartLine"`
	EndLine   int    `json:"EndLine"`
	Code      struct {
		Lines []struct {
			Number      int    `json:"Number"`
			Content     string `json:"Content"`
			IsCause     bool   `json:"IsCause"`
			Annotation  string `json:"Annotation"`
			Truncated   bool   `json:"Truncated"`
			Highlighted string `json:"Highlighted"`
			FirstCause  bool   `json:"FirstCause"`
			LastCause   bool   `json:"LastCause"`
		} `json:"Lines"`
	} `json:"Code"`
	Match string `json:"Match"`
	Layer struct {
		Digest string `json:"Digest"`
		DiffID string `json:"DiffID"`
	} `json:"Layer"`
}

// Month is used in reports parsing. It is contained within a Year and contains multiple trivy.Report(s).
type Month struct {
	Reports map[string]Report
}

// Year is used in reports parsing. It is the top level and contains multiple Month(s).
type Year struct {
	Months map[string]Month
}
