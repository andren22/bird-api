package main
import(
	"fmt"
	"net/http"
	"encoding/json"
	//"io/ioutil"
	//"io"

	//"github.com/gorilla/mux"
)


type Bird struct {
	Species     string `json:"species"`
	Description string `json:"description"`
}

type Dog struct {
	Name string `json:"name"`
	Vacinnes []string `json:"vacinnes"`
}

func getDogHandler(w http.ResponseWriter, r *http.Request){
	fmt.Printf("get dogs call")
	dogs,err:=store.GetDogs()
	dogsListBytes,err:=json.Marshal(dogs)
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(dogsListBytes)
}

func createDogHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("create dog call\n")
	
     var dog Dog
    body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
    fmt.Printf("body read\n")
	checkErr(err)

    err = r.Body.Close()
    checkErr(err)
    fmt.Printf("body closed\n")
    if err = json.Unmarshal(body, &dog); err != nil { fmt.Printf("json error")
        w.Header().Set("Content-Type", "application/json; charset=UTF-8")
        w.WriteHeader(422) // unprocessable entity
        if err := json.NewEncoder(w).Encode(err); err != nil {
           panic(err)
        }
   }
    fmt.Printf("unpacked\n")
   fmt.Printf("Dog name: %s\n",dog.Name)
    for v,i:= range dog.Vacinnes {
    	fmt.Printf("Vacinne %v : %s\n",v,i)
    }
   err = store.CreateDog(&dog)
    checkErr(err)
   w.Header().Set("Content-Type", "application/json; charset=UTF-8")
   w.WriteHeader(http.StatusCreated)

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
	checkErr(err)
	/*
	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}*/

	bird.Species = r.Form.Get("species")
	bird.Description = r.Form.Get("description")

	//birds = append(birds, bird)
	err = store.CreateBird(&bird)

	http.Redirect(w, r, "/assets/", http.StatusFound)
}