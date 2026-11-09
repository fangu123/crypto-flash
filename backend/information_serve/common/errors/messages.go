package errors

var (
	ErrorMessages = map[string]map[string]string{
		"apikey_not_found": {
			"cn": "没有apikey",
			"en": "api key not found",
		},
		"point_not_found": {
			"cn": "没有点位",
			"en": "no trading point found",
		},
		"coin_not_found": {
			"cn": "没有币种",
			"en": "coin not found",
		},
		"not_login": {
			"cn": "未登录",
			"en": "not login in",
		},
		"acquire_success": {
			"cn": "获取成功",
			"en": "acquire success",
		},
		"acquire_fail": {
			"cn": "获取失败",
			"en": "acquire fail",
		},
		//获取市场概览数据失败
		"get_market_overview_fail": {
			"cn": "获取市场概览数据失败",
			"en": "get market overview data fail",
		},
		//获取多空比数据失败
		"get_market_overview_ratio_fail": {
			"cn": "获取多空比数据失败",
			"en": "get market overview ratio data fail",
		},
		//获取财经日历数据失败
		"get_finance_calendar_fail": {
			"cn": "获取财经日历数据失败",
			"en": "get finance calendar data fail",
		},
		//获取新闻快讯数据失败
		"get_news_flash_fail": {
			"cn": "获取资讯快讯数据失败",
			"en": "get news flash data fail",
		},
		//获取最新交易点位失败
		"get_latest_point_fail": {
			"cn": "获取最新交易点位失败",
			"en": "get latest trading point fail",
		},
		//获取新闻列表失败
		"get_news_list_fail": {
			"cn": "获取资讯列表失败",
			"en": "get news list fail",
		},
		//获取新闻详情失败
		"get_news_detail_fail": {
			"cn": "获取资讯详情失败",
			"en": "get news detail fail",
		},
		//新闻ID无效
		"news_id_invalid": {
			"cn": "资讯ID无效",
			"en": "news id invalid",
		},
		//分享新闻失败
		"share_news_fail": {
			"cn": "分享资讯失败",
			"en": "share news fail",
		},
		//分享新闻成功
		"share_news_success": {
			"cn": "分享资讯成功",
			"en": "share news success",
		},
		//收藏成功
		"collect_success": {
			"cn": "收藏成功",
			"en": "collect success",
		},
		//取消收藏成功
		"cancel_collect_success": {
			"cn": "取消收藏成功",
			"en": "cancel collect success",
		},
		//资讯不存在
		"news_not_exist": {
			"cn": "资讯不存在",
			"en": "news not exist",
		},
		//用户ID无效
		"user_id_invalid": {
			"cn": "用户ID无效",
			"en": "user id invalid",
		},
		//已经收藏过该新闻
		"news_already_collected": {
			"cn": "已经收藏过该新闻",
			"en": "news already collected",
		},
		//点位ID无效
		"point_id_invalid": {
			"cn": "点位ID无效",
			"en": "point id invalid",
		},
		//分享点位失败
		"share_point_fail": {
			"cn": "分享点位失败",
			"en": "share point fail",
		},
		//分享点位成功
		"share_point_success": {
			"cn": "分享点位成功",
			"en": "share point success",
		},
		//获取点位列表失败
		"get_point_list_fail": {
			"cn": "获取点位列表失败",
			"en": "get point list fail",
		},
		//获取点位详情失败
		"get_point_detail_fail": {
			"cn": "获取点位详情失败",
			"en": "get point detail fail",
		},
		//收藏点位失败
		"collect_point_fail": {
			"cn": "收藏点位失败",
			"en": "collect point fail",
		},
		//取消收藏点位失败
		"cancel_collect_point_fail": {
			"cn": "取消收藏点位失败",
			"en": "cancel collect point fail",
		},
		//解锁全部点位失败
		"unlock_all_point_fail": {
			"cn": "解锁全部点位失败",
			"en": "unlock all point fail",
		},
		//解锁成功
		"unlock_success": {
			"cn": "解锁成功",
			"en": "unlock success",
		},
		//点位不存在
		"point_not_exist": {
			"cn": "点位不存在",
			"en": "point not exist",
		},
		//已经收藏过该点位
		"point_already_collected": {
			"cn": "已经收藏过该点位",
			"en": "point already collected",
		},
		//用户不存在
		"user_not_exist": {
			"cn": "用户不存在",
			"en": "user not exist",
		},
		//NFT API配置未找到
		"nft_api_config_not_found": {
			"cn": "NFT API配置未找到",
			"en": "nft api config not found",
		},
		//parameter error
		"parameter_error": {
			"cn": "参数错误",
			"en": "parameter error",
		},
		//the email format is incorrect
		"the_email_format_is_incorrect": {
			"cn": "邮箱格式不正确",
			"en": "the email format is incorrect",
		},
		//the password format is incorrect
		"the_password_format_is_incorrect": {
			"cn": "密码格式不正确",
			"en": "the password format is incorrect",
		},
		//密码长度至少12位
		"password_length_at_least_12": {
			"cn": "密码长度至少12位",
			"en": "password length at least 12",
		},
		//密码需包含大写字母、小写字母、数字、特殊字符
		"password_must_contain_uppercase_lowercase_digit_special_character": {
			"cn": "密码需包含大写字母、小写字母、数字、特殊字符",
			"en": "password must contain uppercase, lowercase, digit, special character",
		},
		//registered successfully
		"registered_successfully": {
			"cn": "注册成功",
			"en": "registered successfully",
		},
		//the email has been registered
		"the_email_has_been_registered": {
			"cn": "邮箱已被注册",
			"en": "the email has been registered",
		},
		//exit successful
		"exit_successful": {
			"cn": "退出成功",
			"en": "exit successful",
		},
		//please send the verification code
		"please_send_the_verification_code": {
			"cn": "请发送验证码",
			"en": "please send the verification code",
		},
		//invalid verification code
		"invalid_verification_code": {
			"cn": "验证码无效",
			"en": "invalid verification code",
		},
		//update successfully
		"update_successfully": {
			"cn": "更新成功",
			"en": "update successfully",
		},
		//update failed
		"update_failed": {
			"cn": "更新失败",
			"en": "update failed",
		},
		//verification code sent successfully
		"verification_code_sent_successfully": {
			"cn": "验证码发送成功",
			"en": "verification code sent successfully",
		},
		//password reset successfully
		"password_reset_successfully": {
			"cn": "密码重置成功",
			"en": "password reset successfully",
		},
		//获取消息列表失败
		"get_message_list_fail": {
			"cn": "获取消息列表失败",
			"en": "get message list fail",
		},
		//消息ID无效
		"message_id_invalid": {
			"cn": "消息ID无效",
			"en": "message id invalid",
		},
		//标记消息已读失败
		"mark_message_read_fail": {
			"cn": "标记消息已读失败",
			"en": "mark message read fail",
		},
		//标记成功
		"mark_success": {
			"cn": "标记成功",
			"en": "mark success",
		},
		//获取收藏列表失败
		"get_collection_list_fail": {
			"cn": "获取收藏列表失败",
			"en": "get collection list fail",
		},
		//获取联系方式失败
		"get_contact_information_fail": {
			"cn": "获取联系方式失败",
			"en": "get contact information fail",
		},
		//获取关于我们版本失败
		"get_version_information_fail": {
			"cn": "获取版本信息失败",
			"en": "get version information fail",
		},
		//获取关于我们版本成功
		"get_about_us_version_success": {
			"cn": "获取关于我们版本成功",
			"en": "get about us version success",
		},
		//failed to set 2FA
		"failed_to_set_2FA": {
			"cn": "设置2FA失败",
			"en": "failed to set 2FA",
		},
		//success to set 2FA
		"success_to_set_2FA": {
			"cn": "设置2FA成功",
			"en": "success to set 2FA",
		},
		//email parameter is required
		"email_parameter_is_required": {
			"cn": "邮件参数必传",
			"en": "email parameter is required",
		},
		//绑定钱包地址失败
		"bind_wallet_address_fail": {
			"cn": "绑定钱包地址失败",
			"en": "bind wallet address fail",
		},
		//绑定成功
		"bind_success": {
			"cn": "绑定成功",
			"en": "bind success",
		},
		//获取消息通知设置失败
		"get_message_setting_fail": {
			"cn": "获取消息通知设置失败",
			"en": "get message setting fail",
		},
		//更新消息通知设置失败
		"update_message_setting_fail": {
			"cn": "更新消息通知设置失败",
			"en": "update message setting fail",
		},
		//new password and confirm password do not match
		"new_password_and_confirm_password_do_not_match": {
			"cn": "新密码和确认密码不匹配",
			"en": "new password and confirm password do not match",
		},
		//email does not exist
		"email_does_not_exist": {
			"cn": "邮箱不存在",
			"en": "email does not exist",
		},
		//Verification Code
		"verification_code": {
			"cn": "验证码",
			"en": "Verification Code",
		},
		"email_information_before": {
			"cn": "您的验证码为：",
			"en": "Your verification code is:",
		},
		"email_information_after": {
			"cn": "，有效期为5分钟，请尽快使用。",
			"en": "This code will expire in 5 minutes.",
		},
		//failed to generate 2FA key
		"failed_to_generate_2FA_key": {
			"cn": "生成2FA密钥失败",
			"en": "failed to generate 2FA key",
		},
		//钱包地址不能为空
		"wallet_address_cannot_be_empty": {
			"cn": "钱包地址不能为空",
			"en": "wallet address cannot be empty",
		},
		//无效的消息类型
		"invalid_message_type": {
			"cn": "无效的消息类型",
			"en": "invalid message type",
		},
		//无效的状态值
		"invalid_status_value": {
			"cn": "无效的状态值",
			"en": "invalid status value",
		},
		//更新通知设置失败
		"update_notification_settings_fail": {
			"cn": "更新通知设置失败",
			"en": "update notification settings fail",
		},
		//message not found
		"message_not_found": {
			"cn": "消息未找到",
			"en": "message not found",
		},
		//please binding 2FA verification
		"please_binding_2FA_verification": {
			"cn": "请绑定2FA验证",
			"en": "please binding 2FA verification",
		},
		//please enter the Google verification code
		"please_enter_the_Google_verification_code": {
			"cn": "请输入Google验证码",
			"en": "please enter the Google verification code",
		},
		//google verification code error
		"google_verification_code_error": {
			"cn": "Google验证码错误",
			"en": "google verification code error",
		},
		//please enter the email code
		"please_enter_the_email_code": {
			"cn": "请输入邮箱验证码",
			"en": "please enter the email code",
		},
		//please enter the correct type
		"please_enter_the_correct_type": {
			"cn": "请输入正确的类型",
			"en": "please enter the correct type",
		},
		//failed to enable 2FA
		"failed_to_enable_2FA": {
			"cn": "启用2FA失败",
			"en": "failed to enable 2FA",
		},
		//已解锁点位
		"unlock_point": {
			"cn": "已解锁点位",
			"en": "unlock point",
		},
		//incorrect Username or Password
		"incorrect_username_or_password": {
			"cn": "用户名或密码错误",
			"en": "incorrect Username or Password",
		},
		//wallet_address_already_bound
		"wallet_address_already_bound": {
			"cn": "钱包地址已绑定",
			"en": "wallet address already bound",
		},
		//keyword required
		"keyword_required": {
			"cn": "搜索关键词不能为空",
			"en": "keyword is required",
		},
		//search fail
		"search_fail": {
			"cn": "搜索失败",
			"en": "search fail",
		},
		//coins not found
		"coins_not_found": {
			"cn": "未找到相关币种",
			"en": "coins not found",
		},
		//financial news not found
		"financial_news_not_found": {
			"cn": "未找到相关财经新闻",
			"en": "financial news not found",
		},
		//news flash not found
		"news_flash_not_found": {
			"cn": "未找到相关新闻快讯",
			"en": "news flash not found",
		},
		//trading points not found
		"trading_points_not_found": {
			"cn": "未找到相关交易点位",
			"en": "trading points not found",
		},
		//invalid wallet address
		"invalid_wallet_address": {
			"cn": "无效的钱包地址",
			"en": "invalid wallet address",
		},
		//user creation failed
		"user_creation_failed": {
			"cn": "用户创建失败",
			"en": "user creation failed",
		},
		//token_refreshed_failed
		"token_refreshed_failed": {
			"cn": "刷新token失败",
			"en": "token refreshed failed",
		},
		//token_refreshed_successfully
		"token_refreshed_successfully": {
			"cn": "刷新token成功",
			"en": "token refreshed successfully",
		},
		//bind_failed
		"bind_failed": {
			"cn": "绑定失败",
			"en": "bind failed",
		},
		"unbind_failed": {
			"cn": "解绑失败",
			"en": "unbind failed",
		},
		"unbind_success": {
			"cn": "解绑成功",
			"en": "unbind success",
		},
		//检查账号绑定状态失败
		"check_account_binding_status_failed": {
			"cn": "检查账号绑定状态失败",
			"en": "check account binding status failed",
		},
		//该Google账号已被其他用户绑定
		"google_account_has_been_bound_by_other_users": {
			"cn": "该Google账号已被其他用户绑定",
			"en": "google account has been bound by other users",
		},
		//绑定Google账号失败
		"bind_google_account_failed": {
			"cn": "绑定Google账号失败",
			"en": "bind google account failed",
		},
		//该Microsoft账号已被其他用户绑定
		"microsoft_account_has_been_bound_by_other_users": {
			"cn": "该Microsoft账号已被其他用户绑定",
			"en": "microsoft account has been bound by other users",
		},
		//绑定Microsoft账号失败
		"bind_microsoft_account_failed": {
			"cn": "绑定Microsoft账号失败",
			"en": "bind microsoft account failed",
		},
		//未绑定Google账号
		"not_bound_google_account": {
			"cn": "未绑定Google账号",
			"en": "not bound google account",
		},
		//解绑Google账号失败
		"unbind_google_account_failed": {
			"cn": "解绑Google账号失败",
			"en": "unbind google account failed",
		},
		//未绑定Microsoft账号
		"not_bound_microsoft_account": {
			"cn": "未绑定Microsoft账号",
			"en": "not bound microsoft account",
		},
		//解绑Microsoft账号失败
		"unbind_microsoft_account_failed": {
			"cn": "解绑Microsoft账号失败",
			"en": "unbind microsoft account failed",
		},
		//获取用户信息失败
		"get_user_info_failed": {
			"cn": "获取用户信息失败",
			"en": "get user info failed",
		},
		//invalid_redirect_url
		"invalid_redirect_url": {
			"cn": "无效的重定向URL",
			"en": "invalid redirect url",
		},
		//获取Google认证URL失败
		"get_google_auth_url_failed": {
			"cn": "获取Google认证URL失败",
			"en": "get google auth url failed",
		},
		//获取Microsoft认证URL失败
		"get_microsoft_auth_url_failed": {
			"cn": "获取Microsoft认证URL失败",
			"en": "get microsoft auth url failed",
		},
		//invalid_state
		"invalid_state": {
			"cn": "无效的状态",
			"en": "invalid state",
		},
		//无效的状态参数
		"invalid_state_parameter": {
			"cn": "无效的状态参数",
			"en": "invalid state parameter",
		},
		//token_exchange_failed
		"token_exchange_failed": {
			"cn": "令牌交换失败",
			"en": "token exchange failed",
		},
		//获取Google访问令牌失败
		"get_google_access_token_failed": {
			"cn": "获取Google访问令牌失败",
			"en": "get google access token failed",
		},
		//获取Microsoft访问令牌失败
		"get_microsoft_access_token_failed": {
			"cn": "获取Microsoft访问令牌失败",
			"en": "get microsoft access token failed",
		},
		//获取Google用户信息失败
		"get_google_user_info_failed": {
			"cn": "获取Google用户信息失败",
			"en": "get google user info failed",
		},
		//解析Google用户信息失败
		"parse_google_user_info_failed": {
			"cn": "解析Google用户信息失败",
			"en": "parse google user info failed",
		},
		//获取Microsoft用户信息失败
		"get_microsoft_user_info_failed": {
			"cn": "获取Microsoft用户信息失败",
			"en": "get microsoft user info failed",
		},
		//解析Microsoft用户信息失败
		"parse_microsoft_user_info_failed": {
			"cn": "解析Microsoft用户信息失败",
			"en": "parse microsoft user info failed",
		},
		//上传成功
		"upload_success": {
			"cn": "上传成功",
			"en": "upload success",
		},
		//初始化文件路径失败
		"init_file_path_failed": {
			"cn": "初始化文件路径失败",
			"en": "init file path failed"},
		//上传第三方失败
		"upload_third_party_failed": {
			"cn": "上传第三方失败",
			"en": "upload third party failed",
		},
		//图片不能为空
		"image_cannot_be_empty": {
			"cn": "图片不能为空",
			"en": "image cannot be empty",
		},
		//资讯ID无效
		"invalid_news_id": {
			"cn": "资讯ID无效",
			"en": "invalid news id",
		},
		//父评论ID无效
		"invalid_parent_comment_id": {
			"cn": "父评论ID无效",
			"en": "invalid parent comment id",
		},
		//评论发布成功
		"comment_published_successfully": {
			"cn": "评论发布成功",
			"en": "comment published successfully",
		},
		//评论ID无效
		"invalid_comment_id": {
			"cn": "评论ID无效",
			"en": "invalid comment id",
		},
		//评论删除成功
		"comment_deleted_successfully": {
			"cn": "评论删除成功",
			"en": "comment deleted successfully",
		},
		//点赞成功
		"like_success": {
			"cn": "点赞成功",
			"en": "like success",
		},
		//取消点赞成功
		"cancel_like_success": {
			"cn": "取消点赞成功",
			"en": "cancel like success",
		},
		//评论不存在或无权删除
		"comment_does_not_exist_or_no_permission_to_delete": {
			"cn": "评论不存在或无权删除",
			"en": "comment does not exist or no permission to delete",
		},
		//评论不存在或点赞状态无效
		"comment_does_not_exist_or_like_status_invalid": {
			"cn": "评论不存在或点赞状态无效",
			"en": "comment does not exist or like status invalid",
		},
		//评论内容不能为空
		"comment_content_cannot_be_empty": {
			"cn": "评论内容不能为空",
			"en": "comment content cannot be empty",
		},
		//评论内容不能超过2000个字符
		"comment_content_cannot_be_more_than_2000_characters": {
			"cn": "评论内容不能超过2000个字符",
			"en": "comment content cannot be more than 2000 characters",
		},
		//资讯不存在或未发布
		"news_does_not_exist_or_not_published": {
			"cn": "资讯不存在或未发布",
			"en": "news does not exist or not published",
		},
		//回复的评论不存在
		"reply_to_comment_does_not_exist": {
			"cn": "回复的评论不存在",
			"en": "reply to comment does not exist",
		},
		//点位未解锁，请前往解锁后查看
		"point_not_unlocked": {
			"cn": "点位未解锁，请前往解锁后查看",
			"en": "This point has not been unlocked. Please go to unlock it before viewing.",
		},
	}
)

func GetErrorMessage(key, lang string) string {
	if msg, ok := ErrorMessages[key][lang]; ok {
		return msg
	}
	return ErrorMessages[key]["cn"] // 默认返回中文
}
