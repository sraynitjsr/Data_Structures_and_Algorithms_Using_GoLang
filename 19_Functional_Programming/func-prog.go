package func_prog

type MapperFunc func(int) int

func Map(slice []int, mapper MapperFunc) []int {
   result := make([]int, len(slice))
   for i, v := range slice {
      result[i] = mapper(v)
   }
   return result
}
