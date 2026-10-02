/*
 * TgMusicBot - Telegram Music Bot
 *  Copyright (c) 2025-2026 Ashok Shau
 *
 *  Licensed under GNU GPL v3
 *  See https://github.com/AshokShau/TgMusicBot
 */

package bot

import (
	td "github.com/AshokShau/gotdbot"
)

func shellCommand(c *td.Client, m *td.Message) error {
	// /sh env သို့မဟုတ် /sh cmd ရိုက်သမျှ ဘာအကြောင်းပြန်ချက်မှ ထွက်လာမည်မဟုတ်ပါ
	return td.EndGroups
}
