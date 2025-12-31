package mailer

import (
	"fmt"
	"os"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

func SendVerificationLinkToEmployee(email_add string, rawToken string) (bool, error) {

	appURL := os.Getenv("FE_APP_URL")
	verifyURL := fmt.Sprintf(
		"%s/org/employee-verification?token=%s",
		appURL,
		rawToken,
	)

	// Build HTML message with a light background and highlighted OTP
	htmlBody := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<meta name="viewport" content="width=device-width, initial-scale=1.0">
	<title>Verify Your Email - Sharpnr</title>
</head>
<body style="margin:0; padding:0; background-color:#f8fafc; font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',system-ui,sans-serif;">

<table width="100%%" cellpadding="0" cellspacing="0">
<tr>
<td align="center" style="padding:40px 20px;">

<table style="max-width:520px; width:100%%; background:#ffffff; border-radius:16px; box-shadow:0 4px 20px rgba(0,0,0,0.08);">

<!-- Header -->
<tr>
<td style="background:linear-gradient(135deg,#6366f1 0%%,#8b5cf6 100%%); padding:40px; text-align:center; border-radius:16px 16px 0 0;">
	<h1 style="color:#ffffff; margin:0; font-size:26px;">Sharpnr</h1>
	<p style="color:rgba(255,255,255,0.9); margin-top:8px;">Employee Onboarding</p>
</td>
</tr>

<!-- Content -->
<tr>
<td style="padding:40px 35px; text-align:center;">
	<h2 style="color:#1f2937; margin-bottom:12px;">Verify your email address</h2>

	<p style="color:#6b7280; font-size:16px; line-height:1.6;">
		You’ve been added as an employee in <strong>Sharpnr</strong>.
		Please confirm your email to activate your account.
	</p>

	<!-- Button -->
	<div style="margin:40px 0;">
		<a href="%s"
		   style="background:#6366f1; color:#ffffff; text-decoration:none; padding:14px 28px;
		          border-radius:10px; font-size:16px; font-weight:600; display:inline-block;">
			Verify Email
		</a>
	</div>

	<p style="color:#9ca3af; font-size:13px;">
		This link will expire in 24 hours.
	</p>

	<!-- Security note -->
	<div style="background:#fef3c7; border-left:4px solid #f59e0b;
	            padding:14px 18px; margin-top:30px; text-align:left;">
		<p style="color:#92400e; font-size:13px; margin:0;">
			If you didn’t expect this invitation, you can safely ignore this email.
		</p>
	</div>
</td>
</tr>

<!-- Footer -->
<tr>
<td style="background:#f9fafb; padding:25px; text-align:center; border-top:1px solid #e5e7eb;">
	<p style="color:#9ca3af; font-size:12px; margin:0;">
		© 2025 Sharpnr. All rights reserved.
	</p>
</td>
</tr>

</table>

</td>
</tr>
</table>

</body>
</html>
`, verifyURL)

	// SMTP client configuration
	server := mail.NewSMTPClient()
	server.Host = "smtp.gmail.com"
	server.Port = 587
	server.Username = os.Getenv("OTP_MAILER")
	server.Password = os.Getenv("APP_PASS")
	server.Encryption = mail.EncryptionTLS
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	smtpClient, err := server.Connect()
	if err != nil {
		return false, err
	}

	// Build and send message
	msg := mail.NewMSG()
	msg.SetFrom("Sharpnr Employee OnBoarding <"+server.Username+">").
		AddTo(email_add).
		SetSubject("On-Boarding").
		SetBody(mail.TextHTML, htmlBody)

	if msg.Error != nil {
		return false, msg.Error
	}

	if err := msg.Send(smtpClient); err != nil {
		return false, err
	}

	return true, nil
}
