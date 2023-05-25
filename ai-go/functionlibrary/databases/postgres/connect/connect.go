package connect

type Postgres struct {
	Host     string
	Port     string
	Username string
	Password string
	DB       string
}

func (p *Postgres) PostgresConn() {

}
