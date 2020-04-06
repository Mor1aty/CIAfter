package constant

/**
 * @author 16计算机 Moriaty
 * @version 1.0
 * @copyright ：Moriaty 版权所有 © 2020
 * @date 2020/4/4 21:31
 * @Description TODO
 * 自定义常量
 */
const (

	/**
	 * 执行成功
	 */
	CODE_EXEC_SUCCESS = 200

	/**
	 * 获取成功
	 */
	CODE_OBTAIN_SUCCESS = 201

	/**
	 * 获取数据为空
	 */
	CODE_OBTAIN_NULL = 202

	/**
	 * 执行成功
	 */
	CODE_EXEC_FAILURE = 400

	/**
	 * 获取失败
	 */
	CODE_OBTAIN_FAILURE = 401

	/**
	 * 参数校验未通过
	 */
	CODE_PARAM_FAILURE = 402

	/**
	 * Token 校验未通过
	 */
	CODE_TOKEN_FAILURE = 403

	/**
	 * 权限校验未通过
	 */
	CODE_LEVEL_FAILURE = 404

	/**
	 * 请求方法错误
	 */
	CODE_METHOD_FAILURE = 405

	/**
	 * 服务器异常
	 */
	CODE_SERVER_EXCEPTION = 500

	/**
	 * 返回错误信息
	 */
	MSG_ERROR = "服务器错误"

	/**
	 * 返回成功信息
	 */
	MSG_SUCCESS = "执行成功"

	/**
	 * 返回验证未过信息
	 */
	MSG_TOKEN_FAILED = "Token 过期"

	/**
	 * 返回权限未过信息
	 */
	MSG_LEVEL_FAILED = "权限不足"

	/**
	 * 返回邮件成功信息
	 */
	MSG_MAIL_SUCCESS = "邮件发送成功"

	/**
	 * 返回邮件错误信息
	 */
	MSG_MAIL_ERROR = "邮件发送错误"

	/**
	 * 返回无参数信息
	 */
	MSG_PARAM_EMPTY = "缺少参数"

	/**
	 * 返回参数为空信息
	 */
	MSG_PARAM_NULL = " 参数不能为空"

	/**
	 * 返回参数邮件信息
	 */
	MSG_PARAM_EMAIL = " 不是有效邮件格式"

	/**
	 * 返回参数数字信息
	 */
	MSG_PARAM_NUMBER = " 不是有效纯数字格式"

	/**
	 * 返回参数手机信息
	 */
	MSG_PARAM_PHONE = " 不是有效手机格式"

	/**
	 * 日期格式：yyyy-MM-dd HH:mm:ss
	 */
	DATE_TYPE_01 = "2006-01-02 15:04:06"

	/**
	 * 日期格式：yyyy-MM-dd
	 */
	DATE_TYPE_02 = "2006-01-02"

	/**
	 * 日期格式：HH:mm:ss
	 */
	DATE_TYPE_03 = "15:04:06"

	/**
	 * 日期格式：yyyy-MM-dd HH:mm:ss:SSS
	 */
	DATE_TYPE_04 = "2006-01-02 15:04:06:000"

	/**
	 * 日期格式：HH:mm:ss:SSS
	 */
	DATE_TYPE_05 = "15:04:06:000"

	/**
	 * 邮箱正则
	 */
	REGULAR_MAIL = "[\\w\\.\\-]+@([\\w\\-]+\\.)+[\\w\\-]+"

	/**
	 * 纯数字正则
	 */
	REGULAR_NUMBER = "^(0|[1-9][0-9]*|-[1-9][0-9]*)$"

	/**
	 * 整数正则
	 */
	REGULAR_INTEGER = "^[-\\+]?[\\d]*$"

	/**
	 * 浮点数正则
	 */
	REGULAR_FLOAT = "^[-\\+]?[.\\d]*$"

	/**
	 * 手机正则
	 */
	REGULAR_PHONE = "^((13[0-9])|(14[5|7])|(15([0-3]|[5-9]))|(17[013678])|(18[0,5-9]))\\d{8}$"
)
