import System.Environment


main = do args <- getArgs
          let maxNum = read (head args ) :: Integer
          let new = sum [1..maxNum]^2 - sum ( map (^2) [1..maxNum])
          print new