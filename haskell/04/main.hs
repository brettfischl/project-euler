import System.Environment

main = do 
  let maxPalindrome = maximum [res | a <- [100..999], b <- [100..999], let res = a * b, let r = show res, r == reverse r]
  print maxPalindrome