package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/HeChX/REST_API_Server/model"

	"github.com/boltdb/bolt"
)

type Field struct {
	Fields map[string]interface{}
	Pk     int
}

var db, err = bolt.Open("../database/my.db", 0600, nil)

func main() {
	creataBucket()
	initPlanets()
	initPeople()
	initSpecies()
	initTransports()
	initStarships()
	initVehicles()
	initFilms()

	// db.View(func(tx *bolt.Tx) error {
	// 	b := tx.Bucket([]byte("People"))
	// 	v := b.Get([]byte("1"))
	// 	fmt.Println(string(v))
	// 	return nil
	// })
}
func initSpecies() {
	data, err := ioutil.ReadFile("resourses/species.json")
	if err != nil {
		return
	}
	// people := model.Peoples{}
	fields := []Field{}
	err = json.Unmarshal(data, &fields)
	for i := 0; i < len(fields); i++ {
		insertSpecies(&fields[i])
	}

}

func initPeople() {
	data, err := ioutil.ReadFile("resourses/people.json")
	if err != nil {
		return
	}
	// people := model.Peoples{}
	fields := []Field{}
	err = json.Unmarshal(data, &fields)
	for i := 0; i < len(fields); i++ {
		insertPeople(&fields[i])
	}

}

func initFilms() {
	data, err := ioutil.ReadFile("resourses/films.json")
	if err != nil {
		return
	}
	// people := model.Peoples{}
	fields := []Field{}
	err = json.Unmarshal(data, &fields)
	for i := 0; i < len(fields); i++ {
		insertFilm(&fields[i])
	}

}

func initPlanets() {
	data, err := ioutil.ReadFile("resourses/planets.json")
	if err != nil {
		return
	}
	// people := model.Peoples{}
	fields := []Field{}
	err = json.Unmarshal(data, &fields)
	for i := 0; i < len(fields); i++ {
		insertPlanet(&fields[i])
	}

}

func initTransports() {
	data, err := ioutil.ReadFile("resourses/transport.json")
	if err != nil {
		return
	}
	// people := model.Peoples{}
	fields := []Field{}
	err = json.Unmarshal(data, &fields)
	for i := 0; i < len(fields); i++ {
		insertTransport(&fields[i])
	}

}

func initStarships() {
	data, err := ioutil.ReadFile("resourses/starships.json")
	if err != nil {
		return
	}
	// people := model.Peoples{}
	fields := []Field{}
	err = json.Unmarshal(data, &fields)
	for i := 0; i < len(fields); i++ {
		insertStarship(&fields[i])
	}

}

func initVehicles() {
	data, err := ioutil.ReadFile("resourses/vehicles.json")
	if err != nil {
		return
	}
	// people := model.Peoples{}
	fields := []Field{}
	err = json.Unmarshal(data, &fields)
	for i := 0; i < len(fields); i++ {
		insertVehicle(&fields[i])
	}

}
func creataBucket() {

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Transport"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Vehicle"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("People"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Film"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Planet"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Species"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("Starship"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("User"))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err)
		}
		return nil
	})
}

func insertTransport(field *Field) {
	strId := strconv.Itoa(field.Pk)
	var m map[string]interface{} = field.Fields

	data, _ := json.Marshal(&m)
	transport := model.Transport{}
	_ = json.Unmarshal(data, &transport)

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Transport"))
		data, _ := json.Marshal(&transport)
		err := b.Put([]byte(strId), data)
		return err
	})
}

func insertVehicle(field *Field) {
	strId := strconv.Itoa(field.Pk)
	var transportData []byte
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Transport"))
		transportData = b.Get([]byte(strId))
		return nil
	})

	var m map[string]interface{}
	_ = json.Unmarshal(transportData, &m)

	var vehicleData map[string]interface{} = field.Fields

	m["vehicle_class"] = vehicleData["vehicle_class"]
	m["url"] = "http://localhost:8080/vehicles/" + strId

	pilots := vehicleData["pilots"].([]interface{})
	var peopleURL []string
	for i := 0; i < len(pilots); i++ {
		id := strconv.Itoa(int(pilots[i].(float64)))
		url := "http://localhost:8080/people/" + id
		peopleURL = append(peopleURL, url)
		peopleUpdate(id, m["url"].(string), "vehicles")
	}
	delete(m, "pilots")
	m["pilots"] = peopleURL

	data, _ := json.Marshal(&m)
	vehicle := model.Vehicle{}
	_ = json.Unmarshal(data, &vehicle)

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Vehicle"))
		data, _ := json.Marshal(&vehicle)
		err := b.Put([]byte(strId), data)
		return err
	})
}

func insertStarship(field *Field) {

	strId := strconv.Itoa(field.Pk)
	var transportData []byte
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Transport"))
		transportData = b.Get([]byte(strId))
		return nil
	})

	var transport map[string]interface{}
	_ = json.Unmarshal(transportData, &transport)

	var m map[string]interface{} = field.Fields

	m["edited"] = transport["edited"]
	m["name"] = transport["name"]
	m["model"] = transport["model"]
	m["manufacturer"] = transport["manufacturer"]
	m["cost_in_credits"] = transport["cost_in_credits"]
	m["max_atmosphering_speed"] = transport["max_atmosphering_speed"]
	m["crew"] = transport["crew"]
	m["passengers"] = transport["passengers"]
	m["Cargo_capacity"] = transport["Cargo_capacity"]
	m["consumables"] = transport["consumables"]
	m["created"] = transport["created"]
	m["url"] = "http://localhost:8080/starships/" + strId

	pilots := m["pilots"].([]interface{})
	var peopleURL []string
	for i := 0; i < len(pilots); i++ {
		id := strconv.Itoa(int(pilots[i].(float64)))
		url := "http://localhost:8080/people/" + id
		peopleURL = append(peopleURL, url)
		peopleUpdate(id, m["url"].(string), "starships")
	}
	delete(m, "pilots")
	m["pilots"] = peopleURL

	data, _ := json.Marshal(&m)
	starship := model.Starship{}
	_ = json.Unmarshal(data, &starship)

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Starship"))
		data, _ := json.Marshal(&starship)
		err := b.Put([]byte(strId), data)
		return err
	})
}

func insertPlanet(field *Field) {
	strId := strconv.Itoa(field.Pk)
	var m map[string]interface{} = field.Fields

	m["url"] = "http://localhost:8080/planets/" + strId

	data, _ := json.Marshal(&m)
	planet := model.Planet{}
	_ = json.Unmarshal(data, &planet)

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Planet"))
		data, _ := json.Marshal(&planet)
		err := b.Put([]byte(strId), data)
		return err
	})

}

func insertSpecies(field *Field) {
	strId := strconv.Itoa(field.Pk)
	var m map[string]interface{} = field.Fields

	if m["homeworld"] != nil {
		homeworld := int(m["homeworld"].(float64))
		delete(m, "homeworld")
		m["homeworld"] = "http://localhost:8080/planet/" + strconv.Itoa(homeworld)
	}

	m["url"] = "http://localhost:8080/species/" + strId

	people := m["people"].([]interface{})
	var peopleURL []string

	for i := 0; i < len(people); i++ {
		cid := strconv.Itoa(int(people[i].(float64)))
		url := "http://localhost:8080/people/" + cid
		peopleURL = append(peopleURL, url)
		peopleUpdate(cid, m["url"].(string), "species")
	}
	delete(m, "people")
	m["people"] = peopleURL

	data, _ := json.Marshal(&m)
	species := model.Species{}
	_ = json.Unmarshal(data, &species)

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Species"))
		data, _ := json.Marshal(&species)
		err := b.Put([]byte(strId), data)
		return err
	})

}

func insertPeople(field *Field) {
	strId := strconv.Itoa(field.Pk)
	var m map[string]interface{} = field.Fields

	homeworld := int(m["homeworld"].(float64))
	delete(m, "homeworld")
	m["url"] = "http://localhost:8080/people/" + strId
	m["homeworld"] = "http://localhost:8080/planet/" + strconv.Itoa(homeworld)

	data, _ := json.Marshal(&m)
	people := model.People{}
	_ = json.Unmarshal(data, &people)

	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("People"))
		data, _ := json.Marshal(&people)
		err := b.Put([]byte(strId), data)
		return err
	})

	planetUpdate(strconv.Itoa(homeworld), m["url"].(string), true)

}

func insertFilm(field *Field) {
	strId := strconv.Itoa(field.Pk)
	var m map[string]interface{} = field.Fields

	m["url"] = "http://localhost:8080/films/" + strId
	characters := m["characters"].([]interface{})
	planets := m["planets"].([]interface{})
	species := m["species"].([]interface{})
	starships := m["starships"].([]interface{})
	vehicles := m["vehicles"].([]interface{})
	var charactersURL []string
	var planetsURL []string
	var speciesURL []string
	var starshipsURL []string
	var vehiclesURL []string

	for i := 0; i < len(characters); i++ {
		cid := strconv.Itoa(int(characters[i].(float64)))
		url := "http://localhost:8080/people/" + cid
		charactersURL = append(charactersURL, url)
		peopleUpdate(cid, m["url"].(string), "films")
	}
	delete(m, "characters")
	m["characters"] = charactersURL

	for i := 0; i < len(planets); i++ {
		pid := strconv.Itoa(int(planets[i].(float64)))
		url := "http://localhost:8080/planets/" + pid
		planetsURL = append(planetsURL, url)
		planetUpdate(pid, m["url"].(string), false)
	}
	delete(m, "planets")
	m["planets"] = planetsURL

	for i := 0; i < len(species); i++ {
		sid := strconv.Itoa(int(species[i].(float64)))
		url := "http://localhost:8080/species/" + sid
		speciesURL = append(speciesURL, url)
		speciesUpdate(sid, m["url"].(string))
	}
	delete(m, "species")
	m["species"] = speciesURL

	for i := 0; i < len(starships); i++ {
		sid := strconv.Itoa(int(starships[i].(float64)))
		url := "http://localhost:8080/starships/" + sid
		starshipsURL = append(starshipsURL, url)
		starshipUpdate(sid, m["url"].(string))
	}
	delete(m, "starships")
	m["starships"] = starshipsURL

	for i := 0; i < len(vehicles); i++ {
		sid := strconv.Itoa(int(vehicles[i].(float64)))
		url := "http://localhost:8080/vehicles/" + sid
		vehiclesURL = append(vehiclesURL, url)
		vehicleUpdate(sid, m["url"].(string))
	}
	delete(m, "vehicles")
	m["vehicles"] = vehiclesURL

	data, _ := json.Marshal(&m)
	film := model.Film{}
	_ = json.Unmarshal(data, &film)

	//updata people film url
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Film"))
		data, _ := json.Marshal(&film)
		err := b.Put([]byte(strId), data)
		return err
	})

}

func peopleUpdate(strId string, url string, sel string) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("People"))
		v := b.Get([]byte(strId))
		people := model.People{}
		_ = json.Unmarshal(v, &people)
		switch sel {
		case "films":
			people.Films = append(people.Films, url)
		case "species":
			people.Species = append(people.Species, url)
		case "starships":
			people.Starships = append(people.Starships, url)
		case "vehicles":
			people.Vehicles = append(people.Vehicles, url)

		}

		data, _ := json.Marshal(&people)
		err := b.Put([]byte(strId), data)
		return err
	})

}

func planetUpdate(strId string, url string, flag bool) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Planet"))
		v := b.Get([]byte(strId))
		planet := model.Planet{}
		_ = json.Unmarshal(v, &planet)
		if flag == true {
			planet.Residents = append(planet.Residents, url)
		} else {
			planet.Films = append(planet.Films, url)
		}
		data, _ := json.Marshal(&planet)
		err := b.Put([]byte(strId), data)
		return err
	})

}

func speciesUpdate(strId string, url string) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Species"))
		v := b.Get([]byte(strId))
		species := model.Species{}
		_ = json.Unmarshal(v, &species)
		species.Films = append(species.Films, url)

		data, _ := json.Marshal(&species)
		err := b.Put([]byte(strId), data)
		return err
	})

}

func starshipUpdate(strId string, url string) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Starship"))
		v := b.Get([]byte(strId))
		starship := model.Starship{}
		_ = json.Unmarshal(v, &starship)
		starship.Films = append(starship.Films, url)

		data, _ := json.Marshal(&starship)
		err := b.Put([]byte(strId), data)
		return err
	})

}

func vehicleUpdate(strId string, url string) {
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Vehicle"))
		v := b.Get([]byte(strId))
		vehicle := model.Vehicle{}
		_ = json.Unmarshal(v, &vehicle)
		vehicle.Films = append(vehicle.Films, url)

		data, _ := json.Marshal(&vehicle)
		err := b.Put([]byte(strId), data)
		return err
	})

}
