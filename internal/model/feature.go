package model

type ComponentPermission struct {
	Name         string   `json:"name"`
	ContentTypes []string `json:"content_types"`
}

type Feature struct {
	Name                 string                `json:"name"`
	Identifier           string                `json:"identifier"`
	PermissionName       string                `json:"permission_name"`
	Version              string                `json:"version"`
	Description          string                `json:"description"`
	IndexPath            string                `json:"index_path"`
	ContentType          string                `json:"content_type"`
	Area                 string                `json:"area"`
	Interchangeable      bool                  `json:"interchangeable"`
	ExpiresAt            int64                 `json:"expires_at"`
	NoExpire             bool                  `json:"no_expire"`
	RoleName             string                `json:"role_name"`
	NoteType             string                `json:"note_type"`
	FileType             string                `json:"file_type"`
	SpellCheckControl    bool                  `json:"spellcheckControl"`
	NoMobile             bool                  `json:"no_mobile"`
	Layerable            bool                  `json:"layerable"`
	ComponentPermissions []ComponentPermission `json:"component_permissions"`
}

func NewFeature(feature Feature, roleName string) Feature {
	return Feature{
		ExpiresAt:      8640000000000000,
		NoExpire:       false,
		Name:           feature.Name,
		Identifier:     feature.Identifier,
		PermissionName: feature.PermissionName,
	}
}

func NewEditorFeature(feature Feature, roleName string) Feature {
	return Feature{
		RoleName:    roleName,
		Area:        "editor-editor",
		ContentType: "SN|Component",
		IndexPath:   "dist/index.html",
		ComponentPermissions: []ComponentPermission{
			{
				Name: "stream-context-item",
				ContentTypes: []string{
					"Note",
				},
			},
		},
		Version:           "0.0.0",
		ExpiresAt:         8640000000000000,
		NoExpire:          false,
		Description:       feature.Description,
		Name:              feature.Name,
		Identifier:        feature.Identifier,
		PermissionName:    feature.PermissionName,
		Interchangeable:   feature.Interchangeable,
		NoteType:          feature.NoteType,
		FileType:          feature.FileType,
		SpellCheckControl: feature.SpellCheckControl,
		NoMobile:          feature.NoMobile,
		Layerable:         feature.Layerable,
	}
}

func NewThemeFeature(feature Feature, roleName string) Feature {
	return Feature{
		RoleName:          roleName,
		Area:              "editor-editor",
		ContentType:       "SN|Theme",
		IndexPath:         "dist/dist.css",
		Version:           "0.0.0",
		ExpiresAt:         8640000000000000,
		NoExpire:          false,
		Description:       feature.Description,
		Name:              feature.Name,
		Identifier:        feature.Identifier,
		PermissionName:    feature.PermissionName,
		Interchangeable:   feature.Interchangeable,
		NoteType:          feature.NoteType,
		FileType:          feature.FileType,
		SpellCheckControl: feature.SpellCheckControl,
		NoMobile:          feature.NoMobile,
		Layerable:         feature.Layerable,
	}
}
