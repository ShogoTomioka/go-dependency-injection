## はじめに
　GolangでDependency Injection(DI)を使って実装する手順はどうなのかな〜と思ったので調べて実装してみた。

## Dependency Injection(DI)とは
>依存性の注入（いそんせいのちゅうにゅう、英: Dependency injection）とは、コンポーネント間の依存関係をプログラムのソースコードから排除し、外部の設定ファイルなどで注入できるようにするソフトウェアパターンである。(Wikipediaより)

ということらしいです。DIに関しては、詳しくまとめてくださっている記事があるのでそちらなどで確認していただけるとより理解が深まると思います。[猿でも分かる! Dependency Injection: 依存性の注入](https://qiita.com/hshimo/items/1136087e1c6e5c5b0d9f)

DIを用いるメリットとして、「依存関係を明確にできる」、「モックなどを使ったテストを描きやすくできる」などがあります。それらについても後で触れたいと思います。

## 参考サイト
[Dependency Injection in Go](https://blog.drewolson.org/dependency-injection-in-go)  
[GoにおけるDI](http://inukirom.hatenablog.com/entry/di-in-go)  
[GoとDependency Injection](https://recruit-tech.co.jp/blog/2017/12/11/go_dependency_injection/)  

## 想定するユースケース
ここでは、ユースケースとして、
* ユーザの情報の全てをデータベースから取得  
* ユーザIDを使ってユーザの情報をデータベースから取得  
という二つのシンプルなパターンを想定して行きます。


## 実装

各モジュールの関係は以下の図のような感じでイメージしてます。  
左から右に向かって依存の方向が向いており、各モジュールがDIでInjectionしたい対象ですね。

### ドメイン
まずは、ドメインとなるUserを定義します。
```
type User struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}
```
Userの定義は以上で、シンプルに三つのフィールドを持っているのみです。

### データベース
 データベースへのコネクションを行うために、”sql.DB"型のオブジェクトを返すメソッドを定義します。また、DB接続に必要な設定を取得するためのメソッドも用意します。今回はDBの設定はハードコーディングしちゃいます。
 ```
// DB接続に必要な情報を持った構造体
type Config struct {
	Host     string
	Database string
	DbName   string
	User     string
	Password string
	Port     string
	Protocol string
}

func NewConfig() *Config {
	return &Config{
		Host:     "localhost",
		Database: "sqlite3",
		DbName:   "testdb",
		User:     "root",
		Password: "",
		Port:     "3306",
		Protocol: "tcp",
	}
}

// sql.DBを返却する
func ConnectDatabase(c *Config) (*sql.DB, error) {

	dbConnection := fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
		c.User,
		c.Password,
		c.Protocol,
		c.Host,
		c.Port,
		c.DbName)
	return sql.Open("sqlite3", dbConnection)
}
 ```

今回は抽象化していないですが、DBへの接続関係を抽象化する場合には、後述するuser_repositoryと同じレイヤーにDBへの接続に関係するメソッドをもつinterfaceを定義しておくことで、DBの種類の変更などに対応できる形になると思います。


## User_Repository
次に、実際にsql.DBを使ってクエリを発行し、データベースのユーザ情報を構造体にマッピングして返却する役割を担っているUser_Repositoryを定義して行きます。中身の実装を全部書くと長くなってしまうので、細かい部分は省略
```
type UserRepository struct {
	database *sql.DB
}

func (repository *UserRepository) FindById(uid int) *User {
	query := fmt.Sprintf("SELECT id, name, age WHERE id = %s FROM users;", uid)
	rows, _ := repository.database.Query(query)
	defer rows.Close()
    // 構造体へのマッピング

	return &User{id, name, age}
}

func (repository *UserRepository) FindAll() []*User {
	rows, _ := repository.database.Query(
		"SELECT id, name, age FROM users;")

	defer rows.Close()
    // 構造体へのマッピング
	return users
}

// UserRepositoryのコンストラクタ、必要なのは*sql.DB
func NewUserRepository(database *sql.DB) *UserRepository {
	return &UserRepository{database: database}
}
```

## 



