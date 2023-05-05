package godie
/*
 * goDie -- Simple dice macros for creating systems
 * 
 * when importing this use the pattern import roll "github.com/StephanieSunshine/godie"
 * 
 * MIT License -- Stephanie Sunshine 2023
 */

import (
  "math/rand"
)

func d4() uint {
  return Roll(4)
}

func d6() uint {
  return Roll(6)
}

func d8() uint {
  return Roll(8)
}

func d10() uint {
  return Roll(10)
}

func d12() uint {
  return Roll(12)
}

func d20() uint {
  return Roll(20)
}

func d100() uint {
  return Roll(100)
}

func Roll(max uint) uint {
  return (uint(rand.Uint32()) % max)+1
}
