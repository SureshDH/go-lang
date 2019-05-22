package math

func Avarage(marks []float32) float32{
    avg := 0.0
    for _, value := range marks{
        avg += value
    }
    return avg/len(marks)
}
