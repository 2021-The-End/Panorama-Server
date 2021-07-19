package info

const (
	DBHost          = "localhost"
	User            = "postgres"
	Password        = "rlawnsdn6!"
	Dbname          = "panorama"
	InsQuery        = "INSERT INTO " + "projectcon" + " (title,contents,creaters,imgpaths,summary,grade,created_at,updated_at) VALUES ($1, $2,$3,$4,$5,$6,$7,$8)"
	SelImgpathQuery = "SELECT imgpaths FROM projectcon WHERE id = $1"
	SelCreaterQuery = "SELECT creaters FROM projectcon WHERE id = $1"
)
