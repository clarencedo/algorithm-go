### 子集

假设 nums = [1,2,3]，那么回溯的过程如下：

```
index=0, current=[]       → 加入结果 [[]]
  ├── index=1, current=[1] → 加入结果 [[], [1]]
  │    ├── index=2, current=[1,2] → 加入结果 [[], [1], [1,2]]
  │    │    ├── index=3, current=[1,2,3] → 加入结果 [[], [1], [1,2], [1,2,3]]
  │    │    └── 回溯，current=[1,2]
  │    ├── index=3, current=[1,3] → 加入结果 [[], [1], [1,2], [1,2,3], [1,3]]
  │    └── 回溯，current=[1]
  ├── index=2, current=[2] → 加入结果 [[], [1], [1,2], [1,2,3], [1,3], [2]]
  │    ├── index=3, current=[2,3] → 加入结果 [[], [1], [1,2], [1,2,3], [1,3], [2], [2,3]]
  │    └── 回溯，current=[2]
  ├── index=3, current=[3] → 加入结果 [[], [1], [1,2], [1,2,3], [1,3], [2], [2,3], [3]]
  └── 回溯，current=[]
```

最终得到：

```
[[], [1], [1,2], [1,2,3], [1,3], [2], [2,3], [3]]
```
