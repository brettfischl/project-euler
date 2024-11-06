import System.Environment
primes = 2 : filter (null . tail . primeFactors) [3,5..]

primeFactors n = factor n primes
  where
    factor _ [] = []
    factor n (p:ps)
      | p*p > n = [n]
      | n `mod` p == 0 = p : factor (n `div` p) ps
      | otherwise = factor n ps

main = do args <- getArgs
          let limit = (read $ head args :: Integer)
          print $ last (primeFactors limit)


