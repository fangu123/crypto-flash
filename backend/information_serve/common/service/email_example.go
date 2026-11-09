package service

import (
	"fmt"
)

// SendPaymentNotification 发送支付通知邮件示例
func SendPaymentNotification(toEmail, orderNo, amount string) error {
	subject := fmt.Sprintf("支付通知 - 订单号: %s", orderNo)

	htmlBody := fmt.Sprintf(`
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>支付通知</title>
</head>
<body>
    <div style="font-family: Arial, sans-serif; max-width: 600px; margin: 0 auto; padding: 20px;">
        <h2 style="color: #333;">支付成功通知</h2>
        <p>尊敬的用户：</p>
        <p>您的订单已成功支付，详细信息如下：</p>
        <table style="width: 100%%; border-collapse: collapse; margin: 20px 0;">
            <tr>
                <td style="padding: 10px; border: 1px solid #ddd; background-color: #f5f5f5;"><strong>订单号</strong></td>
                <td style="padding: 10px; border: 1px solid #ddd;">%s</td>
            </tr>
            <tr>
                <td style="padding: 10px; border: 1px solid #ddd; background-color: #f5f5f5;"><strong>支付金额</strong></td>
                <td style="padding: 10px; border: 1px solid #ddd;">%s</td>
            </tr>
            <tr>
                <td style="padding: 10px; border: 1px solid #ddd; background-color: #f5f5f5;"><strong>支付时间</strong></td>
                <td style="padding: 10px; border: 1px solid #ddd;">%s</td>
            </tr>
        </table>
        <p>如有任何疑问，请联系客服。</p>
        <p style="color: #999; font-size: 12px;">此邮件由系统自动发送，请勿直接回复。</p>
    </div>
</body>
</html>`, orderNo, amount, getCurrentTime())

	return SendHTMLEmail([]string{toEmail}, subject, htmlBody)
}

// SendInvoiceReminder 发送发票提醒邮件示例
func SendInvoiceReminder(toEmail, invoiceNo string) error {
	subject := fmt.Sprintf("发票提醒 - 发票号: %s", invoiceNo)

	textBody := fmt.Sprintf(`尊敬的用户：

您的发票 %s 已生成，请及时查收。

如有任何疑问，请联系客服。

此邮件由系统自动发送，请勿直接回复。`, invoiceNo)

	return SendTextEmail([]string{toEmail}, subject, textBody)
}

// getCurrentTime 获取当前时间字符串
func getCurrentTime() string {
	return fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d",
		2024, 1, 1, 12, 0, 0) // 示例时间，实际应使用 time.Now()
}
