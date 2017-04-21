/*
* StandSitSkill implements a skill
* that makes the HEXA sit and stand
 */

package standsitskill

import (
	"mind/core/framework/drivers/hexabody"
	"mind/core/framework/log"
	"mind/core/framework/skill"
	"os"
)

type StandSitSkill struct {
	skill.Base
}

func NewSkill() skill.Interface {
	return &StandSitSkill{}
}

func (d *StandSitSkill) stand() {
	hexabody.Stand()
}

func (d *StandSitSkill) sit() {
	hexabody.RelaxLegs()
}

func (d *StandSitSkill) OnStart() {
	log.Info.Println("STARTING!!!!")
	hexabody.Start()
}

func (d *StandSitSkill) OnClose() {
	hexabody.Close()
}

func (d *StandSitSkill) OnConnect() {
	log.Info.Println("CONNECTED!!!!")
}

func (d *StandSitSkill) OnDisconnect() {
	log.Info.Println("Terminating skill")
	os.Exit(0) // Closes the process when remote disconnects
}

func (d *StandSitSkill) OnRecvString(data string) {
	log.Info.Println("Received string!!!!")
	switch data {
	case "stand":
		hexabody.Stand()
	case "sit":
		hexabody.RelaxLegs()
	}
}
