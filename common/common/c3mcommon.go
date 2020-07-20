package common

import (
	"archive/zip"
	"bufio"
	"bytes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"mime/multipart"
	"path"
	"reflect"

	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"colis/common/log"
	"colis/common/mycrypto"
	"colis/common/mystring"
	"colis/models"

	"github.com/nfnt/resize"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
)

var (
	db              map[string]*mgo.Database
	listCountryFlag map[string]string
	listLocale      map[string]string
	listCountry     map[string]string
)

func init() {
	fmt.Print("init common...")
	initListLocale()
	initCountryFlag()
	initListCountry()
	fmt.Print("done\n")
}
func initListLocale() {
	listLocale = make(map[string]string)
	listLocale["af"] = "af_ZA"
	listLocale["ak"] = "ak_GH"
	listLocale["am"] = "am_ET"
	listLocale["ar"] = "ar_AR"
	listLocale["as"] = "as_IN"
	listLocale["ay"] = "ay_BO"
	listLocale["az"] = "az_AZ"
	listLocale["be"] = "be_BY"
	listLocale["bg"] = "bg_BG"
	listLocale["bn"] = "bn_IN"
	listLocale["br"] = "br_FR"
	listLocale["bs"] = "bs_BA"
	listLocale["ca"] = "ca_ES"
	listLocale["cb"] = "cb_IQ"
	listLocale["ck"] = "ck_US"
	listLocale["co"] = "co_FR"
	listLocale["cs"] = "cs_CZ"
	listLocale["cx"] = "cx_PH"
	listLocale["cy"] = "cy_GB"
	listLocale["da"] = "da_DK"
	listLocale["de"] = "de_DE"
	listLocale["el"] = "el_GR"
	listLocale["en"] = "en_GB"
	listLocale["eo"] = "eo_EO"

	listLocale["es"] = "es_ES"

	listLocale["et"] = "et_EE"
	listLocale["eu"] = "eu_ES"
	listLocale["fa"] = "fa_IR"
	listLocale["fb"] = "fb_LT"
	listLocale["ff"] = "ff_NG"
	listLocale["fi"] = "fi_FI"
	listLocale["fo"] = "fo_FO"
	listLocale["fr"] = "fr_FR"
	listLocale["fy"] = "fy_NL"
	listLocale["ga"] = "ga_IE"
	listLocale["gl"] = "gl_ES"
	listLocale["gn"] = "gn_PY"
	listLocale["gu"] = "gu_IN"
	listLocale["gx"] = "gx_GR"
	listLocale["ha"] = "ha_NG"
	listLocale["he"] = "he_IL"
	listLocale["hi"] = "hi_IN"
	listLocale["hr"] = "hr_HR"
	listLocale["ht"] = "ht_HT"
	listLocale["hu"] = "hu_HU"
	listLocale["hy"] = "hy_AM"
	listLocale["id"] = "id_ID"
	listLocale["ig"] = "ig_NG"
	listLocale["is"] = "is_IS"
	listLocale["it"] = "it_IT"
	listLocale["ja"] = "ja_JP"
	listLocale["jv"] = "jv_ID"
	listLocale["ka"] = "ka_GE"
	listLocale["kk"] = "kk_KZ"
	listLocale["km"] = "km_KH"
	listLocale["kn"] = "kn_IN"
	listLocale["ko"] = "ko_KR"
	listLocale["ku"] = "ku_TR"
	listLocale["ky"] = "ky_KG"
	listLocale["la"] = "la_VA"
	listLocale["lg"] = "lg_UG"
	listLocale["li"] = "li_NL"
	listLocale["ln"] = "ln_CD"
	listLocale["lo"] = "lo_LA"
	listLocale["lt"] = "lt_LT"
	listLocale["lv"] = "lv_LV"
	listLocale["mg"] = "mg_MG"
	listLocale["mi"] = "mi_NZ"
	listLocale["mk"] = "mk_MK"
	listLocale["ml"] = "ml_IN"
	listLocale["mn"] = "mn_MN"
	listLocale["mr"] = "mr_IN"
	listLocale["ms"] = "ms_MY"
	listLocale["mt"] = "mt_MT"
	listLocale["my"] = "my_MM"
	listLocale["nb"] = "nb_NO"
	listLocale["nd"] = "nd_ZW"
	listLocale["ne"] = "ne_NP"

	listLocale["nl"] = "nl_NL"
	listLocale["nn"] = "nn_NO"
	listLocale["ny"] = "ny_MW"
	listLocale["or"] = "or_IN"
	listLocale["pa"] = "pa_IN"
	listLocale["pl"] = "pl_PL"
	listLocale["ps"] = "ps_AF"
	listLocale["pt"] = "pt_PT"
	listLocale["qc"] = "qc_GT"
	listLocale["qu"] = "qu_PE"
	listLocale["rm"] = "rm_CH"
	listLocale["ro"] = "ro_RO"
	listLocale["ru"] = "ru_RU"
	listLocale["rw"] = "rw_RW"
	listLocale["sa"] = "sa_IN"
	listLocale["sc"] = "sc_IT"
	listLocale["se"] = "se_NO"
	listLocale["si"] = "si_LK"
	listLocale["sk"] = "sk_SK"
	listLocale["sl"] = "sl_SI"
	listLocale["sn"] = "sn_ZW"
	listLocale["so"] = "so_SO"
	listLocale["sq"] = "sq_AL"
	listLocale["sr"] = "sr_RS"
	listLocale["sv"] = "sv_SE"
	listLocale["sw"] = "sw_KE"
	listLocale["sy"] = "sy_SY"
	listLocale["sz"] = "sz_PL"
	listLocale["ta"] = "ta_IN"
	listLocale["te"] = "te_IN"
	listLocale["tg"] = "tg_TJ"
	listLocale["th"] = "th_TH"
	listLocale["tk"] = "tk_TM"

	listLocale["tl"] = "tl_ST"
	listLocale["tr"] = "tr_TR"
	listLocale["tt"] = "tt_RU"
	listLocale["tz"] = "tz_MA"
	listLocale["uk"] = "uk_UA"
	listLocale["ur"] = "ur_PK"
	listLocale["uz"] = "uz_UZ"
	listLocale["vi"] = "vi_VN"
	listLocale["wo"] = "wo_SN"
	listLocale["xh"] = "xh_ZA"
	listLocale["yi"] = "yi_DE"
	listLocale["yo"] = "yo_NG"
	listLocale["zh"] = "zh_CN"
	listLocale["zu"] = "zu_ZA"
	listLocale["zz"] = "zz_TR"
}
func GetLangnameByCode(code string) string {
	return listCountry[code]
}
func initListCountry() {
	listCountry = make(map[string]string)
	listCountry["af"] = "Afrikaans"
	listCountry["ak"] = "Akan"
	listCountry["am"] = "Amharic"
	listCountry["ar"] = "Arabic"
	listCountry["as"] = "Assamese"
	listCountry["ay"] = "Aymara"
	listCountry["az"] = "Azerbaijani"
	listCountry["be"] = "Belarusian"
	listCountry["bg"] = "Bulgarian"
	listCountry["bn"] = "Bengali"
	listCountry["br"] = "Breton"
	listCountry["bs"] = "Bosnian"
	listCountry["ca"] = "Catalan"
	listCountry["cb"] = "Sorani Kurdish"
	listCountry["ck"] = "Cherokee"
	listCountry["co"] = "Corsican"
	listCountry["cs"] = "Czech"
	listCountry["cx"] = "Cebuano"
	listCountry["cy"] = "Welsh"
	listCountry["da"] = "Danish"
	listCountry["de"] = "German"
	listCountry["el"] = "Greek"
	listCountry["en"] = "English"
	listCountry["eo"] = "Esperanto"
	listCountry["es"] = "Spanish (Venezuela)"
	listCountry["et"] = "Estonian"
	listCountry["eu"] = "Basque"
	listCountry["fa"] = "Persian"
	listCountry["fb"] = "Leet Speak"
	listCountry["ff"] = "Fulah"
	listCountry["fi"] = "Finnish"
	listCountry["fo"] = "Faroese"
	listCountry["fr"] = "French"
	listCountry["fy"] = "Frisian"
	listCountry["ga"] = "Irish"
	listCountry["gl"] = "Galician"
	listCountry["gn"] = "Guarani"
	listCountry["gu"] = "Gujarati"
	listCountry["gx"] = "Classical Greek"
	listCountry["ha"] = "Hausa"
	listCountry["he"] = "Hebrew"
	listCountry["hi"] = "Hindi"
	listCountry["hr"] = "Croatian"
	listCountry["ht"] = "Haitian Creole"
	listCountry["hu"] = "Hungarian"
	listCountry["hy"] = "Armenian"
	listCountry["id"] = "Indonesian"
	listCountry["ig"] = "Igbo"
	listCountry["is"] = "Icelandic"
	listCountry["it"] = "Italian"
	listCountry["ja"] = "Japanese"
	listCountry["jv"] = "Javanese"
	listCountry["ka"] = "Georgian"
	listCountry["kk"] = "Kazakh"
	listCountry["km"] = "Khmer"
	listCountry["kn"] = "Kannada"
	listCountry["ko"] = "Korean"
	listCountry["ku"] = "Kurdish (Kurmanji)"
	listCountry["ky"] = "Kyrgyz"
	listCountry["la"] = "Latin"
	listCountry["lg"] = "Ganda"
	listCountry["li"] = "Limburgish"
	listCountry["ln"] = "Lingala"
	listCountry["lo"] = "Lao"
	listCountry["lt"] = "Lithuanian"
	listCountry["lv"] = "Latvian"
	listCountry["mg"] = "Malagasy"
	listCountry["mi"] = "Māori"
	listCountry["mk"] = "Macedonian"
	listCountry["ml"] = "Malayalam"
	listCountry["mn"] = "Mongolian"
	listCountry["mr"] = "Marathi"
	listCountry["ms"] = "Malay"
	listCountry["mt"] = "Maltese"
	listCountry["my"] = "Burmese"
	listCountry["nb"] = "Norwegian (bokmal)"
	listCountry["nd"] = "Ndebele"
	listCountry["ne"] = "Nepali"

	listCountry["nl"] = "Dutch"
	listCountry["nn"] = "Norwegian (nynorsk)"
	listCountry["ny"] = "Chewa"
	listCountry["or"] = "Oriya"
	listCountry["pa"] = "Punjabi"
	listCountry["pl"] = "Polish"
	listCountry["ps"] = "Pashto"
	listCountry["pt"] = "Portuguese (Brazil)"
	listCountry["qc"] = "Quiché"
	listCountry["qu"] = "Quechua"
	listCountry["rm"] = "Romansh"
	listCountry["ro"] = "Romanian"
	listCountry["ru"] = "Russian"
	listCountry["rw"] = "Kinyarwanda"
	listCountry["sa"] = "Sanskrit"
	listCountry["sc"] = "Sardinian"
	listCountry["se"] = "Northern Sámi"
	listCountry["si"] = "Sinhala"
	listCountry["sk"] = "Slovak"
	listCountry["sl"] = "Slovenian"
	listCountry["sn"] = "Shona"
	listCountry["so"] = "Somali"
	listCountry["sq"] = "Albanian"
	listCountry["sr"] = "Serbian"
	listCountry["sv"] = "Swedish"
	listCountry["sw"] = "Swahili"
	listCountry["sy"] = "Syriac"
	listCountry["sz"] = "Silesian"
	listCountry["ta"] = "Tamil"
	listCountry["te"] = "Telugu"
	listCountry["tg"] = "Tajik"
	listCountry["th"] = "Thai"
	listCountry["tk"] = "Turkmen"

	listCountry["tl"] = "Klingon"
	listCountry["tr"] = "Turkish"
	listCountry["tt"] = "Tatar"
	listCountry["tz"] = "Tamazight"
	listCountry["uk"] = "Ukrainian"
	listCountry["ur"] = "Urdu"
	listCountry["uz"] = "Uzbek"
	listCountry["vi"] = "Tiếng Việt"
	listCountry["wo"] = "Wolof"
	listCountry["xh"] = "Xhosa"
	listCountry["yi"] = "Yiddish"
	listCountry["yo"] = "Yoruba"
	listCountry["zh"] = "Simplified Chinese (China)"
	listCountry["zu"] = "Zulu"
	listCountry["zz"] = "Zazaki"
}
func initCountryFlag() {
	listCountryFlag = make(map[string]string)
	listCountryFlag["af"] = "za"
	listCountryFlag["ak"] = "gh"
	listCountryFlag["am"] = "et"
	listCountryFlag["ar"] = "ar"
	listCountryFlag["as"] = "in"
	listCountryFlag["ay"] = "bo"
	listCountryFlag["az"] = "az"
	listCountryFlag["be"] = "by"
	listCountryFlag["bg"] = "bg"
	listCountryFlag["bn"] = "in"
	listCountryFlag["br"] = "fr"
	listCountryFlag["bs"] = "ba"
	listCountryFlag["ca"] = "es"
	listCountryFlag["cb"] = "iq"
	listCountryFlag["ck"] = "us"
	listCountryFlag["co"] = "fr"
	listCountryFlag["cs"] = "cz"
	listCountryFlag["cx"] = "ph"
	listCountryFlag["cy"] = "gb"
	listCountryFlag["da"] = "dk"
	listCountryFlag["de"] = "de"
	listCountryFlag["el"] = "gr"
	listCountryFlag["en"] = "gb"
	listCountryFlag["eo"] = "eo"
	listCountryFlag["es"] = "cl"
	listCountryFlag["es"] = "es"
	listCountryFlag["et"] = "ee"
	listCountryFlag["eu"] = "es"
	listCountryFlag["fa"] = "ir"
	listCountryFlag["fb"] = "lt"
	listCountryFlag["ff"] = "ng"
	listCountryFlag["fi"] = "fi"
	listCountryFlag["fo"] = "fo"
	listCountryFlag["fr"] = "fr"
	listCountryFlag["fy"] = "nl"
	listCountryFlag["ga"] = "ie"
	listCountryFlag["gl"] = "es"
	listCountryFlag["gn"] = "py"
	listCountryFlag["gu"] = "in"
	listCountryFlag["gx"] = "gr"
	listCountryFlag["ha"] = "ng"
	listCountryFlag["he"] = "il"
	listCountryFlag["hi"] = "in"
	listCountryFlag["hr"] = "hr"
	listCountryFlag["ht"] = "ht"
	listCountryFlag["hu"] = "hu"
	listCountryFlag["hy"] = "am"
	listCountryFlag["id"] = "id"
	listCountryFlag["ig"] = "ng"
	listCountryFlag["is"] = "is"
	listCountryFlag["it"] = "it"
	listCountryFlag["ja"] = "jp"
	listCountryFlag["jv"] = "id"
	listCountryFlag["ka"] = "ge"
	listCountryFlag["kk"] = "kz"
	listCountryFlag["km"] = "kh"
	listCountryFlag["kn"] = "in"
	listCountryFlag["ko"] = "kr"
	listCountryFlag["ku"] = "tr"
	listCountryFlag["ky"] = "kg"
	listCountryFlag["la"] = "va"
	listCountryFlag["lg"] = "ug"
	listCountryFlag["li"] = "nl"
	listCountryFlag["ln"] = "cd"
	listCountryFlag["lo"] = "la"
	listCountryFlag["lt"] = "lt"
	listCountryFlag["lv"] = "lv"
	listCountryFlag["mg"] = "mg"
	listCountryFlag["mi"] = "nz"
	listCountryFlag["mk"] = "mk"
	listCountryFlag["ml"] = "in"
	listCountryFlag["mn"] = "mn"
	listCountryFlag["mr"] = "in"
	listCountryFlag["ms"] = "my"
	listCountryFlag["mt"] = "mt"
	listCountryFlag["my"] = "mm"
	listCountryFlag["nb"] = "no"
	listCountryFlag["nd"] = "zw"
	listCountryFlag["ne"] = "np"
	listCountryFlag["nl"] = "nl"
	listCountryFlag["nn"] = "no"
	listCountryFlag["ny"] = "mw"
	listCountryFlag["or"] = "in"
	listCountryFlag["pa"] = "in"
	listCountryFlag["pl"] = "pl"
	listCountryFlag["ps"] = "af"
	listCountryFlag["pt"] = "pt"
	listCountryFlag["qc"] = "gt"
	listCountryFlag["qu"] = "pe"
	listCountryFlag["rm"] = "ch"
	listCountryFlag["ro"] = "ro"
	listCountryFlag["ru"] = "ru"
	listCountryFlag["rw"] = "rw"
	listCountryFlag["sa"] = "in"
	listCountryFlag["sc"] = "it"
	listCountryFlag["se"] = "no"
	listCountryFlag["si"] = "lk"
	listCountryFlag["sk"] = "sk"
	listCountryFlag["sl"] = "si"
	listCountryFlag["sn"] = "zw"
	listCountryFlag["so"] = "so"
	listCountryFlag["sq"] = "al"
	listCountryFlag["sr"] = "rs"
	listCountryFlag["sv"] = "se"
	listCountryFlag["sw"] = "ke"
	listCountryFlag["sy"] = "sy"
	listCountryFlag["sz"] = "pl"
	listCountryFlag["ta"] = "in"
	listCountryFlag["te"] = "in"
	listCountryFlag["tg"] = "tj"
	listCountryFlag["th"] = "th"
	listCountryFlag["tk"] = "tm"
	listCountryFlag["tl"] = "st"
	listCountryFlag["tr"] = "tr"
	listCountryFlag["tt"] = "ru"
	listCountryFlag["tz"] = "ma"
	listCountryFlag["uk"] = "ua"
	listCountryFlag["ur"] = "pk"
	listCountryFlag["uz"] = "uz"
	listCountryFlag["vi"] = "vn"
	listCountryFlag["wo"] = "sn"
	listCountryFlag["xh"] = "za"
	listCountryFlag["yi"] = "de"
	listCountryFlag["yo"] = "ng"
	listCountryFlag["zh"] = "cn"
	listCountryFlag["zu"] = "za"
	listCountryFlag["zz"] = "tr"
}

func ConnectDB(dbname string) (db *mgo.Database, strErr string) {
	//get database config from ENV
	hosts := strings.Trim(os.Getenv(strings.ToUpper(dbname)+"_DB_HOST"), " ")
	name := strings.Trim(os.Getenv(strings.ToUpper(dbname)+"_DB_NAME"), " ")
	user := strings.Trim(os.Getenv(strings.ToUpper(dbname)+"_DB_USER"), " ")
	pass := strings.Trim(os.Getenv(strings.ToUpper(dbname)+"_DB_PASS"), " ")
	if hosts == "" || name == "" || user == "" || pass == "" {
		strErr = "Missing config data for database connection"
		return db, strErr
	}

	//call to connection
	mongoDBDialInfo := mgo.DialInfo{
		Addrs:    strings.Split(hosts, ","),
		Timeout:  60 * time.Second,
		Database: name,
		Username: user,
		Password: pass,
	}

	mongoSession, err := mgo.DialWithInfo(&mongoDBDialInfo)

	if CheckError("error when connect db "+name+" with user "+user+" and pass "+pass+" on "+hosts, err) {
		mongoSession.SetMode(mgo.Monotonic, true)
		db = mongoSession.DB(name)
	}
	if err != nil {
		strErr = err.Error()
	}
	return db, strErr
}
func GetSpecialChar() string {
	return `.*?/\n~!@#$%^&*(),.[];'<>"` + "`"
}
func ReturnJsonMessage(status, strerr, strmsg, data string) models.RequestResult {
	if data == "" {
		data = "{}"
	}
	var resp models.RequestResult
	resp.Status = status
	resp.Error = strerr
	resp.Message = strmsg

	resp.Data = data
	return resp
}
func FileCount(path string) int {
	i := 0
	files, err := ioutil.ReadDir(path)
	if err != nil {
		CheckError(path+" filecount error", err)
		return 0
	}
	for _, file := range files {
		if file.IsDir() {
			i += FileCount(path + "/" + file.Name())
		} else {
			i++
		}

	}
	return i
}

//func CheckRequest(uri, useragent, referrer, remoteAddress string) bool {

//	col := db["session"].C("requestUrls")
//	log.Printf("now: %d , check: %d", int(time.Now().Unix()), int(time.Now().Unix())-10)
//	urlcount, err := col.Find(bson.M{"uri": uri, "created": bson.M{"$gt": int(time.Now().Unix()) - 1}}).Count()
//	if CheckError("checkRequest", err) {
//		if urlcount == 0 {
//			//check ip in 3 sec
//			urlcount, err := col.Find(bson.M{"remoteAddress": remoteAddress, "created": bson.M{"$gt": int(time.Now().Unix()) - 3}}).Count()
//			if urlcount < 50 {
//				err = col.Insert(bson.M{"uri": uri, "created": int(time.Now().Unix()), "user-agent": useragent, "referer": referrer, "remoteAddress": remoteAddress})
//				CheckError("checkRequest Insert", err)
//				return true
//			}

//		}
//	}
//	return false
//}

//true: no error
func CheckError(msg string, err error) bool {
	if err != nil {
		log.Debugf(msg+": %s", err.Error())
		return false
	}
	return true
}
func InArray(v interface{}, in interface{}) (ok bool, i int) {
	val := reflect.Indirect(reflect.ValueOf(in))
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for ; i < val.Len(); i++ {
			if ok = v == val.Index(i).Interface(); ok {
				return
			}
		}
	}
	return
}
func ImgResize(imagebytes []byte, w, h uint) ([]byte, string) {
	filetype := http.DetectContentType(imagebytes[:512])
	r := bytes.NewReader(imagebytes)
	imagecontent, _, err := image.Decode(r)
	m := resize.Resize(w, h, imagecontent, resize.NearestNeighbor)
	if err != nil {
		return nil, ""
	}
	var buf bytes.Buffer
	wr := io.Writer(&buf)
	returnext := "jpg"
	if filetype == "image/jpeg" {
		jpeg.Encode(wr, m, nil)
	} else if filetype == "image/gif" {
		gif.Encode(wr, m, nil)
		returnext = "gif"
	} else if filetype == "image/png" {
		png.Encode(wr, m)
		returnext = "png"
	} else {
		returnext = filetype
	}

	return buf.Bytes(), returnext
}

func CheckDomain(requestDomain string) string {

	domainallow := strings.Split(viper.GetString("config.domainallow"), ",")
	requestDomain = strings.Replace(requestDomain, "http://", "", -1)
	requestDomain = strings.Replace(requestDomain, "https://", "", -1)
	requestDomain = strings.Replace(requestDomain, "/", "", -1)

	for i := 0; i < len(domainallow); i++ {
		log.Debugf("%s - %s", domainallow[i], requestDomain)
		if domainallow[i] == requestDomain {
			return requestDomain
			break
		}
	}
	return ""
}

func Fake64() string {
	return mystring.RandString(100)
}

func Code2Flag(code string) string {
	return listCountryFlag[code]
}

// Unzip will decompress a zip archive, moving all files and folders
// within the zip file (parameter 1) to an output directory (parameter 2).
func Unzip(src string, dest string) ([]string, error) {

	var filenames []string

	r, err := zip.OpenReader(src)
	if err != nil {
		return filenames, errors.New("open " + src + ": " + err.Error())
	}
	defer r.Close()
	filetypeAllows := strings.Split(viper.GetString("config.zipfileextallow"), ",")
	filetypeAllowMap := make(map[string]string)
	for _, ext := range filetypeAllows {
		filetypeAllowMap[ext] = ext
	}

	for _, f := range r.File {
		//check file type:
		if filetypeAllowMap[path.Ext(f.Name)] == "" {
			continue
		}
		// Store filename/path for returning and using later on
		fpath := filepath.Join(dest, f.Name)

		// Check for ZipSlip. More Info: http://bit.ly/2MsjAWE
		if !strings.HasPrefix(fpath, filepath.Clean(dest)+string(os.PathSeparator)) {
			return filenames, fmt.Errorf("%s: illegal file path", fpath)
		}

		filenames = append(filenames, fpath)

		if f.FileInfo().IsDir() {
			// Make Folder
			os.MkdirAll(fpath, os.ModePerm)
			continue
		}

		// Make File
		if err = os.MkdirAll(filepath.Dir(fpath), os.ModePerm); err != nil {
			return filenames, errors.New("create " + strings.Replace(fpath, dest, "", 1) + ": " + err.Error())
		}

		outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
		if err != nil {
			return filenames, errors.New("file " + strings.Replace(outFile.Name(), dest, "", 1) + ": " + err.Error())
		}

		rc, err := f.Open()
		if err != nil {
			return filenames, errors.New("file " + strings.Replace(outFile.Name(), dest, "", 1) + ": " + err.Error())
		}

		_, err = io.Copy(outFile, rc)

		// Close the file without defer to close before next iteration of loop
		outFile.Close()
		rc.Close()

		if err != nil {
			return filenames, errors.New("file " + strings.Replace(outFile.Name(), dest, "", 1) + ": " + err.Error())
		}
	}
	return filenames, nil
}

func Zipit(source, target string) error {
	zipfile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	info, err := os.Stat(source)
	if err != nil {
		return nil
	}

	var baseDir string
	if info.IsDir() {
		baseDir = filepath.Base(source)
	}
	//filetype allow:
	filetypeAllows := strings.Split(viper.GetString("config.zipfileextallow"), ",")
	filetypeAllowMap := make(map[string]string)
	for _, ext := range filetypeAllows {
		filetypeAllowMap[ext] = ext
	}
	filepath.Walk(source, func(walkpath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		//check file type:
		if filetypeAllowMap[path.Ext(info.Name())] == "" {
			return nil
		}

		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		if baseDir != "" {
			//header.Name = filepath.Join(baseDir, strings.TrimPrefix(walkpath, source))
			header.Name = strings.TrimPrefix(walkpath, source)
			//remove slash
			header.Name = header.Name[1:]
		}

		if info.IsDir() {
			header.Name += "/"
		} else {
			header.Method = zip.Deflate
		}
		log.Debugf("basedir:%s, walkpath:%s, file: %s, headername:%s", baseDir, walkpath, info.Name(), header.Name)
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		file, err := os.Open(walkpath)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = io.Copy(writer, file)
		return err
	})

	return err
}

// Check if a port is available
func CheckPort(port int) (status bool) {

	// Concatenate a colon and the port
	host := ":" + strconv.Itoa(port)

	// Try to create a server with the port
	server, err := net.Listen("tcp", host)

	// if it fails then the port is likely taken
	if err != nil {
		return false
	}

	// close the server
	server.Close()

	// we successfully used and closed the port
	// so it's now available to be used again
	return true
}

func FolderExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

//alternate code: https://github.com/otiai10/copy/blob/master/copy.go
func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}

func CopyDir(source string, dest string) (err error) {

	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform copy
			err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

func JS1Line(content string) string {
	content = strings.Replace(content, "\r\n", "", -1)
	content = strings.Replace(content, "\n", "", -1)
	content = strings.Replace(content, "\r", "", -1)
	return content
}
func JSMinify(content string) string {
	data := "code=" + viper.GetString("config.minifyKey")
	data += "&text=" + url.QueryEscape(content)

	rtstr, _ := RequestUrl(viper.GetString("config.minify"), "POST", data)
	return rtstr
}

//commpress using lzstring from nodejs
func Base64Compress(content string) string {
	data := "code=" + viper.GetString("config.minifyKey")
	data += "&text=" + url.QueryEscape(content)
	rtstr, _ := RequestUrl(viper.GetString("config.compress"), "POST", data)
	return rtstr
}
func MinifyCompress(content string) string {
	data := "code=" + viper.GetString("config.minifyKey")
	data += "&text=" + url.QueryEscape(content)

	rtstr, _ := RequestUrl(viper.GetString("config.minifycompress"), "POST", data)
	return rtstr
}
func RequestMainService(uri, method, data string) models.RequestResult {
	data = "data=" + mycrypto.EncDat2(data)
	rtstr, resp := RequestUrl(os.Getenv("MAIN_SERVER")+mycrypto.EncDat2(uri), method, data)
	var rs models.RequestResult
	if resp == nil || resp.StatusCode != 200 {
		rs.Status = "0"
		rs.Error = "Request service error. Please contact your administrator."
		return rs
	}
	rtstr = mycrypto.DecodeOld(rtstr, 8)

	json.Unmarshal([]byte(rtstr), &rs)

	if rs.Status == "" {
		rs.Status = "0"
		rs.Error = "Service Response error. Please contact your administrator."
	}
	return rs
}
func RequestBuildService(uri, method, data string) models.RequestResult {

	data = "data=" + mycrypto.EncDat2(data)
	rtstr, resp := RequestUrl(viper.GetString("config.buildserver")+mycrypto.EncodeBK(uri, "name"), method, data)
	var rs models.RequestResult
	if resp == nil || resp.StatusCode != 200 {
		rs.Status = "0"
		rs.Error = "Request service error. Please contact your administrator."

	}
	rtstr = mycrypto.DecodeLight1(rtstr, 5)

	json.Unmarshal([]byte(rtstr), &rs)

	if rs.Status == "" {
		rs.Status = "0"
		rs.Error = "Service Response error. Please contact your administrator."
	}

	return rs
}
func RequestBuildServiceAsync(uri, method, data string) {
	go func(uri, method, data string) {
		RequestBuildService(uri, method, data)
	}(uri, method, data)
}
func RequestUrl(urlrequest, method, data string) (string, *http.Response) {
	var req *http.Request
	var err error
	// payloadBytes, err := json.Marshal(data)
	body := bytes.NewReader([]byte(data))
	if strings.ToLower(method) == "post" {
		if viper.GetString("config.proxy") != "" {

			proxyUrl, _ := url.Parse(viper.GetString("config.proxy"))
			http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
		}

		req, err = http.NewRequest("POST", urlrequest, body)
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		host := viper.GetString("config.hostname")
		req.Header.Set("Origin", host)

		// host = strings.Replace(host, "http://", "", -1)
		// host = strings.Replace(host, "https://", "", -1)
		// host = strings.Replace(host, "/", "", -1)
		// req.Host = host

	} else {
		req, err = http.NewRequest("GET", urlrequest+"?"+data, nil)

	}

	resp, err := http.DefaultClient.Do(req)
	if !CheckError("request api", err) {
		return "", resp
	}

	defer req.Body.Close()

	bodyresp, err := ioutil.ReadAll(resp.Body)
	bodystr := string(bodyresp)
	CheckError("request read data", err)

	return bodystr, resp
}
func RequestUrl2(urlrequest, method string, data url.Values) string {
	var rsp *http.Response
	var err error

	if strings.ToLower(method) == "post" {
		if viper.GetString("config.proxy") != "" {
			proxyUrl, _ := url.Parse(viper.GetString("config.proxy"))
			http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
		}
		rsp, err = http.PostForm(urlrequest, data)
		rsp.Header.Set("Origin", "application/json")

		if !CheckError("request api", err) {
			return ""
		}
	} else {
		rsp, err = http.Get(urlrequest + "?" + data.Encode())
		if !CheckError("request api", err) {
			return ""
		}
	}

	defer rsp.Body.Close()
	rtbyte, err := ioutil.ReadAll(rsp.Body)
	CheckError("request read data", err)
	rtstr := string(rtbyte)
	return rtstr
}

// Creates a new file upload http request with optional extra params
func FileUploadRequest(uri string, params map[string]string, paramName, path string) string {
	rt := ""
	file, err := os.Open(path)
	if err != nil {
		return rt
	}
	defer file.Close()

	// fileContents, err := ioutil.ReadAll(file)
	// if err != nil {
	// 	return rt
	// }
	fi, err := file.Stat()
	if err != nil {
		return rt
	}

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, fi.Name())
	if err != nil {
		return rt
	}
	io.Copy(part, file)

	for key, val := range params {
		_ = writer.WriteField(key, val)
	}
	err = writer.Close()
	if err != nil {
		return rt
	}
	if viper.GetString("config.proxy") != "" {
		proxyUrl, _ := url.Parse(viper.GetString("config.proxy"))
		http.DefaultTransport = &http.Transport{Proxy: http.ProxyURL(proxyUrl)}
	}
	request, err := http.NewRequest("POST", uri, body)
	request.Header.Add("Content-Type", writer.FormDataContentType())
	if err != nil {
		CheckError("Request "+uri+" error", err)
	}
	//client := &http.Client{}
	host := viper.GetString("config.hostname")
	request.Header.Set("Origin", host)

	resp, err := http.DefaultClient.Do(request)
	bodyresp, _ := ioutil.ReadAll(resp.Body)
	rt = string(bodyresp)
	// if err != nil {
	// 	CheckError("Request "+uri+" error", err)
	// } else {
	// 	var bodyContent []byte
	// 	resp.Body.Read(bodyContent)
	// 	resp.Body.Close()
	// 	rt = string(bodyContent)
	// }
	return rt
}

func RequestService2(serviceurl string, data url.Values) string {
	payloadBytes, err := json.Marshal(data)
	body := bytes.NewReader(payloadBytes)
	req, err := http.NewRequest("POST", serviceurl, body)
	if !CheckError("request api", err) {
		return ""
	}
	//req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if !CheckError("request api", err) {
		return ""
	}

	defer resp.Body.Close()

	bodyresp, _ := ioutil.ReadAll(resp.Body)
	bodystr := string(bodyresp)

	if bodystr == "" {
		return ""
	}

	bodystr = mycrypto.Decode4(bodystr)
	// log.Debugf("response decode:%s", bodystr)
	return bodystr
}

func RemoveHTMLComments(content []byte) []byte {
	// https://www.google.com/search?q=regex+html+comments
	// http://stackoverflow.com/a/1084759
	htmlcmt := regexp.MustCompile(`<!--[^>]*-->`)
	return htmlcmt.ReplaceAll(content, []byte(""))
}

func MinifyHTML(html []byte) string {
	// read line by line
	minifiedHTML := ""
	scanner := bufio.NewScanner(bytes.NewReader(RemoveHTMLComments(html)))
	for scanner.Scan() {
		// all leading and trailing white space of each line are removed
		lineTrimmed := strings.TrimSpace(scanner.Text())
		minifiedHTML += lineTrimmed
		if len(lineTrimmed) > 0 {
			// in case of following trimmed line:
			// <div id="foo"
			minifiedHTML += " "
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return minifiedHTML
}

//minify css
func RemoveCStyleComments(content []byte) []byte {
	// http://blog.ostermiller.org/find-comment
	ccmt := regexp.MustCompile(`/\*([^*]|[\r\n]|(\*+([^*/]|[\r\n])))*\*+/`)
	return ccmt.ReplaceAll(content, []byte(""))
}

func RemoveCppStyleComments(content []byte) []byte {
	cppcmt := regexp.MustCompile(`//.*`)
	return cppcmt.ReplaceAll(content, []byte(""))
}

func MinifyCSS(csscontent []byte) string {

	cssAllNoComments := RemoveCStyleComments(csscontent)

	// read line by line
	minifiedCss := ""
	scanner := bufio.NewScanner(bytes.NewReader(cssAllNoComments))
	for scanner.Scan() {
		// all leading and trailing white space of each line are removed
		minifiedCss += strings.TrimSpace(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return minifiedCss
}
func CreateImageFile(path, b64content string) error {
	imagebytes, err := base64.StdEncoding.DecodeString(b64content)
	if len(imagebytes) < 512 {
		return errors.New("not image file: " + string(imagebytes))
	}
	filetype := http.DetectContentType(imagebytes[:512])
	r := bytes.NewReader(imagebytes)
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		errstr := "Cannot open file: " + path + " - " + err.Error()
		return errors.New(errstr)

	}

	if filetype == "image/jpeg" {
		im, err := jpeg.Decode(r)
		if err != nil {

			errstr := "Bad jpg: " + path + " - " + err.Error()
			return errors.New(errstr)
		}
		jpeg.Encode(f, im, nil)
	} else if filetype == "image/gif" {
		im, err := gif.Decode(r)
		if err != nil {
			errstr := "Bad gif: " + path + " - " + err.Error()
			return errors.New(errstr)
		}
		gif.Encode(f, im, nil)

	} else if filetype == "image/png" {
		im, err := png.Decode(r)
		if err != nil {
			errstr := "Bad png: " + path + " - " + err.Error()
			return errors.New(errstr)
		}
		png.Encode(f, im)

	}

	defer f.Close()

	return nil
}
func CreateImageFileOld(path, b64content string) error {
	unbased, err := base64.StdEncoding.DecodeString(b64content)
	if err != nil {
		log.Debugf("Cannot decode b64  %s", err)
		return err
	}
	r := bytes.NewReader(unbased)

	if filepath.Ext(path) == ".png" {
		im, err := png.Decode(r)
		if err != nil {
			errstr := "Bad png: " + path + " - " + err.Error()
			return errors.New(errstr)
		}

		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			errstr := "Cannot open file: " + path + " - " + err.Error()
			return errors.New(errstr)

		}
		png.Encode(f, im)
		f.Close()
	} else if filepath.Ext(path) == ".jpg" || filepath.Ext(path) == ".jpeg" {
		im, err := jpeg.Decode(r)
		if err != nil {

			errstr := "Bad jpg: " + path + " - " + err.Error()
			return errors.New(errstr)
		}

		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {
			errstr := "Cannot open file: " + path + " - " + err.Error()
			return errors.New(errstr)
		}
		jpeg.Encode(f, im, nil)
		f.Close()
	} else if filepath.Ext(path) == ".gif" {
		im, err := gif.Decode(r)
		if err != nil {
			errstr := "Bad gif: " + path + " - " + err.Error()
			return errors.New(errstr)
		}

		f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
		if err != nil {

			errstr := "Cannot open file: " + path + " - " + err.Error()
			return errors.New(errstr)
		}
		gif.Encode(f, im, nil)
		f.Close()
	} else if filepath.Ext(path) == ".svg" {
		i := strings.Index(b64content, ",")
		if i < 0 {
			log.Errorf("no comma")
		}
		// pass reader to NewDecoder
		dec := base64.NewDecoder(base64.StdEncoding, strings.NewReader(b64content[i+1:]))
		b, err := ioutil.ReadAll(dec)
		err = ioutil.WriteFile(path, b, 0777)
		if err != nil {
			errstr := "Bad svg: " + path + " - " + err.Error()
			return errors.New(errstr)
		}

	}
	return nil
}

const (
	encodePath encoding = 1 + iota
	encodeHost
	encodeUserPassword
	encodeQueryComponent
	encodeFragment
)

type encoding int
type EscapeError string

func (e EscapeError) Error() string {
	return "invalid URL escape " + strconv.Quote(string(e))
}

func ishex(c byte) bool {
	switch {
	case '0' <= c && c <= '9':
		return true
	case 'a' <= c && c <= 'f':
		return true
	case 'A' <= c && c <= 'F':
		return true
	}
	return false
}

func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

// Return true if the specified character should be escaped when
// appearing in a URL string, according to RFC 3986.
//
// Please be informed that for now shouldEscape does not check all
// reserved characters correctly. See golang.org/issue/5684.
func shouldEscape(c byte, mode encoding) bool {
	// §2.3 Unreserved characters (alphanum)
	if 'A' <= c && c <= 'Z' || 'a' <= c && c <= 'z' || '0' <= c && c <= '9' {
		return false
	}

	if mode == encodeHost {
		// §3.2.2 Host allows
		//  sub-delims = "!" / "$" / "&" / "'" / "(" / ")" / "*" / "+" / "," / ";" / "="
		// as part of reg-name.
		// We add : because we include :port as part of host.
		// We add [ ] because we include [ipv6]:port as part of host
		switch c {
		case '!', '$', '&', '\'', '(', ')', '*', '+', ',', ';', '=', ':', '[', ']':
			return false
		}
	}

	switch c {
	case '-', '_', '.', '~': // §2.3 Unreserved characters (mark)
		return false

	case '$', '&', '+', ',', '/', ':', ';', '=', '?', '@': // §2.2 Reserved characters (reserved)
		// Different sections of the URL allow a few of
		// the reserved characters to appear unescaped.
		switch mode {
		case encodePath: // §3.3
			// The RFC allows : @ & = + $ but saves / ; , for assigning
			// meaning to individual path segments. This package
			// only manipulates the path as a whole, so we allow those
			// last two as well. That leaves only ? to escape.
			return c == '?'

		case encodeUserPassword: // §3.2.1
			// The RFC allows ';', ':', '&', '=', '+', '$', and ',' in
			// userinfo, so we must escape only '@', '/', and '?'.
			// The parsing of userinfo treats ':' as special so we must escape
			// that too.
			return c == '@' || c == '/' || c == '?' || c == ':'

		case encodeQueryComponent: // §3.4
			// The RFC reserves (so we must escape) everything.
			return true

		case encodeFragment: // §4.1
			// The RFC text is silent but the grammar allows
			// everything, so escape nothing.
			return false
		}
	}

	// Everything else must be escaped.
	return true
}

func escape(s string, mode encoding) string {
	spaceCount, hexCount := 0, 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if shouldEscape(c, mode) {
			if c == ' ' && mode == encodeQueryComponent {
				spaceCount++
			} else {
				hexCount++
			}
		}
	}

	if spaceCount == 0 && hexCount == 0 {
		return s
	}

	t := make([]byte, len(s)+2*hexCount)
	j := 0
	for i := 0; i < len(s); i++ {
		switch c := s[i]; {
		case c == ' ' && mode == encodeQueryComponent:
			t[j] = '+'
			j++
		case shouldEscape(c, mode):
			t[j] = '%'
			t[j+1] = "0123456789ABCDEF"[c>>4]
			t[j+2] = "0123456789ABCDEF"[c&15]
			j += 3
		default:
			t[j] = s[i]
			j++
		}
	}
	return string(t)
}

// unescape unescapes a string; the mode specifies
// which section of the URL string is being unescaped.
func unescape(s string, mode encoding) (string, error) {
	// Count %, check that they're well-formed.
	n := 0
	hasPlus := false
	for i := 0; i < len(s); {
		switch s[i] {
		case '%':
			n++
			if i+2 >= len(s) || !ishex(s[i+1]) || !ishex(s[i+2]) {
				s = s[i:]
				if len(s) > 3 {
					s = s[:3]
				}
				return "", EscapeError(s)
			}
			i += 3
		case '+':
			hasPlus = mode == encodeQueryComponent
			i++
		default:
			i++
		}
	}

	if n == 0 && !hasPlus {
		return s, nil
	}

	t := make([]byte, len(s)-2*n)
	j := 0
	for i := 0; i < len(s); {
		switch s[i] {
		case '%':
			t[j] = unhex(s[i+1])<<4 | unhex(s[i+2])
			j++
			i += 3
		case '+':
			if mode == encodeQueryComponent {
				t[j] = ' '
			} else {
				t[j] = '+'
			}
			j++
			i++
		default:
			t[j] = s[i]
			j++
			i++
		}
	}
	return string(t), nil
}

func EncodeUriComponent(rawString string) string {
	return escape(rawString, encodeFragment)
}

func DecodeUriCompontent(encoded string) (string, error) {
	return unescape(encoded, encodeQueryComponent)
}
