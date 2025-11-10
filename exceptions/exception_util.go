package exceptions

import "fmt"

func TryCatch[T any](try func() (T, error), catch func(error), finally func()) (T, error) {
	// 默认 catch 和 finally 为无操作函数
	if catch == nil {
		catch = func(err error) {}
	}
	if finally == nil {
		finally = func() {}
	}

	// 使用 defer 确保 finally 被执行
	defer finally()

	// 使用 defer 捕获 panic
	defer func() {
		if r := recover(); r != nil {
			// 如果发生 panic，调用 catch 来处理 panic
			catch(fmt.Errorf("panic caught: %v", r))
		}
	}()

	// 执行 try 代码块
	data, err := try()
	if err != nil {
		// 如果有错误，调用 catch 块处理错误
		catch(err)
		var zero T // 返回类型 T 的零值
		return zero, err
	}

	// 没有错误时，返回结果
	return data, nil
}

func Catch(try func()) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Panic Caught:", r)
		}
	}()
	try()
}
