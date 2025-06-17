# 文档(Docs)

```

```

## 查看lvm磁盘占用

```bash
# 查看
df -h

fdisk -l

# 查看磁盘挂载情况
lsblk
# 查看lvm卷组的信息，如果看到 Free PE / Size > 0，表示还有扩容空间。
vgdisplay


# 扩容
#lvextend -L 10G /dev/mapper/ubuntu--vg-ubuntu--lv      //增大或减小至19G
#lvextend -L +10G /dev/mapper/ubuntu--vg-ubuntu--lv     //增加10G
#lvreduce -L -10G /dev/mapper/ubuntu--vg-ubuntu--lv     //减小10G
#lvresize -l  +100%FREE /dev/mapper/ubuntu--vg-ubuntu--lv   //按百分比扩

# 扩容
sudo lvresize -l  +100%FREE /dev/mapper/ubuntu--vg-ubuntu--lv
# 生效
sudo resize2fs /dev/mapper/ubuntu--vg-ubuntu--lv