package main

import (
	"log"
	"strings"
)

func ArgsParser(flags *Flags, args *[]string) {
	for i := 0; i < len(*args); i++ {
		if flags.ModeOnlyPathsActivated {
			flags.Paths = append(flags.Paths, (*args)[i])
		} else if (*args)[i] == "--" {
			flags.ModeOnlyPathsActivated = true
		} else if len((*args)[i]) >= 2 && (*args)[i][0] == '-' {
			if (*args)[i][1] != '-' {
				splicedArgs := strings.Split((*args)[i], "")
				for _, flag := range splicedArgs[1:] {
					switch flag {
					case "l":
						flags.Dash_l = true
					case "R":
						flags.Dash_R = true
					case "a":
						flags.Dash_A = false
						flags.Dash_a = true
					case "r":
						flags.Dash_r = true
					case "t":
						flags.Dash_S = false
						flags.Dash_t = true
					case "S":
						flags.Dash_t = false
						flags.Dash_S = true
					case "A":
						flags.Dash_a = false
						flags.Dash_A = true
					default:
						log.Fatalf("--%s is not a valid option", flag)
					}
				}
			} else {
				log.Fatalf("%s is not a valid option", (*args)[i])
			}

		} else {
			flags.Paths = append(flags.Paths, (*args)[i])
		}
	}
}
