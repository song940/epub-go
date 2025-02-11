package epub

// Opf content.opf
type Opf struct {
	Metadata Metadata   `xml:"metadata" json:"metadata"`
	Manifest []Manifest `xml:"manifest>item" json:"manifest"`
	Spine    Spine      `xml:"spine" json:"spine"`
}

func (p *Opf) GetManifestItem(id string) *Manifest {
	for _, m := range p.Manifest {
		if m.ID == id {
			return &m
		}
	}
	return nil
}

// Metadata metadata
type Metadata struct {
	Title       []string     `xml:"title" json:"title"`
	Language    []string     `xml:"language" json:"language"`
	Identifier  []Identifier `xml:"identifier" json:"identifier"`
	Creator     []Author     `xml:"creator" json:"creator"`
	Subject     []string     `xml:"subject" json:"subject"`
	Description []string     `xml:"description" json:"description"`
	Publisher   []string     `xml:"publisher" json:"publisher"`
	Contributor []Author     `xml:"contributor" json:"contributor"`
	Date        []Date       `xml:"date" json:"date"`
	Type        []string     `xml:"type" json:"type"`
	Format      []string     `xml:"format" json:"format"`
	Source      []string     `xml:"source" json:"source"`
	Relation    []string     `xml:"relation" json:"relation"`
	Coverage    []string     `xml:"coverage" json:"coverage"`
	Rights      []string     `xml:"rights" json:"rights"`
	Meta        []Metafield  `xml:"meta" json:"meta"`
}

// Identifier identifier
type Identifier struct {
	Data   string `xml:",chardata" json:"data"`
	ID     string `xml:"id,attr" json:"id"`
	Scheme string `xml:"scheme,attr" json:"scheme"`
}

// Author author
type Author struct {
	Data   string `xml:",chardata" json:"author"`
	FileAs string `xml:"file-as,attr" json:"file_as"`
	Role   string `xml:"role,attr" json:"role"`
}

// Date date
type Date struct {
	Data  string `xml:",chardata" json:"data"`
	Event string `xml:"event,attr" json:"event"`
}

// Metafield metafield
type Metafield struct {
	Name    string `xml:"name,attr" json:"name"`
	Content string `xml:"content,attr" json:"content"`
}

// Manifest manifest
type Manifest struct {
	ID           string `xml:"id,attr" json:"id"`
	Href         string `xml:"href,attr" json:"href"`
	MediaType    string `xml:"media-type,attr" json:"type"`
	Fallback     string `xml:"media-fallback,attr" json:"fallback"`
	Properties   string `xml:"properties,attr" json:"properties"`
	MediaOverlay string `xml:"media-overlay,attr" json:"overlay"`
}

// Spine spine
type Spine struct {
	ID              string      `xml:"id,attr" json:"id"`
	Toc             string      `xml:"toc,attr" json:"toc"`
	PageProgression string      `xml:"page-progression-direction,attr" json:"progression"`
	Items           []SpineItem `xml:"itemref" json:"items"`
}

// SpineItem spine item
type SpineItem struct {
	IDref      string `xml:"idref,attr" json:"id_ref"`
	Linear     string `xml:"linear,attr" json:"linear"`
	ID         string `xml:"id,attr" json:"id"`
	Properties string `xml:"properties,attr" json:"properties"`
}
