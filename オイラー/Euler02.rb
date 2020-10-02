def fib(n)
  if n == 0 || n == 1
    return 1
  end
  p fib(n - 2) + fib(n - 1)
end


n1, n2 = 1, 2
sequence = []
sequence.push(n1, n2)
# 例外処理

next_number = n1+n2

# フィボナッチ数列を入れまくる
begin
  sequence << next_number
  n1 = n2 # 上の定義に影響及ぼす
  n2 = next_number
  next_number = n1+n2
end while next_number < 4000000

# 偶数を足せ
# 答えの初期化
sum = 0
sequence.each do |number|
  if(number%2 == 0)
    sum += number
  end
end
puts "Answer: " + sum.to_s