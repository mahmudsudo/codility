package lessons
func Solution(A []int, K int) []int {
    // write your code in Go 1.4
    // 13458
    
    
     

    for K > 0{
        
       newArr := make([]int,len(A))
     copy(newArr,A)
        for i:=0;i<=len(A)-1;i++{
            
            if i == 0 {
                A[i]=newArr[len(newArr)-1]

            }else {
               
                A[i]=newArr[i-1]
                
            }
        }
          K-- 
        }
     
    
    return A
}
