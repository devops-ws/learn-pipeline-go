package auth

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/go-session/session"
	"golang.org/x/oauth2"
)

type Auth struct {
	homePage string
	config   oauth2.Config
}

// NewAuth creates a new auth handler
func NewAuth(config oauth2.Config) *Auth {
	return &Auth{
		homePage: "/",
		config:   config,
	}
}

func (a *Auth) LoginHandler(w http.ResponseWriter, r *http.Request) {
	store, err := session.Start(r.Context(), w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if r.Method == http.MethodPost {
		if r.Form == nil {
			if err := r.ParseForm(); err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
		}
		username := r.Form.Get("username")
		fmt.Println("username", username)
		store.Set("LoggedInUserID", username)
		store.Save()

		w.Header().Set("Location", "/auth")
		w.WriteHeader(http.StatusFound)
		return
	}
	outputHTML(w, r, "static/login.html")
}

func (a *Auth) AuthHandler(w http.ResponseWriter, r *http.Request) {
	store, err := session.Start(nil, w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if _, ok := store.Get("LoggedInUserID"); !ok {
		w.Header().Set("Location", "/login")
		w.WriteHeader(http.StatusFound)
		return
	}

	outputHTML(w, r, "static/auth.html")
}

func (a *Auth) Callback(w http.ResponseWriter, r *http.Request) {
	// r.ParseForm()
	// state := r.Form.Get("state")
	// if state != "xyz" {
	// 	http.Error(w, "State invalid", http.StatusBadRequest)
	// 	return
	// }
	// code := r.Form.Get("code")
	// if code == "" {
	// 	http.Error(w, "Code not found", http.StatusBadRequest)
	// 	return
	// }
	// token, err := config.Exchange(r.Context(), code, oauth2.SetAuthURLParam("code_verifier", "s256example"))
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	// // get userInfo, save it to session
	// if userInfo, err := oauth.GetUserInfo("authServerURL", token.AccessToken); err == nil {
	// 	store, err := session.Start(r.Context(), w, r)
	// 	if err != nil {
	// 		http.Error(w, err.Error(), http.StatusInternalServerError)
	// 		return
	// 	}
	// 	store.Set("userinfo", userInfo)
	// 	store.Save()
	// }

	// http.Redirect(w, r, a.homePage, http.StatusFound)
}

func (a *Auth) RequestCode(w http.ResponseWriter, r *http.Request) {
	u := a.config.AuthCodeURL("xyz",
		oauth2.SetAuthURLParam("code_challenge", `genCodeChallengeS256("s256example")`),
		oauth2.SetAuthURLParam("code_challenge_method", "S256"))
	http.Redirect(w, r, u, http.StatusFound)
}

func outputHTML(w http.ResponseWriter, req *http.Request, filename string) {
	absPath, _ := filepath.Abs(filename)
	file, err := os.Open(absPath)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer file.Close()
	fi, _ := file.Stat()
	http.ServeContent(w, req, file.Name(), fi.ModTime(), file)
}
