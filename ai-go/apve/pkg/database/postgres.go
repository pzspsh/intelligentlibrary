package database

type PgOptions struct {
	Host     string `json:"host,omitempty"`
	Port     int    `json:"port,omitempty"`
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Db       string `json:"db,omitempty"`
}

func Postgres() {

}
