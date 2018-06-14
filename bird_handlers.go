package main
import(
	"fmt"
	"net/http"
	"encoding/json"

	//"github.com/gorilla/mux"
)


type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

//var birds []Bird
func getBirdIndex(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/assets/index.html", http.StatusFound)
}

func createBirdPage(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/assets/create.html", http.StatusFound)
}


func getBirdHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("get birds call\n")
	birds, err := store.GetBirds()//new
	
	birdListBytes, err := json.Marshal(birds)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(birdListBytes)
}

func createBirdHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("create bird call\n")
	bird := Bird{}
	err := r.ParseForm()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	//birds = append(birds, bird)
	err = store.CreateBird(&bird)

	http.Redirect(w, r, "/assets/", http.StatusFound)
}