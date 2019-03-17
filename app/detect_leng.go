package app

import "github.com/abadojack/whatlanggo"

var YandexLangMap = map[whatlanggo.Lang]string{
	whatlanggo.Afr: "af",
	whatlanggo.Aka: "ak",
	whatlanggo.Amh: "am",
	whatlanggo.Arb: "ar",
	whatlanggo.Azj: "az", // Azerbaijani iso 639-3 is aze, iso 639-1 az
	whatlanggo.Bel: "be",
	whatlanggo.Ben: "bn",
	whatlanggo.Bho: "bh",
	whatlanggo.Bul: "bg",
	whatlanggo.Ceb: "", // No iso 639-1 code
	whatlanggo.Ces: "cs",
	whatlanggo.Cmn: "zh", // No iso 639-1, but http://www.loc.gov/standards/iso639-2/faq.html#24
	whatlanggo.Dan: "da",
	whatlanggo.Deu: "de",
	whatlanggo.Ell: "el",
	whatlanggo.Eng: "en",
	whatlanggo.Epo: "eo",
	whatlanggo.Est: "et",
	whatlanggo.Fin: "fi",
	whatlanggo.Fra: "fr",
	whatlanggo.Guj: "gu",
	whatlanggo.Hat: "ht",
	whatlanggo.Hau: "ha",
	whatlanggo.Heb: "he",
	whatlanggo.Hin: "hi",
	whatlanggo.Hrv: "hr",
	whatlanggo.Hun: "hu",
	whatlanggo.Ibo: "ig",
	whatlanggo.Ilo: "", // No iso639-1
	whatlanggo.Ind: "id",
	whatlanggo.Ita: "it",
	whatlanggo.Jav: "jv",
	whatlanggo.Jpn: "ja",
	whatlanggo.Kan: "kn",
	whatlanggo.Kat: "ka",
	whatlanggo.Khm: "km",
	whatlanggo.Kin: "rw",
	whatlanggo.Kor: "ko",
	whatlanggo.Kur: "ku",
	whatlanggo.Lav: "lv",
	whatlanggo.Lit: "lt",
	whatlanggo.Mai: "", // No iso639-1
	whatlanggo.Mal: "ml",
	whatlanggo.Mar: "mr",
	whatlanggo.Mkd: "mk",
	whatlanggo.Mlg: "mg",
	whatlanggo.Mya: "my",
	whatlanggo.Nep: "ne",
	whatlanggo.Nld: "nl",
	whatlanggo.Nno: "nn",
	whatlanggo.Nob: "nb",
	whatlanggo.Nya: "ny",
	whatlanggo.Ori: "or",
	whatlanggo.Orm: "om",
	whatlanggo.Pan: "pa",
	whatlanggo.Pes: "", // No iso639-1
	whatlanggo.Pol: "pl",
	whatlanggo.Por: "pt",
	whatlanggo.Ron: "ro",
	whatlanggo.Run: "rn",
	whatlanggo.Rus: "ru",
	whatlanggo.Sin: "si",
	whatlanggo.Skr: "", // No iso639-1
	whatlanggo.Slv: "sl",
	whatlanggo.Sna: "sn",
	whatlanggo.Som: "so",
	whatlanggo.Spa: "es",
	whatlanggo.Srp: "sr",
	whatlanggo.Swe: "sv",
	whatlanggo.Tam: "ta",
	whatlanggo.Tel: "te",
	whatlanggo.Tgl: "tl",
	whatlanggo.Tha: "th",
	whatlanggo.Tir: "ti",
	whatlanggo.Tuk: "tk",
	whatlanggo.Tur: "tr",
	whatlanggo.Uig: "ug",
	whatlanggo.Ukr: "uk",
	whatlanggo.Urd: "ur",
	whatlanggo.Uzb: "uz",
	whatlanggo.Vie: "vi",
	whatlanggo.Ydd: "", // No iso639-1
	whatlanggo.Yor: "yo",
	whatlanggo.Zul: "zu",
}

func getWhiteListLang() map[whatlanggo.Lang]bool {

	whiteList := make(map[whatlanggo.Lang]bool)

	_ = EnvSetting.WhiteListLang

	for _, element := range EnvSetting.WhiteListLang {
		lang := whatlanggo.CodeToLang(element)

		whiteList[lang] = true
	}

	return whiteList
}

func DetectLang(text string) string {
	whiteList := getWhiteListLang()

	options := whatlanggo.Options{
		Whitelist: whiteList,
	}

	info := whatlanggo.DetectLangWithOptions(text, options)

	if val, ok := YandexLangMap[info]; ok {
		return val
	}
	return ""
}

func TargetLanguageCode() {

}
