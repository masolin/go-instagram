package instagram

import (
	"time"

	"fmt"

	"github.com/masolin/go-instagram/models"
	"github.com/masolin/go-instagram/utils"
)

type Instagram struct {
	Username  string
	Password  string
	AgentPool *utils.SuperAgentPool
	Inbox     *models.Inbox
	Thread    *models.Thread
}

func Create(username string, password string, poolSize int, sleep string) (*Instagram, error) {
	if poolSize < 1 {
		poolSize = 1
	}
	pool, err := utils.NewSuperAgentPool(poolSize)
	if err != nil {
		return nil, err
	}

	ig := Instagram{
		Username:  username,
		Password:  password,
		AgentPool: pool,
	}

	ig.Login(sleep)

	ig.Inbox = &models.Inbox{AgentPool: ig.AgentPool}

	return &ig, nil
}

func (ig Instagram) Login(sleep string) {
	sleepDuration, err := time.ParseDuration(sleep)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < ig.AgentPool.Len(); i++ {
		if i != 0 {
			time.Sleep(sleepDuration)
		}

		uuid := utils.GenerateUUID()

		agent := ig.AgentPool.Get()
		defer ig.AgentPool.Put(agent)

		login := models.Login{
			Csrftoken:         "missing",
			DeviceID:          "android-b256317fd493b848",
			UUID:              uuid,
			UserName:          ig.Username,
			Password:          ig.Password,
			LoginAttemptCount: 0,
			Agent:             agent,
		}

		login.Login()
	}
}
