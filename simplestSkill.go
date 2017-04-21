package example

import "mind/core/framework/skill"

/*
To develop a Skill, implement the skill.Base interface which is defined as:

type Interface interface {
    OnStart()             // Called when Skill is started
    OnClose()             // Called when Skill is closed
    OnConnect()           // Called when Skill is connected
    OnDisconnect()        // Called when Skill is disconnected
    OnRecvJSON([]byte)    // Called when remote sent data
    OnRecvString(string)  // Called when remote sent a string
}
*/

type Example struct {
    skill.Base
}

/*
Before using a driver the drivers Start() method has to be invoked as well as Close() when finished.
List of all drivers: https://www.vincross.com/developer/api-reference/mind-sdk-robot-part/drivers/
*/

func NewSkill() skill.Interface {
    return &Example{}
}
