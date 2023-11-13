package func_prog

type MapperFunc func(int) int
type PredicateFunc func(int) bool
type ReducerFunc func(int, int) int

func Map(slice []int, mapper MapperFunc) []int {
   result := make([]int, len(slice))
   for i, v := range slice {
      result[i] = mapper(v)
   }
   return result
}

func Filter(slice []int, predicate PredicateFunc) []int {
   var result []int
   for _, v := range slice {
      if predicate(v) {
         result = append(result, v)
      }
   }
   return result
}

func Reduce(slice []int, reducer ReducerFunc, initial int) int {
   result := initial
   for _, v := range slice {
      result = reducer(result, v)
   }
   return result
}
