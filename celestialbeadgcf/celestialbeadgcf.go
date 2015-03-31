package celestialbeadgcf

import (
	"html/template"
	"net/http"
)

func init() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/aquarius", aquariusHandler)
	http.HandleFunc("/aries", ariesHandler)
	http.HandleFunc("/cancer", cancerHandler)
	http.HandleFunc("/capricorn", capricornHandler)
	http.HandleFunc("/gemini", geminiHandler)
	http.HandleFunc("/leo", leoHandler)
	http.HandleFunc("/libra", libraHandler)
	http.HandleFunc("/lyra", lyraHandler)
	http.HandleFunc("/pisces", piscesHandler)
	http.HandleFunc("/sagittarius", sagittariusHandler)
	http.HandleFunc("/scorpio", scorpioHandler)
	http.HandleFunc("/taurus", taurusHandler)
	http.HandleFunc("/virgo", virgoHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("index").ParseFiles(
	"celestialbeadgcf/base.html", "celestialbeadgcf/index.html"))
	tmpl.ExecuteTemplate(w, "base", "Ursa Minor")
}

func aquariusHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("aquarius").ParseFiles(
	"celestialbeadgcf/base.html", "celestialbeadgcf/aquarius/sadalmelik.html"))
	tmpl.ExecuteTemplate(w, "base", "Aquarius")
}

func ariesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("aries").ParseFiles(
	"celestialbeadgcf/base.html", "celestialbeadgcf/aries/hamal.html"))
	tmpl.ExecuteTemplate(w, "base", "Aries")
}

func cancerHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("cancer").ParseFiles(
	"celestialbeadgcf/base.html", "celestialbeadgcf/cancer/acubens.html"))
	tmpl.ExecuteTemplate(w, "base", "Cancer")
}

func capricornHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("capricorn").ParseFiles(
	"celestialbeadgcf/base.html", "celestialbeadgcf/capricorn/algedi.html"))
	tmpl.ExecuteTemplate(w, "base", "Capricorn")
}

func geminiHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("castorpollux").ParseFiles(
	"celestialbeadgcf/base.html", "celestialbeadgcf/gemini/castorpollux.html"))
	tmpl.ExecuteTemplate(w, "base", "Gemini")
}

func leoHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("leo").ParseFiles(
	"celestialbeadgcf/base.html", "celestialbeadgcf/leo/regulus.html"))
	tmpl.ExecuteTemplate(w, "base", "Leo")
}

func libraHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("libra").ParseFiles(
	"celestialbeadgcf/base.html", "celestialbeadgcf/libra/zubenelgenubi.html"))
	tmpl.ExecuteTemplate(w, "base", "Libra")
}

func lyraHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("lyra").ParseFiles(
	"celestialbeadgcf/base.html", "celestialbeadgcf/lyra/vega.html"))
	tmpl.ExecuteTemplate(w, "base", "Lyra")
}

func piscesHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("pisces").ParseFiles(
	"celestialbeadgcf/base.html", "celestialbeadgcf/pisces/alpherg.html"))
	tmpl.ExecuteTemplate(w, "base", "Pisces")
}

func sagittariusHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("sagittarius").ParseFiles(
	"celestialbeadgcf/base.html", "celestialbeadgcf/sagittarius/rukbat.html"))
	tmpl.ExecuteTemplate(w, "base", "Sagittarius")
}

func scorpioHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("scorpio").ParseFiles(
	"celestialbeadgcf/base.html", "celestialbeadgcf/scorpio/antares.html"))
	tmpl.ExecuteTemplate(w, "base", "Scorpio")
}

func taurusHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("taurus").ParseFiles(
	"celestialbeadgcf/base.html", "celestialbeadgcf/taurus/aldebaran.html"))
	tmpl.ExecuteTemplate(w, "base", "Taurus")
}

func virgoHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.New("virgo").ParseFiles(
	"celestialbeadgcf/base.html", "celestialbeadgcf/virgo/spica.html"))
	tmpl.ExecuteTemplate(w, "base", "Virgo")
}

