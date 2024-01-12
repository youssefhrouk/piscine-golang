package piscine

func Atoi(s string) int {
  rslt := 0
  sign := 1
  for lengh, i := range s {
    if i == '+' {
      if lengh == 0 {
        sign = 1 
      } else {
        return 0
      }
    } else if i == '-' {
      if lengh == 0 {
        sign = -1 
      } else {
        return 0
      }
    }  else if i >= '0' && i <= '9' {
      rslt = rslt*10 + int (i - '0')
    } else {
      return 0
    }
  }
  return rslt*sign
}