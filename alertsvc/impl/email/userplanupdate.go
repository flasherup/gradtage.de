package email

const UserPlanUpdateTemplate =
	`From: {{.FromName}} <{{.From}}>
To: {{.To}}
Subject: {{.Subject}}
MIME-version: 1.0
Content-Type: text/html; charset="UTF-8";
<html><body>
<p>&nbsp;</p>
<!-- [if mso]> <style> .templateContainer { border: 0px none #aaaaaa; background-color: #ffffff; border-radius: 0px; } #brandingContainer { background-color: transparent !important; border: 0; } .templateContainerInner { padding: 0px; } </style> <![endif]--><center>
<table id="bodyTable" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: auto; padding: 0; background-color: #f7f7f7; height: 100%; margin: 0; width: 100%;" border="0" width="100%" cellspacing="0" cellpadding="0" align="center" data-upload-file-url="/ajax/email-editor/file/upload" data-upload-files-url="/ajax/email-editor/files/upload">
<tbody>
<tr>
<td id="bodyCell" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: auto; border-top: 0; height: 100%; margin: 0; width: 100%; padding: 50px 20px 20px 20px;" align="center" valign="top"><!-- [if !mso]><!-->
<div class="templateContainer" style="border: 0 none #aaa; background-color: #fff; border-radius: 0; display: table; width: 600px;">
<div class="templateContainerInner" style="padding: 0;"><!--<![endif]--> <!-- [if mso]> <table border="0" cellpadding="0" cellspacing="0" class="templateContainer" width="600" style="border-collapse:collapse;mso-table-lspace:0;mso-table-rspace:0;"> <tbody> <tr> <td class="templateContainerInner" style="border-collapse:collapse;mso-table-lspace:0;mso-table-rspace:0;"> <![endif]-->
<table style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" border="0" width="100%" cellspacing="0" cellpadding="0">
<tbody>
<tr>
<td style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" align="center" valign="top">
<table class="templateRow" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" border="0" width="100%" cellspacing="0" cellpadding="0">
<tbody>
<tr>
<td class="rowContainer kmFloatLeft" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" valign="top">
<table class="kmButtonCollectionBlock" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" border="0" width="100%" cellspacing="0" cellpadding="0">
<tbody class="kmButtonCollectionOuter">
<tr>
<td class="kmButtonCollectionInner" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed; min-width: 60px; padding: 9 18 9 18; background-color: #ffffff;" align="center" valign="top">
<table class="kmButtonCollectionContentContainer" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" border="0" width="100%" cellspacing="0" cellpadding="0">
<tbody>
<tr>
<td style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" align="left">
<table class="kmButtonCollectionContent" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed; font-family: 'Helvetica Neue', Arial;" width="100%" cellspacing="0" cellpadding="0">
<tbody>
<tr>
<td style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" valign="top">
<table class="kmMobileNoAlign kmDesktopAutoWidth kmMobileAutoWidth" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed; width: 100%; float: left;" border="0" cellspacing="0" cellpadding="0" align="center">
<tbody>
<tr class="kmMobileNoAlign" style="font-size: 0;">
<td class="kmDesktopOnly" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed; padding-right: 10px;" valign="middle">
<table style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" border="0" cellspacing="0" cellpadding="0" align="center">
<tbody>
<tr>
<td style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;"><a style="word-wrap: break-word; max-width: 100%; color: #ddd; font-weight: normal; text-decoration: underline;" href="http://www.gradtage.de"> <img style="border: 0; height: auto; line-height: 100%; outline: none; text-decoration: none; max-width: 100%;" src="https://d3k81ch9hvuctc.cloudfront.net/company/MULsK2/images/bdeced58-c872-4f00-be78-067319e9b79a.png" alt="" width="250" /> </a></td>
</tr>
</tbody>
</table>
</td>
<!-- [if !mso]><!-->
<td class="kmMobileHeaderStackDesktopNone" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed; display: none;" valign="middle">
<table style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" border="0" cellspacing="0" cellpadding="0" align="center">
<tbody>
<tr>
<td style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" align="center"><a style="word-wrap: break-word; max-width: 100%; color: #ddd; font-weight: normal; text-decoration: underline;" href="http://www.gradtage.de"> <img style="border: 0; height: auto; line-height: 100%; outline: none; text-decoration: none; max-width: 100%;" src="https://d3k81ch9hvuctc.cloudfront.net/company/MULsK2/images/bdeced58-c872-4f00-be78-067319e9b79a.png" alt="" width="250" /> </a></td>
</tr>
</tbody>
</table>
</td>
<!--<![endif]--></tr>
</tbody>
</table>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
</tbody>
</table>
<table class="kmDividerBlock" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" border="0" width="100%" cellspacing="0" cellpadding="0">
<tbody class="kmDividerBlockOuter">
<tr>
<td class="kmDividerBlockInner" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed; padding: 18px;">
<table class="kmDividerContent" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed; border-top-width: 1px; border-top-style: solid; border-top-color: #ccc;" border="0" width="100%" cellspacing="0" cellpadding="0">
<tbody>
<tr>
<td style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;">&nbsp;</td>
</tr>
</tbody>
</table>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
<tr>
<td style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" align="center" valign="top">
<table class="templateRow" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" border="0" width="100%" cellspacing="0" cellpadding="0">
<tbody>
<tr>
<td class="rowContainer kmFloatLeft" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" valign="top">
<table class="kmTextBlock" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" border="0" width="100%" cellspacing="0" cellpadding="0">
<tbody class="kmTextBlockOuter">
<tr>
<td class="kmTextBlockInner" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" valign="top">
<table class="kmTextContentContainer" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" border="0" width="100%" cellspacing="0" cellpadding="0" align="left">
<tbody>
<tr>
<td class="kmTextContent" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed; color: #3d3d3d; font-family: 'Helvetica Neue', Arial; font-size: 14px; line-height: 1.3; letter-spacing: 0; text-align: left; max-width: 100%; word-wrap: break-word; padding: 9px 18px 9px 18px;" valign="top">
<h3 style="color: #3d3d3d; display: block; font-family: 'Helvetica Neue', Arial; font-size: 24px; font-style: normal; font-weight: bold; line-height: 1.1; letter-spacing: 0; margin: 0; margin-bottom: 12px; text-align: left;">Guten Tag,</h3>
<p style="margin: 0; padding-bottom: 1em;">und vielen Dank f&uuml;r Ihre Bestellung!</p>
<p style="margin: 0; padding-bottom: 1em;">&Uuml;ber unsere API k&ouml;nnen Sie Gradtagzahlen, Heizgradtage und K&uuml;hlgradtage f&uuml;r viele internationale Standorte tag-genau abfragen.</p>
<p style="margin: 0; padding-bottom: 0;">Hier senden wir Ihnen pers&ouml;nlichen API-Key:</p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
</tbody>
</table>
<table class="kmTextBlock" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" border="0" width="100%" cellspacing="0" cellpadding="0">
<tbody class="kmTextBlockOuter">
<tr>
<td class="kmTextBlockInner" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" valign="top">
<table class="kmTextContentContainer" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" border="0" width="100%" cellspacing="0" cellpadding="0" align="left">
<tbody>
<tr>
<td class="kmTextContent" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed; color: #3d3d3d; font-family: 'Helvetica Neue', Arial; font-size: 14px; line-height: 1.3; letter-spacing: 0; text-align: left; max-width: 100%; word-wrap: break-word; padding: 9px 18px 9px 18px;" valign="top">
<p style="margin: 0; padding-bottom: 0; text-align: center;"><em><strong>{{.Key}}</strong></em></p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
</tbody>
</table>
<table class="kmTextBlock" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" border="0" width="100%" cellspacing="0" cellpadding="0">
<tbody class="kmTextBlockOuter">
<tr>
<td class="kmTextBlockInner" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" valign="top">
<table class="kmTextContentContainer" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" border="0" width="100%" cellspacing="0" cellpadding="0" align="left">
<tbody>
<tr>
<td class="kmTextContent" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed; color: #3d3d3d; font-family: 'Helvetica Neue', Arial; font-size: 14px; line-height: 1.3; letter-spacing: 0; text-align: left; max-width: 100%; word-wrap: break-word; padding: 9px 18px 9px 18px;" valign="top">
<p style="margin: 0; padding-bottom: 1em;">Die Rechnung f&uuml;r Ihre Bestellung und eine Anleitung zur Nutzung des API-Services erhalten Sie in einer separaten Email. Hier finden Sie auch Details zur Ihrem Abonnement wie z.B. gebuchte Leistungen oder Laufzeit.</p>
<p style="margin: 0; padding-bottom: 1em;">Sollten Sie dennoch offene Fragen und technische Schwierigkeiten bei der Einrichtung haben, kontaktieren Sie uns gerne.</p>
<p style="margin: 0; padding-bottom: 1em;">Freundliche Gr&uuml;&szlig;e</p>
<p style="margin: 0; padding-bottom: 1em;"><em>Team GRADTAGE</em></p>
<p style="margin: 0; padding-bottom: 0;">&nbsp;</p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
</tbody>
</table>
<table class="kmImageBlock" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" border="0" width="100%" cellspacing="0" cellpadding="0">
<tbody class="kmImageBlockOuter">
<tr>
<td class="kmImageBlockInner" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed; padding: 0px;" valign="top">
<table class="kmImageContentContainer" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" border="0" width="100%" cellspacing="0" cellpadding="0" align="left">
<tbody>
<tr>
<td class="kmImageContent" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed; padding: 0; font-size: 0;" valign="top"><img class="kmImage" style="border: 0; height: auto; line-height: 100%; outline: none; text-decoration: none; max-width: 100%; padding-bottom: 0; display: inline; vertical-align: top; font-size: 12px; width: 100%;" src="https://d3k81ch9hvuctc.cloudfront.net/assets/email/bottom_shadow_444.png" alt="Shadow" width="600" align="center" /></td>
</tr>
</tbody>
</table>
</td>
</tr>
</tbody>
</table>
<table class="kmImageBlock" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed; min-width: 100%;" border="0" width="100%" cellspacing="0" cellpadding="0">
<tbody class="kmImageBlockOuter">
<tr>
<td class="kmImageBlockInner" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed; padding: 9px; padding-right: 9; padding-left: 9;" valign="top">
<table class="kmImageContentContainer" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed; min-width: 100%;" border="0" width="100%" cellspacing="0" cellpadding="0" align="left">
<tbody>
<tr>
<td class="kmImageContent" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed; padding: 0px 9px 0 9px; font-size: 0; text-align: center;" valign="top"><a style="word-wrap: break-word; max-width: 100%; color: #ddd; font-weight: normal; text-decoration: underline;" href="http://www.gradtage.de" target="_self"> <img class="kmImage" style="border: 0; height: auto; line-height: 100%; outline: none; text-decoration: none; max-width: 75px; padding-bottom: 0; display: inline; vertical-align: top; font-size: 12px; width: 100%; padding: 0; border-width: 50px;" src="https://d3k81ch9hvuctc.cloudfront.net/company/MULsK2/images/7dfa11c7-b3e7-45fd-bd66-f8a9f90f6ca5.png" alt="GRADTAGE Logo" width="75" align="center" /> </a></td>
</tr>
</tbody>
</table>
</td>
</tr>
</tbody>
</table>
<table class="kmTextBlock" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" border="0" width="100%" cellspacing="0" cellpadding="0">
<tbody class="kmTextBlockOuter">
<tr>
<td class="kmTextBlockInner" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" valign="top">
<table class="kmTextContentContainer" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" border="0" width="100%" cellspacing="0" cellpadding="0" align="left">
<tbody>
<tr>
<td class="kmTextContent" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed; color: #3d3d3d; font-family: 'Helvetica Neue', Arial; font-size: 14px; line-height: 1.3; letter-spacing: 0; text-align: left; max-width: 100%; word-wrap: break-word; padding: 9px 18px 9px 18px;" valign="top">
<p style="margin: 0; padding-bottom: 1em; text-align: center;"><a style="word-wrap: break-word; max-width: 100%; color: #ddd; font-weight: normal; text-decoration: underline;" href="http://gradtage.de"><span style="color: #666666;"><strong>GRADTAGE</strong></span></a></p>
<p style="margin: 0; padding-bottom: 0; text-align: center;"><span style="color: #666666;">Regina Nagel<br /> Max-Brauer-Allee 104 | 22765 Hamburg<br /> Telefon: 040 60779041<br /> E-Mail: info@gradtage.de</span></p>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
</tbody>
</table>
</td>
</tr>
</tbody>
</table>
<!-- [if !mso]><!--></div>
</div>
<!--<![endif]--> <!-- [if mso]> </td> </tr> </tbody> </table> <![endif]--> <!-- [if !mso]><!-->
<div class="templateContainer brandingContainer" style="border: 0; background-color: transparent; border-radius: 0; display: table; width: 600px;">
<div class="templateContainerInner" style="padding: 0;"><!--<![endif]--> <!-- [if mso]> <table border="0" cellpadding="0" cellspacing="0" class="templateContainer" id="brandingContainer" width="600" style="border-collapse:collapse;mso-table-lspace:0;mso-table-rspace:0;"> <tbody> <tr> <td class="templateContainerInner" style="border-collapse:collapse;mso-table-lspace:0;mso-table-rspace:0;"> <![endif]-->
<table style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" border="0" width="100%" cellspacing="0" cellpadding="0">
<tbody>
<tr>
<td style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" align="center" valign="top">
<table class="templateRow" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" border="0" width="100%" cellspacing="0" cellpadding="0">
<tbody>
<tr>
<td class="rowContainer kmFloatLeft" style="border-collapse: collapse; mso-table-lspace: 0; mso-table-rspace: 0; table-layout: fixed;" valign="top">&nbsp;</td>
</tr>
</tbody>
</table>
</td>
</tr>
</tbody>
</table>
<!-- [if !mso]><!--></div>
</div>
<!--<![endif]--> <!-- [if mso]> </td> </tr> </tbody> </table> <![endif]--></td>
</tr>
</tbody>
</table>
</center>
</body></html>`
