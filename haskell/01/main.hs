main :: IO ()
main = do
  let val = sum [3, 6 .. 999] + sum [5, 10 .. 999] - sum [15, 30 .. 999]
  print val