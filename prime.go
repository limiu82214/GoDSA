func getPrime(maximum int) []bool {
    if maximum <2 {return []int{} }
    isPrime := make([]bool, maximum+1)
    for i:=2; i<= maximum; i++ { isPrime[i] = true}

    for i:=2; i*i<=maximum; i++ {
        if isPrime[i] {
            for j := i*i; j<=maximum; j+=i {
                isPrime[j] = false
            }
        }
    }

    return isPrime
}

