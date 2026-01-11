package mailer

import (
	"fmt"
	"os"
	"time"

	mail "github.com/xhit/go-simple-mail/v2"
)

// VerifyOtp_Org_Registration generates an OTP and sends a styled HTML email
// to the provided address.  It returns true + the OTP on success.
func VerifyOtp_Org_Registration(email_add string, otp string) error {

	// Build HTML message with a light background and highlighted OTP
	htmlBody := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Sharpnr Verification</title>
</head>
<body style="margin:0; padding:0; background-color:#f8fafc; font-family:-apple-system,BlinkMacSystemFont,'Segoe UI',system-ui,sans-serif;">
    
    <table cellpadding="0" cellspacing="0" border="0" width="100%%">
        <tr>
            <td align="center" style="padding:40px 20px;">
                
                <!-- Main Container -->
                <table cellpadding="0" cellspacing="0" border="0" style="max-width:480px; width:100%%; background:#ffffff; border-radius:16px; box-shadow:0 4px 20px rgba(0,0,0,0.08);">
                    
                    <!-- Header -->
                    <tr>
                        <td style="background:linear-gradient(135deg,#6366f1 0%%,#8b5cf6 100%%); padding:40px 30px; text-align:center; border-radius:16px 16px 0 0;">
                            <h1 style="color:#ffffff; font-size:24px; font-weight:700; margin:0; letter-spacing:-0.5px;">Sharpnr</h1>
                            <p style="color:rgba(255,255,255,0.9); font-size:15px; margin:10px 0 0; font-weight:500;">Email Verification</p>
                        </td>
                    </tr>
                    
                    <!-- Content -->
                    <tr>
                        <td style="padding:50px 40px;">
                            
                            <div style="text-align:center; margin-bottom:40px;">
                                <h2 style="color:#1f2937; font-size:22px; font-weight:600; margin:0 0 12px; line-height:1.3;">Verify Your Organization</h2>
                                <p style="color:#6b7280; font-size:16px; margin:0; line-height:1.5;">
                                    Enter this verification code to complete your registration:
                                </p>
                            </div>
                            
                            <!-- OTP Code -->
                            <div style="background:#f8fafc; border:2px solid #e2e8f0; border-radius:12px; padding:30px; text-align:center; margin:30px 0;">
                                <p style="color:#64748b; font-size:13px; font-weight:600; margin:0 0 15px; text-transform:uppercase; letter-spacing:1px;">Verification Code</p>
                                <div style="background:#ffffff; border:2px solid #6366f1; border-radius:8px; padding:16px; display:inline-block;">
                                    <span style="font-size:28px; font-weight:700; letter-spacing:6px; color:#6366f1; font-family:'Courier New',monospace;">%s</span>
                                </div>
                                <p style="color:#94a3b8; font-size:12px; margin:15px 0 0; font-style:italic;">Expires in 10 minutes</p>
                            </div>
                            
                            <!-- Security Notice -->
                            <div style="background:#fef3c7; border-left:4px solid #f59e0b; padding:16px 20px; border-radius:0 8px 8px 0; margin:30px 0;">
                                <p style="color:#92400e; font-size:13px; margin:0; line-height:1.5;">
                                    <strong>Security:</strong> If you didn't request this code, please ignore this email. Never share this code with anyone.
                                </p>
                            </div>
                            
                        </td>
                    </tr>
                    
                    <!-- Footer -->
                    <tr>
                        <td style="background:#f9fafb; padding:25px 40px; text-align:center; border-radius:0 0 16px 16px; border-top:1px solid #e5e7eb;">
                            <p style="color:#9ca3af; font-size:12px; margin:0;">
                                © 2024 Sharpnr. All rights reserved.
                            </p>
                            <p style="margin:10px 0 0;">
                                <a href="#" style="color:#6366f1; text-decoration:none; font-size:12px; margin:0 8px;">Support</a>
                                <span style="color:#d1d5db;">•</span>
                                <a href="#" style="color:#6366f1; text-decoration:none; font-size:12px; margin:0 8px;">Privacy</a>
                            </p>
                        </td>
                    </tr>
                    
                </table>
                
            </td>
        </tr>
    </table>
    
</body>
</html>
`, otp)

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
		return err
	}

	// Build and send message
	msg := mail.NewMSG()
	msg.SetFrom("Sharpnr Org Verification <"+server.Username+">").
		AddTo(email_add).
		SetSubject("Organization Registration – Verify Your OTP").
		SetBody(mail.TextHTML, htmlBody)

	if msg.Error != nil {
		return msg.Error
	}

	if err := msg.Send(smtpClient); err != nil {
		return err
	}

	return nil
}
