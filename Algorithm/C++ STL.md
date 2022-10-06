## vector（变长数组），倍增的思想，支持比较运算（按字典序）
```cpp
 定义：
     vector <int> a; 定义：一个vector数组a
     vector <int> a(10); 定义：一个长度为10的vector数组a
     vector <int> a(10,3); 定义：一个长度为10的vector数组a，并且所有元素都为3
常用函数：
     size(); 返回元素个数
     empty(); 返回是否是空
     clear(); 清空
     front(); 返回vector的第一个数
     back(); 返回vector的最后一个数
     push_back(); 向vector的最后插入一个数
     pop_back(); 把vector的最后一个数删掉
     begin(); vector的第0个数
     end(); vector的最后一个的数的后面一个数
倍增的思想：
     系统为某一程序分配空间是，所需时间，与空间大小无关，与申请次数有关
遍历方法：
     假设有个vector <int> a;
     第一种：
         for(int i = 0;i < a.size();i ++) cout<<a[i]<<" ";
     第二种：
         for(vector <int>::iterator i = a.begin();i != a.end();i ++) 
		 cout<<*i<<" "; //vector <int>::iterator可以写为auto
     第三种：
         for(auto  x : a) cout<<x<<" ";
```
## pair，支持比较运算，以first为第一关键字，以second为第二关键字（按字典序）
```c++
定义：
     pair <类型,类型> 变量名;    两个类型可以不同
初始化方式：
     假设有个pair <int,string> p;
     第一种：
         p = make_pair(10,"abc");
     第二种：
         p = {10,"abc");
常用函数：
     first(); 第一个元素
     second(); 第二个元素
```
## string（字符串）
```c++
常用函数：
			size()/length();  返回字符串长度
			empty(); 返回字符串是否为空
			clear(); 清空字符串
			substr(起始下标, (子串长度));  返回子串
			c_str();  返回字符串所在字符数组的起始地址，头指针
			erase(); 
				1、erase(起始下标, 结束下标); 包含起始位置，不包含结束位置，
				2、erase (iterator p);即删除迭代器p指向的字符。
				3、erase((起始位置(0), (长度(string::npos)));删除从起始位置的定长字符串。
				erase()函数的返回值是删除后的字符串。
			find(a); 返回字符串字串a的位置，没有返回string::npos
			string::npos;
			
```
## queue队列
```c++
定义：
     queue <类型> 变量名;
常用函数：
     size(); 这个队列的长度
     empty(); 返回这个队列是否为空
     push(); 往队尾插入一个元素
     front(); 返回队头元素
     back(); 返回队尾元素
     pop(); 把队头弹出
     注意：队列没有clear函数！！！
清空：
     变量名 = queue <int> ();
```
## priority_queue（优先队列，堆）
```c++
注意：默认是大根堆！！！
定义：
     大根堆：priority_queue <类型> 变量名;
     小根堆：priority_queue <类型,vecotr <类型>,greater <类型>> 变量名
常用函数：
     size(); 这个堆的长度
     empty(); 返回这个堆是否为空
     push();往堆里插入一个元素
     top(); 返回堆顶元素
     pop(); 弹出堆顶元素
     注意：堆没有clear函数！！！
```
## stack（栈）
```c++
常用函数：
     size(); 这个栈的长度
     empty(); 返回这个栈是否为空
     push(); 向栈顶插入一个元素
     top(); 返回栈顶元素
     pop(); 弹出栈顶元素
```
## deque（双端队列）
```c++
常用函数：
     size(); 这个双端队列的长度
     empty(); 返回这个双端队列是否为空
     clear(); 清空这个双端队列
     front(); 返回第一个元素
     back(); 返回最后一个元素
     push_back(); 向最后插入一个元素
     pop_back(); 弹出最后一个元素
     push_front(); 向队首插入一个元素
     pop_front(); 弹出第一个元素
     begin(); 双端队列的第0个数
     end(); 双端队列的最后一个的数的后面一个数
```
## set，map，multiset，multimap 基于平衡二叉树（红黑树），动态维护有序序列
###  set/multiset
```c++
注意：set不允许元素重复，如果有重复就会被忽略，但multiset允许！！！
常用函数：
      size(); 返回元素个数
      empty(); 返回set是否是空的
      clear(); 清空
      begin(); 第0个数，支持++或--，返回前驱和后继
      end(); 最后一个的数的后面一个数，支持++或--，返回前驱和后继
      insert(); 插入一个数
      find(); 查找一个数
      count(); 返回某一个数的个数
      erase();
          （1）输入是一个数x，删除所有x    O(k + log n)
          （2）输入一个迭代器，删除这个迭代器
      lower_bound(x); 返回大于等于x的最小的数的迭代器
      upper_bound(x); 返回大于x的最小的数的迭代器
```
### map/multimap
```c++
常用函数：
      insert(); 插入一个数，插入的数是一个pair
      erase(); 
          （1）输入是pair
          （2）输入一个迭代器，删除这个迭代器
      find(); 查找一个数
      lower_bound(x); 返回大于等于x的最小的数的迭代器
      upper_bound(x); 返回大于x的最小的数的迭代器
```
## unordered_set，unordered_map，unordered_muliset,unordered_multimap 基于哈希表
```c++
和上面类似，增删改查的时间复杂度是O(1)
不支持lower_bound()和upper_bound()
```
## bitset 压位
```c++
定义：
     bitset <个数> 变量名;
支持：
     ~，&，|，^
     >>，<<
     ==，!=
     []
常用函数：
     count(); 返回某一个数的个数
     any(); 判断是否至少有一个1
     none(); 判断是否全为0
     set(); 把所有位置赋值为1
     set(k,v); 将第k位变成v
     reset(); 把所有位变成0
     flip(); 把所有位取反，等价于~
     flip(k); 把第k位取反
```
___
