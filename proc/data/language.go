package data

// BinSearch checks if needle is in the Haystack with BinarySearch
func BinSearch(haystack []string, needle string) bool {
	if len(haystack) == 0 {
		return false
	}

	if len(haystack) == 1 {
		return haystack[0] == needle
	}

	var front, back int
	front = 0
	back = len(haystack) - 1
	for front <= back {
		middle := (front + back) / 2
		if haystack[middle] < needle {
			front = middle + 1
		} else if haystack[middle] > needle {
			back = middle - 1
		} else {
			return true
		}
	}
	return false
}

// getLanguage maps a given language tag to the nominal value
func getLanguage(choosen string) string {
	if choosen == "" {
		return languages[0]
	} else if BinSearch(languages, choosen) {
		return choosen
	} else {
		return languages[1]
	}
}

// ExtractLanguages extracts the Languages from the BrowserInfo and saves it
func (features *ProcessedFeatures) ExtractLanguages(binfo BrowserInfo) {
	lang := binfo.Languages
	if len(lang) > 0 {
		features.Lang1 = getLanguage(lang[0])
	} else {
		features.Lang1 = languages[0]
	}
	if len(lang) > 1 {
		features.Lang2 = getLanguage(lang[1])
	} else {
		features.Lang2 = languages[0]
	}
	if len(lang) > 2 {
		features.Lang3 = getLanguage(lang[2])
	} else {
		features.Lang3 = languages[0]
	}
	if len(lang) > 3 {
		features.Lang4 = getLanguage(lang[3])
	} else {
		features.Lang4 = languages[0]
	}
	if len(lang) > 4 {
		features.Lang5 = getLanguage(lang[4])
	} else {
		features.Lang5 = languages[0]
	}
	if len(lang) > 5 {
		features.Lang6 = getLanguage(lang[5])
	} else {
		features.Lang6 = languages[0]
	}
	if len(lang) > 6 {
		features.Lang7 = getLanguage(lang[6])
	} else {
		features.Lang7 = languages[0]
	}
	if len(lang) > 7 {
		features.Lang8 = getLanguage(lang[7])
	} else {
		features.Lang8 = languages[0]
	}

	if len(lang) > 8 {
		features.MoreLanguages = 1
	} else {
		features.MoreLanguages = 0
	}
}

// All possible language codes
// Source: https://datahub.io/core/language-codes#data
// This list was made available under the Public Domain Dedication and License v1.0 whose full text can be found at: http://www.opendatacommons.org/licenses/pddl/1.0/
var languages = []string{
	"EMPTY",       // lang was not chosen
	"NOT NOMINAL", // lang is not in this list,
	"af",
	"af-NA",
	"af-ZA",
	"agq",
	"agq-CM",
	"ak",
	"ak-GH",
	"am",
	"am-ET",
	"ar",
	"ar-001",
	"ar-AE",
	"ar-BH",
	"ar-DJ",
	"ar-DZ",
	"ar-EG",
	"ar-EH",
	"ar-ER",
	"ar-IL",
	"ar-IQ",
	"ar-JO",
	"ar-KM",
	"ar-KW",
	"ar-LB",
	"ar-LY",
	"ar-MA",
	"ar-MR",
	"ar-OM",
	"ar-PS",
	"ar-QA",
	"ar-SA",
	"ar-SD",
	"ar-SO",
	"ar-SS",
	"ar-SY",
	"ar-TD",
	"ar-TN",
	"ar-YE",
	"as",
	"as-IN",
	"asa",
	"asa-TZ",
	"ast",
	"ast-ES",
	"az",
	"az-Cyrl",
	"az-Cyrl-AZ",
	"az-Latn",
	"az-Latn-AZ",
	"bas",
	"bas-CM",
	"be",
	"be-BY",
	"bem",
	"bem-ZM",
	"bez",
	"bez-TZ",
	"bg",
	"bg-BG",
	"bm",
	"bm-ML",
	"bn",
	"bn-BD",
	"bn-IN",
	"bo",
	"bo-CN",
	"bo-IN",
	"br",
	"br-FR",
	"brx",
	"brx-IN",
	"bs",
	"bs-Cyrl",
	"bs-Cyrl-BA",
	"bs-Latn",
	"bs-Latn-BA",
	"ca",
	"ca-AD",
	"ca-ES",
	"ca-ES-VALENCIA",
	"ca-FR",
	"ca-IT",
	"ccp",
	"ccp-BD",
	"ccp-IN",
	"ce",
	"ce-RU",
	"ceb",
	"ceb-PH",
	"cgg",
	"cgg-UG",
	"chr",
	"chr-US",
	"ckb",
	"ckb-IQ",
	"ckb-IR",
	"cs",
	"cs-CZ",
	"cu",
	"cu-RU",
	"cy",
	"cy-GB",
	"da",
	"da-DK",
	"da-GL",
	"dav",
	"dav-KE",
	"de",
	"de-AT",
	"de-BE",
	"de-CH",
	"de-DE",
	"de-IT",
	"de-LI",
	"de-LU",
	"dje",
	"dje-NE",
	"dsb",
	"dsb-DE",
	"dua",
	"dua-CM",
	"dyo",
	"dyo-SN",
	"dz",
	"dz-BT",
	"ebu",
	"ebu-KE",
	"ee",
	"ee-GH",
	"ee-TG",
	"el",
	"el-CY",
	"el-GR",
	"en",
	"en-001",
	"en-150",
	"en-AE",
	"en-AG",
	"en-AI",
	"en-AS",
	"en-AT",
	"en-AU",
	"en-BB",
	"en-BE",
	"en-BI",
	"en-BM",
	"en-BS",
	"en-BW",
	"en-BZ",
	"en-CA",
	"en-CC",
	"en-CH",
	"en-CK",
	"en-CM",
	"en-CX",
	"en-CY",
	"en-DE",
	"en-DG",
	"en-DK",
	"en-DM",
	"en-ER",
	"en-FI",
	"en-FJ",
	"en-FK",
	"en-FM",
	"en-GB",
	"en-GD",
	"en-GG",
	"en-GH",
	"en-GI",
	"en-GM",
	"en-GU",
	"en-GY",
	"en-HK",
	"en-IE",
	"en-IL",
	"en-IM",
	"en-IN",
	"en-IO",
	"en-JE",
	"en-JM",
	"en-KE",
	"en-KI",
	"en-KN",
	"en-KY",
	"en-LC",
	"en-LR",
	"en-LS",
	"en-MG",
	"en-MH",
	"en-MO",
	"en-MP",
	"en-MS",
	"en-MT",
	"en-MU",
	"en-MW",
	"en-MY",
	"en-NA",
	"en-NF",
	"en-NG",
	"en-NL",
	"en-NR",
	"en-NU",
	"en-NZ",
	"en-PG",
	"en-PH",
	"en-PK",
	"en-PN",
	"en-PR",
	"en-PW",
	"en-RW",
	"en-SB",
	"en-SC",
	"en-SD",
	"en-SE",
	"en-SG",
	"en-SH",
	"en-SI",
	"en-SL",
	"en-SS",
	"en-SX",
	"en-SZ",
	"en-TC",
	"en-TK",
	"en-TO",
	"en-TT",
	"en-TV",
	"en-TZ",
	"en-UG",
	"en-UM",
	"en-US",
	"en-US-POSIX",
	"en-VC",
	"en-VG",
	"en-VI",
	"en-VU",
	"en-WS",
	"en-ZA",
	"en-ZM",
	"en-ZW",
	"eo",
	"eo-001",
	"es",
	"es-419",
	"es-AR",
	"es-BO",
	"es-BR",
	"es-BZ",
	"es-CL",
	"es-CO",
	"es-CR",
	"es-CU",
	"es-DO",
	"es-EA",
	"es-EC",
	"es-ES",
	"es-GQ",
	"es-GT",
	"es-HN",
	"es-IC",
	"es-MX",
	"es-NI",
	"es-PA",
	"es-PE",
	"es-PH",
	"es-PR",
	"es-PY",
	"es-SV",
	"es-US",
	"es-UY",
	"es-VE",
	"et",
	"et-EE",
	"eu",
	"eu-ES",
	"ewo",
	"ewo-CM",
	"fa",
	"fa-AF",
	"fa-IR",
	"ff",
	"ff-Adlm",
	"ff-Adlm-BF",
	"ff-Adlm-CM",
	"ff-Adlm-GH",
	"ff-Adlm-GM",
	"ff-Adlm-GN",
	"ff-Adlm-GW",
	"ff-Adlm-LR",
	"ff-Adlm-MR",
	"ff-Adlm-NE",
	"ff-Adlm-NG",
	"ff-Adlm-SL",
	"ff-Adlm-SN",
	"ff-Latn",
	"ff-Latn-BF",
	"ff-Latn-CM",
	"ff-Latn-GH",
	"ff-Latn-GM",
	"ff-Latn-GN",
	"ff-Latn-GW",
	"ff-Latn-LR",
	"ff-Latn-MR",
	"ff-Latn-NE",
	"ff-Latn-NG",
	"ff-Latn-SL",
	"ff-Latn-SN",
	"fi",
	"fi-FI",
	"fil",
	"fil-PH",
	"fo",
	"fo-DK",
	"fo-FO",
	"fr",
	"fr-BE",
	"fr-BF",
	"fr-BI",
	"fr-BJ",
	"fr-BL",
	"fr-CA",
	"fr-CD",
	"fr-CF",
	"fr-CG",
	"fr-CH",
	"fr-CI",
	"fr-CM",
	"fr-DJ",
	"fr-DZ",
	"fr-FR",
	"fr-GA",
	"fr-GF",
	"fr-GN",
	"fr-GP",
	"fr-GQ",
	"fr-HT",
	"fr-KM",
	"fr-LU",
	"fr-MA",
	"fr-MC",
	"fr-MF",
	"fr-MG",
	"fr-ML",
	"fr-MQ",
	"fr-MR",
	"fr-MU",
	"fr-NC",
	"fr-NE",
	"fr-PF",
	"fr-PM",
	"fr-RE",
	"fr-RW",
	"fr-SC",
	"fr-SN",
	"fr-SY",
	"fr-TD",
	"fr-TG",
	"fr-TN",
	"fr-VU",
	"fr-WF",
	"fr-YT",
	"fur",
	"fur-IT",
	"fy",
	"fy-NL",
	"ga",
	"ga-GB",
	"ga-IE",
	"gd",
	"gd-GB",
	"gl",
	"gl-ES",
	"gsw",
	"gsw-CH",
	"gsw-FR",
	"gsw-LI",
	"gu",
	"gu-IN",
	"guz",
	"guz-KE",
	"gv",
	"gv-IM",
	"ha",
	"ha-GH",
	"ha-NE",
	"ha-NG",
	"haw",
	"haw-US",
	"he",
	"he-IL",
	"hi",
	"hi-IN",
	"hr",
	"hr-BA",
	"hr-HR",
	"hsb",
	"hsb-DE",
	"hu",
	"hu-HU",
	"hy",
	"hy-AM",
	"ia",
	"ia-001",
	"id",
	"id-ID",
	"ig",
	"ig-NG",
	"ii",
	"ii-CN",
	"is",
	"is-IS",
	"it",
	"it-CH",
	"it-IT",
	"it-SM",
	"it-VA",
	"ja",
	"ja-JP",
	"jgo",
	"jgo-CM",
	"jmc",
	"jmc-TZ",
	"jv",
	"jv-ID",
	"ka",
	"ka-GE",
	"kab",
	"kab-DZ",
	"kam",
	"kam-KE",
	"kde",
	"kde-TZ",
	"kea",
	"kea-CV",
	"khq",
	"khq-ML",
	"ki",
	"ki-KE",
	"kk",
	"kk-KZ",
	"kkj",
	"kkj-CM",
	"kl",
	"kl-GL",
	"kln",
	"kln-KE",
	"km",
	"km-KH",
	"kn",
	"kn-IN",
	"ko",
	"ko-KP",
	"ko-KR",
	"kok",
	"kok-IN",
	"ks",
	"ks-Arab",
	"ks-Arab-IN",
	"ksb",
	"ksb-TZ",
	"ksf",
	"ksf-CM",
	"ksh",
	"ksh-DE",
	"ku",
	"ku-TR",
	"kw",
	"kw-GB",
	"ky",
	"ky-KG",
	"lag",
	"lag-TZ",
	"lb",
	"lb-LU",
	"lg",
	"lg-UG",
	"lkt",
	"lkt-US",
	"ln",
	"ln-AO",
	"ln-CD",
	"ln-CF",
	"ln-CG",
	"lo",
	"lo-LA",
	"lrc",
	"lrc-IQ",
	"lrc-IR",
	"lt",
	"lt-LT",
	"lu",
	"lu-CD",
	"luo",
	"luo-KE",
	"luy",
	"luy-KE",
	"lv",
	"lv-LV",
	"mai",
	"mai-IN",
	"mas",
	"mas-KE",
	"mas-TZ",
	"mer",
	"mer-KE",
	"mfe",
	"mfe-MU",
	"mg",
	"mg-MG",
	"mgh",
	"mgh-MZ",
	"mgo",
	"mgo-CM",
	"mi",
	"mi-NZ",
	"mk",
	"mk-MK",
	"ml",
	"ml-IN",
	"mn",
	"mn-MN",
	"mni",
	"mni-Beng",
	"mni-Beng-IN",
	"mr",
	"mr-IN",
	"ms",
	"ms-BN",
	"ms-ID",
	"ms-MY",
	"ms-SG",
	"mt",
	"mt-MT",
	"mua",
	"mua-CM",
	"my",
	"my-MM",
	"mzn",
	"mzn-IR",
	"naq",
	"naq-NA",
	"nb",
	"nb-NO",
	"nb-SJ",
	"nd",
	"nd-ZW",
	"nds",
	"nds-DE",
	"nds-NL",
	"ne",
	"ne-IN",
	"ne-NP",
	"nl",
	"nl-AW",
	"nl-BE",
	"nl-BQ",
	"nl-CW",
	"nl-NL",
	"nl-SR",
	"nl-SX",
	"nmg",
	"nmg-CM",
	"nn",
	"nn-NO",
	"nnh",
	"nnh-CM",
	"nus",
	"nus-SS",
	"nyn",
	"nyn-UG",
	"om",
	"om-ET",
	"om-KE",
	"or",
	"or-IN",
	"os",
	"os-GE",
	"os-RU",
	"pa",
	"pa-Arab",
	"pa-Arab-PK",
	"pa-Guru",
	"pa-Guru-IN",
	"pcm",
	"pcm-NG",
	"pl",
	"pl-PL",
	"prg",
	"prg-001",
	"ps",
	"ps-AF",
	"ps-PK",
	"pt",
	"pt-AO",
	"pt-BR",
	"pt-CH",
	"pt-CV",
	"pt-GQ",
	"pt-GW",
	"pt-LU",
	"pt-MO",
	"pt-MZ",
	"pt-PT",
	"pt-ST",
	"pt-TL",
	"qu",
	"qu-BO",
	"qu-EC",
	"qu-PE",
	"rm",
	"rm-CH",
	"rn",
	"rn-BI",
	"ro",
	"ro-MD",
	"ro-RO",
	"rof",
	"rof-TZ",
	"root",
	"ru",
	"ru-BY",
	"ru-KG",
	"ru-KZ",
	"ru-MD",
	"ru-RU",
	"ru-UA",
	"rw",
	"rw-RW",
	"rwk",
	"rwk-TZ",
	"sah",
	"sah-RU",
	"saq",
	"saq-KE",
	"sat",
	"sat-Olck",
	"sat-Olck-IN",
	"sbp",
	"sbp-TZ",
	"sd",
	"sd-Arab",
	"sd-Arab-PK",
	"sd-Deva",
	"sd-Deva-IN",
	"se",
	"se-FI",
	"se-NO",
	"se-SE",
	"seh",
	"seh-MZ",
	"ses",
	"ses-ML",
	"sg",
	"sg-CF",
	"shi",
	"shi-Latn",
	"shi-Latn-MA",
	"shi-Tfng",
	"shi-Tfng-MA",
	"si",
	"si-LK",
	"sk",
	"sk-SK",
	"sl",
	"sl-SI",
	"smn",
	"smn-FI",
	"sn",
	"sn-ZW",
	"so",
	"so-DJ",
	"so-ET",
	"so-KE",
	"so-SO",
	"sq",
	"sq-AL",
	"sq-MK",
	"sq-XK",
	"sr",
	"sr-Cyrl",
	"sr-Cyrl-BA",
	"sr-Cyrl-ME",
	"sr-Cyrl-RS",
	"sr-Cyrl-XK",
	"sr-Latn",
	"sr-Latn-BA",
	"sr-Latn-ME",
	"sr-Latn-RS",
	"sr-Latn-XK",
	"su",
	"su-Latn",
	"su-Latn-ID",
	"sv",
	"sv-AX",
	"sv-FI",
	"sv-SE",
	"sw",
	"sw-CD",
	"sw-KE",
	"sw-TZ",
	"sw-UG",
	"ta",
	"ta-IN",
	"ta-LK",
	"ta-MY",
	"ta-SG",
	"te",
	"te-IN",
	"teo",
	"teo-KE",
	"teo-UG",
	"tg",
	"tg-TJ",
	"th",
	"th-TH",
	"ti",
	"ti-ER",
	"ti-ET",
	"tk",
	"tk-TM",
	"to",
	"to-TO",
	"tr",
	"tr-CY",
	"tr-TR",
	"tt",
	"tt-RU",
	"twq",
	"twq-NE",
	"tzm",
	"tzm-MA",
	"ug",
	"ug-CN",
	"uk",
	"uk-UA",
	"ur",
	"ur-IN",
	"ur-PK",
	"uz",
	"uz-Arab",
	"uz-Arab-AF",
	"uz-Cyrl",
	"uz-Cyrl-UZ",
	"uz-Latn",
	"uz-Latn-UZ",
	"vai",
	"vai-Latn",
	"vai-Latn-LR",
	"vai-Vaii",
	"vai-Vaii-LR",
	"vi",
	"vi-VN",
	"vo",
	"vo-001",
	"vun",
	"vun-TZ",
	"wae",
	"wae-CH",
	"wo",
	"wo-SN",
	"xh",
	"xh-ZA",
	"xog",
	"xog-UG",
	"yav",
	"yav-CM",
	"yi",
	"yi-001",
	"yo",
	"yo-BJ",
	"yo-NG",
	"yue",
	"yue-Hans",
	"yue-Hans-CN",
	"yue-Hant",
	"yue-Hant-HK",
	"zgh",
	"zgh-MA",
	"zh",
	"zh-Hans",
	"zh-Hans-CN",
	"zh-Hans-HK",
	"zh-Hans-MO",
	"zh-Hans-SG",
	"zh-Hant",
	"zh-Hant-HK",
	"zh-Hant-MO",
	"zh-Hant-TW",
	"zu",
	"zu-ZA",
}
