# go-check-database-users

go-check-database-users is command line tool for checking the database user

# install

```
go get github.com/hlts2/go-check-database-users
```

## CLI Usage

```
$ go-check-database-users --help

Usage:
  go-check-database-users [flags]
  go-check-database-users [command]

Available Commands:
  help        Help about any command
  ls          List for database users

Flags:
  -H, --Host string           Hostname (default "localhost")
  -n, --account-Host string   Account host name
  -a, --account-name string   Account user name
  -d, --dbms string           Database management system (default "mysql")
  -h, --help                  help for go-check-database-users
  -P, --password string       Password
  -p, --port int              Port (default 3306)
  -u, --user string           Username (default "root")
```

```
Usage:
  go-check-database-users ls [flags]

Flags:
  -h, --help   help for ls

Global Flags:
  -H, --Host string       Hostname (default "localhost")
  -d, --dbms string       Database management system (default "mysql")
  -P, --password string   Password
  -p, --port int          Port (default 3306)
  -u, --user string       Username (default "root")
```

### Example

You can check the existence of the user by executing this command

```
$ go-check-database-users -H 192.168.33.10 -p 3306 -d mysql -u root -P root  -n localhost -a aaaaa
Database User NG: user 'aaaaa'@'localhost' not found
```


You can check the list of users by running this comman

```
$ go-check-database-users ls -H 192.168.33.10 -p 3306 -d mysql -u root -P root
+-----------+------+
|   HOST    | NAME |
+-----------+------+
| %         | root |
| 127.0.0.1 | root |
| ::1       | root |
| localhost | root |
+-----------+------+
```
