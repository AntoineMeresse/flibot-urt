package actions

import actionslist "github.com/AntoineMeresse/flibot-urt/src/actions/actions_list"

var Actions map[string]interface{} = map[string]interface{}{
	"ClientBegin:": actionslist.ClientBegin,
	"say:": actionslist.Say,
	"ClientConnect:": actionslist.ClientConnect,
	"ClientUserinfo:" : actionslist.ClientUserinfo,
}


