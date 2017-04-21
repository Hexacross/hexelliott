/*
*   Template for HEXA skills.
*   @Author: Elliott Ploutz
*
*   Note: you may need to run sudo for these commands.
*   mind scan    // Finds a HEXA on your network.
*   mind set-default-robot [HEXAname]
*   mind init templateSkill
*   cd templateSkill
*   mind run      // Uploads the skill to HEXA
*
*   Now edit the skill, save, build, and run.
*   mind build && mind pack && mind run
*
*   HEXA should now stand up as soon as the skill is started.
*
*
*/

import (
  "mind/core/framework/drivers/hexabody"
  "mind/core/framework/skill"
  "mind/core/framework/drivers/distance"  // Distance sensor
  "mind/core/framework/log"
  "time"
  "math"
)

func (d *templateSkill) OnStart() {
  hexabody.Start()
  hexabody.Stand()
  distance.Start()
}

func (d *templateSkill) OnClose() {
  hexabody.Close()
  distance.Close()
}

// The OnConnect function is called whenever a remote (browser, mobile app or terminal)
// connects to, and disconnects from the Skill
func (d *templateSkill) OnConnect() {
  // HEXA doesn't know which way it's facing (the position of its head).
  // This tells HEXA to move it's head to position 0.
  // The second argument is the duration of the animation, but since HEXA currently does
  // not know how far away position 0 is, it won't be able to move there in a given time.
  // However, for all future head movements it will.
  hexabody.MoveHead(0, 0)
  // HEXA walks at 0.5 cm per second. The fastest speed is 1.3-1.5 cm/s.
  // 0 tells HEXA to walk in the direction it's facing.
  hexabody.WalkContinuously(0, 0.5)

  // This measures distance twice per second.
  // It should be done with Goroutines and channels, but for simplicity
  // we use a for loop for now without error handling code.
  for {
    dist, _ := distance.Value()
    log.Info.Println("Distance in millimeters: ", dist)
    if dist < 500 {                                 // Object near.
      hexabody.StopWalkingContinuously()            // Stop walking.
      hexabody.Relax()
      time.Sleep(2 * time.Second)                   // Sit and relax for 2 seconds.
      direction = newDirection(direction, 180)      // Calculate a new direction (180 degrees).
      hexabody.MoveHead(direction, 0)
      time.Sleep(time.Second)
      hexabody.WalkContinuously(0, 0.5)             // Start walking.
      time.Sleep(time.Second)
    }
    time.Sleep(500 * time.Millisecond)
  }
}

// Tell HEXA to stop!
func (d *templateSkill) OnDisconnect() {
  hexabody.StopWalkingContinuously()
  hexabody.Relax()
}

// For when HEXA's distance measure shows an object too close.
// We can turn clockwise or counter-clockwise with each call.
func newDirection(direction float64, degrees float64) float64 {
  return math.Mod(direction+degrees, 360) * -1
}
