/*
 * TgMusicBot - Telegram Music Bot
 *  Copyright (c) 2025-2026 Ashok Shau
 *
 *  Licensed under GNU GPL v3
 *  See https://github.com/AshokShau/TgMusicBot
 */

package bot

import (
	"ashokshau/tgmusic/internal/config"
	"ashokshau/tgmusic/internal/utils"
	"fmt"
	"runtime"
	"time"

	td "github.com/AshokShau/gotdbot"
)

func pingHandler(c *td.Client, m *td.Message) error {
	deleteCmd(c, m)

	start := time.Now()

	msg, err := m.ReplyText(c, "Pinging… please wait…", nil)
	if err != nil {
		return err
	}

	latency := time.Since(start).Milliseconds()
	uptime := getFormattedDuration(time.Since(startTime))

	response := fmt.Sprintf(
		"<b>📊 System Performance Metrics</b>\n\n"+
			"<b>Bot Latency:</b> <code>%d ms</code>\n"+
			"<b>Uptime:</b> <code>%s</code>\n"+
			"<b>Go Routines:</b> <code>%d</code>\n",
		latency, uptime, runtime.NumGoroutine(),
	)

	_, err = msg.EditText(c, response, &td.EditTextMessageOpts{ParseMode: "HTML"})
	return err
}

func startHandler(c *td.Client, m *td.Message) error {
	chatID := m.ChatId
	go storeChatToDB(chatID)

	deleteCmd(c, m)

	if m.IsPrivate() {
		response := fmt.Sprintf(
			"<img src=\"%s\"/>\n"+
				"<h3>Wᴇʟᴄᴏᴍᴇ, %s!</h3>\n"+
				"<p><b>%s</b> ʟᴇᴛs ʏᴏᴜ sᴛʀᴇᴀᴍ ʜɪɢʜ-ǫᴜᴀʟɪᴛʏ ᴍᴜsɪᴄ ᴀɴᴅ ᴠɪᴅᴇᴏ ᴅɪʀᴇᴄᴛʟʏ ɪɴ Tᴇʟᴇɢʀᴀᴍ ᴠᴏɪᴄᴇ ᴀɴᴅ ᴠɪᴅᴇᴏ ᴄʜᴀᴛs.</p>\n\n"+
				"<p><b>Sᴜᴘᴘᴏʀᴛᴇᴅ ᴘʟᴀᴛғᴏʀᴍs:</b> YᴏᴜTᴜʙᴇ, Sᴘᴏᴛɪғʏ, Aᴘᴘʟᴇ Mᴜsɪᴄ, SᴏᴜɴᴅCʟᴏᴜᴅ, Dᴇᴇᴢᴇʀ, Tᴡɪᴛᴄʜ, ᴀɴᴅ ᴍᴀɴʏ ᴍᴏʀᴇ.</p>\n\n"+
				"<p>Usᴇ ᴛʜᴇ ʙᴜᴛᴛᴏɴs ʙᴇʟᴏᴡ ᴛᴏ ᴀᴅᴅ ᴛʜᴇ ʙᴏᴛ ᴛᴏ ʏᴏᴜʀ ɢʀᴏᴜᴘ ᴏʀ ᴇxᴘʟᴏʀᴇ ᴛʜᴇ ᴀᴠᴀɪʟᴀʙʟᴇ ᴄᴏᴍᴍᴀɴᴅs.</p>",
			config.StartImg,
			firstName(c, m),
			c.Me.FirstName,
		)

		richMessage := &td.InputRichMessage{
			Source: &td.RichMessageSourceHtml{
				Text: response,
			},
		}

		_, err := m.ReplyRichMessage(c, richMessage, &td.SendTextMessageOpts{
			ReplyMarkup: utils.AddMeMarkup(c.Me.Usernames.EditableUsername),
		})

		return err
	}

	uptime := getFormattedDuration(time.Since(startTime))
	htmlText := fmt.Sprintf(
		"<h3>%s ɪs ʀᴇᴀᴅʏ!</h3>\n"+
			"<p><b>Uᴘᴛɪᴍᴇ:</b> <code>%s</code></p>\n"+
			"<p><i>A ғᴇᴀᴛᴜʀᴇ-ʀɪᴄʜ ᴍᴜsɪᴄ ʙᴏᴛ ғᴏʀ ʏᴏᴜʀ ɢʀᴏᴜᴘ ᴠɪᴅᴇᴏ ᴄʜᴀᴛs. Pʟᴀʏ ʏᴏᴜʀ ғᴀᴠᴏʀɪᴛᴇ ᴛʀᴀᴄᴋs sᴇᴀᴍʟᴇssʟʏ..</i></p>\n\n"+
			"<p><tg-button type=\"url\" url=\"%s\">υᴘᴅᴧᴛєs</tg-button> <tg-button type=\"url\" url=\"%s\">sυᴘᴘσʀᴛ</tg-button></p>",
		c.Me.FirstName,
		uptime,
		config.SupportChannel,
		config.SupportGroup,
	)

	richMessage := &td.InputRichMessage{
		Source: &td.RichMessageSourceHtml{
			Text: htmlText,
		},
	}

	_, err := m.ReplyRichMessage(c, richMessage, nil)
	return err
}
