package bean

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/24 17:02
 * @Description TODO
 * 任务发布者 bean
 */

type FileUrlResp struct {
	Code int      `json:"code"`
	Msg  string   `json:"msg"`
	Data FileInfo `json:"data"`
}

type FileInfo struct {
	Download string  `json:"download"`
	Name     string  `json:"name"`
	Size     int64   `json:"size"`
	Edition  float32 `json:"edition"`
}
