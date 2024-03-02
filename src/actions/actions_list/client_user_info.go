package actionslist

import (
	"fmt"
	"strings"

	"github.com/AntoineMeresse/flibot-urt/src/models"
	"github.com/AntoineMeresse/flibot-urt/src/utils"
)

func ClientUserinfo(action_params []string, server models.Server) {
	fmt.Printf("\nClient User Info: %v", action_params)
	if (len(action_params) > 1) {
		// playerNumber := action_params[0]
		infoString := strings.Join(action_params[1:], "")
		infosMap := splitInfos(infoString)

		fmt.Printf("\nInfos: \n%v", infosMap)
	}
	
}

func splitInfos(infos string) map[string]string {
	res := make(map[string]string)

	infoSplit := utils.CleanEmptyElements(strings.Split(infos, "\\"))

	for i:=0; i < len(infoSplit)-1; i+=2 {
		res[infoSplit[i]] = infoSplit[i+1]
	}

	return res;
}