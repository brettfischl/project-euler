import System.Environment

main = do args <- getArgs
          let limit = read $ head args :: Integer
          let max = foldl1 lcm [1..limit]
          print max