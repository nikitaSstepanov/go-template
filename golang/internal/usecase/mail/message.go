package mail

import "fmt"

func activationMessage(code string) string {
	return fmt.Sprintf(
		`
        <html>
            <body>
                <div style="background-color: rgb(40, 39, 39);
                            border-radius: 15px; 
                            margin-left: 0.75%%;
                            padding-bottom: 2%%;">
                    <h1 style="color: whitesmoke;
                               font-family: MailSans, Helvetica, Arial, sans-serif;
                               margin-left: 2%%;margin-bottom: -1.5%%;
                               padding-top: 1%%;
                               text-decoration: none;">
                        Hello!
                    </h1>
                    <p style="color: whitesmoke;
                              font-family: MailSans, Helvetica, Arial, sans-serif;
                              margin-left: 2%%;
                              margin-right: 2%%;
                              font-size: 90%%;">
                        &nbsp;&nbsp;&nbsp;Your email address was used when registering. Your verification code:<br>
                    </p>
                    <p align="center"
                       style="color: whitesmoke;
                              font-family: MailSans, Helvetica, Arial, sans-serif;">
                        %s
                    </p>
                </div>
            </body>
        </html>
        `, code,
	)
}
