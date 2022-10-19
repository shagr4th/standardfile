package service

import (
	"github.com/mdouchement/standardfile/internal/database"
	"github.com/mdouchement/standardfile/internal/model"
)

var Editors = []model.Feature{
	// EDITORS
	{
		Name:            "Spreadsheet",
		Identifier:      "org.standardnotes.standard-sheets",
		PermissionName:  "editor:sheets",
		NoteType:        "spreadsheet",
		FileType:        "json",
		Interchangeable: false,
	},
	{
		Name:              "Minimal Markdown",
		Identifier:        "org.standardnotes.minimal-markdown-editor",
		PermissionName:    "editor:markdown-minimist",
		NoteType:          "markdown",
		FileType:          "md",
		SpellCheckControl: true,
		Interchangeable:   true,
	},
	{
		Name:              "Markdown with Math",
		Identifier:        "org.standardnotes.fancy-markdown-editor",
		PermissionName:    "editor:markdown-math",
		NoteType:          "markdown",
		FileType:          "md",
		SpellCheckControl: true,
		Interchangeable:   true,
	},
	{
		Name:              "Hybrid Markdown",
		Identifier:        "org.standardnotes.advanced-markdown-editor",
		PermissionName:    "editor:markdown-pro",
		NoteType:          "markdown",
		FileType:          "md",
		SpellCheckControl: true,
		Interchangeable:   true,
	},
	{
		Name:              "Advanced Checklist [Beta]",
		Identifier:        "org.standardnotes.advanced-checklist",
		PermissionName:    "editor:advanced-checklist",
		NoteType:          "markdown",
		FileType:          "md",
		SpellCheckControl: false,
		Interchangeable:   true,
	},
	{
		Name:            "Passwords and Tokens",
		Identifier:      "org.standardnotes.token-vault",
		PermissionName:  "editor:token-vault",
		NoteType:        "authentication",
		FileType:        "json",
		Interchangeable: false,
	},
	{
		Name:              "Dynamic Markdown",
		Identifier:        "org.standardnotes.markdown-visual-editor",
		PermissionName:    "editor:markdown-visual",
		NoteType:          "markdown",
		FileType:          "md",
		SpellCheckControl: true,
		Interchangeable:   true,
	},
	{
		Name:              "Classic Rich Text",
		Identifier:        "org.standardnotes.plus-editor",
		PermissionName:    "editor:plus",
		NoteType:          "rich-text",
		FileType:          "html",
		SpellCheckControl: true,
		Interchangeable:   true,
	},
	{
		Name:              "Basic Checklist",
		Identifier:        "org.standardnotes.simple-task-editor",
		PermissionName:    "editor:task-editor",
		NoteType:          "task",
		FileType:          "md",
		SpellCheckControl: true,
		Interchangeable:   false,
	},
	{
		Name:              "Code",
		Identifier:        "org.standardnotes.code-editor",
		PermissionName:    "editor:code-editor",
		NoteType:          "code",
		FileType:          "txt",
		SpellCheckControl: true,
		Interchangeable:   true,
	},
	{
		Name:              "Alternative Rich Text",
		Identifier:        "org.standardnotes.bold-editor",
		PermissionName:    "editor:bold",
		NoteType:          "rich-text",
		FileType:          "html",
		SpellCheckControl: true,
		Interchangeable:   true,
	},
	{
		Name:              "Basic Markdown",
		Identifier:        "org.standardnotes.simple-markdown-editor",
		PermissionName:    "editor:markdown-basic",
		NoteType:          "markdown",
		FileType:          "md",
		SpellCheckControl: true,
		Interchangeable:   true,
	},
}

var Themes = []model.Feature{
	// THEMES
	{
		Name:           "Focus",
		Identifier:     "org.standardnotes.theme-focus",
		PermissionName: "theme:focused",
	},
	{
		Name:           "Autobiography",
		Identifier:     "org.standardnotes.theme-autobiography",
		PermissionName: "theme:autobiography",
	},
	{
		Name:           "Midnight",
		Identifier:     "org.standardnotes.theme-midnight",
		PermissionName: "theme:midnight",
	},
	{
		Name:           "Dynamic",
		Identifier:     "org.standardnotes.theme-dynamic",
		PermissionName: "theme:dynamic",
		NoMobile:       true,
		Layerable:      true,
	},
	{
		Name:           "Solarized Dark",
		Identifier:     "org.standardnotes.theme-solarized-dark",
		PermissionName: "theme:solarized-dark",
	},
	{
		Name:           "Futura",
		Identifier:     "org.standardnotes.theme-futura",
		PermissionName: "theme:futura",
	},
	{
		Name:           "Titanium",
		Identifier:     "org.standardnotes.theme-titanium",
		PermissionName: "theme:titanium",
	},
}

var AppFeatures = []model.Feature{
	{
		Name:           "Focus Mode",
		Identifier:     "org.standardnotes.focus-mode",
		PermissionName: "app:focus-mode",
	},
	{
		Name:           "Tag Nesting",
		Identifier:     "org.standardnotes.tag-nesting",
		PermissionName: "app:tag-nesting",
	},
	{
		Name:           "Smart Filters",
		Identifier:     "org.standardnotes.smart-filters",
		PermissionName: "app:smart-filters",
	},
}

var ServerFeatures = []model.Feature{
	{
		Name:           "Sign-in email alerts",
		Identifier:     "com.standardnotes.sign-in-alerts",
		PermissionName: "server:sign-in-alerts",
	},
	{
		Identifier:     "org.standardnotes.files-max-storage-tier",
		PermissionName: "server:files-max-storage-tier",
	},
	{
		Identifier:     "org.standardnotes.daily-gdrive-backup",
		PermissionName: "server:daily-gdrive-backup",
	},
	{
		Identifier:     "org.standardnotes.daily-onedrive-backup",
		PermissionName: "server:daily-onedrive-backup",
	},
	{
		Name:           "Two factor authentication",
		Identifier:     "org.standardnotes.two-factor-auth",
		PermissionName: "server:two-factor-auth",
	},
	{
		Name:           "Unlimited note history",
		Identifier:     "org.standardnotes.note-history-unlimited",
		PermissionName: "server:note-history-unlimited",
	},
	{
		Identifier:     "org.standardnotes.daily-dropbox-backup",
		PermissionName: "server:daily-dropbox-backup",
	},
	{
		Name:           "Email backups",
		Identifier:     "org.standardnotes.daily-email-backup",
		PermissionName: "server:daily-email-backup",
	},
}

var MiscModules = []model.Feature{
	{
		Name:           "Listed Custom Domain",
		Identifier:     "org.standardnotes.listed-custom-domain",
		PermissionName: "listed:custom-domain",
	},
}

type (
	// A PKCEService is a service used for managing challenges.
	SubscriptionService interface {
		Features(user *model.User, subscription string) []model.Feature
		GetSubscription(user *model.User) *model.Subscription
		UserRoles(user *model.User) []model.SubscriptionRole
	}

	subscriptionService struct {
		db           database.Client
		FilesPath    string
		subscription string
	}
)

// NewSubscriptionService instantiates a new Subscription service.
func NewSubscriptionService(db database.Client, subscription string, params Params) (s SubscriptionService) {
	switch params.APIVersion { // for future API increments
	default:
		s = &subscriptionService{
			db:           db,
			subscription: subscription,
			FilesPath:    "", // TODO for future file support
		}
	}
	return s
}

func (s *subscriptionService) GetSubscription(user *model.User) *model.Subscription {
	if len(s.subscription) > 0 {
		return &model.Subscription{
			ID:               "caf81bfc-bfcf-11ec-b43b-0242ac120008",
			Name:             s.subscription,
			EndsAt:           8640000000000000,
			CreatedAt:        0,
			UpdatedAt:        0,
			Cancelled:        0,
			SubscriptionId:   1,
			SubscriptionType: "regular",
		}
	}

	return nil
}

func (s *subscriptionService) Features(user *model.User, subscription string) []model.Feature {
	var features []model.Feature = make([]model.Feature, 0)

	subscriptionRole := s.subscriptionRole()
	if subscriptionRole != nil {
		for i := 0; i < len(Editors); i++ {
			features = append(features, model.NewEditorFeature(Editors[i], subscriptionRole.Name))
		}
		for i := 0; i < len(Themes); i++ {
			features = append(features, model.NewThemeFeature(Themes[i], subscriptionRole.Name))
		}
		for i := 0; i < len(AppFeatures); i++ {
			features = append(features, model.NewFeature(AppFeatures[i], subscriptionRole.Name))
		}

		if len(s.FilesPath) > 0 {
			features = append(features, model.NewFeature(model.Feature{Name: "Encrypted files",
				Identifier:     "org.standardnotes.files",
				PermissionName: "app:files"}, subscriptionRole.Name))
		}
	}
	return features
}

func (s *subscriptionService) UserRoles(user *model.User) []model.SubscriptionRole {
	userRoles := []model.SubscriptionRole{
		{
			Name: "CORE_USER",
			ID:   "8802d6a3-b97c-4b25-968a-8fb21c65c3a1",
		},
	}
	subscriptionRole := s.subscriptionRole()
	if subscriptionRole != nil {
		userRoles = append(userRoles, *subscriptionRole)
	}
	return userRoles
}

func (s *subscriptionService) subscriptionRole() *model.SubscriptionRole {
	if s.subscription == "PRO_PLAN" {
		return &model.SubscriptionRole{
			Name: "PRO_USER",
			ID:   "8047edbb-a10a-4ff8-8d53-c2cae600a8e8",
		}
	}
	return nil
}
