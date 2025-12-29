    
    // 前序 Preorder: 根→左→右
    // 中序 Inorder: 左→根→右
    // 後序 Postorder: 左→右→根
    // BST 走中序的時候 (因為BST特性的關系)
    //   - 會拿到由小到大排序好的數字其中每次取root值會越取越大
    //   - 所以想要取下一個值(Successor: 最小且>當前值)可以參考中遍歷的方法
    //      - 也就是中值拿完之後的動作：取右子節點一次，然後左節點走到底再取值 
    //      - 也可以思考是比當前大，所以取Right，然後拿最小的值，所以Left走到底
    //      - `r = r.Right; for r.Left != nil { r = r.Left }; return root`
    //      - 若沒有.Right 則往祖先看，找>val的
    //   - 所以想要取上一個值(Predecessor: 最大且<當前值)可以參考中遍歷的方法
    //      - 也就是中值拿完之前的動作：取左子節點一次，然後右節點走到底再取值
    //      - 也可以思考是比當前小，所以取Left，然後拿最大的值，所以Right走到底
    //      - `r = r.Left; for r.Right != nil { r = r.Right }; return root`
    //      - 若沒有.Left 則往祖先看，找<val的
