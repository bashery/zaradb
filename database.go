package dblite

import (
	"io/fs"
	"os"
)

var pi = "pi" // primary index

type Database struct {
	PrimaryIndex int64
	Indexs       map[string]Index
	Name         string
	Collection   string
	//	collections  map[string]Collection
	Pages map[string]*os.File
}

type Collection struct {
	primaryIndex int64
	at           int
}

// NewCollection constracts List of files collection
func NewDatabase(name string) *Database {
	database := &Database{
		Name:       rootPath() + name + slash,
		Collection: "test" + slash,
		Pages:      make(map[string]*os.File, 2),
		Indexs:     make(map[string]Index, 1),
	}
	return database
}

// opnens all collection in Root database folder
func (db *Database) Open() {
	path := db.Name + db.Collection
	//	iLog.Println("opening database ", path)

	var err error
	var files []fs.DirEntry

	files, err = os.ReadDir(path)
	if os.IsNotExist(err) {
		err = os.MkdirAll(db.Name+db.Collection, 0744)
		if err != nil {
			eLog.Println("while mkDir", err)
		}
	}

	_, err = os.Stat(db.Name + db.Collection + pi)
	if os.IsNotExist(err) {
		f, err := os.OpenFile(db.Name+db.Collection+pi, os.O_CREATE|os.O_RDWR, 0644)
		if err != nil {
			eLog.Println("when creating pi ", err)
			return
		}
		f.Close()
	}

	files, err = os.ReadDir(path)
	if err != nil {
		eLog.Printf("while reading dir %s, %v\n\n", path, err)
		return
	}

	//iLog.Printf("reading  %s\n", db.Name+db.Collections)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		page, err := os.OpenFile(path+file.Name(), os.O_RDWR, 0644) //
		if err != nil {
			iLog.Printf("os open file: %s,  %v\n", path+file.Name(), err)
			//break
		}
		// filepath.Join(path, file.Name())
		db.Pages[db.Name+db.Collection+file.Name()] = page

		//	iLog.Printf("%s is ready\n", file.Name())
	}
	if len(db.Pages) < 2 {
		page, err := os.OpenFile(path+"0", os.O_CREATE|os.O_RDWR, 0644) //
		if err != nil {
			iLog.Printf("os open file: %s,  %v\n", path+"0", err)
		}
		// filepath.Join(path, file.Name())
		db.Pages[db.Name+db.Collection+"0"] = page

	}
	// iLog.Println("length of db.Pages is : ", len(db.Pages))
}

// closes All collection
func (db *Database) Close() {
	for _, page := range db.Pages {
		page.Close()
		iLog.Printf("%s closed\n", page.Name())
	}
}

/*

// creates new page and add it to Collections
func (db *Database) NewPage(id int) {
	// TODO
	indexFilePath := db.Name + db.Collection + pi

	filename, _, _ := GetIndex(db.Pages[indexFilePath], id)
	//	iLog.Println("GetIndex from :", indexFilePath)

	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}

	path := filepath.Join(db.Name, db.Collection+strconv.Itoa(id))

	db.Pages[path] = file
	//iLog.Printf("new page is created with %s path\n", path)
}
*/
