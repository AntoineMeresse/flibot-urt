package msg

const (
	// download.go
	DOWNLOAD_START = "Downloading map ^5%s^3: ^2Start"
	DOWNLOAD_OK = "Downloading map ^5%s^3: ^2OK^3 ^7(Size: ^5%s^7, time: ^5%s^3)"
	DOWNLOAD_KO = "Downloading map ^5%s^3: ^1KO^3"
	DOWNLOAD_ALREADY_ON_SERV = "^5%s^3 is ^1already^3 on server !"

	DOWNLOAD_NO_MAP = "No map found that contains (^5%s^3)"
	DOWNLOAD_MULTIPLE_MAP = "^7Maps found (^5%d^7) matching ^5%s^7:"
	DOWNLOAD_MAP_ITEM = "^7|-> ^5%s"
	DOWNLOAD_MAP_ITEM_ALREADY = "^7|-> ^5%s ^7(^2On server^7)"

	// goto.go
	GOTO_NO_LOCATION = "Location (^5%s^3) ^1doesn't^3 exist."

	// map_list.go 
	MAP_LIST = "Server map list [^5%d^3]: "

	// map.go
	MAP_CHANGE = "^7Changing map to %s"

	// removegoto.go
	GOTO_REMOVE = "Location (^5%s^3) has been deleted."
	GOTO_DONT_EXIST = "Location (^5%s^3) ^1doesn't^3 exist."

	// command.go
	NOT_ENOUGH_RIGHTS = "You ^1can't^3 use command ^5%s^3 ^7(required: ^6%d^7 | got: ^1%d^7)"
)