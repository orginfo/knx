package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"

	s3 "github.com/mattn/go-sqlite3"
	"knx/calc"
)

var DB *sql.DB

var sqlDeclarations []string = []string{
	`CREATE TABLE meta (
    key      TEXT NOT NULL DEFAULT '' UNIQUE,
    value    TEXT NOT NULL DEFAULT '')`,

	`CREATE TABLE user (
    id       INTEGER PRIMARY KEY,
    login    TEXT NOT NULL DEFAULT ''UNIQUE,
    name     TEXT NOT NULL DEFAULT '',
    phone    TEXT NOT NULL DEFAULT '',
    position TEXT NOT NULL DEFAULT '',
    comment  TEXT NOT NULL DEFAULT '')`,

	`CREATE TABLE client (
    id       INTEGER PRIMARY KEY,
    name     TEXT NOT NULL DEFAULT '',
    phone    TEXT NOT NULL DEFAULT '',
    comment  TEXT NOT NULL DEFAULT '')`,

	`CREATE INDEX idx_client_name ON client(name)`,

	`CREATE TABLE project (
    id            INTEGER PRIMARY KEY,
    nr            TEXT NOT NULL DEFAULT '' UNIQUE,
    client_id     INTEGER REFERENCES client(id) NOT NULL,
    user_id       INTEGER REFERENCES user(id)   NOT NULL,
    contract_date DATETIME,
    install_date  DATETIME,
    address       TEXT NOT NULL DEFAULT '',
    comment       TEXT NOT NULL DEFAULT '')`,

	`CREATE TABLE tregion (
    id            INTEGER PRIMARY KEY,
    name          TEXT NOT NULL DEFAULT '')`,

	`CREATE TABLE region (
    id            INTEGER PRIMARY KEY,
    description   TEXT NOT NULL DEFAULT '',
    project_id    INTEGER REFERENCES project(id) NOT NULL,
    tregion_id    INTEGER REFERENCES tregion(id) NOT NULL,
    nr            INTEGER NOT NULL DEFAULT 0)`,

	`CREATE TABLE tparam (
    id            INTEGER PRIMARY KEY,
    name          TEXT NOT NULL DEFAULT '',
    description   TEXT NOT NULL DEFAULT '')`,

	`CREATE TABLE param (
    id            INTEGER PRIMARY KEY,
    tparam_id     INTEGER REFERENCES tparam(id) NOT NULL,
    region_id     INTEGER REFERENCES region(id) NOT NULL,
    value         FLOAT NOT NULL DEFAULT 0)`,

	`CREATE TABLE cn_tparam_tregion (
    tparam_id     INTEGER REFERENCES tparam(id)  NOT NULL,
    tregion_id    INTEGER REFERENCES tregion(id) NOT NULL )`,

	`CREATE TABLE tparamvalue (
    tparam_id     INTEGER REFERENCES tparam(id) NOT NULL,
    value         FLOAT NOT NULL DEFAULT 0,
    name          TEXT NOT NULL DEFAULT '')`,

	`CREATE TABLE tresult (
    id            INTEGER PRIMARY KEY,
    name          TEXT NOT NULL DEFAULT '',
    description   TEXT NOT NULL DEFAULT '')`,

	`CREATE TABLE color_scheme (
    id              INTEGER PRIMARY KEY,
    name            TEXT NOT NULL DEFAULT '')`,

	`CREATE TABLE color (
    id              INTEGER PRIMARY KEY,
    name            TEXT NOT NULL DEFAULT '',
    color_scheme_id INTEGER REFERENCES color_scheme(id) NOT NULL,
    value           INTEGER NOT NULL DEFAULT 0)`,

	`CREATE TABLE tnomenclature (
    id              INTEGER PRIMARY KEY,
    name            TEXT NOT NULL DEFAULT '',
    color_scheme_id INTEGER REFERENCES color_scheme(id) )`,

	`CREATE TABLE nomenclature (
    id               INTEGER PRIMARY KEY,
    tnomenclature_id INTEGER REFERENCES tnomenclature(id) NOT NULL,
    vendor_code      TEXT NOT NULL DEFAULT '',
    name             TEXT NOT NULL DEFAULT '',
    measure_unit     TEXT NOT NULL DEFAULT '',
    material         TEXT NOT NULL DEFAULT '',
    thickness        FLOAT NOT NULL DEFAULT 0,
    color_id         INTEGER REFERENCES color(id),
    size             FLOAT NOT NULL DEFAULT 0,
    division         TEXT NOT NULL DEFAULT '',
    division_service_nomenclature_id INTEGER REFERENCES nomenclature(id) )`,

	`CREATE INDEX idx_nomenclature_name ON nomenclature(name)`,

	`CREATE TABLE result (
    id              INTEGER PRIMARY KEY,
    tresult_id      INTEGER REFERENCES tresult(id) NOT NULL,
    region_id       INTEGER REFERENCES region(id) NOT NULL,
    nomenclature_id INTEGER REFERENCES nomenclature(id),
    value           FLOAT NOT NULL DEFAULT 0)`,

	`CREATE TABLE cn_tnomenclature_usefield (
    tnomenclature_id INTEGER REFERENCES tnomenclature(id) ON DELETE CASCADE ON UPDATE CASCADE NOT NULL,
    field_name       TEXT NOT NULL DEFAULT '')`,

	`CREATE TABLE tcomponent (
    id              INTEGER PRIMARY KEY,
    name            TEXT NOT NULL DEFAULT '')`,

	`CREATE TABLE tcalculation (
    id              INTEGER PRIMARY KEY,
    name            TEXT NOT NULL DEFAULT '')`,

	`CREATE TABLE tpart (
    id              INTEGER PRIMARY KEY,
    name            TEXT NOT NULL DEFAULT '',
    tcomponent_id   INTEGER REFERENCES tcomponent(id) NOT NULL,
    tcalculation_id INTEGER REFERENCES tcalculation(id) )`,

	`CREATE TABLE cn_tpart_nomenclature (
    tpart_id        INTEGER REFERENCES tpart(id) NOT NULL,
    nomenclature_id INTEGER REFERENCES nomenclature(id) )`,

	`CREATE TABLE cn_tregion_tcomponent (
    tregion_id      INTEGER REFERENCES tregion(id) NOT NULL,
    tcomponent_id   INTEGER REFERENCES tcomponent(id) NOT NULL )`,

	`CREATE TABLE component (
    id              INTEGER PRIMARY KEY,
    tcomponent_id   INTEGER REFERENCES tcomponent(id) NOT NULL,
    region_id       INTEGER REFERENCES region(id) NOT NULL )`,

	`CREATE TABLE part (
    id              INTEGER PRIMARY KEY,
    tpart_id        INTEGER REFERENCES tpart(id) NOT NULL,
    component_id    INTEGER REFERENCES component(id) NOT NULL,
    nomenclature_id INTEGER REFERENCES nomenclature(id) )`,

	`CREATE TABLE price (
    nomenclature_id INTEGER REFERENCES nomenclature(id) NOT NULL,
    date            DATETIME NOT NULL,
    cost_price      INTEGER NOT NULL DEFAULT 0,
    price           INTEGER NOT NULL DEFAULT 0,
    UNIQUE(nomenclature_id, date) )`,

	`CREATE TABLE cn_tparam_tpart (
    tparam_id       INTEGER REFERENCES tparam(id) NOT NULL,
    tpart_id        INTEGER REFERENCES tpart(id) NOT NULL,
    UNIQUE(tparam_id, tpart_id) )`,

	`CREATE TABLE cn_tparamvalue_nomenclature (
    tparam_id       INTEGER REFERENCES tparam(id),
    value           FLOAT NOT NULL DEFAULT 0,
    nomenclature_id INTEGER REFERENCES nomenclature(id) NOT NULL,
    UNIQUE(tparam_id, value, nomenclature_id) )`,

	`CREATE TABLE cn_tparamvalue_tparamvalue (
    tparam_id           INTEGER REFERENCES tparam(id) NOT NULL,
    value               FLOAT NOT NULL DEFAULT 0,
    dependent_tparam_id INTEGER REFERENCES tparam(id) NOT NULL,
    dependent_value     FLOAT NOT NULL DEFAULT 0,
    CHECK(tparam_id <> dependent_tparam_id),
    UNIQUE(tparam_id,value,dependent_tparam_id, dependent_value) )`,
}

const (
	MetaKeyVersion = "version"
)

var MetaValues map[string]string = map[string]string{
	MetaKeyVersion: "2018-04-19",
}

// convertDB - convert DB from one version to another
func convertDB(db *sql.DB, sourceVersion string, targetVersion string) error {
	// TODO: Make DB converter
	return fmt.Errorf("DB has unsupportable version '%s'. Actual version should be '%s'. Converting DB is not supported!",
		sourceVersion, targetVersion)
}

// createDB - create new DB with actual version
func createDB(db *sql.DB) (err error) {
	// Start transaction
	tx, err := db.Begin()
	if err != nil {
		return
	}

	defer tx.Commit()

	// Create empty tables
	for _, sqlQuery := range sqlDeclarations {
		_, err = tx.Exec(sqlQuery)
		if err != nil {
			return
		}
	}

	// Fill the tables with constant enums
	// [meta]
	for key, value := range MetaValues {
		_, err = tx.Exec("INSERT INTO meta(key,value) VALUES(?,?)", key, value)
		if err != nil {
			return
		}
	}

	// [tregion]
	for id, tregion := range calc.Regions {
		_, err = tx.Exec("INSERT INTO tregion(id,name) VALUES(?,?)", id, tregion.Name)
		if err != nil {
			return
		}
	}

	// [tparam] [tparamvalue] [cn_tparam_tregion]
	for id, tparam := range calc.Params {
		_, err = tx.Exec("INSERT INTO tparam(id,name,description) VALUES(?,?,?)", id, tparam.Name, tparam.Description)
		if err != nil {
			return
		}

		for _, value := range tparam.Values {

			// For color param declaration: use colors from color scheme as values
			if value.Name == calc.ColorParamName {
				cs := calc.ColorSchemes[int(value.Value)]
				for _, color := range cs.Colors {
					_, err = tx.Exec("INSERT INTO tparamvalue(tparam_id,value,name) VALUES(?,?,?)", id, color.Value, color.Name)
					if err != nil {
						return
					}
				}
				continue
			}

			_, err = tx.Exec("INSERT INTO tparamvalue(tparam_id,value,name) VALUES(?,?,?)", id, value.Value, value.Name)
			if err != nil {
				return
			}
		}
		for _, rt := range tparam.RegionTypes {
			_, err = tx.Exec("INSERT INTO cn_tparam_tregion(tparam_id,tregion_id) VALUES(?,?)", id, rt)
			if err != nil {
				return
			}
		}
	}

	// [cn_tparamvalue_tparamvalue]
	for id, tparam := range calc.Params {
		for _, value := range tparam.Values {
			for dependendParamID, dependentValues := range value.DependentParams {
				// For empty value list add -1 means value is not set and is not available to choose
				if len(dependentValues) == 0 {
					_, err = tx.Exec("INSERT INTO cn_tparamvalue_tparamvalue(tparam_id,value,dependent_tparam_id, dependent_value) VALUES(?,?,?,?)",
						id, value.Value, dependendParamID, -1)
					if err != nil {
						return
					}
				}
				// Add all the values, manually set in declaration array
				for _, dependentValue := range dependentValues {
					_, err = tx.Exec("INSERT INTO cn_tparamvalue_tparamvalue(tparam_id,value,dependent_tparam_id, dependent_value) VALUES(?,?,?,?)",
						id, value.Value, dependendParamID, dependentValue)
					if err != nil {
						return
					}
				}
			}
		}
	}

	// [tresult]
	for id, tresult := range calc.Results {
		_, err = tx.Exec("INSERT INTO tresult(id,name,description) VALUES(?,?,?)", id, tresult.Name, tresult.Description)
		if err != nil {
			return
		}
	}

	// [tcalculation]
	for id, tcalc := range calc.MaterialCalculations {
		_, err = tx.Exec("INSERT INTO tcalculation(id,name) VALUES(?,?)", id, tcalc.Name)
		if err != nil {
			return
		}
	}

	// [tcomponent] [tpart]
	for id, tcomp := range calc.Components {
		_, err = tx.Exec("INSERT INTO tcomponent(id, name) VALUES(?, ?)", id, tcomp.Name)
		if err != nil {
			return
		}
		for _, tpart := range tcomp.Parts {
			_, err = tx.Exec("INSERT INTO tpart(tcomponent_id, name, tcalculation_id) VALUES(?,?,?)", id, tpart.Name, tpart.MC)
			if err != nil {
				return
			}
		}
	}

	// [cn_tregion_tcomponent]
	for id, tregion := range calc.Regions {
		for _, compID := range tregion.Components {
			_, err = tx.Exec("INSERT INTO cn_tregion_tcomponent(tregion_id,tcomponent_id) VALUES(?,?)", id, compID)
			if err != nil {
				return
			}
		}
	}

	// [color]
	for _, cs := range calc.ColorSchemes {
		var res sql.Result
		res, err = tx.Exec("INSERT INTO color_scheme(name) VALUES(?)", cs.Name)
		if err != nil {
			return
		}
		id, _ := res.LastInsertId()

		for _, color := range cs.Colors {
			_, err = tx.Exec("INSERT INTO color(name,color_scheme_id,value) VALUES(?,?,?)", color.Name, id, color.Value)
			if err != nil {
				return
			}
		}
	}

	return
}

// initDB - check DB existance and version, create db
func InitDB(dbPath string) (db *sql.DB, err error) {
	db = nil

	// Create all the parent directories
	if err = os.MkdirAll(filepath.Dir(dbPath), os.ModePerm); err != nil {
		return
	}

	// Open the DB
	var dbPathPlusParams string = dbPath + "?_foreign_keys=1"
	db, err = sql.Open("sqlite3", dbPathPlusParams)
	if err != nil {
		return
	}

	// Check if the file exists, if not - create new DB
	if _, err = os.Stat(dbPath); os.IsNotExist(err) {
		err = createDB(db)
		return
	}

	// Check DB-version of existing DB
	rows, err := db.Query("SELECT value FROM meta WHERE key=?", MetaKeyVersion)
	if err != nil {
		return
	}
	defer rows.Close()

	dbVersion := "<Unknown>"
	for rows.Next() {
		err = rows.Scan(&dbVersion)
		if err != nil {
			return
		}
	}

	if dbVersion != MetaValues[MetaKeyVersion] {
		err = convertDB(db, dbVersion, MetaValues[MetaKeyVersion])
		if err != nil {
			return
		}
	}

	err = rows.Err()
	if err != nil {
		return
	}

	return
}

func init() {
	libVersion, libVersionNumber, sourceID := s3.Version()
	log.Printf("SQLITE3 Version: %v %v %v\n", libVersion, libVersionNumber, sourceID)

	// Init DB
	dbPath := "../data/db.sqlite3"
	if len(os.Args) > 1 {
		dbPath = os.Args[1]
	}

	var err error
	DB, err = InitDB(dbPath)
	if err != nil {
		fmt.Printf("An error occured while opening db '%s': %v\n", dbPath, err)
		return
	}
}
