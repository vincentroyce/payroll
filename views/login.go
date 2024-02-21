package views

import(
	"net/http"
	"github.com/uadmin/uadmin"
	"strings"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	type Context struct {
        Err         string
        ErrExists   bool
        OTPRequired bool
        Username    string
        Password    string
    }
    // Call the Context struct.
    c := Context{}

	if r.Method == "POST" {
		username := r.PostFormValue("user")
		username = strings.TrimSpace(strings.ToLower(username))
		password := r.PostFormValue("pass")
		session, _ := uadmin.Login(r, username, password)
		
		uadmin.Trail(uadmin.DEBUG, "Username: %s", username)
        uadmin.Trail(uadmin.DEBUG, "Password: %s", password)
		uadmin.Trail(uadmin.DEBUG, "Session: %s", session)

		if session == nil || !session.User.Active {
            /* Assign the login validation here that will be used for UI displaying. ErrExists and
            Err fields are coming from the Context struct. */
            c.ErrExists = true
            c.Err = "Invalid username/password or inactive user"

        } else {
            cookie, _ := r.Cookie("session")
            if cookie == nil {
                cookie = &http.Cookie{}
            }
            cookie.Name = "session"
            cookie.Value = session.Key
            cookie.Path = "/"
            cookie.SameSite = http.SameSiteStrictMode
            http.SetCookie(w, cookie)

            // If the next value is empty, redirect the page that omits the logout keyword in the last part.
            if r.URL.Query().Get("next") == "" {
                http.Redirect(w, r, strings.TrimSuffix(r.RequestURI, "logout"), http.StatusSeeOther)
                return
            }

            // Redirect to the page depending on the value of the next.
            http.Redirect(w, r, r.URL.Query().Get("next"), http.StatusSeeOther)
            return
        }
 	}
    // Render the login filepath and pass the context data object to the HTML file.
    uadmin.RenderHTML(w, r, "templates/login.html", c)
}