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

func D4() uint {
  return Roll(4)
}

func D6() uint {
  return Roll(6)
}

func D8() uint {
  return Roll(8)
}

func D10() uint {
  return Roll(10)
}

func D12() uint {
  return Roll(12)
}

func D20() uint {
  return Roll(20)
}

func D100() uint {
  return Roll(100)
}

func Roll(max uint) uint {
  return (uint(rand.Uint32()) % max)+1
}
