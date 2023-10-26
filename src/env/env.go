package env

// dev
const DbNameDev = "./sqlite/api.db"
const PrivateKeyPath = "./dev/secrets/id_rsa"
const PublicKeyPath = "./dev/secrets/id_rsa.pub"

// test
const DbNameTest = "./sqlite/test.db"

type Environment struct {
	Env    string
	DbName string
}

var Env *Environment

func SetEnv(e string) {
	if e == "test" {
		Env = &Environment{
			Env:    "test",
			DbName: DbNameTest,
		}
	} else {
		// dev
		Env = &Environment{
			Env:    "dev",
			DbName: DbNameDev,
		}
	}
}
