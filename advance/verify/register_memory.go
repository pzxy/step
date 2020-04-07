package main

func main() {

}
//理解内存性能
//func write_read(long *src, long *dst, long n) {
//	long
//	cnt = n;
//	long
//	val = 0;
//	while(cnt)
//	{
//		*dst = val;
//		val = (*src) + 1;
//		cnt --;
//	}
//}

//write_read(&a[0],&a[1],3)//更快
//write_read(&a[0],&a[0],3)//慢，因为要对相同地址读写操作，下一个操作必须等上一个操作完成