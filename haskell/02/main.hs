import System.Environment

fibs :: [Integer]
fibs = 0 : 1 : zipWith (+) fibs (tail fibs)


main :: IO ()
main = do args <- getArgs
          let limit = (read $ head args :: Integer)
          let e2 = sum[n | n <- takeWhile (< limit) fibs, even n]
          print e2
