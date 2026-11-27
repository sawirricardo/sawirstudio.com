package site

// Lang is a translation of the site, built by scripts/i18n.py under /<Prefix>/.
// Keep this list in the same order as LANGS in scripts/i18n.py.
type Lang struct {
	Code, Prefix, Name string
}

var Langs = []Lang{
	{"zh-Hans", "zh-hans", "简体中文"},
	{"zh-Hant", "zh-hant", "繁體中文"},
	{"ja", "ja", "日本語"},
	{"de", "de", "Deutsch"},
	{"he", "he", "עברית"},
	{"ru", "ru", "Русский"},
	{"fr", "fr", "Français"},
	{"hi", "hi", "हिन्दी"},
	{"nl", "nl", "Nederlands"},
	{"it", "it", "Italiano"},
	{"ar", "ar", "العربية"},
}
